// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AccessControl struct {

	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AccessControl) EncodeValues(key string, v *url.Values) error {
	pb, err := accessControlToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AccessControl) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accessControlPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accessControlFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccessControl) MarshalJSON() ([]byte, error) {
	pb, err := accessControlToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Aggregation string

const AggregationAvg Aggregation = `AVG`

const AggregationCount Aggregation = `COUNT`

const AggregationCountDistinct Aggregation = `COUNT_DISTINCT`

const AggregationMax Aggregation = `MAX`

const AggregationMedian Aggregation = `MEDIAN`

const AggregationMin Aggregation = `MIN`

const AggregationStddev Aggregation = `STDDEV`

const AggregationSum Aggregation = `SUM`

// String representation for [fmt.Print]
func (f *Aggregation) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Aggregation) Set(v string) error {
	switch v {
	case `AVG`, `COUNT`, `COUNT_DISTINCT`, `MAX`, `MEDIAN`, `MIN`, `STDDEV`, `SUM`:
		*f = Aggregation(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVG", "COUNT", "COUNT_DISTINCT", "MAX", "MEDIAN", "MIN", "STDDEV", "SUM"`, v)
	}
}

// Values returns all possible values for Aggregation.
//
// There is no guarantee on the order of the values in the slice.
func (f *Aggregation) Values() []Aggregation {
	return []Aggregation{
		AggregationAvg,
		AggregationCount,
		AggregationCountDistinct,
		AggregationMax,
		AggregationMedian,
		AggregationMin,
		AggregationStddev,
		AggregationSum,
	}
}

// Type always returns Aggregation to satisfy [pflag.Value] interface
func (f *Aggregation) Type() string {
	return "Aggregation"
}

type Alert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition `json:"condition,omitempty"`
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The workspace state of the alert. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// The workspace path of the folder containing the alert.
	// Wire name: 'parent_path'
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	// Wire name: 'state'
	State AlertState `json:"state,omitempty"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	// Wire name: 'trigger_time'
	TriggerTime string `json:"trigger_time,omitempty"`
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Alert) EncodeValues(key string, v *url.Values) error {
	pb, err := alertToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Alert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Alert) MarshalJSON() ([]byte, error) {
	pb, err := alertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertCondition struct {
	// Alert state if result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertState `json:"empty_result_state,omitempty"`
	// Operator used for comparison in alert evaluation.
	// Wire name: 'op'
	Op AlertOperator `json:"op,omitempty"`
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	// Wire name: 'operand'
	Operand *AlertConditionOperand `json:"operand,omitempty"`
	// Threshold value used for comparison in alert evaluation.
	// Wire name: 'threshold'
	Threshold *AlertConditionThreshold `json:"threshold,omitempty"`
}

func (st AlertCondition) EncodeValues(key string, v *url.Values) error {
	pb, err := alertConditionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertCondition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertConditionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertConditionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertCondition) MarshalJSON() ([]byte, error) {
	pb, err := alertConditionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertConditionOperand struct {

	// Wire name: 'column'
	Column *AlertOperandColumn `json:"column,omitempty"`
}

func (st AlertConditionOperand) EncodeValues(key string, v *url.Values) error {
	pb, err := alertConditionOperandToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertConditionOperand) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertConditionOperandPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertConditionOperandFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertConditionOperand) MarshalJSON() ([]byte, error) {
	pb, err := alertConditionOperandToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertConditionThreshold struct {

	// Wire name: 'value'
	Value *AlertOperandValue `json:"value,omitempty"`
}

func (st AlertConditionThreshold) EncodeValues(key string, v *url.Values) error {
	pb, err := alertConditionThresholdToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertConditionThreshold) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertConditionThresholdPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertConditionThresholdFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertConditionThreshold) MarshalJSON() ([]byte, error) {
	pb, err := alertConditionThresholdToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// UNSPECIFIED - default unspecify value for proto enum, do not use it in the
// code UNKNOWN - alert not yet evaluated TRIGGERED - alert is triggered OK -
// alert is not triggered ERROR - alert evaluation failed
type AlertEvaluationState string

const AlertEvaluationStateError AlertEvaluationState = `ERROR`

const AlertEvaluationStateOk AlertEvaluationState = `OK`

const AlertEvaluationStateTriggered AlertEvaluationState = `TRIGGERED`

const AlertEvaluationStateUnknown AlertEvaluationState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *AlertEvaluationState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertEvaluationState) Set(v string) error {
	switch v {
	case `ERROR`, `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = AlertEvaluationState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ERROR", "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for AlertEvaluationState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertEvaluationState) Values() []AlertEvaluationState {
	return []AlertEvaluationState{
		AlertEvaluationStateError,
		AlertEvaluationStateOk,
		AlertEvaluationStateTriggered,
		AlertEvaluationStateUnknown,
	}
}

// Type always returns AlertEvaluationState to satisfy [pflag.Value] interface
func (f *AlertEvaluationState) Type() string {
	return "AlertEvaluationState"
}

type AlertOperandColumn struct {

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertOperandColumn) EncodeValues(key string, v *url.Values) error {
	pb, err := alertOperandColumnToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertOperandColumn) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertOperandColumnPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertOperandColumnFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertOperandColumn) MarshalJSON() ([]byte, error) {
	pb, err := alertOperandColumnToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertOperandValue struct {

	// Wire name: 'bool_value'
	BoolValue bool `json:"bool_value,omitempty"`

	// Wire name: 'double_value'
	DoubleValue float64 `json:"double_value,omitempty"`

	// Wire name: 'string_value'
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertOperandValue) EncodeValues(key string, v *url.Values) error {
	pb, err := alertOperandValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertOperandValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertOperandValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertOperandValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertOperandValue) MarshalJSON() ([]byte, error) {
	pb, err := alertOperandValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertOperator string

const AlertOperatorEqual AlertOperator = `EQUAL`

const AlertOperatorGreaterThan AlertOperator = `GREATER_THAN`

const AlertOperatorGreaterThanOrEqual AlertOperator = `GREATER_THAN_OR_EQUAL`

const AlertOperatorIsNull AlertOperator = `IS_NULL`

const AlertOperatorLessThan AlertOperator = `LESS_THAN`

const AlertOperatorLessThanOrEqual AlertOperator = `LESS_THAN_OR_EQUAL`

const AlertOperatorNotEqual AlertOperator = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *AlertOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertOperator) Set(v string) error {
	switch v {
	case `EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_NULL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = AlertOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "IS_NULL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Values returns all possible values for AlertOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertOperator) Values() []AlertOperator {
	return []AlertOperator{
		AlertOperatorEqual,
		AlertOperatorGreaterThan,
		AlertOperatorGreaterThanOrEqual,
		AlertOperatorIsNull,
		AlertOperatorLessThan,
		AlertOperatorLessThanOrEqual,
		AlertOperatorNotEqual,
	}
}

// Type always returns AlertOperator to satisfy [pflag.Value] interface
func (f *AlertOperator) Type() string {
	return "AlertOperator"
}

// Alert configuration options.
type AlertOptions struct {
	// Name of column in the query result to compare in alert evaluation.
	// Wire name: 'column'
	Column string `json:"column"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string `json:"custom_subject,omitempty"`
	// State that alert evaluates to when query result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertOptionsEmptyResultState `json:"empty_result_state,omitempty"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	// Wire name: 'muted'
	Muted bool `json:"muted,omitempty"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	// Wire name: 'op'
	Op string `json:"op"`
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	// Wire name: 'value'
	Value any `json:"value"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertOptions) EncodeValues(key string, v *url.Values) error {
	pb, err := alertOptionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertOptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertOptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertOptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertOptions) MarshalJSON() ([]byte, error) {
	pb, err := alertOptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// State that alert evaluates to when query result is empty.
type AlertOptionsEmptyResultState string

const AlertOptionsEmptyResultStateOk AlertOptionsEmptyResultState = `ok`

const AlertOptionsEmptyResultStateTriggered AlertOptionsEmptyResultState = `triggered`

const AlertOptionsEmptyResultStateUnknown AlertOptionsEmptyResultState = `unknown`

// String representation for [fmt.Print]
func (f *AlertOptionsEmptyResultState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertOptionsEmptyResultState) Set(v string) error {
	switch v {
	case `ok`, `triggered`, `unknown`:
		*f = AlertOptionsEmptyResultState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ok", "triggered", "unknown"`, v)
	}
}

// Values returns all possible values for AlertOptionsEmptyResultState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertOptionsEmptyResultState) Values() []AlertOptionsEmptyResultState {
	return []AlertOptionsEmptyResultState{
		AlertOptionsEmptyResultStateOk,
		AlertOptionsEmptyResultStateTriggered,
		AlertOptionsEmptyResultStateUnknown,
	}
}

// Type always returns AlertOptionsEmptyResultState to satisfy [pflag.Value] interface
func (f *AlertOptionsEmptyResultState) Type() string {
	return "AlertOptionsEmptyResultState"
}

type AlertQuery struct {
	// The timestamp when this query was created.
	// Wire name: 'created_at'
	CreatedAt string `json:"created_at,omitempty"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Query ID.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	// Wire name: 'is_draft'
	IsDraft bool `json:"is_draft,omitempty"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	// Wire name: 'is_safe'
	IsSafe bool `json:"is_safe,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'options'
	Options *QueryOptions `json:"options,omitempty"`
	// The text of the query to be run.
	// Wire name: 'query'
	Query string `json:"query,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`
	// The timestamp at which this query was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the user who owns the query.
	// Wire name: 'user_id'
	UserId int `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertQuery) EncodeValues(key string, v *url.Values) error {
	pb, err := alertQueryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertQuery) MarshalJSON() ([]byte, error) {
	pb, err := alertQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertState string

const AlertStateOk AlertState = `OK`

const AlertStateTriggered AlertState = `TRIGGERED`

const AlertStateUnknown AlertState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *AlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertState) Set(v string) error {
	switch v {
	case `OK`, `TRIGGERED`, `UNKNOWN`:
		*f = AlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OK", "TRIGGERED", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for AlertState.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertState) Values() []AlertState {
	return []AlertState{
		AlertStateOk,
		AlertStateTriggered,
		AlertStateUnknown,
	}
}

// Type always returns AlertState to satisfy [pflag.Value] interface
func (f *AlertState) Type() string {
	return "AlertState"
}

type AlertV2 struct {
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// Custom description for the alert. support mustache template.
	// Wire name: 'custom_description'
	CustomDescription string `json:"custom_description,omitempty"`
	// Custom summary for the alert. support mustache template.
	// Wire name: 'custom_summary'
	CustomSummary string `json:"custom_summary,omitempty"`
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`

	// Wire name: 'evaluation'
	Evaluation *AlertV2Evaluation `json:"evaluation,omitempty"`
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	// Wire name: 'parent_path'
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string `json:"query_text,omitempty"`
	// The run as username or application ID of service principal. This field is
	// set to "Unavailable" if the user has been deleted. On Create and Update,
	// this field can be set to application ID of an active service principal.
	// Setting this field requires the servicePrincipal/user role. If not
	// specified it'll default to be request user.
	// Wire name: 'run_as_user_name'
	RunAsUserName string `json:"run_as_user_name,omitempty"`

	// Wire name: 'schedule'
	Schedule *CronSchedule `json:"schedule,omitempty"`
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the alert.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertV2) EncodeValues(key string, v *url.Values) error {
	pb, err := alertV2ToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertV2) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertV2Pb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertV2FromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertV2) MarshalJSON() ([]byte, error) {
	pb, err := alertV2ToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertV2Evaluation struct {
	// Operator used for comparison in alert evaluation.
	// Wire name: 'comparison_operator'
	ComparisonOperator ComparisonOperator `json:"comparison_operator,omitempty"`
	// Alert state if result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertEvaluationState `json:"empty_result_state,omitempty"`
	// Timestamp of the last evaluation.
	// Wire name: 'last_evaluated_at'
	LastEvaluatedAt string `json:"last_evaluated_at,omitempty"`
	// User or Notification Destination to notify when alert is triggered.
	// Wire name: 'notification'
	Notification *AlertV2Notification `json:"notification,omitempty"`
	// Source column from result to use to evaluate alert
	// Wire name: 'source'
	Source *AlertV2OperandColumn `json:"source,omitempty"`
	// Latest state of alert evaluation.
	// Wire name: 'state'
	State AlertEvaluationState `json:"state,omitempty"`
	// Threshold to user for alert evaluation, can be a column or a value.
	// Wire name: 'threshold'
	Threshold *AlertV2Operand `json:"threshold,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertV2Evaluation) EncodeValues(key string, v *url.Values) error {
	pb, err := alertV2EvaluationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertV2Evaluation) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertV2EvaluationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertV2EvaluationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertV2Evaluation) MarshalJSON() ([]byte, error) {
	pb, err := alertV2EvaluationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertV2Notification struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'retrigger_seconds'
	RetriggerSeconds int `json:"retrigger_seconds,omitempty"`

	// Wire name: 'subscriptions'
	Subscriptions []AlertV2Subscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertV2Notification) EncodeValues(key string, v *url.Values) error {
	pb, err := alertV2NotificationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertV2Notification) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertV2NotificationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertV2NotificationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertV2Notification) MarshalJSON() ([]byte, error) {
	pb, err := alertV2NotificationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertV2Operand struct {

	// Wire name: 'column'
	Column *AlertV2OperandColumn `json:"column,omitempty"`

	// Wire name: 'value'
	Value *AlertV2OperandValue `json:"value,omitempty"`
}

func (st AlertV2Operand) EncodeValues(key string, v *url.Values) error {
	pb, err := alertV2OperandToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertV2Operand) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertV2OperandPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertV2OperandFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertV2Operand) MarshalJSON() ([]byte, error) {
	pb, err := alertV2OperandToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertV2OperandColumn struct {

	// Wire name: 'aggregation'
	Aggregation Aggregation `json:"aggregation,omitempty"`

	// Wire name: 'display'
	Display string `json:"display,omitempty"`

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertV2OperandColumn) EncodeValues(key string, v *url.Values) error {
	pb, err := alertV2OperandColumnToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertV2OperandColumn) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertV2OperandColumnPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertV2OperandColumnFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertV2OperandColumn) MarshalJSON() ([]byte, error) {
	pb, err := alertV2OperandColumnToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertV2OperandValue struct {

	// Wire name: 'bool_value'
	BoolValue bool `json:"bool_value,omitempty"`

	// Wire name: 'double_value'
	DoubleValue float64 `json:"double_value,omitempty"`

	// Wire name: 'string_value'
	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertV2OperandValue) EncodeValues(key string, v *url.Values) error {
	pb, err := alertV2OperandValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertV2OperandValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertV2OperandValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertV2OperandValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertV2OperandValue) MarshalJSON() ([]byte, error) {
	pb, err := alertV2OperandValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AlertV2Subscription struct {

	// Wire name: 'destination_id'
	DestinationId string `json:"destination_id,omitempty"`

	// Wire name: 'user_email'
	UserEmail string `json:"user_email,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AlertV2Subscription) EncodeValues(key string, v *url.Values) error {
	pb, err := alertV2SubscriptionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AlertV2Subscription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertV2SubscriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertV2SubscriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertV2Subscription) MarshalJSON() ([]byte, error) {
	pb, err := alertV2SubscriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Describes metadata for a particular chunk, within a result set; this
// structure is used both within a manifest, and when fetching individual chunk
// data or links.
type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st BaseChunkInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := baseChunkInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *BaseChunkInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &baseChunkInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := baseChunkInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BaseChunkInfo) MarshalJSON() ([]byte, error) {
	pb, err := baseChunkInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" tf:"-"`
}

func (st CancelExecutionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := cancelExecutionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CancelExecutionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelExecutionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelExecutionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelExecutionRequest) MarshalJSON() ([]byte, error) {
	pb, err := cancelExecutionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel struct {

	// Wire name: 'dbsql_version'
	DbsqlVersion string `json:"dbsql_version,omitempty"`

	// Wire name: 'name'
	Name ChannelName `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Channel) EncodeValues(key string, v *url.Values) error {
	pb, err := channelToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Channel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &channelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := channelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Channel) MarshalJSON() ([]byte, error) {
	pb, err := channelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details about a Channel.
type ChannelInfo struct {
	// DB SQL Version the Channel is mapped to.
	// Wire name: 'dbsql_version'
	DbsqlVersion string `json:"dbsql_version,omitempty"`
	// Name of the channel
	// Wire name: 'name'
	Name ChannelName `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ChannelInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := channelInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ChannelInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &channelInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := channelInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ChannelInfo) MarshalJSON() ([]byte, error) {
	pb, err := channelInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ChannelName string

const ChannelNameChannelNameCurrent ChannelName = `CHANNEL_NAME_CURRENT`

const ChannelNameChannelNameCustom ChannelName = `CHANNEL_NAME_CUSTOM`

const ChannelNameChannelNamePreview ChannelName = `CHANNEL_NAME_PREVIEW`

const ChannelNameChannelNamePrevious ChannelName = `CHANNEL_NAME_PREVIOUS`

// String representation for [fmt.Print]
func (f *ChannelName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ChannelName) Set(v string) error {
	switch v {
	case `CHANNEL_NAME_CURRENT`, `CHANNEL_NAME_CUSTOM`, `CHANNEL_NAME_PREVIEW`, `CHANNEL_NAME_PREVIOUS`:
		*f = ChannelName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CHANNEL_NAME_CURRENT", "CHANNEL_NAME_CUSTOM", "CHANNEL_NAME_PREVIEW", "CHANNEL_NAME_PREVIOUS"`, v)
	}
}

// Values returns all possible values for ChannelName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ChannelName) Values() []ChannelName {
	return []ChannelName{
		ChannelNameChannelNameCurrent,
		ChannelNameChannelNameCustom,
		ChannelNameChannelNamePreview,
		ChannelNameChannelNamePrevious,
	}
}

// Type always returns ChannelName to satisfy [pflag.Value] interface
func (f *ChannelName) Type() string {
	return "ChannelName"
}

type ClientConfig struct {

	// Wire name: 'allow_custom_js_visualizations'
	AllowCustomJsVisualizations bool `json:"allow_custom_js_visualizations,omitempty"`

	// Wire name: 'allow_downloads'
	AllowDownloads bool `json:"allow_downloads,omitempty"`

	// Wire name: 'allow_external_shares'
	AllowExternalShares bool `json:"allow_external_shares,omitempty"`

	// Wire name: 'allow_subscriptions'
	AllowSubscriptions bool `json:"allow_subscriptions,omitempty"`

	// Wire name: 'date_format'
	DateFormat string `json:"date_format,omitempty"`

	// Wire name: 'date_time_format'
	DateTimeFormat string `json:"date_time_format,omitempty"`

	// Wire name: 'disable_publish'
	DisablePublish bool `json:"disable_publish,omitempty"`

	// Wire name: 'enable_legacy_autodetect_types'
	EnableLegacyAutodetectTypes bool `json:"enable_legacy_autodetect_types,omitempty"`

	// Wire name: 'feature_show_permissions_control'
	FeatureShowPermissionsControl bool `json:"feature_show_permissions_control,omitempty"`

	// Wire name: 'hide_plotly_mode_bar'
	HidePlotlyModeBar bool `json:"hide_plotly_mode_bar,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ClientConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := clientConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ClientConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clientConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clientConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClientConfig) MarshalJSON() ([]byte, error) {
	pb, err := clientConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ColumnInfo struct {
	// The name of the column.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The ordinal position of the column (starting at position 0).
	// Wire name: 'position'
	Position int `json:"position,omitempty"`
	// The format of the interval type.
	// Wire name: 'type_interval_type'
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	// Wire name: 'type_name'
	TypeName ColumnInfoTypeName `json:"type_name,omitempty"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	// Wire name: 'type_precision'
	TypePrecision int `json:"type_precision,omitempty"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	// Wire name: 'type_scale'
	TypeScale int `json:"type_scale,omitempty"`
	// The full SQL type specification.
	// Wire name: 'type_text'
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ColumnInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := columnInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ColumnInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &columnInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := columnInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ColumnInfo) MarshalJSON() ([]byte, error) {
	pb, err := columnInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The name of the base data type. This doesn't include details for complex
// types such as STRUCT, MAP or ARRAY.
type ColumnInfoTypeName string

const ColumnInfoTypeNameArray ColumnInfoTypeName = `ARRAY`

const ColumnInfoTypeNameBinary ColumnInfoTypeName = `BINARY`

const ColumnInfoTypeNameBoolean ColumnInfoTypeName = `BOOLEAN`

const ColumnInfoTypeNameByte ColumnInfoTypeName = `BYTE`

const ColumnInfoTypeNameChar ColumnInfoTypeName = `CHAR`

const ColumnInfoTypeNameDate ColumnInfoTypeName = `DATE`

const ColumnInfoTypeNameDecimal ColumnInfoTypeName = `DECIMAL`

const ColumnInfoTypeNameDouble ColumnInfoTypeName = `DOUBLE`

const ColumnInfoTypeNameFloat ColumnInfoTypeName = `FLOAT`

const ColumnInfoTypeNameInt ColumnInfoTypeName = `INT`

const ColumnInfoTypeNameInterval ColumnInfoTypeName = `INTERVAL`

const ColumnInfoTypeNameLong ColumnInfoTypeName = `LONG`

const ColumnInfoTypeNameMap ColumnInfoTypeName = `MAP`

const ColumnInfoTypeNameNull ColumnInfoTypeName = `NULL`

const ColumnInfoTypeNameShort ColumnInfoTypeName = `SHORT`

const ColumnInfoTypeNameString ColumnInfoTypeName = `STRING`

const ColumnInfoTypeNameStruct ColumnInfoTypeName = `STRUCT`

const ColumnInfoTypeNameTimestamp ColumnInfoTypeName = `TIMESTAMP`

const ColumnInfoTypeNameUserDefinedType ColumnInfoTypeName = `USER_DEFINED_TYPE`

// String representation for [fmt.Print]
func (f *ColumnInfoTypeName) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ColumnInfoTypeName) Set(v string) error {
	switch v {
	case `ARRAY`, `BINARY`, `BOOLEAN`, `BYTE`, `CHAR`, `DATE`, `DECIMAL`, `DOUBLE`, `FLOAT`, `INT`, `INTERVAL`, `LONG`, `MAP`, `NULL`, `SHORT`, `STRING`, `STRUCT`, `TIMESTAMP`, `USER_DEFINED_TYPE`:
		*f = ColumnInfoTypeName(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARRAY", "BINARY", "BOOLEAN", "BYTE", "CHAR", "DATE", "DECIMAL", "DOUBLE", "FLOAT", "INT", "INTERVAL", "LONG", "MAP", "NULL", "SHORT", "STRING", "STRUCT", "TIMESTAMP", "USER_DEFINED_TYPE"`, v)
	}
}

// Values returns all possible values for ColumnInfoTypeName.
//
// There is no guarantee on the order of the values in the slice.
func (f *ColumnInfoTypeName) Values() []ColumnInfoTypeName {
	return []ColumnInfoTypeName{
		ColumnInfoTypeNameArray,
		ColumnInfoTypeNameBinary,
		ColumnInfoTypeNameBoolean,
		ColumnInfoTypeNameByte,
		ColumnInfoTypeNameChar,
		ColumnInfoTypeNameDate,
		ColumnInfoTypeNameDecimal,
		ColumnInfoTypeNameDouble,
		ColumnInfoTypeNameFloat,
		ColumnInfoTypeNameInt,
		ColumnInfoTypeNameInterval,
		ColumnInfoTypeNameLong,
		ColumnInfoTypeNameMap,
		ColumnInfoTypeNameNull,
		ColumnInfoTypeNameShort,
		ColumnInfoTypeNameString,
		ColumnInfoTypeNameStruct,
		ColumnInfoTypeNameTimestamp,
		ColumnInfoTypeNameUserDefinedType,
	}
}

// Type always returns ColumnInfoTypeName to satisfy [pflag.Value] interface
func (f *ColumnInfoTypeName) Type() string {
	return "ColumnInfoTypeName"
}

type ComparisonOperator string

const ComparisonOperatorEqual ComparisonOperator = `EQUAL`

const ComparisonOperatorGreaterThan ComparisonOperator = `GREATER_THAN`

const ComparisonOperatorGreaterThanOrEqual ComparisonOperator = `GREATER_THAN_OR_EQUAL`

const ComparisonOperatorIsNotNull ComparisonOperator = `IS_NOT_NULL`

const ComparisonOperatorIsNull ComparisonOperator = `IS_NULL`

const ComparisonOperatorLessThan ComparisonOperator = `LESS_THAN`

const ComparisonOperatorLessThanOrEqual ComparisonOperator = `LESS_THAN_OR_EQUAL`

const ComparisonOperatorNotEqual ComparisonOperator = `NOT_EQUAL`

// String representation for [fmt.Print]
func (f *ComparisonOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComparisonOperator) Set(v string) error {
	switch v {
	case `EQUAL`, `GREATER_THAN`, `GREATER_THAN_OR_EQUAL`, `IS_NOT_NULL`, `IS_NULL`, `LESS_THAN`, `LESS_THAN_OR_EQUAL`, `NOT_EQUAL`:
		*f = ComparisonOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", "IS_NOT_NULL", "IS_NULL", "LESS_THAN", "LESS_THAN_OR_EQUAL", "NOT_EQUAL"`, v)
	}
}

// Values returns all possible values for ComparisonOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		ComparisonOperatorEqual,
		ComparisonOperatorGreaterThan,
		ComparisonOperatorGreaterThanOrEqual,
		ComparisonOperatorIsNotNull,
		ComparisonOperatorIsNull,
		ComparisonOperatorLessThan,
		ComparisonOperatorLessThanOrEqual,
		ComparisonOperatorNotEqual,
	}
}

