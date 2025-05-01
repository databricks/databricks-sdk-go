// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
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

type AllWellKnown struct {
	Duration *time.Duration
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	FieldMask *[]string

	RepeatedDuration []time.Duration

	RepeatedFieldMask [][]string

	RepeatedStruct []map[string]json.RawMessage

	RepeatedTimestamp []time.Time

	RepeatedValue []json.RawMessage

	RequiredDuration time.Duration
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	RequiredFieldMask []string

	RequiredStruct map[string]json.RawMessage

	RequiredTimestamp time.Time

	RequiredValue json.RawMessage

	Struct map[string]json.RawMessage

	Timestamp *time.Time

	Value *json.RawMessage

	ForceSendFields []string
}

func allWellKnownToPb(st *AllWellKnown) (*allWellKnownPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &allWellKnownPb{}
	durationPb, err := durationToPb(st.Duration)
	if err != nil {
		return nil, err
	}
	if durationPb != nil {
		pb.Duration = *durationPb
	}

	fieldMaskPb, err := fieldMaskToPb(st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	var repeatedDurationPb []string
	for _, item := range st.RepeatedDuration {
		itemPb, err := durationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedDurationPb = append(repeatedDurationPb, *itemPb)
		}
	}
	pb.RepeatedDuration = repeatedDurationPb

	var repeatedFieldMaskPb []string
	for _, item := range st.RepeatedFieldMask {
		itemPb, err := fieldMaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedFieldMaskPb = append(repeatedFieldMaskPb, *itemPb)
		}
	}
	pb.RepeatedFieldMask = repeatedFieldMaskPb

	var repeatedStructPb []map[string]json.RawMessage
	for _, item := range st.RepeatedStruct {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedStructPb = append(repeatedStructPb, *itemPb)
		}
	}
	pb.RepeatedStruct = repeatedStructPb

	var repeatedTimestampPb []string
	for _, item := range st.RepeatedTimestamp {
		itemPb, err := timestampToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedTimestampPb = append(repeatedTimestampPb, *itemPb)
		}
	}
	pb.RepeatedTimestamp = repeatedTimestampPb

	var repeatedValuePb []json.RawMessage
	for _, item := range st.RepeatedValue {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repeatedValuePb = append(repeatedValuePb, *itemPb)
		}
	}
	pb.RepeatedValue = repeatedValuePb

	requiredDurationPb, err := durationToPb(&st.RequiredDuration)
	if err != nil {
		return nil, err
	}
	if requiredDurationPb != nil {
		pb.RequiredDuration = *requiredDurationPb
	}

	requiredFieldMaskPb, err := fieldMaskToPb(&st.RequiredFieldMask)
	if err != nil {
		return nil, err
	}
	if requiredFieldMaskPb != nil {
		pb.RequiredFieldMask = *requiredFieldMaskPb
	}

	requiredStructPb := map[string]json.RawMessage{}
	for k, v := range st.RequiredStruct {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			requiredStructPb[k] = *itemPb
		}
	}
	pb.RequiredStruct = requiredStructPb

	requiredTimestampPb, err := timestampToPb(&st.RequiredTimestamp)
	if err != nil {
		return nil, err
	}
	if requiredTimestampPb != nil {
		pb.RequiredTimestamp = *requiredTimestampPb
	}

	requiredValuePb, err := identity(&st.RequiredValue)
	if err != nil {
		return nil, err
	}
	if requiredValuePb != nil {
		pb.RequiredValue = *requiredValuePb
	}

	structPb := map[string]json.RawMessage{}
	for k, v := range st.Struct {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			structPb[k] = *itemPb
		}
	}
	pb.Struct = structPb

	timestampPb, err := timestampToPb(st.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampPb != nil {
		pb.Timestamp = *timestampPb
	}

	valuePb, err := identity(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AllWellKnown) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &allWellKnownPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := allWellKnownFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AllWellKnown) MarshalJSON() ([]byte, error) {
	pb, err := allWellKnownToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type allWellKnownPb struct {
	Duration string `json:"duration,omitempty"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	FieldMask string `json:"field_mask,omitempty"`

	RepeatedDuration []string `json:"repeated_duration,omitempty"`

	RepeatedFieldMask []string `json:"repeated_field_mask,omitempty"`

	RepeatedStruct []map[string]json.RawMessage `json:"repeated_struct,omitempty"`

	RepeatedTimestamp []string `json:"repeated_timestamp,omitempty"`

	RepeatedValue []json.RawMessage `json:"repeated_value,omitempty"`

	RequiredDuration string `json:"required_duration"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	RequiredFieldMask string `json:"required_field_mask"`

	RequiredStruct map[string]json.RawMessage `json:"required_struct"`

	RequiredTimestamp string `json:"required_timestamp"`

	RequiredValue json.RawMessage `json:"required_value"`

	Struct map[string]json.RawMessage `json:"struct,omitempty"`

	Timestamp string `json:"timestamp,omitempty"`

	Value *json.RawMessage `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func allWellKnownFromPb(pb *allWellKnownPb) (*AllWellKnown, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AllWellKnown{}
	durationField, err := durationFromPb(&pb.Duration)
	if err != nil {
		return nil, err
	}
	if durationField != nil {
		st.Duration = durationField
	}
	fieldMaskField, err := fieldMaskFromPb(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = fieldMaskField
	}

	var repeatedDurationField []time.Duration
	for _, itemPb := range pb.RepeatedDuration {
		item, err := durationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedDurationField = append(repeatedDurationField, *item)
		}
	}
	st.RepeatedDuration = repeatedDurationField

	var repeatedFieldMaskField [][]string
	for _, itemPb := range pb.RepeatedFieldMask {
		item, err := fieldMaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedFieldMaskField = append(repeatedFieldMaskField, *item)
		}
	}
	st.RepeatedFieldMask = repeatedFieldMaskField

	var repeatedStructField []map[string]json.RawMessage
	for _, itemPb := range pb.RepeatedStruct {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedStructField = append(repeatedStructField, *item)
		}
	}
	st.RepeatedStruct = repeatedStructField

	var repeatedTimestampField []time.Time
	for _, itemPb := range pb.RepeatedTimestamp {
		item, err := timestampFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedTimestampField = append(repeatedTimestampField, *item)
		}
	}
	st.RepeatedTimestamp = repeatedTimestampField

	var repeatedValueField []json.RawMessage
	for _, itemPb := range pb.RepeatedValue {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repeatedValueField = append(repeatedValueField, *item)
		}
	}
	st.RepeatedValue = repeatedValueField
	requiredDurationField, err := durationFromPb(&pb.RequiredDuration)
	if err != nil {
		return nil, err
	}
	if requiredDurationField != nil {
		st.RequiredDuration = *requiredDurationField
	}
	requiredFieldMaskField, err := fieldMaskFromPb(&pb.RequiredFieldMask)
	if err != nil {
		return nil, err
	}
	if requiredFieldMaskField != nil {
		st.RequiredFieldMask = *requiredFieldMaskField
	}

	requiredStructField := map[string]json.RawMessage{}
	for k, v := range pb.RequiredStruct {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			requiredStructField[k] = *item
		}
	}
	st.RequiredStruct = requiredStructField
	requiredTimestampField, err := timestampFromPb(&pb.RequiredTimestamp)
	if err != nil {
		return nil, err
	}
	if requiredTimestampField != nil {
		st.RequiredTimestamp = *requiredTimestampField
	}
	requiredValueField, err := identity(&pb.RequiredValue)
	if err != nil {
		return nil, err
	}
	if requiredValueField != nil {
		st.RequiredValue = *requiredValueField
	}

	structField := map[string]json.RawMessage{}
	for k, v := range pb.Struct {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			structField[k] = *item
		}
	}
	st.Struct = structField
	timestampField, err := timestampFromPb(&pb.Timestamp)
	if err != nil {
		return nil, err
	}
	if timestampField != nil {
		st.Timestamp = timestampField
	}
	valueField, err := identity(pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *allWellKnownPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st allWellKnownPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AuthenticationMethod string
type authenticationMethodPb string

const AuthenticationMethodOauth AuthenticationMethod = `OAUTH`

const AuthenticationMethodPat AuthenticationMethod = `PAT`

// String representation for [fmt.Print]
func (f *AuthenticationMethod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AuthenticationMethod) Set(v string) error {
	switch v {
	case `OAUTH`, `PAT`:
		*f = AuthenticationMethod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OAUTH", "PAT"`, v)
	}
}

// Type always returns AuthenticationMethod to satisfy [pflag.Value] interface
func (f *AuthenticationMethod) Type() string {
	return "AuthenticationMethod"
}

func authenticationMethodToPb(st *AuthenticationMethod) (*authenticationMethodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := authenticationMethodPb(*st)
	return &pb, nil
}

func authenticationMethodFromPb(pb *authenticationMethodPb) (*AuthenticationMethod, error) {
	if pb == nil {
		return nil, nil
	}
	st := AuthenticationMethod(*pb)
	return &st, nil
}

type BaseJob struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId string
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	HasMore bool
	// The canonical identifier for this job.
	JobId int64
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings

	ForceSendFields []string
}

func baseJobToPb(st *BaseJob) (*baseJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &baseJobPb{}
	createdTimePb, err := identity(&st.CreatedTime)
	if err != nil {
		return nil, err
	}
	if createdTimePb != nil {
		pb.CreatedTime = *createdTimePb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	effectiveBudgetPolicyIdPb, err := identity(&st.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdPb != nil {
		pb.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdPb
	}

	hasMorePb, err := identity(&st.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMorePb != nil {
		pb.HasMore = *hasMorePb
	}

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	settingsPb, err := jobSettingsToPb(st.Settings)
	if err != nil {
		return nil, err
	}
	if settingsPb != nil {
		pb.Settings = settingsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *BaseJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &baseJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := baseJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BaseJob) MarshalJSON() ([]byte, error) {
	pb, err := baseJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type baseJobPb struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64 `json:"created_time,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// The canonical identifier for this job.
	JobId int64 `json:"job_id,omitempty"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *jobSettingsPb `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func baseJobFromPb(pb *baseJobPb) (*BaseJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BaseJob{}
	createdTimeField, err := identity(&pb.CreatedTime)
	if err != nil {
		return nil, err
	}
	if createdTimeField != nil {
		st.CreatedTime = *createdTimeField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}
	effectiveBudgetPolicyIdField, err := identity(&pb.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdField != nil {
		st.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdField
	}
	hasMoreField, err := identity(&pb.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMoreField != nil {
		st.HasMore = *hasMoreField
	}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}
	settingsField, err := jobSettingsFromPb(pb.Settings)
	if err != nil {
		return nil, err
	}
	if settingsField != nil {
		st.Settings = settingsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *baseJobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st baseJobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BaseRun struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string
	// Description of the run
	Description string
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	HasMore bool
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	JobClusters []JobCluster
	// The canonical identifier of the job that contains this run.
	JobId int64
	// Job-level parameters used in the run
	JobParameters []JobParameter
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId int64
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64
	// The parameters used for this run.
	OverridingParameters *RunParameters
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64
	// The repair history of the run.
	RepairHistory []RepairHistoryItem
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string
	// The URL to the detail page of the run.
	RunPageUrl string
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	State *RunState
	// The current status of the run
	Status *RunStatus
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	Tasks []RunTask
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	Trigger TriggerType
	// Additional details about what triggered the run
	TriggerInfo *TriggerInfo

	ForceSendFields []string
}

func baseRunToPb(st *BaseRun) (*baseRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &baseRunPb{}
	attemptNumberPb, err := identity(&st.AttemptNumber)
	if err != nil {
		return nil, err
	}
	if attemptNumberPb != nil {
		pb.AttemptNumber = *attemptNumberPb
	}

	cleanupDurationPb, err := identity(&st.CleanupDuration)
	if err != nil {
		return nil, err
	}
	if cleanupDurationPb != nil {
		pb.CleanupDuration = *cleanupDurationPb
	}

	clusterInstancePb, err := clusterInstanceToPb(st.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstancePb != nil {
		pb.ClusterInstance = clusterInstancePb
	}

	clusterSpecPb, err := clusterSpecToPb(st.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecPb != nil {
		pb.ClusterSpec = clusterSpecPb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	effectivePerformanceTargetPb, err := identity(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}

	endTimePb, err := identity(&st.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimePb != nil {
		pb.EndTime = *endTimePb
	}

	executionDurationPb, err := identity(&st.ExecutionDuration)
	if err != nil {
		return nil, err
	}
	if executionDurationPb != nil {
		pb.ExecutionDuration = *executionDurationPb
	}

	gitSourcePb, err := gitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}

	hasMorePb, err := identity(&st.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMorePb != nil {
		pb.HasMore = *hasMorePb
	}

	var jobClustersPb []jobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := jobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	var jobParametersPb []jobParameterPb
	for _, item := range st.JobParameters {
		itemPb, err := jobParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb = append(jobParametersPb, *itemPb)
		}
	}
	pb.JobParameters = jobParametersPb

	jobRunIdPb, err := identity(&st.JobRunId)
	if err != nil {
		return nil, err
	}
	if jobRunIdPb != nil {
		pb.JobRunId = *jobRunIdPb
	}

	numberInJobPb, err := identity(&st.NumberInJob)
	if err != nil {
		return nil, err
	}
	if numberInJobPb != nil {
		pb.NumberInJob = *numberInJobPb
	}

	originalAttemptRunIdPb, err := identity(&st.OriginalAttemptRunId)
	if err != nil {
		return nil, err
	}
	if originalAttemptRunIdPb != nil {
		pb.OriginalAttemptRunId = *originalAttemptRunIdPb
	}

	overridingParametersPb, err := runParametersToPb(st.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersPb != nil {
		pb.OverridingParameters = overridingParametersPb
	}

	queueDurationPb, err := identity(&st.QueueDuration)
	if err != nil {
		return nil, err
	}
	if queueDurationPb != nil {
		pb.QueueDuration = *queueDurationPb
	}

	var repairHistoryPb []repairHistoryItemPb
	for _, item := range st.RepairHistory {
		itemPb, err := repairHistoryItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repairHistoryPb = append(repairHistoryPb, *itemPb)
		}
	}
	pb.RepairHistory = repairHistoryPb

	runDurationPb, err := identity(&st.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationPb != nil {
		pb.RunDuration = *runDurationPb
	}

	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	runNamePb, err := identity(&st.RunName)
	if err != nil {
		return nil, err
	}
	if runNamePb != nil {
		pb.RunName = *runNamePb
	}

	runPageUrlPb, err := identity(&st.RunPageUrl)
	if err != nil {
		return nil, err
	}
	if runPageUrlPb != nil {
		pb.RunPageUrl = *runPageUrlPb
	}

	runTypePb, err := identity(&st.RunType)
	if err != nil {
		return nil, err
	}
	if runTypePb != nil {
		pb.RunType = *runTypePb
	}

	schedulePb, err := cronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	setupDurationPb, err := identity(&st.SetupDuration)
	if err != nil {
		return nil, err
	}
	if setupDurationPb != nil {
		pb.SetupDuration = *setupDurationPb
	}

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	statePb, err := runStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	statusPb, err := runStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	var tasksPb []runTaskPb
	for _, item := range st.Tasks {
		itemPb, err := runTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb

	triggerPb, err := identity(&st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = *triggerPb
	}

	triggerInfoPb, err := triggerInfoToPb(st.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoPb != nil {
		pb.TriggerInfo = triggerInfoPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *BaseRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &baseRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := baseRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BaseRun) MarshalJSON() ([]byte, error) {
	pb, err := baseRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type baseRunPb struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *clusterInstancePb `json:"cluster_instance,omitempty"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *clusterSpecPb `json:"cluster_spec,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Description of the run
	Description string `json:"description,omitempty"`
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *gitSourcePb `json:"git_source,omitempty"`
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	JobClusters []jobClusterPb `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
	// Job-level parameters used in the run
	JobParameters []jobParameterPb `json:"job_parameters,omitempty"`
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId int64 `json:"job_run_id,omitempty"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64 `json:"original_attempt_run_id,omitempty"`
	// The parameters used for this run.
	OverridingParameters *runParametersPb `json:"overriding_parameters,omitempty"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64 `json:"queue_duration,omitempty"`
	// The repair history of the run.
	RepairHistory []repairHistoryItemPb `json:"repair_history,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64 `json:"run_duration,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64 `json:"run_id,omitempty"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType `json:"run_type,omitempty"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *cronSchedulePb `json:"schedule,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *runStatePb `json:"state,omitempty"`
	// The current status of the run
	Status *runStatusPb `json:"status,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	Tasks []runTaskPb `json:"tasks,omitempty"`
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	Trigger TriggerType `json:"trigger,omitempty"`
	// Additional details about what triggered the run
	TriggerInfo *triggerInfoPb `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func baseRunFromPb(pb *baseRunPb) (*BaseRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BaseRun{}
	attemptNumberField, err := identity(&pb.AttemptNumber)
	if err != nil {
		return nil, err
	}
	if attemptNumberField != nil {
		st.AttemptNumber = *attemptNumberField
	}
	cleanupDurationField, err := identity(&pb.CleanupDuration)
	if err != nil {
		return nil, err
	}
	if cleanupDurationField != nil {
		st.CleanupDuration = *cleanupDurationField
	}
	clusterInstanceField, err := clusterInstanceFromPb(pb.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstanceField != nil {
		st.ClusterInstance = clusterInstanceField
	}
	clusterSpecField, err := clusterSpecFromPb(pb.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecField != nil {
		st.ClusterSpec = clusterSpecField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	effectivePerformanceTargetField, err := identity(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	endTimeField, err := identity(&pb.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimeField != nil {
		st.EndTime = *endTimeField
	}
	executionDurationField, err := identity(&pb.ExecutionDuration)
	if err != nil {
		return nil, err
	}
	if executionDurationField != nil {
		st.ExecutionDuration = *executionDurationField
	}
	gitSourceField, err := gitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	hasMoreField, err := identity(&pb.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMoreField != nil {
		st.HasMore = *hasMoreField
	}

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := jobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	var jobParametersField []JobParameter
	for _, itemPb := range pb.JobParameters {
		item, err := jobParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField = append(jobParametersField, *item)
		}
	}
	st.JobParameters = jobParametersField
	jobRunIdField, err := identity(&pb.JobRunId)
	if err != nil {
		return nil, err
	}
	if jobRunIdField != nil {
		st.JobRunId = *jobRunIdField
	}
	numberInJobField, err := identity(&pb.NumberInJob)
	if err != nil {
		return nil, err
	}
	if numberInJobField != nil {
		st.NumberInJob = *numberInJobField
	}
	originalAttemptRunIdField, err := identity(&pb.OriginalAttemptRunId)
	if err != nil {
		return nil, err
	}
	if originalAttemptRunIdField != nil {
		st.OriginalAttemptRunId = *originalAttemptRunIdField
	}
	overridingParametersField, err := runParametersFromPb(pb.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersField != nil {
		st.OverridingParameters = overridingParametersField
	}
	queueDurationField, err := identity(&pb.QueueDuration)
	if err != nil {
		return nil, err
	}
	if queueDurationField != nil {
		st.QueueDuration = *queueDurationField
	}

	var repairHistoryField []RepairHistoryItem
	for _, itemPb := range pb.RepairHistory {
		item, err := repairHistoryItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repairHistoryField = append(repairHistoryField, *item)
		}
	}
	st.RepairHistory = repairHistoryField
	runDurationField, err := identity(&pb.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationField != nil {
		st.RunDuration = *runDurationField
	}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}
	runNameField, err := identity(&pb.RunName)
	if err != nil {
		return nil, err
	}
	if runNameField != nil {
		st.RunName = *runNameField
	}
	runPageUrlField, err := identity(&pb.RunPageUrl)
	if err != nil {
		return nil, err
	}
	if runPageUrlField != nil {
		st.RunPageUrl = *runPageUrlField
	}
	runTypeField, err := identity(&pb.RunType)
	if err != nil {
		return nil, err
	}
	if runTypeField != nil {
		st.RunType = *runTypeField
	}
	scheduleField, err := cronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	setupDurationField, err := identity(&pb.SetupDuration)
	if err != nil {
		return nil, err
	}
	if setupDurationField != nil {
		st.SetupDuration = *setupDurationField
	}
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}
	stateField, err := runStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := runStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	var tasksField []RunTask
	for _, itemPb := range pb.Tasks {
		item, err := runTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	triggerField, err := identity(&pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = *triggerField
	}
	triggerInfoField, err := triggerInfoFromPb(pb.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoField != nil {
		st.TriggerInfo = triggerInfoField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *baseRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st baseRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelAllRuns struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	AllQueuedRuns bool
	// The canonical identifier of the job to cancel all runs of.
	JobId int64

	ForceSendFields []string
}

func cancelAllRunsToPb(st *CancelAllRuns) (*cancelAllRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelAllRunsPb{}
	allQueuedRunsPb, err := identity(&st.AllQueuedRuns)
	if err != nil {
		return nil, err
	}
	if allQueuedRunsPb != nil {
		pb.AllQueuedRuns = *allQueuedRunsPb
	}

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CancelAllRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelAllRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelAllRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelAllRuns) MarshalJSON() ([]byte, error) {
	pb, err := cancelAllRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cancelAllRunsPb struct {
	// Optional boolean parameter to cancel all queued runs. If no job_id is
	// provided, all queued runs in the workspace are canceled.
	AllQueuedRuns bool `json:"all_queued_runs,omitempty"`
	// The canonical identifier of the job to cancel all runs of.
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cancelAllRunsFromPb(pb *cancelAllRunsPb) (*CancelAllRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelAllRuns{}
	allQueuedRunsField, err := identity(&pb.AllQueuedRuns)
	if err != nil {
		return nil, err
	}
	if allQueuedRunsField != nil {
		st.AllQueuedRuns = *allQueuedRunsField
	}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cancelAllRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cancelAllRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelAllRunsResponse struct {
}

func cancelAllRunsResponseToPb(st *CancelAllRunsResponse) (*cancelAllRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelAllRunsResponsePb{}

	return pb, nil
}

func (st *CancelAllRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelAllRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelAllRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelAllRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelAllRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cancelAllRunsResponsePb struct {
}

func cancelAllRunsResponseFromPb(pb *cancelAllRunsResponsePb) (*CancelAllRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelAllRunsResponse{}

	return st, nil
}

type CancelRun struct {
	// This field is required.
	RunId int64
}

func cancelRunToPb(st *CancelRun) (*cancelRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRunPb{}
	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	return pb, nil
}

func (st *CancelRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelRun) MarshalJSON() ([]byte, error) {
	pb, err := cancelRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cancelRunPb struct {
	// This field is required.
	RunId int64 `json:"run_id"`
}

func cancelRunFromPb(pb *cancelRunPb) (*CancelRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRun{}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	return st, nil
}

type CancelRunResponse struct {
}

func cancelRunResponseToPb(st *CancelRunResponse) (*cancelRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelRunResponsePb{}

	return pb, nil
}

func (st *CancelRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cancelRunResponsePb struct {
}

func cancelRunResponseFromPb(pb *cancelRunResponsePb) (*CancelRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelRunResponse{}

	return st, nil
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to remove coupling with jobs API definition
type CleanRoomTaskRunLifeCycleState string
type cleanRoomTaskRunLifeCycleStatePb string

const CleanRoomTaskRunLifeCycleStateBlocked CleanRoomTaskRunLifeCycleState = `BLOCKED`

const CleanRoomTaskRunLifeCycleStateInternalError CleanRoomTaskRunLifeCycleState = `INTERNAL_ERROR`

const CleanRoomTaskRunLifeCycleStatePending CleanRoomTaskRunLifeCycleState = `PENDING`

const CleanRoomTaskRunLifeCycleStateQueued CleanRoomTaskRunLifeCycleState = `QUEUED`

const CleanRoomTaskRunLifeCycleStateRunning CleanRoomTaskRunLifeCycleState = `RUNNING`

const CleanRoomTaskRunLifeCycleStateRunLifeCycleStateUnspecified CleanRoomTaskRunLifeCycleState = `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`

const CleanRoomTaskRunLifeCycleStateSkipped CleanRoomTaskRunLifeCycleState = `SKIPPED`

const CleanRoomTaskRunLifeCycleStateTerminated CleanRoomTaskRunLifeCycleState = `TERMINATED`

const CleanRoomTaskRunLifeCycleStateTerminating CleanRoomTaskRunLifeCycleState = `TERMINATING`

const CleanRoomTaskRunLifeCycleStateWaitingForRetry CleanRoomTaskRunLifeCycleState = `WAITING_FOR_RETRY`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunLifeCycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunLifeCycleState) Set(v string) error {
	switch v {
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `QUEUED`, `RUNNING`, `RUN_LIFE_CYCLE_STATE_UNSPECIFIED`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = CleanRoomTaskRunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "QUEUED", "RUNNING", "RUN_LIFE_CYCLE_STATE_UNSPECIFIED", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
	}
}

// Type always returns CleanRoomTaskRunLifeCycleState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunLifeCycleState) Type() string {
	return "CleanRoomTaskRunLifeCycleState"
}

func cleanRoomTaskRunLifeCycleStateToPb(st *CleanRoomTaskRunLifeCycleState) (*cleanRoomTaskRunLifeCycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanRoomTaskRunLifeCycleStatePb(*st)
	return &pb, nil
}

func cleanRoomTaskRunLifeCycleStateFromPb(pb *cleanRoomTaskRunLifeCycleStatePb) (*CleanRoomTaskRunLifeCycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomTaskRunLifeCycleState(*pb)
	return &st, nil
}

// Copied from elastic-spark-common/api/messages/runs.proto. Using the original
// definition to avoid cyclic dependency.
type CleanRoomTaskRunResultState string
type cleanRoomTaskRunResultStatePb string

const CleanRoomTaskRunResultStateCanceled CleanRoomTaskRunResultState = `CANCELED`

const CleanRoomTaskRunResultStateDisabled CleanRoomTaskRunResultState = `DISABLED`

const CleanRoomTaskRunResultStateEvicted CleanRoomTaskRunResultState = `EVICTED`

const CleanRoomTaskRunResultStateExcluded CleanRoomTaskRunResultState = `EXCLUDED`

const CleanRoomTaskRunResultStateFailed CleanRoomTaskRunResultState = `FAILED`

const CleanRoomTaskRunResultStateMaximumConcurrentRunsReached CleanRoomTaskRunResultState = `MAXIMUM_CONCURRENT_RUNS_REACHED`

const CleanRoomTaskRunResultStateRunResultStateUnspecified CleanRoomTaskRunResultState = `RUN_RESULT_STATE_UNSPECIFIED`

const CleanRoomTaskRunResultStateSuccess CleanRoomTaskRunResultState = `SUCCESS`

const CleanRoomTaskRunResultStateSuccessWithFailures CleanRoomTaskRunResultState = `SUCCESS_WITH_FAILURES`

const CleanRoomTaskRunResultStateTimedout CleanRoomTaskRunResultState = `TIMEDOUT`

const CleanRoomTaskRunResultStateUpstreamCanceled CleanRoomTaskRunResultState = `UPSTREAM_CANCELED`

const CleanRoomTaskRunResultStateUpstreamEvicted CleanRoomTaskRunResultState = `UPSTREAM_EVICTED`

const CleanRoomTaskRunResultStateUpstreamFailed CleanRoomTaskRunResultState = `UPSTREAM_FAILED`

// String representation for [fmt.Print]
func (f *CleanRoomTaskRunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CleanRoomTaskRunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `DISABLED`, `EVICTED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `RUN_RESULT_STATE_UNSPECIFIED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_EVICTED`, `UPSTREAM_FAILED`:
		*f = CleanRoomTaskRunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "DISABLED", "EVICTED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "RUN_RESULT_STATE_UNSPECIFIED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_EVICTED", "UPSTREAM_FAILED"`, v)
	}
}

// Type always returns CleanRoomTaskRunResultState to satisfy [pflag.Value] interface
func (f *CleanRoomTaskRunResultState) Type() string {
	return "CleanRoomTaskRunResultState"
}

func cleanRoomTaskRunResultStateToPb(st *CleanRoomTaskRunResultState) (*cleanRoomTaskRunResultStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := cleanRoomTaskRunResultStatePb(*st)
	return &pb, nil
}

func cleanRoomTaskRunResultStateFromPb(pb *cleanRoomTaskRunResultStatePb) (*CleanRoomTaskRunResultState, error) {
	if pb == nil {
		return nil, nil
	}
	st := CleanRoomTaskRunResultState(*pb)
	return &st, nil
}

// Stores the run state of the clean rooms notebook task.
type CleanRoomTaskRunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	LifeCycleState CleanRoomTaskRunLifeCycleState
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	ResultState CleanRoomTaskRunResultState
}

func CleanRoomTaskRunStateToPb(st *CleanRoomTaskRunState) (*CleanRoomTaskRunStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &CleanRoomTaskRunStatePb{}
	lifeCycleStatePb, err := identity(&st.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStatePb != nil {
		pb.LifeCycleState = *lifeCycleStatePb
	}

	resultStatePb, err := identity(&st.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStatePb != nil {
		pb.ResultState = *resultStatePb
	}

	return pb, nil
}

func (st *CleanRoomTaskRunState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &CleanRoomTaskRunStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CleanRoomTaskRunStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomTaskRunState) MarshalJSON() ([]byte, error) {
	pb, err := CleanRoomTaskRunStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CleanRoomTaskRunStatePb struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	LifeCycleState CleanRoomTaskRunLifeCycleState `json:"life_cycle_state,omitempty"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	ResultState CleanRoomTaskRunResultState `json:"result_state,omitempty"`
}

func CleanRoomTaskRunStateFromPb(pb *CleanRoomTaskRunStatePb) (*CleanRoomTaskRunState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomTaskRunState{}
	lifeCycleStateField, err := identity(&pb.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStateField != nil {
		st.LifeCycleState = *lifeCycleStateField
	}
	resultStateField, err := identity(&pb.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStateField != nil {
		st.ResultState = *resultStateField
	}

	return st, nil
}

type CleanRoomsNotebookTask struct {
	// The clean room that the notebook belongs to.
	CleanRoomName string
	// Checksum to validate the freshness of the notebook resource (i.e. the
	// notebook being run is the latest version). It can be fetched by calling
	// the :method:cleanroomassets/get API.
	Etag string
	// Base parameters to be used for the clean room notebook job.
	NotebookBaseParameters map[string]string
	// Name of the notebook being run.
	NotebookName string

	ForceSendFields []string
}

func cleanRoomsNotebookTaskToPb(st *CleanRoomsNotebookTask) (*cleanRoomsNotebookTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomsNotebookTaskPb{}
	cleanRoomNamePb, err := identity(&st.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNamePb != nil {
		pb.CleanRoomName = *cleanRoomNamePb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	notebookBaseParametersPb := map[string]string{}
	for k, v := range st.NotebookBaseParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebookBaseParametersPb[k] = *itemPb
		}
	}
	pb.NotebookBaseParameters = notebookBaseParametersPb

	notebookNamePb, err := identity(&st.NotebookName)
	if err != nil {
		return nil, err
	}
	if notebookNamePb != nil {
		pb.NotebookName = *notebookNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CleanRoomsNotebookTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomsNotebookTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomsNotebookTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomsNotebookTask) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomsNotebookTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomsNotebookTaskPb struct {
	// The clean room that the notebook belongs to.
	CleanRoomName string `json:"clean_room_name"`
	// Checksum to validate the freshness of the notebook resource (i.e. the
	// notebook being run is the latest version). It can be fetched by calling
	// the :method:cleanroomassets/get API.
	Etag string `json:"etag,omitempty"`
	// Base parameters to be used for the clean room notebook job.
	NotebookBaseParameters map[string]string `json:"notebook_base_parameters,omitempty"`
	// Name of the notebook being run.
	NotebookName string `json:"notebook_name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cleanRoomsNotebookTaskFromPb(pb *cleanRoomsNotebookTaskPb) (*CleanRoomsNotebookTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomsNotebookTask{}
	cleanRoomNameField, err := identity(&pb.CleanRoomName)
	if err != nil {
		return nil, err
	}
	if cleanRoomNameField != nil {
		st.CleanRoomName = *cleanRoomNameField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	notebookBaseParametersField := map[string]string{}
	for k, v := range pb.NotebookBaseParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notebookBaseParametersField[k] = *item
		}
	}
	st.NotebookBaseParameters = notebookBaseParametersField
	notebookNameField, err := identity(&pb.NotebookName)
	if err != nil {
		return nil, err
	}
	if notebookNameField != nil {
		st.NotebookName = *notebookNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cleanRoomsNotebookTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cleanRoomsNotebookTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput struct {
	// The run state of the clean rooms notebook task.
	CleanRoomJobRunState *CleanRoomTaskRunState
	// The notebook output for the clean room run
	NotebookOutput *NotebookOutput
	// Information on how to access the output schema for the clean room run
	OutputSchemaInfo *OutputSchemaInfo
}

func cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(st *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) (*cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb{}
	cleanRoomJobRunStatePb, err := CleanRoomTaskRunStateToPb(st.CleanRoomJobRunState)
	if err != nil {
		return nil, err
	}
	if cleanRoomJobRunStatePb != nil {
		pb.CleanRoomJobRunState = cleanRoomJobRunStatePb
	}

	notebookOutputPb, err := notebookOutputToPb(st.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputPb != nil {
		pb.NotebookOutput = notebookOutputPb
	}

	outputSchemaInfoPb, err := outputSchemaInfoToPb(st.OutputSchemaInfo)
	if err != nil {
		return nil, err
	}
	if outputSchemaInfoPb != nil {
		pb.OutputSchemaInfo = outputSchemaInfoPb
	}

	return pb, nil
}

func (st *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb struct {
	// The run state of the clean rooms notebook task.
	CleanRoomJobRunState *CleanRoomTaskRunStatePb `json:"clean_room_job_run_state,omitempty"`
	// The notebook output for the clean room run
	NotebookOutput *notebookOutputPb `json:"notebook_output,omitempty"`
	// Information on how to access the output schema for the clean room run
	OutputSchemaInfo *outputSchemaInfoPb `json:"output_schema_info,omitempty"`
}

func cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb *cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb) (*CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput{}
	cleanRoomJobRunStateField, err := CleanRoomTaskRunStateFromPb(pb.CleanRoomJobRunState)
	if err != nil {
		return nil, err
	}
	if cleanRoomJobRunStateField != nil {
		st.CleanRoomJobRunState = cleanRoomJobRunStateField
	}
	notebookOutputField, err := notebookOutputFromPb(pb.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputField != nil {
		st.NotebookOutput = notebookOutputField
	}
	outputSchemaInfoField, err := outputSchemaInfoFromPb(pb.OutputSchemaInfo)
	if err != nil {
		return nil, err
	}
	if outputSchemaInfoField != nil {
		st.OutputSchemaInfo = outputSchemaInfoField
	}

	return st, nil
}

type ClusterInstance struct {
	// The canonical identifier for the cluster used by a run. This field is
	// always available for runs on existing clusters. For runs on new clusters,
	// it becomes available once the cluster is created. This value can be used
	// to view logs by browsing to `/#setting/sparkui/$cluster_id/driver-logs`.
	// The logs continue to be available after the run completes.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	ClusterId string
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	SparkContextId string

	ForceSendFields []string
}

func clusterInstanceToPb(st *ClusterInstance) (*clusterInstancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterInstancePb{}
	clusterIdPb, err := identity(&st.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdPb != nil {
		pb.ClusterId = *clusterIdPb
	}

	sparkContextIdPb, err := identity(&st.SparkContextId)
	if err != nil {
		return nil, err
	}
	if sparkContextIdPb != nil {
		pb.SparkContextId = *sparkContextIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ClusterInstance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterInstancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterInstanceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterInstance) MarshalJSON() ([]byte, error) {
	pb, err := clusterInstanceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type clusterInstancePb struct {
	// The canonical identifier for the cluster used by a run. This field is
	// always available for runs on existing clusters. For runs on new clusters,
	// it becomes available once the cluster is created. This value can be used
	// to view logs by browsing to `/#setting/sparkui/$cluster_id/driver-logs`.
	// The logs continue to be available after the run completes.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	ClusterId string `json:"cluster_id,omitempty"`
	// The canonical identifier for the Spark context used by a run. This field
	// is filled in once the run begins execution. This value can be used to
	// view the Spark UI by browsing to
	// `/#setting/sparkui/$cluster_id/$spark_context_id`. The Spark UI continues
	// to be available after the run has completed.
	//
	// The response won’t include this field if the identifier is not
	// available yet.
	SparkContextId string `json:"spark_context_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterInstanceFromPb(pb *clusterInstancePb) (*ClusterInstance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterInstance{}
	clusterIdField, err := identity(&pb.ClusterId)
	if err != nil {
		return nil, err
	}
	if clusterIdField != nil {
		st.ClusterId = *clusterIdField
	}
	sparkContextIdField, err := identity(&pb.SparkContextId)
	if err != nil {
		return nil, err
	}
	if sparkContextIdField != nil {
		st.SparkContextId = *sparkContextIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterInstancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterInstancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterSpec struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec

	ForceSendFields []string
}

func clusterSpecToPb(st *ClusterSpec) (*clusterSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterSpecPb{}
	existingClusterIdPb, err := identity(&st.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdPb != nil {
		pb.ExistingClusterId = *existingClusterIdPb
	}

	jobClusterKeyPb, err := identity(&st.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyPb != nil {
		pb.JobClusterKey = *jobClusterKeyPb
	}

	var librariesPb []compute.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterSpecPb struct {
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.LibraryPb `json:"libraries,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpecPb `json:"new_cluster,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterSpecFromPb(pb *clusterSpecPb) (*ClusterSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterSpec{}
	existingClusterIdField, err := identity(&pb.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdField != nil {
		st.ExistingClusterId = *existingClusterIdField
	}
	jobClusterKeyField, err := identity(&pb.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyField != nil {
		st.JobClusterKey = *jobClusterKeyField
	}

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComputeConfig struct {
	// IDof the GPU pool to use.
	GpuNodePoolId string
	// GPU type.
	GpuType string
	// Number of GPUs.
	NumGpus int

	ForceSendFields []string
}

func computeConfigToPb(st *ComputeConfig) (*computeConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &computeConfigPb{}
	gpuNodePoolIdPb, err := identity(&st.GpuNodePoolId)
	if err != nil {
		return nil, err
	}
	if gpuNodePoolIdPb != nil {
		pb.GpuNodePoolId = *gpuNodePoolIdPb
	}

	gpuTypePb, err := identity(&st.GpuType)
	if err != nil {
		return nil, err
	}
	if gpuTypePb != nil {
		pb.GpuType = *gpuTypePb
	}

	numGpusPb, err := identity(&st.NumGpus)
	if err != nil {
		return nil, err
	}
	if numGpusPb != nil {
		pb.NumGpus = *numGpusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ComputeConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &computeConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := computeConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ComputeConfig) MarshalJSON() ([]byte, error) {
	pb, err := computeConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type computeConfigPb struct {
	// IDof the GPU pool to use.
	GpuNodePoolId string `json:"gpu_node_pool_id"`
	// GPU type.
	GpuType string `json:"gpu_type,omitempty"`
	// Number of GPUs.
	NumGpus int `json:"num_gpus"`

	ForceSendFields []string `json:"-" url:"-"`
}

func computeConfigFromPb(pb *computeConfigPb) (*ComputeConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComputeConfig{}
	gpuNodePoolIdField, err := identity(&pb.GpuNodePoolId)
	if err != nil {
		return nil, err
	}
	if gpuNodePoolIdField != nil {
		st.GpuNodePoolId = *gpuNodePoolIdField
	}
	gpuTypeField, err := identity(&pb.GpuType)
	if err != nil {
		return nil, err
	}
	if gpuTypeField != nil {
		st.GpuType = *gpuTypeField
	}
	numGpusField, err := identity(&pb.NumGpus)
	if err != nil {
		return nil, err
	}
	if numGpusField != nil {
		st.NumGpus = *numGpusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *computeConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st computeConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Condition string
type conditionPb string

const ConditionAllUpdated Condition = `ALL_UPDATED`

const ConditionAnyUpdated Condition = `ANY_UPDATED`

// String representation for [fmt.Print]
func (f *Condition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Condition) Set(v string) error {
	switch v {
	case `ALL_UPDATED`, `ANY_UPDATED`:
		*f = Condition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_UPDATED", "ANY_UPDATED"`, v)
	}
}

// Type always returns Condition to satisfy [pflag.Value] interface
func (f *Condition) Type() string {
	return "Condition"
}

func conditionToPb(st *Condition) (*conditionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := conditionPb(*st)
	return &pb, nil
}

func conditionFromPb(pb *conditionPb) (*Condition, error) {
	if pb == nil {
		return nil, nil
	}
	st := Condition(*pb)
	return &st, nil
}

type ConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left string
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	Op ConditionTaskOp
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right string
}

func conditionTaskToPb(st *ConditionTask) (*conditionTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &conditionTaskPb{}
	leftPb, err := identity(&st.Left)
	if err != nil {
		return nil, err
	}
	if leftPb != nil {
		pb.Left = *leftPb
	}

	opPb, err := identity(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}

	rightPb, err := identity(&st.Right)
	if err != nil {
		return nil, err
	}
	if rightPb != nil {
		pb.Right = *rightPb
	}

	return pb, nil
}

func (st *ConditionTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &conditionTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := conditionTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ConditionTask) MarshalJSON() ([]byte, error) {
	pb, err := conditionTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type conditionTaskPb struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left string `json:"left"`
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	Op ConditionTaskOp `json:"op"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right string `json:"right"`
}

func conditionTaskFromPb(pb *conditionTaskPb) (*ConditionTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ConditionTask{}
	leftField, err := identity(&pb.Left)
	if err != nil {
		return nil, err
	}
	if leftField != nil {
		st.Left = *leftField
	}
	opField, err := identity(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	rightField, err := identity(&pb.Right)
	if err != nil {
		return nil, err
	}
	if rightField != nil {
		st.Right = *rightField
	}

	return st, nil
}

// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
// operands. This means that `“12.0” == “12”` will evaluate to `false`.
// * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`
// operators perform numeric comparison of their operands. `“12.0” >=
// “12”` will evaluate to `true`, `“10.0” >= “12”` will evaluate to
// `false`.
//
// The boolean comparison to task values can be implemented with operators
// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it will
// be serialized to `“true”` or `“false”` for the comparison.
type ConditionTaskOp string
type conditionTaskOpPb string

const ConditionTaskOpEqualTo ConditionTaskOp = `EQUAL_TO`

const ConditionTaskOpGreaterThan ConditionTaskOp = `GREATER_THAN`

const ConditionTaskOpGreaterThanOrEqual ConditionTaskOp = `GREATER_THAN_OR_EQUAL`

const ConditionTaskOpLessThan ConditionTaskOp = `LESS_THAN`

const ConditionTaskOpLessThanOrEqual ConditionTaskOp = `LESS_THAN_OR_EQUAL`

const ConditionTaskOpNotEqual ConditionTaskOp = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *ConditionTaskOp) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ConditionTaskOp) Set(v string) error {
	switch v {
	case `EQUAL_TO`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = ConditionTaskOp(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL_TO", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Type always returns ConditionTaskOp to satisfy [pflag.Value] interface
func (f *ConditionTaskOp) Type() string {
	return "ConditionTaskOp"
}

func conditionTaskOpToPb(st *ConditionTaskOp) (*conditionTaskOpPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := conditionTaskOpPb(*st)
	return &pb, nil
}

func conditionTaskOpFromPb(pb *conditionTaskOpPb) (*ConditionTaskOp, error) {
	if pb == nil {
		return nil, nil
	}
	st := ConditionTaskOp(*pb)
	return &st, nil
}

type Continuous struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	PauseStatus PauseStatus
}

func continuousToPb(st *Continuous) (*continuousPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &continuousPb{}
	pauseStatusPb, err := identity(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}

	return pb, nil
}

func (st *Continuous) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &continuousPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := continuousFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Continuous) MarshalJSON() ([]byte, error) {
	pb, err := continuousToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type continuousPb struct {
	// Indicate whether the continuous execution of the job is paused or not.
	// Defaults to UNPAUSED.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
}

func continuousFromPb(pb *continuousPb) (*Continuous, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Continuous{}
	pauseStatusField, err := identity(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}

	return st, nil
}

type CreateJob struct {
	// List of permissions to set on the job.
	AccessControlList []JobAccessControlRequest
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId string
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description string
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	Environments []JobEnvironment
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	MaxConcurrentRuns int
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings
	// Job-level parameter definitions
	Parameters []JobParameterDefinition
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget
	// The queue settings of the job.
	Queue *QueueSettings
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	RunAs *JobRunAs
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	Tasks []Task
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications

	WellKnown *AllWellKnown

	ForceSendFields []string
}

func createJobToPb(st *CreateJob) (*createJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createJobPb{}

	var accessControlListPb []jobAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := jobAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	continuousPb, err := continuousToPb(st.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousPb != nil {
		pb.Continuous = continuousPb
	}

	deploymentPb, err := jobDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	editModePb, err := identity(&st.EditMode)
	if err != nil {
		return nil, err
	}
	if editModePb != nil {
		pb.EditMode = *editModePb
	}

	emailNotificationsPb, err := jobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	var environmentsPb []jobEnvironmentPb
	for _, item := range st.Environments {
		itemPb, err := jobEnvironmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentsPb = append(environmentsPb, *itemPb)
		}
	}
	pb.Environments = environmentsPb

	formatPb, err := identity(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}

	gitSourcePb, err := gitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}

	healthPb, err := jobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	var jobClustersPb []jobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := jobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb

	maxConcurrentRunsPb, err := identity(&st.MaxConcurrentRuns)
	if err != nil {
		return nil, err
	}
	if maxConcurrentRunsPb != nil {
		pb.MaxConcurrentRuns = *maxConcurrentRunsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	notificationSettingsPb, err := jobNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	var parametersPb []jobParameterDefinitionPb
	for _, item := range st.Parameters {
		itemPb, err := jobParameterDefinitionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	performanceTargetPb, err := identity(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}

	queuePb, err := queueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}

	runAsPb, err := jobRunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}

	schedulePb, err := cronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	tagsPb := map[string]string{}
	for k, v := range st.Tags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb[k] = *itemPb
		}
	}
	pb.Tags = tagsPb

	var tasksPb []taskPb
	for _, item := range st.Tasks {
		itemPb, err := taskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb

	timeoutSecondsPb, err := identity(&st.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsPb != nil {
		pb.TimeoutSeconds = *timeoutSecondsPb
	}

	triggerPb, err := triggerSettingsToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	webhookNotificationsPb, err := webhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	wellKnownPb, err := allWellKnownToPb(st.WellKnown)
	if err != nil {
		return nil, err
	}
	if wellKnownPb != nil {
		pb.WellKnown = wellKnownPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateJob) MarshalJSON() ([]byte, error) {
	pb, err := createJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createJobPb struct {
	// List of permissions to set on the job.
	AccessControlList []jobAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *continuousPb `json:"continuous,omitempty"`
	// Deployment information for jobs managed by external sources.
	Deployment *jobDeploymentPb `json:"deployment,omitempty"`
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description string `json:"description,omitempty"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode `json:"edit_mode,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *jobEmailNotificationsPb `json:"email_notifications,omitempty"`
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	Environments []jobEnvironmentPb `json:"environments,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `json:"format,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *gitSourcePb `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *jobsHealthRulesPb `json:"health,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []jobClusterPb `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *jobNotificationSettingsPb `json:"notification_settings,omitempty"`
	// Job-level parameter definitions
	Parameters []jobParameterDefinitionPb `json:"parameters,omitempty"`
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// The queue settings of the job.
	Queue *queueSettingsPb `json:"queue,omitempty"`
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	RunAs *jobRunAsPb `json:"run_as,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *cronSchedulePb `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	Tasks []taskPb `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *triggerSettingsPb `json:"trigger,omitempty"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *webhookNotificationsPb `json:"webhook_notifications,omitempty"`

	WellKnown *allWellKnownPb `json:"well_known,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createJobFromPb(pb *createJobPb) (*CreateJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateJob{}

	var accessControlListField []JobAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := jobAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	continuousField, err := continuousFromPb(pb.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousField != nil {
		st.Continuous = continuousField
	}
	deploymentField, err := jobDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	editModeField, err := identity(&pb.EditMode)
	if err != nil {
		return nil, err
	}
	if editModeField != nil {
		st.EditMode = *editModeField
	}
	emailNotificationsField, err := jobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}

	var environmentsField []JobEnvironment
	for _, itemPb := range pb.Environments {
		item, err := jobEnvironmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentsField = append(environmentsField, *item)
		}
	}
	st.Environments = environmentsField
	formatField, err := identity(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	gitSourceField, err := gitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	healthField, err := jobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := jobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	maxConcurrentRunsField, err := identity(&pb.MaxConcurrentRuns)
	if err != nil {
		return nil, err
	}
	if maxConcurrentRunsField != nil {
		st.MaxConcurrentRuns = *maxConcurrentRunsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	notificationSettingsField, err := jobNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}

	var parametersField []JobParameterDefinition
	for _, itemPb := range pb.Parameters {
		item, err := jobParameterDefinitionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	performanceTargetField, err := identity(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	queueField, err := queueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}
	runAsField, err := jobRunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	scheduleField, err := cronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}

	tagsField := map[string]string{}
	for k, v := range pb.Tags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField[k] = *item
		}
	}
	st.Tags = tagsField

	var tasksField []Task
	for _, itemPb := range pb.Tasks {
		item, err := taskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	timeoutSecondsField, err := identity(&pb.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsField != nil {
		st.TimeoutSeconds = *timeoutSecondsField
	}
	triggerField, err := triggerSettingsFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}
	webhookNotificationsField, err := webhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}
	wellKnownField, err := allWellKnownFromPb(pb.WellKnown)
	if err != nil {
		return nil, err
	}
	if wellKnownField != nil {
		st.WellKnown = wellKnownField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createJobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createJobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Job was created successfully
type CreateResponse struct {
	// The canonical identifier for the newly created job.
	JobId int64

	ForceSendFields []string
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
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
	// The canonical identifier for the newly created job.
	JobId int64 `json:"job_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
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

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus PauseStatus
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string
}

func cronScheduleToPb(st *CronSchedule) (*cronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronSchedulePb{}
	pauseStatusPb, err := identity(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}

	quartzCronExpressionPb, err := identity(&st.QuartzCronExpression)
	if err != nil {
		return nil, err
	}
	if quartzCronExpressionPb != nil {
		pb.QuartzCronExpression = *quartzCronExpressionPb
	}

	timezoneIdPb, err := identity(&st.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdPb != nil {
		pb.TimezoneId = *timezoneIdPb
	}

	return pb, nil
}

func (st *CronSchedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cronSchedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cronScheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CronSchedule) MarshalJSON() ([]byte, error) {
	pb, err := cronScheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cronSchedulePb struct {
	// Indicate whether this schedule is paused or not.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// A Cron expression using Quartz syntax that describes the schedule for a
	// job. See [Cron Trigger] for details. This field is required.
	//
	// [Cron Trigger]: http://www.quartz-scheduler.org/documentation/quartz-2.3.0/tutorials/crontrigger.html
	QuartzCronExpression string `json:"quartz_cron_expression"`
	// A Java timezone ID. The schedule for a job is resolved with respect to
	// this timezone. See [Java TimeZone] for details. This field is required.
	//
	// [Java TimeZone]: https://docs.oracle.com/javase/7/docs/api/java/util/TimeZone.html
	TimezoneId string `json:"timezone_id"`
}

func cronScheduleFromPb(pb *cronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	pauseStatusField, err := identity(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	quartzCronExpressionField, err := identity(&pb.QuartzCronExpression)
	if err != nil {
		return nil, err
	}
	if quartzCronExpressionField != nil {
		st.QuartzCronExpression = *quartzCronExpressionField
	}
	timezoneIdField, err := identity(&pb.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdField != nil {
		st.TimezoneId = *timezoneIdField
	}

	return st, nil
}

type DashboardPageSnapshot struct {
	PageDisplayName string

	WidgetErrorDetails []WidgetErrorDetail

	ForceSendFields []string
}

func dashboardPageSnapshotToPb(st *DashboardPageSnapshot) (*dashboardPageSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPageSnapshotPb{}
	pageDisplayNamePb, err := identity(&st.PageDisplayName)
	if err != nil {
		return nil, err
	}
	if pageDisplayNamePb != nil {
		pb.PageDisplayName = *pageDisplayNamePb
	}

	var widgetErrorDetailsPb []widgetErrorDetailPb
	for _, item := range st.WidgetErrorDetails {
		itemPb, err := widgetErrorDetailToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			widgetErrorDetailsPb = append(widgetErrorDetailsPb, *itemPb)
		}
	}
	pb.WidgetErrorDetails = widgetErrorDetailsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DashboardPageSnapshot) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardPageSnapshotPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardPageSnapshotFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardPageSnapshot) MarshalJSON() ([]byte, error) {
	pb, err := dashboardPageSnapshotToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dashboardPageSnapshotPb struct {
	PageDisplayName string `json:"page_display_name,omitempty"`

	WidgetErrorDetails []widgetErrorDetailPb `json:"widget_error_details,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardPageSnapshotFromPb(pb *dashboardPageSnapshotPb) (*DashboardPageSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardPageSnapshot{}
	pageDisplayNameField, err := identity(&pb.PageDisplayName)
	if err != nil {
		return nil, err
	}
	if pageDisplayNameField != nil {
		st.PageDisplayName = *pageDisplayNameField
	}

	var widgetErrorDetailsField []WidgetErrorDetail
	for _, itemPb := range pb.WidgetErrorDetails {
		item, err := widgetErrorDetailFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			widgetErrorDetailsField = append(widgetErrorDetailsField, *item)
		}
	}
	st.WidgetErrorDetails = widgetErrorDetailsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardPageSnapshotPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardPageSnapshotPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Configures the Lakeview Dashboard job task type.
type DashboardTask struct {
	// The identifier of the dashboard to refresh.
	DashboardId string
	// Optional: subscription configuration for sending the dashboard snapshot.
	Subscription *Subscription
	// Optional: The warehouse id to execute the dashboard with for the
	// schedule. If not specified, the default warehouse of the dashboard will
	// be used.
	WarehouseId string

	ForceSendFields []string
}

func dashboardTaskToPb(st *DashboardTask) (*dashboardTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardTaskPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	subscriptionPb, err := subscriptionToPb(st.Subscription)
	if err != nil {
		return nil, err
	}
	if subscriptionPb != nil {
		pb.Subscription = subscriptionPb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DashboardTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardTask) MarshalJSON() ([]byte, error) {
	pb, err := dashboardTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dashboardTaskPb struct {
	// The identifier of the dashboard to refresh.
	DashboardId string `json:"dashboard_id,omitempty"`
	// Optional: subscription configuration for sending the dashboard snapshot.
	Subscription *subscriptionPb `json:"subscription,omitempty"`
	// Optional: The warehouse id to execute the dashboard with for the
	// schedule. If not specified, the default warehouse of the dashboard will
	// be used.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardTaskFromPb(pb *dashboardTaskPb) (*DashboardTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardTask{}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	subscriptionField, err := subscriptionFromPb(pb.Subscription)
	if err != nil {
		return nil, err
	}
	if subscriptionField != nil {
		st.Subscription = subscriptionField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardTaskOutput struct {
	// Should only be populated for manual PDF download jobs.
	PageSnapshots []DashboardPageSnapshot
}

func dashboardTaskOutputToPb(st *DashboardTaskOutput) (*dashboardTaskOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardTaskOutputPb{}

	var pageSnapshotsPb []dashboardPageSnapshotPb
	for _, item := range st.PageSnapshots {
		itemPb, err := dashboardPageSnapshotToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pageSnapshotsPb = append(pageSnapshotsPb, *itemPb)
		}
	}
	pb.PageSnapshots = pageSnapshotsPb

	return pb, nil
}

func (st *DashboardTaskOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardTaskOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardTaskOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardTaskOutput) MarshalJSON() ([]byte, error) {
	pb, err := dashboardTaskOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dashboardTaskOutputPb struct {
	// Should only be populated for manual PDF download jobs.
	PageSnapshots []dashboardPageSnapshotPb `json:"page_snapshots,omitempty"`
}

func dashboardTaskOutputFromPb(pb *dashboardTaskOutputPb) (*DashboardTaskOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardTaskOutput{}

	var pageSnapshotsField []DashboardPageSnapshot
	for _, itemPb := range pb.PageSnapshots {
		item, err := dashboardPageSnapshotFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pageSnapshotsField = append(pageSnapshotsField, *item)
		}
	}
	st.PageSnapshots = pageSnapshotsField

	return st, nil
}

type DbtOutput struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders map[string]string
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink string

	ForceSendFields []string
}

func dbtOutputToPb(st *DbtOutput) (*dbtOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dbtOutputPb{}

	artifactsHeadersPb := map[string]string{}
	for k, v := range st.ArtifactsHeaders {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			artifactsHeadersPb[k] = *itemPb
		}
	}
	pb.ArtifactsHeaders = artifactsHeadersPb

	artifactsLinkPb, err := identity(&st.ArtifactsLink)
	if err != nil {
		return nil, err
	}
	if artifactsLinkPb != nil {
		pb.ArtifactsLink = *artifactsLinkPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DbtOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dbtOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dbtOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DbtOutput) MarshalJSON() ([]byte, error) {
	pb, err := dbtOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dbtOutputPb struct {
	// An optional map of headers to send when retrieving the artifact from the
	// `artifacts_link`.
	ArtifactsHeaders map[string]string `json:"artifacts_headers,omitempty"`
	// A pre-signed URL to download the (compressed) dbt artifacts. This link is
	// valid for a limited time (30 minutes). This information is only available
	// after the run has finished.
	ArtifactsLink string `json:"artifacts_link,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dbtOutputFromPb(pb *dbtOutputPb) (*DbtOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtOutput{}

	artifactsHeadersField := map[string]string{}
	for k, v := range pb.ArtifactsHeaders {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			artifactsHeadersField[k] = *item
		}
	}
	st.ArtifactsHeaders = artifactsHeadersField
	artifactsLinkField, err := identity(&pb.ArtifactsLink)
	if err != nil {
		return nil, err
	}
	if artifactsLinkField != nil {
		st.ArtifactsLink = *artifactsLinkField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dbtOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dbtOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DbtTask struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	Catalog string
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	Commands []string
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	ProfilesDirectory string
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	ProjectDirectory string
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	Schema string
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in Databricks workspace. * `GIT`:
	// Project is located in cloud Git provider.
	Source Source
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	WarehouseId string

	ForceSendFields []string
}

func dbtTaskToPb(st *DbtTask) (*dbtTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dbtTaskPb{}
	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	var commandsPb []string
	for _, item := range st.Commands {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			commandsPb = append(commandsPb, *itemPb)
		}
	}
	pb.Commands = commandsPb

	profilesDirectoryPb, err := identity(&st.ProfilesDirectory)
	if err != nil {
		return nil, err
	}
	if profilesDirectoryPb != nil {
		pb.ProfilesDirectory = *profilesDirectoryPb
	}

	projectDirectoryPb, err := identity(&st.ProjectDirectory)
	if err != nil {
		return nil, err
	}
	if projectDirectoryPb != nil {
		pb.ProjectDirectory = *projectDirectoryPb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	sourcePb, err := identity(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DbtTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dbtTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dbtTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DbtTask) MarshalJSON() ([]byte, error) {
	pb, err := dbtTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type dbtTaskPb struct {
	// Optional name of the catalog to use. The value is the top level in the
	// 3-level namespace of Unity Catalog (catalog / schema / relation). The
	// catalog value can only be specified if a warehouse_id is specified.
	// Requires dbt-databricks >= 1.1.1.
	Catalog string `json:"catalog,omitempty"`
	// A list of dbt commands to execute. All commands must start with `dbt`.
	// This parameter must not be empty. A maximum of up to 10 commands can be
	// provided.
	Commands []string `json:"commands"`
	// Optional (relative) path to the profiles directory. Can only be specified
	// if no warehouse_id is specified. If no warehouse_id is specified and this
	// folder is unset, the root directory is used.
	ProfilesDirectory string `json:"profiles_directory,omitempty"`
	// Path to the project directory. Optional for Git sourced tasks, in which
	// case if no value is provided, the root of the Git repository is used.
	ProjectDirectory string `json:"project_directory,omitempty"`
	// Optional schema to write to. This parameter is only used when a
	// warehouse_id is also provided. If not provided, the `default` schema is
	// used.
	Schema string `json:"schema,omitempty"`
	// Optional location type of the project directory. When set to `WORKSPACE`,
	// the project will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the project will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: Project is located in Databricks workspace. * `GIT`:
	// Project is located in cloud Git provider.
	Source Source `json:"source,omitempty"`
	// ID of the SQL warehouse to connect to. If provided, we automatically
	// generate and provide the profile and connection details to dbt. It can be
	// overridden on a per-command basis by using the `--profiles-dir` command
	// line argument.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dbtTaskFromPb(pb *dbtTaskPb) (*DbtTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DbtTask{}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}

	var commandsField []string
	for _, itemPb := range pb.Commands {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			commandsField = append(commandsField, *item)
		}
	}
	st.Commands = commandsField
	profilesDirectoryField, err := identity(&pb.ProfilesDirectory)
	if err != nil {
		return nil, err
	}
	if profilesDirectoryField != nil {
		st.ProfilesDirectory = *profilesDirectoryField
	}
	projectDirectoryField, err := identity(&pb.ProjectDirectory)
	if err != nil {
		return nil, err
	}
	if projectDirectoryField != nil {
		st.ProjectDirectory = *projectDirectoryField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	sourceField, err := identity(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dbtTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dbtTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteJob struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId int64
}

func deleteJobToPb(st *DeleteJob) (*deleteJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteJobPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	return pb, nil
}

func (st *DeleteJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteJob) MarshalJSON() ([]byte, error) {
	pb, err := deleteJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteJobPb struct {
	// The canonical identifier of the job to delete. This field is required.
	JobId int64 `json:"job_id"`
}

func deleteJobFromPb(pb *deleteJobPb) (*DeleteJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteJob{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

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

type DeleteRun struct {
	// ID of the run to delete.
	RunId int64
}

func deleteRunToPb(st *DeleteRun) (*deleteRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunPb{}
	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	return pb, nil
}

func (st *DeleteRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRun) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRunPb struct {
	// ID of the run to delete.
	RunId int64 `json:"run_id"`
}

func deleteRunFromPb(pb *deleteRunPb) (*DeleteRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRun{}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	return st, nil
}

type DeleteRunResponse struct {
}

func deleteRunResponseToPb(st *DeleteRunResponse) (*deleteRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunResponsePb{}

	return pb, nil
}

func (st *DeleteRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRunResponsePb struct {
}

func deleteRunResponseFromPb(pb *deleteRunResponsePb) (*DeleteRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRunResponse{}

	return st, nil
}

// Represents a change to the job cluster's settings that would be required for
// the job clusters to become compliant with their policies.
type EnforcePolicyComplianceForJobResponseJobClusterSettingsChange struct {
	// The field where this change would be made, prepended with the job cluster
	// key.
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

func enforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(st *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) (*enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb{}
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

func (st *EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforcePolicyComplianceForJobResponseJobClusterSettingsChange) MarshalJSON() ([]byte, error) {
	pb, err := enforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb struct {
	// The field where this change would be made, prepended with the job cluster
	// key.
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

func enforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(pb *enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) (*EnforcePolicyComplianceForJobResponseJobClusterSettingsChange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceForJobResponseJobClusterSettingsChange{}
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

func (st *enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnforcePolicyComplianceRequest struct {
	// The ID of the job you want to enforce policy compliance on.
	JobId int64
	// If set, previews changes made to the job to comply with its policy, but
	// does not update the job.
	ValidateOnly bool

	ForceSendFields []string
}

func enforcePolicyComplianceRequestToPb(st *EnforcePolicyComplianceRequest) (*enforcePolicyComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforcePolicyComplianceRequestPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
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

func (st *EnforcePolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforcePolicyComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforcePolicyComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforcePolicyComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := enforcePolicyComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type enforcePolicyComplianceRequestPb struct {
	// The ID of the job you want to enforce policy compliance on.
	JobId int64 `json:"job_id"`
	// If set, previews changes made to the job to comply with its policy, but
	// does not update the job.
	ValidateOnly bool `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforcePolicyComplianceRequestFromPb(pb *enforcePolicyComplianceRequestPb) (*EnforcePolicyComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceRequest{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
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

func (st *enforcePolicyComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforcePolicyComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnforcePolicyComplianceResponse struct {
	// Whether any changes have been made to the job cluster settings for the
	// job to become compliant with its policies.
	HasChanges bool
	// A list of job cluster changes that have been made to the job’s cluster
	// settings in order for all job clusters to become compliant with their
	// policies.
	JobClusterChanges []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange
	// Updated job settings after policy enforcement. Policy enforcement only
	// applies to job clusters that are created when running the job (which are
	// specified in new_cluster) and does not apply to existing all-purpose
	// clusters. Updated job settings are derived by applying policy default
	// values to the existing job clusters in order to satisfy policy
	// requirements.
	Settings *JobSettings

	ForceSendFields []string
}

func enforcePolicyComplianceResponseToPb(st *EnforcePolicyComplianceResponse) (*enforcePolicyComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enforcePolicyComplianceResponsePb{}
	hasChangesPb, err := identity(&st.HasChanges)
	if err != nil {
		return nil, err
	}
	if hasChangesPb != nil {
		pb.HasChanges = *hasChangesPb
	}

	var jobClusterChangesPb []enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb
	for _, item := range st.JobClusterChanges {
		itemPb, err := enforcePolicyComplianceForJobResponseJobClusterSettingsChangeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClusterChangesPb = append(jobClusterChangesPb, *itemPb)
		}
	}
	pb.JobClusterChanges = jobClusterChangesPb

	settingsPb, err := jobSettingsToPb(st.Settings)
	if err != nil {
		return nil, err
	}
	if settingsPb != nil {
		pb.Settings = settingsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *EnforcePolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enforcePolicyComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enforcePolicyComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnforcePolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := enforcePolicyComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type enforcePolicyComplianceResponsePb struct {
	// Whether any changes have been made to the job cluster settings for the
	// job to become compliant with its policies.
	HasChanges bool `json:"has_changes,omitempty"`
	// A list of job cluster changes that have been made to the job’s cluster
	// settings in order for all job clusters to become compliant with their
	// policies.
	JobClusterChanges []enforcePolicyComplianceForJobResponseJobClusterSettingsChangePb `json:"job_cluster_changes,omitempty"`
	// Updated job settings after policy enforcement. Policy enforcement only
	// applies to job clusters that are created when running the job (which are
	// specified in new_cluster) and does not apply to existing all-purpose
	// clusters. Updated job settings are derived by applying policy default
	// values to the existing job clusters in order to satisfy policy
	// requirements.
	Settings *jobSettingsPb `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enforcePolicyComplianceResponseFromPb(pb *enforcePolicyComplianceResponsePb) (*EnforcePolicyComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnforcePolicyComplianceResponse{}
	hasChangesField, err := identity(&pb.HasChanges)
	if err != nil {
		return nil, err
	}
	if hasChangesField != nil {
		st.HasChanges = *hasChangesField
	}

	var jobClusterChangesField []EnforcePolicyComplianceForJobResponseJobClusterSettingsChange
	for _, itemPb := range pb.JobClusterChanges {
		item, err := enforcePolicyComplianceForJobResponseJobClusterSettingsChangeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClusterChangesField = append(jobClusterChangesField, *item)
		}
	}
	st.JobClusterChanges = jobClusterChangesField
	settingsField, err := jobSettingsFromPb(pb.Settings)
	if err != nil {
		return nil, err
	}
	if settingsField != nil {
		st.Settings = settingsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enforcePolicyComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enforcePolicyComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Run was exported successfully.
type ExportRunOutput struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	Views []ViewItem
}

func exportRunOutputToPb(st *ExportRunOutput) (*exportRunOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportRunOutputPb{}

	var viewsPb []viewItemPb
	for _, item := range st.Views {
		itemPb, err := viewItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			viewsPb = append(viewsPb, *itemPb)
		}
	}
	pb.Views = viewsPb

	return pb, nil
}

func (st *ExportRunOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportRunOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportRunOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportRunOutput) MarshalJSON() ([]byte, error) {
	pb, err := exportRunOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type exportRunOutputPb struct {
	// The exported content in HTML format (one for every view item). To extract
	// the HTML notebook from the JSON response, download and run this [Python
	// script].
	//
	// [Python script]: https://docs.databricks.com/en/_static/examples/extract.py
	Views []viewItemPb `json:"views,omitempty"`
}

func exportRunOutputFromPb(pb *exportRunOutputPb) (*ExportRunOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRunOutput{}

	var viewsField []ViewItem
	for _, itemPb := range pb.Views {
		item, err := viewItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			viewsField = append(viewsField, *item)
		}
	}
	st.Views = viewsField

	return st, nil
}

// Export and retrieve a job run
type ExportRunRequest struct {
	// The canonical identifier for the run. This field is required.
	RunId int64
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport ViewsToExport
}

func exportRunRequestToPb(st *ExportRunRequest) (*exportRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportRunRequestPb{}
	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	viewsToExportPb, err := identity(&st.ViewsToExport)
	if err != nil {
		return nil, err
	}
	if viewsToExportPb != nil {
		pb.ViewsToExport = *viewsToExportPb
	}

	return pb, nil
}

func (st *ExportRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := exportRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type exportRunRequestPb struct {
	// The canonical identifier for the run. This field is required.
	RunId int64 `json:"-" url:"run_id"`
	// Which views to export (CODE, DASHBOARDS, or ALL). Defaults to CODE.
	ViewsToExport ViewsToExport `json:"-" url:"views_to_export,omitempty"`
}

func exportRunRequestFromPb(pb *exportRunRequestPb) (*ExportRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportRunRequest{}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}
	viewsToExportField, err := identity(&pb.ViewsToExport)
	if err != nil {
		return nil, err
	}
	if viewsToExportField != nil {
		st.ViewsToExport = *viewsToExportField
	}

	return st, nil
}

type FileArrivalTriggerConfiguration struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	MinTimeBetweenTriggersSeconds int
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	Url string
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	WaitAfterLastChangeSeconds int

	ForceSendFields []string
}

func fileArrivalTriggerConfigurationToPb(st *FileArrivalTriggerConfiguration) (*fileArrivalTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileArrivalTriggerConfigurationPb{}
	minTimeBetweenTriggersSecondsPb, err := identity(&st.MinTimeBetweenTriggersSeconds)
	if err != nil {
		return nil, err
	}
	if minTimeBetweenTriggersSecondsPb != nil {
		pb.MinTimeBetweenTriggersSeconds = *minTimeBetweenTriggersSecondsPb
	}

	urlPb, err := identity(&st.Url)
	if err != nil {
		return nil, err
	}
	if urlPb != nil {
		pb.Url = *urlPb
	}

	waitAfterLastChangeSecondsPb, err := identity(&st.WaitAfterLastChangeSeconds)
	if err != nil {
		return nil, err
	}
	if waitAfterLastChangeSecondsPb != nil {
		pb.WaitAfterLastChangeSeconds = *waitAfterLastChangeSecondsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FileArrivalTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileArrivalTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileArrivalTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileArrivalTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := fileArrivalTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type fileArrivalTriggerConfigurationPb struct {
	// If set, the trigger starts a run only after the specified amount of time
	// passed since the last time the trigger fired. The minimum allowed value
	// is 60 seconds
	MinTimeBetweenTriggersSeconds int `json:"min_time_between_triggers_seconds,omitempty"`
	// URL to be monitored for file arrivals. The path must point to the root or
	// a subpath of the external location.
	Url string `json:"url"`
	// If set, the trigger starts a run only after no file activity has occurred
	// for the specified amount of time. This makes it possible to wait for a
	// batch of incoming files to arrive before triggering a run. The minimum
	// allowed value is 60 seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileArrivalTriggerConfigurationFromPb(pb *fileArrivalTriggerConfigurationPb) (*FileArrivalTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileArrivalTriggerConfiguration{}
	minTimeBetweenTriggersSecondsField, err := identity(&pb.MinTimeBetweenTriggersSeconds)
	if err != nil {
		return nil, err
	}
	if minTimeBetweenTriggersSecondsField != nil {
		st.MinTimeBetweenTriggersSeconds = *minTimeBetweenTriggersSecondsField
	}
	urlField, err := identity(&pb.Url)
	if err != nil {
		return nil, err
	}
	if urlField != nil {
		st.Url = *urlField
	}
	waitAfterLastChangeSecondsField, err := identity(&pb.WaitAfterLastChangeSeconds)
	if err != nil {
		return nil, err
	}
	if waitAfterLastChangeSecondsField != nil {
		st.WaitAfterLastChangeSeconds = *waitAfterLastChangeSecondsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileArrivalTriggerConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileArrivalTriggerConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForEachStats struct {
	// Sample of 3 most common error messages occurred during the iteration.
	ErrorMessageStats []ForEachTaskErrorMessageStats
	// Describes stats of the iteration. Only latest retries are considered.
	TaskRunStats *ForEachTaskTaskRunStats
}

func forEachStatsToPb(st *ForEachStats) (*forEachStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachStatsPb{}

	var errorMessageStatsPb []forEachTaskErrorMessageStatsPb
	for _, item := range st.ErrorMessageStats {
		itemPb, err := forEachTaskErrorMessageStatsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			errorMessageStatsPb = append(errorMessageStatsPb, *itemPb)
		}
	}
	pb.ErrorMessageStats = errorMessageStatsPb

	taskRunStatsPb, err := forEachTaskTaskRunStatsToPb(st.TaskRunStats)
	if err != nil {
		return nil, err
	}
	if taskRunStatsPb != nil {
		pb.TaskRunStats = taskRunStatsPb
	}

	return pb, nil
}

func (st *ForEachStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachStats) MarshalJSON() ([]byte, error) {
	pb, err := forEachStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type forEachStatsPb struct {
	// Sample of 3 most common error messages occurred during the iteration.
	ErrorMessageStats []forEachTaskErrorMessageStatsPb `json:"error_message_stats,omitempty"`
	// Describes stats of the iteration. Only latest retries are considered.
	TaskRunStats *forEachTaskTaskRunStatsPb `json:"task_run_stats,omitempty"`
}

func forEachStatsFromPb(pb *forEachStatsPb) (*ForEachStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachStats{}

	var errorMessageStatsField []ForEachTaskErrorMessageStats
	for _, itemPb := range pb.ErrorMessageStats {
		item, err := forEachTaskErrorMessageStatsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			errorMessageStatsField = append(errorMessageStatsField, *item)
		}
	}
	st.ErrorMessageStats = errorMessageStatsField
	taskRunStatsField, err := forEachTaskTaskRunStatsFromPb(pb.TaskRunStats)
	if err != nil {
		return nil, err
	}
	if taskRunStatsField != nil {
		st.TaskRunStats = taskRunStatsField
	}

	return st, nil
}

type ForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency int
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string
	// Configuration for the task that will be run for each element in the array
	Task Task

	ForceSendFields []string
}

func forEachTaskToPb(st *ForEachTask) (*forEachTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachTaskPb{}
	concurrencyPb, err := identity(&st.Concurrency)
	if err != nil {
		return nil, err
	}
	if concurrencyPb != nil {
		pb.Concurrency = *concurrencyPb
	}

	inputsPb, err := identity(&st.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsPb != nil {
		pb.Inputs = *inputsPb
	}

	taskPb, err := taskToPb(&st.Task)
	if err != nil {
		return nil, err
	}
	if taskPb != nil {
		pb.Task = *taskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ForEachTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachTask) MarshalJSON() ([]byte, error) {
	pb, err := forEachTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type forEachTaskPb struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency int `json:"concurrency,omitempty"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string `json:"inputs"`
	// Configuration for the task that will be run for each element in the array
	Task taskPb `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forEachTaskFromPb(pb *forEachTaskPb) (*ForEachTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTask{}
	concurrencyField, err := identity(&pb.Concurrency)
	if err != nil {
		return nil, err
	}
	if concurrencyField != nil {
		st.Concurrency = *concurrencyField
	}
	inputsField, err := identity(&pb.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsField != nil {
		st.Inputs = *inputsField
	}
	taskField, err := taskFromPb(&pb.Task)
	if err != nil {
		return nil, err
	}
	if taskField != nil {
		st.Task = *taskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forEachTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forEachTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForEachTaskErrorMessageStats struct {
	// Describes the count of such error message encountered during the
	// iterations.
	Count int
	// Describes the error message occured during the iterations.
	ErrorMessage string
	// Describes the termination reason for the error message.
	TerminationCategory string

	ForceSendFields []string
}

func forEachTaskErrorMessageStatsToPb(st *ForEachTaskErrorMessageStats) (*forEachTaskErrorMessageStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachTaskErrorMessageStatsPb{}
	countPb, err := identity(&st.Count)
	if err != nil {
		return nil, err
	}
	if countPb != nil {
		pb.Count = *countPb
	}

	errorMessagePb, err := identity(&st.ErrorMessage)
	if err != nil {
		return nil, err
	}
	if errorMessagePb != nil {
		pb.ErrorMessage = *errorMessagePb
	}

	terminationCategoryPb, err := identity(&st.TerminationCategory)
	if err != nil {
		return nil, err
	}
	if terminationCategoryPb != nil {
		pb.TerminationCategory = *terminationCategoryPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ForEachTaskErrorMessageStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachTaskErrorMessageStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachTaskErrorMessageStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachTaskErrorMessageStats) MarshalJSON() ([]byte, error) {
	pb, err := forEachTaskErrorMessageStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type forEachTaskErrorMessageStatsPb struct {
	// Describes the count of such error message encountered during the
	// iterations.
	Count int `json:"count,omitempty"`
	// Describes the error message occured during the iterations.
	ErrorMessage string `json:"error_message,omitempty"`
	// Describes the termination reason for the error message.
	TerminationCategory string `json:"termination_category,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forEachTaskErrorMessageStatsFromPb(pb *forEachTaskErrorMessageStatsPb) (*ForEachTaskErrorMessageStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTaskErrorMessageStats{}
	countField, err := identity(&pb.Count)
	if err != nil {
		return nil, err
	}
	if countField != nil {
		st.Count = *countField
	}
	errorMessageField, err := identity(&pb.ErrorMessage)
	if err != nil {
		return nil, err
	}
	if errorMessageField != nil {
		st.ErrorMessage = *errorMessageField
	}
	terminationCategoryField, err := identity(&pb.TerminationCategory)
	if err != nil {
		return nil, err
	}
	if terminationCategoryField != nil {
		st.TerminationCategory = *terminationCategoryField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forEachTaskErrorMessageStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forEachTaskErrorMessageStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForEachTaskTaskRunStats struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	ActiveIterations int
	// Describes the number of failed and succeeded iteration runs.
	CompletedIterations int
	// Describes the number of failed iteration runs.
	FailedIterations int
	// Describes the number of iteration runs that have been scheduled.
	ScheduledIterations int
	// Describes the number of succeeded iteration runs.
	SucceededIterations int
	// Describes the length of the list of items to iterate over.
	TotalIterations int

	ForceSendFields []string
}

func forEachTaskTaskRunStatsToPb(st *ForEachTaskTaskRunStats) (*forEachTaskTaskRunStatsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forEachTaskTaskRunStatsPb{}
	activeIterationsPb, err := identity(&st.ActiveIterations)
	if err != nil {
		return nil, err
	}
	if activeIterationsPb != nil {
		pb.ActiveIterations = *activeIterationsPb
	}

	completedIterationsPb, err := identity(&st.CompletedIterations)
	if err != nil {
		return nil, err
	}
	if completedIterationsPb != nil {
		pb.CompletedIterations = *completedIterationsPb
	}

	failedIterationsPb, err := identity(&st.FailedIterations)
	if err != nil {
		return nil, err
	}
	if failedIterationsPb != nil {
		pb.FailedIterations = *failedIterationsPb
	}

	scheduledIterationsPb, err := identity(&st.ScheduledIterations)
	if err != nil {
		return nil, err
	}
	if scheduledIterationsPb != nil {
		pb.ScheduledIterations = *scheduledIterationsPb
	}

	succeededIterationsPb, err := identity(&st.SucceededIterations)
	if err != nil {
		return nil, err
	}
	if succeededIterationsPb != nil {
		pb.SucceededIterations = *succeededIterationsPb
	}

	totalIterationsPb, err := identity(&st.TotalIterations)
	if err != nil {
		return nil, err
	}
	if totalIterationsPb != nil {
		pb.TotalIterations = *totalIterationsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ForEachTaskTaskRunStats) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forEachTaskTaskRunStatsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forEachTaskTaskRunStatsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForEachTaskTaskRunStats) MarshalJSON() ([]byte, error) {
	pb, err := forEachTaskTaskRunStatsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type forEachTaskTaskRunStatsPb struct {
	// Describes the iteration runs having an active lifecycle state or an
	// active run sub state.
	ActiveIterations int `json:"active_iterations,omitempty"`
	// Describes the number of failed and succeeded iteration runs.
	CompletedIterations int `json:"completed_iterations,omitempty"`
	// Describes the number of failed iteration runs.
	FailedIterations int `json:"failed_iterations,omitempty"`
	// Describes the number of iteration runs that have been scheduled.
	ScheduledIterations int `json:"scheduled_iterations,omitempty"`
	// Describes the number of succeeded iteration runs.
	SucceededIterations int `json:"succeeded_iterations,omitempty"`
	// Describes the length of the list of items to iterate over.
	TotalIterations int `json:"total_iterations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forEachTaskTaskRunStatsFromPb(pb *forEachTaskTaskRunStatsPb) (*ForEachTaskTaskRunStats, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForEachTaskTaskRunStats{}
	activeIterationsField, err := identity(&pb.ActiveIterations)
	if err != nil {
		return nil, err
	}
	if activeIterationsField != nil {
		st.ActiveIterations = *activeIterationsField
	}
	completedIterationsField, err := identity(&pb.CompletedIterations)
	if err != nil {
		return nil, err
	}
	if completedIterationsField != nil {
		st.CompletedIterations = *completedIterationsField
	}
	failedIterationsField, err := identity(&pb.FailedIterations)
	if err != nil {
		return nil, err
	}
	if failedIterationsField != nil {
		st.FailedIterations = *failedIterationsField
	}
	scheduledIterationsField, err := identity(&pb.ScheduledIterations)
	if err != nil {
		return nil, err
	}
	if scheduledIterationsField != nil {
		st.ScheduledIterations = *scheduledIterationsField
	}
	succeededIterationsField, err := identity(&pb.SucceededIterations)
	if err != nil {
		return nil, err
	}
	if succeededIterationsField != nil {
		st.SucceededIterations = *succeededIterationsField
	}
	totalIterationsField, err := identity(&pb.TotalIterations)
	if err != nil {
		return nil, err
	}
	if totalIterationsField != nil {
		st.TotalIterations = *totalIterationsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forEachTaskTaskRunStatsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forEachTaskTaskRunStatsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Format string
type formatPb string

const FormatMultiTask Format = `MULTI_TASK`

const FormatSingleTask Format = `SINGLE_TASK`

// String representation for [fmt.Print]
func (f *Format) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Format) Set(v string) error {
	switch v {
	case `MULTI_TASK`, `SINGLE_TASK`:
		*f = Format(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MULTI_TASK", "SINGLE_TASK"`, v)
	}
}

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

func formatToPb(st *Format) (*formatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := formatPb(*st)
	return &pb, nil
}

func formatFromPb(pb *formatPb) (*Format, error) {
	if pb == nil {
		return nil, nil
	}
	st := Format(*pb)
	return &st, nil
}

type GenAiComputeTask struct {
	// Command launcher to run the actual script, e.g. bash, python etc.
	Command string

	Compute *ComputeConfig
	// Runtime image
	DlRuntimeImage string
	// Optional string containing the name of the MLflow experiment to log the
	// run to. If name is not found, backend will create the mlflow experiment
	// using the name.
	MlflowExperimentName string
	// Optional location type of the training script. When set to `WORKSPACE`,
	// the script will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the script will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`: Script
	// is located in Databricks workspace. * `GIT`: Script is located in cloud
	// Git provider.
	Source Source
	// The training script file path to be executed. Cloud file URIs (such as
	// dbfs:/, s3:/, adls:/, gcs:/) and workspace paths are supported. For
	// python files stored in the Databricks workspace, the path must be
	// absolute and begin with `/`. For files stored in a remote repository, the
	// path must be relative. This field is required.
	TrainingScriptPath string
	// Optional string containing model parameters passed to the training script
	// in yaml format. If present, then the content in yaml_parameters_file_path
	// will be ignored.
	YamlParameters string
	// Optional path to a YAML file containing model parameters passed to the
	// training script.
	YamlParametersFilePath string

	ForceSendFields []string
}

func genAiComputeTaskToPb(st *GenAiComputeTask) (*genAiComputeTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genAiComputeTaskPb{}
	commandPb, err := identity(&st.Command)
	if err != nil {
		return nil, err
	}
	if commandPb != nil {
		pb.Command = *commandPb
	}

	computePb, err := computeConfigToPb(st.Compute)
	if err != nil {
		return nil, err
	}
	if computePb != nil {
		pb.Compute = computePb
	}

	dlRuntimeImagePb, err := identity(&st.DlRuntimeImage)
	if err != nil {
		return nil, err
	}
	if dlRuntimeImagePb != nil {
		pb.DlRuntimeImage = *dlRuntimeImagePb
	}

	mlflowExperimentNamePb, err := identity(&st.MlflowExperimentName)
	if err != nil {
		return nil, err
	}
	if mlflowExperimentNamePb != nil {
		pb.MlflowExperimentName = *mlflowExperimentNamePb
	}

	sourcePb, err := identity(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	trainingScriptPathPb, err := identity(&st.TrainingScriptPath)
	if err != nil {
		return nil, err
	}
	if trainingScriptPathPb != nil {
		pb.TrainingScriptPath = *trainingScriptPathPb
	}

	yamlParametersPb, err := identity(&st.YamlParameters)
	if err != nil {
		return nil, err
	}
	if yamlParametersPb != nil {
		pb.YamlParameters = *yamlParametersPb
	}

	yamlParametersFilePathPb, err := identity(&st.YamlParametersFilePath)
	if err != nil {
		return nil, err
	}
	if yamlParametersFilePathPb != nil {
		pb.YamlParametersFilePath = *yamlParametersFilePathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GenAiComputeTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genAiComputeTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genAiComputeTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenAiComputeTask) MarshalJSON() ([]byte, error) {
	pb, err := genAiComputeTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type genAiComputeTaskPb struct {
	// Command launcher to run the actual script, e.g. bash, python etc.
	Command string `json:"command,omitempty"`

	Compute *computeConfigPb `json:"compute,omitempty"`
	// Runtime image
	DlRuntimeImage string `json:"dl_runtime_image"`
	// Optional string containing the name of the MLflow experiment to log the
	// run to. If name is not found, backend will create the mlflow experiment
	// using the name.
	MlflowExperimentName string `json:"mlflow_experiment_name,omitempty"`
	// Optional location type of the training script. When set to `WORKSPACE`,
	// the script will be retrieved from the local Databricks workspace. When
	// set to `GIT`, the script will be retrieved from a Git repository defined
	// in `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`: Script
	// is located in Databricks workspace. * `GIT`: Script is located in cloud
	// Git provider.
	Source Source `json:"source,omitempty"`
	// The training script file path to be executed. Cloud file URIs (such as
	// dbfs:/, s3:/, adls:/, gcs:/) and workspace paths are supported. For
	// python files stored in the Databricks workspace, the path must be
	// absolute and begin with `/`. For files stored in a remote repository, the
	// path must be relative. This field is required.
	TrainingScriptPath string `json:"training_script_path,omitempty"`
	// Optional string containing model parameters passed to the training script
	// in yaml format. If present, then the content in yaml_parameters_file_path
	// will be ignored.
	YamlParameters string `json:"yaml_parameters,omitempty"`
	// Optional path to a YAML file containing model parameters passed to the
	// training script.
	YamlParametersFilePath string `json:"yaml_parameters_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genAiComputeTaskFromPb(pb *genAiComputeTaskPb) (*GenAiComputeTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenAiComputeTask{}
	commandField, err := identity(&pb.Command)
	if err != nil {
		return nil, err
	}
	if commandField != nil {
		st.Command = *commandField
	}
	computeField, err := computeConfigFromPb(pb.Compute)
	if err != nil {
		return nil, err
	}
	if computeField != nil {
		st.Compute = computeField
	}
	dlRuntimeImageField, err := identity(&pb.DlRuntimeImage)
	if err != nil {
		return nil, err
	}
	if dlRuntimeImageField != nil {
		st.DlRuntimeImage = *dlRuntimeImageField
	}
	mlflowExperimentNameField, err := identity(&pb.MlflowExperimentName)
	if err != nil {
		return nil, err
	}
	if mlflowExperimentNameField != nil {
		st.MlflowExperimentName = *mlflowExperimentNameField
	}
	sourceField, err := identity(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	trainingScriptPathField, err := identity(&pb.TrainingScriptPath)
	if err != nil {
		return nil, err
	}
	if trainingScriptPathField != nil {
		st.TrainingScriptPath = *trainingScriptPathField
	}
	yamlParametersField, err := identity(&pb.YamlParameters)
	if err != nil {
		return nil, err
	}
	if yamlParametersField != nil {
		st.YamlParameters = *yamlParametersField
	}
	yamlParametersFilePathField, err := identity(&pb.YamlParametersFilePath)
	if err != nil {
		return nil, err
	}
	if yamlParametersFilePathField != nil {
		st.YamlParametersFilePath = *yamlParametersFilePathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genAiComputeTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genAiComputeTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get job permission levels
type GetJobPermissionLevelsRequest struct {
	// The job for which to get or manage permissions.
	JobId string
}

func getJobPermissionLevelsRequestToPb(st *GetJobPermissionLevelsRequest) (*getJobPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobPermissionLevelsRequestPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	return pb, nil
}

func (st *GetJobPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getJobPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getJobPermissionLevelsRequestPb struct {
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

func getJobPermissionLevelsRequestFromPb(pb *getJobPermissionLevelsRequestPb) (*GetJobPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionLevelsRequest{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	return st, nil
}

type GetJobPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []JobPermissionsDescription
}

func getJobPermissionLevelsResponseToPb(st *GetJobPermissionLevelsResponse) (*getJobPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobPermissionLevelsResponsePb{}

	var permissionLevelsPb []jobPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := jobPermissionsDescriptionToPb(&item)
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

func (st *GetJobPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getJobPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getJobPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []jobPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getJobPermissionLevelsResponseFromPb(pb *getJobPermissionLevelsResponsePb) (*GetJobPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionLevelsResponse{}

	var permissionLevelsField []JobPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := jobPermissionsDescriptionFromPb(&itemPb)
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

// Get job permissions
type GetJobPermissionsRequest struct {
	// The job for which to get or manage permissions.
	JobId string
}

func getJobPermissionsRequestToPb(st *GetJobPermissionsRequest) (*getJobPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobPermissionsRequestPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	return pb, nil
}

func (st *GetJobPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getJobPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getJobPermissionsRequestPb struct {
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

func getJobPermissionsRequestFromPb(pb *getJobPermissionsRequestPb) (*GetJobPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobPermissionsRequest{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	return st, nil
}

// Get a single job
type GetJobRequest struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId int64
	// Use `next_page_token` returned from the previous GetJob response to
	// request the next page of the job's array properties.
	PageToken string

	ForceSendFields []string
}

func getJobRequestToPb(st *GetJobRequest) (*getJobRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getJobRequestPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
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

func (st *GetJobRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getJobRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getJobRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetJobRequest) MarshalJSON() ([]byte, error) {
	pb, err := getJobRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getJobRequestPb struct {
	// The canonical identifier of the job to retrieve information about. This
	// field is required.
	JobId int64 `json:"-" url:"job_id"`
	// Use `next_page_token` returned from the previous GetJob response to
	// request the next page of the job's array properties.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getJobRequestFromPb(pb *getJobRequestPb) (*GetJobRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetJobRequest{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
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

func (st *getJobRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getJobRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get job policy compliance
type GetPolicyComplianceRequest struct {
	// The ID of the job whose compliance status you are requesting.
	JobId int64
}

func getPolicyComplianceRequestToPb(st *GetPolicyComplianceRequest) (*getPolicyComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPolicyComplianceRequestPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	return pb, nil
}

func (st *GetPolicyComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPolicyComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPolicyComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPolicyComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPolicyComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getPolicyComplianceRequestPb struct {
	// The ID of the job whose compliance status you are requesting.
	JobId int64 `json:"-" url:"job_id"`
}

func getPolicyComplianceRequestFromPb(pb *getPolicyComplianceRequestPb) (*GetPolicyComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyComplianceRequest{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	return st, nil
}

type GetPolicyComplianceResponse struct {
	// Whether the job is compliant with its policies or not. Jobs could be out
	// of compliance if a policy they are using was updated after the job was
	// last edited and some of its job clusters no longer comply with their
	// updated policies.
	IsCompliant bool
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations map[string]string

	ForceSendFields []string
}

func getPolicyComplianceResponseToPb(st *GetPolicyComplianceResponse) (*getPolicyComplianceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPolicyComplianceResponsePb{}
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

func (st *GetPolicyComplianceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPolicyComplianceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPolicyComplianceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPolicyComplianceResponse) MarshalJSON() ([]byte, error) {
	pb, err := getPolicyComplianceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getPolicyComplianceResponsePb struct {
	// Whether the job is compliant with its policies or not. Jobs could be out
	// of compliance if a policy they are using was updated after the job was
	// last edited and some of its job clusters no longer comply with their
	// updated policies.
	IsCompliant bool `json:"is_compliant,omitempty"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPolicyComplianceResponseFromPb(pb *getPolicyComplianceResponsePb) (*GetPolicyComplianceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPolicyComplianceResponse{}
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

func (st *getPolicyComplianceResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPolicyComplianceResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get the output for a single run
type GetRunOutputRequest struct {
	// The canonical identifier for the run.
	RunId int64
}

func getRunOutputRequestToPb(st *GetRunOutputRequest) (*getRunOutputRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunOutputRequestPb{}
	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	return pb, nil
}

func (st *GetRunOutputRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunOutputRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunOutputRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunOutputRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRunOutputRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRunOutputRequestPb struct {
	// The canonical identifier for the run.
	RunId int64 `json:"-" url:"run_id"`
}

func getRunOutputRequestFromPb(pb *getRunOutputRequestPb) (*GetRunOutputRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunOutputRequest{}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	return st, nil
}

// Get a single job run
type GetRunRequest struct {
	// Whether to include the repair history in the response.
	IncludeHistory bool
	// Whether to include resolved parameter values in the response.
	IncludeResolvedValues bool
	// Use `next_page_token` returned from the previous GetRun response to
	// request the next page of the run's array properties.
	PageToken string
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId int64

	ForceSendFields []string
}

func getRunRequestToPb(st *GetRunRequest) (*getRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunRequestPb{}
	includeHistoryPb, err := identity(&st.IncludeHistory)
	if err != nil {
		return nil, err
	}
	if includeHistoryPb != nil {
		pb.IncludeHistory = *includeHistoryPb
	}

	includeResolvedValuesPb, err := identity(&st.IncludeResolvedValues)
	if err != nil {
		return nil, err
	}
	if includeResolvedValuesPb != nil {
		pb.IncludeResolvedValues = *includeResolvedValuesPb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRunRequestPb struct {
	// Whether to include the repair history in the response.
	IncludeHistory bool `json:"-" url:"include_history,omitempty"`
	// Whether to include resolved parameter values in the response.
	IncludeResolvedValues bool `json:"-" url:"include_resolved_values,omitempty"`
	// Use `next_page_token` returned from the previous GetRun response to
	// request the next page of the run's array properties.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The canonical identifier of the run for which to retrieve the metadata.
	// This field is required.
	RunId int64 `json:"-" url:"run_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRunRequestFromPb(pb *getRunRequestPb) (*GetRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunRequest{}
	includeHistoryField, err := identity(&pb.IncludeHistory)
	if err != nil {
		return nil, err
	}
	if includeHistoryField != nil {
		st.IncludeHistory = *includeHistoryField
	}
	includeResolvedValuesField, err := identity(&pb.IncludeResolvedValues)
	if err != nil {
		return nil, err
	}
	if includeResolvedValuesField != nil {
		st.IncludeResolvedValues = *includeResolvedValuesField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRunRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRunRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GitProvider string
type gitProviderPb string

const GitProviderAwsCodeCommit GitProvider = `awsCodeCommit`

const GitProviderAzureDevOpsServices GitProvider = `azureDevOpsServices`

const GitProviderBitbucketCloud GitProvider = `bitbucketCloud`

const GitProviderBitbucketServer GitProvider = `bitbucketServer`

const GitProviderGitHub GitProvider = `gitHub`

const GitProviderGitHubEnterprise GitProvider = `gitHubEnterprise`

const GitProviderGitLab GitProvider = `gitLab`

const GitProviderGitLabEnterpriseEdition GitProvider = `gitLabEnterpriseEdition`

// String representation for [fmt.Print]
func (f *GitProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GitProvider) Set(v string) error {
	switch v {
	case `awsCodeCommit`, `azureDevOpsServices`, `bitbucketCloud`, `bitbucketServer`, `gitHub`, `gitHubEnterprise`, `gitLab`, `gitLabEnterpriseEdition`:
		*f = GitProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "awsCodeCommit", "azureDevOpsServices", "bitbucketCloud", "bitbucketServer", "gitHub", "gitHubEnterprise", "gitLab", "gitLabEnterpriseEdition"`, v)
	}
}

// Type always returns GitProvider to satisfy [pflag.Value] interface
func (f *GitProvider) Type() string {
	return "GitProvider"
}

func gitProviderToPb(st *GitProvider) (*gitProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := gitProviderPb(*st)
	return &pb, nil
}

func gitProviderFromPb(pb *gitProviderPb) (*GitProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := GitProvider(*pb)
	return &st, nil
}

// Read-only state of the remote repository at the time the job was run. This
// field is only included on job runs.
type GitSnapshot struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	UsedCommit string

	ForceSendFields []string
}

func gitSnapshotToPb(st *GitSnapshot) (*gitSnapshotPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gitSnapshotPb{}
	usedCommitPb, err := identity(&st.UsedCommit)
	if err != nil {
		return nil, err
	}
	if usedCommitPb != nil {
		pb.UsedCommit = *usedCommitPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GitSnapshot) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gitSnapshotPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gitSnapshotFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GitSnapshot) MarshalJSON() ([]byte, error) {
	pb, err := gitSnapshotToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type gitSnapshotPb struct {
	// Commit that was used to execute the run. If git_branch was specified,
	// this points to the HEAD of the branch at the time of the run; if git_tag
	// was specified, this points to the commit the tag points to.
	UsedCommit string `json:"used_commit,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gitSnapshotFromPb(pb *gitSnapshotPb) (*GitSnapshot, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GitSnapshot{}
	usedCommitField, err := identity(&pb.UsedCommit)
	if err != nil {
		return nil, err
	}
	if usedCommitField != nil {
		st.UsedCommit = *usedCommitField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gitSnapshotPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gitSnapshotPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// An optional specification for a remote Git repository containing the source
// code used by tasks. Version-controlled source code is supported by notebook,
// dbt, Python script, and SQL File tasks.
//
// If `git_source` is set, these tasks retrieve the file from the remote
// repository by default. However, this behavior can be overridden by setting
// `source` to `WORKSPACE` on the task.
//
// Note: dbt and SQL File tasks support only version-controlled sources. If dbt
// or SQL File tasks are used, `git_source` must be defined on the job.
type GitSource struct {
	// Name of the branch to be checked out and used by this job. This field
	// cannot be specified in conjunction with git_tag or git_commit.
	GitBranch string
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	GitCommit string
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider GitProvider
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	GitSnapshot *GitSnapshot
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	GitTag string
	// URL of the repository to be cloned by this job.
	GitUrl string
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	JobSource *JobSource

	ForceSendFields []string
}

func gitSourceToPb(st *GitSource) (*gitSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &gitSourcePb{}
	gitBranchPb, err := identity(&st.GitBranch)
	if err != nil {
		return nil, err
	}
	if gitBranchPb != nil {
		pb.GitBranch = *gitBranchPb
	}

	gitCommitPb, err := identity(&st.GitCommit)
	if err != nil {
		return nil, err
	}
	if gitCommitPb != nil {
		pb.GitCommit = *gitCommitPb
	}

	gitProviderPb, err := identity(&st.GitProvider)
	if err != nil {
		return nil, err
	}
	if gitProviderPb != nil {
		pb.GitProvider = *gitProviderPb
	}

	gitSnapshotPb, err := gitSnapshotToPb(st.GitSnapshot)
	if err != nil {
		return nil, err
	}
	if gitSnapshotPb != nil {
		pb.GitSnapshot = gitSnapshotPb
	}

	gitTagPb, err := identity(&st.GitTag)
	if err != nil {
		return nil, err
	}
	if gitTagPb != nil {
		pb.GitTag = *gitTagPb
	}

	gitUrlPb, err := identity(&st.GitUrl)
	if err != nil {
		return nil, err
	}
	if gitUrlPb != nil {
		pb.GitUrl = *gitUrlPb
	}

	jobSourcePb, err := jobSourceToPb(st.JobSource)
	if err != nil {
		return nil, err
	}
	if jobSourcePb != nil {
		pb.JobSource = jobSourcePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GitSource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &gitSourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := gitSourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GitSource) MarshalJSON() ([]byte, error) {
	pb, err := gitSourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type gitSourcePb struct {
	// Name of the branch to be checked out and used by this job. This field
	// cannot be specified in conjunction with git_tag or git_commit.
	GitBranch string `json:"git_branch,omitempty"`
	// Commit to be checked out and used by this job. This field cannot be
	// specified in conjunction with git_branch or git_tag.
	GitCommit string `json:"git_commit,omitempty"`
	// Unique identifier of the service used to host the Git repository. The
	// value is case insensitive.
	GitProvider GitProvider `json:"git_provider"`
	// Read-only state of the remote repository at the time the job was run.
	// This field is only included on job runs.
	GitSnapshot *gitSnapshotPb `json:"git_snapshot,omitempty"`
	// Name of the tag to be checked out and used by this job. This field cannot
	// be specified in conjunction with git_branch or git_commit.
	GitTag string `json:"git_tag,omitempty"`
	// URL of the repository to be cloned by this job.
	GitUrl string `json:"git_url"`
	// The source of the job specification in the remote repository when the job
	// is source controlled.
	JobSource *jobSourcePb `json:"job_source,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func gitSourceFromPb(pb *gitSourcePb) (*GitSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GitSource{}
	gitBranchField, err := identity(&pb.GitBranch)
	if err != nil {
		return nil, err
	}
	if gitBranchField != nil {
		st.GitBranch = *gitBranchField
	}
	gitCommitField, err := identity(&pb.GitCommit)
	if err != nil {
		return nil, err
	}
	if gitCommitField != nil {
		st.GitCommit = *gitCommitField
	}
	gitProviderField, err := identity(&pb.GitProvider)
	if err != nil {
		return nil, err
	}
	if gitProviderField != nil {
		st.GitProvider = *gitProviderField
	}
	gitSnapshotField, err := gitSnapshotFromPb(pb.GitSnapshot)
	if err != nil {
		return nil, err
	}
	if gitSnapshotField != nil {
		st.GitSnapshot = gitSnapshotField
	}
	gitTagField, err := identity(&pb.GitTag)
	if err != nil {
		return nil, err
	}
	if gitTagField != nil {
		st.GitTag = *gitTagField
	}
	gitUrlField, err := identity(&pb.GitUrl)
	if err != nil {
		return nil, err
	}
	if gitUrlField != nil {
		st.GitUrl = *gitUrlField
	}
	jobSourceField, err := jobSourceFromPb(pb.JobSource)
	if err != nil {
		return nil, err
	}
	if jobSourceField != nil {
		st.JobSource = jobSourceField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *gitSourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st gitSourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Job was retrieved successfully.
type Job struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId string
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	HasMore bool
	// The canonical identifier for this job.
	JobId int64
	// A token that can be used to list the next page of array properties.
	NextPageToken string
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	RunAsUserName string
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *JobSettings

	ForceSendFields []string
}

func jobToPb(st *Job) (*jobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPb{}
	createdTimePb, err := identity(&st.CreatedTime)
	if err != nil {
		return nil, err
	}
	if createdTimePb != nil {
		pb.CreatedTime = *createdTimePb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	effectiveBudgetPolicyIdPb, err := identity(&st.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdPb != nil {
		pb.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdPb
	}

	hasMorePb, err := identity(&st.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMorePb != nil {
		pb.HasMore = *hasMorePb
	}

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	runAsUserNamePb, err := identity(&st.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNamePb != nil {
		pb.RunAsUserName = *runAsUserNamePb
	}

	settingsPb, err := jobSettingsToPb(st.Settings)
	if err != nil {
		return nil, err
	}
	if settingsPb != nil {
		pb.Settings = settingsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Job) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Job) MarshalJSON() ([]byte, error) {
	pb, err := jobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobPb struct {
	// The time at which this job was created in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC).
	CreatedTime int64 `json:"created_time,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The id of the budget policy used by this job for cost attribution
	// purposes. This may be set through (in order of precedence): 1. Budget
	// admins through the account or workspace console 2. Jobs UI in the job
	// details page and Jobs API using `budget_policy_id` 3. Inferred default
	// based on accessible budget policies of the run_as identity on job
	// creation or modification.
	EffectiveBudgetPolicyId string `json:"effective_budget_policy_id,omitempty"`
	// Indicates if the job has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/get endpoint.
	// It is only relevant for API 2.2 :method:jobs/list requests with
	// `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// The canonical identifier for this job.
	JobId int64 `json:"job_id,omitempty"`
	// A token that can be used to list the next page of array properties.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The email of an active workspace user or the application ID of a service
	// principal that the job runs as. This value can be changed by setting the
	// `run_as` field when creating or updating a job.
	//
	// By default, `run_as_user_name` is based on the current job settings and
	// is set to the creator of the job if job access control is disabled or to
	// the user with the `is_owner` permission if job access control is enabled.
	RunAsUserName string `json:"run_as_user_name,omitempty"`
	// Settings for this job and all of its runs. These settings can be updated
	// using the `resetJob` method.
	Settings *jobSettingsPb `json:"settings,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobFromPb(pb *jobPb) (*Job, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Job{}
	createdTimeField, err := identity(&pb.CreatedTime)
	if err != nil {
		return nil, err
	}
	if createdTimeField != nil {
		st.CreatedTime = *createdTimeField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}
	effectiveBudgetPolicyIdField, err := identity(&pb.EffectiveBudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if effectiveBudgetPolicyIdField != nil {
		st.EffectiveBudgetPolicyId = *effectiveBudgetPolicyIdField
	}
	hasMoreField, err := identity(&pb.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMoreField != nil {
		st.HasMore = *hasMoreField
	}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
	runAsUserNameField, err := identity(&pb.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNameField != nil {
		st.RunAsUserName = *runAsUserNameField
	}
	settingsField, err := jobSettingsFromPb(pb.Settings)
	if err != nil {
		return nil, err
	}
	if settingsField != nil {
		st.Settings = settingsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobAccessControlRequest struct {
	// name of the group
	GroupName string
	// Permission level
	PermissionLevel JobPermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func jobAccessControlRequestToPb(st *JobAccessControlRequest) (*jobAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobAccessControlRequestPb{}
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

func (st *JobAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := jobAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobAccessControlRequestFromPb(pb *jobAccessControlRequestPb) (*JobAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobAccessControlRequest{}
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

func (st *jobAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobAccessControlResponse struct {
	// All permissions.
	AllPermissions []JobPermission
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

func jobAccessControlResponseToPb(st *JobAccessControlResponse) (*jobAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobAccessControlResponsePb{}

	var allPermissionsPb []jobPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := jobPermissionToPb(&item)
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

func (st *JobAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := jobAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []jobPermissionPb `json:"all_permissions,omitempty"`
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

func jobAccessControlResponseFromPb(pb *jobAccessControlResponsePb) (*JobAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobAccessControlResponse{}

	var allPermissionsField []JobPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := jobPermissionFromPb(&itemPb)
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

func (st *jobAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobCluster struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey string
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster compute.ClusterSpec
}

func jobClusterToPb(st *JobCluster) (*jobClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobClusterPb{}
	jobClusterKeyPb, err := identity(&st.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyPb != nil {
		pb.JobClusterKey = *jobClusterKeyPb
	}

	newClusterPb, err := compute.ClusterSpecToPb(&st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = *newClusterPb
	}

	return pb, nil
}

func (st *JobCluster) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobClusterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobClusterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobCluster) MarshalJSON() ([]byte, error) {
	pb, err := jobClusterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobClusterPb struct {
	// A unique name for the job cluster. This field is required and must be
	// unique within the job. `JobTaskSettings` may refer to this field to
	// determine which cluster to launch for the task execution.
	JobClusterKey string `json:"job_cluster_key"`
	// If new_cluster, a description of a cluster that is created for each task.
	NewCluster compute.ClusterSpecPb `json:"new_cluster"`
}

func jobClusterFromPb(pb *jobClusterPb) (*JobCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobCluster{}
	jobClusterKeyField, err := identity(&pb.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyField != nil {
		st.JobClusterKey = *jobClusterKeyField
	}
	newClusterField, err := compute.ClusterSpecFromPb(&pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = *newClusterField
	}

	return st, nil
}

type JobCompliance struct {
	// Whether this job is in compliance with the latest version of its policy.
	IsCompliant bool
	// Canonical unique identifier for a job.
	JobId int64
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations map[string]string

	ForceSendFields []string
}

func jobComplianceToPb(st *JobCompliance) (*jobCompliancePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobCompliancePb{}
	isCompliantPb, err := identity(&st.IsCompliant)
	if err != nil {
		return nil, err
	}
	if isCompliantPb != nil {
		pb.IsCompliant = *isCompliantPb
	}

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
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

func (st *JobCompliance) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobCompliancePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobComplianceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobCompliance) MarshalJSON() ([]byte, error) {
	pb, err := jobComplianceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobCompliancePb struct {
	// Whether this job is in compliance with the latest version of its policy.
	IsCompliant bool `json:"is_compliant,omitempty"`
	// Canonical unique identifier for a job.
	JobId int64 `json:"job_id"`
	// An object containing key-value mappings representing the first 200 policy
	// validation errors. The keys indicate the path where the policy validation
	// error is occurring. An identifier for the job cluster is prepended to the
	// path. The values indicate an error message describing the policy
	// validation error.
	Violations map[string]string `json:"violations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobComplianceFromPb(pb *jobCompliancePb) (*JobCompliance, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobCompliance{}
	isCompliantField, err := identity(&pb.IsCompliant)
	if err != nil {
		return nil, err
	}
	if isCompliantField != nil {
		st.IsCompliant = *isCompliantField
	}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
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

func (st *jobCompliancePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobCompliancePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobDeployment struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	Kind JobDeploymentKind
	// Path of the file that contains deployment metadata.
	MetadataFilePath string

	ForceSendFields []string
}

func jobDeploymentToPb(st *JobDeployment) (*jobDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobDeploymentPb{}
	kindPb, err := identity(&st.Kind)
	if err != nil {
		return nil, err
	}
	if kindPb != nil {
		pb.Kind = *kindPb
	}

	metadataFilePathPb, err := identity(&st.MetadataFilePath)
	if err != nil {
		return nil, err
	}
	if metadataFilePathPb != nil {
		pb.MetadataFilePath = *metadataFilePathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *JobDeployment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobDeploymentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobDeploymentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobDeployment) MarshalJSON() ([]byte, error) {
	pb, err := jobDeploymentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobDeploymentPb struct {
	// The kind of deployment that manages the job.
	//
	// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
	Kind JobDeploymentKind `json:"kind"`
	// Path of the file that contains deployment metadata.
	MetadataFilePath string `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobDeploymentFromPb(pb *jobDeploymentPb) (*JobDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobDeployment{}
	kindField, err := identity(&pb.Kind)
	if err != nil {
		return nil, err
	}
	if kindField != nil {
		st.Kind = *kindField
	}
	metadataFilePathField, err := identity(&pb.MetadataFilePath)
	if err != nil {
		return nil, err
	}
	if metadataFilePathField != nil {
		st.MetadataFilePath = *metadataFilePathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// * `BUNDLE`: The job is managed by Databricks Asset Bundle.
type JobDeploymentKind string
type jobDeploymentKindPb string

// The job is managed by Databricks Asset Bundle.
const JobDeploymentKindBundle JobDeploymentKind = `BUNDLE`

// String representation for [fmt.Print]
func (f *JobDeploymentKind) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobDeploymentKind) Set(v string) error {
	switch v {
	case `BUNDLE`:
		*f = JobDeploymentKind(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUNDLE"`, v)
	}
}

// Type always returns JobDeploymentKind to satisfy [pflag.Value] interface
func (f *JobDeploymentKind) Type() string {
	return "JobDeploymentKind"
}

func jobDeploymentKindToPb(st *JobDeploymentKind) (*jobDeploymentKindPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobDeploymentKindPb(*st)
	return &pb, nil
}

func jobDeploymentKindFromPb(pb *jobDeploymentKindPb) (*JobDeploymentKind, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobDeploymentKind(*pb)
	return &st, nil
}

// Edit mode of the job.
//
// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
// `EDITABLE`: The job is in an editable state and can be modified.
type JobEditMode string
type jobEditModePb string

// The job is in an editable state and can be modified.
const JobEditModeEditable JobEditMode = `EDITABLE`

// The job is in a locked UI state and cannot be modified.
const JobEditModeUiLocked JobEditMode = `UI_LOCKED`

// String representation for [fmt.Print]
func (f *JobEditMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobEditMode) Set(v string) error {
	switch v {
	case `EDITABLE`, `UI_LOCKED`:
		*f = JobEditMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EDITABLE", "UI_LOCKED"`, v)
	}
}

// Type always returns JobEditMode to satisfy [pflag.Value] interface
func (f *JobEditMode) Type() string {
	return "JobEditMode"
}

func jobEditModeToPb(st *JobEditMode) (*jobEditModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobEditModePb(*st)
	return &pb, nil
}

func jobEditModeFromPb(pb *jobEditModePb) (*JobEditMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobEditMode(*pb)
	return &st, nil
}

type JobEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns bool
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []string
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string

	ForceSendFields []string
}

func jobEmailNotificationsToPb(st *JobEmailNotifications) (*jobEmailNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobEmailNotificationsPb{}
	noAlertForSkippedRunsPb, err := identity(&st.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsPb != nil {
		pb.NoAlertForSkippedRuns = *noAlertForSkippedRunsPb
	}

	var onDurationWarningThresholdExceededPb []string
	for _, item := range st.OnDurationWarningThresholdExceeded {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onDurationWarningThresholdExceededPb = append(onDurationWarningThresholdExceededPb, *itemPb)
		}
	}
	pb.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededPb

	var onFailurePb []string
	for _, item := range st.OnFailure {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onFailurePb = append(onFailurePb, *itemPb)
		}
	}
	pb.OnFailure = onFailurePb

	var onStartPb []string
	for _, item := range st.OnStart {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStartPb = append(onStartPb, *itemPb)
		}
	}
	pb.OnStart = onStartPb

	var onStreamingBacklogExceededPb []string
	for _, item := range st.OnStreamingBacklogExceeded {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStreamingBacklogExceededPb = append(onStreamingBacklogExceededPb, *itemPb)
		}
	}
	pb.OnStreamingBacklogExceeded = onStreamingBacklogExceededPb

	var onSuccessPb []string
	for _, item := range st.OnSuccess {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onSuccessPb = append(onSuccessPb, *itemPb)
		}
	}
	pb.OnSuccess = onSuccessPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *JobEmailNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobEmailNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobEmailNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobEmailNotifications) MarshalJSON() ([]byte, error) {
	pb, err := jobEmailNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobEmailNotificationsPb struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []string `json:"on_streaming_backlog_exceeded,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobEmailNotificationsFromPb(pb *jobEmailNotificationsPb) (*JobEmailNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobEmailNotifications{}
	noAlertForSkippedRunsField, err := identity(&pb.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsField != nil {
		st.NoAlertForSkippedRuns = *noAlertForSkippedRunsField
	}

	var onDurationWarningThresholdExceededField []string
	for _, itemPb := range pb.OnDurationWarningThresholdExceeded {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onDurationWarningThresholdExceededField = append(onDurationWarningThresholdExceededField, *item)
		}
	}
	st.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededField

	var onFailureField []string
	for _, itemPb := range pb.OnFailure {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onFailureField = append(onFailureField, *item)
		}
	}
	st.OnFailure = onFailureField

	var onStartField []string
	for _, itemPb := range pb.OnStart {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStartField = append(onStartField, *item)
		}
	}
	st.OnStart = onStartField

	var onStreamingBacklogExceededField []string
	for _, itemPb := range pb.OnStreamingBacklogExceeded {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStreamingBacklogExceededField = append(onStreamingBacklogExceededField, *item)
		}
	}
	st.OnStreamingBacklogExceeded = onStreamingBacklogExceededField

	var onSuccessField []string
	for _, itemPb := range pb.OnSuccess {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onSuccessField = append(onSuccessField, *item)
		}
	}
	st.OnSuccess = onSuccessField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobEmailNotificationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobEmailNotificationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobEnvironment struct {
	// The key of an environment. It has to be unique within a job.
	EnvironmentKey string
	// The environment entity used to preserve serverless environment side
	// panel, jobs' environment for non-notebook task, and DLT's environment for
	// classic and serverless pipelines. (Note: DLT uses a copied version of the
	// Environment proto below, at
	// //spark/pipelines/api/protos/copied/libraries-environments-copy.proto) In
	// this minimal environment spec, only pip dependencies are supported.
	Spec *compute.Environment
}

func jobEnvironmentToPb(st *JobEnvironment) (*jobEnvironmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobEnvironmentPb{}
	environmentKeyPb, err := identity(&st.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyPb != nil {
		pb.EnvironmentKey = *environmentKeyPb
	}

	specPb, err := compute.EnvironmentToPb(st.Spec)
	if err != nil {
		return nil, err
	}
	if specPb != nil {
		pb.Spec = specPb
	}

	return pb, nil
}

func (st *JobEnvironment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobEnvironmentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobEnvironmentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobEnvironment) MarshalJSON() ([]byte, error) {
	pb, err := jobEnvironmentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobEnvironmentPb struct {
	// The key of an environment. It has to be unique within a job.
	EnvironmentKey string `json:"environment_key"`
	// The environment entity used to preserve serverless environment side
	// panel, jobs' environment for non-notebook task, and DLT's environment for
	// classic and serverless pipelines. (Note: DLT uses a copied version of the
	// Environment proto below, at
	// //spark/pipelines/api/protos/copied/libraries-environments-copy.proto) In
	// this minimal environment spec, only pip dependencies are supported.
	Spec *compute.EnvironmentPb `json:"spec,omitempty"`
}

func jobEnvironmentFromPb(pb *jobEnvironmentPb) (*JobEnvironment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobEnvironment{}
	environmentKeyField, err := identity(&pb.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyField != nil {
		st.EnvironmentKey = *environmentKeyField
	}
	specField, err := compute.EnvironmentFromPb(pb.Spec)
	if err != nil {
		return nil, err
	}
	if specField != nil {
		st.Spec = specField
	}

	return st, nil
}

type JobNotificationSettings struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool

	ForceSendFields []string
}

func jobNotificationSettingsToPb(st *JobNotificationSettings) (*jobNotificationSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobNotificationSettingsPb{}
	noAlertForCanceledRunsPb, err := identity(&st.NoAlertForCanceledRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForCanceledRunsPb != nil {
		pb.NoAlertForCanceledRuns = *noAlertForCanceledRunsPb
	}

	noAlertForSkippedRunsPb, err := identity(&st.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsPb != nil {
		pb.NoAlertForSkippedRuns = *noAlertForSkippedRunsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *JobNotificationSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobNotificationSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobNotificationSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobNotificationSettings) MarshalJSON() ([]byte, error) {
	pb, err := jobNotificationSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobNotificationSettingsPb struct {
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobNotificationSettingsFromPb(pb *jobNotificationSettingsPb) (*JobNotificationSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobNotificationSettings{}
	noAlertForCanceledRunsField, err := identity(&pb.NoAlertForCanceledRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForCanceledRunsField != nil {
		st.NoAlertForCanceledRuns = *noAlertForCanceledRunsField
	}
	noAlertForSkippedRunsField, err := identity(&pb.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsField != nil {
		st.NoAlertForSkippedRuns = *noAlertForSkippedRunsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobNotificationSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobNotificationSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobParameter struct {
	// The optional default value of the parameter
	Default string
	// The name of the parameter
	Name string
	// The value used in the run
	Value string

	ForceSendFields []string
}

func jobParameterToPb(st *JobParameter) (*jobParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobParameterPb{}
	defaultPb, err := identity(&st.Default)
	if err != nil {
		return nil, err
	}
	if defaultPb != nil {
		pb.Default = *defaultPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
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

func (st *JobParameter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobParameterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobParameterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobParameter) MarshalJSON() ([]byte, error) {
	pb, err := jobParameterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobParameterPb struct {
	// The optional default value of the parameter
	Default string `json:"default,omitempty"`
	// The name of the parameter
	Name string `json:"name,omitempty"`
	// The value used in the run
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobParameterFromPb(pb *jobParameterPb) (*JobParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobParameter{}
	defaultField, err := identity(&pb.Default)
	if err != nil {
		return nil, err
	}
	if defaultField != nil {
		st.Default = *defaultField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
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

func (st *jobParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobParameterDefinition struct {
	// Default value of the parameter.
	Default string
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	Name string
}

func jobParameterDefinitionToPb(st *JobParameterDefinition) (*jobParameterDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobParameterDefinitionPb{}
	defaultPb, err := identity(&st.Default)
	if err != nil {
		return nil, err
	}
	if defaultPb != nil {
		pb.Default = *defaultPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
}

func (st *JobParameterDefinition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobParameterDefinitionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobParameterDefinitionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobParameterDefinition) MarshalJSON() ([]byte, error) {
	pb, err := jobParameterDefinitionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobParameterDefinitionPb struct {
	// Default value of the parameter.
	Default string `json:"default"`
	// The name of the defined parameter. May only contain alphanumeric
	// characters, `_`, `-`, and `.`
	Name string `json:"name"`
}

func jobParameterDefinitionFromPb(pb *jobParameterDefinitionPb) (*JobParameterDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobParameterDefinition{}
	defaultField, err := identity(&pb.Default)
	if err != nil {
		return nil, err
	}
	if defaultField != nil {
		st.Default = *defaultField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type JobPermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel JobPermissionLevel

	ForceSendFields []string
}

func jobPermissionToPb(st *JobPermission) (*jobPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionPb{}
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

func (st *JobPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermission) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobPermissionFromPb(pb *jobPermissionPb) (*JobPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermission{}
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

func (st *jobPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type JobPermissionLevel string
type jobPermissionLevelPb string

const JobPermissionLevelCanManage JobPermissionLevel = `CAN_MANAGE`

const JobPermissionLevelCanManageRun JobPermissionLevel = `CAN_MANAGE_RUN`

const JobPermissionLevelCanView JobPermissionLevel = `CAN_VIEW`

const JobPermissionLevelIsOwner JobPermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *JobPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MANAGE_RUN`, `CAN_VIEW`, `IS_OWNER`:
		*f = JobPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MANAGE_RUN", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Type always returns JobPermissionLevel to satisfy [pflag.Value] interface
func (f *JobPermissionLevel) Type() string {
	return "JobPermissionLevel"
}

func jobPermissionLevelToPb(st *JobPermissionLevel) (*jobPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobPermissionLevelPb(*st)
	return &pb, nil
}

func jobPermissionLevelFromPb(pb *jobPermissionLevelPb) (*JobPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobPermissionLevel(*pb)
	return &st, nil
}

type JobPermissions struct {
	AccessControlList []JobAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
}

func jobPermissionsToPb(st *JobPermissions) (*jobPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionsPb{}

	var accessControlListPb []jobAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := jobAccessControlResponseToPb(&item)
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

func (st *JobPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermissions) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobPermissionsPb struct {
	AccessControlList []jobAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobPermissionsFromPb(pb *jobPermissionsPb) (*JobPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissions{}

	var accessControlListField []JobAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := jobAccessControlResponseFromPb(&itemPb)
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

func (st *jobPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobPermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel JobPermissionLevel

	ForceSendFields []string
}

func jobPermissionsDescriptionToPb(st *JobPermissionsDescription) (*jobPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionsDescriptionPb{}
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

func (st *JobPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel JobPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobPermissionsDescriptionFromPb(pb *jobPermissionsDescriptionPb) (*JobPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissionsDescription{}
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

func (st *jobPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobPermissionsRequest struct {
	AccessControlList []JobAccessControlRequest
	// The job for which to get or manage permissions.
	JobId string
}

func jobPermissionsRequestToPb(st *JobPermissionsRequest) (*jobPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobPermissionsRequestPb{}

	var accessControlListPb []jobAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := jobAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	return pb, nil
}

func (st *JobPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := jobPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobPermissionsRequestPb struct {
	AccessControlList []jobAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The job for which to get or manage permissions.
	JobId string `json:"-" url:"-"`
}

func jobPermissionsRequestFromPb(pb *jobPermissionsRequestPb) (*JobPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobPermissionsRequest{}

	var accessControlListField []JobAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := jobAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	return st, nil
}

// Write-only setting. Specifies the user or service principal that the job runs
// as. If not specified, the job runs as the user who created the job.
//
// Either `user_name` or `service_principal_name` should be specified. If not,
// an error is thrown.
type JobRunAs struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	UserName string

	ForceSendFields []string
}

func jobRunAsToPb(st *JobRunAs) (*jobRunAsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobRunAsPb{}
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

func (st *JobRunAs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobRunAsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobRunAsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobRunAs) MarshalJSON() ([]byte, error) {
	pb, err := jobRunAsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobRunAsPb struct {
	// Application ID of an active service principal. Setting this field
	// requires the `servicePrincipal/user` role.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// The email of an active workspace user. Non-admin users can only set this
	// field to their own email.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobRunAsFromPb(pb *jobRunAsPb) (*JobRunAs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobRunAs{}
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

func (st *jobRunAsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobRunAsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobSettings struct {
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId string
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *Continuous
	// Deployment information for jobs managed by external sources.
	Deployment *JobDeployment
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description string
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *JobEmailNotifications
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	Environments []JobEnvironment
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []JobCluster
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	MaxConcurrentRuns int
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *JobNotificationSettings
	// Job-level parameter definitions
	Parameters []JobParameterDefinition
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget
	// The queue settings of the job.
	Queue *QueueSettings
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	RunAs *JobRunAs
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *CronSchedule
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	Tasks []Task
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *TriggerSettings
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string
}

func jobSettingsToPb(st *JobSettings) (*jobSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSettingsPb{}
	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	continuousPb, err := continuousToPb(st.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousPb != nil {
		pb.Continuous = continuousPb
	}

	deploymentPb, err := jobDeploymentToPb(st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = deploymentPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	editModePb, err := identity(&st.EditMode)
	if err != nil {
		return nil, err
	}
	if editModePb != nil {
		pb.EditMode = *editModePb
	}

	emailNotificationsPb, err := jobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	var environmentsPb []jobEnvironmentPb
	for _, item := range st.Environments {
		itemPb, err := jobEnvironmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentsPb = append(environmentsPb, *itemPb)
		}
	}
	pb.Environments = environmentsPb

	formatPb, err := identity(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}

	gitSourcePb, err := gitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}

	healthPb, err := jobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	var jobClustersPb []jobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := jobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb

	maxConcurrentRunsPb, err := identity(&st.MaxConcurrentRuns)
	if err != nil {
		return nil, err
	}
	if maxConcurrentRunsPb != nil {
		pb.MaxConcurrentRuns = *maxConcurrentRunsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	notificationSettingsPb, err := jobNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	var parametersPb []jobParameterDefinitionPb
	for _, item := range st.Parameters {
		itemPb, err := jobParameterDefinitionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	performanceTargetPb, err := identity(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}

	queuePb, err := queueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}

	runAsPb, err := jobRunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}

	schedulePb, err := cronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	tagsPb := map[string]string{}
	for k, v := range st.Tags {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb[k] = *itemPb
		}
	}
	pb.Tags = tagsPb

	var tasksPb []taskPb
	for _, item := range st.Tasks {
		itemPb, err := taskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb

	timeoutSecondsPb, err := identity(&st.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsPb != nil {
		pb.TimeoutSeconds = *timeoutSecondsPb
	}

	triggerPb, err := triggerSettingsToPb(st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = triggerPb
	}

	webhookNotificationsPb, err := webhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *JobSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSettings) MarshalJSON() ([]byte, error) {
	pb, err := jobSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobSettingsPb struct {
	// The id of the user specified budget policy to use for this job. If not
	// specified, a default budget policy may be applied when creating or
	// modifying the job. See `effective_budget_policy_id` for the budget policy
	// used by this workload.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// An optional continuous property for this job. The continuous property
	// will ensure that there is always one run executing. Only one of
	// `schedule` and `continuous` can be used.
	Continuous *continuousPb `json:"continuous,omitempty"`
	// Deployment information for jobs managed by external sources.
	Deployment *jobDeploymentPb `json:"deployment,omitempty"`
	// An optional description for the job. The maximum length is 27700
	// characters in UTF-8 encoding.
	Description string `json:"description,omitempty"`
	// Edit mode of the job.
	//
	// * `UI_LOCKED`: The job is in a locked UI state and cannot be modified. *
	// `EDITABLE`: The job is in an editable state and can be modified.
	EditMode JobEditMode `json:"edit_mode,omitempty"`
	// An optional set of email addresses that is notified when runs of this job
	// begin or complete as well as when this job is deleted.
	EmailNotifications *jobEmailNotificationsPb `json:"email_notifications,omitempty"`
	// A list of task execution environment specifications that can be
	// referenced by serverless tasks of this job. An environment is required to
	// be present for serverless tasks. For serverless notebook tasks, the
	// environment is accessible in the notebook environment panel. For other
	// serverless tasks, the task environment is required to be specified using
	// environment_key in the task settings.
	Environments []jobEnvironmentPb `json:"environments,omitempty"`
	// Used to tell what is the format of the job. This field is ignored in
	// Create/Update/Reset calls. When using the Jobs API 2.1 this value is
	// always set to `"MULTI_TASK"`.
	Format Format `json:"format,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *gitSourcePb `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *jobsHealthRulesPb `json:"health,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings.
	JobClusters []jobClusterPb `json:"job_clusters,omitempty"`
	// An optional maximum allowed number of concurrent runs of the job. Set
	// this value if you want to be able to execute multiple runs of the same
	// job concurrently. This is useful for example if you trigger your job on a
	// frequent schedule and want to allow consecutive runs to overlap with each
	// other, or if you want to trigger multiple runs which differ by their
	// input parameters. This setting affects only new runs. For example,
	// suppose the job’s concurrency is 4 and there are 4 concurrent active
	// runs. Then setting the concurrency to 3 won’t kill any of the active
	// runs. However, from then on, new runs are skipped unless there are fewer
	// than 3 active runs. This value cannot exceed 1000. Setting this value to
	// `0` causes all new runs to be skipped.
	MaxConcurrentRuns int `json:"max_concurrent_runs,omitempty"`
	// An optional name for the job. The maximum length is 4096 bytes in UTF-8
	// encoding.
	Name string `json:"name,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// job.
	NotificationSettings *jobNotificationSettingsPb `json:"notification_settings,omitempty"`
	// Job-level parameter definitions
	Parameters []jobParameterDefinitionPb `json:"parameters,omitempty"`
	// The performance mode on a serverless job. This field determines the level
	// of compute performance or cost-efficiency for the run.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// The queue settings of the job.
	Queue *queueSettingsPb `json:"queue,omitempty"`
	// Write-only setting. Specifies the user or service principal that the job
	// runs as. If not specified, the job runs as the user who created the job.
	//
	// Either `user_name` or `service_principal_name` should be specified. If
	// not, an error is thrown.
	RunAs *jobRunAsPb `json:"run_as,omitempty"`
	// An optional periodic schedule for this job. The default behavior is that
	// the job only runs when triggered by clicking “Run Now” in the Jobs UI
	// or sending an API request to `runNow`.
	Schedule *cronSchedulePb `json:"schedule,omitempty"`
	// A map of tags associated with the job. These are forwarded to the cluster
	// as cluster tags for jobs clusters, and are subject to the same
	// limitations as cluster tags. A maximum of 25 tags can be added to the
	// job.
	Tags map[string]string `json:"tags,omitempty"`
	// A list of task specifications to be executed by this job. It supports up
	// to 1000 elements in write endpoints (:method:jobs/create,
	// :method:jobs/reset, :method:jobs/update, :method:jobs/submit). Read
	// endpoints return only 100 tasks. If more than 100 tasks are available,
	// you can paginate through them using :method:jobs/get. Use the
	// `next_page_token` field at the object root to determine if more results
	// are available.
	Tasks []taskPb `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A configuration to trigger a run when certain conditions are met. The
	// default behavior is that the job runs only when triggered by clicking
	// “Run Now” in the Jobs UI or sending an API request to `runNow`.
	Trigger *triggerSettingsPb `json:"trigger,omitempty"`
	// A collection of system notification IDs to notify when runs of this job
	// begin or complete.
	WebhookNotifications *webhookNotificationsPb `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobSettingsFromPb(pb *jobSettingsPb) (*JobSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSettings{}
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	continuousField, err := continuousFromPb(pb.Continuous)
	if err != nil {
		return nil, err
	}
	if continuousField != nil {
		st.Continuous = continuousField
	}
	deploymentField, err := jobDeploymentFromPb(pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = deploymentField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	editModeField, err := identity(&pb.EditMode)
	if err != nil {
		return nil, err
	}
	if editModeField != nil {
		st.EditMode = *editModeField
	}
	emailNotificationsField, err := jobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}

	var environmentsField []JobEnvironment
	for _, itemPb := range pb.Environments {
		item, err := jobEnvironmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentsField = append(environmentsField, *item)
		}
	}
	st.Environments = environmentsField
	formatField, err := identity(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	gitSourceField, err := gitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	healthField, err := jobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := jobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	maxConcurrentRunsField, err := identity(&pb.MaxConcurrentRuns)
	if err != nil {
		return nil, err
	}
	if maxConcurrentRunsField != nil {
		st.MaxConcurrentRuns = *maxConcurrentRunsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	notificationSettingsField, err := jobNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}

	var parametersField []JobParameterDefinition
	for _, itemPb := range pb.Parameters {
		item, err := jobParameterDefinitionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	performanceTargetField, err := identity(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	queueField, err := queueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}
	runAsField, err := jobRunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	scheduleField, err := cronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}

	tagsField := map[string]string{}
	for k, v := range pb.Tags {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField[k] = *item
		}
	}
	st.Tags = tagsField

	var tasksField []Task
	for _, itemPb := range pb.Tasks {
		item, err := taskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	timeoutSecondsField, err := identity(&pb.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsField != nil {
		st.TimeoutSeconds = *timeoutSecondsField
	}
	triggerField, err := triggerSettingsFromPb(pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = triggerField
	}
	webhookNotificationsField, err := webhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The source of the job specification in the remote repository when the job is
// source controlled.
type JobSource struct {
	// Dirty state indicates the job is not fully synced with the job
	// specification in the remote repository.
	//
	// Possible values are: * `NOT_SYNCED`: The job is not yet synced with the
	// remote job specification. Import the remote job specification from UI to
	// make the job fully synced. * `DISCONNECTED`: The job is temporary
	// disconnected from the remote job specification and is allowed for live
	// edit. Import the remote job specification again from UI to make the job
	// fully synced.
	DirtyState JobSourceDirtyState
	// Name of the branch which the job is imported from.
	ImportFromGitBranch string
	// Path of the job YAML file that contains the job specification.
	JobConfigPath string
}

func jobSourceToPb(st *JobSource) (*jobSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSourcePb{}
	dirtyStatePb, err := identity(&st.DirtyState)
	if err != nil {
		return nil, err
	}
	if dirtyStatePb != nil {
		pb.DirtyState = *dirtyStatePb
	}

	importFromGitBranchPb, err := identity(&st.ImportFromGitBranch)
	if err != nil {
		return nil, err
	}
	if importFromGitBranchPb != nil {
		pb.ImportFromGitBranch = *importFromGitBranchPb
	}

	jobConfigPathPb, err := identity(&st.JobConfigPath)
	if err != nil {
		return nil, err
	}
	if jobConfigPathPb != nil {
		pb.JobConfigPath = *jobConfigPathPb
	}

	return pb, nil
}

func (st *JobSource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSource) MarshalJSON() ([]byte, error) {
	pb, err := jobSourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobSourcePb struct {
	// Dirty state indicates the job is not fully synced with the job
	// specification in the remote repository.
	//
	// Possible values are: * `NOT_SYNCED`: The job is not yet synced with the
	// remote job specification. Import the remote job specification from UI to
	// make the job fully synced. * `DISCONNECTED`: The job is temporary
	// disconnected from the remote job specification and is allowed for live
	// edit. Import the remote job specification again from UI to make the job
	// fully synced.
	DirtyState JobSourceDirtyState `json:"dirty_state,omitempty"`
	// Name of the branch which the job is imported from.
	ImportFromGitBranch string `json:"import_from_git_branch"`
	// Path of the job YAML file that contains the job specification.
	JobConfigPath string `json:"job_config_path"`
}

func jobSourceFromPb(pb *jobSourcePb) (*JobSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSource{}
	dirtyStateField, err := identity(&pb.DirtyState)
	if err != nil {
		return nil, err
	}
	if dirtyStateField != nil {
		st.DirtyState = *dirtyStateField
	}
	importFromGitBranchField, err := identity(&pb.ImportFromGitBranch)
	if err != nil {
		return nil, err
	}
	if importFromGitBranchField != nil {
		st.ImportFromGitBranch = *importFromGitBranchField
	}
	jobConfigPathField, err := identity(&pb.JobConfigPath)
	if err != nil {
		return nil, err
	}
	if jobConfigPathField != nil {
		st.JobConfigPath = *jobConfigPathField
	}

	return st, nil
}

// Dirty state indicates the job is not fully synced with the job specification
// in the remote repository.
//
// Possible values are: * `NOT_SYNCED`: The job is not yet synced with the
// remote job specification. Import the remote job specification from UI to make
// the job fully synced. * `DISCONNECTED`: The job is temporary disconnected
// from the remote job specification and is allowed for live edit. Import the
// remote job specification again from UI to make the job fully synced.
type JobSourceDirtyState string
type jobSourceDirtyStatePb string

// The job is temporary disconnected from the remote job specification and is
// allowed for live edit. Import the remote job specification again from UI to
// make the job fully synced.
const JobSourceDirtyStateDisconnected JobSourceDirtyState = `DISCONNECTED`

// The job is not yet synced with the remote job specification. Import the
// remote job specification from UI to make the job fully synced.
const JobSourceDirtyStateNotSynced JobSourceDirtyState = `NOT_SYNCED`

// String representation for [fmt.Print]
func (f *JobSourceDirtyState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobSourceDirtyState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `NOT_SYNCED`:
		*f = JobSourceDirtyState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "NOT_SYNCED"`, v)
	}
}

// Type always returns JobSourceDirtyState to satisfy [pflag.Value] interface
func (f *JobSourceDirtyState) Type() string {
	return "JobSourceDirtyState"
}

func jobSourceDirtyStateToPb(st *JobSourceDirtyState) (*jobSourceDirtyStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobSourceDirtyStatePb(*st)
	return &pb, nil
}

func jobSourceDirtyStateFromPb(pb *jobSourceDirtyStatePb) (*JobSourceDirtyState, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobSourceDirtyState(*pb)
	return &st, nil
}

// Specifies the health metric that is being evaluated for a particular health
// rule.
//
// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data waiting
// to be consumed across all streams. This metric is in Public Preview. *
// `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset lag across all
// streams. This metric is in Public Preview. * `STREAMING_BACKLOG_SECONDS`: An
// estimate of the maximum consumer delay across all streams. This metric is in
// Public Preview. * `STREAMING_BACKLOG_FILES`: An estimate of the maximum
// number of outstanding files across all streams. This metric is in Public
// Preview.
type JobsHealthMetric string
type jobsHealthMetricPb string

// Expected total time for a run in seconds.
const JobsHealthMetricRunDurationSeconds JobsHealthMetric = `RUN_DURATION_SECONDS`

// An estimate of the maximum bytes of data waiting to be consumed across all
// streams. This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogBytes JobsHealthMetric = `STREAMING_BACKLOG_BYTES`

// An estimate of the maximum number of outstanding files across all streams.
// This metric is in Public Preview.
const JobsHealthMetricStreamingBacklogFiles JobsHealthMetric = `STREAMING_BACKLOG_FILES`

// An estimate of the maximum offset lag across all streams. This metric is in
// Public Preview.
const JobsHealthMetricStreamingBacklogRecords JobsHealthMetric = `STREAMING_BACKLOG_RECORDS`

// An estimate of the maximum consumer delay across all streams. This metric is
// in Public Preview.
const JobsHealthMetricStreamingBacklogSeconds JobsHealthMetric = `STREAMING_BACKLOG_SECONDS`

// String representation for [fmt.Print]
func (f *JobsHealthMetric) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobsHealthMetric) Set(v string) error {
	switch v {
	case `RUN_DURATION_SECONDS`, `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_FILES`, `STREAMING_BACKLOG_RECORDS`, `STREAMING_BACKLOG_SECONDS`:
		*f = JobsHealthMetric(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "RUN_DURATION_SECONDS", "STREAMING_BACKLOG_BYTES", "STREAMING_BACKLOG_FILES", "STREAMING_BACKLOG_RECORDS", "STREAMING_BACKLOG_SECONDS"`, v)
	}
}

// Type always returns JobsHealthMetric to satisfy [pflag.Value] interface
func (f *JobsHealthMetric) Type() string {
	return "JobsHealthMetric"
}

func jobsHealthMetricToPb(st *JobsHealthMetric) (*jobsHealthMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobsHealthMetricPb(*st)
	return &pb, nil
}

func jobsHealthMetricFromPb(pb *jobsHealthMetricPb) (*JobsHealthMetric, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobsHealthMetric(*pb)
	return &st, nil
}

// Specifies the operator used to compare the health metric value with the
// specified threshold.
type JobsHealthOperator string
type jobsHealthOperatorPb string

const JobsHealthOperatorGreaterThan JobsHealthOperator = `GREATER_THAN`

// String representation for [fmt.Print]
func (f *JobsHealthOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *JobsHealthOperator) Set(v string) error {
	switch v {
	case `GREATER_THAN`:
		*f = JobsHealthOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GREATER_THAN"`, v)
	}
}

// Type always returns JobsHealthOperator to satisfy [pflag.Value] interface
func (f *JobsHealthOperator) Type() string {
	return "JobsHealthOperator"
}

func jobsHealthOperatorToPb(st *JobsHealthOperator) (*jobsHealthOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := jobsHealthOperatorPb(*st)
	return &pb, nil
}

func jobsHealthOperatorFromPb(pb *jobsHealthOperatorPb) (*JobsHealthOperator, error) {
	if pb == nil {
		return nil, nil
	}
	st := JobsHealthOperator(*pb)
	return &st, nil
}

type JobsHealthRule struct {
	// Specifies the health metric that is being evaluated for a particular
	// health rule.
	//
	// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
	// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data
	// waiting to be consumed across all streams. This metric is in Public
	// Preview. * `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset
	// lag across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_SECONDS`: An estimate of the maximum consumer delay
	// across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_FILES`: An estimate of the maximum number of
	// outstanding files across all streams. This metric is in Public Preview.
	Metric JobsHealthMetric
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	Op JobsHealthOperator
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	Value int64
}

func jobsHealthRuleToPb(st *JobsHealthRule) (*jobsHealthRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobsHealthRulePb{}
	metricPb, err := identity(&st.Metric)
	if err != nil {
		return nil, err
	}
	if metricPb != nil {
		pb.Metric = *metricPb
	}

	opPb, err := identity(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}

	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	return pb, nil
}

func (st *JobsHealthRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobsHealthRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobsHealthRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobsHealthRule) MarshalJSON() ([]byte, error) {
	pb, err := jobsHealthRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobsHealthRulePb struct {
	// Specifies the health metric that is being evaluated for a particular
	// health rule.
	//
	// * `RUN_DURATION_SECONDS`: Expected total time for a run in seconds. *
	// `STREAMING_BACKLOG_BYTES`: An estimate of the maximum bytes of data
	// waiting to be consumed across all streams. This metric is in Public
	// Preview. * `STREAMING_BACKLOG_RECORDS`: An estimate of the maximum offset
	// lag across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_SECONDS`: An estimate of the maximum consumer delay
	// across all streams. This metric is in Public Preview. *
	// `STREAMING_BACKLOG_FILES`: An estimate of the maximum number of
	// outstanding files across all streams. This metric is in Public Preview.
	Metric JobsHealthMetric `json:"metric"`
	// Specifies the operator used to compare the health metric value with the
	// specified threshold.
	Op JobsHealthOperator `json:"op"`
	// Specifies the threshold value that the health metric should obey to
	// satisfy the health rule.
	Value int64 `json:"value"`
}

func jobsHealthRuleFromPb(pb *jobsHealthRulePb) (*JobsHealthRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobsHealthRule{}
	metricField, err := identity(&pb.Metric)
	if err != nil {
		return nil, err
	}
	if metricField != nil {
		st.Metric = *metricField
	}
	opField, err := identity(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

	return st, nil
}

// An optional set of health rules that can be defined for this job.
type JobsHealthRules struct {
	Rules []JobsHealthRule
}

func jobsHealthRulesToPb(st *JobsHealthRules) (*jobsHealthRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobsHealthRulesPb{}

	var rulesPb []jobsHealthRulePb
	for _, item := range st.Rules {
		itemPb, err := jobsHealthRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rulesPb = append(rulesPb, *itemPb)
		}
	}
	pb.Rules = rulesPb

	return pb, nil
}

func (st *JobsHealthRules) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobsHealthRulesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobsHealthRulesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobsHealthRules) MarshalJSON() ([]byte, error) {
	pb, err := jobsHealthRulesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobsHealthRulesPb struct {
	Rules []jobsHealthRulePb `json:"rules,omitempty"`
}

func jobsHealthRulesFromPb(pb *jobsHealthRulesPb) (*JobsHealthRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobsHealthRules{}

	var rulesField []JobsHealthRule
	for _, itemPb := range pb.Rules {
		item, err := jobsHealthRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rulesField = append(rulesField, *item)
		}
	}
	st.Rules = rulesField

	return st, nil
}

type ListJobComplianceForPolicyResponse struct {
	// A list of jobs and their policy compliance statuses.
	Jobs []JobCompliance
	// This field represents the pagination token to retrieve the next page of
	// results. If this field is not in the response, it means no further
	// results for the request.
	NextPageToken string
	// This field represents the pagination token to retrieve the previous page
	// of results. If this field is not in the response, it means no further
	// results for the request.
	PrevPageToken string

	ForceSendFields []string
}

func listJobComplianceForPolicyResponseToPb(st *ListJobComplianceForPolicyResponse) (*listJobComplianceForPolicyResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobComplianceForPolicyResponsePb{}

	var jobsPb []jobCompliancePb
	for _, item := range st.Jobs {
		itemPb, err := jobComplianceToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobsPb = append(jobsPb, *itemPb)
		}
	}
	pb.Jobs = jobsPb

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

func (st *ListJobComplianceForPolicyResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobComplianceForPolicyResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobComplianceForPolicyResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobComplianceForPolicyResponse) MarshalJSON() ([]byte, error) {
	pb, err := listJobComplianceForPolicyResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listJobComplianceForPolicyResponsePb struct {
	// A list of jobs and their policy compliance statuses.
	Jobs []jobCompliancePb `json:"jobs,omitempty"`
	// This field represents the pagination token to retrieve the next page of
	// results. If this field is not in the response, it means no further
	// results for the request.
	NextPageToken string `json:"next_page_token,omitempty"`
	// This field represents the pagination token to retrieve the previous page
	// of results. If this field is not in the response, it means no further
	// results for the request.
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listJobComplianceForPolicyResponseFromPb(pb *listJobComplianceForPolicyResponsePb) (*ListJobComplianceForPolicyResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobComplianceForPolicyResponse{}

	var jobsField []JobCompliance
	for _, itemPb := range pb.Jobs {
		item, err := jobComplianceFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobsField = append(jobsField, *item)
		}
	}
	st.Jobs = jobsField
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

func (st *listJobComplianceForPolicyResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobComplianceForPolicyResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List job policy compliance
type ListJobComplianceRequest struct {
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

func listJobComplianceRequestToPb(st *ListJobComplianceRequest) (*listJobComplianceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobComplianceRequestPb{}
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

func (st *ListJobComplianceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobComplianceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobComplianceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobComplianceRequest) MarshalJSON() ([]byte, error) {
	pb, err := listJobComplianceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listJobComplianceRequestPb struct {
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

func listJobComplianceRequestFromPb(pb *listJobComplianceRequestPb) (*ListJobComplianceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobComplianceRequest{}
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

func (st *listJobComplianceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobComplianceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List jobs
type ListJobsRequest struct {
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/get to
	// paginate through all tasks and clusters.
	ExpandTasks bool
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 100. The default value is 20.
	Limit int
	// A filter on the list based on the exact (case insensitive) job name.
	Name string
	// The offset of the first job to return, relative to the most recently
	// created job. Deprecated since June 2023. Use `page_token` to iterate
	// through the pages instead.
	Offset int
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of jobs respectively.
	PageToken string

	ForceSendFields []string
}

func listJobsRequestToPb(st *ListJobsRequest) (*listJobsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobsRequestPb{}
	expandTasksPb, err := identity(&st.ExpandTasks)
	if err != nil {
		return nil, err
	}
	if expandTasksPb != nil {
		pb.ExpandTasks = *expandTasksPb
	}

	limitPb, err := identity(&st.Limit)
	if err != nil {
		return nil, err
	}
	if limitPb != nil {
		pb.Limit = *limitPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	offsetPb, err := identity(&st.Offset)
	if err != nil {
		return nil, err
	}
	if offsetPb != nil {
		pb.Offset = *offsetPb
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

func (st *ListJobsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listJobsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listJobsRequestPb struct {
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/get to
	// paginate through all tasks and clusters.
	ExpandTasks bool `json:"-" url:"expand_tasks,omitempty"`
	// The number of jobs to return. This value must be greater than 0 and less
	// or equal to 100. The default value is 20.
	Limit int `json:"-" url:"limit,omitempty"`
	// A filter on the list based on the exact (case insensitive) job name.
	Name string `json:"-" url:"name,omitempty"`
	// The offset of the first job to return, relative to the most recently
	// created job. Deprecated since June 2023. Use `page_token` to iterate
	// through the pages instead.
	Offset int `json:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of jobs respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listJobsRequestFromPb(pb *listJobsRequestPb) (*ListJobsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobsRequest{}
	expandTasksField, err := identity(&pb.ExpandTasks)
	if err != nil {
		return nil, err
	}
	if expandTasksField != nil {
		st.ExpandTasks = *expandTasksField
	}
	limitField, err := identity(&pb.Limit)
	if err != nil {
		return nil, err
	}
	if limitField != nil {
		st.Limit = *limitField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	offsetField, err := identity(&pb.Offset)
	if err != nil {
		return nil, err
	}
	if offsetField != nil {
		st.Offset = *offsetField
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

func (st *listJobsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List of jobs was retrieved successfully.
type ListJobsResponse struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	HasMore bool
	// The list of jobs. Only included in the response if there are jobs to
	// list.
	Jobs []BaseJob
	// A token that can be used to list the next page of jobs (if applicable).
	NextPageToken string
	// A token that can be used to list the previous page of jobs (if
	// applicable).
	PrevPageToken string

	ForceSendFields []string
}

func listJobsResponseToPb(st *ListJobsResponse) (*listJobsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listJobsResponsePb{}
	hasMorePb, err := identity(&st.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMorePb != nil {
		pb.HasMore = *hasMorePb
	}

	var jobsPb []baseJobPb
	for _, item := range st.Jobs {
		itemPb, err := baseJobToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobsPb = append(jobsPb, *itemPb)
		}
	}
	pb.Jobs = jobsPb

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

func (st *ListJobsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listJobsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listJobsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListJobsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listJobsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listJobsResponsePb struct {
	// If true, additional jobs matching the provided filter are available for
	// listing.
	HasMore bool `json:"has_more,omitempty"`
	// The list of jobs. Only included in the response if there are jobs to
	// list.
	Jobs []baseJobPb `json:"jobs,omitempty"`
	// A token that can be used to list the next page of jobs (if applicable).
	NextPageToken string `json:"next_page_token,omitempty"`
	// A token that can be used to list the previous page of jobs (if
	// applicable).
	PrevPageToken string `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listJobsResponseFromPb(pb *listJobsResponsePb) (*ListJobsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListJobsResponse{}
	hasMoreField, err := identity(&pb.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMoreField != nil {
		st.HasMore = *hasMoreField
	}

	var jobsField []BaseJob
	for _, itemPb := range pb.Jobs {
		item, err := baseJobFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobsField = append(jobsField, *item)
		}
	}
	st.Jobs = jobsField
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

func (st *listJobsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listJobsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List job runs
type ListRunsRequest struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `QUEUED`, `PENDING`, `RUNNING`, or `TERMINATING`. This field
	// cannot be `true` when completed_only is `true`.
	ActiveOnly bool
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	CompletedOnly bool
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/getrun to
	// paginate through all tasks and clusters.
	ExpandTasks bool
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	JobId int64
	// The number of runs to return. This value must be greater than 0 and less
	// than 25. The default value is 20. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	Limit int
	// The offset of the first run to return, relative to the most recent run.
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	Offset int
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of runs respectively.
	PageToken string
	// The type of runs to return. For a description of run types, see
	// :method:jobs/getRun.
	RunType RunType
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	StartTimeFrom int64
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	StartTimeTo int64

	ForceSendFields []string
}

func listRunsRequestToPb(st *ListRunsRequest) (*listRunsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRunsRequestPb{}
	activeOnlyPb, err := identity(&st.ActiveOnly)
	if err != nil {
		return nil, err
	}
	if activeOnlyPb != nil {
		pb.ActiveOnly = *activeOnlyPb
	}

	completedOnlyPb, err := identity(&st.CompletedOnly)
	if err != nil {
		return nil, err
	}
	if completedOnlyPb != nil {
		pb.CompletedOnly = *completedOnlyPb
	}

	expandTasksPb, err := identity(&st.ExpandTasks)
	if err != nil {
		return nil, err
	}
	if expandTasksPb != nil {
		pb.ExpandTasks = *expandTasksPb
	}

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

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

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	runTypePb, err := identity(&st.RunType)
	if err != nil {
		return nil, err
	}
	if runTypePb != nil {
		pb.RunType = *runTypePb
	}

	startTimeFromPb, err := identity(&st.StartTimeFrom)
	if err != nil {
		return nil, err
	}
	if startTimeFromPb != nil {
		pb.StartTimeFrom = *startTimeFromPb
	}

	startTimeToPb, err := identity(&st.StartTimeTo)
	if err != nil {
		return nil, err
	}
	if startTimeToPb != nil {
		pb.StartTimeTo = *startTimeToPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListRunsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRunsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRunsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRunsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listRunsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRunsRequestPb struct {
	// If active_only is `true`, only active runs are included in the results;
	// otherwise, lists both active and completed runs. An active run is a run
	// in the `QUEUED`, `PENDING`, `RUNNING`, or `TERMINATING`. This field
	// cannot be `true` when completed_only is `true`.
	ActiveOnly bool `json:"-" url:"active_only,omitempty"`
	// If completed_only is `true`, only completed runs are included in the
	// results; otherwise, lists both active and completed runs. This field
	// cannot be `true` when active_only is `true`.
	CompletedOnly bool `json:"-" url:"completed_only,omitempty"`
	// Whether to include task and cluster details in the response. Note that
	// only the first 100 elements will be shown. Use :method:jobs/getrun to
	// paginate through all tasks and clusters.
	ExpandTasks bool `json:"-" url:"expand_tasks,omitempty"`
	// The job for which to list runs. If omitted, the Jobs service lists runs
	// from all jobs.
	JobId int64 `json:"-" url:"job_id,omitempty"`
	// The number of runs to return. This value must be greater than 0 and less
	// than 25. The default value is 20. If a request specifies a limit of 0,
	// the service instead uses the maximum limit.
	Limit int `json:"-" url:"limit,omitempty"`
	// The offset of the first run to return, relative to the most recent run.
	// Deprecated since June 2023. Use `page_token` to iterate through the pages
	// instead.
	Offset int `json:"-" url:"offset,omitempty"`
	// Use `next_page_token` or `prev_page_token` returned from the previous
	// request to list the next or previous page of runs respectively.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The type of runs to return. For a description of run types, see
	// :method:jobs/getRun.
	RunType RunType `json:"-" url:"run_type,omitempty"`
	// Show runs that started _at or after_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_to_ to filter
	// by a time range.
	StartTimeFrom int64 `json:"-" url:"start_time_from,omitempty"`
	// Show runs that started _at or before_ this value. The value must be a UTC
	// timestamp in milliseconds. Can be combined with _start_time_from_ to
	// filter by a time range.
	StartTimeTo int64 `json:"-" url:"start_time_to,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRunsRequestFromPb(pb *listRunsRequestPb) (*ListRunsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRunsRequest{}
	activeOnlyField, err := identity(&pb.ActiveOnly)
	if err != nil {
		return nil, err
	}
	if activeOnlyField != nil {
		st.ActiveOnly = *activeOnlyField
	}
	completedOnlyField, err := identity(&pb.CompletedOnly)
	if err != nil {
		return nil, err
	}
	if completedOnlyField != nil {
		st.CompletedOnly = *completedOnlyField
	}
	expandTasksField, err := identity(&pb.ExpandTasks)
	if err != nil {
		return nil, err
	}
	if expandTasksField != nil {
		st.ExpandTasks = *expandTasksField
	}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}
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
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	runTypeField, err := identity(&pb.RunType)
	if err != nil {
		return nil, err
	}
	if runTypeField != nil {
		st.RunType = *runTypeField
	}
	startTimeFromField, err := identity(&pb.StartTimeFrom)
	if err != nil {
		return nil, err
	}
	if startTimeFromField != nil {
		st.StartTimeFrom = *startTimeFromField
	}
	startTimeToField, err := identity(&pb.StartTimeTo)
	if err != nil {
		return nil, err
	}
	if startTimeToField != nil {
		st.StartTimeTo = *startTimeToField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRunsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRunsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List of runs was retrieved successfully.
type ListRunsResponse struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	HasMore bool
	// A token that can be used to list the next page of runs (if applicable).
	NextPageToken string
	// A token that can be used to list the previous page of runs (if
	// applicable).
	PrevPageToken string
	// A list of runs, from most recently started to least. Only included in the
	// response if there are runs to list.
	Runs []BaseRun

	ForceSendFields []string
}

func listRunsResponseToPb(st *ListRunsResponse) (*listRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRunsResponsePb{}
	hasMorePb, err := identity(&st.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMorePb != nil {
		pb.HasMore = *hasMorePb
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

	var runsPb []baseRunPb
	for _, item := range st.Runs {
		itemPb, err := baseRunToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			runsPb = append(runsPb, *itemPb)
		}
	}
	pb.Runs = runsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRunsResponsePb struct {
	// If true, additional runs matching the provided filter are available for
	// listing.
	HasMore bool `json:"has_more,omitempty"`
	// A token that can be used to list the next page of runs (if applicable).
	NextPageToken string `json:"next_page_token,omitempty"`
	// A token that can be used to list the previous page of runs (if
	// applicable).
	PrevPageToken string `json:"prev_page_token,omitempty"`
	// A list of runs, from most recently started to least. Only included in the
	// response if there are runs to list.
	Runs []baseRunPb `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRunsResponseFromPb(pb *listRunsResponsePb) (*ListRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRunsResponse{}
	hasMoreField, err := identity(&pb.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMoreField != nil {
		st.HasMore = *hasMoreField
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

	var runsField []BaseRun
	for _, itemPb := range pb.Runs {
		item, err := baseRunFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			runsField = append(runsField, *item)
		}
	}
	st.Runs = runsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotebookOutput struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	Result string
	// Whether or not the result was truncated.
	Truncated bool

	ForceSendFields []string
}

func notebookOutputToPb(st *NotebookOutput) (*notebookOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookOutputPb{}
	resultPb, err := identity(&st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = *resultPb
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

func (st *NotebookOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notebookOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notebookOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NotebookOutput) MarshalJSON() ([]byte, error) {
	pb, err := notebookOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type notebookOutputPb struct {
	// The value passed to
	// [dbutils.notebook.exit()](/notebooks/notebook-workflows.html#notebook-workflows-exit).
	// Databricks restricts this API to return the first 5 MB of the value. For
	// a larger result, your job can store the results in a cloud storage
	// service. This field is absent if `dbutils.notebook.exit()` was never
	// called.
	Result string `json:"result,omitempty"`
	// Whether or not the result was truncated.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookOutputFromPb(pb *notebookOutputPb) (*NotebookOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookOutput{}
	resultField, err := identity(&pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = *resultField
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

func (st *notebookOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotebookTask struct {
	// Base parameters to be used for each run of this job. If the run is
	// initiated by a call to :method:jobs/run Now with parameters specified,
	// the two parameters maps are merged. If the same key is specified in
	// `base_parameters` and in `run-now`, the value from `run-now` is used. Use
	// [Task parameter variables] to set parameters containing information about
	// job runs.
	//
	// If the notebook takes a parameter that is not specified in the job’s
	// `base_parameters` or the `run-now` override parameters, the default value
	// from the notebook is used.
	//
	// Retrieve these parameters in a notebook using [dbutils.widgets.get].
	//
	// The JSON representation of this field cannot exceed 1MB.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-widgets
	BaseParameters map[string]string
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath string
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local Databricks workspace. When set
	// to `GIT`, the notebook will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`:
	// Notebook is located in Databricks workspace. * `GIT`: Notebook is located
	// in cloud Git provider.
	Source Source
	// Optional `warehouse_id` to run the notebook on a SQL warehouse. Classic
	// SQL warehouses are NOT supported, please use serverless or pro SQL
	// warehouses.
	//
	// Note that SQL warehouses only support SQL cells; if the notebook contains
	// non-SQL cells, the run will fail.
	WarehouseId string

	ForceSendFields []string
}

func notebookTaskToPb(st *NotebookTask) (*notebookTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookTaskPb{}

	baseParametersPb := map[string]string{}
	for k, v := range st.BaseParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			baseParametersPb[k] = *itemPb
		}
	}
	pb.BaseParameters = baseParametersPb

	notebookPathPb, err := identity(&st.NotebookPath)
	if err != nil {
		return nil, err
	}
	if notebookPathPb != nil {
		pb.NotebookPath = *notebookPathPb
	}

	sourcePb, err := identity(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *NotebookTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notebookTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notebookTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NotebookTask) MarshalJSON() ([]byte, error) {
	pb, err := notebookTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type notebookTaskPb struct {
	// Base parameters to be used for each run of this job. If the run is
	// initiated by a call to :method:jobs/run Now with parameters specified,
	// the two parameters maps are merged. If the same key is specified in
	// `base_parameters` and in `run-now`, the value from `run-now` is used. Use
	// [Task parameter variables] to set parameters containing information about
	// job runs.
	//
	// If the notebook takes a parameter that is not specified in the job’s
	// `base_parameters` or the `run-now` override parameters, the default value
	// from the notebook is used.
	//
	// Retrieve these parameters in a notebook using [dbutils.widgets.get].
	//
	// The JSON representation of this field cannot exceed 1MB.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-widgets
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
	// The path of the notebook to be run in the Databricks workspace or remote
	// repository. For notebooks stored in the Databricks workspace, the path
	// must be absolute and begin with a slash. For notebooks stored in a remote
	// repository, the path must be relative. This field is required.
	NotebookPath string `json:"notebook_path"`
	// Optional location type of the notebook. When set to `WORKSPACE`, the
	// notebook will be retrieved from the local Databricks workspace. When set
	// to `GIT`, the notebook will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise. * `WORKSPACE`:
	// Notebook is located in Databricks workspace. * `GIT`: Notebook is located
	// in cloud Git provider.
	Source Source `json:"source,omitempty"`
	// Optional `warehouse_id` to run the notebook on a SQL warehouse. Classic
	// SQL warehouses are NOT supported, please use serverless or pro SQL
	// warehouses.
	//
	// Note that SQL warehouses only support SQL cells; if the notebook contains
	// non-SQL cells, the run will fail.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookTaskFromPb(pb *notebookTaskPb) (*NotebookTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookTask{}

	baseParametersField := map[string]string{}
	for k, v := range pb.BaseParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			baseParametersField[k] = *item
		}
	}
	st.BaseParameters = baseParametersField
	notebookPathField, err := identity(&pb.NotebookPath)
	if err != nil {
		return nil, err
	}
	if notebookPathField != nil {
		st.NotebookPath = *notebookPathField
	}
	sourceField, err := identity(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notebookTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Stores the catalog name, schema name, and the output schema expiration time
// for the clean room run.
type OutputSchemaInfo struct {
	CatalogName string
	// The expiration time for the output schema as a Unix timestamp in
	// milliseconds.
	ExpirationTime int64

	SchemaName string

	ForceSendFields []string
}

func outputSchemaInfoToPb(st *OutputSchemaInfo) (*outputSchemaInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &outputSchemaInfoPb{}
	catalogNamePb, err := identity(&st.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	expirationTimePb, err := identity(&st.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimePb != nil {
		pb.ExpirationTime = *expirationTimePb
	}

	schemaNamePb, err := identity(&st.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *OutputSchemaInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &outputSchemaInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := outputSchemaInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OutputSchemaInfo) MarshalJSON() ([]byte, error) {
	pb, err := outputSchemaInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type outputSchemaInfoPb struct {
	CatalogName string `json:"catalog_name,omitempty"`
	// The expiration time for the output schema as a Unix timestamp in
	// milliseconds.
	ExpirationTime int64 `json:"expiration_time,omitempty"`

	SchemaName string `json:"schema_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func outputSchemaInfoFromPb(pb *outputSchemaInfoPb) (*OutputSchemaInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OutputSchemaInfo{}
	catalogNameField, err := identity(&pb.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	expirationTimeField, err := identity(&pb.ExpirationTime)
	if err != nil {
		return nil, err
	}
	if expirationTimeField != nil {
		st.ExpirationTime = *expirationTimeField
	}
	schemaNameField, err := identity(&pb.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *outputSchemaInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st outputSchemaInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PauseStatus string
type pauseStatusPb string

const PauseStatusPaused PauseStatus = `PAUSED`

const PauseStatusUnpaused PauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *PauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = PauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Type always returns PauseStatus to satisfy [pflag.Value] interface
func (f *PauseStatus) Type() string {
	return "PauseStatus"
}

func pauseStatusToPb(st *PauseStatus) (*pauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := pauseStatusPb(*st)
	return &pb, nil
}

func pauseStatusFromPb(pb *pauseStatusPb) (*PauseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := PauseStatus(*pb)
	return &st, nil
}

// PerformanceTarget defines how performant (lower latency) or cost efficient
// the execution of run on serverless compute should be. The performance mode on
// the job or pipeline should map to a performance setting that is passed to
// Cluster Manager (see cluster-common PerformanceTarget).
type PerformanceTarget string
type performanceTargetPb string

const PerformanceTargetPerformanceOptimized PerformanceTarget = `PERFORMANCE_OPTIMIZED`

const PerformanceTargetStandard PerformanceTarget = `STANDARD`

// String representation for [fmt.Print]
func (f *PerformanceTarget) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PerformanceTarget) Set(v string) error {
	switch v {
	case `PERFORMANCE_OPTIMIZED`, `STANDARD`:
		*f = PerformanceTarget(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PERFORMANCE_OPTIMIZED", "STANDARD"`, v)
	}
}

// Type always returns PerformanceTarget to satisfy [pflag.Value] interface
func (f *PerformanceTarget) Type() string {
	return "PerformanceTarget"
}

func performanceTargetToPb(st *PerformanceTarget) (*performanceTargetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := performanceTargetPb(*st)
	return &pb, nil
}

func performanceTargetFromPb(pb *performanceTargetPb) (*PerformanceTarget, error) {
	if pb == nil {
		return nil, nil
	}
	st := PerformanceTarget(*pb)
	return &st, nil
}

type PeriodicTriggerConfiguration struct {
	// The interval at which the trigger should run.
	Interval int
	// The unit of time for the interval.
	Unit PeriodicTriggerConfigurationTimeUnit
}

func periodicTriggerConfigurationToPb(st *PeriodicTriggerConfiguration) (*periodicTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &periodicTriggerConfigurationPb{}
	intervalPb, err := identity(&st.Interval)
	if err != nil {
		return nil, err
	}
	if intervalPb != nil {
		pb.Interval = *intervalPb
	}

	unitPb, err := identity(&st.Unit)
	if err != nil {
		return nil, err
	}
	if unitPb != nil {
		pb.Unit = *unitPb
	}

	return pb, nil
}

func (st *PeriodicTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &periodicTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := periodicTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PeriodicTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := periodicTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type periodicTriggerConfigurationPb struct {
	// The interval at which the trigger should run.
	Interval int `json:"interval"`
	// The unit of time for the interval.
	Unit PeriodicTriggerConfigurationTimeUnit `json:"unit"`
}

func periodicTriggerConfigurationFromPb(pb *periodicTriggerConfigurationPb) (*PeriodicTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PeriodicTriggerConfiguration{}
	intervalField, err := identity(&pb.Interval)
	if err != nil {
		return nil, err
	}
	if intervalField != nil {
		st.Interval = *intervalField
	}
	unitField, err := identity(&pb.Unit)
	if err != nil {
		return nil, err
	}
	if unitField != nil {
		st.Unit = *unitField
	}

	return st, nil
}

type PeriodicTriggerConfigurationTimeUnit string
type periodicTriggerConfigurationTimeUnitPb string

const PeriodicTriggerConfigurationTimeUnitDays PeriodicTriggerConfigurationTimeUnit = `DAYS`

const PeriodicTriggerConfigurationTimeUnitHours PeriodicTriggerConfigurationTimeUnit = `HOURS`

const PeriodicTriggerConfigurationTimeUnitWeeks PeriodicTriggerConfigurationTimeUnit = `WEEKS`

// String representation for [fmt.Print]
func (f *PeriodicTriggerConfigurationTimeUnit) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PeriodicTriggerConfigurationTimeUnit) Set(v string) error {
	switch v {
	case `DAYS`, `HOURS`, `WEEKS`:
		*f = PeriodicTriggerConfigurationTimeUnit(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAYS", "HOURS", "WEEKS"`, v)
	}
}

// Type always returns PeriodicTriggerConfigurationTimeUnit to satisfy [pflag.Value] interface
func (f *PeriodicTriggerConfigurationTimeUnit) Type() string {
	return "PeriodicTriggerConfigurationTimeUnit"
}

func periodicTriggerConfigurationTimeUnitToPb(st *PeriodicTriggerConfigurationTimeUnit) (*periodicTriggerConfigurationTimeUnitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := periodicTriggerConfigurationTimeUnitPb(*st)
	return &pb, nil
}

func periodicTriggerConfigurationTimeUnitFromPb(pb *periodicTriggerConfigurationTimeUnitPb) (*PeriodicTriggerConfigurationTimeUnit, error) {
	if pb == nil {
		return nil, nil
	}
	st := PeriodicTriggerConfigurationTimeUnit(*pb)
	return &st, nil
}

type PipelineParams struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool

	ForceSendFields []string
}

func pipelineParamsToPb(st *PipelineParams) (*pipelineParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineParamsPb{}
	fullRefreshPb, err := identity(&st.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshPb != nil {
		pb.FullRefresh = *fullRefreshPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PipelineParams) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineParamsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineParamsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineParams) MarshalJSON() ([]byte, error) {
	pb, err := pipelineParamsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type pipelineParamsPb struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineParamsFromPb(pb *pipelineParamsPb) (*PipelineParams, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineParams{}
	fullRefreshField, err := identity(&pb.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshField != nil {
		st.FullRefresh = *fullRefreshField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineParamsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineParamsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineTask struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool
	// The full name of the pipeline task to execute.
	PipelineId string

	ForceSendFields []string
}

func pipelineTaskToPb(st *PipelineTask) (*pipelineTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineTaskPb{}
	fullRefreshPb, err := identity(&st.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshPb != nil {
		pb.FullRefresh = *fullRefreshPb
	}

	pipelineIdPb, err := identity(&st.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdPb != nil {
		pb.PipelineId = *pipelineIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PipelineTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pipelineTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pipelineTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PipelineTask) MarshalJSON() ([]byte, error) {
	pb, err := pipelineTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type pipelineTaskPb struct {
	// If true, triggers a full refresh on the delta live table.
	FullRefresh bool `json:"full_refresh,omitempty"`
	// The full name of the pipeline task to execute.
	PipelineId string `json:"pipeline_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineTaskFromPb(pb *pipelineTaskPb) (*PipelineTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineTask{}
	fullRefreshField, err := identity(&pb.FullRefresh)
	if err != nil {
		return nil, err
	}
	if fullRefreshField != nil {
		st.FullRefresh = *fullRefreshField
	}
	pipelineIdField, err := identity(&pb.PipelineId)
	if err != nil {
		return nil, err
	}
	if pipelineIdField != nil {
		st.PipelineId = *pipelineIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PowerBiModel struct {
	// How the published Power BI model authenticates to Databricks
	AuthenticationMethod AuthenticationMethod
	// The name of the Power BI model
	ModelName string
	// Whether to overwrite existing Power BI models
	OverwriteExisting bool
	// The default storage mode of the Power BI model
	StorageMode StorageMode
	// The name of the Power BI workspace of the model
	WorkspaceName string

	ForceSendFields []string
}

func powerBiModelToPb(st *PowerBiModel) (*powerBiModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &powerBiModelPb{}
	authenticationMethodPb, err := identity(&st.AuthenticationMethod)
	if err != nil {
		return nil, err
	}
	if authenticationMethodPb != nil {
		pb.AuthenticationMethod = *authenticationMethodPb
	}

	modelNamePb, err := identity(&st.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNamePb != nil {
		pb.ModelName = *modelNamePb
	}

	overwriteExistingPb, err := identity(&st.OverwriteExisting)
	if err != nil {
		return nil, err
	}
	if overwriteExistingPb != nil {
		pb.OverwriteExisting = *overwriteExistingPb
	}

	storageModePb, err := identity(&st.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModePb != nil {
		pb.StorageMode = *storageModePb
	}

	workspaceNamePb, err := identity(&st.WorkspaceName)
	if err != nil {
		return nil, err
	}
	if workspaceNamePb != nil {
		pb.WorkspaceName = *workspaceNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PowerBiModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &powerBiModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := powerBiModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PowerBiModel) MarshalJSON() ([]byte, error) {
	pb, err := powerBiModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type powerBiModelPb struct {
	// How the published Power BI model authenticates to Databricks
	AuthenticationMethod AuthenticationMethod `json:"authentication_method,omitempty"`
	// The name of the Power BI model
	ModelName string `json:"model_name,omitempty"`
	// Whether to overwrite existing Power BI models
	OverwriteExisting bool `json:"overwrite_existing,omitempty"`
	// The default storage mode of the Power BI model
	StorageMode StorageMode `json:"storage_mode,omitempty"`
	// The name of the Power BI workspace of the model
	WorkspaceName string `json:"workspace_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func powerBiModelFromPb(pb *powerBiModelPb) (*PowerBiModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiModel{}
	authenticationMethodField, err := identity(&pb.AuthenticationMethod)
	if err != nil {
		return nil, err
	}
	if authenticationMethodField != nil {
		st.AuthenticationMethod = *authenticationMethodField
	}
	modelNameField, err := identity(&pb.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNameField != nil {
		st.ModelName = *modelNameField
	}
	overwriteExistingField, err := identity(&pb.OverwriteExisting)
	if err != nil {
		return nil, err
	}
	if overwriteExistingField != nil {
		st.OverwriteExisting = *overwriteExistingField
	}
	storageModeField, err := identity(&pb.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModeField != nil {
		st.StorageMode = *storageModeField
	}
	workspaceNameField, err := identity(&pb.WorkspaceName)
	if err != nil {
		return nil, err
	}
	if workspaceNameField != nil {
		st.WorkspaceName = *workspaceNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *powerBiModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st powerBiModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PowerBiTable struct {
	// The catalog name in Databricks
	Catalog string
	// The table name in Databricks
	Name string
	// The schema name in Databricks
	Schema string
	// The Power BI storage mode of the table
	StorageMode StorageMode

	ForceSendFields []string
}

func powerBiTableToPb(st *PowerBiTable) (*powerBiTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &powerBiTablePb{}
	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	storageModePb, err := identity(&st.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModePb != nil {
		pb.StorageMode = *storageModePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PowerBiTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &powerBiTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := powerBiTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PowerBiTable) MarshalJSON() ([]byte, error) {
	pb, err := powerBiTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type powerBiTablePb struct {
	// The catalog name in Databricks
	Catalog string `json:"catalog,omitempty"`
	// The table name in Databricks
	Name string `json:"name,omitempty"`
	// The schema name in Databricks
	Schema string `json:"schema,omitempty"`
	// The Power BI storage mode of the table
	StorageMode StorageMode `json:"storage_mode,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func powerBiTableFromPb(pb *powerBiTablePb) (*PowerBiTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiTable{}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	storageModeField, err := identity(&pb.StorageMode)
	if err != nil {
		return nil, err
	}
	if storageModeField != nil {
		st.StorageMode = *storageModeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *powerBiTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st powerBiTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PowerBiTask struct {
	// The resource name of the UC connection to authenticate from Databricks to
	// Power BI
	ConnectionResourceName string
	// The semantic model to update
	PowerBiModel *PowerBiModel
	// Whether the model should be refreshed after the update
	RefreshAfterUpdate bool
	// The tables to be exported to Power BI
	Tables []PowerBiTable
	// The SQL warehouse ID to use as the Power BI data source
	WarehouseId string

	ForceSendFields []string
}

func powerBiTaskToPb(st *PowerBiTask) (*powerBiTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &powerBiTaskPb{}
	connectionResourceNamePb, err := identity(&st.ConnectionResourceName)
	if err != nil {
		return nil, err
	}
	if connectionResourceNamePb != nil {
		pb.ConnectionResourceName = *connectionResourceNamePb
	}

	powerBiModelPb, err := powerBiModelToPb(st.PowerBiModel)
	if err != nil {
		return nil, err
	}
	if powerBiModelPb != nil {
		pb.PowerBiModel = powerBiModelPb
	}

	refreshAfterUpdatePb, err := identity(&st.RefreshAfterUpdate)
	if err != nil {
		return nil, err
	}
	if refreshAfterUpdatePb != nil {
		pb.RefreshAfterUpdate = *refreshAfterUpdatePb
	}

	var tablesPb []powerBiTablePb
	for _, item := range st.Tables {
		itemPb, err := powerBiTableToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tablesPb = append(tablesPb, *itemPb)
		}
	}
	pb.Tables = tablesPb

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PowerBiTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &powerBiTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := powerBiTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PowerBiTask) MarshalJSON() ([]byte, error) {
	pb, err := powerBiTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type powerBiTaskPb struct {
	// The resource name of the UC connection to authenticate from Databricks to
	// Power BI
	ConnectionResourceName string `json:"connection_resource_name,omitempty"`
	// The semantic model to update
	PowerBiModel *powerBiModelPb `json:"power_bi_model,omitempty"`
	// Whether the model should be refreshed after the update
	RefreshAfterUpdate bool `json:"refresh_after_update,omitempty"`
	// The tables to be exported to Power BI
	Tables []powerBiTablePb `json:"tables,omitempty"`
	// The SQL warehouse ID to use as the Power BI data source
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func powerBiTaskFromPb(pb *powerBiTaskPb) (*PowerBiTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PowerBiTask{}
	connectionResourceNameField, err := identity(&pb.ConnectionResourceName)
	if err != nil {
		return nil, err
	}
	if connectionResourceNameField != nil {
		st.ConnectionResourceName = *connectionResourceNameField
	}
	powerBiModelField, err := powerBiModelFromPb(pb.PowerBiModel)
	if err != nil {
		return nil, err
	}
	if powerBiModelField != nil {
		st.PowerBiModel = powerBiModelField
	}
	refreshAfterUpdateField, err := identity(&pb.RefreshAfterUpdate)
	if err != nil {
		return nil, err
	}
	if refreshAfterUpdateField != nil {
		st.RefreshAfterUpdate = *refreshAfterUpdateField
	}

	var tablesField []PowerBiTable
	for _, itemPb := range pb.Tables {
		item, err := powerBiTableFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tablesField = append(tablesField, *item)
		}
	}
	st.Tables = tablesField
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *powerBiTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st powerBiTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PythonWheelTask struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	EntryPoint string
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	NamedParameters map[string]string
	// Name of the package to execute
	PackageName string
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	Parameters []string
}

func pythonWheelTaskToPb(st *PythonWheelTask) (*pythonWheelTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pythonWheelTaskPb{}
	entryPointPb, err := identity(&st.EntryPoint)
	if err != nil {
		return nil, err
	}
	if entryPointPb != nil {
		pb.EntryPoint = *entryPointPb
	}

	namedParametersPb := map[string]string{}
	for k, v := range st.NamedParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			namedParametersPb[k] = *itemPb
		}
	}
	pb.NamedParameters = namedParametersPb

	packageNamePb, err := identity(&st.PackageName)
	if err != nil {
		return nil, err
	}
	if packageNamePb != nil {
		pb.PackageName = *packageNamePb
	}

	var parametersPb []string
	for _, item := range st.Parameters {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	return pb, nil
}

func (st *PythonWheelTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pythonWheelTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pythonWheelTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PythonWheelTask) MarshalJSON() ([]byte, error) {
	pb, err := pythonWheelTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type pythonWheelTaskPb struct {
	// Named entry point to use, if it does not exist in the metadata of the
	// package it executes the function from the package directly using
	// `$packageName.$entryPoint()`
	EntryPoint string `json:"entry_point"`
	// Command-line parameters passed to Python wheel task in the form of
	// `["--name=task", "--data=dbfs:/path/to/data.json"]`. Leave it empty if
	// `parameters` is not null.
	NamedParameters map[string]string `json:"named_parameters,omitempty"`
	// Name of the package to execute
	PackageName string `json:"package_name"`
	// Command-line parameters passed to Python wheel task. Leave it empty if
	// `named_parameters` is not null.
	Parameters []string `json:"parameters,omitempty"`
}

func pythonWheelTaskFromPb(pb *pythonWheelTaskPb) (*PythonWheelTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PythonWheelTask{}
	entryPointField, err := identity(&pb.EntryPoint)
	if err != nil {
		return nil, err
	}
	if entryPointField != nil {
		st.EntryPoint = *entryPointField
	}

	namedParametersField := map[string]string{}
	for k, v := range pb.NamedParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			namedParametersField[k] = *item
		}
	}
	st.NamedParameters = namedParametersField
	packageNameField, err := identity(&pb.PackageName)
	if err != nil {
		return nil, err
	}
	if packageNameField != nil {
		st.PackageName = *packageNameField
	}

	var parametersField []string
	for _, itemPb := range pb.Parameters {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField

	return st, nil
}

type QueueDetails struct {
	// The reason for queuing the run. * `ACTIVE_RUNS_LIMIT_REACHED`: The run
	// was queued due to reaching the workspace limit of active task runs. *
	// `MAX_CONCURRENT_RUNS_REACHED`: The run was queued due to reaching the
	// per-job limit of concurrent job runs. *
	// `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`: The run was queued due to reaching
	// the workspace limit of active run job tasks.
	Code QueueDetailsCodeCode
	// A descriptive message with the queuing details. This field is
	// unstructured, and its exact format is subject to change.
	Message string

	ForceSendFields []string
}

func queueDetailsToPb(st *QueueDetails) (*queueDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queueDetailsPb{}
	codePb, err := identity(&st.Code)
	if err != nil {
		return nil, err
	}
	if codePb != nil {
		pb.Code = *codePb
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

func (st *QueueDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queueDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queueDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueueDetails) MarshalJSON() ([]byte, error) {
	pb, err := queueDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type queueDetailsPb struct {
	// The reason for queuing the run. * `ACTIVE_RUNS_LIMIT_REACHED`: The run
	// was queued due to reaching the workspace limit of active task runs. *
	// `MAX_CONCURRENT_RUNS_REACHED`: The run was queued due to reaching the
	// per-job limit of concurrent job runs. *
	// `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`: The run was queued due to reaching
	// the workspace limit of active run job tasks.
	Code QueueDetailsCodeCode `json:"code,omitempty"`
	// A descriptive message with the queuing details. This field is
	// unstructured, and its exact format is subject to change.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queueDetailsFromPb(pb *queueDetailsPb) (*QueueDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueueDetails{}
	codeField, err := identity(&pb.Code)
	if err != nil {
		return nil, err
	}
	if codeField != nil {
		st.Code = *codeField
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

func (st *queueDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queueDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The reason for queuing the run. * `ACTIVE_RUNS_LIMIT_REACHED`: The run was
// queued due to reaching the workspace limit of active task runs. *
// `MAX_CONCURRENT_RUNS_REACHED`: The run was queued due to reaching the per-job
// limit of concurrent job runs. * `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`: The run
// was queued due to reaching the workspace limit of active run job tasks.
type QueueDetailsCodeCode string
type queueDetailsCodeCodePb string

// The run was queued due to reaching the workspace limit of active task runs.
const QueueDetailsCodeCodeActiveRunsLimitReached QueueDetailsCodeCode = `ACTIVE_RUNS_LIMIT_REACHED`

// The run was queued due to reaching the workspace limit of active run job
// tasks.
const QueueDetailsCodeCodeActiveRunJobTasksLimitReached QueueDetailsCodeCode = `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`

// The run was queued due to reaching the per-job limit of concurrent job runs.
const QueueDetailsCodeCodeMaxConcurrentRunsReached QueueDetailsCodeCode = `MAX_CONCURRENT_RUNS_REACHED`

// String representation for [fmt.Print]
func (f *QueueDetailsCodeCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueueDetailsCodeCode) Set(v string) error {
	switch v {
	case `ACTIVE_RUNS_LIMIT_REACHED`, `ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED`, `MAX_CONCURRENT_RUNS_REACHED`:
		*f = QueueDetailsCodeCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_RUNS_LIMIT_REACHED", "ACTIVE_RUN_JOB_TASKS_LIMIT_REACHED", "MAX_CONCURRENT_RUNS_REACHED"`, v)
	}
}

// Type always returns QueueDetailsCodeCode to satisfy [pflag.Value] interface
func (f *QueueDetailsCodeCode) Type() string {
	return "QueueDetailsCodeCode"
}

func queueDetailsCodeCodeToPb(st *QueueDetailsCodeCode) (*queueDetailsCodeCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := queueDetailsCodeCodePb(*st)
	return &pb, nil
}

func queueDetailsCodeCodeFromPb(pb *queueDetailsCodeCodePb) (*QueueDetailsCodeCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueueDetailsCodeCode(*pb)
	return &st, nil
}

type QueueSettings struct {
	// If true, enable queueing for the job. This is a required field.
	Enabled bool
}

func queueSettingsToPb(st *QueueSettings) (*queueSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queueSettingsPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	return pb, nil
}

func (st *QueueSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queueSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queueSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueueSettings) MarshalJSON() ([]byte, error) {
	pb, err := queueSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type queueSettingsPb struct {
	// If true, enable queueing for the job. This is a required field.
	Enabled bool `json:"enabled"`
}

func queueSettingsFromPb(pb *queueSettingsPb) (*QueueSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueueSettings{}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}

	return st, nil
}

type RepairHistoryItem struct {
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget
	// The end time of the (repaired) run.
	EndTime int64
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id int64
	// The start time of the (repaired) run.
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	State *RunState
	// The current status of the run
	Status *RunStatus
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	TaskRunIds []int64
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	Type RepairHistoryItemType

	ForceSendFields []string
}

func repairHistoryItemToPb(st *RepairHistoryItem) (*repairHistoryItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repairHistoryItemPb{}
	effectivePerformanceTargetPb, err := identity(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}

	endTimePb, err := identity(&st.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimePb != nil {
		pb.EndTime = *endTimePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	statePb, err := runStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	statusPb, err := runStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	var taskRunIdsPb []int64
	for _, item := range st.TaskRunIds {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			taskRunIdsPb = append(taskRunIdsPb, *itemPb)
		}
	}
	pb.TaskRunIds = taskRunIdsPb

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

func (st *RepairHistoryItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repairHistoryItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repairHistoryItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepairHistoryItem) MarshalJSON() ([]byte, error) {
	pb, err := repairHistoryItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type repairHistoryItemPb struct {
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// The end time of the (repaired) run.
	EndTime int64 `json:"end_time,omitempty"`
	// The ID of the repair. Only returned for the items that represent a repair
	// in `repair_history`.
	Id int64 `json:"id,omitempty"`
	// The start time of the (repaired) run.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *runStatePb `json:"state,omitempty"`
	// The current status of the run
	Status *runStatusPb `json:"status,omitempty"`
	// The run IDs of the task runs that ran as part of this repair history
	// item.
	TaskRunIds []int64 `json:"task_run_ids,omitempty"`
	// The repair history item type. Indicates whether a run is the original run
	// or a repair run.
	Type RepairHistoryItemType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repairHistoryItemFromPb(pb *repairHistoryItemPb) (*RepairHistoryItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairHistoryItem{}
	effectivePerformanceTargetField, err := identity(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	endTimeField, err := identity(&pb.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimeField != nil {
		st.EndTime = *endTimeField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}
	stateField, err := runStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := runStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	var taskRunIdsField []int64
	for _, itemPb := range pb.TaskRunIds {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			taskRunIdsField = append(taskRunIdsField, *item)
		}
	}
	st.TaskRunIds = taskRunIdsField
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

func (st *repairHistoryItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repairHistoryItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The repair history item type. Indicates whether a run is the original run or
// a repair run.
type RepairHistoryItemType string
type repairHistoryItemTypePb string

const RepairHistoryItemTypeOriginal RepairHistoryItemType = `ORIGINAL`

const RepairHistoryItemTypeRepair RepairHistoryItemType = `REPAIR`

// String representation for [fmt.Print]
func (f *RepairHistoryItemType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RepairHistoryItemType) Set(v string) error {
	switch v {
	case `ORIGINAL`, `REPAIR`:
		*f = RepairHistoryItemType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ORIGINAL", "REPAIR"`, v)
	}
}

// Type always returns RepairHistoryItemType to satisfy [pflag.Value] interface
func (f *RepairHistoryItemType) Type() string {
	return "RepairHistoryItemType"
}

func repairHistoryItemTypeToPb(st *RepairHistoryItemType) (*repairHistoryItemTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := repairHistoryItemTypePb(*st)
	return &pb, nil
}

func repairHistoryItemTypeFromPb(pb *repairHistoryItemTypePb) (*RepairHistoryItemType, error) {
	if pb == nil {
		return nil, nil
	}
	st := RepairHistoryItemType(*pb)
	return &st, nil
}

type RepairRun struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId int64
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams

	PythonNamedParams map[string]string
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	RerunAllFailedTasks bool
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	RerunDependentTasks bool
	// The task keys of the task runs to repair.
	RerunTasks []string
	// The job run ID of the run to repair. The run must not be in progress.
	RunId int64
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string

	ForceSendFields []string
}

func repairRunToPb(st *RepairRun) (*repairRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repairRunPb{}

	var dbtCommandsPb []string
	for _, item := range st.DbtCommands {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dbtCommandsPb = append(dbtCommandsPb, *itemPb)
		}
	}
	pb.DbtCommands = dbtCommandsPb

	var jarParamsPb []string
	for _, item := range st.JarParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jarParamsPb = append(jarParamsPb, *itemPb)
		}
	}
	pb.JarParams = jarParamsPb

	jobParametersPb := map[string]string{}
	for k, v := range st.JobParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb[k] = *itemPb
		}
	}
	pb.JobParameters = jobParametersPb

	latestRepairIdPb, err := identity(&st.LatestRepairId)
	if err != nil {
		return nil, err
	}
	if latestRepairIdPb != nil {
		pb.LatestRepairId = *latestRepairIdPb
	}

	notebookParamsPb := map[string]string{}
	for k, v := range st.NotebookParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebookParamsPb[k] = *itemPb
		}
	}
	pb.NotebookParams = notebookParamsPb

	performanceTargetPb, err := identity(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}

	pipelineParamsPb, err := pipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}

	pythonNamedParamsPb := map[string]string{}
	for k, v := range st.PythonNamedParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonNamedParamsPb[k] = *itemPb
		}
	}
	pb.PythonNamedParams = pythonNamedParamsPb

	var pythonParamsPb []string
	for _, item := range st.PythonParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonParamsPb = append(pythonParamsPb, *itemPb)
		}
	}
	pb.PythonParams = pythonParamsPb

	rerunAllFailedTasksPb, err := identity(&st.RerunAllFailedTasks)
	if err != nil {
		return nil, err
	}
	if rerunAllFailedTasksPb != nil {
		pb.RerunAllFailedTasks = *rerunAllFailedTasksPb
	}

	rerunDependentTasksPb, err := identity(&st.RerunDependentTasks)
	if err != nil {
		return nil, err
	}
	if rerunDependentTasksPb != nil {
		pb.RerunDependentTasks = *rerunDependentTasksPb
	}

	var rerunTasksPb []string
	for _, item := range st.RerunTasks {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rerunTasksPb = append(rerunTasksPb, *itemPb)
		}
	}
	pb.RerunTasks = rerunTasksPb

	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	var sparkSubmitParamsPb []string
	for _, item := range st.SparkSubmitParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkSubmitParamsPb = append(sparkSubmitParamsPb, *itemPb)
		}
	}
	pb.SparkSubmitParams = sparkSubmitParamsPb

	sqlParamsPb := map[string]string{}
	for k, v := range st.SqlParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlParamsPb[k] = *itemPb
		}
	}
	pb.SqlParams = sqlParamsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RepairRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repairRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repairRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepairRun) MarshalJSON() ([]byte, error) {
	pb, err := repairRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type repairRunPb struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string `json:"jar_params,omitempty"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	// The ID of the latest repair. This parameter is not required when
	// repairing a run for the first time, but must be provided on subsequent
	// requests to repair the same run.
	LatestRepairId int64 `json:"latest_repair_id,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *pipelineParamsPb `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string `json:"python_params,omitempty"`
	// If true, repair all failed tasks. Only one of `rerun_tasks` or
	// `rerun_all_failed_tasks` can be used.
	RerunAllFailedTasks bool `json:"rerun_all_failed_tasks,omitempty"`
	// If true, repair all tasks that depend on the tasks in `rerun_tasks`, even
	// if they were previously successful. Can be also used in combination with
	// `rerun_all_failed_tasks`.
	RerunDependentTasks bool `json:"rerun_dependent_tasks,omitempty"`
	// The task keys of the task runs to repair.
	RerunTasks []string `json:"rerun_tasks,omitempty"`
	// The job run ID of the run to repair. The run must not be in progress.
	RunId int64 `json:"run_id"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repairRunFromPb(pb *repairRunPb) (*RepairRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairRun{}

	var dbtCommandsField []string
	for _, itemPb := range pb.DbtCommands {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dbtCommandsField = append(dbtCommandsField, *item)
		}
	}
	st.DbtCommands = dbtCommandsField

	var jarParamsField []string
	for _, itemPb := range pb.JarParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jarParamsField = append(jarParamsField, *item)
		}
	}
	st.JarParams = jarParamsField

	jobParametersField := map[string]string{}
	for k, v := range pb.JobParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField[k] = *item
		}
	}
	st.JobParameters = jobParametersField
	latestRepairIdField, err := identity(&pb.LatestRepairId)
	if err != nil {
		return nil, err
	}
	if latestRepairIdField != nil {
		st.LatestRepairId = *latestRepairIdField
	}

	notebookParamsField := map[string]string{}
	for k, v := range pb.NotebookParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notebookParamsField[k] = *item
		}
	}
	st.NotebookParams = notebookParamsField
	performanceTargetField, err := identity(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	pipelineParamsField, err := pipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}

	pythonNamedParamsField := map[string]string{}
	for k, v := range pb.PythonNamedParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonNamedParamsField[k] = *item
		}
	}
	st.PythonNamedParams = pythonNamedParamsField

	var pythonParamsField []string
	for _, itemPb := range pb.PythonParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonParamsField = append(pythonParamsField, *item)
		}
	}
	st.PythonParams = pythonParamsField
	rerunAllFailedTasksField, err := identity(&pb.RerunAllFailedTasks)
	if err != nil {
		return nil, err
	}
	if rerunAllFailedTasksField != nil {
		st.RerunAllFailedTasks = *rerunAllFailedTasksField
	}
	rerunDependentTasksField, err := identity(&pb.RerunDependentTasks)
	if err != nil {
		return nil, err
	}
	if rerunDependentTasksField != nil {
		st.RerunDependentTasks = *rerunDependentTasksField
	}

	var rerunTasksField []string
	for _, itemPb := range pb.RerunTasks {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rerunTasksField = append(rerunTasksField, *item)
		}
	}
	st.RerunTasks = rerunTasksField
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	var sparkSubmitParamsField []string
	for _, itemPb := range pb.SparkSubmitParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkSubmitParamsField = append(sparkSubmitParamsField, *item)
		}
	}
	st.SparkSubmitParams = sparkSubmitParamsField

	sqlParamsField := map[string]string{}
	for k, v := range pb.SqlParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlParamsField[k] = *item
		}
	}
	st.SqlParams = sqlParamsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repairRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repairRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Run repair was initiated.
type RepairRunResponse struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId int64

	ForceSendFields []string
}

func repairRunResponseToPb(st *RepairRunResponse) (*repairRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repairRunResponsePb{}
	repairIdPb, err := identity(&st.RepairId)
	if err != nil {
		return nil, err
	}
	if repairIdPb != nil {
		pb.RepairId = *repairIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RepairRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repairRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repairRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepairRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := repairRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type repairRunResponsePb struct {
	// The ID of the repair. Must be provided in subsequent repairs using the
	// `latest_repair_id` field to ensure sequential repairs.
	RepairId int64 `json:"repair_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func repairRunResponseFromPb(pb *repairRunResponsePb) (*RepairRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepairRunResponse{}
	repairIdField, err := identity(&pb.RepairId)
	if err != nil {
		return nil, err
	}
	if repairIdField != nil {
		st.RepairId = *repairIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *repairRunResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st repairRunResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResetJob struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId int64
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	NewSettings JobSettings
}

func resetJobToPb(st *ResetJob) (*resetJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resetJobPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	newSettingsPb, err := jobSettingsToPb(&st.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsPb != nil {
		pb.NewSettings = *newSettingsPb
	}

	return pb, nil
}

func (st *ResetJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resetJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resetJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResetJob) MarshalJSON() ([]byte, error) {
	pb, err := resetJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resetJobPb struct {
	// The canonical identifier of the job to reset. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings of the job. These settings completely replace the old
	// settings.
	//
	// Changes to the field `JobBaseSettings.timeout_seconds` are applied to
	// active runs. Changes to other fields are applied to future runs only.
	NewSettings jobSettingsPb `json:"new_settings"`
}

func resetJobFromPb(pb *resetJobPb) (*ResetJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResetJob{}
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}
	newSettingsField, err := jobSettingsFromPb(&pb.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsField != nil {
		st.NewSettings = *newSettingsField
	}

	return st, nil
}

type ResetResponse struct {
}

func resetResponseToPb(st *ResetResponse) (*resetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resetResponsePb{}

	return pb, nil
}

func (st *ResetResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resetResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resetResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResetResponse) MarshalJSON() ([]byte, error) {
	pb, err := resetResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resetResponsePb struct {
}

func resetResponseFromPb(pb *resetResponsePb) (*ResetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResetResponse{}

	return st, nil
}

type ResolvedConditionTaskValues struct {
	Left string

	Right string

	ForceSendFields []string
}

func resolvedConditionTaskValuesToPb(st *ResolvedConditionTaskValues) (*resolvedConditionTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedConditionTaskValuesPb{}
	leftPb, err := identity(&st.Left)
	if err != nil {
		return nil, err
	}
	if leftPb != nil {
		pb.Left = *leftPb
	}

	rightPb, err := identity(&st.Right)
	if err != nil {
		return nil, err
	}
	if rightPb != nil {
		pb.Right = *rightPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ResolvedConditionTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedConditionTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedConditionTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedConditionTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedConditionTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedConditionTaskValuesPb struct {
	Left string `json:"left,omitempty"`

	Right string `json:"right,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resolvedConditionTaskValuesFromPb(pb *resolvedConditionTaskValuesPb) (*ResolvedConditionTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedConditionTaskValues{}
	leftField, err := identity(&pb.Left)
	if err != nil {
		return nil, err
	}
	if leftField != nil {
		st.Left = *leftField
	}
	rightField, err := identity(&pb.Right)
	if err != nil {
		return nil, err
	}
	if rightField != nil {
		st.Right = *rightField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resolvedConditionTaskValuesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resolvedConditionTaskValuesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResolvedDbtTaskValues struct {
	Commands []string
}

func resolvedDbtTaskValuesToPb(st *ResolvedDbtTaskValues) (*resolvedDbtTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedDbtTaskValuesPb{}

	var commandsPb []string
	for _, item := range st.Commands {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			commandsPb = append(commandsPb, *itemPb)
		}
	}
	pb.Commands = commandsPb

	return pb, nil
}

func (st *ResolvedDbtTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedDbtTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedDbtTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedDbtTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedDbtTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedDbtTaskValuesPb struct {
	Commands []string `json:"commands,omitempty"`
}

func resolvedDbtTaskValuesFromPb(pb *resolvedDbtTaskValuesPb) (*ResolvedDbtTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedDbtTaskValues{}

	var commandsField []string
	for _, itemPb := range pb.Commands {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			commandsField = append(commandsField, *item)
		}
	}
	st.Commands = commandsField

	return st, nil
}

type ResolvedNotebookTaskValues struct {
	BaseParameters map[string]string
}

func resolvedNotebookTaskValuesToPb(st *ResolvedNotebookTaskValues) (*resolvedNotebookTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedNotebookTaskValuesPb{}

	baseParametersPb := map[string]string{}
	for k, v := range st.BaseParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			baseParametersPb[k] = *itemPb
		}
	}
	pb.BaseParameters = baseParametersPb

	return pb, nil
}

func (st *ResolvedNotebookTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedNotebookTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedNotebookTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedNotebookTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedNotebookTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedNotebookTaskValuesPb struct {
	BaseParameters map[string]string `json:"base_parameters,omitempty"`
}

func resolvedNotebookTaskValuesFromPb(pb *resolvedNotebookTaskValuesPb) (*ResolvedNotebookTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedNotebookTaskValues{}

	baseParametersField := map[string]string{}
	for k, v := range pb.BaseParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			baseParametersField[k] = *item
		}
	}
	st.BaseParameters = baseParametersField

	return st, nil
}

type ResolvedParamPairValues struct {
	Parameters map[string]string
}

func resolvedParamPairValuesToPb(st *ResolvedParamPairValues) (*resolvedParamPairValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedParamPairValuesPb{}

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

	return pb, nil
}

func (st *ResolvedParamPairValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedParamPairValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedParamPairValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedParamPairValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedParamPairValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedParamPairValuesPb struct {
	Parameters map[string]string `json:"parameters,omitempty"`
}

func resolvedParamPairValuesFromPb(pb *resolvedParamPairValuesPb) (*ResolvedParamPairValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedParamPairValues{}

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

	return st, nil
}

type ResolvedPythonWheelTaskValues struct {
	NamedParameters map[string]string

	Parameters []string
}

func resolvedPythonWheelTaskValuesToPb(st *ResolvedPythonWheelTaskValues) (*resolvedPythonWheelTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedPythonWheelTaskValuesPb{}

	namedParametersPb := map[string]string{}
	for k, v := range st.NamedParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			namedParametersPb[k] = *itemPb
		}
	}
	pb.NamedParameters = namedParametersPb

	var parametersPb []string
	for _, item := range st.Parameters {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	return pb, nil
}

func (st *ResolvedPythonWheelTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedPythonWheelTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedPythonWheelTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedPythonWheelTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedPythonWheelTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedPythonWheelTaskValuesPb struct {
	NamedParameters map[string]string `json:"named_parameters,omitempty"`

	Parameters []string `json:"parameters,omitempty"`
}

func resolvedPythonWheelTaskValuesFromPb(pb *resolvedPythonWheelTaskValuesPb) (*ResolvedPythonWheelTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedPythonWheelTaskValues{}

	namedParametersField := map[string]string{}
	for k, v := range pb.NamedParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			namedParametersField[k] = *item
		}
	}
	st.NamedParameters = namedParametersField

	var parametersField []string
	for _, itemPb := range pb.Parameters {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField

	return st, nil
}

type ResolvedRunJobTaskValues struct {
	JobParameters map[string]string

	Parameters map[string]string
}

func resolvedRunJobTaskValuesToPb(st *ResolvedRunJobTaskValues) (*resolvedRunJobTaskValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedRunJobTaskValuesPb{}

	jobParametersPb := map[string]string{}
	for k, v := range st.JobParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb[k] = *itemPb
		}
	}
	pb.JobParameters = jobParametersPb

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

	return pb, nil
}

func (st *ResolvedRunJobTaskValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedRunJobTaskValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedRunJobTaskValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedRunJobTaskValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedRunJobTaskValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedRunJobTaskValuesPb struct {
	JobParameters map[string]string `json:"job_parameters,omitempty"`

	Parameters map[string]string `json:"parameters,omitempty"`
}

func resolvedRunJobTaskValuesFromPb(pb *resolvedRunJobTaskValuesPb) (*ResolvedRunJobTaskValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedRunJobTaskValues{}

	jobParametersField := map[string]string{}
	for k, v := range pb.JobParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField[k] = *item
		}
	}
	st.JobParameters = jobParametersField

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

	return st, nil
}

type ResolvedStringParamsValues struct {
	Parameters []string
}

func resolvedStringParamsValuesToPb(st *ResolvedStringParamsValues) (*resolvedStringParamsValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedStringParamsValuesPb{}

	var parametersPb []string
	for _, item := range st.Parameters {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	return pb, nil
}

func (st *ResolvedStringParamsValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedStringParamsValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedStringParamsValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedStringParamsValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedStringParamsValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedStringParamsValuesPb struct {
	Parameters []string `json:"parameters,omitempty"`
}

func resolvedStringParamsValuesFromPb(pb *resolvedStringParamsValuesPb) (*ResolvedStringParamsValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedStringParamsValues{}

	var parametersField []string
	for _, itemPb := range pb.Parameters {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField

	return st, nil
}

type ResolvedValues struct {
	ConditionTask *ResolvedConditionTaskValues

	DbtTask *ResolvedDbtTaskValues

	NotebookTask *ResolvedNotebookTaskValues

	PythonWheelTask *ResolvedPythonWheelTaskValues

	RunJobTask *ResolvedRunJobTaskValues

	SimulationTask *ResolvedParamPairValues

	SparkJarTask *ResolvedStringParamsValues

	SparkPythonTask *ResolvedStringParamsValues

	SparkSubmitTask *ResolvedStringParamsValues

	SqlTask *ResolvedParamPairValues
}

func resolvedValuesToPb(st *ResolvedValues) (*resolvedValuesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resolvedValuesPb{}
	conditionTaskPb, err := resolvedConditionTaskValuesToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}

	dbtTaskPb, err := resolvedDbtTaskValuesToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}

	notebookTaskPb, err := resolvedNotebookTaskValuesToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}

	pythonWheelTaskPb, err := resolvedPythonWheelTaskValuesToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}

	runJobTaskPb, err := resolvedRunJobTaskValuesToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}

	simulationTaskPb, err := resolvedParamPairValuesToPb(st.SimulationTask)
	if err != nil {
		return nil, err
	}
	if simulationTaskPb != nil {
		pb.SimulationTask = simulationTaskPb
	}

	sparkJarTaskPb, err := resolvedStringParamsValuesToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}

	sparkPythonTaskPb, err := resolvedStringParamsValuesToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}

	sparkSubmitTaskPb, err := resolvedStringParamsValuesToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}

	sqlTaskPb, err := resolvedParamPairValuesToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}

	return pb, nil
}

func (st *ResolvedValues) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resolvedValuesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resolvedValuesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResolvedValues) MarshalJSON() ([]byte, error) {
	pb, err := resolvedValuesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type resolvedValuesPb struct {
	ConditionTask *resolvedConditionTaskValuesPb `json:"condition_task,omitempty"`

	DbtTask *resolvedDbtTaskValuesPb `json:"dbt_task,omitempty"`

	NotebookTask *resolvedNotebookTaskValuesPb `json:"notebook_task,omitempty"`

	PythonWheelTask *resolvedPythonWheelTaskValuesPb `json:"python_wheel_task,omitempty"`

	RunJobTask *resolvedRunJobTaskValuesPb `json:"run_job_task,omitempty"`

	SimulationTask *resolvedParamPairValuesPb `json:"simulation_task,omitempty"`

	SparkJarTask *resolvedStringParamsValuesPb `json:"spark_jar_task,omitempty"`

	SparkPythonTask *resolvedStringParamsValuesPb `json:"spark_python_task,omitempty"`

	SparkSubmitTask *resolvedStringParamsValuesPb `json:"spark_submit_task,omitempty"`

	SqlTask *resolvedParamPairValuesPb `json:"sql_task,omitempty"`
}

func resolvedValuesFromPb(pb *resolvedValuesPb) (*ResolvedValues, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResolvedValues{}
	conditionTaskField, err := resolvedConditionTaskValuesFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dbtTaskField, err := resolvedDbtTaskValuesFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}
	notebookTaskField, err := resolvedNotebookTaskValuesFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	pythonWheelTaskField, err := resolvedPythonWheelTaskValuesFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	runJobTaskField, err := resolvedRunJobTaskValuesFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	simulationTaskField, err := resolvedParamPairValuesFromPb(pb.SimulationTask)
	if err != nil {
		return nil, err
	}
	if simulationTaskField != nil {
		st.SimulationTask = simulationTaskField
	}
	sparkJarTaskField, err := resolvedStringParamsValuesFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := resolvedStringParamsValuesFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := resolvedStringParamsValuesFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := resolvedParamPairValuesFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}

	return st, nil
}

// Run was retrieved successfully
type Run struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *ClusterSpec
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string
	// Description of the run
	Description string
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	HasMore bool
	// Only populated by for-each iterations. The parent for-each task is
	// located in tasks array.
	Iterations []RunTask
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	JobClusters []JobCluster
	// The canonical identifier of the job that contains this run.
	JobId int64
	// Job-level parameters used in the run
	JobParameters []JobParameter
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId int64
	// A token that can be used to list the next page of array properties.
	NextPageToken string
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64
	// The parameters used for this run.
	OverridingParameters *RunParameters
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64
	// The repair history of the run.
	RepairHistory []RepairHistoryItem
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string
	// The URL to the detail page of the run.
	RunPageUrl string
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *CronSchedule
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	State *RunState
	// The current status of the run
	Status *RunStatus
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	Tasks []RunTask
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	Trigger TriggerType
	// Additional details about what triggered the run
	TriggerInfo *TriggerInfo

	ForceSendFields []string
}

func runToPb(st *Run) (*runPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runPb{}
	attemptNumberPb, err := identity(&st.AttemptNumber)
	if err != nil {
		return nil, err
	}
	if attemptNumberPb != nil {
		pb.AttemptNumber = *attemptNumberPb
	}

	cleanupDurationPb, err := identity(&st.CleanupDuration)
	if err != nil {
		return nil, err
	}
	if cleanupDurationPb != nil {
		pb.CleanupDuration = *cleanupDurationPb
	}

	clusterInstancePb, err := clusterInstanceToPb(st.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstancePb != nil {
		pb.ClusterInstance = clusterInstancePb
	}

	clusterSpecPb, err := clusterSpecToPb(st.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecPb != nil {
		pb.ClusterSpec = clusterSpecPb
	}

	creatorUserNamePb, err := identity(&st.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNamePb != nil {
		pb.CreatorUserName = *creatorUserNamePb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	effectivePerformanceTargetPb, err := identity(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}

	endTimePb, err := identity(&st.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimePb != nil {
		pb.EndTime = *endTimePb
	}

	executionDurationPb, err := identity(&st.ExecutionDuration)
	if err != nil {
		return nil, err
	}
	if executionDurationPb != nil {
		pb.ExecutionDuration = *executionDurationPb
	}

	gitSourcePb, err := gitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}

	hasMorePb, err := identity(&st.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMorePb != nil {
		pb.HasMore = *hasMorePb
	}

	var iterationsPb []runTaskPb
	for _, item := range st.Iterations {
		itemPb, err := runTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			iterationsPb = append(iterationsPb, *itemPb)
		}
	}
	pb.Iterations = iterationsPb

	var jobClustersPb []jobClusterPb
	for _, item := range st.JobClusters {
		itemPb, err := jobClusterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobClustersPb = append(jobClustersPb, *itemPb)
		}
	}
	pb.JobClusters = jobClustersPb

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	var jobParametersPb []jobParameterPb
	for _, item := range st.JobParameters {
		itemPb, err := jobParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb = append(jobParametersPb, *itemPb)
		}
	}
	pb.JobParameters = jobParametersPb

	jobRunIdPb, err := identity(&st.JobRunId)
	if err != nil {
		return nil, err
	}
	if jobRunIdPb != nil {
		pb.JobRunId = *jobRunIdPb
	}

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	numberInJobPb, err := identity(&st.NumberInJob)
	if err != nil {
		return nil, err
	}
	if numberInJobPb != nil {
		pb.NumberInJob = *numberInJobPb
	}

	originalAttemptRunIdPb, err := identity(&st.OriginalAttemptRunId)
	if err != nil {
		return nil, err
	}
	if originalAttemptRunIdPb != nil {
		pb.OriginalAttemptRunId = *originalAttemptRunIdPb
	}

	overridingParametersPb, err := runParametersToPb(st.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersPb != nil {
		pb.OverridingParameters = overridingParametersPb
	}

	queueDurationPb, err := identity(&st.QueueDuration)
	if err != nil {
		return nil, err
	}
	if queueDurationPb != nil {
		pb.QueueDuration = *queueDurationPb
	}

	var repairHistoryPb []repairHistoryItemPb
	for _, item := range st.RepairHistory {
		itemPb, err := repairHistoryItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			repairHistoryPb = append(repairHistoryPb, *itemPb)
		}
	}
	pb.RepairHistory = repairHistoryPb

	runDurationPb, err := identity(&st.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationPb != nil {
		pb.RunDuration = *runDurationPb
	}

	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	runNamePb, err := identity(&st.RunName)
	if err != nil {
		return nil, err
	}
	if runNamePb != nil {
		pb.RunName = *runNamePb
	}

	runPageUrlPb, err := identity(&st.RunPageUrl)
	if err != nil {
		return nil, err
	}
	if runPageUrlPb != nil {
		pb.RunPageUrl = *runPageUrlPb
	}

	runTypePb, err := identity(&st.RunType)
	if err != nil {
		return nil, err
	}
	if runTypePb != nil {
		pb.RunType = *runTypePb
	}

	schedulePb, err := cronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	setupDurationPb, err := identity(&st.SetupDuration)
	if err != nil {
		return nil, err
	}
	if setupDurationPb != nil {
		pb.SetupDuration = *setupDurationPb
	}

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	statePb, err := runStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	statusPb, err := runStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	var tasksPb []runTaskPb
	for _, item := range st.Tasks {
		itemPb, err := runTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb

	triggerPb, err := identity(&st.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerPb != nil {
		pb.Trigger = *triggerPb
	}

	triggerInfoPb, err := triggerInfoToPb(st.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoPb != nil {
		pb.TriggerInfo = triggerInfoPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Run) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Run) MarshalJSON() ([]byte, error) {
	pb, err := runToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runPb struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *clusterInstancePb `json:"cluster_instance,omitempty"`
	// A snapshot of the job’s cluster specification when this run was
	// created.
	ClusterSpec *clusterSpecPb `json:"cluster_spec,omitempty"`
	// The creator user name. This field won’t be included in the response if
	// the user has already been deleted.
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// Description of the run
	Description string `json:"description,omitempty"`
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *gitSourcePb `json:"git_source,omitempty"`
	// Indicates if the run has more array properties (`tasks`, `job_clusters`)
	// that are not shown. They can be accessed via :method:jobs/getrun
	// endpoint. It is only relevant for API 2.2 :method:jobs/listruns requests
	// with `expand_tasks=true`.
	HasMore bool `json:"has_more,omitempty"`
	// Only populated by for-each iterations. The parent for-each task is
	// located in tasks array.
	Iterations []runTaskPb `json:"iterations,omitempty"`
	// A list of job cluster specifications that can be shared and reused by
	// tasks of this job. Libraries cannot be declared in a shared job cluster.
	// You must declare dependent libraries in task settings. If more than 100
	// job clusters are available, you can paginate through them using
	// :method:jobs/getrun.
	JobClusters []jobClusterPb `json:"job_clusters,omitempty"`
	// The canonical identifier of the job that contains this run.
	JobId int64 `json:"job_id,omitempty"`
	// Job-level parameters used in the run
	JobParameters []jobParameterPb `json:"job_parameters,omitempty"`
	// ID of the job run that this run belongs to. For legacy and single-task
	// job runs the field is populated with the job run ID. For task runs, the
	// field is populated with the ID of the job run that the task run belongs
	// to.
	JobRunId int64 `json:"job_run_id,omitempty"`
	// A token that can be used to list the next page of array properties.
	NextPageToken string `json:"next_page_token,omitempty"`
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// If this run is a retry of a prior run attempt, this field contains the
	// run_id of the original attempt; otherwise, it is the same as the run_id.
	OriginalAttemptRunId int64 `json:"original_attempt_run_id,omitempty"`
	// The parameters used for this run.
	OverridingParameters *runParametersPb `json:"overriding_parameters,omitempty"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64 `json:"queue_duration,omitempty"`
	// The repair history of the run.
	RepairHistory []repairHistoryItemPb `json:"repair_history,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64 `json:"run_duration,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	RunId int64 `json:"run_id,omitempty"`
	// An optional name for the run. The maximum length is 4096 bytes in UTF-8
	// encoding.
	RunName string `json:"run_name,omitempty"`
	// The URL to the detail page of the run.
	RunPageUrl string `json:"run_page_url,omitempty"`
	// The type of a run. * `JOB_RUN`: Normal job run. A run created with
	// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
	// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
	// :method:jobs/submit.
	//
	// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
	RunType RunType `json:"run_type,omitempty"`
	// The cron schedule that triggered this run if it was triggered by the
	// periodic scheduler.
	Schedule *cronSchedulePb `json:"schedule,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *runStatePb `json:"state,omitempty"`
	// The current status of the run
	Status *runStatusPb `json:"status,omitempty"`
	// The list of tasks performed by the run. Each task has its own `run_id`
	// which you can use to call `JobsGetOutput` to retrieve the run resutls. If
	// more than 100 tasks are available, you can paginate through them using
	// :method:jobs/getrun. Use the `next_page_token` field at the object root
	// to determine if more results are available.
	Tasks []runTaskPb `json:"tasks,omitempty"`
	// The type of trigger that fired this run.
	//
	// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
	// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
	// occurs you triggered a single run on demand through the UI or the API. *
	// `RETRY`: Indicates a run that is triggered as a retry of a previously
	// failed run. This occurs when you request to re-run the job in case of
	// failures. * `RUN_JOB_TASK`: Indicates a run that is triggered using a Run
	// Job task. * `FILE_ARRIVAL`: Indicates a run that is triggered by a file
	// arrival. * `TABLE`: Indicates a run that is triggered by a table update.
	// * `CONTINUOUS_RESTART`: Indicates a run created by user to manually
	// restart a continuous job run.
	Trigger TriggerType `json:"trigger,omitempty"`
	// Additional details about what triggered the run
	TriggerInfo *triggerInfoPb `json:"trigger_info,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runFromPb(pb *runPb) (*Run, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Run{}
	attemptNumberField, err := identity(&pb.AttemptNumber)
	if err != nil {
		return nil, err
	}
	if attemptNumberField != nil {
		st.AttemptNumber = *attemptNumberField
	}
	cleanupDurationField, err := identity(&pb.CleanupDuration)
	if err != nil {
		return nil, err
	}
	if cleanupDurationField != nil {
		st.CleanupDuration = *cleanupDurationField
	}
	clusterInstanceField, err := clusterInstanceFromPb(pb.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstanceField != nil {
		st.ClusterInstance = clusterInstanceField
	}
	clusterSpecField, err := clusterSpecFromPb(pb.ClusterSpec)
	if err != nil {
		return nil, err
	}
	if clusterSpecField != nil {
		st.ClusterSpec = clusterSpecField
	}
	creatorUserNameField, err := identity(&pb.CreatorUserName)
	if err != nil {
		return nil, err
	}
	if creatorUserNameField != nil {
		st.CreatorUserName = *creatorUserNameField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	effectivePerformanceTargetField, err := identity(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	endTimeField, err := identity(&pb.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimeField != nil {
		st.EndTime = *endTimeField
	}
	executionDurationField, err := identity(&pb.ExecutionDuration)
	if err != nil {
		return nil, err
	}
	if executionDurationField != nil {
		st.ExecutionDuration = *executionDurationField
	}
	gitSourceField, err := gitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	hasMoreField, err := identity(&pb.HasMore)
	if err != nil {
		return nil, err
	}
	if hasMoreField != nil {
		st.HasMore = *hasMoreField
	}

	var iterationsField []RunTask
	for _, itemPb := range pb.Iterations {
		item, err := runTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			iterationsField = append(iterationsField, *item)
		}
	}
	st.Iterations = iterationsField

	var jobClustersField []JobCluster
	for _, itemPb := range pb.JobClusters {
		item, err := jobClusterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobClustersField = append(jobClustersField, *item)
		}
	}
	st.JobClusters = jobClustersField
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	var jobParametersField []JobParameter
	for _, itemPb := range pb.JobParameters {
		item, err := jobParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField = append(jobParametersField, *item)
		}
	}
	st.JobParameters = jobParametersField
	jobRunIdField, err := identity(&pb.JobRunId)
	if err != nil {
		return nil, err
	}
	if jobRunIdField != nil {
		st.JobRunId = *jobRunIdField
	}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}
	numberInJobField, err := identity(&pb.NumberInJob)
	if err != nil {
		return nil, err
	}
	if numberInJobField != nil {
		st.NumberInJob = *numberInJobField
	}
	originalAttemptRunIdField, err := identity(&pb.OriginalAttemptRunId)
	if err != nil {
		return nil, err
	}
	if originalAttemptRunIdField != nil {
		st.OriginalAttemptRunId = *originalAttemptRunIdField
	}
	overridingParametersField, err := runParametersFromPb(pb.OverridingParameters)
	if err != nil {
		return nil, err
	}
	if overridingParametersField != nil {
		st.OverridingParameters = overridingParametersField
	}
	queueDurationField, err := identity(&pb.QueueDuration)
	if err != nil {
		return nil, err
	}
	if queueDurationField != nil {
		st.QueueDuration = *queueDurationField
	}

	var repairHistoryField []RepairHistoryItem
	for _, itemPb := range pb.RepairHistory {
		item, err := repairHistoryItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			repairHistoryField = append(repairHistoryField, *item)
		}
	}
	st.RepairHistory = repairHistoryField
	runDurationField, err := identity(&pb.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationField != nil {
		st.RunDuration = *runDurationField
	}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}
	runNameField, err := identity(&pb.RunName)
	if err != nil {
		return nil, err
	}
	if runNameField != nil {
		st.RunName = *runNameField
	}
	runPageUrlField, err := identity(&pb.RunPageUrl)
	if err != nil {
		return nil, err
	}
	if runPageUrlField != nil {
		st.RunPageUrl = *runPageUrlField
	}
	runTypeField, err := identity(&pb.RunType)
	if err != nil {
		return nil, err
	}
	if runTypeField != nil {
		st.RunType = *runTypeField
	}
	scheduleField, err := cronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	setupDurationField, err := identity(&pb.SetupDuration)
	if err != nil {
		return nil, err
	}
	if setupDurationField != nil {
		st.SetupDuration = *setupDurationField
	}
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}
	stateField, err := runStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := runStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	var tasksField []RunTask
	for _, itemPb := range pb.Tasks {
		item, err := runTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	triggerField, err := identity(&pb.Trigger)
	if err != nil {
		return nil, err
	}
	if triggerField != nil {
		st.Trigger = *triggerField
	}
	triggerInfoField, err := triggerInfoFromPb(pb.TriggerInfo)
	if err != nil {
		return nil, err
	}
	if triggerInfoField != nil {
		st.TriggerInfo = triggerInfoField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunConditionTask struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left string
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	Op ConditionTaskOp
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	Outcome string
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right string

	ForceSendFields []string
}

func runConditionTaskToPb(st *RunConditionTask) (*runConditionTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runConditionTaskPb{}
	leftPb, err := identity(&st.Left)
	if err != nil {
		return nil, err
	}
	if leftPb != nil {
		pb.Left = *leftPb
	}

	opPb, err := identity(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}

	outcomePb, err := identity(&st.Outcome)
	if err != nil {
		return nil, err
	}
	if outcomePb != nil {
		pb.Outcome = *outcomePb
	}

	rightPb, err := identity(&st.Right)
	if err != nil {
		return nil, err
	}
	if rightPb != nil {
		pb.Right = *rightPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunConditionTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runConditionTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runConditionTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunConditionTask) MarshalJSON() ([]byte, error) {
	pb, err := runConditionTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runConditionTaskPb struct {
	// The left operand of the condition task. Can be either a string value or a
	// job state or parameter reference.
	Left string `json:"left"`
	// * `EQUAL_TO`, `NOT_EQUAL` operators perform string comparison of their
	// operands. This means that `“12.0” == “12”` will evaluate to
	// `false`. * `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `LESS_THAN`,
	// `LESS_THAN_OR_EQUAL` operators perform numeric comparison of their
	// operands. `“12.0” >= “12”` will evaluate to `true`, `“10.0”
	// >= “12”` will evaluate to `false`.
	//
	// The boolean comparison to task values can be implemented with operators
	// `EQUAL_TO`, `NOT_EQUAL`. If a task value was set to a boolean value, it
	// will be serialized to `“true”` or `“false”` for the comparison.
	Op ConditionTaskOp `json:"op"`
	// The condition expression evaluation result. Filled in if the task was
	// successfully completed. Can be `"true"` or `"false"`
	Outcome string `json:"outcome,omitempty"`
	// The right operand of the condition task. Can be either a string value or
	// a job state or parameter reference.
	Right string `json:"right"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runConditionTaskFromPb(pb *runConditionTaskPb) (*RunConditionTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunConditionTask{}
	leftField, err := identity(&pb.Left)
	if err != nil {
		return nil, err
	}
	if leftField != nil {
		st.Left = *leftField
	}
	opField, err := identity(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	outcomeField, err := identity(&pb.Outcome)
	if err != nil {
		return nil, err
	}
	if outcomeField != nil {
		st.Outcome = *outcomeField
	}
	rightField, err := identity(&pb.Right)
	if err != nil {
		return nil, err
	}
	if rightField != nil {
		st.Right = *rightField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runConditionTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runConditionTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunForEachTask struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency int
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string
	// Read only field. Populated for GetRun and ListRuns RPC calls and stores
	// the execution stats of an For each task
	Stats *ForEachStats
	// Configuration for the task that will be run for each element in the array
	Task Task

	ForceSendFields []string
}

func runForEachTaskToPb(st *RunForEachTask) (*runForEachTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runForEachTaskPb{}
	concurrencyPb, err := identity(&st.Concurrency)
	if err != nil {
		return nil, err
	}
	if concurrencyPb != nil {
		pb.Concurrency = *concurrencyPb
	}

	inputsPb, err := identity(&st.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsPb != nil {
		pb.Inputs = *inputsPb
	}

	statsPb, err := forEachStatsToPb(st.Stats)
	if err != nil {
		return nil, err
	}
	if statsPb != nil {
		pb.Stats = statsPb
	}

	taskPb, err := taskToPb(&st.Task)
	if err != nil {
		return nil, err
	}
	if taskPb != nil {
		pb.Task = *taskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunForEachTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runForEachTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runForEachTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunForEachTask) MarshalJSON() ([]byte, error) {
	pb, err := runForEachTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runForEachTaskPb struct {
	// An optional maximum allowed number of concurrent runs of the task. Set
	// this value if you want to be able to execute multiple runs of the task
	// concurrently.
	Concurrency int `json:"concurrency,omitempty"`
	// Array for task to iterate on. This can be a JSON string or a reference to
	// an array parameter.
	Inputs string `json:"inputs"`
	// Read only field. Populated for GetRun and ListRuns RPC calls and stores
	// the execution stats of an For each task
	Stats *forEachStatsPb `json:"stats,omitempty"`
	// Configuration for the task that will be run for each element in the array
	Task taskPb `json:"task"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runForEachTaskFromPb(pb *runForEachTaskPb) (*RunForEachTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunForEachTask{}
	concurrencyField, err := identity(&pb.Concurrency)
	if err != nil {
		return nil, err
	}
	if concurrencyField != nil {
		st.Concurrency = *concurrencyField
	}
	inputsField, err := identity(&pb.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsField != nil {
		st.Inputs = *inputsField
	}
	statsField, err := forEachStatsFromPb(pb.Stats)
	if err != nil {
		return nil, err
	}
	if statsField != nil {
		st.Stats = statsField
	}
	taskField, err := taskFromPb(&pb.Task)
	if err != nil {
		return nil, err
	}
	if taskField != nil {
		st.Task = *taskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runForEachTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runForEachTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// An optional value indicating the condition that determines whether the task
// should be run once its dependencies have been completed. When omitted,
// defaults to `ALL_SUCCESS`.
//
// Possible values are: * `ALL_SUCCESS`: All dependencies have executed and
// succeeded * `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
// `NONE_FAILED`: None of the dependencies have failed and at least one was
// executed * `ALL_DONE`: All dependencies have been completed *
// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
// dependencies have failed
type RunIf string
type runIfPb string

// All dependencies have been completed
const RunIfAllDone RunIf = `ALL_DONE`

// ALl dependencies have failed
const RunIfAllFailed RunIf = `ALL_FAILED`

// All dependencies have executed and succeeded
const RunIfAllSuccess RunIf = `ALL_SUCCESS`

// At least one dependency failed
const RunIfAtLeastOneFailed RunIf = `AT_LEAST_ONE_FAILED`

// At least one dependency has succeeded
const RunIfAtLeastOneSuccess RunIf = `AT_LEAST_ONE_SUCCESS`

// None of the dependencies have failed and at least one was executed
const RunIfNoneFailed RunIf = `NONE_FAILED`

// String representation for [fmt.Print]
func (f *RunIf) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunIf) Set(v string) error {
	switch v {
	case `ALL_DONE`, `ALL_FAILED`, `ALL_SUCCESS`, `AT_LEAST_ONE_FAILED`, `AT_LEAST_ONE_SUCCESS`, `NONE_FAILED`:
		*f = RunIf(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_DONE", "ALL_FAILED", "ALL_SUCCESS", "AT_LEAST_ONE_FAILED", "AT_LEAST_ONE_SUCCESS", "NONE_FAILED"`, v)
	}
}

// Type always returns RunIf to satisfy [pflag.Value] interface
func (f *RunIf) Type() string {
	return "RunIf"
}

func runIfToPb(st *RunIf) (*runIfPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runIfPb(*st)
	return &pb, nil
}

func runIfFromPb(pb *runIfPb) (*RunIf, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunIf(*pb)
	return &st, nil
}

type RunJobOutput struct {
	// The run id of the triggered job run
	RunId int64

	ForceSendFields []string
}

func runJobOutputToPb(st *RunJobOutput) (*runJobOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runJobOutputPb{}
	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunJobOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runJobOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runJobOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunJobOutput) MarshalJSON() ([]byte, error) {
	pb, err := runJobOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runJobOutputPb struct {
	// The run id of the triggered job run
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runJobOutputFromPb(pb *runJobOutputPb) (*RunJobOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunJobOutput{}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runJobOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runJobOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunJobTask struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string
	// ID of the job to trigger.
	JobId int64
	// Job-level parameters used to trigger the job.
	JobParameters map[string]string
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams

	PythonNamedParams map[string]string
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string
}

func runJobTaskToPb(st *RunJobTask) (*runJobTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runJobTaskPb{}

	var dbtCommandsPb []string
	for _, item := range st.DbtCommands {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dbtCommandsPb = append(dbtCommandsPb, *itemPb)
		}
	}
	pb.DbtCommands = dbtCommandsPb

	var jarParamsPb []string
	for _, item := range st.JarParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jarParamsPb = append(jarParamsPb, *itemPb)
		}
	}
	pb.JarParams = jarParamsPb

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	jobParametersPb := map[string]string{}
	for k, v := range st.JobParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb[k] = *itemPb
		}
	}
	pb.JobParameters = jobParametersPb

	notebookParamsPb := map[string]string{}
	for k, v := range st.NotebookParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebookParamsPb[k] = *itemPb
		}
	}
	pb.NotebookParams = notebookParamsPb

	pipelineParamsPb, err := pipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}

	pythonNamedParamsPb := map[string]string{}
	for k, v := range st.PythonNamedParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonNamedParamsPb[k] = *itemPb
		}
	}
	pb.PythonNamedParams = pythonNamedParamsPb

	var pythonParamsPb []string
	for _, item := range st.PythonParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonParamsPb = append(pythonParamsPb, *itemPb)
		}
	}
	pb.PythonParams = pythonParamsPb

	var sparkSubmitParamsPb []string
	for _, item := range st.SparkSubmitParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkSubmitParamsPb = append(sparkSubmitParamsPb, *itemPb)
		}
	}
	pb.SparkSubmitParams = sparkSubmitParamsPb

	sqlParamsPb := map[string]string{}
	for k, v := range st.SqlParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlParamsPb[k] = *itemPb
		}
	}
	pb.SqlParams = sqlParamsPb

	return pb, nil
}

func (st *RunJobTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runJobTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runJobTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunJobTask) MarshalJSON() ([]byte, error) {
	pb, err := runJobTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runJobTaskPb struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string `json:"jar_params,omitempty"`
	// ID of the job to trigger.
	JobId int64 `json:"job_id"`
	// Job-level parameters used to trigger the job.
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *pipelineParamsPb `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`
}

func runJobTaskFromPb(pb *runJobTaskPb) (*RunJobTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunJobTask{}

	var dbtCommandsField []string
	for _, itemPb := range pb.DbtCommands {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dbtCommandsField = append(dbtCommandsField, *item)
		}
	}
	st.DbtCommands = dbtCommandsField

	var jarParamsField []string
	for _, itemPb := range pb.JarParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jarParamsField = append(jarParamsField, *item)
		}
	}
	st.JarParams = jarParamsField
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	jobParametersField := map[string]string{}
	for k, v := range pb.JobParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField[k] = *item
		}
	}
	st.JobParameters = jobParametersField

	notebookParamsField := map[string]string{}
	for k, v := range pb.NotebookParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notebookParamsField[k] = *item
		}
	}
	st.NotebookParams = notebookParamsField
	pipelineParamsField, err := pipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}

	pythonNamedParamsField := map[string]string{}
	for k, v := range pb.PythonNamedParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonNamedParamsField[k] = *item
		}
	}
	st.PythonNamedParams = pythonNamedParamsField

	var pythonParamsField []string
	for _, itemPb := range pb.PythonParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonParamsField = append(pythonParamsField, *item)
		}
	}
	st.PythonParams = pythonParamsField

	var sparkSubmitParamsField []string
	for _, itemPb := range pb.SparkSubmitParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkSubmitParamsField = append(sparkSubmitParamsField, *item)
		}
	}
	st.SparkSubmitParams = sparkSubmitParamsField

	sqlParamsField := map[string]string{}
	for k, v := range pb.SqlParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlParamsField[k] = *item
		}
	}
	st.SqlParams = sqlParamsField

	return st, nil
}

// A value indicating the run's lifecycle state. The possible values are: *
// `QUEUED`: The run is queued. * `PENDING`: The run is waiting to be executed
// while the cluster and execution context are being prepared. * `RUNNING`: The
// task of this run is being executed. * `TERMINATING`: The task of this run has
// completed, and the cluster and execution context are being cleaned up. *
// `TERMINATED`: The task of this run has completed, and the cluster and
// execution context have been cleaned up. This state is terminal. * `SKIPPED`:
// This run was aborted because a previous run of the same job was already
// active. This state is terminal. * `INTERNAL_ERROR`: An exceptional state that
// indicates a failure in the Jobs service, such as network failure over a long
// period. If a run on a new cluster ends in the `INTERNAL_ERROR` state, the
// Jobs service terminates the cluster as soon as possible. This state is
// terminal. * `BLOCKED`: The run is blocked on an upstream dependency. *
// `WAITING_FOR_RETRY`: The run is waiting for a retry.
type RunLifeCycleState string
type runLifeCycleStatePb string

// The run is blocked on an upstream dependency.
const RunLifeCycleStateBlocked RunLifeCycleState = `BLOCKED`

// An exceptional state that indicates a failure in the Jobs service, such as
// network failure over a long period. If a run on a new cluster ends in the
// `INTERNAL_ERROR` state, the Jobs service terminates the cluster as soon as
// possible. This state is terminal.
const RunLifeCycleStateInternalError RunLifeCycleState = `INTERNAL_ERROR`

// The run is waiting to be executed while the cluster and execution context are
// being prepared.
const RunLifeCycleStatePending RunLifeCycleState = `PENDING`

// The run is queued.
const RunLifeCycleStateQueued RunLifeCycleState = `QUEUED`

// The task of this run is being executed.
const RunLifeCycleStateRunning RunLifeCycleState = `RUNNING`

// This run was aborted because a previous run of the same job was already
// active. This state is terminal.
const RunLifeCycleStateSkipped RunLifeCycleState = `SKIPPED`

// The task of this run has completed, and the cluster and execution context
// have been cleaned up. This state is terminal.
const RunLifeCycleStateTerminated RunLifeCycleState = `TERMINATED`

// The task of this run has completed, and the cluster and execution context are
// being cleaned up.
const RunLifeCycleStateTerminating RunLifeCycleState = `TERMINATING`

// The run is waiting for a retry.
const RunLifeCycleStateWaitingForRetry RunLifeCycleState = `WAITING_FOR_RETRY`

// String representation for [fmt.Print]
func (f *RunLifeCycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunLifeCycleState) Set(v string) error {
	switch v {
	case `BLOCKED`, `INTERNAL_ERROR`, `PENDING`, `QUEUED`, `RUNNING`, `SKIPPED`, `TERMINATED`, `TERMINATING`, `WAITING_FOR_RETRY`:
		*f = RunLifeCycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "INTERNAL_ERROR", "PENDING", "QUEUED", "RUNNING", "SKIPPED", "TERMINATED", "TERMINATING", "WAITING_FOR_RETRY"`, v)
	}
}

// Type always returns RunLifeCycleState to satisfy [pflag.Value] interface
func (f *RunLifeCycleState) Type() string {
	return "RunLifeCycleState"
}

func runLifeCycleStateToPb(st *RunLifeCycleState) (*runLifeCycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runLifeCycleStatePb(*st)
	return &pb, nil
}

func runLifeCycleStateFromPb(pb *runLifeCycleStatePb) (*RunLifeCycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunLifeCycleState(*pb)
	return &st, nil
}

// The current state of the run.
type RunLifecycleStateV2State string
type runLifecycleStateV2StatePb string

const RunLifecycleStateV2StateBlocked RunLifecycleStateV2State = `BLOCKED`

const RunLifecycleStateV2StatePending RunLifecycleStateV2State = `PENDING`

const RunLifecycleStateV2StateQueued RunLifecycleStateV2State = `QUEUED`

const RunLifecycleStateV2StateRunning RunLifecycleStateV2State = `RUNNING`

const RunLifecycleStateV2StateTerminated RunLifecycleStateV2State = `TERMINATED`

const RunLifecycleStateV2StateTerminating RunLifecycleStateV2State = `TERMINATING`

const RunLifecycleStateV2StateWaiting RunLifecycleStateV2State = `WAITING`

// String representation for [fmt.Print]
func (f *RunLifecycleStateV2State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunLifecycleStateV2State) Set(v string) error {
	switch v {
	case `BLOCKED`, `PENDING`, `QUEUED`, `RUNNING`, `TERMINATED`, `TERMINATING`, `WAITING`:
		*f = RunLifecycleStateV2State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCKED", "PENDING", "QUEUED", "RUNNING", "TERMINATED", "TERMINATING", "WAITING"`, v)
	}
}

// Type always returns RunLifecycleStateV2State to satisfy [pflag.Value] interface
func (f *RunLifecycleStateV2State) Type() string {
	return "RunLifecycleStateV2State"
}

func runLifecycleStateV2StateToPb(st *RunLifecycleStateV2State) (*runLifecycleStateV2StatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runLifecycleStateV2StatePb(*st)
	return &pb, nil
}

func runLifecycleStateV2StateFromPb(pb *runLifecycleStateV2StatePb) (*RunLifecycleStateV2State, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunLifecycleStateV2State(*pb)
	return &st, nil
}

type RunNow struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string
	// An optional token to guarantee the idempotency of job run requests. If a
	// run with the provided token already exists, the request does not create a
	// new run but returns the ID of the existing run instead. If a run with the
	// provided token is deleted, an error is returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	IdempotencyToken string
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string
	// The ID of the job to be executed
	JobId int64
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string
	// A list of task keys to run inside of the job. If this field is not
	// provided, all tasks in the job will be run.
	Only []string
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams

	PythonNamedParams map[string]string
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string
	// The queue settings of the run.
	Queue *QueueSettings
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string

	ForceSendFields []string
}

func runNowToPb(st *RunNow) (*runNowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runNowPb{}

	var dbtCommandsPb []string
	for _, item := range st.DbtCommands {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dbtCommandsPb = append(dbtCommandsPb, *itemPb)
		}
	}
	pb.DbtCommands = dbtCommandsPb

	idempotencyTokenPb, err := identity(&st.IdempotencyToken)
	if err != nil {
		return nil, err
	}
	if idempotencyTokenPb != nil {
		pb.IdempotencyToken = *idempotencyTokenPb
	}

	var jarParamsPb []string
	for _, item := range st.JarParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jarParamsPb = append(jarParamsPb, *itemPb)
		}
	}
	pb.JarParams = jarParamsPb

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	jobParametersPb := map[string]string{}
	for k, v := range st.JobParameters {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jobParametersPb[k] = *itemPb
		}
	}
	pb.JobParameters = jobParametersPb

	notebookParamsPb := map[string]string{}
	for k, v := range st.NotebookParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebookParamsPb[k] = *itemPb
		}
	}
	pb.NotebookParams = notebookParamsPb

	var onlyPb []string
	for _, item := range st.Only {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onlyPb = append(onlyPb, *itemPb)
		}
	}
	pb.Only = onlyPb

	performanceTargetPb, err := identity(&st.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetPb != nil {
		pb.PerformanceTarget = *performanceTargetPb
	}

	pipelineParamsPb, err := pipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}

	pythonNamedParamsPb := map[string]string{}
	for k, v := range st.PythonNamedParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonNamedParamsPb[k] = *itemPb
		}
	}
	pb.PythonNamedParams = pythonNamedParamsPb

	var pythonParamsPb []string
	for _, item := range st.PythonParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonParamsPb = append(pythonParamsPb, *itemPb)
		}
	}
	pb.PythonParams = pythonParamsPb

	queuePb, err := queueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}

	var sparkSubmitParamsPb []string
	for _, item := range st.SparkSubmitParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkSubmitParamsPb = append(sparkSubmitParamsPb, *itemPb)
		}
	}
	pb.SparkSubmitParams = sparkSubmitParamsPb

	sqlParamsPb := map[string]string{}
	for k, v := range st.SqlParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlParamsPb[k] = *itemPb
		}
	}
	pb.SqlParams = sqlParamsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunNow) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runNowPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runNowFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunNow) MarshalJSON() ([]byte, error) {
	pb, err := runNowToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runNowPb struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// An optional token to guarantee the idempotency of job run requests. If a
	// run with the provided token already exists, the request does not create a
	// new run but returns the ID of the existing run instead. If a run with the
	// provided token is deleted, an error is returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string `json:"jar_params,omitempty"`
	// The ID of the job to be executed
	JobId int64 `json:"job_id"`
	// Job-level parameters used in the run. for example `"param":
	// "overriding_val"`
	JobParameters map[string]string `json:"job_parameters,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// A list of task keys to run inside of the job. If this field is not
	// provided, all tasks in the job will be run.
	Only []string `json:"only,omitempty"`
	// The performance mode on a serverless job. The performance target
	// determines the level of compute performance or cost-efficiency for the
	// run. This field overrides the performance target defined on the job
	// level.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	PerformanceTarget PerformanceTarget `json:"performance_target,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *pipelineParamsPb `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string `json:"python_params,omitempty"`
	// The queue settings of the run.
	Queue *queueSettingsPb `json:"queue,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runNowFromPb(pb *runNowPb) (*RunNow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunNow{}

	var dbtCommandsField []string
	for _, itemPb := range pb.DbtCommands {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dbtCommandsField = append(dbtCommandsField, *item)
		}
	}
	st.DbtCommands = dbtCommandsField
	idempotencyTokenField, err := identity(&pb.IdempotencyToken)
	if err != nil {
		return nil, err
	}
	if idempotencyTokenField != nil {
		st.IdempotencyToken = *idempotencyTokenField
	}

	var jarParamsField []string
	for _, itemPb := range pb.JarParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jarParamsField = append(jarParamsField, *item)
		}
	}
	st.JarParams = jarParamsField
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}

	jobParametersField := map[string]string{}
	for k, v := range pb.JobParameters {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jobParametersField[k] = *item
		}
	}
	st.JobParameters = jobParametersField

	notebookParamsField := map[string]string{}
	for k, v := range pb.NotebookParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notebookParamsField[k] = *item
		}
	}
	st.NotebookParams = notebookParamsField

	var onlyField []string
	for _, itemPb := range pb.Only {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onlyField = append(onlyField, *item)
		}
	}
	st.Only = onlyField
	performanceTargetField, err := identity(&pb.PerformanceTarget)
	if err != nil {
		return nil, err
	}
	if performanceTargetField != nil {
		st.PerformanceTarget = *performanceTargetField
	}
	pipelineParamsField, err := pipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}

	pythonNamedParamsField := map[string]string{}
	for k, v := range pb.PythonNamedParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonNamedParamsField[k] = *item
		}
	}
	st.PythonNamedParams = pythonNamedParamsField

	var pythonParamsField []string
	for _, itemPb := range pb.PythonParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonParamsField = append(pythonParamsField, *item)
		}
	}
	st.PythonParams = pythonParamsField
	queueField, err := queueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}

	var sparkSubmitParamsField []string
	for _, itemPb := range pb.SparkSubmitParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkSubmitParamsField = append(sparkSubmitParamsField, *item)
		}
	}
	st.SparkSubmitParams = sparkSubmitParamsField

	sqlParamsField := map[string]string{}
	for k, v := range pb.SqlParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlParamsField[k] = *item
		}
	}
	st.SqlParams = sqlParamsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runNowPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runNowPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Run was started successfully.
type RunNowResponse struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64
	// The globally unique ID of the newly triggered run.
	RunId int64

	ForceSendFields []string
}

func runNowResponseToPb(st *RunNowResponse) (*runNowResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runNowResponsePb{}
	numberInJobPb, err := identity(&st.NumberInJob)
	if err != nil {
		return nil, err
	}
	if numberInJobPb != nil {
		pb.NumberInJob = *numberInJobPb
	}

	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunNowResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runNowResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runNowResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunNowResponse) MarshalJSON() ([]byte, error) {
	pb, err := runNowResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runNowResponsePb struct {
	// A unique identifier for this job run. This is set to the same value as
	// `run_id`.
	NumberInJob int64 `json:"number_in_job,omitempty"`
	// The globally unique ID of the newly triggered run.
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runNowResponseFromPb(pb *runNowResponsePb) (*RunNowResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunNowResponse{}
	numberInJobField, err := identity(&pb.NumberInJob)
	if err != nil {
		return nil, err
	}
	if numberInJobField != nil {
		st.NumberInJob = *numberInJobField
	}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runNowResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runNowResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Run output was retrieved successfully.
type RunOutput struct {
	// The output of a clean rooms notebook task, if available
	CleanRoomsNotebookOutput *CleanRoomsNotebookTaskCleanRoomsNotebookTaskOutput
	// The output of a dashboard task, if available
	DashboardOutput *DashboardTaskOutput
	// The output of a dbt task, if available.
	DbtOutput *DbtOutput
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	Error string
	// If there was an error executing the run, this field contains any
	// available stack traces.
	ErrorTrace string

	Info string
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	Logs string
	// Whether the logs are truncated.
	LogsTruncated bool
	// All details of the run except for its output.
	Metadata *Run
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. Databricks restricts this API
	// to return the first 5 MB of the output. To return a larger result, use
	// the [ClusterLogConf] field to configure log storage for the job cluster.
	//
	// [ClusterLogConf]: https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlogconf
	NotebookOutput *NotebookOutput
	// The output of a run job task, if available
	RunJobOutput *RunJobOutput
	// The output of a SQL task, if available.
	SqlOutput *SqlOutput

	ForceSendFields []string
}

func runOutputToPb(st *RunOutput) (*runOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runOutputPb{}
	cleanRoomsNotebookOutputPb, err := cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputToPb(st.CleanRoomsNotebookOutput)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookOutputPb != nil {
		pb.CleanRoomsNotebookOutput = cleanRoomsNotebookOutputPb
	}

	dashboardOutputPb, err := dashboardTaskOutputToPb(st.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputPb != nil {
		pb.DashboardOutput = dashboardOutputPb
	}

	dbtOutputPb, err := dbtOutputToPb(st.DbtOutput)
	if err != nil {
		return nil, err
	}
	if dbtOutputPb != nil {
		pb.DbtOutput = dbtOutputPb
	}

	errorPb, err := identity(&st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = *errorPb
	}

	errorTracePb, err := identity(&st.ErrorTrace)
	if err != nil {
		return nil, err
	}
	if errorTracePb != nil {
		pb.ErrorTrace = *errorTracePb
	}

	infoPb, err := identity(&st.Info)
	if err != nil {
		return nil, err
	}
	if infoPb != nil {
		pb.Info = *infoPb
	}

	logsPb, err := identity(&st.Logs)
	if err != nil {
		return nil, err
	}
	if logsPb != nil {
		pb.Logs = *logsPb
	}

	logsTruncatedPb, err := identity(&st.LogsTruncated)
	if err != nil {
		return nil, err
	}
	if logsTruncatedPb != nil {
		pb.LogsTruncated = *logsTruncatedPb
	}

	metadataPb, err := runToPb(st.Metadata)
	if err != nil {
		return nil, err
	}
	if metadataPb != nil {
		pb.Metadata = metadataPb
	}

	notebookOutputPb, err := notebookOutputToPb(st.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputPb != nil {
		pb.NotebookOutput = notebookOutputPb
	}

	runJobOutputPb, err := runJobOutputToPb(st.RunJobOutput)
	if err != nil {
		return nil, err
	}
	if runJobOutputPb != nil {
		pb.RunJobOutput = runJobOutputPb
	}

	sqlOutputPb, err := sqlOutputToPb(st.SqlOutput)
	if err != nil {
		return nil, err
	}
	if sqlOutputPb != nil {
		pb.SqlOutput = sqlOutputPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunOutput) MarshalJSON() ([]byte, error) {
	pb, err := runOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runOutputPb struct {
	// The output of a clean rooms notebook task, if available
	CleanRoomsNotebookOutput *cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputPb `json:"clean_rooms_notebook_output,omitempty"`
	// The output of a dashboard task, if available
	DashboardOutput *dashboardTaskOutputPb `json:"dashboard_output,omitempty"`
	// The output of a dbt task, if available.
	DbtOutput *dbtOutputPb `json:"dbt_output,omitempty"`
	// An error message indicating why a task failed or why output is not
	// available. The message is unstructured, and its exact format is subject
	// to change.
	Error string `json:"error,omitempty"`
	// If there was an error executing the run, this field contains any
	// available stack traces.
	ErrorTrace string `json:"error_trace,omitempty"`

	Info string `json:"info,omitempty"`
	// The output from tasks that write to standard streams (stdout/stderr) such
	// as spark_jar_task, spark_python_task, python_wheel_task.
	//
	// It's not supported for the notebook_task, pipeline_task or
	// spark_submit_task.
	//
	// Databricks restricts this API to return the last 5 MB of these logs.
	Logs string `json:"logs,omitempty"`
	// Whether the logs are truncated.
	LogsTruncated bool `json:"logs_truncated,omitempty"`
	// All details of the run except for its output.
	Metadata *runPb `json:"metadata,omitempty"`
	// The output of a notebook task, if available. A notebook task that
	// terminates (either successfully or with a failure) without calling
	// `dbutils.notebook.exit()` is considered to have an empty output. This
	// field is set but its result value is empty. Databricks restricts this API
	// to return the first 5 MB of the output. To return a larger result, use
	// the [ClusterLogConf] field to configure log storage for the job cluster.
	//
	// [ClusterLogConf]: https://docs.databricks.com/dev-tools/api/latest/clusters.html#clusterlogconf
	NotebookOutput *notebookOutputPb `json:"notebook_output,omitempty"`
	// The output of a run job task, if available
	RunJobOutput *runJobOutputPb `json:"run_job_output,omitempty"`
	// The output of a SQL task, if available.
	SqlOutput *sqlOutputPb `json:"sql_output,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runOutputFromPb(pb *runOutputPb) (*RunOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunOutput{}
	cleanRoomsNotebookOutputField, err := cleanRoomsNotebookTaskCleanRoomsNotebookTaskOutputFromPb(pb.CleanRoomsNotebookOutput)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookOutputField != nil {
		st.CleanRoomsNotebookOutput = cleanRoomsNotebookOutputField
	}
	dashboardOutputField, err := dashboardTaskOutputFromPb(pb.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputField != nil {
		st.DashboardOutput = dashboardOutputField
	}
	dbtOutputField, err := dbtOutputFromPb(pb.DbtOutput)
	if err != nil {
		return nil, err
	}
	if dbtOutputField != nil {
		st.DbtOutput = dbtOutputField
	}
	errorField, err := identity(&pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = *errorField
	}
	errorTraceField, err := identity(&pb.ErrorTrace)
	if err != nil {
		return nil, err
	}
	if errorTraceField != nil {
		st.ErrorTrace = *errorTraceField
	}
	infoField, err := identity(&pb.Info)
	if err != nil {
		return nil, err
	}
	if infoField != nil {
		st.Info = *infoField
	}
	logsField, err := identity(&pb.Logs)
	if err != nil {
		return nil, err
	}
	if logsField != nil {
		st.Logs = *logsField
	}
	logsTruncatedField, err := identity(&pb.LogsTruncated)
	if err != nil {
		return nil, err
	}
	if logsTruncatedField != nil {
		st.LogsTruncated = *logsTruncatedField
	}
	metadataField, err := runFromPb(pb.Metadata)
	if err != nil {
		return nil, err
	}
	if metadataField != nil {
		st.Metadata = metadataField
	}
	notebookOutputField, err := notebookOutputFromPb(pb.NotebookOutput)
	if err != nil {
		return nil, err
	}
	if notebookOutputField != nil {
		st.NotebookOutput = notebookOutputField
	}
	runJobOutputField, err := runJobOutputFromPb(pb.RunJobOutput)
	if err != nil {
		return nil, err
	}
	if runJobOutputField != nil {
		st.RunJobOutput = runJobOutputField
	}
	sqlOutputField, err := sqlOutputFromPb(pb.SqlOutput)
	if err != nil {
		return nil, err
	}
	if sqlOutputField != nil {
		st.SqlOutput = sqlOutputField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunParameters struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *PipelineParams

	PythonNamedParams map[string]string
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string
}

func runParametersToPb(st *RunParameters) (*runParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runParametersPb{}

	var dbtCommandsPb []string
	for _, item := range st.DbtCommands {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dbtCommandsPb = append(dbtCommandsPb, *itemPb)
		}
	}
	pb.DbtCommands = dbtCommandsPb

	var jarParamsPb []string
	for _, item := range st.JarParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			jarParamsPb = append(jarParamsPb, *itemPb)
		}
	}
	pb.JarParams = jarParamsPb

	notebookParamsPb := map[string]string{}
	for k, v := range st.NotebookParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			notebookParamsPb[k] = *itemPb
		}
	}
	pb.NotebookParams = notebookParamsPb

	pipelineParamsPb, err := pipelineParamsToPb(st.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsPb != nil {
		pb.PipelineParams = pipelineParamsPb
	}

	pythonNamedParamsPb := map[string]string{}
	for k, v := range st.PythonNamedParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonNamedParamsPb[k] = *itemPb
		}
	}
	pb.PythonNamedParams = pythonNamedParamsPb

	var pythonParamsPb []string
	for _, item := range st.PythonParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			pythonParamsPb = append(pythonParamsPb, *itemPb)
		}
	}
	pb.PythonParams = pythonParamsPb

	var sparkSubmitParamsPb []string
	for _, item := range st.SparkSubmitParams {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sparkSubmitParamsPb = append(sparkSubmitParamsPb, *itemPb)
		}
	}
	pb.SparkSubmitParams = sparkSubmitParamsPb

	sqlParamsPb := map[string]string{}
	for k, v := range st.SqlParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlParamsPb[k] = *itemPb
		}
	}
	pb.SqlParams = sqlParamsPb

	return pb, nil
}

func (st *RunParameters) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runParametersPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runParametersFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunParameters) MarshalJSON() ([]byte, error) {
	pb, err := runParametersToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runParametersPb struct {
	// An array of commands to execute for jobs with the dbt task, for example
	// `"dbt_commands": ["dbt deps", "dbt seed", "dbt deps", "dbt seed", "dbt
	// run"]`
	DbtCommands []string `json:"dbt_commands,omitempty"`
	// A list of parameters for jobs with Spark JAR tasks, for example
	// `"jar_params": ["john doe", "35"]`. The parameters are used to invoke the
	// main function of the main class specified in the Spark JAR task. If not
	// specified upon `run-now`, it defaults to an empty list. jar_params cannot
	// be specified in conjunction with notebook_params. The JSON representation
	// of this field (for example `{"jar_params":["john doe","35"]}`) cannot
	// exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	JarParams []string `json:"jar_params,omitempty"`
	// A map from keys to values for jobs with notebook task, for example
	// `"notebook_params": {"name": "john doe", "age": "35"}`. The map is passed
	// to the notebook and is accessible through the [dbutils.widgets.get]
	// function.
	//
	// If not specified upon `run-now`, the triggered run uses the job’s base
	// parameters.
	//
	// notebook_params cannot be specified in conjunction with jar_params.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// The JSON representation of this field (for example
	// `{"notebook_params":{"name":"john doe","age":"35"}}`) cannot exceed
	// 10,000 bytes.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	// [dbutils.widgets.get]: https://docs.databricks.com/dev-tools/databricks-utils.html
	NotebookParams map[string]string `json:"notebook_params,omitempty"`
	// Controls whether the pipeline should perform a full refresh
	PipelineParams *pipelineParamsPb `json:"pipeline_params,omitempty"`

	PythonNamedParams map[string]string `json:"python_named_params,omitempty"`
	// A list of parameters for jobs with Python tasks, for example
	// `"python_params": ["john doe", "35"]`. The parameters are passed to
	// Python file as command-line parameters. If specified upon `run-now`, it
	// would overwrite the parameters specified in job setting. The JSON
	// representation of this field (for example `{"python_params":["john
	// doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	PythonParams []string `json:"python_params,omitempty"`
	// A list of parameters for jobs with spark submit task, for example
	// `"spark_submit_params": ["--class",
	// "org.apache.spark.examples.SparkPi"]`. The parameters are passed to
	// spark-submit script as command-line parameters. If specified upon
	// `run-now`, it would overwrite the parameters specified in job setting.
	// The JSON representation of this field (for example
	// `{"python_params":["john doe","35"]}`) cannot exceed 10,000 bytes.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs
	//
	// Important
	//
	// These parameters accept only Latin characters (ASCII character set).
	// Using non-ASCII characters returns an error. Examples of invalid,
	// non-ASCII characters are Chinese, Japanese kanjis, and emojis.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	SparkSubmitParams []string `json:"spark_submit_params,omitempty"`
	// A map from keys to values for jobs with SQL task, for example
	// `"sql_params": {"name": "john doe", "age": "35"}`. The SQL alert task
	// does not support custom parameters.
	SqlParams map[string]string `json:"sql_params,omitempty"`
}

func runParametersFromPb(pb *runParametersPb) (*RunParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunParameters{}

	var dbtCommandsField []string
	for _, itemPb := range pb.DbtCommands {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dbtCommandsField = append(dbtCommandsField, *item)
		}
	}
	st.DbtCommands = dbtCommandsField

	var jarParamsField []string
	for _, itemPb := range pb.JarParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			jarParamsField = append(jarParamsField, *item)
		}
	}
	st.JarParams = jarParamsField

	notebookParamsField := map[string]string{}
	for k, v := range pb.NotebookParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			notebookParamsField[k] = *item
		}
	}
	st.NotebookParams = notebookParamsField
	pipelineParamsField, err := pipelineParamsFromPb(pb.PipelineParams)
	if err != nil {
		return nil, err
	}
	if pipelineParamsField != nil {
		st.PipelineParams = pipelineParamsField
	}

	pythonNamedParamsField := map[string]string{}
	for k, v := range pb.PythonNamedParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonNamedParamsField[k] = *item
		}
	}
	st.PythonNamedParams = pythonNamedParamsField

	var pythonParamsField []string
	for _, itemPb := range pb.PythonParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			pythonParamsField = append(pythonParamsField, *item)
		}
	}
	st.PythonParams = pythonParamsField

	var sparkSubmitParamsField []string
	for _, itemPb := range pb.SparkSubmitParams {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sparkSubmitParamsField = append(sparkSubmitParamsField, *item)
		}
	}
	st.SparkSubmitParams = sparkSubmitParamsField

	sqlParamsField := map[string]string{}
	for k, v := range pb.SqlParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlParamsField[k] = *item
		}
	}
	st.SqlParams = sqlParamsField

	return st, nil
}

// A value indicating the run's result. The possible values are: * `SUCCESS`:
// The task completed successfully. * `FAILED`: The task completed with an
// error. * `TIMEDOUT`: The run was stopped after reaching the timeout. *
// `CANCELED`: The run was canceled at user request. *
// `MAXIMUM_CONCURRENT_RUNS_REACHED`: The run was skipped because the maximum
// concurrent runs were reached. * `EXCLUDED`: The run was skipped because the
// necessary conditions were not met. * `SUCCESS_WITH_FAILURES`: The job run
// completed successfully with some failures; leaf tasks were successful. *
// `UPSTREAM_FAILED`: The run was skipped because of an upstream failure. *
// `UPSTREAM_CANCELED`: The run was skipped because an upstream task was
// canceled. * `DISABLED`: The run was skipped because it was disabled
// explicitly by the user.
type RunResultState string
type runResultStatePb string

// The run was canceled at user request.
const RunResultStateCanceled RunResultState = `CANCELED`

// The run was skipped because it was disabled explicitly by the user.
const RunResultStateDisabled RunResultState = `DISABLED`

// The run was skipped because the necessary conditions were not met.
const RunResultStateExcluded RunResultState = `EXCLUDED`

// The task completed with an error.
const RunResultStateFailed RunResultState = `FAILED`

// The run was skipped because the maximum concurrent runs were reached.
const RunResultStateMaximumConcurrentRunsReached RunResultState = `MAXIMUM_CONCURRENT_RUNS_REACHED`

// The task completed successfully.
const RunResultStateSuccess RunResultState = `SUCCESS`

// The job run completed successfully with some failures; leaf tasks were
// successful.
const RunResultStateSuccessWithFailures RunResultState = `SUCCESS_WITH_FAILURES`

// The run was stopped after reaching the timeout.
const RunResultStateTimedout RunResultState = `TIMEDOUT`

// The run was skipped because an upstream task was canceled.
const RunResultStateUpstreamCanceled RunResultState = `UPSTREAM_CANCELED`

// The run was skipped because of an upstream failure.
const RunResultStateUpstreamFailed RunResultState = `UPSTREAM_FAILED`

// String representation for [fmt.Print]
func (f *RunResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunResultState) Set(v string) error {
	switch v {
	case `CANCELED`, `DISABLED`, `EXCLUDED`, `FAILED`, `MAXIMUM_CONCURRENT_RUNS_REACHED`, `SUCCESS`, `SUCCESS_WITH_FAILURES`, `TIMEDOUT`, `UPSTREAM_CANCELED`, `UPSTREAM_FAILED`:
		*f = RunResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "DISABLED", "EXCLUDED", "FAILED", "MAXIMUM_CONCURRENT_RUNS_REACHED", "SUCCESS", "SUCCESS_WITH_FAILURES", "TIMEDOUT", "UPSTREAM_CANCELED", "UPSTREAM_FAILED"`, v)
	}
}

// Type always returns RunResultState to satisfy [pflag.Value] interface
func (f *RunResultState) Type() string {
	return "RunResultState"
}

func runResultStateToPb(st *RunResultState) (*runResultStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runResultStatePb(*st)
	return &pb, nil
}

func runResultStateFromPb(pb *runResultStatePb) (*RunResultState, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunResultState(*pb)
	return &st, nil
}

// The current state of the run.
type RunState struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	LifeCycleState RunLifeCycleState
	// The reason indicating why the run was queued.
	QueueReason string
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	ResultState RunResultState
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage string
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	UserCancelledOrTimedout bool

	ForceSendFields []string
}

func runStateToPb(st *RunState) (*runStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runStatePb{}
	lifeCycleStatePb, err := identity(&st.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStatePb != nil {
		pb.LifeCycleState = *lifeCycleStatePb
	}

	queueReasonPb, err := identity(&st.QueueReason)
	if err != nil {
		return nil, err
	}
	if queueReasonPb != nil {
		pb.QueueReason = *queueReasonPb
	}

	resultStatePb, err := identity(&st.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStatePb != nil {
		pb.ResultState = *resultStatePb
	}

	stateMessagePb, err := identity(&st.StateMessage)
	if err != nil {
		return nil, err
	}
	if stateMessagePb != nil {
		pb.StateMessage = *stateMessagePb
	}

	userCancelledOrTimedoutPb, err := identity(&st.UserCancelledOrTimedout)
	if err != nil {
		return nil, err
	}
	if userCancelledOrTimedoutPb != nil {
		pb.UserCancelledOrTimedout = *userCancelledOrTimedoutPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunState) MarshalJSON() ([]byte, error) {
	pb, err := runStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runStatePb struct {
	// A value indicating the run's current lifecycle state. This field is
	// always available in the response. Note: Additional states might be
	// introduced in future releases.
	LifeCycleState RunLifeCycleState `json:"life_cycle_state,omitempty"`
	// The reason indicating why the run was queued.
	QueueReason string `json:"queue_reason,omitempty"`
	// A value indicating the run's result. This field is only available for
	// terminal lifecycle states. Note: Additional states might be introduced in
	// future releases.
	ResultState RunResultState `json:"result_state,omitempty"`
	// A descriptive message for the current state. This field is unstructured,
	// and its exact format is subject to change.
	StateMessage string `json:"state_message,omitempty"`
	// A value indicating whether a run was canceled manually by a user or by
	// the scheduler because the run timed out.
	UserCancelledOrTimedout bool `json:"user_cancelled_or_timedout,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runStateFromPb(pb *runStatePb) (*RunState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunState{}
	lifeCycleStateField, err := identity(&pb.LifeCycleState)
	if err != nil {
		return nil, err
	}
	if lifeCycleStateField != nil {
		st.LifeCycleState = *lifeCycleStateField
	}
	queueReasonField, err := identity(&pb.QueueReason)
	if err != nil {
		return nil, err
	}
	if queueReasonField != nil {
		st.QueueReason = *queueReasonField
	}
	resultStateField, err := identity(&pb.ResultState)
	if err != nil {
		return nil, err
	}
	if resultStateField != nil {
		st.ResultState = *resultStateField
	}
	stateMessageField, err := identity(&pb.StateMessage)
	if err != nil {
		return nil, err
	}
	if stateMessageField != nil {
		st.StateMessage = *stateMessageField
	}
	userCancelledOrTimedoutField, err := identity(&pb.UserCancelledOrTimedout)
	if err != nil {
		return nil, err
	}
	if userCancelledOrTimedoutField != nil {
		st.UserCancelledOrTimedout = *userCancelledOrTimedoutField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runStatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runStatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The current status of the run
type RunStatus struct {
	// If the run was queued, details about the reason for queuing the run.
	QueueDetails *QueueDetails
	// The current state of the run.
	State RunLifecycleStateV2State
	// If the run is in a TERMINATING or TERMINATED state, details about the
	// reason for terminating the run.
	TerminationDetails *TerminationDetails
}

func runStatusToPb(st *RunStatus) (*runStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runStatusPb{}
	queueDetailsPb, err := queueDetailsToPb(st.QueueDetails)
	if err != nil {
		return nil, err
	}
	if queueDetailsPb != nil {
		pb.QueueDetails = queueDetailsPb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	terminationDetailsPb, err := terminationDetailsToPb(st.TerminationDetails)
	if err != nil {
		return nil, err
	}
	if terminationDetailsPb != nil {
		pb.TerminationDetails = terminationDetailsPb
	}

	return pb, nil
}

func (st *RunStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunStatus) MarshalJSON() ([]byte, error) {
	pb, err := runStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runStatusPb struct {
	// If the run was queued, details about the reason for queuing the run.
	QueueDetails *queueDetailsPb `json:"queue_details,omitempty"`
	// The current state of the run.
	State RunLifecycleStateV2State `json:"state,omitempty"`
	// If the run is in a TERMINATING or TERMINATED state, details about the
	// reason for terminating the run.
	TerminationDetails *terminationDetailsPb `json:"termination_details,omitempty"`
}

func runStatusFromPb(pb *runStatusPb) (*RunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunStatus{}
	queueDetailsField, err := queueDetailsFromPb(pb.QueueDetails)
	if err != nil {
		return nil, err
	}
	if queueDetailsField != nil {
		st.QueueDetails = queueDetailsField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	terminationDetailsField, err := terminationDetailsFromPb(pb.TerminationDetails)
	if err != nil {
		return nil, err
	}
	if terminationDetailsField != nil {
		st.TerminationDetails = terminationDetailsField
	}

	return st, nil
}

// Used when outputting a child run, in GetRun or ListRuns.
type RunTask struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask *CleanRoomsNotebookTask
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *ClusterInstance
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *RunConditionTask
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *DashboardTask
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *DbtTask
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency
	// An optional description for this task.
	Description string
	// Deprecated, field was never used in production.
	Disabled bool
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey string
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *RunForEachTask

	GenAiComputeTask *GenAiComputeTask
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks. If `git_source` is set,
	// these tasks retrieve the file from the remote repository by default.
	// However, this behavior can be overridden by setting `source` to
	// `WORKSPACE` on the task. Note: dbt and SQL File tasks support only
	// version-controlled sources. If dbt or SQL File tasks are used,
	// `git_source` must be defined on the job.
	GitSource *GitSource
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *NotebookTask
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *TaskNotificationSettings
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *PipelineTask
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *PowerBiTask
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *PythonWheelTask
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64
	// Parameter values including resolved references
	ResolvedValues *ResolvedValues
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64
	// The ID of the task run.
	RunId int64
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *RunJobTask

	RunPageUrl string
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *SparkJarTask
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *SparkPythonTask
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *SparkSubmitTask
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *SqlTask
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64
	// Deprecated. Please use the `status` field instead.
	State *RunState
	// The current status of the run
	Status *RunStatus
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string
}

func runTaskToPb(st *RunTask) (*runTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runTaskPb{}
	attemptNumberPb, err := identity(&st.AttemptNumber)
	if err != nil {
		return nil, err
	}
	if attemptNumberPb != nil {
		pb.AttemptNumber = *attemptNumberPb
	}

	cleanRoomsNotebookTaskPb, err := cleanRoomsNotebookTaskToPb(st.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskPb != nil {
		pb.CleanRoomsNotebookTask = cleanRoomsNotebookTaskPb
	}

	cleanupDurationPb, err := identity(&st.CleanupDuration)
	if err != nil {
		return nil, err
	}
	if cleanupDurationPb != nil {
		pb.CleanupDuration = *cleanupDurationPb
	}

	clusterInstancePb, err := clusterInstanceToPb(st.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstancePb != nil {
		pb.ClusterInstance = clusterInstancePb
	}

	conditionTaskPb, err := runConditionTaskToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}

	dashboardTaskPb, err := dashboardTaskToPb(st.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskPb != nil {
		pb.DashboardTask = dashboardTaskPb
	}

	dbtTaskPb, err := dbtTaskToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}

	var dependsOnPb []taskDependencyPb
	for _, item := range st.DependsOn {
		itemPb, err := taskDependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependsOnPb = append(dependsOnPb, *itemPb)
		}
	}
	pb.DependsOn = dependsOnPb

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	disabledPb, err := identity(&st.Disabled)
	if err != nil {
		return nil, err
	}
	if disabledPb != nil {
		pb.Disabled = *disabledPb
	}

	effectivePerformanceTargetPb, err := identity(&st.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetPb != nil {
		pb.EffectivePerformanceTarget = *effectivePerformanceTargetPb
	}

	emailNotificationsPb, err := jobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	endTimePb, err := identity(&st.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimePb != nil {
		pb.EndTime = *endTimePb
	}

	environmentKeyPb, err := identity(&st.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyPb != nil {
		pb.EnvironmentKey = *environmentKeyPb
	}

	executionDurationPb, err := identity(&st.ExecutionDuration)
	if err != nil {
		return nil, err
	}
	if executionDurationPb != nil {
		pb.ExecutionDuration = *executionDurationPb
	}

	existingClusterIdPb, err := identity(&st.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdPb != nil {
		pb.ExistingClusterId = *existingClusterIdPb
	}

	forEachTaskPb, err := runForEachTaskToPb(st.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskPb != nil {
		pb.ForEachTask = forEachTaskPb
	}

	genAiComputeTaskPb, err := genAiComputeTaskToPb(st.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskPb != nil {
		pb.GenAiComputeTask = genAiComputeTaskPb
	}

	gitSourcePb, err := gitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}

	jobClusterKeyPb, err := identity(&st.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyPb != nil {
		pb.JobClusterKey = *jobClusterKeyPb
	}

	var librariesPb []compute.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}

	notebookTaskPb, err := notebookTaskToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}

	notificationSettingsPb, err := taskNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	pipelineTaskPb, err := pipelineTaskToPb(st.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskPb != nil {
		pb.PipelineTask = pipelineTaskPb
	}

	powerBiTaskPb, err := powerBiTaskToPb(st.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskPb != nil {
		pb.PowerBiTask = powerBiTaskPb
	}

	pythonWheelTaskPb, err := pythonWheelTaskToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}

	queueDurationPb, err := identity(&st.QueueDuration)
	if err != nil {
		return nil, err
	}
	if queueDurationPb != nil {
		pb.QueueDuration = *queueDurationPb
	}

	resolvedValuesPb, err := resolvedValuesToPb(st.ResolvedValues)
	if err != nil {
		return nil, err
	}
	if resolvedValuesPb != nil {
		pb.ResolvedValues = resolvedValuesPb
	}

	runDurationPb, err := identity(&st.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationPb != nil {
		pb.RunDuration = *runDurationPb
	}

	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	runIfPb, err := identity(&st.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfPb != nil {
		pb.RunIf = *runIfPb
	}

	runJobTaskPb, err := runJobTaskToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}

	runPageUrlPb, err := identity(&st.RunPageUrl)
	if err != nil {
		return nil, err
	}
	if runPageUrlPb != nil {
		pb.RunPageUrl = *runPageUrlPb
	}

	setupDurationPb, err := identity(&st.SetupDuration)
	if err != nil {
		return nil, err
	}
	if setupDurationPb != nil {
		pb.SetupDuration = *setupDurationPb
	}

	sparkJarTaskPb, err := sparkJarTaskToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}

	sparkPythonTaskPb, err := sparkPythonTaskToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}

	sparkSubmitTaskPb, err := sparkSubmitTaskToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}

	sqlTaskPb, err := sqlTaskToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	statePb, err := runStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	statusPb, err := runStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	taskKeyPb, err := identity(&st.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyPb != nil {
		pb.TaskKey = *taskKeyPb
	}

	timeoutSecondsPb, err := identity(&st.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsPb != nil {
		pb.TimeoutSeconds = *timeoutSecondsPb
	}

	webhookNotificationsPb, err := webhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunTask) MarshalJSON() ([]byte, error) {
	pb, err := runTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runTaskPb struct {
	// The sequence number of this run attempt for a triggered job run. The
	// initial attempt of a run has an attempt_number of 0. If the initial run
	// attempt fails, and the job has a retry policy (`max_retries` > 0),
	// subsequent runs are created with an `original_attempt_run_id` of the
	// original attempt’s ID and an incrementing `attempt_number`. Runs are
	// retried only until they succeed, and the maximum `attempt_number` is the
	// same as the `max_retries` value for the job.
	AttemptNumber int `json:"attempt_number,omitempty"`
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask *cleanRoomsNotebookTaskPb `json:"clean_rooms_notebook_task,omitempty"`
	// The time in milliseconds it took to terminate the cluster and clean up
	// any associated artifacts. The duration of a task run is the sum of the
	// `setup_duration`, `execution_duration`, and the `cleanup_duration`. The
	// `cleanup_duration` field is set to 0 for multitask job runs. The total
	// duration of a multitask job run is the value of the `run_duration` field.
	CleanupDuration int64 `json:"cleanup_duration,omitempty"`
	// The cluster used for this run. If the run is specified to use a new
	// cluster, this field is set once the Jobs service has requested a cluster
	// for the run.
	ClusterInstance *clusterInstancePb `json:"cluster_instance,omitempty"`
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *runConditionTaskPb `json:"condition_task,omitempty"`
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *dashboardTaskPb `json:"dashboard_task,omitempty"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *dbtTaskPb `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []taskDependencyPb `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// Deprecated, field was never used in production.
	Disabled bool `json:"disabled,omitempty"`
	// The actual performance target used by the serverless run during
	// execution. This can differ from the client-set performance target on the
	// request depending on whether the performance mode is supported by the job
	// type.
	//
	// * `STANDARD`: Enables cost-efficient execution of serverless workloads. *
	// `PERFORMANCE_OPTIMIZED`: Prioritizes fast startup and execution times
	// through rapid scaling and optimized cluster performance.
	EffectivePerformanceTarget PerformanceTarget `json:"effective_performance_target,omitempty"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *jobEmailNotificationsPb `json:"email_notifications,omitempty"`
	// The time at which this run ended in epoch milliseconds (milliseconds
	// since 1/1/1970 UTC). This field is set to 0 if the job is still running.
	EndTime int64 `json:"end_time,omitempty"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey string `json:"environment_key,omitempty"`
	// The time in milliseconds it took to execute the commands in the JAR or
	// notebook until they completed, failed, timed out, were cancelled, or
	// encountered an unexpected error. The duration of a task run is the sum of
	// the `setup_duration`, `execution_duration`, and the `cleanup_duration`.
	// The `execution_duration` field is set to 0 for multitask job runs. The
	// total duration of a multitask job run is the value of the `run_duration`
	// field.
	ExecutionDuration int64 `json:"execution_duration,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *runForEachTaskPb `json:"for_each_task,omitempty"`

	GenAiComputeTask *genAiComputeTaskPb `json:"gen_ai_compute_task,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks. If `git_source` is set,
	// these tasks retrieve the file from the remote repository by default.
	// However, this behavior can be overridden by setting `source` to
	// `WORKSPACE` on the task. Note: dbt and SQL File tasks support only
	// version-controlled sources. If dbt or SQL File tasks are used,
	// `git_source` must be defined on the job.
	GitSource *gitSourcePb `json:"git_source,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.LibraryPb `json:"libraries,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpecPb `json:"new_cluster,omitempty"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *notebookTaskPb `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *taskNotificationSettingsPb `json:"notification_settings,omitempty"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *pipelineTaskPb `json:"pipeline_task,omitempty"`
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *powerBiTaskPb `json:"power_bi_task,omitempty"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *pythonWheelTaskPb `json:"python_wheel_task,omitempty"`
	// The time in milliseconds that the run has spent in the queue.
	QueueDuration int64 `json:"queue_duration,omitempty"`
	// Parameter values including resolved references
	ResolvedValues *resolvedValuesPb `json:"resolved_values,omitempty"`
	// The time in milliseconds it took the job run and all of its repairs to
	// finish.
	RunDuration int64 `json:"run_duration,omitempty"`
	// The ID of the task run.
	RunId int64 `json:"run_id,omitempty"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `json:"run_if,omitempty"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *runJobTaskPb `json:"run_job_task,omitempty"`

	RunPageUrl string `json:"run_page_url,omitempty"`
	// The time in milliseconds it took to set up the cluster. For runs that run
	// on new clusters this is the cluster creation time, for runs that run on
	// existing clusters this time should be very short. The duration of a task
	// run is the sum of the `setup_duration`, `execution_duration`, and the
	// `cleanup_duration`. The `setup_duration` field is set to 0 for multitask
	// job runs. The total duration of a multitask job run is the value of the
	// `run_duration` field.
	SetupDuration int64 `json:"setup_duration,omitempty"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *sparkJarTaskPb `json:"spark_jar_task,omitempty"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *sparkPythonTaskPb `json:"spark_python_task,omitempty"`
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *sparkSubmitTaskPb `json:"spark_submit_task,omitempty"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *sqlTaskPb `json:"sql_task,omitempty"`
	// The time at which this run was started in epoch milliseconds
	// (milliseconds since 1/1/1970 UTC). This may not be the time when the job
	// task starts executing, for example, if the job is scheduled to run on a
	// new cluster, this is the time the cluster creation call is issued.
	StartTime int64 `json:"start_time,omitempty"`
	// Deprecated. Please use the `status` field instead.
	State *runStatePb `json:"state,omitempty"`
	// The current status of the run
	Status *runStatusPb `json:"status,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *webhookNotificationsPb `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runTaskFromPb(pb *runTaskPb) (*RunTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunTask{}
	attemptNumberField, err := identity(&pb.AttemptNumber)
	if err != nil {
		return nil, err
	}
	if attemptNumberField != nil {
		st.AttemptNumber = *attemptNumberField
	}
	cleanRoomsNotebookTaskField, err := cleanRoomsNotebookTaskFromPb(pb.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskField != nil {
		st.CleanRoomsNotebookTask = cleanRoomsNotebookTaskField
	}
	cleanupDurationField, err := identity(&pb.CleanupDuration)
	if err != nil {
		return nil, err
	}
	if cleanupDurationField != nil {
		st.CleanupDuration = *cleanupDurationField
	}
	clusterInstanceField, err := clusterInstanceFromPb(pb.ClusterInstance)
	if err != nil {
		return nil, err
	}
	if clusterInstanceField != nil {
		st.ClusterInstance = clusterInstanceField
	}
	conditionTaskField, err := runConditionTaskFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dashboardTaskField, err := dashboardTaskFromPb(pb.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskField != nil {
		st.DashboardTask = dashboardTaskField
	}
	dbtTaskField, err := dbtTaskFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}

	var dependsOnField []TaskDependency
	for _, itemPb := range pb.DependsOn {
		item, err := taskDependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependsOnField = append(dependsOnField, *item)
		}
	}
	st.DependsOn = dependsOnField
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	disabledField, err := identity(&pb.Disabled)
	if err != nil {
		return nil, err
	}
	if disabledField != nil {
		st.Disabled = *disabledField
	}
	effectivePerformanceTargetField, err := identity(&pb.EffectivePerformanceTarget)
	if err != nil {
		return nil, err
	}
	if effectivePerformanceTargetField != nil {
		st.EffectivePerformanceTarget = *effectivePerformanceTargetField
	}
	emailNotificationsField, err := jobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}
	endTimeField, err := identity(&pb.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimeField != nil {
		st.EndTime = *endTimeField
	}
	environmentKeyField, err := identity(&pb.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyField != nil {
		st.EnvironmentKey = *environmentKeyField
	}
	executionDurationField, err := identity(&pb.ExecutionDuration)
	if err != nil {
		return nil, err
	}
	if executionDurationField != nil {
		st.ExecutionDuration = *executionDurationField
	}
	existingClusterIdField, err := identity(&pb.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdField != nil {
		st.ExistingClusterId = *existingClusterIdField
	}
	forEachTaskField, err := runForEachTaskFromPb(pb.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskField != nil {
		st.ForEachTask = forEachTaskField
	}
	genAiComputeTaskField, err := genAiComputeTaskFromPb(pb.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskField != nil {
		st.GenAiComputeTask = genAiComputeTaskField
	}
	gitSourceField, err := gitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	jobClusterKeyField, err := identity(&pb.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyField != nil {
		st.JobClusterKey = *jobClusterKeyField
	}

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}
	notebookTaskField, err := notebookTaskFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	notificationSettingsField, err := taskNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	pipelineTaskField, err := pipelineTaskFromPb(pb.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskField != nil {
		st.PipelineTask = pipelineTaskField
	}
	powerBiTaskField, err := powerBiTaskFromPb(pb.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskField != nil {
		st.PowerBiTask = powerBiTaskField
	}
	pythonWheelTaskField, err := pythonWheelTaskFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	queueDurationField, err := identity(&pb.QueueDuration)
	if err != nil {
		return nil, err
	}
	if queueDurationField != nil {
		st.QueueDuration = *queueDurationField
	}
	resolvedValuesField, err := resolvedValuesFromPb(pb.ResolvedValues)
	if err != nil {
		return nil, err
	}
	if resolvedValuesField != nil {
		st.ResolvedValues = resolvedValuesField
	}
	runDurationField, err := identity(&pb.RunDuration)
	if err != nil {
		return nil, err
	}
	if runDurationField != nil {
		st.RunDuration = *runDurationField
	}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}
	runIfField, err := identity(&pb.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfField != nil {
		st.RunIf = *runIfField
	}
	runJobTaskField, err := runJobTaskFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	runPageUrlField, err := identity(&pb.RunPageUrl)
	if err != nil {
		return nil, err
	}
	if runPageUrlField != nil {
		st.RunPageUrl = *runPageUrlField
	}
	setupDurationField, err := identity(&pb.SetupDuration)
	if err != nil {
		return nil, err
	}
	if setupDurationField != nil {
		st.SetupDuration = *setupDurationField
	}
	sparkJarTaskField, err := sparkJarTaskFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := sparkPythonTaskFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := sparkSubmitTaskFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := sqlTaskFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}
	stateField, err := runStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	statusField, err := runStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}
	taskKeyField, err := identity(&pb.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyField != nil {
		st.TaskKey = *taskKeyField
	}
	timeoutSecondsField, err := identity(&pb.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsField != nil {
		st.TimeoutSeconds = *timeoutSecondsField
	}
	webhookNotificationsField, err := webhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of a run. * `JOB_RUN`: Normal job run. A run created with
// :method:jobs/runNow. * `WORKFLOW_RUN`: Workflow run. A run created with
// [dbutils.notebook.run]. * `SUBMIT_RUN`: Submit run. A run created with
// :method:jobs/submit.
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
type RunType string
type runTypePb string

// Normal job run. A run created with :method:jobs/runNow.
const RunTypeJobRun RunType = `JOB_RUN`

// Submit run. A run created with :method:jobs/submit.
const RunTypeSubmitRun RunType = `SUBMIT_RUN`

// Workflow run. A run created with [dbutils.notebook.run].
//
// [dbutils.notebook.run]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-workflow
const RunTypeWorkflowRun RunType = `WORKFLOW_RUN`

// String representation for [fmt.Print]
func (f *RunType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunType) Set(v string) error {
	switch v {
	case `JOB_RUN`, `SUBMIT_RUN`, `WORKFLOW_RUN`:
		*f = RunType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "JOB_RUN", "SUBMIT_RUN", "WORKFLOW_RUN"`, v)
	}
}

// Type always returns RunType to satisfy [pflag.Value] interface
func (f *RunType) Type() string {
	return "RunType"
}

func runTypeToPb(st *RunType) (*runTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runTypePb(*st)
	return &pb, nil
}

func runTypeFromPb(pb *runTypePb) (*RunType, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunType(*pb)
	return &st, nil
}

// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL file
// will be retrieved\ from the local Databricks workspace. When set to `GIT`,
// the SQL file will be retrieved from a Git repository defined in `git_source`.
// If the value is empty, the task will use `GIT` if `git_source` is defined and
// `WORKSPACE` otherwise.
//
// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL file
// is located in cloud Git provider.
type Source string
type sourcePb string

// SQL file is located in cloud Git provider.
const SourceGit Source = `GIT`

// SQL file is located in <Databricks> workspace.
const SourceWorkspace Source = `WORKSPACE`

// String representation for [fmt.Print]
func (f *Source) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Source) Set(v string) error {
	switch v {
	case `GIT`, `WORKSPACE`:
		*f = Source(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "GIT", "WORKSPACE"`, v)
	}
}

// Type always returns Source to satisfy [pflag.Value] interface
func (f *Source) Type() string {
	return "Source"
}

func sourceToPb(st *Source) (*sourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sourcePb(*st)
	return &pb, nil
}

func sourceFromPb(pb *sourcePb) (*Source, error) {
	if pb == nil {
		return nil, nil
	}
	st := Source(*pb)
	return &st, nil
}

type SparkJarTask struct {
	// Deprecated since 04/2016. Provide a `jar` through the `libraries` field
	// instead. For an example, see :method:jobs/create.
	JarUri string
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	MainClassName string
	// Parameters passed to the main method.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string
	// Deprecated. A value of `false` is no longer supported.
	RunAsRepl bool

	ForceSendFields []string
}

func sparkJarTaskToPb(st *SparkJarTask) (*sparkJarTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkJarTaskPb{}
	jarUriPb, err := identity(&st.JarUri)
	if err != nil {
		return nil, err
	}
	if jarUriPb != nil {
		pb.JarUri = *jarUriPb
	}

	mainClassNamePb, err := identity(&st.MainClassName)
	if err != nil {
		return nil, err
	}
	if mainClassNamePb != nil {
		pb.MainClassName = *mainClassNamePb
	}

	var parametersPb []string
	for _, item := range st.Parameters {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	runAsReplPb, err := identity(&st.RunAsRepl)
	if err != nil {
		return nil, err
	}
	if runAsReplPb != nil {
		pb.RunAsRepl = *runAsReplPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SparkJarTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkJarTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkJarTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkJarTask) MarshalJSON() ([]byte, error) {
	pb, err := sparkJarTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sparkJarTaskPb struct {
	// Deprecated since 04/2016. Provide a `jar` through the `libraries` field
	// instead. For an example, see :method:jobs/create.
	JarUri string `json:"jar_uri,omitempty"`
	// The full name of the class containing the main method to be executed.
	// This class must be contained in a JAR provided as a library.
	//
	// The code must use `SparkContext.getOrCreate` to obtain a Spark context;
	// otherwise, runs of the job fail.
	MainClassName string `json:"main_class_name,omitempty"`
	// Parameters passed to the main method.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
	// Deprecated. A value of `false` is no longer supported.
	RunAsRepl bool `json:"run_as_repl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sparkJarTaskFromPb(pb *sparkJarTaskPb) (*SparkJarTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkJarTask{}
	jarUriField, err := identity(&pb.JarUri)
	if err != nil {
		return nil, err
	}
	if jarUriField != nil {
		st.JarUri = *jarUriField
	}
	mainClassNameField, err := identity(&pb.MainClassName)
	if err != nil {
		return nil, err
	}
	if mainClassNameField != nil {
		st.MainClassName = *mainClassNameField
	}

	var parametersField []string
	for _, itemPb := range pb.Parameters {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	runAsReplField, err := identity(&pb.RunAsRepl)
	if err != nil {
		return nil, err
	}
	if runAsReplField != nil {
		st.RunAsRepl = *runAsReplField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sparkJarTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sparkJarTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SparkPythonTask struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	PythonFile string
	// Optional location type of the Python file. When set to `WORKSPACE` or not
	// specified, the file will be retrieved from the local Databricks workspace
	// or cloud location (if the `python_file` has a URI format). When set to
	// `GIT`, the Python file will be retrieved from a Git repository defined in
	// `git_source`.
	//
	// * `WORKSPACE`: The Python file is located in a Databricks workspace or at
	// a cloud filesystem URI. * `GIT`: The Python file is located in a remote
	// Git repository.
	Source Source
}

func sparkPythonTaskToPb(st *SparkPythonTask) (*sparkPythonTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkPythonTaskPb{}

	var parametersPb []string
	for _, item := range st.Parameters {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	pythonFilePb, err := identity(&st.PythonFile)
	if err != nil {
		return nil, err
	}
	if pythonFilePb != nil {
		pb.PythonFile = *pythonFilePb
	}

	sourcePb, err := identity(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	return pb, nil
}

func (st *SparkPythonTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkPythonTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkPythonTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkPythonTask) MarshalJSON() ([]byte, error) {
	pb, err := sparkPythonTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sparkPythonTaskPb struct {
	// Command line parameters passed to the Python file.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
	// The Python file to be executed. Cloud file URIs (such as dbfs:/, s3:/,
	// adls:/, gcs:/) and workspace paths are supported. For python files stored
	// in the Databricks workspace, the path must be absolute and begin with
	// `/`. For files stored in a remote repository, the path must be relative.
	// This field is required.
	PythonFile string `json:"python_file"`
	// Optional location type of the Python file. When set to `WORKSPACE` or not
	// specified, the file will be retrieved from the local Databricks workspace
	// or cloud location (if the `python_file` has a URI format). When set to
	// `GIT`, the Python file will be retrieved from a Git repository defined in
	// `git_source`.
	//
	// * `WORKSPACE`: The Python file is located in a Databricks workspace or at
	// a cloud filesystem URI. * `GIT`: The Python file is located in a remote
	// Git repository.
	Source Source `json:"source,omitempty"`
}

func sparkPythonTaskFromPb(pb *sparkPythonTaskPb) (*SparkPythonTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkPythonTask{}

	var parametersField []string
	for _, itemPb := range pb.Parameters {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	pythonFileField, err := identity(&pb.PythonFile)
	if err != nil {
		return nil, err
	}
	if pythonFileField != nil {
		st.PythonFile = *pythonFileField
	}
	sourceField, err := identity(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}

	return st, nil
}

type SparkSubmitTask struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string
}

func sparkSubmitTaskToPb(st *SparkSubmitTask) (*sparkSubmitTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sparkSubmitTaskPb{}

	var parametersPb []string
	for _, item := range st.Parameters {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	return pb, nil
}

func (st *SparkSubmitTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sparkSubmitTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sparkSubmitTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SparkSubmitTask) MarshalJSON() ([]byte, error) {
	pb, err := sparkSubmitTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sparkSubmitTaskPb struct {
	// Command-line parameters passed to spark submit.
	//
	// Use [Task parameter variables] to set parameters containing information
	// about job runs.
	//
	// [Task parameter variables]: https://docs.databricks.com/jobs.html#parameter-variables
	Parameters []string `json:"parameters,omitempty"`
}

func sparkSubmitTaskFromPb(pb *sparkSubmitTaskPb) (*SparkSubmitTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SparkSubmitTask{}

	var parametersField []string
	for _, itemPb := range pb.Parameters {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField

	return st, nil
}

type SqlAlertOutput struct {
	// The state of the SQL alert.
	//
	// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
	// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled
	// trigger conditions
	AlertState SqlAlertState
	// The link to find the output results.
	OutputLink string
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	QueryText string
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput
	// The canonical identifier of the SQL warehouse.
	WarehouseId string

	ForceSendFields []string
}

func sqlAlertOutputToPb(st *SqlAlertOutput) (*sqlAlertOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlAlertOutputPb{}
	alertStatePb, err := identity(&st.AlertState)
	if err != nil {
		return nil, err
	}
	if alertStatePb != nil {
		pb.AlertState = *alertStatePb
	}

	outputLinkPb, err := identity(&st.OutputLink)
	if err != nil {
		return nil, err
	}
	if outputLinkPb != nil {
		pb.OutputLink = *outputLinkPb
	}

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	var sqlStatementsPb []sqlStatementOutputPb
	for _, item := range st.SqlStatements {
		itemPb, err := sqlStatementOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlStatementsPb = append(sqlStatementsPb, *itemPb)
		}
	}
	pb.SqlStatements = sqlStatementsPb

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SqlAlertOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlAlertOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlAlertOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlAlertOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlAlertOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlAlertOutputPb struct {
	// The state of the SQL alert.
	//
	// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
	// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled
	// trigger conditions
	AlertState SqlAlertState `json:"alert_state,omitempty"`
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// The text of the SQL query. Can Run permission of the SQL query associated
	// with the SQL alert is required to view this field.
	QueryText string `json:"query_text,omitempty"`
	// Information about SQL statements executed in the run.
	SqlStatements []sqlStatementOutputPb `json:"sql_statements,omitempty"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlAlertOutputFromPb(pb *sqlAlertOutputPb) (*SqlAlertOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlAlertOutput{}
	alertStateField, err := identity(&pb.AlertState)
	if err != nil {
		return nil, err
	}
	if alertStateField != nil {
		st.AlertState = *alertStateField
	}
	outputLinkField, err := identity(&pb.OutputLink)
	if err != nil {
		return nil, err
	}
	if outputLinkField != nil {
		st.OutputLink = *outputLinkField
	}
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}

	var sqlStatementsField []SqlStatementOutput
	for _, itemPb := range pb.SqlStatements {
		item, err := sqlStatementOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlStatementsField = append(sqlStatementsField, *item)
		}
	}
	st.SqlStatements = sqlStatementsField
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlAlertOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlAlertOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The state of the SQL alert.
//
// * UNKNOWN: alert yet to be evaluated * OK: alert evaluated and did not
// fulfill trigger conditions * TRIGGERED: alert evaluated and fulfilled trigger
// conditions
type SqlAlertState string
type sqlAlertStatePb string

const SqlAlertStateOk SqlAlertState = `OK`

const SqlAlertStateTriggered SqlAlertState = `TRIGGERED`

const SqlAlertStateUnknown SqlAlertState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *SqlAlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SqlAlertState) Set(v string) error {
	switch v {
	case `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = SqlAlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Type always returns SqlAlertState to satisfy [pflag.Value] interface
func (f *SqlAlertState) Type() string {
	return "SqlAlertState"
}

func sqlAlertStateToPb(st *SqlAlertState) (*sqlAlertStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlAlertStatePb(*st)
	return &pb, nil
}

func sqlAlertStateFromPb(pb *sqlAlertStatePb) (*SqlAlertState, error) {
	if pb == nil {
		return nil, nil
	}
	st := SqlAlertState(*pb)
	return &st, nil
}

type SqlDashboardOutput struct {
	// The canonical identifier of the SQL warehouse.
	WarehouseId string
	// Widgets executed in the run. Only SQL query based widgets are listed.
	Widgets []SqlDashboardWidgetOutput

	ForceSendFields []string
}

func sqlDashboardOutputToPb(st *SqlDashboardOutput) (*sqlDashboardOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlDashboardOutputPb{}
	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	var widgetsPb []sqlDashboardWidgetOutputPb
	for _, item := range st.Widgets {
		itemPb, err := sqlDashboardWidgetOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			widgetsPb = append(widgetsPb, *itemPb)
		}
	}
	pb.Widgets = widgetsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SqlDashboardOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlDashboardOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlDashboardOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlDashboardOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlDashboardOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlDashboardOutputPb struct {
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`
	// Widgets executed in the run. Only SQL query based widgets are listed.
	Widgets []sqlDashboardWidgetOutputPb `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlDashboardOutputFromPb(pb *sqlDashboardOutputPb) (*SqlDashboardOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlDashboardOutput{}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	var widgetsField []SqlDashboardWidgetOutput
	for _, itemPb := range pb.Widgets {
		item, err := sqlDashboardWidgetOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			widgetsField = append(widgetsField, *item)
		}
	}
	st.Widgets = widgetsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlDashboardOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlDashboardOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlDashboardWidgetOutput struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	EndTime int64
	// The information about the error when execution fails.
	Error *SqlOutputError
	// The link to find the output results.
	OutputLink string
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	StartTime int64
	// The execution status of the SQL widget.
	Status SqlDashboardWidgetOutputStatus
	// The canonical identifier of the SQL widget.
	WidgetId string
	// The title of the SQL widget.
	WidgetTitle string

	ForceSendFields []string
}

func sqlDashboardWidgetOutputToPb(st *SqlDashboardWidgetOutput) (*sqlDashboardWidgetOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlDashboardWidgetOutputPb{}
	endTimePb, err := identity(&st.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimePb != nil {
		pb.EndTime = *endTimePb
	}

	errorPb, err := sqlOutputErrorToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}

	outputLinkPb, err := identity(&st.OutputLink)
	if err != nil {
		return nil, err
	}
	if outputLinkPb != nil {
		pb.OutputLink = *outputLinkPb
	}

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	widgetIdPb, err := identity(&st.WidgetId)
	if err != nil {
		return nil, err
	}
	if widgetIdPb != nil {
		pb.WidgetId = *widgetIdPb
	}

	widgetTitlePb, err := identity(&st.WidgetTitle)
	if err != nil {
		return nil, err
	}
	if widgetTitlePb != nil {
		pb.WidgetTitle = *widgetTitlePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SqlDashboardWidgetOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlDashboardWidgetOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlDashboardWidgetOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlDashboardWidgetOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlDashboardWidgetOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlDashboardWidgetOutputPb struct {
	// Time (in epoch milliseconds) when execution of the SQL widget ends.
	EndTime int64 `json:"end_time,omitempty"`
	// The information about the error when execution fails.
	Error *sqlOutputErrorPb `json:"error,omitempty"`
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// Time (in epoch milliseconds) when execution of the SQL widget starts.
	StartTime int64 `json:"start_time,omitempty"`
	// The execution status of the SQL widget.
	Status SqlDashboardWidgetOutputStatus `json:"status,omitempty"`
	// The canonical identifier of the SQL widget.
	WidgetId string `json:"widget_id,omitempty"`
	// The title of the SQL widget.
	WidgetTitle string `json:"widget_title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlDashboardWidgetOutputFromPb(pb *sqlDashboardWidgetOutputPb) (*SqlDashboardWidgetOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlDashboardWidgetOutput{}
	endTimeField, err := identity(&pb.EndTime)
	if err != nil {
		return nil, err
	}
	if endTimeField != nil {
		st.EndTime = *endTimeField
	}
	errorField, err := sqlOutputErrorFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	outputLinkField, err := identity(&pb.OutputLink)
	if err != nil {
		return nil, err
	}
	if outputLinkField != nil {
		st.OutputLink = *outputLinkField
	}
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	widgetIdField, err := identity(&pb.WidgetId)
	if err != nil {
		return nil, err
	}
	if widgetIdField != nil {
		st.WidgetId = *widgetIdField
	}
	widgetTitleField, err := identity(&pb.WidgetTitle)
	if err != nil {
		return nil, err
	}
	if widgetTitleField != nil {
		st.WidgetTitle = *widgetTitleField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlDashboardWidgetOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlDashboardWidgetOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlDashboardWidgetOutputStatus string
type sqlDashboardWidgetOutputStatusPb string

const SqlDashboardWidgetOutputStatusCancelled SqlDashboardWidgetOutputStatus = `CANCELLED`

const SqlDashboardWidgetOutputStatusFailed SqlDashboardWidgetOutputStatus = `FAILED`

const SqlDashboardWidgetOutputStatusPending SqlDashboardWidgetOutputStatus = `PENDING`

const SqlDashboardWidgetOutputStatusRunning SqlDashboardWidgetOutputStatus = `RUNNING`

const SqlDashboardWidgetOutputStatusSuccess SqlDashboardWidgetOutputStatus = `SUCCESS`

// String representation for [fmt.Print]
func (f *SqlDashboardWidgetOutputStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SqlDashboardWidgetOutputStatus) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCESS`:
		*f = SqlDashboardWidgetOutputStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCESS"`, v)
	}
}

// Type always returns SqlDashboardWidgetOutputStatus to satisfy [pflag.Value] interface
func (f *SqlDashboardWidgetOutputStatus) Type() string {
	return "SqlDashboardWidgetOutputStatus"
}

func sqlDashboardWidgetOutputStatusToPb(st *SqlDashboardWidgetOutputStatus) (*sqlDashboardWidgetOutputStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlDashboardWidgetOutputStatusPb(*st)
	return &pb, nil
}

func sqlDashboardWidgetOutputStatusFromPb(pb *sqlDashboardWidgetOutputStatusPb) (*SqlDashboardWidgetOutputStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SqlDashboardWidgetOutputStatus(*pb)
	return &st, nil
}

type SqlOutput struct {
	// The output of a SQL alert task, if available.
	AlertOutput *SqlAlertOutput
	// The output of a SQL dashboard task, if available.
	DashboardOutput *SqlDashboardOutput
	// The output of a SQL query task, if available.
	QueryOutput *SqlQueryOutput
}

func sqlOutputToPb(st *SqlOutput) (*sqlOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlOutputPb{}
	alertOutputPb, err := sqlAlertOutputToPb(st.AlertOutput)
	if err != nil {
		return nil, err
	}
	if alertOutputPb != nil {
		pb.AlertOutput = alertOutputPb
	}

	dashboardOutputPb, err := sqlDashboardOutputToPb(st.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputPb != nil {
		pb.DashboardOutput = dashboardOutputPb
	}

	queryOutputPb, err := sqlQueryOutputToPb(st.QueryOutput)
	if err != nil {
		return nil, err
	}
	if queryOutputPb != nil {
		pb.QueryOutput = queryOutputPb
	}

	return pb, nil
}

func (st *SqlOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlOutputPb struct {
	// The output of a SQL alert task, if available.
	AlertOutput *sqlAlertOutputPb `json:"alert_output,omitempty"`
	// The output of a SQL dashboard task, if available.
	DashboardOutput *sqlDashboardOutputPb `json:"dashboard_output,omitempty"`
	// The output of a SQL query task, if available.
	QueryOutput *sqlQueryOutputPb `json:"query_output,omitempty"`
}

func sqlOutputFromPb(pb *sqlOutputPb) (*SqlOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlOutput{}
	alertOutputField, err := sqlAlertOutputFromPb(pb.AlertOutput)
	if err != nil {
		return nil, err
	}
	if alertOutputField != nil {
		st.AlertOutput = alertOutputField
	}
	dashboardOutputField, err := sqlDashboardOutputFromPb(pb.DashboardOutput)
	if err != nil {
		return nil, err
	}
	if dashboardOutputField != nil {
		st.DashboardOutput = dashboardOutputField
	}
	queryOutputField, err := sqlQueryOutputFromPb(pb.QueryOutput)
	if err != nil {
		return nil, err
	}
	if queryOutputField != nil {
		st.QueryOutput = queryOutputField
	}

	return st, nil
}

type SqlOutputError struct {
	// The error message when execution fails.
	Message string

	ForceSendFields []string
}

func sqlOutputErrorToPb(st *SqlOutputError) (*sqlOutputErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlOutputErrorPb{}
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

func (st *SqlOutputError) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlOutputErrorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlOutputErrorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlOutputError) MarshalJSON() ([]byte, error) {
	pb, err := sqlOutputErrorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlOutputErrorPb struct {
	// The error message when execution fails.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlOutputErrorFromPb(pb *sqlOutputErrorPb) (*SqlOutputError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlOutputError{}
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

func (st *sqlOutputErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlOutputErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlQueryOutput struct {
	EndpointId string
	// The link to find the output results.
	OutputLink string
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	QueryText string
	// Information about SQL statements executed in the run.
	SqlStatements []SqlStatementOutput
	// The canonical identifier of the SQL warehouse.
	WarehouseId string

	ForceSendFields []string
}

func sqlQueryOutputToPb(st *SqlQueryOutput) (*sqlQueryOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlQueryOutputPb{}
	endpointIdPb, err := identity(&st.EndpointId)
	if err != nil {
		return nil, err
	}
	if endpointIdPb != nil {
		pb.EndpointId = *endpointIdPb
	}

	outputLinkPb, err := identity(&st.OutputLink)
	if err != nil {
		return nil, err
	}
	if outputLinkPb != nil {
		pb.OutputLink = *outputLinkPb
	}

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	var sqlStatementsPb []sqlStatementOutputPb
	for _, item := range st.SqlStatements {
		itemPb, err := sqlStatementOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			sqlStatementsPb = append(sqlStatementsPb, *itemPb)
		}
	}
	pb.SqlStatements = sqlStatementsPb

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SqlQueryOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlQueryOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlQueryOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlQueryOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlQueryOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlQueryOutputPb struct {
	EndpointId string `json:"endpoint_id,omitempty"`
	// The link to find the output results.
	OutputLink string `json:"output_link,omitempty"`
	// The text of the SQL query. Can Run permission of the SQL query is
	// required to view this field.
	QueryText string `json:"query_text,omitempty"`
	// Information about SQL statements executed in the run.
	SqlStatements []sqlStatementOutputPb `json:"sql_statements,omitempty"`
	// The canonical identifier of the SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlQueryOutputFromPb(pb *sqlQueryOutputPb) (*SqlQueryOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlQueryOutput{}
	endpointIdField, err := identity(&pb.EndpointId)
	if err != nil {
		return nil, err
	}
	if endpointIdField != nil {
		st.EndpointId = *endpointIdField
	}
	outputLinkField, err := identity(&pb.OutputLink)
	if err != nil {
		return nil, err
	}
	if outputLinkField != nil {
		st.OutputLink = *outputLinkField
	}
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}

	var sqlStatementsField []SqlStatementOutput
	for _, itemPb := range pb.SqlStatements {
		item, err := sqlStatementOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			sqlStatementsField = append(sqlStatementsField, *item)
		}
	}
	st.SqlStatements = sqlStatementsField
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlQueryOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlQueryOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlStatementOutput struct {
	// A key that can be used to look up query details.
	LookupKey string

	ForceSendFields []string
}

func sqlStatementOutputToPb(st *SqlStatementOutput) (*sqlStatementOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlStatementOutputPb{}
	lookupKeyPb, err := identity(&st.LookupKey)
	if err != nil {
		return nil, err
	}
	if lookupKeyPb != nil {
		pb.LookupKey = *lookupKeyPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SqlStatementOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlStatementOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlStatementOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlStatementOutput) MarshalJSON() ([]byte, error) {
	pb, err := sqlStatementOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlStatementOutputPb struct {
	// A key that can be used to look up query details.
	LookupKey string `json:"lookup_key,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlStatementOutputFromPb(pb *sqlStatementOutputPb) (*SqlStatementOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlStatementOutput{}
	lookupKeyField, err := identity(&pb.LookupKey)
	if err != nil {
		return nil, err
	}
	if lookupKeyField != nil {
		st.LookupKey = *lookupKeyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlStatementOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlStatementOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlTask struct {
	// If alert, indicates that this job must refresh a SQL alert.
	Alert *SqlTaskAlert
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	Dashboard *SqlTaskDashboard
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	File *SqlTaskFile
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	Parameters map[string]string
	// If query, indicates that this job must execute a SQL query.
	Query *SqlTaskQuery
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	WarehouseId string
}

func sqlTaskToPb(st *SqlTask) (*sqlTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskPb{}
	alertPb, err := sqlTaskAlertToPb(st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = alertPb
	}

	dashboardPb, err := sqlTaskDashboardToPb(st.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardPb != nil {
		pb.Dashboard = dashboardPb
	}

	filePb, err := sqlTaskFileToPb(st.File)
	if err != nil {
		return nil, err
	}
	if filePb != nil {
		pb.File = filePb
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

	queryPb, err := sqlTaskQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

	return pb, nil
}

func (st *SqlTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTask) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlTaskPb struct {
	// If alert, indicates that this job must refresh a SQL alert.
	Alert *sqlTaskAlertPb `json:"alert,omitempty"`
	// If dashboard, indicates that this job must refresh a SQL dashboard.
	Dashboard *sqlTaskDashboardPb `json:"dashboard,omitempty"`
	// If file, indicates that this job runs a SQL file in a remote Git
	// repository.
	File *sqlTaskFilePb `json:"file,omitempty"`
	// Parameters to be used for each run of this job. The SQL alert task does
	// not support custom parameters.
	Parameters map[string]string `json:"parameters,omitempty"`
	// If query, indicates that this job must execute a SQL query.
	Query *sqlTaskQueryPb `json:"query,omitempty"`
	// The canonical identifier of the SQL warehouse. Recommended to use with
	// serverless or pro SQL warehouses. Classic SQL warehouses are only
	// supported for SQL alert, dashboard and query tasks and are limited to
	// scheduled single-task jobs.
	WarehouseId string `json:"warehouse_id"`
}

func sqlTaskFromPb(pb *sqlTaskPb) (*SqlTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTask{}
	alertField, err := sqlTaskAlertFromPb(pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = alertField
	}
	dashboardField, err := sqlTaskDashboardFromPb(pb.Dashboard)
	if err != nil {
		return nil, err
	}
	if dashboardField != nil {
		st.Dashboard = dashboardField
	}
	fileField, err := sqlTaskFileFromPb(pb.File)
	if err != nil {
		return nil, err
	}
	if fileField != nil {
		st.File = fileField
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
	queryField, err := sqlTaskQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	return st, nil
}

type SqlTaskAlert struct {
	// The canonical identifier of the SQL alert.
	AlertId string
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions bool
	// If specified, alert notifications are sent to subscribers.
	Subscriptions []SqlTaskSubscription

	ForceSendFields []string
}

func sqlTaskAlertToPb(st *SqlTaskAlert) (*sqlTaskAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskAlertPb{}
	alertIdPb, err := identity(&st.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdPb != nil {
		pb.AlertId = *alertIdPb
	}

	pauseSubscriptionsPb, err := identity(&st.PauseSubscriptions)
	if err != nil {
		return nil, err
	}
	if pauseSubscriptionsPb != nil {
		pb.PauseSubscriptions = *pauseSubscriptionsPb
	}

	var subscriptionsPb []sqlTaskSubscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := sqlTaskSubscriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subscriptionsPb = append(subscriptionsPb, *itemPb)
		}
	}
	pb.Subscriptions = subscriptionsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SqlTaskAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskAlert) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlTaskAlertPb struct {
	// The canonical identifier of the SQL alert.
	AlertId string `json:"alert_id"`
	// If true, the alert notifications are not sent to subscribers.
	PauseSubscriptions bool `json:"pause_subscriptions,omitempty"`
	// If specified, alert notifications are sent to subscribers.
	Subscriptions []sqlTaskSubscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlTaskAlertFromPb(pb *sqlTaskAlertPb) (*SqlTaskAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskAlert{}
	alertIdField, err := identity(&pb.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdField != nil {
		st.AlertId = *alertIdField
	}
	pauseSubscriptionsField, err := identity(&pb.PauseSubscriptions)
	if err != nil {
		return nil, err
	}
	if pauseSubscriptionsField != nil {
		st.PauseSubscriptions = *pauseSubscriptionsField
	}

	var subscriptionsField []SqlTaskSubscription
	for _, itemPb := range pb.Subscriptions {
		item, err := sqlTaskSubscriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subscriptionsField = append(subscriptionsField, *item)
		}
	}
	st.Subscriptions = subscriptionsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlTaskAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlTaskAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlTaskDashboard struct {
	// Subject of the email sent to subscribers of this task.
	CustomSubject string
	// The canonical identifier of the SQL dashboard.
	DashboardId string
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	PauseSubscriptions bool
	// If specified, dashboard snapshots are sent to subscriptions.
	Subscriptions []SqlTaskSubscription

	ForceSendFields []string
}

func sqlTaskDashboardToPb(st *SqlTaskDashboard) (*sqlTaskDashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskDashboardPb{}
	customSubjectPb, err := identity(&st.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectPb != nil {
		pb.CustomSubject = *customSubjectPb
	}

	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	pauseSubscriptionsPb, err := identity(&st.PauseSubscriptions)
	if err != nil {
		return nil, err
	}
	if pauseSubscriptionsPb != nil {
		pb.PauseSubscriptions = *pauseSubscriptionsPb
	}

	var subscriptionsPb []sqlTaskSubscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := sqlTaskSubscriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subscriptionsPb = append(subscriptionsPb, *itemPb)
		}
	}
	pb.Subscriptions = subscriptionsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SqlTaskDashboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskDashboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskDashboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskDashboard) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskDashboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlTaskDashboardPb struct {
	// Subject of the email sent to subscribers of this task.
	CustomSubject string `json:"custom_subject,omitempty"`
	// The canonical identifier of the SQL dashboard.
	DashboardId string `json:"dashboard_id"`
	// If true, the dashboard snapshot is not taken, and emails are not sent to
	// subscribers.
	PauseSubscriptions bool `json:"pause_subscriptions,omitempty"`
	// If specified, dashboard snapshots are sent to subscriptions.
	Subscriptions []sqlTaskSubscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlTaskDashboardFromPb(pb *sqlTaskDashboardPb) (*SqlTaskDashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskDashboard{}
	customSubjectField, err := identity(&pb.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectField != nil {
		st.CustomSubject = *customSubjectField
	}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	pauseSubscriptionsField, err := identity(&pb.PauseSubscriptions)
	if err != nil {
		return nil, err
	}
	if pauseSubscriptionsField != nil {
		st.PauseSubscriptions = *pauseSubscriptionsField
	}

	var subscriptionsField []SqlTaskSubscription
	for _, itemPb := range pb.Subscriptions {
		item, err := sqlTaskSubscriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subscriptionsField = append(subscriptionsField, *item)
		}
	}
	st.Subscriptions = subscriptionsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlTaskDashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlTaskDashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlTaskFile struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	Path string
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local Databricks workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL
	// file is located in cloud Git provider.
	Source Source
}

func sqlTaskFileToPb(st *SqlTaskFile) (*sqlTaskFilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskFilePb{}
	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	sourcePb, err := identity(&st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = *sourcePb
	}

	return pb, nil
}

func (st *SqlTaskFile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskFilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskFileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskFile) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskFileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlTaskFilePb struct {
	// Path of the SQL file. Must be relative if the source is a remote Git
	// repository and absolute for workspace paths.
	Path string `json:"path"`
	// Optional location type of the SQL file. When set to `WORKSPACE`, the SQL
	// file will be retrieved from the local Databricks workspace. When set to
	// `GIT`, the SQL file will be retrieved from a Git repository defined in
	// `git_source`. If the value is empty, the task will use `GIT` if
	// `git_source` is defined and `WORKSPACE` otherwise.
	//
	// * `WORKSPACE`: SQL file is located in Databricks workspace. * `GIT`: SQL
	// file is located in cloud Git provider.
	Source Source `json:"source,omitempty"`
}

func sqlTaskFileFromPb(pb *sqlTaskFilePb) (*SqlTaskFile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskFile{}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}
	sourceField, err := identity(&pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = *sourceField
	}

	return st, nil
}

type SqlTaskQuery struct {
	// The canonical identifier of the SQL query.
	QueryId string
}

func sqlTaskQueryToPb(st *SqlTaskQuery) (*sqlTaskQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskQueryPb{}
	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	return pb, nil
}

func (st *SqlTaskQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskQuery) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlTaskQueryPb struct {
	// The canonical identifier of the SQL query.
	QueryId string `json:"query_id"`
}

func sqlTaskQueryFromPb(pb *sqlTaskQueryPb) (*SqlTaskQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskQuery{}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}

	return st, nil
}

type SqlTaskSubscription struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	DestinationId string
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	UserName string

	ForceSendFields []string
}

func sqlTaskSubscriptionToPb(st *SqlTaskSubscription) (*sqlTaskSubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlTaskSubscriptionPb{}
	destinationIdPb, err := identity(&st.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdPb != nil {
		pb.DestinationId = *destinationIdPb
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

func (st *SqlTaskSubscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlTaskSubscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlTaskSubscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlTaskSubscription) MarshalJSON() ([]byte, error) {
	pb, err := sqlTaskSubscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sqlTaskSubscriptionPb struct {
	// The canonical identifier of the destination to receive email
	// notification. This parameter is mutually exclusive with user_name. You
	// cannot set both destination_id and user_name for subscription
	// notifications.
	DestinationId string `json:"destination_id,omitempty"`
	// The user name to receive the subscription email. This parameter is
	// mutually exclusive with destination_id. You cannot set both
	// destination_id and user_name for subscription notifications.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlTaskSubscriptionFromPb(pb *sqlTaskSubscriptionPb) (*SqlTaskSubscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlTaskSubscription{}
	destinationIdField, err := identity(&pb.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdField != nil {
		st.DestinationId = *destinationIdField
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

func (st *sqlTaskSubscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlTaskSubscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StorageMode string
type storageModePb string

const StorageModeDirectQuery StorageMode = `DIRECT_QUERY`

const StorageModeDual StorageMode = `DUAL`

const StorageModeImport StorageMode = `IMPORT`

// String representation for [fmt.Print]
func (f *StorageMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StorageMode) Set(v string) error {
	switch v {
	case `DIRECT_QUERY`, `DUAL`, `IMPORT`:
		*f = StorageMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DIRECT_QUERY", "DUAL", "IMPORT"`, v)
	}
}

// Type always returns StorageMode to satisfy [pflag.Value] interface
func (f *StorageMode) Type() string {
	return "StorageMode"
}

func storageModeToPb(st *StorageMode) (*storageModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := storageModePb(*st)
	return &pb, nil
}

func storageModeFromPb(pb *storageModePb) (*StorageMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := StorageMode(*pb)
	return &st, nil
}

type SubmitRun struct {
	// List of permissions to set on the job.
	AccessControlList []JobAccessControlRequest
	// The user specified id of the budget policy to use for this one-time run.
	// If not specified, the run will be not be attributed to any budget policy.
	BudgetPolicyId string
	// An optional set of email addresses notified when the run begins or
	// completes.
	EmailNotifications *JobEmailNotifications
	// A list of task execution environment specifications that can be
	// referenced by tasks of this run.
	Environments []JobEnvironment
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *GitSource
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules
	// An optional token that can be used to guarantee the idempotency of job
	// run requests. If a run with the provided token already exists, the
	// request does not create a new run but returns the ID of the existing run
	// instead. If a run with the provided token is deleted, an error is
	// returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	IdempotencyToken string
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	NotificationSettings *JobNotificationSettings
	// The queue settings of the one-time run.
	Queue *QueueSettings
	// Specifies the user or service principal that the job runs as. If not
	// specified, the job runs as the user who submits the request.
	RunAs *JobRunAs
	// An optional name for the run. The default value is `Untitled`.
	RunName string

	Tasks []SubmitTask
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string
}

func submitRunToPb(st *SubmitRun) (*submitRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &submitRunPb{}

	var accessControlListPb []jobAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := jobAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	emailNotificationsPb, err := jobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	var environmentsPb []jobEnvironmentPb
	for _, item := range st.Environments {
		itemPb, err := jobEnvironmentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentsPb = append(environmentsPb, *itemPb)
		}
	}
	pb.Environments = environmentsPb

	gitSourcePb, err := gitSourceToPb(st.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourcePb != nil {
		pb.GitSource = gitSourcePb
	}

	healthPb, err := jobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	idempotencyTokenPb, err := identity(&st.IdempotencyToken)
	if err != nil {
		return nil, err
	}
	if idempotencyTokenPb != nil {
		pb.IdempotencyToken = *idempotencyTokenPb
	}

	notificationSettingsPb, err := jobNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	queuePb, err := queueSettingsToPb(st.Queue)
	if err != nil {
		return nil, err
	}
	if queuePb != nil {
		pb.Queue = queuePb
	}

	runAsPb, err := jobRunAsToPb(st.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsPb != nil {
		pb.RunAs = runAsPb
	}

	runNamePb, err := identity(&st.RunName)
	if err != nil {
		return nil, err
	}
	if runNamePb != nil {
		pb.RunName = *runNamePb
	}

	var tasksPb []submitTaskPb
	for _, item := range st.Tasks {
		itemPb, err := submitTaskToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tasksPb = append(tasksPb, *itemPb)
		}
	}
	pb.Tasks = tasksPb

	timeoutSecondsPb, err := identity(&st.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsPb != nil {
		pb.TimeoutSeconds = *timeoutSecondsPb
	}

	webhookNotificationsPb, err := webhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SubmitRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &submitRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := submitRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubmitRun) MarshalJSON() ([]byte, error) {
	pb, err := submitRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type submitRunPb struct {
	// List of permissions to set on the job.
	AccessControlList []jobAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The user specified id of the budget policy to use for this one-time run.
	// If not specified, the run will be not be attributed to any budget policy.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// An optional set of email addresses notified when the run begins or
	// completes.
	EmailNotifications *jobEmailNotificationsPb `json:"email_notifications,omitempty"`
	// A list of task execution environment specifications that can be
	// referenced by tasks of this run.
	Environments []jobEnvironmentPb `json:"environments,omitempty"`
	// An optional specification for a remote Git repository containing the
	// source code used by tasks. Version-controlled source code is supported by
	// notebook, dbt, Python script, and SQL File tasks.
	//
	// If `git_source` is set, these tasks retrieve the file from the remote
	// repository by default. However, this behavior can be overridden by
	// setting `source` to `WORKSPACE` on the task.
	//
	// Note: dbt and SQL File tasks support only version-controlled sources. If
	// dbt or SQL File tasks are used, `git_source` must be defined on the job.
	GitSource *gitSourcePb `json:"git_source,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *jobsHealthRulesPb `json:"health,omitempty"`
	// An optional token that can be used to guarantee the idempotency of job
	// run requests. If a run with the provided token already exists, the
	// request does not create a new run but returns the ID of the existing run
	// instead. If a run with the provided token is deleted, an error is
	// returned.
	//
	// If you specify the idempotency token, upon failure you can retry until
	// the request succeeds. Databricks guarantees that exactly one run is
	// launched with that idempotency token.
	//
	// This token must have at most 64 characters.
	//
	// For more information, see [How to ensure idempotency for jobs].
	//
	// [How to ensure idempotency for jobs]: https://kb.databricks.com/jobs/jobs-idempotency.html
	IdempotencyToken string `json:"idempotency_token,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// run.
	NotificationSettings *jobNotificationSettingsPb `json:"notification_settings,omitempty"`
	// The queue settings of the one-time run.
	Queue *queueSettingsPb `json:"queue,omitempty"`
	// Specifies the user or service principal that the job runs as. If not
	// specified, the job runs as the user who submits the request.
	RunAs *jobRunAsPb `json:"run_as,omitempty"`
	// An optional name for the run. The default value is `Untitled`.
	RunName string `json:"run_name,omitempty"`

	Tasks []submitTaskPb `json:"tasks,omitempty"`
	// An optional timeout applied to each run of this job. A value of `0` means
	// no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes.
	WebhookNotifications *webhookNotificationsPb `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func submitRunFromPb(pb *submitRunPb) (*SubmitRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitRun{}

	var accessControlListField []JobAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := jobAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	emailNotificationsField, err := jobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}

	var environmentsField []JobEnvironment
	for _, itemPb := range pb.Environments {
		item, err := jobEnvironmentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentsField = append(environmentsField, *item)
		}
	}
	st.Environments = environmentsField
	gitSourceField, err := gitSourceFromPb(pb.GitSource)
	if err != nil {
		return nil, err
	}
	if gitSourceField != nil {
		st.GitSource = gitSourceField
	}
	healthField, err := jobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	idempotencyTokenField, err := identity(&pb.IdempotencyToken)
	if err != nil {
		return nil, err
	}
	if idempotencyTokenField != nil {
		st.IdempotencyToken = *idempotencyTokenField
	}
	notificationSettingsField, err := jobNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	queueField, err := queueSettingsFromPb(pb.Queue)
	if err != nil {
		return nil, err
	}
	if queueField != nil {
		st.Queue = queueField
	}
	runAsField, err := jobRunAsFromPb(pb.RunAs)
	if err != nil {
		return nil, err
	}
	if runAsField != nil {
		st.RunAs = runAsField
	}
	runNameField, err := identity(&pb.RunName)
	if err != nil {
		return nil, err
	}
	if runNameField != nil {
		st.RunName = *runNameField
	}

	var tasksField []SubmitTask
	for _, itemPb := range pb.Tasks {
		item, err := submitTaskFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tasksField = append(tasksField, *item)
		}
	}
	st.Tasks = tasksField
	timeoutSecondsField, err := identity(&pb.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsField != nil {
		st.TimeoutSeconds = *timeoutSecondsField
	}
	webhookNotificationsField, err := webhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *submitRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st submitRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Run was created and started successfully.
type SubmitRunResponse struct {
	// The canonical identifier for the newly submitted run.
	RunId int64

	ForceSendFields []string
}

func submitRunResponseToPb(st *SubmitRunResponse) (*submitRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &submitRunResponsePb{}
	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SubmitRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &submitRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := submitRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubmitRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := submitRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type submitRunResponsePb struct {
	// The canonical identifier for the newly submitted run.
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func submitRunResponseFromPb(pb *submitRunResponsePb) (*SubmitRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitRunResponse{}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *submitRunResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st submitRunResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubmitTask struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask *CleanRoomsNotebookTask
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *ConditionTask
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *DashboardTask
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *DbtTask
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []TaskDependency
	// An optional description for this task.
	Description string
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *JobEmailNotifications
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey string
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *ForEachTask

	GenAiComputeTask *GenAiComputeTask
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *NotebookTask
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *TaskNotificationSettings
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *PipelineTask
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *PowerBiTask
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *PythonWheelTask
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *RunJobTask
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *SparkJarTask
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *SparkPythonTask
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *SparkSubmitTask
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *SqlTask
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string
}

func submitTaskToPb(st *SubmitTask) (*submitTaskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &submitTaskPb{}
	cleanRoomsNotebookTaskPb, err := cleanRoomsNotebookTaskToPb(st.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskPb != nil {
		pb.CleanRoomsNotebookTask = cleanRoomsNotebookTaskPb
	}

	conditionTaskPb, err := conditionTaskToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}

	dashboardTaskPb, err := dashboardTaskToPb(st.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskPb != nil {
		pb.DashboardTask = dashboardTaskPb
	}

	dbtTaskPb, err := dbtTaskToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}

	var dependsOnPb []taskDependencyPb
	for _, item := range st.DependsOn {
		itemPb, err := taskDependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependsOnPb = append(dependsOnPb, *itemPb)
		}
	}
	pb.DependsOn = dependsOnPb

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	emailNotificationsPb, err := jobEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	environmentKeyPb, err := identity(&st.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyPb != nil {
		pb.EnvironmentKey = *environmentKeyPb
	}

	existingClusterIdPb, err := identity(&st.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdPb != nil {
		pb.ExistingClusterId = *existingClusterIdPb
	}

	forEachTaskPb, err := forEachTaskToPb(st.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskPb != nil {
		pb.ForEachTask = forEachTaskPb
	}

	genAiComputeTaskPb, err := genAiComputeTaskToPb(st.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskPb != nil {
		pb.GenAiComputeTask = genAiComputeTaskPb
	}

	healthPb, err := jobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	var librariesPb []compute.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}

	notebookTaskPb, err := notebookTaskToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}

	notificationSettingsPb, err := taskNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	pipelineTaskPb, err := pipelineTaskToPb(st.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskPb != nil {
		pb.PipelineTask = pipelineTaskPb
	}

	powerBiTaskPb, err := powerBiTaskToPb(st.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskPb != nil {
		pb.PowerBiTask = powerBiTaskPb
	}

	pythonWheelTaskPb, err := pythonWheelTaskToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}

	runIfPb, err := identity(&st.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfPb != nil {
		pb.RunIf = *runIfPb
	}

	runJobTaskPb, err := runJobTaskToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}

	sparkJarTaskPb, err := sparkJarTaskToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}

	sparkPythonTaskPb, err := sparkPythonTaskToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}

	sparkSubmitTaskPb, err := sparkSubmitTaskToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}

	sqlTaskPb, err := sqlTaskToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}

	taskKeyPb, err := identity(&st.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyPb != nil {
		pb.TaskKey = *taskKeyPb
	}

	timeoutSecondsPb, err := identity(&st.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsPb != nil {
		pb.TimeoutSeconds = *timeoutSecondsPb
	}

	webhookNotificationsPb, err := webhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SubmitTask) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &submitTaskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := submitTaskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubmitTask) MarshalJSON() ([]byte, error) {
	pb, err := submitTaskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type submitTaskPb struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask *cleanRoomsNotebookTaskPb `json:"clean_rooms_notebook_task,omitempty"`
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *conditionTaskPb `json:"condition_task,omitempty"`
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *dashboardTaskPb `json:"dashboard_task,omitempty"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *dbtTaskPb `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete successfully before
	// executing this task. The key is `task_key`, and the value is the name
	// assigned to the dependent task.
	DependsOn []taskDependencyPb `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// An optional set of email addresses notified when the task run begins or
	// completes. The default behavior is to not send any emails.
	EmailNotifications *jobEmailNotificationsPb `json:"email_notifications,omitempty"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey string `json:"environment_key,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *forEachTaskPb `json:"for_each_task,omitempty"`

	GenAiComputeTask *genAiComputeTaskPb `json:"gen_ai_compute_task,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *jobsHealthRulesPb `json:"health,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.LibraryPb `json:"libraries,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpecPb `json:"new_cluster,omitempty"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *notebookTaskPb `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task run.
	NotificationSettings *taskNotificationSettingsPb `json:"notification_settings,omitempty"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *pipelineTaskPb `json:"pipeline_task,omitempty"`
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *powerBiTaskPb `json:"power_bi_task,omitempty"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *pythonWheelTaskPb `json:"python_wheel_task,omitempty"`
	// An optional value indicating the condition that determines whether the
	// task should be run once its dependencies have been completed. When
	// omitted, defaults to `ALL_SUCCESS`. See :method:jobs/create for a list of
	// possible values.
	RunIf RunIf `json:"run_if,omitempty"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *runJobTaskPb `json:"run_job_task,omitempty"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *sparkJarTaskPb `json:"spark_jar_task,omitempty"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *sparkPythonTaskPb `json:"spark_python_task,omitempty"`
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *sparkSubmitTaskPb `json:"spark_submit_task,omitempty"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *sqlTaskPb `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when the run begins or
	// completes. The default behavior is to not send any system notifications.
	// Task webhooks respect the task notification settings.
	WebhookNotifications *webhookNotificationsPb `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func submitTaskFromPb(pb *submitTaskPb) (*SubmitTask, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubmitTask{}
	cleanRoomsNotebookTaskField, err := cleanRoomsNotebookTaskFromPb(pb.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskField != nil {
		st.CleanRoomsNotebookTask = cleanRoomsNotebookTaskField
	}
	conditionTaskField, err := conditionTaskFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dashboardTaskField, err := dashboardTaskFromPb(pb.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskField != nil {
		st.DashboardTask = dashboardTaskField
	}
	dbtTaskField, err := dbtTaskFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}

	var dependsOnField []TaskDependency
	for _, itemPb := range pb.DependsOn {
		item, err := taskDependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependsOnField = append(dependsOnField, *item)
		}
	}
	st.DependsOn = dependsOnField
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	emailNotificationsField, err := jobEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}
	environmentKeyField, err := identity(&pb.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyField != nil {
		st.EnvironmentKey = *environmentKeyField
	}
	existingClusterIdField, err := identity(&pb.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdField != nil {
		st.ExistingClusterId = *existingClusterIdField
	}
	forEachTaskField, err := forEachTaskFromPb(pb.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskField != nil {
		st.ForEachTask = forEachTaskField
	}
	genAiComputeTaskField, err := genAiComputeTaskFromPb(pb.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskField != nil {
		st.GenAiComputeTask = genAiComputeTaskField
	}
	healthField, err := jobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}
	notebookTaskField, err := notebookTaskFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	notificationSettingsField, err := taskNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	pipelineTaskField, err := pipelineTaskFromPb(pb.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskField != nil {
		st.PipelineTask = pipelineTaskField
	}
	powerBiTaskField, err := powerBiTaskFromPb(pb.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskField != nil {
		st.PowerBiTask = powerBiTaskField
	}
	pythonWheelTaskField, err := pythonWheelTaskFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	runIfField, err := identity(&pb.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfField != nil {
		st.RunIf = *runIfField
	}
	runJobTaskField, err := runJobTaskFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	sparkJarTaskField, err := sparkJarTaskFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := sparkPythonTaskFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := sparkSubmitTaskFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := sqlTaskFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}
	taskKeyField, err := identity(&pb.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyField != nil {
		st.TaskKey = *taskKeyField
	}
	timeoutSecondsField, err := identity(&pb.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsField != nil {
		st.TimeoutSeconds = *timeoutSecondsField
	}
	webhookNotificationsField, err := webhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *submitTaskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st submitTaskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Subscription struct {
	// Optional: Allows users to specify a custom subject line on the email sent
	// to subscribers.
	CustomSubject string
	// When true, the subscription will not send emails.
	Paused bool
	// The list of subscribers to send the snapshot of the dashboard to.
	Subscribers []SubscriptionSubscriber

	ForceSendFields []string
}

func subscriptionToPb(st *Subscription) (*subscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionPb{}
	customSubjectPb, err := identity(&st.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectPb != nil {
		pb.CustomSubject = *customSubjectPb
	}

	pausedPb, err := identity(&st.Paused)
	if err != nil {
		return nil, err
	}
	if pausedPb != nil {
		pb.Paused = *pausedPb
	}

	var subscribersPb []subscriptionSubscriberPb
	for _, item := range st.Subscribers {
		itemPb, err := subscriptionSubscriberToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subscribersPb = append(subscribersPb, *itemPb)
		}
	}
	pb.Subscribers = subscribersPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Subscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Subscription) MarshalJSON() ([]byte, error) {
	pb, err := subscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type subscriptionPb struct {
	// Optional: Allows users to specify a custom subject line on the email sent
	// to subscribers.
	CustomSubject string `json:"custom_subject,omitempty"`
	// When true, the subscription will not send emails.
	Paused bool `json:"paused,omitempty"`
	// The list of subscribers to send the snapshot of the dashboard to.
	Subscribers []subscriptionSubscriberPb `json:"subscribers,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func subscriptionFromPb(pb *subscriptionPb) (*Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Subscription{}
	customSubjectField, err := identity(&pb.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectField != nil {
		st.CustomSubject = *customSubjectField
	}
	pausedField, err := identity(&pb.Paused)
	if err != nil {
		return nil, err
	}
	if pausedField != nil {
		st.Paused = *pausedField
	}

	var subscribersField []SubscriptionSubscriber
	for _, itemPb := range pb.Subscribers {
		item, err := subscriptionSubscriberFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subscribersField = append(subscribersField, *item)
		}
	}
	st.Subscribers = subscribersField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *subscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st subscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SubscriptionSubscriber struct {
	// A snapshot of the dashboard will be sent to the destination when the
	// `destination_id` field is present.
	DestinationId string
	// A snapshot of the dashboard will be sent to the user's email when the
	// `user_name` field is present.
	UserName string

	ForceSendFields []string
}

func subscriptionSubscriberToPb(st *SubscriptionSubscriber) (*subscriptionSubscriberPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &subscriptionSubscriberPb{}
	destinationIdPb, err := identity(&st.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdPb != nil {
		pb.DestinationId = *destinationIdPb
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

func (st *SubscriptionSubscriber) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &subscriptionSubscriberPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := subscriptionSubscriberFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SubscriptionSubscriber) MarshalJSON() ([]byte, error) {
	pb, err := subscriptionSubscriberToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type subscriptionSubscriberPb struct {
	// A snapshot of the dashboard will be sent to the destination when the
	// `destination_id` field is present.
	DestinationId string `json:"destination_id,omitempty"`
	// A snapshot of the dashboard will be sent to the user's email when the
	// `user_name` field is present.
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func subscriptionSubscriberFromPb(pb *subscriptionSubscriberPb) (*SubscriptionSubscriber, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SubscriptionSubscriber{}
	destinationIdField, err := identity(&pb.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdField != nil {
		st.DestinationId = *destinationIdField
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

func (st *subscriptionSubscriberPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st subscriptionSubscriberPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableUpdateTriggerConfiguration struct {
	// The table(s) condition based on which to trigger a job run.
	Condition Condition
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	MinTimeBetweenTriggersSeconds int
	// A list of Delta tables to monitor for changes. The table name must be in
	// the format `catalog_name.schema_name.table_name`.
	TableNames []string
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	WaitAfterLastChangeSeconds int

	ForceSendFields []string
}

func tableUpdateTriggerConfigurationToPb(st *TableUpdateTriggerConfiguration) (*tableUpdateTriggerConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableUpdateTriggerConfigurationPb{}
	conditionPb, err := identity(&st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = *conditionPb
	}

	minTimeBetweenTriggersSecondsPb, err := identity(&st.MinTimeBetweenTriggersSeconds)
	if err != nil {
		return nil, err
	}
	if minTimeBetweenTriggersSecondsPb != nil {
		pb.MinTimeBetweenTriggersSeconds = *minTimeBetweenTriggersSecondsPb
	}

	var tableNamesPb []string
	for _, item := range st.TableNames {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tableNamesPb = append(tableNamesPb, *itemPb)
		}
	}
	pb.TableNames = tableNamesPb

	waitAfterLastChangeSecondsPb, err := identity(&st.WaitAfterLastChangeSeconds)
	if err != nil {
		return nil, err
	}
	if waitAfterLastChangeSecondsPb != nil {
		pb.WaitAfterLastChangeSeconds = *waitAfterLastChangeSecondsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TableUpdateTriggerConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tableUpdateTriggerConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableUpdateTriggerConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TableUpdateTriggerConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := tableUpdateTriggerConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tableUpdateTriggerConfigurationPb struct {
	// The table(s) condition based on which to trigger a job run.
	Condition Condition `json:"condition,omitempty"`
	// If set, the trigger starts a run only after the specified amount of time
	// has passed since the last time the trigger fired. The minimum allowed
	// value is 60 seconds.
	MinTimeBetweenTriggersSeconds int `json:"min_time_between_triggers_seconds,omitempty"`
	// A list of Delta tables to monitor for changes. The table name must be in
	// the format `catalog_name.schema_name.table_name`.
	TableNames []string `json:"table_names,omitempty"`
	// If set, the trigger starts a run only after no table updates have
	// occurred for the specified time and can be used to wait for a series of
	// table updates before triggering a run. The minimum allowed value is 60
	// seconds.
	WaitAfterLastChangeSeconds int `json:"wait_after_last_change_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableUpdateTriggerConfigurationFromPb(pb *tableUpdateTriggerConfigurationPb) (*TableUpdateTriggerConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableUpdateTriggerConfiguration{}
	conditionField, err := identity(&pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = *conditionField
	}
	minTimeBetweenTriggersSecondsField, err := identity(&pb.MinTimeBetweenTriggersSeconds)
	if err != nil {
		return nil, err
	}
	if minTimeBetweenTriggersSecondsField != nil {
		st.MinTimeBetweenTriggersSeconds = *minTimeBetweenTriggersSecondsField
	}

	var tableNamesField []string
	for _, itemPb := range pb.TableNames {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tableNamesField = append(tableNamesField, *item)
		}
	}
	st.TableNames = tableNamesField
	waitAfterLastChangeSecondsField, err := identity(&pb.WaitAfterLastChangeSeconds)
	if err != nil {
		return nil, err
	}
	if waitAfterLastChangeSecondsField != nil {
		st.WaitAfterLastChangeSeconds = *waitAfterLastChangeSecondsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableUpdateTriggerConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableUpdateTriggerConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Task struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask *CleanRoomsNotebookTask
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *ConditionTask
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *DashboardTask
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *DbtTask
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	DependsOn []TaskDependency
	// An optional description for this task.
	Description string
	// An option to disable auto optimization in serverless
	DisableAutoOptimization bool
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *TaskEmailNotifications
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey string
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *ForEachTask

	GenAiComputeTask *GenAiComputeTask
	// An optional set of health rules that can be defined for this job.
	Health *JobsHealthRules
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.Library
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries int
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis int
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpec
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *NotebookTask
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	NotificationSettings *TaskNotificationSettings
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *PipelineTask
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *PowerBiTask
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *PythonWheelTask
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout bool
	// An optional value specifying the condition determining whether the task
	// is run once its dependencies have been completed.
	//
	// * `ALL_SUCCESS`: All dependencies have executed and succeeded *
	// `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
	// `NONE_FAILED`: None of the dependencies have failed and at least one was
	// executed * `ALL_DONE`: All dependencies have been completed *
	// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
	// dependencies have failed
	RunIf RunIf
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *RunJobTask
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *SparkJarTask
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *SparkPythonTask
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *SparkSubmitTask
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *SqlTask
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	WebhookNotifications *WebhookNotifications

	ForceSendFields []string
}

func taskToPb(st *Task) (*taskPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskPb{}
	cleanRoomsNotebookTaskPb, err := cleanRoomsNotebookTaskToPb(st.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskPb != nil {
		pb.CleanRoomsNotebookTask = cleanRoomsNotebookTaskPb
	}

	conditionTaskPb, err := conditionTaskToPb(st.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskPb != nil {
		pb.ConditionTask = conditionTaskPb
	}

	dashboardTaskPb, err := dashboardTaskToPb(st.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskPb != nil {
		pb.DashboardTask = dashboardTaskPb
	}

	dbtTaskPb, err := dbtTaskToPb(st.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskPb != nil {
		pb.DbtTask = dbtTaskPb
	}

	var dependsOnPb []taskDependencyPb
	for _, item := range st.DependsOn {
		itemPb, err := taskDependencyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dependsOnPb = append(dependsOnPb, *itemPb)
		}
	}
	pb.DependsOn = dependsOnPb

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	disableAutoOptimizationPb, err := identity(&st.DisableAutoOptimization)
	if err != nil {
		return nil, err
	}
	if disableAutoOptimizationPb != nil {
		pb.DisableAutoOptimization = *disableAutoOptimizationPb
	}

	emailNotificationsPb, err := taskEmailNotificationsToPb(st.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsPb != nil {
		pb.EmailNotifications = emailNotificationsPb
	}

	environmentKeyPb, err := identity(&st.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyPb != nil {
		pb.EnvironmentKey = *environmentKeyPb
	}

	existingClusterIdPb, err := identity(&st.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdPb != nil {
		pb.ExistingClusterId = *existingClusterIdPb
	}

	forEachTaskPb, err := forEachTaskToPb(st.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskPb != nil {
		pb.ForEachTask = forEachTaskPb
	}

	genAiComputeTaskPb, err := genAiComputeTaskToPb(st.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskPb != nil {
		pb.GenAiComputeTask = genAiComputeTaskPb
	}

	healthPb, err := jobsHealthRulesToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	jobClusterKeyPb, err := identity(&st.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyPb != nil {
		pb.JobClusterKey = *jobClusterKeyPb
	}

	var librariesPb []compute.LibraryPb
	for _, item := range st.Libraries {
		itemPb, err := compute.LibraryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			librariesPb = append(librariesPb, *itemPb)
		}
	}
	pb.Libraries = librariesPb

	maxRetriesPb, err := identity(&st.MaxRetries)
	if err != nil {
		return nil, err
	}
	if maxRetriesPb != nil {
		pb.MaxRetries = *maxRetriesPb
	}

	minRetryIntervalMillisPb, err := identity(&st.MinRetryIntervalMillis)
	if err != nil {
		return nil, err
	}
	if minRetryIntervalMillisPb != nil {
		pb.MinRetryIntervalMillis = *minRetryIntervalMillisPb
	}

	newClusterPb, err := compute.ClusterSpecToPb(st.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterPb != nil {
		pb.NewCluster = newClusterPb
	}

	notebookTaskPb, err := notebookTaskToPb(st.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskPb != nil {
		pb.NotebookTask = notebookTaskPb
	}

	notificationSettingsPb, err := taskNotificationSettingsToPb(st.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsPb != nil {
		pb.NotificationSettings = notificationSettingsPb
	}

	pipelineTaskPb, err := pipelineTaskToPb(st.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskPb != nil {
		pb.PipelineTask = pipelineTaskPb
	}

	powerBiTaskPb, err := powerBiTaskToPb(st.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskPb != nil {
		pb.PowerBiTask = powerBiTaskPb
	}

	pythonWheelTaskPb, err := pythonWheelTaskToPb(st.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskPb != nil {
		pb.PythonWheelTask = pythonWheelTaskPb
	}

	retryOnTimeoutPb, err := identity(&st.RetryOnTimeout)
	if err != nil {
		return nil, err
	}
	if retryOnTimeoutPb != nil {
		pb.RetryOnTimeout = *retryOnTimeoutPb
	}

	runIfPb, err := identity(&st.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfPb != nil {
		pb.RunIf = *runIfPb
	}

	runJobTaskPb, err := runJobTaskToPb(st.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskPb != nil {
		pb.RunJobTask = runJobTaskPb
	}

	sparkJarTaskPb, err := sparkJarTaskToPb(st.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskPb != nil {
		pb.SparkJarTask = sparkJarTaskPb
	}

	sparkPythonTaskPb, err := sparkPythonTaskToPb(st.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskPb != nil {
		pb.SparkPythonTask = sparkPythonTaskPb
	}

	sparkSubmitTaskPb, err := sparkSubmitTaskToPb(st.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskPb != nil {
		pb.SparkSubmitTask = sparkSubmitTaskPb
	}

	sqlTaskPb, err := sqlTaskToPb(st.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskPb != nil {
		pb.SqlTask = sqlTaskPb
	}

	taskKeyPb, err := identity(&st.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyPb != nil {
		pb.TaskKey = *taskKeyPb
	}

	timeoutSecondsPb, err := identity(&st.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsPb != nil {
		pb.TimeoutSeconds = *timeoutSecondsPb
	}

	webhookNotificationsPb, err := webhookNotificationsToPb(st.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsPb != nil {
		pb.WebhookNotifications = webhookNotificationsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Task) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Task) MarshalJSON() ([]byte, error) {
	pb, err := taskToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type taskPb struct {
	// The task runs a [clean rooms] notebook when the
	// `clean_rooms_notebook_task` field is present.
	//
	// [clean rooms]: https://docs.databricks.com/en/clean-rooms/index.html
	CleanRoomsNotebookTask *cleanRoomsNotebookTaskPb `json:"clean_rooms_notebook_task,omitempty"`
	// The task evaluates a condition that can be used to control the execution
	// of other tasks when the `condition_task` field is present. The condition
	// task does not require a cluster to execute and does not support retries
	// or notifications.
	ConditionTask *conditionTaskPb `json:"condition_task,omitempty"`
	// The task refreshes a dashboard and sends a snapshot to subscribers.
	DashboardTask *dashboardTaskPb `json:"dashboard_task,omitempty"`
	// The task runs one or more dbt commands when the `dbt_task` field is
	// present. The dbt task requires both Databricks SQL and the ability to use
	// a serverless or a pro SQL warehouse.
	DbtTask *dbtTaskPb `json:"dbt_task,omitempty"`
	// An optional array of objects specifying the dependency graph of the task.
	// All tasks specified in this field must complete before executing this
	// task. The task will run only if the `run_if` condition is true. The key
	// is `task_key`, and the value is the name assigned to the dependent task.
	DependsOn []taskDependencyPb `json:"depends_on,omitempty"`
	// An optional description for this task.
	Description string `json:"description,omitempty"`
	// An option to disable auto optimization in serverless
	DisableAutoOptimization bool `json:"disable_auto_optimization,omitempty"`
	// An optional set of email addresses that is notified when runs of this
	// task begin or complete as well as when this task is deleted. The default
	// behavior is to not send any emails.
	EmailNotifications *taskEmailNotificationsPb `json:"email_notifications,omitempty"`
	// The key that references an environment spec in a job. This field is
	// required for Python script, Python wheel and dbt tasks when using
	// serverless compute.
	EnvironmentKey string `json:"environment_key,omitempty"`
	// If existing_cluster_id, the ID of an existing cluster that is used for
	// all runs. When running jobs or tasks on an existing cluster, you may need
	// to manually restart the cluster if it stops responding. We suggest
	// running jobs and tasks on new clusters for greater reliability
	ExistingClusterId string `json:"existing_cluster_id,omitempty"`
	// The task executes a nested task for every input provided when the
	// `for_each_task` field is present.
	ForEachTask *forEachTaskPb `json:"for_each_task,omitempty"`

	GenAiComputeTask *genAiComputeTaskPb `json:"gen_ai_compute_task,omitempty"`
	// An optional set of health rules that can be defined for this job.
	Health *jobsHealthRulesPb `json:"health,omitempty"`
	// If job_cluster_key, this task is executed reusing the cluster specified
	// in `job.settings.job_clusters`.
	JobClusterKey string `json:"job_cluster_key,omitempty"`
	// An optional list of libraries to be installed on the cluster. The default
	// value is an empty list.
	Libraries []compute.LibraryPb `json:"libraries,omitempty"`
	// An optional maximum number of times to retry an unsuccessful run. A run
	// is considered to be unsuccessful if it completes with the `FAILED`
	// result_state or `INTERNAL_ERROR` `life_cycle_state`. The value `-1` means
	// to retry indefinitely and the value `0` means to never retry.
	MaxRetries int `json:"max_retries,omitempty"`
	// An optional minimal interval in milliseconds between the start of the
	// failed run and the subsequent retry run. The default behavior is that
	// unsuccessful runs are immediately retried.
	MinRetryIntervalMillis int `json:"min_retry_interval_millis,omitempty"`
	// If new_cluster, a description of a new cluster that is created for each
	// run.
	NewCluster *compute.ClusterSpecPb `json:"new_cluster,omitempty"`
	// The task runs a notebook when the `notebook_task` field is present.
	NotebookTask *notebookTaskPb `json:"notebook_task,omitempty"`
	// Optional notification settings that are used when sending notifications
	// to each of the `email_notifications` and `webhook_notifications` for this
	// task.
	NotificationSettings *taskNotificationSettingsPb `json:"notification_settings,omitempty"`
	// The task triggers a pipeline update when the `pipeline_task` field is
	// present. Only pipelines configured to use triggered more are supported.
	PipelineTask *pipelineTaskPb `json:"pipeline_task,omitempty"`
	// The task triggers a Power BI semantic model update when the
	// `power_bi_task` field is present.
	PowerBiTask *powerBiTaskPb `json:"power_bi_task,omitempty"`
	// The task runs a Python wheel when the `python_wheel_task` field is
	// present.
	PythonWheelTask *pythonWheelTaskPb `json:"python_wheel_task,omitempty"`
	// An optional policy to specify whether to retry a job when it times out.
	// The default behavior is to not retry on timeout.
	RetryOnTimeout bool `json:"retry_on_timeout,omitempty"`
	// An optional value specifying the condition determining whether the task
	// is run once its dependencies have been completed.
	//
	// * `ALL_SUCCESS`: All dependencies have executed and succeeded *
	// `AT_LEAST_ONE_SUCCESS`: At least one dependency has succeeded *
	// `NONE_FAILED`: None of the dependencies have failed and at least one was
	// executed * `ALL_DONE`: All dependencies have been completed *
	// `AT_LEAST_ONE_FAILED`: At least one dependency failed * `ALL_FAILED`: ALl
	// dependencies have failed
	RunIf RunIf `json:"run_if,omitempty"`
	// The task triggers another job when the `run_job_task` field is present.
	RunJobTask *runJobTaskPb `json:"run_job_task,omitempty"`
	// The task runs a JAR when the `spark_jar_task` field is present.
	SparkJarTask *sparkJarTaskPb `json:"spark_jar_task,omitempty"`
	// The task runs a Python file when the `spark_python_task` field is
	// present.
	SparkPythonTask *sparkPythonTaskPb `json:"spark_python_task,omitempty"`
	// (Legacy) The task runs the spark-submit script when the
	// `spark_submit_task` field is present. This task can run only on new
	// clusters and is not compatible with serverless compute.
	//
	// In the `new_cluster` specification, `libraries` and `spark_conf` are not
	// supported. Instead, use `--jars` and `--py-files` to add Java and Python
	// libraries and `--conf` to set the Spark configurations.
	//
	// `master`, `deploy-mode`, and `executor-cores` are automatically
	// configured by Databricks; you _cannot_ specify them in parameters.
	//
	// By default, the Spark submit job uses all available memory (excluding
	// reserved memory for Databricks services). You can set `--driver-memory`,
	// and `--executor-memory` to a smaller value to leave some room for
	// off-heap usage.
	//
	// The `--jars`, `--py-files`, `--files` arguments support DBFS and S3
	// paths.
	SparkSubmitTask *sparkSubmitTaskPb `json:"spark_submit_task,omitempty"`
	// The task runs a SQL query or file, or it refreshes a SQL alert or a
	// legacy SQL dashboard when the `sql_task` field is present.
	SqlTask *sqlTaskPb `json:"sql_task,omitempty"`
	// A unique name for the task. This field is used to refer to this task from
	// other tasks. This field is required and must be unique within its parent
	// job. On Update or Reset, this field is used to reference the tasks to be
	// updated or reset.
	TaskKey string `json:"task_key"`
	// An optional timeout applied to each run of this job task. A value of `0`
	// means no timeout.
	TimeoutSeconds int `json:"timeout_seconds,omitempty"`
	// A collection of system notification IDs to notify when runs of this task
	// begin or complete. The default behavior is to not send any system
	// notifications.
	WebhookNotifications *webhookNotificationsPb `json:"webhook_notifications,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskFromPb(pb *taskPb) (*Task, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Task{}
	cleanRoomsNotebookTaskField, err := cleanRoomsNotebookTaskFromPb(pb.CleanRoomsNotebookTask)
	if err != nil {
		return nil, err
	}
	if cleanRoomsNotebookTaskField != nil {
		st.CleanRoomsNotebookTask = cleanRoomsNotebookTaskField
	}
	conditionTaskField, err := conditionTaskFromPb(pb.ConditionTask)
	if err != nil {
		return nil, err
	}
	if conditionTaskField != nil {
		st.ConditionTask = conditionTaskField
	}
	dashboardTaskField, err := dashboardTaskFromPb(pb.DashboardTask)
	if err != nil {
		return nil, err
	}
	if dashboardTaskField != nil {
		st.DashboardTask = dashboardTaskField
	}
	dbtTaskField, err := dbtTaskFromPb(pb.DbtTask)
	if err != nil {
		return nil, err
	}
	if dbtTaskField != nil {
		st.DbtTask = dbtTaskField
	}

	var dependsOnField []TaskDependency
	for _, itemPb := range pb.DependsOn {
		item, err := taskDependencyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dependsOnField = append(dependsOnField, *item)
		}
	}
	st.DependsOn = dependsOnField
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	disableAutoOptimizationField, err := identity(&pb.DisableAutoOptimization)
	if err != nil {
		return nil, err
	}
	if disableAutoOptimizationField != nil {
		st.DisableAutoOptimization = *disableAutoOptimizationField
	}
	emailNotificationsField, err := taskEmailNotificationsFromPb(pb.EmailNotifications)
	if err != nil {
		return nil, err
	}
	if emailNotificationsField != nil {
		st.EmailNotifications = emailNotificationsField
	}
	environmentKeyField, err := identity(&pb.EnvironmentKey)
	if err != nil {
		return nil, err
	}
	if environmentKeyField != nil {
		st.EnvironmentKey = *environmentKeyField
	}
	existingClusterIdField, err := identity(&pb.ExistingClusterId)
	if err != nil {
		return nil, err
	}
	if existingClusterIdField != nil {
		st.ExistingClusterId = *existingClusterIdField
	}
	forEachTaskField, err := forEachTaskFromPb(pb.ForEachTask)
	if err != nil {
		return nil, err
	}
	if forEachTaskField != nil {
		st.ForEachTask = forEachTaskField
	}
	genAiComputeTaskField, err := genAiComputeTaskFromPb(pb.GenAiComputeTask)
	if err != nil {
		return nil, err
	}
	if genAiComputeTaskField != nil {
		st.GenAiComputeTask = genAiComputeTaskField
	}
	healthField, err := jobsHealthRulesFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	jobClusterKeyField, err := identity(&pb.JobClusterKey)
	if err != nil {
		return nil, err
	}
	if jobClusterKeyField != nil {
		st.JobClusterKey = *jobClusterKeyField
	}

	var librariesField []compute.Library
	for _, itemPb := range pb.Libraries {
		item, err := compute.LibraryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			librariesField = append(librariesField, *item)
		}
	}
	st.Libraries = librariesField
	maxRetriesField, err := identity(&pb.MaxRetries)
	if err != nil {
		return nil, err
	}
	if maxRetriesField != nil {
		st.MaxRetries = *maxRetriesField
	}
	minRetryIntervalMillisField, err := identity(&pb.MinRetryIntervalMillis)
	if err != nil {
		return nil, err
	}
	if minRetryIntervalMillisField != nil {
		st.MinRetryIntervalMillis = *minRetryIntervalMillisField
	}
	newClusterField, err := compute.ClusterSpecFromPb(pb.NewCluster)
	if err != nil {
		return nil, err
	}
	if newClusterField != nil {
		st.NewCluster = newClusterField
	}
	notebookTaskField, err := notebookTaskFromPb(pb.NotebookTask)
	if err != nil {
		return nil, err
	}
	if notebookTaskField != nil {
		st.NotebookTask = notebookTaskField
	}
	notificationSettingsField, err := taskNotificationSettingsFromPb(pb.NotificationSettings)
	if err != nil {
		return nil, err
	}
	if notificationSettingsField != nil {
		st.NotificationSettings = notificationSettingsField
	}
	pipelineTaskField, err := pipelineTaskFromPb(pb.PipelineTask)
	if err != nil {
		return nil, err
	}
	if pipelineTaskField != nil {
		st.PipelineTask = pipelineTaskField
	}
	powerBiTaskField, err := powerBiTaskFromPb(pb.PowerBiTask)
	if err != nil {
		return nil, err
	}
	if powerBiTaskField != nil {
		st.PowerBiTask = powerBiTaskField
	}
	pythonWheelTaskField, err := pythonWheelTaskFromPb(pb.PythonWheelTask)
	if err != nil {
		return nil, err
	}
	if pythonWheelTaskField != nil {
		st.PythonWheelTask = pythonWheelTaskField
	}
	retryOnTimeoutField, err := identity(&pb.RetryOnTimeout)
	if err != nil {
		return nil, err
	}
	if retryOnTimeoutField != nil {
		st.RetryOnTimeout = *retryOnTimeoutField
	}
	runIfField, err := identity(&pb.RunIf)
	if err != nil {
		return nil, err
	}
	if runIfField != nil {
		st.RunIf = *runIfField
	}
	runJobTaskField, err := runJobTaskFromPb(pb.RunJobTask)
	if err != nil {
		return nil, err
	}
	if runJobTaskField != nil {
		st.RunJobTask = runJobTaskField
	}
	sparkJarTaskField, err := sparkJarTaskFromPb(pb.SparkJarTask)
	if err != nil {
		return nil, err
	}
	if sparkJarTaskField != nil {
		st.SparkJarTask = sparkJarTaskField
	}
	sparkPythonTaskField, err := sparkPythonTaskFromPb(pb.SparkPythonTask)
	if err != nil {
		return nil, err
	}
	if sparkPythonTaskField != nil {
		st.SparkPythonTask = sparkPythonTaskField
	}
	sparkSubmitTaskField, err := sparkSubmitTaskFromPb(pb.SparkSubmitTask)
	if err != nil {
		return nil, err
	}
	if sparkSubmitTaskField != nil {
		st.SparkSubmitTask = sparkSubmitTaskField
	}
	sqlTaskField, err := sqlTaskFromPb(pb.SqlTask)
	if err != nil {
		return nil, err
	}
	if sqlTaskField != nil {
		st.SqlTask = sqlTaskField
	}
	taskKeyField, err := identity(&pb.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyField != nil {
		st.TaskKey = *taskKeyField
	}
	timeoutSecondsField, err := identity(&pb.TimeoutSeconds)
	if err != nil {
		return nil, err
	}
	if timeoutSecondsField != nil {
		st.TimeoutSeconds = *timeoutSecondsField
	}
	webhookNotificationsField, err := webhookNotificationsFromPb(pb.WebhookNotifications)
	if err != nil {
		return nil, err
	}
	if webhookNotificationsField != nil {
		st.WebhookNotifications = webhookNotificationsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskDependency struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	Outcome string
	// The name of the task this task depends on.
	TaskKey string

	ForceSendFields []string
}

func taskDependencyToPb(st *TaskDependency) (*taskDependencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskDependencyPb{}
	outcomePb, err := identity(&st.Outcome)
	if err != nil {
		return nil, err
	}
	if outcomePb != nil {
		pb.Outcome = *outcomePb
	}

	taskKeyPb, err := identity(&st.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyPb != nil {
		pb.TaskKey = *taskKeyPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TaskDependency) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskDependencyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskDependencyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskDependency) MarshalJSON() ([]byte, error) {
	pb, err := taskDependencyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type taskDependencyPb struct {
	// Can only be specified on condition task dependencies. The outcome of the
	// dependent task that must be met for this task to run.
	Outcome string `json:"outcome,omitempty"`
	// The name of the task this task depends on.
	TaskKey string `json:"task_key"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskDependencyFromPb(pb *taskDependencyPb) (*TaskDependency, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskDependency{}
	outcomeField, err := identity(&pb.Outcome)
	if err != nil {
		return nil, err
	}
	if outcomeField != nil {
		st.Outcome = *outcomeField
	}
	taskKeyField, err := identity(&pb.TaskKey)
	if err != nil {
		return nil, err
	}
	if taskKeyField != nil {
		st.TaskKey = *taskKeyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskDependencyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskDependencyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskEmailNotifications struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns bool
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []string
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string

	ForceSendFields []string
}

func taskEmailNotificationsToPb(st *TaskEmailNotifications) (*taskEmailNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskEmailNotificationsPb{}
	noAlertForSkippedRunsPb, err := identity(&st.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsPb != nil {
		pb.NoAlertForSkippedRuns = *noAlertForSkippedRunsPb
	}

	var onDurationWarningThresholdExceededPb []string
	for _, item := range st.OnDurationWarningThresholdExceeded {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onDurationWarningThresholdExceededPb = append(onDurationWarningThresholdExceededPb, *itemPb)
		}
	}
	pb.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededPb

	var onFailurePb []string
	for _, item := range st.OnFailure {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onFailurePb = append(onFailurePb, *itemPb)
		}
	}
	pb.OnFailure = onFailurePb

	var onStartPb []string
	for _, item := range st.OnStart {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStartPb = append(onStartPb, *itemPb)
		}
	}
	pb.OnStart = onStartPb

	var onStreamingBacklogExceededPb []string
	for _, item := range st.OnStreamingBacklogExceeded {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStreamingBacklogExceededPb = append(onStreamingBacklogExceededPb, *itemPb)
		}
	}
	pb.OnStreamingBacklogExceeded = onStreamingBacklogExceededPb

	var onSuccessPb []string
	for _, item := range st.OnSuccess {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onSuccessPb = append(onSuccessPb, *itemPb)
		}
	}
	pb.OnSuccess = onSuccessPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TaskEmailNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskEmailNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskEmailNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskEmailNotifications) MarshalJSON() ([]byte, error) {
	pb, err := taskEmailNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type taskEmailNotificationsPb struct {
	// If true, do not send email to recipients specified in `on_failure` if the
	// run is skipped. This field is `deprecated`. Please use the
	// `notification_settings.no_alert_for_skipped_runs` field.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`
	// A list of email addresses to be notified when the duration of a run
	// exceeds the threshold specified for the `RUN_DURATION_SECONDS` metric in
	// the `health` field. If no rule for the `RUN_DURATION_SECONDS` metric is
	// specified in the `health` field for the job, notifications are not sent.
	OnDurationWarningThresholdExceeded []string `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// A list of email addresses to be notified when a run unsuccessfully
	// completes. A run is considered to have completed unsuccessfully if it
	// ends with an `INTERNAL_ERROR` `life_cycle_state` or a `FAILED`, or
	// `TIMED_OUT` result_state. If this is not specified on job creation,
	// reset, or update the list is empty, and notifications are not sent.
	OnFailure []string `json:"on_failure,omitempty"`
	// A list of email addresses to be notified when a run begins. If not
	// specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnStart []string `json:"on_start,omitempty"`
	// A list of email addresses to notify when any streaming backlog thresholds
	// are exceeded for any stream. Streaming backlog thresholds can be set in
	// the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes.
	OnStreamingBacklogExceeded []string `json:"on_streaming_backlog_exceeded,omitempty"`
	// A list of email addresses to be notified when a run successfully
	// completes. A run is considered to have completed successfully if it ends
	// with a `TERMINATED` `life_cycle_state` and a `SUCCESS` result_state. If
	// not specified on job creation, reset, or update, the list is empty, and
	// notifications are not sent.
	OnSuccess []string `json:"on_success,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskEmailNotificationsFromPb(pb *taskEmailNotificationsPb) (*TaskEmailNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskEmailNotifications{}
	noAlertForSkippedRunsField, err := identity(&pb.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsField != nil {
		st.NoAlertForSkippedRuns = *noAlertForSkippedRunsField
	}

	var onDurationWarningThresholdExceededField []string
	for _, itemPb := range pb.OnDurationWarningThresholdExceeded {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onDurationWarningThresholdExceededField = append(onDurationWarningThresholdExceededField, *item)
		}
	}
	st.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededField

	var onFailureField []string
	for _, itemPb := range pb.OnFailure {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onFailureField = append(onFailureField, *item)
		}
	}
	st.OnFailure = onFailureField

	var onStartField []string
	for _, itemPb := range pb.OnStart {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStartField = append(onStartField, *item)
		}
	}
	st.OnStart = onStartField

	var onStreamingBacklogExceededField []string
	for _, itemPb := range pb.OnStreamingBacklogExceeded {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStreamingBacklogExceededField = append(onStreamingBacklogExceededField, *item)
		}
	}
	st.OnStreamingBacklogExceeded = onStreamingBacklogExceededField

	var onSuccessField []string
	for _, itemPb := range pb.OnSuccess {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onSuccessField = append(onSuccessField, *item)
		}
	}
	st.OnSuccess = onSuccessField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskEmailNotificationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskEmailNotificationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskNotificationSettings struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	AlertOnLastAttempt bool
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool

	ForceSendFields []string
}

func taskNotificationSettingsToPb(st *TaskNotificationSettings) (*taskNotificationSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskNotificationSettingsPb{}
	alertOnLastAttemptPb, err := identity(&st.AlertOnLastAttempt)
	if err != nil {
		return nil, err
	}
	if alertOnLastAttemptPb != nil {
		pb.AlertOnLastAttempt = *alertOnLastAttemptPb
	}

	noAlertForCanceledRunsPb, err := identity(&st.NoAlertForCanceledRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForCanceledRunsPb != nil {
		pb.NoAlertForCanceledRuns = *noAlertForCanceledRunsPb
	}

	noAlertForSkippedRunsPb, err := identity(&st.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsPb != nil {
		pb.NoAlertForSkippedRuns = *noAlertForSkippedRunsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TaskNotificationSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskNotificationSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskNotificationSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskNotificationSettings) MarshalJSON() ([]byte, error) {
	pb, err := taskNotificationSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type taskNotificationSettingsPb struct {
	// If true, do not send notifications to recipients specified in `on_start`
	// for the retried runs and do not send notifications to recipients
	// specified in `on_failure` until the last retry of the run.
	AlertOnLastAttempt bool `json:"alert_on_last_attempt,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is canceled.
	NoAlertForCanceledRuns bool `json:"no_alert_for_canceled_runs,omitempty"`
	// If true, do not send notifications to recipients specified in
	// `on_failure` if the run is skipped.
	NoAlertForSkippedRuns bool `json:"no_alert_for_skipped_runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskNotificationSettingsFromPb(pb *taskNotificationSettingsPb) (*TaskNotificationSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskNotificationSettings{}
	alertOnLastAttemptField, err := identity(&pb.AlertOnLastAttempt)
	if err != nil {
		return nil, err
	}
	if alertOnLastAttemptField != nil {
		st.AlertOnLastAttempt = *alertOnLastAttemptField
	}
	noAlertForCanceledRunsField, err := identity(&pb.NoAlertForCanceledRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForCanceledRunsField != nil {
		st.NoAlertForCanceledRuns = *noAlertForCanceledRunsField
	}
	noAlertForSkippedRunsField, err := identity(&pb.NoAlertForSkippedRuns)
	if err != nil {
		return nil, err
	}
	if noAlertForSkippedRunsField != nil {
		st.NoAlertForSkippedRuns = *noAlertForSkippedRunsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskNotificationSettingsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskNotificationSettingsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The code indicates why the run was terminated. Additional codes might be
// introduced in future releases. * `SUCCESS`: The run was completed
// successfully. * `USER_CANCELED`: The run was successfully canceled during
// execution by a user. * `CANCELED`: The run was canceled during execution by
// the Databricks platform; for example, if the maximum run duration was
// exceeded. * `SKIPPED`: Run was never executed, for example, if the upstream
// task run failed, the dependency type condition was not met, or there were no
// material tasks to execute. * `INTERNAL_ERROR`: The run encountered an
// unexpected error. Refer to the state message for further details. *
// `DRIVER_ERROR`: The run encountered an error while communicating with the
// Spark Driver. * `CLUSTER_ERROR`: The run failed due to a cluster error. Refer
// to the state message for further details. * `REPOSITORY_CHECKOUT_FAILED`:
// Failed to complete the checkout due to an error when communicating with the
// third party service. * `INVALID_CLUSTER_REQUEST`: The run failed because it
// issued an invalid request to start the cluster. *
// `WORKSPACE_RUN_LIMIT_EXCEEDED`: The workspace has reached the quota for the
// maximum number of concurrent active runs. Consider scheduling the runs over a
// larger time frame. * `FEATURE_DISABLED`: The run failed because it tried to
// access a feature unavailable for the workspace. *
// `CLUSTER_REQUEST_LIMIT_EXCEEDED`: The number of cluster creation, start, and
// upsize requests have exceeded the allotted rate limit. Consider spreading the
// run execution over a larger time frame. * `STORAGE_ACCESS_ERROR`: The run
// failed due to an error when accessing the customer blob storage. Refer to the
// state message for further details. * `RUN_EXECUTION_ERROR`: The run was
// completed with task failures. For more details, refer to the state message or
// run output. * `UNAUTHORIZED_ERROR`: The run failed due to a permission issue
// while accessing a resource. Refer to the state message for further details. *
// `LIBRARY_INSTALLATION_ERROR`: The run failed while installing the
// user-requested library. Refer to the state message for further details. The
// causes might include, but are not limited to: The provided library is
// invalid, there are insufficient permissions to install the library, and so
// forth. * `MAX_CONCURRENT_RUNS_EXCEEDED`: The scheduled run exceeds the limit
// of maximum concurrent runs set for the job. * `MAX_SPARK_CONTEXTS_EXCEEDED`:
// The run is scheduled on a cluster that has already reached the maximum number
// of contexts it is configured to create. See: [Link]. * `RESOURCE_NOT_FOUND`:
// A resource necessary for run execution does not exist. Refer to the state
// message for further details. * `INVALID_RUN_CONFIGURATION`: The run failed
// due to an invalid configuration. Refer to the state message for further
// details. * `CLOUD_FAILURE`: The run failed due to a cloud provider issue.
// Refer to the state message for further details. *
// `MAX_JOB_QUEUE_SIZE_EXCEEDED`: The run was skipped due to reaching the job
// level queue size limit. * `DISABLED`: The run was never executed because it
// was disabled explicitly by the user.
//
// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
type TerminationCodeCode string
type terminationCodeCodePb string

const TerminationCodeCodeBudgetPolicyLimitExceeded TerminationCodeCode = `BUDGET_POLICY_LIMIT_EXCEEDED`

// The run was canceled during execution by the <Databricks> platform; for
// example, if the maximum run duration was exceeded.
const TerminationCodeCodeCanceled TerminationCodeCode = `CANCELED`

// The run failed due to a cloud provider issue. Refer to the state message for
// further details.
const TerminationCodeCodeCloudFailure TerminationCodeCode = `CLOUD_FAILURE`

// The run failed due to a cluster error. Refer to the state message for further
// details.
const TerminationCodeCodeClusterError TerminationCodeCode = `CLUSTER_ERROR`

// The number of cluster creation, start, and upsize requests have exceeded the
// allotted rate limit. Consider spreading the run execution over a larger time
// frame.
const TerminationCodeCodeClusterRequestLimitExceeded TerminationCodeCode = `CLUSTER_REQUEST_LIMIT_EXCEEDED`

// The run was never executed because it was disabled explicitly by the user.
const TerminationCodeCodeDisabled TerminationCodeCode = `DISABLED`

// The run encountered an error while communicating with the Spark Driver.
const TerminationCodeCodeDriverError TerminationCodeCode = `DRIVER_ERROR`

// The run failed because it tried to access a feature unavailable for the
// workspace.
const TerminationCodeCodeFeatureDisabled TerminationCodeCode = `FEATURE_DISABLED`

// The run encountered an unexpected error. Refer to the state message for
// further details.
const TerminationCodeCodeInternalError TerminationCodeCode = `INTERNAL_ERROR`

// The run failed because it issued an invalid request to start the cluster.
const TerminationCodeCodeInvalidClusterRequest TerminationCodeCode = `INVALID_CLUSTER_REQUEST`

// The run failed due to an invalid configuration. Refer to the state message
// for further details.
const TerminationCodeCodeInvalidRunConfiguration TerminationCodeCode = `INVALID_RUN_CONFIGURATION`

// The run failed while installing the user-requested library. Refer to the
// state message for further details. The causes might include, but are not
// limited to: The provided library is invalid, there are insufficient
// permissions to install the library, and so forth.
const TerminationCodeCodeLibraryInstallationError TerminationCodeCode = `LIBRARY_INSTALLATION_ERROR`

// The scheduled run exceeds the limit of maximum concurrent runs set for the
// job.
const TerminationCodeCodeMaxConcurrentRunsExceeded TerminationCodeCode = `MAX_CONCURRENT_RUNS_EXCEEDED`

// The run was skipped due to reaching the job level queue size limit.
const TerminationCodeCodeMaxJobQueueSizeExceeded TerminationCodeCode = `MAX_JOB_QUEUE_SIZE_EXCEEDED`

// The run is scheduled on a cluster that has already reached the maximum number
// of contexts it is configured to create. See: [Link].
//
// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
const TerminationCodeCodeMaxSparkContextsExceeded TerminationCodeCode = `MAX_SPARK_CONTEXTS_EXCEEDED`

// Failed to complete the checkout due to an error when communicating with the
// third party service.
const TerminationCodeCodeRepositoryCheckoutFailed TerminationCodeCode = `REPOSITORY_CHECKOUT_FAILED`

// A resource necessary for run execution does not exist. Refer to the state
// message for further details.
const TerminationCodeCodeResourceNotFound TerminationCodeCode = `RESOURCE_NOT_FOUND`

// The run was completed with task failures. For more details, refer to the
// state message or run output.
const TerminationCodeCodeRunExecutionError TerminationCodeCode = `RUN_EXECUTION_ERROR`

// Run was never executed, for example, if the upstream task run failed, the
// dependency type condition was not met, or there were no material tasks to
// execute.
const TerminationCodeCodeSkipped TerminationCodeCode = `SKIPPED`

// The run failed due to an error when accessing the customer blob storage.
// Refer to the state message for further details.
const TerminationCodeCodeStorageAccessError TerminationCodeCode = `STORAGE_ACCESS_ERROR`

// The run was completed successfully.
const TerminationCodeCodeSuccess TerminationCodeCode = `SUCCESS`

// The run failed due to a permission issue while accessing a resource. Refer to
// the state message for further details.
const TerminationCodeCodeUnauthorizedError TerminationCodeCode = `UNAUTHORIZED_ERROR`

// The run was successfully canceled during execution by a user.
const TerminationCodeCodeUserCanceled TerminationCodeCode = `USER_CANCELED`

// The workspace has reached the quota for the maximum number of concurrent
// active runs. Consider scheduling the runs over a larger time frame.
const TerminationCodeCodeWorkspaceRunLimitExceeded TerminationCodeCode = `WORKSPACE_RUN_LIMIT_EXCEEDED`

// String representation for [fmt.Print]
func (f *TerminationCodeCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationCodeCode) Set(v string) error {
	switch v {
	case `BUDGET_POLICY_LIMIT_EXCEEDED`, `CANCELED`, `CLOUD_FAILURE`, `CLUSTER_ERROR`, `CLUSTER_REQUEST_LIMIT_EXCEEDED`, `DISABLED`, `DRIVER_ERROR`, `FEATURE_DISABLED`, `INTERNAL_ERROR`, `INVALID_CLUSTER_REQUEST`, `INVALID_RUN_CONFIGURATION`, `LIBRARY_INSTALLATION_ERROR`, `MAX_CONCURRENT_RUNS_EXCEEDED`, `MAX_JOB_QUEUE_SIZE_EXCEEDED`, `MAX_SPARK_CONTEXTS_EXCEEDED`, `REPOSITORY_CHECKOUT_FAILED`, `RESOURCE_NOT_FOUND`, `RUN_EXECUTION_ERROR`, `SKIPPED`, `STORAGE_ACCESS_ERROR`, `SUCCESS`, `UNAUTHORIZED_ERROR`, `USER_CANCELED`, `WORKSPACE_RUN_LIMIT_EXCEEDED`:
		*f = TerminationCodeCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BUDGET_POLICY_LIMIT_EXCEEDED", "CANCELED", "CLOUD_FAILURE", "CLUSTER_ERROR", "CLUSTER_REQUEST_LIMIT_EXCEEDED", "DISABLED", "DRIVER_ERROR", "FEATURE_DISABLED", "INTERNAL_ERROR", "INVALID_CLUSTER_REQUEST", "INVALID_RUN_CONFIGURATION", "LIBRARY_INSTALLATION_ERROR", "MAX_CONCURRENT_RUNS_EXCEEDED", "MAX_JOB_QUEUE_SIZE_EXCEEDED", "MAX_SPARK_CONTEXTS_EXCEEDED", "REPOSITORY_CHECKOUT_FAILED", "RESOURCE_NOT_FOUND", "RUN_EXECUTION_ERROR", "SKIPPED", "STORAGE_ACCESS_ERROR", "SUCCESS", "UNAUTHORIZED_ERROR", "USER_CANCELED", "WORKSPACE_RUN_LIMIT_EXCEEDED"`, v)
	}
}

// Type always returns TerminationCodeCode to satisfy [pflag.Value] interface
func (f *TerminationCodeCode) Type() string {
	return "TerminationCodeCode"
}

func terminationCodeCodeToPb(st *TerminationCodeCode) (*terminationCodeCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := terminationCodeCodePb(*st)
	return &pb, nil
}

func terminationCodeCodeFromPb(pb *terminationCodeCodePb) (*TerminationCodeCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationCodeCode(*pb)
	return &st, nil
}

type TerminationDetails struct {
	// The code indicates why the run was terminated. Additional codes might be
	// introduced in future releases. * `SUCCESS`: The run was completed
	// successfully. * `USER_CANCELED`: The run was successfully canceled during
	// execution by a user. * `CANCELED`: The run was canceled during execution
	// by the Databricks platform; for example, if the maximum run duration was
	// exceeded. * `SKIPPED`: Run was never executed, for example, if the
	// upstream task run failed, the dependency type condition was not met, or
	// there were no material tasks to execute. * `INTERNAL_ERROR`: The run
	// encountered an unexpected error. Refer to the state message for further
	// details. * `DRIVER_ERROR`: The run encountered an error while
	// communicating with the Spark Driver. * `CLUSTER_ERROR`: The run failed
	// due to a cluster error. Refer to the state message for further details. *
	// `REPOSITORY_CHECKOUT_FAILED`: Failed to complete the checkout due to an
	// error when communicating with the third party service. *
	// `INVALID_CLUSTER_REQUEST`: The run failed because it issued an invalid
	// request to start the cluster. * `WORKSPACE_RUN_LIMIT_EXCEEDED`: The
	// workspace has reached the quota for the maximum number of concurrent
	// active runs. Consider scheduling the runs over a larger time frame. *
	// `FEATURE_DISABLED`: The run failed because it tried to access a feature
	// unavailable for the workspace. * `CLUSTER_REQUEST_LIMIT_EXCEEDED`: The
	// number of cluster creation, start, and upsize requests have exceeded the
	// allotted rate limit. Consider spreading the run execution over a larger
	// time frame. * `STORAGE_ACCESS_ERROR`: The run failed due to an error when
	// accessing the customer blob storage. Refer to the state message for
	// further details. * `RUN_EXECUTION_ERROR`: The run was completed with task
	// failures. For more details, refer to the state message or run output. *
	// `UNAUTHORIZED_ERROR`: The run failed due to a permission issue while
	// accessing a resource. Refer to the state message for further details. *
	// `LIBRARY_INSTALLATION_ERROR`: The run failed while installing the
	// user-requested library. Refer to the state message for further details.
	// The causes might include, but are not limited to: The provided library is
	// invalid, there are insufficient permissions to install the library, and
	// so forth. * `MAX_CONCURRENT_RUNS_EXCEEDED`: The scheduled run exceeds the
	// limit of maximum concurrent runs set for the job. *
	// `MAX_SPARK_CONTEXTS_EXCEEDED`: The run is scheduled on a cluster that has
	// already reached the maximum number of contexts it is configured to
	// create. See: [Link]. * `RESOURCE_NOT_FOUND`: A resource necessary for run
	// execution does not exist. Refer to the state message for further details.
	// * `INVALID_RUN_CONFIGURATION`: The run failed due to an invalid
	// configuration. Refer to the state message for further details. *
	// `CLOUD_FAILURE`: The run failed due to a cloud provider issue. Refer to
	// the state message for further details. * `MAX_JOB_QUEUE_SIZE_EXCEEDED`:
	// The run was skipped due to reaching the job level queue size limit. *
	// `DISABLED`: The run was never executed because it was disabled explicitly
	// by the user.
	//
	// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
	Code TerminationCodeCode
	// A descriptive message with the termination details. This field is
	// unstructured and the format might change.
	Message string
	// * `SUCCESS`: The run terminated without any issues * `INTERNAL_ERROR`: An
	// error occurred in the Databricks platform. Please look at the [status
	// page] or contact support if the issue persists. * `CLIENT_ERROR`: The run
	// was terminated because of an error caused by user input or the job
	// configuration. * `CLOUD_FAILURE`: The run was terminated because of an
	// issue with your cloud provider.
	//
	// [status page]: https://status.databricks.com/
	Type TerminationTypeType

	ForceSendFields []string
}

func terminationDetailsToPb(st *TerminationDetails) (*terminationDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &terminationDetailsPb{}
	codePb, err := identity(&st.Code)
	if err != nil {
		return nil, err
	}
	if codePb != nil {
		pb.Code = *codePb
	}

	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
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

func (st *TerminationDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &terminationDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := terminationDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TerminationDetails) MarshalJSON() ([]byte, error) {
	pb, err := terminationDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type terminationDetailsPb struct {
	// The code indicates why the run was terminated. Additional codes might be
	// introduced in future releases. * `SUCCESS`: The run was completed
	// successfully. * `USER_CANCELED`: The run was successfully canceled during
	// execution by a user. * `CANCELED`: The run was canceled during execution
	// by the Databricks platform; for example, if the maximum run duration was
	// exceeded. * `SKIPPED`: Run was never executed, for example, if the
	// upstream task run failed, the dependency type condition was not met, or
	// there were no material tasks to execute. * `INTERNAL_ERROR`: The run
	// encountered an unexpected error. Refer to the state message for further
	// details. * `DRIVER_ERROR`: The run encountered an error while
	// communicating with the Spark Driver. * `CLUSTER_ERROR`: The run failed
	// due to a cluster error. Refer to the state message for further details. *
	// `REPOSITORY_CHECKOUT_FAILED`: Failed to complete the checkout due to an
	// error when communicating with the third party service. *
	// `INVALID_CLUSTER_REQUEST`: The run failed because it issued an invalid
	// request to start the cluster. * `WORKSPACE_RUN_LIMIT_EXCEEDED`: The
	// workspace has reached the quota for the maximum number of concurrent
	// active runs. Consider scheduling the runs over a larger time frame. *
	// `FEATURE_DISABLED`: The run failed because it tried to access a feature
	// unavailable for the workspace. * `CLUSTER_REQUEST_LIMIT_EXCEEDED`: The
	// number of cluster creation, start, and upsize requests have exceeded the
	// allotted rate limit. Consider spreading the run execution over a larger
	// time frame. * `STORAGE_ACCESS_ERROR`: The run failed due to an error when
	// accessing the customer blob storage. Refer to the state message for
	// further details. * `RUN_EXECUTION_ERROR`: The run was completed with task
	// failures. For more details, refer to the state message or run output. *
	// `UNAUTHORIZED_ERROR`: The run failed due to a permission issue while
	// accessing a resource. Refer to the state message for further details. *
	// `LIBRARY_INSTALLATION_ERROR`: The run failed while installing the
	// user-requested library. Refer to the state message for further details.
	// The causes might include, but are not limited to: The provided library is
	// invalid, there are insufficient permissions to install the library, and
	// so forth. * `MAX_CONCURRENT_RUNS_EXCEEDED`: The scheduled run exceeds the
	// limit of maximum concurrent runs set for the job. *
	// `MAX_SPARK_CONTEXTS_EXCEEDED`: The run is scheduled on a cluster that has
	// already reached the maximum number of contexts it is configured to
	// create. See: [Link]. * `RESOURCE_NOT_FOUND`: A resource necessary for run
	// execution does not exist. Refer to the state message for further details.
	// * `INVALID_RUN_CONFIGURATION`: The run failed due to an invalid
	// configuration. Refer to the state message for further details. *
	// `CLOUD_FAILURE`: The run failed due to a cloud provider issue. Refer to
	// the state message for further details. * `MAX_JOB_QUEUE_SIZE_EXCEEDED`:
	// The run was skipped due to reaching the job level queue size limit. *
	// `DISABLED`: The run was never executed because it was disabled explicitly
	// by the user.
	//
	// [Link]: https://kb.databricks.com/en_US/notebooks/too-many-execution-contexts-are-open-right-now
	Code TerminationCodeCode `json:"code,omitempty"`
	// A descriptive message with the termination details. This field is
	// unstructured and the format might change.
	Message string `json:"message,omitempty"`
	// * `SUCCESS`: The run terminated without any issues * `INTERNAL_ERROR`: An
	// error occurred in the Databricks platform. Please look at the [status
	// page] or contact support if the issue persists. * `CLIENT_ERROR`: The run
	// was terminated because of an error caused by user input or the job
	// configuration. * `CLOUD_FAILURE`: The run was terminated because of an
	// issue with your cloud provider.
	//
	// [status page]: https://status.databricks.com/
	Type TerminationTypeType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func terminationDetailsFromPb(pb *terminationDetailsPb) (*TerminationDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TerminationDetails{}
	codeField, err := identity(&pb.Code)
	if err != nil {
		return nil, err
	}
	if codeField != nil {
		st.Code = *codeField
	}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
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

func (st *terminationDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st terminationDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// * `SUCCESS`: The run terminated without any issues * `INTERNAL_ERROR`: An
// error occurred in the Databricks platform. Please look at the [status page]
// or contact support if the issue persists. * `CLIENT_ERROR`: The run was
// terminated because of an error caused by user input or the job configuration.
// * `CLOUD_FAILURE`: The run was terminated because of an issue with your cloud
// provider.
//
// [status page]: https://status.databricks.com/
type TerminationTypeType string
type terminationTypeTypePb string

// The run was terminated because of an error caused by user input or the job
// configuration.
const TerminationTypeTypeClientError TerminationTypeType = `CLIENT_ERROR`

// The run was terminated because of an issue with your cloud provider.
const TerminationTypeTypeCloudFailure TerminationTypeType = `CLOUD_FAILURE`

// An error occurred in the <Databricks> platform. Please look at the [status
// page] or contact support if the issue persists.
//
// [status page]: https://status.databricks.com/
const TerminationTypeTypeInternalError TerminationTypeType = `INTERNAL_ERROR`

// The run terminated without any issues
const TerminationTypeTypeSuccess TerminationTypeType = `SUCCESS`

// String representation for [fmt.Print]
func (f *TerminationTypeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationTypeType) Set(v string) error {
	switch v {
	case `CLIENT_ERROR`, `CLOUD_FAILURE`, `INTERNAL_ERROR`, `SUCCESS`:
		*f = TerminationTypeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLIENT_ERROR", "CLOUD_FAILURE", "INTERNAL_ERROR", "SUCCESS"`, v)
	}
}

// Type always returns TerminationTypeType to satisfy [pflag.Value] interface
func (f *TerminationTypeType) Type() string {
	return "TerminationTypeType"
}

func terminationTypeTypeToPb(st *TerminationTypeType) (*terminationTypeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := terminationTypeTypePb(*st)
	return &pb, nil
}

func terminationTypeTypeFromPb(pb *terminationTypeTypePb) (*TerminationTypeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationTypeType(*pb)
	return &st, nil
}

// Additional details about what triggered the run
type TriggerInfo struct {
	// The run id of the Run Job task run
	RunId int64

	ForceSendFields []string
}

func triggerInfoToPb(st *TriggerInfo) (*triggerInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &triggerInfoPb{}
	runIdPb, err := identity(&st.RunId)
	if err != nil {
		return nil, err
	}
	if runIdPb != nil {
		pb.RunId = *runIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TriggerInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &triggerInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := triggerInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TriggerInfo) MarshalJSON() ([]byte, error) {
	pb, err := triggerInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type triggerInfoPb struct {
	// The run id of the Run Job task run
	RunId int64 `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func triggerInfoFromPb(pb *triggerInfoPb) (*TriggerInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggerInfo{}
	runIdField, err := identity(&pb.RunId)
	if err != nil {
		return nil, err
	}
	if runIdField != nil {
		st.RunId = *runIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *triggerInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st triggerInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TriggerSettings struct {
	// File arrival trigger settings.
	FileArrival *FileArrivalTriggerConfiguration
	// Whether this trigger is paused or not.
	PauseStatus PauseStatus
	// Periodic trigger settings.
	Periodic *PeriodicTriggerConfiguration
	// Old table trigger settings name. Deprecated in favor of `table_update`.
	Table *TableUpdateTriggerConfiguration

	TableUpdate *TableUpdateTriggerConfiguration
}

func triggerSettingsToPb(st *TriggerSettings) (*triggerSettingsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &triggerSettingsPb{}
	fileArrivalPb, err := fileArrivalTriggerConfigurationToPb(st.FileArrival)
	if err != nil {
		return nil, err
	}
	if fileArrivalPb != nil {
		pb.FileArrival = fileArrivalPb
	}

	pauseStatusPb, err := identity(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}

	periodicPb, err := periodicTriggerConfigurationToPb(st.Periodic)
	if err != nil {
		return nil, err
	}
	if periodicPb != nil {
		pb.Periodic = periodicPb
	}

	tablePb, err := tableUpdateTriggerConfigurationToPb(st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = tablePb
	}

	tableUpdatePb, err := tableUpdateTriggerConfigurationToPb(st.TableUpdate)
	if err != nil {
		return nil, err
	}
	if tableUpdatePb != nil {
		pb.TableUpdate = tableUpdatePb
	}

	return pb, nil
}

func (st *TriggerSettings) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &triggerSettingsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := triggerSettingsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TriggerSettings) MarshalJSON() ([]byte, error) {
	pb, err := triggerSettingsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type triggerSettingsPb struct {
	// File arrival trigger settings.
	FileArrival *fileArrivalTriggerConfigurationPb `json:"file_arrival,omitempty"`
	// Whether this trigger is paused or not.
	PauseStatus PauseStatus `json:"pause_status,omitempty"`
	// Periodic trigger settings.
	Periodic *periodicTriggerConfigurationPb `json:"periodic,omitempty"`
	// Old table trigger settings name. Deprecated in favor of `table_update`.
	Table *tableUpdateTriggerConfigurationPb `json:"table,omitempty"`

	TableUpdate *tableUpdateTriggerConfigurationPb `json:"table_update,omitempty"`
}

func triggerSettingsFromPb(pb *triggerSettingsPb) (*TriggerSettings, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TriggerSettings{}
	fileArrivalField, err := fileArrivalTriggerConfigurationFromPb(pb.FileArrival)
	if err != nil {
		return nil, err
	}
	if fileArrivalField != nil {
		st.FileArrival = fileArrivalField
	}
	pauseStatusField, err := identity(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	periodicField, err := periodicTriggerConfigurationFromPb(pb.Periodic)
	if err != nil {
		return nil, err
	}
	if periodicField != nil {
		st.Periodic = periodicField
	}
	tableField, err := tableUpdateTriggerConfigurationFromPb(pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = tableField
	}
	tableUpdateField, err := tableUpdateTriggerConfigurationFromPb(pb.TableUpdate)
	if err != nil {
		return nil, err
	}
	if tableUpdateField != nil {
		st.TableUpdate = tableUpdateField
	}

	return st, nil
}

// The type of trigger that fired this run.
//
// * `PERIODIC`: Schedules that periodically trigger runs, such as a cron
// scheduler. * `ONE_TIME`: One time triggers that fire a single run. This
// occurs you triggered a single run on demand through the UI or the API. *
// `RETRY`: Indicates a run that is triggered as a retry of a previously failed
// run. This occurs when you request to re-run the job in case of failures. *
// `RUN_JOB_TASK`: Indicates a run that is triggered using a Run Job task. *
// `FILE_ARRIVAL`: Indicates a run that is triggered by a file arrival. *
// `TABLE`: Indicates a run that is triggered by a table update. *
// `CONTINUOUS_RESTART`: Indicates a run created by user to manually restart a
// continuous job run.
type TriggerType string
type triggerTypePb string

// Indicates a run that is triggered by a file arrival.
const TriggerTypeFileArrival TriggerType = `FILE_ARRIVAL`

// One time triggers that fire a single run. This occurs you triggered a single
// run on demand through the UI or the API.
const TriggerTypeOneTime TriggerType = `ONE_TIME`

// Schedules that periodically trigger runs, such as a cron scheduler.
const TriggerTypePeriodic TriggerType = `PERIODIC`

// Indicates a run that is triggered as a retry of a previously failed run. This
// occurs when you request to re-run the job in case of failures.
const TriggerTypeRetry TriggerType = `RETRY`

// Indicates a run that is triggered using a Run Job task.
const TriggerTypeRunJobTask TriggerType = `RUN_JOB_TASK`

// Indicates a run that is triggered by a table update.
const TriggerTypeTable TriggerType = `TABLE`

// String representation for [fmt.Print]
func (f *TriggerType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TriggerType) Set(v string) error {
	switch v {
	case `FILE_ARRIVAL`, `ONE_TIME`, `PERIODIC`, `RETRY`, `RUN_JOB_TASK`, `TABLE`:
		*f = TriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FILE_ARRIVAL", "ONE_TIME", "PERIODIC", "RETRY", "RUN_JOB_TASK", "TABLE"`, v)
	}
}

// Type always returns TriggerType to satisfy [pflag.Value] interface
func (f *TriggerType) Type() string {
	return "TriggerType"
}

func triggerTypeToPb(st *TriggerType) (*triggerTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := triggerTypePb(*st)
	return &pb, nil
}

func triggerTypeFromPb(pb *triggerTypePb) (*TriggerType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TriggerType(*pb)
	return &st, nil
}

type UpdateJob struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported, except for tasks and job clusters (`tasks/task_1`). This
	// field is optional.
	FieldsToRemove []string
	// The canonical identifier of the job to update. This field is required.
	JobId int64
	// The new settings for the job.
	//
	// Top-level fields specified in `new_settings` are completely replaced,
	// except for arrays which are merged. That is, new and existing entries are
	// completely replaced based on the respective key fields, i.e. `task_key`
	// or `job_cluster_key`, while previous entries are kept.
	//
	// Partially updating nested fields is not supported.
	//
	// Changes to the field `JobSettings.timeout_seconds` are applied to active
	// runs. Changes to other fields are applied to future runs only.
	NewSettings *JobSettings
}

func updateJobToPb(st *UpdateJob) (*updateJobPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateJobPb{}

	var fieldsToRemovePb []string
	for _, item := range st.FieldsToRemove {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			fieldsToRemovePb = append(fieldsToRemovePb, *itemPb)
		}
	}
	pb.FieldsToRemove = fieldsToRemovePb

	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	newSettingsPb, err := jobSettingsToPb(st.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsPb != nil {
		pb.NewSettings = newSettingsPb
	}

	return pb, nil
}

func (st *UpdateJob) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateJobPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateJobFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateJob) MarshalJSON() ([]byte, error) {
	pb, err := updateJobToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateJobPb struct {
	// Remove top-level fields in the job settings. Removing nested fields is
	// not supported, except for tasks and job clusters (`tasks/task_1`). This
	// field is optional.
	FieldsToRemove []string `json:"fields_to_remove,omitempty"`
	// The canonical identifier of the job to update. This field is required.
	JobId int64 `json:"job_id"`
	// The new settings for the job.
	//
	// Top-level fields specified in `new_settings` are completely replaced,
	// except for arrays which are merged. That is, new and existing entries are
	// completely replaced based on the respective key fields, i.e. `task_key`
	// or `job_cluster_key`, while previous entries are kept.
	//
	// Partially updating nested fields is not supported.
	//
	// Changes to the field `JobSettings.timeout_seconds` are applied to active
	// runs. Changes to other fields are applied to future runs only.
	NewSettings *jobSettingsPb `json:"new_settings,omitempty"`
}

func updateJobFromPb(pb *updateJobPb) (*UpdateJob, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateJob{}

	var fieldsToRemoveField []string
	for _, itemPb := range pb.FieldsToRemove {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			fieldsToRemoveField = append(fieldsToRemoveField, *item)
		}
	}
	st.FieldsToRemove = fieldsToRemoveField
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}
	newSettingsField, err := jobSettingsFromPb(pb.NewSettings)
	if err != nil {
		return nil, err
	}
	if newSettingsField != nil {
		st.NewSettings = newSettingsField
	}

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

type ViewItem struct {
	// Content of the view.
	Content string
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	Name string
	// Type of the view item.
	Type ViewType

	ForceSendFields []string
}

func viewItemToPb(st *ViewItem) (*viewItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &viewItemPb{}
	contentPb, err := identity(&st.Content)
	if err != nil {
		return nil, err
	}
	if contentPb != nil {
		pb.Content = *contentPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
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

func (st *ViewItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &viewItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := viewItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ViewItem) MarshalJSON() ([]byte, error) {
	pb, err := viewItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type viewItemPb struct {
	// Content of the view.
	Content string `json:"content,omitempty"`
	// Name of the view item. In the case of code view, it would be the
	// notebook’s name. In the case of dashboard view, it would be the
	// dashboard’s name.
	Name string `json:"name,omitempty"`
	// Type of the view item.
	Type ViewType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func viewItemFromPb(pb *viewItemPb) (*ViewItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ViewItem{}
	contentField, err := identity(&pb.Content)
	if err != nil {
		return nil, err
	}
	if contentField != nil {
		st.Content = *contentField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
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

func (st *viewItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st viewItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// * `NOTEBOOK`: Notebook view item. * `DASHBOARD`: Dashboard view item.
type ViewType string
type viewTypePb string

// Dashboard view item.
const ViewTypeDashboard ViewType = `DASHBOARD`

// Notebook view item.
const ViewTypeNotebook ViewType = `NOTEBOOK`

// String representation for [fmt.Print]
func (f *ViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewType) Set(v string) error {
	switch v {
	case `DASHBOARD`, `NOTEBOOK`:
		*f = ViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DASHBOARD", "NOTEBOOK"`, v)
	}
}

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}

func viewTypeToPb(st *ViewType) (*viewTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := viewTypePb(*st)
	return &pb, nil
}

func viewTypeFromPb(pb *viewTypePb) (*ViewType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ViewType(*pb)
	return &st, nil
}

// * `CODE`: Code view of the notebook. * `DASHBOARDS`: All dashboard views of
// the notebook. * `ALL`: All views of the notebook.
type ViewsToExport string
type viewsToExportPb string

// All views of the notebook.
const ViewsToExportAll ViewsToExport = `ALL`

// Code view of the notebook.
const ViewsToExportCode ViewsToExport = `CODE`

// All dashboard views of the notebook.
const ViewsToExportDashboards ViewsToExport = `DASHBOARDS`

// String representation for [fmt.Print]
func (f *ViewsToExport) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewsToExport) Set(v string) error {
	switch v {
	case `ALL`, `CODE`, `DASHBOARDS`:
		*f = ViewsToExport(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL", "CODE", "DASHBOARDS"`, v)
	}
}

// Type always returns ViewsToExport to satisfy [pflag.Value] interface
func (f *ViewsToExport) Type() string {
	return "ViewsToExport"
}

func viewsToExportToPb(st *ViewsToExport) (*viewsToExportPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := viewsToExportPb(*st)
	return &pb, nil
}

func viewsToExportFromPb(pb *viewsToExportPb) (*ViewsToExport, error) {
	if pb == nil {
		return nil, nil
	}
	st := ViewsToExport(*pb)
	return &st, nil
}

type Webhook struct {
	Id string
}

func webhookToPb(st *Webhook) (*webhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &webhookPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	return pb, nil
}

func (st *Webhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &webhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := webhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Webhook) MarshalJSON() ([]byte, error) {
	pb, err := webhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type webhookPb struct {
	Id string `json:"id"`
}

func webhookFromPb(pb *webhookPb) (*Webhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Webhook{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

type WebhookNotifications struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	OnDurationWarningThresholdExceeded []Webhook
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	OnFailure []Webhook
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	OnStart []Webhook
	// An optional list of system notification IDs to call when any streaming
	// backlog thresholds are exceeded for any stream. Streaming backlog
	// thresholds can be set in the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes. A maximum of 3 destinations
	// can be specified for the `on_streaming_backlog_exceeded` property.
	OnStreamingBacklogExceeded []Webhook
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	OnSuccess []Webhook
}

func webhookNotificationsToPb(st *WebhookNotifications) (*webhookNotificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &webhookNotificationsPb{}

	var onDurationWarningThresholdExceededPb []webhookPb
	for _, item := range st.OnDurationWarningThresholdExceeded {
		itemPb, err := webhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onDurationWarningThresholdExceededPb = append(onDurationWarningThresholdExceededPb, *itemPb)
		}
	}
	pb.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededPb

	var onFailurePb []webhookPb
	for _, item := range st.OnFailure {
		itemPb, err := webhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onFailurePb = append(onFailurePb, *itemPb)
		}
	}
	pb.OnFailure = onFailurePb

	var onStartPb []webhookPb
	for _, item := range st.OnStart {
		itemPb, err := webhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStartPb = append(onStartPb, *itemPb)
		}
	}
	pb.OnStart = onStartPb

	var onStreamingBacklogExceededPb []webhookPb
	for _, item := range st.OnStreamingBacklogExceeded {
		itemPb, err := webhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onStreamingBacklogExceededPb = append(onStreamingBacklogExceededPb, *itemPb)
		}
	}
	pb.OnStreamingBacklogExceeded = onStreamingBacklogExceededPb

	var onSuccessPb []webhookPb
	for _, item := range st.OnSuccess {
		itemPb, err := webhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onSuccessPb = append(onSuccessPb, *itemPb)
		}
	}
	pb.OnSuccess = onSuccessPb

	return pb, nil
}

func (st *WebhookNotifications) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &webhookNotificationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := webhookNotificationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WebhookNotifications) MarshalJSON() ([]byte, error) {
	pb, err := webhookNotificationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type webhookNotificationsPb struct {
	// An optional list of system notification IDs to call when the duration of
	// a run exceeds the threshold specified for the `RUN_DURATION_SECONDS`
	// metric in the `health` field. A maximum of 3 destinations can be
	// specified for the `on_duration_warning_threshold_exceeded` property.
	OnDurationWarningThresholdExceeded []webhookPb `json:"on_duration_warning_threshold_exceeded,omitempty"`
	// An optional list of system notification IDs to call when the run fails. A
	// maximum of 3 destinations can be specified for the `on_failure` property.
	OnFailure []webhookPb `json:"on_failure,omitempty"`
	// An optional list of system notification IDs to call when the run starts.
	// A maximum of 3 destinations can be specified for the `on_start` property.
	OnStart []webhookPb `json:"on_start,omitempty"`
	// An optional list of system notification IDs to call when any streaming
	// backlog thresholds are exceeded for any stream. Streaming backlog
	// thresholds can be set in the `health` field using the following metrics:
	// `STREAMING_BACKLOG_BYTES`, `STREAMING_BACKLOG_RECORDS`,
	// `STREAMING_BACKLOG_SECONDS`, or `STREAMING_BACKLOG_FILES`. Alerting is
	// based on the 10-minute average of these metrics. If the issue persists,
	// notifications are resent every 30 minutes. A maximum of 3 destinations
	// can be specified for the `on_streaming_backlog_exceeded` property.
	OnStreamingBacklogExceeded []webhookPb `json:"on_streaming_backlog_exceeded,omitempty"`
	// An optional list of system notification IDs to call when the run
	// completes successfully. A maximum of 3 destinations can be specified for
	// the `on_success` property.
	OnSuccess []webhookPb `json:"on_success,omitempty"`
}

func webhookNotificationsFromPb(pb *webhookNotificationsPb) (*WebhookNotifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WebhookNotifications{}

	var onDurationWarningThresholdExceededField []Webhook
	for _, itemPb := range pb.OnDurationWarningThresholdExceeded {
		item, err := webhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onDurationWarningThresholdExceededField = append(onDurationWarningThresholdExceededField, *item)
		}
	}
	st.OnDurationWarningThresholdExceeded = onDurationWarningThresholdExceededField

	var onFailureField []Webhook
	for _, itemPb := range pb.OnFailure {
		item, err := webhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onFailureField = append(onFailureField, *item)
		}
	}
	st.OnFailure = onFailureField

	var onStartField []Webhook
	for _, itemPb := range pb.OnStart {
		item, err := webhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStartField = append(onStartField, *item)
		}
	}
	st.OnStart = onStartField

	var onStreamingBacklogExceededField []Webhook
	for _, itemPb := range pb.OnStreamingBacklogExceeded {
		item, err := webhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onStreamingBacklogExceededField = append(onStreamingBacklogExceededField, *item)
		}
	}
	st.OnStreamingBacklogExceeded = onStreamingBacklogExceededField

	var onSuccessField []Webhook
	for _, itemPb := range pb.OnSuccess {
		item, err := webhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onSuccessField = append(onSuccessField, *item)
		}
	}
	st.OnSuccess = onSuccessField

	return st, nil
}

type WidgetErrorDetail struct {
	Message string

	ForceSendFields []string
}

func widgetErrorDetailToPb(st *WidgetErrorDetail) (*widgetErrorDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetErrorDetailPb{}
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

func (st *WidgetErrorDetail) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &widgetErrorDetailPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := widgetErrorDetailFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WidgetErrorDetail) MarshalJSON() ([]byte, error) {
	pb, err := widgetErrorDetailToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type widgetErrorDetailPb struct {
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func widgetErrorDetailFromPb(pb *widgetErrorDetailPb) (*WidgetErrorDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WidgetErrorDetail{}
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

func (st *widgetErrorDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st widgetErrorDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
