// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

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

type AccessControl struct {
	GroupName string
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionLevel PermissionLevel

	UserName string

	ForceSendFields []string
}

func accessControlToPb(st *AccessControl) (*accessControlPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accessControlPb{}
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
	Condition *AlertCondition
	// The timestamp indicating when the alert was created.
	CreateTime string
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string
	// The display name of the alert.
	DisplayName string
	// UUID identifying the alert.
	Id string
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState LifecycleState
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string
	// The workspace path of the folder containing the alert.
	ParentPath string
	// UUID of the query attached to the alert.
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State AlertState
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime string
	// The timestamp indicating when the alert was updated.
	UpdateTime string

	ForceSendFields []string
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

	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	customBodyPb, err := identity(&st.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyPb != nil {
		pb.CustomBody = *customBodyPb
	}

	customSubjectPb, err := identity(&st.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectPb != nil {
		pb.CustomSubject = *customSubjectPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lifecycleStatePb, err := identity(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}

	notifyOnOkPb, err := identity(&st.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkPb != nil {
		pb.NotifyOnOk = *notifyOnOkPb
	}

	ownerUserNamePb, err := identity(&st.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNamePb != nil {
		pb.OwnerUserName = *ownerUserNamePb
	}

	parentPathPb, err := identity(&st.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathPb != nil {
		pb.ParentPath = *parentPathPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	secondsToRetriggerPb, err := identity(&st.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerPb != nil {
		pb.SecondsToRetrigger = *secondsToRetriggerPb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	triggerTimePb, err := identity(&st.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimePb != nil {
		pb.TriggerTime = *triggerTimePb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

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
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	customBodyField, err := identity(&pb.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyField != nil {
		st.CustomBody = *customBodyField
	}
	customSubjectField, err := identity(&pb.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectField != nil {
		st.CustomSubject = *customSubjectField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lifecycleStateField, err := identity(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	notifyOnOkField, err := identity(&pb.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkField != nil {
		st.NotifyOnOk = *notifyOnOkField
	}
	ownerUserNameField, err := identity(&pb.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNameField != nil {
		st.OwnerUserName = *ownerUserNameField
	}
	parentPathField, err := identity(&pb.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathField != nil {
		st.ParentPath = *parentPathField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	secondsToRetriggerField, err := identity(&pb.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerField != nil {
		st.SecondsToRetrigger = *secondsToRetriggerField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	triggerTimeField, err := identity(&pb.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimeField != nil {
		st.TriggerTime = *triggerTimeField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

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
	EmptyResultState AlertState
	// Operator used for comparison in alert evaluation.
	Op AlertOperator
	// Name of the column from the query result to use for comparison in alert
	// evaluation.
	Operand *AlertConditionOperand
	// Threshold value used for comparison in alert evaluation.
	Threshold *AlertConditionThreshold
}

func alertConditionToPb(st *AlertCondition) (*alertConditionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConditionPb{}
	emptyResultStatePb, err := identity(&st.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStatePb != nil {
		pb.EmptyResultState = *emptyResultStatePb
	}

	opPb, err := identity(&st.Op)
	if err != nil {
		return nil, err
	}
	if opPb != nil {
		pb.Op = *opPb
	}

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
	emptyResultStateField, err := identity(&pb.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStateField != nil {
		st.EmptyResultState = *emptyResultStateField
	}
	opField, err := identity(&pb.Op)
	if err != nil {
		return nil, err
	}
	if opField != nil {
		st.Op = *opField
	}
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
	Name string

	ForceSendFields []string
}

func alertOperandColumnToPb(st *AlertOperandColumn) (*alertOperandColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertOperandColumnPb{}
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

func (st *alertOperandColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertOperandColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertOperandValue struct {
	BoolValue bool

	DoubleValue float64

	StringValue string

	ForceSendFields []string
}

func alertOperandValueToPb(st *AlertOperandValue) (*alertOperandValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertOperandValuePb{}
	boolValuePb, err := identity(&st.BoolValue)
	if err != nil {
		return nil, err
	}
	if boolValuePb != nil {
		pb.BoolValue = *boolValuePb
	}

	doubleValuePb, err := identity(&st.DoubleValue)
	if err != nil {
		return nil, err
	}
	if doubleValuePb != nil {
		pb.DoubleValue = *doubleValuePb
	}

	stringValuePb, err := identity(&st.StringValue)
	if err != nil {
		return nil, err
	}
	if stringValuePb != nil {
		pb.StringValue = *stringValuePb
	}

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
	boolValueField, err := identity(&pb.BoolValue)
	if err != nil {
		return nil, err
	}
	if boolValueField != nil {
		st.BoolValue = *boolValueField
	}
	doubleValueField, err := identity(&pb.DoubleValue)
	if err != nil {
		return nil, err
	}
	if doubleValueField != nil {
		st.DoubleValue = *doubleValueField
	}
	stringValueField, err := identity(&pb.StringValue)
	if err != nil {
		return nil, err
	}
	if stringValueField != nil {
		st.StringValue = *stringValueField
	}

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
	Column string
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string
	// Custom subject of alert notification, if it exists. This includes email
	// subject, Slack notification header, etc. See [here] for custom templating
	// instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string
	// State that alert evaluates to when query result is empty.
	EmptyResultState AlertOptionsEmptyResultState
	// Whether or not the alert is muted. If an alert is muted, it will not
	// notify users and notification destinations when triggered.
	Muted bool
	// Operator used to compare in alert evaluation: `>`, `>=`, `<`, `<=`, `==`,
	// `!=`
	Op string
	// Value used to compare in alert evaluation. Supported types include
	// strings (eg. 'foobar'), floats (eg. 123.4), and booleans (true).
	Value any

	ForceSendFields []string
}

func alertOptionsToPb(st *AlertOptions) (*alertOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertOptionsPb{}
	columnPb, err := identity(&st.Column)
	if err != nil {
		return nil, err
	}
	if columnPb != nil {
		pb.Column = *columnPb
	}

	customBodyPb, err := identity(&st.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyPb != nil {
		pb.CustomBody = *customBodyPb
	}

	customSubjectPb, err := identity(&st.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectPb != nil {
		pb.CustomSubject = *customSubjectPb
	}

	emptyResultStatePb, err := identity(&st.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStatePb != nil {
		pb.EmptyResultState = *emptyResultStatePb
	}

	mutedPb, err := identity(&st.Muted)
	if err != nil {
		return nil, err
	}
	if mutedPb != nil {
		pb.Muted = *mutedPb
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
	columnField, err := identity(&pb.Column)
	if err != nil {
		return nil, err
	}
	if columnField != nil {
		st.Column = *columnField
	}
	customBodyField, err := identity(&pb.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyField != nil {
		st.CustomBody = *customBodyField
	}
	customSubjectField, err := identity(&pb.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectField != nil {
		st.CustomSubject = *customSubjectField
	}
	emptyResultStateField, err := identity(&pb.EmptyResultState)
	if err != nil {
		return nil, err
	}
	if emptyResultStateField != nil {
		st.EmptyResultState = *emptyResultStateField
	}
	mutedField, err := identity(&pb.Muted)
	if err != nil {
		return nil, err
	}
	if mutedField != nil {
		st.Muted = *mutedField
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
	CreatedAt string
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// Query ID.
	Id string
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived bool
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft bool
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe bool
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string

	Options *QueryOptions
	// The text of the query to be run.
	Query string

	Tags []string
	// The timestamp at which this query was last updated.
	UpdatedAt string
	// The ID of the user who owns the query.
	UserId int

	ForceSendFields []string
}

func alertQueryToPb(st *AlertQuery) (*alertQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertQueryPb{}
	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	dataSourceIdPb, err := identity(&st.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdPb != nil {
		pb.DataSourceId = *dataSourceIdPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	isArchivedPb, err := identity(&st.IsArchived)
	if err != nil {
		return nil, err
	}
	if isArchivedPb != nil {
		pb.IsArchived = *isArchivedPb
	}

	isDraftPb, err := identity(&st.IsDraft)
	if err != nil {
		return nil, err
	}
	if isDraftPb != nil {
		pb.IsDraft = *isDraftPb
	}

	isSafePb, err := identity(&st.IsSafe)
	if err != nil {
		return nil, err
	}
	if isSafePb != nil {
		pb.IsSafe = *isSafePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb, err := queryOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	queryPb, err := identity(&st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = *queryPb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	userIdPb, err := identity(&st.UserId)
	if err != nil {
		return nil, err
	}
	if userIdPb != nil {
		pb.UserId = *userIdPb
	}

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
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	dataSourceIdField, err := identity(&pb.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdField != nil {
		st.DataSourceId = *dataSourceIdField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	isArchivedField, err := identity(&pb.IsArchived)
	if err != nil {
		return nil, err
	}
	if isArchivedField != nil {
		st.IsArchived = *isArchivedField
	}
	isDraftField, err := identity(&pb.IsDraft)
	if err != nil {
		return nil, err
	}
	if isDraftField != nil {
		st.IsDraft = *isDraftField
	}
	isSafeField, err := identity(&pb.IsSafe)
	if err != nil {
		return nil, err
	}
	if isSafeField != nil {
		st.IsSafe = *isSafeField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	optionsField, err := queryOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	queryField, err := identity(&pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = *queryField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	userIdField, err := identity(&pb.UserId)
	if err != nil {
		return nil, err
	}
	if userIdField != nil {
		st.UserId = *userIdField
	}

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
	CreateTime *time.Time
	// Custom description for the alert. support mustache template.
	CustomDescription string
	// Custom summary for the alert. support mustache template.
	CustomSummary string
	// The display name of the alert.
	DisplayName string

	Evaluation *AlertV2Evaluation
	// UUID identifying the alert.
	Id string
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string
	// The workspace path of the folder containing the alert. Can only be set on
	// create, and cannot be updated.
	ParentPath string
	// Text of the query to be run.
	QueryText string
	// The run as username. This field is set to "Unavailable" if the user has
	// been deleted.
	RunAsUserName string

	Schedule *CronSchedule
	// The timestamp indicating when the alert was updated.
	UpdateTime *time.Time
	// ID of the SQL warehouse attached to the alert.
	WarehouseId string

	ForceSendFields []string
}

func alertV2ToPb(st *AlertV2) (*alertV2Pb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2Pb{}
	createTimePb, err := timestampToPb(st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	customDescriptionPb, err := identity(&st.CustomDescription)
	if err != nil {
		return nil, err
	}
	if customDescriptionPb != nil {
		pb.CustomDescription = *customDescriptionPb
	}

	customSummaryPb, err := identity(&st.CustomSummary)
	if err != nil {
		return nil, err
	}
	if customSummaryPb != nil {
		pb.CustomSummary = *customSummaryPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	evaluationPb, err := alertV2EvaluationToPb(st.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationPb != nil {
		pb.Evaluation = evaluationPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lifecycleStatePb, err := identity(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}

	ownerUserNamePb, err := identity(&st.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNamePb != nil {
		pb.OwnerUserName = *ownerUserNamePb
	}

	parentPathPb, err := identity(&st.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathPb != nil {
		pb.ParentPath = *parentPathPb
	}

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	runAsUserNamePb, err := identity(&st.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNamePb != nil {
		pb.RunAsUserName = *runAsUserNamePb
	}

	schedulePb, err := cronScheduleToPb(st.Schedule)
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
	createTimeField, err := timestampFromPb(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = createTimeField
	}
	customDescriptionField, err := identity(&pb.CustomDescription)
	if err != nil {
		return nil, err
	}
	if customDescriptionField != nil {
		st.CustomDescription = *customDescriptionField
	}
	customSummaryField, err := identity(&pb.CustomSummary)
	if err != nil {
		return nil, err
	}
	if customSummaryField != nil {
		st.CustomSummary = *customSummaryField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	evaluationField, err := alertV2EvaluationFromPb(pb.Evaluation)
	if err != nil {
		return nil, err
	}
	if evaluationField != nil {
		st.Evaluation = evaluationField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lifecycleStateField, err := identity(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	ownerUserNameField, err := identity(&pb.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNameField != nil {
		st.OwnerUserName = *ownerUserNameField
	}
	parentPathField, err := identity(&pb.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathField != nil {
		st.ParentPath = *parentPathField
	}
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}
	runAsUserNameField, err := identity(&pb.RunAsUserName)
	if err != nil {
		return nil, err
	}
	if runAsUserNameField != nil {
		st.RunAsUserName = *runAsUserNameField
	}
	scheduleField, err := cronScheduleFromPb(pb.Schedule)
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

func (st *alertV2Pb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2Pb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2Evaluation struct {
	// Operator used for comparison in alert evaluation.
	ComparisonOperator ComparisonOperator
	// Alert state if result is empty.
	EmptyResultState AlertEvaluationState
	// Timestamp of the last evaluation.
	LastEvaluatedAt *time.Time
	// User or Notification Destination to notify when alert is triggered.
	Notification *AlertV2Notification
	// Source column from result to use to evaluate alert
	Source *AlertV2OperandColumn
	// Latest state of alert evaluation.
	State AlertEvaluationState
	// Threshold to user for alert evaluation, can be a column or a value.
	Threshold *AlertV2Operand

	ForceSendFields []string
}

func alertV2EvaluationToPb(st *AlertV2Evaluation) (*alertV2EvaluationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2EvaluationPb{}
	comparisonOperatorPb, err := identity(&st.ComparisonOperator)
	if err != nil {
		return nil, err
	}
	if comparisonOperatorPb != nil {
		pb.ComparisonOperator = *comparisonOperatorPb
	}

	emptyResultStatePb, err := identity(&st.EmptyResultState)
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

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

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
	comparisonOperatorField, err := identity(&pb.ComparisonOperator)
	if err != nil {
		return nil, err
	}
	if comparisonOperatorField != nil {
		st.ComparisonOperator = *comparisonOperatorField
	}
	emptyResultStateField, err := identity(&pb.EmptyResultState)
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
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
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
	NotifyOnOk bool
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	RetriggerSeconds int

	Subscriptions []AlertV2Subscription

	ForceSendFields []string
}

func alertV2NotificationToPb(st *AlertV2Notification) (*alertV2NotificationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2NotificationPb{}
	notifyOnOkPb, err := identity(&st.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkPb != nil {
		pb.NotifyOnOk = *notifyOnOkPb
	}

	retriggerSecondsPb, err := identity(&st.RetriggerSeconds)
	if err != nil {
		return nil, err
	}
	if retriggerSecondsPb != nil {
		pb.RetriggerSeconds = *retriggerSecondsPb
	}

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
	notifyOnOkField, err := identity(&pb.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkField != nil {
		st.NotifyOnOk = *notifyOnOkField
	}
	retriggerSecondsField, err := identity(&pb.RetriggerSeconds)
	if err != nil {
		return nil, err
	}
	if retriggerSecondsField != nil {
		st.RetriggerSeconds = *retriggerSecondsField
	}

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
	Column *AlertV2OperandColumn

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
	Aggregation Aggregation

	Display string

	Name string

	ForceSendFields []string
}

func alertV2OperandColumnToPb(st *AlertV2OperandColumn) (*alertV2OperandColumnPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2OperandColumnPb{}
	aggregationPb, err := identity(&st.Aggregation)
	if err != nil {
		return nil, err
	}
	if aggregationPb != nil {
		pb.Aggregation = *aggregationPb
	}

	displayPb, err := identity(&st.Display)
	if err != nil {
		return nil, err
	}
	if displayPb != nil {
		pb.Display = *displayPb
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
	aggregationField, err := identity(&pb.Aggregation)
	if err != nil {
		return nil, err
	}
	if aggregationField != nil {
		st.Aggregation = *aggregationField
	}
	displayField, err := identity(&pb.Display)
	if err != nil {
		return nil, err
	}
	if displayField != nil {
		st.Display = *displayField
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

func (st *alertV2OperandColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2OperandColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2OperandValue struct {
	BoolValue bool

	DoubleValue float64

	StringValue string

	ForceSendFields []string
}

func alertV2OperandValueToPb(st *AlertV2OperandValue) (*alertV2OperandValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2OperandValuePb{}
	boolValuePb, err := identity(&st.BoolValue)
	if err != nil {
		return nil, err
	}
	if boolValuePb != nil {
		pb.BoolValue = *boolValuePb
	}

	doubleValuePb, err := identity(&st.DoubleValue)
	if err != nil {
		return nil, err
	}
	if doubleValuePb != nil {
		pb.DoubleValue = *doubleValuePb
	}

	stringValuePb, err := identity(&st.StringValue)
	if err != nil {
		return nil, err
	}
	if stringValuePb != nil {
		pb.StringValue = *stringValuePb
	}

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
	boolValueField, err := identity(&pb.BoolValue)
	if err != nil {
		return nil, err
	}
	if boolValueField != nil {
		st.BoolValue = *boolValueField
	}
	doubleValueField, err := identity(&pb.DoubleValue)
	if err != nil {
		return nil, err
	}
	if doubleValueField != nil {
		st.DoubleValue = *doubleValueField
	}
	stringValueField, err := identity(&pb.StringValue)
	if err != nil {
		return nil, err
	}
	if stringValueField != nil {
		st.StringValue = *stringValueField
	}

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
	DestinationId string

	UserEmail string

	ForceSendFields []string
}

func alertV2SubscriptionToPb(st *AlertV2Subscription) (*alertV2SubscriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2SubscriptionPb{}
	destinationIdPb, err := identity(&st.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdPb != nil {
		pb.DestinationId = *destinationIdPb
	}

	userEmailPb, err := identity(&st.UserEmail)
	if err != nil {
		return nil, err
	}
	if userEmailPb != nil {
		pb.UserEmail = *userEmailPb
	}

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
	destinationIdField, err := identity(&pb.DestinationId)
	if err != nil {
		return nil, err
	}
	if destinationIdField != nil {
		st.DestinationId = *destinationIdField
	}
	userEmailField, err := identity(&pb.UserEmail)
	if err != nil {
		return nil, err
	}
	if userEmailField != nil {
		st.UserEmail = *userEmailField
	}

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
	ByteCount int64
	// The position within the sequence of result set chunks.
	ChunkIndex int
	// The number of rows within the result chunk.
	RowCount int64
	// The starting row offset within the result set.
	RowOffset int64

	ForceSendFields []string
}

func baseChunkInfoToPb(st *BaseChunkInfo) (*baseChunkInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &baseChunkInfoPb{}
	byteCountPb, err := identity(&st.ByteCount)
	if err != nil {
		return nil, err
	}
	if byteCountPb != nil {
		pb.ByteCount = *byteCountPb
	}

	chunkIndexPb, err := identity(&st.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexPb != nil {
		pb.ChunkIndex = *chunkIndexPb
	}

	rowCountPb, err := identity(&st.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountPb != nil {
		pb.RowCount = *rowCountPb
	}

	rowOffsetPb, err := identity(&st.RowOffset)
	if err != nil {
		return nil, err
	}
	if rowOffsetPb != nil {
		pb.RowOffset = *rowOffsetPb
	}

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
	byteCountField, err := identity(&pb.ByteCount)
	if err != nil {
		return nil, err
	}
	if byteCountField != nil {
		st.ByteCount = *byteCountField
	}
	chunkIndexField, err := identity(&pb.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexField != nil {
		st.ChunkIndex = *chunkIndexField
	}
	rowCountField, err := identity(&pb.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountField != nil {
		st.RowCount = *rowCountField
	}
	rowOffsetField, err := identity(&pb.RowOffset)
	if err != nil {
		return nil, err
	}
	if rowOffsetField != nil {
		st.RowOffset = *rowOffsetField
	}

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
	StatementId string
}

func cancelExecutionRequestToPb(st *CancelExecutionRequest) (*cancelExecutionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelExecutionRequestPb{}
	statementIdPb, err := identity(&st.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdPb != nil {
		pb.StatementId = *statementIdPb
	}

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
	statementIdField, err := identity(&pb.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdField != nil {
		st.StatementId = *statementIdField
	}

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
	DbsqlVersion string

	Name ChannelName

	ForceSendFields []string
}

func channelToPb(st *Channel) (*channelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &channelPb{}
	dbsqlVersionPb, err := identity(&st.DbsqlVersion)
	if err != nil {
		return nil, err
	}
	if dbsqlVersionPb != nil {
		pb.DbsqlVersion = *dbsqlVersionPb
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
	dbsqlVersionField, err := identity(&pb.DbsqlVersion)
	if err != nil {
		return nil, err
	}
	if dbsqlVersionField != nil {
		st.DbsqlVersion = *dbsqlVersionField
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

func (st *channelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st channelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Details about a Channel.
type ChannelInfo struct {
	// DB SQL Version the Channel is mapped to.
	DbsqlVersion string
	// Name of the channel
	Name ChannelName

	ForceSendFields []string
}

func channelInfoToPb(st *ChannelInfo) (*channelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &channelInfoPb{}
	dbsqlVersionPb, err := identity(&st.DbsqlVersion)
	if err != nil {
		return nil, err
	}
	if dbsqlVersionPb != nil {
		pb.DbsqlVersion = *dbsqlVersionPb
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
	dbsqlVersionField, err := identity(&pb.DbsqlVersion)
	if err != nil {
		return nil, err
	}
	if dbsqlVersionField != nil {
		st.DbsqlVersion = *dbsqlVersionField
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
	AllowCustomJsVisualizations bool

	AllowDownloads bool

	AllowExternalShares bool

	AllowSubscriptions bool

	DateFormat string

	DateTimeFormat string

	DisablePublish bool

	EnableLegacyAutodetectTypes bool

	FeatureShowPermissionsControl bool

	HidePlotlyModeBar bool

	ForceSendFields []string
}

func clientConfigToPb(st *ClientConfig) (*clientConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clientConfigPb{}
	allowCustomJsVisualizationsPb, err := identity(&st.AllowCustomJsVisualizations)
	if err != nil {
		return nil, err
	}
	if allowCustomJsVisualizationsPb != nil {
		pb.AllowCustomJsVisualizations = *allowCustomJsVisualizationsPb
	}

	allowDownloadsPb, err := identity(&st.AllowDownloads)
	if err != nil {
		return nil, err
	}
	if allowDownloadsPb != nil {
		pb.AllowDownloads = *allowDownloadsPb
	}

	allowExternalSharesPb, err := identity(&st.AllowExternalShares)
	if err != nil {
		return nil, err
	}
	if allowExternalSharesPb != nil {
		pb.AllowExternalShares = *allowExternalSharesPb
	}

	allowSubscriptionsPb, err := identity(&st.AllowSubscriptions)
	if err != nil {
		return nil, err
	}
	if allowSubscriptionsPb != nil {
		pb.AllowSubscriptions = *allowSubscriptionsPb
	}

	dateFormatPb, err := identity(&st.DateFormat)
	if err != nil {
		return nil, err
	}
	if dateFormatPb != nil {
		pb.DateFormat = *dateFormatPb
	}

	dateTimeFormatPb, err := identity(&st.DateTimeFormat)
	if err != nil {
		return nil, err
	}
	if dateTimeFormatPb != nil {
		pb.DateTimeFormat = *dateTimeFormatPb
	}

	disablePublishPb, err := identity(&st.DisablePublish)
	if err != nil {
		return nil, err
	}
	if disablePublishPb != nil {
		pb.DisablePublish = *disablePublishPb
	}

	enableLegacyAutodetectTypesPb, err := identity(&st.EnableLegacyAutodetectTypes)
	if err != nil {
		return nil, err
	}
	if enableLegacyAutodetectTypesPb != nil {
		pb.EnableLegacyAutodetectTypes = *enableLegacyAutodetectTypesPb
	}

	featureShowPermissionsControlPb, err := identity(&st.FeatureShowPermissionsControl)
	if err != nil {
		return nil, err
	}
	if featureShowPermissionsControlPb != nil {
		pb.FeatureShowPermissionsControl = *featureShowPermissionsControlPb
	}

	hidePlotlyModeBarPb, err := identity(&st.HidePlotlyModeBar)
	if err != nil {
		return nil, err
	}
	if hidePlotlyModeBarPb != nil {
		pb.HidePlotlyModeBar = *hidePlotlyModeBarPb
	}

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
	allowCustomJsVisualizationsField, err := identity(&pb.AllowCustomJsVisualizations)
	if err != nil {
		return nil, err
	}
	if allowCustomJsVisualizationsField != nil {
		st.AllowCustomJsVisualizations = *allowCustomJsVisualizationsField
	}
	allowDownloadsField, err := identity(&pb.AllowDownloads)
	if err != nil {
		return nil, err
	}
	if allowDownloadsField != nil {
		st.AllowDownloads = *allowDownloadsField
	}
	allowExternalSharesField, err := identity(&pb.AllowExternalShares)
	if err != nil {
		return nil, err
	}
	if allowExternalSharesField != nil {
		st.AllowExternalShares = *allowExternalSharesField
	}
	allowSubscriptionsField, err := identity(&pb.AllowSubscriptions)
	if err != nil {
		return nil, err
	}
	if allowSubscriptionsField != nil {
		st.AllowSubscriptions = *allowSubscriptionsField
	}
	dateFormatField, err := identity(&pb.DateFormat)
	if err != nil {
		return nil, err
	}
	if dateFormatField != nil {
		st.DateFormat = *dateFormatField
	}
	dateTimeFormatField, err := identity(&pb.DateTimeFormat)
	if err != nil {
		return nil, err
	}
	if dateTimeFormatField != nil {
		st.DateTimeFormat = *dateTimeFormatField
	}
	disablePublishField, err := identity(&pb.DisablePublish)
	if err != nil {
		return nil, err
	}
	if disablePublishField != nil {
		st.DisablePublish = *disablePublishField
	}
	enableLegacyAutodetectTypesField, err := identity(&pb.EnableLegacyAutodetectTypes)
	if err != nil {
		return nil, err
	}
	if enableLegacyAutodetectTypesField != nil {
		st.EnableLegacyAutodetectTypes = *enableLegacyAutodetectTypesField
	}
	featureShowPermissionsControlField, err := identity(&pb.FeatureShowPermissionsControl)
	if err != nil {
		return nil, err
	}
	if featureShowPermissionsControlField != nil {
		st.FeatureShowPermissionsControl = *featureShowPermissionsControlField
	}
	hidePlotlyModeBarField, err := identity(&pb.HidePlotlyModeBar)
	if err != nil {
		return nil, err
	}
	if hidePlotlyModeBarField != nil {
		st.HidePlotlyModeBar = *hidePlotlyModeBarField
	}

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
	Name string
	// The ordinal position of the column (starting at position 0).
	Position int
	// The format of the interval type.
	TypeIntervalType string
	// The name of the base data type. This doesn't include details for complex
	// types such as STRUCT, MAP or ARRAY.
	TypeName ColumnInfoTypeName
	// Specifies the number of digits in a number. This applies to the DECIMAL
	// type.
	TypePrecision int
	// Specifies the number of digits to the right of the decimal point in a
	// number. This applies to the DECIMAL type.
	TypeScale int
	// The full SQL type specification.
	TypeText string

	ForceSendFields []string
}

func columnInfoToPb(st *ColumnInfo) (*columnInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &columnInfoPb{}
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

	typeIntervalTypePb, err := identity(&st.TypeIntervalType)
	if err != nil {
		return nil, err
	}
	if typeIntervalTypePb != nil {
		pb.TypeIntervalType = *typeIntervalTypePb
	}

	typeNamePb, err := identity(&st.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNamePb != nil {
		pb.TypeName = *typeNamePb
	}

	typePrecisionPb, err := identity(&st.TypePrecision)
	if err != nil {
		return nil, err
	}
	if typePrecisionPb != nil {
		pb.TypePrecision = *typePrecisionPb
	}

	typeScalePb, err := identity(&st.TypeScale)
	if err != nil {
		return nil, err
	}
	if typeScalePb != nil {
		pb.TypeScale = *typeScalePb
	}

	typeTextPb, err := identity(&st.TypeText)
	if err != nil {
		return nil, err
	}
	if typeTextPb != nil {
		pb.TypeText = *typeTextPb
	}

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
	typeIntervalTypeField, err := identity(&pb.TypeIntervalType)
	if err != nil {
		return nil, err
	}
	if typeIntervalTypeField != nil {
		st.TypeIntervalType = *typeIntervalTypeField
	}
	typeNameField, err := identity(&pb.TypeName)
	if err != nil {
		return nil, err
	}
	if typeNameField != nil {
		st.TypeName = *typeNameField
	}
	typePrecisionField, err := identity(&pb.TypePrecision)
	if err != nil {
		return nil, err
	}
	if typePrecisionField != nil {
		st.TypePrecision = *typePrecisionField
	}
	typeScaleField, err := identity(&pb.TypeScale)
	if err != nil {
		return nil, err
	}
	if typeScaleField != nil {
		st.TypeScale = *typeScaleField
	}
	typeTextField, err := identity(&pb.TypeText)
	if err != nil {
		return nil, err
	}
	if typeTextField != nil {
		st.TypeText = *typeTextField
	}

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
	Name string
	// Alert configuration options.
	Options AlertOptions
	// The identifier of the workspace folder containing the object.
	Parent string
	// Query ID.
	QueryId string
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int

	ForceSendFields []string
}

func createAlertToPb(st *CreateAlert) (*createAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb, err := alertOptionsToPb(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	parentPb, err := identity(&st.Parent)
	if err != nil {
		return nil, err
	}
	if parentPb != nil {
		pb.Parent = *parentPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	rearmPb, err := identity(&st.Rearm)
	if err != nil {
		return nil, err
	}
	if rearmPb != nil {
		pb.Rearm = *rearmPb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	optionsField, err := alertOptionsFromPb(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	parentField, err := identity(&pb.Parent)
	if err != nil {
		return nil, err
	}
	if parentField != nil {
		st.Parent = *parentField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	rearmField, err := identity(&pb.Rearm)
	if err != nil {
		return nil, err
	}
	if rearmField != nil {
		st.Rearm = *rearmField
	}

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
	Alert *CreateAlertRequestAlert
	// If true, automatically resolve alert display name conflicts. Otherwise,
	// fail the request if the alert's display name conflicts with an existing
	// alert's display name.
	AutoResolveDisplayName bool

	ForceSendFields []string
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

	autoResolveDisplayNamePb, err := identity(&st.AutoResolveDisplayName)
	if err != nil {
		return nil, err
	}
	if autoResolveDisplayNamePb != nil {
		pb.AutoResolveDisplayName = *autoResolveDisplayNamePb
	}

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
	autoResolveDisplayNameField, err := identity(&pb.AutoResolveDisplayName)
	if err != nil {
		return nil, err
	}
	if autoResolveDisplayNameField != nil {
		st.AutoResolveDisplayName = *autoResolveDisplayNameField
	}

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
	Condition *AlertCondition
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string
	// The display name of the alert.
	DisplayName string
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool
	// The workspace path of the folder containing the alert.
	ParentPath string
	// UUID of the query attached to the alert.
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int

	ForceSendFields []string
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

	customBodyPb, err := identity(&st.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyPb != nil {
		pb.CustomBody = *customBodyPb
	}

	customSubjectPb, err := identity(&st.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectPb != nil {
		pb.CustomSubject = *customSubjectPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	notifyOnOkPb, err := identity(&st.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkPb != nil {
		pb.NotifyOnOk = *notifyOnOkPb
	}

	parentPathPb, err := identity(&st.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathPb != nil {
		pb.ParentPath = *parentPathPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	secondsToRetriggerPb, err := identity(&st.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerPb != nil {
		pb.SecondsToRetrigger = *secondsToRetriggerPb
	}

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
	customBodyField, err := identity(&pb.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyField != nil {
		st.CustomBody = *customBodyField
	}
	customSubjectField, err := identity(&pb.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectField != nil {
		st.CustomSubject = *customSubjectField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	notifyOnOkField, err := identity(&pb.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkField != nil {
		st.NotifyOnOk = *notifyOnOkField
	}
	parentPathField, err := identity(&pb.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathField != nil {
		st.ParentPath = *parentPathField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	secondsToRetriggerField, err := identity(&pb.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerField != nil {
		st.SecondsToRetrigger = *secondsToRetriggerField
	}

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
	AutoResolveDisplayName bool

	Query *CreateQueryRequestQuery

	ForceSendFields []string
}

func createQueryRequestToPb(st *CreateQueryRequest) (*createQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQueryRequestPb{}
	autoResolveDisplayNamePb, err := identity(&st.AutoResolveDisplayName)
	if err != nil {
		return nil, err
	}
	if autoResolveDisplayNamePb != nil {
		pb.AutoResolveDisplayName = *autoResolveDisplayNamePb
	}

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
	autoResolveDisplayNameField, err := identity(&pb.AutoResolveDisplayName)
	if err != nil {
		return nil, err
	}
	if autoResolveDisplayNameField != nil {
		st.AutoResolveDisplayName = *autoResolveDisplayNameField
	}
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
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	Catalog string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string
	// List of query parameter definitions.
	Parameters []QueryParameter
	// Workspace path of the workspace folder containing the object.
	ParentPath string
	// Text of the query to be run.
	QueryText string
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	Schema string

	Tags []string
	// ID of the SQL warehouse attached to the query.
	WarehouseId string

	ForceSendFields []string
}

func createQueryRequestQueryToPb(st *CreateQueryRequestQuery) (*createQueryRequestQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQueryRequestQueryPb{}
	applyAutoLimitPb, err := identity(&st.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitPb != nil {
		pb.ApplyAutoLimit = *applyAutoLimitPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

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

	parentPathPb, err := identity(&st.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathPb != nil {
		pb.ParentPath = *parentPathPb
	}

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	runAsModePb, err := identity(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

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
	applyAutoLimitField, err := identity(&pb.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitField != nil {
		st.ApplyAutoLimit = *applyAutoLimitField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}

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
	parentPathField, err := identity(&pb.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathField != nil {
		st.ParentPath = *parentPathField
	}
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}
	runAsModeField, err := identity(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
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
	Description string
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name string
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any
	// The identifier returned by :method:queries/create
	QueryId string
	// The type of visualization: chart, table, pivot table, and so on.
	Type string

	ForceSendFields []string
}

func createQueryVisualizationsLegacyRequestToPb(st *CreateQueryVisualizationsLegacyRequest) (*createQueryVisualizationsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQueryVisualizationsLegacyRequestPb{}
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

	optionsPb, err := identity(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
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
	optionsField, err := identity(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
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

func (st *createQueryVisualizationsLegacyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createQueryVisualizationsLegacyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateVisualizationRequest struct {
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
	DisplayName string
	// UUID of the query that the visualization is attached to.
	QueryId string
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string
	// The type of visualization: counter, table, funnel, and so on.
	Type string

	ForceSendFields []string
}

func createVisualizationRequestVisualizationToPb(st *CreateVisualizationRequestVisualization) (*createVisualizationRequestVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVisualizationRequestVisualizationPb{}
	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	serializedOptionsPb, err := identity(&st.SerializedOptions)
	if err != nil {
		return nil, err
	}
	if serializedOptionsPb != nil {
		pb.SerializedOptions = *serializedOptionsPb
	}

	serializedQueryPlanPb, err := identity(&st.SerializedQueryPlan)
	if err != nil {
		return nil, err
	}
	if serializedQueryPlanPb != nil {
		pb.SerializedQueryPlan = *serializedQueryPlanPb
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
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	serializedOptionsField, err := identity(&pb.SerializedOptions)
	if err != nil {
		return nil, err
	}
	if serializedOptionsField != nil {
		st.SerializedOptions = *serializedOptionsField
	}
	serializedQueryPlanField, err := identity(&pb.SerializedQueryPlan)
	if err != nil {
		return nil, err
	}
	if serializedQueryPlanField != nil {
		st.SerializedQueryPlan = *serializedQueryPlanField
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
	AutoStopMins int
	// Channel Details
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string
	// warehouse creator name
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType CreateWarehouseRequestWarehouseType

	ForceSendFields []string
}

func createWarehouseRequestToPb(st *CreateWarehouseRequest) (*createWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWarehouseRequestPb{}
	autoStopMinsPb, err := identity(&st.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsPb != nil {
		pb.AutoStopMins = *autoStopMinsPb
	}

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	clusterSizePb, err := identity(&st.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizePb != nil {
		pb.ClusterSize = *clusterSizePb
	}

	creatorNamePb, err := identity(&st.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNamePb != nil {
		pb.CreatorName = *creatorNamePb
	}

	enablePhotonPb, err := identity(&st.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonPb != nil {
		pb.EnablePhoton = *enablePhotonPb
	}

	enableServerlessComputePb, err := identity(&st.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputePb != nil {
		pb.EnableServerlessCompute = *enableServerlessComputePb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	maxNumClustersPb, err := identity(&st.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersPb != nil {
		pb.MaxNumClusters = *maxNumClustersPb
	}

	minNumClustersPb, err := identity(&st.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersPb != nil {
		pb.MinNumClusters = *minNumClustersPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	spotInstancePolicyPb, err := identity(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	warehouseTypePb, err := identity(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

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
	autoStopMinsField, err := identity(&pb.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsField != nil {
		st.AutoStopMins = *autoStopMinsField
	}
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	clusterSizeField, err := identity(&pb.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizeField != nil {
		st.ClusterSize = *clusterSizeField
	}
	creatorNameField, err := identity(&pb.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNameField != nil {
		st.CreatorName = *creatorNameField
	}
	enablePhotonField, err := identity(&pb.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonField != nil {
		st.EnablePhoton = *enablePhotonField
	}
	enableServerlessComputeField, err := identity(&pb.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputeField != nil {
		st.EnableServerlessCompute = *enableServerlessComputeField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	maxNumClustersField, err := identity(&pb.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersField != nil {
		st.MaxNumClusters = *maxNumClustersField
	}
	minNumClustersField, err := identity(&pb.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersField != nil {
		st.MinNumClusters = *minNumClustersField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	spotInstancePolicyField, err := identity(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := identity(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

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
	Id string

	ForceSendFields []string
}

func createWarehouseResponseToPb(st *CreateWarehouseResponse) (*createWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWarehouseResponsePb{}
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

func (st *createWarehouseResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createWarehouseResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateWidget struct {
	// Dashboard ID returned by :method:dashboards/create.
	DashboardId string
	// Widget ID returned by :method:dashboardwidgets/create
	Id string

	Options WidgetOptions
	// If this is a textbox widget, the application displays this text. This
	// field is ignored if the widget contains a visualization in the
	// `visualization` field.
	Text string
	// Query Vizualization ID returned by :method:queryvisualizations/create.
	VisualizationId string
	// Width of a widget
	Width int

	ForceSendFields []string
}

func createWidgetToPb(st *CreateWidget) (*createWidgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWidgetPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	optionsPb, err := widgetOptionsToPb(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	textPb, err := identity(&st.Text)
	if err != nil {
		return nil, err
	}
	if textPb != nil {
		pb.Text = *textPb
	}

	visualizationIdPb, err := identity(&st.VisualizationId)
	if err != nil {
		return nil, err
	}
	if visualizationIdPb != nil {
		pb.VisualizationId = *visualizationIdPb
	}

	widthPb, err := identity(&st.Width)
	if err != nil {
		return nil, err
	}
	if widthPb != nil {
		pb.Width = *widthPb
	}

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
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	optionsField, err := widgetOptionsFromPb(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	textField, err := identity(&pb.Text)
	if err != nil {
		return nil, err
	}
	if textField != nil {
		st.Text = *textField
	}
	visualizationIdField, err := identity(&pb.VisualizationId)
	if err != nil {
		return nil, err
	}
	if visualizationIdField != nil {
		st.VisualizationId = *visualizationIdField
	}
	widthField, err := identity(&pb.Width)
	if err != nil {
		return nil, err
	}
	if widthField != nil {
		st.Width = *widthField
	}

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
	PauseStatus SchedulePauseStatus
	// A cron expression using quartz syntax that specifies the schedule for
	// this pipeline. Should use the quartz format described here:
	// http://www.quartz-scheduler.org/documentation/quartz-2.1.7/tutorials/tutorial-lesson-06.html
	QuartzCronSchedule string
	// A Java timezone id. The schedule will be resolved using this timezone.
	// This will be combined with the quartz_cron_schedule to determine the
	// schedule. See
	// https://docs.databricks.com/sql/language-manual/sql-ref-syntax-aux-conf-mgmt-set-timezone.html
	// for details.
	TimezoneId string

	ForceSendFields []string
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

	quartzCronSchedulePb, err := identity(&st.QuartzCronSchedule)
	if err != nil {
		return nil, err
	}
	if quartzCronSchedulePb != nil {
		pb.QuartzCronSchedule = *quartzCronSchedulePb
	}

	timezoneIdPb, err := identity(&st.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdPb != nil {
		pb.TimezoneId = *timezoneIdPb
	}

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
	pauseStatusField, err := identity(&pb.PauseStatus)
	if err != nil {
		return nil, err
	}
	if pauseStatusField != nil {
		st.PauseStatus = *pauseStatusField
	}
	quartzCronScheduleField, err := identity(&pb.QuartzCronSchedule)
	if err != nil {
		return nil, err
	}
	if quartzCronScheduleField != nil {
		st.QuartzCronSchedule = *quartzCronScheduleField
	}
	timezoneIdField, err := identity(&pb.TimezoneId)
	if err != nil {
		return nil, err
	}
	if timezoneIdField != nil {
		st.TimezoneId = *timezoneIdField
	}

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
	CanEdit bool
	// Timestamp when this dashboard was created.
	CreatedAt string
	// In the web application, query filters that share a name are coupled to a
	// single selection box if this value is `true`.
	DashboardFiltersEnabled bool
	// The ID for this dashboard.
	Id string
	// Indicates whether a dashboard is trashed. Trashed dashboards won't appear
	// in list views. If this boolean is `true`, the `options` property for this
	// dashboard includes a `moved_to_trash_at` timestamp. Items in trash are
	// permanently deleted after 30 days.
	IsArchived bool
	// Whether a dashboard is a draft. Draft dashboards only appear in list
	// views for their owners.
	IsDraft bool
	// Indicates whether this query object appears in the current user's
	// favorites list. This flag determines whether the star icon for favorites
	// is selected.
	IsFavorite bool
	// The title of the dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string

	Options *DashboardOptions
	// The identifier of the workspace folder containing the object.
	Parent string
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel
	// URL slug. Usually mirrors the query name with dashes (`-`) instead of
	// spaces. Appears in the URL for this query.
	Slug string

	Tags []string
	// Timestamp when this dashboard was last updated.
	UpdatedAt string

	User *User
	// The ID of the user who owns the dashboard.
	UserId int

	Widgets []Widget

	ForceSendFields []string
}

func dashboardToPb(st *Dashboard) (*dashboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPb{}
	canEditPb, err := identity(&st.CanEdit)
	if err != nil {
		return nil, err
	}
	if canEditPb != nil {
		pb.CanEdit = *canEditPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	dashboardFiltersEnabledPb, err := identity(&st.DashboardFiltersEnabled)
	if err != nil {
		return nil, err
	}
	if dashboardFiltersEnabledPb != nil {
		pb.DashboardFiltersEnabled = *dashboardFiltersEnabledPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	isArchivedPb, err := identity(&st.IsArchived)
	if err != nil {
		return nil, err
	}
	if isArchivedPb != nil {
		pb.IsArchived = *isArchivedPb
	}

	isDraftPb, err := identity(&st.IsDraft)
	if err != nil {
		return nil, err
	}
	if isDraftPb != nil {
		pb.IsDraft = *isDraftPb
	}

	isFavoritePb, err := identity(&st.IsFavorite)
	if err != nil {
		return nil, err
	}
	if isFavoritePb != nil {
		pb.IsFavorite = *isFavoritePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb, err := dashboardOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	parentPb, err := identity(&st.Parent)
	if err != nil {
		return nil, err
	}
	if parentPb != nil {
		pb.Parent = *parentPb
	}

	permissionTierPb, err := identity(&st.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierPb != nil {
		pb.PermissionTier = *permissionTierPb
	}

	slugPb, err := identity(&st.Slug)
	if err != nil {
		return nil, err
	}
	if slugPb != nil {
		pb.Slug = *slugPb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	userPb, err := userToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}

	userIdPb, err := identity(&st.UserId)
	if err != nil {
		return nil, err
	}
	if userIdPb != nil {
		pb.UserId = *userIdPb
	}

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
	canEditField, err := identity(&pb.CanEdit)
	if err != nil {
		return nil, err
	}
	if canEditField != nil {
		st.CanEdit = *canEditField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	dashboardFiltersEnabledField, err := identity(&pb.DashboardFiltersEnabled)
	if err != nil {
		return nil, err
	}
	if dashboardFiltersEnabledField != nil {
		st.DashboardFiltersEnabled = *dashboardFiltersEnabledField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	isArchivedField, err := identity(&pb.IsArchived)
	if err != nil {
		return nil, err
	}
	if isArchivedField != nil {
		st.IsArchived = *isArchivedField
	}
	isDraftField, err := identity(&pb.IsDraft)
	if err != nil {
		return nil, err
	}
	if isDraftField != nil {
		st.IsDraft = *isDraftField
	}
	isFavoriteField, err := identity(&pb.IsFavorite)
	if err != nil {
		return nil, err
	}
	if isFavoriteField != nil {
		st.IsFavorite = *isFavoriteField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	optionsField, err := dashboardOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	parentField, err := identity(&pb.Parent)
	if err != nil {
		return nil, err
	}
	if parentField != nil {
		st.Parent = *parentField
	}
	permissionTierField, err := identity(&pb.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierField != nil {
		st.PermissionTier = *permissionTierField
	}
	slugField, err := identity(&pb.Slug)
	if err != nil {
		return nil, err
	}
	if slugField != nil {
		st.Slug = *slugField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	userField, err := userFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}
	userIdField, err := identity(&pb.UserId)
	if err != nil {
		return nil, err
	}
	if userIdField != nil {
		st.UserId = *userIdField
	}

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
	DashboardId string
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole

	Tags []string

	ForceSendFields []string
}

func dashboardEditContentToPb(st *DashboardEditContent) (*dashboardEditContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardEditContentPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	runAsRolePb, err := identity(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

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
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	runAsRoleField, err := identity(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

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
	MovedToTrashAt string

	ForceSendFields []string
}

func dashboardOptionsToPb(st *DashboardOptions) (*dashboardOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardOptionsPb{}
	movedToTrashAtPb, err := identity(&st.MovedToTrashAt)
	if err != nil {
		return nil, err
	}
	if movedToTrashAtPb != nil {
		pb.MovedToTrashAt = *movedToTrashAtPb
	}

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
	movedToTrashAtField, err := identity(&pb.MovedToTrashAt)
	if err != nil {
		return nil, err
	}
	if movedToTrashAtField != nil {
		st.MovedToTrashAt = *movedToTrashAtField
	}

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
	DashboardFiltersEnabled bool
	// Indicates whether this dashboard object should appear in the current
	// user's favorites list.
	IsFavorite bool
	// The title of this dashboard that appears in list views and at the top of
	// the dashboard page.
	Name string
	// The identifier of the workspace folder containing the object.
	Parent string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole

	Tags []string

	ForceSendFields []string
}

func dashboardPostContentToPb(st *DashboardPostContent) (*dashboardPostContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardPostContentPb{}
	dashboardFiltersEnabledPb, err := identity(&st.DashboardFiltersEnabled)
	if err != nil {
		return nil, err
	}
	if dashboardFiltersEnabledPb != nil {
		pb.DashboardFiltersEnabled = *dashboardFiltersEnabledPb
	}

	isFavoritePb, err := identity(&st.IsFavorite)
	if err != nil {
		return nil, err
	}
	if isFavoritePb != nil {
		pb.IsFavorite = *isFavoritePb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	parentPb, err := identity(&st.Parent)
	if err != nil {
		return nil, err
	}
	if parentPb != nil {
		pb.Parent = *parentPb
	}

	runAsRolePb, err := identity(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

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
	dashboardFiltersEnabledField, err := identity(&pb.DashboardFiltersEnabled)
	if err != nil {
		return nil, err
	}
	if dashboardFiltersEnabledField != nil {
		st.DashboardFiltersEnabled = *dashboardFiltersEnabledField
	}
	isFavoriteField, err := identity(&pb.IsFavorite)
	if err != nil {
		return nil, err
	}
	if isFavoriteField != nil {
		st.IsFavorite = *isFavoriteField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	parentField, err := identity(&pb.Parent)
	if err != nil {
		return nil, err
	}
	if parentField != nil {
		st.Parent = *parentField
	}
	runAsRoleField, err := identity(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

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
	Id string
	// The string name of this data source / SQL warehouse as it appears in the
	// Databricks SQL web application.
	Name string
	// Reserved for internal use.
	PauseReason string
	// Reserved for internal use.
	Paused int
	// Reserved for internal use.
	SupportsAutoLimit bool
	// Reserved for internal use.
	Syntax string
	// The type of data source. For SQL warehouses, this will be
	// `databricks_internal`.
	Type string
	// Reserved for internal use.
	ViewOnly bool
	// The ID of the associated SQL warehouse, if this data source is backed by
	// a SQL warehouse.
	WarehouseId string

	ForceSendFields []string
}

func dataSourceToPb(st *DataSource) (*dataSourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataSourcePb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pauseReasonPb, err := identity(&st.PauseReason)
	if err != nil {
		return nil, err
	}
	if pauseReasonPb != nil {
		pb.PauseReason = *pauseReasonPb
	}

	pausedPb, err := identity(&st.Paused)
	if err != nil {
		return nil, err
	}
	if pausedPb != nil {
		pb.Paused = *pausedPb
	}

	supportsAutoLimitPb, err := identity(&st.SupportsAutoLimit)
	if err != nil {
		return nil, err
	}
	if supportsAutoLimitPb != nil {
		pb.SupportsAutoLimit = *supportsAutoLimitPb
	}

	syntaxPb, err := identity(&st.Syntax)
	if err != nil {
		return nil, err
	}
	if syntaxPb != nil {
		pb.Syntax = *syntaxPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	viewOnlyPb, err := identity(&st.ViewOnly)
	if err != nil {
		return nil, err
	}
	if viewOnlyPb != nil {
		pb.ViewOnly = *viewOnlyPb
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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	pauseReasonField, err := identity(&pb.PauseReason)
	if err != nil {
		return nil, err
	}
	if pauseReasonField != nil {
		st.PauseReason = *pauseReasonField
	}
	pausedField, err := identity(&pb.Paused)
	if err != nil {
		return nil, err
	}
	if pausedField != nil {
		st.Paused = *pausedField
	}
	supportsAutoLimitField, err := identity(&pb.SupportsAutoLimit)
	if err != nil {
		return nil, err
	}
	if supportsAutoLimitField != nil {
		st.SupportsAutoLimit = *supportsAutoLimitField
	}
	syntaxField, err := identity(&pb.Syntax)
	if err != nil {
		return nil, err
	}
	if syntaxField != nil {
		st.Syntax = *syntaxField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}
	viewOnlyField, err := identity(&pb.ViewOnly)
	if err != nil {
		return nil, err
	}
	if viewOnlyField != nil {
		st.ViewOnly = *viewOnlyField
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
	End string

	Start string
}

func dateRangeToPb(st *DateRange) (*dateRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dateRangePb{}
	endPb, err := identity(&st.End)
	if err != nil {
		return nil, err
	}
	if endPb != nil {
		pb.End = *endPb
	}

	startPb, err := identity(&st.Start)
	if err != nil {
		return nil, err
	}
	if startPb != nil {
		pb.Start = *startPb
	}

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
	endField, err := identity(&pb.End)
	if err != nil {
		return nil, err
	}
	if endField != nil {
		st.End = *endField
	}
	startField, err := identity(&pb.Start)
	if err != nil {
		return nil, err
	}
	if startField != nil {
		st.Start = *startField
	}

	return st, nil
}

type DateRangeValue struct {
	// Manually specified date-time range value.
	DateRangeValue *DateRange
	// Dynamic date-time range value based on current date-time.
	DynamicDateRangeValue DateRangeValueDynamicDateRange
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision DatePrecision

	StartDayOfWeek int

	ForceSendFields []string
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

	dynamicDateRangeValuePb, err := identity(&st.DynamicDateRangeValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateRangeValuePb != nil {
		pb.DynamicDateRangeValue = *dynamicDateRangeValuePb
	}

	precisionPb, err := identity(&st.Precision)
	if err != nil {
		return nil, err
	}
	if precisionPb != nil {
		pb.Precision = *precisionPb
	}

	startDayOfWeekPb, err := identity(&st.StartDayOfWeek)
	if err != nil {
		return nil, err
	}
	if startDayOfWeekPb != nil {
		pb.StartDayOfWeek = *startDayOfWeekPb
	}

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
	dynamicDateRangeValueField, err := identity(&pb.DynamicDateRangeValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateRangeValueField != nil {
		st.DynamicDateRangeValue = *dynamicDateRangeValueField
	}
	precisionField, err := identity(&pb.Precision)
	if err != nil {
		return nil, err
	}
	if precisionField != nil {
		st.Precision = *precisionField
	}
	startDayOfWeekField, err := identity(&pb.StartDayOfWeek)
	if err != nil {
		return nil, err
	}
	if startDayOfWeekField != nil {
		st.StartDayOfWeek = *startDayOfWeekField
	}

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
	DateValue string
	// Dynamic date-time value based on current date-time.
	DynamicDateValue DateValueDynamicDate
	// Date-time precision to format the value into when the query is run.
	// Defaults to DAY_PRECISION (YYYY-MM-DD).
	Precision DatePrecision

	ForceSendFields []string
}

func dateValueToPb(st *DateValue) (*dateValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dateValuePb{}
	dateValuePb, err := identity(&st.DateValue)
	if err != nil {
		return nil, err
	}
	if dateValuePb != nil {
		pb.DateValue = *dateValuePb
	}

	dynamicDateValuePb, err := identity(&st.DynamicDateValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateValuePb != nil {
		pb.DynamicDateValue = *dynamicDateValuePb
	}

	precisionPb, err := identity(&st.Precision)
	if err != nil {
		return nil, err
	}
	if precisionPb != nil {
		pb.Precision = *precisionPb
	}

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
	dateValueField, err := identity(&pb.DateValue)
	if err != nil {
		return nil, err
	}
	if dateValueField != nil {
		st.DateValue = *dateValueField
	}
	dynamicDateValueField, err := identity(&pb.DynamicDateValue)
	if err != nil {
		return nil, err
	}
	if dynamicDateValueField != nil {
		st.DynamicDateValue = *dynamicDateValueField
	}
	precisionField, err := identity(&pb.Precision)
	if err != nil {
		return nil, err
	}
	if precisionField != nil {
		st.Precision = *precisionField
	}

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
	AlertId string
}

func deleteAlertsLegacyRequestToPb(st *DeleteAlertsLegacyRequest) (*deleteAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAlertsLegacyRequestPb{}
	alertIdPb, err := identity(&st.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdPb != nil {
		pb.AlertId = *alertIdPb
	}

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
	alertIdField, err := identity(&pb.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdField != nil {
		st.AlertId = *alertIdField
	}

	return st, nil
}

// Remove a dashboard
type DeleteDashboardRequest struct {
	DashboardId string
}

func deleteDashboardRequestToPb(st *DeleteDashboardRequest) (*deleteDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

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
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

// Remove widget
type DeleteDashboardWidgetRequest struct {
	// Widget ID returned by :method:dashboardwidgets/create
	Id string
}

func deleteDashboardWidgetRequestToPb(st *DeleteDashboardWidgetRequest) (*deleteDashboardWidgetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardWidgetRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

// Delete a query
type DeleteQueriesLegacyRequest struct {
	QueryId string
}

func deleteQueriesLegacyRequestToPb(st *DeleteQueriesLegacyRequest) (*deleteQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQueriesLegacyRequestPb{}
	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

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
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}

	return st, nil
}

// Remove visualization
type DeleteQueryVisualizationsLegacyRequest struct {
	// Widget ID returned by :method:queryvizualisations/create
	Id string
}

func deleteQueryVisualizationsLegacyRequestToPb(st *DeleteQueryVisualizationsLegacyRequest) (*deleteQueryVisualizationsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQueryVisualizationsLegacyRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
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

// Remove a visualization
type DeleteVisualizationRequest struct {
	Id string
}

func deleteVisualizationRequestToPb(st *DeleteVisualizationRequest) (*deleteVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVisualizationRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

// Delete a warehouse
type DeleteWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id string
}

func deleteWarehouseRequestToPb(st *DeleteWarehouseRequest) (*deleteWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWarehouseRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

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
	AlertId string
	// Name of the alert.
	Name string
	// Alert configuration options.
	Options AlertOptions
	// Query ID.
	QueryId string
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int

	ForceSendFields []string
}

func editAlertToPb(st *EditAlert) (*editAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editAlertPb{}
	alertIdPb, err := identity(&st.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdPb != nil {
		pb.AlertId = *alertIdPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb, err := alertOptionsToPb(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	rearmPb, err := identity(&st.Rearm)
	if err != nil {
		return nil, err
	}
	if rearmPb != nil {
		pb.Rearm = *rearmPb
	}

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
	alertIdField, err := identity(&pb.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdField != nil {
		st.AlertId = *alertIdField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	optionsField, err := alertOptionsFromPb(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	rearmField, err := identity(&pb.Rearm)
	if err != nil {
		return nil, err
	}
	if rearmField != nil {
		st.Rearm = *rearmField
	}

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
	AutoStopMins int
	// Channel Details
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string
	// warehouse creator name
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute.
	EnableServerlessCompute bool
	// Required. Id of the warehouse to configure.
	Id string
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType EditWarehouseRequestWarehouseType

	ForceSendFields []string
}

func editWarehouseRequestToPb(st *EditWarehouseRequest) (*editWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editWarehouseRequestPb{}
	autoStopMinsPb, err := identity(&st.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsPb != nil {
		pb.AutoStopMins = *autoStopMinsPb
	}

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	clusterSizePb, err := identity(&st.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizePb != nil {
		pb.ClusterSize = *clusterSizePb
	}

	creatorNamePb, err := identity(&st.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNamePb != nil {
		pb.CreatorName = *creatorNamePb
	}

	enablePhotonPb, err := identity(&st.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonPb != nil {
		pb.EnablePhoton = *enablePhotonPb
	}

	enableServerlessComputePb, err := identity(&st.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputePb != nil {
		pb.EnableServerlessCompute = *enableServerlessComputePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	maxNumClustersPb, err := identity(&st.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersPb != nil {
		pb.MaxNumClusters = *maxNumClustersPb
	}

	minNumClustersPb, err := identity(&st.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersPb != nil {
		pb.MinNumClusters = *minNumClustersPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	spotInstancePolicyPb, err := identity(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	warehouseTypePb, err := identity(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

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
	autoStopMinsField, err := identity(&pb.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsField != nil {
		st.AutoStopMins = *autoStopMinsField
	}
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	clusterSizeField, err := identity(&pb.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizeField != nil {
		st.ClusterSize = *clusterSizeField
	}
	creatorNameField, err := identity(&pb.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNameField != nil {
		st.CreatorName = *creatorNameField
	}
	enablePhotonField, err := identity(&pb.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonField != nil {
		st.EnablePhoton = *enablePhotonField
	}
	enableServerlessComputeField, err := identity(&pb.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputeField != nil {
		st.EnableServerlessCompute = *enableServerlessComputeField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	maxNumClustersField, err := identity(&pb.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersField != nil {
		st.MaxNumClusters = *maxNumClustersField
	}
	minNumClustersField, err := identity(&pb.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersField != nil {
		st.MinNumClusters = *minNumClustersField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	spotInstancePolicyField, err := identity(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := identity(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

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
	Key string

	Value string

	ForceSendFields []string
}

func endpointConfPairToPb(st *EndpointConfPair) (*endpointConfPairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointConfPairPb{}
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

func (st *endpointConfPairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointConfPairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointHealth struct {
	// Details about errors that are causing current degraded/failed status.
	Details string
	// The reason for failure to bring up clusters for this warehouse. This is
	// available when status is 'FAILED' and sometimes when it is DEGRADED.
	FailureReason *TerminationReason
	// Deprecated. split into summary and details for security
	Message string
	// Health status of the warehouse.
	Status Status
	// A short summary of the health status in case of degraded/failed
	// warehouses.
	Summary string

	ForceSendFields []string
}

func endpointHealthToPb(st *EndpointHealth) (*endpointHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointHealthPb{}
	detailsPb, err := identity(&st.Details)
	if err != nil {
		return nil, err
	}
	if detailsPb != nil {
		pb.Details = *detailsPb
	}

	failureReasonPb, err := terminationReasonToPb(st.FailureReason)
	if err != nil {
		return nil, err
	}
	if failureReasonPb != nil {
		pb.FailureReason = failureReasonPb
	}

	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	summaryPb, err := identity(&st.Summary)
	if err != nil {
		return nil, err
	}
	if summaryPb != nil {
		pb.Summary = *summaryPb
	}

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
	detailsField, err := identity(&pb.Details)
	if err != nil {
		return nil, err
	}
	if detailsField != nil {
		st.Details = *detailsField
	}
	failureReasonField, err := terminationReasonFromPb(pb.FailureReason)
	if err != nil {
		return nil, err
	}
	if failureReasonField != nil {
		st.FailureReason = failureReasonField
	}
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	summaryField, err := identity(&pb.Summary)
	if err != nil {
		return nil, err
	}
	if summaryField != nil {
		st.Summary = *summaryField
	}

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
	AutoStopMins int
	// Channel Details
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string
	// warehouse creator name
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *EndpointHealth
	// unique identifier for warehouse
	Id string
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string
	// the jdbc connection string for this warehouse
	JdbcUrl string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions int64
	// current number of clusters running for the service
	NumClusters int
	// ODBC parameters for the SQL warehouse
	OdbcParams *OdbcParams
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy
	// State of the warehouse
	State State
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType EndpointInfoWarehouseType

	ForceSendFields []string
}

func endpointInfoToPb(st *EndpointInfo) (*endpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointInfoPb{}
	autoStopMinsPb, err := identity(&st.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsPb != nil {
		pb.AutoStopMins = *autoStopMinsPb
	}

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	clusterSizePb, err := identity(&st.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizePb != nil {
		pb.ClusterSize = *clusterSizePb
	}

	creatorNamePb, err := identity(&st.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNamePb != nil {
		pb.CreatorName = *creatorNamePb
	}

	enablePhotonPb, err := identity(&st.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonPb != nil {
		pb.EnablePhoton = *enablePhotonPb
	}

	enableServerlessComputePb, err := identity(&st.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputePb != nil {
		pb.EnableServerlessCompute = *enableServerlessComputePb
	}

	healthPb, err := endpointHealthToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	jdbcUrlPb, err := identity(&st.JdbcUrl)
	if err != nil {
		return nil, err
	}
	if jdbcUrlPb != nil {
		pb.JdbcUrl = *jdbcUrlPb
	}

	maxNumClustersPb, err := identity(&st.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersPb != nil {
		pb.MaxNumClusters = *maxNumClustersPb
	}

	minNumClustersPb, err := identity(&st.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersPb != nil {
		pb.MinNumClusters = *minNumClustersPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	numActiveSessionsPb, err := identity(&st.NumActiveSessions)
	if err != nil {
		return nil, err
	}
	if numActiveSessionsPb != nil {
		pb.NumActiveSessions = *numActiveSessionsPb
	}

	numClustersPb, err := identity(&st.NumClusters)
	if err != nil {
		return nil, err
	}
	if numClustersPb != nil {
		pb.NumClusters = *numClustersPb
	}

	odbcParamsPb, err := odbcParamsToPb(st.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsPb != nil {
		pb.OdbcParams = odbcParamsPb
	}

	spotInstancePolicyPb, err := identity(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	warehouseTypePb, err := identity(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

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
	autoStopMinsField, err := identity(&pb.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsField != nil {
		st.AutoStopMins = *autoStopMinsField
	}
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	clusterSizeField, err := identity(&pb.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizeField != nil {
		st.ClusterSize = *clusterSizeField
	}
	creatorNameField, err := identity(&pb.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNameField != nil {
		st.CreatorName = *creatorNameField
	}
	enablePhotonField, err := identity(&pb.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonField != nil {
		st.EnablePhoton = *enablePhotonField
	}
	enableServerlessComputeField, err := identity(&pb.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputeField != nil {
		st.EnableServerlessCompute = *enableServerlessComputeField
	}
	healthField, err := endpointHealthFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	jdbcUrlField, err := identity(&pb.JdbcUrl)
	if err != nil {
		return nil, err
	}
	if jdbcUrlField != nil {
		st.JdbcUrl = *jdbcUrlField
	}
	maxNumClustersField, err := identity(&pb.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersField != nil {
		st.MaxNumClusters = *maxNumClustersField
	}
	minNumClustersField, err := identity(&pb.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersField != nil {
		st.MinNumClusters = *minNumClustersField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	numActiveSessionsField, err := identity(&pb.NumActiveSessions)
	if err != nil {
		return nil, err
	}
	if numActiveSessionsField != nil {
		st.NumActiveSessions = *numActiveSessionsField
	}
	numClustersField, err := identity(&pb.NumClusters)
	if err != nil {
		return nil, err
	}
	if numClustersField != nil {
		st.NumClusters = *numClustersField
	}
	odbcParamsField, err := odbcParamsFromPb(pb.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsField != nil {
		st.OdbcParams = odbcParamsField
	}
	spotInstancePolicyField, err := identity(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := identity(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

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
	Key string

	Value string

	ForceSendFields []string
}

func endpointTagPairToPb(st *EndpointTagPair) (*endpointTagPairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagPairPb{}
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

func (st *endpointTagPairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointTagPairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointTags struct {
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
	EnumOptions string
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions *MultiValuesOptions
	// List of selected query parameter values.
	Values []string

	ForceSendFields []string
}

func enumValueToPb(st *EnumValue) (*enumValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enumValuePb{}
	enumOptionsPb, err := identity(&st.EnumOptions)
	if err != nil {
		return nil, err
	}
	if enumOptionsPb != nil {
		pb.EnumOptions = *enumOptionsPb
	}

	multiValuesOptionsPb, err := multiValuesOptionsToPb(st.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsPb != nil {
		pb.MultiValuesOptions = multiValuesOptionsPb
	}

	var valuesPb []string
	for _, item := range st.Values {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			valuesPb = append(valuesPb, *itemPb)
		}
	}
	pb.Values = valuesPb

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
	enumOptionsField, err := identity(&pb.EnumOptions)
	if err != nil {
		return nil, err
	}
	if enumOptionsField != nil {
		st.EnumOptions = *enumOptionsField
	}
	multiValuesOptionsField, err := multiValuesOptionsFromPb(pb.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsField != nil {
		st.MultiValuesOptions = multiValuesOptionsField
	}

	var valuesField []string
	for _, itemPb := range pb.Values {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			valuesField = append(valuesField, *item)
		}
	}
	st.Values = valuesField

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
	ByteLimit int64
	// Sets default catalog for statement execution, similar to [`USE CATALOG`]
	// in SQL.
	//
	// [`USE CATALOG`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-catalog.html
	Catalog string

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
	Format Format
	// When `wait_timeout > 0s`, the call will block up to the specified time.
	// If the statement execution doesn't finish within this time,
	// `on_wait_timeout` determines whether the execution should continue or be
	// canceled. When set to `CONTINUE`, the statement execution continues
	// asynchronously and the call returns a statement ID which can be used for
	// polling with :method:statementexecution/getStatement. When set to
	// `CANCEL`, the statement execution is canceled and the call returns with a
	// `CANCELED` state.
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
	Parameters []StatementParameterListItem
	// Applies the given row limit to the statement's result set, but unlike the
	// `LIMIT` clause in SQL, it also sets the `truncated` field in the response
	// to indicate whether the result was trimmed due to the limit or not.
	RowLimit int64
	// Sets default schema for statement execution, similar to [`USE SCHEMA`] in
	// SQL.
	//
	// [`USE SCHEMA`]: https://docs.databricks.com/sql/language-manual/sql-ref-syntax-ddl-use-schema.html
	Schema string
	// The SQL statement to execute. The statement can optionally be
	// parameterized, see `parameters`.
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
	WaitTimeout string
	// Warehouse upon which to execute a statement. See also [What are SQL
	// warehouses?]
	//
	// [What are SQL warehouses?]: https://docs.databricks.com/sql/admin/warehouse-type.html
	WarehouseId string

	ForceSendFields []string
}

func executeStatementRequestToPb(st *ExecuteStatementRequest) (*executeStatementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &executeStatementRequestPb{}
	byteLimitPb, err := identity(&st.ByteLimit)
	if err != nil {
		return nil, err
	}
	if byteLimitPb != nil {
		pb.ByteLimit = *byteLimitPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	dispositionPb, err := identity(&st.Disposition)
	if err != nil {
		return nil, err
	}
	if dispositionPb != nil {
		pb.Disposition = *dispositionPb
	}

	formatPb, err := identity(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}

	onWaitTimeoutPb, err := identity(&st.OnWaitTimeout)
	if err != nil {
		return nil, err
	}
	if onWaitTimeoutPb != nil {
		pb.OnWaitTimeout = *onWaitTimeoutPb
	}

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

	rowLimitPb, err := identity(&st.RowLimit)
	if err != nil {
		return nil, err
	}
	if rowLimitPb != nil {
		pb.RowLimit = *rowLimitPb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	statementPb, err := identity(&st.Statement)
	if err != nil {
		return nil, err
	}
	if statementPb != nil {
		pb.Statement = *statementPb
	}

	waitTimeoutPb, err := identity(&st.WaitTimeout)
	if err != nil {
		return nil, err
	}
	if waitTimeoutPb != nil {
		pb.WaitTimeout = *waitTimeoutPb
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
	byteLimitField, err := identity(&pb.ByteLimit)
	if err != nil {
		return nil, err
	}
	if byteLimitField != nil {
		st.ByteLimit = *byteLimitField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	dispositionField, err := identity(&pb.Disposition)
	if err != nil {
		return nil, err
	}
	if dispositionField != nil {
		st.Disposition = *dispositionField
	}
	formatField, err := identity(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	onWaitTimeoutField, err := identity(&pb.OnWaitTimeout)
	if err != nil {
		return nil, err
	}
	if onWaitTimeoutField != nil {
		st.OnWaitTimeout = *onWaitTimeoutField
	}

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
	rowLimitField, err := identity(&pb.RowLimit)
	if err != nil {
		return nil, err
	}
	if rowLimitField != nil {
		st.RowLimit = *rowLimitField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}
	statementField, err := identity(&pb.Statement)
	if err != nil {
		return nil, err
	}
	if statementField != nil {
		st.Statement = *statementField
	}
	waitTimeoutField, err := identity(&pb.WaitTimeout)
	if err != nil {
		return nil, err
	}
	if waitTimeoutField != nil {
		st.WaitTimeout = *waitTimeoutField
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
	ByteCount int64
	// The position within the sequence of result set chunks.
	ChunkIndex int
	// Indicates the date-time that the given external link will expire and
	// becomes invalid, after which point a new `external_link` must be
	// requested.
	Expiration string

	ExternalLink string
	// HTTP headers that must be included with a GET request to the
	// `external_link`. Each header is provided as a key-value pair. Headers are
	// typically used to pass a decryption key to the external service. The
	// values of these headers should be considered sensitive and the client
	// should not expose these values in a log.
	HttpHeaders map[string]string
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex int
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink string
	// The number of rows within the result chunk.
	RowCount int64
	// The starting row offset within the result set.
	RowOffset int64

	ForceSendFields []string
}

func externalLinkToPb(st *ExternalLink) (*externalLinkPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalLinkPb{}
	byteCountPb, err := identity(&st.ByteCount)
	if err != nil {
		return nil, err
	}
	if byteCountPb != nil {
		pb.ByteCount = *byteCountPb
	}

	chunkIndexPb, err := identity(&st.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexPb != nil {
		pb.ChunkIndex = *chunkIndexPb
	}

	expirationPb, err := identity(&st.Expiration)
	if err != nil {
		return nil, err
	}
	if expirationPb != nil {
		pb.Expiration = *expirationPb
	}

	externalLinkPb, err := identity(&st.ExternalLink)
	if err != nil {
		return nil, err
	}
	if externalLinkPb != nil {
		pb.ExternalLink = *externalLinkPb
	}

	httpHeadersPb := map[string]string{}
	for k, v := range st.HttpHeaders {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			httpHeadersPb[k] = *itemPb
		}
	}
	pb.HttpHeaders = httpHeadersPb

	nextChunkIndexPb, err := identity(&st.NextChunkIndex)
	if err != nil {
		return nil, err
	}
	if nextChunkIndexPb != nil {
		pb.NextChunkIndex = *nextChunkIndexPb
	}

	nextChunkInternalLinkPb, err := identity(&st.NextChunkInternalLink)
	if err != nil {
		return nil, err
	}
	if nextChunkInternalLinkPb != nil {
		pb.NextChunkInternalLink = *nextChunkInternalLinkPb
	}

	rowCountPb, err := identity(&st.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountPb != nil {
		pb.RowCount = *rowCountPb
	}

	rowOffsetPb, err := identity(&st.RowOffset)
	if err != nil {
		return nil, err
	}
	if rowOffsetPb != nil {
		pb.RowOffset = *rowOffsetPb
	}

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
	byteCountField, err := identity(&pb.ByteCount)
	if err != nil {
		return nil, err
	}
	if byteCountField != nil {
		st.ByteCount = *byteCountField
	}
	chunkIndexField, err := identity(&pb.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexField != nil {
		st.ChunkIndex = *chunkIndexField
	}
	expirationField, err := identity(&pb.Expiration)
	if err != nil {
		return nil, err
	}
	if expirationField != nil {
		st.Expiration = *expirationField
	}
	externalLinkField, err := identity(&pb.ExternalLink)
	if err != nil {
		return nil, err
	}
	if externalLinkField != nil {
		st.ExternalLink = *externalLinkField
	}

	httpHeadersField := map[string]string{}
	for k, v := range pb.HttpHeaders {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			httpHeadersField[k] = *item
		}
	}
	st.HttpHeaders = httpHeadersField
	nextChunkIndexField, err := identity(&pb.NextChunkIndex)
	if err != nil {
		return nil, err
	}
	if nextChunkIndexField != nil {
		st.NextChunkIndex = *nextChunkIndexField
	}
	nextChunkInternalLinkField, err := identity(&pb.NextChunkInternalLink)
	if err != nil {
		return nil, err
	}
	if nextChunkInternalLinkField != nil {
		st.NextChunkInternalLink = *nextChunkInternalLinkField
	}
	rowCountField, err := identity(&pb.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountField != nil {
		st.RowCount = *rowCountField
	}
	rowOffsetField, err := identity(&pb.RowOffset)
	if err != nil {
		return nil, err
	}
	if rowOffsetField != nil {
		st.RowOffset = *rowOffsetField
	}

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
	AlertId string
	// The canonical identifier for this Lakeview dashboard
	DashboardId string
	// The canonical identifier for this Genie space
	GenieSpaceId string

	JobInfo *ExternalQuerySourceJobInfo
	// The canonical identifier for this legacy dashboard
	LegacyDashboardId string
	// The canonical identifier for this notebook
	NotebookId string
	// The canonical identifier for this SQL query
	SqlQueryId string

	ForceSendFields []string
}

func externalQuerySourceToPb(st *ExternalQuerySource) (*externalQuerySourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalQuerySourcePb{}
	alertIdPb, err := identity(&st.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdPb != nil {
		pb.AlertId = *alertIdPb
	}

	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

	genieSpaceIdPb, err := identity(&st.GenieSpaceId)
	if err != nil {
		return nil, err
	}
	if genieSpaceIdPb != nil {
		pb.GenieSpaceId = *genieSpaceIdPb
	}

	jobInfoPb, err := externalQuerySourceJobInfoToPb(st.JobInfo)
	if err != nil {
		return nil, err
	}
	if jobInfoPb != nil {
		pb.JobInfo = jobInfoPb
	}

	legacyDashboardIdPb, err := identity(&st.LegacyDashboardId)
	if err != nil {
		return nil, err
	}
	if legacyDashboardIdPb != nil {
		pb.LegacyDashboardId = *legacyDashboardIdPb
	}

	notebookIdPb, err := identity(&st.NotebookId)
	if err != nil {
		return nil, err
	}
	if notebookIdPb != nil {
		pb.NotebookId = *notebookIdPb
	}

	sqlQueryIdPb, err := identity(&st.SqlQueryId)
	if err != nil {
		return nil, err
	}
	if sqlQueryIdPb != nil {
		pb.SqlQueryId = *sqlQueryIdPb
	}

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
	alertIdField, err := identity(&pb.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdField != nil {
		st.AlertId = *alertIdField
	}
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}
	genieSpaceIdField, err := identity(&pb.GenieSpaceId)
	if err != nil {
		return nil, err
	}
	if genieSpaceIdField != nil {
		st.GenieSpaceId = *genieSpaceIdField
	}
	jobInfoField, err := externalQuerySourceJobInfoFromPb(pb.JobInfo)
	if err != nil {
		return nil, err
	}
	if jobInfoField != nil {
		st.JobInfo = jobInfoField
	}
	legacyDashboardIdField, err := identity(&pb.LegacyDashboardId)
	if err != nil {
		return nil, err
	}
	if legacyDashboardIdField != nil {
		st.LegacyDashboardId = *legacyDashboardIdField
	}
	notebookIdField, err := identity(&pb.NotebookId)
	if err != nil {
		return nil, err
	}
	if notebookIdField != nil {
		st.NotebookId = *notebookIdField
	}
	sqlQueryIdField, err := identity(&pb.SqlQueryId)
	if err != nil {
		return nil, err
	}
	if sqlQueryIdField != nil {
		st.SqlQueryId = *sqlQueryIdField
	}

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
	JobId string
	// The canonical identifier of the run. This ID is unique across all runs of
	// all jobs.
	JobRunId string
	// The canonical identifier of the task run.
	JobTaskRunId string

	ForceSendFields []string
}

func externalQuerySourceJobInfoToPb(st *ExternalQuerySourceJobInfo) (*externalQuerySourceJobInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalQuerySourceJobInfoPb{}
	jobIdPb, err := identity(&st.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdPb != nil {
		pb.JobId = *jobIdPb
	}

	jobRunIdPb, err := identity(&st.JobRunId)
	if err != nil {
		return nil, err
	}
	if jobRunIdPb != nil {
		pb.JobRunId = *jobRunIdPb
	}

	jobTaskRunIdPb, err := identity(&st.JobTaskRunId)
	if err != nil {
		return nil, err
	}
	if jobTaskRunIdPb != nil {
		pb.JobTaskRunId = *jobTaskRunIdPb
	}

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
	jobIdField, err := identity(&pb.JobId)
	if err != nil {
		return nil, err
	}
	if jobIdField != nil {
		st.JobId = *jobIdField
	}
	jobRunIdField, err := identity(&pb.JobRunId)
	if err != nil {
		return nil, err
	}
	if jobRunIdField != nil {
		st.JobRunId = *jobRunIdField
	}
	jobTaskRunIdField, err := identity(&pb.JobTaskRunId)
	if err != nil {
		return nil, err
	}
	if jobTaskRunIdField != nil {
		st.JobTaskRunId = *jobTaskRunIdField
	}

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
	Id string
}

func getAlertRequestToPb(st *GetAlertRequest) (*getAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

// Get an alert
type GetAlertV2Request struct {
	Id string
}

func getAlertV2RequestToPb(st *GetAlertV2Request) (*getAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertV2RequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

// Get an alert
type GetAlertsLegacyRequest struct {
	AlertId string
}

func getAlertsLegacyRequestToPb(st *GetAlertsLegacyRequest) (*getAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertsLegacyRequestPb{}
	alertIdPb, err := identity(&st.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdPb != nil {
		pb.AlertId = *alertIdPb
	}

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
	alertIdField, err := identity(&pb.AlertId)
	if err != nil {
		return nil, err
	}
	if alertIdField != nil {
		st.AlertId = *alertIdField
	}

	return st, nil
}

// Retrieve a definition
type GetDashboardRequest struct {
	DashboardId string
}

func getDashboardRequestToPb(st *GetDashboardRequest) (*getDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDashboardRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

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
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

// Get object ACL
type GetDbsqlPermissionRequest struct {
	// Object ID. An ACL is returned for the object with this UUID.
	ObjectId string
	// The type of object permissions to check.
	ObjectType ObjectTypePlural
}

func getDbsqlPermissionRequestToPb(st *GetDbsqlPermissionRequest) (*getDbsqlPermissionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDbsqlPermissionRequestPb{}
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

	return st, nil
}

// Get a query definition.
type GetQueriesLegacyRequest struct {
	QueryId string
}

func getQueriesLegacyRequestToPb(st *GetQueriesLegacyRequest) (*getQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQueriesLegacyRequestPb{}
	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

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
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}

	return st, nil
}

// Get a query
type GetQueryRequest struct {
	Id string
}

func getQueryRequestToPb(st *GetQueryRequest) (*getQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQueryRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

type GetResponse struct {
	AccessControlList []AccessControl
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId string
	// A singular noun object type.
	ObjectType ObjectType

	ForceSendFields []string
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
	StatementId string
}

func getStatementRequestToPb(st *GetStatementRequest) (*getStatementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatementRequestPb{}
	statementIdPb, err := identity(&st.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdPb != nil {
		pb.StatementId = *statementIdPb
	}

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
	statementIdField, err := identity(&pb.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdField != nil {
		st.StatementId = *statementIdField
	}

	return st, nil
}

// Get result chunk by index
type GetStatementResultChunkNRequest struct {
	ChunkIndex int
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string
}

func getStatementResultChunkNRequestToPb(st *GetStatementResultChunkNRequest) (*getStatementResultChunkNRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatementResultChunkNRequestPb{}
	chunkIndexPb, err := identity(&st.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexPb != nil {
		pb.ChunkIndex = *chunkIndexPb
	}

	statementIdPb, err := identity(&st.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdPb != nil {
		pb.StatementId = *statementIdPb
	}

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
	chunkIndexField, err := identity(&pb.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexField != nil {
		st.ChunkIndex = *chunkIndexField
	}
	statementIdField, err := identity(&pb.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdField != nil {
		st.StatementId = *statementIdField
	}

	return st, nil
}

// Get SQL warehouse permission levels
type GetWarehousePermissionLevelsRequest struct {
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string
}

func getWarehousePermissionLevelsRequestToPb(st *GetWarehousePermissionLevelsRequest) (*getWarehousePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionLevelsRequestPb{}
	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

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
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	return st, nil
}

type GetWarehousePermissionLevelsResponse struct {
	// Specific permission levels
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
	WarehouseId string
}

func getWarehousePermissionsRequestToPb(st *GetWarehousePermissionsRequest) (*getWarehousePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionsRequestPb{}
	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

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
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	return st, nil
}

// Get warehouse info
type GetWarehouseRequest struct {
	// Required. Id of the SQL warehouse.
	Id string
}

func getWarehouseRequestToPb(st *GetWarehouseRequest) (*getWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehouseRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

type GetWarehouseResponse struct {
	// The amount of time in minutes that a SQL warehouse must be idle (i.e., no
	// RUNNING queries) before it is automatically stopped.
	//
	// Supported values: - Must be == 0 or >= 10 mins - 0 indicates no autostop.
	//
	// Defaults to 120 mins
	AutoStopMins int
	// Channel Details
	Channel *Channel
	// Size of the clusters allocated for this warehouse. Increasing the size of
	// a spark cluster allows you to run larger queries on it. If you want to
	// increase the number of concurrent queries, please tune max_num_clusters.
	//
	// Supported values: - 2X-Small - X-Small - Small - Medium - Large - X-Large
	// - 2X-Large - 3X-Large - 4X-Large
	ClusterSize string
	// warehouse creator name
	CreatorName string
	// Configures whether the warehouse should use Photon optimized clusters.
	//
	// Defaults to false.
	EnablePhoton bool
	// Configures whether the warehouse should use serverless compute
	EnableServerlessCompute bool
	// Optional health status. Assume the warehouse is healthy if this field is
	// not set.
	Health *EndpointHealth
	// unique identifier for warehouse
	Id string
	// Deprecated. Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string
	// the jdbc connection string for this warehouse
	JdbcUrl string
	// Maximum number of clusters that the autoscaler will create to handle
	// concurrent queries.
	//
	// Supported values: - Must be >= min_num_clusters - Must be <= 30.
	//
	// Defaults to min_clusters if unset.
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
	MinNumClusters int
	// Logical name for the cluster.
	//
	// Supported values: - Must be unique within an org. - Must be less than 100
	// characters.
	Name string
	// Deprecated. current number of active sessions for the warehouse
	NumActiveSessions int64
	// current number of clusters running for the service
	NumClusters int
	// ODBC parameters for the SQL warehouse
	OdbcParams *OdbcParams
	// Configurations whether the warehouse should use spot instances.
	SpotInstancePolicy SpotInstancePolicy
	// State of the warehouse
	State State
	// A set of key-value pairs that will be tagged on all resources (e.g., AWS
	// instances and EBS volumes) associated with this SQL warehouse.
	//
	// Supported values: - Number of tags < 45.
	Tags *EndpointTags
	// Warehouse type: `PRO` or `CLASSIC`. If you want to use serverless
	// compute, you must set to `PRO` and also set the field
	// `enable_serverless_compute` to `true`.
	WarehouseType GetWarehouseResponseWarehouseType

	ForceSendFields []string
}

func getWarehouseResponseToPb(st *GetWarehouseResponse) (*getWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehouseResponsePb{}
	autoStopMinsPb, err := identity(&st.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsPb != nil {
		pb.AutoStopMins = *autoStopMinsPb
	}

	channelPb, err := channelToPb(st.Channel)
	if err != nil {
		return nil, err
	}
	if channelPb != nil {
		pb.Channel = channelPb
	}

	clusterSizePb, err := identity(&st.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizePb != nil {
		pb.ClusterSize = *clusterSizePb
	}

	creatorNamePb, err := identity(&st.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNamePb != nil {
		pb.CreatorName = *creatorNamePb
	}

	enablePhotonPb, err := identity(&st.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonPb != nil {
		pb.EnablePhoton = *enablePhotonPb
	}

	enableServerlessComputePb, err := identity(&st.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputePb != nil {
		pb.EnableServerlessCompute = *enableServerlessComputePb
	}

	healthPb, err := endpointHealthToPb(st.Health)
	if err != nil {
		return nil, err
	}
	if healthPb != nil {
		pb.Health = healthPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	jdbcUrlPb, err := identity(&st.JdbcUrl)
	if err != nil {
		return nil, err
	}
	if jdbcUrlPb != nil {
		pb.JdbcUrl = *jdbcUrlPb
	}

	maxNumClustersPb, err := identity(&st.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersPb != nil {
		pb.MaxNumClusters = *maxNumClustersPb
	}

	minNumClustersPb, err := identity(&st.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersPb != nil {
		pb.MinNumClusters = *minNumClustersPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	numActiveSessionsPb, err := identity(&st.NumActiveSessions)
	if err != nil {
		return nil, err
	}
	if numActiveSessionsPb != nil {
		pb.NumActiveSessions = *numActiveSessionsPb
	}

	numClustersPb, err := identity(&st.NumClusters)
	if err != nil {
		return nil, err
	}
	if numClustersPb != nil {
		pb.NumClusters = *numClustersPb
	}

	odbcParamsPb, err := odbcParamsToPb(st.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsPb != nil {
		pb.OdbcParams = odbcParamsPb
	}

	spotInstancePolicyPb, err := identity(&st.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyPb != nil {
		pb.SpotInstancePolicy = *spotInstancePolicyPb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	tagsPb, err := endpointTagsToPb(st.Tags)
	if err != nil {
		return nil, err
	}
	if tagsPb != nil {
		pb.Tags = tagsPb
	}

	warehouseTypePb, err := identity(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

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
	autoStopMinsField, err := identity(&pb.AutoStopMins)
	if err != nil {
		return nil, err
	}
	if autoStopMinsField != nil {
		st.AutoStopMins = *autoStopMinsField
	}
	channelField, err := channelFromPb(pb.Channel)
	if err != nil {
		return nil, err
	}
	if channelField != nil {
		st.Channel = channelField
	}
	clusterSizeField, err := identity(&pb.ClusterSize)
	if err != nil {
		return nil, err
	}
	if clusterSizeField != nil {
		st.ClusterSize = *clusterSizeField
	}
	creatorNameField, err := identity(&pb.CreatorName)
	if err != nil {
		return nil, err
	}
	if creatorNameField != nil {
		st.CreatorName = *creatorNameField
	}
	enablePhotonField, err := identity(&pb.EnablePhoton)
	if err != nil {
		return nil, err
	}
	if enablePhotonField != nil {
		st.EnablePhoton = *enablePhotonField
	}
	enableServerlessComputeField, err := identity(&pb.EnableServerlessCompute)
	if err != nil {
		return nil, err
	}
	if enableServerlessComputeField != nil {
		st.EnableServerlessCompute = *enableServerlessComputeField
	}
	healthField, err := endpointHealthFromPb(pb.Health)
	if err != nil {
		return nil, err
	}
	if healthField != nil {
		st.Health = healthField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	jdbcUrlField, err := identity(&pb.JdbcUrl)
	if err != nil {
		return nil, err
	}
	if jdbcUrlField != nil {
		st.JdbcUrl = *jdbcUrlField
	}
	maxNumClustersField, err := identity(&pb.MaxNumClusters)
	if err != nil {
		return nil, err
	}
	if maxNumClustersField != nil {
		st.MaxNumClusters = *maxNumClustersField
	}
	minNumClustersField, err := identity(&pb.MinNumClusters)
	if err != nil {
		return nil, err
	}
	if minNumClustersField != nil {
		st.MinNumClusters = *minNumClustersField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	numActiveSessionsField, err := identity(&pb.NumActiveSessions)
	if err != nil {
		return nil, err
	}
	if numActiveSessionsField != nil {
		st.NumActiveSessions = *numActiveSessionsField
	}
	numClustersField, err := identity(&pb.NumClusters)
	if err != nil {
		return nil, err
	}
	if numClustersField != nil {
		st.NumClusters = *numClustersField
	}
	odbcParamsField, err := odbcParamsFromPb(pb.OdbcParams)
	if err != nil {
		return nil, err
	}
	if odbcParamsField != nil {
		st.OdbcParams = odbcParamsField
	}
	spotInstancePolicyField, err := identity(&pb.SpotInstancePolicy)
	if err != nil {
		return nil, err
	}
	if spotInstancePolicyField != nil {
		st.SpotInstancePolicy = *spotInstancePolicyField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	tagsField, err := endpointTagsFromPb(pb.Tags)
	if err != nil {
		return nil, err
	}
	if tagsField != nil {
		st.Tags = tagsField
	}
	warehouseTypeField, err := identity(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

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
	Channel *Channel
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig []EndpointConfPair
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes []WarehouseTypePair
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount string
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string
	// Security policy for warehouses
	SecurityPolicy GetWorkspaceWarehouseConfigResponseSecurityPolicy
	// SQL configuration parameters
	SqlConfigurationParameters *RepeatedEndpointConfPairs

	ForceSendFields []string
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

	googleServiceAccountPb, err := identity(&st.GoogleServiceAccount)
	if err != nil {
		return nil, err
	}
	if googleServiceAccountPb != nil {
		pb.GoogleServiceAccount = *googleServiceAccountPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	securityPolicyPb, err := identity(&st.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyPb != nil {
		pb.SecurityPolicy = *securityPolicyPb
	}

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
	googleServiceAccountField, err := identity(&pb.GoogleServiceAccount)
	if err != nil {
		return nil, err
	}
	if googleServiceAccountField != nil {
		st.GoogleServiceAccount = *googleServiceAccountField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	securityPolicyField, err := identity(&pb.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyField != nil {
		st.SecurityPolicy = *securityPolicyField
	}
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
	CreatedAt string
	// Alert ID.
	Id string
	// Timestamp when the alert was last triggered.
	LastTriggeredAt string
	// Name of the alert.
	Name string
	// Alert configuration options.
	Options *AlertOptions
	// The identifier of the workspace folder containing the object.
	Parent string

	Query *AlertQuery
	// Number of seconds after being triggered before the alert rearms itself
	// and can be triggered again. If `null`, alert will never be triggered
	// again.
	Rearm int
	// State of the alert. Possible values are: `unknown` (yet to be evaluated),
	// `triggered` (evaluated and fulfilled trigger conditions), or `ok`
	// (evaluated and did not fulfill trigger conditions).
	State LegacyAlertState
	// Timestamp when the alert was last updated.
	UpdatedAt string

	User *User

	ForceSendFields []string
}

func legacyAlertToPb(st *LegacyAlert) (*legacyAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &legacyAlertPb{}
	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastTriggeredAtPb, err := identity(&st.LastTriggeredAt)
	if err != nil {
		return nil, err
	}
	if lastTriggeredAtPb != nil {
		pb.LastTriggeredAt = *lastTriggeredAtPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb, err := alertOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	parentPb, err := identity(&st.Parent)
	if err != nil {
		return nil, err
	}
	if parentPb != nil {
		pb.Parent = *parentPb
	}

	queryPb, err := alertQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	rearmPb, err := identity(&st.Rearm)
	if err != nil {
		return nil, err
	}
	if rearmPb != nil {
		pb.Rearm = *rearmPb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

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
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastTriggeredAtField, err := identity(&pb.LastTriggeredAt)
	if err != nil {
		return nil, err
	}
	if lastTriggeredAtField != nil {
		st.LastTriggeredAt = *lastTriggeredAtField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	optionsField, err := alertOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	parentField, err := identity(&pb.Parent)
	if err != nil {
		return nil, err
	}
	if parentField != nil {
		st.Parent = *parentField
	}
	queryField, err := alertQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	rearmField, err := identity(&pb.Rearm)
	if err != nil {
		return nil, err
	}
	if rearmField != nil {
		st.Rearm = *rearmField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
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
	CanEdit bool
	// The timestamp when this query was created.
	CreatedAt string
	// Data source ID maps to the ID of the data source used by the resource and
	// is distinct from the warehouse ID. [Learn more]
	//
	// [Learn more]: https://docs.databricks.com/api/workspace/datasources/list
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// Query ID.
	Id string
	// Indicates whether the query is trashed. Trashed queries can't be used in
	// dashboards, or appear in search results. If this boolean is `true`, the
	// `options` property for this query includes a `moved_to_trash_at`
	// timestamp. Trashed queries are permanently deleted after 30 days.
	IsArchived bool
	// Whether the query is a draft. Draft queries only appear in list views for
	// their owners. Visualizations from draft queries cannot appear on
	// dashboards.
	IsDraft bool
	// Whether this query object appears in the current user's favorites list.
	// This flag determines whether the star icon for favorites is selected.
	IsFavorite bool
	// Text parameter types are not safe from SQL injection for all types of
	// data source. Set this Boolean parameter to `true` if a query either does
	// not use any text type parameters or uses a data source type where text
	// type parameters are handled safely.
	IsSafe bool

	LastModifiedBy *User
	// The ID of the user who last saved changes to this query.
	LastModifiedById int
	// If there is a cached result for this query and user, this field includes
	// the query result ID. If this query uses parameters, this field is always
	// null.
	LatestQueryDataId string
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string

	Options *QueryOptions
	// The identifier of the workspace folder containing the object.
	Parent string
	// * `CAN_VIEW`: Can view the query * `CAN_RUN`: Can run the query *
	// `CAN_EDIT`: Can edit the query * `CAN_MANAGE`: Can manage the query
	PermissionTier PermissionLevel
	// The text of the query to be run.
	Query string
	// A SHA-256 hash of the query text along with the authenticated user ID.
	QueryHash string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole

	Tags []string
	// The timestamp at which this query was last updated.
	UpdatedAt string

	User *User
	// The ID of the user who owns the query.
	UserId int

	Visualizations []LegacyVisualization

	ForceSendFields []string
}

func legacyQueryToPb(st *LegacyQuery) (*legacyQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &legacyQueryPb{}
	canEditPb, err := identity(&st.CanEdit)
	if err != nil {
		return nil, err
	}
	if canEditPb != nil {
		pb.CanEdit = *canEditPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	dataSourceIdPb, err := identity(&st.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdPb != nil {
		pb.DataSourceId = *dataSourceIdPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	isArchivedPb, err := identity(&st.IsArchived)
	if err != nil {
		return nil, err
	}
	if isArchivedPb != nil {
		pb.IsArchived = *isArchivedPb
	}

	isDraftPb, err := identity(&st.IsDraft)
	if err != nil {
		return nil, err
	}
	if isDraftPb != nil {
		pb.IsDraft = *isDraftPb
	}

	isFavoritePb, err := identity(&st.IsFavorite)
	if err != nil {
		return nil, err
	}
	if isFavoritePb != nil {
		pb.IsFavorite = *isFavoritePb
	}

	isSafePb, err := identity(&st.IsSafe)
	if err != nil {
		return nil, err
	}
	if isSafePb != nil {
		pb.IsSafe = *isSafePb
	}

	lastModifiedByPb, err := userToPb(st.LastModifiedBy)
	if err != nil {
		return nil, err
	}
	if lastModifiedByPb != nil {
		pb.LastModifiedBy = lastModifiedByPb
	}

	lastModifiedByIdPb, err := identity(&st.LastModifiedById)
	if err != nil {
		return nil, err
	}
	if lastModifiedByIdPb != nil {
		pb.LastModifiedById = *lastModifiedByIdPb
	}

	latestQueryDataIdPb, err := identity(&st.LatestQueryDataId)
	if err != nil {
		return nil, err
	}
	if latestQueryDataIdPb != nil {
		pb.LatestQueryDataId = *latestQueryDataIdPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb, err := queryOptionsToPb(st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = optionsPb
	}

	parentPb, err := identity(&st.Parent)
	if err != nil {
		return nil, err
	}
	if parentPb != nil {
		pb.Parent = *parentPb
	}

	permissionTierPb, err := identity(&st.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierPb != nil {
		pb.PermissionTier = *permissionTierPb
	}

	queryPb, err := identity(&st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = *queryPb
	}

	queryHashPb, err := identity(&st.QueryHash)
	if err != nil {
		return nil, err
	}
	if queryHashPb != nil {
		pb.QueryHash = *queryHashPb
	}

	runAsRolePb, err := identity(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	userPb, err := userToPb(st.User)
	if err != nil {
		return nil, err
	}
	if userPb != nil {
		pb.User = userPb
	}

	userIdPb, err := identity(&st.UserId)
	if err != nil {
		return nil, err
	}
	if userIdPb != nil {
		pb.UserId = *userIdPb
	}

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
	canEditField, err := identity(&pb.CanEdit)
	if err != nil {
		return nil, err
	}
	if canEditField != nil {
		st.CanEdit = *canEditField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	dataSourceIdField, err := identity(&pb.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdField != nil {
		st.DataSourceId = *dataSourceIdField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	isArchivedField, err := identity(&pb.IsArchived)
	if err != nil {
		return nil, err
	}
	if isArchivedField != nil {
		st.IsArchived = *isArchivedField
	}
	isDraftField, err := identity(&pb.IsDraft)
	if err != nil {
		return nil, err
	}
	if isDraftField != nil {
		st.IsDraft = *isDraftField
	}
	isFavoriteField, err := identity(&pb.IsFavorite)
	if err != nil {
		return nil, err
	}
	if isFavoriteField != nil {
		st.IsFavorite = *isFavoriteField
	}
	isSafeField, err := identity(&pb.IsSafe)
	if err != nil {
		return nil, err
	}
	if isSafeField != nil {
		st.IsSafe = *isSafeField
	}
	lastModifiedByField, err := userFromPb(pb.LastModifiedBy)
	if err != nil {
		return nil, err
	}
	if lastModifiedByField != nil {
		st.LastModifiedBy = lastModifiedByField
	}
	lastModifiedByIdField, err := identity(&pb.LastModifiedById)
	if err != nil {
		return nil, err
	}
	if lastModifiedByIdField != nil {
		st.LastModifiedById = *lastModifiedByIdField
	}
	latestQueryDataIdField, err := identity(&pb.LatestQueryDataId)
	if err != nil {
		return nil, err
	}
	if latestQueryDataIdField != nil {
		st.LatestQueryDataId = *latestQueryDataIdField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	optionsField, err := queryOptionsFromPb(pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = optionsField
	}
	parentField, err := identity(&pb.Parent)
	if err != nil {
		return nil, err
	}
	if parentField != nil {
		st.Parent = *parentField
	}
	permissionTierField, err := identity(&pb.PermissionTier)
	if err != nil {
		return nil, err
	}
	if permissionTierField != nil {
		st.PermissionTier = *permissionTierField
	}
	queryField, err := identity(&pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = *queryField
	}
	queryHashField, err := identity(&pb.QueryHash)
	if err != nil {
		return nil, err
	}
	if queryHashField != nil {
		st.QueryHash = *queryHashField
	}
	runAsRoleField, err := identity(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	userField, err := userFromPb(pb.User)
	if err != nil {
		return nil, err
	}
	if userField != nil {
		st.User = userField
	}
	userIdField, err := identity(&pb.UserId)
	if err != nil {
		return nil, err
	}
	if userIdField != nil {
		st.UserId = *userIdField
	}

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
	CreatedAt string
	// A short description of this visualization. This is not displayed in the
	// UI.
	Description string
	// The UUID for this visualization.
	Id string
	// The name of the visualization that appears on dashboards and the query
	// screen.
	Name string
	// The options object varies widely from one visualization type to the next
	// and is unsupported. Databricks does not recommend modifying visualization
	// settings in JSON.
	Options any

	Query *LegacyQuery
	// The type of visualization: chart, table, pivot table, and so on.
	Type string

	UpdatedAt string

	ForceSendFields []string
}

func legacyVisualizationToPb(st *LegacyVisualization) (*legacyVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &legacyVisualizationPb{}
	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	optionsPb, err := identity(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	queryPb, err := legacyQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

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
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	optionsField, err := identity(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	queryField, err := legacyQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}

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
	PageSize int

	PageToken string

	ForceSendFields []string
}

func listAlertsRequestToPb(st *ListAlertsRequest) (*listAlertsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsRequestPb{}
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
	NextPageToken string

	Results []ListAlertsResponseAlert

	ForceSendFields []string
}

func listAlertsResponseToPb(st *ListAlertsResponse) (*listAlertsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	Condition *AlertCondition
	// The timestamp indicating when the alert was created.
	CreateTime string
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string
	// The display name of the alert.
	DisplayName string
	// UUID identifying the alert.
	Id string
	// The workspace state of the alert. Used for tracking trashed status.
	LifecycleState LifecycleState
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string
	// UUID of the query attached to the alert.
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int
	// Current state of the alert's trigger status. This field is set to UNKNOWN
	// if the alert has not yet been evaluated or ran into an error during the
	// last evaluation.
	State AlertState
	// Timestamp when the alert was last triggered, if the alert has been
	// triggered before.
	TriggerTime string
	// The timestamp indicating when the alert was updated.
	UpdateTime string

	ForceSendFields []string
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

	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	customBodyPb, err := identity(&st.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyPb != nil {
		pb.CustomBody = *customBodyPb
	}

	customSubjectPb, err := identity(&st.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectPb != nil {
		pb.CustomSubject = *customSubjectPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lifecycleStatePb, err := identity(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}

	notifyOnOkPb, err := identity(&st.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkPb != nil {
		pb.NotifyOnOk = *notifyOnOkPb
	}

	ownerUserNamePb, err := identity(&st.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNamePb != nil {
		pb.OwnerUserName = *ownerUserNamePb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	secondsToRetriggerPb, err := identity(&st.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerPb != nil {
		pb.SecondsToRetrigger = *secondsToRetriggerPb
	}

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	triggerTimePb, err := identity(&st.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimePb != nil {
		pb.TriggerTime = *triggerTimePb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

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
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	customBodyField, err := identity(&pb.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyField != nil {
		st.CustomBody = *customBodyField
	}
	customSubjectField, err := identity(&pb.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectField != nil {
		st.CustomSubject = *customSubjectField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lifecycleStateField, err := identity(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	notifyOnOkField, err := identity(&pb.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkField != nil {
		st.NotifyOnOk = *notifyOnOkField
	}
	ownerUserNameField, err := identity(&pb.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNameField != nil {
		st.OwnerUserName = *ownerUserNameField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	secondsToRetriggerField, err := identity(&pb.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerField != nil {
		st.SecondsToRetrigger = *secondsToRetriggerField
	}
	stateField, err := identity(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}
	triggerTimeField, err := identity(&pb.TriggerTime)
	if err != nil {
		return nil, err
	}
	if triggerTimeField != nil {
		st.TriggerTime = *triggerTimeField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

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
	PageSize int

	PageToken string

	ForceSendFields []string
}

func listAlertsV2RequestToPb(st *ListAlertsV2Request) (*listAlertsV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsV2RequestPb{}
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
	NextPageToken string

	Results []AlertV2

	ForceSendFields []string
}

func listAlertsV2ResponseToPb(st *ListAlertsV2Response) (*listAlertsV2ResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsV2ResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	var resultsPb []alertV2Pb
	for _, item := range st.Results {
		itemPb, err := alertV2ToPb(&item)
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

	Results []alertV2Pb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsV2ResponseFromPb(pb *listAlertsV2ResponsePb) (*ListAlertsV2Response, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsV2Response{}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var resultsField []AlertV2
	for _, itemPb := range pb.Results {
		item, err := alertV2FromPb(&itemPb)
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

// Get dashboard objects
type ListDashboardsRequest struct {
	// Name of dashboard attribute to order by.
	Order ListOrder
	// Page number to retrieve.
	Page int
	// Number of dashboards to return per page.
	PageSize int
	// Full text search term.
	Q string

	ForceSendFields []string
}

func listDashboardsRequestToPb(st *ListDashboardsRequest) (*listDashboardsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDashboardsRequestPb{}
	orderPb, err := identity(&st.Order)
	if err != nil {
		return nil, err
	}
	if orderPb != nil {
		pb.Order = *orderPb
	}

	pagePb, err := identity(&st.Page)
	if err != nil {
		return nil, err
	}
	if pagePb != nil {
		pb.Page = *pagePb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	qPb, err := identity(&st.Q)
	if err != nil {
		return nil, err
	}
	if qPb != nil {
		pb.Q = *qPb
	}

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
	orderField, err := identity(&pb.Order)
	if err != nil {
		return nil, err
	}
	if orderField != nil {
		st.Order = *orderField
	}
	pageField, err := identity(&pb.Page)
	if err != nil {
		return nil, err
	}
	if pageField != nil {
		st.Page = *pageField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	qField, err := identity(&pb.Q)
	if err != nil {
		return nil, err
	}
	if qField != nil {
		st.Q = *qField
	}

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
	Order string
	// Page number to retrieve.
	Page int
	// Number of queries to return per page.
	PageSize int
	// Full text search term
	Q string

	ForceSendFields []string
}

func listQueriesLegacyRequestToPb(st *ListQueriesLegacyRequest) (*listQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueriesLegacyRequestPb{}
	orderPb, err := identity(&st.Order)
	if err != nil {
		return nil, err
	}
	if orderPb != nil {
		pb.Order = *orderPb
	}

	pagePb, err := identity(&st.Page)
	if err != nil {
		return nil, err
	}
	if pagePb != nil {
		pb.Page = *pagePb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	qPb, err := identity(&st.Q)
	if err != nil {
		return nil, err
	}
	if qPb != nil {
		pb.Q = *qPb
	}

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
	orderField, err := identity(&pb.Order)
	if err != nil {
		return nil, err
	}
	if orderField != nil {
		st.Order = *orderField
	}
	pageField, err := identity(&pb.Page)
	if err != nil {
		return nil, err
	}
	if pageField != nil {
		st.Page = *pageField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	qField, err := identity(&pb.Q)
	if err != nil {
		return nil, err
	}
	if qField != nil {
		st.Q = *qField
	}

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
	PageSize int

	PageToken string

	ForceSendFields []string
}

func listQueriesRequestToPb(st *ListQueriesRequest) (*listQueriesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueriesRequestPb{}
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
	HasNextPage bool
	// A token that can be used to get the next page of results.
	NextPageToken string

	Res []QueryInfo

	ForceSendFields []string
}

func listQueriesResponseToPb(st *ListQueriesResponse) (*listQueriesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueriesResponsePb{}
	hasNextPagePb, err := identity(&st.HasNextPage)
	if err != nil {
		return nil, err
	}
	if hasNextPagePb != nil {
		pb.HasNextPage = *hasNextPagePb
	}

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	hasNextPageField, err := identity(&pb.HasNextPage)
	if err != nil {
		return nil, err
	}
	if hasNextPageField != nil {
		st.HasNextPage = *hasNextPageField
	}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	FilterBy *QueryFilter
	// Whether to include the query metrics with each query. Only use this for a
	// small subset of queries (max_results). Defaults to false.
	IncludeMetrics bool
	// Limit the number of results returned in one page. Must be less than 1000
	// and the default is 100.
	MaxResults int
	// A token that can be used to get the next page of results. The token can
	// contains characters that need to be encoded before using it in a URL. For
	// example, the character '+' needs to be replaced by %2B. This field is
	// optional.
	PageToken string

	ForceSendFields []string
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

	includeMetricsPb, err := identity(&st.IncludeMetrics)
	if err != nil {
		return nil, err
	}
	if includeMetricsPb != nil {
		pb.IncludeMetrics = *includeMetricsPb
	}

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
	includeMetricsField, err := identity(&pb.IncludeMetrics)
	if err != nil {
		return nil, err
	}
	if includeMetricsField != nil {
		st.IncludeMetrics = *includeMetricsField
	}
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

func (st *listQueryHistoryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueryHistoryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueryObjectsResponse struct {
	NextPageToken string

	Results []ListQueryObjectsResponseQuery

	ForceSendFields []string
}

func listQueryObjectsResponseToPb(st *ListQueryObjectsResponse) (*listQueryObjectsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueryObjectsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	Catalog string
	// Timestamp when this query was created.
	CreateTime string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string
	// UUID identifying the query.
	Id string
	// Username of the user who last saved changes to this query.
	LastModifierUserName string
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState
	// Username of the user that owns the query.
	OwnerUserName string
	// List of query parameter definitions.
	Parameters []QueryParameter
	// Text of the query to be run.
	QueryText string
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	Schema string

	Tags []string
	// Timestamp when this query was last updated.
	UpdateTime string
	// ID of the SQL warehouse attached to the query.
	WarehouseId string

	ForceSendFields []string
}

func listQueryObjectsResponseQueryToPb(st *ListQueryObjectsResponseQuery) (*listQueryObjectsResponseQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueryObjectsResponseQueryPb{}
	applyAutoLimitPb, err := identity(&st.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitPb != nil {
		pb.ApplyAutoLimit = *applyAutoLimitPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastModifierUserNamePb, err := identity(&st.LastModifierUserName)
	if err != nil {
		return nil, err
	}
	if lastModifierUserNamePb != nil {
		pb.LastModifierUserName = *lastModifierUserNamePb
	}

	lifecycleStatePb, err := identity(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}

	ownerUserNamePb, err := identity(&st.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNamePb != nil {
		pb.OwnerUserName = *ownerUserNamePb
	}

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

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	runAsModePb, err := identity(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
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
	applyAutoLimitField, err := identity(&pb.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitField != nil {
		st.ApplyAutoLimit = *applyAutoLimitField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastModifierUserNameField, err := identity(&pb.LastModifierUserName)
	if err != nil {
		return nil, err
	}
	if lastModifierUserNameField != nil {
		st.LastModifierUserName = *lastModifierUserNameField
	}
	lifecycleStateField, err := identity(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	ownerUserNameField, err := identity(&pb.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNameField != nil {
		st.OwnerUserName = *ownerUserNameField
	}

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
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}
	runAsModeField, err := identity(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
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

func (st *listQueryObjectsResponseQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueryObjectsResponseQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListResponse struct {
	// The total number of dashboards.
	Count int
	// The current page being displayed.
	Page int
	// The number of dashboards per page.
	PageSize int
	// List of dashboards returned.
	Results []Dashboard

	ForceSendFields []string
}

func listResponseToPb(st *ListResponse) (*listResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listResponsePb{}
	countPb, err := identity(&st.Count)
	if err != nil {
		return nil, err
	}
	if countPb != nil {
		pb.Count = *countPb
	}

	pagePb, err := identity(&st.Page)
	if err != nil {
		return nil, err
	}
	if pagePb != nil {
		pb.Page = *pagePb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

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
	countField, err := identity(&pb.Count)
	if err != nil {
		return nil, err
	}
	if countField != nil {
		st.Count = *countField
	}
	pageField, err := identity(&pb.Page)
	if err != nil {
		return nil, err
	}
	if pageField != nil {
		st.Page = *pageField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}

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
	Id string

	PageSize int

	PageToken string

	ForceSendFields []string
}

func listVisualizationsForQueryRequestToPb(st *ListVisualizationsForQueryRequest) (*listVisualizationsForQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVisualizationsForQueryRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
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
	NextPageToken string

	Results []Visualization

	ForceSendFields []string
}

func listVisualizationsForQueryResponseToPb(st *ListVisualizationsForQueryResponse) (*listVisualizationsForQueryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVisualizationsForQueryResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

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
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

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
	RunAsUserId int

	ForceSendFields []string
}

func listWarehousesRequestToPb(st *ListWarehousesRequest) (*listWarehousesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWarehousesRequestPb{}
	runAsUserIdPb, err := identity(&st.RunAsUserId)
	if err != nil {
		return nil, err
	}
	if runAsUserIdPb != nil {
		pb.RunAsUserId = *runAsUserIdPb
	}

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
	runAsUserIdField, err := identity(&pb.RunAsUserId)
	if err != nil {
		return nil, err
	}
	if runAsUserIdField != nil {
		st.RunAsUserId = *runAsUserIdField
	}

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
	Prefix string
	// Character that separates each selected parameter value. Defaults to a
	// comma.
	Separator string
	// Character that suffixes each selected parameter value.
	Suffix string

	ForceSendFields []string
}

func multiValuesOptionsToPb(st *MultiValuesOptions) (*multiValuesOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &multiValuesOptionsPb{}
	prefixPb, err := identity(&st.Prefix)
	if err != nil {
		return nil, err
	}
	if prefixPb != nil {
		pb.Prefix = *prefixPb
	}

	separatorPb, err := identity(&st.Separator)
	if err != nil {
		return nil, err
	}
	if separatorPb != nil {
		pb.Separator = *separatorPb
	}

	suffixPb, err := identity(&st.Suffix)
	if err != nil {
		return nil, err
	}
	if suffixPb != nil {
		pb.Suffix = *suffixPb
	}

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
	prefixField, err := identity(&pb.Prefix)
	if err != nil {
		return nil, err
	}
	if prefixField != nil {
		st.Prefix = *prefixField
	}
	separatorField, err := identity(&pb.Separator)
	if err != nil {
		return nil, err
	}
	if separatorField != nil {
		st.Separator = *separatorField
	}
	suffixField, err := identity(&pb.Suffix)
	if err != nil {
		return nil, err
	}
	if suffixField != nil {
		st.Suffix = *suffixField
	}

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
	Value float64

	ForceSendFields []string
}

func numericValueToPb(st *NumericValue) (*numericValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &numericValuePb{}
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
	Hostname string

	Path string

	Port int

	Protocol string

	ForceSendFields []string
}

func odbcParamsToPb(st *OdbcParams) (*odbcParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &odbcParamsPb{}
	hostnamePb, err := identity(&st.Hostname)
	if err != nil {
		return nil, err
	}
	if hostnamePb != nil {
		pb.Hostname = *hostnamePb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	portPb, err := identity(&st.Port)
	if err != nil {
		return nil, err
	}
	if portPb != nil {
		pb.Port = *portPb
	}

	protocolPb, err := identity(&st.Protocol)
	if err != nil {
		return nil, err
	}
	if protocolPb != nil {
		pb.Protocol = *protocolPb
	}

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
	hostnameField, err := identity(&pb.Hostname)
	if err != nil {
		return nil, err
	}
	if hostnameField != nil {
		st.Hostname = *hostnameField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}
	portField, err := identity(&pb.Port)
	if err != nil {
		return nil, err
	}
	if portField != nil {
		st.Port = *portField
	}
	protocolField, err := identity(&pb.Protocol)
	if err != nil {
		return nil, err
	}
	if protocolField != nil {
		st.Protocol = *protocolField
	}

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
	EnumOptions string
	// If specified, allows multiple values to be selected for this parameter.
	// Only applies to dropdown list and query-based dropdown list parameters.
	MultiValuesOptions *MultiValuesOptions
	// The literal parameter marker that appears between double curly braces in
	// the query text.
	Name string
	// The UUID of the query that provides the parameter values. Only applies
	// for query-based dropdown list parameters.
	QueryId string
	// The text displayed in a parameter picking widget.
	Title string
	// Parameters can have several different types.
	Type ParameterType
	// The default value for this parameter.
	Value any

	ForceSendFields []string
}

func parameterToPb(st *Parameter) (*parameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &parameterPb{}
	enumOptionsPb, err := identity(&st.EnumOptions)
	if err != nil {
		return nil, err
	}
	if enumOptionsPb != nil {
		pb.EnumOptions = *enumOptionsPb
	}

	multiValuesOptionsPb, err := multiValuesOptionsToPb(st.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsPb != nil {
		pb.MultiValuesOptions = multiValuesOptionsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	titlePb, err := identity(&st.Title)
	if err != nil {
		return nil, err
	}
	if titlePb != nil {
		pb.Title = *titlePb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
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
	enumOptionsField, err := identity(&pb.EnumOptions)
	if err != nil {
		return nil, err
	}
	if enumOptionsField != nil {
		st.EnumOptions = *enumOptionsField
	}
	multiValuesOptionsField, err := multiValuesOptionsFromPb(pb.MultiValuesOptions)
	if err != nil {
		return nil, err
	}
	if multiValuesOptionsField != nil {
		st.MultiValuesOptions = multiValuesOptionsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	titleField, err := identity(&pb.Title)
	if err != nil {
		return nil, err
	}
	if titleField != nil {
		st.Title = *titleField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
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
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	Catalog string
	// Timestamp when this query was created.
	CreateTime string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string
	// UUID identifying the query.
	Id string
	// Username of the user who last saved changes to this query.
	LastModifierUserName string
	// Indicates whether the query is trashed.
	LifecycleState LifecycleState
	// Username of the user that owns the query.
	OwnerUserName string
	// List of query parameter definitions.
	Parameters []QueryParameter
	// Workspace path of the workspace folder containing the object.
	ParentPath string
	// Text of the query to be run.
	QueryText string
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	Schema string

	Tags []string
	// Timestamp when this query was last updated.
	UpdateTime string
	// ID of the SQL warehouse attached to the query.
	WarehouseId string

	ForceSendFields []string
}

func queryToPb(st *Query) (*queryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryPb{}
	applyAutoLimitPb, err := identity(&st.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitPb != nil {
		pb.ApplyAutoLimit = *applyAutoLimitPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastModifierUserNamePb, err := identity(&st.LastModifierUserName)
	if err != nil {
		return nil, err
	}
	if lastModifierUserNamePb != nil {
		pb.LastModifierUserName = *lastModifierUserNamePb
	}

	lifecycleStatePb, err := identity(&st.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStatePb != nil {
		pb.LifecycleState = *lifecycleStatePb
	}

	ownerUserNamePb, err := identity(&st.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNamePb != nil {
		pb.OwnerUserName = *ownerUserNamePb
	}

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

	parentPathPb, err := identity(&st.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathPb != nil {
		pb.ParentPath = *parentPathPb
	}

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	runAsModePb, err := identity(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
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
	applyAutoLimitField, err := identity(&pb.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitField != nil {
		st.ApplyAutoLimit = *applyAutoLimitField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastModifierUserNameField, err := identity(&pb.LastModifierUserName)
	if err != nil {
		return nil, err
	}
	if lastModifierUserNameField != nil {
		st.LastModifierUserName = *lastModifierUserNameField
	}
	lifecycleStateField, err := identity(&pb.LifecycleState)
	if err != nil {
		return nil, err
	}
	if lifecycleStateField != nil {
		st.LifecycleState = *lifecycleStateField
	}
	ownerUserNameField, err := identity(&pb.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNameField != nil {
		st.OwnerUserName = *ownerUserNameField
	}

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
	parentPathField, err := identity(&pb.ParentPath)
	if err != nil {
		return nil, err
	}
	if parentPathField != nil {
		st.ParentPath = *parentPathField
	}
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}
	runAsModeField, err := identity(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
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

func (st *queryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryBackedValue struct {
	// If specified, allows multiple values to be selected for this parameter.
	MultiValuesOptions *MultiValuesOptions
	// UUID of the query that provides the parameter values.
	QueryId string
	// List of selected query parameter values.
	Values []string

	ForceSendFields []string
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

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	var valuesPb []string
	for _, item := range st.Values {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			valuesPb = append(valuesPb, *itemPb)
		}
	}
	pb.Values = valuesPb

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
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}

	var valuesField []string
	for _, itemPb := range pb.Values {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			valuesField = append(valuesField, *item)
		}
	}
	st.Values = valuesField

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
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any
	// The text of the query to be run.
	Query string

	QueryId string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole

	Tags []string

	ForceSendFields []string
}

func queryEditContentToPb(st *QueryEditContent) (*queryEditContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryEditContentPb{}
	dataSourceIdPb, err := identity(&st.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdPb != nil {
		pb.DataSourceId = *dataSourceIdPb
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

	optionsPb, err := identity(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	queryPb, err := identity(&st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = *queryPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	runAsRolePb, err := identity(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

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
	dataSourceIdField, err := identity(&pb.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdField != nil {
		st.DataSourceId = *dataSourceIdField
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
	optionsField, err := identity(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	queryField, err := identity(&pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = *queryField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	runAsRoleField, err := identity(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

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
	QueryStartTimeRange *TimeRange
	// A list of statement IDs.
	StatementIds []string

	Statuses []QueryStatus
	// A list of user IDs who ran the queries.
	UserIds []int64
	// A list of warehouse IDs.
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

	var statementIdsPb []string
	for _, item := range st.StatementIds {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			statementIdsPb = append(statementIdsPb, *itemPb)
		}
	}
	pb.StatementIds = statementIdsPb

	var statusesPb []QueryStatus
	for _, item := range st.Statuses {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			statusesPb = append(statusesPb, *itemPb)
		}
	}
	pb.Statuses = statusesPb

	var userIdsPb []int64
	for _, item := range st.UserIds {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			userIdsPb = append(userIdsPb, *itemPb)
		}
	}
	pb.UserIds = userIdsPb

	var warehouseIdsPb []string
	for _, item := range st.WarehouseIds {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			warehouseIdsPb = append(warehouseIdsPb, *itemPb)
		}
	}
	pb.WarehouseIds = warehouseIdsPb

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

	var statementIdsField []string
	for _, itemPb := range pb.StatementIds {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			statementIdsField = append(statementIdsField, *item)
		}
	}
	st.StatementIds = statementIdsField

	var statusesField []QueryStatus
	for _, itemPb := range pb.Statuses {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			statusesField = append(statusesField, *item)
		}
	}
	st.Statuses = statusesField

	var userIdsField []int64
	for _, itemPb := range pb.UserIds {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			userIdsField = append(userIdsField, *item)
		}
	}
	st.UserIds = userIdsField

	var warehouseIdsField []string
	for _, itemPb := range pb.WarehouseIds {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			warehouseIdsField = append(warehouseIdsField, *item)
		}
	}
	st.WarehouseIds = warehouseIdsField

	return st, nil
}

type QueryInfo struct {
	// SQL Warehouse channel information at the time of query execution
	ChannelUsed *ChannelInfo
	// Client application that ran the statement. For example: Databricks SQL
	// Editor, Tableau, and Power BI. This field is derived from information
	// provided by client applications. While values are expected to remain
	// static over time, this cannot be guaranteed.
	ClientApplication string
	// Total execution time of the statement ( excluding result fetch time ).
	Duration int64
	// Alias for `warehouse_id`.
	EndpointId string
	// Message describing why the query could not complete.
	ErrorMessage string
	// The ID of the user whose credentials were used to run the query.
	ExecutedAsUserId int64
	// The email address or username of the user whose credentials were used to
	// run the query.
	ExecutedAsUserName string
	// The time execution of the query ended.
	ExecutionEndTimeMs int64
	// Whether more updates for the query are expected.
	IsFinal bool
	// A key that can be used to look up query details.
	LookupKey string
	// Metrics about query execution.
	Metrics *QueryMetrics
	// Whether plans exist for the execution, or the reason why they are missing
	PlansState PlansState
	// The time the query ended.
	QueryEndTimeMs int64
	// The query ID.
	QueryId string
	// A struct that contains key-value pairs representing Databricks entities
	// that were involved in the execution of this statement, such as jobs,
	// notebooks, or dashboards. This field only records Databricks entities.
	QuerySource *ExternalQuerySource
	// The time the query started.
	QueryStartTimeMs int64
	// The text of the query.
	QueryText string
	// The number of results returned by the query.
	RowsProduced int64
	// URL to the Spark UI query plan.
	SparkUiUrl string
	// Type of statement for this query
	StatementType QueryStatementType
	// Query status with one the following values:
	//
	// - `QUEUED`: Query has been received and queued. - `RUNNING`: Query has
	// started. - `CANCELED`: Query has been cancelled by the user. - `FAILED`:
	// Query has failed. - `FINISHED`: Query has completed.
	Status QueryStatus
	// The ID of the user who ran the query.
	UserId int64
	// The email address or username of the user who ran the query.
	UserName string
	// Warehouse ID.
	WarehouseId string

	ForceSendFields []string
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

	clientApplicationPb, err := identity(&st.ClientApplication)
	if err != nil {
		return nil, err
	}
	if clientApplicationPb != nil {
		pb.ClientApplication = *clientApplicationPb
	}

	durationPb, err := identity(&st.Duration)
	if err != nil {
		return nil, err
	}
	if durationPb != nil {
		pb.Duration = *durationPb
	}

	endpointIdPb, err := identity(&st.EndpointId)
	if err != nil {
		return nil, err
	}
	if endpointIdPb != nil {
		pb.EndpointId = *endpointIdPb
	}

	errorMessagePb, err := identity(&st.ErrorMessage)
	if err != nil {
		return nil, err
	}
	if errorMessagePb != nil {
		pb.ErrorMessage = *errorMessagePb
	}

	executedAsUserIdPb, err := identity(&st.ExecutedAsUserId)
	if err != nil {
		return nil, err
	}
	if executedAsUserIdPb != nil {
		pb.ExecutedAsUserId = *executedAsUserIdPb
	}

	executedAsUserNamePb, err := identity(&st.ExecutedAsUserName)
	if err != nil {
		return nil, err
	}
	if executedAsUserNamePb != nil {
		pb.ExecutedAsUserName = *executedAsUserNamePb
	}

	executionEndTimeMsPb, err := identity(&st.ExecutionEndTimeMs)
	if err != nil {
		return nil, err
	}
	if executionEndTimeMsPb != nil {
		pb.ExecutionEndTimeMs = *executionEndTimeMsPb
	}

	isFinalPb, err := identity(&st.IsFinal)
	if err != nil {
		return nil, err
	}
	if isFinalPb != nil {
		pb.IsFinal = *isFinalPb
	}

	lookupKeyPb, err := identity(&st.LookupKey)
	if err != nil {
		return nil, err
	}
	if lookupKeyPb != nil {
		pb.LookupKey = *lookupKeyPb
	}

	metricsPb, err := queryMetricsToPb(st.Metrics)
	if err != nil {
		return nil, err
	}
	if metricsPb != nil {
		pb.Metrics = metricsPb
	}

	plansStatePb, err := identity(&st.PlansState)
	if err != nil {
		return nil, err
	}
	if plansStatePb != nil {
		pb.PlansState = *plansStatePb
	}

	queryEndTimeMsPb, err := identity(&st.QueryEndTimeMs)
	if err != nil {
		return nil, err
	}
	if queryEndTimeMsPb != nil {
		pb.QueryEndTimeMs = *queryEndTimeMsPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	querySourcePb, err := externalQuerySourceToPb(st.QuerySource)
	if err != nil {
		return nil, err
	}
	if querySourcePb != nil {
		pb.QuerySource = querySourcePb
	}

	queryStartTimeMsPb, err := identity(&st.QueryStartTimeMs)
	if err != nil {
		return nil, err
	}
	if queryStartTimeMsPb != nil {
		pb.QueryStartTimeMs = *queryStartTimeMsPb
	}

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	rowsProducedPb, err := identity(&st.RowsProduced)
	if err != nil {
		return nil, err
	}
	if rowsProducedPb != nil {
		pb.RowsProduced = *rowsProducedPb
	}

	sparkUiUrlPb, err := identity(&st.SparkUiUrl)
	if err != nil {
		return nil, err
	}
	if sparkUiUrlPb != nil {
		pb.SparkUiUrl = *sparkUiUrlPb
	}

	statementTypePb, err := identity(&st.StatementType)
	if err != nil {
		return nil, err
	}
	if statementTypePb != nil {
		pb.StatementType = *statementTypePb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	userIdPb, err := identity(&st.UserId)
	if err != nil {
		return nil, err
	}
	if userIdPb != nil {
		pb.UserId = *userIdPb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
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
	clientApplicationField, err := identity(&pb.ClientApplication)
	if err != nil {
		return nil, err
	}
	if clientApplicationField != nil {
		st.ClientApplication = *clientApplicationField
	}
	durationField, err := identity(&pb.Duration)
	if err != nil {
		return nil, err
	}
	if durationField != nil {
		st.Duration = *durationField
	}
	endpointIdField, err := identity(&pb.EndpointId)
	if err != nil {
		return nil, err
	}
	if endpointIdField != nil {
		st.EndpointId = *endpointIdField
	}
	errorMessageField, err := identity(&pb.ErrorMessage)
	if err != nil {
		return nil, err
	}
	if errorMessageField != nil {
		st.ErrorMessage = *errorMessageField
	}
	executedAsUserIdField, err := identity(&pb.ExecutedAsUserId)
	if err != nil {
		return nil, err
	}
	if executedAsUserIdField != nil {
		st.ExecutedAsUserId = *executedAsUserIdField
	}
	executedAsUserNameField, err := identity(&pb.ExecutedAsUserName)
	if err != nil {
		return nil, err
	}
	if executedAsUserNameField != nil {
		st.ExecutedAsUserName = *executedAsUserNameField
	}
	executionEndTimeMsField, err := identity(&pb.ExecutionEndTimeMs)
	if err != nil {
		return nil, err
	}
	if executionEndTimeMsField != nil {
		st.ExecutionEndTimeMs = *executionEndTimeMsField
	}
	isFinalField, err := identity(&pb.IsFinal)
	if err != nil {
		return nil, err
	}
	if isFinalField != nil {
		st.IsFinal = *isFinalField
	}
	lookupKeyField, err := identity(&pb.LookupKey)
	if err != nil {
		return nil, err
	}
	if lookupKeyField != nil {
		st.LookupKey = *lookupKeyField
	}
	metricsField, err := queryMetricsFromPb(pb.Metrics)
	if err != nil {
		return nil, err
	}
	if metricsField != nil {
		st.Metrics = metricsField
	}
	plansStateField, err := identity(&pb.PlansState)
	if err != nil {
		return nil, err
	}
	if plansStateField != nil {
		st.PlansState = *plansStateField
	}
	queryEndTimeMsField, err := identity(&pb.QueryEndTimeMs)
	if err != nil {
		return nil, err
	}
	if queryEndTimeMsField != nil {
		st.QueryEndTimeMs = *queryEndTimeMsField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	querySourceField, err := externalQuerySourceFromPb(pb.QuerySource)
	if err != nil {
		return nil, err
	}
	if querySourceField != nil {
		st.QuerySource = querySourceField
	}
	queryStartTimeMsField, err := identity(&pb.QueryStartTimeMs)
	if err != nil {
		return nil, err
	}
	if queryStartTimeMsField != nil {
		st.QueryStartTimeMs = *queryStartTimeMsField
	}
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}
	rowsProducedField, err := identity(&pb.RowsProduced)
	if err != nil {
		return nil, err
	}
	if rowsProducedField != nil {
		st.RowsProduced = *rowsProducedField
	}
	sparkUiUrlField, err := identity(&pb.SparkUiUrl)
	if err != nil {
		return nil, err
	}
	if sparkUiUrlField != nil {
		st.SparkUiUrl = *sparkUiUrlField
	}
	statementTypeField, err := identity(&pb.StatementType)
	if err != nil {
		return nil, err
	}
	if statementTypeField != nil {
		st.StatementType = *statementTypeField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	userIdField, err := identity(&pb.UserId)
	if err != nil {
		return nil, err
	}
	if userIdField != nil {
		st.UserId = *userIdField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
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

func (st *queryInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryList struct {
	// The total number of queries.
	Count int
	// The page number that is currently displayed.
	Page int
	// The number of queries per page.
	PageSize int
	// List of queries returned.
	Results []LegacyQuery

	ForceSendFields []string
}

func queryListToPb(st *QueryList) (*queryListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryListPb{}
	countPb, err := identity(&st.Count)
	if err != nil {
		return nil, err
	}
	if countPb != nil {
		pb.Count = *countPb
	}

	pagePb, err := identity(&st.Page)
	if err != nil {
		return nil, err
	}
	if pagePb != nil {
		pb.Page = *pagePb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

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
	countField, err := identity(&pb.Count)
	if err != nil {
		return nil, err
	}
	if countField != nil {
		st.Count = *countField
	}
	pageField, err := identity(&pb.Page)
	if err != nil {
		return nil, err
	}
	if pageField != nil {
		st.Page = *pageField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}

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
	CompilationTimeMs int64
	// Time spent executing the query, in milliseconds.
	ExecutionTimeMs int64
	// Total amount of data sent over the network between executor nodes during
	// shuffle, in bytes.
	NetworkSentBytes int64
	// Timestamp of when the query was enqueued waiting while the warehouse was
	// at max load. This field is optional and will not appear if the query
	// skipped the overloading queue.
	OverloadingQueueStartTimestamp int64
	// Total execution time for all individual Photon query engine tasks in the
	// query, in milliseconds.
	PhotonTotalTimeMs int64
	// Timestamp of when the query was enqueued waiting for a cluster to be
	// provisioned for the warehouse. This field is optional and will not appear
	// if the query skipped the provisioning queue.
	ProvisioningQueueStartTimestamp int64
	// Total number of bytes in all tables not read due to pruning
	PrunedBytes int64
	// Total number of files from all tables not read due to pruning
	PrunedFilesCount int64
	// Timestamp of when the underlying compute started compilation of the
	// query.
	QueryCompilationStartTimestamp int64
	// Total size of data read by the query, in bytes.
	ReadBytes int64
	// Size of persistent data read from the cache, in bytes.
	ReadCacheBytes int64
	// Number of files read after pruning
	ReadFilesCount int64
	// Number of partitions read after pruning.
	ReadPartitionsCount int64
	// Size of persistent data read from cloud object storage on your cloud
	// tenant, in bytes.
	ReadRemoteBytes int64
	// Time spent fetching the query results after the execution finished, in
	// milliseconds.
	ResultFetchTimeMs int64
	// `true` if the query result was fetched from cache, `false` otherwise.
	ResultFromCache bool
	// Total number of rows returned by the query.
	RowsProducedCount int64
	// Total number of rows read by the query.
	RowsReadCount int64
	// Size of data temporarily written to disk while executing the query, in
	// bytes.
	SpillToDiskBytes int64
	// Sum of execution time for all of the query’s tasks, in milliseconds.
	TaskTotalTimeMs int64
	// Total execution time of the query from the client’s point of view, in
	// milliseconds.
	TotalTimeMs int64
	// Size pf persistent data written to cloud object storage in your cloud
	// tenant, in bytes.
	WriteRemoteBytes int64

	ForceSendFields []string
}

func queryMetricsToPb(st *QueryMetrics) (*queryMetricsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryMetricsPb{}
	compilationTimeMsPb, err := identity(&st.CompilationTimeMs)
	if err != nil {
		return nil, err
	}
	if compilationTimeMsPb != nil {
		pb.CompilationTimeMs = *compilationTimeMsPb
	}

	executionTimeMsPb, err := identity(&st.ExecutionTimeMs)
	if err != nil {
		return nil, err
	}
	if executionTimeMsPb != nil {
		pb.ExecutionTimeMs = *executionTimeMsPb
	}

	networkSentBytesPb, err := identity(&st.NetworkSentBytes)
	if err != nil {
		return nil, err
	}
	if networkSentBytesPb != nil {
		pb.NetworkSentBytes = *networkSentBytesPb
	}

	overloadingQueueStartTimestampPb, err := identity(&st.OverloadingQueueStartTimestamp)
	if err != nil {
		return nil, err
	}
	if overloadingQueueStartTimestampPb != nil {
		pb.OverloadingQueueStartTimestamp = *overloadingQueueStartTimestampPb
	}

	photonTotalTimeMsPb, err := identity(&st.PhotonTotalTimeMs)
	if err != nil {
		return nil, err
	}
	if photonTotalTimeMsPb != nil {
		pb.PhotonTotalTimeMs = *photonTotalTimeMsPb
	}

	provisioningQueueStartTimestampPb, err := identity(&st.ProvisioningQueueStartTimestamp)
	if err != nil {
		return nil, err
	}
	if provisioningQueueStartTimestampPb != nil {
		pb.ProvisioningQueueStartTimestamp = *provisioningQueueStartTimestampPb
	}

	prunedBytesPb, err := identity(&st.PrunedBytes)
	if err != nil {
		return nil, err
	}
	if prunedBytesPb != nil {
		pb.PrunedBytes = *prunedBytesPb
	}

	prunedFilesCountPb, err := identity(&st.PrunedFilesCount)
	if err != nil {
		return nil, err
	}
	if prunedFilesCountPb != nil {
		pb.PrunedFilesCount = *prunedFilesCountPb
	}

	queryCompilationStartTimestampPb, err := identity(&st.QueryCompilationStartTimestamp)
	if err != nil {
		return nil, err
	}
	if queryCompilationStartTimestampPb != nil {
		pb.QueryCompilationStartTimestamp = *queryCompilationStartTimestampPb
	}

	readBytesPb, err := identity(&st.ReadBytes)
	if err != nil {
		return nil, err
	}
	if readBytesPb != nil {
		pb.ReadBytes = *readBytesPb
	}

	readCacheBytesPb, err := identity(&st.ReadCacheBytes)
	if err != nil {
		return nil, err
	}
	if readCacheBytesPb != nil {
		pb.ReadCacheBytes = *readCacheBytesPb
	}

	readFilesCountPb, err := identity(&st.ReadFilesCount)
	if err != nil {
		return nil, err
	}
	if readFilesCountPb != nil {
		pb.ReadFilesCount = *readFilesCountPb
	}

	readPartitionsCountPb, err := identity(&st.ReadPartitionsCount)
	if err != nil {
		return nil, err
	}
	if readPartitionsCountPb != nil {
		pb.ReadPartitionsCount = *readPartitionsCountPb
	}

	readRemoteBytesPb, err := identity(&st.ReadRemoteBytes)
	if err != nil {
		return nil, err
	}
	if readRemoteBytesPb != nil {
		pb.ReadRemoteBytes = *readRemoteBytesPb
	}

	resultFetchTimeMsPb, err := identity(&st.ResultFetchTimeMs)
	if err != nil {
		return nil, err
	}
	if resultFetchTimeMsPb != nil {
		pb.ResultFetchTimeMs = *resultFetchTimeMsPb
	}

	resultFromCachePb, err := identity(&st.ResultFromCache)
	if err != nil {
		return nil, err
	}
	if resultFromCachePb != nil {
		pb.ResultFromCache = *resultFromCachePb
	}

	rowsProducedCountPb, err := identity(&st.RowsProducedCount)
	if err != nil {
		return nil, err
	}
	if rowsProducedCountPb != nil {
		pb.RowsProducedCount = *rowsProducedCountPb
	}

	rowsReadCountPb, err := identity(&st.RowsReadCount)
	if err != nil {
		return nil, err
	}
	if rowsReadCountPb != nil {
		pb.RowsReadCount = *rowsReadCountPb
	}

	spillToDiskBytesPb, err := identity(&st.SpillToDiskBytes)
	if err != nil {
		return nil, err
	}
	if spillToDiskBytesPb != nil {
		pb.SpillToDiskBytes = *spillToDiskBytesPb
	}

	taskTotalTimeMsPb, err := identity(&st.TaskTotalTimeMs)
	if err != nil {
		return nil, err
	}
	if taskTotalTimeMsPb != nil {
		pb.TaskTotalTimeMs = *taskTotalTimeMsPb
	}

	totalTimeMsPb, err := identity(&st.TotalTimeMs)
	if err != nil {
		return nil, err
	}
	if totalTimeMsPb != nil {
		pb.TotalTimeMs = *totalTimeMsPb
	}

	writeRemoteBytesPb, err := identity(&st.WriteRemoteBytes)
	if err != nil {
		return nil, err
	}
	if writeRemoteBytesPb != nil {
		pb.WriteRemoteBytes = *writeRemoteBytesPb
	}

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
	compilationTimeMsField, err := identity(&pb.CompilationTimeMs)
	if err != nil {
		return nil, err
	}
	if compilationTimeMsField != nil {
		st.CompilationTimeMs = *compilationTimeMsField
	}
	executionTimeMsField, err := identity(&pb.ExecutionTimeMs)
	if err != nil {
		return nil, err
	}
	if executionTimeMsField != nil {
		st.ExecutionTimeMs = *executionTimeMsField
	}
	networkSentBytesField, err := identity(&pb.NetworkSentBytes)
	if err != nil {
		return nil, err
	}
	if networkSentBytesField != nil {
		st.NetworkSentBytes = *networkSentBytesField
	}
	overloadingQueueStartTimestampField, err := identity(&pb.OverloadingQueueStartTimestamp)
	if err != nil {
		return nil, err
	}
	if overloadingQueueStartTimestampField != nil {
		st.OverloadingQueueStartTimestamp = *overloadingQueueStartTimestampField
	}
	photonTotalTimeMsField, err := identity(&pb.PhotonTotalTimeMs)
	if err != nil {
		return nil, err
	}
	if photonTotalTimeMsField != nil {
		st.PhotonTotalTimeMs = *photonTotalTimeMsField
	}
	provisioningQueueStartTimestampField, err := identity(&pb.ProvisioningQueueStartTimestamp)
	if err != nil {
		return nil, err
	}
	if provisioningQueueStartTimestampField != nil {
		st.ProvisioningQueueStartTimestamp = *provisioningQueueStartTimestampField
	}
	prunedBytesField, err := identity(&pb.PrunedBytes)
	if err != nil {
		return nil, err
	}
	if prunedBytesField != nil {
		st.PrunedBytes = *prunedBytesField
	}
	prunedFilesCountField, err := identity(&pb.PrunedFilesCount)
	if err != nil {
		return nil, err
	}
	if prunedFilesCountField != nil {
		st.PrunedFilesCount = *prunedFilesCountField
	}
	queryCompilationStartTimestampField, err := identity(&pb.QueryCompilationStartTimestamp)
	if err != nil {
		return nil, err
	}
	if queryCompilationStartTimestampField != nil {
		st.QueryCompilationStartTimestamp = *queryCompilationStartTimestampField
	}
	readBytesField, err := identity(&pb.ReadBytes)
	if err != nil {
		return nil, err
	}
	if readBytesField != nil {
		st.ReadBytes = *readBytesField
	}
	readCacheBytesField, err := identity(&pb.ReadCacheBytes)
	if err != nil {
		return nil, err
	}
	if readCacheBytesField != nil {
		st.ReadCacheBytes = *readCacheBytesField
	}
	readFilesCountField, err := identity(&pb.ReadFilesCount)
	if err != nil {
		return nil, err
	}
	if readFilesCountField != nil {
		st.ReadFilesCount = *readFilesCountField
	}
	readPartitionsCountField, err := identity(&pb.ReadPartitionsCount)
	if err != nil {
		return nil, err
	}
	if readPartitionsCountField != nil {
		st.ReadPartitionsCount = *readPartitionsCountField
	}
	readRemoteBytesField, err := identity(&pb.ReadRemoteBytes)
	if err != nil {
		return nil, err
	}
	if readRemoteBytesField != nil {
		st.ReadRemoteBytes = *readRemoteBytesField
	}
	resultFetchTimeMsField, err := identity(&pb.ResultFetchTimeMs)
	if err != nil {
		return nil, err
	}
	if resultFetchTimeMsField != nil {
		st.ResultFetchTimeMs = *resultFetchTimeMsField
	}
	resultFromCacheField, err := identity(&pb.ResultFromCache)
	if err != nil {
		return nil, err
	}
	if resultFromCacheField != nil {
		st.ResultFromCache = *resultFromCacheField
	}
	rowsProducedCountField, err := identity(&pb.RowsProducedCount)
	if err != nil {
		return nil, err
	}
	if rowsProducedCountField != nil {
		st.RowsProducedCount = *rowsProducedCountField
	}
	rowsReadCountField, err := identity(&pb.RowsReadCount)
	if err != nil {
		return nil, err
	}
	if rowsReadCountField != nil {
		st.RowsReadCount = *rowsReadCountField
	}
	spillToDiskBytesField, err := identity(&pb.SpillToDiskBytes)
	if err != nil {
		return nil, err
	}
	if spillToDiskBytesField != nil {
		st.SpillToDiskBytes = *spillToDiskBytesField
	}
	taskTotalTimeMsField, err := identity(&pb.TaskTotalTimeMs)
	if err != nil {
		return nil, err
	}
	if taskTotalTimeMsField != nil {
		st.TaskTotalTimeMs = *taskTotalTimeMsField
	}
	totalTimeMsField, err := identity(&pb.TotalTimeMs)
	if err != nil {
		return nil, err
	}
	if totalTimeMsField != nil {
		st.TotalTimeMs = *totalTimeMsField
	}
	writeRemoteBytesField, err := identity(&pb.WriteRemoteBytes)
	if err != nil {
		return nil, err
	}
	if writeRemoteBytesField != nil {
		st.WriteRemoteBytes = *writeRemoteBytesField
	}

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
	Catalog string
	// The timestamp when this query was moved to trash. Only present when the
	// `is_archived` property is `true`. Trashed items are deleted after thirty
	// days.
	MovedToTrashAt string

	Parameters []Parameter
	// The name of the schema to execute this query in.
	Schema string

	ForceSendFields []string
}

func queryOptionsToPb(st *QueryOptions) (*queryOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryOptionsPb{}
	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	movedToTrashAtPb, err := identity(&st.MovedToTrashAt)
	if err != nil {
		return nil, err
	}
	if movedToTrashAtPb != nil {
		pb.MovedToTrashAt = *movedToTrashAtPb
	}

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

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

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
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	movedToTrashAtField, err := identity(&pb.MovedToTrashAt)
	if err != nil {
		return nil, err
	}
	if movedToTrashAtField != nil {
		st.MovedToTrashAt = *movedToTrashAtField
	}

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
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}

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
	DateRangeValue *DateRangeValue
	// Date query parameter value. Can only specify one of `dynamic_date_value`
	// or `date_value`.
	DateValue *DateValue
	// Dropdown query parameter value.
	EnumValue *EnumValue
	// Literal parameter marker that appears between double curly braces in the
	// query text.
	Name string
	// Numeric query parameter value.
	NumericValue *NumericValue
	// Query-based dropdown query parameter value.
	QueryBackedValue *QueryBackedValue
	// Text query parameter value.
	TextValue *TextValue
	// Text displayed in the user-facing parameter widget in the UI.
	Title string

	ForceSendFields []string
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

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

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

	titlePb, err := identity(&st.Title)
	if err != nil {
		return nil, err
	}
	if titlePb != nil {
		pb.Title = *titlePb
	}

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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
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
	titleField, err := identity(&pb.Title)
	if err != nil {
		return nil, err
	}
	if titleField != nil {
		st.Title = *titleField
	}

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
	DataSourceId string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// The title of this query that appears in list views, widget headings, and
	// on the query page.
	Name string
	// Exclusively used for storing a list parameter definitions. A parameter is
	// an object with `title`, `name`, `type`, and `value` properties. The
	// `value` field here is the default value. It can be overridden at runtime.
	Options any
	// The identifier of the workspace folder containing the object.
	Parent string
	// The text of the query to be run.
	Query string
	// Sets the **Run as** role for the object. Must be set to one of `"viewer"`
	// (signifying "run as viewer" behavior) or `"owner"` (signifying "run as
	// owner" behavior)
	RunAsRole RunAsRole

	Tags []string

	ForceSendFields []string
}

func queryPostContentToPb(st *QueryPostContent) (*queryPostContentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryPostContentPb{}
	dataSourceIdPb, err := identity(&st.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdPb != nil {
		pb.DataSourceId = *dataSourceIdPb
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

	optionsPb, err := identity(&st.Options)
	if err != nil {
		return nil, err
	}
	if optionsPb != nil {
		pb.Options = *optionsPb
	}

	parentPb, err := identity(&st.Parent)
	if err != nil {
		return nil, err
	}
	if parentPb != nil {
		pb.Parent = *parentPb
	}

	queryPb, err := identity(&st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = *queryPb
	}

	runAsRolePb, err := identity(&st.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRolePb != nil {
		pb.RunAsRole = *runAsRolePb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

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
	dataSourceIdField, err := identity(&pb.DataSourceId)
	if err != nil {
		return nil, err
	}
	if dataSourceIdField != nil {
		st.DataSourceId = *dataSourceIdField
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
	optionsField, err := identity(&pb.Options)
	if err != nil {
		return nil, err
	}
	if optionsField != nil {
		st.Options = *optionsField
	}
	parentField, err := identity(&pb.Parent)
	if err != nil {
		return nil, err
	}
	if parentField != nil {
		st.Parent = *parentField
	}
	queryField, err := identity(&pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = *queryField
	}
	runAsRoleField, err := identity(&pb.RunAsRole)
	if err != nil {
		return nil, err
	}
	if runAsRoleField != nil {
		st.RunAsRole = *runAsRoleField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

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
	ConfigPair []EndpointConfPair

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
	DashboardId string
}

func restoreDashboardRequestToPb(st *RestoreDashboardRequest) (*restoreDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreDashboardRequestPb{}
	dashboardIdPb, err := identity(&st.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdPb != nil {
		pb.DashboardId = *dashboardIdPb
	}

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
	dashboardIdField, err := identity(&pb.DashboardId)
	if err != nil {
		return nil, err
	}
	if dashboardIdField != nil {
		st.DashboardId = *dashboardIdField
	}

	return st, nil
}

// Restore a query
type RestoreQueriesLegacyRequest struct {
	QueryId string
}

func restoreQueriesLegacyRequestToPb(st *RestoreQueriesLegacyRequest) (*restoreQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreQueriesLegacyRequestPb{}
	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

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
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}

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
	ByteCount int64
	// The position within the sequence of result set chunks.
	ChunkIndex int
	// The `JSON_ARRAY` format is an array of arrays of values, where each
	// non-null value is formatted as a string. Null values are encoded as JSON
	// `null`.
	DataArray [][]string

	ExternalLinks []ExternalLink
	// When fetching, provides the `chunk_index` for the _next_ chunk. If
	// absent, indicates there are no more chunks. The next chunk can be fetched
	// with a :method:statementexecution/getStatementResultChunkN request.
	NextChunkIndex int
	// When fetching, provides a link to fetch the _next_ chunk. If absent,
	// indicates there are no more chunks. This link is an absolute `path` to be
	// joined with your `$DATABRICKS_HOST`, and should be treated as an opaque
	// link. This is an alternative to using `next_chunk_index`.
	NextChunkInternalLink string
	// The number of rows within the result chunk.
	RowCount int64
	// The starting row offset within the result set.
	RowOffset int64

	ForceSendFields []string
}

func resultDataToPb(st *ResultData) (*resultDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultDataPb{}
	byteCountPb, err := identity(&st.ByteCount)
	if err != nil {
		return nil, err
	}
	if byteCountPb != nil {
		pb.ByteCount = *byteCountPb
	}

	chunkIndexPb, err := identity(&st.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexPb != nil {
		pb.ChunkIndex = *chunkIndexPb
	}

	var dataArrayPb [][]string
	for _, item := range st.DataArray {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataArrayPb = append(dataArrayPb, *itemPb)
		}
	}
	pb.DataArray = dataArrayPb

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

	nextChunkIndexPb, err := identity(&st.NextChunkIndex)
	if err != nil {
		return nil, err
	}
	if nextChunkIndexPb != nil {
		pb.NextChunkIndex = *nextChunkIndexPb
	}

	nextChunkInternalLinkPb, err := identity(&st.NextChunkInternalLink)
	if err != nil {
		return nil, err
	}
	if nextChunkInternalLinkPb != nil {
		pb.NextChunkInternalLink = *nextChunkInternalLinkPb
	}

	rowCountPb, err := identity(&st.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountPb != nil {
		pb.RowCount = *rowCountPb
	}

	rowOffsetPb, err := identity(&st.RowOffset)
	if err != nil {
		return nil, err
	}
	if rowOffsetPb != nil {
		pb.RowOffset = *rowOffsetPb
	}

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
	byteCountField, err := identity(&pb.ByteCount)
	if err != nil {
		return nil, err
	}
	if byteCountField != nil {
		st.ByteCount = *byteCountField
	}
	chunkIndexField, err := identity(&pb.ChunkIndex)
	if err != nil {
		return nil, err
	}
	if chunkIndexField != nil {
		st.ChunkIndex = *chunkIndexField
	}

	var dataArrayField [][]string
	for _, itemPb := range pb.DataArray {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataArrayField = append(dataArrayField, *item)
		}
	}
	st.DataArray = dataArrayField

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
	nextChunkIndexField, err := identity(&pb.NextChunkIndex)
	if err != nil {
		return nil, err
	}
	if nextChunkIndexField != nil {
		st.NextChunkIndex = *nextChunkIndexField
	}
	nextChunkInternalLinkField, err := identity(&pb.NextChunkInternalLink)
	if err != nil {
		return nil, err
	}
	if nextChunkInternalLinkField != nil {
		st.NextChunkInternalLink = *nextChunkInternalLinkField
	}
	rowCountField, err := identity(&pb.RowCount)
	if err != nil {
		return nil, err
	}
	if rowCountField != nil {
		st.RowCount = *rowCountField
	}
	rowOffsetField, err := identity(&pb.RowOffset)
	if err != nil {
		return nil, err
	}
	if rowOffsetField != nil {
		st.RowOffset = *rowOffsetField
	}

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
	Chunks []BaseChunkInfo

	Format Format
	// The schema is an ordered list of column descriptions.
	Schema *ResultSchema
	// The total number of bytes in the result set. This field is not available
	// when using `INLINE` disposition.
	TotalByteCount int64
	// The total number of chunks that the result set has been divided into.
	TotalChunkCount int
	// The total number of rows in the result set.
	TotalRowCount int64
	// Indicates whether the result is truncated due to `row_limit` or
	// `byte_limit`.
	Truncated bool

	ForceSendFields []string
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

	formatPb, err := identity(&st.Format)
	if err != nil {
		return nil, err
	}
	if formatPb != nil {
		pb.Format = *formatPb
	}

	schemaPb, err := resultSchemaToPb(st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = schemaPb
	}

	totalByteCountPb, err := identity(&st.TotalByteCount)
	if err != nil {
		return nil, err
	}
	if totalByteCountPb != nil {
		pb.TotalByteCount = *totalByteCountPb
	}

	totalChunkCountPb, err := identity(&st.TotalChunkCount)
	if err != nil {
		return nil, err
	}
	if totalChunkCountPb != nil {
		pb.TotalChunkCount = *totalChunkCountPb
	}

	totalRowCountPb, err := identity(&st.TotalRowCount)
	if err != nil {
		return nil, err
	}
	if totalRowCountPb != nil {
		pb.TotalRowCount = *totalRowCountPb
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
	formatField, err := identity(&pb.Format)
	if err != nil {
		return nil, err
	}
	if formatField != nil {
		st.Format = *formatField
	}
	schemaField, err := resultSchemaFromPb(pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = schemaField
	}
	totalByteCountField, err := identity(&pb.TotalByteCount)
	if err != nil {
		return nil, err
	}
	if totalByteCountField != nil {
		st.TotalByteCount = *totalByteCountField
	}
	totalChunkCountField, err := identity(&pb.TotalChunkCount)
	if err != nil {
		return nil, err
	}
	if totalChunkCountField != nil {
		st.TotalChunkCount = *totalChunkCountField
	}
	totalRowCountField, err := identity(&pb.TotalRowCount)
	if err != nil {
		return nil, err
	}
	if totalRowCountField != nil {
		st.TotalRowCount = *totalRowCountField
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

func (st *resultManifestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultManifestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The schema is an ordered list of column descriptions.
type ResultSchema struct {
	ColumnCount int

	Columns []ColumnInfo

	ForceSendFields []string
}

func resultSchemaToPb(st *ResultSchema) (*resultSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultSchemaPb{}
	columnCountPb, err := identity(&st.ColumnCount)
	if err != nil {
		return nil, err
	}
	if columnCountPb != nil {
		pb.ColumnCount = *columnCountPb
	}

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
	columnCountField, err := identity(&pb.ColumnCount)
	if err != nil {
		return nil, err
	}
	if columnCountField != nil {
		st.ColumnCount = *columnCountField
	}

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
	ErrorCode ServiceErrorCode
	// A brief summary of the error condition.
	Message string

	ForceSendFields []string
}

func serviceErrorToPb(st *ServiceError) (*serviceErrorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &serviceErrorPb{}
	errorCodePb, err := identity(&st.ErrorCode)
	if err != nil {
		return nil, err
	}
	if errorCodePb != nil {
		pb.ErrorCode = *errorCodePb
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
	errorCodeField, err := identity(&pb.ErrorCode)
	if err != nil {
		return nil, err
	}
	if errorCodeField != nil {
		st.ErrorCode = *errorCodeField
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
	AccessControlList []AccessControl
	// Object ID. The ACL for the object with this UUID is overwritten by this
	// request's POST content.
	ObjectId string
	// The type of object permission to set.
	ObjectType ObjectTypePlural
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

	return st, nil
}

type SetResponse struct {
	AccessControlList []AccessControl
	// An object's type and UUID, separated by a forward slash (/) character.
	ObjectId string
	// A singular noun object type.
	ObjectType ObjectType

	ForceSendFields []string
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

func (st *setResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetWorkspaceWarehouseConfigRequest struct {
	// Optional: Channel selection details
	Channel *Channel
	// Deprecated: Use sql_configuration_parameters
	ConfigParam *RepeatedEndpointConfPairs
	// Spark confs for external hive metastore configuration JSON serialized
	// size must be less than <= 512K
	DataAccessConfig []EndpointConfPair
	// List of Warehouse Types allowed in this workspace (limits allowed value
	// of the type field in CreateWarehouse and EditWarehouse). Note: Some types
	// cannot be disabled, they don't need to be specified in
	// SetWorkspaceWarehouseConfig. Note: Disabling a type may cause existing
	// warehouses to be converted to another type. Used by frontend to save
	// specific type availability in the warehouse create and edit form UI.
	EnabledWarehouseTypes []WarehouseTypePair
	// Deprecated: Use sql_configuration_parameters
	GlobalParam *RepeatedEndpointConfPairs
	// GCP only: Google Service Account used to pass to cluster to access Google
	// Cloud Storage
	GoogleServiceAccount string
	// AWS Only: Instance profile used to pass IAM role to the cluster
	InstanceProfileArn string
	// Security policy for warehouses
	SecurityPolicy SetWorkspaceWarehouseConfigRequestSecurityPolicy
	// SQL configuration parameters
	SqlConfigurationParameters *RepeatedEndpointConfPairs

	ForceSendFields []string
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

	googleServiceAccountPb, err := identity(&st.GoogleServiceAccount)
	if err != nil {
		return nil, err
	}
	if googleServiceAccountPb != nil {
		pb.GoogleServiceAccount = *googleServiceAccountPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	securityPolicyPb, err := identity(&st.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyPb != nil {
		pb.SecurityPolicy = *securityPolicyPb
	}

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
	googleServiceAccountField, err := identity(&pb.GoogleServiceAccount)
	if err != nil {
		return nil, err
	}
	if googleServiceAccountField != nil {
		st.GoogleServiceAccount = *googleServiceAccountField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	securityPolicyField, err := identity(&pb.SecurityPolicy)
	if err != nil {
		return nil, err
	}
	if securityPolicyField != nil {
		st.SecurityPolicy = *securityPolicyField
	}
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
	Id string
}

func startRequestToPb(st *StartRequest) (*startRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

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
	Name string
	// The data type, given as a string. For example: `INT`, `STRING`,
	// `DECIMAL(10,2)`. If no type is given the type is assumed to be `STRING`.
	// Complex types, such as `ARRAY`, `MAP`, and `STRUCT` are not supported.
	// For valid types, refer to the section [Data types] of the SQL language
	// reference.
	//
	// [Data types]: https://docs.databricks.com/sql/language-manual/functions/cast.html
	Type string
	// The value to substitute, represented as a string. If omitted, the value
	// is interpreted as NULL.
	Value string

	ForceSendFields []string
}

func statementParameterListItemToPb(st *StatementParameterListItem) (*statementParameterListItemPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &statementParameterListItemPb{}
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

func (st *statementParameterListItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st statementParameterListItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StatementResponse struct {
	// The result manifest provides schema and metadata for the result set.
	Manifest *ResultManifest

	Result *ResultData
	// The statement ID is returned upon successfully submitting a SQL
	// statement, and is a required reference for all subsequent calls.
	StatementId string
	// The status response includes execution state and if relevant, error
	// information.
	Status *StatementStatus

	ForceSendFields []string
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

	statementIdPb, err := identity(&st.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdPb != nil {
		pb.StatementId = *statementIdPb
	}

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
	statementIdField, err := identity(&pb.StatementId)
	if err != nil {
		return nil, err
	}
	if statementIdField != nil {
		st.StatementId = *statementIdField
	}
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
	Error *ServiceError
	// Statement execution state: - `PENDING`: waiting for warehouse -
	// `RUNNING`: running - `SUCCEEDED`: execution was successful, result data
	// available for fetch - `FAILED`: execution failed; reason for failure
	// described in accomanying error message - `CANCELED`: user canceled; can
	// come from explicit cancel call, or timeout with `on_wait_timeout=CANCEL`
	// - `CLOSED`: execution successful, and statement closed; result no longer
	// available for fetch
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

	statePb, err := identity(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

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
	stateField, err := identity(&pb.State)
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
	Id string
}

func stopRequestToPb(st *StopRequest) (*stopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

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
	Message SuccessMessage
}

func successToPb(st *Success) (*successPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &successPb{}
	messagePb, err := identity(&st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = *messagePb
	}

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
	messageField, err := identity(&pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = *messageField
	}

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
	Value string

	ForceSendFields []string
}

func textValueToPb(st *TextValue) (*textValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &textValuePb{}
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

func (st *textValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st textValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TimeRange struct {
	// The end time in milliseconds.
	EndTimeMs int64
	// The start time in milliseconds.
	StartTimeMs int64

	ForceSendFields []string
}

func timeRangeToPb(st *TimeRange) (*timeRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &timeRangePb{}
	endTimeMsPb, err := identity(&st.EndTimeMs)
	if err != nil {
		return nil, err
	}
	if endTimeMsPb != nil {
		pb.EndTimeMs = *endTimeMsPb
	}

	startTimeMsPb, err := identity(&st.StartTimeMs)
	if err != nil {
		return nil, err
	}
	if startTimeMsPb != nil {
		pb.StartTimeMs = *startTimeMsPb
	}

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
	endTimeMsField, err := identity(&pb.EndTimeMs)
	if err != nil {
		return nil, err
	}
	if endTimeMsField != nil {
		st.EndTimeMs = *endTimeMsField
	}
	startTimeMsField, err := identity(&pb.StartTimeMs)
	if err != nil {
		return nil, err
	}
	if startTimeMsField != nil {
		st.StartTimeMs = *startTimeMsField
	}

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
	NewOwner string

	ForceSendFields []string
}

func transferOwnershipObjectIdToPb(st *TransferOwnershipObjectId) (*transferOwnershipObjectIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transferOwnershipObjectIdPb{}
	newOwnerPb, err := identity(&st.NewOwner)
	if err != nil {
		return nil, err
	}
	if newOwnerPb != nil {
		pb.NewOwner = *newOwnerPb
	}

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
	newOwnerField, err := identity(&pb.NewOwner)
	if err != nil {
		return nil, err
	}
	if newOwnerField != nil {
		st.NewOwner = *newOwnerField
	}

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
	NewOwner string
	// The ID of the object on which to change ownership.
	ObjectId TransferOwnershipObjectId
	// The type of object on which to change ownership.
	ObjectType OwnableObjectType

	ForceSendFields []string
}

func transferOwnershipRequestToPb(st *TransferOwnershipRequest) (*transferOwnershipRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transferOwnershipRequestPb{}
	newOwnerPb, err := identity(&st.NewOwner)
	if err != nil {
		return nil, err
	}
	if newOwnerPb != nil {
		pb.NewOwner = *newOwnerPb
	}

	objectIdPb, err := transferOwnershipObjectIdToPb(&st.ObjectId)
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
	newOwnerField, err := identity(&pb.NewOwner)
	if err != nil {
		return nil, err
	}
	if newOwnerField != nil {
		st.NewOwner = *newOwnerField
	}
	objectIdField, err := transferOwnershipObjectIdFromPb(&pb.ObjectId)
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

func (st *transferOwnershipRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st transferOwnershipRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete an alert
type TrashAlertRequest struct {
	Id string
}

func trashAlertRequestToPb(st *TrashAlertRequest) (*trashAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashAlertRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

// Delete an alert
type TrashAlertV2Request struct {
	Id string
}

func trashAlertV2RequestToPb(st *TrashAlertV2Request) (*trashAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashAlertV2RequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

// Delete a query
type TrashQueryRequest struct {
	Id string
}

func trashQueryRequestToPb(st *TrashQueryRequest) (*trashQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashQueryRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
}

type UpdateAlertRequest struct {
	Alert *UpdateAlertRequestAlert

	Id string
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

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
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

type UpdateAlertRequestAlert struct {
	// Trigger conditions of the alert.
	Condition *AlertCondition
	// Custom body of alert notification, if it exists. See [here] for custom
	// templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomBody string
	// Custom subject of alert notification, if it exists. This can include
	// email subject entries and Slack notification headers, for example. See
	// [here] for custom templating instructions.
	//
	// [here]: https://docs.databricks.com/sql/user/alerts/index.html
	CustomSubject string
	// The display name of the alert.
	DisplayName string
	// Whether to notify alert subscribers when alert returns back to normal.
	NotifyOnOk bool
	// The owner's username. This field is set to "Unavailable" if the user has
	// been deleted.
	OwnerUserName string
	// UUID of the query attached to the alert.
	QueryId string
	// Number of seconds an alert must wait after being triggered to rearm
	// itself. After rearming, it can be triggered again. If 0 or not specified,
	// the alert will not be triggered again.
	SecondsToRetrigger int

	ForceSendFields []string
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

	customBodyPb, err := identity(&st.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyPb != nil {
		pb.CustomBody = *customBodyPb
	}

	customSubjectPb, err := identity(&st.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectPb != nil {
		pb.CustomSubject = *customSubjectPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	notifyOnOkPb, err := identity(&st.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkPb != nil {
		pb.NotifyOnOk = *notifyOnOkPb
	}

	ownerUserNamePb, err := identity(&st.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNamePb != nil {
		pb.OwnerUserName = *ownerUserNamePb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	secondsToRetriggerPb, err := identity(&st.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerPb != nil {
		pb.SecondsToRetrigger = *secondsToRetriggerPb
	}

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
	customBodyField, err := identity(&pb.CustomBody)
	if err != nil {
		return nil, err
	}
	if customBodyField != nil {
		st.CustomBody = *customBodyField
	}
	customSubjectField, err := identity(&pb.CustomSubject)
	if err != nil {
		return nil, err
	}
	if customSubjectField != nil {
		st.CustomSubject = *customSubjectField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	notifyOnOkField, err := identity(&pb.NotifyOnOk)
	if err != nil {
		return nil, err
	}
	if notifyOnOkField != nil {
		st.NotifyOnOk = *notifyOnOkField
	}
	ownerUserNameField, err := identity(&pb.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNameField != nil {
		st.OwnerUserName = *ownerUserNameField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	secondsToRetriggerField, err := identity(&pb.SecondsToRetrigger)
	if err != nil {
		return nil, err
	}
	if secondsToRetriggerField != nil {
		st.SecondsToRetrigger = *secondsToRetriggerField
	}

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
	Alert *AlertV2
	// UUID identifying the alert.
	Id string
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
	UpdateMask []string
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

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
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
	Id string

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
	UpdateMask string
}

func updateQueryRequestToPb(st *UpdateQueryRequest) (*updateQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateQueryRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	queryPb, err := updateQueryRequestQueryToPb(st.Query)
	if err != nil {
		return nil, err
	}
	if queryPb != nil {
		pb.Query = queryPb
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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	queryField, err := updateQueryRequestQueryFromPb(pb.Query)
	if err != nil {
		return nil, err
	}
	if queryField != nil {
		st.Query = queryField
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

type UpdateQueryRequestQuery struct {
	// Whether to apply a 1000 row limit to the query result.
	ApplyAutoLimit bool
	// Name of the catalog where this query will be executed.
	Catalog string
	// General description that conveys additional information about this query
	// such as usage notes.
	Description string
	// Display name of the query that appears in list views, widget headings,
	// and on the query page.
	DisplayName string
	// Username of the user that owns the query.
	OwnerUserName string
	// List of query parameter definitions.
	Parameters []QueryParameter
	// Text of the query to be run.
	QueryText string
	// Sets the "Run as" role for the object.
	RunAsMode RunAsMode
	// Name of the schema where this query will be executed.
	Schema string

	Tags []string
	// ID of the SQL warehouse attached to the query.
	WarehouseId string

	ForceSendFields []string
}

func updateQueryRequestQueryToPb(st *UpdateQueryRequestQuery) (*updateQueryRequestQueryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateQueryRequestQueryPb{}
	applyAutoLimitPb, err := identity(&st.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitPb != nil {
		pb.ApplyAutoLimit = *applyAutoLimitPb
	}

	catalogPb, err := identity(&st.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogPb != nil {
		pb.Catalog = *catalogPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	ownerUserNamePb, err := identity(&st.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNamePb != nil {
		pb.OwnerUserName = *ownerUserNamePb
	}

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

	queryTextPb, err := identity(&st.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextPb != nil {
		pb.QueryText = *queryTextPb
	}

	runAsModePb, err := identity(&st.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModePb != nil {
		pb.RunAsMode = *runAsModePb
	}

	schemaPb, err := identity(&st.Schema)
	if err != nil {
		return nil, err
	}
	if schemaPb != nil {
		pb.Schema = *schemaPb
	}

	var tagsPb []string
	for _, item := range st.Tags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

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
	applyAutoLimitField, err := identity(&pb.ApplyAutoLimit)
	if err != nil {
		return nil, err
	}
	if applyAutoLimitField != nil {
		st.ApplyAutoLimit = *applyAutoLimitField
	}
	catalogField, err := identity(&pb.Catalog)
	if err != nil {
		return nil, err
	}
	if catalogField != nil {
		st.Catalog = *catalogField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	ownerUserNameField, err := identity(&pb.OwnerUserName)
	if err != nil {
		return nil, err
	}
	if ownerUserNameField != nil {
		st.OwnerUserName = *ownerUserNameField
	}

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
	queryTextField, err := identity(&pb.QueryText)
	if err != nil {
		return nil, err
	}
	if queryTextField != nil {
		st.QueryText = *queryTextField
	}
	runAsModeField, err := identity(&pb.RunAsMode)
	if err != nil {
		return nil, err
	}
	if runAsModeField != nil {
		st.RunAsMode = *runAsModeField
	}
	schemaField, err := identity(&pb.Schema)
	if err != nil {
		return nil, err
	}
	if schemaField != nil {
		st.Schema = *schemaField
	}

	var tagsField []string
	for _, itemPb := range pb.Tags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
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
	Id string
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

	Visualization *UpdateVisualizationRequestVisualization
}

func updateVisualizationRequestToPb(st *UpdateVisualizationRequest) (*updateVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVisualizationRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	updateMaskPb, err := identity(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	updateMaskField, err := identity(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}
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
	DisplayName string
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string
	// The type of visualization: counter, table, funnel, and so on.
	Type string

	ForceSendFields []string
}

func updateVisualizationRequestVisualizationToPb(st *UpdateVisualizationRequestVisualization) (*updateVisualizationRequestVisualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVisualizationRequestVisualizationPb{}
	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	serializedOptionsPb, err := identity(&st.SerializedOptions)
	if err != nil {
		return nil, err
	}
	if serializedOptionsPb != nil {
		pb.SerializedOptions = *serializedOptionsPb
	}

	serializedQueryPlanPb, err := identity(&st.SerializedQueryPlan)
	if err != nil {
		return nil, err
	}
	if serializedQueryPlanPb != nil {
		pb.SerializedQueryPlan = *serializedQueryPlanPb
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
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	serializedOptionsField, err := identity(&pb.SerializedOptions)
	if err != nil {
		return nil, err
	}
	if serializedOptionsField != nil {
		st.SerializedOptions = *serializedOptionsField
	}
	serializedQueryPlanField, err := identity(&pb.SerializedQueryPlan)
	if err != nil {
		return nil, err
	}
	if serializedQueryPlanField != nil {
		st.SerializedQueryPlan = *serializedQueryPlanField
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

func (st *updateVisualizationRequestVisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateVisualizationRequestVisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type User struct {
	Email string

	Id int

	Name string

	ForceSendFields []string
}

func userToPb(st *User) (*userPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &userPb{}
	emailPb, err := identity(&st.Email)
	if err != nil {
		return nil, err
	}
	if emailPb != nil {
		pb.Email = *emailPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
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
	emailField, err := identity(&pb.Email)
	if err != nil {
		return nil, err
	}
	if emailField != nil {
		st.Email = *emailField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
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

func (st *userPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st userPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Visualization struct {
	// The timestamp indicating when the visualization was created.
	CreateTime string
	// The display name of the visualization.
	DisplayName string
	// UUID identifying the visualization.
	Id string
	// UUID of the query that the visualization is attached to.
	QueryId string
	// The visualization options varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying
	// visualization options directly.
	SerializedOptions string
	// The visualization query plan varies widely from one visualization type to
	// the next and is unsupported. Databricks does not recommend modifying the
	// visualization query plan directly.
	SerializedQueryPlan string
	// The type of visualization: counter, table, funnel, and so on.
	Type string
	// The timestamp indicating when the visualization was updated.
	UpdateTime string

	ForceSendFields []string
}

func visualizationToPb(st *Visualization) (*visualizationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &visualizationPb{}
	createTimePb, err := identity(&st.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	queryIdPb, err := identity(&st.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdPb != nil {
		pb.QueryId = *queryIdPb
	}

	serializedOptionsPb, err := identity(&st.SerializedOptions)
	if err != nil {
		return nil, err
	}
	if serializedOptionsPb != nil {
		pb.SerializedOptions = *serializedOptionsPb
	}

	serializedQueryPlanPb, err := identity(&st.SerializedQueryPlan)
	if err != nil {
		return nil, err
	}
	if serializedQueryPlanPb != nil {
		pb.SerializedQueryPlan = *serializedQueryPlanPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	updateTimePb, err := identity(&st.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

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
	createTimeField, err := identity(&pb.CreateTime)
	if err != nil {
		return nil, err
	}
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	queryIdField, err := identity(&pb.QueryId)
	if err != nil {
		return nil, err
	}
	if queryIdField != nil {
		st.QueryId = *queryIdField
	}
	serializedOptionsField, err := identity(&pb.SerializedOptions)
	if err != nil {
		return nil, err
	}
	if serializedOptionsField != nil {
		st.SerializedOptions = *serializedOptionsField
	}
	serializedQueryPlanField, err := identity(&pb.SerializedQueryPlan)
	if err != nil {
		return nil, err
	}
	if serializedQueryPlanField != nil {
		st.SerializedQueryPlan = *serializedQueryPlanField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}
	updateTimeField, err := identity(&pb.UpdateTime)
	if err != nil {
		return nil, err
	}
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

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
	GroupName string
	// Permission level
	PermissionLevel WarehousePermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func warehouseAccessControlRequestToPb(st *WarehouseAccessControlRequest) (*warehouseAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehouseAccessControlRequestPb{}
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

func (st *warehouseAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehouseAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehouseAccessControlResponse struct {
	// All permissions.
	AllPermissions []WarehousePermission
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

func (st *warehouseAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehouseAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel WarehousePermissionLevel

	ForceSendFields []string
}

func warehousePermissionToPb(st *WarehousePermission) (*warehousePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionPb{}
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
	AccessControlList []WarehouseAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
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

func (st *warehousePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehousePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel WarehousePermissionLevel

	ForceSendFields []string
}

func warehousePermissionsDescriptionToPb(st *WarehousePermissionsDescription) (*warehousePermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionsDescriptionPb{}
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

func (st *warehousePermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st warehousePermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionsRequest struct {
	AccessControlList []WarehouseAccessControlRequest
	// The SQL warehouse for which to get or manage permissions.
	WarehouseId string
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

	warehouseIdPb, err := identity(&st.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdPb != nil {
		pb.WarehouseId = *warehouseIdPb
	}

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
	warehouseIdField, err := identity(&pb.WarehouseId)
	if err != nil {
		return nil, err
	}
	if warehouseIdField != nil {
		st.WarehouseId = *warehouseIdField
	}

	return st, nil
}

type WarehouseTypePair struct {
	// If set to false the specific warehouse type will not be be allowed as a
	// value for warehouse_type in CreateWarehouse and EditWarehouse
	Enabled bool
	// Warehouse type: `PRO` or `CLASSIC`.
	WarehouseType WarehouseTypePairWarehouseType

	ForceSendFields []string
}

func warehouseTypePairToPb(st *WarehouseTypePair) (*warehouseTypePairPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehouseTypePairPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	warehouseTypePb, err := identity(&st.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypePb != nil {
		pb.WarehouseType = *warehouseTypePb
	}

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
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	warehouseTypeField, err := identity(&pb.WarehouseType)
	if err != nil {
		return nil, err
	}
	if warehouseTypeField != nil {
		st.WarehouseType = *warehouseTypeField
	}

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
	Id string

	Options *WidgetOptions
	// The visualization description API changes frequently and is unsupported.
	// You can duplicate a visualization by copying description objects received
	// _from the API_ and then using them to create a new one with a POST
	// request to the same endpoint. Databricks does not recommend constructing
	// ad-hoc visualizations entirely in JSON.
	Visualization *LegacyVisualization
	// Unused field.
	Width int

	ForceSendFields []string
}

func widgetToPb(st *Widget) (*widgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

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

	widthPb, err := identity(&st.Width)
	if err != nil {
		return nil, err
	}
	if widthPb != nil {
		pb.Width = *widthPb
	}

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
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
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
	widthField, err := identity(&pb.Width)
	if err != nil {
		return nil, err
	}
	if widthField != nil {
		st.Width = *widthField
	}

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
	CreatedAt string
	// Custom description of the widget
	Description string
	// Whether this widget is hidden on the dashboard.
	IsHidden bool
	// How parameters used by the visualization in this widget relate to other
	// widgets on the dashboard. Databricks does not recommend modifying this
	// definition in JSON.
	ParameterMappings any
	// Coordinates of this widget on a dashboard. This portion of the API
	// changes frequently and is unsupported.
	Position *WidgetPosition
	// Custom title of the widget
	Title string
	// Timestamp of the last time this object was updated.
	UpdatedAt string

	ForceSendFields []string
}

func widgetOptionsToPb(st *WidgetOptions) (*widgetOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetOptionsPb{}
	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	isHiddenPb, err := identity(&st.IsHidden)
	if err != nil {
		return nil, err
	}
	if isHiddenPb != nil {
		pb.IsHidden = *isHiddenPb
	}

	parameterMappingsPb, err := identity(&st.ParameterMappings)
	if err != nil {
		return nil, err
	}
	if parameterMappingsPb != nil {
		pb.ParameterMappings = *parameterMappingsPb
	}

	positionPb, err := widgetPositionToPb(st.Position)
	if err != nil {
		return nil, err
	}
	if positionPb != nil {
		pb.Position = positionPb
	}

	titlePb, err := identity(&st.Title)
	if err != nil {
		return nil, err
	}
	if titlePb != nil {
		pb.Title = *titlePb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

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
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	isHiddenField, err := identity(&pb.IsHidden)
	if err != nil {
		return nil, err
	}
	if isHiddenField != nil {
		st.IsHidden = *isHiddenField
	}
	parameterMappingsField, err := identity(&pb.ParameterMappings)
	if err != nil {
		return nil, err
	}
	if parameterMappingsField != nil {
		st.ParameterMappings = *parameterMappingsField
	}
	positionField, err := widgetPositionFromPb(pb.Position)
	if err != nil {
		return nil, err
	}
	if positionField != nil {
		st.Position = positionField
	}
	titleField, err := identity(&pb.Title)
	if err != nil {
		return nil, err
	}
	if titleField != nil {
		st.Title = *titleField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}

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
	AutoHeight bool
	// column in the dashboard grid. Values start with 0
	Col int
	// row in the dashboard grid. Values start with 0
	Row int
	// width of the widget measured in dashboard grid cells
	SizeX int
	// height of the widget measured in dashboard grid cells
	SizeY int

	ForceSendFields []string
}

func widgetPositionToPb(st *WidgetPosition) (*widgetPositionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetPositionPb{}
	autoHeightPb, err := identity(&st.AutoHeight)
	if err != nil {
		return nil, err
	}
	if autoHeightPb != nil {
		pb.AutoHeight = *autoHeightPb
	}

	colPb, err := identity(&st.Col)
	if err != nil {
		return nil, err
	}
	if colPb != nil {
		pb.Col = *colPb
	}

	rowPb, err := identity(&st.Row)
	if err != nil {
		return nil, err
	}
	if rowPb != nil {
		pb.Row = *rowPb
	}

	sizeXPb, err := identity(&st.SizeX)
	if err != nil {
		return nil, err
	}
	if sizeXPb != nil {
		pb.SizeX = *sizeXPb
	}

	sizeYPb, err := identity(&st.SizeY)
	if err != nil {
		return nil, err
	}
	if sizeYPb != nil {
		pb.SizeY = *sizeYPb
	}

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
	autoHeightField, err := identity(&pb.AutoHeight)
	if err != nil {
		return nil, err
	}
	if autoHeightField != nil {
		st.AutoHeight = *autoHeightField
	}
	colField, err := identity(&pb.Col)
	if err != nil {
		return nil, err
	}
	if colField != nil {
		st.Col = *colField
	}
	rowField, err := identity(&pb.Row)
	if err != nil {
		return nil, err
	}
	if rowField != nil {
		st.Row = *rowField
	}
	sizeXField, err := identity(&pb.SizeX)
	if err != nil {
		return nil, err
	}
	if sizeXField != nil {
		st.SizeX = *sizeXField
	}
	sizeYField, err := identity(&pb.SizeY)
	if err != nil {
		return nil, err
	}
	if sizeYField != nil {
		st.SizeY = *sizeYField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *widgetPositionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st widgetPositionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