// Type always returns ComparisonOperator to satisfy [pflag.Value] interface
func (f *ComparisonOperator) Type() string {
	return "ComparisonOperator"
}

type CreateAlert struct {
	// Name of the alert.
	// Wire name: 'name'
	Name string `json:"name"`
	// Alert configuration options.
	// Wire name: 'options'
	Options AlertOptions `json:"options"`
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string `json:"parent,omitempty"`
	// Query ID.
	// Wire name: 'query_id'
	QueryId string `json:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm int `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateAlert) EncodeValues(key string, v *url.Values) error {
	pb, err := createAlertToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAlert) MarshalJSON() ([]byte, error) {
	pb, err := createAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateAlertRequest struct {

	// Wire name: 'alert'
	Alert *CreateAlertRequestAlert `json:"alert,omitempty"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateAlertRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createAlertRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateAlertRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAlertRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAlertRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAlertRequest) MarshalJSON() ([]byte, error) {
	pb, err := createAlertRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition `json:"condition,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The workspace path of the folder containing the alert.
	// Wire name: 'parent_path'
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateAlertRequestAlert) EncodeValues(key string, v *url.Values) error {
	pb, err := createAlertRequestAlertToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateAlertRequestAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAlertRequestAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAlertRequestAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAlertRequestAlert) MarshalJSON() ([]byte, error) {
	pb, err := createAlertRequestAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateAlertV2Request struct {

	// Wire name: 'alert'
	Alert AlertV2 `json:"alert"`
}

func (st CreateAlertV2Request) EncodeValues(key string, v *url.Values) error {
	pb, err := createAlertV2RequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateAlertV2Request) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAlertV2RequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAlertV2RequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAlertV2Request) MarshalJSON() ([]byte, error) {
	pb, err := createAlertV2RequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateQueryRequest struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`

	// Wire name: 'query'
	Query *CreateQueryRequestQuery `json:"query,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateQueryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createQueryRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := createQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Workspace path of the workspace folder containing the object.
	// Wire name: 'parent_path'
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateQueryRequestQuery) EncodeValues(key string, v *url.Values) error {
	pb, err := createQueryRequestQueryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateQueryRequestQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createQueryRequestQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createQueryRequestQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateQueryRequestQuery) MarshalJSON() ([]byte, error) {
	pb, err := createQueryRequestQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	// Wire name: 'options'
	Options any `json:"options"`
	// The identifier returned by :method:queries/create
	// Wire name: 'query_id'
	QueryId string `json:"query_id"`
	// The type of visualization: chart, table, pivot table, and so on.
	// Wire name: 'type'
	Type string `json:"type"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateQueryVisualizationsLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createQueryVisualizationsLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateQueryVisualizationsLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createQueryVisualizationsLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createQueryVisualizationsLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateQueryVisualizationsLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createQueryVisualizationsLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateVisualizationRequest struct {

	// Wire name: 'visualization'
	Visualization *CreateVisualizationRequestVisualization `json:"visualization,omitempty"`
}

func (st CreateVisualizationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createVisualizationRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateVisualizationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createVisualizationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createVisualizationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateVisualizationRequest) MarshalJSON() ([]byte, error) {
	pb, err := createVisualizationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID of the query that the visualization is attached to.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateVisualizationRequestVisualization) EncodeValues(key string, v *url.Values) error {
	pb, err := createVisualizationRequestVisualizationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateVisualizationRequestVisualization) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createVisualizationRequestVisualizationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createVisualizationRequestVisualizationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateVisualizationRequestVisualization) MarshalJSON() ([]byte, error) {
	pb, err := createVisualizationRequestVisualizationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be >= 0 mins for serverless warehouses - Must be
	// == 0 or >= 10 mins for non-serverless warehouses - 0 indicates no
	// autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	// Wire name: 'min_num_clusters'
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags `json:"tags,omitempty"`

	// Wire name: 'warehouse_type'
	WarehouseType CreateWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateWarehouseRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createWarehouseRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateWarehouseRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createWarehouseRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createWarehouseRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateWarehouseRequest) MarshalJSON() ([]byte, error) {
	pb, err := createWarehouseRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type CreateWarehouseRequestWarehouseType string

const CreateWarehouseRequestWarehouseTypeClassic CreateWarehouseRequestWarehouseType = `CLASSIC`

const CreateWarehouseRequestWarehouseTypePro CreateWarehouseRequestWarehouseType = `PRO`

const CreateWarehouseRequestWarehouseTypeTypeUnspecified CreateWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *CreateWarehouseRequestWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreateWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = CreateWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for CreateWarehouseRequestWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *CreateWarehouseRequestWarehouseType) Values() []CreateWarehouseRequestWarehouseType {
	return []CreateWarehouseRequestWarehouseType{
		CreateWarehouseRequestWarehouseTypeClassic,
		CreateWarehouseRequestWarehouseTypePro,
		CreateWarehouseRequestWarehouseTypeTypeUnspecified,
	}
}

// Type always returns CreateWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *CreateWarehouseRequestWarehouseType) Type() string {
	return "CreateWarehouseRequestWarehouseType"
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateWarehouseResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createWarehouseResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateWarehouseResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createWarehouseResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createWarehouseResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateWarehouseResponse) MarshalJSON() ([]byte, error) {
	pb, err := createWarehouseResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id"`

	// Wire name: 'options'
	Options WidgetOptions `json:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	// Wire name: 'text'
	Text string `json:"text,omitempty"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	// Wire name: 'visualization_id'
	VisualizationId string `json:"visualization_id,omitempty"`
	// Width of a widget
	// Wire name: 'width'
	Width int `json:"width"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateWidget) EncodeValues(key string, v *url.Values) error {
	pb, err := createWidgetToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateWidget) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createWidgetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createWidgetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateWidget) MarshalJSON() ([]byte, error) {
	pb, err := createWidgetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`
	// A cron expression using quartz syntax that specifies the schedule for
	// this pipeline. Should use the quartz format described here:
	// http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
	// Wire name: 'quartz_cron_schedule'
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`
	// A Java timezone id. The schedule will be resolved using this timezone.
	// This will be combined with the quartz_cron_schedule to determine the
	// schedule. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details.
	// Wire name: 'timezone_id'
	TimezoneId string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CronSchedule) EncodeValues(key string, v *url.Values) error {
	pb, err := cronScheduleToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	// Wire name: 'can_edit'
	CanEdit bool `json:"can_edit,omitempty"`
	// Timestamp when this dashboard was created.
	// Wire name: 'created_at'
	CreatedAt string `json:"created_at,omitempty"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	// Wire name: 'dashboard_filters_enabled'
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled,omitempty"`
	// The ID for this dashboard.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	// Wire name: 'is_draft'
	IsDraft bool `json:"is_draft,omitempty"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	// Wire name: 'is_favorite'
	IsFavorite bool `json:"is_favorite,omitempty"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'options'
	Options *DashboardOptions `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string `json:"parent,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_tier'
	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	// Wire name: 'slug'
	Slug string `json:"slug,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`
	// Timestamp when this dashboard was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string `json:"updated_at,omitempty"`

	// Wire name: 'user'
	User *User `json:"user,omitempty"`
	// The ID of the user who owns the dashboard.
	// Wire name: 'user_id'
	UserId int `json:"user_id,omitempty"`

	// Wire name: 'widgets'
	Widgets []Widget `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Dashboard) EncodeValues(key string, v *url.Values) error {
	pb, err := dashboardToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Dashboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Dashboard) MarshalJSON() ([]byte, error) {
	pb, err := dashboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DashboardEditContent struct {
	DashboardId string `json:"-" tf:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DashboardEditContent) EncodeValues(key string, v *url.Values) error {
	pb, err := dashboardEditContentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DashboardEditContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardEditContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardEditContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardEditContent) MarshalJSON() ([]byte, error) {
	pb, err := dashboardEditContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	// Wire name: 'moved_to_trash_at'
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DashboardOptions) EncodeValues(key string, v *url.Values) error {
	pb, err := dashboardOptionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DashboardOptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardOptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardOptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardOptions) MarshalJSON() ([]byte, error) {
	pb, err := dashboardOptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DashboardPostContent struct {
	// Indicates whether the dashboard filters are enabled
	// Wire name: 'dashboard_filters_enabled'
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled,omitempty"`
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	// Wire name: 'is_favorite'
	IsFavorite bool `json:"is_favorite,omitempty"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string `json:"name"`
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string `json:"parent,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DashboardPostContent) EncodeValues(key string, v *url.Values) error {
	pb, err := dashboardPostContentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DashboardPostContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardPostContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardPostContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardPostContent) MarshalJSON() ([]byte, error) {
	pb, err := dashboardPostContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Reserved for internal use.
	// Wire name: 'pause_reason'
	PauseReason string `json:"pause_reason,omitempty"`
	// Reserved for internal use.
	// Wire name: 'paused'
	Paused int `json:"paused,omitempty"`
	// Reserved for internal use.
	// Wire name: 'supports_auto_limit'
	SupportsAutoLimit bool `json:"supports_auto_limit,omitempty"`
	// Reserved for internal use.
	// Wire name: 'syntax'
	Syntax string `json:"syntax,omitempty"`
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	// Wire name: 'type'
	Type string `json:"type,omitempty"`
	// Reserved for internal use.
	// Wire name: 'view_only'
	ViewOnly bool `json:"view_only,omitempty"`
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DataSource) EncodeValues(key string, v *url.Values) error {
	pb, err := dataSourceToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DataSource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dataSourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dataSourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DataSource) MarshalJSON() ([]byte, error) {
	pb, err := dataSourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DatePrecision string

const DatePrecisionDayPrecision DatePrecision = `DAY_PRECISION`

const DatePrecisionMinutePrecision DatePrecision = `MINUTE_PRECISION`

const DatePrecisionSecondPrecision DatePrecision = `SECOND_PRECISION`

// String representation for [fmt.Print]
func (f *DatePrecision) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DatePrecision) Set(v string) error {
	switch v {
	case `DAY_PRECISION`, `MINUTE_PRECISION`, `SECOND_PRECISION`:
		*f = DatePrecision(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DAY_PRECISION", "MINUTE_PRECISION", "SECOND_PRECISION"`, v)
	}
}

// Values returns all possible values for DatePrecision.
//
// There is no guarantee on the order of the values in the slice.
func (f *DatePrecision) Values() []DatePrecision {
	return []DatePrecision{
		DatePrecisionDayPrecision,
		DatePrecisionMinutePrecision,
		DatePrecisionSecondPrecision,
	}
}

// Type always returns DatePrecision to satisfy [pflag.Value] interface
func (f *DatePrecision) Type() string {
	return "DatePrecision"
}

type DateRange struct {

	// Wire name: 'end'
	End string `json:"end"`

	// Wire name: 'start'
	Start string `json:"start"`
}

func (st DateRange) EncodeValues(key string, v *url.Values) error {
	pb, err := dateRangeToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DateRange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dateRangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dateRangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DateRange) MarshalJSON() ([]byte, error) {
	pb, err := dateRangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DateRangeValue struct {
	// Manually specified date-time range value.
	// Wire name: 'date_range_value'
	DateRangeValue *DateRange `json:"date_range_value,omitempty"`
	// Dynamic date-time range value based on current date-time.
	// Wire name: 'dynamic_date_range_value'
	DynamicDateRangeValue DateRangeValueDynamicDateRange `json:"dynamic_date_range_value,omitempty"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	// Wire name: 'precision'
	Precision DatePrecision `json:"precision,omitempty"`

	// Wire name: 'start_day_of_week'
	StartDayOfWeek int `json:"start_day_of_week,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DateRangeValue) EncodeValues(key string, v *url.Values) error {
	pb, err := dateRangeValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DateRangeValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dateRangeValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dateRangeValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DateRangeValue) MarshalJSON() ([]byte, error) {
	pb, err := dateRangeValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DateRangeValueDynamicDateRange string

const DateRangeValueDynamicDateRangeLast12Months DateRangeValueDynamicDateRange = `LAST_12_MONTHS`

const DateRangeValueDynamicDateRangeLast14Days DateRangeValueDynamicDateRange = `LAST_14_DAYS`

const DateRangeValueDynamicDateRangeLast24Hours DateRangeValueDynamicDateRange = `LAST_24_HOURS`

const DateRangeValueDynamicDateRangeLast30Days DateRangeValueDynamicDateRange = `LAST_30_DAYS`

const DateRangeValueDynamicDateRangeLast60Days DateRangeValueDynamicDateRange = `LAST_60_DAYS`

const DateRangeValueDynamicDateRangeLast7Days DateRangeValueDynamicDateRange = `LAST_7_DAYS`

const DateRangeValueDynamicDateRangeLast8Hours DateRangeValueDynamicDateRange = `LAST_8_HOURS`

const DateRangeValueDynamicDateRangeLast90Days DateRangeValueDynamicDateRange = `LAST_90_DAYS`

const DateRangeValueDynamicDateRangeLastHour DateRangeValueDynamicDateRange = `LAST_HOUR`

const DateRangeValueDynamicDateRangeLastMonth DateRangeValueDynamicDateRange = `LAST_MONTH`

const DateRangeValueDynamicDateRangeLastWeek DateRangeValueDynamicDateRange = `LAST_WEEK`

const DateRangeValueDynamicDateRangeLastYear DateRangeValueDynamicDateRange = `LAST_YEAR`

const DateRangeValueDynamicDateRangeThisMonth DateRangeValueDynamicDateRange = `THIS_MONTH`

const DateRangeValueDynamicDateRangeThisWeek DateRangeValueDynamicDateRange = `THIS_WEEK`

const DateRangeValueDynamicDateRangeThisYear DateRangeValueDynamicDateRange = `THIS_YEAR`

const DateRangeValueDynamicDateRangeToday DateRangeValueDynamicDateRange = `TODAY`

const DateRangeValueDynamicDateRangeYesterday DateRangeValueDynamicDateRange = `YESTERDAY`

// String representation for [fmt.Print]
func (f *DateRangeValueDynamicDateRange) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DateRangeValueDynamicDateRange) Set(v string) error {
	switch v {
	case `LAST_12_MONTHS`, `LAST_14_DAYS`, `LAST_24_HOURS`, `LAST_30_DAYS`, `LAST_60_DAYS`, `LAST_7_DAYS`, `LAST_8_HOURS`, `LAST_90_DAYS`, `LAST_HOUR`, `LAST_MONTH`, `LAST_WEEK`, `LAST_YEAR`, `THIS_MONTH`, `THIS_WEEK`, `THIS_YEAR`, `TODAY`, `YESTERDAY`:
		*f = DateRangeValueDynamicDateRange(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LAST_12_MONTHS", "LAST_14_DAYS", "LAST_24_HOURS", "LAST_30_DAYS", "LAST_60_DAYS", "LAST_7_DAYS", "LAST_8_HOURS", "LAST_90_DAYS", "LAST_HOUR", "LAST_MONTH", "LAST_WEEK", "LAST_YEAR", "THIS_MONTH", "THIS_WEEK", "THIS_YEAR", "TODAY", "YESTERDAY"`, v)
	}
}

// Values returns all possible values for DateRangeValueDynamicDateRange.
//
// There is no guarantee on the order of the values in the slice.
func (f *DateRangeValueDynamicDateRange) Values() []DateRangeValueDynamicDateRange {
	return []DateRangeValueDynamicDateRange{
		DateRangeValueDynamicDateRangeLast12Months,
		DateRangeValueDynamicDateRangeLast14Days,
		DateRangeValueDynamicDateRangeLast24Hours,
		DateRangeValueDynamicDateRangeLast30Days,
		DateRangeValueDynamicDateRangeLast60Days,
		DateRangeValueDynamicDateRangeLast7Days,
		DateRangeValueDynamicDateRangeLast8Hours,
		DateRangeValueDynamicDateRangeLast90Days,
		DateRangeValueDynamicDateRangeLastHour,
		DateRangeValueDynamicDateRangeLastMonth,
		DateRangeValueDynamicDateRangeLastWeek,
		DateRangeValueDynamicDateRangeLastYear,
		DateRangeValueDynamicDateRangeThisMonth,
		DateRangeValueDynamicDateRangeThisWeek,
		DateRangeValueDynamicDateRangeThisYear,
		DateRangeValueDynamicDateRangeToday,
		DateRangeValueDynamicDateRangeYesterday,
	}
}

// Type always returns DateRangeValueDynamicDateRange to satisfy [pflag.Value] interface
func (f *DateRangeValueDynamicDateRange) Type() string {
	return "DateRangeValueDynamicDateRange"
}

type DateValue struct {
	// Manually specified date-time value.
	// Wire name: 'date_value'
	DateValue string `json:"date_value,omitempty"`
	// Dynamic date-time value based on current date-time.
	// Wire name: 'dynamic_date_value'
	DynamicDateValue DateValueDynamicDate `json:"dynamic_date_value,omitempty"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	// Wire name: 'precision'
	Precision DatePrecision `json:"precision,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DateValue) EncodeValues(key string, v *url.Values) error {
	pb, err := dateValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DateValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dateValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dateValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DateValue) MarshalJSON() ([]byte, error) {
	pb, err := dateValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DateValueDynamicDate string

const DateValueDynamicDateNow DateValueDynamicDate = `NOW`

const DateValueDynamicDateYesterday DateValueDynamicDate = `YESTERDAY`

// String representation for [fmt.Print]
func (f *DateValueDynamicDate) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DateValueDynamicDate) Set(v string) error {
	switch v {
	case `NOW`, `YESTERDAY`:
		*f = DateValueDynamicDate(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NOW", "YESTERDAY"`, v)
	}
}

// Values returns all possible values for DateValueDynamicDate.
//
// There is no guarantee on the order of the values in the slice.
func (f *DateValueDynamicDate) Values() []DateValueDynamicDate {
	return []DateValueDynamicDate{
		DateValueDynamicDateNow,
		DateValueDynamicDateYesterday,
	}
}

// Type always returns DateValueDynamicDate to satisfy [pflag.Value] interface
func (f *DateValueDynamicDate) Type() string {
	return "DateValueDynamicDate"
}

type DeleteAlertsLegacyRequest struct {
	AlertId string `json:"-" tf:"-"`
}

func (st DeleteAlertsLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAlertsLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteAlertsLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAlertsLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAlertsLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAlertsLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAlertsLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDashboardRequest struct {
	DashboardId string `json:"-" tf:"-"`
}

func (st DeleteDashboardRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDashboardRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id string `json:"-" tf:"-"`
}

func (st DeleteDashboardWidgetRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDashboardWidgetRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDashboardWidgetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDashboardWidgetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDashboardWidgetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDashboardWidgetRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDashboardWidgetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteQueriesLegacyRequest struct {
	QueryId string `json:"-" tf:"-"`
}

func (st DeleteQueriesLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteQueriesLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteQueriesLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteQueriesLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteQueriesLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteQueriesLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvisualizations/create
	Id string `json:"-" tf:"-"`
}

func (st DeleteQueryVisualizationsLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteQueryVisualizationsLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteQueryVisualizationsLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteQueryVisualizationsLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteQueryVisualizationsLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteQueryVisualizationsLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteQueryVisualizationsLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteVisualizationRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st DeleteVisualizationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteVisualizationRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteVisualizationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteVisualizationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteVisualizationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteVisualizationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteVisualizationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" tf:"-"`
}

func (st DeleteWarehouseRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteWarehouseRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteWarehouseRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWarehouseRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWarehouseRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWarehouseRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteWarehouseRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Disposition string

const DispositionExternalLinks Disposition = `EXTERNAL_LINKS`

const DispositionInline Disposition = `INLINE`

// String representation for [fmt.Print]
func (f *Disposition) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Disposition) Set(v string) error {
	switch v {
	case `EXTERNAL_LINKS`, `INLINE`:
		*f = Disposition(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EXTERNAL_LINKS", "INLINE"`, v)
	}
}

// Values returns all possible values for Disposition.
//
// There is no guarantee on the order of the values in the slice.
func (f *Disposition) Values() []Disposition {
	return []Disposition{
		DispositionExternalLinks,
		DispositionInline,
	}
}

// Type always returns Disposition to satisfy [pflag.Value] interface
func (f *Disposition) Type() string {
	return "Disposition"
}

type EditAlert struct {
	AlertId string `json:"-" tf:"-"`
	// Name of the alert.
	// Wire name: 'name'
	Name string `json:"name"`
	// Alert configuration options.
	// Wire name: 'options'
	Options AlertOptions `json:"options"`
	// Query ID.
	// Wire name: 'query_id'
	QueryId string `json:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm int `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EditAlert) EncodeValues(key string, v *url.Values) error {
	pb, err := editAlertToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EditAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditAlert) MarshalJSON() ([]byte, error) {
	pb, err := editAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute.
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Required. Id of the warehouse to configure.
	Id string `json:"-" tf:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	// Wire name: 'min_num_clusters'
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags `json:"tags,omitempty"`

	// Wire name: 'warehouse_type'
	WarehouseType EditWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EditWarehouseRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := editWarehouseRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EditWarehouseRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editWarehouseRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editWarehouseRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditWarehouseRequest) MarshalJSON() ([]byte, error) {
	pb, err := editWarehouseRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type EditWarehouseRequestWarehouseType string

const EditWarehouseRequestWarehouseTypeClassic EditWarehouseRequestWarehouseType = `CLASSIC`

const EditWarehouseRequestWarehouseTypePro EditWarehouseRequestWarehouseType = `PRO`

const EditWarehouseRequestWarehouseTypeTypeUnspecified EditWarehouseRequestWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *EditWarehouseRequestWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EditWarehouseRequestWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = EditWarehouseRequestWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for EditWarehouseRequestWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EditWarehouseRequestWarehouseType) Values() []EditWarehouseRequestWarehouseType {
	return []EditWarehouseRequestWarehouseType{
		EditWarehouseRequestWarehouseTypeClassic,
		EditWarehouseRequestWarehouseTypePro,
		EditWarehouseRequestWarehouseTypeTypeUnspecified,
	}
}

// Type always returns EditWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *EditWarehouseRequestWarehouseType) Type() string {
	return "EditWarehouseRequestWarehouseType"
}

type EndpointConfPair struct {

	// Wire name: 'key'
	Key string `json:"key,omitempty"`

	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EndpointConfPair) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointConfPairToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointConfPair) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointConfPairPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointConfPairFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointConfPair) MarshalJSON() ([]byte, error) {
	pb, err := endpointConfPairToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	// Wire name: 'details'
	Details string `json:"details,omitempty"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	// Wire name: 'failure_reason'
	FailureReason *TerminationReason `json:"failure_reason,omitempty"`
	// Deprecated. split into summary and details for security
	// Wire name: 'message'
	Message string `json:"message,omitempty"`

	// Wire name: 'status'
	Status Status `json:"status,omitempty"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	// Wire name: 'summary'
	Summary string `json:"summary,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EndpointHealth) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointHealthToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointHealth) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointHealthPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointHealthFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointHealth) MarshalJSON() ([]byte, error) {
	pb, err := endpointHealthToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	// Wire name: 'health'
	Health *EndpointHealth `json:"health,omitempty"`
	// unique identifier for warehouse
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// the jdbc connection string for this warehouse
	// Wire name: 'jdbc_url'
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	// Wire name: 'min_num_clusters'
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Deprecated. current number of active sessions for the warehouse
	// Wire name: 'num_active_sessions'
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// current number of clusters running for the service
	// Wire name: 'num_clusters'
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the SQL warehouse
	// Wire name: 'odbc_params'
	OdbcParams *OdbcParams `json:"odbc_params,omitempty"`

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`

	// Wire name: 'state'
	State State `json:"state,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags `json:"tags,omitempty"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	// Wire name: 'warehouse_type'
	WarehouseType EndpointInfoWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EndpointInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointInfo) MarshalJSON() ([]byte, error) {
	pb, err := endpointInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type EndpointInfoWarehouseType string

const EndpointInfoWarehouseTypeClassic EndpointInfoWarehouseType = `CLASSIC`

const EndpointInfoWarehouseTypePro EndpointInfoWarehouseType = `PRO`

const EndpointInfoWarehouseTypeTypeUnspecified EndpointInfoWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *EndpointInfoWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointInfoWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = EndpointInfoWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for EndpointInfoWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointInfoWarehouseType) Values() []EndpointInfoWarehouseType {
	return []EndpointInfoWarehouseType{
		EndpointInfoWarehouseTypeClassic,
		EndpointInfoWarehouseTypePro,
		EndpointInfoWarehouseTypeTypeUnspecified,
	}
}

// Type always returns EndpointInfoWarehouseType to satisfy [pflag.Value] interface
func (f *EndpointInfoWarehouseType) Type() string {
	return "EndpointInfoWarehouseType"
}

type EndpointTagPair struct {

	// Wire name: 'key'
	Key string `json:"key,omitempty"`

	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EndpointTagPair) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointTagPairToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointTagPair) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointTagPairPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointTagPairFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointTagPair) MarshalJSON() ([]byte, error) {
	pb, err := endpointTagPairToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointTags struct {

	// Wire name: 'custom_tags'
	CustomTags []EndpointTagPair `json:"custom_tags,omitempty"`
}

func (st EndpointTags) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointTagsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointTags) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointTagsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointTagsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointTags) MarshalJSON() ([]byte, error) {
	pb, err := endpointTagsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnumValue struct {
	// List of valid query parameter values, newline delimited.
	// Wire name: 'enum_options'
	EnumOptions string `json:"enum_options,omitempty"`
	// If specified, allows multiple values to be selected for this parameter.
	// Wire name: 'multi_values_options'
	MultiValuesOptions *MultiValuesOptions `json:"multi_values_options,omitempty"`
	// List of selected query parameter values.
	// Wire name: 'values'
	Values []string `json:"values,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EnumValue) EncodeValues(key string, v *url.Values) error {
	pb, err := enumValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EnumValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enumValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enumValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnumValue) MarshalJSON() ([]byte, error) {
	pb, err := enumValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	// Wire name: 'byte_limit'
	ByteLimit int64 `json:"byte_limit,omitempty"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`

	// Wire name: 'disposition'
	Disposition Disposition `json:"disposition,omitempty"`
	// Statement execution supports three result formats: `JSON_ARRAY`
	// (default), `ARROW_STREAM`, and `CSV`.
	//
	// Important: The formats `ARROW_STREAM` and `CSV` are supported only with
	// `EXTERNAL_LINKS` disposition. `JSON_ARRAY` is supported in `INLINE` and
	// `EXTERNAL_LINKS` disposition.
	//
	// When specifying `format=JSON_ARRAY`, result data will be formatted as an
	// array of arrays of values, where each value is either the *string
	// representation* of a value, or `null`. For example, the output of `SELECT
	// concat('id-', id) AS strCol, id AS intCol, null AS nullCol FROM range(3)`
	// would look like this:
	//
	// ``` [ [ "id-1", "1", null ], [ "id-2", "2", null ], [ "id-3", "3", null
	// ], ] ```
	//
	// When specifying `format=JSON_ARRAY` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result contains compact JSON with no indentation or
	// extra whitespace.
	//
	// When specifying `format=ARROW_STREAM` and `disposition=EXTERNAL_LINKS`,
	// each chunk in the result will be formatted as Apache Arrow Stream. See
	// the [Apache Arrow streaming format].
	//
	// When specifying `format=CSV` and `disposition=EXTERNAL_LINKS`, each chunk
	// in the result will be a CSV according to [RFC 4180] standard. All the
	// columns values will have *string representation* similar to the
	// `JSON_ARRAY` format, and `null` values will be encoded as “null”.
	// Only the first chunk in the result would contain a header row with column
	// names. For example, the output of `SELECT concat('id-', id) AS strCol, id
	// AS intCol, null as nullCol FROM range(3)` would look like this:
	//
	// ``` strCol,intCol,nullCol id-1,1,null id-2,2,null id-3,3,null ```
	//
	// [Apache Arrow streaming format]: https://arrow.apache.org/docs/format/Columnar.html#ipc-streaming-format
	// [RFC 4180]: https://www.rfc-editor.org/rfc/rfc4180
	// Wire name: 'format'
	Format Format `json:"format,omitempty"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	// Wire name: 'on_wait_timeout'
	OnWaitTimeout ExecuteStatementRequestOnWaitTimeout `json:"on_wait_timeout,omitempty"`
	// A list of parameters to pass into a SQL statement containing parameter
	// markers. A parameter consists of a name, a value, and optionally a type.
	// To represent a NULL value, the `value` field may be omitted or set to
	// `null` explicitly. If the `type` field is omitted, the value is
	// interpreted as a string.
	//
	// If the type is given, parameters will be checked for type correctness
	// according to the given type. A value is correct if the provided string
	// can be converted to the requested type using the `cast` function. The
	// exact semantics are described in the section [`cast` function] of the SQL
	// language reference.
	//
	// For example, the following statement contains two parameters, `my_name`
	// and `my_date`:
	//
	// SELECT * FROM my_table WHERE name = :my_name AND date = :my_date
	//
	// The parameters can be passed in the request body as follows:
	//
	// { ..., "statement": "SELECT * FROM my_table WHERE name = :my_name AND
	// date = :my_date", "parameters": [ { "name": "my_name", "value": "the
	// name" }, { "name": "my_date", "value": "2020-01-01", "type": "DATE" } ] }
	//
	// Currently, positional parameters denoted by a `?` marker are not
	// supported by the Databricks SQL Statement Execution API.
	//
	// Also see the section [Parameter markers] of the SQL language reference.
	//
	// [Parameter markers]: https://docs.databricks.com/sql/language-manual/sql-ref-parameter-marker.html
	// [`cast` function]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	// Wire name: 'parameters'
	Parameters []StatementParameterListItem `json:"parameters,omitempty"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	// Wire name: 'row_limit'
	RowLimit int64 `json:"row_limit,omitempty"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`. The maximum query text size is 16 MiB.
	// Wire name: 'statement'
	Statement string `json:"statement"`
	// The time in seconds the call will wait for the statement's result set as
	// `Ns`, where `N` can be set to 0 or to a value between 5 and 50.
	//
	// When set to `0s`, the statement will execute in asynchronous mode and the
	// call will not wait for the execution to finish. In this case, the call
	// returns directly with `PENDING` state and a statement ID which can be
	// used for polling with :method:statementexecution/getStatement.
	//
	// When set between 5 and 50 seconds, the call will behave synchronously up
	// to this timeout and wait for the statement execution to finish. If the
	// execution finishes within this time, the call returns immediately with a
	// manifest and result data (or a `FAILED` state in case of an execution
	// error). If the statement takes longer to execute, `on_wait_timeout`
	// determines what should happen after the timeout is reached.
	// Wire name: 'wait_timeout'
	WaitTimeout string `json:"wait_timeout,omitempty"`
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExecuteStatementRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := executeStatementRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExecuteStatementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &executeStatementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := executeStatementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExecuteStatementRequest) MarshalJSON() ([]byte, error) {
	pb, err := executeStatementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// When `wait_timeout > 0s`, the call will block up to the specified time. If
// the statement execution doesn't finish within this time, `on_wait_timeout`
// determines whether the execution should continue or be canceled. When set to
// `CONTINUE`, the statement execution continues asynchronously and the call
// returns a statement ID which can be used for polling with
// :method:statementexecution/getStatement. When set to `CANCEL`, the statement
// execution is canceled and the call returns with a `CANCELED` state.
type ExecuteStatementRequestOnWaitTimeout string

const ExecuteStatementRequestOnWaitTimeoutCancel ExecuteStatementRequestOnWaitTimeout = `CANCEL`

const ExecuteStatementRequestOnWaitTimeoutContinue ExecuteStatementRequestOnWaitTimeout = `CONTINUE`

// String representation for [fmt.Print]
func (f *ExecuteStatementRequestOnWaitTimeout) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExecuteStatementRequestOnWaitTimeout) Set(v string) error {
	switch v {
	case `CANCEL`, `CONTINUE`:
		*f = ExecuteStatementRequestOnWaitTimeout(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCEL", "CONTINUE"`, v)
	}
}

// Values returns all possible values for ExecuteStatementRequestOnWaitTimeout.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExecuteStatementRequestOnWaitTimeout) Values() []ExecuteStatementRequestOnWaitTimeout {
	return []ExecuteStatementRequestOnWaitTimeout{
		ExecuteStatementRequestOnWaitTimeoutCancel,
		ExecuteStatementRequestOnWaitTimeoutContinue,
	}
}

// Type always returns ExecuteStatementRequestOnWaitTimeout to satisfy [pflag.Value] interface
func (f *ExecuteStatementRequestOnWaitTimeout) Type() string {
	return "ExecuteStatementRequestOnWaitTimeout"
}

type ExternalLink struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int `json:"chunk_index,omitempty"`
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	// Wire name: 'expiration'
	Expiration string `json:"expiration,omitempty"`

	// Wire name: 'external_link'
	ExternalLink string `json:"external_link,omitempty"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	// Wire name: 'http_headers'
	HttpHeaders map[string]string `json:"http_headers,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	// Wire name: 'next_chunk_index'
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	// Wire name: 'next_chunk_internal_link'
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalLink) EncodeValues(key string, v *url.Values) error {
	pb, err := externalLinkToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalLink) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalLinkPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalLinkFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalLink) MarshalJSON() ([]byte, error) {
	pb, err := externalLinkToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalQuerySource struct {
	// The canonical identifier for this SQL alert
	// Wire name: 'alert_id'
	AlertId string `json:"alert_id,omitempty"`
	// The canonical identifier for this Lakeview dashboard
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The canonical identifier for this Genie space
	// Wire name: 'genie_space_id'
	GenieSpaceId string `json:"genie_space_id,omitempty"`

	// Wire name: 'job_info'
	JobInfo *ExternalQuerySourceJobInfo `json:"job_info,omitempty"`
	// The canonical identifier for this legacy dashboard
	// Wire name: 'legacy_dashboard_id'
	LegacyDashboardId string `json:"legacy_dashboard_id,omitempty"`
	// The canonical identifier for this notebook
	// Wire name: 'notebook_id'
	NotebookId string `json:"notebook_id,omitempty"`
	// The canonical identifier for this SQL query
	// Wire name: 'sql_query_id'
	SqlQueryId string `json:"sql_query_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalQuerySource) EncodeValues(key string, v *url.Values) error {
	pb, err := externalQuerySourceToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalQuerySource) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalQuerySourcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalQuerySourceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalQuerySource) MarshalJSON() ([]byte, error) {
	pb, err := externalQuerySourceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalQuerySourceJobInfo struct {
	// The canonical identifier for this job.
	// Wire name: 'job_id'
	JobId string `json:"job_id,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	// Wire name: 'job_run_id'
	JobRunId string `json:"job_run_id,omitempty"`
	// The canonical identifier of the task run.
	// Wire name: 'job_task_run_id'
	JobTaskRunId string `json:"job_task_run_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalQuerySourceJobInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := externalQuerySourceJobInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalQuerySourceJobInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalQuerySourceJobInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalQuerySourceJobInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalQuerySourceJobInfo) MarshalJSON() ([]byte, error) {
	pb, err := externalQuerySourceJobInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Format string

const FormatArrowStream Format = `ARROW_STREAM`

const FormatCsv Format = `CSV`

const FormatJsonArray Format = `JSON_ARRAY`

// String representation for [fmt.Print]
func (f *Format) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Format) Set(v string) error {
	switch v {
	case `ARROW_STREAM`, `CSV`, `JSON_ARRAY`:
		*f = Format(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARROW_STREAM", "CSV", "JSON_ARRAY"`, v)
	}
}

// Values returns all possible values for Format.
//
// There is no guarantee on the order of the values in the slice.
func (f *Format) Values() []Format {
	return []Format{
		FormatArrowStream,
		FormatCsv,
		FormatJsonArray,
	}
}

// Type always returns Format to satisfy [pflag.Value] interface
func (f *Format) Type() string {
	return "Format"
}

type GetAlertRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st GetAlertRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAlertRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetAlertRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAlertRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAlertRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAlertRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAlertRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetAlertV2Request struct {
	Id string `json:"-" tf:"-"`
}

func (st GetAlertV2Request) EncodeValues(key string, v *url.Values) error {
	pb, err := getAlertV2RequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetAlertV2Request) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAlertV2RequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAlertV2RequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAlertV2Request) MarshalJSON() ([]byte, error) {
	pb, err := getAlertV2RequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetAlertsLegacyRequest struct {
	AlertId string `json:"-" tf:"-"`
}

func (st GetAlertsLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAlertsLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetAlertsLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAlertsLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAlertsLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAlertsLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAlertsLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetDashboardRequest struct {
	DashboardId string `json:"-" tf:"-"`
}

func (st GetDashboardRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDashboardRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId string `json:"-" tf:"-"`
	// The type of object permissions to check.
	ObjectType ObjectTypePlural `json:"-" tf:"-"`
}

func (st GetDbsqlPermissionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDbsqlPermissionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetDbsqlPermissionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDbsqlPermissionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDbsqlPermissionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDbsqlPermissionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDbsqlPermissionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetQueriesLegacyRequest struct {
	QueryId string `json:"-" tf:"-"`
}

func (st GetQueriesLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getQueriesLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getQueriesLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getQueriesLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetQueriesLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getQueriesLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetQueryRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st GetQueryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getQueryRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := getQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetResponse struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	// An object's type and UUID, separated by a forward slash (/) character.
	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`
	// A singular noun object type.
	// Wire name: 'object_type'
	ObjectType ObjectType `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetResponse) MarshalJSON() ([]byte, error) {
	pb, err := getResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" tf:"-"`
}

func (st GetStatementRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getStatementRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetStatementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStatementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStatementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStatementRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStatementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetStatementResultChunkNRequest struct {
	ChunkIndex int `json:"-" tf:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" tf:"-"`
}

func (st GetStatementResultChunkNRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getStatementResultChunkNRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetStatementResultChunkNRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStatementResultChunkNRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStatementResultChunkNRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStatementResultChunkNRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStatementResultChunkNRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" tf:"-"`
}

func (st GetWarehousePermissionLevelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWarehousePermissionLevelsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWarehousePermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWarehousePermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWarehousePermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWarehousePermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWarehousePermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWarehousePermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []WarehousePermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetWarehousePermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getWarehousePermissionLevelsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWarehousePermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWarehousePermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWarehousePermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWarehousePermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getWarehousePermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" tf:"-"`
}

func (st GetWarehousePermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWarehousePermissionsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWarehousePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWarehousePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWarehousePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWarehousePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWarehousePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" tf:"-"`
}

func (st GetWarehouseRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWarehouseRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWarehouseRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWarehouseRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWarehouseRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWarehouseRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWarehouseRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	// Wire name: 'health'
	Health *EndpointHealth `json:"health,omitempty"`
	// unique identifier for warehouse
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// the jdbc connection string for this warehouse
	// Wire name: 'jdbc_url'
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int `json:"max_num_clusters,omitempty"`
	// Minimum number of available clusters that will be maintained for this SQL
	// warehouse. Increasing this will ensure that a larger number of clusters
	// are always running and therefore may reduce the cold start time for new
	// queries. This is similar to reserved vs. revocable cores in a resource
	// manager.
	//
	// Supported values: - Must be > 0 - Must be <= min(max_num_clusters, 30)
	//
	// Defaults to 1
	// Wire name: 'min_num_clusters'
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Deprecated. current number of active sessions for the warehouse
	// Wire name: 'num_active_sessions'
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// current number of clusters running for the service
	// Wire name: 'num_clusters'
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the SQL warehouse
	// Wire name: 'odbc_params'
	OdbcParams *OdbcParams `json:"odbc_params,omitempty"`

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`

	// Wire name: 'state'
	State State `json:"state,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags `json:"tags,omitempty"`

	// Wire name: 'warehouse_type'
	WarehouseType GetWarehouseResponseWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetWarehouseResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getWarehouseResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWarehouseResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWarehouseResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWarehouseResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWarehouseResponse) MarshalJSON() ([]byte, error) {
	pb, err := getWarehouseResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type GetWarehouseResponseWarehouseType string

const GetWarehouseResponseWarehouseTypeClassic GetWarehouseResponseWarehouseType = `CLASSIC`

const GetWarehouseResponseWarehouseTypePro GetWarehouseResponseWarehouseType = `PRO`

const GetWarehouseResponseWarehouseTypeTypeUnspecified GetWarehouseResponseWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *GetWarehouseResponseWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetWarehouseResponseWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = GetWarehouseResponseWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for GetWarehouseResponseWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetWarehouseResponseWarehouseType) Values() []GetWarehouseResponseWarehouseType {
	return []GetWarehouseResponseWarehouseType{
		GetWarehouseResponseWarehouseTypeClassic,
		GetWarehouseResponseWarehouseTypePro,
		GetWarehouseResponseWarehouseTypeTypeUnspecified,
	}
}

// Type always returns GetWarehouseResponseWarehouseType to satisfy [pflag.Value] interface
func (f *GetWarehouseResponseWarehouseType) Type() string {
	return "GetWarehouseResponseWarehouseType"
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	// Wire name: 'channel'
	Channel *Channel `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'config_param'
	ConfigParam *RepeatedEndpointConfPairs `json:"config_param,omitempty"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	// Wire name: 'data_access_config'
	DataAccessConfig []EndpointConfPair `json:"data_access_config,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	// Wire name: 'enabled_warehouse_types'
	EnabledWarehouseTypes []WarehouseTypePair `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'global_param'
	GlobalParam *RepeatedEndpointConfPairs `json:"global_param,omitempty"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	// Wire name: 'google_service_account'
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for warehouses
	// Wire name: 'security_policy'
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy `json:"security_policy,omitempty"`
	// SQL configuration parameters
	// Wire name: 'sql_configuration_parameters'
	SqlConfigurationParameters *RepeatedEndpointConfPairs `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetWorkspaceWarehouseConfigResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceWarehouseConfigResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetWorkspaceWarehouseConfigResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceWarehouseConfigResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceWarehouseConfigResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceWarehouseConfigResponse) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceWarehouseConfigResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Security policy for warehouses
type GetWorkspaceWarehouseConfigResponseSecurityPolicy string

const GetWorkspaceWarehouseConfigResponseSecurityPolicyDataAccessControl GetWorkspaceWarehouseConfigResponseSecurityPolicy = `DATA_ACCESS_CONTROL`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyNone GetWorkspaceWarehouseConfigResponseSecurityPolicy = `NONE`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyPassthrough GetWorkspaceWarehouseConfigResponseSecurityPolicy = `PASSTHROUGH`

// String representation for [fmt.Print]
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Set(v string) error {
	switch v {
	case `DATA_ACCESS_CONTROL`, `NONE`, `PASSTHROUGH`:
		*f = GetWorkspaceWarehouseConfigResponseSecurityPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_ACCESS_CONTROL", "NONE", "PASSTHROUGH"`, v)
	}
}

// Values returns all possible values for GetWorkspaceWarehouseConfigResponseSecurityPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Values() []GetWorkspaceWarehouseConfigResponseSecurityPolicy {
	return []GetWorkspaceWarehouseConfigResponseSecurityPolicy{
		GetWorkspaceWarehouseConfigResponseSecurityPolicyDataAccessControl,
		GetWorkspaceWarehouseConfigResponseSecurityPolicyNone,
		GetWorkspaceWarehouseConfigResponseSecurityPolicyPassthrough,
	}
}

// Type always returns GetWorkspaceWarehouseConfigResponseSecurityPolicy to satisfy [pflag.Value] interface
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Type() string {
	return "GetWorkspaceWarehouseConfigResponseSecurityPolicy"
}

type LegacyAlert struct {
	// Timestamp when the alert was created.
	// Wire name: 'created_at'
	CreatedAt string `json:"created_at,omitempty"`
	// Alert ID.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Timestamp when the alert was last triggered.
	// Wire name: 'last_triggered_at'
	LastTriggeredAt string `json:"last_triggered_at,omitempty"`
	// Name of the alert.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Alert configuration options.
	// Wire name: 'options'
	Options *AlertOptions `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string `json:"parent,omitempty"`

	// Wire name: 'query'
	Query *AlertQuery `json:"query,omitempty"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm int `json:"rearm,omitempty"`
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	// Wire name: 'state'
	State LegacyAlertState `json:"state,omitempty"`
	// Timestamp when the alert was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string `json:"updated_at,omitempty"`

	// Wire name: 'user'
	User *User `json:"user,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LegacyAlert) EncodeValues(key string, v *url.Values) error {
	pb, err := legacyAlertToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *LegacyAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &legacyAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := legacyAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LegacyAlert) MarshalJSON() ([]byte, error) {
	pb, err := legacyAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LegacyAlertState string

const LegacyAlertStateOk LegacyAlertState = `ok`

const LegacyAlertStateTriggered LegacyAlertState = `triggered`

const LegacyAlertStateUnknown LegacyAlertState = `unknown`

// String representation for [fmt.Print]
func (f *LegacyAlertState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LegacyAlertState) Set(v string) error {
	switch v {
	case `ok`, `triggered`, `unknown`:
		*f = LegacyAlertState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ok", "triggered", "unknown"`, v)
	}
}

// Values returns all possible values for LegacyAlertState.
//
// There is no guarantee on the order of the values in the slice.
func (f *LegacyAlertState) Values() []LegacyAlertState {
	return []LegacyAlertState{
		LegacyAlertStateOk,
		LegacyAlertStateTriggered,
		LegacyAlertStateUnknown,
	}
}

// Type always returns LegacyAlertState to satisfy [pflag.Value] interface
func (f *LegacyAlertState) Type() string {
	return "LegacyAlertState"
}

type LegacyQuery struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	// Wire name: 'can_edit'
	CanEdit bool `json:"can_edit,omitempty"`
	// The timestamp when this query was created.
	// Wire name: 'created_at'
	CreatedAt string `json:"created_at,omitempty"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Query ID.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	// Wire name: 'is_draft'
	IsDraft bool `json:"is_draft,omitempty"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	// Wire name: 'is_favorite'
	IsFavorite bool `json:"is_favorite,omitempty"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	// Wire name: 'is_safe'
	IsSafe bool `json:"is_safe,omitempty"`

	// Wire name: 'last_modified_by'
	LastModifiedBy *User `json:"last_modified_by,omitempty"`
	// The ID of the user who last saved changes to this query.
	// Wire name: 'last_modified_by_id'
	LastModifiedById int `json:"last_modified_by_id,omitempty"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	// Wire name: 'latest_query_data_id'
	LatestQueryDataId string `json:"latest_query_data_id,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'options'
	Options *QueryOptions `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string `json:"parent,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_tier'
	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// The text of the query to be run.
	// Wire name: 'query'
	Query string `json:"query,omitempty"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	// Wire name: 'query_hash'
	QueryHash string `json:"query_hash,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`
	// The timestamp at which this query was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string `json:"updated_at,omitempty"`

	// Wire name: 'user'
	User *User `json:"user,omitempty"`
	// The ID of the user who owns the query.
	// Wire name: 'user_id'
	UserId int `json:"user_id,omitempty"`

	// Wire name: 'visualizations'
	Visualizations []LegacyVisualization `json:"visualizations,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LegacyQuery) EncodeValues(key string, v *url.Values) error {
	pb, err := legacyQueryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *LegacyQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &legacyQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := legacyQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LegacyQuery) MarshalJSON() ([]byte, error) {
	pb, err := legacyQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization struct {

	// Wire name: 'created_at'
	CreatedAt string `json:"created_at,omitempty"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// The UUID for this visualization.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	// Wire name: 'options'
	Options any `json:"options,omitempty"`

	// Wire name: 'query'
	Query *LegacyQuery `json:"query,omitempty"`
	// The type of visualization: chart, table, pivot table, and so on.
	// Wire name: 'type'
	Type string `json:"type,omitempty"`

	// Wire name: 'updated_at'
	UpdatedAt string `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LegacyVisualization) EncodeValues(key string, v *url.Values) error {
	pb, err := legacyVisualizationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *LegacyVisualization) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &legacyVisualizationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := legacyVisualizationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LegacyVisualization) MarshalJSON() ([]byte, error) {
	pb, err := legacyVisualizationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LifecycleState string

const LifecycleStateActive LifecycleState = `ACTIVE`

const LifecycleStateTrashed LifecycleState = `TRASHED`

// String representation for [fmt.Print]
func (f *LifecycleState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LifecycleState) Set(v string) error {
	switch v {
	case `ACTIVE`, `TRASHED`:
		*f = LifecycleState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "TRASHED"`, v)
	}
}

// Values returns all possible values for LifecycleState.
//
// There is no guarantee on the order of the values in the slice.
func (f *LifecycleState) Values() []LifecycleState {
	return []LifecycleState{
		LifecycleStateActive,
		LifecycleStateTrashed,
	}
}

// Type always returns LifecycleState to satisfy [pflag.Value] interface
func (f *LifecycleState) Type() string {
	return "LifecycleState"
}

type ListAlertsRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAlertsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listAlertsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListAlertsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAlertsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAlertsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAlertsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAlertsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAlertsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'results'
	Results []ListAlertsResponseAlert `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAlertsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listAlertsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListAlertsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAlertsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAlertsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAlertsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listAlertsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAlertsResponseAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition `json:"condition,omitempty"`
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The workspace state of the alert. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	// Wire name: 'state'
	State AlertState `json:"state,omitempty"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	// Wire name: 'trigger_time'
	TriggerTime string `json:"trigger_time,omitempty"`
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAlertsResponseAlert) EncodeValues(key string, v *url.Values) error {
	pb, err := listAlertsResponseAlertToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListAlertsResponseAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAlertsResponseAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAlertsResponseAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAlertsResponseAlert) MarshalJSON() ([]byte, error) {
	pb, err := listAlertsResponseAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAlertsV2Request struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAlertsV2Request) EncodeValues(key string, v *url.Values) error {
	pb, err := listAlertsV2RequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListAlertsV2Request) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAlertsV2RequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAlertsV2RequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAlertsV2Request) MarshalJSON() ([]byte, error) {
	pb, err := listAlertsV2RequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListAlertsV2Response struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'results'
	Results []AlertV2 `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAlertsV2Response) EncodeValues(key string, v *url.Values) error {
	pb, err := listAlertsV2ResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListAlertsV2Response) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAlertsV2ResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAlertsV2ResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAlertsV2Response) MarshalJSON() ([]byte, error) {
	pb, err := listAlertsV2ResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	Order ListOrder `json:"-" tf:"-"`
	// Page number to retrieve.
	Page int `json:"-" tf:"-"`
	// Number of dashboards to return per page.
	PageSize int `json:"-" tf:"-"`
	// Full text search term.
	Q string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListDashboardsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listDashboardsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListDashboardsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDashboardsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDashboardsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDashboardsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listDashboardsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListOrder string

const ListOrderCreatedAt ListOrder = `created_at`

const ListOrderName ListOrder = `name`

// String representation for [fmt.Print]
func (f *ListOrder) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListOrder) Set(v string) error {
	switch v {
	case `created_at`, `name`:
		*f = ListOrder(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "created_at", "name"`, v)
	}
}

