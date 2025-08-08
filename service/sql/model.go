// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/sql/sqlpb"
)

type AccessControl struct {

	// Wire name: 'group_name'
	GroupName string ``
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel ``

	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AccessControl) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccessControl) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AccessControlToPb(st *AccessControl) (*sqlpb.AccessControlPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AccessControlPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := PermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AccessControlFromPb(pb *sqlpb.AccessControlPb) (*AccessControl, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccessControl{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := PermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func AggregationToPb(st *Aggregation) (*sqlpb.AggregationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.AggregationPb(*st)
	return &pb, nil
}

func AggregationFromPb(pb *sqlpb.AggregationPb) (*Aggregation, error) {
	if pb == nil {
		return nil, nil
	}
	st := Aggregation(*pb)
	return &st, nil
}

type Alert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition ``
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string ``
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string ``
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string ``
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string ``
	// The workspace state of the alert. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState ``
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool ``
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string ``
	// The workspace path of the folder containing the alert.
	// Wire name: 'parent_path'
	ParentPath string ``
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string ``
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int ``
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	// Wire name: 'state'
	State AlertState ``
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	// Wire name: 'trigger_time'
	TriggerTime *time.Time ``
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime      *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *Alert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Alert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertToPb(st *Alert) (*sqlpb.AlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertPb{}
	conditionPb, err := AlertConditionToPb(st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = conditionPb
	}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.CustomBody = st.CustomBody
	pb.CustomSubject = st.CustomSubject
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id
	lifecycleStatePb, err := LifecycleStateToPb(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}
	pb.NotifyOnOk = st.NotifyOnOk
	pb.OwnerUserName = st.OwnerUserName
	pb.ParentPath = st.ParentPath
	pb.QueryId = st.QueryId
	pb.SecondsToRetrigger = st.SecondsToRetrigger
	statePb, err := AlertStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	triggerTimePb, err := timestampToPb(st.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimePb != nil {
		pb.TriggerTime = *triggerTimePb
	}
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertFromPb(pb *sqlpb.AlertPb) (*Alert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Alert{}
	conditionField, err := AlertConditionFromPb(pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = conditionField
	}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	lifecycleStateField, err := LifecycleStateFromPb(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	st.NotifyOnOk = pb.NotifyOnOk
	st.OwnerUserName = pb.OwnerUserName
	st.ParentPath = pb.ParentPath
	st.QueryId = pb.QueryId
	st.SecondsToRetrigger = pb.SecondsToRetrigger
	stateField, err := AlertStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	triggerTimeField, err := timestampFromPb(&pb.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimeField != nil {
		st.TriggerTime = triggerTimeField
	}
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AlertCondition struct {
	// Alert state if result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertState ``
	// Operator used for comparison in alert evaluation.
	// Wire name: 'op'
	Op AlertOperator ``
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	// Wire name: 'operand'
	Operand *AlertConditionOperand ``
	// Threshold value used for comparison in alert evaluation.
	// Wire name: 'threshold'
	Threshold *AlertConditionThreshold ``
}

func AlertConditionToPb(st *AlertCondition) (*sqlpb.AlertConditionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertConditionPb{}
	emptyResultStatePb, err := AlertStateToPb(&st.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStatePb != nil {
		pb.EmptyResultState = *emptyResultStatePb
	}
	opPb, err := AlertOperatorToPb(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}
	operandPb, err := AlertConditionOperandToPb(st.Operand)
	if err != nil {
		return nil, err
	}
	if operandPb != nil {
		pb.Operand = operandPb
	}
	thresholdPb, err := AlertConditionThresholdToPb(st.Threshold)
	if err != nil {
		return nil, err
	}
	if thresholdPb != nil {
		pb.Threshold = thresholdPb
	}

	return pb, nil
}

func AlertConditionFromPb(pb *sqlpb.AlertConditionPb) (*AlertCondition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertCondition{}
	emptyResultStateField, err := AlertStateFromPb(&pb.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStateField != nil {
		st.EmptyResultState = *emptyResultStateField
	}
	opField, err := AlertOperatorFromPb(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
	operandField, err := AlertConditionOperandFromPb(pb.Operand)
	if err != nil {
		return nil, err
	}
	if operandField != nil {
		st.Operand = operandField
	}
	thresholdField, err := AlertConditionThresholdFromPb(pb.Threshold)
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
	Column *AlertOperandColumn ``
}

func AlertConditionOperandToPb(st *AlertConditionOperand) (*sqlpb.AlertConditionOperandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertConditionOperandPb{}
	columnPb, err := AlertOperandColumnToPb(st.Column)
	if err != nil {
		return nil, err
	}
	if columnPb != nil {
		pb.Column = columnPb
	}

	return pb, nil
}

func AlertConditionOperandFromPb(pb *sqlpb.AlertConditionOperandPb) (*AlertConditionOperand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConditionOperand{}
	columnField, err := AlertOperandColumnFromPb(pb.Column)
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
	Value *AlertOperandValue ``
}

func AlertConditionThresholdToPb(st *AlertConditionThreshold) (*sqlpb.AlertConditionThresholdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertConditionThresholdPb{}
	valuePb, err := AlertOperandValueToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	return pb, nil
}

func AlertConditionThresholdFromPb(pb *sqlpb.AlertConditionThresholdPb) (*AlertConditionThreshold, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConditionThreshold{}
	valueField, err := AlertOperandValueFromPb(pb.Value)
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

func AlertEvaluationStateToPb(st *AlertEvaluationState) (*sqlpb.AlertEvaluationStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.AlertEvaluationStatePb(*st)
	return &pb, nil
}

func AlertEvaluationStateFromPb(pb *sqlpb.AlertEvaluationStatePb) (*AlertEvaluationState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertEvaluationState(*pb)
	return &st, nil
}

type AlertOperandColumn struct {

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertOperandColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertOperandColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertOperandColumnToPb(st *AlertOperandColumn) (*sqlpb.AlertOperandColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertOperandColumnPb{}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertOperandColumnFromPb(pb *sqlpb.AlertOperandColumnPb) (*AlertOperandColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertOperandColumn{}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AlertOperandValue struct {

	// Wire name: 'bool_value'
	BoolValue bool ``

	// Wire name: 'double_value'
	DoubleValue float64 ``

	// Wire name: 'string_value'
	StringValue     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertOperandValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertOperandValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertOperandValueToPb(st *AlertOperandValue) (*sqlpb.AlertOperandValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertOperandValuePb{}
	pb.BoolValue = st.BoolValue
	pb.DoubleValue = st.DoubleValue
	pb.StringValue = st.StringValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertOperandValueFromPb(pb *sqlpb.AlertOperandValuePb) (*AlertOperandValue, error) {
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

func AlertOperatorToPb(st *AlertOperator) (*sqlpb.AlertOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.AlertOperatorPb(*st)
	return &pb, nil
}

func AlertOperatorFromPb(pb *sqlpb.AlertOperatorPb) (*AlertOperator, error) {
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
	Column string ``
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string ``
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string ``
	// State that alert evaluates to when query result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertOptionsEmptyResultState ``
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	// Wire name: 'muted'
	Muted bool ``
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	// Wire name: 'op'
	Op string ``
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	// Wire name: 'value'
	Value           any      ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertOptionsToPb(st *AlertOptions) (*sqlpb.AlertOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertOptionsPb{}
	pb.Column = st.Column
	pb.CustomBody = st.CustomBody
	pb.CustomSubject = st.CustomSubject
	emptyResultStatePb, err := AlertOptionsEmptyResultStateToPb(&st.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStatePb != nil {
		pb.EmptyResultState = *emptyResultStatePb
	}
	pb.Muted = st.Muted
	pb.Op = st.Op
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertOptionsFromPb(pb *sqlpb.AlertOptionsPb) (*AlertOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertOptions{}
	st.Column = pb.Column
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	emptyResultStateField, err := AlertOptionsEmptyResultStateFromPb(&pb.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStateField != nil {
		st.EmptyResultState = *emptyResultStateField
	}
	st.Muted = pb.Muted
	st.Op = pb.Op
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func AlertOptionsEmptyResultStateToPb(st *AlertOptionsEmptyResultState) (*sqlpb.AlertOptionsEmptyResultStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.AlertOptionsEmptyResultStatePb(*st)
	return &pb, nil
}

func AlertOptionsEmptyResultStateFromPb(pb *sqlpb.AlertOptionsEmptyResultStatePb) (*AlertOptionsEmptyResultState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertOptionsEmptyResultState(*pb)
	return &st, nil
}

type AlertQuery struct {
	// The timestamp when this query was created.
	// Wire name: 'created_at'
	CreatedAt string ``
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// Query ID.
	// Wire name: 'id'
	Id string ``
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool ``
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	// Wire name: 'is_draft'
	IsDraft bool ``
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	// Wire name: 'is_safe'
	IsSafe bool ``
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'options'
	Options *QueryOptions ``
	// The text of the query to be run.
	// Wire name: 'query'
	Query string ``

	// Wire name: 'tags'
	Tags []string ``
	// The timestamp at which this query was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string ``
	// The ID of the user who owns the query.
	// Wire name: 'user_id'
	UserId          int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertQueryToPb(st *AlertQuery) (*sqlpb.AlertQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertQueryPb{}
	pb.CreatedAt = st.CreatedAt
	pb.DataSourceId = st.DataSourceId
	pb.Description = st.Description
	pb.Id = st.Id
	pb.IsArchived = st.IsArchived
	pb.IsDraft = st.IsDraft
	pb.IsSafe = st.IsSafe
	pb.Name = st.Name
	optionsPb, err := QueryOptionsToPb(st.Options)
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

func AlertQueryFromPb(pb *sqlpb.AlertQueryPb) (*AlertQuery, error) {
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
	optionsField, err := QueryOptionsFromPb(pb.Options)
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

func AlertStateToPb(st *AlertState) (*sqlpb.AlertStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.AlertStatePb(*st)
	return &pb, nil
}

func AlertStateFromPb(pb *sqlpb.AlertStatePb) (*AlertState, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertState(*pb)
	return &st, nil
}

type AlertV2 struct {
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// Custom description for the alert. support mustache template.
	// Wire name: 'custom_description'
	CustomDescription string ``
	// Custom summary for the alert. support mustache template.
	// Wire name: 'custom_summary'
	CustomSummary string ``
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string ``

	// Wire name: 'evaluation'
	Evaluation *AlertV2Evaluation ``
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string ``
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState ``
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string ``
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	// Wire name: 'parent_path'
	ParentPath string ``
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string ``
	// The run as username or application ID of service principal. On Create and
	// Update, this field can be set to application ID of an active service
	// principal. Setting this field requires the servicePrincipal/user role.
	// Wire name: 'run_as_user_name'
	RunAsUserName string ``

	// Wire name: 'schedule'
	Schedule *CronSchedule ``
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime *time.Time ``
	// ID of the SQL warehouse attached to the alert.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertV2) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertV2ToPb(st *AlertV2) (*sqlpb.AlertV2Pb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertV2Pb{}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.CustomDescription = st.CustomDescription
	pb.CustomSummary = st.CustomSummary
	pb.DisplayName = st.DisplayName
	evaluationPb, err := AlertV2EvaluationToPb(st.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationPb != nil {
		pb.Evaluation = evaluationPb
	}
	pb.Id = st.Id
	lifecycleStatePb, err := LifecycleStateToPb(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}
	pb.OwnerUserName = st.OwnerUserName
	pb.ParentPath = st.ParentPath
	pb.QueryText = st.QueryText
	pb.RunAsUserName = st.RunAsUserName
	schedulePb, err := CronScheduleToPb(st.Schedule)
	if err != nil {
		return nil, err
	}
	if schedulePb != nil {
		pb.Schedule = schedulePb
	}
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertV2FromPb(pb *sqlpb.AlertV2Pb) (*AlertV2, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2{}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.CustomDescription = pb.CustomDescription
	st.CustomSummary = pb.CustomSummary
	st.DisplayName = pb.DisplayName
	evaluationField, err := AlertV2EvaluationFromPb(pb.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationField != nil {
		st.Evaluation = evaluationField
	}
	st.Id = pb.Id
	lifecycleStateField, err := LifecycleStateFromPb(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	st.OwnerUserName = pb.OwnerUserName
	st.ParentPath = pb.ParentPath
	st.QueryText = pb.QueryText
	st.RunAsUserName = pb.RunAsUserName
	scheduleField, err := CronScheduleFromPb(pb.Schedule)
	if err != nil {
		return nil, err
	}
	if scheduleField != nil {
		st.Schedule = scheduleField
	}
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AlertV2Evaluation struct {
	// Operator used for comparison in alert evaluation.
	// Wire name: 'comparison_operator'
	ComparisonOperator ComparisonOperator ``
	// Alert state if result is empty.
	// Wire name: 'empty_result_state'
	EmptyResultState AlertEvaluationState ``
	// Timestamp of the last evaluation.
	// Wire name: 'last_evaluated_at'
	LastEvaluatedAt *time.Time ``
	// User or Notification Destination to notify when alert is triggered.
	// Wire name: 'notification'
	Notification *AlertV2Notification ``
	// Source column from result to use to evaluate alert
	// Wire name: 'source'
	Source *AlertV2OperandColumn ``
	// Latest state of alert evaluation.
	// Wire name: 'state'
	State AlertEvaluationState ``
	// Threshold to user for alert evaluation, can be a column or a value.
	// Wire name: 'threshold'
	Threshold       *AlertV2Operand ``
	ForceSendFields []string        `tf:"-"`
}

func (s *AlertV2Evaluation) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2Evaluation) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertV2EvaluationToPb(st *AlertV2Evaluation) (*sqlpb.AlertV2EvaluationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertV2EvaluationPb{}
	comparisonOperatorPb, err := ComparisonOperatorToPb(&st.ComparisonOperator)
	if err != nil {
		return nil, err
	}
	if comparisonOperatorPb != nil {
		pb.ComparisonOperator = *comparisonOperatorPb
	}
	emptyResultStatePb, err := AlertEvaluationStateToPb(&st.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStatePb != nil {
		pb.EmptyResultState = *emptyResultStatePb
	}
	lastEvaluatedAtPb, err := timestampToPb(st.LastEvaluatedAt)
	if err != nil {
		return nil, err
	}
	if lastEvaluatedAtPb != nil {
		pb.LastEvaluatedAt = *lastEvaluatedAtPb
	}
	notificationPb, err := AlertV2NotificationToPb(st.Notification)
	if err != nil {
		return nil, err
	}
	if notificationPb != nil {
		pb.Notification = notificationPb
	}
	sourcePb, err := AlertV2OperandColumnToPb(st.Source)
	if err != nil {
		return nil, err
	}
	if sourcePb != nil {
		pb.Source = sourcePb
	}
	statePb, err := AlertEvaluationStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	thresholdPb, err := AlertV2OperandToPb(st.Threshold)
	if err != nil {
		return nil, err
	}
	if thresholdPb != nil {
		pb.Threshold = thresholdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertV2EvaluationFromPb(pb *sqlpb.AlertV2EvaluationPb) (*AlertV2Evaluation, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Evaluation{}
	comparisonOperatorField, err := ComparisonOperatorFromPb(&pb.ComparisonOperator)
	if err != nil {
		return nil, err
	}
	if comparisonOperatorField != nil {
		st.ComparisonOperator = *comparisonOperatorField
	}
	emptyResultStateField, err := AlertEvaluationStateFromPb(&pb.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStateField != nil {
		st.EmptyResultState = *emptyResultStateField
	}
	lastEvaluatedAtField, err := timestampFromPb(&pb.LastEvaluatedAt)
	if err != nil {
		return nil, err
	}
	if lastEvaluatedAtField != nil {
		st.LastEvaluatedAt = lastEvaluatedAtField
	}
	notificationField, err := AlertV2NotificationFromPb(pb.Notification)
	if err != nil {
		return nil, err
	}
	if notificationField != nil {
		st.Notification = notificationField
	}
	sourceField, err := AlertV2OperandColumnFromPb(pb.Source)
	if err != nil {
		return nil, err
	}
	if sourceField != nil {
		st.Source = sourceField
	}
	stateField, err := AlertEvaluationStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	thresholdField, err := AlertV2OperandFromPb(pb.Threshold)
	if err != nil {
		return nil, err
	}
	if thresholdField != nil {
		st.Threshold = thresholdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AlertV2Notification struct {
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool ``
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'retrigger_seconds'
	RetriggerSeconds int ``

	// Wire name: 'subscriptions'
	Subscriptions   []AlertV2Subscription ``
	ForceSendFields []string              `tf:"-"`
}

func (s *AlertV2Notification) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2Notification) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertV2NotificationToPb(st *AlertV2Notification) (*sqlpb.AlertV2NotificationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertV2NotificationPb{}
	pb.NotifyOnOk = st.NotifyOnOk
	pb.RetriggerSeconds = st.RetriggerSeconds

	var subscriptionsPb []sqlpb.AlertV2SubscriptionPb
	for _, item := range st.Subscriptions {
		itemPb, err := AlertV2SubscriptionToPb(&item)
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

func AlertV2NotificationFromPb(pb *sqlpb.AlertV2NotificationPb) (*AlertV2Notification, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Notification{}
	st.NotifyOnOk = pb.NotifyOnOk
	st.RetriggerSeconds = pb.RetriggerSeconds

	var subscriptionsField []AlertV2Subscription
	for _, itemPb := range pb.Subscriptions {
		item, err := AlertV2SubscriptionFromPb(&itemPb)
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

type AlertV2Operand struct {

	// Wire name: 'column'
	Column *AlertV2OperandColumn ``

	// Wire name: 'value'
	Value *AlertV2OperandValue ``
}

func AlertV2OperandToPb(st *AlertV2Operand) (*sqlpb.AlertV2OperandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertV2OperandPb{}
	columnPb, err := AlertV2OperandColumnToPb(st.Column)
	if err != nil {
		return nil, err
	}
	if columnPb != nil {
		pb.Column = columnPb
	}
	valuePb, err := AlertV2OperandValueToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	return pb, nil
}

func AlertV2OperandFromPb(pb *sqlpb.AlertV2OperandPb) (*AlertV2Operand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Operand{}
	columnField, err := AlertV2OperandColumnFromPb(pb.Column)
	if err != nil {
		return nil, err
	}
	if columnField != nil {
		st.Column = columnField
	}
	valueField, err := AlertV2OperandValueFromPb(pb.Value)
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
	Aggregation Aggregation ``

	// Wire name: 'display'
	Display string ``

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertV2OperandColumn) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2OperandColumn) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertV2OperandColumnToPb(st *AlertV2OperandColumn) (*sqlpb.AlertV2OperandColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertV2OperandColumnPb{}
	aggregationPb, err := AggregationToPb(&st.Aggregation)
	if err != nil {
		return nil, err
	}
	if aggregationPb != nil {
		pb.Aggregation = *aggregationPb
	}
	pb.Display = st.Display
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertV2OperandColumnFromPb(pb *sqlpb.AlertV2OperandColumnPb) (*AlertV2OperandColumn, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2OperandColumn{}
	aggregationField, err := AggregationFromPb(&pb.Aggregation)
	if err != nil {
		return nil, err
	}
	if aggregationField != nil {
		st.Aggregation = *aggregationField
	}
	st.Display = pb.Display
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AlertV2OperandValue struct {

	// Wire name: 'bool_value'
	BoolValue bool ``

	// Wire name: 'double_value'
	DoubleValue float64 ``

	// Wire name: 'string_value'
	StringValue     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertV2OperandValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2OperandValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertV2OperandValueToPb(st *AlertV2OperandValue) (*sqlpb.AlertV2OperandValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertV2OperandValuePb{}
	pb.BoolValue = st.BoolValue
	pb.DoubleValue = st.DoubleValue
	pb.StringValue = st.StringValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertV2OperandValueFromPb(pb *sqlpb.AlertV2OperandValuePb) (*AlertV2OperandValue, error) {
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

type AlertV2Subscription struct {

	// Wire name: 'destination_id'
	DestinationId string ``

	// Wire name: 'user_email'
	UserEmail       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AlertV2Subscription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertV2Subscription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AlertV2SubscriptionToPb(st *AlertV2Subscription) (*sqlpb.AlertV2SubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.AlertV2SubscriptionPb{}
	pb.DestinationId = st.DestinationId
	pb.UserEmail = st.UserEmail

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AlertV2SubscriptionFromPb(pb *sqlpb.AlertV2SubscriptionPb) (*AlertV2Subscription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Subscription{}
	st.DestinationId = pb.DestinationId
	st.UserEmail = pb.UserEmail

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Describes metadata for a particular chunk, within a result set; this
// structure is used both within a manifest, and when fetching individual chunk
// data or links.
type BaseChunkInfo struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64 ``
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int ``
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64 ``
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset       int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *BaseChunkInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BaseChunkInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func BaseChunkInfoToPb(st *BaseChunkInfo) (*sqlpb.BaseChunkInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.BaseChunkInfoPb{}
	pb.ByteCount = st.ByteCount
	pb.ChunkIndex = st.ChunkIndex
	pb.RowCount = st.RowCount
	pb.RowOffset = st.RowOffset

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func BaseChunkInfoFromPb(pb *sqlpb.BaseChunkInfoPb) (*BaseChunkInfo, error) {
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

type CancelExecutionRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string `tf:"-"`
}

func CancelExecutionRequestToPb(st *CancelExecutionRequest) (*sqlpb.CancelExecutionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CancelExecutionRequestPb{}
	pb.StatementId = st.StatementId

	return pb, nil
}

func CancelExecutionRequestFromPb(pb *sqlpb.CancelExecutionRequestPb) (*CancelExecutionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelExecutionRequest{}
	st.StatementId = pb.StatementId

	return st, nil
}

// Configures the channel name and DBSQL version of the warehouse.
// CHANNEL_NAME_CUSTOM should be chosen only when `dbsql_version` is specified.
type Channel struct {

	// Wire name: 'dbsql_version'
	DbsqlVersion string ``

	// Wire name: 'name'
	Name            ChannelName ``
	ForceSendFields []string    `tf:"-"`
}

func (s *Channel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Channel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ChannelToPb(st *Channel) (*sqlpb.ChannelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ChannelPb{}
	pb.DbsqlVersion = st.DbsqlVersion
	namePb, err := ChannelNameToPb(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ChannelFromPb(pb *sqlpb.ChannelPb) (*Channel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Channel{}
	st.DbsqlVersion = pb.DbsqlVersion
	nameField, err := ChannelNameFromPb(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Details about a Channel.
type ChannelInfo struct {
	// DB SQL Version the Channel is mapped to.
	// Wire name: 'dbsql_version'
	DbsqlVersion string ``
	// Name of the channel
	// Wire name: 'name'
	Name            ChannelName ``
	ForceSendFields []string    `tf:"-"`
}

func (s *ChannelInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ChannelInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ChannelInfoToPb(st *ChannelInfo) (*sqlpb.ChannelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ChannelInfoPb{}
	pb.DbsqlVersion = st.DbsqlVersion
	namePb, err := ChannelNameToPb(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ChannelInfoFromPb(pb *sqlpb.ChannelInfoPb) (*ChannelInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChannelInfo{}
	st.DbsqlVersion = pb.DbsqlVersion
	nameField, err := ChannelNameFromPb(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ChannelNameToPb(st *ChannelName) (*sqlpb.ChannelNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ChannelNamePb(*st)
	return &pb, nil
}

func ChannelNameFromPb(pb *sqlpb.ChannelNamePb) (*ChannelName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ChannelName(*pb)
	return &st, nil
}

type ClientConfig struct {

	// Wire name: 'allow_custom_js_visualizations'
	AllowCustomJsVisualizations bool ``

	// Wire name: 'allow_downloads'
	AllowDownloads bool ``

	// Wire name: 'allow_external_shares'
	AllowExternalShares bool ``

	// Wire name: 'allow_subscriptions'
	AllowSubscriptions bool ``

	// Wire name: 'date_format'
	DateFormat string ``

	// Wire name: 'date_time_format'
	DateTimeFormat string ``

	// Wire name: 'disable_publish'
	DisablePublish bool ``

	// Wire name: 'enable_legacy_autodetect_types'
	EnableLegacyAutodetectTypes bool ``

	// Wire name: 'feature_show_permissions_control'
	FeatureShowPermissionsControl bool ``

	// Wire name: 'hide_plotly_mode_bar'
	HidePlotlyModeBar bool     ``
	ForceSendFields   []string `tf:"-"`
}

func (s *ClientConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClientConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ClientConfigToPb(st *ClientConfig) (*sqlpb.ClientConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ClientConfigPb{}
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

func ClientConfigFromPb(pb *sqlpb.ClientConfigPb) (*ClientConfig, error) {
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

type ColumnInfo struct {
	// The name of the column.
	// Wire name: 'name'
	Name string ``
	// The ordinal position of the column (starting at position 0).
	// Wire name: 'position'
	Position int ``
	// The format of the interval type.
	// Wire name: 'type_interval_type'
	TypeIntervalType string ``
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	// Wire name: 'type_name'
	TypeName ColumnInfoTypeName ``
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	// Wire name: 'type_precision'
	TypePrecision int ``
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	// Wire name: 'type_scale'
	TypeScale int ``
	// The full SQL type specification.
	// Wire name: 'type_text'
	TypeText        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ColumnInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ColumnInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ColumnInfoToPb(st *ColumnInfo) (*sqlpb.ColumnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ColumnInfoPb{}
	pb.Name = st.Name
	pb.Position = st.Position
	pb.TypeIntervalType = st.TypeIntervalType
	typeNamePb, err := ColumnInfoTypeNameToPb(&st.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}
	pb.TypePrecision = st.TypePrecision
	pb.TypeScale = st.TypeScale
	pb.TypeText = st.TypeText

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ColumnInfoFromPb(pb *sqlpb.ColumnInfoPb) (*ColumnInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ColumnInfo{}
	st.Name = pb.Name
	st.Position = pb.Position
	st.TypeIntervalType = pb.TypeIntervalType
	typeNameField, err := ColumnInfoTypeNameFromPb(&pb.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	st.TypePrecision = pb.TypePrecision
	st.TypeScale = pb.TypeScale
	st.TypeText = pb.TypeText

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ColumnInfoTypeNameToPb(st *ColumnInfoTypeName) (*sqlpb.ColumnInfoTypeNamePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ColumnInfoTypeNamePb(*st)
	return &pb, nil
}

func ColumnInfoTypeNameFromPb(pb *sqlpb.ColumnInfoTypeNamePb) (*ColumnInfoTypeName, error) {
	if pb == nil {
		return nil, nil
	}
	st := ColumnInfoTypeName(*pb)
	return &st, nil
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

func ComparisonOperatorToPb(st *ComparisonOperator) (*sqlpb.ComparisonOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ComparisonOperatorPb(*st)
	return &pb, nil
}

func ComparisonOperatorFromPb(pb *sqlpb.ComparisonOperatorPb) (*ComparisonOperator, error) {
	if pb == nil {
		return nil, nil
	}
	st := ComparisonOperator(*pb)
	return &st, nil
}

type CreateAlert struct {
	// Name of the alert.
	// Wire name: 'name'
	Name string ``
	// Alert configuration options.
	// Wire name: 'options'
	Options AlertOptions ``
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string ``
	// Query ID.
	// Wire name: 'query_id'
	QueryId string ``
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm           int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateAlertToPb(st *CreateAlert) (*sqlpb.CreateAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateAlertPb{}
	pb.Name = st.Name
	optionsPb, err := AlertOptionsToPb(&st.Options)
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

func CreateAlertFromPb(pb *sqlpb.CreateAlertPb) (*CreateAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlert{}
	st.Name = pb.Name
	optionsField, err := AlertOptionsFromPb(&pb.Options)
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

type CreateAlertRequest struct {

	// Wire name: 'alert'
	Alert *CreateAlertRequestAlert ``
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool     ``
	ForceSendFields        []string `tf:"-"`
}

func (s *CreateAlertRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAlertRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateAlertRequestToPb(st *CreateAlertRequest) (*sqlpb.CreateAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateAlertRequestPb{}
	alertPb, err := CreateAlertRequestAlertToPb(st.Alert)
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

func CreateAlertRequestFromPb(pb *sqlpb.CreateAlertRequestPb) (*CreateAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertRequest{}
	alertField, err := CreateAlertRequestAlertFromPb(pb.Alert)
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

type CreateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition ``
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string ``
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string ``
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string ``
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool ``
	// The workspace path of the folder containing the alert.
	// Wire name: 'parent_path'
	ParentPath string ``
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string ``
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int      ``
	ForceSendFields    []string `tf:"-"`
}

func (s *CreateAlertRequestAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAlertRequestAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateAlertRequestAlertToPb(st *CreateAlertRequestAlert) (*sqlpb.CreateAlertRequestAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateAlertRequestAlertPb{}
	conditionPb, err := AlertConditionToPb(st.Condition)
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

func CreateAlertRequestAlertFromPb(pb *sqlpb.CreateAlertRequestAlertPb) (*CreateAlertRequestAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertRequestAlert{}
	conditionField, err := AlertConditionFromPb(pb.Condition)
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

type CreateAlertV2Request struct {

	// Wire name: 'alert'
	Alert AlertV2 ``
}

func CreateAlertV2RequestToPb(st *CreateAlertV2Request) (*sqlpb.CreateAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateAlertV2RequestPb{}
	alertPb, err := AlertV2ToPb(&st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = *alertPb
	}

	return pb, nil
}

func CreateAlertV2RequestFromPb(pb *sqlpb.CreateAlertV2RequestPb) (*CreateAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertV2Request{}
	alertField, err := AlertV2FromPb(&pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = *alertField
	}

	return st, nil
}

type CreateQueryRequest struct {
	// If true, automatically resolve query display name conflicts. Otherwise,
	// fail the request if the query's display name conflicts with an existing
	// query's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool ``

	// Wire name: 'query'
	Query           *CreateQueryRequestQuery ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *CreateQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateQueryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateQueryRequestToPb(st *CreateQueryRequest) (*sqlpb.CreateQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateQueryRequestPb{}
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName
	queryPb, err := CreateQueryRequestQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateQueryRequestFromPb(pb *sqlpb.CreateQueryRequestPb) (*CreateQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateQueryRequest{}
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName
	queryField, err := CreateQueryRequestQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool ``
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string ``
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter ``
	// Workspace path of the workspace folder containing the object.
	// Wire name: 'parent_path'
	ParentPath string ``
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string ``
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode ``
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string ``

	// Wire name: 'tags'
	Tags []string ``
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateQueryRequestQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateQueryRequestQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateQueryRequestQueryToPb(st *CreateQueryRequestQuery) (*sqlpb.CreateQueryRequestQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateQueryRequestQueryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit
	pb.Catalog = st.Catalog
	pb.Description = st.Description
	pb.DisplayName = st.DisplayName

	var parametersPb []sqlpb.QueryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := QueryParameterToPb(&item)
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
	runAsModePb, err := RunAsModeToPb(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}
	pb.Schema = st.Schema
	pb.Tags = st.Tags
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateQueryRequestQueryFromPb(pb *sqlpb.CreateQueryRequestQueryPb) (*CreateQueryRequestQuery, error) {
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
		item, err := QueryParameterFromPb(&itemPb)
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
	runAsModeField, err := RunAsModeFromPb(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Add visualization to a query
type CreateQueryVisualizationsLegacyRequest struct {
	// A short description of this visualization. This is not displayed in the
	// UI.
	// Wire name: 'description'
	Description string ``
	// The name of the visualization that appears on dashboards and the query
	// screen.
	// Wire name: 'name'
	Name string ``
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	// Wire name: 'options'
	Options any ``
	// The identifier returned by :method:queries/create
	// Wire name: 'query_id'
	QueryId string ``
	// The type of visualization: chart, table, pivot table, and so on.
	// Wire name: 'type'
	Type            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateQueryVisualizationsLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateQueryVisualizationsLegacyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateQueryVisualizationsLegacyRequestToPb(st *CreateQueryVisualizationsLegacyRequest) (*sqlpb.CreateQueryVisualizationsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateQueryVisualizationsLegacyRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name
	pb.Options = st.Options
	pb.QueryId = st.QueryId
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateQueryVisualizationsLegacyRequestFromPb(pb *sqlpb.CreateQueryVisualizationsLegacyRequestPb) (*CreateQueryVisualizationsLegacyRequest, error) {
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

type CreateVisualizationRequest struct {

	// Wire name: 'visualization'
	Visualization *CreateVisualizationRequestVisualization ``
}

func CreateVisualizationRequestToPb(st *CreateVisualizationRequest) (*sqlpb.CreateVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateVisualizationRequestPb{}
	visualizationPb, err := CreateVisualizationRequestVisualizationToPb(st.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationPb != nil {
		pb.Visualization = visualizationPb
	}

	return pb, nil
}

func CreateVisualizationRequestFromPb(pb *sqlpb.CreateVisualizationRequestPb) (*CreateVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVisualizationRequest{}
	visualizationField, err := CreateVisualizationRequestVisualizationFromPb(pb.Visualization)
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
	DisplayName string ``
	// UUID of the query that the visualization is attached to.
	// Wire name: 'query_id'
	QueryId string ``
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string ``
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string ``
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateVisualizationRequestVisualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateVisualizationRequestVisualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateVisualizationRequestVisualizationToPb(st *CreateVisualizationRequestVisualization) (*sqlpb.CreateVisualizationRequestVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateVisualizationRequestVisualizationPb{}
	pb.DisplayName = st.DisplayName
	pb.QueryId = st.QueryId
	pb.SerializedOptions = st.SerializedOptions
	pb.SerializedQueryPlan = st.SerializedQueryPlan
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateVisualizationRequestVisualizationFromPb(pb *sqlpb.CreateVisualizationRequestVisualizationPb) (*CreateVisualizationRequestVisualization, error) {
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
	AutoStopMins int ``
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel ``
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string ``
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string ``
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool ``
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool ``
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int ``
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
	MinNumClusters int ``
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy ``
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags ``

	// Wire name: 'warehouse_type'
	WarehouseType   CreateWarehouseRequestWarehouseType ``
	ForceSendFields []string                            `tf:"-"`
}

func (s *CreateWarehouseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWarehouseRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateWarehouseRequestToPb(st *CreateWarehouseRequest) (*sqlpb.CreateWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateWarehouseRequestPb{}
	pb.AutoStopMins = st.AutoStopMins
	channelPb, err := ChannelToPb(st.Channel)
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
	spotInstancePolicyPb, err := SpotInstancePolicyToPb(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}
	tagsPb, err := EndpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}
	warehouseTypePb, err := CreateWarehouseRequestWarehouseTypeToPb(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateWarehouseRequestFromPb(pb *sqlpb.CreateWarehouseRequestPb) (*CreateWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWarehouseRequest{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := ChannelFromPb(pb.Channel)
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
	spotInstancePolicyField, err := SpotInstancePolicyFromPb(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	tagsField, err := EndpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := CreateWarehouseRequestWarehouseTypeFromPb(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func CreateWarehouseRequestWarehouseTypeToPb(st *CreateWarehouseRequestWarehouseType) (*sqlpb.CreateWarehouseRequestWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.CreateWarehouseRequestWarehouseTypePb(*st)
	return &pb, nil
}

func CreateWarehouseRequestWarehouseTypeFromPb(pb *sqlpb.CreateWarehouseRequestWarehouseTypePb) (*CreateWarehouseRequestWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := CreateWarehouseRequestWarehouseType(*pb)
	return &st, nil
}

type CreateWarehouseResponse struct {
	// Id for the SQL warehouse. This value is unique across all SQL warehouses.
	// Wire name: 'id'
	Id              string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateWarehouseResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWarehouseResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateWarehouseResponseToPb(st *CreateWarehouseResponse) (*sqlpb.CreateWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateWarehouseResponsePb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateWarehouseResponseFromPb(pb *sqlpb.CreateWarehouseResponsePb) (*CreateWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWarehouseResponse{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	// Wire name: 'dashboard_id'
	DashboardId string ``

	// Wire name: 'options'
	Options WidgetOptions ``
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	// Wire name: 'text'
	Text string ``
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	// Wire name: 'visualization_id'
	VisualizationId string ``
	// Width of a widget
	// Wire name: 'width'
	Width           int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *CreateWidget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateWidget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateWidgetToPb(st *CreateWidget) (*sqlpb.CreateWidgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CreateWidgetPb{}
	pb.DashboardId = st.DashboardId
	optionsPb, err := WidgetOptionsToPb(&st.Options)
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

func CreateWidgetFromPb(pb *sqlpb.CreateWidgetPb) (*CreateWidget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWidget{}
	st.DashboardId = pb.DashboardId
	optionsField, err := WidgetOptionsFromPb(&pb.Options)
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

type CronSchedule struct {
	// Indicate whether this schedule is paused or not.
	// Wire name: 'pause_status'
	PauseStatus SchedulePauseStatus ``
	// A cron expression using quartz syntax that specifies the schedule for
	// this pipeline. Should use the quartz format described here:
	// http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
	// Wire name: 'quartz_cron_schedule'
	QuartzCronSchedule string ``
	// A Java timezone id. The schedule will be resolved using this timezone.
	// This will be combined with the quartz_cron_schedule to determine the
	// schedule. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details.
	// Wire name: 'timezone_id'
	TimezoneId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *CronSchedule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CronSchedule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CronScheduleToPb(st *CronSchedule) (*sqlpb.CronSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.CronSchedulePb{}
	pauseStatusPb, err := SchedulePauseStatusToPb(&st.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusPb != nil {
		pb.PauseStatus = *pauseStatusPb
	}
	pb.QuartzCronSchedule = st.QuartzCronSchedule
	pb.TimezoneId = st.TimezoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CronScheduleFromPb(pb *sqlpb.CronSchedulePb) (*CronSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronSchedule{}
	pauseStatusField, err := SchedulePauseStatusFromPb(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	st.QuartzCronSchedule = pb.QuartzCronSchedule
	st.TimezoneId = pb.TimezoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A JSON representing a dashboard containing widgets of visualizations and text
// boxes.
type Dashboard struct {
	// Whether the authenticated user can edit the query definition.
	// Wire name: 'can_edit'
	CanEdit bool ``
	// Timestamp when this dashboard was created.
	// Wire name: 'created_at'
	CreatedAt string ``
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	// Wire name: 'dashboard_filters_enabled'
	DashboardFiltersEnabled bool ``
	// The ID for this dashboard.
	// Wire name: 'id'
	Id string ``
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool ``
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	// Wire name: 'is_draft'
	IsDraft bool ``
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	// Wire name: 'is_favorite'
	IsFavorite bool ``
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'options'
	Options *DashboardOptions ``
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string ``
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_tier'
	PermissionTier PermissionLevel ``
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	// Wire name: 'slug'
	Slug string ``

	// Wire name: 'tags'
	Tags []string ``
	// Timestamp when this dashboard was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string ``

	// Wire name: 'user'
	User *User ``
	// The ID of the user who owns the dashboard.
	// Wire name: 'user_id'
	UserId int ``

	// Wire name: 'widgets'
	Widgets         []Widget ``
	ForceSendFields []string `tf:"-"`
}

func (s *Dashboard) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Dashboard) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DashboardToPb(st *Dashboard) (*sqlpb.DashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DashboardPb{}
	pb.CanEdit = st.CanEdit
	pb.CreatedAt = st.CreatedAt
	pb.DashboardFiltersEnabled = st.DashboardFiltersEnabled
	pb.Id = st.Id
	pb.IsArchived = st.IsArchived
	pb.IsDraft = st.IsDraft
	pb.IsFavorite = st.IsFavorite
	pb.Name = st.Name
	optionsPb, err := DashboardOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}
	pb.Parent = st.Parent
	permissionTierPb, err := PermissionLevelToPb(&st.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierPb != nil {
		pb.PermissionTier = *permissionTierPb
	}
	pb.Slug = st.Slug
	pb.Tags = st.Tags
	pb.UpdatedAt = st.UpdatedAt
	userPb, err := UserToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}
	pb.UserId = st.UserId

	var widgetsPb []sqlpb.WidgetPb
	for _, item := range st.Widgets {
		itemPb, err := WidgetToPb(&item)
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

func DashboardFromPb(pb *sqlpb.DashboardPb) (*Dashboard, error) {
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
	optionsField, err := DashboardOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	st.Parent = pb.Parent
	permissionTierField, err := PermissionLevelFromPb(&pb.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierField != nil {
		st.PermissionTier = *permissionTierField
	}
	st.Slug = pb.Slug
	st.Tags = pb.Tags
	st.UpdatedAt = pb.UpdatedAt
	userField, err := UserFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}
	st.UserId = pb.UserId

	var widgetsField []Widget
	for _, itemPb := range pb.Widgets {
		item, err := WidgetFromPb(&itemPb)
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

type DashboardEditContent struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	// Wire name: 'name'
	Name string ``
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole ``

	// Wire name: 'tags'
	Tags            []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *DashboardEditContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardEditContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DashboardEditContentToPb(st *DashboardEditContent) (*sqlpb.DashboardEditContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DashboardEditContentPb{}
	pb.DashboardId = st.DashboardId
	pb.Name = st.Name
	runAsRolePb, err := RunAsRoleToPb(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DashboardEditContentFromPb(pb *sqlpb.DashboardEditContentPb) (*DashboardEditContent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardEditContent{}
	st.DashboardId = pb.DashboardId
	st.Name = pb.Name
	runAsRoleField, err := RunAsRoleFromPb(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DashboardOptions struct {
	// The timestamp when this dashboard was moved to trash. Only present when
	// the `is_archived` property is `true`. Trashed items are deleted after
	// thirty days.
	// Wire name: 'moved_to_trash_at'
	MovedToTrashAt  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DashboardOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DashboardOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DashboardOptionsToPb(st *DashboardOptions) (*sqlpb.DashboardOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DashboardOptionsPb{}
	pb.MovedToTrashAt = st.MovedToTrashAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DashboardOptionsFromPb(pb *sqlpb.DashboardOptionsPb) (*DashboardOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardOptions{}
	st.MovedToTrashAt = pb.MovedToTrashAt

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// A JSON object representing a DBSQL data source / SQL warehouse.
type DataSource struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'id'
	Id string ``
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	// Wire name: 'name'
	Name string ``
	// Reserved for internal use.
	// Wire name: 'pause_reason'
	PauseReason string ``
	// Reserved for internal use.
	// Wire name: 'paused'
	Paused int ``
	// Reserved for internal use.
	// Wire name: 'supports_auto_limit'
	SupportsAutoLimit bool ``
	// Reserved for internal use.
	// Wire name: 'syntax'
	Syntax string ``
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	// Wire name: 'type'
	Type string ``
	// Reserved for internal use.
	// Wire name: 'view_only'
	ViewOnly bool ``
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DataSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DataSourceToPb(st *DataSource) (*sqlpb.DataSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DataSourcePb{}
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

func DataSourceFromPb(pb *sqlpb.DataSourcePb) (*DataSource, error) {
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

func DatePrecisionToPb(st *DatePrecision) (*sqlpb.DatePrecisionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.DatePrecisionPb(*st)
	return &pb, nil
}

func DatePrecisionFromPb(pb *sqlpb.DatePrecisionPb) (*DatePrecision, error) {
	if pb == nil {
		return nil, nil
	}
	st := DatePrecision(*pb)
	return &st, nil
}

type DateRange struct {

	// Wire name: 'end'
	End string ``

	// Wire name: 'start'
	Start string ``
}

func DateRangeToPb(st *DateRange) (*sqlpb.DateRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DateRangePb{}
	pb.End = st.End
	pb.Start = st.Start

	return pb, nil
}

func DateRangeFromPb(pb *sqlpb.DateRangePb) (*DateRange, error) {
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
	DateRangeValue *DateRange ``
	// Dynamic date-time range value based on current date-time.
	// Wire name: 'dynamic_date_range_value'
	DynamicDateRangeValue DateRangeValueDynamicDateRange ``
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	// Wire name: 'precision'
	Precision DatePrecision ``

	// Wire name: 'start_day_of_week'
	StartDayOfWeek  int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *DateRangeValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DateRangeValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DateRangeValueToPb(st *DateRangeValue) (*sqlpb.DateRangeValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DateRangeValuePb{}
	dateRangeValuePb, err := DateRangeToPb(st.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValuePb != nil {
		pb.DateRangeValue = dateRangeValuePb
	}
	dynamicDateRangeValuePb, err := DateRangeValueDynamicDateRangeToPb(&st.DynamicDateRangeValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateRangeValuePb != nil {
		pb.DynamicDateRangeValue = *dynamicDateRangeValuePb
	}
	precisionPb, err := DatePrecisionToPb(&st.Precision)
	if err != nil {
		return nil, err
	}
	if precisionPb != nil {
		pb.Precision = *precisionPb
	}
	pb.StartDayOfWeek = st.StartDayOfWeek

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DateRangeValueFromPb(pb *sqlpb.DateRangeValuePb) (*DateRangeValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DateRangeValue{}
	dateRangeValueField, err := DateRangeFromPb(pb.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValueField != nil {
		st.DateRangeValue = dateRangeValueField
	}
	dynamicDateRangeValueField, err := DateRangeValueDynamicDateRangeFromPb(&pb.DynamicDateRangeValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateRangeValueField != nil {
		st.DynamicDateRangeValue = *dynamicDateRangeValueField
	}
	precisionField, err := DatePrecisionFromPb(&pb.Precision)
	if err != nil {
		return nil, err
	}
	if precisionField != nil {
		st.Precision = *precisionField
	}
	st.StartDayOfWeek = pb.StartDayOfWeek

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func DateRangeValueDynamicDateRangeToPb(st *DateRangeValueDynamicDateRange) (*sqlpb.DateRangeValueDynamicDateRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.DateRangeValueDynamicDateRangePb(*st)
	return &pb, nil
}

func DateRangeValueDynamicDateRangeFromPb(pb *sqlpb.DateRangeValueDynamicDateRangePb) (*DateRangeValueDynamicDateRange, error) {
	if pb == nil {
		return nil, nil
	}
	st := DateRangeValueDynamicDateRange(*pb)
	return &st, nil
}

type DateValue struct {
	// Manually specified date-time value.
	// Wire name: 'date_value'
	DateValue string ``
	// Dynamic date-time value based on current date-time.
	// Wire name: 'dynamic_date_value'
	DynamicDateValue DateValueDynamicDate ``
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	// Wire name: 'precision'
	Precision       DatePrecision ``
	ForceSendFields []string      `tf:"-"`
}

func (s *DateValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DateValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DateValueToPb(st *DateValue) (*sqlpb.DateValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DateValuePb{}
	pb.DateValue = st.DateValue
	dynamicDateValuePb, err := DateValueDynamicDateToPb(&st.DynamicDateValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateValuePb != nil {
		pb.DynamicDateValue = *dynamicDateValuePb
	}
	precisionPb, err := DatePrecisionToPb(&st.Precision)
	if err != nil {
		return nil, err
	}
	if precisionPb != nil {
		pb.Precision = *precisionPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DateValueFromPb(pb *sqlpb.DateValuePb) (*DateValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DateValue{}
	st.DateValue = pb.DateValue
	dynamicDateValueField, err := DateValueDynamicDateFromPb(&pb.DynamicDateValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateValueField != nil {
		st.DynamicDateValue = *dynamicDateValueField
	}
	precisionField, err := DatePrecisionFromPb(&pb.Precision)
	if err != nil {
		return nil, err
	}
	if precisionField != nil {
		st.Precision = *precisionField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func DateValueDynamicDateToPb(st *DateValueDynamicDate) (*sqlpb.DateValueDynamicDatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.DateValueDynamicDatePb(*st)
	return &pb, nil
}

func DateValueDynamicDateFromPb(pb *sqlpb.DateValueDynamicDatePb) (*DateValueDynamicDate, error) {
	if pb == nil {
		return nil, nil
	}
	st := DateValueDynamicDate(*pb)
	return &st, nil
}

type DeleteAlertsLegacyRequest struct {

	// Wire name: 'alert_id'
	AlertId string `tf:"-"`
}

func DeleteAlertsLegacyRequestToPb(st *DeleteAlertsLegacyRequest) (*sqlpb.DeleteAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DeleteAlertsLegacyRequestPb{}
	pb.AlertId = st.AlertId

	return pb, nil
}

func DeleteAlertsLegacyRequestFromPb(pb *sqlpb.DeleteAlertsLegacyRequestPb) (*DeleteAlertsLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAlertsLegacyRequest{}
	st.AlertId = pb.AlertId

	return st, nil
}

type DeleteDashboardRequest struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func DeleteDashboardRequestToPb(st *DeleteDashboardRequest) (*sqlpb.DeleteDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DeleteDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func DeleteDashboardRequestFromPb(pb *sqlpb.DeleteDashboardRequestPb) (*DeleteDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteDashboardWidgetRequestToPb(st *DeleteDashboardWidgetRequest) (*sqlpb.DeleteDashboardWidgetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DeleteDashboardWidgetRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteDashboardWidgetRequestFromPb(pb *sqlpb.DeleteDashboardWidgetRequestPb) (*DeleteDashboardWidgetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardWidgetRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteQueriesLegacyRequest struct {

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
}

func DeleteQueriesLegacyRequestToPb(st *DeleteQueriesLegacyRequest) (*sqlpb.DeleteQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DeleteQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
}

func DeleteQueriesLegacyRequestFromPb(pb *sqlpb.DeleteQueriesLegacyRequestPb) (*DeleteQueriesLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQueriesLegacyRequest{}
	st.QueryId = pb.QueryId

	return st, nil
}

type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvisualizations/create
	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteQueryVisualizationsLegacyRequestToPb(st *DeleteQueryVisualizationsLegacyRequest) (*sqlpb.DeleteQueryVisualizationsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DeleteQueryVisualizationsLegacyRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteQueryVisualizationsLegacyRequestFromPb(pb *sqlpb.DeleteQueryVisualizationsLegacyRequestPb) (*DeleteQueryVisualizationsLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQueryVisualizationsLegacyRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteVisualizationRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteVisualizationRequestToPb(st *DeleteVisualizationRequest) (*sqlpb.DeleteVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DeleteVisualizationRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteVisualizationRequestFromPb(pb *sqlpb.DeleteVisualizationRequestPb) (*DeleteVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteVisualizationRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func DeleteWarehouseRequestToPb(st *DeleteWarehouseRequest) (*sqlpb.DeleteWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.DeleteWarehouseRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteWarehouseRequestFromPb(pb *sqlpb.DeleteWarehouseRequestPb) (*DeleteWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWarehouseRequest{}
	st.Id = pb.Id

	return st, nil
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

func DispositionToPb(st *Disposition) (*sqlpb.DispositionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.DispositionPb(*st)
	return &pb, nil
}

func DispositionFromPb(pb *sqlpb.DispositionPb) (*Disposition, error) {
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
	Name string ``
	// Alert configuration options.
	// Wire name: 'options'
	Options AlertOptions ``
	// Query ID.
	// Wire name: 'query_id'
	QueryId string ``
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm           int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *EditAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EditAlertToPb(st *EditAlert) (*sqlpb.EditAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EditAlertPb{}
	pb.AlertId = st.AlertId
	pb.Name = st.Name
	optionsPb, err := AlertOptionsToPb(&st.Options)
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

func EditAlertFromPb(pb *sqlpb.EditAlertPb) (*EditAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditAlert{}
	st.AlertId = pb.AlertId
	st.Name = pb.Name
	optionsField, err := AlertOptionsFromPb(&pb.Options)
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

type EditWarehouseRequest struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int ``
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel ``
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string ``
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string ``
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool ``
	// Configures whether the warehouse should use serverless compute.
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool ``
	// Required. Id of the warehouse to configure.
	// Wire name: 'id'
	Id string `tf:"-"`
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int ``
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
	MinNumClusters int ``
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy ``
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags ``

	// Wire name: 'warehouse_type'
	WarehouseType   EditWarehouseRequestWarehouseType ``
	ForceSendFields []string                          `tf:"-"`
}

func (s *EditWarehouseRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EditWarehouseRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EditWarehouseRequestToPb(st *EditWarehouseRequest) (*sqlpb.EditWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EditWarehouseRequestPb{}
	pb.AutoStopMins = st.AutoStopMins
	channelPb, err := ChannelToPb(st.Channel)
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
	spotInstancePolicyPb, err := SpotInstancePolicyToPb(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}
	tagsPb, err := EndpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}
	warehouseTypePb, err := EditWarehouseRequestWarehouseTypeToPb(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EditWarehouseRequestFromPb(pb *sqlpb.EditWarehouseRequestPb) (*EditWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditWarehouseRequest{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := ChannelFromPb(pb.Channel)
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
	spotInstancePolicyField, err := SpotInstancePolicyFromPb(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	tagsField, err := EndpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := EditWarehouseRequestWarehouseTypeFromPb(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func EditWarehouseRequestWarehouseTypeToPb(st *EditWarehouseRequestWarehouseType) (*sqlpb.EditWarehouseRequestWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.EditWarehouseRequestWarehouseTypePb(*st)
	return &pb, nil
}

func EditWarehouseRequestWarehouseTypeFromPb(pb *sqlpb.EditWarehouseRequestWarehouseTypePb) (*EditWarehouseRequestWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EditWarehouseRequestWarehouseType(*pb)
	return &st, nil
}

type EndpointConfPair struct {

	// Wire name: 'key'
	Key string ``

	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *EndpointConfPair) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointConfPair) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointConfPairToPb(st *EndpointConfPair) (*sqlpb.EndpointConfPairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EndpointConfPairPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointConfPairFromPb(pb *sqlpb.EndpointConfPairPb) (*EndpointConfPair, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointConfPair{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	// Wire name: 'details'
	Details string ``
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	// Wire name: 'failure_reason'
	FailureReason *TerminationReason ``
	// Deprecated. split into summary and details for security
	// Wire name: 'message'
	Message string ``

	// Wire name: 'status'
	Status Status ``
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	// Wire name: 'summary'
	Summary         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *EndpointHealth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointHealth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointHealthToPb(st *EndpointHealth) (*sqlpb.EndpointHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EndpointHealthPb{}
	pb.Details = st.Details
	failureReasonPb, err := TerminationReasonToPb(st.FailureReason)
	if err != nil {
		return nil, err
	}
	if failureReasonPb != nil {
		pb.FailureReason = failureReasonPb
	}
	pb.Message = st.Message
	statusPb, err := StatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.Summary = st.Summary

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointHealthFromPb(pb *sqlpb.EndpointHealthPb) (*EndpointHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointHealth{}
	st.Details = pb.Details
	failureReasonField, err := TerminationReasonFromPb(pb.FailureReason)
	if err != nil {
		return nil, err
	}
	if failureReasonField != nil {
		st.FailureReason = failureReasonField
	}
	st.Message = pb.Message
	statusField, err := StatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.Summary = pb.Summary

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EndpointInfo struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	// Wire name: 'auto_stop_mins'
	AutoStopMins int ``
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel ``
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string ``
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string ``
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool ``
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool ``
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	// Wire name: 'health'
	Health *EndpointHealth ``
	// unique identifier for warehouse
	// Wire name: 'id'
	Id string ``
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// the jdbc connection string for this warehouse
	// Wire name: 'jdbc_url'
	JdbcUrl string ``
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int ``
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
	MinNumClusters int ``
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string ``
	// Deprecated. current number of active sessions for the warehouse
	// Wire name: 'num_active_sessions'
	NumActiveSessions int64 ``
	// current number of clusters running for the service
	// Wire name: 'num_clusters'
	NumClusters int ``
	// ODBC parameters for the SQL warehouse
	// Wire name: 'odbc_params'
	OdbcParams *OdbcParams ``

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy ``

	// Wire name: 'state'
	State State ``
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags ``
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	// Wire name: 'warehouse_type'
	WarehouseType   EndpointInfoWarehouseType ``
	ForceSendFields []string                  `tf:"-"`
}

func (s *EndpointInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointInfoToPb(st *EndpointInfo) (*sqlpb.EndpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EndpointInfoPb{}
	pb.AutoStopMins = st.AutoStopMins
	channelPb, err := ChannelToPb(st.Channel)
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
	healthPb, err := EndpointHealthToPb(st.Health)
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
	odbcParamsPb, err := OdbcParamsToPb(st.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsPb != nil {
		pb.OdbcParams = odbcParamsPb
	}
	spotInstancePolicyPb, err := SpotInstancePolicyToPb(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}
	statePb, err := StateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	tagsPb, err := EndpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}
	warehouseTypePb, err := EndpointInfoWarehouseTypeToPb(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointInfoFromPb(pb *sqlpb.EndpointInfoPb) (*EndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointInfo{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := ChannelFromPb(pb.Channel)
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
	healthField, err := EndpointHealthFromPb(pb.Health)
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
	odbcParamsField, err := OdbcParamsFromPb(pb.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsField != nil {
		st.OdbcParams = odbcParamsField
	}
	spotInstancePolicyField, err := SpotInstancePolicyFromPb(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	stateField, err := StateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	tagsField, err := EndpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := EndpointInfoWarehouseTypeFromPb(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func EndpointInfoWarehouseTypeToPb(st *EndpointInfoWarehouseType) (*sqlpb.EndpointInfoWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.EndpointInfoWarehouseTypePb(*st)
	return &pb, nil
}

func EndpointInfoWarehouseTypeFromPb(pb *sqlpb.EndpointInfoWarehouseTypePb) (*EndpointInfoWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointInfoWarehouseType(*pb)
	return &st, nil
}

type EndpointTagPair struct {

	// Wire name: 'key'
	Key string ``

	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *EndpointTagPair) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointTagPair) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointTagPairToPb(st *EndpointTagPair) (*sqlpb.EndpointTagPairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EndpointTagPairPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointTagPairFromPb(pb *sqlpb.EndpointTagPairPb) (*EndpointTagPair, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTagPair{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EndpointTags struct {

	// Wire name: 'custom_tags'
	CustomTags []EndpointTagPair ``
}

func EndpointTagsToPb(st *EndpointTags) (*sqlpb.EndpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EndpointTagsPb{}

	var customTagsPb []sqlpb.EndpointTagPairPb
	for _, item := range st.CustomTags {
		itemPb, err := EndpointTagPairToPb(&item)
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

func EndpointTagsFromPb(pb *sqlpb.EndpointTagsPb) (*EndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTags{}

	var customTagsField []EndpointTagPair
	for _, itemPb := range pb.CustomTags {
		item, err := EndpointTagPairFromPb(&itemPb)
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
	EnumOptions string ``
	// If specified, allows multiple values to be selected for this parameter.
	// Wire name: 'multi_values_options'
	MultiValuesOptions *MultiValuesOptions ``
	// List of selected query parameter values.
	// Wire name: 'values'
	Values          []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *EnumValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnumValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EnumValueToPb(st *EnumValue) (*sqlpb.EnumValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.EnumValuePb{}
	pb.EnumOptions = st.EnumOptions
	multiValuesOptionsPb, err := MultiValuesOptionsToPb(st.MultiValuesOptions)
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

func EnumValueFromPb(pb *sqlpb.EnumValuePb) (*EnumValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnumValue{}
	st.EnumOptions = pb.EnumOptions
	multiValuesOptionsField, err := MultiValuesOptionsFromPb(pb.MultiValuesOptions)
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

type ExecuteStatementRequest struct {
	// Applies the given byte limit to the statement's result size. Byte counts
	// are based on internal data representations and might not match the final
	// size in the requested `format`. If the result was truncated due to the
	// byte limit, then `truncated` in the response is set to `true`. When using
	// `EXTERNAL_LINKS` disposition, a default `byte_limit` of 100 GiB is
	// applied if `byte_limit` is not explcitly set.
	// Wire name: 'byte_limit'
	ByteLimit int64 ``
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	// Wire name: 'catalog'
	Catalog string ``

	// Wire name: 'disposition'
	Disposition Disposition ``
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
	Format Format ``
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
	// Wire name: 'on_wait_timeout'
	OnWaitTimeout ExecuteStatementRequestOnWaitTimeout ``
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
	Parameters []StatementParameterListItem ``
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	// Wire name: 'row_limit'
	RowLimit int64 ``
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	// Wire name: 'schema'
	Schema string ``
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`. The maximum query text size is 16 MiB.
	// Wire name: 'statement'
	Statement string ``
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
	WaitTimeout string ``
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExecuteStatementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExecuteStatementRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExecuteStatementRequestToPb(st *ExecuteStatementRequest) (*sqlpb.ExecuteStatementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ExecuteStatementRequestPb{}
	pb.ByteLimit = st.ByteLimit
	pb.Catalog = st.Catalog
	dispositionPb, err := DispositionToPb(&st.Disposition)
	if err != nil {
		return nil, err
	}
	if dispositionPb != nil {
		pb.Disposition = *dispositionPb
	}
	formatPb, err := FormatToPb(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}
	onWaitTimeoutPb, err := ExecuteStatementRequestOnWaitTimeoutToPb(&st.OnWaitTimeout)
	if err != nil {
		return nil, err
	}
	if onWaitTimeoutPb != nil {
		pb.OnWaitTimeout = *onWaitTimeoutPb
	}

	var parametersPb []sqlpb.StatementParameterListItemPb
	for _, item := range st.Parameters {
		itemPb, err := StatementParameterListItemToPb(&item)
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

func ExecuteStatementRequestFromPb(pb *sqlpb.ExecuteStatementRequestPb) (*ExecuteStatementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExecuteStatementRequest{}
	st.ByteLimit = pb.ByteLimit
	st.Catalog = pb.Catalog
	dispositionField, err := DispositionFromPb(&pb.Disposition)
	if err != nil {
		return nil, err
	}
	if dispositionField != nil {
		st.Disposition = *dispositionField
	}
	formatField, err := FormatFromPb(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	onWaitTimeoutField, err := ExecuteStatementRequestOnWaitTimeoutFromPb(&pb.OnWaitTimeout)
	if err != nil {
		return nil, err
	}
	if onWaitTimeoutField != nil {
		st.OnWaitTimeout = *onWaitTimeoutField
	}

	var parametersField []StatementParameterListItem
	for _, itemPb := range pb.Parameters {
		item, err := StatementParameterListItemFromPb(&itemPb)
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

func ExecuteStatementRequestOnWaitTimeoutToPb(st *ExecuteStatementRequestOnWaitTimeout) (*sqlpb.ExecuteStatementRequestOnWaitTimeoutPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ExecuteStatementRequestOnWaitTimeoutPb(*st)
	return &pb, nil
}

func ExecuteStatementRequestOnWaitTimeoutFromPb(pb *sqlpb.ExecuteStatementRequestOnWaitTimeoutPb) (*ExecuteStatementRequestOnWaitTimeout, error) {
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
	ByteCount int64 ``
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int ``
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	// Wire name: 'expiration'
	Expiration string ``

	// Wire name: 'external_link'
	ExternalLink string ``
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	// Wire name: 'http_headers'
	HttpHeaders map[string]string ``
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	// Wire name: 'next_chunk_index'
	NextChunkIndex int ``
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	// Wire name: 'next_chunk_internal_link'
	NextChunkInternalLink string ``
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64 ``
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset       int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalLink) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalLink) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalLinkToPb(st *ExternalLink) (*sqlpb.ExternalLinkPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ExternalLinkPb{}
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

func ExternalLinkFromPb(pb *sqlpb.ExternalLinkPb) (*ExternalLink, error) {
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

type ExternalQuerySource struct {
	// The canonical identifier for this SQL alert
	// Wire name: 'alert_id'
	AlertId string ``
	// The canonical identifier for this Lakeview dashboard
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// The canonical identifier for this Genie space
	// Wire name: 'genie_space_id'
	GenieSpaceId string ``

	// Wire name: 'job_info'
	JobInfo *ExternalQuerySourceJobInfo ``
	// The canonical identifier for this legacy dashboard
	// Wire name: 'legacy_dashboard_id'
	LegacyDashboardId string ``
	// The canonical identifier for this notebook
	// Wire name: 'notebook_id'
	NotebookId string ``
	// The canonical identifier for this SQL query
	// Wire name: 'sql_query_id'
	SqlQueryId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalQuerySource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalQuerySource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalQuerySourceToPb(st *ExternalQuerySource) (*sqlpb.ExternalQuerySourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ExternalQuerySourcePb{}
	pb.AlertId = st.AlertId
	pb.DashboardId = st.DashboardId
	pb.GenieSpaceId = st.GenieSpaceId
	jobInfoPb, err := ExternalQuerySourceJobInfoToPb(st.JobInfo)
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

func ExternalQuerySourceFromPb(pb *sqlpb.ExternalQuerySourcePb) (*ExternalQuerySource, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalQuerySource{}
	st.AlertId = pb.AlertId
	st.DashboardId = pb.DashboardId
	st.GenieSpaceId = pb.GenieSpaceId
	jobInfoField, err := ExternalQuerySourceJobInfoFromPb(pb.JobInfo)
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

type ExternalQuerySourceJobInfo struct {
	// The canonical identifier for this job.
	// Wire name: 'job_id'
	JobId string ``
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	// Wire name: 'job_run_id'
	JobRunId string ``
	// The canonical identifier of the task run.
	// Wire name: 'job_task_run_id'
	JobTaskRunId    string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalQuerySourceJobInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalQuerySourceJobInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalQuerySourceJobInfoToPb(st *ExternalQuerySourceJobInfo) (*sqlpb.ExternalQuerySourceJobInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ExternalQuerySourceJobInfoPb{}
	pb.JobId = st.JobId
	pb.JobRunId = st.JobRunId
	pb.JobTaskRunId = st.JobTaskRunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalQuerySourceJobInfoFromPb(pb *sqlpb.ExternalQuerySourceJobInfoPb) (*ExternalQuerySourceJobInfo, error) {
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

func FormatToPb(st *Format) (*sqlpb.FormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.FormatPb(*st)
	return &pb, nil
}

func FormatFromPb(pb *sqlpb.FormatPb) (*Format, error) {
	if pb == nil {
		return nil, nil
	}
	st := Format(*pb)
	return &st, nil
}

type GetAlertRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetAlertRequestToPb(st *GetAlertRequest) (*sqlpb.GetAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetAlertRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetAlertRequestFromPb(pb *sqlpb.GetAlertRequestPb) (*GetAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAlertRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetAlertV2Request struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetAlertV2RequestToPb(st *GetAlertV2Request) (*sqlpb.GetAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetAlertV2RequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetAlertV2RequestFromPb(pb *sqlpb.GetAlertV2RequestPb) (*GetAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAlertV2Request{}
	st.Id = pb.Id

	return st, nil
}

type GetAlertsLegacyRequest struct {

	// Wire name: 'alert_id'
	AlertId string `tf:"-"`
}

func GetAlertsLegacyRequestToPb(st *GetAlertsLegacyRequest) (*sqlpb.GetAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetAlertsLegacyRequestPb{}
	pb.AlertId = st.AlertId

	return pb, nil
}

func GetAlertsLegacyRequestFromPb(pb *sqlpb.GetAlertsLegacyRequestPb) (*GetAlertsLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAlertsLegacyRequest{}
	st.AlertId = pb.AlertId

	return st, nil
}

type GetDashboardRequest struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func GetDashboardRequestToPb(st *GetDashboardRequest) (*sqlpb.GetDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func GetDashboardRequestFromPb(pb *sqlpb.GetDashboardRequestPb) (*GetDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	// Wire name: 'objectId'
	ObjectId string `tf:"-"`
	// The type of object permissions to check.
	// Wire name: 'objectType'
	ObjectType ObjectTypePlural `tf:"-"`
}

func GetDbsqlPermissionRequestToPb(st *GetDbsqlPermissionRequest) (*sqlpb.GetDbsqlPermissionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetDbsqlPermissionRequestPb{}
	pb.ObjectId = st.ObjectId
	objectTypePb, err := ObjectTypePluralToPb(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	return pb, nil
}

func GetDbsqlPermissionRequestFromPb(pb *sqlpb.GetDbsqlPermissionRequestPb) (*GetDbsqlPermissionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDbsqlPermissionRequest{}
	st.ObjectId = pb.ObjectId
	objectTypeField, err := ObjectTypePluralFromPb(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	return st, nil
}

type GetQueriesLegacyRequest struct {

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
}

func GetQueriesLegacyRequestToPb(st *GetQueriesLegacyRequest) (*sqlpb.GetQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
}

func GetQueriesLegacyRequestFromPb(pb *sqlpb.GetQueriesLegacyRequestPb) (*GetQueriesLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQueriesLegacyRequest{}
	st.QueryId = pb.QueryId

	return st, nil
}

type GetQueryRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetQueryRequestToPb(st *GetQueryRequest) (*sqlpb.GetQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetQueryRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetQueryRequestFromPb(pb *sqlpb.GetQueryRequestPb) (*GetQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQueryRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetResponse struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl ``
	// An object's type and UUID, separated by a forward slash (/) character.
	// Wire name: 'object_id'
	ObjectId string ``
	// A singular noun object type.
	// Wire name: 'object_type'
	ObjectType      ObjectType ``
	ForceSendFields []string   `tf:"-"`
}

func (s *GetResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetResponseToPb(st *GetResponse) (*sqlpb.GetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetResponsePb{}

	var accessControlListPb []sqlpb.AccessControlPb
	for _, item := range st.AccessControlList {
		itemPb, err := AccessControlToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	objectTypePb, err := ObjectTypeToPb(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetResponseFromPb(pb *sqlpb.GetResponsePb) (*GetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetResponse{}

	var accessControlListField []AccessControl
	for _, itemPb := range pb.AccessControlList {
		item, err := AccessControlFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	objectTypeField, err := ObjectTypeFromPb(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type GetStatementRequest struct {
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string `tf:"-"`
}

func GetStatementRequestToPb(st *GetStatementRequest) (*sqlpb.GetStatementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetStatementRequestPb{}
	pb.StatementId = st.StatementId

	return pb, nil
}

func GetStatementRequestFromPb(pb *sqlpb.GetStatementRequestPb) (*GetStatementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatementRequest{}
	st.StatementId = pb.StatementId

	return st, nil
}

type GetStatementResultChunkNRequest struct {

	// Wire name: 'chunk_index'
	ChunkIndex int `tf:"-"`
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string `tf:"-"`
}

func GetStatementResultChunkNRequestToPb(st *GetStatementResultChunkNRequest) (*sqlpb.GetStatementResultChunkNRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetStatementResultChunkNRequestPb{}
	pb.ChunkIndex = st.ChunkIndex
	pb.StatementId = st.StatementId

	return pb, nil
}

func GetStatementResultChunkNRequestFromPb(pb *sqlpb.GetStatementResultChunkNRequestPb) (*GetStatementResultChunkNRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatementResultChunkNRequest{}
	st.ChunkIndex = pb.ChunkIndex
	st.StatementId = pb.StatementId

	return st, nil
}

type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	// Wire name: 'warehouse_id'
	WarehouseId string `tf:"-"`
}

func GetWarehousePermissionLevelsRequestToPb(st *GetWarehousePermissionLevelsRequest) (*sqlpb.GetWarehousePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetWarehousePermissionLevelsRequestPb{}
	pb.WarehouseId = st.WarehouseId

	return pb, nil
}

func GetWarehousePermissionLevelsRequestFromPb(pb *sqlpb.GetWarehousePermissionLevelsRequestPb) (*GetWarehousePermissionLevelsRequest, error) {
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
	PermissionLevels []WarehousePermissionsDescription ``
}

func GetWarehousePermissionLevelsResponseToPb(st *GetWarehousePermissionLevelsResponse) (*sqlpb.GetWarehousePermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetWarehousePermissionLevelsResponsePb{}

	var permissionLevelsPb []sqlpb.WarehousePermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := WarehousePermissionsDescriptionToPb(&item)
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

func GetWarehousePermissionLevelsResponseFromPb(pb *sqlpb.GetWarehousePermissionLevelsResponsePb) (*GetWarehousePermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehousePermissionLevelsResponse{}

	var permissionLevelsField []WarehousePermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := WarehousePermissionsDescriptionFromPb(&itemPb)
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

type GetWarehousePermissionsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	// Wire name: 'warehouse_id'
	WarehouseId string `tf:"-"`
}

func GetWarehousePermissionsRequestToPb(st *GetWarehousePermissionsRequest) (*sqlpb.GetWarehousePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetWarehousePermissionsRequestPb{}
	pb.WarehouseId = st.WarehouseId

	return pb, nil
}

func GetWarehousePermissionsRequestFromPb(pb *sqlpb.GetWarehousePermissionsRequestPb) (*GetWarehousePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehousePermissionsRequest{}
	st.WarehouseId = pb.WarehouseId

	return st, nil
}

type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func GetWarehouseRequestToPb(st *GetWarehouseRequest) (*sqlpb.GetWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetWarehouseRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetWarehouseRequestFromPb(pb *sqlpb.GetWarehouseRequestPb) (*GetWarehouseRequest, error) {
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
	AutoStopMins int ``
	// Channel Details
	// Wire name: 'channel'
	Channel *Channel ``
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	// Wire name: 'cluster_size'
	ClusterSize string ``
	// warehouse creator name
	// Wire name: 'creator_name'
	CreatorName string ``
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	// Wire name: 'enable_photon'
	EnablePhoton bool ``
	// Configures whether the warehouse should use serverless compute
	// Wire name: 'enable_serverless_compute'
	EnableServerlessCompute bool ``
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	// Wire name: 'health'
	Health *EndpointHealth ``
	// unique identifier for warehouse
	// Wire name: 'id'
	Id string ``
	// Deprecated. Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// the jdbc connection string for this warehouse
	// Wire name: 'jdbc_url'
	JdbcUrl string ``
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
	// Wire name: 'max_num_clusters'
	MaxNumClusters int ``
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
	MinNumClusters int ``
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	// Wire name: 'name'
	Name string ``
	// Deprecated. current number of active sessions for the warehouse
	// Wire name: 'num_active_sessions'
	NumActiveSessions int64 ``
	// current number of clusters running for the service
	// Wire name: 'num_clusters'
	NumClusters int ``
	// ODBC parameters for the SQL warehouse
	// Wire name: 'odbc_params'
	OdbcParams *OdbcParams ``

	// Wire name: 'spot_instance_policy'
	SpotInstancePolicy SpotInstancePolicy ``

	// Wire name: 'state'
	State State ``
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	// Wire name: 'tags'
	Tags *EndpointTags ``

	// Wire name: 'warehouse_type'
	WarehouseType   GetWarehouseResponseWarehouseType ``
	ForceSendFields []string                          `tf:"-"`
}

func (s *GetWarehouseResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetWarehouseResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetWarehouseResponseToPb(st *GetWarehouseResponse) (*sqlpb.GetWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetWarehouseResponsePb{}
	pb.AutoStopMins = st.AutoStopMins
	channelPb, err := ChannelToPb(st.Channel)
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
	healthPb, err := EndpointHealthToPb(st.Health)
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
	odbcParamsPb, err := OdbcParamsToPb(st.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsPb != nil {
		pb.OdbcParams = odbcParamsPb
	}
	spotInstancePolicyPb, err := SpotInstancePolicyToPb(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}
	statePb, err := StateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	tagsPb, err := EndpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}
	warehouseTypePb, err := GetWarehouseResponseWarehouseTypeToPb(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetWarehouseResponseFromPb(pb *sqlpb.GetWarehouseResponsePb) (*GetWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehouseResponse{}
	st.AutoStopMins = pb.AutoStopMins
	channelField, err := ChannelFromPb(pb.Channel)
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
	healthField, err := EndpointHealthFromPb(pb.Health)
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
	odbcParamsField, err := OdbcParamsFromPb(pb.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsField != nil {
		st.OdbcParams = odbcParamsField
	}
	spotInstancePolicyField, err := SpotInstancePolicyFromPb(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	stateField, err := StateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	tagsField, err := EndpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := GetWarehouseResponseWarehouseTypeFromPb(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func GetWarehouseResponseWarehouseTypeToPb(st *GetWarehouseResponseWarehouseType) (*sqlpb.GetWarehouseResponseWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.GetWarehouseResponseWarehouseTypePb(*st)
	return &pb, nil
}

func GetWarehouseResponseWarehouseTypeFromPb(pb *sqlpb.GetWarehouseResponseWarehouseTypePb) (*GetWarehouseResponseWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetWarehouseResponseWarehouseType(*pb)
	return &st, nil
}

type GetWorkspaceWarehouseConfigResponse struct {
	// Optional: Channel selection details
	// Wire name: 'channel'
	Channel *Channel ``
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'config_param'
	ConfigParam *RepeatedEndpointConfPairs ``
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	// Wire name: 'data_access_config'
	DataAccessConfig []EndpointConfPair ``
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	// Wire name: 'enabled_warehouse_types'
	EnabledWarehouseTypes []WarehouseTypePair ``
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'global_param'
	GlobalParam *RepeatedEndpointConfPairs ``
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	// Wire name: 'google_service_account'
	GoogleServiceAccount string ``
	// AWS Only: Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// Security policy for warehouses
	// Wire name: 'security_policy'
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy ``
	// SQL configuration parameters
	// Wire name: 'sql_configuration_parameters'
	SqlConfigurationParameters *RepeatedEndpointConfPairs ``
	ForceSendFields            []string                   `tf:"-"`
}

func (s *GetWorkspaceWarehouseConfigResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetWorkspaceWarehouseConfigResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GetWorkspaceWarehouseConfigResponseToPb(st *GetWorkspaceWarehouseConfigResponse) (*sqlpb.GetWorkspaceWarehouseConfigResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.GetWorkspaceWarehouseConfigResponsePb{}
	channelPb, err := ChannelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}
	configParamPb, err := RepeatedEndpointConfPairsToPb(st.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamPb != nil {
		pb.ConfigParam = configParamPb
	}

	var dataAccessConfigPb []sqlpb.EndpointConfPairPb
	for _, item := range st.DataAccessConfig {
		itemPb, err := EndpointConfPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataAccessConfigPb = append(dataAccessConfigPb, *itemPb)
		}
	}
	pb.DataAccessConfig = dataAccessConfigPb

	var enabledWarehouseTypesPb []sqlpb.WarehouseTypePairPb
	for _, item := range st.EnabledWarehouseTypes {
		itemPb, err := WarehouseTypePairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			enabledWarehouseTypesPb = append(enabledWarehouseTypesPb, *itemPb)
		}
	}
	pb.EnabledWarehouseTypes = enabledWarehouseTypesPb
	globalParamPb, err := RepeatedEndpointConfPairsToPb(st.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamPb != nil {
		pb.GlobalParam = globalParamPb
	}
	pb.GoogleServiceAccount = st.GoogleServiceAccount
	pb.InstanceProfileArn = st.InstanceProfileArn
	securityPolicyPb, err := GetWorkspaceWarehouseConfigResponseSecurityPolicyToPb(&st.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyPb != nil {
		pb.SecurityPolicy = *securityPolicyPb
	}
	sqlConfigurationParametersPb, err := RepeatedEndpointConfPairsToPb(st.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersPb != nil {
		pb.SqlConfigurationParameters = sqlConfigurationParametersPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GetWorkspaceWarehouseConfigResponseFromPb(pb *sqlpb.GetWorkspaceWarehouseConfigResponsePb) (*GetWorkspaceWarehouseConfigResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceWarehouseConfigResponse{}
	channelField, err := ChannelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	configParamField, err := RepeatedEndpointConfPairsFromPb(pb.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamField != nil {
		st.ConfigParam = configParamField
	}

	var dataAccessConfigField []EndpointConfPair
	for _, itemPb := range pb.DataAccessConfig {
		item, err := EndpointConfPairFromPb(&itemPb)
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
		item, err := WarehouseTypePairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			enabledWarehouseTypesField = append(enabledWarehouseTypesField, *item)
		}
	}
	st.EnabledWarehouseTypes = enabledWarehouseTypesField
	globalParamField, err := RepeatedEndpointConfPairsFromPb(pb.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamField != nil {
		st.GlobalParam = globalParamField
	}
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.InstanceProfileArn = pb.InstanceProfileArn
	securityPolicyField, err := GetWorkspaceWarehouseConfigResponseSecurityPolicyFromPb(&pb.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyField != nil {
		st.SecurityPolicy = *securityPolicyField
	}
	sqlConfigurationParametersField, err := RepeatedEndpointConfPairsFromPb(pb.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersField != nil {
		st.SqlConfigurationParameters = sqlConfigurationParametersField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func GetWorkspaceWarehouseConfigResponseSecurityPolicyToPb(st *GetWorkspaceWarehouseConfigResponseSecurityPolicy) (*sqlpb.GetWorkspaceWarehouseConfigResponseSecurityPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.GetWorkspaceWarehouseConfigResponseSecurityPolicyPb(*st)
	return &pb, nil
}

func GetWorkspaceWarehouseConfigResponseSecurityPolicyFromPb(pb *sqlpb.GetWorkspaceWarehouseConfigResponseSecurityPolicyPb) (*GetWorkspaceWarehouseConfigResponseSecurityPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := GetWorkspaceWarehouseConfigResponseSecurityPolicy(*pb)
	return &st, nil
}

type LegacyAlert struct {
	// Timestamp when the alert was created.
	// Wire name: 'created_at'
	CreatedAt string ``
	// Alert ID.
	// Wire name: 'id'
	Id string ``
	// Timestamp when the alert was last triggered.
	// Wire name: 'last_triggered_at'
	LastTriggeredAt string ``
	// Name of the alert.
	// Wire name: 'name'
	Name string ``
	// Alert configuration options.
	// Wire name: 'options'
	Options *AlertOptions ``
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string ``

	// Wire name: 'query'
	Query *AlertQuery ``
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	// Wire name: 'rearm'
	Rearm int ``
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	// Wire name: 'state'
	State LegacyAlertState ``
	// Timestamp when the alert was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string ``

	// Wire name: 'user'
	User            *User    ``
	ForceSendFields []string `tf:"-"`
}

func (s *LegacyAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LegacyAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func LegacyAlertToPb(st *LegacyAlert) (*sqlpb.LegacyAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.LegacyAlertPb{}
	pb.CreatedAt = st.CreatedAt
	pb.Id = st.Id
	pb.LastTriggeredAt = st.LastTriggeredAt
	pb.Name = st.Name
	optionsPb, err := AlertOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}
	pb.Parent = st.Parent
	queryPb, err := AlertQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}
	pb.Rearm = st.Rearm
	statePb, err := LegacyAlertStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	pb.UpdatedAt = st.UpdatedAt
	userPb, err := UserToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func LegacyAlertFromPb(pb *sqlpb.LegacyAlertPb) (*LegacyAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LegacyAlert{}
	st.CreatedAt = pb.CreatedAt
	st.Id = pb.Id
	st.LastTriggeredAt = pb.LastTriggeredAt
	st.Name = pb.Name
	optionsField, err := AlertOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	st.Parent = pb.Parent
	queryField, err := AlertQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	st.Rearm = pb.Rearm
	stateField, err := LegacyAlertStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	st.UpdatedAt = pb.UpdatedAt
	userField, err := UserFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func LegacyAlertStateToPb(st *LegacyAlertState) (*sqlpb.LegacyAlertStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.LegacyAlertStatePb(*st)
	return &pb, nil
}

func LegacyAlertStateFromPb(pb *sqlpb.LegacyAlertStatePb) (*LegacyAlertState, error) {
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
	CanEdit bool ``
	// The timestamp when this query was created.
	// Wire name: 'created_at'
	CreatedAt string ``
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// Query ID.
	// Wire name: 'id'
	Id string ``
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	// Wire name: 'is_archived'
	IsArchived bool ``
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	// Wire name: 'is_draft'
	IsDraft bool ``
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	// Wire name: 'is_favorite'
	IsFavorite bool ``
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	// Wire name: 'is_safe'
	IsSafe bool ``

	// Wire name: 'last_modified_by'
	LastModifiedBy *User ``
	// The ID of the user who last saved changes to this query.
	// Wire name: 'last_modified_by_id'
	LastModifiedById int ``
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	// Wire name: 'latest_query_data_id'
	LatestQueryDataId string ``
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string ``

	// Wire name: 'options'
	Options *QueryOptions ``
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string ``
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	// Wire name: 'permission_tier'
	PermissionTier PermissionLevel ``
	// The text of the query to be run.
	// Wire name: 'query'
	Query string ``
	// A SHA-256 hash of the query text along with the authenticated user ID.
	// Wire name: 'query_hash'
	QueryHash string ``
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole ``

	// Wire name: 'tags'
	Tags []string ``
	// The timestamp at which this query was last updated.
	// Wire name: 'updated_at'
	UpdatedAt string ``

	// Wire name: 'user'
	User *User ``
	// The ID of the user who owns the query.
	// Wire name: 'user_id'
	UserId int ``

	// Wire name: 'visualizations'
	Visualizations  []LegacyVisualization ``
	ForceSendFields []string              `tf:"-"`
}

func (s *LegacyQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LegacyQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func LegacyQueryToPb(st *LegacyQuery) (*sqlpb.LegacyQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.LegacyQueryPb{}
	pb.CanEdit = st.CanEdit
	pb.CreatedAt = st.CreatedAt
	pb.DataSourceId = st.DataSourceId
	pb.Description = st.Description
	pb.Id = st.Id
	pb.IsArchived = st.IsArchived
	pb.IsDraft = st.IsDraft
	pb.IsFavorite = st.IsFavorite
	pb.IsSafe = st.IsSafe
	lastModifiedByPb, err := UserToPb(st.LastModifiedBy)
	if err != nil {
		return nil, err
	}
	if lastModifiedByPb != nil {
		pb.LastModifiedBy = lastModifiedByPb
	}
	pb.LastModifiedById = st.LastModifiedById
	pb.LatestQueryDataId = st.LatestQueryDataId
	pb.Name = st.Name
	optionsPb, err := QueryOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}
	pb.Parent = st.Parent
	permissionTierPb, err := PermissionLevelToPb(&st.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierPb != nil {
		pb.PermissionTier = *permissionTierPb
	}
	pb.Query = st.Query
	pb.QueryHash = st.QueryHash
	runAsRolePb, err := RunAsRoleToPb(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}
	pb.Tags = st.Tags
	pb.UpdatedAt = st.UpdatedAt
	userPb, err := UserToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}
	pb.UserId = st.UserId

	var visualizationsPb []sqlpb.LegacyVisualizationPb
	for _, item := range st.Visualizations {
		itemPb, err := LegacyVisualizationToPb(&item)
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

func LegacyQueryFromPb(pb *sqlpb.LegacyQueryPb) (*LegacyQuery, error) {
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
	lastModifiedByField, err := UserFromPb(pb.LastModifiedBy)
	if err != nil {
		return nil, err
	}
	if lastModifiedByField != nil {
		st.LastModifiedBy = lastModifiedByField
	}
	st.LastModifiedById = pb.LastModifiedById
	st.LatestQueryDataId = pb.LatestQueryDataId
	st.Name = pb.Name
	optionsField, err := QueryOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	st.Parent = pb.Parent
	permissionTierField, err := PermissionLevelFromPb(&pb.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierField != nil {
		st.PermissionTier = *permissionTierField
	}
	st.Query = pb.Query
	st.QueryHash = pb.QueryHash
	runAsRoleField, err := RunAsRoleFromPb(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}
	st.Tags = pb.Tags
	st.UpdatedAt = pb.UpdatedAt
	userField, err := UserFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}
	st.UserId = pb.UserId

	var visualizationsField []LegacyVisualization
	for _, itemPb := range pb.Visualizations {
		item, err := LegacyVisualizationFromPb(&itemPb)
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

// The visualization description API changes frequently and is unsupported. You
// can duplicate a visualization by copying description objects received _from
// the API_ and then using them to create a new one with a POST request to the
// same endpoint. Databricks does not recommend constructing ad-hoc
// visualizations entirely in JSON.
type LegacyVisualization struct {

	// Wire name: 'created_at'
	CreatedAt string ``
	// A short description of this visualization. This is not displayed in the
	// UI.
	// Wire name: 'description'
	Description string ``
	// The UUID for this visualization.
	// Wire name: 'id'
	Id string ``
	// The name of the visualization that appears on dashboards and the query
	// screen.
	// Wire name: 'name'
	Name string ``
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	// Wire name: 'options'
	Options any ``

	// Wire name: 'query'
	Query *LegacyQuery ``
	// The type of visualization: chart, table, pivot table, and so on.
	// Wire name: 'type'
	Type string ``

	// Wire name: 'updated_at'
	UpdatedAt       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *LegacyVisualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LegacyVisualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func LegacyVisualizationToPb(st *LegacyVisualization) (*sqlpb.LegacyVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.LegacyVisualizationPb{}
	pb.CreatedAt = st.CreatedAt
	pb.Description = st.Description
	pb.Id = st.Id
	pb.Name = st.Name
	pb.Options = st.Options
	queryPb, err := LegacyQueryToPb(st.Query)
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

func LegacyVisualizationFromPb(pb *sqlpb.LegacyVisualizationPb) (*LegacyVisualization, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LegacyVisualization{}
	st.CreatedAt = pb.CreatedAt
	st.Description = pb.Description
	st.Id = pb.Id
	st.Name = pb.Name
	st.Options = pb.Options
	queryField, err := LegacyQueryFromPb(pb.Query)
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

func LifecycleStateToPb(st *LifecycleState) (*sqlpb.LifecycleStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.LifecycleStatePb(*st)
	return &pb, nil
}

func LifecycleStateFromPb(pb *sqlpb.LifecycleStatePb) (*LifecycleState, error) {
	if pb == nil {
		return nil, nil
	}
	st := LifecycleState(*pb)
	return &st, nil
}

type ListAlertsRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListAlertsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAlertsRequestToPb(st *ListAlertsRequest) (*sqlpb.ListAlertsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListAlertsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListAlertsRequestFromPb(pb *sqlpb.ListAlertsRequestPb) (*ListAlertsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAlertsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'results'
	Results         []ListAlertsResponseAlert ``
	ForceSendFields []string                  `tf:"-"`
}

func (s *ListAlertsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAlertsResponseToPb(st *ListAlertsResponse) (*sqlpb.ListAlertsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListAlertsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []sqlpb.ListAlertsResponseAlertPb
	for _, item := range st.Results {
		itemPb, err := ListAlertsResponseAlertToPb(&item)
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

func ListAlertsResponseFromPb(pb *sqlpb.ListAlertsResponsePb) (*ListAlertsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsResponse{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []ListAlertsResponseAlert
	for _, itemPb := range pb.Results {
		item, err := ListAlertsResponseAlertFromPb(&itemPb)
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

type ListAlertsResponseAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition ``
	// The timestamp indicating when the alert was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string ``
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string ``
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string ``
	// UUID identifying the alert.
	// Wire name: 'id'
	Id string ``
	// The workspace state of the alert. Used for tracking trashed status.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState ``
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool ``
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string ``
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string ``
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int ``
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	// Wire name: 'state'
	State AlertState ``
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	// Wire name: 'trigger_time'
	TriggerTime *time.Time ``
	// The timestamp indicating when the alert was updated.
	// Wire name: 'update_time'
	UpdateTime      *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *ListAlertsResponseAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsResponseAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAlertsResponseAlertToPb(st *ListAlertsResponseAlert) (*sqlpb.ListAlertsResponseAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListAlertsResponseAlertPb{}
	conditionPb, err := AlertConditionToPb(st.Condition)
	if err != nil {
		return nil, err
	}
	if conditionPb != nil {
		pb.Condition = conditionPb
	}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.CustomBody = st.CustomBody
	pb.CustomSubject = st.CustomSubject
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id
	lifecycleStatePb, err := LifecycleStateToPb(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}
	pb.NotifyOnOk = st.NotifyOnOk
	pb.OwnerUserName = st.OwnerUserName
	pb.QueryId = st.QueryId
	pb.SecondsToRetrigger = st.SecondsToRetrigger
	statePb, err := AlertStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}
	triggerTimePb, err := timestampToPb(st.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimePb != nil {
		pb.TriggerTime = *triggerTimePb
	}
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListAlertsResponseAlertFromPb(pb *sqlpb.ListAlertsResponseAlertPb) (*ListAlertsResponseAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsResponseAlert{}
	conditionField, err := AlertConditionFromPb(pb.Condition)
	if err != nil {
		return nil, err
	}
	if conditionField != nil {
		st.Condition = conditionField
	}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.CustomBody = pb.CustomBody
	st.CustomSubject = pb.CustomSubject
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	lifecycleStateField, err := LifecycleStateFromPb(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	st.NotifyOnOk = pb.NotifyOnOk
	st.OwnerUserName = pb.OwnerUserName
	st.QueryId = pb.QueryId
	st.SecondsToRetrigger = pb.SecondsToRetrigger
	stateField, err := AlertStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	triggerTimeField, err := timestampFromPb(&pb.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimeField != nil {
		st.TriggerTime = triggerTimeField
	}
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAlertsV2Request struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListAlertsV2Request) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsV2Request) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAlertsV2RequestToPb(st *ListAlertsV2Request) (*sqlpb.ListAlertsV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListAlertsV2RequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListAlertsV2RequestFromPb(pb *sqlpb.ListAlertsV2RequestPb) (*ListAlertsV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsV2Request{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListAlertsV2Response struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'results'
	Results         []AlertV2 ``
	ForceSendFields []string  `tf:"-"`
}

func (s *ListAlertsV2Response) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAlertsV2Response) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListAlertsV2ResponseToPb(st *ListAlertsV2Response) (*sqlpb.ListAlertsV2ResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListAlertsV2ResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []sqlpb.AlertV2Pb
	for _, item := range st.Results {
		itemPb, err := AlertV2ToPb(&item)
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

func ListAlertsV2ResponseFromPb(pb *sqlpb.ListAlertsV2ResponsePb) (*ListAlertsV2Response, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsV2Response{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []AlertV2
	for _, itemPb := range pb.Results {
		item, err := AlertV2FromPb(&itemPb)
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
	Q               string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListDashboardsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListDashboardsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListDashboardsRequestToPb(st *ListDashboardsRequest) (*sqlpb.ListDashboardsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListDashboardsRequestPb{}
	orderPb, err := ListOrderToPb(&st.Order)
	if err != nil {
		return nil, err
	}
	if orderPb != nil {
		pb.Order = *orderPb
	}
	pb.Page = st.Page
	pb.PageSize = st.PageSize
	pb.Q = st.Q

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListDashboardsRequestFromPb(pb *sqlpb.ListDashboardsRequestPb) (*ListDashboardsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDashboardsRequest{}
	orderField, err := ListOrderFromPb(&pb.Order)
	if err != nil {
		return nil, err
	}
	if orderField != nil {
		st.Order = *orderField
	}
	st.Page = pb.Page
	st.PageSize = pb.PageSize
	st.Q = pb.Q

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ListOrderToPb(st *ListOrder) (*sqlpb.ListOrderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ListOrderPb(*st)
	return &pb, nil
}

func ListOrderFromPb(pb *sqlpb.ListOrderPb) (*ListOrder, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListOrder(*pb)
	return &st, nil
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
	Q               string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListQueriesLegacyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueriesLegacyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQueriesLegacyRequestToPb(st *ListQueriesLegacyRequest) (*sqlpb.ListQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListQueriesLegacyRequestPb{}
	pb.Order = st.Order
	pb.Page = st.Page
	pb.PageSize = st.PageSize
	pb.Q = st.Q

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListQueriesLegacyRequestFromPb(pb *sqlpb.ListQueriesLegacyRequestPb) (*ListQueriesLegacyRequest, error) {
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

type ListQueriesRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListQueriesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueriesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQueriesRequestToPb(st *ListQueriesRequest) (*sqlpb.ListQueriesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListQueriesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListQueriesRequestFromPb(pb *sqlpb.ListQueriesRequestPb) (*ListQueriesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueriesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListQueriesResponse struct {
	// Whether there is another page of results.
	// Wire name: 'has_next_page'
	HasNextPage bool ``
	// A token that can be used to get the next page of results.
	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'res'
	Res             []QueryInfo ``
	ForceSendFields []string    `tf:"-"`
}

func (s *ListQueriesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueriesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQueriesResponseToPb(st *ListQueriesResponse) (*sqlpb.ListQueriesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListQueriesResponsePb{}
	pb.HasNextPage = st.HasNextPage
	pb.NextPageToken = st.NextPageToken

	var resPb []sqlpb.QueryInfoPb
	for _, item := range st.Res {
		itemPb, err := QueryInfoToPb(&item)
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

func ListQueriesResponseFromPb(pb *sqlpb.ListQueriesResponsePb) (*ListQueriesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueriesResponse{}
	st.HasNextPage = pb.HasNextPage
	st.NextPageToken = pb.NextPageToken

	var resField []QueryInfo
	for _, itemPb := range pb.Res {
		item, err := QueryInfoFromPb(&itemPb)
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

type ListQueryHistoryRequest struct {
	// An optional filter object to limit query history results. Accepts
	// parameters such as user IDs, endpoint IDs, and statuses to narrow the
	// returned data. In a URL, the parameters of this filter are specified with
	// dot notation. For example: `filter_by.statement_ids`.
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
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListQueryHistoryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueryHistoryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQueryHistoryRequestToPb(st *ListQueryHistoryRequest) (*sqlpb.ListQueryHistoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListQueryHistoryRequestPb{}
	filterByPb, err := QueryFilterToPb(st.FilterBy)
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

func ListQueryHistoryRequestFromPb(pb *sqlpb.ListQueryHistoryRequestPb) (*ListQueryHistoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryHistoryRequest{}
	filterByField, err := QueryFilterFromPb(pb.FilterBy)
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

type ListQueryObjectsResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'results'
	Results         []ListQueryObjectsResponseQuery ``
	ForceSendFields []string                        `tf:"-"`
}

func (s *ListQueryObjectsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueryObjectsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQueryObjectsResponseToPb(st *ListQueryObjectsResponse) (*sqlpb.ListQueryObjectsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListQueryObjectsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []sqlpb.ListQueryObjectsResponseQueryPb
	for _, item := range st.Results {
		itemPb, err := ListQueryObjectsResponseQueryToPb(&item)
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

func ListQueryObjectsResponseFromPb(pb *sqlpb.ListQueryObjectsResponsePb) (*ListQueryObjectsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryObjectsResponse{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []ListQueryObjectsResponseQuery
	for _, itemPb := range pb.Results {
		item, err := ListQueryObjectsResponseQueryFromPb(&itemPb)
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

type ListQueryObjectsResponseQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool ``
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string ``
	// Timestamp when this query was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string ``
	// UUID identifying the query.
	// Wire name: 'id'
	Id string ``
	// Username of the user who last saved changes to this query.
	// Wire name: 'last_modifier_user_name'
	LastModifierUserName string ``
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState ``
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string ``
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter ``
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string ``
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode ``
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string ``

	// Wire name: 'tags'
	Tags []string ``
	// Timestamp when this query was last updated.
	// Wire name: 'update_time'
	UpdateTime *time.Time ``
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ListQueryObjectsResponseQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQueryObjectsResponseQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListQueryObjectsResponseQueryToPb(st *ListQueryObjectsResponseQuery) (*sqlpb.ListQueryObjectsResponseQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListQueryObjectsResponseQueryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit
	pb.Catalog = st.Catalog
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.Description = st.Description
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id
	pb.LastModifierUserName = st.LastModifierUserName
	lifecycleStatePb, err := LifecycleStateToPb(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}
	pb.OwnerUserName = st.OwnerUserName

	var parametersPb []sqlpb.QueryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := QueryParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb
	pb.QueryText = st.QueryText
	runAsModePb, err := RunAsModeToPb(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}
	pb.Schema = st.Schema
	pb.Tags = st.Tags
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListQueryObjectsResponseQueryFromPb(pb *sqlpb.ListQueryObjectsResponseQueryPb) (*ListQueryObjectsResponseQuery, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryObjectsResponseQuery{}
	st.ApplyAutoLimit = pb.ApplyAutoLimit
	st.Catalog = pb.Catalog
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.Description = pb.Description
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.LastModifierUserName = pb.LastModifierUserName
	lifecycleStateField, err := LifecycleStateFromPb(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	st.OwnerUserName = pb.OwnerUserName

	var parametersField []QueryParameter
	for _, itemPb := range pb.Parameters {
		item, err := QueryParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.QueryText = pb.QueryText
	runAsModeField, err := RunAsModeFromPb(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListResponse struct {
	// The total number of dashboards.
	// Wire name: 'count'
	Count int ``
	// The current page being displayed.
	// Wire name: 'page'
	Page int ``
	// The number of dashboards per page.
	// Wire name: 'page_size'
	PageSize int ``
	// List of dashboards returned.
	// Wire name: 'results'
	Results         []Dashboard ``
	ForceSendFields []string    `tf:"-"`
}

func (s *ListResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListResponseToPb(st *ListResponse) (*sqlpb.ListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListResponsePb{}
	pb.Count = st.Count
	pb.Page = st.Page
	pb.PageSize = st.PageSize

	var resultsPb []sqlpb.DashboardPb
	for _, item := range st.Results {
		itemPb, err := DashboardToPb(&item)
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

func ListResponseFromPb(pb *sqlpb.ListResponsePb) (*ListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListResponse{}
	st.Count = pb.Count
	st.Page = pb.Page
	st.PageSize = pb.PageSize

	var resultsField []Dashboard
	for _, itemPb := range pb.Results {
		item, err := DashboardFromPb(&itemPb)
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

type ListVisualizationsForQueryRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListVisualizationsForQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVisualizationsForQueryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListVisualizationsForQueryRequestToPb(st *ListVisualizationsForQueryRequest) (*sqlpb.ListVisualizationsForQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListVisualizationsForQueryRequestPb{}
	pb.Id = st.Id
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListVisualizationsForQueryRequestFromPb(pb *sqlpb.ListVisualizationsForQueryRequestPb) (*ListVisualizationsForQueryRequest, error) {
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

type ListVisualizationsForQueryResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'results'
	Results         []Visualization ``
	ForceSendFields []string        `tf:"-"`
}

func (s *ListVisualizationsForQueryResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListVisualizationsForQueryResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListVisualizationsForQueryResponseToPb(st *ListVisualizationsForQueryResponse) (*sqlpb.ListVisualizationsForQueryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListVisualizationsForQueryResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []sqlpb.VisualizationPb
	for _, item := range st.Results {
		itemPb, err := VisualizationToPb(&item)
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

func ListVisualizationsForQueryResponseFromPb(pb *sqlpb.ListVisualizationsForQueryResponsePb) (*ListVisualizationsForQueryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVisualizationsForQueryResponse{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []Visualization
	for _, itemPb := range pb.Results {
		item, err := VisualizationFromPb(&itemPb)
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

type ListWarehousesRequest struct {
	// Service Principal which will be used to fetch the list of warehouses. If
	// not specified, the user from the session header is used.
	// Wire name: 'run_as_user_id'
	RunAsUserId     int      `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (s *ListWarehousesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWarehousesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ListWarehousesRequestToPb(st *ListWarehousesRequest) (*sqlpb.ListWarehousesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListWarehousesRequestPb{}
	pb.RunAsUserId = st.RunAsUserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ListWarehousesRequestFromPb(pb *sqlpb.ListWarehousesRequestPb) (*ListWarehousesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWarehousesRequest{}
	st.RunAsUserId = pb.RunAsUserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ListWarehousesResponse struct {
	// A list of warehouses and their configurations.
	// Wire name: 'warehouses'
	Warehouses []EndpointInfo ``
}

func ListWarehousesResponseToPb(st *ListWarehousesResponse) (*sqlpb.ListWarehousesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ListWarehousesResponsePb{}

	var warehousesPb []sqlpb.EndpointInfoPb
	for _, item := range st.Warehouses {
		itemPb, err := EndpointInfoToPb(&item)
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

func ListWarehousesResponseFromPb(pb *sqlpb.ListWarehousesResponsePb) (*ListWarehousesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWarehousesResponse{}

	var warehousesField []EndpointInfo
	for _, itemPb := range pb.Warehouses {
		item, err := EndpointInfoFromPb(&itemPb)
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
	Prefix string ``
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	// Wire name: 'separator'
	Separator string ``
	// Character that suffixes each selected parameter value.
	// Wire name: 'suffix'
	Suffix          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *MultiValuesOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MultiValuesOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func MultiValuesOptionsToPb(st *MultiValuesOptions) (*sqlpb.MultiValuesOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.MultiValuesOptionsPb{}
	pb.Prefix = st.Prefix
	pb.Separator = st.Separator
	pb.Suffix = st.Suffix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func MultiValuesOptionsFromPb(pb *sqlpb.MultiValuesOptionsPb) (*MultiValuesOptions, error) {
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

type NumericValue struct {

	// Wire name: 'value'
	Value           float64  ``
	ForceSendFields []string `tf:"-"`
}

func (s *NumericValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NumericValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func NumericValueToPb(st *NumericValue) (*sqlpb.NumericValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.NumericValuePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func NumericValueFromPb(pb *sqlpb.NumericValuePb) (*NumericValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NumericValue{}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ObjectTypeToPb(st *ObjectType) (*sqlpb.ObjectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ObjectTypePb(*st)
	return &pb, nil
}

func ObjectTypeFromPb(pb *sqlpb.ObjectTypePb) (*ObjectType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ObjectType(*pb)
	return &st, nil
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

func ObjectTypePluralToPb(st *ObjectTypePlural) (*sqlpb.ObjectTypePluralPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ObjectTypePluralPb(*st)
	return &pb, nil
}

func ObjectTypePluralFromPb(pb *sqlpb.ObjectTypePluralPb) (*ObjectTypePlural, error) {
	if pb == nil {
		return nil, nil
	}
	st := ObjectTypePlural(*pb)
	return &st, nil
}

type OdbcParams struct {

	// Wire name: 'hostname'
	Hostname string ``

	// Wire name: 'path'
	Path string ``

	// Wire name: 'port'
	Port int ``

	// Wire name: 'protocol'
	Protocol        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *OdbcParams) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OdbcParams) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func OdbcParamsToPb(st *OdbcParams) (*sqlpb.OdbcParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.OdbcParamsPb{}
	pb.Hostname = st.Hostname
	pb.Path = st.Path
	pb.Port = st.Port
	pb.Protocol = st.Protocol

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func OdbcParamsFromPb(pb *sqlpb.OdbcParamsPb) (*OdbcParams, error) {
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

func OwnableObjectTypeToPb(st *OwnableObjectType) (*sqlpb.OwnableObjectTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.OwnableObjectTypePb(*st)
	return &pb, nil
}

func OwnableObjectTypeFromPb(pb *sqlpb.OwnableObjectTypePb) (*OwnableObjectType, error) {
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
	EnumOptions string ``
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	// Wire name: 'multiValuesOptions'
	MultiValuesOptions *MultiValuesOptions ``
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	// Wire name: 'name'
	Name string ``
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	// Wire name: 'queryId'
	QueryId string ``
	// The text displayed in a parameter picking widget.
	// Wire name: 'title'
	Title string ``
	// Parameters can have several different types.
	// Wire name: 'type'
	Type ParameterType ``
	// The default value for this parameter.
	// Wire name: 'value'
	Value           any      ``
	ForceSendFields []string `tf:"-"`
}

func (s *Parameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Parameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ParameterToPb(st *Parameter) (*sqlpb.ParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ParameterPb{}
	pb.EnumOptions = st.EnumOptions
	multiValuesOptionsPb, err := MultiValuesOptionsToPb(st.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsPb != nil {
		pb.MultiValuesOptions = multiValuesOptionsPb
	}
	pb.Name = st.Name
	pb.QueryId = st.QueryId
	pb.Title = st.Title
	typePb, err := ParameterTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ParameterFromPb(pb *sqlpb.ParameterPb) (*Parameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Parameter{}
	st.EnumOptions = pb.EnumOptions
	multiValuesOptionsField, err := MultiValuesOptionsFromPb(pb.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsField != nil {
		st.MultiValuesOptions = multiValuesOptionsField
	}
	st.Name = pb.Name
	st.QueryId = pb.QueryId
	st.Title = pb.Title
	typeField, err := ParameterTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ParameterTypeToPb(st *ParameterType) (*sqlpb.ParameterTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ParameterTypePb(*st)
	return &pb, nil
}

func ParameterTypeFromPb(pb *sqlpb.ParameterTypePb) (*ParameterType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ParameterType(*pb)
	return &st, nil
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

func PermissionLevelToPb(st *PermissionLevel) (*sqlpb.PermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.PermissionLevelPb(*st)
	return &pb, nil
}

func PermissionLevelFromPb(pb *sqlpb.PermissionLevelPb) (*PermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PermissionLevel(*pb)
	return &st, nil
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

func PlansStateToPb(st *PlansState) (*sqlpb.PlansStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.PlansStatePb(*st)
	return &pb, nil
}

func PlansStateFromPb(pb *sqlpb.PlansStatePb) (*PlansState, error) {
	if pb == nil {
		return nil, nil
	}
	st := PlansState(*pb)
	return &st, nil
}

type Query struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool ``
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string ``
	// Timestamp when this query was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string ``
	// UUID identifying the query.
	// Wire name: 'id'
	Id string ``
	// Username of the user who last saved changes to this query.
	// Wire name: 'last_modifier_user_name'
	LastModifierUserName string ``
	// Indicates whether the query is trashed.
	// Wire name: 'lifecycle_state'
	LifecycleState LifecycleState ``
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string ``
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter ``
	// Workspace path of the workspace folder containing the object.
	// Wire name: 'parent_path'
	ParentPath string ``
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string ``
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode ``
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string ``

	// Wire name: 'tags'
	Tags []string ``
	// Timestamp when this query was last updated.
	// Wire name: 'update_time'
	UpdateTime *time.Time ``
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *Query) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Query) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryToPb(st *Query) (*sqlpb.QueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit
	pb.Catalog = st.Catalog
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.Description = st.Description
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id
	pb.LastModifierUserName = st.LastModifierUserName
	lifecycleStatePb, err := LifecycleStateToPb(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}
	pb.OwnerUserName = st.OwnerUserName

	var parametersPb []sqlpb.QueryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := QueryParameterToPb(&item)
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
	runAsModePb, err := RunAsModeToPb(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}
	pb.Schema = st.Schema
	pb.Tags = st.Tags
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryFromPb(pb *sqlpb.QueryPb) (*Query, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Query{}
	st.ApplyAutoLimit = pb.ApplyAutoLimit
	st.Catalog = pb.Catalog
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.Description = pb.Description
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.LastModifierUserName = pb.LastModifierUserName
	lifecycleStateField, err := LifecycleStateFromPb(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	st.OwnerUserName = pb.OwnerUserName

	var parametersField []QueryParameter
	for _, itemPb := range pb.Parameters {
		item, err := QueryParameterFromPb(&itemPb)
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
	runAsModeField, err := RunAsModeFromPb(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	// Wire name: 'multi_values_options'
	MultiValuesOptions *MultiValuesOptions ``
	// UUID of the query that provides the parameter values.
	// Wire name: 'query_id'
	QueryId string ``
	// List of selected query parameter values.
	// Wire name: 'values'
	Values          []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryBackedValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryBackedValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryBackedValueToPb(st *QueryBackedValue) (*sqlpb.QueryBackedValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryBackedValuePb{}
	multiValuesOptionsPb, err := MultiValuesOptionsToPb(st.MultiValuesOptions)
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

func QueryBackedValueFromPb(pb *sqlpb.QueryBackedValuePb) (*QueryBackedValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryBackedValue{}
	multiValuesOptionsField, err := MultiValuesOptionsFromPb(pb.MultiValuesOptions)
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

type QueryEditContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string ``
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	// Wire name: 'options'
	Options any ``
	// The text of the query to be run.
	// Wire name: 'query'
	Query string ``

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole ``

	// Wire name: 'tags'
	Tags            []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryEditContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEditContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryEditContentToPb(st *QueryEditContent) (*sqlpb.QueryEditContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryEditContentPb{}
	pb.DataSourceId = st.DataSourceId
	pb.Description = st.Description
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Query = st.Query
	pb.QueryId = st.QueryId
	runAsRolePb, err := RunAsRoleToPb(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryEditContentFromPb(pb *sqlpb.QueryEditContentPb) (*QueryEditContent, error) {
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
	runAsRoleField, err := RunAsRoleFromPb(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type QueryFilter struct {
	// A range filter for query submitted time. The time range must be less than
	// or equal to 30 days.
	// Wire name: 'query_start_time_range'
	QueryStartTimeRange *TimeRange ``
	// A list of statement IDs.
	// Wire name: 'statement_ids'
	StatementIds []string ``
	// A list of statuses (QUEUED, RUNNING, CANCELED, FAILED, FINISHED) to match
	// query results. Corresponds to the `status` field in the response.
	// Filtering for multiple statuses is not recommended. Instead, opt to
	// filter by a single status multiple times and then combine the results.
	// Wire name: 'statuses'
	Statuses []QueryStatus ``
	// A list of user IDs who ran the queries.
	// Wire name: 'user_ids'
	UserIds []int64 ``
	// A list of warehouse IDs.
	// Wire name: 'warehouse_ids'
	WarehouseIds []string ``
}

func QueryFilterToPb(st *QueryFilter) (*sqlpb.QueryFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryFilterPb{}
	queryStartTimeRangePb, err := TimeRangeToPb(st.QueryStartTimeRange)
	if err != nil {
		return nil, err
	}
	if queryStartTimeRangePb != nil {
		pb.QueryStartTimeRange = queryStartTimeRangePb
	}
	pb.StatementIds = st.StatementIds

	var statusesPb []sqlpb.QueryStatusPb
	for _, item := range st.Statuses {
		itemPb, err := QueryStatusToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			statusesPb = append(statusesPb, *itemPb)
		}
	}
	pb.Statuses = statusesPb
	pb.UserIds = st.UserIds
	pb.WarehouseIds = st.WarehouseIds

	return pb, nil
}

func QueryFilterFromPb(pb *sqlpb.QueryFilterPb) (*QueryFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryFilter{}
	queryStartTimeRangeField, err := TimeRangeFromPb(pb.QueryStartTimeRange)
	if err != nil {
		return nil, err
	}
	if queryStartTimeRangeField != nil {
		st.QueryStartTimeRange = queryStartTimeRangeField
	}
	st.StatementIds = pb.StatementIds

	var statusesField []QueryStatus
	for _, itemPb := range pb.Statuses {
		item, err := QueryStatusFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			statusesField = append(statusesField, *item)
		}
	}
	st.Statuses = statusesField
	st.UserIds = pb.UserIds
	st.WarehouseIds = pb.WarehouseIds

	return st, nil
}

type QueryInfo struct {
	// SQL Warehouse channel information at the time of query execution
	// Wire name: 'channel_used'
	ChannelUsed *ChannelInfo ``
	// Client application that ran the statement. For example: Databricks SQL
	// Editor, Tableau, and Power BI. This field is derived from information
	// provided by client applications. While values are expected to remain
	// static over time, this cannot be guaranteed.
	// Wire name: 'client_application'
	ClientApplication string ``
	// Total time of the statement execution. This value does not include the
	// time taken to retrieve the results, which can result in a discrepancy
	// between this value and the start-to-finish wall-clock time.
	// Wire name: 'duration'
	Duration int64 ``
	// Alias for `warehouse_id`.
	// Wire name: 'endpoint_id'
	EndpointId string ``
	// Message describing why the query could not complete.
	// Wire name: 'error_message'
	ErrorMessage string ``
	// The ID of the user whose credentials were used to run the query.
	// Wire name: 'executed_as_user_id'
	ExecutedAsUserId int64 ``
	// The email address or username of the user whose credentials were used to
	// run the query.
	// Wire name: 'executed_as_user_name'
	ExecutedAsUserName string ``
	// The time execution of the query ended.
	// Wire name: 'execution_end_time_ms'
	ExecutionEndTimeMs int64 ``
	// Whether more updates for the query are expected.
	// Wire name: 'is_final'
	IsFinal bool ``
	// A key that can be used to look up query details.
	// Wire name: 'lookup_key'
	LookupKey string ``
	// Metrics about query execution.
	// Wire name: 'metrics'
	Metrics *QueryMetrics ``
	// Whether plans exist for the execution, or the reason why they are missing
	// Wire name: 'plans_state'
	PlansState PlansState ``
	// The time the query ended.
	// Wire name: 'query_end_time_ms'
	QueryEndTimeMs int64 ``
	// The query ID.
	// Wire name: 'query_id'
	QueryId string ``
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	// Wire name: 'query_source'
	QuerySource *ExternalQuerySource ``
	// The time the query started.
	// Wire name: 'query_start_time_ms'
	QueryStartTimeMs int64 ``
	// The text of the query.
	// Wire name: 'query_text'
	QueryText string ``
	// The number of results returned by the query.
	// Wire name: 'rows_produced'
	RowsProduced int64 ``
	// URL to the Spark UI query plan.
	// Wire name: 'spark_ui_url'
	SparkUiUrl string ``
	// Type of statement for this query
	// Wire name: 'statement_type'
	StatementType QueryStatementType ``
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	// Wire name: 'status'
	Status QueryStatus ``
	// The ID of the user who ran the query.
	// Wire name: 'user_id'
	UserId int64 ``
	// The email address or username of the user who ran the query.
	// Wire name: 'user_name'
	UserName string ``
	// Warehouse ID.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryInfoToPb(st *QueryInfo) (*sqlpb.QueryInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryInfoPb{}
	channelUsedPb, err := ChannelInfoToPb(st.ChannelUsed)
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
	metricsPb, err := QueryMetricsToPb(st.Metrics)
	if err != nil {
		return nil, err
	}
	if metricsPb != nil {
		pb.Metrics = metricsPb
	}
	plansStatePb, err := PlansStateToPb(&st.PlansState)
	if err != nil {
		return nil, err
	}
	if plansStatePb != nil {
		pb.PlansState = *plansStatePb
	}
	pb.QueryEndTimeMs = st.QueryEndTimeMs
	pb.QueryId = st.QueryId
	querySourcePb, err := ExternalQuerySourceToPb(st.QuerySource)
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
	statementTypePb, err := QueryStatementTypeToPb(&st.StatementType)
	if err != nil {
		return nil, err
	}
	if statementTypePb != nil {
		pb.StatementType = *statementTypePb
	}
	statusPb, err := QueryStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.UserId = st.UserId
	pb.UserName = st.UserName
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryInfoFromPb(pb *sqlpb.QueryInfoPb) (*QueryInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryInfo{}
	channelUsedField, err := ChannelInfoFromPb(pb.ChannelUsed)
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
	metricsField, err := QueryMetricsFromPb(pb.Metrics)
	if err != nil {
		return nil, err
	}
	if metricsField != nil {
		st.Metrics = metricsField
	}
	plansStateField, err := PlansStateFromPb(&pb.PlansState)
	if err != nil {
		return nil, err
	}
	if plansStateField != nil {
		st.PlansState = *plansStateField
	}
	st.QueryEndTimeMs = pb.QueryEndTimeMs
	st.QueryId = pb.QueryId
	querySourceField, err := ExternalQuerySourceFromPb(pb.QuerySource)
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
	statementTypeField, err := QueryStatementTypeFromPb(&pb.StatementType)
	if err != nil {
		return nil, err
	}
	if statementTypeField != nil {
		st.StatementType = *statementTypeField
	}
	statusField, err := QueryStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.UserId = pb.UserId
	st.UserName = pb.UserName
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type QueryList struct {
	// The total number of queries.
	// Wire name: 'count'
	Count int ``
	// The page number that is currently displayed.
	// Wire name: 'page'
	Page int ``
	// The number of queries per page.
	// Wire name: 'page_size'
	PageSize int ``
	// List of queries returned.
	// Wire name: 'results'
	Results         []LegacyQuery ``
	ForceSendFields []string      `tf:"-"`
}

func (s *QueryList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryListToPb(st *QueryList) (*sqlpb.QueryListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryListPb{}
	pb.Count = st.Count
	pb.Page = st.Page
	pb.PageSize = st.PageSize

	var resultsPb []sqlpb.LegacyQueryPb
	for _, item := range st.Results {
		itemPb, err := LegacyQueryToPb(&item)
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

func QueryListFromPb(pb *sqlpb.QueryListPb) (*QueryList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryList{}
	st.Count = pb.Count
	st.Page = pb.Page
	st.PageSize = pb.PageSize

	var resultsField []LegacyQuery
	for _, itemPb := range pb.Results {
		item, err := LegacyQueryFromPb(&itemPb)
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

// A query metric that encapsulates a set of measurements for a single query.
// Metrics come from the driver and are stored in the history service database.
type QueryMetrics struct {
	// Time spent loading metadata and optimizing the query, in milliseconds.
	// Wire name: 'compilation_time_ms'
	CompilationTimeMs int64 ``
	// Time spent executing the query, in milliseconds.
	// Wire name: 'execution_time_ms'
	ExecutionTimeMs int64 ``
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	// Wire name: 'network_sent_bytes'
	NetworkSentBytes int64 ``
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	// Wire name: 'overloading_queue_start_timestamp'
	OverloadingQueueStartTimestamp int64 ``
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	// Wire name: 'photon_total_time_ms'
	PhotonTotalTimeMs int64 ``
	// projected remaining work to be done aggregated across all stages in the
	// query, in milliseconds
	// Wire name: 'projected_remaining_task_total_time_ms'
	ProjectedRemainingTaskTotalTimeMs int64 ``
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	// Wire name: 'provisioning_queue_start_timestamp'
	ProvisioningQueueStartTimestamp int64 ``
	// Total number of bytes in all tables not read due to pruning
	// Wire name: 'pruned_bytes'
	PrunedBytes int64 ``
	// Total number of files from all tables not read due to pruning
	// Wire name: 'pruned_files_count'
	PrunedFilesCount int64 ``
	// Timestamp of when the underlying compute started compilation of the
	// query.
	// Wire name: 'query_compilation_start_timestamp'
	QueryCompilationStartTimestamp int64 ``
	// Total size of data read by the query, in bytes.
	// Wire name: 'read_bytes'
	ReadBytes int64 ``
	// Size of persistent data read from the cache, in bytes.
	// Wire name: 'read_cache_bytes'
	ReadCacheBytes int64 ``
	// Number of files read after pruning
	// Wire name: 'read_files_count'
	ReadFilesCount int64 ``
	// Number of partitions read after pruning.
	// Wire name: 'read_partitions_count'
	ReadPartitionsCount int64 ``
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	// Wire name: 'read_remote_bytes'
	ReadRemoteBytes int64 ``
	// number of remaining tasks to complete this is based on the current status
	// and could be bigger or smaller in the future based on future updates
	// Wire name: 'remaining_task_count'
	RemainingTaskCount int64 ``
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	// Wire name: 'result_fetch_time_ms'
	ResultFetchTimeMs int64 ``
	// `true` if the query result was fetched from cache, `false` otherwise.
	// Wire name: 'result_from_cache'
	ResultFromCache bool ``
	// Total number of rows returned by the query.
	// Wire name: 'rows_produced_count'
	RowsProducedCount int64 ``
	// Total number of rows read by the query.
	// Wire name: 'rows_read_count'
	RowsReadCount int64 ``
	// number of remaining tasks to complete, calculated by autoscaler
	// StatementAnalysis.scala deprecated: use remaining_task_count instead
	// Wire name: 'runnable_tasks'
	RunnableTasks int64 ``
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	// Wire name: 'spill_to_disk_bytes'
	SpillToDiskBytes int64 ``
	// sum of task times completed in a range of wall clock time, approximated
	// to a configurable number of points aggregated over all stages and jobs in
	// the query (based on task_total_time_ms)
	// Wire name: 'task_time_over_time_range'
	TaskTimeOverTimeRange *TaskTimeOverRange ``
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	// Wire name: 'task_total_time_ms'
	TaskTotalTimeMs int64 ``
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	// Wire name: 'total_time_ms'
	TotalTimeMs int64 ``
	// remaining work to be done across all stages in the query, calculated by
	// autoscaler StatementAnalysis.scala, in milliseconds deprecated: using
	// projected_remaining_task_total_time_ms instead
	// Wire name: 'work_to_be_done'
	WorkToBeDone int64 ``
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	// Wire name: 'write_remote_bytes'
	WriteRemoteBytes int64    ``
	ForceSendFields  []string `tf:"-"`
}

func (s *QueryMetrics) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryMetrics) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryMetricsToPb(st *QueryMetrics) (*sqlpb.QueryMetricsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryMetricsPb{}
	pb.CompilationTimeMs = st.CompilationTimeMs
	pb.ExecutionTimeMs = st.ExecutionTimeMs
	pb.NetworkSentBytes = st.NetworkSentBytes
	pb.OverloadingQueueStartTimestamp = st.OverloadingQueueStartTimestamp
	pb.PhotonTotalTimeMs = st.PhotonTotalTimeMs
	pb.ProjectedRemainingTaskTotalTimeMs = st.ProjectedRemainingTaskTotalTimeMs
	pb.ProvisioningQueueStartTimestamp = st.ProvisioningQueueStartTimestamp
	pb.PrunedBytes = st.PrunedBytes
	pb.PrunedFilesCount = st.PrunedFilesCount
	pb.QueryCompilationStartTimestamp = st.QueryCompilationStartTimestamp
	pb.ReadBytes = st.ReadBytes
	pb.ReadCacheBytes = st.ReadCacheBytes
	pb.ReadFilesCount = st.ReadFilesCount
	pb.ReadPartitionsCount = st.ReadPartitionsCount
	pb.ReadRemoteBytes = st.ReadRemoteBytes
	pb.RemainingTaskCount = st.RemainingTaskCount
	pb.ResultFetchTimeMs = st.ResultFetchTimeMs
	pb.ResultFromCache = st.ResultFromCache
	pb.RowsProducedCount = st.RowsProducedCount
	pb.RowsReadCount = st.RowsReadCount
	pb.RunnableTasks = st.RunnableTasks
	pb.SpillToDiskBytes = st.SpillToDiskBytes
	taskTimeOverTimeRangePb, err := TaskTimeOverRangeToPb(st.TaskTimeOverTimeRange)
	if err != nil {
		return nil, err
	}
	if taskTimeOverTimeRangePb != nil {
		pb.TaskTimeOverTimeRange = taskTimeOverTimeRangePb
	}
	pb.TaskTotalTimeMs = st.TaskTotalTimeMs
	pb.TotalTimeMs = st.TotalTimeMs
	pb.WorkToBeDone = st.WorkToBeDone
	pb.WriteRemoteBytes = st.WriteRemoteBytes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryMetricsFromPb(pb *sqlpb.QueryMetricsPb) (*QueryMetrics, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryMetrics{}
	st.CompilationTimeMs = pb.CompilationTimeMs
	st.ExecutionTimeMs = pb.ExecutionTimeMs
	st.NetworkSentBytes = pb.NetworkSentBytes
	st.OverloadingQueueStartTimestamp = pb.OverloadingQueueStartTimestamp
	st.PhotonTotalTimeMs = pb.PhotonTotalTimeMs
	st.ProjectedRemainingTaskTotalTimeMs = pb.ProjectedRemainingTaskTotalTimeMs
	st.ProvisioningQueueStartTimestamp = pb.ProvisioningQueueStartTimestamp
	st.PrunedBytes = pb.PrunedBytes
	st.PrunedFilesCount = pb.PrunedFilesCount
	st.QueryCompilationStartTimestamp = pb.QueryCompilationStartTimestamp
	st.ReadBytes = pb.ReadBytes
	st.ReadCacheBytes = pb.ReadCacheBytes
	st.ReadFilesCount = pb.ReadFilesCount
	st.ReadPartitionsCount = pb.ReadPartitionsCount
	st.ReadRemoteBytes = pb.ReadRemoteBytes
	st.RemainingTaskCount = pb.RemainingTaskCount
	st.ResultFetchTimeMs = pb.ResultFetchTimeMs
	st.ResultFromCache = pb.ResultFromCache
	st.RowsProducedCount = pb.RowsProducedCount
	st.RowsReadCount = pb.RowsReadCount
	st.RunnableTasks = pb.RunnableTasks
	st.SpillToDiskBytes = pb.SpillToDiskBytes
	taskTimeOverTimeRangeField, err := TaskTimeOverRangeFromPb(pb.TaskTimeOverTimeRange)
	if err != nil {
		return nil, err
	}
	if taskTimeOverTimeRangeField != nil {
		st.TaskTimeOverTimeRange = taskTimeOverTimeRangeField
	}
	st.TaskTotalTimeMs = pb.TaskTotalTimeMs
	st.TotalTimeMs = pb.TotalTimeMs
	st.WorkToBeDone = pb.WorkToBeDone
	st.WriteRemoteBytes = pb.WriteRemoteBytes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type QueryOptions struct {
	// The name of the catalog to execute this query in.
	// Wire name: 'catalog'
	Catalog string ``
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	// Wire name: 'moved_to_trash_at'
	MovedToTrashAt string ``

	// Wire name: 'parameters'
	Parameters []Parameter ``
	// The name of the schema to execute this query in.
	// Wire name: 'schema'
	Schema          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryOptionsToPb(st *QueryOptions) (*sqlpb.QueryOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryOptionsPb{}
	pb.Catalog = st.Catalog
	pb.MovedToTrashAt = st.MovedToTrashAt

	var parametersPb []sqlpb.ParameterPb
	for _, item := range st.Parameters {
		itemPb, err := ParameterToPb(&item)
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

func QueryOptionsFromPb(pb *sqlpb.QueryOptionsPb) (*QueryOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryOptions{}
	st.Catalog = pb.Catalog
	st.MovedToTrashAt = pb.MovedToTrashAt

	var parametersField []Parameter
	for _, itemPb := range pb.Parameters {
		item, err := ParameterFromPb(&itemPb)
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

type QueryParameter struct {
	// Date-range query parameter value. Can only specify one of
	// `dynamic_date_range_value` or `date_range_value`.
	// Wire name: 'date_range_value'
	DateRangeValue *DateRangeValue ``
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	// Wire name: 'date_value'
	DateValue *DateValue ``
	// Dropdown query parameter value.
	// Wire name: 'enum_value'
	EnumValue *EnumValue ``
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	// Wire name: 'name'
	Name string ``
	// Numeric query parameter value.
	// Wire name: 'numeric_value'
	NumericValue *NumericValue ``
	// Query-based dropdown query parameter value.
	// Wire name: 'query_backed_value'
	QueryBackedValue *QueryBackedValue ``
	// Text query parameter value.
	// Wire name: 'text_value'
	TextValue *TextValue ``
	// Text displayed in the user-facing parameter widget in the UI.
	// Wire name: 'title'
	Title           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryParameter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryParameter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryParameterToPb(st *QueryParameter) (*sqlpb.QueryParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryParameterPb{}
	dateRangeValuePb, err := DateRangeValueToPb(st.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValuePb != nil {
		pb.DateRangeValue = dateRangeValuePb
	}
	dateValuePb, err := DateValueToPb(st.DateValue)
	if err != nil {
		return nil, err
	}
	if dateValuePb != nil {
		pb.DateValue = dateValuePb
	}
	enumValuePb, err := EnumValueToPb(st.EnumValue)
	if err != nil {
		return nil, err
	}
	if enumValuePb != nil {
		pb.EnumValue = enumValuePb
	}
	pb.Name = st.Name
	numericValuePb, err := NumericValueToPb(st.NumericValue)
	if err != nil {
		return nil, err
	}
	if numericValuePb != nil {
		pb.NumericValue = numericValuePb
	}
	queryBackedValuePb, err := QueryBackedValueToPb(st.QueryBackedValue)
	if err != nil {
		return nil, err
	}
	if queryBackedValuePb != nil {
		pb.QueryBackedValue = queryBackedValuePb
	}
	textValuePb, err := TextValueToPb(st.TextValue)
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

func QueryParameterFromPb(pb *sqlpb.QueryParameterPb) (*QueryParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryParameter{}
	dateRangeValueField, err := DateRangeValueFromPb(pb.DateRangeValue)
	if err != nil {
		return nil, err
	}
	if dateRangeValueField != nil {
		st.DateRangeValue = dateRangeValueField
	}
	dateValueField, err := DateValueFromPb(pb.DateValue)
	if err != nil {
		return nil, err
	}
	if dateValueField != nil {
		st.DateValue = dateValueField
	}
	enumValueField, err := EnumValueFromPb(pb.EnumValue)
	if err != nil {
		return nil, err
	}
	if enumValueField != nil {
		st.EnumValue = enumValueField
	}
	st.Name = pb.Name
	numericValueField, err := NumericValueFromPb(pb.NumericValue)
	if err != nil {
		return nil, err
	}
	if numericValueField != nil {
		st.NumericValue = numericValueField
	}
	queryBackedValueField, err := QueryBackedValueFromPb(pb.QueryBackedValue)
	if err != nil {
		return nil, err
	}
	if queryBackedValueField != nil {
		st.QueryBackedValue = queryBackedValueField
	}
	textValueField, err := TextValueFromPb(pb.TextValue)
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

type QueryPostContent struct {
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	// Wire name: 'data_source_id'
	DataSourceId string ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	// Wire name: 'name'
	Name string ``
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	// Wire name: 'options'
	Options any ``
	// The identifier of the workspace folder containing the object.
	// Wire name: 'parent'
	Parent string ``
	// The text of the query to be run.
	// Wire name: 'query'
	Query string ``
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	// Wire name: 'run_as_role'
	RunAsRole RunAsRole ``

	// Wire name: 'tags'
	Tags            []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryPostContent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryPostContent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryPostContentToPb(st *QueryPostContent) (*sqlpb.QueryPostContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.QueryPostContentPb{}
	pb.DataSourceId = st.DataSourceId
	pb.Description = st.Description
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Parent = st.Parent
	pb.Query = st.Query
	runAsRolePb, err := RunAsRoleToPb(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}
	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryPostContentFromPb(pb *sqlpb.QueryPostContentPb) (*QueryPostContent, error) {
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
	runAsRoleField, err := RunAsRoleFromPb(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func QueryStatementTypeToPb(st *QueryStatementType) (*sqlpb.QueryStatementTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.QueryStatementTypePb(*st)
	return &pb, nil
}

func QueryStatementTypeFromPb(pb *sqlpb.QueryStatementTypePb) (*QueryStatementType, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueryStatementType(*pb)
	return &st, nil
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

func QueryStatusToPb(st *QueryStatus) (*sqlpb.QueryStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.QueryStatusPb(*st)
	return &pb, nil
}

func QueryStatusFromPb(pb *sqlpb.QueryStatusPb) (*QueryStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueryStatus(*pb)
	return &st, nil
}

type RepeatedEndpointConfPairs struct {
	// Deprecated: Use configuration_pairs
	// Wire name: 'config_pair'
	ConfigPair []EndpointConfPair ``

	// Wire name: 'configuration_pairs'
	ConfigurationPairs []EndpointConfPair ``
}

func RepeatedEndpointConfPairsToPb(st *RepeatedEndpointConfPairs) (*sqlpb.RepeatedEndpointConfPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.RepeatedEndpointConfPairsPb{}

	var configPairPb []sqlpb.EndpointConfPairPb
	for _, item := range st.ConfigPair {
		itemPb, err := EndpointConfPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			configPairPb = append(configPairPb, *itemPb)
		}
	}
	pb.ConfigPair = configPairPb

	var configurationPairsPb []sqlpb.EndpointConfPairPb
	for _, item := range st.ConfigurationPairs {
		itemPb, err := EndpointConfPairToPb(&item)
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

func RepeatedEndpointConfPairsFromPb(pb *sqlpb.RepeatedEndpointConfPairsPb) (*RepeatedEndpointConfPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepeatedEndpointConfPairs{}

	var configPairField []EndpointConfPair
	for _, itemPb := range pb.ConfigPair {
		item, err := EndpointConfPairFromPb(&itemPb)
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
		item, err := EndpointConfPairFromPb(&itemPb)
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

type RestoreDashboardRequest struct {

	// Wire name: 'dashboard_id'
	DashboardId string `tf:"-"`
}

func RestoreDashboardRequestToPb(st *RestoreDashboardRequest) (*sqlpb.RestoreDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.RestoreDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
}

func RestoreDashboardRequestFromPb(pb *sqlpb.RestoreDashboardRequestPb) (*RestoreDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreDashboardRequest{}
	st.DashboardId = pb.DashboardId

	return st, nil
}

type RestoreQueriesLegacyRequest struct {

	// Wire name: 'query_id'
	QueryId string `tf:"-"`
}

func RestoreQueriesLegacyRequestToPb(st *RestoreQueriesLegacyRequest) (*sqlpb.RestoreQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.RestoreQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
}

func RestoreQueriesLegacyRequestFromPb(pb *sqlpb.RestoreQueriesLegacyRequestPb) (*RestoreQueriesLegacyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreQueriesLegacyRequest{}
	st.QueryId = pb.QueryId

	return st, nil
}

type ResultData struct {
	// The number of bytes in the result chunk. This field is not available when
	// using `INLINE` disposition.
	// Wire name: 'byte_count'
	ByteCount int64 ``
	// The position within the sequence of result set chunks.
	// Wire name: 'chunk_index'
	ChunkIndex int ``
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	// Wire name: 'data_array'
	DataArray [][]string ``

	// Wire name: 'external_links'
	ExternalLinks []ExternalLink ``
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	// Wire name: 'next_chunk_index'
	NextChunkIndex int ``
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	// Wire name: 'next_chunk_internal_link'
	NextChunkInternalLink string ``
	// The number of rows within the result chunk.
	// Wire name: 'row_count'
	RowCount int64 ``
	// The starting row offset within the result set.
	// Wire name: 'row_offset'
	RowOffset       int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *ResultData) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultData) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ResultDataToPb(st *ResultData) (*sqlpb.ResultDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ResultDataPb{}
	pb.ByteCount = st.ByteCount
	pb.ChunkIndex = st.ChunkIndex
	pb.DataArray = st.DataArray

	var externalLinksPb []sqlpb.ExternalLinkPb
	for _, item := range st.ExternalLinks {
		itemPb, err := ExternalLinkToPb(&item)
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

func ResultDataFromPb(pb *sqlpb.ResultDataPb) (*ResultData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultData{}
	st.ByteCount = pb.ByteCount
	st.ChunkIndex = pb.ChunkIndex
	st.DataArray = pb.DataArray

	var externalLinksField []ExternalLink
	for _, itemPb := range pb.ExternalLinks {
		item, err := ExternalLinkFromPb(&itemPb)
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

// The result manifest provides schema and metadata for the result set.
type ResultManifest struct {
	// Array of result set chunk metadata.
	// Wire name: 'chunks'
	Chunks []BaseChunkInfo ``

	// Wire name: 'format'
	Format Format ``

	// Wire name: 'schema'
	Schema *ResultSchema ``
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	// Wire name: 'total_byte_count'
	TotalByteCount int64 ``
	// The total number of chunks that the result set has been divided into.
	// Wire name: 'total_chunk_count'
	TotalChunkCount int ``
	// The total number of rows in the result set.
	// Wire name: 'total_row_count'
	TotalRowCount int64 ``
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	// Wire name: 'truncated'
	Truncated       bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *ResultManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ResultManifestToPb(st *ResultManifest) (*sqlpb.ResultManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ResultManifestPb{}

	var chunksPb []sqlpb.BaseChunkInfoPb
	for _, item := range st.Chunks {
		itemPb, err := BaseChunkInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			chunksPb = append(chunksPb, *itemPb)
		}
	}
	pb.Chunks = chunksPb
	formatPb, err := FormatToPb(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}
	schemaPb, err := ResultSchemaToPb(st.Schema)
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

func ResultManifestFromPb(pb *sqlpb.ResultManifestPb) (*ResultManifest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultManifest{}

	var chunksField []BaseChunkInfo
	for _, itemPb := range pb.Chunks {
		item, err := BaseChunkInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			chunksField = append(chunksField, *item)
		}
	}
	st.Chunks = chunksField
	formatField, err := FormatFromPb(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	schemaField, err := ResultSchemaFromPb(pb.Schema)
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

// The schema is an ordered list of column descriptions.
type ResultSchema struct {

	// Wire name: 'column_count'
	ColumnCount int ``

	// Wire name: 'columns'
	Columns         []ColumnInfo ``
	ForceSendFields []string     `tf:"-"`
}

func (s *ResultSchema) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ResultSchema) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ResultSchemaToPb(st *ResultSchema) (*sqlpb.ResultSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ResultSchemaPb{}
	pb.ColumnCount = st.ColumnCount

	var columnsPb []sqlpb.ColumnInfoPb
	for _, item := range st.Columns {
		itemPb, err := ColumnInfoToPb(&item)
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

func ResultSchemaFromPb(pb *sqlpb.ResultSchemaPb) (*ResultSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultSchema{}
	st.ColumnCount = pb.ColumnCount

	var columnsField []ColumnInfo
	for _, itemPb := range pb.Columns {
		item, err := ColumnInfoFromPb(&itemPb)
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

func RunAsModeToPb(st *RunAsMode) (*sqlpb.RunAsModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.RunAsModePb(*st)
	return &pb, nil
}

func RunAsModeFromPb(pb *sqlpb.RunAsModePb) (*RunAsMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunAsMode(*pb)
	return &st, nil
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

func RunAsRoleToPb(st *RunAsRole) (*sqlpb.RunAsRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.RunAsRolePb(*st)
	return &pb, nil
}

func RunAsRoleFromPb(pb *sqlpb.RunAsRolePb) (*RunAsRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunAsRole(*pb)
	return &st, nil
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

func SchedulePauseStatusToPb(st *SchedulePauseStatus) (*sqlpb.SchedulePauseStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.SchedulePauseStatusPb(*st)
	return &pb, nil
}

func SchedulePauseStatusFromPb(pb *sqlpb.SchedulePauseStatusPb) (*SchedulePauseStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := SchedulePauseStatus(*pb)
	return &st, nil
}

type ServiceError struct {

	// Wire name: 'error_code'
	ErrorCode ServiceErrorCode ``
	// A brief summary of the error condition.
	// Wire name: 'message'
	Message         string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServiceError) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServiceError) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServiceErrorToPb(st *ServiceError) (*sqlpb.ServiceErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.ServiceErrorPb{}
	errorCodePb, err := ServiceErrorCodeToPb(&st.ErrorCode)
	if err != nil {
		return nil, err
	}
	if errorCodePb != nil {
		pb.ErrorCode = *errorCodePb
	}
	pb.Message = st.Message

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServiceErrorFromPb(pb *sqlpb.ServiceErrorPb) (*ServiceError, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServiceError{}
	errorCodeField, err := ServiceErrorCodeFromPb(&pb.ErrorCode)
	if err != nil {
		return nil, err
	}
	if errorCodeField != nil {
		st.ErrorCode = *errorCodeField
	}
	st.Message = pb.Message

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ServiceErrorCodeToPb(st *ServiceErrorCode) (*sqlpb.ServiceErrorCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.ServiceErrorCodePb(*st)
	return &pb, nil
}

func ServiceErrorCodeFromPb(pb *sqlpb.ServiceErrorCodePb) (*ServiceErrorCode, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServiceErrorCode(*pb)
	return &st, nil
}

// Set object ACL
type SetRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl ``
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	// Wire name: 'objectId'
	ObjectId string `tf:"-"`
	// The type of object permission to set.
	// Wire name: 'objectType'
	ObjectType ObjectTypePlural `tf:"-"`
}

func SetRequestToPb(st *SetRequest) (*sqlpb.SetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.SetRequestPb{}

	var accessControlListPb []sqlpb.AccessControlPb
	for _, item := range st.AccessControlList {
		itemPb, err := AccessControlToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	objectTypePb, err := ObjectTypePluralToPb(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	return pb, nil
}

func SetRequestFromPb(pb *sqlpb.SetRequestPb) (*SetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetRequest{}

	var accessControlListField []AccessControl
	for _, itemPb := range pb.AccessControlList {
		item, err := AccessControlFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	objectTypeField, err := ObjectTypePluralFromPb(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	return st, nil
}

type SetResponse struct {

	// Wire name: 'access_control_list'
	AccessControlList []AccessControl ``
	// An object's type and UUID, separated by a forward slash (/) character.
	// Wire name: 'object_id'
	ObjectId string ``
	// A singular noun object type.
	// Wire name: 'object_type'
	ObjectType      ObjectType ``
	ForceSendFields []string   `tf:"-"`
}

func (s *SetResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SetResponseToPb(st *SetResponse) (*sqlpb.SetResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.SetResponsePb{}

	var accessControlListPb []sqlpb.AccessControlPb
	for _, item := range st.AccessControlList {
		itemPb, err := AccessControlToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	objectTypePb, err := ObjectTypeToPb(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SetResponseFromPb(pb *sqlpb.SetResponsePb) (*SetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetResponse{}

	var accessControlListField []AccessControl
	for _, itemPb := range pb.AccessControlList {
		item, err := AccessControlFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	objectTypeField, err := ObjectTypeFromPb(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	// Wire name: 'channel'
	Channel *Channel ``
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'config_param'
	ConfigParam *RepeatedEndpointConfPairs ``
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	// Wire name: 'data_access_config'
	DataAccessConfig []EndpointConfPair ``
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	// Wire name: 'enabled_warehouse_types'
	EnabledWarehouseTypes []WarehouseTypePair ``
	// Deprecated: Use sql_configuration_parameters
	// Wire name: 'global_param'
	GlobalParam *RepeatedEndpointConfPairs ``
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	// Wire name: 'google_service_account'
	GoogleServiceAccount string ``
	// AWS Only: Instance profile used to pass IAM role to the cluster
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// Security policy for warehouses
	// Wire name: 'security_policy'
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy ``
	// SQL configuration parameters
	// Wire name: 'sql_configuration_parameters'
	SqlConfigurationParameters *RepeatedEndpointConfPairs ``
	ForceSendFields            []string                   `tf:"-"`
}

func (s *SetWorkspaceWarehouseConfigRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SetWorkspaceWarehouseConfigRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func SetWorkspaceWarehouseConfigRequestToPb(st *SetWorkspaceWarehouseConfigRequest) (*sqlpb.SetWorkspaceWarehouseConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.SetWorkspaceWarehouseConfigRequestPb{}
	channelPb, err := ChannelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}
	configParamPb, err := RepeatedEndpointConfPairsToPb(st.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamPb != nil {
		pb.ConfigParam = configParamPb
	}

	var dataAccessConfigPb []sqlpb.EndpointConfPairPb
	for _, item := range st.DataAccessConfig {
		itemPb, err := EndpointConfPairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataAccessConfigPb = append(dataAccessConfigPb, *itemPb)
		}
	}
	pb.DataAccessConfig = dataAccessConfigPb

	var enabledWarehouseTypesPb []sqlpb.WarehouseTypePairPb
	for _, item := range st.EnabledWarehouseTypes {
		itemPb, err := WarehouseTypePairToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			enabledWarehouseTypesPb = append(enabledWarehouseTypesPb, *itemPb)
		}
	}
	pb.EnabledWarehouseTypes = enabledWarehouseTypesPb
	globalParamPb, err := RepeatedEndpointConfPairsToPb(st.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamPb != nil {
		pb.GlobalParam = globalParamPb
	}
	pb.GoogleServiceAccount = st.GoogleServiceAccount
	pb.InstanceProfileArn = st.InstanceProfileArn
	securityPolicyPb, err := SetWorkspaceWarehouseConfigRequestSecurityPolicyToPb(&st.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyPb != nil {
		pb.SecurityPolicy = *securityPolicyPb
	}
	sqlConfigurationParametersPb, err := RepeatedEndpointConfPairsToPb(st.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersPb != nil {
		pb.SqlConfigurationParameters = sqlConfigurationParametersPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func SetWorkspaceWarehouseConfigRequestFromPb(pb *sqlpb.SetWorkspaceWarehouseConfigRequestPb) (*SetWorkspaceWarehouseConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetWorkspaceWarehouseConfigRequest{}
	channelField, err := ChannelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	configParamField, err := RepeatedEndpointConfPairsFromPb(pb.ConfigParam)
	if err != nil {
		return nil, err
	}
	if configParamField != nil {
		st.ConfigParam = configParamField
	}

	var dataAccessConfigField []EndpointConfPair
	for _, itemPb := range pb.DataAccessConfig {
		item, err := EndpointConfPairFromPb(&itemPb)
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
		item, err := WarehouseTypePairFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			enabledWarehouseTypesField = append(enabledWarehouseTypesField, *item)
		}
	}
	st.EnabledWarehouseTypes = enabledWarehouseTypesField
	globalParamField, err := RepeatedEndpointConfPairsFromPb(pb.GlobalParam)
	if err != nil {
		return nil, err
	}
	if globalParamField != nil {
		st.GlobalParam = globalParamField
	}
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.InstanceProfileArn = pb.InstanceProfileArn
	securityPolicyField, err := SetWorkspaceWarehouseConfigRequestSecurityPolicyFromPb(&pb.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyField != nil {
		st.SecurityPolicy = *securityPolicyField
	}
	sqlConfigurationParametersField, err := RepeatedEndpointConfPairsFromPb(pb.SqlConfigurationParameters)
	if err != nil {
		return nil, err
	}
	if sqlConfigurationParametersField != nil {
		st.SqlConfigurationParameters = sqlConfigurationParametersField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func SetWorkspaceWarehouseConfigRequestSecurityPolicyToPb(st *SetWorkspaceWarehouseConfigRequestSecurityPolicy) (*sqlpb.SetWorkspaceWarehouseConfigRequestSecurityPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.SetWorkspaceWarehouseConfigRequestSecurityPolicyPb(*st)
	return &pb, nil
}

func SetWorkspaceWarehouseConfigRequestSecurityPolicyFromPb(pb *sqlpb.SetWorkspaceWarehouseConfigRequestSecurityPolicyPb) (*SetWorkspaceWarehouseConfigRequestSecurityPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := SetWorkspaceWarehouseConfigRequestSecurityPolicy(*pb)
	return &st, nil
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

func SpotInstancePolicyToPb(st *SpotInstancePolicy) (*sqlpb.SpotInstancePolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.SpotInstancePolicyPb(*st)
	return &pb, nil
}

func SpotInstancePolicyFromPb(pb *sqlpb.SpotInstancePolicyPb) (*SpotInstancePolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := SpotInstancePolicy(*pb)
	return &st, nil
}

type StartRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func StartRequestToPb(st *StartRequest) (*sqlpb.StartRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.StartRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func StartRequestFromPb(pb *sqlpb.StartRequestPb) (*StartRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartRequest{}
	st.Id = pb.Id

	return st, nil
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

func StateToPb(st *State) (*sqlpb.StatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.StatePb(*st)
	return &pb, nil
}

func StateFromPb(pb *sqlpb.StatePb) (*State, error) {
	if pb == nil {
		return nil, nil
	}
	st := State(*pb)
	return &st, nil
}

type StatementParameterListItem struct {
	// The name of a parameter marker to be substituted in the statement.
	// Wire name: 'name'
	Name string ``
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	// Wire name: 'type'
	Type string ``
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *StatementParameterListItem) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementParameterListItem) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StatementParameterListItemToPb(st *StatementParameterListItem) (*sqlpb.StatementParameterListItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.StatementParameterListItemPb{}
	pb.Name = st.Name
	pb.Type = st.Type
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func StatementParameterListItemFromPb(pb *sqlpb.StatementParameterListItemPb) (*StatementParameterListItem, error) {
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

type StatementResponse struct {

	// Wire name: 'manifest'
	Manifest *ResultManifest ``

	// Wire name: 'result'
	Result *ResultData ``
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	// Wire name: 'statement_id'
	StatementId string ``

	// Wire name: 'status'
	Status          *StatementStatus ``
	ForceSendFields []string         `tf:"-"`
}

func (s *StatementResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StatementResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func StatementResponseToPb(st *StatementResponse) (*sqlpb.StatementResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.StatementResponsePb{}
	manifestPb, err := ResultManifestToPb(st.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestPb != nil {
		pb.Manifest = manifestPb
	}
	resultPb, err := ResultDataToPb(st.Result)
	if err != nil {
		return nil, err
	}
	if resultPb != nil {
		pb.Result = resultPb
	}
	pb.StatementId = st.StatementId
	statusPb, err := StatementStatusToPb(st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = statusPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func StatementResponseFromPb(pb *sqlpb.StatementResponsePb) (*StatementResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StatementResponse{}
	manifestField, err := ResultManifestFromPb(pb.Manifest)
	if err != nil {
		return nil, err
	}
	if manifestField != nil {
		st.Manifest = manifestField
	}
	resultField, err := ResultDataFromPb(pb.Result)
	if err != nil {
		return nil, err
	}
	if resultField != nil {
		st.Result = resultField
	}
	st.StatementId = pb.StatementId
	statusField, err := StatementStatusFromPb(pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = statusField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func StatementStateToPb(st *StatementState) (*sqlpb.StatementStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.StatementStatePb(*st)
	return &pb, nil
}

func StatementStateFromPb(pb *sqlpb.StatementStatePb) (*StatementState, error) {
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
	Error *ServiceError ``

	// Wire name: 'state'
	State StatementState ``
}

func StatementStatusToPb(st *StatementStatus) (*sqlpb.StatementStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.StatementStatusPb{}
	errorPb, err := ServiceErrorToPb(st.Error)
	if err != nil {
		return nil, err
	}
	if errorPb != nil {
		pb.Error = errorPb
	}
	statePb, err := StatementStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	return pb, nil
}

func StatementStatusFromPb(pb *sqlpb.StatementStatusPb) (*StatementStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StatementStatus{}
	errorField, err := ServiceErrorFromPb(pb.Error)
	if err != nil {
		return nil, err
	}
	if errorField != nil {
		st.Error = errorField
	}
	stateField, err := StatementStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	return st, nil
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

func StatusToPb(st *Status) (*sqlpb.StatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.StatusPb(*st)
	return &pb, nil
}

func StatusFromPb(pb *sqlpb.StatusPb) (*Status, error) {
	if pb == nil {
		return nil, nil
	}
	st := Status(*pb)
	return &st, nil
}

type StopRequest struct {
	// Required. Id of the SQL warehouse.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func StopRequestToPb(st *StopRequest) (*sqlpb.StopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.StopRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func StopRequestFromPb(pb *sqlpb.StopRequestPb) (*StopRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopRequest{}
	st.Id = pb.Id

	return st, nil
}

type Success struct {

	// Wire name: 'message'
	Message SuccessMessage ``
}

func SuccessToPb(st *Success) (*sqlpb.SuccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.SuccessPb{}
	messagePb, err := SuccessMessageToPb(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	return pb, nil
}

func SuccessFromPb(pb *sqlpb.SuccessPb) (*Success, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Success{}
	messageField, err := SuccessMessageFromPb(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}

	return st, nil
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

func SuccessMessageToPb(st *SuccessMessage) (*sqlpb.SuccessMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.SuccessMessagePb(*st)
	return &pb, nil
}

func SuccessMessageFromPb(pb *sqlpb.SuccessMessagePb) (*SuccessMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := SuccessMessage(*pb)
	return &st, nil
}

type TaskTimeOverRange struct {

	// Wire name: 'entries'
	Entries []TaskTimeOverRangeEntry ``
	// interval length for all entries (difference in start time and end time of
	// an entry range) the same for all entries start time of first interval is
	// query_start_time_ms
	// Wire name: 'interval'
	Interval        int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *TaskTimeOverRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskTimeOverRange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TaskTimeOverRangeToPb(st *TaskTimeOverRange) (*sqlpb.TaskTimeOverRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TaskTimeOverRangePb{}

	var entriesPb []sqlpb.TaskTimeOverRangeEntryPb
	for _, item := range st.Entries {
		itemPb, err := TaskTimeOverRangeEntryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			entriesPb = append(entriesPb, *itemPb)
		}
	}
	pb.Entries = entriesPb
	pb.Interval = st.Interval

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TaskTimeOverRangeFromPb(pb *sqlpb.TaskTimeOverRangePb) (*TaskTimeOverRange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskTimeOverRange{}

	var entriesField []TaskTimeOverRangeEntry
	for _, itemPb := range pb.Entries {
		item, err := TaskTimeOverRangeEntryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			entriesField = append(entriesField, *item)
		}
	}
	st.Entries = entriesField
	st.Interval = pb.Interval

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TaskTimeOverRangeEntry struct {
	// total task completion time in this time range, aggregated over all stages
	// and jobs in the query
	// Wire name: 'task_completed_time_ms'
	TaskCompletedTimeMs int64    ``
	ForceSendFields     []string `tf:"-"`
}

func (s *TaskTimeOverRangeEntry) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TaskTimeOverRangeEntry) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TaskTimeOverRangeEntryToPb(st *TaskTimeOverRangeEntry) (*sqlpb.TaskTimeOverRangeEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TaskTimeOverRangeEntryPb{}
	pb.TaskCompletedTimeMs = st.TaskCompletedTimeMs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TaskTimeOverRangeEntryFromPb(pb *sqlpb.TaskTimeOverRangeEntryPb) (*TaskTimeOverRangeEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskTimeOverRangeEntry{}
	st.TaskCompletedTimeMs = pb.TaskCompletedTimeMs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func TerminationReasonToPb(st *TerminationReason) (*sqlpb.TerminationReasonPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TerminationReasonPb{}
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

func TerminationReasonFromPb(pb *sqlpb.TerminationReasonPb) (*TerminationReason, error) {
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

func TerminationReasonCodeToPb(st *TerminationReasonCode) (*sqlpb.TerminationReasonCodePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.TerminationReasonCodePb(*st)
	return &pb, nil
}

func TerminationReasonCodeFromPb(pb *sqlpb.TerminationReasonCodePb) (*TerminationReasonCode, error) {
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

func TerminationReasonTypeToPb(st *TerminationReasonType) (*sqlpb.TerminationReasonTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.TerminationReasonTypePb(*st)
	return &pb, nil
}

func TerminationReasonTypeFromPb(pb *sqlpb.TerminationReasonTypePb) (*TerminationReasonType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TerminationReasonType(*pb)
	return &st, nil
}

type TextValue struct {

	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *TextValue) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TextValue) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TextValueToPb(st *TextValue) (*sqlpb.TextValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TextValuePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TextValueFromPb(pb *sqlpb.TextValuePb) (*TextValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TextValue{}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TimeRange struct {
	// The end time in milliseconds.
	// Wire name: 'end_time_ms'
	EndTimeMs int64 ``
	// The start time in milliseconds.
	// Wire name: 'start_time_ms'
	StartTimeMs     int64    ``
	ForceSendFields []string `tf:"-"`
}

func (s *TimeRange) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TimeRange) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TimeRangeToPb(st *TimeRange) (*sqlpb.TimeRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TimeRangePb{}
	pb.EndTimeMs = st.EndTimeMs
	pb.StartTimeMs = st.StartTimeMs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TimeRangeFromPb(pb *sqlpb.TimeRangePb) (*TimeRange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TimeRange{}
	st.EndTimeMs = pb.EndTimeMs
	st.StartTimeMs = pb.StartTimeMs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TransferOwnershipObjectId struct {
	// Email address for the new owner, who must exist in the workspace.
	// Wire name: 'new_owner'
	NewOwner        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *TransferOwnershipObjectId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransferOwnershipObjectId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TransferOwnershipObjectIdToPb(st *TransferOwnershipObjectId) (*sqlpb.TransferOwnershipObjectIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TransferOwnershipObjectIdPb{}
	pb.NewOwner = st.NewOwner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TransferOwnershipObjectIdFromPb(pb *sqlpb.TransferOwnershipObjectIdPb) (*TransferOwnershipObjectId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransferOwnershipObjectId{}
	st.NewOwner = pb.NewOwner

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TransferOwnershipRequest struct {
	// Email address for the new owner, who must exist in the workspace.
	// Wire name: 'new_owner'
	NewOwner string ``
	// The ID of the object on which to change ownership.
	// Wire name: 'objectId'
	ObjectId TransferOwnershipObjectId `tf:"-"`
	// The type of object on which to change ownership.
	// Wire name: 'objectType'
	ObjectType      OwnableObjectType `tf:"-"`
	ForceSendFields []string          `tf:"-"`
}

func (s *TransferOwnershipRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TransferOwnershipRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func TransferOwnershipRequestToPb(st *TransferOwnershipRequest) (*sqlpb.TransferOwnershipRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TransferOwnershipRequestPb{}
	pb.NewOwner = st.NewOwner
	objectIdPb, err := TransferOwnershipObjectIdToPb(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}
	objectTypePb, err := OwnableObjectTypeToPb(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func TransferOwnershipRequestFromPb(pb *sqlpb.TransferOwnershipRequestPb) (*TransferOwnershipRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransferOwnershipRequest{}
	st.NewOwner = pb.NewOwner
	objectIdField, err := TransferOwnershipObjectIdFromPb(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := OwnableObjectTypeFromPb(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type TrashAlertRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func TrashAlertRequestToPb(st *TrashAlertRequest) (*sqlpb.TrashAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TrashAlertRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func TrashAlertRequestFromPb(pb *sqlpb.TrashAlertRequestPb) (*TrashAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashAlertRequest{}
	st.Id = pb.Id

	return st, nil
}

type TrashAlertV2Request struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func TrashAlertV2RequestToPb(st *TrashAlertV2Request) (*sqlpb.TrashAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TrashAlertV2RequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func TrashAlertV2RequestFromPb(pb *sqlpb.TrashAlertV2RequestPb) (*TrashAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashAlertV2Request{}
	st.Id = pb.Id

	return st, nil
}

type TrashQueryRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func TrashQueryRequestToPb(st *TrashQueryRequest) (*sqlpb.TrashQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.TrashQueryRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func TrashQueryRequestFromPb(pb *sqlpb.TrashQueryRequestPb) (*TrashQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrashQueryRequest{}
	st.Id = pb.Id

	return st, nil
}

type UpdateAlertRequest struct {

	// Wire name: 'alert'
	Alert *UpdateAlertRequestAlert ``
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool ``

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
	UpdateMask      []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateAlertRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAlertRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateAlertRequestToPb(st *UpdateAlertRequest) (*sqlpb.UpdateAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateAlertRequestPb{}
	alertPb, err := UpdateAlertRequestAlertToPb(st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = alertPb
	}
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName
	pb.Id = st.Id
	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateAlertRequestFromPb(pb *sqlpb.UpdateAlertRequestPb) (*UpdateAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertRequest{}
	alertField, err := UpdateAlertRequestAlertFromPb(pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = alertField
	}
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName
	st.Id = pb.Id
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	// Wire name: 'condition'
	Condition *AlertCondition ``
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_body'
	CustomBody string ``
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	// Wire name: 'custom_subject'
	CustomSubject string ``
	// The display name of the alert.
	// Wire name: 'display_name'
	DisplayName string ``
	// Whether to notify alert subscribers when alert returns back to normal.
	// Wire name: 'notify_on_ok'
	NotifyOnOk bool ``
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	// Wire name: 'owner_user_name'
	OwnerUserName string ``
	// UUID of the query attached to the alert.
	// Wire name: 'query_id'
	QueryId string ``
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	// Wire name: 'seconds_to_retrigger'
	SecondsToRetrigger int      ``
	ForceSendFields    []string `tf:"-"`
}

func (s *UpdateAlertRequestAlert) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAlertRequestAlert) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateAlertRequestAlertToPb(st *UpdateAlertRequestAlert) (*sqlpb.UpdateAlertRequestAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateAlertRequestAlertPb{}
	conditionPb, err := AlertConditionToPb(st.Condition)
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

func UpdateAlertRequestAlertFromPb(pb *sqlpb.UpdateAlertRequestAlertPb) (*UpdateAlertRequestAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertRequestAlert{}
	conditionField, err := AlertConditionFromPb(pb.Condition)
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

type UpdateAlertV2Request struct {

	// Wire name: 'alert'
	Alert AlertV2 ``
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
	UpdateMask []string `tf:"-"`
}

func UpdateAlertV2RequestToPb(st *UpdateAlertV2Request) (*sqlpb.UpdateAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateAlertV2RequestPb{}
	alertPb, err := AlertV2ToPb(&st.Alert)
	if err != nil {
		return nil, err
	}
	if alertPb != nil {
		pb.Alert = *alertPb
	}
	pb.Id = st.Id
	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	return pb, nil
}

func UpdateAlertV2RequestFromPb(pb *sqlpb.UpdateAlertV2RequestPb) (*UpdateAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertV2Request{}
	alertField, err := AlertV2FromPb(&pb.Alert)
	if err != nil {
		return nil, err
	}
	if alertField != nil {
		st.Alert = *alertField
	}
	st.Id = pb.Id
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	return st, nil
}

type UpdateQueryRequest struct {
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	// Wire name: 'auto_resolve_display_name'
	AutoResolveDisplayName bool ``

	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'query'
	Query *UpdateQueryRequestQuery ``
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
	UpdateMask      []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateQueryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateQueryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateQueryRequestToPb(st *UpdateQueryRequest) (*sqlpb.UpdateQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateQueryRequestPb{}
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName
	pb.Id = st.Id
	queryPb, err := UpdateQueryRequestQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}
	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateQueryRequestFromPb(pb *sqlpb.UpdateQueryRequestPb) (*UpdateQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateQueryRequest{}
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName
	st.Id = pb.Id
	queryField, err := UpdateQueryRequestQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type UpdateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	// Wire name: 'apply_auto_limit'
	ApplyAutoLimit bool ``
	// Name of the catalog where this query will be executed.
	// Wire name: 'catalog'
	Catalog string ``
	// General description that conveys additional information about this query
	// such as usage notes.
	// Wire name: 'description'
	Description string ``
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	// Wire name: 'display_name'
	DisplayName string ``
	// Username of the user that owns the query.
	// Wire name: 'owner_user_name'
	OwnerUserName string ``
	// List of query parameter definitions.
	// Wire name: 'parameters'
	Parameters []QueryParameter ``
	// Text of the query to be run.
	// Wire name: 'query_text'
	QueryText string ``
	// Sets the "Run as" role for the object.
	// Wire name: 'run_as_mode'
	RunAsMode RunAsMode ``
	// Name of the schema where this query will be executed.
	// Wire name: 'schema'
	Schema string ``

	// Wire name: 'tags'
	Tags []string ``
	// ID of the SQL warehouse attached to the query.
	// Wire name: 'warehouse_id'
	WarehouseId     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateQueryRequestQuery) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateQueryRequestQuery) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateQueryRequestQueryToPb(st *UpdateQueryRequestQuery) (*sqlpb.UpdateQueryRequestQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateQueryRequestQueryPb{}
	pb.ApplyAutoLimit = st.ApplyAutoLimit
	pb.Catalog = st.Catalog
	pb.Description = st.Description
	pb.DisplayName = st.DisplayName
	pb.OwnerUserName = st.OwnerUserName

	var parametersPb []sqlpb.QueryParameterPb
	for _, item := range st.Parameters {
		itemPb, err := QueryParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			parametersPb = append(parametersPb, *itemPb)
		}
	}
	pb.Parameters = parametersPb
	pb.QueryText = st.QueryText
	runAsModePb, err := RunAsModeToPb(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}
	pb.Schema = st.Schema
	pb.Tags = st.Tags
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateQueryRequestQueryFromPb(pb *sqlpb.UpdateQueryRequestQueryPb) (*UpdateQueryRequestQuery, error) {
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
		item, err := QueryParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			parametersField = append(parametersField, *item)
		}
	}
	st.Parameters = parametersField
	st.QueryText = pb.QueryText
	runAsModeField, err := RunAsModeFromPb(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	st.Schema = pb.Schema
	st.Tags = pb.Tags
	st.WarehouseId = pb.WarehouseId

	st.ForceSendFields = pb.ForceSendFields
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
	UpdateMask []string ``

	// Wire name: 'visualization'
	Visualization *UpdateVisualizationRequestVisualization ``
}

func UpdateVisualizationRequestToPb(st *UpdateVisualizationRequest) (*sqlpb.UpdateVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateVisualizationRequestPb{}
	pb.Id = st.Id
	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}
	visualizationPb, err := UpdateVisualizationRequestVisualizationToPb(st.Visualization)
	if err != nil {
		return nil, err
	}
	if visualizationPb != nil {
		pb.Visualization = visualizationPb
	}

	return pb, nil
}

func UpdateVisualizationRequestFromPb(pb *sqlpb.UpdateVisualizationRequestPb) (*UpdateVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateVisualizationRequest{}
	st.Id = pb.Id
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}
	visualizationField, err := UpdateVisualizationRequestVisualizationFromPb(pb.Visualization)
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
	DisplayName string ``
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string ``
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string ``
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateVisualizationRequestVisualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateVisualizationRequestVisualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateVisualizationRequestVisualizationToPb(st *UpdateVisualizationRequestVisualization) (*sqlpb.UpdateVisualizationRequestVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateVisualizationRequestVisualizationPb{}
	pb.DisplayName = st.DisplayName
	pb.SerializedOptions = st.SerializedOptions
	pb.SerializedQueryPlan = st.SerializedQueryPlan
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UpdateVisualizationRequestVisualizationFromPb(pb *sqlpb.UpdateVisualizationRequestVisualizationPb) (*UpdateVisualizationRequestVisualization, error) {
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

type UpdateWidgetRequest struct {
	// Dashboard ID returned by :method:dashboards/create.
	// Wire name: 'dashboard_id'
	DashboardId string ``
	// Widget ID returned by :method:dashboardwidgets/create
	// Wire name: 'id'
	Id string `tf:"-"`

	// Wire name: 'options'
	Options WidgetOptions ``
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	// Wire name: 'text'
	Text string ``
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	// Wire name: 'visualization_id'
	VisualizationId string ``
	// Width of a widget
	// Wire name: 'width'
	Width           int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *UpdateWidgetRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateWidgetRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UpdateWidgetRequestToPb(st *UpdateWidgetRequest) (*sqlpb.UpdateWidgetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UpdateWidgetRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Id = st.Id
	optionsPb, err := WidgetOptionsToPb(&st.Options)
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

func UpdateWidgetRequestFromPb(pb *sqlpb.UpdateWidgetRequestPb) (*UpdateWidgetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWidgetRequest{}
	st.DashboardId = pb.DashboardId
	st.Id = pb.Id
	optionsField, err := WidgetOptionsFromPb(&pb.Options)
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

type User struct {

	// Wire name: 'email'
	Email string ``

	// Wire name: 'id'
	Id int ``

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *User) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s User) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func UserToPb(st *User) (*sqlpb.UserPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.UserPb{}
	pb.Email = st.Email
	pb.Id = st.Id
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func UserFromPb(pb *sqlpb.UserPb) (*User, error) {
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

type Visualization struct {
	// The timestamp indicating when the visualization was created.
	// Wire name: 'create_time'
	CreateTime *time.Time ``
	// The display name of the visualization.
	// Wire name: 'display_name'
	DisplayName string ``
	// UUID identifying the visualization.
	// Wire name: 'id'
	Id string ``
	// UUID of the query that the visualization is attached to.
	// Wire name: 'query_id'
	QueryId string ``
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	// Wire name: 'serialized_options'
	SerializedOptions string ``
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	// Wire name: 'serialized_query_plan'
	SerializedQueryPlan string ``
	// The type of visualization: counter, table, funnel, and so on.
	// Wire name: 'type'
	Type string ``
	// The timestamp indicating when the visualization was updated.
	// Wire name: 'update_time'
	UpdateTime      *time.Time ``
	ForceSendFields []string   `tf:"-"`
}

func (s *Visualization) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Visualization) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func VisualizationToPb(st *Visualization) (*sqlpb.VisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.VisualizationPb{}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id
	pb.QueryId = st.QueryId
	pb.SerializedOptions = st.SerializedOptions
	pb.SerializedQueryPlan = st.SerializedQueryPlan
	pb.Type = st.Type
	updateTimePb, err := timestampToPb(st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func VisualizationFromPb(pb *sqlpb.VisualizationPb) (*Visualization, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Visualization{}
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id
	st.QueryId = pb.QueryId
	st.SerializedOptions = pb.SerializedOptions
	st.SerializedQueryPlan = pb.SerializedQueryPlan
	st.Type = pb.Type
	updateTimeField, err := timestampFromPb(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WarehouseAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *WarehouseAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehouseAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WarehouseAccessControlRequestToPb(st *WarehouseAccessControlRequest) (*sqlpb.WarehouseAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WarehouseAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := WarehousePermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WarehouseAccessControlRequestFromPb(pb *sqlpb.WarehouseAccessControlRequestPb) (*WarehouseAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehouseAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := WarehousePermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []WarehousePermission ``
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

func (s *WarehouseAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehouseAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WarehouseAccessControlResponseToPb(st *WarehouseAccessControlResponse) (*sqlpb.WarehouseAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WarehouseAccessControlResponsePb{}

	var allPermissionsPb []sqlpb.WarehousePermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := WarehousePermissionToPb(&item)
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

func WarehouseAccessControlResponseFromPb(pb *sqlpb.WarehouseAccessControlResponsePb) (*WarehouseAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehouseAccessControlResponse{}

	var allPermissionsField []WarehousePermission
	for _, itemPb := range pb.AllPermissions {
		item, err := WarehousePermissionFromPb(&itemPb)
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

type WarehousePermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *WarehousePermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehousePermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WarehousePermissionToPb(st *WarehousePermission) (*sqlpb.WarehousePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WarehousePermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := WarehousePermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WarehousePermissionFromPb(pb *sqlpb.WarehousePermissionPb) (*WarehousePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := WarehousePermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func WarehousePermissionLevelToPb(st *WarehousePermissionLevel) (*sqlpb.WarehousePermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.WarehousePermissionLevelPb(*st)
	return &pb, nil
}

func WarehousePermissionLevelFromPb(pb *sqlpb.WarehousePermissionLevelPb) (*WarehousePermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := WarehousePermissionLevel(*pb)
	return &st, nil
}

type WarehousePermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []WarehouseAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *WarehousePermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehousePermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WarehousePermissionsToPb(st *WarehousePermissions) (*sqlpb.WarehousePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WarehousePermissionsPb{}

	var accessControlListPb []sqlpb.WarehouseAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := WarehouseAccessControlResponseToPb(&item)
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

func WarehousePermissionsFromPb(pb *sqlpb.WarehousePermissionsPb) (*WarehousePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissions{}

	var accessControlListField []WarehouseAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := WarehouseAccessControlResponseFromPb(&itemPb)
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

type WarehousePermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel WarehousePermissionLevel ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *WarehousePermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehousePermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WarehousePermissionsDescriptionToPb(st *WarehousePermissionsDescription) (*sqlpb.WarehousePermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WarehousePermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := WarehousePermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WarehousePermissionsDescriptionFromPb(pb *sqlpb.WarehousePermissionsDescriptionPb) (*WarehousePermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := WarehousePermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type WarehousePermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []WarehouseAccessControlRequest ``
	// The SQL warehouse for which to get or manage permissions.
	// Wire name: 'warehouse_id'
	WarehouseId string `tf:"-"`
}

func WarehousePermissionsRequestToPb(st *WarehousePermissionsRequest) (*sqlpb.WarehousePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WarehousePermissionsRequestPb{}

	var accessControlListPb []sqlpb.WarehouseAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := WarehouseAccessControlRequestToPb(&item)
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

func WarehousePermissionsRequestFromPb(pb *sqlpb.WarehousePermissionsRequestPb) (*WarehousePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissionsRequest{}

	var accessControlListField []WarehouseAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := WarehouseAccessControlRequestFromPb(&itemPb)
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
	Enabled bool ``
	// Warehouse type: `PRO` or `CLASSIC`.
	// Wire name: 'warehouse_type'
	WarehouseType   WarehouseTypePairWarehouseType ``
	ForceSendFields []string                       `tf:"-"`
}

func (s *WarehouseTypePair) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WarehouseTypePair) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WarehouseTypePairToPb(st *WarehouseTypePair) (*sqlpb.WarehouseTypePairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WarehouseTypePairPb{}
	pb.Enabled = st.Enabled
	warehouseTypePb, err := WarehouseTypePairWarehouseTypeToPb(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WarehouseTypePairFromPb(pb *sqlpb.WarehouseTypePairPb) (*WarehouseTypePair, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehouseTypePair{}
	st.Enabled = pb.Enabled
	warehouseTypeField, err := WarehouseTypePairWarehouseTypeFromPb(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func WarehouseTypePairWarehouseTypeToPb(st *WarehouseTypePairWarehouseType) (*sqlpb.WarehouseTypePairWarehouseTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sqlpb.WarehouseTypePairWarehouseTypePb(*st)
	return &pb, nil
}

func WarehouseTypePairWarehouseTypeFromPb(pb *sqlpb.WarehouseTypePairWarehouseTypePb) (*WarehouseTypePairWarehouseType, error) {
	if pb == nil {
		return nil, nil
	}
	st := WarehouseTypePairWarehouseType(*pb)
	return &st, nil
}

type Widget struct {
	// The unique ID for this widget.
	// Wire name: 'id'
	Id string ``

	// Wire name: 'options'
	Options *WidgetOptions ``
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	// Wire name: 'visualization'
	Visualization *LegacyVisualization ``
	// Unused field.
	// Wire name: 'width'
	Width           int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *Widget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Widget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WidgetToPb(st *Widget) (*sqlpb.WidgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WidgetPb{}
	pb.Id = st.Id
	optionsPb, err := WidgetOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}
	visualizationPb, err := LegacyVisualizationToPb(st.Visualization)
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

func WidgetFromPb(pb *sqlpb.WidgetPb) (*Widget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Widget{}
	st.Id = pb.Id
	optionsField, err := WidgetOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	visualizationField, err := LegacyVisualizationFromPb(pb.Visualization)
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

type WidgetOptions struct {
	// Timestamp when this object was created
	// Wire name: 'created_at'
	CreatedAt string ``
	// Custom description of the widget
	// Wire name: 'description'
	Description string ``
	// Whether this widget is hidden on the dashboard.
	// Wire name: 'isHidden'
	IsHidden bool ``
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	// Wire name: 'parameterMappings'
	ParameterMappings any ``
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	// Wire name: 'position'
	Position *WidgetPosition ``
	// Custom title of the widget
	// Wire name: 'title'
	Title string ``
	// Timestamp of the last time this object was updated.
	// Wire name: 'updated_at'
	UpdatedAt       string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *WidgetOptions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WidgetOptions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WidgetOptionsToPb(st *WidgetOptions) (*sqlpb.WidgetOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WidgetOptionsPb{}
	pb.CreatedAt = st.CreatedAt
	pb.Description = st.Description
	pb.IsHidden = st.IsHidden
	pb.ParameterMappings = st.ParameterMappings
	positionPb, err := WidgetPositionToPb(st.Position)
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

func WidgetOptionsFromPb(pb *sqlpb.WidgetOptionsPb) (*WidgetOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WidgetOptions{}
	st.CreatedAt = pb.CreatedAt
	st.Description = pb.Description
	st.IsHidden = pb.IsHidden
	st.ParameterMappings = pb.ParameterMappings
	positionField, err := WidgetPositionFromPb(pb.Position)
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

// Coordinates of this widget on a dashboard. This portion of the API changes
// frequently and is unsupported.
type WidgetPosition struct {
	// reserved for internal use
	// Wire name: 'autoHeight'
	AutoHeight bool ``
	// column in the dashboard grid. Values start with 0
	// Wire name: 'col'
	Col int ``
	// row in the dashboard grid. Values start with 0
	// Wire name: 'row'
	Row int ``
	// width of the widget measured in dashboard grid cells
	// Wire name: 'sizeX'
	SizeX int ``
	// height of the widget measured in dashboard grid cells
	// Wire name: 'sizeY'
	SizeY           int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *WidgetPosition) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s WidgetPosition) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func WidgetPositionToPb(st *WidgetPosition) (*sqlpb.WidgetPositionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlpb.WidgetPositionPb{}
	pb.AutoHeight = st.AutoHeight
	pb.Col = st.Col
	pb.Row = st.Row
	pb.SizeX = st.SizeX
	pb.SizeY = st.SizeY

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func WidgetPositionFromPb(pb *sqlpb.WidgetPositionPb) (*WidgetPosition, error) {
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
