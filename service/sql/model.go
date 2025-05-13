// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccessControl struct {

	// Wire name: 'group_name'
	GroupName string
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel

	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func accessControlToPb(st *AccessControl) (*accessControlPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accessControlPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type accessControlPb struct {
	GroupName string `json:"group_name,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`

	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func accessControlFromPb(pb *accessControlPb) (*AccessControl, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccessControl{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *accessControlPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st accessControlPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Aggregation string
type aggregationPb string

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

// Type always returns Aggregation to satisfy [pflag.Value] interface
func (f *Aggregation) Type() string {
	return "Aggregation"
}

func aggregationToPb(st *Aggregation) (*aggregationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := aggregationPb(*st)
	return &pb, nil
}

func aggregationFromPb(pb *aggregationPb) (*Aggregation, error) {
	if pb == nil {
		return nil, nil
	}
	st := Aggregation(*pb)
	return &st, nil
}

type Alert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime string
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string
	// The workspace state of the alert. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// The workspace path of the folder containing the alert.
	// Wire name: 'parent_path'
	ParentPath string
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	// Wire name: 'state'
	State AlertState
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	// Wire name: 'trigger_time'
	TriggerTime string
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime string

	ForceSendFields []string `tf:"-"`
}

func alertToPb(st *Alert) (*alertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertPb{}
	conditionPb, err := alertConditionToPb(st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = conditionPb
	}

	pb.CreateTime = st.CreateTime

	pb.CustomBody = st.CustomBody

	pb.CustomSubject = st.CustomSubject

	pb.DisplayName = st.DisplayName

	pb.Id = st.Id

	pb.LifecycleState = st.LifecycleState

	pb.NotifyOnOk = st.NotifyOnOk

	pb.OwnerUserName = st.OwnerUserName

	pb.ParentPath = st.ParentPath

	pb.QueryId = st.QueryId

	pb.SecondsToRetrigger = st.SecondsToRetrigger

	pb.State = st.State

	pb.TriggerTime = st.TriggerTime

	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertPb struct {
	// Trigger conditions of the alert.
	Condition *alertConditionPb `json:"condition,omitempty"`
	// The timestamp indicating when the alert was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the alert.
	Id string `json:"id,omitempty"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// The workspace path of the folder containing the alert.
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the query attached to the alert.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State AlertState `json:"state,omitempty"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime string `json:"trigger_time,omitempty"`
	// The timestamp indicating when the alert was updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertFromPb(pb *alertPb) (*Alert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Alert{}
	conditionField, err := alertConditionFromPb(pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = conditionField
	}
	st.CreateTime = pb.CreateTime
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.LifecycleState = pb.LifecycleState
	st.NotifyOnOk = pb.NotifyOnOk
	st.OwnerUserName = pb.OwnerUserName
	st.ParentPath = pb.ParentPath
	st.QueryId = pb.QueryId
	st.SecondsToRetrigger = pb.SecondsToRetrigger
	st.State = pb.State
	st.TriggerTime = pb.TriggerTime
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertCondition struct {
	// Alert state if result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertState
	// Operator used for comparison in alert evaluation.
	// Wire name: 'op'
	Op AlertOperator
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	// Wire name: 'operand'
	Operand *AlertConditionOperand
	// Threshold value used for comparison in alert evaluation.
	// Wire name: 'threshold'
	Threshold *AlertConditionThreshold
}

func alertConditionToPb(st *AlertCondition) (*alertConditionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConditionPb{}
	pb.EmptyResultState = st.EmptyResultState

	pb.Op = st.Op

	operandPb, err := alertConditionOperandToPb(st.Operand)
	if err != nil {
		return nil, err
	}
	if operandPb != nil {
		pb.Operand = operandPb
	}

	thresholdPb, err := alertConditionThresholdToPb(st.Threshold)
	if err != nil {
		return nil, err
	}
	if thresholdPb != nil {
		pb.Threshold = thresholdPb
	}

	return pb, nil
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

type alertConditionPb struct {
	// Alert state if result is empty.
	EmptyResultState AlertState `json:"empty_result_state,omitempty"`
	// Operator used for comparison in alert evaluation.
	Op AlertOperator `json:"op,omitempty"`
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	Operand *alertConditionOperandPb `json:"operand,omitempty"`
	// Threshold value used for comparison in alert evaluation.
	Threshold *alertConditionThresholdPb `json:"threshold,omitempty"`
}

func alertConditionFromPb(pb *alertConditionPb) (*AlertCondition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertCondition{}
	st.EmptyResultState = pb.EmptyResultState
	st.Op = pb.Op
	operandField, err := alertConditionOperandFromPb(pb.Operand)
	if err != nil {
		return nil, err
	}
	if operandField != nil {
		st.Operand = operandField
	}
	thresholdField, err := alertConditionThresholdFromPb(pb.Threshold)
	if err != nil {
		return nil, err
	}
	if thresholdField != nil {
		st.Threshold = thresholdField
	}

	return st, nil
}

type AlertConditionOperand struct {

	// Wire name: 'column'
	Column *AlertOperandColumn
}

func alertConditionOperandToPb(st *AlertConditionOperand) (*alertConditionOperandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConditionOperandPb{}
	columnPb, err := alertOperandColumnToPb(st.Column)
	if err != nil {
		return nil, err
	}
	if columnPb != nil {
		pb.Column = columnPb
	}

	return pb, nil
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

type alertConditionOperandPb struct {
	Column *alertOperandColumnPb `json:"column,omitempty"`
}

func alertConditionOperandFromPb(pb *alertConditionOperandPb) (*AlertConditionOperand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConditionOperand{}
	columnField, err := alertOperandColumnFromPb(pb.Column)
	if err != nil {
		return nil, err
	}
	if columnField != nil {
		st.Column = columnField
	}

	return st, nil
}

type AlertConditionThreshold struct {

	// Wire name: 'value'
	Value *AlertOperandValue
}

func alertConditionThresholdToPb(st *AlertConditionThreshold) (*alertConditionThresholdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConditionThresholdPb{}
	valuePb, err := alertOperandValueToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	return pb, nil
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

type alertConditionThresholdPb struct {
	Value *alertOperandValuePb `json:"value,omitempty"`
}

func alertConditionThresholdFromPb(pb *alertConditionThresholdPb) (*AlertConditionThreshold, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConditionThreshold{}
	valueField, err := alertOperandValueFromPb(pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = valueField
	}

	return st, nil
}

// UNSPECIFIED - default unspecify value for proto enum, do not use it in the
// code UNKNOWN - alert not yet evaluated TRIGGERED - alert is triggered OK -
// alert is not triggered ERROR - alert evaluation failed
type AlertEvaluationState string
type alertEvaluationStatePb string

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

// Type always returns AlertEvaluationState to satisfy [pflag.Value] interface
func (f *AlertEvaluationState) Type() string {
	return "AlertEvaluationState"
}

func alertEvaluationStateToPb(st *AlertEvaluationState) (*alertEvaluationStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := alertEvaluationStatePb(*st)
	return &pb, nil
}

func alertEvaluationStateFromPb(pb *alertEvaluationStatePb) (*AlertEvaluationState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertEvaluationState(*pb)
	return &st, nil
}

type AlertOperandColumn struct {

	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func alertOperandColumnToPb(st *AlertOperandColumn) (*alertOperandColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertOperandColumnPb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertOperandColumnPb struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertOperandColumnFromPb(pb *alertOperandColumnPb) (*AlertOperandColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertOperandColumn{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertOperandColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertOperandColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertOperandValue struct {

	// Wire name: 'bool_value'
	BoolValue bool

	// Wire name: 'double_value'
	DoubleValue float64

	// Wire name: 'string_value'
	StringValue string

	ForceSendFields []string `tf:"-"`
}

func alertOperandValueToPb(st *AlertOperandValue) (*alertOperandValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertOperandValuePb{}
	pb.BoolValue = st.BoolValue

	pb.DoubleValue = st.DoubleValue

	pb.StringValue = st.StringValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertOperandValuePb struct {
	BoolValue bool `json:"bool_value,omitempty"`

	DoubleValue float64 `json:"double_value,omitempty"`

	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertOperandValueFromPb(pb *alertOperandValuePb) (*AlertOperandValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertOperandValue{}
	st.BoolValue = pb.BoolValue
	st.DoubleValue = pb.DoubleValue
	st.StringValue = pb.StringValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertOperandValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertOperandValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertOperator string
type alertOperatorPb string

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

// Type always returns AlertOperator to satisfy [pflag.Value] interface
func (f *AlertOperator) Type() string {
	return "AlertOperator"
}

func alertOperatorToPb(st *AlertOperator) (*alertOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := alertOperatorPb(*st)
	return &pb, nil
}

func alertOperatorFromPb(pb *alertOperatorPb) (*AlertOperator, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertOperator(*pb)
	return &st, nil
}

// Alert configuration options.
type AlertOptions struct {
	// Name of column in the query result to compare in alert evaluation.
	// Wire name: 'column'
	Column string
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string
	// State that alert evaluates to when query result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertOptionsEmptyResultState
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	// Wire name: 'muted'
	Muted bool
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	// Wire name: 'op'
	Op string
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	// Wire name: 'value'
	Value any

	ForceSendFields []string `tf:"-"`
}

func alertOptionsToPb(st *AlertOptions) (*alertOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertOptionsPb{}
	pb.Column = st.Column

	pb.CustomBody = st.CustomBody

	pb.CustomSubject = st.CustomSubject

	pb.EmptyResultState = st.EmptyResultState

	pb.Muted = st.Muted

	pb.Op = st.Op

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertOptionsPb struct {
	// Name of column in the query result to compare in alert evaluation.
	Column string `json:"column"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string `json:"custom_subject,omitempty"`
	// State that alert evaluates to when query result is empty.
	EmptyResultState AlertOptionsEmptyResultState `json:"empty_result_state,omitempty"`
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	Muted bool `json:"muted,omitempty"`
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op string `json:"op"`
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	Value any `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertOptionsFromPb(pb *alertOptionsPb) (*AlertOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertOptions{}
	st.Column = pb.Column
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	st.EmptyResultState = pb.EmptyResultState
	st.Muted = pb.Muted
	st.Op = pb.Op
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// State that alert evaluates to when query result is empty.
type AlertOptionsEmptyResultState string
type alertOptionsEmptyResultStatePb string

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

// Type always returns AlertOptionsEmptyResultState to satisfy [pflag.Value] interface
func (f *AlertOptionsEmptyResultState) Type() string {
	return "AlertOptionsEmptyResultState"
}

func alertOptionsEmptyResultStateToPb(st *AlertOptionsEmptyResultState) (*alertOptionsEmptyResultStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := alertOptionsEmptyResultStatePb(*st)
	return &pb, nil
}

func alertOptionsEmptyResultStateFromPb(pb *alertOptionsEmptyResultStatePb) (*AlertOptionsEmptyResultState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertOptionsEmptyResultState(*pb)
	return &st, nil
}

type AlertQuery struct {
	// The timestamp when this query was created.
	// Wire name: 'created_at'
	CreatedAt string
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// Query ID.
	// Wire name: 'id'
	Id string
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	// Wire name: 'is_draft'
	IsDraft bool
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	// Wire name: 'is_safe'
	IsSafe bool
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string

	// Wire name: 'options'
	Options *QueryOptions
	// The text of the query to be run.
	// Wire name: 'query'
	Query string

	// Wire name: 'tags'
	Tags []string
	// The timestamp at which this query was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string
	// The ID of the user who owns the query.
	// Wire name: 'user_id'
	UserId int

	ForceSendFields []string `tf:"-"`
}

func alertQueryToPb(st *AlertQuery) (*alertQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertQueryPb{}
	pb.CreatedAt = st.CreatedAt

	pb.DataSourceId = st.DataSourceId

	pb.Description = st.Description

	pb.Id = st.Id

	pb.IsArchived = st.IsArchived

	pb.IsDraft = st.IsDraft

	pb.IsSafe = st.IsSafe

	pb.Name = st.Name

	optionsPb, err := queryOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	pb.Query = st.Query

	pb.Tags = st.Tags

	pb.UpdatedAt = st.UpdatedAt

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertQueryPb struct {
	// The timestamp when this query was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Query ID.
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft bool `json:"is_draft,omitempty"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe bool `json:"is_safe,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`

	Options *queryOptionsPb `json:"options,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// The timestamp at which this query was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The ID of the user who owns the query.
	UserId int `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertQueryFromPb(pb *alertQueryPb) (*AlertQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertQuery{}
	st.CreatedAt = pb.CreatedAt
	st.DataSourceId = pb.DataSourceId
	st.Description = pb.Description
	st.Id = pb.Id
	st.IsArchived = pb.IsArchived
	st.IsDraft = pb.IsDraft
	st.IsSafe = pb.IsSafe
	st.Name = pb.Name
	optionsField, err := queryOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	st.Query = pb.Query
	st.Tags = pb.Tags
	st.UpdatedAt = pb.UpdatedAt
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertState string
type alertStatePb string

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

// Type always returns AlertState to satisfy [pflag.Value] interface
func (f *AlertState) Type() string {
	return "AlertState"
}

func alertStateToPb(st *AlertState) (*alertStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := alertStatePb(*st)
	return &pb, nil
}

func alertStateFromPb(pb *alertStatePb) (*AlertState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertState(*pb)
	return &st, nil
}

type AlertV2 struct {
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime string
	// Custom description for the alert. support mustache template.
	// Wire name: 'custom_description'
	CustomDescription string
	// Custom summary for the alert. support mustache template.
	// Wire name: 'custom_summary'
	CustomSummary string
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string

	// Wire name: 'evaluation'
	Evaluation *AlertV2Evaluation
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	// Wire name: 'parent_path'
	ParentPath string
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string
	// The run as username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'run_as_user_name'
	RunAsUserName string

	// Wire name: 'schedule'
	Schedule *CronSchedule
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime string
	// ID of the SQL warehouse attached to the alert.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func alertV2ToPb(st *AlertV2) (*alertV2Pb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2Pb{}
	pb.CreateTime = st.CreateTime

	pb.CustomDescription = st.CustomDescription

	pb.CustomSummary = st.CustomSummary

	pb.DisplayName = st.DisplayName

	evaluationPb, err := alertV2EvaluationToPb(st.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationPb != nil {
		pb.Evaluation = evaluationPb
	}

	pb.Id = st.Id

	pb.LifecycleState = st.LifecycleState

	pb.OwnerUserName = st.OwnerUserName

	pb.ParentPath = st.ParentPath

	pb.QueryText = st.QueryText

	pb.RunAsUserName = st.RunAsUserName

	schedulePb, err := cronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	pb.UpdateTime = st.UpdateTime

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertV2Pb struct {
	// The timestamp indicating when the alert was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom description for the alert. support mustache template.
	CustomDescription string `json:"custom_description,omitempty"`
	// Custom summary for the alert. support mustache template.
	CustomSummary string `json:"custom_summary,omitempty"`
	// The display name of the alert.
	DisplayName string `json:"display_name,omitempty"`

	Evaluation *alertV2EvaluationPb `json:"evaluation,omitempty"`
	// UUID identifying the alert.
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// The run as username. This field is set to "Unavailable" if the user has
	// been deleted.
	RunAsUserName string `json:"run_as_user_name,omitempty"`

	Schedule *cronSchedulePb `json:"schedule,omitempty"`
	// The timestamp indicating when the alert was updated.
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the alert.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertV2FromPb(pb *alertV2Pb) (*AlertV2, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2{}
	st.CreateTime = pb.CreateTime
	st.CustomDescription = pb.CustomDescription
	st.CustomSummary = pb.CustomSummary
	st.DisplayName = pb.DisplayName
	evaluationField, err := alertV2EvaluationFromPb(pb.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationField != nil {
		st.Evaluation = evaluationField
	}
	st.Id = pb.Id
	st.LifecycleState = pb.LifecycleState
	st.OwnerUserName = pb.OwnerUserName
	st.ParentPath = pb.ParentPath
	st.QueryText = pb.QueryText
	st.RunAsUserName = pb.RunAsUserName
	scheduleField, err := cronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.UpdateTime = pb.UpdateTime
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertV2Pb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2Pb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2Evaluation struct {
	// Operator used for comparison in alert evaluation.
	// Wire name: 'comparison_operator'
	ComparisonOperator ComparisonOperator
	// Alert state if result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertEvaluationState
	// Timestamp of the last evaluation.
	// Wire name: 'last_evaluated_at'
	LastEvaluatedAt string
	// User or Notification Destination to notify when alert is triggered.
	// Wire name: 'notification'
	Notification *AlertV2Notification
	// Source column from result to use to evaluate alert
	// Wire name: 'source'
	Source *AlertV2OperandColumn
	// Latest state of alert evaluation.
	// Wire name: 'state'
	State AlertEvaluationState
	// Threshold to user for alert evaluation, can be a column or a value.
	// Wire name: 'threshold'
	Threshold *AlertV2Operand

	ForceSendFields []string `tf:"-"`
}

func alertV2EvaluationToPb(st *AlertV2Evaluation) (*alertV2EvaluationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2EvaluationPb{}
	pb.ComparisonOperator = st.ComparisonOperator

	pb.EmptyResultState = st.EmptyResultState

	pb.LastEvaluatedAt = st.LastEvaluatedAt

	notificationPb, err := alertV2NotificationToPb(st.Notification)
	if err != nil {
		return nil, err
	}
	if notificationPb != nil {
		pb.Notification = notificationPb
	}

	sourcePb, err := alertV2OperandColumnToPb(st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = sourcePb
	}

	pb.State = st.State

	thresholdPb, err := alertV2OperandToPb(st.Threshold)
	if err != nil {
		return nil, err
	}
	if thresholdPb != nil {
		pb.Threshold = thresholdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertV2EvaluationPb struct {
	// Operator used for comparison in alert evaluation.
	ComparisonOperator ComparisonOperator `json:"comparison_operator,omitempty"`
	// Alert state if result is empty.
	EmptyResultState AlertEvaluationState `json:"empty_result_state,omitempty"`
	// Timestamp of the last evaluation.
	LastEvaluatedAt string `json:"last_evaluated_at,omitempty"`
	// User or Notification Destination to notify when alert is triggered.
	Notification *alertV2NotificationPb `json:"notification,omitempty"`
	// Source column from result to use to evaluate alert
	Source *alertV2OperandColumnPb `json:"source,omitempty"`
	// Latest state of alert evaluation.
	State AlertEvaluationState `json:"state,omitempty"`
	// Threshold to user for alert evaluation, can be a column or a value.
	Threshold *alertV2OperandPb `json:"threshold,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertV2EvaluationFromPb(pb *alertV2EvaluationPb) (*AlertV2Evaluation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Evaluation{}
	st.ComparisonOperator = pb.ComparisonOperator
	st.EmptyResultState = pb.EmptyResultState
	st.LastEvaluatedAt = pb.LastEvaluatedAt
	notificationField, err := alertV2NotificationFromPb(pb.Notification)
	if err != nil {
		return nil, err
	}
	if notificationField != nil {
		st.Notification = notificationField
	}
	sourceField, err := alertV2OperandColumnFromPb(pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = sourceField
	}
	st.State = pb.State
	thresholdField, err := alertV2OperandFromPb(pb.Threshold)
	if err != nil {
		return nil, err
	}
	if thresholdField != nil {
		st.Threshold = thresholdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertV2EvaluationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2EvaluationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2Notification struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'retrigger_seconds'
	RetriggerSeconds int

	// Wire name: 'subscriptions'
	Subscriptions []AlertV2Subscription

	ForceSendFields []string `tf:"-"`
}

func alertV2NotificationToPb(st *AlertV2Notification) (*alertV2NotificationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2NotificationPb{}
	pb.NotifyOnOk = st.NotifyOnOk

	pb.RetriggerSeconds = st.RetriggerSeconds

	var subscriptionsPb []alertV2SubscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := alertV2SubscriptionToPb(&item)
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

type alertV2NotificationPb struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	RetriggerSeconds int `json:"retrigger_seconds,omitempty"`

	Subscriptions []alertV2SubscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertV2NotificationFromPb(pb *alertV2NotificationPb) (*AlertV2Notification, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Notification{}
	st.NotifyOnOk = pb.NotifyOnOk
	st.RetriggerSeconds = pb.RetriggerSeconds

	var subscriptionsField []AlertV2Subscription
	for _, itemPb := range pb.Subscriptions {
		item, err := alertV2SubscriptionFromPb(&itemPb)
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

func (st *alertV2NotificationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2NotificationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2Operand struct {

	// Wire name: 'column'
	Column *AlertV2OperandColumn

	// Wire name: 'value'
	Value *AlertV2OperandValue
}

func alertV2OperandToPb(st *AlertV2Operand) (*alertV2OperandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2OperandPb{}
	columnPb, err := alertV2OperandColumnToPb(st.Column)
	if err != nil {
		return nil, err
	}
	if columnPb != nil {
		pb.Column = columnPb
	}

	valuePb, err := alertV2OperandValueToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	return pb, nil
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

type alertV2OperandPb struct {
	Column *alertV2OperandColumnPb `json:"column,omitempty"`

	Value *alertV2OperandValuePb `json:"value,omitempty"`
}

func alertV2OperandFromPb(pb *alertV2OperandPb) (*AlertV2Operand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Operand{}
	columnField, err := alertV2OperandColumnFromPb(pb.Column)
	if err != nil {
		return nil, err
	}
	if columnField != nil {
		st.Column = columnField
	}
	valueField, err := alertV2OperandValueFromPb(pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = valueField
	}

	return st, nil
}

type AlertV2OperandColumn struct {

	// Wire name: 'aggregation'
	Aggregation Aggregation

	// Wire name: 'display'
	Display string

	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func alertV2OperandColumnToPb(st *AlertV2OperandColumn) (*alertV2OperandColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2OperandColumnPb{}
	pb.Aggregation = st.Aggregation

	pb.Display = st.Display

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertV2OperandColumnPb struct {
	Aggregation Aggregation `json:"aggregation,omitempty"`

	Display string `json:"display,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertV2OperandColumnFromPb(pb *alertV2OperandColumnPb) (*AlertV2OperandColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2OperandColumn{}
	st.Aggregation = pb.Aggregation
	st.Display = pb.Display
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertV2OperandColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2OperandColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2OperandValue struct {

	// Wire name: 'bool_value'
	BoolValue bool

	// Wire name: 'double_value'
	DoubleValue float64

	// Wire name: 'string_value'
	StringValue string

	ForceSendFields []string `tf:"-"`
}

func alertV2OperandValueToPb(st *AlertV2OperandValue) (*alertV2OperandValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2OperandValuePb{}
	pb.BoolValue = st.BoolValue

	pb.DoubleValue = st.DoubleValue

	pb.StringValue = st.StringValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertV2OperandValuePb struct {
	BoolValue bool `json:"bool_value,omitempty"`

	DoubleValue float64 `json:"double_value,omitempty"`

	StringValue string `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertV2OperandValueFromPb(pb *alertV2OperandValuePb) (*AlertV2OperandValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2OperandValue{}
	st.BoolValue = pb.BoolValue
	st.DoubleValue = pb.DoubleValue
	st.StringValue = pb.StringValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertV2OperandValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2OperandValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2Subscription struct {

	// Wire name: 'destination_id'
	DestinationId string

	// Wire name: 'user_email'
	UserEmail string

	ForceSendFields []string `tf:"-"`
}

func alertV2SubscriptionToPb(st *AlertV2Subscription) (*alertV2SubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2SubscriptionPb{}
	pb.DestinationId = st.DestinationId

	pb.UserEmail = st.UserEmail

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type alertV2SubscriptionPb struct {
	DestinationId string `json:"destination_id,omitempty"`

	UserEmail string `json:"user_email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertV2SubscriptionFromPb(pb *alertV2SubscriptionPb) (*AlertV2Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Subscription{}
	st.DestinationId = pb.DestinationId
	st.UserEmail = pb.UserEmail

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertV2SubscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2SubscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Describes metadata for a particular chunk, within a result set; this
// structure is used both within a manifest, and when fetching individual chunk
// data or links.
type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset int64

	ForceSendFields []string `tf:"-"`
}

func baseChunkInfoToPb(st *BaseChunkInfo) (*baseChunkInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &baseChunkInfoPb{}
	pb.ByteCount = st.ByteCount

	pb.ChunkIndex = st.ChunkIndex

	pb.RowCount = st.RowCount

	pb.RowOffset = st.RowOffset

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type baseChunkInfoPb struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func baseChunkInfoFromPb(pb *baseChunkInfoPb) (*BaseChunkInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BaseChunkInfo{}
	st.ByteCount = pb.ByteCount
	st.ChunkIndex = pb.ChunkIndex
	st.RowCount = pb.RowCount
	st.RowOffset = pb.RowOffset

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *baseChunkInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st baseChunkInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Cancel statement execution
type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string `tf:"-"`
}

func cancelExecutionRequestToPb(st *CancelExecutionRequest) (*cancelExecutionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelExecutionRequestPb{}
	pb.StatementId = st.StatementId

	return pb, nil
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

type cancelExecutionRequestPb struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
}

func cancelExecutionRequestFromPb(pb *cancelExecutionRequestPb) (*CancelExecutionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelExecutionRequest{}
	st.StatementId = pb.StatementId

	return st, nil
}

type CancelExecutionResponse struct {
}

func cancelExecutionResponseToPb(st *CancelExecutionResponse) (*cancelExecutionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelExecutionResponsePb{}

	return pb, nil
}

func (st *CancelExecutionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelExecutionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelExecutionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelExecutionResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelExecutionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type cancelExecutionResponsePb struct {
}

func cancelExecutionResponseFromPb(pb *cancelExecutionResponsePb) (*CancelExecutionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelExecutionResponse{}

	return st, nil
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel struct {

	// Wire name: 'dbsql_version'
	DbsqlVersion string

	// Wire name: 'name'
	Name ChannelName

	ForceSendFields []string `tf:"-"`
}

func channelToPb(st *Channel) (*channelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &channelPb{}
	pb.DbsqlVersion = st.DbsqlVersion

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type channelPb struct {
	DbsqlVersion string `json:"dbsql_version,omitempty"`

	Name ChannelName `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func channelFromPb(pb *channelPb) (*Channel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Channel{}
	st.DbsqlVersion = pb.DbsqlVersion
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *channelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st channelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Details about a Channel.
type ChannelInfo struct {
	// DB SQL Version the Channel is mapped to.
	// Wire name: 'dbsql_version'
	DbsqlVersion string
	// Name of the channel
	// Wire name: 'name'
	Name ChannelName

	ForceSendFields []string `tf:"-"`
}

func channelInfoToPb(st *ChannelInfo) (*channelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &channelInfoPb{}
	pb.DbsqlVersion = st.DbsqlVersion

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type channelInfoPb struct {
	// DB SQL Version the Channel is mapped to.
	DbsqlVersion string `json:"dbsql_version,omitempty"`
	// Name of the channel
	Name ChannelName `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func channelInfoFromPb(pb *channelInfoPb) (*ChannelInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChannelInfo{}
	st.DbsqlVersion = pb.DbsqlVersion
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *channelInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st channelInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ChannelName string
type channelNamePb string

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

// Type always returns ChannelName to satisfy [pflag.Value] interface
func (f *ChannelName) Type() string {
	return "ChannelName"
}

func channelNameToPb(st *ChannelName) (*channelNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := channelNamePb(*st)
	return &pb, nil
}

func channelNameFromPb(pb *channelNamePb) (*ChannelName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ChannelName(*pb)
	return &st, nil
}

type ClientConfig struct {

	// Wire name: 'allow_custom_js_visualizations'
	AllowCustomJsVisualizations bool

	// Wire name: 'allow_downloads'
	AllowDownloads bool

	// Wire name: 'allow_external_shares'
	AllowExternalShares bool

	// Wire name: 'allow_subscriptions'
	AllowSubscriptions bool

	// Wire name: 'date_format'
	DateFormat string

	// Wire name: 'date_time_format'
	DateTimeFormat string

	// Wire name: 'disable_publish'
	DisablePublish bool

	// Wire name: 'enable_legacy_autodetect_types'
	EnableLegacyAutodetectTypes bool

	// Wire name: 'feature_show_permissions_control'
	FeatureShowPermissionsControl bool

	// Wire name: 'hide_plotly_mode_bar'
	HidePlotlyModeBar bool

	ForceSendFields []string `tf:"-"`
}

func clientConfigToPb(st *ClientConfig) (*clientConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clientConfigPb{}
	pb.AllowCustomJsVisualizations = st.AllowCustomJsVisualizations

	pb.AllowDownloads = st.AllowDownloads

	pb.AllowExternalShares = st.AllowExternalShares

	pb.AllowSubscriptions = st.AllowSubscriptions

	pb.DateFormat = st.DateFormat

	pb.DateTimeFormat = st.DateTimeFormat

	pb.DisablePublish = st.DisablePublish

	pb.EnableLegacyAutodetectTypes = st.EnableLegacyAutodetectTypes

	pb.FeatureShowPermissionsControl = st.FeatureShowPermissionsControl

	pb.HidePlotlyModeBar = st.HidePlotlyModeBar

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clientConfigPb struct {
	AllowCustomJsVisualizations bool `json:"allow_custom_js_visualizations,omitempty"`

	AllowDownloads bool `json:"allow_downloads,omitempty"`

	AllowExternalShares bool `json:"allow_external_shares,omitempty"`

	AllowSubscriptions bool `json:"allow_subscriptions,omitempty"`

	DateFormat string `json:"date_format,omitempty"`

	DateTimeFormat string `json:"date_time_format,omitempty"`

	DisablePublish bool `json:"disable_publish,omitempty"`

	EnableLegacyAutodetectTypes bool `json:"enable_legacy_autodetect_types,omitempty"`

	FeatureShowPermissionsControl bool `json:"feature_show_permissions_control,omitempty"`

	HidePlotlyModeBar bool `json:"hide_plotly_mode_bar,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clientConfigFromPb(pb *clientConfigPb) (*ClientConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClientConfig{}
	st.AllowCustomJsVisualizations = pb.AllowCustomJsVisualizations
	st.AllowDownloads = pb.AllowDownloads
	st.AllowExternalShares = pb.AllowExternalShares
	st.AllowSubscriptions = pb.AllowSubscriptions
	st.DateFormat = pb.DateFormat
	st.DateTimeFormat = pb.DateTimeFormat
	st.DisablePublish = pb.DisablePublish
	st.EnableLegacyAutodetectTypes = pb.EnableLegacyAutodetectTypes
	st.FeatureShowPermissionsControl = pb.FeatureShowPermissionsControl
	st.HidePlotlyModeBar = pb.HidePlotlyModeBar

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clientConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clientConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnInfo struct {
	// The name of the column.
	// Wire name: 'name'
	Name string
	// The ordinal position of the column (starting at position 0).
	// Wire name: 'position'
	Position int
	// The format of the interval type.
	// Wire name: 'type_interval_type'
	TypeIntervalType string
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	// Wire name: 'type_name'
	TypeName ColumnInfoTypeName
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	// Wire name: 'type_precision'
	TypePrecision int
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	// Wire name: 'type_scale'
	TypeScale int
	// The full SQL type specification.
	// Wire name: 'type_text'
	TypeText string

	ForceSendFields []string `tf:"-"`
}

func columnInfoToPb(st *ColumnInfo) (*columnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnInfoPb{}
	pb.Name = st.Name

	pb.Position = st.Position

	pb.TypeIntervalType = st.TypeIntervalType

	pb.TypeName = st.TypeName

	pb.TypePrecision = st.TypePrecision

	pb.TypeScale = st.TypeScale

	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type columnInfoPb struct {
	// The name of the column.
	Name string `json:"name,omitempty"`
	// The ordinal position of the column (starting at position 0).
	Position int `json:"position,omitempty"`
	// The format of the interval type.
	TypeIntervalType string `json:"type_interval_type,omitempty"`
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName ColumnInfoTypeName `json:"type_name,omitempty"`
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision int `json:"type_precision,omitempty"`
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale int `json:"type_scale,omitempty"`
	// The full SQL type specification.
	TypeText string `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func columnInfoFromPb(pb *columnInfoPb) (*ColumnInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnInfo{}
	st.Name = pb.Name
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	st.TypeName = pb.TypeName
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *columnInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st columnInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The name of the base data type. This doesn't include details for complex
// types such as STRUCT, MAP or ARRAY.
type ColumnInfoTypeName string
type columnInfoTypeNamePb string

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

// Type always returns ColumnInfoTypeName to satisfy [pflag.Value] interface
func (f *ColumnInfoTypeName) Type() string {
	return "ColumnInfoTypeName"
}

func columnInfoTypeNameToPb(st *ColumnInfoTypeName) (*columnInfoTypeNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := columnInfoTypeNamePb(*st)
	return &pb, nil
}

func columnInfoTypeNameFromPb(pb *columnInfoTypeNamePb) (*ColumnInfoTypeName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ColumnInfoTypeName(*pb)
	return &st, nil
}

type ComparisonOperator string
type comparisonOperatorPb string

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

// Type always returns ComparisonOperator to satisfy [pflag.Value] interface
func (f *ComparisonOperator) Type() string {
	return "ComparisonOperator"
}

func comparisonOperatorToPb(st *ComparisonOperator) (*comparisonOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := comparisonOperatorPb(*st)
	return &pb, nil
}

func comparisonOperatorFromPb(pb *comparisonOperatorPb) (*ComparisonOperator, error) {
	if pb == nil {
		return nil, nil
	}
	st := ComparisonOperator(*pb)
	return &st, nil
}

type CreateAlert struct {
	// Name of the alert.
	// Wire name: 'name'
	Name string
	// Alert configuration options.
	// Wire name: 'options'
	Options AlertOptions
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string
	// Query ID.
	// Wire name: 'query_id'
	QueryId string
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm int

	ForceSendFields []string `tf:"-"`
}

func createAlertToPb(st *CreateAlert) (*createAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertPb{}
	pb.Name = st.Name

	optionsPb, err := alertOptionsToPb(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	pb.Parent = st.Parent

	pb.QueryId = st.QueryId

	pb.Rearm = st.Rearm

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createAlertPb struct {
	// Name of the alert.
	Name string `json:"name"`
	// Alert configuration options.
	Options alertOptionsPb `json:"options"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// Query ID.
	QueryId string `json:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAlertFromPb(pb *createAlertPb) (*CreateAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlert{}
	st.Name = pb.Name
	optionsField, err := alertOptionsFromPb(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	st.Parent = pb.Parent
	st.QueryId = pb.QueryId
	st.Rearm = pb.Rearm

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateAlertRequest struct {

	// Wire name: 'alert'
	Alert *CreateAlertRequestAlert
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool

	ForceSendFields []string `tf:"-"`
}

func createAlertRequestToPb(st *CreateAlertRequest) (*createAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertRequestPb{}
	alertPb, err := createAlertRequestAlertToPb(st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = alertPb
	}

	pb.AutoResolveDisplayName = st.AutoResolveDisplayName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createAlertRequestPb struct {
	Alert *createAlertRequestAlertPb `json:"alert,omitempty"`
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAlertRequestFromPb(pb *createAlertRequestPb) (*CreateAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertRequest{}
	alertField, err := createAlertRequestAlertFromPb(pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = alertField
	}
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAlertRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAlertRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool
	// The workspace path of the folder containing the alert.
	// Wire name: 'parent_path'
	ParentPath string
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int

	ForceSendFields []string `tf:"-"`
}

func createAlertRequestAlertToPb(st *CreateAlertRequestAlert) (*createAlertRequestAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertRequestAlertPb{}
	conditionPb, err := alertConditionToPb(st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = conditionPb
	}

	pb.CustomBody = st.CustomBody

	pb.CustomSubject = st.CustomSubject

	pb.DisplayName = st.DisplayName

	pb.NotifyOnOk = st.NotifyOnOk

	pb.ParentPath = st.ParentPath

	pb.QueryId = st.QueryId

	pb.SecondsToRetrigger = st.SecondsToRetrigger

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createAlertRequestAlertPb struct {
	// Trigger conditions of the alert.
	Condition *alertConditionPb `json:"condition,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The workspace path of the folder containing the alert.
	ParentPath string `json:"parent_path,omitempty"`
	// UUID of the query attached to the alert.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAlertRequestAlertFromPb(pb *createAlertRequestAlertPb) (*CreateAlertRequestAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertRequestAlert{}
	conditionField, err := alertConditionFromPb(pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = conditionField
	}
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	st.DisplayName = pb.DisplayName
	st.NotifyOnOk = pb.NotifyOnOk
	st.ParentPath = pb.ParentPath
	st.QueryId = pb.QueryId
	st.SecondsToRetrigger = pb.SecondsToRetrigger

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAlertRequestAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAlertRequestAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateAlertV2Request struct {

	// Wire name: 'alert'
	Alert *AlertV2
}

func createAlertV2RequestToPb(st *CreateAlertV2Request) (*createAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertV2RequestPb{}
	alertPb, err := alertV2ToPb(st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = alertPb
	}

	return pb, nil
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

type createAlertV2RequestPb struct {
	Alert *alertV2Pb `json:"alert,omitempty"`
}

func createAlertV2RequestFromPb(pb *createAlertV2RequestPb) (*CreateAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertV2Request{}
	alertField, err := alertV2FromPb(pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = alertField
	}

	return st, nil
}

type CreateQueryRequest struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool

	// Wire name: 'query'
	Query *CreateQueryRequestQuery

	ForceSendFields []string `tf:"-"`
}

func createQueryRequestToPb(st *CreateQueryRequest) (*createQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQueryRequestPb{}
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName

	queryPb, err := createQueryRequestQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createQueryRequestPb struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	AutoResolveDisplayName bool `json:"auto_resolve_display_name,omitempty"`

	Query *createQueryRequestQueryPb `json:"query,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createQueryRequestFromPb(pb *createQueryRequestPb) (*CreateQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateQueryRequest{}
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName
	queryField, err := createQueryRequestQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter
	// Workspace path of the workspace folder containing the object.
	// Wire name: 'parent_path'
	ParentPath string
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string

	// Wire name: 'tags'
	Tags []string
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func createQueryRequestQueryToPb(st *CreateQueryRequestQuery) (*createQueryRequestQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQueryRequestQueryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit

	pb.Catalog = st.Catalog

	pb.Description = st.Description

	pb.DisplayName = st.DisplayName

	var parametersPb []queryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := queryParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	pb.ParentPath = st.ParentPath

	pb.QueryText = st.QueryText

	pb.RunAsMode = st.RunAsMode

	pb.Schema = st.Schema

	pb.Tags = st.Tags

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createQueryRequestQueryPb struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// List of query parameter definitions.
	Parameters []queryParameterPb `json:"parameters,omitempty"`
	// Workspace path of the workspace folder containing the object.
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createQueryRequestQueryFromPb(pb *createQueryRequestQueryPb) (*CreateQueryRequestQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateQueryRequestQuery{}
	st.ApplyAutoLimit = pb.ApplyAutoLimit
	st.Catalog = pb.Catalog
	st.Description = pb.Description
	st.DisplayName = pb.DisplayName

	var parametersField []QueryParameter
	for _, itemPb := range pb.Parameters {
		item, err := queryParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.ParentPath = pb.ParentPath
	st.QueryText = pb.QueryText
	st.RunAsMode = pb.RunAsMode
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createQueryRequestQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createQueryRequestQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	// Wire name: 'description'
	Description string
	// The name of the visualization that appears on dashboards and the query
	// screen.
	// Wire name: 'name'
	Name string
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	// Wire name: 'options'
	Options any
	// The identifier returned by :method:queries/create
	// Wire name: 'query_id'
	QueryId string
	// The type of visualization: chart, table, pivot table, and so on.
	// Wire name: 'type'
	Type string

	ForceSendFields []string `tf:"-"`
}

func createQueryVisualizationsLegacyRequestToPb(st *CreateQueryVisualizationsLegacyRequest) (*createQueryVisualizationsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQueryVisualizationsLegacyRequestPb{}
	pb.Description = st.Description

	pb.Name = st.Name

	pb.Options = st.Options

	pb.QueryId = st.QueryId

	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createQueryVisualizationsLegacyRequestPb struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description string `json:"description,omitempty"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name string `json:"name,omitempty"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `json:"options"`
	// The identifier returned by :method:queries/create
	QueryId string `json:"query_id"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type string `json:"type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createQueryVisualizationsLegacyRequestFromPb(pb *createQueryVisualizationsLegacyRequestPb) (*CreateQueryVisualizationsLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateQueryVisualizationsLegacyRequest{}
	st.Description = pb.Description
	st.Name = pb.Name
	st.Options = pb.Options
	st.QueryId = pb.QueryId
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createQueryVisualizationsLegacyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createQueryVisualizationsLegacyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateVisualizationRequest struct {

	// Wire name: 'visualization'
	Visualization *CreateVisualizationRequestVisualization
}

func createVisualizationRequestToPb(st *CreateVisualizationRequest) (*createVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVisualizationRequestPb{}
	visualizationPb, err := createVisualizationRequestVisualizationToPb(st.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationPb != nil {
		pb.Visualization = visualizationPb
	}

	return pb, nil
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

type createVisualizationRequestPb struct {
	Visualization *createVisualizationRequestVisualizationPb `json:"visualization,omitempty"`
}

func createVisualizationRequestFromPb(pb *createVisualizationRequestPb) (*CreateVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVisualizationRequest{}
	visualizationField, err := createVisualizationRequestVisualizationFromPb(pb.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationField != nil {
		st.Visualization = visualizationField
	}

	return st, nil
}

type CreateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	// Wire name: 'display_name'
	DisplayName string
	// UUID of the query that the visualization is attached to.
	// Wire name: 'query_id'
	QueryId string
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type string

	ForceSendFields []string `tf:"-"`
}

func createVisualizationRequestVisualizationToPb(st *CreateVisualizationRequestVisualization) (*createVisualizationRequestVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVisualizationRequestVisualizationPb{}
	pb.DisplayName = st.DisplayName

	pb.QueryId = st.QueryId

	pb.SerializedOptions = st.SerializedOptions

	pb.SerializedQueryPlan = st.SerializedQueryPlan

	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createVisualizationRequestVisualizationPb struct {
	// The display name of the visualization.
	DisplayName string `json:"display_name,omitempty"`
	// UUID of the query that the visualization is attached to.
	QueryId string `json:"query_id,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createVisualizationRequestVisualizationFromPb(pb *createVisualizationRequestVisualizationPb) (*CreateVisualizationRequestVisualization, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVisualizationRequestVisualization{}
	st.DisplayName = pb.DisplayName
	st.QueryId = pb.QueryId
	st.SerializedOptions = pb.SerializedOptions
	st.SerializedQueryPlan = pb.SerializedQueryPlan
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createVisualizationRequestVisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createVisualizationRequestVisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	AutoStopMins int
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string
	// Configurations whether the warehouse should use spot instances.
	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	// Wire name: 'warehouse_type'
	WarehouseType CreateWarehouseRequestWarehouseType

	ForceSendFields []string `tf:"-"`
}

func createWarehouseRequestToPb(st *CreateWarehouseRequest) (*createWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWarehouseRequestPb{}
	pb.AutoStopMins = st.AutoStopMins

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	pb.ClusterSize = st.ClusterSize

	pb.CreatorName = st.CreatorName

	pb.EnablePhoton = st.EnablePhoton

	pb.EnableServerlessCompute = st.EnableServerlessCompute

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.MaxNumClusters = st.MaxNumClusters

	pb.MinNumClusters = st.MinNumClusters

	pb.Name = st.Name

	pb.SpotInstancePolicy = st.SpotInstancePolicy

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createWarehouseRequestPb struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be >= 0 mins for serverless warehouses - Must be
	// == 0 or >= 10 mins for non-serverless warehouses - 0 indicates no
	// autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *channelPb `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *endpointTagsPb `json:"tags,omitempty"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType CreateWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createWarehouseRequestFromPb(pb *createWarehouseRequestPb) (*CreateWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWarehouseRequest{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	st.ClusterSize = pb.ClusterSize
	st.CreatorName = pb.CreatorName
	st.EnablePhoton = pb.EnablePhoton
	st.EnableServerlessCompute = pb.EnableServerlessCompute
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxNumClusters = pb.MaxNumClusters
	st.MinNumClusters = pb.MinNumClusters
	st.Name = pb.Name
	st.SpotInstancePolicy = pb.SpotInstancePolicy
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	st.WarehouseType = pb.WarehouseType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createWarehouseRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createWarehouseRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type CreateWarehouseRequestWarehouseType string
type createWarehouseRequestWarehouseTypePb string

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

// Type always returns CreateWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *CreateWarehouseRequestWarehouseType) Type() string {
	return "CreateWarehouseRequestWarehouseType"
}

func createWarehouseRequestWarehouseTypeToPb(st *CreateWarehouseRequestWarehouseType) (*createWarehouseRequestWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := createWarehouseRequestWarehouseTypePb(*st)
	return &pb, nil
}

func createWarehouseRequestWarehouseTypeFromPb(pb *createWarehouseRequestWarehouseTypePb) (*CreateWarehouseRequestWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateWarehouseRequestWarehouseType(*pb)
	return &st, nil
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	// Wire name: 'id'
	Id string

	ForceSendFields []string `tf:"-"`
}

func createWarehouseResponseToPb(st *CreateWarehouseResponse) (*createWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWarehouseResponsePb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createWarehouseResponsePb struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createWarehouseResponseFromPb(pb *createWarehouseResponsePb) (*CreateWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWarehouseResponse{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createWarehouseResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createWarehouseResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	// Wire name: 'dashboard_id'
	DashboardId string
	// Widget ID returned by :method:dashboardwidgets/create
	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'options'
	Options WidgetOptions
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	// Wire name: 'text'
	Text string
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	// Wire name: 'visualization_id'
	VisualizationId string
	// Width of a widget
	// Wire name: 'width'
	Width int

	ForceSendFields []string `tf:"-"`
}

func createWidgetToPb(st *CreateWidget) (*createWidgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWidgetPb{}
	pb.DashboardId = st.DashboardId

	pb.Id = st.Id

	optionsPb, err := widgetOptionsToPb(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	pb.Text = st.Text

	pb.VisualizationId = st.VisualizationId

	pb.Width = st.Width

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createWidgetPb struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId string `json:"dashboard_id"`
	// Widget ID returned by :method:dashboardwidgets/create
	Id string `json:"-" url:"-"`

	Options widgetOptionsPb `json:"options"`
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text string `json:"text,omitempty"`
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId string `json:"visualization_id,omitempty"`
	// Width of a widget
	Width int `json:"width"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createWidgetFromPb(pb *createWidgetPb) (*CreateWidget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWidget{}
	st.DashboardId = pb.DashboardId
	st.Id = pb.Id
	optionsField, err := widgetOptionsFromPb(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	st.Text = pb.Text
	st.VisualizationId = pb.VisualizationId
	st.Width = pb.Width

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createWidgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createWidgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus SchedulePauseStatus
	// A cron expression using quartz syntax that specifies the schedule for
	// this pipeline. Should use the quartz format described here:
	// http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
	// Wire name: 'quartz_cron_schedule'
	QuartzCronSchedule string
	// A Java timezone id. The schedule will be resolved using this timezone.
	// This will be combined with the quartz_cron_schedule to determine the
	// schedule. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details.
	// Wire name: 'timezone_id'
	TimezoneId string

	ForceSendFields []string `tf:"-"`
}

func cronScheduleToPb(st *CronSchedule) (*cronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronSchedulePb{}
	pb.PauseStatus = st.PauseStatus

	pb.QuartzCronSchedule = st.QuartzCronSchedule

	pb.TimezoneId = st.TimezoneId

	pb.ForceSendFields = st.ForceSendFields
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
	PauseStatus SchedulePauseStatus `json:"pause_status,omitempty"`
	// A cron expression using quartz syntax that specifies the schedule for
	// this pipeline. Should use the quartz format described here:
	// http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`
	// A Java timezone id. The schedule will be resolved using this timezone.
	// This will be combined with the quartz_cron_schedule to determine the
	// schedule. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details.
	TimezoneId string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cronScheduleFromPb(pb *cronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	st.PauseStatus = pb.PauseStatus
	st.QuartzCronSchedule = pb.QuartzCronSchedule
	st.TimezoneId = pb.TimezoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cronSchedulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cronSchedulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	// Wire name: 'can_edit'
	CanEdit bool
	// Timestamp when this dashboard was created.
	// Wire name: 'created_at'
	CreatedAt string
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	// Wire name: 'dashboard_filters_enabled'
	DashboardFiltersEnabled bool
	// The ID for this dashboard.
	// Wire name: 'id'
	Id string
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	// Wire name: 'is_draft'
	IsDraft bool
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	// Wire name: 'is_favorite'
	IsFavorite bool
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string

	// Wire name: 'options'
	Options *DashboardOptions
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_tier'
	PermissionTier PermissionLevel
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	// Wire name: 'slug'
	Slug string

	// Wire name: 'tags'
	Tags []string
	// Timestamp when this dashboard was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string

	// Wire name: 'user'
	User *User
	// The ID of the user who owns the dashboard.
	// Wire name: 'user_id'
	UserId int

	// Wire name: 'widgets'
	Widgets []Widget

	ForceSendFields []string `tf:"-"`
}

func dashboardToPb(st *Dashboard) (*dashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPb{}
	pb.CanEdit = st.CanEdit

	pb.CreatedAt = st.CreatedAt

	pb.DashboardFiltersEnabled = st.DashboardFiltersEnabled

	pb.Id = st.Id

	pb.IsArchived = st.IsArchived

	pb.IsDraft = st.IsDraft

	pb.IsFavorite = st.IsFavorite

	pb.Name = st.Name

	optionsPb, err := dashboardOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	pb.Parent = st.Parent

	pb.PermissionTier = st.PermissionTier

	pb.Slug = st.Slug

	pb.Tags = st.Tags

	pb.UpdatedAt = st.UpdatedAt

	userPb, err := userToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}

	pb.UserId = st.UserId

	var widgetsPb []widgetPb
	for _, item := range st.Widgets {
		itemPb, err := widgetToPb(&item)
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

type dashboardPb struct {
	// Whether the authenticated user can edit the query definition.
	CanEdit bool `json:"can_edit,omitempty"`
	// Timestamp when this dashboard was created.
	CreatedAt string `json:"created_at,omitempty"`
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled,omitempty"`
	// The ID for this dashboard.
	Id string `json:"id,omitempty"`
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft bool `json:"is_draft,omitempty"`
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite bool `json:"is_favorite,omitempty"`
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string `json:"name,omitempty"`

	Options *dashboardOptionsPb `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug string `json:"slug,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// Timestamp when this dashboard was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *userPb `json:"user,omitempty"`
	// The ID of the user who owns the dashboard.
	UserId int `json:"user_id,omitempty"`

	Widgets []widgetPb `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardFromPb(pb *dashboardPb) (*Dashboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dashboard{}
	st.CanEdit = pb.CanEdit
	st.CreatedAt = pb.CreatedAt
	st.DashboardFiltersEnabled = pb.DashboardFiltersEnabled
	st.Id = pb.Id
	st.IsArchived = pb.IsArchived
	st.IsDraft = pb.IsDraft
	st.IsFavorite = pb.IsFavorite
	st.Name = pb.Name
	optionsField, err := dashboardOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	st.Parent = pb.Parent
	st.PermissionTier = pb.PermissionTier
	st.Slug = pb.Slug
	st.Tags = pb.Tags
	st.UpdatedAt = pb.UpdatedAt
	userField, err := userFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}
	st.UserId = pb.UserId

	var widgetsField []Widget
	for _, itemPb := range pb.Widgets {
		item, err := widgetFromPb(&itemPb)
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

func (st *dashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardEditContent struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole

	// Wire name: 'tags'
	Tags []string

	ForceSendFields []string `tf:"-"`
}

func dashboardEditContentToPb(st *DashboardEditContent) (*dashboardEditContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardEditContentPb{}
	pb.DashboardId = st.DashboardId

	pb.Name = st.Name

	pb.RunAsRole = st.RunAsRole

	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dashboardEditContentPb struct {
	DashboardId string `json:"-" url:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string `json:"name,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardEditContentFromPb(pb *dashboardEditContentPb) (*DashboardEditContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardEditContent{}
	st.DashboardId = pb.DashboardId
	st.Name = pb.Name
	st.RunAsRole = pb.RunAsRole
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardEditContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardEditContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	// Wire name: 'moved_to_trash_at'
	MovedToTrashAt string

	ForceSendFields []string `tf:"-"`
}

func dashboardOptionsToPb(st *DashboardOptions) (*dashboardOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardOptionsPb{}
	pb.MovedToTrashAt = st.MovedToTrashAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dashboardOptionsPb struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardOptionsFromPb(pb *dashboardOptionsPb) (*DashboardOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardOptions{}
	st.MovedToTrashAt = pb.MovedToTrashAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardPostContent struct {
	// Indicates whether the dashboard filters are enabled
	// Wire name: 'dashboard_filters_enabled'
	DashboardFiltersEnabled bool
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	// Wire name: 'is_favorite'
	IsFavorite bool
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole

	// Wire name: 'tags'
	Tags []string

	ForceSendFields []string `tf:"-"`
}

func dashboardPostContentToPb(st *DashboardPostContent) (*dashboardPostContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPostContentPb{}
	pb.DashboardFiltersEnabled = st.DashboardFiltersEnabled

	pb.IsFavorite = st.IsFavorite

	pb.Name = st.Name

	pb.Parent = st.Parent

	pb.RunAsRole = st.RunAsRole

	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dashboardPostContentPb struct {
	// Indicates whether the dashboard filters are enabled
	DashboardFiltersEnabled bool `json:"dashboard_filters_enabled,omitempty"`
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	IsFavorite bool `json:"is_favorite,omitempty"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string `json:"name"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardPostContentFromPb(pb *dashboardPostContentPb) (*DashboardPostContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardPostContent{}
	st.DashboardFiltersEnabled = pb.DashboardFiltersEnabled
	st.IsFavorite = pb.IsFavorite
	st.Name = pb.Name
	st.Parent = pb.Parent
	st.RunAsRole = pb.RunAsRole
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardPostContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardPostContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'id'
	Id string
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	// Wire name: 'name'
	Name string
	// Reserved for internal use.
	// Wire name: 'pause_reason'
	PauseReason string
	// Reserved for internal use.
	// Wire name: 'paused'
	Paused int
	// Reserved for internal use.
	// Wire name: 'supports_auto_limit'
	SupportsAutoLimit bool
	// Reserved for internal use.
	// Wire name: 'syntax'
	Syntax string
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	// Wire name: 'type'
	Type string
	// Reserved for internal use.
	// Wire name: 'view_only'
	ViewOnly bool
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func dataSourceToPb(st *DataSource) (*dataSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataSourcePb{}
	pb.Id = st.Id

	pb.Name = st.Name

	pb.PauseReason = st.PauseReason

	pb.Paused = st.Paused

	pb.SupportsAutoLimit = st.SupportsAutoLimit

	pb.Syntax = st.Syntax

	pb.Type = st.Type

	pb.ViewOnly = st.ViewOnly

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dataSourcePb struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	Id string `json:"id,omitempty"`
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name string `json:"name,omitempty"`
	// Reserved for internal use.
	PauseReason string `json:"pause_reason,omitempty"`
	// Reserved for internal use.
	Paused int `json:"paused,omitempty"`
	// Reserved for internal use.
	SupportsAutoLimit bool `json:"supports_auto_limit,omitempty"`
	// Reserved for internal use.
	Syntax string `json:"syntax,omitempty"`
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	Type string `json:"type,omitempty"`
	// Reserved for internal use.
	ViewOnly bool `json:"view_only,omitempty"`
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataSourceFromPb(pb *dataSourcePb) (*DataSource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataSource{}
	st.Id = pb.Id
	st.Name = pb.Name
	st.PauseReason = pb.PauseReason
	st.Paused = pb.Paused
	st.SupportsAutoLimit = pb.SupportsAutoLimit
	st.Syntax = pb.Syntax
	st.Type = pb.Type
	st.ViewOnly = pb.ViewOnly
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dataSourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dataSourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatePrecision string
type datePrecisionPb string

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

// Type always returns DatePrecision to satisfy [pflag.Value] interface
func (f *DatePrecision) Type() string {
	return "DatePrecision"
}

func datePrecisionToPb(st *DatePrecision) (*datePrecisionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := datePrecisionPb(*st)
	return &pb, nil
}

func datePrecisionFromPb(pb *datePrecisionPb) (*DatePrecision, error) {
	if pb == nil {
		return nil, nil
	}
	st := DatePrecision(*pb)
	return &st, nil
}

type DateRange struct {

	// Wire name: 'end'
	End string

	// Wire name: 'start'
	Start string
}

func dateRangeToPb(st *DateRange) (*dateRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dateRangePb{}
	pb.End = st.End

	pb.Start = st.Start

	return pb, nil
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

type dateRangePb struct {
	End string `json:"end"`

	Start string `json:"start"`
}

func dateRangeFromPb(pb *dateRangePb) (*DateRange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DateRange{}
	st.End = pb.End
	st.Start = pb.Start

	return st, nil
}

type DateRangeValue struct {
	// Manually specified date-time range value.
	// Wire name: 'date_range_value'
	DateRangeValue *DateRange
	// Dynamic date-time range value based on current date-time.
	// Wire name: 'dynamic_date_range_value'
	DynamicDateRangeValue DateRangeValueDynamicDateRange
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	// Wire name: 'precision'
	Precision DatePrecision

	// Wire name: 'start_day_of_week'
	StartDayOfWeek int

	ForceSendFields []string `tf:"-"`
}

func dateRangeValueToPb(st *DateRangeValue) (*dateRangeValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dateRangeValuePb{}
	dateRangeValuePb, err := dateRangeToPb(st.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValuePb != nil {
		pb.DateRangeValue = dateRangeValuePb
	}

	pb.DynamicDateRangeValue = st.DynamicDateRangeValue

	pb.Precision = st.Precision

	pb.StartDayOfWeek = st.StartDayOfWeek

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dateRangeValuePb struct {
	// Manually specified date-time range value.
	DateRangeValue *dateRangePb `json:"date_range_value,omitempty"`
	// Dynamic date-time range value based on current date-time.
	DynamicDateRangeValue DateRangeValueDynamicDateRange `json:"dynamic_date_range_value,omitempty"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision DatePrecision `json:"precision,omitempty"`

	StartDayOfWeek int `json:"start_day_of_week,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dateRangeValueFromPb(pb *dateRangeValuePb) (*DateRangeValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DateRangeValue{}
	dateRangeValueField, err := dateRangeFromPb(pb.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValueField != nil {
		st.DateRangeValue = dateRangeValueField
	}
	st.DynamicDateRangeValue = pb.DynamicDateRangeValue
	st.Precision = pb.Precision
	st.StartDayOfWeek = pb.StartDayOfWeek

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dateRangeValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dateRangeValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DateRangeValueDynamicDateRange string
type dateRangeValueDynamicDateRangePb string

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

// Type always returns DateRangeValueDynamicDateRange to satisfy [pflag.Value] interface
func (f *DateRangeValueDynamicDateRange) Type() string {
	return "DateRangeValueDynamicDateRange"
}

func dateRangeValueDynamicDateRangeToPb(st *DateRangeValueDynamicDateRange) (*dateRangeValueDynamicDateRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dateRangeValueDynamicDateRangePb(*st)
	return &pb, nil
}

func dateRangeValueDynamicDateRangeFromPb(pb *dateRangeValueDynamicDateRangePb) (*DateRangeValueDynamicDateRange, error) {
	if pb == nil {
		return nil, nil
	}
	st := DateRangeValueDynamicDateRange(*pb)
	return &st, nil
}

type DateValue struct {
	// Manually specified date-time value.
	// Wire name: 'date_value'
	DateValue string
	// Dynamic date-time value based on current date-time.
	// Wire name: 'dynamic_date_value'
	DynamicDateValue DateValueDynamicDate
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	// Wire name: 'precision'
	Precision DatePrecision

	ForceSendFields []string `tf:"-"`
}

func dateValueToPb(st *DateValue) (*dateValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dateValuePb{}
	pb.DateValue = st.DateValue

	pb.DynamicDateValue = st.DynamicDateValue

	pb.Precision = st.Precision

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dateValuePb struct {
	// Manually specified date-time value.
	DateValue string `json:"date_value,omitempty"`
	// Dynamic date-time value based on current date-time.
	DynamicDateValue DateValueDynamicDate `json:"dynamic_date_value,omitempty"`
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision DatePrecision `json:"precision,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dateValueFromPb(pb *dateValuePb) (*DateValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DateValue{}
	st.DateValue = pb.DateValue
	st.DynamicDateValue = pb.DynamicDateValue
	st.Precision = pb.Precision

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dateValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dateValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DateValueDynamicDate string
type dateValueDynamicDatePb string

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

// Type always returns DateValueDynamicDate to satisfy [pflag.Value] interface
func (f *DateValueDynamicDate) Type() string {
	return "DateValueDynamicDate"
}

func dateValueDynamicDateToPb(st *DateValueDynamicDate) (*dateValueDynamicDatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dateValueDynamicDatePb(*st)
	return &pb, nil
}

func dateValueDynamicDateFromPb(pb *dateValueDynamicDatePb) (*DateValueDynamicDate, error) {
	if pb == nil {
		return nil, nil
	}
	st := DateValueDynamicDate(*pb)
	return &st, nil
}

// Delete an alert
type DeleteAlertsLegacyRequest struct {

	// Wire name: 'alert_id'
	AlertId string `tf:"-"`
}

func deleteAlertsLegacyRequestToPb(st *DeleteAlertsLegacyRequest) (*deleteAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAlertsLegacyRequestPb{}
	pb.AlertId = st.AlertId

	return pb, nil
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

type deleteAlertsLegacyRequestPb struct {
	AlertId string `json:"-" url:"-"`
}

func deleteAlertsLegacyRequestFromPb(pb *deleteAlertsLegacyRequestPb) (*DeleteAlertsLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAlertsLegacyRequest{}
	st.AlertId = pb.AlertId

	return st, nil
}

// Remove a dashboard
type DeleteDashboardRequest struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func deleteDashboardRequestToPb(st *DeleteDashboardRequest) (*deleteDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
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

type deleteDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func deleteDashboardRequestFromPb(pb *deleteDashboardRequestPb) (*DeleteDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

// Remove widget
type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteDashboardWidgetRequestToPb(st *DeleteDashboardWidgetRequest) (*deleteDashboardWidgetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardWidgetRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteDashboardWidgetRequestPb struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id string `json:"-" url:"-"`
}

func deleteDashboardWidgetRequestFromPb(pb *deleteDashboardWidgetRequestPb) (*DeleteDashboardWidgetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardWidgetRequest{}
	st.Id = pb.Id

	return st, nil
}

// Delete a query
type DeleteQueriesLegacyRequest struct {

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
}

func deleteQueriesLegacyRequestToPb(st *DeleteQueriesLegacyRequest) (*deleteQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
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

type deleteQueriesLegacyRequestPb struct {
	QueryId string `json:"-" url:"-"`
}

func deleteQueriesLegacyRequestFromPb(pb *deleteQueriesLegacyRequestPb) (*DeleteQueriesLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQueriesLegacyRequest{}
	st.QueryId = pb.QueryId

	return st, nil
}

// Remove visualization
type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvizualisations/create
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteQueryVisualizationsLegacyRequestToPb(st *DeleteQueryVisualizationsLegacyRequest) (*deleteQueryVisualizationsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQueryVisualizationsLegacyRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteQueryVisualizationsLegacyRequestPb struct {
	// Widget ID returned by :method:queryvizualisations/create
	Id string `json:"-" url:"-"`
}

func deleteQueryVisualizationsLegacyRequestFromPb(pb *deleteQueryVisualizationsLegacyRequestPb) (*DeleteQueryVisualizationsLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQueryVisualizationsLegacyRequest{}
	st.Id = pb.Id

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

// Remove a visualization
type DeleteVisualizationRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteVisualizationRequestToPb(st *DeleteVisualizationRequest) (*deleteVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVisualizationRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteVisualizationRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteVisualizationRequestFromPb(pb *deleteVisualizationRequestPb) (*DeleteVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteVisualizationRequest{}
	st.Id = pb.Id

	return st, nil
}

// Delete a warehouse
type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteWarehouseRequestToPb(st *DeleteWarehouseRequest) (*deleteWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWarehouseRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type deleteWarehouseRequestPb struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func deleteWarehouseRequestFromPb(pb *deleteWarehouseRequestPb) (*DeleteWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWarehouseRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteWarehouseResponse struct {
}

func deleteWarehouseResponseToPb(st *DeleteWarehouseResponse) (*deleteWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWarehouseResponsePb{}

	return pb, nil
}

func (st *DeleteWarehouseResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWarehouseResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWarehouseResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWarehouseResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteWarehouseResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteWarehouseResponsePb struct {
}

func deleteWarehouseResponseFromPb(pb *deleteWarehouseResponsePb) (*DeleteWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWarehouseResponse{}

	return st, nil
}

type Disposition string
type dispositionPb string

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

// Type always returns Disposition to satisfy [pflag.Value] interface
func (f *Disposition) Type() string {
	return "Disposition"
}

func dispositionToPb(st *Disposition) (*dispositionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := dispositionPb(*st)
	return &pb, nil
}

func dispositionFromPb(pb *dispositionPb) (*Disposition, error) {
	if pb == nil {
		return nil, nil
	}
	st := Disposition(*pb)
	return &st, nil
}

type EditAlert struct {

	// Wire name: 'alert_id'
	AlertId string `tf:"-"`
	// Name of the alert.
	// Wire name: 'name'
	Name string
	// Alert configuration options.
	// Wire name: 'options'
	Options AlertOptions
	// Query ID.
	// Wire name: 'query_id'
	QueryId string
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm int

	ForceSendFields []string `tf:"-"`
}

func editAlertToPb(st *EditAlert) (*editAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editAlertPb{}
	pb.AlertId = st.AlertId

	pb.Name = st.Name

	optionsPb, err := alertOptionsToPb(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	pb.QueryId = st.QueryId

	pb.Rearm = st.Rearm

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type editAlertPb struct {
	AlertId string `json:"-" url:"-"`
	// Name of the alert.
	Name string `json:"name"`
	// Alert configuration options.
	Options alertOptionsPb `json:"options"`
	// Query ID.
	QueryId string `json:"query_id"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editAlertFromPb(pb *editAlertPb) (*EditAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditAlert{}
	st.AlertId = pb.AlertId
	st.Name = pb.Name
	optionsField, err := alertOptionsFromPb(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	st.QueryId = pb.QueryId
	st.Rearm = pb.Rearm

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute.
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool
	// Required. Id of the warehouse to configure.
	// Wire name: 'id'
	Id string `tf:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string
	// Configurations whether the warehouse should use spot instances.
	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	// Wire name: 'warehouse_type'
	WarehouseType EditWarehouseRequestWarehouseType

	ForceSendFields []string `tf:"-"`
}

func editWarehouseRequestToPb(st *EditWarehouseRequest) (*editWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editWarehouseRequestPb{}
	pb.AutoStopMins = st.AutoStopMins

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	pb.ClusterSize = st.ClusterSize

	pb.CreatorName = st.CreatorName

	pb.EnablePhoton = st.EnablePhoton

	pb.EnableServerlessCompute = st.EnableServerlessCompute

	pb.Id = st.Id

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.MaxNumClusters = st.MaxNumClusters

	pb.MinNumClusters = st.MinNumClusters

	pb.Name = st.Name

	pb.SpotInstancePolicy = st.SpotInstancePolicy

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type editWarehouseRequestPb struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *channelPb `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute.
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Required. Id of the warehouse to configure.
	Id string `json:"-" url:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *endpointTagsPb `json:"tags,omitempty"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType EditWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editWarehouseRequestFromPb(pb *editWarehouseRequestPb) (*EditWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditWarehouseRequest{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	st.ClusterSize = pb.ClusterSize
	st.CreatorName = pb.CreatorName
	st.EnablePhoton = pb.EnablePhoton
	st.EnableServerlessCompute = pb.EnableServerlessCompute
	st.Id = pb.Id
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxNumClusters = pb.MaxNumClusters
	st.MinNumClusters = pb.MinNumClusters
	st.Name = pb.Name
	st.SpotInstancePolicy = pb.SpotInstancePolicy
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	st.WarehouseType = pb.WarehouseType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editWarehouseRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editWarehouseRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type EditWarehouseRequestWarehouseType string
type editWarehouseRequestWarehouseTypePb string

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

// Type always returns EditWarehouseRequestWarehouseType to satisfy [pflag.Value] interface
func (f *EditWarehouseRequestWarehouseType) Type() string {
	return "EditWarehouseRequestWarehouseType"
}

func editWarehouseRequestWarehouseTypeToPb(st *EditWarehouseRequestWarehouseType) (*editWarehouseRequestWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := editWarehouseRequestWarehouseTypePb(*st)
	return &pb, nil
}

func editWarehouseRequestWarehouseTypeFromPb(pb *editWarehouseRequestWarehouseTypePb) (*EditWarehouseRequestWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EditWarehouseRequestWarehouseType(*pb)
	return &st, nil
}

type EditWarehouseResponse struct {
}

func editWarehouseResponseToPb(st *EditWarehouseResponse) (*editWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editWarehouseResponsePb{}

	return pb, nil
}

func (st *EditWarehouseResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &editWarehouseResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := editWarehouseResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EditWarehouseResponse) MarshalJSON() ([]byte, error) {
	pb, err := editWarehouseResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type editWarehouseResponsePb struct {
}

func editWarehouseResponseFromPb(pb *editWarehouseResponsePb) (*EditWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditWarehouseResponse{}

	return st, nil
}

// Represents an empty message, similar to google.protobuf.Empty, which is not
// available in the firm right now.
type Empty struct {
}

func emptyToPb(st *Empty) (*emptyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &emptyPb{}

	return pb, nil
}

func (st *Empty) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &emptyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := emptyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Empty) MarshalJSON() ([]byte, error) {
	pb, err := emptyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type emptyPb struct {
}

func emptyFromPb(pb *emptyPb) (*Empty, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Empty{}

	return st, nil
}

type EndpointConfPair struct {

	// Wire name: 'key'
	Key string

	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func endpointConfPairToPb(st *EndpointConfPair) (*endpointConfPairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointConfPairPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type endpointConfPairPb struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointConfPairFromPb(pb *endpointConfPairPb) (*EndpointConfPair, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointConfPair{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointConfPairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointConfPairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	// Wire name: 'details'
	Details string
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	// Wire name: 'failure_reason'
	FailureReason *TerminationReason
	// Deprecated. split into summary and details for security
	// Wire name: 'message'
	Message string
	// Health status of the warehouse.
	// Wire name: 'status'
	Status Status
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	// Wire name: 'summary'
	Summary string

	ForceSendFields []string `tf:"-"`
}

func endpointHealthToPb(st *EndpointHealth) (*endpointHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointHealthPb{}
	pb.Details = st.Details

	failureReasonPb, err := terminationReasonToPb(st.FailureReason)
	if err != nil {
		return nil, err
	}
	if failureReasonPb != nil {
		pb.FailureReason = failureReasonPb
	}

	pb.Message = st.Message

	pb.Status = st.Status

	pb.Summary = st.Summary

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type endpointHealthPb struct {
	// Details about errors that are causing current degraded/failed status.
	Details string `json:"details,omitempty"`
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason *terminationReasonPb `json:"failure_reason,omitempty"`
	// Deprecated. split into summary and details for security
	Message string `json:"message,omitempty"`
	// Health status of the warehouse.
	Status Status `json:"status,omitempty"`
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary string `json:"summary,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointHealthFromPb(pb *endpointHealthPb) (*EndpointHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointHealth{}
	st.Details = pb.Details
	failureReasonField, err := terminationReasonFromPb(pb.FailureReason)
	if err != nil {
		return nil, err
	}
	if failureReasonField != nil {
		st.FailureReason = failureReasonField
	}
	st.Message = pb.Message
	st.Status = pb.Status
	st.Summary = pb.Summary

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointHealthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointHealthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	// Wire name: 'health'
	Health *EndpointHealth
	// unique identifier for warehouse
	// Wire name: 'id'
	Id string
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string
	// the jdbc connection string for this warehouse
	// Wire name: 'jdbc_url'
	JdbcUrl string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string
	// Deprecated. current number of active sessions for the warehouse
	// Wire name: 'num_active_sessions'
	NumActiveSessions int64
	// current number of clusters running for the service
	// Wire name: 'num_clusters'
	NumClusters int
	// ODBC parameters for the SQL warehouse
	// Wire name: 'odbc_params'
	OdbcParams *OdbcParams
	// Configurations whether the warehouse should use spot instances.
	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy
	// State of the warehouse
	// Wire name: 'state'
	State State
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	// Wire name: 'warehouse_type'
	WarehouseType EndpointInfoWarehouseType

	ForceSendFields []string `tf:"-"`
}

func endpointInfoToPb(st *EndpointInfo) (*endpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointInfoPb{}
	pb.AutoStopMins = st.AutoStopMins

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	pb.ClusterSize = st.ClusterSize

	pb.CreatorName = st.CreatorName

	pb.EnablePhoton = st.EnablePhoton

	pb.EnableServerlessCompute = st.EnableServerlessCompute

	healthPb, err := endpointHealthToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	pb.Id = st.Id

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.JdbcUrl = st.JdbcUrl

	pb.MaxNumClusters = st.MaxNumClusters

	pb.MinNumClusters = st.MinNumClusters

	pb.Name = st.Name

	pb.NumActiveSessions = st.NumActiveSessions

	pb.NumClusters = st.NumClusters

	odbcParamsPb, err := odbcParamsToPb(st.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsPb != nil {
		pb.OdbcParams = odbcParamsPb
	}

	pb.SpotInstancePolicy = st.SpotInstancePolicy

	pb.State = st.State

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type endpointInfoPb struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *channelPb `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *endpointHealthPb `json:"health,omitempty"`
	// unique identifier for warehouse
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// the jdbc connection string for this warehouse
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// current number of clusters running for the service
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the SQL warehouse
	OdbcParams *odbcParamsPb `json:"odbc_params,omitempty"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// State of the warehouse
	State State `json:"state,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *endpointTagsPb `json:"tags,omitempty"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType EndpointInfoWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointInfoFromPb(pb *endpointInfoPb) (*EndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointInfo{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	st.ClusterSize = pb.ClusterSize
	st.CreatorName = pb.CreatorName
	st.EnablePhoton = pb.EnablePhoton
	st.EnableServerlessCompute = pb.EnableServerlessCompute
	healthField, err := endpointHealthFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	st.Id = pb.Id
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.JdbcUrl = pb.JdbcUrl
	st.MaxNumClusters = pb.MaxNumClusters
	st.MinNumClusters = pb.MinNumClusters
	st.Name = pb.Name
	st.NumActiveSessions = pb.NumActiveSessions
	st.NumClusters = pb.NumClusters
	odbcParamsField, err := odbcParamsFromPb(pb.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsField != nil {
		st.OdbcParams = odbcParamsField
	}
	st.SpotInstancePolicy = pb.SpotInstancePolicy
	st.State = pb.State
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	st.WarehouseType = pb.WarehouseType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type EndpointInfoWarehouseType string
type endpointInfoWarehouseTypePb string

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

// Type always returns EndpointInfoWarehouseType to satisfy [pflag.Value] interface
func (f *EndpointInfoWarehouseType) Type() string {
	return "EndpointInfoWarehouseType"
}

func endpointInfoWarehouseTypeToPb(st *EndpointInfoWarehouseType) (*endpointInfoWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := endpointInfoWarehouseTypePb(*st)
	return &pb, nil
}

func endpointInfoWarehouseTypeFromPb(pb *endpointInfoWarehouseTypePb) (*EndpointInfoWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointInfoWarehouseType(*pb)
	return &st, nil
}

type EndpointTagPair struct {

	// Wire name: 'key'
	Key string

	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func endpointTagPairToPb(st *EndpointTagPair) (*endpointTagPairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagPairPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type endpointTagPairPb struct {
	Key string `json:"key,omitempty"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointTagPairFromPb(pb *endpointTagPairPb) (*EndpointTagPair, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTagPair{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointTagPairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointTagPairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointTags struct {

	// Wire name: 'custom_tags'
	CustomTags []EndpointTagPair
}

func endpointTagsToPb(st *EndpointTags) (*endpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagsPb{}

	var customTagsPb []endpointTagPairPb
	for _, item := range st.CustomTags {
		itemPb, err := endpointTagPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb = append(customTagsPb, *itemPb)
		}
	}
	pb.CustomTags = customTagsPb

	return pb, nil
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

type endpointTagsPb struct {
	CustomTags []endpointTagPairPb `json:"custom_tags,omitempty"`
}

func endpointTagsFromPb(pb *endpointTagsPb) (*EndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTags{}

	var customTagsField []EndpointTagPair
	for _, itemPb := range pb.CustomTags {
		item, err := endpointTagPairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField = append(customTagsField, *item)
		}
	}
	st.CustomTags = customTagsField

	return st, nil
}

type EnumValue struct {
	// List of valid query parameter values, newline delimited.
	// Wire name: 'enum_options'
	EnumOptions string
	// If specified, allows multiple values to be selected for this parameter.
	// Wire name: 'multi_values_options'
	MultiValuesOptions *MultiValuesOptions
	// List of selected query parameter values.
	// Wire name: 'values'
	Values []string

	ForceSendFields []string `tf:"-"`
}

func enumValueToPb(st *EnumValue) (*enumValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enumValuePb{}
	pb.EnumOptions = st.EnumOptions

	multiValuesOptionsPb, err := multiValuesOptionsToPb(st.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsPb != nil {
		pb.MultiValuesOptions = multiValuesOptionsPb
	}

	pb.Values = st.Values

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enumValuePb struct {
	// List of valid query parameter values, newline delimited.
	EnumOptions string `json:"enum_options,omitempty"`
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions *multiValuesOptionsPb `json:"multi_values_options,omitempty"`
	// List of selected query parameter values.
	Values []string `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enumValueFromPb(pb *enumValuePb) (*EnumValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnumValue{}
	st.EnumOptions = pb.EnumOptions
	multiValuesOptionsField, err := multiValuesOptionsFromPb(pb.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsField != nil {
		st.MultiValuesOptions = multiValuesOptionsField
	}
	st.Values = pb.Values

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enumValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enumValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	// Wire name: 'byte_limit'
	ByteLimit int64
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	// Wire name: 'catalog'
	Catalog string

	// Wire name: 'disposition'
	Disposition Disposition
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
	Format Format
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	// Wire name: 'on_wait_timeout'
	OnWaitTimeout ExecuteStatementRequestOnWaitTimeout
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
	Parameters []StatementParameterListItem
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	// Wire name: 'row_limit'
	RowLimit int64
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	// Wire name: 'schema'
	Schema string
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`.
	// Wire name: 'statement'
	Statement string
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
	WaitTimeout string
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func executeStatementRequestToPb(st *ExecuteStatementRequest) (*executeStatementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &executeStatementRequestPb{}
	pb.ByteLimit = st.ByteLimit

	pb.Catalog = st.Catalog

	pb.Disposition = st.Disposition

	pb.Format = st.Format

	pb.OnWaitTimeout = st.OnWaitTimeout

	var parametersPb []statementParameterListItemPb
	for _, item := range st.Parameters {
		itemPb, err := statementParameterListItemToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	pb.RowLimit = st.RowLimit

	pb.Schema = st.Schema

	pb.Statement = st.Statement

	pb.WaitTimeout = st.WaitTimeout

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type executeStatementRequestPb struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	ByteLimit int64 `json:"byte_limit,omitempty"`
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog string `json:"catalog,omitempty"`

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
	Format Format `json:"format,omitempty"`
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
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
	Parameters []statementParameterListItemPb `json:"parameters,omitempty"`
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit int64 `json:"row_limit,omitempty"`
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema string `json:"schema,omitempty"`
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`.
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
	WaitTimeout string `json:"wait_timeout,omitempty"`
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId string `json:"warehouse_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func executeStatementRequestFromPb(pb *executeStatementRequestPb) (*ExecuteStatementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExecuteStatementRequest{}
	st.ByteLimit = pb.ByteLimit
	st.Catalog = pb.Catalog
	st.Disposition = pb.Disposition
	st.Format = pb.Format
	st.OnWaitTimeout = pb.OnWaitTimeout

	var parametersField []StatementParameterListItem
	for _, itemPb := range pb.Parameters {
		item, err := statementParameterListItemFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.RowLimit = pb.RowLimit
	st.Schema = pb.Schema
	st.Statement = pb.Statement
	st.WaitTimeout = pb.WaitTimeout
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *executeStatementRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st executeStatementRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// When `wait_timeout > 0s`, the call will block up to the specified time. If
// the statement execution doesn't finish within this time, `on_wait_timeout`
// determines whether the execution should continue or be canceled. When set to
// `CONTINUE`, the statement execution continues asynchronously and the call
// returns a statement ID which can be used for polling with
// :method:statementexecution/getStatement. When set to `CANCEL`, the statement
// execution is canceled and the call returns with a `CANCELED` state.
type ExecuteStatementRequestOnWaitTimeout string
type executeStatementRequestOnWaitTimeoutPb string

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

// Type always returns ExecuteStatementRequestOnWaitTimeout to satisfy [pflag.Value] interface
func (f *ExecuteStatementRequestOnWaitTimeout) Type() string {
	return "ExecuteStatementRequestOnWaitTimeout"
}

func executeStatementRequestOnWaitTimeoutToPb(st *ExecuteStatementRequestOnWaitTimeout) (*executeStatementRequestOnWaitTimeoutPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := executeStatementRequestOnWaitTimeoutPb(*st)
	return &pb, nil
}

func executeStatementRequestOnWaitTimeoutFromPb(pb *executeStatementRequestOnWaitTimeoutPb) (*ExecuteStatementRequestOnWaitTimeout, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExecuteStatementRequestOnWaitTimeout(*pb)
	return &st, nil
}

type ExternalLink struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	// Wire name: 'expiration'
	Expiration string

	// Wire name: 'external_link'
	ExternalLink string
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	// Wire name: 'http_headers'
	HttpHeaders map[string]string
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	// Wire name: 'next_chunk_index'
	NextChunkIndex int
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	// Wire name: 'next_chunk_internal_link'
	NextChunkInternalLink string
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset int64

	ForceSendFields []string `tf:"-"`
}

func externalLinkToPb(st *ExternalLink) (*externalLinkPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalLinkPb{}
	pb.ByteCount = st.ByteCount

	pb.ChunkIndex = st.ChunkIndex

	pb.Expiration = st.Expiration

	pb.ExternalLink = st.ExternalLink

	pb.HttpHeaders = st.HttpHeaders

	pb.NextChunkIndex = st.NextChunkIndex

	pb.NextChunkInternalLink = st.NextChunkInternalLink

	pb.RowCount = st.RowCount

	pb.RowOffset = st.RowOffset

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type externalLinkPb struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	Expiration string `json:"expiration,omitempty"`

	ExternalLink string `json:"external_link,omitempty"`
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders map[string]string `json:"http_headers,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalLinkFromPb(pb *externalLinkPb) (*ExternalLink, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalLink{}
	st.ByteCount = pb.ByteCount
	st.ChunkIndex = pb.ChunkIndex
	st.Expiration = pb.Expiration
	st.ExternalLink = pb.ExternalLink
	st.HttpHeaders = pb.HttpHeaders
	st.NextChunkIndex = pb.NextChunkIndex
	st.NextChunkInternalLink = pb.NextChunkInternalLink
	st.RowCount = pb.RowCount
	st.RowOffset = pb.RowOffset

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalLinkPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalLinkPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalQuerySource struct {
	// The canonical identifier for this SQL alert
	// Wire name: 'alert_id'
	AlertId string
	// The canonical identifier for this Lakeview dashboard
	// Wire name: 'dashboard_id'
	DashboardId string
	// The canonical identifier for this Genie space
	// Wire name: 'genie_space_id'
	GenieSpaceId string

	// Wire name: 'job_info'
	JobInfo *ExternalQuerySourceJobInfo
	// The canonical identifier for this legacy dashboard
	// Wire name: 'legacy_dashboard_id'
	LegacyDashboardId string
	// The canonical identifier for this notebook
	// Wire name: 'notebook_id'
	NotebookId string
	// The canonical identifier for this SQL query
	// Wire name: 'sql_query_id'
	SqlQueryId string

	ForceSendFields []string `tf:"-"`
}

func externalQuerySourceToPb(st *ExternalQuerySource) (*externalQuerySourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalQuerySourcePb{}
	pb.AlertId = st.AlertId

	pb.DashboardId = st.DashboardId

	pb.GenieSpaceId = st.GenieSpaceId

	jobInfoPb, err := externalQuerySourceJobInfoToPb(st.JobInfo)
	if err != nil {
		return nil, err
	}
	if jobInfoPb != nil {
		pb.JobInfo = jobInfoPb
	}

	pb.LegacyDashboardId = st.LegacyDashboardId

	pb.NotebookId = st.NotebookId

	pb.SqlQueryId = st.SqlQueryId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type externalQuerySourcePb struct {
	// The canonical identifier for this SQL alert
	AlertId string `json:"alert_id,omitempty"`
	// The canonical identifier for this Lakeview dashboard
	DashboardId string `json:"dashboard_id,omitempty"`
	// The canonical identifier for this Genie space
	GenieSpaceId string `json:"genie_space_id,omitempty"`

	JobInfo *externalQuerySourceJobInfoPb `json:"job_info,omitempty"`
	// The canonical identifier for this legacy dashboard
	LegacyDashboardId string `json:"legacy_dashboard_id,omitempty"`
	// The canonical identifier for this notebook
	NotebookId string `json:"notebook_id,omitempty"`
	// The canonical identifier for this SQL query
	SqlQueryId string `json:"sql_query_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalQuerySourceFromPb(pb *externalQuerySourcePb) (*ExternalQuerySource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalQuerySource{}
	st.AlertId = pb.AlertId
	st.DashboardId = pb.DashboardId
	st.GenieSpaceId = pb.GenieSpaceId
	jobInfoField, err := externalQuerySourceJobInfoFromPb(pb.JobInfo)
	if err != nil {
		return nil, err
	}
	if jobInfoField != nil {
		st.JobInfo = jobInfoField
	}
	st.LegacyDashboardId = pb.LegacyDashboardId
	st.NotebookId = pb.NotebookId
	st.SqlQueryId = pb.SqlQueryId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalQuerySourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalQuerySourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalQuerySourceJobInfo struct {
	// The canonical identifier for this job.
	// Wire name: 'job_id'
	JobId string
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	// Wire name: 'job_run_id'
	JobRunId string
	// The canonical identifier of the task run.
	// Wire name: 'job_task_run_id'
	JobTaskRunId string

	ForceSendFields []string `tf:"-"`
}

func externalQuerySourceJobInfoToPb(st *ExternalQuerySourceJobInfo) (*externalQuerySourceJobInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalQuerySourceJobInfoPb{}
	pb.JobId = st.JobId

	pb.JobRunId = st.JobRunId

	pb.JobTaskRunId = st.JobTaskRunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type externalQuerySourceJobInfoPb struct {
	// The canonical identifier for this job.
	JobId string `json:"job_id,omitempty"`
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	JobRunId string `json:"job_run_id,omitempty"`
	// The canonical identifier of the task run.
	JobTaskRunId string `json:"job_task_run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalQuerySourceJobInfoFromPb(pb *externalQuerySourceJobInfoPb) (*ExternalQuerySourceJobInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalQuerySourceJobInfo{}
	st.JobId = pb.JobId
	st.JobRunId = pb.JobRunId
	st.JobTaskRunId = pb.JobTaskRunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalQuerySourceJobInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalQuerySourceJobInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Format string
type formatPb string

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

// Get an alert
type GetAlertRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func getAlertRequestToPb(st *GetAlertRequest) (*getAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getAlertRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getAlertRequestFromPb(pb *getAlertRequestPb) (*GetAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAlertRequest{}
	st.Id = pb.Id

	return st, nil
}

// Get an alert
type GetAlertV2Request struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func getAlertV2RequestToPb(st *GetAlertV2Request) (*getAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertV2RequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getAlertV2RequestPb struct {
	Id string `json:"-" url:"-"`
}

func getAlertV2RequestFromPb(pb *getAlertV2RequestPb) (*GetAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAlertV2Request{}
	st.Id = pb.Id

	return st, nil
}

// Get an alert
type GetAlertsLegacyRequest struct {

	// Wire name: 'alert_id'
	AlertId string `tf:"-"`
}

func getAlertsLegacyRequestToPb(st *GetAlertsLegacyRequest) (*getAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertsLegacyRequestPb{}
	pb.AlertId = st.AlertId

	return pb, nil
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

type getAlertsLegacyRequestPb struct {
	AlertId string `json:"-" url:"-"`
}

func getAlertsLegacyRequestFromPb(pb *getAlertsLegacyRequestPb) (*GetAlertsLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAlertsLegacyRequest{}
	st.AlertId = pb.AlertId

	return st, nil
}

// Retrieve a definition
type GetDashboardRequest struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func getDashboardRequestToPb(st *GetDashboardRequest) (*getDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
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

type getDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func getDashboardRequestFromPb(pb *getDashboardRequestPb) (*GetDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

// Get object ACL
type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	// Wire name: 'objectId'
	ObjectId string `tf:"-"`
	// The type of object permissions to check.
	// Wire name: 'objectType'
	ObjectType ObjectTypePlural `tf:"-"`
}

func getDbsqlPermissionRequestToPb(st *GetDbsqlPermissionRequest) (*getDbsqlPermissionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDbsqlPermissionRequestPb{}
	pb.ObjectId = st.ObjectId

	pb.ObjectType = st.ObjectType

	return pb, nil
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

type getDbsqlPermissionRequestPb struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId string `json:"-" url:"-"`
	// The type of object permissions to check.
	ObjectType ObjectTypePlural `json:"-" url:"-"`
}

func getDbsqlPermissionRequestFromPb(pb *getDbsqlPermissionRequestPb) (*GetDbsqlPermissionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDbsqlPermissionRequest{}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

// Get a query definition.
type GetQueriesLegacyRequest struct {

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
}

func getQueriesLegacyRequestToPb(st *GetQueriesLegacyRequest) (*getQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
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

type getQueriesLegacyRequestPb struct {
	QueryId string `json:"-" url:"-"`
}

func getQueriesLegacyRequestFromPb(pb *getQueriesLegacyRequestPb) (*GetQueriesLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQueriesLegacyRequest{}
	st.QueryId = pb.QueryId

	return st, nil
}

// Get a query
type GetQueryRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func getQueryRequestToPb(st *GetQueryRequest) (*getQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQueryRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getQueryRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getQueryRequestFromPb(pb *getQueryRequestPb) (*GetQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQueryRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetResponse struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl
	// An object's type and UUID, separated by a forward slash (/) character.
	// Wire name: 'object_id'
	ObjectId string
	// A singular noun object type.
	// Wire name: 'object_type'
	ObjectType ObjectType

	ForceSendFields []string `tf:"-"`
}

func getResponseToPb(st *GetResponse) (*getResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getResponsePb{}

	var accessControlListPb []accessControlPb
	for _, item := range st.AccessControlList {
		itemPb, err := accessControlToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getResponsePb struct {
	AccessControlList []accessControlPb `json:"access_control_list,omitempty"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId string `json:"object_id,omitempty"`
	// A singular noun object type.
	ObjectType ObjectType `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getResponseFromPb(pb *getResponsePb) (*GetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetResponse{}

	var accessControlListField []AccessControl
	for _, itemPb := range pb.AccessControlList {
		item, err := accessControlFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get status, manifest, and result first chunk
type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string `tf:"-"`
}

func getStatementRequestToPb(st *GetStatementRequest) (*getStatementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatementRequestPb{}
	pb.StatementId = st.StatementId

	return pb, nil
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

type getStatementRequestPb struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
}

func getStatementRequestFromPb(pb *getStatementRequestPb) (*GetStatementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatementRequest{}
	st.StatementId = pb.StatementId

	return st, nil
}

// Get result chunk by index
type GetStatementResultChunkNRequest struct {

	// Wire name: 'chunk_index'
	ChunkIndex int `tf:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string `tf:"-"`
}

func getStatementResultChunkNRequestToPb(st *GetStatementResultChunkNRequest) (*getStatementResultChunkNRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatementResultChunkNRequestPb{}
	pb.ChunkIndex = st.ChunkIndex

	pb.StatementId = st.StatementId

	return pb, nil
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

type getStatementResultChunkNRequestPb struct {
	ChunkIndex int `json:"-" url:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"-" url:"-"`
}

func getStatementResultChunkNRequestFromPb(pb *getStatementResultChunkNRequestPb) (*GetStatementResultChunkNRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatementResultChunkNRequest{}
	st.ChunkIndex = pb.ChunkIndex
	st.StatementId = pb.StatementId

	return st, nil
}

// Get SQL warehouse permission levels
type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	// Wire name: 'warehouse_id'
	WarehouseId string `tf:"-"`
}

func getWarehousePermissionLevelsRequestToPb(st *GetWarehousePermissionLevelsRequest) (*getWarehousePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionLevelsRequestPb{}
	pb.WarehouseId = st.WarehouseId

	return pb, nil
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

type getWarehousePermissionLevelsRequestPb struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" url:"-"`
}

func getWarehousePermissionLevelsRequestFromPb(pb *getWarehousePermissionLevelsRequestPb) (*GetWarehousePermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehousePermissionLevelsRequest{}
	st.WarehouseId = pb.WarehouseId

	return st, nil
}

type GetWarehousePermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []WarehousePermissionsDescription
}

func getWarehousePermissionLevelsResponseToPb(st *GetWarehousePermissionLevelsResponse) (*getWarehousePermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionLevelsResponsePb{}

	var permissionLevelsPb []warehousePermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := warehousePermissionsDescriptionToPb(&item)
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

type getWarehousePermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []warehousePermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getWarehousePermissionLevelsResponseFromPb(pb *getWarehousePermissionLevelsResponsePb) (*GetWarehousePermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehousePermissionLevelsResponse{}

	var permissionLevelsField []WarehousePermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := warehousePermissionsDescriptionFromPb(&itemPb)
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

// Get SQL warehouse permissions
type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	// Wire name: 'warehouse_id'
	WarehouseId string `tf:"-"`
}

func getWarehousePermissionsRequestToPb(st *GetWarehousePermissionsRequest) (*getWarehousePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionsRequestPb{}
	pb.WarehouseId = st.WarehouseId

	return pb, nil
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

type getWarehousePermissionsRequestPb struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" url:"-"`
}

func getWarehousePermissionsRequestFromPb(pb *getWarehousePermissionsRequestPb) (*GetWarehousePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehousePermissionsRequest{}
	st.WarehouseId = pb.WarehouseId

	return st, nil
}

// Get warehouse info
type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func getWarehouseRequestToPb(st *GetWarehouseRequest) (*getWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehouseRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type getWarehouseRequestPb struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func getWarehouseRequestFromPb(pb *getWarehouseRequestPb) (*GetWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehouseRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	// Wire name: 'health'
	Health *EndpointHealth
	// unique identifier for warehouse
	// Wire name: 'id'
	Id string
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string
	// the jdbc connection string for this warehouse
	// Wire name: 'jdbc_url'
	JdbcUrl string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string
	// Deprecated. current number of active sessions for the warehouse
	// Wire name: 'num_active_sessions'
	NumActiveSessions int64
	// current number of clusters running for the service
	// Wire name: 'num_clusters'
	NumClusters int
	// ODBC parameters for the SQL warehouse
	// Wire name: 'odbc_params'
	OdbcParams *OdbcParams
	// Configurations whether the warehouse should use spot instances.
	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy
	// State of the warehouse
	// Wire name: 'state'
	State State
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	// Wire name: 'warehouse_type'
	WarehouseType GetWarehouseResponseWarehouseType

	ForceSendFields []string `tf:"-"`
}

func getWarehouseResponseToPb(st *GetWarehouseResponse) (*getWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehouseResponsePb{}
	pb.AutoStopMins = st.AutoStopMins

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	pb.ClusterSize = st.ClusterSize

	pb.CreatorName = st.CreatorName

	pb.EnablePhoton = st.EnablePhoton

	pb.EnableServerlessCompute = st.EnableServerlessCompute

	healthPb, err := endpointHealthToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	pb.Id = st.Id

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.JdbcUrl = st.JdbcUrl

	pb.MaxNumClusters = st.MaxNumClusters

	pb.MinNumClusters = st.MinNumClusters

	pb.Name = st.Name

	pb.NumActiveSessions = st.NumActiveSessions

	pb.NumClusters = st.NumClusters

	odbcParamsPb, err := odbcParamsToPb(st.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsPb != nil {
		pb.OdbcParams = odbcParamsPb
	}

	pb.SpotInstancePolicy = st.SpotInstancePolicy

	pb.State = st.State

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getWarehouseResponsePb struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int `json:"auto_stop_mins,omitempty"`
	// Channel Details
	Channel *channelPb `json:"channel,omitempty"`
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string `json:"cluster_size,omitempty"`
	// warehouse creator name
	CreatorName string `json:"creator_name,omitempty"`
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool `json:"enable_photon,omitempty"`
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool `json:"enable_serverless_compute,omitempty"`
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *endpointHealthPb `json:"health,omitempty"`
	// unique identifier for warehouse
	Id string `json:"id,omitempty"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// the jdbc connection string for this warehouse
	JdbcUrl string `json:"jdbc_url,omitempty"`
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int `json:"min_num_clusters,omitempty"`
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string `json:"name,omitempty"`
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions int64 `json:"num_active_sessions,omitempty"`
	// current number of clusters running for the service
	NumClusters int `json:"num_clusters,omitempty"`
	// ODBC parameters for the SQL warehouse
	OdbcParams *odbcParamsPb `json:"odbc_params,omitempty"`
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy `json:"spot_instance_policy,omitempty"`
	// State of the warehouse
	State State `json:"state,omitempty"`
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *endpointTagsPb `json:"tags,omitempty"`
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType GetWarehouseResponseWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getWarehouseResponseFromPb(pb *getWarehouseResponsePb) (*GetWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehouseResponse{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	st.ClusterSize = pb.ClusterSize
	st.CreatorName = pb.CreatorName
	st.EnablePhoton = pb.EnablePhoton
	st.EnableServerlessCompute = pb.EnableServerlessCompute
	healthField, err := endpointHealthFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	st.Id = pb.Id
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.JdbcUrl = pb.JdbcUrl
	st.MaxNumClusters = pb.MaxNumClusters
	st.MinNumClusters = pb.MinNumClusters
	st.Name = pb.Name
	st.NumActiveSessions = pb.NumActiveSessions
	st.NumClusters = pb.NumClusters
	odbcParamsField, err := odbcParamsFromPb(pb.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsField != nil {
		st.OdbcParams = odbcParamsField
	}
	st.SpotInstancePolicy = pb.SpotInstancePolicy
	st.State = pb.State
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	st.WarehouseType = pb.WarehouseType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getWarehouseResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getWarehouseResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless compute,
// you must set to `PRO` and also set the field `enable_serverless_compute` to
// `true`.
type GetWarehouseResponseWarehouseType string
type getWarehouseResponseWarehouseTypePb string

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

// Type always returns GetWarehouseResponseWarehouseType to satisfy [pflag.Value] interface
func (f *GetWarehouseResponseWarehouseType) Type() string {
	return "GetWarehouseResponseWarehouseType"
}

func getWarehouseResponseWarehouseTypeToPb(st *GetWarehouseResponseWarehouseType) (*getWarehouseResponseWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := getWarehouseResponseWarehouseTypePb(*st)
	return &pb, nil
}

func getWarehouseResponseWarehouseTypeFromPb(pb *getWarehouseResponseWarehouseTypePb) (*GetWarehouseResponseWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetWarehouseResponseWarehouseType(*pb)
	return &st, nil
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	// Wire name: 'channel'
	Channel *Channel
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'config_param'
	ConfigParam *RepeatedEndpointConfPairs
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	// Wire name: 'data_access_config'
	DataAccessConfig []EndpointConfPair
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	// Wire name: 'enabled_warehouse_types'
	EnabledWarehouseTypes []WarehouseTypePair
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'global_param'
	GlobalParam *RepeatedEndpointConfPairs
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	// Wire name: 'google_service_account'
	GoogleServiceAccount string
	// AWS Only: Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string
	// Security policy for warehouses
	// Wire name: 'security_policy'
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy
	// SQL configuration parameters
	// Wire name: 'sql_configuration_parameters'
	SqlConfigurationParameters *RepeatedEndpointConfPairs

	ForceSendFields []string `tf:"-"`
}

func getWorkspaceWarehouseConfigResponseToPb(st *GetWorkspaceWarehouseConfigResponse) (*getWorkspaceWarehouseConfigResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceWarehouseConfigResponsePb{}
	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	configParamPb, err := repeatedEndpointConfPairsToPb(st.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamPb != nil {
		pb.ConfigParam = configParamPb
	}

	var dataAccessConfigPb []endpointConfPairPb
	for _, item := range st.DataAccessConfig {
		itemPb, err := endpointConfPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataAccessConfigPb = append(dataAccessConfigPb, *itemPb)
		}
	}
	pb.DataAccessConfig = dataAccessConfigPb

	var enabledWarehouseTypesPb []warehouseTypePairPb
	for _, item := range st.EnabledWarehouseTypes {
		itemPb, err := warehouseTypePairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			enabledWarehouseTypesPb = append(enabledWarehouseTypesPb, *itemPb)
		}
	}
	pb.EnabledWarehouseTypes = enabledWarehouseTypesPb

	globalParamPb, err := repeatedEndpointConfPairsToPb(st.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamPb != nil {
		pb.GlobalParam = globalParamPb
	}

	pb.GoogleServiceAccount = st.GoogleServiceAccount

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.SecurityPolicy = st.SecurityPolicy

	sqlConfigurationParametersPb, err := repeatedEndpointConfPairsToPb(st.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersPb != nil {
		pb.SqlConfigurationParameters = sqlConfigurationParametersPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getWorkspaceWarehouseConfigResponsePb struct {
	// Optional: Channel selection details
	Channel *channelPb `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *repeatedEndpointConfPairsPb `json:"config_param,omitempty"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig []endpointConfPairPb `json:"data_access_config,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes []warehouseTypePairPb `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *repeatedEndpointConfPairsPb `json:"global_param,omitempty"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for warehouses
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy `json:"security_policy,omitempty"`
	// SQL configuration parameters
	SqlConfigurationParameters *repeatedEndpointConfPairsPb `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getWorkspaceWarehouseConfigResponseFromPb(pb *getWorkspaceWarehouseConfigResponsePb) (*GetWorkspaceWarehouseConfigResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceWarehouseConfigResponse{}
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	configParamField, err := repeatedEndpointConfPairsFromPb(pb.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamField != nil {
		st.ConfigParam = configParamField
	}

	var dataAccessConfigField []EndpointConfPair
	for _, itemPb := range pb.DataAccessConfig {
		item, err := endpointConfPairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataAccessConfigField = append(dataAccessConfigField, *item)
		}
	}
	st.DataAccessConfig = dataAccessConfigField

	var enabledWarehouseTypesField []WarehouseTypePair
	for _, itemPb := range pb.EnabledWarehouseTypes {
		item, err := warehouseTypePairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			enabledWarehouseTypesField = append(enabledWarehouseTypesField, *item)
		}
	}
	st.EnabledWarehouseTypes = enabledWarehouseTypesField
	globalParamField, err := repeatedEndpointConfPairsFromPb(pb.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamField != nil {
		st.GlobalParam = globalParamField
	}
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.SecurityPolicy = pb.SecurityPolicy
	sqlConfigurationParametersField, err := repeatedEndpointConfPairsFromPb(pb.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersField != nil {
		st.SqlConfigurationParameters = sqlConfigurationParametersField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getWorkspaceWarehouseConfigResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getWorkspaceWarehouseConfigResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Security policy for warehouses
type GetWorkspaceWarehouseConfigResponseSecurityPolicy string
type getWorkspaceWarehouseConfigResponseSecurityPolicyPb string

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

// Type always returns GetWorkspaceWarehouseConfigResponseSecurityPolicy to satisfy [pflag.Value] interface
func (f *GetWorkspaceWarehouseConfigResponseSecurityPolicy) Type() string {
	return "GetWorkspaceWarehouseConfigResponseSecurityPolicy"
}

func getWorkspaceWarehouseConfigResponseSecurityPolicyToPb(st *GetWorkspaceWarehouseConfigResponseSecurityPolicy) (*getWorkspaceWarehouseConfigResponseSecurityPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := getWorkspaceWarehouseConfigResponseSecurityPolicyPb(*st)
	return &pb, nil
}

func getWorkspaceWarehouseConfigResponseSecurityPolicyFromPb(pb *getWorkspaceWarehouseConfigResponseSecurityPolicyPb) (*GetWorkspaceWarehouseConfigResponseSecurityPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetWorkspaceWarehouseConfigResponseSecurityPolicy(*pb)
	return &st, nil
}

type LegacyAlert struct {
	// Timestamp when the alert was created.
	// Wire name: 'created_at'
	CreatedAt string
	// Alert ID.
	// Wire name: 'id'
	Id string
	// Timestamp when the alert was last triggered.
	// Wire name: 'last_triggered_at'
	LastTriggeredAt string
	// Name of the alert.
	// Wire name: 'name'
	Name string
	// Alert configuration options.
	// Wire name: 'options'
	Options *AlertOptions
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string

	// Wire name: 'query'
	Query *AlertQuery
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm int
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	// Wire name: 'state'
	State LegacyAlertState
	// Timestamp when the alert was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string

	// Wire name: 'user'
	User *User

	ForceSendFields []string `tf:"-"`
}

func legacyAlertToPb(st *LegacyAlert) (*legacyAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &legacyAlertPb{}
	pb.CreatedAt = st.CreatedAt

	pb.Id = st.Id

	pb.LastTriggeredAt = st.LastTriggeredAt

	pb.Name = st.Name

	optionsPb, err := alertOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	pb.Parent = st.Parent

	queryPb, err := alertQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	pb.Rearm = st.Rearm

	pb.State = st.State

	pb.UpdatedAt = st.UpdatedAt

	userPb, err := userToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type legacyAlertPb struct {
	// Timestamp when the alert was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Alert ID.
	Id string `json:"id,omitempty"`
	// Timestamp when the alert was last triggered.
	LastTriggeredAt string `json:"last_triggered_at,omitempty"`
	// Name of the alert.
	Name string `json:"name,omitempty"`
	// Alert configuration options.
	Options *alertOptionsPb `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`

	Query *alertQueryPb `json:"query,omitempty"`
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int `json:"rearm,omitempty"`
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	State LegacyAlertState `json:"state,omitempty"`
	// Timestamp when the alert was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *userPb `json:"user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func legacyAlertFromPb(pb *legacyAlertPb) (*LegacyAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LegacyAlert{}
	st.CreatedAt = pb.CreatedAt
	st.Id = pb.Id
	st.LastTriggeredAt = pb.LastTriggeredAt
	st.Name = pb.Name
	optionsField, err := alertOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	st.Parent = pb.Parent
	queryField, err := alertQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	st.Rearm = pb.Rearm
	st.State = pb.State
	st.UpdatedAt = pb.UpdatedAt
	userField, err := userFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *legacyAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st legacyAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// State of the alert. Possible values are: `unknown` (yet to be evaluated),
// `triggered` (evaluated and fulfilled trigger conditions), or `ok` (evaluated
// and did not fulfill trigger conditions).
type LegacyAlertState string
type legacyAlertStatePb string

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

// Type always returns LegacyAlertState to satisfy [pflag.Value] interface
func (f *LegacyAlertState) Type() string {
	return "LegacyAlertState"
}

func legacyAlertStateToPb(st *LegacyAlertState) (*legacyAlertStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := legacyAlertStatePb(*st)
	return &pb, nil
}

func legacyAlertStateFromPb(pb *legacyAlertStatePb) (*LegacyAlertState, error) {
	if pb == nil {
		return nil, nil
	}
	st := LegacyAlertState(*pb)
	return &st, nil
}

type LegacyQuery struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	// Wire name: 'can_edit'
	CanEdit bool
	// The timestamp when this query was created.
	// Wire name: 'created_at'
	CreatedAt string
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// Query ID.
	// Wire name: 'id'
	Id string
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	// Wire name: 'is_draft'
	IsDraft bool
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	// Wire name: 'is_favorite'
	IsFavorite bool
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	// Wire name: 'is_safe'
	IsSafe bool

	// Wire name: 'last_modified_by'
	LastModifiedBy *User
	// The ID of the user who last saved changes to this query.
	// Wire name: 'last_modified_by_id'
	LastModifiedById int
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	// Wire name: 'latest_query_data_id'
	LatestQueryDataId string
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string

	// Wire name: 'options'
	Options *QueryOptions
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_tier'
	PermissionTier PermissionLevel
	// The text of the query to be run.
	// Wire name: 'query'
	Query string
	// A SHA-256 hash of the query text along with the authenticated user ID.
	// Wire name: 'query_hash'
	QueryHash string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole

	// Wire name: 'tags'
	Tags []string
	// The timestamp at which this query was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string

	// Wire name: 'user'
	User *User
	// The ID of the user who owns the query.
	// Wire name: 'user_id'
	UserId int

	// Wire name: 'visualizations'
	Visualizations []LegacyVisualization

	ForceSendFields []string `tf:"-"`
}

func legacyQueryToPb(st *LegacyQuery) (*legacyQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &legacyQueryPb{}
	pb.CanEdit = st.CanEdit

	pb.CreatedAt = st.CreatedAt

	pb.DataSourceId = st.DataSourceId

	pb.Description = st.Description

	pb.Id = st.Id

	pb.IsArchived = st.IsArchived

	pb.IsDraft = st.IsDraft

	pb.IsFavorite = st.IsFavorite

	pb.IsSafe = st.IsSafe

	lastModifiedByPb, err := userToPb(st.LastModifiedBy)
	if err != nil {
		return nil, err
	}
	if lastModifiedByPb != nil {
		pb.LastModifiedBy = lastModifiedByPb
	}

	pb.LastModifiedById = st.LastModifiedById

	pb.LatestQueryDataId = st.LatestQueryDataId

	pb.Name = st.Name

	optionsPb, err := queryOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	pb.Parent = st.Parent

	pb.PermissionTier = st.PermissionTier

	pb.Query = st.Query

	pb.QueryHash = st.QueryHash

	pb.RunAsRole = st.RunAsRole

	pb.Tags = st.Tags

	pb.UpdatedAt = st.UpdatedAt

	userPb, err := userToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}

	pb.UserId = st.UserId

	var visualizationsPb []legacyVisualizationPb
	for _, item := range st.Visualizations {
		itemPb, err := legacyVisualizationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			visualizationsPb = append(visualizationsPb, *itemPb)
		}
	}
	pb.Visualizations = visualizationsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type legacyQueryPb struct {
	// Describes whether the authenticated user is allowed to edit the
	// definition of this query.
	CanEdit bool `json:"can_edit,omitempty"`
	// The timestamp when this query was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Query ID.
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived bool `json:"is_archived,omitempty"`
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft bool `json:"is_draft,omitempty"`
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite bool `json:"is_favorite,omitempty"`
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe bool `json:"is_safe,omitempty"`

	LastModifiedBy *userPb `json:"last_modified_by,omitempty"`
	// The ID of the user who last saved changes to this query.
	LastModifiedById int `json:"last_modified_by_id,omitempty"`
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId string `json:"latest_query_data_id,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`

	Options *queryOptionsPb `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel `json:"permission_tier,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash string `json:"query_hash,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// The timestamp at which this query was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	User *userPb `json:"user,omitempty"`
	// The ID of the user who owns the query.
	UserId int `json:"user_id,omitempty"`

	Visualizations []legacyVisualizationPb `json:"visualizations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func legacyQueryFromPb(pb *legacyQueryPb) (*LegacyQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LegacyQuery{}
	st.CanEdit = pb.CanEdit
	st.CreatedAt = pb.CreatedAt
	st.DataSourceId = pb.DataSourceId
	st.Description = pb.Description
	st.Id = pb.Id
	st.IsArchived = pb.IsArchived
	st.IsDraft = pb.IsDraft
	st.IsFavorite = pb.IsFavorite
	st.IsSafe = pb.IsSafe
	lastModifiedByField, err := userFromPb(pb.LastModifiedBy)
	if err != nil {
		return nil, err
	}
	if lastModifiedByField != nil {
		st.LastModifiedBy = lastModifiedByField
	}
	st.LastModifiedById = pb.LastModifiedById
	st.LatestQueryDataId = pb.LatestQueryDataId
	st.Name = pb.Name
	optionsField, err := queryOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	st.Parent = pb.Parent
	st.PermissionTier = pb.PermissionTier
	st.Query = pb.Query
	st.QueryHash = pb.QueryHash
	st.RunAsRole = pb.RunAsRole
	st.Tags = pb.Tags
	st.UpdatedAt = pb.UpdatedAt
	userField, err := userFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}
	st.UserId = pb.UserId

	var visualizationsField []LegacyVisualization
	for _, itemPb := range pb.Visualizations {
		item, err := legacyVisualizationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			visualizationsField = append(visualizationsField, *item)
		}
	}
	st.Visualizations = visualizationsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *legacyQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st legacyQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization struct {

	// Wire name: 'created_at'
	CreatedAt string
	// A short description of this visualization. This is not displayed in the
	// UI.
	// Wire name: 'description'
	Description string
	// The UUID for this visualization.
	// Wire name: 'id'
	Id string
	// The name of the visualization that appears on dashboards and the query
	// screen.
	// Wire name: 'name'
	Name string
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	// Wire name: 'options'
	Options any

	// Wire name: 'query'
	Query *LegacyQuery
	// The type of visualization: chart, table, pivot table, and so on.
	// Wire name: 'type'
	Type string

	// Wire name: 'updated_at'
	UpdatedAt string

	ForceSendFields []string `tf:"-"`
}

func legacyVisualizationToPb(st *LegacyVisualization) (*legacyVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &legacyVisualizationPb{}
	pb.CreatedAt = st.CreatedAt

	pb.Description = st.Description

	pb.Id = st.Id

	pb.Name = st.Name

	pb.Options = st.Options

	queryPb, err := legacyQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	pb.Type = st.Type

	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type legacyVisualizationPb struct {
	CreatedAt string `json:"created_at,omitempty"`
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description string `json:"description,omitempty"`
	// The UUID for this visualization.
	Id string `json:"id,omitempty"`
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name string `json:"name,omitempty"`
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any `json:"options,omitempty"`

	Query *legacyQueryPb `json:"query,omitempty"`
	// The type of visualization: chart, table, pivot table, and so on.
	Type string `json:"type,omitempty"`

	UpdatedAt string `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func legacyVisualizationFromPb(pb *legacyVisualizationPb) (*LegacyVisualization, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LegacyVisualization{}
	st.CreatedAt = pb.CreatedAt
	st.Description = pb.Description
	st.Id = pb.Id
	st.Name = pb.Name
	st.Options = pb.Options
	queryField, err := legacyQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	st.Type = pb.Type
	st.UpdatedAt = pb.UpdatedAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *legacyVisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st legacyVisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LifecycleState string
type lifecycleStatePb string

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

// Type always returns LifecycleState to satisfy [pflag.Value] interface
func (f *LifecycleState) Type() string {
	return "LifecycleState"
}

func lifecycleStateToPb(st *LifecycleState) (*lifecycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := lifecycleStatePb(*st)
	return &pb, nil
}

func lifecycleStateFromPb(pb *lifecycleStatePb) (*LifecycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := LifecycleState(*pb)
	return &st, nil
}

// List alerts
type ListAlertsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listAlertsRequestToPb(st *ListAlertsRequest) (*listAlertsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsRequestPb{}
	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAlertsRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsRequestFromPb(pb *listAlertsRequestPb) (*ListAlertsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'results'
	Results []ListAlertsResponseAlert

	ForceSendFields []string `tf:"-"`
}

func listAlertsResponseToPb(st *ListAlertsResponse) (*listAlertsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []listAlertsResponseAlertPb
	for _, item := range st.Results {
		itemPb, err := listAlertsResponseAlertToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAlertsResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Results []listAlertsResponseAlertPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsResponseFromPb(pb *listAlertsResponsePb) (*ListAlertsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsResponse{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []ListAlertsResponseAlert
	for _, itemPb := range pb.Results {
		item, err := listAlertsResponseAlertFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsResponseAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime string
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string
	// The workspace state of the alert. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	// Wire name: 'state'
	State AlertState
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	// Wire name: 'trigger_time'
	TriggerTime string
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime string

	ForceSendFields []string `tf:"-"`
}

func listAlertsResponseAlertToPb(st *ListAlertsResponseAlert) (*listAlertsResponseAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsResponseAlertPb{}
	conditionPb, err := alertConditionToPb(st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = conditionPb
	}

	pb.CreateTime = st.CreateTime

	pb.CustomBody = st.CustomBody

	pb.CustomSubject = st.CustomSubject

	pb.DisplayName = st.DisplayName

	pb.Id = st.Id

	pb.LifecycleState = st.LifecycleState

	pb.NotifyOnOk = st.NotifyOnOk

	pb.OwnerUserName = st.OwnerUserName

	pb.QueryId = st.QueryId

	pb.SecondsToRetrigger = st.SecondsToRetrigger

	pb.State = st.State

	pb.TriggerTime = st.TriggerTime

	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAlertsResponseAlertPb struct {
	// Trigger conditions of the alert.
	Condition *alertConditionPb `json:"condition,omitempty"`
	// The timestamp indicating when the alert was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the alert.
	Id string `json:"id,omitempty"`
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// UUID of the query attached to the alert.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State AlertState `json:"state,omitempty"`
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime string `json:"trigger_time,omitempty"`
	// The timestamp indicating when the alert was updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsResponseAlertFromPb(pb *listAlertsResponseAlertPb) (*ListAlertsResponseAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsResponseAlert{}
	conditionField, err := alertConditionFromPb(pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = conditionField
	}
	st.CreateTime = pb.CreateTime
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.LifecycleState = pb.LifecycleState
	st.NotifyOnOk = pb.NotifyOnOk
	st.OwnerUserName = pb.OwnerUserName
	st.QueryId = pb.QueryId
	st.SecondsToRetrigger = pb.SecondsToRetrigger
	st.State = pb.State
	st.TriggerTime = pb.TriggerTime
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsResponseAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsResponseAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List alerts
type ListAlertsV2Request struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listAlertsV2RequestToPb(st *ListAlertsV2Request) (*listAlertsV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsV2RequestPb{}
	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAlertsV2RequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsV2RequestFromPb(pb *listAlertsV2RequestPb) (*ListAlertsV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsV2Request{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsV2RequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsV2RequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsV2Response struct {

	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'results'
	Results []ListAlertsV2ResponseAlert

	ForceSendFields []string `tf:"-"`
}

func listAlertsV2ResponseToPb(st *ListAlertsV2Response) (*listAlertsV2ResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsV2ResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []listAlertsV2ResponseAlertPb
	for _, item := range st.Results {
		itemPb, err := listAlertsV2ResponseAlertToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listAlertsV2ResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Results []listAlertsV2ResponseAlertPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsV2ResponseFromPb(pb *listAlertsV2ResponsePb) (*ListAlertsV2Response, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsV2Response{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []ListAlertsV2ResponseAlert
	for _, itemPb := range pb.Results {
		item, err := listAlertsV2ResponseAlertFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsV2ResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsV2ResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsV2ResponseAlert struct {
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime string
	// Custom description for the alert. support mustache template.
	// Wire name: 'custom_description'
	CustomDescription string
	// Custom summary for the alert. support mustache template.
	// Wire name: 'custom_summary'
	CustomSummary string
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string

	// Wire name: 'evaluation'
	Evaluation *AlertV2Evaluation
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string
	// The run as username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'run_as_user_name'
	RunAsUserName string

	// Wire name: 'schedule'
	Schedule *CronSchedule
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime string
	// ID of the SQL warehouse attached to the alert.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func listAlertsV2ResponseAlertToPb(st *ListAlertsV2ResponseAlert) (*listAlertsV2ResponseAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsV2ResponseAlertPb{}
	pb.CreateTime = st.CreateTime

	pb.CustomDescription = st.CustomDescription

	pb.CustomSummary = st.CustomSummary

	pb.DisplayName = st.DisplayName

	evaluationPb, err := alertV2EvaluationToPb(st.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationPb != nil {
		pb.Evaluation = evaluationPb
	}

	pb.Id = st.Id

	pb.LifecycleState = st.LifecycleState

	pb.OwnerUserName = st.OwnerUserName

	pb.QueryText = st.QueryText

	pb.RunAsUserName = st.RunAsUserName

	schedulePb, err := cronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}

	pb.UpdateTime = st.UpdateTime

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListAlertsV2ResponseAlert) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAlertsV2ResponseAlertPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAlertsV2ResponseAlertFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAlertsV2ResponseAlert) MarshalJSON() ([]byte, error) {
	pb, err := listAlertsV2ResponseAlertToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAlertsV2ResponseAlertPb struct {
	// The timestamp indicating when the alert was created.
	CreateTime string `json:"create_time,omitempty"`
	// Custom description for the alert. support mustache template.
	CustomDescription string `json:"custom_description,omitempty"`
	// Custom summary for the alert. support mustache template.
	CustomSummary string `json:"custom_summary,omitempty"`
	// The display name of the alert.
	DisplayName string `json:"display_name,omitempty"`

	Evaluation *alertV2EvaluationPb `json:"evaluation,omitempty"`
	// UUID identifying the alert.
	Id string `json:"id,omitempty"`
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// The run as username. This field is set to "Unavailable" if the user has
	// been deleted.
	RunAsUserName string `json:"run_as_user_name,omitempty"`

	Schedule *cronSchedulePb `json:"schedule,omitempty"`
	// The timestamp indicating when the alert was updated.
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the alert.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsV2ResponseAlertFromPb(pb *listAlertsV2ResponseAlertPb) (*ListAlertsV2ResponseAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsV2ResponseAlert{}
	st.CreateTime = pb.CreateTime
	st.CustomDescription = pb.CustomDescription
	st.CustomSummary = pb.CustomSummary
	st.DisplayName = pb.DisplayName
	evaluationField, err := alertV2EvaluationFromPb(pb.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationField != nil {
		st.Evaluation = evaluationField
	}
	st.Id = pb.Id
	st.LifecycleState = pb.LifecycleState
	st.OwnerUserName = pb.OwnerUserName
	st.QueryText = pb.QueryText
	st.RunAsUserName = pb.RunAsUserName
	scheduleField, err := cronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	st.UpdateTime = pb.UpdateTime
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsV2ResponseAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsV2ResponseAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get dashboard objects
type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	// Wire name: 'order'
	Order ListOrder `tf:"-"`
	// Page number to retrieve.
	// Wire name: 'page'
	Page int `tf:"-"`
	// Number of dashboards to return per page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Full text search term.
	// Wire name: 'q'
	Q string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listDashboardsRequestToPb(st *ListDashboardsRequest) (*listDashboardsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDashboardsRequestPb{}
	pb.Order = st.Order

	pb.Page = st.Page

	pb.PageSize = st.PageSize

	pb.Q = st.Q

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listDashboardsRequestPb struct {
	// Name of dashboard attribute to order by.
	Order ListOrder `json:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page int `json:"-" url:"page,omitempty"`
	// Number of dashboards to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Full text search term.
	Q string `json:"-" url:"q,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDashboardsRequestFromPb(pb *listDashboardsRequestPb) (*ListDashboardsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsRequest{}
	st.Order = pb.Order
	st.Page = pb.Page
	st.PageSize = pb.PageSize
	st.Q = pb.Q

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDashboardsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDashboardsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListOrder string
type listOrderPb string

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

// Type always returns ListOrder to satisfy [pflag.Value] interface
func (f *ListOrder) Type() string {
	return "ListOrder"
}

func listOrderToPb(st *ListOrder) (*listOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listOrderPb(*st)
	return &pb, nil
}

func listOrderFromPb(pb *listOrderPb) (*ListOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListOrder(*pb)
	return &st, nil
}

// Get a list of queries
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
	// Wire name: 'order'
	Order string `tf:"-"`
	// Page number to retrieve.
	// Wire name: 'page'
	Page int `tf:"-"`
	// Number of queries to return per page.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Full text search term
	// Wire name: 'q'
	Q string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listQueriesLegacyRequestToPb(st *ListQueriesLegacyRequest) (*listQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueriesLegacyRequestPb{}
	pb.Order = st.Order

	pb.Page = st.Page

	pb.PageSize = st.PageSize

	pb.Q = st.Q

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listQueriesLegacyRequestPb struct {
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
	Order string `json:"-" url:"order,omitempty"`
	// Page number to retrieve.
	Page int `json:"-" url:"page,omitempty"`
	// Number of queries to return per page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Full text search term
	Q string `json:"-" url:"q,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueriesLegacyRequestFromPb(pb *listQueriesLegacyRequestPb) (*ListQueriesLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueriesLegacyRequest{}
	st.Order = pb.Order
	st.Page = pb.Page
	st.PageSize = pb.PageSize
	st.Q = pb.Q

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueriesLegacyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueriesLegacyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List queries
type ListQueriesRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listQueriesRequestToPb(st *ListQueriesRequest) (*listQueriesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueriesRequestPb{}
	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listQueriesRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueriesRequestFromPb(pb *listQueriesRequestPb) (*ListQueriesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueriesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueriesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueriesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	// Wire name: 'has_next_page'
	HasNextPage bool
	// A token that can be used to get the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'res'
	Res []QueryInfo

	ForceSendFields []string `tf:"-"`
}

func listQueriesResponseToPb(st *ListQueriesResponse) (*listQueriesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueriesResponsePb{}
	pb.HasNextPage = st.HasNextPage

	pb.NextPageToken = st.NextPageToken

	var resPb []queryInfoPb
	for _, item := range st.Res {
		itemPb, err := queryInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resPb = append(resPb, *itemPb)
		}
	}
	pb.Res = resPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listQueriesResponsePb struct {
	// Whether there is another page of results.
	HasNextPage bool `json:"has_next_page,omitempty"`
	// A token that can be used to get the next page of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Res []queryInfoPb `json:"res,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueriesResponseFromPb(pb *listQueriesResponsePb) (*ListQueriesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueriesResponse{}
	st.HasNextPage = pb.HasNextPage
	st.NextPageToken = pb.NextPageToken

	var resField []QueryInfo
	for _, itemPb := range pb.Res {
		item, err := queryInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resField = append(resField, *item)
		}
	}
	st.Res = resField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueriesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueriesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List Queries
type ListQueryHistoryRequest struct {
	// A filter to limit query history results. This field is optional.
	// Wire name: 'filter_by'
	FilterBy *QueryFilter `tf:"-"`
	// Whether to include the query metrics with each query. Only use this for a
	// small subset of queries (max_results). Defaults to false.
	// Wire name: 'include_metrics'
	IncludeMetrics bool `tf:"-"`
	// Limit the number of results returned in one page. Must be less than 1000
	// and the default is 100.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// A token that can be used to get the next page of results. The token can
	// contains characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B. This field is
	// optional.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listQueryHistoryRequestToPb(st *ListQueryHistoryRequest) (*listQueryHistoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueryHistoryRequestPb{}
	filterByPb, err := queryFilterToPb(st.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByPb != nil {
		pb.FilterBy = filterByPb
	}

	pb.IncludeMetrics = st.IncludeMetrics

	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listQueryHistoryRequestPb struct {
	// A filter to limit query history results. This field is optional.
	FilterBy *queryFilterPb `json:"-" url:"filter_by,omitempty"`
	// Whether to include the query metrics with each query. Only use this for a
	// small subset of queries (max_results). Defaults to false.
	IncludeMetrics bool `json:"-" url:"include_metrics,omitempty"`
	// Limit the number of results returned in one page. Must be less than 1000
	// and the default is 100.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// A token that can be used to get the next page of results. The token can
	// contains characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B. This field is
	// optional.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueryHistoryRequestFromPb(pb *listQueryHistoryRequestPb) (*ListQueryHistoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryHistoryRequest{}
	filterByField, err := queryFilterFromPb(pb.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByField != nil {
		st.FilterBy = filterByField
	}
	st.IncludeMetrics = pb.IncludeMetrics
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueryHistoryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueryHistoryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueryObjectsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'results'
	Results []ListQueryObjectsResponseQuery

	ForceSendFields []string `tf:"-"`
}

func listQueryObjectsResponseToPb(st *ListQueryObjectsResponse) (*listQueryObjectsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueryObjectsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []listQueryObjectsResponseQueryPb
	for _, item := range st.Results {
		itemPb, err := listQueryObjectsResponseQueryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listQueryObjectsResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Results []listQueryObjectsResponseQueryPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueryObjectsResponseFromPb(pb *listQueryObjectsResponsePb) (*ListQueryObjectsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryObjectsResponse{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []ListQueryObjectsResponseQuery
	for _, itemPb := range pb.Results {
		item, err := listQueryObjectsResponseQueryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueryObjectsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueryObjectsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueryObjectsResponseQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string
	// Timestamp when this query was created.
	// Wire name: 'create_time'
	CreateTime string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying the query.
	// Wire name: 'id'
	Id string
	// Username of the user who last saved changes to this query.
	// Wire name: 'last_modifier_user_name'
	LastModifierUserName string
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string

	// Wire name: 'tags'
	Tags []string
	// Timestamp when this query was last updated.
	// Wire name: 'update_time'
	UpdateTime string
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func listQueryObjectsResponseQueryToPb(st *ListQueryObjectsResponseQuery) (*listQueryObjectsResponseQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueryObjectsResponseQueryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit

	pb.Catalog = st.Catalog

	pb.CreateTime = st.CreateTime

	pb.Description = st.Description

	pb.DisplayName = st.DisplayName

	pb.Id = st.Id

	pb.LastModifierUserName = st.LastModifierUserName

	pb.LifecycleState = st.LifecycleState

	pb.OwnerUserName = st.OwnerUserName

	var parametersPb []queryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := queryParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	pb.QueryText = st.QueryText

	pb.RunAsMode = st.RunAsMode

	pb.Schema = st.Schema

	pb.Tags = st.Tags

	pb.UpdateTime = st.UpdateTime

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listQueryObjectsResponseQueryPb struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// Timestamp when this query was created.
	CreateTime string `json:"create_time,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the query.
	Id string `json:"id,omitempty"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName string `json:"last_modifier_user_name,omitempty"`
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Username of the user that owns the query.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	Parameters []queryParameterPb `json:"parameters,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// Timestamp when this query was last updated.
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueryObjectsResponseQueryFromPb(pb *listQueryObjectsResponseQueryPb) (*ListQueryObjectsResponseQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryObjectsResponseQuery{}
	st.ApplyAutoLimit = pb.ApplyAutoLimit
	st.Catalog = pb.Catalog
	st.CreateTime = pb.CreateTime
	st.Description = pb.Description
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.LastModifierUserName = pb.LastModifierUserName
	st.LifecycleState = pb.LifecycleState
	st.OwnerUserName = pb.OwnerUserName

	var parametersField []QueryParameter
	for _, itemPb := range pb.Parameters {
		item, err := queryParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.QueryText = pb.QueryText
	st.RunAsMode = pb.RunAsMode
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	st.UpdateTime = pb.UpdateTime
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueryObjectsResponseQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueryObjectsResponseQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListResponse struct {
	// The total number of dashboards.
	// Wire name: 'count'
	Count int
	// The current page being displayed.
	// Wire name: 'page'
	Page int
	// The number of dashboards per page.
	// Wire name: 'page_size'
	PageSize int
	// List of dashboards returned.
	// Wire name: 'results'
	Results []Dashboard

	ForceSendFields []string `tf:"-"`
}

func listResponseToPb(st *ListResponse) (*listResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listResponsePb{}
	pb.Count = st.Count

	pb.Page = st.Page

	pb.PageSize = st.PageSize

	var resultsPb []dashboardPb
	for _, item := range st.Results {
		itemPb, err := dashboardToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listResponsePb struct {
	// The total number of dashboards.
	Count int `json:"count,omitempty"`
	// The current page being displayed.
	Page int `json:"page,omitempty"`
	// The number of dashboards per page.
	PageSize int `json:"page_size,omitempty"`
	// List of dashboards returned.
	Results []dashboardPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listResponseFromPb(pb *listResponsePb) (*ListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListResponse{}
	st.Count = pb.Count
	st.Page = pb.Page
	st.PageSize = pb.PageSize

	var resultsField []Dashboard
	for _, itemPb := range pb.Results {
		item, err := dashboardFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List visualizations on a query
type ListVisualizationsForQueryRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listVisualizationsForQueryRequestToPb(st *ListVisualizationsForQueryRequest) (*listVisualizationsForQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVisualizationsForQueryRequestPb{}
	pb.Id = st.Id

	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listVisualizationsForQueryRequestPb struct {
	Id string `json:"-" url:"-"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVisualizationsForQueryRequestFromPb(pb *listVisualizationsForQueryRequestPb) (*ListVisualizationsForQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVisualizationsForQueryRequest{}
	st.Id = pb.Id
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVisualizationsForQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVisualizationsForQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListVisualizationsForQueryResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'results'
	Results []Visualization

	ForceSendFields []string `tf:"-"`
}

func listVisualizationsForQueryResponseToPb(st *ListVisualizationsForQueryResponse) (*listVisualizationsForQueryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVisualizationsForQueryResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []visualizationPb
	for _, item := range st.Results {
		itemPb, err := visualizationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listVisualizationsForQueryResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Results []visualizationPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVisualizationsForQueryResponseFromPb(pb *listVisualizationsForQueryResponsePb) (*ListVisualizationsForQueryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVisualizationsForQueryResponse{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []Visualization
	for _, itemPb := range pb.Results {
		item, err := visualizationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVisualizationsForQueryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVisualizationsForQueryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List warehouses
type ListWarehousesRequest struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	// Wire name: 'run_as_user_id'
	RunAsUserId int `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listWarehousesRequestToPb(st *ListWarehousesRequest) (*listWarehousesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWarehousesRequestPb{}
	pb.RunAsUserId = st.RunAsUserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listWarehousesRequestPb struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	RunAsUserId int `json:"-" url:"run_as_user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listWarehousesRequestFromPb(pb *listWarehousesRequestPb) (*ListWarehousesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWarehousesRequest{}
	st.RunAsUserId = pb.RunAsUserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listWarehousesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listWarehousesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListWarehousesResponse struct {
	// A list of warehouses and their configurations.
	// Wire name: 'warehouses'
	Warehouses []EndpointInfo
}

func listWarehousesResponseToPb(st *ListWarehousesResponse) (*listWarehousesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWarehousesResponsePb{}

	var warehousesPb []endpointInfoPb
	for _, item := range st.Warehouses {
		itemPb, err := endpointInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			warehousesPb = append(warehousesPb, *itemPb)
		}
	}
	pb.Warehouses = warehousesPb

	return pb, nil
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

type listWarehousesResponsePb struct {
	// A list of warehouses and their configurations.
	Warehouses []endpointInfoPb `json:"warehouses,omitempty"`
}

func listWarehousesResponseFromPb(pb *listWarehousesResponsePb) (*ListWarehousesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWarehousesResponse{}

	var warehousesField []EndpointInfo
	for _, itemPb := range pb.Warehouses {
		item, err := endpointInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			warehousesField = append(warehousesField, *item)
		}
	}
	st.Warehouses = warehousesField

	return st, nil
}

type MultiValuesOptions struct {
	// Character that prefixes each selected parameter value.
	// Wire name: 'prefix'
	Prefix string
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	// Wire name: 'separator'
	Separator string
	// Character that suffixes each selected parameter value.
	// Wire name: 'suffix'
	Suffix string

	ForceSendFields []string `tf:"-"`
}

func multiValuesOptionsToPb(st *MultiValuesOptions) (*multiValuesOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &multiValuesOptionsPb{}
	pb.Prefix = st.Prefix

	pb.Separator = st.Separator

	pb.Suffix = st.Suffix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type multiValuesOptionsPb struct {
	// Character that prefixes each selected parameter value.
	Prefix string `json:"prefix,omitempty"`
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator string `json:"separator,omitempty"`
	// Character that suffixes each selected parameter value.
	Suffix string `json:"suffix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func multiValuesOptionsFromPb(pb *multiValuesOptionsPb) (*MultiValuesOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MultiValuesOptions{}
	st.Prefix = pb.Prefix
	st.Separator = pb.Separator
	st.Suffix = pb.Suffix

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *multiValuesOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st multiValuesOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NumericValue struct {

	// Wire name: 'value'
	Value float64

	ForceSendFields []string `tf:"-"`
}

func numericValueToPb(st *NumericValue) (*numericValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &numericValuePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type numericValuePb struct {
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func numericValueFromPb(pb *numericValuePb) (*NumericValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NumericValue{}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *numericValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st numericValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A singular noun object type.
type ObjectType string
type objectTypePb string

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

// Type always returns ObjectType to satisfy [pflag.Value] interface
func (f *ObjectType) Type() string {
	return "ObjectType"
}

func objectTypeToPb(st *ObjectType) (*objectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := objectTypePb(*st)
	return &pb, nil
}

func objectTypeFromPb(pb *objectTypePb) (*ObjectType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ObjectType(*pb)
	return &st, nil
}

// Always a plural of the object type.
type ObjectTypePlural string
type objectTypePluralPb string

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

// Type always returns ObjectTypePlural to satisfy [pflag.Value] interface
func (f *ObjectTypePlural) Type() string {
	return "ObjectTypePlural"
}

func objectTypePluralToPb(st *ObjectTypePlural) (*objectTypePluralPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := objectTypePluralPb(*st)
	return &pb, nil
}

func objectTypePluralFromPb(pb *objectTypePluralPb) (*ObjectTypePlural, error) {
	if pb == nil {
		return nil, nil
	}
	st := ObjectTypePlural(*pb)
	return &st, nil
}

type OdbcParams struct {

	// Wire name: 'hostname'
	Hostname string

	// Wire name: 'path'
	Path string

	// Wire name: 'port'
	Port int

	// Wire name: 'protocol'
	Protocol string

	ForceSendFields []string `tf:"-"`
}

func odbcParamsToPb(st *OdbcParams) (*odbcParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &odbcParamsPb{}
	pb.Hostname = st.Hostname

	pb.Path = st.Path

	pb.Port = st.Port

	pb.Protocol = st.Protocol

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type odbcParamsPb struct {
	Hostname string `json:"hostname,omitempty"`

	Path string `json:"path,omitempty"`

	Port int `json:"port,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func odbcParamsFromPb(pb *odbcParamsPb) (*OdbcParams, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OdbcParams{}
	st.Hostname = pb.Hostname
	st.Path = pb.Path
	st.Port = pb.Port
	st.Protocol = pb.Protocol

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *odbcParamsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st odbcParamsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The singular form of the type of object which can be owned.
type OwnableObjectType string
type ownableObjectTypePb string

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

// Type always returns OwnableObjectType to satisfy [pflag.Value] interface
func (f *OwnableObjectType) Type() string {
	return "OwnableObjectType"
}

func ownableObjectTypeToPb(st *OwnableObjectType) (*ownableObjectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := ownableObjectTypePb(*st)
	return &pb, nil
}

func ownableObjectTypeFromPb(pb *ownableObjectTypePb) (*OwnableObjectType, error) {
	if pb == nil {
		return nil, nil
	}
	st := OwnableObjectType(*pb)
	return &st, nil
}

type Parameter struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	// Wire name: 'enumOptions'
	EnumOptions string
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	// Wire name: 'multiValuesOptions'
	MultiValuesOptions *MultiValuesOptions
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	// Wire name: 'name'
	Name string
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	// Wire name: 'queryId'
	QueryId string
	// The text displayed in a parameter picking widget.
	// Wire name: 'title'
	Title string
	// Parameters can have several different types.
	// Wire name: 'type'
	Type ParameterType
	// The default value for this parameter.
	// Wire name: 'value'
	Value any

	ForceSendFields []string `tf:"-"`
}

func parameterToPb(st *Parameter) (*parameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &parameterPb{}
	pb.EnumOptions = st.EnumOptions

	multiValuesOptionsPb, err := multiValuesOptionsToPb(st.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsPb != nil {
		pb.MultiValuesOptions = multiValuesOptionsPb
	}

	pb.Name = st.Name

	pb.QueryId = st.QueryId

	pb.Title = st.Title

	pb.Type = st.Type

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type parameterPb struct {
	// List of valid parameter values, newline delimited. Only applies for
	// dropdown list parameters.
	EnumOptions string `json:"enumOptions,omitempty"`
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions *multiValuesOptionsPb `json:"multiValuesOptions,omitempty"`
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name string `json:"name,omitempty"`
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId string `json:"queryId,omitempty"`
	// The text displayed in a parameter picking widget.
	Title string `json:"title,omitempty"`
	// Parameters can have several different types.
	Type ParameterType `json:"type,omitempty"`
	// The default value for this parameter.
	Value any `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func parameterFromPb(pb *parameterPb) (*Parameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Parameter{}
	st.EnumOptions = pb.EnumOptions
	multiValuesOptionsField, err := multiValuesOptionsFromPb(pb.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsField != nil {
		st.MultiValuesOptions = multiValuesOptionsField
	}
	st.Name = pb.Name
	st.QueryId = pb.QueryId
	st.Title = pb.Title
	st.Type = pb.Type
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *parameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st parameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Parameters can have several different types.
type ParameterType string
type parameterTypePb string

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

// Type always returns ParameterType to satisfy [pflag.Value] interface
func (f *ParameterType) Type() string {
	return "ParameterType"
}

func parameterTypeToPb(st *ParameterType) (*parameterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := parameterTypePb(*st)
	return &pb, nil
}

func parameterTypeFromPb(pb *parameterTypePb) (*ParameterType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ParameterType(*pb)
	return &st, nil
}

// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query * `CAN_EDIT`:
// Can edit the query * `CAN_MANAGE`: Can manage the query
type PermissionLevel string
type permissionLevelPb string

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

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

func permissionLevelToPb(st *PermissionLevel) (*permissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := permissionLevelPb(*st)
	return &pb, nil
}

func permissionLevelFromPb(pb *permissionLevelPb) (*PermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PermissionLevel(*pb)
	return &st, nil
}

// Possible Reasons for which we have not saved plans in the database
type PlansState string
type plansStatePb string

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

// Type always returns PlansState to satisfy [pflag.Value] interface
func (f *PlansState) Type() string {
	return "PlansState"
}

func plansStateToPb(st *PlansState) (*plansStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := plansStatePb(*st)
	return &pb, nil
}

func plansStateFromPb(pb *plansStatePb) (*PlansState, error) {
	if pb == nil {
		return nil, nil
	}
	st := PlansState(*pb)
	return &st, nil
}

type Query struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string
	// Timestamp when this query was created.
	// Wire name: 'create_time'
	CreateTime string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying the query.
	// Wire name: 'id'
	Id string
	// Username of the user who last saved changes to this query.
	// Wire name: 'last_modifier_user_name'
	LastModifierUserName string
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter
	// Workspace path of the workspace folder containing the object.
	// Wire name: 'parent_path'
	ParentPath string
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string

	// Wire name: 'tags'
	Tags []string
	// Timestamp when this query was last updated.
	// Wire name: 'update_time'
	UpdateTime string
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func queryToPb(st *Query) (*queryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit

	pb.Catalog = st.Catalog

	pb.CreateTime = st.CreateTime

	pb.Description = st.Description

	pb.DisplayName = st.DisplayName

	pb.Id = st.Id

	pb.LastModifierUserName = st.LastModifierUserName

	pb.LifecycleState = st.LifecycleState

	pb.OwnerUserName = st.OwnerUserName

	var parametersPb []queryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := queryParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	pb.ParentPath = st.ParentPath

	pb.QueryText = st.QueryText

	pb.RunAsMode = st.RunAsMode

	pb.Schema = st.Schema

	pb.Tags = st.Tags

	pb.UpdateTime = st.UpdateTime

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryPb struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// Timestamp when this query was created.
	CreateTime string `json:"create_time,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the query.
	Id string `json:"id,omitempty"`
	// Username of the user who last saved changes to this query.
	LastModifierUserName string `json:"last_modifier_user_name,omitempty"`
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState `json:"lifecycle_state,omitempty"`
	// Username of the user that owns the query.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	Parameters []queryParameterPb `json:"parameters,omitempty"`
	// Workspace path of the workspace folder containing the object.
	ParentPath string `json:"parent_path,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// Timestamp when this query was last updated.
	UpdateTime string `json:"update_time,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryFromPb(pb *queryPb) (*Query, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Query{}
	st.ApplyAutoLimit = pb.ApplyAutoLimit
	st.Catalog = pb.Catalog
	st.CreateTime = pb.CreateTime
	st.Description = pb.Description
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.LastModifierUserName = pb.LastModifierUserName
	st.LifecycleState = pb.LifecycleState
	st.OwnerUserName = pb.OwnerUserName

	var parametersField []QueryParameter
	for _, itemPb := range pb.Parameters {
		item, err := queryParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.ParentPath = pb.ParentPath
	st.QueryText = pb.QueryText
	st.RunAsMode = pb.RunAsMode
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	st.UpdateTime = pb.UpdateTime
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	// Wire name: 'multi_values_options'
	MultiValuesOptions *MultiValuesOptions
	// UUID of the query that provides the parameter values.
	// Wire name: 'query_id'
	QueryId string
	// List of selected query parameter values.
	// Wire name: 'values'
	Values []string

	ForceSendFields []string `tf:"-"`
}

func queryBackedValueToPb(st *QueryBackedValue) (*queryBackedValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryBackedValuePb{}
	multiValuesOptionsPb, err := multiValuesOptionsToPb(st.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsPb != nil {
		pb.MultiValuesOptions = multiValuesOptionsPb
	}

	pb.QueryId = st.QueryId

	pb.Values = st.Values

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryBackedValuePb struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions *multiValuesOptionsPb `json:"multi_values_options,omitempty"`
	// UUID of the query that provides the parameter values.
	QueryId string `json:"query_id,omitempty"`
	// List of selected query parameter values.
	Values []string `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryBackedValueFromPb(pb *queryBackedValuePb) (*QueryBackedValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryBackedValue{}
	multiValuesOptionsField, err := multiValuesOptionsFromPb(pb.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsField != nil {
		st.MultiValuesOptions = multiValuesOptionsField
	}
	st.QueryId = pb.QueryId
	st.Values = pb.Values

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryBackedValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryBackedValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryEditContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	// Wire name: 'options'
	Options any
	// The text of the query to be run.
	// Wire name: 'query'
	Query string

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole

	// Wire name: 'tags'
	Tags []string

	ForceSendFields []string `tf:"-"`
}

func queryEditContentToPb(st *QueryEditContent) (*queryEditContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryEditContentPb{}
	pb.DataSourceId = st.DataSourceId

	pb.Description = st.Description

	pb.Name = st.Name

	pb.Options = st.Options

	pb.Query = st.Query

	pb.QueryId = st.QueryId

	pb.RunAsRole = st.RunAsRole

	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryEditContentPb struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `json:"options,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`

	QueryId string `json:"-" url:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryEditContentFromPb(pb *queryEditContentPb) (*QueryEditContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryEditContent{}
	st.DataSourceId = pb.DataSourceId
	st.Description = pb.Description
	st.Name = pb.Name
	st.Options = pb.Options
	st.Query = pb.Query
	st.QueryId = pb.QueryId
	st.RunAsRole = pb.RunAsRole
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryEditContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryEditContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryFilter struct {
	// A range filter for query submitted time. The time range must be <= 30
	// days.
	// Wire name: 'query_start_time_range'
	QueryStartTimeRange *TimeRange
	// A list of statement IDs.
	// Wire name: 'statement_ids'
	StatementIds []string

	// Wire name: 'statuses'
	Statuses []QueryStatus
	// A list of user IDs who ran the queries.
	// Wire name: 'user_ids'
	UserIds []int64
	// A list of warehouse IDs.
	// Wire name: 'warehouse_ids'
	WarehouseIds []string
}

func queryFilterToPb(st *QueryFilter) (*queryFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryFilterPb{}
	queryStartTimeRangePb, err := timeRangeToPb(st.QueryStartTimeRange)
	if err != nil {
		return nil, err
	}
	if queryStartTimeRangePb != nil {
		pb.QueryStartTimeRange = queryStartTimeRangePb
	}

	pb.StatementIds = st.StatementIds

	pb.Statuses = st.Statuses

	pb.UserIds = st.UserIds

	pb.WarehouseIds = st.WarehouseIds

	return pb, nil
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

type queryFilterPb struct {
	// A range filter for query submitted time. The time range must be <= 30
	// days.
	QueryStartTimeRange *timeRangePb `json:"query_start_time_range,omitempty" url:"query_start_time_range,omitempty"`
	// A list of statement IDs.
	StatementIds []string `json:"statement_ids,omitempty" url:"statement_ids,omitempty"`

	Statuses []QueryStatus `json:"statuses,omitempty" url:"statuses,omitempty"`
	// A list of user IDs who ran the queries.
	UserIds []int64 `json:"user_ids,omitempty" url:"user_ids,omitempty"`
	// A list of warehouse IDs.
	WarehouseIds []string `json:"warehouse_ids,omitempty" url:"warehouse_ids,omitempty"`
}

func queryFilterFromPb(pb *queryFilterPb) (*QueryFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryFilter{}
	queryStartTimeRangeField, err := timeRangeFromPb(pb.QueryStartTimeRange)
	if err != nil {
		return nil, err
	}
	if queryStartTimeRangeField != nil {
		st.QueryStartTimeRange = queryStartTimeRangeField
	}
	st.StatementIds = pb.StatementIds
	st.Statuses = pb.Statuses
	st.UserIds = pb.UserIds
	st.WarehouseIds = pb.WarehouseIds

	return st, nil
}

type QueryInfo struct {
	// SQL Warehouse channel information at the time of query execution
	// Wire name: 'channel_used'
	ChannelUsed *ChannelInfo
	// Client application that ran the statement. For example: Databricks SQL
	// Editor, Tableau, and Power BI. This field is derived from information
	// provided by client applications. While values are expected to remain
	// static over time, this cannot be guaranteed.
	// Wire name: 'client_application'
	ClientApplication string
	// Total execution time of the statement ( excluding result fetch time ).
	// Wire name: 'duration'
	Duration int64
	// Alias for `warehouse_id`.
	// Wire name: 'endpoint_id'
	EndpointId string
	// Message describing why the query could not complete.
	// Wire name: 'error_message'
	ErrorMessage string
	// The ID of the user whose credentials were used to run the query.
	// Wire name: 'executed_as_user_id'
	ExecutedAsUserId int64
	// The email address or username of the user whose credentials were used to
	// run the query.
	// Wire name: 'executed_as_user_name'
	ExecutedAsUserName string
	// The time execution of the query ended.
	// Wire name: 'execution_end_time_ms'
	ExecutionEndTimeMs int64
	// Whether more updates for the query are expected.
	// Wire name: 'is_final'
	IsFinal bool
	// A key that can be used to look up query details.
	// Wire name: 'lookup_key'
	LookupKey string
	// Metrics about query execution.
	// Wire name: 'metrics'
	Metrics *QueryMetrics
	// Whether plans exist for the execution, or the reason why they are missing
	// Wire name: 'plans_state'
	PlansState PlansState
	// The time the query ended.
	// Wire name: 'query_end_time_ms'
	QueryEndTimeMs int64
	// The query ID.
	// Wire name: 'query_id'
	QueryId string
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	// Wire name: 'query_source'
	QuerySource *ExternalQuerySource
	// The time the query started.
	// Wire name: 'query_start_time_ms'
	QueryStartTimeMs int64
	// The text of the query.
	// Wire name: 'query_text'
	QueryText string
	// The number of results returned by the query.
	// Wire name: 'rows_produced'
	RowsProduced int64
	// URL to the Spark UI query plan.
	// Wire name: 'spark_ui_url'
	SparkUiUrl string
	// Type of statement for this query
	// Wire name: 'statement_type'
	StatementType QueryStatementType
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	// Wire name: 'status'
	Status QueryStatus
	// The ID of the user who ran the query.
	// Wire name: 'user_id'
	UserId int64
	// The email address or username of the user who ran the query.
	// Wire name: 'user_name'
	UserName string
	// Warehouse ID.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func queryInfoToPb(st *QueryInfo) (*queryInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryInfoPb{}
	channelUsedPb, err := channelInfoToPb(st.ChannelUsed)
	if err != nil {
		return nil, err
	}
	if channelUsedPb != nil {
		pb.ChannelUsed = channelUsedPb
	}

	pb.ClientApplication = st.ClientApplication

	pb.Duration = st.Duration

	pb.EndpointId = st.EndpointId

	pb.ErrorMessage = st.ErrorMessage

	pb.ExecutedAsUserId = st.ExecutedAsUserId

	pb.ExecutedAsUserName = st.ExecutedAsUserName

	pb.ExecutionEndTimeMs = st.ExecutionEndTimeMs

	pb.IsFinal = st.IsFinal

	pb.LookupKey = st.LookupKey

	metricsPb, err := queryMetricsToPb(st.Metrics)
	if err != nil {
		return nil, err
	}
	if metricsPb != nil {
		pb.Metrics = metricsPb
	}

	pb.PlansState = st.PlansState

	pb.QueryEndTimeMs = st.QueryEndTimeMs

	pb.QueryId = st.QueryId

	querySourcePb, err := externalQuerySourceToPb(st.QuerySource)
	if err != nil {
		return nil, err
	}
	if querySourcePb != nil {
		pb.QuerySource = querySourcePb
	}

	pb.QueryStartTimeMs = st.QueryStartTimeMs

	pb.QueryText = st.QueryText

	pb.RowsProduced = st.RowsProduced

	pb.SparkUiUrl = st.SparkUiUrl

	pb.StatementType = st.StatementType

	pb.Status = st.Status

	pb.UserId = st.UserId

	pb.UserName = st.UserName

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryInfoPb struct {
	// SQL Warehouse channel information at the time of query execution
	ChannelUsed *channelInfoPb `json:"channel_used,omitempty"`
	// Client application that ran the statement. For example: Databricks SQL
	// Editor, Tableau, and Power BI. This field is derived from information
	// provided by client applications. While values are expected to remain
	// static over time, this cannot be guaranteed.
	ClientApplication string `json:"client_application,omitempty"`
	// Total execution time of the statement ( excluding result fetch time ).
	Duration int64 `json:"duration,omitempty"`
	// Alias for `warehouse_id`.
	EndpointId string `json:"endpoint_id,omitempty"`
	// Message describing why the query could not complete.
	ErrorMessage string `json:"error_message,omitempty"`
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId int64 `json:"executed_as_user_id,omitempty"`
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName string `json:"executed_as_user_name,omitempty"`
	// The time execution of the query ended.
	ExecutionEndTimeMs int64 `json:"execution_end_time_ms,omitempty"`
	// Whether more updates for the query are expected.
	IsFinal bool `json:"is_final,omitempty"`
	// A key that can be used to look up query details.
	LookupKey string `json:"lookup_key,omitempty"`
	// Metrics about query execution.
	Metrics *queryMetricsPb `json:"metrics,omitempty"`
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState PlansState `json:"plans_state,omitempty"`
	// The time the query ended.
	QueryEndTimeMs int64 `json:"query_end_time_ms,omitempty"`
	// The query ID.
	QueryId string `json:"query_id,omitempty"`
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	QuerySource *externalQuerySourcePb `json:"query_source,omitempty"`
	// The time the query started.
	QueryStartTimeMs int64 `json:"query_start_time_ms,omitempty"`
	// The text of the query.
	QueryText string `json:"query_text,omitempty"`
	// The number of results returned by the query.
	RowsProduced int64 `json:"rows_produced,omitempty"`
	// URL to the Spark UI query plan.
	SparkUiUrl string `json:"spark_ui_url,omitempty"`
	// Type of statement for this query
	StatementType QueryStatementType `json:"statement_type,omitempty"`
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	Status QueryStatus `json:"status,omitempty"`
	// The ID of the user who ran the query.
	UserId int64 `json:"user_id,omitempty"`
	// The email address or username of the user who ran the query.
	UserName string `json:"user_name,omitempty"`
	// Warehouse ID.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryInfoFromPb(pb *queryInfoPb) (*QueryInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryInfo{}
	channelUsedField, err := channelInfoFromPb(pb.ChannelUsed)
	if err != nil {
		return nil, err
	}
	if channelUsedField != nil {
		st.ChannelUsed = channelUsedField
	}
	st.ClientApplication = pb.ClientApplication
	st.Duration = pb.Duration
	st.EndpointId = pb.EndpointId
	st.ErrorMessage = pb.ErrorMessage
	st.ExecutedAsUserId = pb.ExecutedAsUserId
	st.ExecutedAsUserName = pb.ExecutedAsUserName
	st.ExecutionEndTimeMs = pb.ExecutionEndTimeMs
	st.IsFinal = pb.IsFinal
	st.LookupKey = pb.LookupKey
	metricsField, err := queryMetricsFromPb(pb.Metrics)
	if err != nil {
		return nil, err
	}
	if metricsField != nil {
		st.Metrics = metricsField
	}
	st.PlansState = pb.PlansState
	st.QueryEndTimeMs = pb.QueryEndTimeMs
	st.QueryId = pb.QueryId
	querySourceField, err := externalQuerySourceFromPb(pb.QuerySource)
	if err != nil {
		return nil, err
	}
	if querySourceField != nil {
		st.QuerySource = querySourceField
	}
	st.QueryStartTimeMs = pb.QueryStartTimeMs
	st.QueryText = pb.QueryText
	st.RowsProduced = pb.RowsProduced
	st.SparkUiUrl = pb.SparkUiUrl
	st.StatementType = pb.StatementType
	st.Status = pb.Status
	st.UserId = pb.UserId
	st.UserName = pb.UserName
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryList struct {
	// The total number of queries.
	// Wire name: 'count'
	Count int
	// The page number that is currently displayed.
	// Wire name: 'page'
	Page int
	// The number of queries per page.
	// Wire name: 'page_size'
	PageSize int
	// List of queries returned.
	// Wire name: 'results'
	Results []LegacyQuery

	ForceSendFields []string `tf:"-"`
}

func queryListToPb(st *QueryList) (*queryListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryListPb{}
	pb.Count = st.Count

	pb.Page = st.Page

	pb.PageSize = st.PageSize

	var resultsPb []legacyQueryPb
	for _, item := range st.Results {
		itemPb, err := legacyQueryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryListPb struct {
	// The total number of queries.
	Count int `json:"count,omitempty"`
	// The page number that is currently displayed.
	Page int `json:"page,omitempty"`
	// The number of queries per page.
	PageSize int `json:"page_size,omitempty"`
	// List of queries returned.
	Results []legacyQueryPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryListFromPb(pb *queryListPb) (*QueryList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryList{}
	st.Count = pb.Count
	st.Page = pb.Page
	st.PageSize = pb.PageSize

	var resultsField []LegacyQuery
	for _, itemPb := range pb.Results {
		item, err := legacyQueryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	// Wire name: 'compilation_time_ms'
	CompilationTimeMs int64
	// Time spent executing the query, in milliseconds.
	// Wire name: 'execution_time_ms'
	ExecutionTimeMs int64
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	// Wire name: 'network_sent_bytes'
	NetworkSentBytes int64
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	// Wire name: 'overloading_queue_start_timestamp'
	OverloadingQueueStartTimestamp int64
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	// Wire name: 'photon_total_time_ms'
	PhotonTotalTimeMs int64
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	// Wire name: 'provisioning_queue_start_timestamp'
	ProvisioningQueueStartTimestamp int64
	// Total number of bytes in all tables not read due to pruning
	// Wire name: 'pruned_bytes'
	PrunedBytes int64
	// Total number of files from all tables not read due to pruning
	// Wire name: 'pruned_files_count'
	PrunedFilesCount int64
	// Timestamp of when the underlying compute started compilation of the
	// query.
	// Wire name: 'query_compilation_start_timestamp'
	QueryCompilationStartTimestamp int64
	// Total size of data read by the query, in bytes.
	// Wire name: 'read_bytes'
	ReadBytes int64
	// Size of persistent data read from the cache, in bytes.
	// Wire name: 'read_cache_bytes'
	ReadCacheBytes int64
	// Number of files read after pruning
	// Wire name: 'read_files_count'
	ReadFilesCount int64
	// Number of partitions read after pruning.
	// Wire name: 'read_partitions_count'
	ReadPartitionsCount int64
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	// Wire name: 'read_remote_bytes'
	ReadRemoteBytes int64
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	// Wire name: 'result_fetch_time_ms'
	ResultFetchTimeMs int64
	// `true` if the query result was fetched from cache, `false` otherwise.
	// Wire name: 'result_from_cache'
	ResultFromCache bool
	// Total number of rows returned by the query.
	// Wire name: 'rows_produced_count'
	RowsProducedCount int64
	// Total number of rows read by the query.
	// Wire name: 'rows_read_count'
	RowsReadCount int64
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	// Wire name: 'spill_to_disk_bytes'
	SpillToDiskBytes int64
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	// Wire name: 'task_total_time_ms'
	TaskTotalTimeMs int64
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	// Wire name: 'total_time_ms'
	TotalTimeMs int64
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	// Wire name: 'write_remote_bytes'
	WriteRemoteBytes int64

	ForceSendFields []string `tf:"-"`
}

func queryMetricsToPb(st *QueryMetrics) (*queryMetricsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryMetricsPb{}
	pb.CompilationTimeMs = st.CompilationTimeMs

	pb.ExecutionTimeMs = st.ExecutionTimeMs

	pb.NetworkSentBytes = st.NetworkSentBytes

	pb.OverloadingQueueStartTimestamp = st.OverloadingQueueStartTimestamp

	pb.PhotonTotalTimeMs = st.PhotonTotalTimeMs

	pb.ProvisioningQueueStartTimestamp = st.ProvisioningQueueStartTimestamp

	pb.PrunedBytes = st.PrunedBytes

	pb.PrunedFilesCount = st.PrunedFilesCount

	pb.QueryCompilationStartTimestamp = st.QueryCompilationStartTimestamp

	pb.ReadBytes = st.ReadBytes

	pb.ReadCacheBytes = st.ReadCacheBytes

	pb.ReadFilesCount = st.ReadFilesCount

	pb.ReadPartitionsCount = st.ReadPartitionsCount

	pb.ReadRemoteBytes = st.ReadRemoteBytes

	pb.ResultFetchTimeMs = st.ResultFetchTimeMs

	pb.ResultFromCache = st.ResultFromCache

	pb.RowsProducedCount = st.RowsProducedCount

	pb.RowsReadCount = st.RowsReadCount

	pb.SpillToDiskBytes = st.SpillToDiskBytes

	pb.TaskTotalTimeMs = st.TaskTotalTimeMs

	pb.TotalTimeMs = st.TotalTimeMs

	pb.WriteRemoteBytes = st.WriteRemoteBytes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryMetricsPb struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	CompilationTimeMs int64 `json:"compilation_time_ms,omitempty"`
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs int64 `json:"execution_time_ms,omitempty"`
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	NetworkSentBytes int64 `json:"network_sent_bytes,omitempty"`
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	OverloadingQueueStartTimestamp int64 `json:"overloading_queue_start_timestamp,omitempty"`
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs int64 `json:"photon_total_time_ms,omitempty"`
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	ProvisioningQueueStartTimestamp int64 `json:"provisioning_queue_start_timestamp,omitempty"`
	// Total number of bytes in all tables not read due to pruning
	PrunedBytes int64 `json:"pruned_bytes,omitempty"`
	// Total number of files from all tables not read due to pruning
	PrunedFilesCount int64 `json:"pruned_files_count,omitempty"`
	// Timestamp of when the underlying compute started compilation of the
	// query.
	QueryCompilationStartTimestamp int64 `json:"query_compilation_start_timestamp,omitempty"`
	// Total size of data read by the query, in bytes.
	ReadBytes int64 `json:"read_bytes,omitempty"`
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes int64 `json:"read_cache_bytes,omitempty"`
	// Number of files read after pruning
	ReadFilesCount int64 `json:"read_files_count,omitempty"`
	// Number of partitions read after pruning.
	ReadPartitionsCount int64 `json:"read_partitions_count,omitempty"`
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes int64 `json:"read_remote_bytes,omitempty"`
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs int64 `json:"result_fetch_time_ms,omitempty"`
	// `true` if the query result was fetched from cache, `false` otherwise.
	ResultFromCache bool `json:"result_from_cache,omitempty"`
	// Total number of rows returned by the query.
	RowsProducedCount int64 `json:"rows_produced_count,omitempty"`
	// Total number of rows read by the query.
	RowsReadCount int64 `json:"rows_read_count,omitempty"`
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes int64 `json:"spill_to_disk_bytes,omitempty"`
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs int64 `json:"task_total_time_ms,omitempty"`
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs int64 `json:"total_time_ms,omitempty"`
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes int64 `json:"write_remote_bytes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryMetricsFromPb(pb *queryMetricsPb) (*QueryMetrics, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryMetrics{}
	st.CompilationTimeMs = pb.CompilationTimeMs
	st.ExecutionTimeMs = pb.ExecutionTimeMs
	st.NetworkSentBytes = pb.NetworkSentBytes
	st.OverloadingQueueStartTimestamp = pb.OverloadingQueueStartTimestamp
	st.PhotonTotalTimeMs = pb.PhotonTotalTimeMs
	st.ProvisioningQueueStartTimestamp = pb.ProvisioningQueueStartTimestamp
	st.PrunedBytes = pb.PrunedBytes
	st.PrunedFilesCount = pb.PrunedFilesCount
	st.QueryCompilationStartTimestamp = pb.QueryCompilationStartTimestamp
	st.ReadBytes = pb.ReadBytes
	st.ReadCacheBytes = pb.ReadCacheBytes
	st.ReadFilesCount = pb.ReadFilesCount
	st.ReadPartitionsCount = pb.ReadPartitionsCount
	st.ReadRemoteBytes = pb.ReadRemoteBytes
	st.ResultFetchTimeMs = pb.ResultFetchTimeMs
	st.ResultFromCache = pb.ResultFromCache
	st.RowsProducedCount = pb.RowsProducedCount
	st.RowsReadCount = pb.RowsReadCount
	st.SpillToDiskBytes = pb.SpillToDiskBytes
	st.TaskTotalTimeMs = pb.TaskTotalTimeMs
	st.TotalTimeMs = pb.TotalTimeMs
	st.WriteRemoteBytes = pb.WriteRemoteBytes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryMetricsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryMetricsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryOptions struct {
	// The name of the catalog to execute this query in.
	// Wire name: 'catalog'
	Catalog string
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	// Wire name: 'moved_to_trash_at'
	MovedToTrashAt string

	// Wire name: 'parameters'
	Parameters []Parameter
	// The name of the schema to execute this query in.
	// Wire name: 'schema'
	Schema string

	ForceSendFields []string `tf:"-"`
}

func queryOptionsToPb(st *QueryOptions) (*queryOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryOptionsPb{}
	pb.Catalog = st.Catalog

	pb.MovedToTrashAt = st.MovedToTrashAt

	var parametersPb []parameterPb
	for _, item := range st.Parameters {
		itemPb, err := parameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	pb.Schema = st.Schema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryOptionsPb struct {
	// The name of the catalog to execute this query in.
	Catalog string `json:"catalog,omitempty"`
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	Parameters []parameterPb `json:"parameters,omitempty"`
	// The name of the schema to execute this query in.
	Schema string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryOptionsFromPb(pb *queryOptionsPb) (*QueryOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryOptions{}
	st.Catalog = pb.Catalog
	st.MovedToTrashAt = pb.MovedToTrashAt

	var parametersField []Parameter
	for _, itemPb := range pb.Parameters {
		item, err := parameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.Schema = pb.Schema

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryParameter struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	// Wire name: 'date_range_value'
	DateRangeValue *DateRangeValue
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	// Wire name: 'date_value'
	DateValue *DateValue
	// Dropdown query parameter value.
	// Wire name: 'enum_value'
	EnumValue *EnumValue
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	// Wire name: 'name'
	Name string
	// Numeric query parameter value.
	// Wire name: 'numeric_value'
	NumericValue *NumericValue
	// Query-based dropdown query parameter value.
	// Wire name: 'query_backed_value'
	QueryBackedValue *QueryBackedValue
	// Text query parameter value.
	// Wire name: 'text_value'
	TextValue *TextValue
	// Text displayed in the user-facing parameter widget in the UI.
	// Wire name: 'title'
	Title string

	ForceSendFields []string `tf:"-"`
}

func queryParameterToPb(st *QueryParameter) (*queryParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryParameterPb{}
	dateRangeValuePb, err := dateRangeValueToPb(st.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValuePb != nil {
		pb.DateRangeValue = dateRangeValuePb
	}

	dateValuePb, err := dateValueToPb(st.DateValue)
	if err != nil {
		return nil, err
	}
	if dateValuePb != nil {
		pb.DateValue = dateValuePb
	}

	enumValuePb, err := enumValueToPb(st.EnumValue)
	if err != nil {
		return nil, err
	}
	if enumValuePb != nil {
		pb.EnumValue = enumValuePb
	}

	pb.Name = st.Name

	numericValuePb, err := numericValueToPb(st.NumericValue)
	if err != nil {
		return nil, err
	}
	if numericValuePb != nil {
		pb.NumericValue = numericValuePb
	}

	queryBackedValuePb, err := queryBackedValueToPb(st.QueryBackedValue)
	if err != nil {
		return nil, err
	}
	if queryBackedValuePb != nil {
		pb.QueryBackedValue = queryBackedValuePb
	}

	textValuePb, err := textValueToPb(st.TextValue)
	if err != nil {
		return nil, err
	}
	if textValuePb != nil {
		pb.TextValue = textValuePb
	}

	pb.Title = st.Title

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryParameterPb struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	DateRangeValue *dateRangeValuePb `json:"date_range_value,omitempty"`
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	DateValue *dateValuePb `json:"date_value,omitempty"`
	// Dropdown query parameter value.
	EnumValue *enumValuePb `json:"enum_value,omitempty"`
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	Name string `json:"name,omitempty"`
	// Numeric query parameter value.
	NumericValue *numericValuePb `json:"numeric_value,omitempty"`
	// Query-based dropdown query parameter value.
	QueryBackedValue *queryBackedValuePb `json:"query_backed_value,omitempty"`
	// Text query parameter value.
	TextValue *textValuePb `json:"text_value,omitempty"`
	// Text displayed in the user-facing parameter widget in the UI.
	Title string `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryParameterFromPb(pb *queryParameterPb) (*QueryParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryParameter{}
	dateRangeValueField, err := dateRangeValueFromPb(pb.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValueField != nil {
		st.DateRangeValue = dateRangeValueField
	}
	dateValueField, err := dateValueFromPb(pb.DateValue)
	if err != nil {
		return nil, err
	}
	if dateValueField != nil {
		st.DateValue = dateValueField
	}
	enumValueField, err := enumValueFromPb(pb.EnumValue)
	if err != nil {
		return nil, err
	}
	if enumValueField != nil {
		st.EnumValue = enumValueField
	}
	st.Name = pb.Name
	numericValueField, err := numericValueFromPb(pb.NumericValue)
	if err != nil {
		return nil, err
	}
	if numericValueField != nil {
		st.NumericValue = numericValueField
	}
	queryBackedValueField, err := queryBackedValueFromPb(pb.QueryBackedValue)
	if err != nil {
		return nil, err
	}
	if queryBackedValueField != nil {
		st.QueryBackedValue = queryBackedValueField
	}
	textValueField, err := textValueFromPb(pb.TextValue)
	if err != nil {
		return nil, err
	}
	if textValueField != nil {
		st.TextValue = textValueField
	}
	st.Title = pb.Title

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryPostContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	// Wire name: 'options'
	Options any
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string
	// The text of the query to be run.
	// Wire name: 'query'
	Query string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole

	// Wire name: 'tags'
	Tags []string

	ForceSendFields []string `tf:"-"`
}

func queryPostContentToPb(st *QueryPostContent) (*queryPostContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryPostContentPb{}
	pb.DataSourceId = st.DataSourceId

	pb.Description = st.Description

	pb.Name = st.Name

	pb.Options = st.Options

	pb.Parent = st.Parent

	pb.Query = st.Query

	pb.RunAsRole = st.RunAsRole

	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryPostContentPb struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string `json:"data_source_id,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string `json:"name,omitempty"`
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any `json:"options,omitempty"`
	// The identifier of the workspace folder containing the object.
	Parent string `json:"parent,omitempty"`
	// The text of the query to be run.
	Query string `json:"query,omitempty"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole `json:"run_as_role,omitempty"`

	Tags []string `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryPostContentFromPb(pb *queryPostContentPb) (*QueryPostContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryPostContent{}
	st.DataSourceId = pb.DataSourceId
	st.Description = pb.Description
	st.Name = pb.Name
	st.Options = pb.Options
	st.Parent = pb.Parent
	st.Query = pb.Query
	st.RunAsRole = pb.RunAsRole
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryPostContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryPostContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryStatementType string
type queryStatementTypePb string

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

// Type always returns QueryStatementType to satisfy [pflag.Value] interface
func (f *QueryStatementType) Type() string {
	return "QueryStatementType"
}

func queryStatementTypeToPb(st *QueryStatementType) (*queryStatementTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := queryStatementTypePb(*st)
	return &pb, nil
}

func queryStatementTypeFromPb(pb *queryStatementTypePb) (*QueryStatementType, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueryStatementType(*pb)
	return &st, nil
}

// Statuses which are also used by OperationStatus in runtime
type QueryStatus string
type queryStatusPb string

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

// Type always returns QueryStatus to satisfy [pflag.Value] interface
func (f *QueryStatus) Type() string {
	return "QueryStatus"
}

func queryStatusToPb(st *QueryStatus) (*queryStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := queryStatusPb(*st)
	return &pb, nil
}

func queryStatusFromPb(pb *queryStatusPb) (*QueryStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueryStatus(*pb)
	return &st, nil
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	// Wire name: 'config_pair'
	ConfigPair []EndpointConfPair

	// Wire name: 'configuration_pairs'
	ConfigurationPairs []EndpointConfPair
}

func repeatedEndpointConfPairsToPb(st *RepeatedEndpointConfPairs) (*repeatedEndpointConfPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repeatedEndpointConfPairsPb{}

	var configPairPb []endpointConfPairPb
	for _, item := range st.ConfigPair {
		itemPb, err := endpointConfPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			configPairPb = append(configPairPb, *itemPb)
		}
	}
	pb.ConfigPair = configPairPb

	var configurationPairsPb []endpointConfPairPb
	for _, item := range st.ConfigurationPairs {
		itemPb, err := endpointConfPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			configurationPairsPb = append(configurationPairsPb, *itemPb)
		}
	}
	pb.ConfigurationPairs = configurationPairsPb

	return pb, nil
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

type repeatedEndpointConfPairsPb struct {
	// Deprecated: Use configuration_pairs
	ConfigPair []endpointConfPairPb `json:"config_pair,omitempty"`

	ConfigurationPairs []endpointConfPairPb `json:"configuration_pairs,omitempty"`
}

func repeatedEndpointConfPairsFromPb(pb *repeatedEndpointConfPairsPb) (*RepeatedEndpointConfPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepeatedEndpointConfPairs{}

	var configPairField []EndpointConfPair
	for _, itemPb := range pb.ConfigPair {
		item, err := endpointConfPairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			configPairField = append(configPairField, *item)
		}
	}
	st.ConfigPair = configPairField

	var configurationPairsField []EndpointConfPair
	for _, itemPb := range pb.ConfigurationPairs {
		item, err := endpointConfPairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			configurationPairsField = append(configurationPairsField, *item)
		}
	}
	st.ConfigurationPairs = configurationPairsField

	return st, nil
}

// Restore a dashboard
type RestoreDashboardRequest struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func restoreDashboardRequestToPb(st *RestoreDashboardRequest) (*restoreDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
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

type restoreDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

func restoreDashboardRequestFromPb(pb *restoreDashboardRequestPb) (*RestoreDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

// Restore a query
type RestoreQueriesLegacyRequest struct {

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
}

func restoreQueriesLegacyRequestToPb(st *RestoreQueriesLegacyRequest) (*restoreQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
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

type restoreQueriesLegacyRequestPb struct {
	QueryId string `json:"-" url:"-"`
}

func restoreQueriesLegacyRequestFromPb(pb *restoreQueriesLegacyRequestPb) (*RestoreQueriesLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreQueriesLegacyRequest{}
	st.QueryId = pb.QueryId

	return st, nil
}

type RestoreResponse struct {
}

func restoreResponseToPb(st *RestoreResponse) (*restoreResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreResponsePb{}

	return pb, nil
}

func (st *RestoreResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreResponse) MarshalJSON() ([]byte, error) {
	pb, err := restoreResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type restoreResponsePb struct {
}

func restoreResponseFromPb(pb *restoreResponsePb) (*RestoreResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreResponse{}

	return st, nil
}

type ResultData struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	// Wire name: 'data_array'
	DataArray [][]string

	// Wire name: 'external_links'
	ExternalLinks []ExternalLink
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	// Wire name: 'next_chunk_index'
	NextChunkIndex int
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	// Wire name: 'next_chunk_internal_link'
	NextChunkInternalLink string
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset int64

	ForceSendFields []string `tf:"-"`
}

func resultDataToPb(st *ResultData) (*resultDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultDataPb{}
	pb.ByteCount = st.ByteCount

	pb.ChunkIndex = st.ChunkIndex

	pb.DataArray = st.DataArray

	var externalLinksPb []externalLinkPb
	for _, item := range st.ExternalLinks {
		itemPb, err := externalLinkToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			externalLinksPb = append(externalLinksPb, *itemPb)
		}
	}
	pb.ExternalLinks = externalLinksPb

	pb.NextChunkIndex = st.NextChunkIndex

	pb.NextChunkInternalLink = st.NextChunkInternalLink

	pb.RowCount = st.RowCount

	pb.RowOffset = st.RowOffset

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resultDataPb struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	ByteCount int64 `json:"byte_count,omitempty"`
	// The position within the sequence of result set chunks.
	ChunkIndex int `json:"chunk_index,omitempty"`
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray [][]string `json:"data_array,omitempty"`

	ExternalLinks []externalLinkPb `json:"external_links,omitempty"`
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex int `json:"next_chunk_index,omitempty"`
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink string `json:"next_chunk_internal_link,omitempty"`
	// The number of rows within the result chunk.
	RowCount int64 `json:"row_count,omitempty"`
	// The starting row offset within the result set.
	RowOffset int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultDataFromPb(pb *resultDataPb) (*ResultData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultData{}
	st.ByteCount = pb.ByteCount
	st.ChunkIndex = pb.ChunkIndex
	st.DataArray = pb.DataArray

	var externalLinksField []ExternalLink
	for _, itemPb := range pb.ExternalLinks {
		item, err := externalLinkFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			externalLinksField = append(externalLinksField, *item)
		}
	}
	st.ExternalLinks = externalLinksField
	st.NextChunkIndex = pb.NextChunkIndex
	st.NextChunkInternalLink = pb.NextChunkInternalLink
	st.RowCount = pb.RowCount
	st.RowOffset = pb.RowOffset

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultDataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultDataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	// Wire name: 'chunks'
	Chunks []BaseChunkInfo

	// Wire name: 'format'
	Format Format
	// The schema is an ordered list of column descriptions.
	// Wire name: 'schema'
	Schema *ResultSchema
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	// Wire name: 'total_byte_count'
	TotalByteCount int64
	// The total number of chunks that the result set has been divided into.
	// Wire name: 'total_chunk_count'
	TotalChunkCount int
	// The total number of rows in the result set.
	// Wire name: 'total_row_count'
	TotalRowCount int64
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	// Wire name: 'truncated'
	Truncated bool

	ForceSendFields []string `tf:"-"`
}

func resultManifestToPb(st *ResultManifest) (*resultManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultManifestPb{}

	var chunksPb []baseChunkInfoPb
	for _, item := range st.Chunks {
		itemPb, err := baseChunkInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			chunksPb = append(chunksPb, *itemPb)
		}
	}
	pb.Chunks = chunksPb

	pb.Format = st.Format

	schemaPb, err := resultSchemaToPb(st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = schemaPb
	}

	pb.TotalByteCount = st.TotalByteCount

	pb.TotalChunkCount = st.TotalChunkCount

	pb.TotalRowCount = st.TotalRowCount

	pb.Truncated = st.Truncated

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resultManifestPb struct {
	// Array of result set chunk metadata.
	Chunks []baseChunkInfoPb `json:"chunks,omitempty"`

	Format Format `json:"format,omitempty"`
	// The schema is an ordered list of column descriptions.
	Schema *resultSchemaPb `json:"schema,omitempty"`
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount int64 `json:"total_byte_count,omitempty"`
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount int `json:"total_chunk_count,omitempty"`
	// The total number of rows in the result set.
	TotalRowCount int64 `json:"total_row_count,omitempty"`
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated bool `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultManifestFromPb(pb *resultManifestPb) (*ResultManifest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultManifest{}

	var chunksField []BaseChunkInfo
	for _, itemPb := range pb.Chunks {
		item, err := baseChunkInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			chunksField = append(chunksField, *item)
		}
	}
	st.Chunks = chunksField
	st.Format = pb.Format
	schemaField, err := resultSchemaFromPb(pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = schemaField
	}
	st.TotalByteCount = pb.TotalByteCount
	st.TotalChunkCount = pb.TotalChunkCount
	st.TotalRowCount = pb.TotalRowCount
	st.Truncated = pb.Truncated

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultManifestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultManifestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {

	// Wire name: 'column_count'
	ColumnCount int

	// Wire name: 'columns'
	Columns []ColumnInfo

	ForceSendFields []string `tf:"-"`
}

func resultSchemaToPb(st *ResultSchema) (*resultSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultSchemaPb{}
	pb.ColumnCount = st.ColumnCount

	var columnsPb []columnInfoPb
	for _, item := range st.Columns {
		itemPb, err := columnInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type resultSchemaPb struct {
	ColumnCount int `json:"column_count,omitempty"`

	Columns []columnInfoPb `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultSchemaFromPb(pb *resultSchemaPb) (*ResultSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultSchema{}
	st.ColumnCount = pb.ColumnCount

	var columnsField []ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := columnInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunAsMode string
type runAsModePb string

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

// Type always returns RunAsMode to satisfy [pflag.Value] interface
func (f *RunAsMode) Type() string {
	return "RunAsMode"
}

func runAsModeToPb(st *RunAsMode) (*runAsModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runAsModePb(*st)
	return &pb, nil
}

func runAsModeFromPb(pb *runAsModePb) (*RunAsMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunAsMode(*pb)
	return &st, nil
}

// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as owner"
// behavior)
type RunAsRole string
type runAsRolePb string

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

// Type always returns RunAsRole to satisfy [pflag.Value] interface
func (f *RunAsRole) Type() string {
	return "RunAsRole"
}

func runAsRoleToPb(st *RunAsRole) (*runAsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runAsRolePb(*st)
	return &pb, nil
}

func runAsRoleFromPb(pb *runAsRolePb) (*RunAsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunAsRole(*pb)
	return &st, nil
}

type SchedulePauseStatus string
type schedulePauseStatusPb string

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

// Type always returns SchedulePauseStatus to satisfy [pflag.Value] interface
func (f *SchedulePauseStatus) Type() string {
	return "SchedulePauseStatus"
}

func schedulePauseStatusToPb(st *SchedulePauseStatus) (*schedulePauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := schedulePauseStatusPb(*st)
	return &pb, nil
}

func schedulePauseStatusFromPb(pb *schedulePauseStatusPb) (*SchedulePauseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SchedulePauseStatus(*pb)
	return &st, nil
}

type ServiceError struct {

	// Wire name: 'error_code'
	ErrorCode ServiceErrorCode
	// A brief summary of the error condition.
	// Wire name: 'message'
	Message string

	ForceSendFields []string `tf:"-"`
}

func serviceErrorToPb(st *ServiceError) (*serviceErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &serviceErrorPb{}
	pb.ErrorCode = st.ErrorCode

	pb.Message = st.Message

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type serviceErrorPb struct {
	ErrorCode ServiceErrorCode `json:"error_code,omitempty"`
	// A brief summary of the error condition.
	Message string `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func serviceErrorFromPb(pb *serviceErrorPb) (*ServiceError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServiceError{}
	st.ErrorCode = pb.ErrorCode
	st.Message = pb.Message

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *serviceErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st serviceErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServiceErrorCode string
type serviceErrorCodePb string

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

// Type always returns ServiceErrorCode to satisfy [pflag.Value] interface
func (f *ServiceErrorCode) Type() string {
	return "ServiceErrorCode"
}

func serviceErrorCodeToPb(st *ServiceErrorCode) (*serviceErrorCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := serviceErrorCodePb(*st)
	return &pb, nil
}

func serviceErrorCodeFromPb(pb *serviceErrorCodePb) (*ServiceErrorCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServiceErrorCode(*pb)
	return &st, nil
}

// Set object ACL
type SetRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	// Wire name: 'objectId'
	ObjectId string `tf:"-"`
	// The type of object permission to set.
	// Wire name: 'objectType'
	ObjectType ObjectTypePlural `tf:"-"`
}

func setRequestToPb(st *SetRequest) (*setRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setRequestPb{}

	var accessControlListPb []accessControlPb
	for _, item := range st.AccessControlList {
		itemPb, err := accessControlToPb(&item)
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

	return pb, nil
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

type setRequestPb struct {
	AccessControlList []accessControlPb `json:"access_control_list,omitempty"`
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId string `json:"-" url:"-"`
	// The type of object permission to set.
	ObjectType ObjectTypePlural `json:"-" url:"-"`
}

func setRequestFromPb(pb *setRequestPb) (*SetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetRequest{}

	var accessControlListField []AccessControl
	for _, itemPb := range pb.AccessControlList {
		item, err := accessControlFromPb(&itemPb)
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

	return st, nil
}

type SetResponse struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl
	// An object's type and UUID, separated by a forward slash (/) character.
	// Wire name: 'object_id'
	ObjectId string
	// A singular noun object type.
	// Wire name: 'object_type'
	ObjectType ObjectType

	ForceSendFields []string `tf:"-"`
}

func setResponseToPb(st *SetResponse) (*setResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setResponsePb{}

	var accessControlListPb []accessControlPb
	for _, item := range st.AccessControlList {
		itemPb, err := accessControlToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type setResponsePb struct {
	AccessControlList []accessControlPb `json:"access_control_list,omitempty"`
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId string `json:"object_id,omitempty"`
	// A singular noun object type.
	ObjectType ObjectType `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setResponseFromPb(pb *setResponsePb) (*SetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetResponse{}

	var accessControlListField []AccessControl
	for _, itemPb := range pb.AccessControlList {
		item, err := accessControlFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *setResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	// Wire name: 'channel'
	Channel *Channel
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'config_param'
	ConfigParam *RepeatedEndpointConfPairs
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	// Wire name: 'data_access_config'
	DataAccessConfig []EndpointConfPair
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	// Wire name: 'enabled_warehouse_types'
	EnabledWarehouseTypes []WarehouseTypePair
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'global_param'
	GlobalParam *RepeatedEndpointConfPairs
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	// Wire name: 'google_service_account'
	GoogleServiceAccount string
	// AWS Only: Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string
	// Security policy for warehouses
	// Wire name: 'security_policy'
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy
	// SQL configuration parameters
	// Wire name: 'sql_configuration_parameters'
	SqlConfigurationParameters *RepeatedEndpointConfPairs

	ForceSendFields []string `tf:"-"`
}

func setWorkspaceWarehouseConfigRequestToPb(st *SetWorkspaceWarehouseConfigRequest) (*setWorkspaceWarehouseConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setWorkspaceWarehouseConfigRequestPb{}
	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	configParamPb, err := repeatedEndpointConfPairsToPb(st.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamPb != nil {
		pb.ConfigParam = configParamPb
	}

	var dataAccessConfigPb []endpointConfPairPb
	for _, item := range st.DataAccessConfig {
		itemPb, err := endpointConfPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataAccessConfigPb = append(dataAccessConfigPb, *itemPb)
		}
	}
	pb.DataAccessConfig = dataAccessConfigPb

	var enabledWarehouseTypesPb []warehouseTypePairPb
	for _, item := range st.EnabledWarehouseTypes {
		itemPb, err := warehouseTypePairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			enabledWarehouseTypesPb = append(enabledWarehouseTypesPb, *itemPb)
		}
	}
	pb.EnabledWarehouseTypes = enabledWarehouseTypesPb

	globalParamPb, err := repeatedEndpointConfPairsToPb(st.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamPb != nil {
		pb.GlobalParam = globalParamPb
	}

	pb.GoogleServiceAccount = st.GoogleServiceAccount

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.SecurityPolicy = st.SecurityPolicy

	sqlConfigurationParametersPb, err := repeatedEndpointConfPairsToPb(st.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersPb != nil {
		pb.SqlConfigurationParameters = sqlConfigurationParametersPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type setWorkspaceWarehouseConfigRequestPb struct {
	// Optional: Channel selection details
	Channel *channelPb `json:"channel,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *repeatedEndpointConfPairsPb `json:"config_param,omitempty"`
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig []endpointConfPairPb `json:"data_access_config,omitempty"`
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes []warehouseTypePairPb `json:"enabled_warehouse_types,omitempty"`
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *repeatedEndpointConfPairsPb `json:"global_param,omitempty"`
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount string `json:"google_service_account,omitempty"`
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// Security policy for warehouses
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy `json:"security_policy,omitempty"`
	// SQL configuration parameters
	SqlConfigurationParameters *repeatedEndpointConfPairsPb `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setWorkspaceWarehouseConfigRequestFromPb(pb *setWorkspaceWarehouseConfigRequestPb) (*SetWorkspaceWarehouseConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetWorkspaceWarehouseConfigRequest{}
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	configParamField, err := repeatedEndpointConfPairsFromPb(pb.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamField != nil {
		st.ConfigParam = configParamField
	}

	var dataAccessConfigField []EndpointConfPair
	for _, itemPb := range pb.DataAccessConfig {
		item, err := endpointConfPairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataAccessConfigField = append(dataAccessConfigField, *item)
		}
	}
	st.DataAccessConfig = dataAccessConfigField

	var enabledWarehouseTypesField []WarehouseTypePair
	for _, itemPb := range pb.EnabledWarehouseTypes {
		item, err := warehouseTypePairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			enabledWarehouseTypesField = append(enabledWarehouseTypesField, *item)
		}
	}
	st.EnabledWarehouseTypes = enabledWarehouseTypesField
	globalParamField, err := repeatedEndpointConfPairsFromPb(pb.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamField != nil {
		st.GlobalParam = globalParamField
	}
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.SecurityPolicy = pb.SecurityPolicy
	sqlConfigurationParametersField, err := repeatedEndpointConfPairsFromPb(pb.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersField != nil {
		st.SqlConfigurationParameters = sqlConfigurationParametersField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *setWorkspaceWarehouseConfigRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setWorkspaceWarehouseConfigRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Security policy for warehouses
type SetWorkspaceWarehouseConfigRequestSecurityPolicy string
type setWorkspaceWarehouseConfigRequestSecurityPolicyPb string

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

// Type always returns SetWorkspaceWarehouseConfigRequestSecurityPolicy to satisfy [pflag.Value] interface
func (f *SetWorkspaceWarehouseConfigRequestSecurityPolicy) Type() string {
	return "SetWorkspaceWarehouseConfigRequestSecurityPolicy"
}

func setWorkspaceWarehouseConfigRequestSecurityPolicyToPb(st *SetWorkspaceWarehouseConfigRequestSecurityPolicy) (*setWorkspaceWarehouseConfigRequestSecurityPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := setWorkspaceWarehouseConfigRequestSecurityPolicyPb(*st)
	return &pb, nil
}

func setWorkspaceWarehouseConfigRequestSecurityPolicyFromPb(pb *setWorkspaceWarehouseConfigRequestSecurityPolicyPb) (*SetWorkspaceWarehouseConfigRequestSecurityPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := SetWorkspaceWarehouseConfigRequestSecurityPolicy(*pb)
	return &st, nil
}

type SetWorkspaceWarehouseConfigResponse struct {
}

func setWorkspaceWarehouseConfigResponseToPb(st *SetWorkspaceWarehouseConfigResponse) (*setWorkspaceWarehouseConfigResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setWorkspaceWarehouseConfigResponsePb{}

	return pb, nil
}

func (st *SetWorkspaceWarehouseConfigResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setWorkspaceWarehouseConfigResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setWorkspaceWarehouseConfigResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetWorkspaceWarehouseConfigResponse) MarshalJSON() ([]byte, error) {
	pb, err := setWorkspaceWarehouseConfigResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setWorkspaceWarehouseConfigResponsePb struct {
}

func setWorkspaceWarehouseConfigResponseFromPb(pb *setWorkspaceWarehouseConfigResponsePb) (*SetWorkspaceWarehouseConfigResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetWorkspaceWarehouseConfigResponse{}

	return st, nil
}

// Configurations whether the warehouse should use spot instances.
type SpotInstancePolicy string
type spotInstancePolicyPb string

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

// Type always returns SpotInstancePolicy to satisfy [pflag.Value] interface
func (f *SpotInstancePolicy) Type() string {
	return "SpotInstancePolicy"
}

func spotInstancePolicyToPb(st *SpotInstancePolicy) (*spotInstancePolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := spotInstancePolicyPb(*st)
	return &pb, nil
}

func spotInstancePolicyFromPb(pb *spotInstancePolicyPb) (*SpotInstancePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := SpotInstancePolicy(*pb)
	return &st, nil
}

// Start a warehouse
type StartRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func startRequestToPb(st *StartRequest) (*startRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type startRequestPb struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func startRequestFromPb(pb *startRequestPb) (*StartRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartRequest{}
	st.Id = pb.Id

	return st, nil
}

type StartWarehouseResponse struct {
}

func startWarehouseResponseToPb(st *StartWarehouseResponse) (*startWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startWarehouseResponsePb{}

	return pb, nil
}

func (st *StartWarehouseResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startWarehouseResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startWarehouseResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartWarehouseResponse) MarshalJSON() ([]byte, error) {
	pb, err := startWarehouseResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type startWarehouseResponsePb struct {
}

func startWarehouseResponseFromPb(pb *startWarehouseResponsePb) (*StartWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartWarehouseResponse{}

	return st, nil
}

// State of the warehouse
type State string
type statePb string

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

type StatementParameterListItem struct {
	// The name of a parameter marker to be substituted in the statement.
	// Wire name: 'name'
	Name string
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	// Wire name: 'type'
	Type string
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func statementParameterListItemToPb(st *StatementParameterListItem) (*statementParameterListItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &statementParameterListItemPb{}
	pb.Name = st.Name

	pb.Type = st.Type

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type statementParameterListItemPb struct {
	// The name of a parameter marker to be substituted in the statement.
	Name string `json:"name"`
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type string `json:"type,omitempty"`
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func statementParameterListItemFromPb(pb *statementParameterListItemPb) (*StatementParameterListItem, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StatementParameterListItem{}
	st.Name = pb.Name
	st.Type = pb.Type
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *statementParameterListItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st statementParameterListItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StatementResponse struct {
	// The result manifest provides schema and metadata for the result set.
	// Wire name: 'manifest'
	Manifest *ResultManifest

	// Wire name: 'result'
	Result *ResultData
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string
	// The status response includes execution state and if relevant, error
	// information.
	// Wire name: 'status'
	Status *StatementStatus

	ForceSendFields []string `tf:"-"`
}

func StatementResponseToPb(st *StatementResponse) (*StatementResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &StatementResponsePb{}
	manifestPb, err := resultManifestToPb(st.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestPb != nil {
		pb.Manifest = manifestPb
	}

	resultPb, err := resultDataToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}

	pb.StatementId = st.StatementId

	statusPb, err := statementStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *StatementResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &StatementResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := StatementResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StatementResponse) MarshalJSON() ([]byte, error) {
	pb, err := StatementResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StatementResponsePb struct {
	// The result manifest provides schema and metadata for the result set.
	Manifest *resultManifestPb `json:"manifest,omitempty"`

	Result *resultDataPb `json:"result,omitempty"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string `json:"statement_id,omitempty"`
	// The status response includes execution state and if relevant, error
	// information.
	Status *statementStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func StatementResponseFromPb(pb *StatementResponsePb) (*StatementResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StatementResponse{}
	manifestField, err := resultManifestFromPb(pb.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestField != nil {
		st.Manifest = manifestField
	}
	resultField, err := resultDataFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}
	st.StatementId = pb.StatementId
	statusField, err := statementStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *StatementResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StatementResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Statement execution state: - `PENDING`: waiting for warehouse - `RUNNING`:
// running - `SUCCEEDED`: execution was successful, result data available for
// fetch - `FAILED`: execution failed; reason for failure described in
// accomanying error message - `CANCELED`: user canceled; can come from explicit
// cancel call, or timeout with `on_wait_timeout=CANCEL` - `CLOSED`: execution
// successful, and statement closed; result no longer available for fetch
type StatementState string
type statementStatePb string

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

// Type always returns StatementState to satisfy [pflag.Value] interface
func (f *StatementState) Type() string {
	return "StatementState"
}

func statementStateToPb(st *StatementState) (*statementStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := statementStatePb(*st)
	return &pb, nil
}

func statementStateFromPb(pb *statementStatePb) (*StatementState, error) {
	if pb == nil {
		return nil, nil
	}
	st := StatementState(*pb)
	return &st, nil
}

// The status response includes execution state and if relevant, error
// information.
type StatementStatus struct {

	// Wire name: 'error'
	Error *ServiceError
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accomanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	// Wire name: 'state'
	State StatementState
}

func statementStatusToPb(st *StatementStatus) (*statementStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &statementStatusPb{}
	errorPb, err := serviceErrorToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}

	pb.State = st.State

	return pb, nil
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

type statementStatusPb struct {
	Error *serviceErrorPb `json:"error,omitempty"`
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accomanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
	State StatementState `json:"state,omitempty"`
}

func statementStatusFromPb(pb *statementStatusPb) (*StatementStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StatementStatus{}
	errorField, err := serviceErrorFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	st.State = pb.State

	return st, nil
}

// Health status of the warehouse.
type Status string
type statusPb string

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

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

func statusToPb(st *Status) (*statusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := statusPb(*st)
	return &pb, nil
}

func statusFromPb(pb *statusPb) (*Status, error) {
	if pb == nil {
		return nil, nil
	}
	st := Status(*pb)
	return &st, nil
}

// Stop a warehouse
type StopRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func stopRequestToPb(st *StopRequest) (*stopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type stopRequestPb struct {
	// Required. Id of the SQL warehouse.
	Id string `json:"-" url:"-"`
}

func stopRequestFromPb(pb *stopRequestPb) (*StopRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopRequest{}
	st.Id = pb.Id

	return st, nil
}

type StopWarehouseResponse struct {
}

func stopWarehouseResponseToPb(st *StopWarehouseResponse) (*stopWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopWarehouseResponsePb{}

	return pb, nil
}

func (st *StopWarehouseResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stopWarehouseResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stopWarehouseResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StopWarehouseResponse) MarshalJSON() ([]byte, error) {
	pb, err := stopWarehouseResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type stopWarehouseResponsePb struct {
}

func stopWarehouseResponseFromPb(pb *stopWarehouseResponsePb) (*StopWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopWarehouseResponse{}

	return st, nil
}

type Success struct {

	// Wire name: 'message'
	Message SuccessMessage
}

func successToPb(st *Success) (*successPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &successPb{}
	pb.Message = st.Message

	return pb, nil
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

type successPb struct {
	Message SuccessMessage `json:"message,omitempty"`
}

func successFromPb(pb *successPb) (*Success, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Success{}
	st.Message = pb.Message

	return st, nil
}

type SuccessMessage string
type successMessagePb string

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

// Type always returns SuccessMessage to satisfy [pflag.Value] interface
func (f *SuccessMessage) Type() string {
	return "SuccessMessage"
}

func successMessageToPb(st *SuccessMessage) (*successMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := successMessagePb(*st)
	return &pb, nil
}

func successMessageFromPb(pb *successMessagePb) (*SuccessMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := SuccessMessage(*pb)
	return &st, nil
}

type TerminationReason struct {
	// status code indicating why the cluster was terminated
	// Wire name: 'code'
	Code TerminationReasonCode
	// list of parameters that provide additional information about why the
	// cluster was terminated
	// Wire name: 'parameters'
	Parameters map[string]string
	// type of the termination
	// Wire name: 'type'
	Type TerminationReasonType
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
	st.Code = pb.Code
	st.Parameters = pb.Parameters
	st.Type = pb.Type

	return st, nil
}

// status code indicating why the cluster was terminated
type TerminationReasonCode string
type terminationReasonCodePb string

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

type TextValue struct {

	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func textValueToPb(st *TextValue) (*textValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &textValuePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type textValuePb struct {
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func textValueFromPb(pb *textValuePb) (*TextValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TextValue{}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *textValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st textValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TimeRange struct {
	// The end time in milliseconds.
	// Wire name: 'end_time_ms'
	EndTimeMs int64
	// The start time in milliseconds.
	// Wire name: 'start_time_ms'
	StartTimeMs int64

	ForceSendFields []string `tf:"-"`
}

func timeRangeToPb(st *TimeRange) (*timeRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &timeRangePb{}
	pb.EndTimeMs = st.EndTimeMs

	pb.StartTimeMs = st.StartTimeMs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type timeRangePb struct {
	// The end time in milliseconds.
	EndTimeMs int64 `json:"end_time_ms,omitempty" url:"end_time_ms,omitempty"`
	// The start time in milliseconds.
	StartTimeMs int64 `json:"start_time_ms,omitempty" url:"start_time_ms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func timeRangeFromPb(pb *timeRangePb) (*TimeRange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TimeRange{}
	st.EndTimeMs = pb.EndTimeMs
	st.StartTimeMs = pb.StartTimeMs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *timeRangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st timeRangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	// Wire name: 'new_owner'
	NewOwner string

	ForceSendFields []string `tf:"-"`
}

func transferOwnershipObjectIdToPb(st *TransferOwnershipObjectId) (*transferOwnershipObjectIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transferOwnershipObjectIdPb{}
	pb.NewOwner = st.NewOwner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type transferOwnershipObjectIdPb struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner string `json:"new_owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func transferOwnershipObjectIdFromPb(pb *transferOwnershipObjectIdPb) (*TransferOwnershipObjectId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransferOwnershipObjectId{}
	st.NewOwner = pb.NewOwner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *transferOwnershipObjectIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st transferOwnershipObjectIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Transfer object ownership
type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	// Wire name: 'new_owner'
	NewOwner string
	// The ID of the object on which to change ownership.
	// Wire name: 'objectId'
	ObjectId TransferOwnershipObjectId `tf:"-"`
	// The type of object on which to change ownership.
	// Wire name: 'objectType'
	ObjectType OwnableObjectType `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func transferOwnershipRequestToPb(st *TransferOwnershipRequest) (*transferOwnershipRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transferOwnershipRequestPb{}
	pb.NewOwner = st.NewOwner

	objectIdPb, err := transferOwnershipObjectIdToPb(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type transferOwnershipRequestPb struct {
	// Email address for the new owner, who must exist in the workspace.
	NewOwner string `json:"new_owner,omitempty"`
	// The ID of the object on which to change ownership.
	ObjectId transferOwnershipObjectIdPb `json:"-" url:"-"`
	// The type of object on which to change ownership.
	ObjectType OwnableObjectType `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func transferOwnershipRequestFromPb(pb *transferOwnershipRequestPb) (*TransferOwnershipRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransferOwnershipRequest{}
	st.NewOwner = pb.NewOwner
	objectIdField, err := transferOwnershipObjectIdFromPb(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *transferOwnershipRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st transferOwnershipRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete an alert
type TrashAlertRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func trashAlertRequestToPb(st *TrashAlertRequest) (*trashAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashAlertRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type trashAlertRequestPb struct {
	Id string `json:"-" url:"-"`
}

func trashAlertRequestFromPb(pb *trashAlertRequestPb) (*TrashAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashAlertRequest{}
	st.Id = pb.Id

	return st, nil
}

// Delete an alert
type TrashAlertV2Request struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func trashAlertV2RequestToPb(st *TrashAlertV2Request) (*trashAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashAlertV2RequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type trashAlertV2RequestPb struct {
	Id string `json:"-" url:"-"`
}

func trashAlertV2RequestFromPb(pb *trashAlertV2RequestPb) (*TrashAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashAlertV2Request{}
	st.Id = pb.Id

	return st, nil
}

// Delete a query
type TrashQueryRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func trashQueryRequestToPb(st *TrashQueryRequest) (*trashQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashQueryRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

type trashQueryRequestPb struct {
	Id string `json:"-" url:"-"`
}

func trashQueryRequestFromPb(pb *trashQueryRequestPb) (*TrashQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashQueryRequest{}
	st.Id = pb.Id

	return st, nil
}

type UpdateAlertRequest struct {

	// Wire name: 'alert'
	Alert *UpdateAlertRequestAlert

	// Wire name: 'id'
	Id string `tf:"-"`
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
	UpdateMask string
}

func updateAlertRequestToPb(st *UpdateAlertRequest) (*updateAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAlertRequestPb{}
	alertPb, err := updateAlertRequestAlertToPb(st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = alertPb
	}

	pb.Id = st.Id

	pb.UpdateMask = st.UpdateMask

	return pb, nil
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

type updateAlertRequestPb struct {
	Alert *updateAlertRequestAlertPb `json:"alert,omitempty"`

	Id string `json:"-" url:"-"`
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

func updateAlertRequestFromPb(pb *updateAlertRequestPb) (*UpdateAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertRequest{}
	alertField, err := updateAlertRequestAlertFromPb(pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = alertField
	}
	st.Id = pb.Id
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

type UpdateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int

	ForceSendFields []string `tf:"-"`
}

func updateAlertRequestAlertToPb(st *UpdateAlertRequestAlert) (*updateAlertRequestAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAlertRequestAlertPb{}
	conditionPb, err := alertConditionToPb(st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = conditionPb
	}

	pb.CustomBody = st.CustomBody

	pb.CustomSubject = st.CustomSubject

	pb.DisplayName = st.DisplayName

	pb.NotifyOnOk = st.NotifyOnOk

	pb.OwnerUserName = st.OwnerUserName

	pb.QueryId = st.QueryId

	pb.SecondsToRetrigger = st.SecondsToRetrigger

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateAlertRequestAlertPb struct {
	// Trigger conditions of the alert.
	Condition *alertConditionPb `json:"condition,omitempty"`
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string `json:"custom_body,omitempty"`
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string `json:"custom_subject,omitempty"`
	// The display name of the alert.
	DisplayName string `json:"display_name,omitempty"`
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool `json:"notify_on_ok,omitempty"`
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// UUID of the query attached to the alert.
	QueryId string `json:"query_id,omitempty"`
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateAlertRequestAlertFromPb(pb *updateAlertRequestAlertPb) (*UpdateAlertRequestAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertRequestAlert{}
	conditionField, err := alertConditionFromPb(pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = conditionField
	}
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	st.DisplayName = pb.DisplayName
	st.NotifyOnOk = pb.NotifyOnOk
	st.OwnerUserName = pb.OwnerUserName
	st.QueryId = pb.QueryId
	st.SecondsToRetrigger = pb.SecondsToRetrigger

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateAlertRequestAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateAlertRequestAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateAlertV2Request struct {

	// Wire name: 'alert'
	Alert *AlertV2
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string `tf:"-"`
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
	UpdateMask string
}

func updateAlertV2RequestToPb(st *UpdateAlertV2Request) (*updateAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAlertV2RequestPb{}
	alertPb, err := alertV2ToPb(st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = alertPb
	}

	pb.Id = st.Id

	pb.UpdateMask = st.UpdateMask

	return pb, nil
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

type updateAlertV2RequestPb struct {
	Alert *alertV2Pb `json:"alert,omitempty"`
	// UUID identifying the alert.
	Id string `json:"-" url:"-"`
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

func updateAlertV2RequestFromPb(pb *updateAlertV2RequestPb) (*UpdateAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertV2Request{}
	alertField, err := alertV2FromPb(pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = alertField
	}
	st.Id = pb.Id
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

type UpdateQueryRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'query'
	Query *UpdateQueryRequestQuery
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
	UpdateMask string
}

func updateQueryRequestToPb(st *UpdateQueryRequest) (*updateQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateQueryRequestPb{}
	pb.Id = st.Id

	queryPb, err := updateQueryRequestQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	pb.UpdateMask = st.UpdateMask

	return pb, nil
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

type updateQueryRequestPb struct {
	Id string `json:"-" url:"-"`

	Query *updateQueryRequestQueryPb `json:"query,omitempty"`
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

func updateQueryRequestFromPb(pb *updateQueryRequestPb) (*UpdateQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateQueryRequest{}
	st.Id = pb.Id
	queryField, err := updateQueryRequestQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

type UpdateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string

	// Wire name: 'tags'
	Tags []string
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId string

	ForceSendFields []string `tf:"-"`
}

func updateQueryRequestQueryToPb(st *UpdateQueryRequestQuery) (*updateQueryRequestQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateQueryRequestQueryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit

	pb.Catalog = st.Catalog

	pb.Description = st.Description

	pb.DisplayName = st.DisplayName

	pb.OwnerUserName = st.OwnerUserName

	var parametersPb []queryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := queryParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb

	pb.QueryText = st.QueryText

	pb.RunAsMode = st.RunAsMode

	pb.Schema = st.Schema

	pb.Tags = st.Tags

	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateQueryRequestQueryPb struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool `json:"apply_auto_limit,omitempty"`
	// Name of the catalog where this query will be executed.
	Catalog string `json:"catalog,omitempty"`
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string `json:"description,omitempty"`
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string `json:"display_name,omitempty"`
	// Username of the user that owns the query.
	OwnerUserName string `json:"owner_user_name,omitempty"`
	// List of query parameter definitions.
	Parameters []queryParameterPb `json:"parameters,omitempty"`
	// Text of the query to be run.
	QueryText string `json:"query_text,omitempty"`
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode `json:"run_as_mode,omitempty"`
	// Name of the schema where this query will be executed.
	Schema string `json:"schema,omitempty"`

	Tags []string `json:"tags,omitempty"`
	// ID of the SQL warehouse attached to the query.
	WarehouseId string `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateQueryRequestQueryFromPb(pb *updateQueryRequestQueryPb) (*UpdateQueryRequestQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateQueryRequestQuery{}
	st.ApplyAutoLimit = pb.ApplyAutoLimit
	st.Catalog = pb.Catalog
	st.Description = pb.Description
	st.DisplayName = pb.DisplayName
	st.OwnerUserName = pb.OwnerUserName

	var parametersField []QueryParameter
	for _, itemPb := range pb.Parameters {
		item, err := queryParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.QueryText = pb.QueryText
	st.RunAsMode = pb.RunAsMode
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateQueryRequestQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateQueryRequestQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type UpdateVisualizationRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
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
	UpdateMask string

	// Wire name: 'visualization'
	Visualization *UpdateVisualizationRequestVisualization
}

func updateVisualizationRequestToPb(st *UpdateVisualizationRequest) (*updateVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVisualizationRequestPb{}
	pb.Id = st.Id

	pb.UpdateMask = st.UpdateMask

	visualizationPb, err := updateVisualizationRequestVisualizationToPb(st.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationPb != nil {
		pb.Visualization = visualizationPb
	}

	return pb, nil
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

type updateVisualizationRequestPb struct {
	Id string `json:"-" url:"-"`
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

	Visualization *updateVisualizationRequestVisualizationPb `json:"visualization,omitempty"`
}

func updateVisualizationRequestFromPb(pb *updateVisualizationRequestPb) (*UpdateVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateVisualizationRequest{}
	st.Id = pb.Id
	st.UpdateMask = pb.UpdateMask
	visualizationField, err := updateVisualizationRequestVisualizationFromPb(pb.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationField != nil {
		st.Visualization = visualizationField
	}

	return st, nil
}

type UpdateVisualizationRequestVisualization struct {
	// The display name of the visualization.
	// Wire name: 'display_name'
	DisplayName string
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type string

	ForceSendFields []string `tf:"-"`
}

func updateVisualizationRequestVisualizationToPb(st *UpdateVisualizationRequestVisualization) (*updateVisualizationRequestVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVisualizationRequestVisualizationPb{}
	pb.DisplayName = st.DisplayName

	pb.SerializedOptions = st.SerializedOptions

	pb.SerializedQueryPlan = st.SerializedQueryPlan

	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateVisualizationRequestVisualizationPb struct {
	// The display name of the visualization.
	DisplayName string `json:"display_name,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateVisualizationRequestVisualizationFromPb(pb *updateVisualizationRequestVisualizationPb) (*UpdateVisualizationRequestVisualization, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateVisualizationRequestVisualization{}
	st.DisplayName = pb.DisplayName
	st.SerializedOptions = pb.SerializedOptions
	st.SerializedQueryPlan = pb.SerializedQueryPlan
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateVisualizationRequestVisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateVisualizationRequestVisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type User struct {

	// Wire name: 'email'
	Email string

	// Wire name: 'id'
	Id int

	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func userToPb(st *User) (*userPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &userPb{}
	pb.Email = st.Email

	pb.Id = st.Id

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type userPb struct {
	Email string `json:"email,omitempty"`

	Id int `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func userFromPb(pb *userPb) (*User, error) {
	if pb == nil {
		return nil, nil
	}
	st := &User{}
	st.Email = pb.Email
	st.Id = pb.Id
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *userPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st userPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Visualization struct {
	// The timestamp indicating when the visualization was created.
	// Wire name: 'create_time'
	CreateTime string
	// The display name of the visualization.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying the visualization.
	// Wire name: 'id'
	Id string
	// UUID of the query that the visualization is attached to.
	// Wire name: 'query_id'
	QueryId string
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type string
	// The timestamp indicating when the visualization was updated.
	// Wire name: 'update_time'
	UpdateTime string

	ForceSendFields []string `tf:"-"`
}

func visualizationToPb(st *Visualization) (*visualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &visualizationPb{}
	pb.CreateTime = st.CreateTime

	pb.DisplayName = st.DisplayName

	pb.Id = st.Id

	pb.QueryId = st.QueryId

	pb.SerializedOptions = st.SerializedOptions

	pb.SerializedQueryPlan = st.SerializedQueryPlan

	pb.Type = st.Type

	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type visualizationPb struct {
	// The timestamp indicating when the visualization was created.
	CreateTime string `json:"create_time,omitempty"`
	// The display name of the visualization.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying the visualization.
	Id string `json:"id,omitempty"`
	// UUID of the query that the visualization is attached to.
	QueryId string `json:"query_id,omitempty"`
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string `json:"serialized_options,omitempty"`
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	// The type of visualization: counter, table, funnel, and so on.
	Type string `json:"type,omitempty"`
	// The timestamp indicating when the visualization was updated.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func visualizationFromPb(pb *visualizationPb) (*Visualization, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Visualization{}
	st.CreateTime = pb.CreateTime
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.QueryId = pb.QueryId
	st.SerializedOptions = pb.SerializedOptions
	st.SerializedQueryPlan = pb.SerializedQueryPlan
	st.Type = pb.Type
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *visualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st visualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehouseAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func warehouseAccessControlRequestToPb(st *WarehouseAccessControlRequest) (*warehouseAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehouseAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type warehouseAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func warehouseAccessControlRequestFromPb(pb *warehouseAccessControlRequestPb) (*WarehouseAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehouseAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *warehouseAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehouseAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []WarehousePermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func warehouseAccessControlResponseToPb(st *WarehouseAccessControlResponse) (*warehouseAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehouseAccessControlResponsePb{}

	var allPermissionsPb []warehousePermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := warehousePermissionToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type warehouseAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []warehousePermissionPb `json:"all_permissions,omitempty"`
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

func warehouseAccessControlResponseFromPb(pb *warehouseAccessControlResponsePb) (*WarehouseAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehouseAccessControlResponse{}

	var allPermissionsField []WarehousePermission
	for _, itemPb := range pb.AllPermissions {
		item, err := warehousePermissionFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *warehouseAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehouseAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel

	ForceSendFields []string `tf:"-"`
}

func warehousePermissionToPb(st *WarehousePermission) (*warehousePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type warehousePermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func warehousePermissionFromPb(pb *warehousePermissionPb) (*WarehousePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *warehousePermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehousePermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type WarehousePermissionLevel string
type warehousePermissionLevelPb string

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

// Type always returns WarehousePermissionLevel to satisfy [pflag.Value] interface
func (f *WarehousePermissionLevel) Type() string {
	return "WarehousePermissionLevel"
}

func warehousePermissionLevelToPb(st *WarehousePermissionLevel) (*warehousePermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := warehousePermissionLevelPb(*st)
	return &pb, nil
}

func warehousePermissionLevelFromPb(pb *warehousePermissionLevelPb) (*WarehousePermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := WarehousePermissionLevel(*pb)
	return &st, nil
}

type WarehousePermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []WarehouseAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func warehousePermissionsToPb(st *WarehousePermissions) (*warehousePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionsPb{}

	var accessControlListPb []warehouseAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := warehouseAccessControlResponseToPb(&item)
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

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type warehousePermissionsPb struct {
	AccessControlList []warehouseAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func warehousePermissionsFromPb(pb *warehousePermissionsPb) (*WarehousePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissions{}

	var accessControlListField []WarehouseAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := warehouseAccessControlResponseFromPb(&itemPb)
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *warehousePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehousePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel

	ForceSendFields []string `tf:"-"`
}

func warehousePermissionsDescriptionToPb(st *WarehousePermissionsDescription) (*warehousePermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type warehousePermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel WarehousePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func warehousePermissionsDescriptionFromPb(pb *warehousePermissionsDescriptionPb) (*WarehousePermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *warehousePermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehousePermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []WarehouseAccessControlRequest
	// The SQL warehouse for which to get or manage permissions.
	// Wire name: 'warehouse_id'
	WarehouseId string `tf:"-"`
}

func warehousePermissionsRequestToPb(st *WarehousePermissionsRequest) (*warehousePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionsRequestPb{}

	var accessControlListPb []warehouseAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := warehouseAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	pb.WarehouseId = st.WarehouseId

	return pb, nil
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

type warehousePermissionsRequestPb struct {
	AccessControlList []warehouseAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string `json:"-" url:"-"`
}

func warehousePermissionsRequestFromPb(pb *warehousePermissionsRequestPb) (*WarehousePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissionsRequest{}

	var accessControlListField []WarehouseAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := warehouseAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.WarehouseId = pb.WarehouseId

	return st, nil
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	// Wire name: 'enabled'
	Enabled bool
	// Warehouse type: `PRO` or `CLASSIC`.
	// Wire name: 'warehouse_type'
	WarehouseType WarehouseTypePairWarehouseType

	ForceSendFields []string `tf:"-"`
}

func warehouseTypePairToPb(st *WarehouseTypePair) (*warehouseTypePairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehouseTypePairPb{}
	pb.Enabled = st.Enabled

	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type warehouseTypePairPb struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled bool `json:"enabled,omitempty"`
	// Warehouse type: `PRO` or `CLASSIC`.
	WarehouseType WarehouseTypePairWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func warehouseTypePairFromPb(pb *warehouseTypePairPb) (*WarehouseTypePair, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehouseTypePair{}
	st.Enabled = pb.Enabled
	st.WarehouseType = pb.WarehouseType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *warehouseTypePairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehouseTypePairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Warehouse type: `PRO` or `CLASSIC`.
type WarehouseTypePairWarehouseType string
type warehouseTypePairWarehouseTypePb string

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

// Type always returns WarehouseTypePairWarehouseType to satisfy [pflag.Value] interface
func (f *WarehouseTypePairWarehouseType) Type() string {
	return "WarehouseTypePairWarehouseType"
}

func warehouseTypePairWarehouseTypeToPb(st *WarehouseTypePairWarehouseType) (*warehouseTypePairWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := warehouseTypePairWarehouseTypePb(*st)
	return &pb, nil
}

func warehouseTypePairWarehouseTypeFromPb(pb *warehouseTypePairWarehouseTypePb) (*WarehouseTypePairWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := WarehouseTypePairWarehouseType(*pb)
	return &st, nil
}

type Widget struct {
	// The unique ID for this widget.
	// Wire name: 'id'
	Id string

	// Wire name: 'options'
	Options *WidgetOptions
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	// Wire name: 'visualization'
	Visualization *LegacyVisualization
	// Unused field.
	// Wire name: 'width'
	Width int

	ForceSendFields []string `tf:"-"`
}

func widgetToPb(st *Widget) (*widgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetPb{}
	pb.Id = st.Id

	optionsPb, err := widgetOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	visualizationPb, err := legacyVisualizationToPb(st.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationPb != nil {
		pb.Visualization = visualizationPb
	}

	pb.Width = st.Width

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type widgetPb struct {
	// The unique ID for this widget.
	Id string `json:"id,omitempty"`

	Options *widgetOptionsPb `json:"options,omitempty"`
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization *legacyVisualizationPb `json:"visualization,omitempty"`
	// Unused field.
	Width int `json:"width,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func widgetFromPb(pb *widgetPb) (*Widget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Widget{}
	st.Id = pb.Id
	optionsField, err := widgetOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	visualizationField, err := legacyVisualizationFromPb(pb.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationField != nil {
		st.Visualization = visualizationField
	}
	st.Width = pb.Width

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *widgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st widgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WidgetOptions struct {
	// Timestamp when this object was created
	// Wire name: 'created_at'
	CreatedAt string
	// Custom description of the widget
	// Wire name: 'description'
	Description string
	// Whether this widget is hidden on the dashboard.
	// Wire name: 'isHidden'
	IsHidden bool
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	// Wire name: 'parameterMappings'
	ParameterMappings any
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	// Wire name: 'position'
	Position *WidgetPosition
	// Custom title of the widget
	// Wire name: 'title'
	Title string
	// Timestamp of the last time this object was updated.
	// Wire name: 'updated_at'
	UpdatedAt string

	ForceSendFields []string `tf:"-"`
}

func widgetOptionsToPb(st *WidgetOptions) (*widgetOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetOptionsPb{}
	pb.CreatedAt = st.CreatedAt

	pb.Description = st.Description

	pb.IsHidden = st.IsHidden

	pb.ParameterMappings = st.ParameterMappings

	positionPb, err := widgetPositionToPb(st.Position)
	if err != nil {
		return nil, err
	}
	if positionPb != nil {
		pb.Position = positionPb
	}

	pb.Title = st.Title

	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type widgetOptionsPb struct {
	// Timestamp when this object was created
	CreatedAt string `json:"created_at,omitempty"`
	// Custom description of the widget
	Description string `json:"description,omitempty"`
	// Whether this widget is hidden on the dashboard.
	IsHidden bool `json:"isHidden,omitempty"`
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings any `json:"parameterMappings,omitempty"`
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position *widgetPositionPb `json:"position,omitempty"`
	// Custom title of the widget
	Title string `json:"title,omitempty"`
	// Timestamp of the last time this object was updated.
	UpdatedAt string `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func widgetOptionsFromPb(pb *widgetOptionsPb) (*WidgetOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WidgetOptions{}
	st.CreatedAt = pb.CreatedAt
	st.Description = pb.Description
	st.IsHidden = pb.IsHidden
	st.ParameterMappings = pb.ParameterMappings
	positionField, err := widgetPositionFromPb(pb.Position)
	if err != nil {
		return nil, err
	}
	if positionField != nil {
		st.Position = positionField
	}
	st.Title = pb.Title
	st.UpdatedAt = pb.UpdatedAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *widgetOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st widgetOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	// Wire name: 'autoHeight'
	AutoHeight bool
	// column in the dashboard grid. Values start with 0
	// Wire name: 'col'
	Col int
	// row in the dashboard grid. Values start with 0
	// Wire name: 'row'
	Row int
	// width of the widget measured in dashboard grid cells
	// Wire name: 'sizeX'
	SizeX int
	// height of the widget measured in dashboard grid cells
	// Wire name: 'sizeY'
	SizeY int

	ForceSendFields []string `tf:"-"`
}

func widgetPositionToPb(st *WidgetPosition) (*widgetPositionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetPositionPb{}
	pb.AutoHeight = st.AutoHeight

	pb.Col = st.Col

	pb.Row = st.Row

	pb.SizeX = st.SizeX

	pb.SizeY = st.SizeY

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type widgetPositionPb struct {
	// reserved for internal use
	AutoHeight bool `json:"autoHeight,omitempty"`
	// column in the dashboard grid. Values start with 0
	Col int `json:"col,omitempty"`
	// row in the dashboard grid. Values start with 0
	Row int `json:"row,omitempty"`
	// width of the widget measured in dashboard grid cells
	SizeX int `json:"sizeX,omitempty"`
	// height of the widget measured in dashboard grid cells
	SizeY int `json:"sizeY,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func widgetPositionFromPb(pb *widgetPositionPb) (*WidgetPosition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WidgetPosition{}
	st.AutoHeight = pb.AutoHeight
	st.Col = pb.Col
	st.Row = pb.Row
	st.SizeX = pb.SizeX
	st.SizeY = pb.SizeY

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *widgetPositionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st widgetPositionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
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