// Values returns all possible values for ListOrder.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListOrder) Values() []ListOrder {
	return []ListOrder{
		ListOrderCreatedAt,
		ListOrderName,
	}
}

// Type always returns ListOrder to satisfy [pflag.Value] interface
func (f *ListOrder) Type() string {
	return "ListOrder"
}

type ListQueriesLegacyRequest struct {
	// Name of query attribute to order by. Default sort order is ascending.
	// Append a dash (`-`) to order descending instead.
	//
	// - `name`: The name of the query.
	//
	// - `created_at`: The timestamp the query was created.
	//
	// - `runtime`: The time it took to run this query. This is blank for
	// parameterized queries. A blank value is treated as the highest value for
	// sorting.
	//
	// - `executed_at`: The timestamp when the query was last run.
	//
	// - `created_by`: The user name of the user that created the query.
	Order string `json:"-" tf:"-"`
	// Page number to retrieve.
	Page int `json:"-" tf:"-"`
	// Number of queries to return per page.
	PageSize int `json:"-" tf:"-"`
	// Full text search term
	Q string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQueriesLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listQueriesLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQueriesLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQueriesLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQueriesLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := listQueriesLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListQueriesRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQueriesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listQueriesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListQueriesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQueriesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQueriesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQueriesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listQueriesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	// Wire name: 'has_next_page'
	HasNextPage bool `json:"has_next_page,omitempty"`
	// A token that can be used to get the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'res'
	Res []QueryInfo `json:"res,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQueriesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listQueriesResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListQueriesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQueriesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQueriesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQueriesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listQueriesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListQueryHistoryRequest struct {
	// An optional filter object to limit query history results. Accepts
	// parameters such as user IDs, endpoint IDs, and statuses to narrow the
	// returned data. In a URL, the parameters of this filter are specified with
	// dot notation. For example: `filter_by.statement_ids`.
	FilterBy *QueryFilter `json:"-" tf:"-"`
	// Whether to include the query metrics with each query. Only use this for a
	// small subset of queries (max_results). Defaults to false.
	IncludeMetrics bool `json:"-" tf:"-"`
	// Limit the number of results returned in one page. Must be less than 1000
	// and the default is 100.
	MaxResults int `json:"-" tf:"-"`
	// A token that can be used to get the next page of results. The token can
	// contains characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B. This field is
	// optional.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQueryHistoryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listQueryHistoryRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListQueryHistoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQueryHistoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQueryHistoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQueryHistoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := listQueryHistoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListQueryObjectsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'results'
	Results []ListQueryObjectsResponseQuery `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQueryObjectsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listQueryObjectsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListQueryObjectsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQueryObjectsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQueryObjectsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQueryObjectsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listQueryObjectsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListQueryObjectsResponseQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// Timestamp when this query was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the query.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Username of the user who last saved changes to this query.
	// Wire name: 'last_modifier_user_name'
	LastModifierUserName string `json:"last_modifier_user_name,omitempty"`
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`
	// Timestamp when this query was last updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListQueryObjectsResponseQuery) EncodeValues(key string, v *url.Values) error {
	pb, err := listQueryObjectsResponseQueryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListQueryObjectsResponseQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQueryObjectsResponseQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQueryObjectsResponseQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQueryObjectsResponseQuery) MarshalJSON() ([]byte, error) {
	pb, err := listQueryObjectsResponseQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListResponse struct {
	// The total number of dashboards.
	// Wire name: 'count'
	Count int `json:"count,omitempty"`
	// The current page being displayed.
	// Wire name: 'page'
	Page int `json:"page,omitempty"`
	// The number of dashboards per page.
	// Wire name: 'page_size'
	PageSize int `json:"page_size,omitempty"`
	// List of dashboards returned.
	// Wire name: 'results'
	Results []Dashboard `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListResponse) MarshalJSON() ([]byte, error) {
	pb, err := listResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListVisualizationsForQueryRequest struct {
	Id string `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListVisualizationsForQueryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listVisualizationsForQueryRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListVisualizationsForQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listVisualizationsForQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listVisualizationsForQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListVisualizationsForQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := listVisualizationsForQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListVisualizationsForQueryResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'results'
	Results []Visualization `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListVisualizationsForQueryResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listVisualizationsForQueryResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListVisualizationsForQueryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listVisualizationsForQueryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listVisualizationsForQueryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListVisualizationsForQueryResponse) MarshalJSON() ([]byte, error) {
	pb, err := listVisualizationsForQueryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListWarehousesRequest struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	RunAsUserId int `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListWarehousesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listWarehousesRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListWarehousesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listWarehousesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listWarehousesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListWarehousesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listWarehousesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListWarehousesResponse struct {
	// A list of warehouses and their configurations.
	// Wire name: 'warehouses'
	Warehouses []EndpointInfo `json:"warehouses,omitempty"`
}

func (st ListWarehousesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listWarehousesResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListWarehousesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listWarehousesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listWarehousesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListWarehousesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listWarehousesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MultiValuesOptions struct {
	// Character that prefixes each selected parameter value.
	// Wire name: 'prefix'
	Prefix string `json:"prefix,omitempty"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	// Wire name: 'separator'
	Separator string `json:"separator,omitempty"`
	// Character that suffixes each selected parameter value.
	// Wire name: 'suffix'
	Suffix string `json:"suffix,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MultiValuesOptions) EncodeValues(key string, v *url.Values) error {
	pb, err := multiValuesOptionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *MultiValuesOptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &multiValuesOptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := multiValuesOptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MultiValuesOptions) MarshalJSON() ([]byte, error) {
	pb, err := multiValuesOptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NumericValue struct {

	// Wire name: 'value'
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NumericValue) EncodeValues(key string, v *url.Values) error {
	pb, err := numericValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *NumericValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &numericValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := numericValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NumericValue) MarshalJSON() ([]byte, error) {
	pb, err := numericValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A singular noun object type.
type ObjectType string

const ObjectTypeAlert ObjectType = `alert`

const ObjectTypeDashboard ObjectType = `dashboard`

const ObjectTypeDataSource ObjectType = `data_source`

const ObjectTypeQuery ObjectType = `query`

// String representation for [fmt.Print]
func (f *ObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ObjectType) Set(v string) error {
	switch v {
	case `alert`, `dashboard`, `data_source`, `query`:
		*f = ObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "alert", "dashboard", "data_source", "query"`, v)
	}
}

// Values returns all possible values for ObjectType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ObjectType) Values() []ObjectType {
	return []ObjectType{
		ObjectTypeAlert,
		ObjectTypeDashboard,
		ObjectTypeDataSource,
		ObjectTypeQuery,
	}
}

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (f *ObjectType) Type() string {
	return "ObjectType"
}

// Always a plural of the object type.
type ObjectTypePlural string

const ObjectTypePluralAlerts ObjectTypePlural = `alerts`

const ObjectTypePluralDashboards ObjectTypePlural = `dashboards`

const ObjectTypePluralDataSources ObjectTypePlural = `data_sources`

const ObjectTypePluralQueries ObjectTypePlural = `queries`

// String representation for [fmt.Print]
func (f *ObjectTypePlural) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ObjectTypePlural) Set(v string) error {
	switch v {
	case `alerts`, `dashboards`, `data_sources`, `queries`:
		*f = ObjectTypePlural(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "alerts", "dashboards", "data_sources", "queries"`, v)
	}
}

// Values returns all possible values for ObjectTypePlural.
//
// There is no guarantee on the order of the values in the slice.
func (f *ObjectTypePlural) Values() []ObjectTypePlural {
	return []ObjectTypePlural{
		ObjectTypePluralAlerts,
		ObjectTypePluralDashboards,
		ObjectTypePluralDataSources,
		ObjectTypePluralQueries,
	}
}

// Type always returns ObjectTypePlural to satisfy [pflag.Value] interface
func (f *ObjectTypePlural) Type() string {
	return "ObjectTypePlural"
}

type OdbcParams struct {

	// Wire name: 'hostname'
	Hostname string `json:"hostname,omitempty"`

	// Wire name: 'path'
	Path string `json:"path,omitempty"`

	// Wire name: 'port'
	Port int `json:"port,omitempty"`

	// Wire name: 'protocol'
	Protocol string `json:"protocol,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OdbcParams) EncodeValues(key string, v *url.Values) error {
	pb, err := odbcParamsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *OdbcParams) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &odbcParamsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := odbcParamsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OdbcParams) MarshalJSON() ([]byte, error) {
	pb, err := odbcParamsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type OwnableObjectType string

const OwnableObjectTypeAlert OwnableObjectType = `alert`

const OwnableObjectTypeDashboard OwnableObjectType = `dashboard`

const OwnableObjectTypeQuery OwnableObjectType = `query`

// String representation for [fmt.Print]
func (f *OwnableObjectType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OwnableObjectType) Set(v string) error {
	switch v {
	case `alert`, `dashboard`, `query`:
		*f = OwnableObjectType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "alert", "dashboard", "query"`, v)
	}
}

// Values returns all possible values for OwnableObjectType.
//
// There is no guarantee on the order of the values in the slice.
func (f *OwnableObjectType) Values() []OwnableObjectType {
	return []OwnableObjectType{
		OwnableObjectTypeAlert,
		OwnableObjectTypeDashboard,
		OwnableObjectTypeQuery,
	}
}

// Type always returns OwnableObjectType to satisfy [pflag.Value] interface
func (f *OwnableObjectType) Type() string {
	return "OwnableObjectType"
}

type Parameter struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	// Wire name: 'enumOptions'
	EnumOptions string `json:"enumOptions,omitempty"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	// Wire name: 'multiValuesOptions'
	MultiValuesOptions *MultiValuesOptions `json:"multiValuesOptions,omitempty"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	// Wire name: 'queryId'
	QueryId string `json:"queryId,omitempty"`
	// The text displayed in a parameter picking widget.
	// Wire name: 'title'
	Title string `json:"title,omitempty"`
	// Parameters can have several different types.
	// Wire name: 'type'
	Type ParameterType `json:"type,omitempty"`
	// The default value for this parameter.
	// Wire name: 'value'
	Value any `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Parameter) EncodeValues(key string, v *url.Values) error {
	pb, err := parameterToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Parameter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &parameterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := parameterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Parameter) MarshalJSON() ([]byte, error) {
	pb, err := parameterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ParameterType string

const ParameterTypeDatetime ParameterType = `datetime`

const ParameterTypeEnum ParameterType = `enum`

const ParameterTypeNumber ParameterType = `number`

const ParameterTypeQuery ParameterType = `query`

const ParameterTypeText ParameterType = `text`

// String representation for [fmt.Print]
func (f *ParameterType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ParameterType) Set(v string) error {
	switch v {
	case `datetime`, `enum`, `number`, `query`, `text`:
		*f = ParameterType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "datetime", "enum", "number", "query", "text"`, v)
	}
}

// Values returns all possible values for ParameterType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ParameterType) Values() []ParameterType {
	return []ParameterType{
		ParameterTypeDatetime,
		ParameterTypeEnum,
		ParameterTypeNumber,
		ParameterTypeQuery,
		ParameterTypeText,
	}
}

// Type always returns ParameterType to satisfy [pflag.Value] interface
func (f *ParameterType) Type() string {
	return "ParameterType"
}

// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query * `CAN_EDIT`:
// Can edit the query * `CAN_MANAGE`: Can manage the query
type PermissionLevel string

// Can edit the query
const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

// Can manage the query
const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

// Can run the query
const PermissionLevelCanRun PermissionLevel = `CAN_RUN`

// Can view the query
const PermissionLevelCanView PermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_RUN`, `CAN_VIEW`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_RUN", "CAN_VIEW"`, v)
	}
}

// Values returns all possible values for PermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PermissionLevel) Values() []PermissionLevel {
	return []PermissionLevel{
		PermissionLevelCanEdit,
		PermissionLevelCanManage,
		PermissionLevelCanRun,
		PermissionLevelCanView,
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

// Possible Reasons for which we have not saved plans in the database
type PlansState string

const PlansStateEmpty PlansState = `EMPTY`

const PlansStateExists PlansState = `EXISTS`

const PlansStateIgnoredLargePlansSize PlansState = `IGNORED_LARGE_PLANS_SIZE`

const PlansStateIgnoredSmallDuration PlansState = `IGNORED_SMALL_DURATION`

const PlansStateIgnoredSparkPlanType PlansState = `IGNORED_SPARK_PLAN_TYPE`

const PlansStateUnknown PlansState = `UNKNOWN`

// String representation for [fmt.Print]
func (f *PlansState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PlansState) Set(v string) error {
	switch v {
	case `EMPTY`, `EXISTS`, `IGNORED_LARGE_PLANS_SIZE`, `IGNORED_SMALL_DURATION`, `IGNORED_SPARK_PLAN_TYPE`, `UNKNOWN`:
		*f = PlansState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMPTY", "EXISTS", "IGNORED_LARGE_PLANS_SIZE", "IGNORED_SMALL_DURATION", "IGNORED_SPARK_PLAN_TYPE", "UNKNOWN"`, v)
	}
}

// Values returns all possible values for PlansState.
//
// There is no guarantee on the order of the values in the slice.
func (f *PlansState) Values() []PlansState {
	return []PlansState{
		PlansStateEmpty,
		PlansStateExists,
		PlansStateIgnoredLargePlansSize,
		PlansStateIgnoredSmallDuration,
		PlansStateIgnoredSparkPlanType,
		PlansStateUnknown,
	}
}

// Type always returns PlansState to satisfy [pflag.Value] interface
func (f *PlansState) Type() string {
	return "PlansState"
}

type Query struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// Timestamp when this query was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the query.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Username of the user who last saved changes to this query.
	// Wire name: 'last_modifier_user_name'
	LastModifierUserName string `json:"last_modifier_user_name,omitempty"`
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Workspace path of the workspace folder containing the object.
	// Wire name: 'parent_path'
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`
	// Timestamp when this query was last updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Query) EncodeValues(key string, v *url.Values) error {
	pb, err := queryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Query) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Query) MarshalJSON() ([]byte, error) {
	pb, err := queryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	// Wire name: 'multi_values_options'
	MultiValuesOptions *MultiValuesOptions `json:"multi_values_options,omitempty"`
	// UUID of the query that provides the parameter values.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// List of selected query parameter values.
	// Wire name: 'values'
	Values []string `json:"values,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryBackedValue) EncodeValues(key string, v *url.Values) error {
	pb, err := queryBackedValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryBackedValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryBackedValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryBackedValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryBackedValue) MarshalJSON() ([]byte, error) {
	pb, err := queryBackedValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryEditContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	// Wire name: 'options'
	Options any `json:"options,omitempty"`
	// The text of the query to be run.
	// Wire name: 'query'
	Query string `json:"query,omitempty"`

	QueryId string `json:"-" tf:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryEditContent) EncodeValues(key string, v *url.Values) error {
	pb, err := queryEditContentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryEditContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryEditContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryEditContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryEditContent) MarshalJSON() ([]byte, error) {
	pb, err := queryEditContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryFilter struct {
	// A range filter for query submitted time. The time range must be less than
	// or equal to 30 days.
	// Wire name: 'query_start_time_range'
	QueryStartTimeRange *TimeRange `json:"query_start_time_range,omitempty"`
	// A list of statement IDs.
	// Wire name: 'statement_ids'
	StatementIds []string `json:"statement_ids,omitempty"`
	// A list of statuses (QUEUED, RUNNING, CANCELED, FAILED, FINISHED) to match
	// query results. Corresponds to the `status` field in the response.
	// Filtering for multiple statuses is not recommended. Instead, opt to
	// filter by a single status multiple times and then combine the results.
	// Wire name: 'statuses'
	Statuses []QueryStatus `json:"statuses,omitempty"`
	// A list of user IDs who ran the queries.
	// Wire name: 'user_ids'
	UserIds []int64 `json:"user_ids,omitempty"`
	// A list of warehouse IDs.
	// Wire name: 'warehouse_ids'
	WarehouseIds []string `json:"warehouse_ids,omitempty"`
}

func (st QueryFilter) EncodeValues(key string, v *url.Values) error {
	pb, err := queryFilterToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryFilter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryFilterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryFilterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryFilter) MarshalJSON() ([]byte, error) {
	pb, err := queryFilterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryInfo struct {
	// SQL Warehouse channel information at the time of query execution
	// Wire name: 'channel_used'
	ChannelUsed *ChannelInfo `json:"channel_used,omitempty"`
	// Client application that ran the statement. For example: Databricks SQL
	// Editor, Tableau, and Power BI. This field is derived from information
	// provided by client applications. While values are expected to remain
	// static over time, this cannot be guaranteed.
	// Wire name: 'client_application'
	ClientApplication string `json:"client_application,omitempty"`
	// Total time of the statement execution. This value does not include the
	// time taken to retrieve the results, which can result in a discrepancy
	// between this value and the start-to-finish wall-clock time.
	// Wire name: 'duration'
	Duration int64 `json:"duration,omitempty"`
	// Alias for `warehouse_id`.
	// Wire name: 'endpoint_id'
	EndpointId string `json:"endpoint_id,omitempty"`
	// Message describing why the query could not complete.
	// Wire name: 'error_message'
	ErrorMessage string `json:"error_message,omitempty"`
	// The ID of the user whose credentials were used to run the query.
	// Wire name: 'executed_as_user_id'
	ExecutedAsUserId int64 `json:"executed_as_user_id,omitempty"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	// Wire name: 'executed_as_user_name'
	ExecutedAsUserName string `json:"executed_as_user_name,omitempty"`
	// The time execution of the query ended.
	// Wire name: 'execution_end_time_ms'
	ExecutionEndTimeMs int64 `json:"execution_end_time_ms,omitempty"`
	// Whether more updates for the query are expected.
	// Wire name: 'is_final'
	IsFinal bool `json:"is_final,omitempty"`
	// A key that can be used to look up query details.
	// Wire name: 'lookup_key'
	LookupKey string `json:"lookup_key,omitempty"`
	// Metrics about query execution.
	// Wire name: 'metrics'
	Metrics *QueryMetrics `json:"metrics,omitempty"`
	// Whether plans exist for the execution, or the reason why they are missing
	// Wire name: 'plans_state'
	PlansState PlansState `json:"plans_state,omitempty"`
	// The time the query ended.
	// Wire name: 'query_end_time_ms'
	QueryEndTimeMs int64 `json:"query_end_time_ms,omitempty"`
	// The query ID.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	// Wire name: 'query_source'
	QuerySource *ExternalQuerySource `json:"query_source,omitempty"`
	// The time the query started.
	// Wire name: 'query_start_time_ms'
	QueryStartTimeMs int64 `json:"query_start_time_ms,omitempty"`
	// The text of the query.
	// Wire name: 'query_text'
	QueryText string `json:"query_text,omitempty"`
	// The number of results returned by the query.
	// Wire name: 'rows_produced'
	RowsProduced int64 `json:"rows_produced,omitempty"`
	// URL to the Spark UI query plan.
	// Wire name: 'spark_ui_url'
	SparkUiUrl string `json:"spark_ui_url,omitempty"`
	// Type of statement for this query
	// Wire name: 'statement_type'
	StatementType QueryStatementType `json:"statement_type,omitempty"`
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	// Wire name: 'status'
	Status QueryStatus `json:"status,omitempty"`
	// The ID of the user who ran the query.
	// Wire name: 'user_id'
	UserId int64 `json:"user_id,omitempty"`
	// The email address or username of the user who ran the query.
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`
	// Warehouse ID.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := queryInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryInfo) MarshalJSON() ([]byte, error) {
	pb, err := queryInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryList struct {
	// The total number of queries.
	// Wire name: 'count'
	Count int `json:"count,omitempty"`
	// The page number that is currently displayed.
	// Wire name: 'page'
	Page int `json:"page,omitempty"`
	// The number of queries per page.
	// Wire name: 'page_size'
	PageSize int `json:"page_size,omitempty"`
	// List of queries returned.
	// Wire name: 'results'
	Results []LegacyQuery `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryList) EncodeValues(key string, v *url.Values) error {
	pb, err := queryListToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryList) MarshalJSON() ([]byte, error) {
	pb, err := queryListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	// Wire name: 'compilation_time_ms'
	CompilationTimeMs int64 `json:"compilation_time_ms,omitempty"`
	// Time spent executing the query, in milliseconds.
	// Wire name: 'execution_time_ms'
	ExecutionTimeMs int64 `json:"execution_time_ms,omitempty"`
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	// Wire name: 'network_sent_bytes'
	NetworkSentBytes int64 `json:"network_sent_bytes,omitempty"`
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	// Wire name: 'overloading_queue_start_timestamp'
	OverloadingQueueStartTimestamp int64 `json:"overloading_queue_start_timestamp,omitempty"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	// Wire name: 'photon_total_time_ms'
	PhotonTotalTimeMs int64 `json:"photon_total_time_ms,omitempty"`
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	// Wire name: 'provisioning_queue_start_timestamp'
	ProvisioningQueueStartTimestamp int64 `json:"provisioning_queue_start_timestamp,omitempty"`
	// Total number of bytes in all tables not read due to pruning
	// Wire name: 'pruned_bytes'
	PrunedBytes int64 `json:"pruned_bytes,omitempty"`
	// Total number of files from all tables not read due to pruning
	// Wire name: 'pruned_files_count'
	PrunedFilesCount int64 `json:"pruned_files_count,omitempty"`
	// Timestamp of when the underlying compute started compilation of the
	// query.
	// Wire name: 'query_compilation_start_timestamp'
	QueryCompilationStartTimestamp int64 `json:"query_compilation_start_timestamp,omitempty"`
	// Total size of data read by the query, in bytes.
	// Wire name: 'read_bytes'
	ReadBytes int64 `json:"read_bytes,omitempty"`
	// Size of persistent data read from the cache, in bytes.
	// Wire name: 'read_cache_bytes'
	ReadCacheBytes int64 `json:"read_cache_bytes,omitempty"`
	// Number of files read after pruning
	// Wire name: 'read_files_count'
	ReadFilesCount int64 `json:"read_files_count,omitempty"`
	// Number of partitions read after pruning.
	// Wire name: 'read_partitions_count'
	ReadPartitionsCount int64 `json:"read_partitions_count,omitempty"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	// Wire name: 'read_remote_bytes'
	ReadRemoteBytes int64 `json:"read_remote_bytes,omitempty"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	// Wire name: 'result_fetch_time_ms'
	ResultFetchTimeMs int64 `json:"result_fetch_time_ms,omitempty"`
	// `true` if the query result was fetched from cache, `false` otherwise.
	// Wire name: 'result_from_cache'
	ResultFromCache bool `json:"result_from_cache,omitempty"`
	// Total number of rows returned by the query.
	// Wire name: 'rows_produced_count'
	RowsProducedCount int64 `json:"rows_produced_count,omitempty"`
	// Total number of rows read by the query.
	// Wire name: 'rows_read_count'
	RowsReadCount int64 `json:"rows_read_count,omitempty"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	// Wire name: 'spill_to_disk_bytes'
	SpillToDiskBytes int64 `json:"spill_to_disk_bytes,omitempty"`
	// sum of task times completed in a range of wall clock time, approximated
	// to a configurable number of points aggregated over all stages and jobs in
	// the query (based on task_total_time_ms)
	// Wire name: 'task_time_over_time_range'
	TaskTimeOverTimeRange *TaskTimeOverRange `json:"task_time_over_time_range,omitempty"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	// Wire name: 'task_total_time_ms'
	TaskTotalTimeMs int64 `json:"task_total_time_ms,omitempty"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	// Wire name: 'total_time_ms'
	TotalTimeMs int64 `json:"total_time_ms,omitempty"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	// Wire name: 'write_remote_bytes'
	WriteRemoteBytes int64 `json:"write_remote_bytes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryMetrics) EncodeValues(key string, v *url.Values) error {
	pb, err := queryMetricsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryMetrics) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryMetricsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryMetricsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryMetrics) MarshalJSON() ([]byte, error) {
	pb, err := queryMetricsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryOptions struct {
	// The name of the catalog to execute this query in.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	// Wire name: 'moved_to_trash_at'
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	// Wire name: 'parameters'
	Parameters []Parameter `json:"parameters,omitempty"`
	// The name of the schema to execute this query in.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryOptions) EncodeValues(key string, v *url.Values) error {
	pb, err := queryOptionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryOptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryOptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryOptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryOptions) MarshalJSON() ([]byte, error) {
	pb, err := queryOptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryParameter struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	// Wire name: 'date_range_value'
	DateRangeValue *DateRangeValue `json:"date_range_value,omitempty"`
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	// Wire name: 'date_value'
	DateValue *DateValue `json:"date_value,omitempty"`
	// Dropdown query parameter value.
	// Wire name: 'enum_value'
	EnumValue *EnumValue `json:"enum_value,omitempty"`
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Numeric query parameter value.
	// Wire name: 'numeric_value'
	NumericValue *NumericValue `json:"numeric_value,omitempty"`
	// Query-based dropdown query parameter value.
	// Wire name: 'query_backed_value'
	QueryBackedValue *QueryBackedValue `json:"query_backed_value,omitempty"`
	// Text query parameter value.
	// Wire name: 'text_value'
	TextValue *TextValue `json:"text_value,omitempty"`
	// Text displayed in the user-facing parameter widget in the UI.
	// Wire name: 'title'
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryParameter) EncodeValues(key string, v *url.Values) error {
	pb, err := queryParameterToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryParameter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryParameterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryParameterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryParameter) MarshalJSON() ([]byte, error) {
	pb, err := queryParameterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryPostContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	// Wire name: 'options'
	Options any `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string `json:"parent,omitempty"`
	// The text of the query to be run.
	// Wire name: 'query'
	Query string `json:"query,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryPostContent) EncodeValues(key string, v *url.Values) error {
	pb, err := queryPostContentToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryPostContent) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryPostContentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryPostContentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryPostContent) MarshalJSON() ([]byte, error) {
	pb, err := queryPostContentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryStatementType string

const QueryStatementTypeAlter QueryStatementType = `ALTER`

const QueryStatementTypeAnalyze QueryStatementType = `ANALYZE`

const QueryStatementTypeCopy QueryStatementType = `COPY`

const QueryStatementTypeCreate QueryStatementType = `CREATE`

const QueryStatementTypeDelete QueryStatementType = `DELETE`

const QueryStatementTypeDescribe QueryStatementType = `DESCRIBE`

const QueryStatementTypeDrop QueryStatementType = `DROP`

const QueryStatementTypeExplain QueryStatementType = `EXPLAIN`

const QueryStatementTypeGrant QueryStatementType = `GRANT`

const QueryStatementTypeInsert QueryStatementType = `INSERT`

const QueryStatementTypeMerge QueryStatementType = `MERGE`

const QueryStatementTypeOptimize QueryStatementType = `OPTIMIZE`

const QueryStatementTypeOther QueryStatementType = `OTHER`

const QueryStatementTypeRefresh QueryStatementType = `REFRESH`

const QueryStatementTypeReplace QueryStatementType = `REPLACE`

const QueryStatementTypeRevoke QueryStatementType = `REVOKE`

const QueryStatementTypeSelect QueryStatementType = `SELECT`

const QueryStatementTypeSet QueryStatementType = `SET`

const QueryStatementTypeShow QueryStatementType = `SHOW`

const QueryStatementTypeTruncate QueryStatementType = `TRUNCATE`

const QueryStatementTypeUpdate QueryStatementType = `UPDATE`

const QueryStatementTypeUse QueryStatementType = `USE`

// String representation for [fmt.Print]
func (f *QueryStatementType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueryStatementType) Set(v string) error {
	switch v {
	case `ALTER`, `ANALYZE`, `COPY`, `CREATE`, `DELETE`, `DESCRIBE`, `DROP`, `EXPLAIN`, `GRANT`, `INSERT`, `MERGE`, `OPTIMIZE`, `OTHER`, `REFRESH`, `REPLACE`, `REVOKE`, `SELECT`, `SET`, `SHOW`, `TRUNCATE`, `UPDATE`, `USE`:
		*f = QueryStatementType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALTER", "ANALYZE", "COPY", "CREATE", "DELETE", "DESCRIBE", "DROP", "EXPLAIN", "GRANT", "INSERT", "MERGE", "OPTIMIZE", "OTHER", "REFRESH", "REPLACE", "REVOKE", "SELECT", "SET", "SHOW", "TRUNCATE", "UPDATE", "USE"`, v)
	}
}

// Values returns all possible values for QueryStatementType.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueryStatementType) Values() []QueryStatementType {
	return []QueryStatementType{
		QueryStatementTypeAlter,
		QueryStatementTypeAnalyze,
		QueryStatementTypeCopy,
		QueryStatementTypeCreate,
		QueryStatementTypeDelete,
		QueryStatementTypeDescribe,
		QueryStatementTypeDrop,
		QueryStatementTypeExplain,
		QueryStatementTypeGrant,
		QueryStatementTypeInsert,
		QueryStatementTypeMerge,
		QueryStatementTypeOptimize,
		QueryStatementTypeOther,
		QueryStatementTypeRefresh,
		QueryStatementTypeReplace,
		QueryStatementTypeRevoke,
		QueryStatementTypeSelect,
		QueryStatementTypeSet,
		QueryStatementTypeShow,
		QueryStatementTypeTruncate,
		QueryStatementTypeUpdate,
		QueryStatementTypeUse,
	}
}

// Type always returns QueryStatementType to satisfy [pflag.Value] interface
func (f *QueryStatementType) Type() string {
	return "QueryStatementType"
}

// Statuses which are also used by OperationStatus in runtime
type QueryStatus string

const QueryStatusCanceled QueryStatus = `CANCELED`

const QueryStatusCompiled QueryStatus = `COMPILED`

const QueryStatusCompiling QueryStatus = `COMPILING`

const QueryStatusFailed QueryStatus = `FAILED`

const QueryStatusFinished QueryStatus = `FINISHED`

const QueryStatusQueued QueryStatus = `QUEUED`

const QueryStatusRunning QueryStatus = `RUNNING`

const QueryStatusStarted QueryStatus = `STARTED`

// String representation for [fmt.Print]
func (f *QueryStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueryStatus) Set(v string) error {
	switch v {
	case `CANCELED`, `COMPILED`, `COMPILING`, `FAILED`, `FINISHED`, `QUEUED`, `RUNNING`, `STARTED`:
		*f = QueryStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "COMPILED", "COMPILING", "FAILED", "FINISHED", "QUEUED", "RUNNING", "STARTED"`, v)
	}
}

// Values returns all possible values for QueryStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueryStatus) Values() []QueryStatus {
	return []QueryStatus{
		QueryStatusCanceled,
		QueryStatusCompiled,
		QueryStatusCompiling,
		QueryStatusFailed,
		QueryStatusFinished,
		QueryStatusQueued,
		QueryStatusRunning,
		QueryStatusStarted,
	}
}

// Type always returns QueryStatus to satisfy [pflag.Value] interface
func (f *QueryStatus) Type() string {
	return "QueryStatus"
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	// Wire name: 'config_pair'
	ConfigPair []EndpointConfPair `json:"config_pair,omitempty"`

	// Wire name: 'configuration_pairs'
	ConfigurationPairs []EndpointConfPair `json:"configuration_pairs,omitempty"`
}

func (st RepeatedEndpointConfPairs) EncodeValues(key string, v *url.Values) error {
	pb, err := repeatedEndpointConfPairsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RepeatedEndpointConfPairs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &repeatedEndpointConfPairsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := repeatedEndpointConfPairsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RepeatedEndpointConfPairs) MarshalJSON() ([]byte, error) {
	pb, err := repeatedEndpointConfPairsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreDashboardRequest struct {
	DashboardId string `json:"-" tf:"-"`
}

func (st RestoreDashboardRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := restoreDashboardRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RestoreDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := restoreDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestoreQueriesLegacyRequest struct {
	QueryId string `json:"-" tf:"-"`
}

func (st RestoreQueriesLegacyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := restoreQueriesLegacyRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RestoreQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreQueriesLegacyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreQueriesLegacyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreQueriesLegacyRequest) MarshalJSON() ([]byte, error) {
	pb, err := restoreQueriesLegacyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ResultData struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	// Wire name: 'data_array'
	DataArray [][]string `json:"data_array,omitempty"`

	// Wire name: 'external_links'
	ExternalLinks []ExternalLink `json:"external_links,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	// Wire name: 'next_chunk_index'
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	// Wire name: 'next_chunk_internal_link'
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ResultData) EncodeValues(key string, v *url.Values) error {
	pb, err := resultDataToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ResultData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resultDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resultDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResultData) MarshalJSON() ([]byte, error) {
	pb, err := resultDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	// Wire name: 'chunks'
	Chunks []BaseChunkInfo `json:"chunks,omitempty"`

	// Wire name: 'format'
	Format Format `json:"format,omitempty"`

	// Wire name: 'schema'
	Schema *ResultSchema `json:"schema,omitempty"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	// Wire name: 'total_byte_count'
	TotalByteCount int64 `json:"total_byte_count,omitempty"`
	// The total number of chunks that the result set has been divided into.
	// Wire name: 'total_chunk_count'
	TotalChunkCount int `json:"total_chunk_count,omitempty"`
	// The total number of rows in the result set.
	// Wire name: 'total_row_count'
	TotalRowCount int64 `json:"total_row_count,omitempty"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	// Wire name: 'truncated'
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ResultManifest) EncodeValues(key string, v *url.Values) error {
	pb, err := resultManifestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ResultManifest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resultManifestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resultManifestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResultManifest) MarshalJSON() ([]byte, error) {
	pb, err := resultManifestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {

	// Wire name: 'column_count'
	ColumnCount int `json:"column_count,omitempty"`

	// Wire name: 'columns'
	Columns []ColumnInfo `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ResultSchema) EncodeValues(key string, v *url.Values) error {
	pb, err := resultSchemaToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ResultSchema) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &resultSchemaPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := resultSchemaFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ResultSchema) MarshalJSON() ([]byte, error) {
	pb, err := resultSchemaToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RunAsMode string

const RunAsModeOwner RunAsMode = `OWNER`

const RunAsModeViewer RunAsMode = `VIEWER`

// String representation for [fmt.Print]
func (f *RunAsMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunAsMode) Set(v string) error {
	switch v {
	case `OWNER`, `VIEWER`:
		*f = RunAsMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "OWNER", "VIEWER"`, v)
	}
}

// Values returns all possible values for RunAsMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunAsMode) Values() []RunAsMode {
	return []RunAsMode{
		RunAsModeOwner,
		RunAsModeViewer,
	}
}

// Type always returns RunAsMode to satisfy [pflag.Value] interface
func (f *RunAsMode) Type() string {
	return "RunAsMode"
}

type RunAsRole string

const RunAsRoleOwner RunAsRole = `owner`

const RunAsRoleViewer RunAsRole = `viewer`

// String representation for [fmt.Print]
func (f *RunAsRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunAsRole) Set(v string) error {
	switch v {
	case `owner`, `viewer`:
		*f = RunAsRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "owner", "viewer"`, v)
	}
}

// Values returns all possible values for RunAsRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunAsRole) Values() []RunAsRole {
	return []RunAsRole{
		RunAsRoleOwner,
		RunAsRoleViewer,
	}
}

// Type always returns RunAsRole to satisfy [pflag.Value] interface
func (f *RunAsRole) Type() string {
	return "RunAsRole"
}

type SchedulePauseStatus string

const SchedulePauseStatusPaused SchedulePauseStatus = `PAUSED`

const SchedulePauseStatusUnpaused SchedulePauseStatus = `UNPAUSED`

// String representation for [fmt.Print]
func (f *SchedulePauseStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SchedulePauseStatus) Set(v string) error {
	switch v {
	case `PAUSED`, `UNPAUSED`:
		*f = SchedulePauseStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "PAUSED", "UNPAUSED"`, v)
	}
}

// Values returns all possible values for SchedulePauseStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *SchedulePauseStatus) Values() []SchedulePauseStatus {
	return []SchedulePauseStatus{
		SchedulePauseStatusPaused,
		SchedulePauseStatusUnpaused,
	}
}

// Type always returns SchedulePauseStatus to satisfy [pflag.Value] interface
func (f *SchedulePauseStatus) Type() string {
	return "SchedulePauseStatus"
}

type ServiceError struct {

	// Wire name: 'error_code'
	ErrorCode ServiceErrorCode `json:"error_code,omitempty"`
	// A brief summary of the error condition.
	// Wire name: 'message'
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServiceError) EncodeValues(key string, v *url.Values) error {
	pb, err := serviceErrorToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServiceError) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &serviceErrorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := serviceErrorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServiceError) MarshalJSON() ([]byte, error) {
	pb, err := serviceErrorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServiceErrorCode string

const ServiceErrorCodeAborted ServiceErrorCode = `ABORTED`

const ServiceErrorCodeAlreadyExists ServiceErrorCode = `ALREADY_EXISTS`

const ServiceErrorCodeBadRequest ServiceErrorCode = `BAD_REQUEST`

const ServiceErrorCodeCancelled ServiceErrorCode = `CANCELLED`

const ServiceErrorCodeDeadlineExceeded ServiceErrorCode = `DEADLINE_EXCEEDED`

const ServiceErrorCodeInternalError ServiceErrorCode = `INTERNAL_ERROR`

const ServiceErrorCodeIoError ServiceErrorCode = `IO_ERROR`

const ServiceErrorCodeNotFound ServiceErrorCode = `NOT_FOUND`

const ServiceErrorCodeResourceExhausted ServiceErrorCode = `RESOURCE_EXHAUSTED`

const ServiceErrorCodeServiceUnderMaintenance ServiceErrorCode = `SERVICE_UNDER_MAINTENANCE`

const ServiceErrorCodeTemporarilyUnavailable ServiceErrorCode = `TEMPORARILY_UNAVAILABLE`

const ServiceErrorCodeUnauthenticated ServiceErrorCode = `UNAUTHENTICATED`

const ServiceErrorCodeUnknown ServiceErrorCode = `UNKNOWN`

const ServiceErrorCodeWorkspaceTemporarilyUnavailable ServiceErrorCode = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

// String representation for [fmt.Print]
func (f *ServiceErrorCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServiceErrorCode) Set(v string) error {
	switch v {
	case `ABORTED`, `ALREADY_EXISTS`, `BAD_REQUEST`, `CANCELLED`, `DEADLINE_EXCEEDED`, `INTERNAL_ERROR`, `IO_ERROR`, `NOT_FOUND`, `RESOURCE_EXHAUSTED`, `SERVICE_UNDER_MAINTENANCE`, `TEMPORARILY_UNAVAILABLE`, `UNAUTHENTICATED`, `UNKNOWN`, `WORKSPACE_TEMPORARILY_UNAVAILABLE`:
		*f = ServiceErrorCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABORTED", "ALREADY_EXISTS", "BAD_REQUEST", "CANCELLED", "DEADLINE_EXCEEDED", "INTERNAL_ERROR", "IO_ERROR", "NOT_FOUND", "RESOURCE_EXHAUSTED", "SERVICE_UNDER_MAINTENANCE", "TEMPORARILY_UNAVAILABLE", "UNAUTHENTICATED", "UNKNOWN", "WORKSPACE_TEMPORARILY_UNAVAILABLE"`, v)
	}
}

// Values returns all possible values for ServiceErrorCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServiceErrorCode) Values() []ServiceErrorCode {
	return []ServiceErrorCode{
		ServiceErrorCodeAborted,
		ServiceErrorCodeAlreadyExists,
		ServiceErrorCodeBadRequest,
		ServiceErrorCodeCancelled,
		ServiceErrorCodeDeadlineExceeded,
		ServiceErrorCodeInternalError,
		ServiceErrorCodeIoError,
		ServiceErrorCodeNotFound,
		ServiceErrorCodeResourceExhausted,
		ServiceErrorCodeServiceUnderMaintenance,
		ServiceErrorCodeTemporarilyUnavailable,
		ServiceErrorCodeUnauthenticated,
		ServiceErrorCodeUnknown,
		ServiceErrorCodeWorkspaceTemporarilyUnavailable,
	}
}

// Type always returns ServiceErrorCode to satisfy [pflag.Value] interface
func (f *ServiceErrorCode) Type() string {
	return "ServiceErrorCode"
}

// Set object ACL
type SetRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId string `json:"-" tf:"-"`
	// The type of object permission to set.
	ObjectType ObjectTypePlural `json:"-" tf:"-"`
}

func (st SetRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := setRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetRequest) MarshalJSON() ([]byte, error) {
	pb, err := setRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetResponse struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	// An object's type and UUID, separated by a forward slash (/) character.
	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`
	// A singular noun object type.
	// Wire name: 'object_type'
	ObjectType ObjectType `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SetResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := setResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SetResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetResponse) MarshalJSON() ([]byte, error) {
	pb, err := setResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	// Wire name: 'channel'
	Channel *Channel `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'config_param'
	ConfigParam *RepeatedEndpointConfPairs `json:"config_param,omitempty"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	// Wire name: 'data_access_config'
	DataAccessConfig []EndpointConfPair `json:"data_access_config,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	// Wire name: 'enabled_warehouse_types'
	EnabledWarehouseTypes []WarehouseTypePair `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'global_param'
	GlobalParam *RepeatedEndpointConfPairs `json:"global_param,omitempty"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	// Wire name: 'google_service_account'
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for warehouses
	// Wire name: 'security_policy'
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy `json:"security_policy,omitempty"`
	// SQL configuration parameters
	// Wire name: 'sql_configuration_parameters'
	SqlConfigurationParameters *RepeatedEndpointConfPairs `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SetWorkspaceWarehouseConfigRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := setWorkspaceWarehouseConfigRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SetWorkspaceWarehouseConfigRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setWorkspaceWarehouseConfigRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setWorkspaceWarehouseConfigRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetWorkspaceWarehouseConfigRequest) MarshalJSON() ([]byte, error) {
	pb, err := setWorkspaceWarehouseConfigRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Security policy for warehouses
type SetWorkspaceWarehouseConfigRequestSecurityPolicy string

const SetWorkspaceWarehouseConfigRequestSecurityPolicyDataAccessControl SetWorkspaceWarehouseConfigRequestSecurityPolicy = `DATA_ACCESS_CONTROL`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyNone SetWorkspaceWarehouseConfigRequestSecurityPolicy = `NONE`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyPassthrough SetWorkspaceWarehouseConfigRequestSecurityPolicy = `PASSTHROUGH`

// String representation for [fmt.Print]
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Set(v string) error {
	switch v {
	case `DATA_ACCESS_CONTROL`, `NONE`, `PASSTHROUGH`:
		*f = SetWorkspaceWarehouseConfigRequestSecurityPolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DATA_ACCESS_CONTROL", "NONE", "PASSTHROUGH"`, v)
	}
}

// Values returns all possible values for SetWorkspaceWarehouseConfigRequestSecurityPolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Values() []SetWorkspaceWarehouseConfigRequestSecurityPolicy {
	return []SetWorkspaceWarehouseConfigRequestSecurityPolicy{
		SetWorkspaceWarehouseConfigRequestSecurityPolicyDataAccessControl,
		SetWorkspaceWarehouseConfigRequestSecurityPolicyNone,
		SetWorkspaceWarehouseConfigRequestSecurityPolicyPassthrough,
	}
}

// Type always returns SetWorkspaceWarehouseConfigRequestSecurityPolicy to satisfy [pflag.Value] interface
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Type() string {
	return "SetWorkspaceWarehouseConfigRequestSecurityPolicy"
}

// Configurations whether the warehouse should use spot instances.
type SpotInstancePolicy string

const SpotInstancePolicyCostOptimized SpotInstancePolicy = `COST_OPTIMIZED`

const SpotInstancePolicyPolicyUnspecified SpotInstancePolicy = `POLICY_UNSPECIFIED`

const SpotInstancePolicyReliabilityOptimized SpotInstancePolicy = `RELIABILITY_OPTIMIZED`

// String representation for [fmt.Print]
func (f *SpotInstancePolicy) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SpotInstancePolicy) Set(v string) error {
	switch v {
	case `COST_OPTIMIZED`, `POLICY_UNSPECIFIED`, `RELIABILITY_OPTIMIZED`:
		*f = SpotInstancePolicy(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COST_OPTIMIZED", "POLICY_UNSPECIFIED", "RELIABILITY_OPTIMIZED"`, v)
	}
}

// Values returns all possible values for SpotInstancePolicy.
//
// There is no guarantee on the order of the values in the slice.
func (f *SpotInstancePolicy) Values() []SpotInstancePolicy {
	return []SpotInstancePolicy{
		SpotInstancePolicyCostOptimized,
		SpotInstancePolicyPolicyUnspecified,
		SpotInstancePolicyReliabilityOptimized,
	}
}

// Type always returns SpotInstancePolicy to satisfy [pflag.Value] interface
func (f *SpotInstancePolicy) Type() string {
	return "SpotInstancePolicy"
}

type StartRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" tf:"-"`
}

func (st StartRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := startRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StartRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartRequest) MarshalJSON() ([]byte, error) {
	pb, err := startRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// State of the warehouse
type State string

const StateDeleted State = `DELETED`

const StateDeleting State = `DELETING`

const StateRunning State = `RUNNING`

const StateStarting State = `STARTING`

const StateStopped State = `STOPPED`

const StateStopping State = `STOPPING`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `DELETED`, `DELETING`, `RUNNING`, `STARTING`, `STOPPED`, `STOPPING`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETED", "DELETING", "RUNNING", "STARTING", "STOPPED", "STOPPING"`, v)
	}
}

// Values returns all possible values for State.
//
// There is no guarantee on the order of the values in the slice.
func (f *State) Values() []State {
	return []State{
		StateDeleted,
		StateDeleting,
		StateRunning,
		StateStarting,
		StateStopped,
		StateStopping,
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

type StatementParameterListItem struct {
	// The name of a parameter marker to be substituted in the statement.
	// Wire name: 'name'
	Name string `json:"name"`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	// Wire name: 'type'
	Type string `json:"type,omitempty"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st StatementParameterListItem) EncodeValues(key string, v *url.Values) error {
	pb, err := statementParameterListItemToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StatementParameterListItem) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &statementParameterListItemPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := statementParameterListItemFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StatementParameterListItem) MarshalJSON() ([]byte, error) {
	pb, err := statementParameterListItemToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StatementResponse struct {

	// Wire name: 'manifest'
	Manifest *ResultManifest `json:"manifest,omitempty"`

	// Wire name: 'result'
	Result *ResultData `json:"result,omitempty"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string `json:"statement_id,omitempty"`

	// Wire name: 'status'
	Status *StatementStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st StatementResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := statementResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StatementResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &statementResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := statementResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StatementResponse) MarshalJSON() ([]byte, error) {
	pb, err := statementResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Statement execution state: - `PENDING`: waiting for warehouse - `RUNNING`:
// running - `SUCCEEDED`: execution was successful, result data available for
// fetch - `FAILED`: execution failed; reason for failure described in
// accomanying error message - `CANCELED`: user canceled; can come from explicit
// cancel call, or timeout with `on_wait_timeout=CANCEL` - `CLOSED`: execution
// successful, and statement closed; result no longer available for fetch
type StatementState string

// user canceled; can come from explicit cancel call, or timeout with
// `on_wait_timeout=CANCEL`
const StatementStateCanceled StatementState = `CANCELED`

// execution successful, and statement closed; result no longer available for
// fetch
const StatementStateClosed StatementState = `CLOSED`

// execution failed; reason for failure described in accomanying error message
const StatementStateFailed StatementState = `FAILED`

// waiting for warehouse
const StatementStatePending StatementState = `PENDING`

// running
const StatementStateRunning StatementState = `RUNNING`

// execution was successful, result data available for fetch
const StatementStateSucceeded StatementState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *StatementState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *StatementState) Set(v string) error {
	switch v {
	case `CANCELED`, `CLOSED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = StatementState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELED", "CLOSED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Values returns all possible values for StatementState.
//
// There is no guarantee on the order of the values in the slice.
func (f *StatementState) Values() []StatementState {
	return []StatementState{
		StatementStateCanceled,
		StatementStateClosed,
		StatementStateFailed,
		StatementStatePending,
		StatementStateRunning,
		StatementStateSucceeded,
	}
}

// Type always returns StatementState to satisfy [pflag.Value] interface
func (f *StatementState) Type() string {
	return "StatementState"
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {

	// Wire name: 'error'
	Error *ServiceError `json:"error,omitempty"`

	// Wire name: 'state'
	State StatementState `json:"state,omitempty"`
}

func (st StatementStatus) EncodeValues(key string, v *url.Values) error {
	pb, err := statementStatusToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StatementStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &statementStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := statementStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StatementStatus) MarshalJSON() ([]byte, error) {
	pb, err := statementStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Health status of the warehouse.
type Status string

const StatusDegraded Status = `DEGRADED`

const StatusFailed Status = `FAILED`

const StatusHealthy Status = `HEALTHY`

const StatusStatusUnspecified Status = `STATUS_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *Status) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Status) Set(v string) error {
	switch v {
	case `DEGRADED`, `FAILED`, `HEALTHY`, `STATUS_UNSPECIFIED`:
		*f = Status(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEGRADED", "FAILED", "HEALTHY", "STATUS_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for Status.
//
// There is no guarantee on the order of the values in the slice.
func (f *Status) Values() []Status {
	return []Status{
		StatusDegraded,
		StatusFailed,
		StatusHealthy,
		StatusStatusUnspecified,
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

type StopRequest struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" tf:"-"`
}

func (st StopRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := stopRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *StopRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stopRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stopRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StopRequest) MarshalJSON() ([]byte, error) {
	pb, err := stopRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Success struct {

	// Wire name: 'message'
	Message SuccessMessage `json:"message,omitempty"`
}

func (st Success) EncodeValues(key string, v *url.Values) error {
	pb, err := successToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Success) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &successPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := successFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Success) MarshalJSON() ([]byte, error) {
	pb, err := successToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SuccessMessage string

const SuccessMessageSuccess SuccessMessage = `Success`

// String representation for [fmt.Print]
func (f *SuccessMessage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SuccessMessage) Set(v string) error {
	switch v {
	case `Success`:
		*f = SuccessMessage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Success"`, v)
	}
}

// Values returns all possible values for SuccessMessage.
//
// There is no guarantee on the order of the values in the slice.
func (f *SuccessMessage) Values() []SuccessMessage {
	return []SuccessMessage{
		SuccessMessageSuccess,
	}
}

// Type always returns SuccessMessage to satisfy [pflag.Value] interface
func (f *SuccessMessage) Type() string {
	return "SuccessMessage"
}

type TaskTimeOverRange struct {

	// Wire name: 'entries'
	Entries []TaskTimeOverRangeEntry `json:"entries,omitempty"`
	// interval length for all entries (difference in start time and end time of
	// an entry range) the same for all entries start time of first interval is
	// query_start_time_ms
	// Wire name: 'interval'
	Interval int64 `json:"interval,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TaskTimeOverRange) EncodeValues(key string, v *url.Values) error {
	pb, err := taskTimeOverRangeToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TaskTimeOverRange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskTimeOverRangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskTimeOverRangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskTimeOverRange) MarshalJSON() ([]byte, error) {
	pb, err := taskTimeOverRangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TaskTimeOverRangeEntry struct {
	// total task completion time in this time range, aggregated over all stages
	// and jobs in the query
	// Wire name: 'task_completed_time_ms'
	TaskCompletedTimeMs int64 `json:"task_completed_time_ms,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TaskTimeOverRangeEntry) EncodeValues(key string, v *url.Values) error {
	pb, err := taskTimeOverRangeEntryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TaskTimeOverRangeEntry) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &taskTimeOverRangeEntryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := taskTimeOverRangeEntryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TaskTimeOverRangeEntry) MarshalJSON() ([]byte, error) {
	pb, err := taskTimeOverRangeEntryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

func (st TerminationReason) EncodeValues(key string, v *url.Values) error {
	pb, err := terminationReasonToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// status code indicating why the cluster was terminated
type TerminationReasonCode string

const TerminationReasonCodeAbuseDetected TerminationReasonCode = `ABUSE_DETECTED`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCode = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCode = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCode = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCode = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCode = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCode = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCode = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCode = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCode = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCode = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCode = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCode = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCode = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCode = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCode = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCode = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCode = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCode = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCode = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCode = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCode = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCode = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCode = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeCommunicationLost TerminationReasonCode = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCode = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCode = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCode = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCode = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCode = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDriverUnreachable TerminationReasonCode = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCode = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCode = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCode = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCode = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCode = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCode = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCode = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCode = `INACTIVITY`

const TerminationReasonCodeInitScriptFailure TerminationReasonCode = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCode = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCode = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInternalError TerminationReasonCode = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCode = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCode = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCode = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCode = `JOB_FINISHED`

const TerminationReasonCodeK8sAutoscalingFailure TerminationReasonCode = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeK8sDbrClusterLaunchTimeout TerminationReasonCode = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCode = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCode = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCode = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCode = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCode = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCode = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodeRequestRejected TerminationReasonCode = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCode = `REQUEST_THROTTLED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCode = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCode = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCode = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCode = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCode = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCode = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCode = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCode = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCode = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCode = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCode = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCode = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCode = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCode = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCode = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnknown TerminationReasonCode = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCode = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCode = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUserRequest TerminationReasonCode = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCode = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCode = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCode = `WORKSPACE_CONFIGURATION_ERROR`

// String representation for [fmt.Print]
func (f *TerminationReasonCode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TerminationReasonCode) Set(v string) error {
	switch v {
	case `ABUSE_DETECTED`, `ATTACH_PROJECT_FAILURE`, `AWS_AUTHORIZATION_FAILURE`, `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`, `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`, `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`, `AWS_REQUEST_LIMIT_EXCEEDED`, `AWS_UNSUPPORTED_FAILURE`, `AZURE_BYOK_KEY_PERMISSION_FAILURE`, `AZURE_EPHEMERAL_DISK_FAILURE`, `AZURE_INVALID_DEPLOYMENT_TEMPLATE`, `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`, `AZURE_QUOTA_EXCEEDED_EXCEPTION`, `AZURE_RESOURCE_MANAGER_THROTTLING`, `AZURE_RESOURCE_PROVIDER_THROTTLING`, `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`, `AZURE_VM_EXTENSION_FAILURE`, `AZURE_VNET_CONFIGURATION_FAILURE`, `BOOTSTRAP_TIMEOUT`, `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`, `CLOUD_PROVIDER_DISK_SETUP_FAILURE`, `CLOUD_PROVIDER_LAUNCH_FAILURE`, `CLOUD_PROVIDER_RESOURCE_STOCKOUT`, `CLOUD_PROVIDER_SHUTDOWN`, `COMMUNICATION_LOST`, `CONTAINER_LAUNCH_FAILURE`, `CONTROL_PLANE_REQUEST_FAILURE`, `DATABASE_CONNECTION_FAILURE`, `DBFS_COMPONENT_UNHEALTHY`, `DOCKER_IMAGE_PULL_FAILURE`, `DRIVER_UNREACHABLE`, `DRIVER_UNRESPONSIVE`, `EXECUTION_COMPONENT_UNHEALTHY`, `GCP_QUOTA_EXCEEDED`, `GCP_SERVICE_ACCOUNT_DELETED`, `GLOBAL_INIT_SCRIPT_FAILURE`, `HIVE_METASTORE_PROVISIONING_FAILURE`, `IMAGE_PULL_PERMISSION_DENIED`, `INACTIVITY`, `INIT_SCRIPT_FAILURE`, `INSTANCE_POOL_CLUSTER_FAILURE`, `INSTANCE_UNREACHABLE`, `INTERNAL_ERROR`, `INVALID_ARGUMENT`, `INVALID_SPARK_IMAGE`, `IP_EXHAUSTION_FAILURE`, `JOB_FINISHED`, `K8S_AUTOSCALING_FAILURE`, `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`, `METASTORE_COMPONENT_UNHEALTHY`, `NEPHOS_RESOURCE_MANAGEMENT`, `NETWORK_CONFIGURATION_FAILURE`, `NFS_MOUNT_FAILURE`, `NPIP_TUNNEL_SETUP_FAILURE`, `NPIP_TUNNEL_TOKEN_FAILURE`, `REQUEST_REJECTED`, `REQUEST_THROTTLED`, `SECRET_RESOLUTION_ERROR`, `SECURITY_DAEMON_REGISTRATION_EXCEPTION`, `SELF_BOOTSTRAP_FAILURE`, `SKIPPED_SLOW_NODES`, `SLOW_IMAGE_DOWNLOAD`, `SPARK_ERROR`, `SPARK_IMAGE_DOWNLOAD_FAILURE`, `SPARK_STARTUP_FAILURE`, `SPOT_INSTANCE_TERMINATION`, `STORAGE_DOWNLOAD_FAILURE`, `STS_CLIENT_SETUP_FAILURE`, `SUBNET_EXHAUSTED_FAILURE`, `TEMPORARILY_UNAVAILABLE`, `TRIAL_EXPIRED`, `UNEXPECTED_LAUNCH_FAILURE`, `UNKNOWN`, `UNSUPPORTED_INSTANCE_TYPE`, `UPDATE_INSTANCE_PROFILE_FAILURE`, `USER_REQUEST`, `WORKER_SETUP_FAILURE`, `WORKSPACE_CANCELLED_ERROR`, `WORKSPACE_CONFIGURATION_ERROR`:
		*f = TerminationReasonCode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ABUSE_DETECTED", "ATTACH_PROJECT_FAILURE", "AWS_AUTHORIZATION_FAILURE", "AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE", "AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE", "AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE", "AWS_REQUEST_LIMIT_EXCEEDED", "AWS_UNSUPPORTED_FAILURE", "AZURE_BYOK_KEY_PERMISSION_FAILURE", "AZURE_EPHEMERAL_DISK_FAILURE", "AZURE_INVALID_DEPLOYMENT_TEMPLATE", "AZURE_OPERATION_NOT_ALLOWED_EXCEPTION", "AZURE_QUOTA_EXCEEDED_EXCEPTION", "AZURE_RESOURCE_MANAGER_THROTTLING", "AZURE_RESOURCE_PROVIDER_THROTTLING", "AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE", "AZURE_VM_EXTENSION_FAILURE", "AZURE_VNET_CONFIGURATION_FAILURE", "BOOTSTRAP_TIMEOUT", "BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION", "CLOUD_PROVIDER_DISK_SETUP_FAILURE", "CLOUD_PROVIDER_LAUNCH_FAILURE", "CLOUD_PROVIDER_RESOURCE_STOCKOUT", "CLOUD_PROVIDER_SHUTDOWN", "COMMUNICATION_LOST", "CONTAINER_LAUNCH_FAILURE", "CONTROL_PLANE_REQUEST_FAILURE", "DATABASE_CONNECTION_FAILURE", "DBFS_COMPONENT_UNHEALTHY", "DOCKER_IMAGE_PULL_FAILURE", "DRIVER_UNREACHABLE", "DRIVER_UNRESPONSIVE", "EXECUTION_COMPONENT_UNHEALTHY", "GCP_QUOTA_EXCEEDED", "GCP_SERVICE_ACCOUNT_DELETED", "GLOBAL_INIT_SCRIPT_FAILURE", "HIVE_METASTORE_PROVISIONING_FAILURE", "IMAGE_PULL_PERMISSION_DENIED", "INACTIVITY", "INIT_SCRIPT_FAILURE", "INSTANCE_POOL_CLUSTER_FAILURE", "INSTANCE_UNREACHABLE", "INTERNAL_ERROR", "INVALID_ARGUMENT", "INVALID_SPARK_IMAGE", "IP_EXHAUSTION_FAILURE", "JOB_FINISHED", "K8S_AUTOSCALING_FAILURE", "K8S_DBR_CLUSTER_LAUNCH_TIMEOUT", "METASTORE_COMPONENT_UNHEALTHY", "NEPHOS_RESOURCE_MANAGEMENT", "NETWORK_CONFIGURATION_FAILURE", "NFS_MOUNT_FAILURE", "NPIP_TUNNEL_SETUP_FAILURE", "NPIP_TUNNEL_TOKEN_FAILURE", "REQUEST_REJECTED", "REQUEST_THROTTLED", "SECRET_RESOLUTION_ERROR", "SECURITY_DAEMON_REGISTRATION_EXCEPTION", "SELF_BOOTSTRAP_FAILURE", "SKIPPED_SLOW_NODES", "SLOW_IMAGE_DOWNLOAD", "SPARK_ERROR", "SPARK_IMAGE_DOWNLOAD_FAILURE", "SPARK_STARTUP_FAILURE", "SPOT_INSTANCE_TERMINATION", "STORAGE_DOWNLOAD_FAILURE", "STS_CLIENT_SETUP_FAILURE", "SUBNET_EXHAUSTED_FAILURE", "TEMPORARILY_UNAVAILABLE", "TRIAL_EXPIRED", "UNEXPECTED_LAUNCH_FAILURE", "UNKNOWN", "UNSUPPORTED_INSTANCE_TYPE", "UPDATE_INSTANCE_PROFILE_FAILURE", "USER_REQUEST", "WORKER_SETUP_FAILURE", "WORKSPACE_CANCELLED_ERROR", "WORKSPACE_CONFIGURATION_ERROR"`, v)
	}
}

// Values returns all possible values for TerminationReasonCode.
//
// There is no guarantee on the order of the values in the slice.
func (f *TerminationReasonCode) Values() []TerminationReasonCode {
	return []TerminationReasonCode{
		TerminationReasonCodeAbuseDetected,
		TerminationReasonCodeAttachProjectFailure,
		TerminationReasonCodeAwsAuthorizationFailure,
		TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure,
		TerminationReasonCodeAwsInsufficientInstanceCapacityFailure,
		TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure,
		TerminationReasonCodeAwsRequestLimitExceeded,
		TerminationReasonCodeAwsUnsupportedFailure,
		TerminationReasonCodeAzureByokKeyPermissionFailure,
		TerminationReasonCodeAzureEphemeralDiskFailure,
		TerminationReasonCodeAzureInvalidDeploymentTemplate,
		TerminationReasonCodeAzureOperationNotAllowedException,
		TerminationReasonCodeAzureQuotaExceededException,
		TerminationReasonCodeAzureResourceManagerThrottling,
		TerminationReasonCodeAzureResourceProviderThrottling,
		TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure,
		TerminationReasonCodeAzureVmExtensionFailure,
		TerminationReasonCodeAzureVnetConfigurationFailure,
		TerminationReasonCodeBootstrapTimeout,
		TerminationReasonCodeBootstrapTimeoutCloudProviderException,
		TerminationReasonCodeCloudProviderDiskSetupFailure,
		TerminationReasonCodeCloudProviderLaunchFailure,
		TerminationReasonCodeCloudProviderResourceStockout,
		TerminationReasonCodeCloudProviderShutdown,
		TerminationReasonCodeCommunicationLost,
		TerminationReasonCodeContainerLaunchFailure,
		TerminationReasonCodeControlPlaneRequestFailure,
		TerminationReasonCodeDatabaseConnectionFailure,
		TerminationReasonCodeDbfsComponentUnhealthy,
		TerminationReasonCodeDockerImagePullFailure,
		TerminationReasonCodeDriverUnreachable,
		TerminationReasonCodeDriverUnresponsive,
		TerminationReasonCodeExecutionComponentUnhealthy,
		TerminationReasonCodeGcpQuotaExceeded,
		TerminationReasonCodeGcpServiceAccountDeleted,
		TerminationReasonCodeGlobalInitScriptFailure,
		TerminationReasonCodeHiveMetastoreProvisioningFailure,
		TerminationReasonCodeImagePullPermissionDenied,
		TerminationReasonCodeInactivity,
		TerminationReasonCodeInitScriptFailure,
		TerminationReasonCodeInstancePoolClusterFailure,
		TerminationReasonCodeInstanceUnreachable,
		TerminationReasonCodeInternalError,
		TerminationReasonCodeInvalidArgument,
		TerminationReasonCodeInvalidSparkImage,
		TerminationReasonCodeIpExhaustionFailure,
		TerminationReasonCodeJobFinished,
		TerminationReasonCodeK8sAutoscalingFailure,
		TerminationReasonCodeK8sDbrClusterLaunchTimeout,
		TerminationReasonCodeMetastoreComponentUnhealthy,
		TerminationReasonCodeNephosResourceManagement,
		TerminationReasonCodeNetworkConfigurationFailure,
		TerminationReasonCodeNfsMountFailure,
		TerminationReasonCodeNpipTunnelSetupFailure,
		TerminationReasonCodeNpipTunnelTokenFailure,
		TerminationReasonCodeRequestRejected,
		TerminationReasonCodeRequestThrottled,
		TerminationReasonCodeSecretResolutionError,
		TerminationReasonCodeSecurityDaemonRegistrationException,
		TerminationReasonCodeSelfBootstrapFailure,
		TerminationReasonCodeSkippedSlowNodes,
		TerminationReasonCodeSlowImageDownload,
		TerminationReasonCodeSparkError,
		TerminationReasonCodeSparkImageDownloadFailure,
		TerminationReasonCodeSparkStartupFailure,
		TerminationReasonCodeSpotInstanceTermination,
		TerminationReasonCodeStorageDownloadFailure,
		TerminationReasonCodeStsClientSetupFailure,
		TerminationReasonCodeSubnetExhaustedFailure,
		TerminationReasonCodeTemporarilyUnavailable,
		TerminationReasonCodeTrialExpired,
		TerminationReasonCodeUnexpectedLaunchFailure,
		TerminationReasonCodeUnknown,
		TerminationReasonCodeUnsupportedInstanceType,
		TerminationReasonCodeUpdateInstanceProfileFailure,
		TerminationReasonCodeUserRequest,
		TerminationReasonCodeWorkerSetupFailure,
		TerminationReasonCodeWorkspaceCancelledError,
		TerminationReasonCodeWorkspaceConfigurationError,
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

type TextValue struct {

	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TextValue) EncodeValues(key string, v *url.Values) error {
	pb, err := textValueToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TextValue) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &textValuePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := textValueFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TextValue) MarshalJSON() ([]byte, error) {
	pb, err := textValueToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TimeRange struct {
	// The end time in milliseconds.
	// Wire name: 'end_time_ms'
	EndTimeMs int64 `json:"end_time_ms,omitempty"`
	// The start time in milliseconds.
	// Wire name: 'start_time_ms'
	StartTimeMs int64 `json:"start_time_ms,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TimeRange) EncodeValues(key string, v *url.Values) error {
	pb, err := timeRangeToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TimeRange) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &timeRangePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := timeRangeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TimeRange) MarshalJSON() ([]byte, error) {
	pb, err := timeRangeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	// Wire name: 'new_owner'
	NewOwner string `json:"new_owner,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TransferOwnershipObjectId) EncodeValues(key string, v *url.Values) error {
	pb, err := transferOwnershipObjectIdToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TransferOwnershipObjectId) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transferOwnershipObjectIdPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transferOwnershipObjectIdFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransferOwnershipObjectId) MarshalJSON() ([]byte, error) {
	pb, err := transferOwnershipObjectIdToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	// Wire name: 'new_owner'
	NewOwner string `json:"new_owner,omitempty"`
	// The ID of the object on which to change ownership.
	ObjectId TransferOwnershipObjectId `json:"-" tf:"-"`
	// The type of object on which to change ownership.
	ObjectType OwnableObjectType `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TransferOwnershipRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := transferOwnershipRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TransferOwnershipRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transferOwnershipRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transferOwnershipRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransferOwnershipRequest) MarshalJSON() ([]byte, error) {
	pb, err := transferOwnershipRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TrashAlertRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st TrashAlertRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := trashAlertRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TrashAlertRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &trashAlertRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := trashAlertRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TrashAlertRequest) MarshalJSON() ([]byte, error) {
	pb, err := trashAlertRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TrashAlertV2Request struct {
	Id string `json:"-" tf:"-"`
}

func (st TrashAlertV2Request) EncodeValues(key string, v *url.Values) error {
	pb, err := trashAlertV2RequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TrashAlertV2Request) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &trashAlertV2RequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := trashAlertV2RequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TrashAlertV2Request) MarshalJSON() ([]byte, error) {
	pb, err := trashAlertV2RequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TrashQueryRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st TrashQueryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := trashQueryRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TrashQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &trashQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := trashQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TrashQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := trashQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateAlertRequest struct {

	// Wire name: 'alert'
	Alert *UpdateAlertRequestAlert `json:"alert,omitempty"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`

	Id string `json:"-" tf:"-"`
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

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateAlertRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateAlertRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateAlertRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAlertRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAlertRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAlertRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateAlertRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition `json:"condition,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateAlertRequestAlert) EncodeValues(key string, v *url.Values) error {
	pb, err := updateAlertRequestAlertToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateAlertRequestAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAlertRequestAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAlertRequestAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAlertRequestAlert) MarshalJSON() ([]byte, error) {
	pb, err := updateAlertRequestAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateAlertV2Request struct {

	// Wire name: 'alert'
	Alert AlertV2 `json:"alert"`
	// UUID identifying the alert.
	Id string `json:"-" tf:"-"`
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
	UpdateMask string `json:"-" tf:"-"`
}

func (st UpdateAlertV2Request) EncodeValues(key string, v *url.Values) error {
	pb, err := updateAlertV2RequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateAlertV2Request) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAlertV2RequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAlertV2RequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAlertV2Request) MarshalJSON() ([]byte, error) {
	pb, err := updateAlertV2RequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateQueryRequest struct {
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`

	Id string `json:"-" tf:"-"`

	// Wire name: 'query'
	Query *UpdateQueryRequestQuery `json:"query,omitempty"`
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

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateQueryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateQueryRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateQueryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateQueryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateQueryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateQueryRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateQueryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string `json:"catalog,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter `json:"parameters,omitempty"`
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string `json:"schema,omitempty"`

	// Wire name: 'tags'
	Tags []string `json:"tags,omitempty"`
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateQueryRequestQuery) EncodeValues(key string, v *url.Values) error {
	pb, err := updateQueryRequestQueryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateQueryRequestQuery) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateQueryRequestQueryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateQueryRequestQueryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateQueryRequestQuery) MarshalJSON() ([]byte, error) {
	pb, err := updateQueryRequestQueryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateVisualizationRequest struct {
	Id string `json:"-" tf:"-"`
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

	// Wire name: 'visualization'
	Visualization *UpdateVisualizationRequestVisualization `json:"visualization,omitempty"`
}

func (st UpdateVisualizationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateVisualizationRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateVisualizationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateVisualizationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateVisualizationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateVisualizationRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateVisualizationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateVisualizationRequestVisualization) EncodeValues(key string, v *url.Values) error {
	pb, err := updateVisualizationRequestVisualizationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateVisualizationRequestVisualization) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateVisualizationRequestVisualizationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateVisualizationRequestVisualizationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateVisualizationRequestVisualization) MarshalJSON() ([]byte, error) {
	pb, err := updateVisualizationRequestVisualizationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateWidgetRequest struct {
	// Dashboard ID returned by :method:dashboards/create.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id"`
	// Widget ID returned by :method:dashboardwidgets/create
	Id string `json:"-" tf:"-"`

	// Wire name: 'options'
	Options WidgetOptions `json:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	// Wire name: 'text'
	Text string `json:"text,omitempty"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	// Wire name: 'visualization_id'
	VisualizationId string `json:"visualization_id,omitempty"`
	// Width of a widget
	// Wire name: 'width'
	Width int `json:"width"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateWidgetRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateWidgetRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateWidgetRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWidgetRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWidgetRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWidgetRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateWidgetRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type User struct {

	// Wire name: 'email'
	Email string `json:"email,omitempty"`

	// Wire name: 'id'
	Id int `json:"id,omitempty"`

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st User) EncodeValues(key string, v *url.Values) error {
	pb, err := userToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *User) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &userPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := userFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st User) MarshalJSON() ([]byte, error) {
	pb, err := userToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Visualization struct {
	// The timestamp indicating when the visualization was created.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// The display name of the visualization.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the visualization.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// UUID of the query that the visualization is attached to.
	// Wire name: 'query_id'
	QueryId string `json:"query_id,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type string `json:"type,omitempty"`
	// The timestamp indicating when the visualization was updated.
	// Wire name: 'update_time'
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Visualization) EncodeValues(key string, v *url.Values) error {
	pb, err := visualizationToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Visualization) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &visualizationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := visualizationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Visualization) MarshalJSON() ([]byte, error) {
	pb, err := visualizationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WarehouseAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WarehouseAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := warehouseAccessControlRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WarehouseAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &warehouseAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := warehouseAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WarehouseAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := warehouseAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []WarehousePermission `json:"all_permissions,omitempty"`
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

func (st WarehouseAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := warehouseAccessControlResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WarehouseAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &warehouseAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := warehouseAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WarehouseAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := warehouseAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WarehousePermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WarehousePermission) EncodeValues(key string, v *url.Values) error {
	pb, err := warehousePermissionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WarehousePermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &warehousePermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := warehousePermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WarehousePermission) MarshalJSON() ([]byte, error) {
	pb, err := warehousePermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Permission level
type WarehousePermissionLevel string

const WarehousePermissionLevelCanManage WarehousePermissionLevel = `CAN_MANAGE`

const WarehousePermissionLevelCanMonitor WarehousePermissionLevel = `CAN_MONITOR`

const WarehousePermissionLevelCanUse WarehousePermissionLevel = `CAN_USE`

const WarehousePermissionLevelCanView WarehousePermissionLevel = `CAN_VIEW`

const WarehousePermissionLevelIsOwner WarehousePermissionLevel = `IS_OWNER`

// String representation for [fmt.Print]
func (f *WarehousePermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarehousePermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_MONITOR`, `CAN_USE`, `CAN_VIEW`, `IS_OWNER`:
		*f = WarehousePermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_MONITOR", "CAN_USE", "CAN_VIEW", "IS_OWNER"`, v)
	}
}

// Values returns all possible values for WarehousePermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *WarehousePermissionLevel) Values() []WarehousePermissionLevel {
	return []WarehousePermissionLevel{
		WarehousePermissionLevelCanManage,
		WarehousePermissionLevelCanMonitor,
		WarehousePermissionLevelCanUse,
		WarehousePermissionLevelCanView,
		WarehousePermissionLevelIsOwner,
	}
}

// Type always returns WarehousePermissionLevel to satisfy [pflag.Value] interface
func (f *WarehousePermissionLevel) Type() string {
	return "WarehousePermissionLevel"
}

type WarehousePermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []WarehouseAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WarehousePermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := warehousePermissionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WarehousePermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &warehousePermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := warehousePermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WarehousePermissions) MarshalJSON() ([]byte, error) {
	pb, err := warehousePermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WarehousePermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WarehousePermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := warehousePermissionsDescriptionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WarehousePermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &warehousePermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := warehousePermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WarehousePermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := warehousePermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WarehousePermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []WarehouseAccessControlRequest `json:"access_control_list,omitempty"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" tf:"-"`
}

func (st WarehousePermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := warehousePermissionsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WarehousePermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &warehousePermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := warehousePermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WarehousePermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := warehousePermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// Warehouse type: `PRO` or `CLASSIC`.
	// Wire name: 'warehouse_type'
	WarehouseType WarehouseTypePairWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WarehouseTypePair) EncodeValues(key string, v *url.Values) error {
	pb, err := warehouseTypePairToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WarehouseTypePair) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &warehouseTypePairPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := warehouseTypePairFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WarehouseTypePair) MarshalJSON() ([]byte, error) {
	pb, err := warehouseTypePairToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Warehouse type: `PRO` or `CLASSIC`.
type WarehouseTypePairWarehouseType string

const WarehouseTypePairWarehouseTypeClassic WarehouseTypePairWarehouseType = `CLASSIC`

const WarehouseTypePairWarehouseTypePro WarehouseTypePairWarehouseType = `PRO`

const WarehouseTypePairWarehouseTypeTypeUnspecified WarehouseTypePairWarehouseType = `TYPE_UNSPECIFIED`

// String representation for [fmt.Print]
func (f *WarehouseTypePairWarehouseType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *WarehouseTypePairWarehouseType) Set(v string) error {
	switch v {
	case `CLASSIC`, `PRO`, `TYPE_UNSPECIFIED`:
		*f = WarehouseTypePairWarehouseType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CLASSIC", "PRO", "TYPE_UNSPECIFIED"`, v)
	}
}

// Values returns all possible values for WarehouseTypePairWarehouseType.
//
// There is no guarantee on the order of the values in the slice.
func (f *WarehouseTypePairWarehouseType) Values() []WarehouseTypePairWarehouseType {
	return []WarehouseTypePairWarehouseType{
		WarehouseTypePairWarehouseTypeClassic,
		WarehouseTypePairWarehouseTypePro,
		WarehouseTypePairWarehouseTypeTypeUnspecified,
	}
}

// Type always returns WarehouseTypePairWarehouseType to satisfy [pflag.Value] interface
func (f *WarehouseTypePairWarehouseType) Type() string {
	return "WarehouseTypePairWarehouseType"
}

type Widget struct {
	// The unique ID for this widget.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	// Wire name: 'options'
	Options *WidgetOptions `json:"options,omitempty"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	// Wire name: 'visualization'
	Visualization *LegacyVisualization `json:"visualization,omitempty"`
	// Unused field.
	// Wire name: 'width'
	Width int `json:"width,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Widget) EncodeValues(key string, v *url.Values) error {
	pb, err := widgetToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Widget) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &widgetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := widgetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Widget) MarshalJSON() ([]byte, error) {
	pb, err := widgetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WidgetOptions struct {
	// Timestamp when this object was created
	// Wire name: 'created_at'
	CreatedAt string `json:"created_at,omitempty"`
	// Custom description of the widget
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Whether this widget is hidden on the dashboard.
	// Wire name: 'isHidden'
	IsHidden bool `json:"isHidden,omitempty"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	// Wire name: 'parameterMappings'
	ParameterMappings any `json:"parameterMappings,omitempty"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	// Wire name: 'position'
	Position *WidgetPosition `json:"position,omitempty"`
	// Custom title of the widget
	// Wire name: 'title'
	Title string `json:"title,omitempty"`
	// Timestamp of the last time this object was updated.
	// Wire name: 'updated_at'
	UpdatedAt string `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WidgetOptions) EncodeValues(key string, v *url.Values) error {
	pb, err := widgetOptionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WidgetOptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &widgetOptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := widgetOptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WidgetOptions) MarshalJSON() ([]byte, error) {
	pb, err := widgetOptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	// Wire name: 'autoHeight'
	AutoHeight bool `json:"autoHeight,omitempty"`
	// column in the dashboard grid. Values start with 0
	// Wire name: 'col'
	Col int `json:"col,omitempty"`
	// row in the dashboard grid. Values start with 0
	// Wire name: 'row'
	Row int `json:"row,omitempty"`
	// width of the widget measured in dashboard grid cells
	// Wire name: 'sizeX'
	SizeX int `json:"sizeX,omitempty"`
	// height of the widget measured in dashboard grid cells
	// Wire name: 'sizeY'
	SizeY int `json:"sizeY,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WidgetPosition) EncodeValues(key string, v *url.Values) error {
	pb, err := widgetPositionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *WidgetPosition) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &widgetPositionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := widgetPositionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WidgetPosition) MarshalJSON() ([]byte, error) {
	pb, err := widgetPositionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
