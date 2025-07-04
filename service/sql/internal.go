// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sql

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

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

type accessControlPb struct {
	GroupName       string          `json:"group_name,omitempty"`
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	UserName        string          `json:"user_name,omitempty"`

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

func alertToPb(st *Alert) (*alertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertPb{}
	pb.Condition = st.Condition
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

type alertPb struct {
	Condition          *AlertCondition `json:"condition,omitempty"`
	CreateTime         string          `json:"create_time,omitempty"`
	CustomBody         string          `json:"custom_body,omitempty"`
	CustomSubject      string          `json:"custom_subject,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	Id                 string          `json:"id,omitempty"`
	LifecycleState     LifecycleState  `json:"lifecycle_state,omitempty"`
	NotifyOnOk         bool            `json:"notify_on_ok,omitempty"`
	OwnerUserName      string          `json:"owner_user_name,omitempty"`
	ParentPath         string          `json:"parent_path,omitempty"`
	QueryId            string          `json:"query_id,omitempty"`
	SecondsToRetrigger int             `json:"seconds_to_retrigger,omitempty"`
	State              AlertState      `json:"state,omitempty"`
	TriggerTime        string          `json:"trigger_time,omitempty"`
	UpdateTime         string          `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertFromPb(pb *alertPb) (*Alert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Alert{}
	st.Condition = pb.Condition
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

func alertConditionToPb(st *AlertCondition) (*alertConditionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConditionPb{}
	pb.EmptyResultState = st.EmptyResultState
	pb.Op = st.Op
	pb.Operand = st.Operand
	pb.Threshold = st.Threshold

	return pb, nil
}

type alertConditionPb struct {
	EmptyResultState AlertState               `json:"empty_result_state,omitempty"`
	Op               AlertOperator            `json:"op,omitempty"`
	Operand          *AlertConditionOperand   `json:"operand,omitempty"`
	Threshold        *AlertConditionThreshold `json:"threshold,omitempty"`
}

func alertConditionFromPb(pb *alertConditionPb) (*AlertCondition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertCondition{}
	st.EmptyResultState = pb.EmptyResultState
	st.Op = pb.Op
	st.Operand = pb.Operand
	st.Threshold = pb.Threshold

	return st, nil
}

func alertConditionOperandToPb(st *AlertConditionOperand) (*alertConditionOperandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConditionOperandPb{}
	pb.Column = st.Column

	return pb, nil
}

type alertConditionOperandPb struct {
	Column *AlertOperandColumn `json:"column,omitempty"`
}

func alertConditionOperandFromPb(pb *alertConditionOperandPb) (*AlertConditionOperand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConditionOperand{}
	st.Column = pb.Column

	return st, nil
}

func alertConditionThresholdToPb(st *AlertConditionThreshold) (*alertConditionThresholdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConditionThresholdPb{}
	pb.Value = st.Value

	return pb, nil
}

type alertConditionThresholdPb struct {
	Value *AlertOperandValue `json:"value,omitempty"`
}

func alertConditionThresholdFromPb(pb *alertConditionThresholdPb) (*AlertConditionThreshold, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConditionThreshold{}
	st.Value = pb.Value

	return st, nil
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

type alertOperandValuePb struct {
	BoolValue   bool    `json:"bool_value,omitempty"`
	DoubleValue float64 `json:"double_value,omitempty"`
	StringValue string  `json:"string_value,omitempty"`

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

type alertOptionsPb struct {
	Column           string                       `json:"column"`
	CustomBody       string                       `json:"custom_body,omitempty"`
	CustomSubject    string                       `json:"custom_subject,omitempty"`
	EmptyResultState AlertOptionsEmptyResultState `json:"empty_result_state,omitempty"`
	Muted            bool                         `json:"muted,omitempty"`
	Op               string                       `json:"op"`
	Value            any                          `json:"value"`

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
	pb.Options = st.Options
	pb.Query = st.Query
	pb.Tags = st.Tags
	pb.UpdatedAt = st.UpdatedAt
	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type alertQueryPb struct {
	CreatedAt    string        `json:"created_at,omitempty"`
	DataSourceId string        `json:"data_source_id,omitempty"`
	Description  string        `json:"description,omitempty"`
	Id           string        `json:"id,omitempty"`
	IsArchived   bool          `json:"is_archived,omitempty"`
	IsDraft      bool          `json:"is_draft,omitempty"`
	IsSafe       bool          `json:"is_safe,omitempty"`
	Name         string        `json:"name,omitempty"`
	Options      *QueryOptions `json:"options,omitempty"`
	Query        string        `json:"query,omitempty"`
	Tags         []string      `json:"tags,omitempty"`
	UpdatedAt    string        `json:"updated_at,omitempty"`
	UserId       int           `json:"user_id,omitempty"`

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
	st.Options = pb.Options
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

func alertV2ToPb(st *AlertV2) (*alertV2Pb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2Pb{}
	pb.CreateTime = st.CreateTime
	pb.CustomDescription = st.CustomDescription
	pb.CustomSummary = st.CustomSummary
	pb.DisplayName = st.DisplayName
	pb.Evaluation = st.Evaluation
	pb.Id = st.Id
	pb.LifecycleState = st.LifecycleState
	pb.OwnerUserName = st.OwnerUserName
	pb.ParentPath = st.ParentPath
	pb.QueryText = st.QueryText
	pb.RunAsUserName = st.RunAsUserName
	pb.Schedule = st.Schedule
	pb.UpdateTime = st.UpdateTime
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type alertV2Pb struct {
	CreateTime        string             `json:"create_time,omitempty"`
	CustomDescription string             `json:"custom_description,omitempty"`
	CustomSummary     string             `json:"custom_summary,omitempty"`
	DisplayName       string             `json:"display_name,omitempty"`
	Evaluation        *AlertV2Evaluation `json:"evaluation,omitempty"`
	Id                string             `json:"id,omitempty"`
	LifecycleState    LifecycleState     `json:"lifecycle_state,omitempty"`
	OwnerUserName     string             `json:"owner_user_name,omitempty"`
	ParentPath        string             `json:"parent_path,omitempty"`
	QueryText         string             `json:"query_text,omitempty"`
	RunAsUserName     string             `json:"run_as_user_name,omitempty"`
	Schedule          *CronSchedule      `json:"schedule,omitempty"`
	UpdateTime        string             `json:"update_time,omitempty"`
	WarehouseId       string             `json:"warehouse_id,omitempty"`

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
	st.Evaluation = pb.Evaluation
	st.Id = pb.Id
	st.LifecycleState = pb.LifecycleState
	st.OwnerUserName = pb.OwnerUserName
	st.ParentPath = pb.ParentPath
	st.QueryText = pb.QueryText
	st.RunAsUserName = pb.RunAsUserName
	st.Schedule = pb.Schedule
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

func alertV2EvaluationToPb(st *AlertV2Evaluation) (*alertV2EvaluationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2EvaluationPb{}
	pb.ComparisonOperator = st.ComparisonOperator
	pb.EmptyResultState = st.EmptyResultState
	pb.LastEvaluatedAt = st.LastEvaluatedAt
	pb.Notification = st.Notification
	pb.Source = st.Source
	pb.State = st.State
	pb.Threshold = st.Threshold

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type alertV2EvaluationPb struct {
	ComparisonOperator ComparisonOperator    `json:"comparison_operator,omitempty"`
	EmptyResultState   AlertEvaluationState  `json:"empty_result_state,omitempty"`
	LastEvaluatedAt    string                `json:"last_evaluated_at,omitempty"`
	Notification       *AlertV2Notification  `json:"notification,omitempty"`
	Source             *AlertV2OperandColumn `json:"source,omitempty"`
	State              AlertEvaluationState  `json:"state,omitempty"`
	Threshold          *AlertV2Operand       `json:"threshold,omitempty"`

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
	st.Notification = pb.Notification
	st.Source = pb.Source
	st.State = pb.State
	st.Threshold = pb.Threshold

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertV2EvaluationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2EvaluationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func alertV2NotificationToPb(st *AlertV2Notification) (*alertV2NotificationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2NotificationPb{}
	pb.NotifyOnOk = st.NotifyOnOk
	pb.RetriggerSeconds = st.RetriggerSeconds
	pb.Subscriptions = st.Subscriptions

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type alertV2NotificationPb struct {
	NotifyOnOk       bool                  `json:"notify_on_ok,omitempty"`
	RetriggerSeconds int                   `json:"retrigger_seconds,omitempty"`
	Subscriptions    []AlertV2Subscription `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertV2NotificationFromPb(pb *alertV2NotificationPb) (*AlertV2Notification, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Notification{}
	st.NotifyOnOk = pb.NotifyOnOk
	st.RetriggerSeconds = pb.RetriggerSeconds
	st.Subscriptions = pb.Subscriptions

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertV2NotificationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertV2NotificationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func alertV2OperandToPb(st *AlertV2Operand) (*alertV2OperandPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertV2OperandPb{}
	pb.Column = st.Column
	pb.Value = st.Value

	return pb, nil
}

type alertV2OperandPb struct {
	Column *AlertV2OperandColumn `json:"column,omitempty"`
	Value  *AlertV2OperandValue  `json:"value,omitempty"`
}

func alertV2OperandFromPb(pb *alertV2OperandPb) (*AlertV2Operand, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertV2Operand{}
	st.Column = pb.Column
	st.Value = pb.Value

	return st, nil
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

type alertV2OperandColumnPb struct {
	Aggregation Aggregation `json:"aggregation,omitempty"`
	Display     string      `json:"display,omitempty"`
	Name        string      `json:"name,omitempty"`

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

type alertV2OperandValuePb struct {
	BoolValue   bool    `json:"bool_value,omitempty"`
	DoubleValue float64 `json:"double_value,omitempty"`
	StringValue string  `json:"string_value,omitempty"`

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

type alertV2SubscriptionPb struct {
	DestinationId string `json:"destination_id,omitempty"`
	UserEmail     string `json:"user_email,omitempty"`

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

type baseChunkInfoPb struct {
	ByteCount  int64 `json:"byte_count,omitempty"`
	ChunkIndex int   `json:"chunk_index,omitempty"`
	RowCount   int64 `json:"row_count,omitempty"`
	RowOffset  int64 `json:"row_offset,omitempty"`

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

func cancelExecutionRequestToPb(st *CancelExecutionRequest) (*cancelExecutionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelExecutionRequestPb{}
	pb.StatementId = st.StatementId

	return pb, nil
}

type cancelExecutionRequestPb struct {
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

type channelPb struct {
	DbsqlVersion string      `json:"dbsql_version,omitempty"`
	Name         ChannelName `json:"name,omitempty"`

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

type channelInfoPb struct {
	DbsqlVersion string      `json:"dbsql_version,omitempty"`
	Name         ChannelName `json:"name,omitempty"`

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

type clientConfigPb struct {
	AllowCustomJsVisualizations   bool   `json:"allow_custom_js_visualizations,omitempty"`
	AllowDownloads                bool   `json:"allow_downloads,omitempty"`
	AllowExternalShares           bool   `json:"allow_external_shares,omitempty"`
	AllowSubscriptions            bool   `json:"allow_subscriptions,omitempty"`
	DateFormat                    string `json:"date_format,omitempty"`
	DateTimeFormat                string `json:"date_time_format,omitempty"`
	DisablePublish                bool   `json:"disable_publish,omitempty"`
	EnableLegacyAutodetectTypes   bool   `json:"enable_legacy_autodetect_types,omitempty"`
	FeatureShowPermissionsControl bool   `json:"feature_show_permissions_control,omitempty"`
	HidePlotlyModeBar             bool   `json:"hide_plotly_mode_bar,omitempty"`

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

type columnInfoPb struct {
	Name             string             `json:"name,omitempty"`
	Position         int                `json:"position,omitempty"`
	TypeIntervalType string             `json:"type_interval_type,omitempty"`
	TypeName         ColumnInfoTypeName `json:"type_name,omitempty"`
	TypePrecision    int                `json:"type_precision,omitempty"`
	TypeScale        int                `json:"type_scale,omitempty"`
	TypeText         string             `json:"type_text,omitempty"`

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

func createAlertToPb(st *CreateAlert) (*createAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertPb{}
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Parent = st.Parent
	pb.QueryId = st.QueryId
	pb.Rearm = st.Rearm

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createAlertPb struct {
	Name    string       `json:"name"`
	Options AlertOptions `json:"options"`
	Parent  string       `json:"parent,omitempty"`
	QueryId string       `json:"query_id"`
	Rearm   int          `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAlertFromPb(pb *createAlertPb) (*CreateAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlert{}
	st.Name = pb.Name
	st.Options = pb.Options
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

func createAlertRequestToPb(st *CreateAlertRequest) (*createAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertRequestPb{}
	pb.Alert = st.Alert
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createAlertRequestPb struct {
	Alert                  *CreateAlertRequestAlert `json:"alert,omitempty"`
	AutoResolveDisplayName bool                     `json:"auto_resolve_display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAlertRequestFromPb(pb *createAlertRequestPb) (*CreateAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertRequest{}
	st.Alert = pb.Alert
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

func createAlertRequestAlertToPb(st *CreateAlertRequestAlert) (*createAlertRequestAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertRequestAlertPb{}
	pb.Condition = st.Condition
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

type createAlertRequestAlertPb struct {
	Condition          *AlertCondition `json:"condition,omitempty"`
	CustomBody         string          `json:"custom_body,omitempty"`
	CustomSubject      string          `json:"custom_subject,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	NotifyOnOk         bool            `json:"notify_on_ok,omitempty"`
	ParentPath         string          `json:"parent_path,omitempty"`
	QueryId            string          `json:"query_id,omitempty"`
	SecondsToRetrigger int             `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAlertRequestAlertFromPb(pb *createAlertRequestAlertPb) (*CreateAlertRequestAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertRequestAlert{}
	st.Condition = pb.Condition
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

func createAlertV2RequestToPb(st *CreateAlertV2Request) (*createAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAlertV2RequestPb{}
	pb.Alert = st.Alert

	return pb, nil
}

type createAlertV2RequestPb struct {
	Alert AlertV2 `json:"alert"`
}

func createAlertV2RequestFromPb(pb *createAlertV2RequestPb) (*CreateAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAlertV2Request{}
	st.Alert = pb.Alert

	return st, nil
}

func createQueryRequestToPb(st *CreateQueryRequest) (*createQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQueryRequestPb{}
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName
	pb.Query = st.Query

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createQueryRequestPb struct {
	AutoResolveDisplayName bool                     `json:"auto_resolve_display_name,omitempty"`
	Query                  *CreateQueryRequestQuery `json:"query,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createQueryRequestFromPb(pb *createQueryRequestPb) (*CreateQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateQueryRequest{}
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName
	st.Query = pb.Query

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	pb.Parameters = st.Parameters
	pb.ParentPath = st.ParentPath
	pb.QueryText = st.QueryText
	pb.RunAsMode = st.RunAsMode
	pb.Schema = st.Schema
	pb.Tags = st.Tags
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createQueryRequestQueryPb struct {
	ApplyAutoLimit bool             `json:"apply_auto_limit,omitempty"`
	Catalog        string           `json:"catalog,omitempty"`
	Description    string           `json:"description,omitempty"`
	DisplayName    string           `json:"display_name,omitempty"`
	Parameters     []QueryParameter `json:"parameters,omitempty"`
	ParentPath     string           `json:"parent_path,omitempty"`
	QueryText      string           `json:"query_text,omitempty"`
	RunAsMode      RunAsMode        `json:"run_as_mode,omitempty"`
	Schema         string           `json:"schema,omitempty"`
	Tags           []string         `json:"tags,omitempty"`
	WarehouseId    string           `json:"warehouse_id,omitempty"`

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
	st.Parameters = pb.Parameters
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

type createQueryVisualizationsLegacyRequestPb struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Options     any    `json:"options"`
	QueryId     string `json:"query_id"`
	Type        string `json:"type"`

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

func createVisualizationRequestToPb(st *CreateVisualizationRequest) (*createVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createVisualizationRequestPb{}
	pb.Visualization = st.Visualization

	return pb, nil
}

type createVisualizationRequestPb struct {
	Visualization *CreateVisualizationRequestVisualization `json:"visualization,omitempty"`
}

func createVisualizationRequestFromPb(pb *createVisualizationRequestPb) (*CreateVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateVisualizationRequest{}
	st.Visualization = pb.Visualization

	return st, nil
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

type createVisualizationRequestVisualizationPb struct {
	DisplayName         string `json:"display_name,omitempty"`
	QueryId             string `json:"query_id,omitempty"`
	SerializedOptions   string `json:"serialized_options,omitempty"`
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	Type                string `json:"type,omitempty"`

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

func createWarehouseRequestToPb(st *CreateWarehouseRequest) (*createWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWarehouseRequestPb{}
	pb.AutoStopMins = st.AutoStopMins
	pb.Channel = st.Channel
	pb.ClusterSize = st.ClusterSize
	pb.CreatorName = st.CreatorName
	pb.EnablePhoton = st.EnablePhoton
	pb.EnableServerlessCompute = st.EnableServerlessCompute
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.MaxNumClusters = st.MaxNumClusters
	pb.MinNumClusters = st.MinNumClusters
	pb.Name = st.Name
	pb.SpotInstancePolicy = st.SpotInstancePolicy
	pb.Tags = st.Tags
	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createWarehouseRequestPb struct {
	AutoStopMins            int                                 `json:"auto_stop_mins,omitempty"`
	Channel                 *Channel                            `json:"channel,omitempty"`
	ClusterSize             string                              `json:"cluster_size,omitempty"`
	CreatorName             string                              `json:"creator_name,omitempty"`
	EnablePhoton            bool                                `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                                `json:"enable_serverless_compute,omitempty"`
	InstanceProfileArn      string                              `json:"instance_profile_arn,omitempty"`
	MaxNumClusters          int                                 `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                                 `json:"min_num_clusters,omitempty"`
	Name                    string                              `json:"name,omitempty"`
	SpotInstancePolicy      SpotInstancePolicy                  `json:"spot_instance_policy,omitempty"`
	Tags                    *EndpointTags                       `json:"tags,omitempty"`
	WarehouseType           CreateWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createWarehouseRequestFromPb(pb *createWarehouseRequestPb) (*CreateWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWarehouseRequest{}
	st.AutoStopMins = pb.AutoStopMins
	st.Channel = pb.Channel
	st.ClusterSize = pb.ClusterSize
	st.CreatorName = pb.CreatorName
	st.EnablePhoton = pb.EnablePhoton
	st.EnableServerlessCompute = pb.EnableServerlessCompute
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxNumClusters = pb.MaxNumClusters
	st.MinNumClusters = pb.MinNumClusters
	st.Name = pb.Name
	st.SpotInstancePolicy = pb.SpotInstancePolicy
	st.Tags = pb.Tags
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

func createWarehouseResponseToPb(st *CreateWarehouseResponse) (*createWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWarehouseResponsePb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createWarehouseResponsePb struct {
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

func createWidgetToPb(st *CreateWidget) (*createWidgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWidgetPb{}
	pb.DashboardId = st.DashboardId
	pb.Options = st.Options
	pb.Text = st.Text
	pb.VisualizationId = st.VisualizationId
	pb.Width = st.Width

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createWidgetPb struct {
	DashboardId     string        `json:"dashboard_id"`
	Options         WidgetOptions `json:"options"`
	Text            string        `json:"text,omitempty"`
	VisualizationId string        `json:"visualization_id,omitempty"`
	Width           int           `json:"width"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createWidgetFromPb(pb *createWidgetPb) (*CreateWidget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWidget{}
	st.DashboardId = pb.DashboardId
	st.Options = pb.Options
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

type cronSchedulePb struct {
	PauseStatus        SchedulePauseStatus `json:"pause_status,omitempty"`
	QuartzCronSchedule string              `json:"quartz_cron_schedule,omitempty"`
	TimezoneId         string              `json:"timezone_id,omitempty"`

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
	pb.Options = st.Options
	pb.Parent = st.Parent
	pb.PermissionTier = st.PermissionTier
	pb.Slug = st.Slug
	pb.Tags = st.Tags
	pb.UpdatedAt = st.UpdatedAt
	pb.User = st.User
	pb.UserId = st.UserId
	pb.Widgets = st.Widgets

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dashboardPb struct {
	CanEdit                 bool              `json:"can_edit,omitempty"`
	CreatedAt               string            `json:"created_at,omitempty"`
	DashboardFiltersEnabled bool              `json:"dashboard_filters_enabled,omitempty"`
	Id                      string            `json:"id,omitempty"`
	IsArchived              bool              `json:"is_archived,omitempty"`
	IsDraft                 bool              `json:"is_draft,omitempty"`
	IsFavorite              bool              `json:"is_favorite,omitempty"`
	Name                    string            `json:"name,omitempty"`
	Options                 *DashboardOptions `json:"options,omitempty"`
	Parent                  string            `json:"parent,omitempty"`
	PermissionTier          PermissionLevel   `json:"permission_tier,omitempty"`
	Slug                    string            `json:"slug,omitempty"`
	Tags                    []string          `json:"tags,omitempty"`
	UpdatedAt               string            `json:"updated_at,omitempty"`
	User                    *User             `json:"user,omitempty"`
	UserId                  int               `json:"user_id,omitempty"`
	Widgets                 []Widget          `json:"widgets,omitempty"`

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
	st.Options = pb.Options
	st.Parent = pb.Parent
	st.PermissionTier = pb.PermissionTier
	st.Slug = pb.Slug
	st.Tags = pb.Tags
	st.UpdatedAt = pb.UpdatedAt
	st.User = pb.User
	st.UserId = pb.UserId
	st.Widgets = pb.Widgets

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type dashboardEditContentPb struct {
	DashboardId string    `json:"-" url:"-"`
	Name        string    `json:"name,omitempty"`
	RunAsRole   RunAsRole `json:"run_as_role,omitempty"`
	Tags        []string  `json:"tags,omitempty"`

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

func dashboardOptionsToPb(st *DashboardOptions) (*dashboardOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardOptionsPb{}
	pb.MovedToTrashAt = st.MovedToTrashAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dashboardOptionsPb struct {
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

type dashboardPostContentPb struct {
	DashboardFiltersEnabled bool      `json:"dashboard_filters_enabled,omitempty"`
	IsFavorite              bool      `json:"is_favorite,omitempty"`
	Name                    string    `json:"name"`
	Parent                  string    `json:"parent,omitempty"`
	RunAsRole               RunAsRole `json:"run_as_role,omitempty"`
	Tags                    []string  `json:"tags,omitempty"`

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

type dataSourcePb struct {
	Id                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	PauseReason       string `json:"pause_reason,omitempty"`
	Paused            int    `json:"paused,omitempty"`
	SupportsAutoLimit bool   `json:"supports_auto_limit,omitempty"`
	Syntax            string `json:"syntax,omitempty"`
	Type              string `json:"type,omitempty"`
	ViewOnly          bool   `json:"view_only,omitempty"`
	WarehouseId       string `json:"warehouse_id,omitempty"`

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

func dateRangeToPb(st *DateRange) (*dateRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dateRangePb{}
	pb.End = st.End
	pb.Start = st.Start

	return pb, nil
}

type dateRangePb struct {
	End   string `json:"end"`
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

func dateRangeValueToPb(st *DateRangeValue) (*dateRangeValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dateRangeValuePb{}
	pb.DateRangeValue = st.DateRangeValue
	pb.DynamicDateRangeValue = st.DynamicDateRangeValue
	pb.Precision = st.Precision
	pb.StartDayOfWeek = st.StartDayOfWeek

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dateRangeValuePb struct {
	DateRangeValue        *DateRange                     `json:"date_range_value,omitempty"`
	DynamicDateRangeValue DateRangeValueDynamicDateRange `json:"dynamic_date_range_value,omitempty"`
	Precision             DatePrecision                  `json:"precision,omitempty"`
	StartDayOfWeek        int                            `json:"start_day_of_week,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dateRangeValueFromPb(pb *dateRangeValuePb) (*DateRangeValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DateRangeValue{}
	st.DateRangeValue = pb.DateRangeValue
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

type dateValuePb struct {
	DateValue        string               `json:"date_value,omitempty"`
	DynamicDateValue DateValueDynamicDate `json:"dynamic_date_value,omitempty"`
	Precision        DatePrecision        `json:"precision,omitempty"`

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

func deleteAlertsLegacyRequestToPb(st *DeleteAlertsLegacyRequest) (*deleteAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAlertsLegacyRequestPb{}
	pb.AlertId = st.AlertId

	return pb, nil
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

func deleteDashboardRequestToPb(st *DeleteDashboardRequest) (*deleteDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
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

func deleteDashboardWidgetRequestToPb(st *DeleteDashboardWidgetRequest) (*deleteDashboardWidgetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardWidgetRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteDashboardWidgetRequestPb struct {
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

func deleteQueriesLegacyRequestToPb(st *DeleteQueriesLegacyRequest) (*deleteQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
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

func deleteQueryVisualizationsLegacyRequestToPb(st *DeleteQueryVisualizationsLegacyRequest) (*deleteQueryVisualizationsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQueryVisualizationsLegacyRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteQueryVisualizationsLegacyRequestPb struct {
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

func deleteVisualizationRequestToPb(st *DeleteVisualizationRequest) (*deleteVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteVisualizationRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

func deleteWarehouseRequestToPb(st *DeleteWarehouseRequest) (*deleteWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWarehouseRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteWarehouseRequestPb struct {
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

func editAlertToPb(st *EditAlert) (*editAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editAlertPb{}
	pb.AlertId = st.AlertId
	pb.Name = st.Name
	pb.Options = st.Options
	pb.QueryId = st.QueryId
	pb.Rearm = st.Rearm

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type editAlertPb struct {
	AlertId string       `json:"-" url:"-"`
	Name    string       `json:"name"`
	Options AlertOptions `json:"options"`
	QueryId string       `json:"query_id"`
	Rearm   int          `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editAlertFromPb(pb *editAlertPb) (*EditAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditAlert{}
	st.AlertId = pb.AlertId
	st.Name = pb.Name
	st.Options = pb.Options
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

func editWarehouseRequestToPb(st *EditWarehouseRequest) (*editWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editWarehouseRequestPb{}
	pb.AutoStopMins = st.AutoStopMins
	pb.Channel = st.Channel
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
	pb.Tags = st.Tags
	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type editWarehouseRequestPb struct {
	AutoStopMins            int                               `json:"auto_stop_mins,omitempty"`
	Channel                 *Channel                          `json:"channel,omitempty"`
	ClusterSize             string                            `json:"cluster_size,omitempty"`
	CreatorName             string                            `json:"creator_name,omitempty"`
	EnablePhoton            bool                              `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                              `json:"enable_serverless_compute,omitempty"`
	Id                      string                            `json:"-" url:"-"`
	InstanceProfileArn      string                            `json:"instance_profile_arn,omitempty"`
	MaxNumClusters          int                               `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                               `json:"min_num_clusters,omitempty"`
	Name                    string                            `json:"name,omitempty"`
	SpotInstancePolicy      SpotInstancePolicy                `json:"spot_instance_policy,omitempty"`
	Tags                    *EndpointTags                     `json:"tags,omitempty"`
	WarehouseType           EditWarehouseRequestWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editWarehouseRequestFromPb(pb *editWarehouseRequestPb) (*EditWarehouseRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditWarehouseRequest{}
	st.AutoStopMins = pb.AutoStopMins
	st.Channel = pb.Channel
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
	st.Tags = pb.Tags
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

type endpointConfPairPb struct {
	Key   string `json:"key,omitempty"`
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

func endpointHealthToPb(st *EndpointHealth) (*endpointHealthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointHealthPb{}
	pb.Details = st.Details
	pb.FailureReason = st.FailureReason
	pb.Message = st.Message
	pb.Status = st.Status
	pb.Summary = st.Summary

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type endpointHealthPb struct {
	Details       string             `json:"details,omitempty"`
	FailureReason *TerminationReason `json:"failure_reason,omitempty"`
	Message       string             `json:"message,omitempty"`
	Status        Status             `json:"status,omitempty"`
	Summary       string             `json:"summary,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointHealthFromPb(pb *endpointHealthPb) (*EndpointHealth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointHealth{}
	st.Details = pb.Details
	st.FailureReason = pb.FailureReason
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

func endpointInfoToPb(st *EndpointInfo) (*endpointInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointInfoPb{}
	pb.AutoStopMins = st.AutoStopMins
	pb.Channel = st.Channel
	pb.ClusterSize = st.ClusterSize
	pb.CreatorName = st.CreatorName
	pb.EnablePhoton = st.EnablePhoton
	pb.EnableServerlessCompute = st.EnableServerlessCompute
	pb.Health = st.Health
	pb.Id = st.Id
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.JdbcUrl = st.JdbcUrl
	pb.MaxNumClusters = st.MaxNumClusters
	pb.MinNumClusters = st.MinNumClusters
	pb.Name = st.Name
	pb.NumActiveSessions = st.NumActiveSessions
	pb.NumClusters = st.NumClusters
	pb.OdbcParams = st.OdbcParams
	pb.SpotInstancePolicy = st.SpotInstancePolicy
	pb.State = st.State
	pb.Tags = st.Tags
	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type endpointInfoPb struct {
	AutoStopMins            int                       `json:"auto_stop_mins,omitempty"`
	Channel                 *Channel                  `json:"channel,omitempty"`
	ClusterSize             string                    `json:"cluster_size,omitempty"`
	CreatorName             string                    `json:"creator_name,omitempty"`
	EnablePhoton            bool                      `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                      `json:"enable_serverless_compute,omitempty"`
	Health                  *EndpointHealth           `json:"health,omitempty"`
	Id                      string                    `json:"id,omitempty"`
	InstanceProfileArn      string                    `json:"instance_profile_arn,omitempty"`
	JdbcUrl                 string                    `json:"jdbc_url,omitempty"`
	MaxNumClusters          int                       `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                       `json:"min_num_clusters,omitempty"`
	Name                    string                    `json:"name,omitempty"`
	NumActiveSessions       int64                     `json:"num_active_sessions,omitempty"`
	NumClusters             int                       `json:"num_clusters,omitempty"`
	OdbcParams              *OdbcParams               `json:"odbc_params,omitempty"`
	SpotInstancePolicy      SpotInstancePolicy        `json:"spot_instance_policy,omitempty"`
	State                   State                     `json:"state,omitempty"`
	Tags                    *EndpointTags             `json:"tags,omitempty"`
	WarehouseType           EndpointInfoWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointInfoFromPb(pb *endpointInfoPb) (*EndpointInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointInfo{}
	st.AutoStopMins = pb.AutoStopMins
	st.Channel = pb.Channel
	st.ClusterSize = pb.ClusterSize
	st.CreatorName = pb.CreatorName
	st.EnablePhoton = pb.EnablePhoton
	st.EnableServerlessCompute = pb.EnableServerlessCompute
	st.Health = pb.Health
	st.Id = pb.Id
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.JdbcUrl = pb.JdbcUrl
	st.MaxNumClusters = pb.MaxNumClusters
	st.MinNumClusters = pb.MinNumClusters
	st.Name = pb.Name
	st.NumActiveSessions = pb.NumActiveSessions
	st.NumClusters = pb.NumClusters
	st.OdbcParams = pb.OdbcParams
	st.SpotInstancePolicy = pb.SpotInstancePolicy
	st.State = pb.State
	st.Tags = pb.Tags
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

type endpointTagPairPb struct {
	Key   string `json:"key,omitempty"`
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

func endpointTagsToPb(st *EndpointTags) (*endpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagsPb{}
	pb.CustomTags = st.CustomTags

	return pb, nil
}

type endpointTagsPb struct {
	CustomTags []EndpointTagPair `json:"custom_tags,omitempty"`
}

func endpointTagsFromPb(pb *endpointTagsPb) (*EndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTags{}
	st.CustomTags = pb.CustomTags

	return st, nil
}

func enumValueToPb(st *EnumValue) (*enumValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enumValuePb{}
	pb.EnumOptions = st.EnumOptions
	pb.MultiValuesOptions = st.MultiValuesOptions
	pb.Values = st.Values

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enumValuePb struct {
	EnumOptions        string              `json:"enum_options,omitempty"`
	MultiValuesOptions *MultiValuesOptions `json:"multi_values_options,omitempty"`
	Values             []string            `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enumValueFromPb(pb *enumValuePb) (*EnumValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnumValue{}
	st.EnumOptions = pb.EnumOptions
	st.MultiValuesOptions = pb.MultiValuesOptions
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
	pb.Parameters = st.Parameters
	pb.RowLimit = st.RowLimit
	pb.Schema = st.Schema
	pb.Statement = st.Statement
	pb.WaitTimeout = st.WaitTimeout
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type executeStatementRequestPb struct {
	ByteLimit     int64                                `json:"byte_limit,omitempty"`
	Catalog       string                               `json:"catalog,omitempty"`
	Disposition   Disposition                          `json:"disposition,omitempty"`
	Format        Format                               `json:"format,omitempty"`
	OnWaitTimeout ExecuteStatementRequestOnWaitTimeout `json:"on_wait_timeout,omitempty"`
	Parameters    []StatementParameterListItem         `json:"parameters,omitempty"`
	RowLimit      int64                                `json:"row_limit,omitempty"`
	Schema        string                               `json:"schema,omitempty"`
	Statement     string                               `json:"statement"`
	WaitTimeout   string                               `json:"wait_timeout,omitempty"`
	WarehouseId   string                               `json:"warehouse_id"`

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
	st.Parameters = pb.Parameters
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

type externalLinkPb struct {
	ByteCount             int64             `json:"byte_count,omitempty"`
	ChunkIndex            int               `json:"chunk_index,omitempty"`
	Expiration            string            `json:"expiration,omitempty"`
	ExternalLink          string            `json:"external_link,omitempty"`
	HttpHeaders           map[string]string `json:"http_headers,omitempty"`
	NextChunkIndex        int               `json:"next_chunk_index,omitempty"`
	NextChunkInternalLink string            `json:"next_chunk_internal_link,omitempty"`
	RowCount              int64             `json:"row_count,omitempty"`
	RowOffset             int64             `json:"row_offset,omitempty"`

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

func externalQuerySourceToPb(st *ExternalQuerySource) (*externalQuerySourcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalQuerySourcePb{}
	pb.AlertId = st.AlertId
	pb.DashboardId = st.DashboardId
	pb.GenieSpaceId = st.GenieSpaceId
	pb.JobInfo = st.JobInfo
	pb.LegacyDashboardId = st.LegacyDashboardId
	pb.NotebookId = st.NotebookId
	pb.SqlQueryId = st.SqlQueryId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type externalQuerySourcePb struct {
	AlertId           string                      `json:"alert_id,omitempty"`
	DashboardId       string                      `json:"dashboard_id,omitempty"`
	GenieSpaceId      string                      `json:"genie_space_id,omitempty"`
	JobInfo           *ExternalQuerySourceJobInfo `json:"job_info,omitempty"`
	LegacyDashboardId string                      `json:"legacy_dashboard_id,omitempty"`
	NotebookId        string                      `json:"notebook_id,omitempty"`
	SqlQueryId        string                      `json:"sql_query_id,omitempty"`

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
	st.JobInfo = pb.JobInfo
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

type externalQuerySourceJobInfoPb struct {
	JobId        string `json:"job_id,omitempty"`
	JobRunId     string `json:"job_run_id,omitempty"`
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

func getAlertRequestToPb(st *GetAlertRequest) (*getAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

func getAlertV2RequestToPb(st *GetAlertV2Request) (*getAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertV2RequestPb{}
	pb.Id = st.Id

	return pb, nil
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

func getAlertsLegacyRequestToPb(st *GetAlertsLegacyRequest) (*getAlertsLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAlertsLegacyRequestPb{}
	pb.AlertId = st.AlertId

	return pb, nil
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

func getDashboardRequestToPb(st *GetDashboardRequest) (*getDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
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

func getDbsqlPermissionRequestToPb(st *GetDbsqlPermissionRequest) (*getDbsqlPermissionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDbsqlPermissionRequestPb{}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

type getDbsqlPermissionRequestPb struct {
	ObjectId   string           `json:"-" url:"-"`
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

func getQueriesLegacyRequestToPb(st *GetQueriesLegacyRequest) (*getQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
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

func getQueryRequestToPb(st *GetQueryRequest) (*getQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQueryRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

func getResponseToPb(st *GetResponse) (*getResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getResponsePb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getResponsePb struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	ObjectId          string          `json:"object_id,omitempty"`
	ObjectType        ObjectType      `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getResponseFromPb(pb *getResponsePb) (*GetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetResponse{}
	st.AccessControlList = pb.AccessControlList
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

func getStatementRequestToPb(st *GetStatementRequest) (*getStatementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatementRequestPb{}
	pb.StatementId = st.StatementId

	return pb, nil
}

type getStatementRequestPb struct {
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

func getStatementResultChunkNRequestToPb(st *GetStatementResultChunkNRequest) (*getStatementResultChunkNRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatementResultChunkNRequestPb{}
	pb.ChunkIndex = st.ChunkIndex
	pb.StatementId = st.StatementId

	return pb, nil
}

type getStatementResultChunkNRequestPb struct {
	ChunkIndex  int    `json:"-" url:"-"`
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

func getWarehousePermissionLevelsRequestToPb(st *GetWarehousePermissionLevelsRequest) (*getWarehousePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionLevelsRequestPb{}
	pb.WarehouseId = st.WarehouseId

	return pb, nil
}

type getWarehousePermissionLevelsRequestPb struct {
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

func getWarehousePermissionLevelsResponseToPb(st *GetWarehousePermissionLevelsResponse) (*getWarehousePermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getWarehousePermissionLevelsResponsePb struct {
	PermissionLevels []WarehousePermissionsDescription `json:"permission_levels,omitempty"`
}

func getWarehousePermissionLevelsResponseFromPb(pb *getWarehousePermissionLevelsResponsePb) (*GetWarehousePermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehousePermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getWarehousePermissionsRequestToPb(st *GetWarehousePermissionsRequest) (*getWarehousePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehousePermissionsRequestPb{}
	pb.WarehouseId = st.WarehouseId

	return pb, nil
}

type getWarehousePermissionsRequestPb struct {
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

func getWarehouseRequestToPb(st *GetWarehouseRequest) (*getWarehouseRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehouseRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getWarehouseRequestPb struct {
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

func getWarehouseResponseToPb(st *GetWarehouseResponse) (*getWarehouseResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWarehouseResponsePb{}
	pb.AutoStopMins = st.AutoStopMins
	pb.Channel = st.Channel
	pb.ClusterSize = st.ClusterSize
	pb.CreatorName = st.CreatorName
	pb.EnablePhoton = st.EnablePhoton
	pb.EnableServerlessCompute = st.EnableServerlessCompute
	pb.Health = st.Health
	pb.Id = st.Id
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.JdbcUrl = st.JdbcUrl
	pb.MaxNumClusters = st.MaxNumClusters
	pb.MinNumClusters = st.MinNumClusters
	pb.Name = st.Name
	pb.NumActiveSessions = st.NumActiveSessions
	pb.NumClusters = st.NumClusters
	pb.OdbcParams = st.OdbcParams
	pb.SpotInstancePolicy = st.SpotInstancePolicy
	pb.State = st.State
	pb.Tags = st.Tags
	pb.WarehouseType = st.WarehouseType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getWarehouseResponsePb struct {
	AutoStopMins            int                               `json:"auto_stop_mins,omitempty"`
	Channel                 *Channel                          `json:"channel,omitempty"`
	ClusterSize             string                            `json:"cluster_size,omitempty"`
	CreatorName             string                            `json:"creator_name,omitempty"`
	EnablePhoton            bool                              `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                              `json:"enable_serverless_compute,omitempty"`
	Health                  *EndpointHealth                   `json:"health,omitempty"`
	Id                      string                            `json:"id,omitempty"`
	InstanceProfileArn      string                            `json:"instance_profile_arn,omitempty"`
	JdbcUrl                 string                            `json:"jdbc_url,omitempty"`
	MaxNumClusters          int                               `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                               `json:"min_num_clusters,omitempty"`
	Name                    string                            `json:"name,omitempty"`
	NumActiveSessions       int64                             `json:"num_active_sessions,omitempty"`
	NumClusters             int                               `json:"num_clusters,omitempty"`
	OdbcParams              *OdbcParams                       `json:"odbc_params,omitempty"`
	SpotInstancePolicy      SpotInstancePolicy                `json:"spot_instance_policy,omitempty"`
	State                   State                             `json:"state,omitempty"`
	Tags                    *EndpointTags                     `json:"tags,omitempty"`
	WarehouseType           GetWarehouseResponseWarehouseType `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getWarehouseResponseFromPb(pb *getWarehouseResponsePb) (*GetWarehouseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWarehouseResponse{}
	st.AutoStopMins = pb.AutoStopMins
	st.Channel = pb.Channel
	st.ClusterSize = pb.ClusterSize
	st.CreatorName = pb.CreatorName
	st.EnablePhoton = pb.EnablePhoton
	st.EnableServerlessCompute = pb.EnableServerlessCompute
	st.Health = pb.Health
	st.Id = pb.Id
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.JdbcUrl = pb.JdbcUrl
	st.MaxNumClusters = pb.MaxNumClusters
	st.MinNumClusters = pb.MinNumClusters
	st.Name = pb.Name
	st.NumActiveSessions = pb.NumActiveSessions
	st.NumClusters = pb.NumClusters
	st.OdbcParams = pb.OdbcParams
	st.SpotInstancePolicy = pb.SpotInstancePolicy
	st.State = pb.State
	st.Tags = pb.Tags
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

func getWorkspaceWarehouseConfigResponseToPb(st *GetWorkspaceWarehouseConfigResponse) (*getWorkspaceWarehouseConfigResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceWarehouseConfigResponsePb{}
	pb.Channel = st.Channel
	pb.ConfigParam = st.ConfigParam
	pb.DataAccessConfig = st.DataAccessConfig
	pb.EnabledWarehouseTypes = st.EnabledWarehouseTypes
	pb.GlobalParam = st.GlobalParam
	pb.GoogleServiceAccount = st.GoogleServiceAccount
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.SecurityPolicy = st.SecurityPolicy
	pb.SqlConfigurationParameters = st.SqlConfigurationParameters

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getWorkspaceWarehouseConfigResponsePb struct {
	Channel                    *Channel                                          `json:"channel,omitempty"`
	ConfigParam                *RepeatedEndpointConfPairs                        `json:"config_param,omitempty"`
	DataAccessConfig           []EndpointConfPair                                `json:"data_access_config,omitempty"`
	EnabledWarehouseTypes      []WarehouseTypePair                               `json:"enabled_warehouse_types,omitempty"`
	GlobalParam                *RepeatedEndpointConfPairs                        `json:"global_param,omitempty"`
	GoogleServiceAccount       string                                            `json:"google_service_account,omitempty"`
	InstanceProfileArn         string                                            `json:"instance_profile_arn,omitempty"`
	SecurityPolicy             GetWorkspaceWarehouseConfigResponseSecurityPolicy `json:"security_policy,omitempty"`
	SqlConfigurationParameters *RepeatedEndpointConfPairs                        `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getWorkspaceWarehouseConfigResponseFromPb(pb *getWorkspaceWarehouseConfigResponsePb) (*GetWorkspaceWarehouseConfigResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceWarehouseConfigResponse{}
	st.Channel = pb.Channel
	st.ConfigParam = pb.ConfigParam
	st.DataAccessConfig = pb.DataAccessConfig
	st.EnabledWarehouseTypes = pb.EnabledWarehouseTypes
	st.GlobalParam = pb.GlobalParam
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.SecurityPolicy = pb.SecurityPolicy
	st.SqlConfigurationParameters = pb.SqlConfigurationParameters

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getWorkspaceWarehouseConfigResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getWorkspaceWarehouseConfigResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	pb.Options = st.Options
	pb.Parent = st.Parent
	pb.Query = st.Query
	pb.Rearm = st.Rearm
	pb.State = st.State
	pb.UpdatedAt = st.UpdatedAt
	pb.User = st.User

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type legacyAlertPb struct {
	CreatedAt       string           `json:"created_at,omitempty"`
	Id              string           `json:"id,omitempty"`
	LastTriggeredAt string           `json:"last_triggered_at,omitempty"`
	Name            string           `json:"name,omitempty"`
	Options         *AlertOptions    `json:"options,omitempty"`
	Parent          string           `json:"parent,omitempty"`
	Query           *AlertQuery      `json:"query,omitempty"`
	Rearm           int              `json:"rearm,omitempty"`
	State           LegacyAlertState `json:"state,omitempty"`
	UpdatedAt       string           `json:"updated_at,omitempty"`
	User            *User            `json:"user,omitempty"`

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
	st.Options = pb.Options
	st.Parent = pb.Parent
	st.Query = pb.Query
	st.Rearm = pb.Rearm
	st.State = pb.State
	st.UpdatedAt = pb.UpdatedAt
	st.User = pb.User

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *legacyAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st legacyAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	pb.LastModifiedBy = st.LastModifiedBy
	pb.LastModifiedById = st.LastModifiedById
	pb.LatestQueryDataId = st.LatestQueryDataId
	pb.Name = st.Name
	pb.Options = st.Options
	pb.Parent = st.Parent
	pb.PermissionTier = st.PermissionTier
	pb.Query = st.Query
	pb.QueryHash = st.QueryHash
	pb.RunAsRole = st.RunAsRole
	pb.Tags = st.Tags
	pb.UpdatedAt = st.UpdatedAt
	pb.User = st.User
	pb.UserId = st.UserId
	pb.Visualizations = st.Visualizations

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type legacyQueryPb struct {
	CanEdit           bool                  `json:"can_edit,omitempty"`
	CreatedAt         string                `json:"created_at,omitempty"`
	DataSourceId      string                `json:"data_source_id,omitempty"`
	Description       string                `json:"description,omitempty"`
	Id                string                `json:"id,omitempty"`
	IsArchived        bool                  `json:"is_archived,omitempty"`
	IsDraft           bool                  `json:"is_draft,omitempty"`
	IsFavorite        bool                  `json:"is_favorite,omitempty"`
	IsSafe            bool                  `json:"is_safe,omitempty"`
	LastModifiedBy    *User                 `json:"last_modified_by,omitempty"`
	LastModifiedById  int                   `json:"last_modified_by_id,omitempty"`
	LatestQueryDataId string                `json:"latest_query_data_id,omitempty"`
	Name              string                `json:"name,omitempty"`
	Options           *QueryOptions         `json:"options,omitempty"`
	Parent            string                `json:"parent,omitempty"`
	PermissionTier    PermissionLevel       `json:"permission_tier,omitempty"`
	Query             string                `json:"query,omitempty"`
	QueryHash         string                `json:"query_hash,omitempty"`
	RunAsRole         RunAsRole             `json:"run_as_role,omitempty"`
	Tags              []string              `json:"tags,omitempty"`
	UpdatedAt         string                `json:"updated_at,omitempty"`
	User              *User                 `json:"user,omitempty"`
	UserId            int                   `json:"user_id,omitempty"`
	Visualizations    []LegacyVisualization `json:"visualizations,omitempty"`

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
	st.LastModifiedBy = pb.LastModifiedBy
	st.LastModifiedById = pb.LastModifiedById
	st.LatestQueryDataId = pb.LatestQueryDataId
	st.Name = pb.Name
	st.Options = pb.Options
	st.Parent = pb.Parent
	st.PermissionTier = pb.PermissionTier
	st.Query = pb.Query
	st.QueryHash = pb.QueryHash
	st.RunAsRole = pb.RunAsRole
	st.Tags = pb.Tags
	st.UpdatedAt = pb.UpdatedAt
	st.User = pb.User
	st.UserId = pb.UserId
	st.Visualizations = pb.Visualizations

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *legacyQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st legacyQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	pb.Query = st.Query
	pb.Type = st.Type
	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type legacyVisualizationPb struct {
	CreatedAt   string       `json:"created_at,omitempty"`
	Description string       `json:"description,omitempty"`
	Id          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Options     any          `json:"options,omitempty"`
	Query       *LegacyQuery `json:"query,omitempty"`
	Type        string       `json:"type,omitempty"`
	UpdatedAt   string       `json:"updated_at,omitempty"`

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
	st.Query = pb.Query
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

type listAlertsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
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

func listAlertsResponseToPb(st *ListAlertsResponse) (*listAlertsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAlertsResponsePb struct {
	NextPageToken string                    `json:"next_page_token,omitempty"`
	Results       []ListAlertsResponseAlert `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsResponseFromPb(pb *listAlertsResponsePb) (*ListAlertsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listAlertsResponseAlertToPb(st *ListAlertsResponseAlert) (*listAlertsResponseAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsResponseAlertPb{}
	pb.Condition = st.Condition
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

type listAlertsResponseAlertPb struct {
	Condition          *AlertCondition `json:"condition,omitempty"`
	CreateTime         string          `json:"create_time,omitempty"`
	CustomBody         string          `json:"custom_body,omitempty"`
	CustomSubject      string          `json:"custom_subject,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	Id                 string          `json:"id,omitempty"`
	LifecycleState     LifecycleState  `json:"lifecycle_state,omitempty"`
	NotifyOnOk         bool            `json:"notify_on_ok,omitempty"`
	OwnerUserName      string          `json:"owner_user_name,omitempty"`
	QueryId            string          `json:"query_id,omitempty"`
	SecondsToRetrigger int             `json:"seconds_to_retrigger,omitempty"`
	State              AlertState      `json:"state,omitempty"`
	TriggerTime        string          `json:"trigger_time,omitempty"`
	UpdateTime         string          `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsResponseAlertFromPb(pb *listAlertsResponseAlertPb) (*ListAlertsResponseAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsResponseAlert{}
	st.Condition = pb.Condition
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

type listAlertsV2RequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
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

func listAlertsV2ResponseToPb(st *ListAlertsV2Response) (*listAlertsV2ResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAlertsV2ResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAlertsV2ResponsePb struct {
	NextPageToken string    `json:"next_page_token,omitempty"`
	Results       []AlertV2 `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAlertsV2ResponseFromPb(pb *listAlertsV2ResponsePb) (*ListAlertsV2Response, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAlertsV2Response{}
	st.NextPageToken = pb.NextPageToken
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAlertsV2ResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAlertsV2ResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type listDashboardsRequestPb struct {
	Order    ListOrder `json:"-" url:"order,omitempty"`
	Page     int       `json:"-" url:"page,omitempty"`
	PageSize int       `json:"-" url:"page_size,omitempty"`
	Q        string    `json:"-" url:"q,omitempty"`

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

type listQueriesLegacyRequestPb struct {
	Order    string `json:"-" url:"order,omitempty"`
	Page     int    `json:"-" url:"page,omitempty"`
	PageSize int    `json:"-" url:"page_size,omitempty"`
	Q        string `json:"-" url:"q,omitempty"`

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

type listQueriesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
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

func listQueriesResponseToPb(st *ListQueriesResponse) (*listQueriesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueriesResponsePb{}
	pb.HasNextPage = st.HasNextPage
	pb.NextPageToken = st.NextPageToken
	pb.Res = st.Res

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQueriesResponsePb struct {
	HasNextPage   bool        `json:"has_next_page,omitempty"`
	NextPageToken string      `json:"next_page_token,omitempty"`
	Res           []QueryInfo `json:"res,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueriesResponseFromPb(pb *listQueriesResponsePb) (*ListQueriesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueriesResponse{}
	st.HasNextPage = pb.HasNextPage
	st.NextPageToken = pb.NextPageToken
	st.Res = pb.Res

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueriesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueriesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listQueryHistoryRequestToPb(st *ListQueryHistoryRequest) (*listQueryHistoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueryHistoryRequestPb{}
	pb.FilterBy = st.FilterBy
	pb.IncludeMetrics = st.IncludeMetrics
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQueryHistoryRequestPb struct {
	FilterBy       *QueryFilter `json:"-" url:"filter_by,omitempty"`
	IncludeMetrics bool         `json:"-" url:"include_metrics,omitempty"`
	MaxResults     int          `json:"-" url:"max_results,omitempty"`
	PageToken      string       `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueryHistoryRequestFromPb(pb *listQueryHistoryRequestPb) (*ListQueryHistoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryHistoryRequest{}
	st.FilterBy = pb.FilterBy
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

func listQueryObjectsResponseToPb(st *ListQueryObjectsResponse) (*listQueryObjectsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQueryObjectsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQueryObjectsResponsePb struct {
	NextPageToken string                          `json:"next_page_token,omitempty"`
	Results       []ListQueryObjectsResponseQuery `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQueryObjectsResponseFromPb(pb *listQueryObjectsResponsePb) (*ListQueryObjectsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQueryObjectsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQueryObjectsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQueryObjectsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	pb.Parameters = st.Parameters
	pb.QueryText = st.QueryText
	pb.RunAsMode = st.RunAsMode
	pb.Schema = st.Schema
	pb.Tags = st.Tags
	pb.UpdateTime = st.UpdateTime
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQueryObjectsResponseQueryPb struct {
	ApplyAutoLimit       bool             `json:"apply_auto_limit,omitempty"`
	Catalog              string           `json:"catalog,omitempty"`
	CreateTime           string           `json:"create_time,omitempty"`
	Description          string           `json:"description,omitempty"`
	DisplayName          string           `json:"display_name,omitempty"`
	Id                   string           `json:"id,omitempty"`
	LastModifierUserName string           `json:"last_modifier_user_name,omitempty"`
	LifecycleState       LifecycleState   `json:"lifecycle_state,omitempty"`
	OwnerUserName        string           `json:"owner_user_name,omitempty"`
	Parameters           []QueryParameter `json:"parameters,omitempty"`
	QueryText            string           `json:"query_text,omitempty"`
	RunAsMode            RunAsMode        `json:"run_as_mode,omitempty"`
	Schema               string           `json:"schema,omitempty"`
	Tags                 []string         `json:"tags,omitempty"`
	UpdateTime           string           `json:"update_time,omitempty"`
	WarehouseId          string           `json:"warehouse_id,omitempty"`

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
	st.Parameters = pb.Parameters
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

func listResponseToPb(st *ListResponse) (*listResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listResponsePb{}
	pb.Count = st.Count
	pb.Page = st.Page
	pb.PageSize = st.PageSize
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listResponsePb struct {
	Count    int         `json:"count,omitempty"`
	Page     int         `json:"page,omitempty"`
	PageSize int         `json:"page_size,omitempty"`
	Results  []Dashboard `json:"results,omitempty"`

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
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type listVisualizationsForQueryRequestPb struct {
	Id        string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
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

func listVisualizationsForQueryResponseToPb(st *ListVisualizationsForQueryResponse) (*listVisualizationsForQueryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listVisualizationsForQueryResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listVisualizationsForQueryResponsePb struct {
	NextPageToken string          `json:"next_page_token,omitempty"`
	Results       []Visualization `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listVisualizationsForQueryResponseFromPb(pb *listVisualizationsForQueryResponsePb) (*ListVisualizationsForQueryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListVisualizationsForQueryResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listVisualizationsForQueryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listVisualizationsForQueryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type listWarehousesRequestPb struct {
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

func listWarehousesResponseToPb(st *ListWarehousesResponse) (*listWarehousesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWarehousesResponsePb{}
	pb.Warehouses = st.Warehouses

	return pb, nil
}

type listWarehousesResponsePb struct {
	Warehouses []EndpointInfo `json:"warehouses,omitempty"`
}

func listWarehousesResponseFromPb(pb *listWarehousesResponsePb) (*ListWarehousesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWarehousesResponse{}
	st.Warehouses = pb.Warehouses

	return st, nil
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

type multiValuesOptionsPb struct {
	Prefix    string `json:"prefix,omitempty"`
	Separator string `json:"separator,omitempty"`
	Suffix    string `json:"suffix,omitempty"`

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

func numericValueToPb(st *NumericValue) (*numericValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &numericValuePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type odbcParamsPb struct {
	Hostname string `json:"hostname,omitempty"`
	Path     string `json:"path,omitempty"`
	Port     int    `json:"port,omitempty"`
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

func parameterToPb(st *Parameter) (*parameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &parameterPb{}
	pb.EnumOptions = st.EnumOptions
	pb.MultiValuesOptions = st.MultiValuesOptions
	pb.Name = st.Name
	pb.QueryId = st.QueryId
	pb.Title = st.Title
	pb.Type = st.Type
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type parameterPb struct {
	EnumOptions        string              `json:"enumOptions,omitempty"`
	MultiValuesOptions *MultiValuesOptions `json:"multiValuesOptions,omitempty"`
	Name               string              `json:"name,omitempty"`
	QueryId            string              `json:"queryId,omitempty"`
	Title              string              `json:"title,omitempty"`
	Type               ParameterType       `json:"type,omitempty"`
	Value              any                 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func parameterFromPb(pb *parameterPb) (*Parameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Parameter{}
	st.EnumOptions = pb.EnumOptions
	st.MultiValuesOptions = pb.MultiValuesOptions
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
	pb.Parameters = st.Parameters
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

type queryPb struct {
	ApplyAutoLimit       bool             `json:"apply_auto_limit,omitempty"`
	Catalog              string           `json:"catalog,omitempty"`
	CreateTime           string           `json:"create_time,omitempty"`
	Description          string           `json:"description,omitempty"`
	DisplayName          string           `json:"display_name,omitempty"`
	Id                   string           `json:"id,omitempty"`
	LastModifierUserName string           `json:"last_modifier_user_name,omitempty"`
	LifecycleState       LifecycleState   `json:"lifecycle_state,omitempty"`
	OwnerUserName        string           `json:"owner_user_name,omitempty"`
	Parameters           []QueryParameter `json:"parameters,omitempty"`
	ParentPath           string           `json:"parent_path,omitempty"`
	QueryText            string           `json:"query_text,omitempty"`
	RunAsMode            RunAsMode        `json:"run_as_mode,omitempty"`
	Schema               string           `json:"schema,omitempty"`
	Tags                 []string         `json:"tags,omitempty"`
	UpdateTime           string           `json:"update_time,omitempty"`
	WarehouseId          string           `json:"warehouse_id,omitempty"`

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
	st.Parameters = pb.Parameters
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

func queryBackedValueToPb(st *QueryBackedValue) (*queryBackedValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryBackedValuePb{}
	pb.MultiValuesOptions = st.MultiValuesOptions
	pb.QueryId = st.QueryId
	pb.Values = st.Values

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryBackedValuePb struct {
	MultiValuesOptions *MultiValuesOptions `json:"multi_values_options,omitempty"`
	QueryId            string              `json:"query_id,omitempty"`
	Values             []string            `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryBackedValueFromPb(pb *queryBackedValuePb) (*QueryBackedValue, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryBackedValue{}
	st.MultiValuesOptions = pb.MultiValuesOptions
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

type queryEditContentPb struct {
	DataSourceId string    `json:"data_source_id,omitempty"`
	Description  string    `json:"description,omitempty"`
	Name         string    `json:"name,omitempty"`
	Options      any       `json:"options,omitempty"`
	Query        string    `json:"query,omitempty"`
	QueryId      string    `json:"-" url:"-"`
	RunAsRole    RunAsRole `json:"run_as_role,omitempty"`
	Tags         []string  `json:"tags,omitempty"`

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

func queryFilterToPb(st *QueryFilter) (*queryFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryFilterPb{}
	pb.QueryStartTimeRange = st.QueryStartTimeRange
	pb.StatementIds = st.StatementIds
	pb.Statuses = st.Statuses
	pb.UserIds = st.UserIds
	pb.WarehouseIds = st.WarehouseIds

	return pb, nil
}

type queryFilterPb struct {
	QueryStartTimeRange *TimeRange    `json:"query_start_time_range,omitempty" url:"query_start_time_range,omitempty"`
	StatementIds        []string      `json:"statement_ids,omitempty" url:"statement_ids,omitempty"`
	Statuses            []QueryStatus `json:"statuses,omitempty" url:"statuses,omitempty"`
	UserIds             []int64       `json:"user_ids,omitempty" url:"user_ids,omitempty"`
	WarehouseIds        []string      `json:"warehouse_ids,omitempty" url:"warehouse_ids,omitempty"`
}

func queryFilterFromPb(pb *queryFilterPb) (*QueryFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryFilter{}
	st.QueryStartTimeRange = pb.QueryStartTimeRange
	st.StatementIds = pb.StatementIds
	st.Statuses = pb.Statuses
	st.UserIds = pb.UserIds
	st.WarehouseIds = pb.WarehouseIds

	return st, nil
}

func queryInfoToPb(st *QueryInfo) (*queryInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryInfoPb{}
	pb.ChannelUsed = st.ChannelUsed
	pb.ClientApplication = st.ClientApplication
	pb.Duration = st.Duration
	pb.EndpointId = st.EndpointId
	pb.ErrorMessage = st.ErrorMessage
	pb.ExecutedAsUserId = st.ExecutedAsUserId
	pb.ExecutedAsUserName = st.ExecutedAsUserName
	pb.ExecutionEndTimeMs = st.ExecutionEndTimeMs
	pb.IsFinal = st.IsFinal
	pb.LookupKey = st.LookupKey
	pb.Metrics = st.Metrics
	pb.PlansState = st.PlansState
	pb.QueryEndTimeMs = st.QueryEndTimeMs
	pb.QueryId = st.QueryId
	pb.QuerySource = st.QuerySource
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

type queryInfoPb struct {
	ChannelUsed        *ChannelInfo         `json:"channel_used,omitempty"`
	ClientApplication  string               `json:"client_application,omitempty"`
	Duration           int64                `json:"duration,omitempty"`
	EndpointId         string               `json:"endpoint_id,omitempty"`
	ErrorMessage       string               `json:"error_message,omitempty"`
	ExecutedAsUserId   int64                `json:"executed_as_user_id,omitempty"`
	ExecutedAsUserName string               `json:"executed_as_user_name,omitempty"`
	ExecutionEndTimeMs int64                `json:"execution_end_time_ms,omitempty"`
	IsFinal            bool                 `json:"is_final,omitempty"`
	LookupKey          string               `json:"lookup_key,omitempty"`
	Metrics            *QueryMetrics        `json:"metrics,omitempty"`
	PlansState         PlansState           `json:"plans_state,omitempty"`
	QueryEndTimeMs     int64                `json:"query_end_time_ms,omitempty"`
	QueryId            string               `json:"query_id,omitempty"`
	QuerySource        *ExternalQuerySource `json:"query_source,omitempty"`
	QueryStartTimeMs   int64                `json:"query_start_time_ms,omitempty"`
	QueryText          string               `json:"query_text,omitempty"`
	RowsProduced       int64                `json:"rows_produced,omitempty"`
	SparkUiUrl         string               `json:"spark_ui_url,omitempty"`
	StatementType      QueryStatementType   `json:"statement_type,omitempty"`
	Status             QueryStatus          `json:"status,omitempty"`
	UserId             int64                `json:"user_id,omitempty"`
	UserName           string               `json:"user_name,omitempty"`
	WarehouseId        string               `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryInfoFromPb(pb *queryInfoPb) (*QueryInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryInfo{}
	st.ChannelUsed = pb.ChannelUsed
	st.ClientApplication = pb.ClientApplication
	st.Duration = pb.Duration
	st.EndpointId = pb.EndpointId
	st.ErrorMessage = pb.ErrorMessage
	st.ExecutedAsUserId = pb.ExecutedAsUserId
	st.ExecutedAsUserName = pb.ExecutedAsUserName
	st.ExecutionEndTimeMs = pb.ExecutionEndTimeMs
	st.IsFinal = pb.IsFinal
	st.LookupKey = pb.LookupKey
	st.Metrics = pb.Metrics
	st.PlansState = pb.PlansState
	st.QueryEndTimeMs = pb.QueryEndTimeMs
	st.QueryId = pb.QueryId
	st.QuerySource = pb.QuerySource
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

func queryListToPb(st *QueryList) (*queryListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryListPb{}
	pb.Count = st.Count
	pb.Page = st.Page
	pb.PageSize = st.PageSize
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryListPb struct {
	Count    int           `json:"count,omitempty"`
	Page     int           `json:"page,omitempty"`
	PageSize int           `json:"page_size,omitempty"`
	Results  []LegacyQuery `json:"results,omitempty"`

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
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	pb.TaskTimeOverTimeRange = st.TaskTimeOverTimeRange
	pb.TaskTotalTimeMs = st.TaskTotalTimeMs
	pb.TotalTimeMs = st.TotalTimeMs
	pb.WriteRemoteBytes = st.WriteRemoteBytes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryMetricsPb struct {
	CompilationTimeMs               int64              `json:"compilation_time_ms,omitempty"`
	ExecutionTimeMs                 int64              `json:"execution_time_ms,omitempty"`
	NetworkSentBytes                int64              `json:"network_sent_bytes,omitempty"`
	OverloadingQueueStartTimestamp  int64              `json:"overloading_queue_start_timestamp,omitempty"`
	PhotonTotalTimeMs               int64              `json:"photon_total_time_ms,omitempty"`
	ProvisioningQueueStartTimestamp int64              `json:"provisioning_queue_start_timestamp,omitempty"`
	PrunedBytes                     int64              `json:"pruned_bytes,omitempty"`
	PrunedFilesCount                int64              `json:"pruned_files_count,omitempty"`
	QueryCompilationStartTimestamp  int64              `json:"query_compilation_start_timestamp,omitempty"`
	ReadBytes                       int64              `json:"read_bytes,omitempty"`
	ReadCacheBytes                  int64              `json:"read_cache_bytes,omitempty"`
	ReadFilesCount                  int64              `json:"read_files_count,omitempty"`
	ReadPartitionsCount             int64              `json:"read_partitions_count,omitempty"`
	ReadRemoteBytes                 int64              `json:"read_remote_bytes,omitempty"`
	ResultFetchTimeMs               int64              `json:"result_fetch_time_ms,omitempty"`
	ResultFromCache                 bool               `json:"result_from_cache,omitempty"`
	RowsProducedCount               int64              `json:"rows_produced_count,omitempty"`
	RowsReadCount                   int64              `json:"rows_read_count,omitempty"`
	SpillToDiskBytes                int64              `json:"spill_to_disk_bytes,omitempty"`
	TaskTimeOverTimeRange           *TaskTimeOverRange `json:"task_time_over_time_range,omitempty"`
	TaskTotalTimeMs                 int64              `json:"task_total_time_ms,omitempty"`
	TotalTimeMs                     int64              `json:"total_time_ms,omitempty"`
	WriteRemoteBytes                int64              `json:"write_remote_bytes,omitempty"`

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
	st.TaskTimeOverTimeRange = pb.TaskTimeOverTimeRange
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

func queryOptionsToPb(st *QueryOptions) (*queryOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryOptionsPb{}
	pb.Catalog = st.Catalog
	pb.MovedToTrashAt = st.MovedToTrashAt
	pb.Parameters = st.Parameters
	pb.Schema = st.Schema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryOptionsPb struct {
	Catalog        string      `json:"catalog,omitempty"`
	MovedToTrashAt string      `json:"moved_to_trash_at,omitempty"`
	Parameters     []Parameter `json:"parameters,omitempty"`
	Schema         string      `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryOptionsFromPb(pb *queryOptionsPb) (*QueryOptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryOptions{}
	st.Catalog = pb.Catalog
	st.MovedToTrashAt = pb.MovedToTrashAt
	st.Parameters = pb.Parameters
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

func queryParameterToPb(st *QueryParameter) (*queryParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryParameterPb{}
	pb.DateRangeValue = st.DateRangeValue
	pb.DateValue = st.DateValue
	pb.EnumValue = st.EnumValue
	pb.Name = st.Name
	pb.NumericValue = st.NumericValue
	pb.QueryBackedValue = st.QueryBackedValue
	pb.TextValue = st.TextValue
	pb.Title = st.Title

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryParameterPb struct {
	DateRangeValue   *DateRangeValue   `json:"date_range_value,omitempty"`
	DateValue        *DateValue        `json:"date_value,omitempty"`
	EnumValue        *EnumValue        `json:"enum_value,omitempty"`
	Name             string            `json:"name,omitempty"`
	NumericValue     *NumericValue     `json:"numeric_value,omitempty"`
	QueryBackedValue *QueryBackedValue `json:"query_backed_value,omitempty"`
	TextValue        *TextValue        `json:"text_value,omitempty"`
	Title            string            `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryParameterFromPb(pb *queryParameterPb) (*QueryParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryParameter{}
	st.DateRangeValue = pb.DateRangeValue
	st.DateValue = pb.DateValue
	st.EnumValue = pb.EnumValue
	st.Name = pb.Name
	st.NumericValue = pb.NumericValue
	st.QueryBackedValue = pb.QueryBackedValue
	st.TextValue = pb.TextValue
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

type queryPostContentPb struct {
	DataSourceId string    `json:"data_source_id,omitempty"`
	Description  string    `json:"description,omitempty"`
	Name         string    `json:"name,omitempty"`
	Options      any       `json:"options,omitempty"`
	Parent       string    `json:"parent,omitempty"`
	Query        string    `json:"query,omitempty"`
	RunAsRole    RunAsRole `json:"run_as_role,omitempty"`
	Tags         []string  `json:"tags,omitempty"`

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

func repeatedEndpointConfPairsToPb(st *RepeatedEndpointConfPairs) (*repeatedEndpointConfPairsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &repeatedEndpointConfPairsPb{}
	pb.ConfigPair = st.ConfigPair
	pb.ConfigurationPairs = st.ConfigurationPairs

	return pb, nil
}

type repeatedEndpointConfPairsPb struct {
	ConfigPair         []EndpointConfPair `json:"config_pair,omitempty"`
	ConfigurationPairs []EndpointConfPair `json:"configuration_pairs,omitempty"`
}

func repeatedEndpointConfPairsFromPb(pb *repeatedEndpointConfPairsPb) (*RepeatedEndpointConfPairs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RepeatedEndpointConfPairs{}
	st.ConfigPair = pb.ConfigPair
	st.ConfigurationPairs = pb.ConfigurationPairs

	return st, nil
}

func restoreDashboardRequestToPb(st *RestoreDashboardRequest) (*restoreDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreDashboardRequestPb{}
	pb.DashboardId = st.DashboardId

	return pb, nil
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

func restoreQueriesLegacyRequestToPb(st *RestoreQueriesLegacyRequest) (*restoreQueriesLegacyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreQueriesLegacyRequestPb{}
	pb.QueryId = st.QueryId

	return pb, nil
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

func resultDataToPb(st *ResultData) (*resultDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultDataPb{}
	pb.ByteCount = st.ByteCount
	pb.ChunkIndex = st.ChunkIndex
	pb.DataArray = st.DataArray
	pb.ExternalLinks = st.ExternalLinks
	pb.NextChunkIndex = st.NextChunkIndex
	pb.NextChunkInternalLink = st.NextChunkInternalLink
	pb.RowCount = st.RowCount
	pb.RowOffset = st.RowOffset

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resultDataPb struct {
	ByteCount             int64          `json:"byte_count,omitempty"`
	ChunkIndex            int            `json:"chunk_index,omitempty"`
	DataArray             [][]string     `json:"data_array,omitempty"`
	ExternalLinks         []ExternalLink `json:"external_links,omitempty"`
	NextChunkIndex        int            `json:"next_chunk_index,omitempty"`
	NextChunkInternalLink string         `json:"next_chunk_internal_link,omitempty"`
	RowCount              int64          `json:"row_count,omitempty"`
	RowOffset             int64          `json:"row_offset,omitempty"`

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
	st.ExternalLinks = pb.ExternalLinks
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

func resultManifestToPb(st *ResultManifest) (*resultManifestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultManifestPb{}
	pb.Chunks = st.Chunks
	pb.Format = st.Format
	pb.Schema = st.Schema
	pb.TotalByteCount = st.TotalByteCount
	pb.TotalChunkCount = st.TotalChunkCount
	pb.TotalRowCount = st.TotalRowCount
	pb.Truncated = st.Truncated

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resultManifestPb struct {
	Chunks          []BaseChunkInfo `json:"chunks,omitempty"`
	Format          Format          `json:"format,omitempty"`
	Schema          *ResultSchema   `json:"schema,omitempty"`
	TotalByteCount  int64           `json:"total_byte_count,omitempty"`
	TotalChunkCount int             `json:"total_chunk_count,omitempty"`
	TotalRowCount   int64           `json:"total_row_count,omitempty"`
	Truncated       bool            `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultManifestFromPb(pb *resultManifestPb) (*ResultManifest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultManifest{}
	st.Chunks = pb.Chunks
	st.Format = pb.Format
	st.Schema = pb.Schema
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

func resultSchemaToPb(st *ResultSchema) (*resultSchemaPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &resultSchemaPb{}
	pb.ColumnCount = st.ColumnCount
	pb.Columns = st.Columns

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type resultSchemaPb struct {
	ColumnCount int          `json:"column_count,omitempty"`
	Columns     []ColumnInfo `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func resultSchemaFromPb(pb *resultSchemaPb) (*ResultSchema, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ResultSchema{}
	st.ColumnCount = pb.ColumnCount
	st.Columns = pb.Columns

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *resultSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st resultSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type serviceErrorPb struct {
	ErrorCode ServiceErrorCode `json:"error_code,omitempty"`
	Message   string           `json:"message,omitempty"`

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

func setRequestToPb(st *SetRequest) (*setRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

type setRequestPb struct {
	AccessControlList []AccessControl  `json:"access_control_list,omitempty"`
	ObjectId          string           `json:"-" url:"-"`
	ObjectType        ObjectTypePlural `json:"-" url:"-"`
}

func setRequestFromPb(pb *setRequestPb) (*SetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetRequest{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

func setResponseToPb(st *SetResponse) (*setResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setResponsePb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type setResponsePb struct {
	AccessControlList []AccessControl `json:"access_control_list,omitempty"`
	ObjectId          string          `json:"object_id,omitempty"`
	ObjectType        ObjectType      `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setResponseFromPb(pb *setResponsePb) (*SetResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetResponse{}
	st.AccessControlList = pb.AccessControlList
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

func setWorkspaceWarehouseConfigRequestToPb(st *SetWorkspaceWarehouseConfigRequest) (*setWorkspaceWarehouseConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setWorkspaceWarehouseConfigRequestPb{}
	pb.Channel = st.Channel
	pb.ConfigParam = st.ConfigParam
	pb.DataAccessConfig = st.DataAccessConfig
	pb.EnabledWarehouseTypes = st.EnabledWarehouseTypes
	pb.GlobalParam = st.GlobalParam
	pb.GoogleServiceAccount = st.GoogleServiceAccount
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.SecurityPolicy = st.SecurityPolicy
	pb.SqlConfigurationParameters = st.SqlConfigurationParameters

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type setWorkspaceWarehouseConfigRequestPb struct {
	Channel                    *Channel                                         `json:"channel,omitempty"`
	ConfigParam                *RepeatedEndpointConfPairs                       `json:"config_param,omitempty"`
	DataAccessConfig           []EndpointConfPair                               `json:"data_access_config,omitempty"`
	EnabledWarehouseTypes      []WarehouseTypePair                              `json:"enabled_warehouse_types,omitempty"`
	GlobalParam                *RepeatedEndpointConfPairs                       `json:"global_param,omitempty"`
	GoogleServiceAccount       string                                           `json:"google_service_account,omitempty"`
	InstanceProfileArn         string                                           `json:"instance_profile_arn,omitempty"`
	SecurityPolicy             SetWorkspaceWarehouseConfigRequestSecurityPolicy `json:"security_policy,omitempty"`
	SqlConfigurationParameters *RepeatedEndpointConfPairs                       `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setWorkspaceWarehouseConfigRequestFromPb(pb *setWorkspaceWarehouseConfigRequestPb) (*SetWorkspaceWarehouseConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetWorkspaceWarehouseConfigRequest{}
	st.Channel = pb.Channel
	st.ConfigParam = pb.ConfigParam
	st.DataAccessConfig = pb.DataAccessConfig
	st.EnabledWarehouseTypes = pb.EnabledWarehouseTypes
	st.GlobalParam = pb.GlobalParam
	st.GoogleServiceAccount = pb.GoogleServiceAccount
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.SecurityPolicy = pb.SecurityPolicy
	st.SqlConfigurationParameters = pb.SqlConfigurationParameters

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *setWorkspaceWarehouseConfigRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setWorkspaceWarehouseConfigRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func startRequestToPb(st *StartRequest) (*startRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type startRequestPb struct {
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

type statementParameterListItemPb struct {
	Name  string `json:"name"`
	Type  string `json:"type,omitempty"`
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

func statementResponseToPb(st *StatementResponse) (*statementResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &statementResponsePb{}
	pb.Manifest = st.Manifest
	pb.Result = st.Result
	pb.StatementId = st.StatementId
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type statementResponsePb struct {
	Manifest    *ResultManifest  `json:"manifest,omitempty"`
	Result      *ResultData      `json:"result,omitempty"`
	StatementId string           `json:"statement_id,omitempty"`
	Status      *StatementStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func statementResponseFromPb(pb *statementResponsePb) (*StatementResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StatementResponse{}
	st.Manifest = pb.Manifest
	st.Result = pb.Result
	st.StatementId = pb.StatementId
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *statementResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st statementResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func statementStatusToPb(st *StatementStatus) (*statementStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &statementStatusPb{}
	pb.Error = st.Error
	pb.State = st.State

	return pb, nil
}

type statementStatusPb struct {
	Error *ServiceError  `json:"error,omitempty"`
	State StatementState `json:"state,omitempty"`
}

func statementStatusFromPb(pb *statementStatusPb) (*StatementStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StatementStatus{}
	st.Error = pb.Error
	st.State = pb.State

	return st, nil
}

func stopRequestToPb(st *StopRequest) (*stopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type stopRequestPb struct {
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

func successToPb(st *Success) (*successPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &successPb{}
	pb.Message = st.Message

	return pb, nil
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

func taskTimeOverRangeToPb(st *TaskTimeOverRange) (*taskTimeOverRangePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskTimeOverRangePb{}
	pb.Entries = st.Entries
	pb.Interval = st.Interval

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type taskTimeOverRangePb struct {
	Entries  []TaskTimeOverRangeEntry `json:"entries,omitempty"`
	Interval int64                    `json:"interval,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskTimeOverRangeFromPb(pb *taskTimeOverRangePb) (*TaskTimeOverRange, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskTimeOverRange{}
	st.Entries = pb.Entries
	st.Interval = pb.Interval

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskTimeOverRangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskTimeOverRangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func taskTimeOverRangeEntryToPb(st *TaskTimeOverRangeEntry) (*taskTimeOverRangeEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &taskTimeOverRangeEntryPb{}
	pb.TaskCompletedTimeMs = st.TaskCompletedTimeMs

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type taskTimeOverRangeEntryPb struct {
	TaskCompletedTimeMs int64 `json:"task_completed_time_ms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func taskTimeOverRangeEntryFromPb(pb *taskTimeOverRangeEntryPb) (*TaskTimeOverRangeEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TaskTimeOverRangeEntry{}
	st.TaskCompletedTimeMs = pb.TaskCompletedTimeMs

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *taskTimeOverRangeEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st taskTimeOverRangeEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

func textValueToPb(st *TextValue) (*textValuePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &textValuePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type timeRangePb struct {
	EndTimeMs   int64 `json:"end_time_ms,omitempty" url:"end_time_ms,omitempty"`
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

func transferOwnershipObjectIdToPb(st *TransferOwnershipObjectId) (*transferOwnershipObjectIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transferOwnershipObjectIdPb{}
	pb.NewOwner = st.NewOwner

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type transferOwnershipObjectIdPb struct {
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

func transferOwnershipRequestToPb(st *TransferOwnershipRequest) (*transferOwnershipRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transferOwnershipRequestPb{}
	pb.NewOwner = st.NewOwner
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type transferOwnershipRequestPb struct {
	NewOwner   string                    `json:"new_owner,omitempty"`
	ObjectId   TransferOwnershipObjectId `json:"-" url:"-"`
	ObjectType OwnableObjectType         `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func transferOwnershipRequestFromPb(pb *transferOwnershipRequestPb) (*TransferOwnershipRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransferOwnershipRequest{}
	st.NewOwner = pb.NewOwner
	st.ObjectId = pb.ObjectId
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

func trashAlertRequestToPb(st *TrashAlertRequest) (*trashAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashAlertRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

func trashAlertV2RequestToPb(st *TrashAlertV2Request) (*trashAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashAlertV2RequestPb{}
	pb.Id = st.Id

	return pb, nil
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

func trashQueryRequestToPb(st *TrashQueryRequest) (*trashQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trashQueryRequestPb{}
	pb.Id = st.Id

	return pb, nil
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

func updateAlertRequestToPb(st *UpdateAlertRequest) (*updateAlertRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAlertRequestPb{}
	pb.Alert = st.Alert
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName
	pb.Id = st.Id
	pb.UpdateMask = st.UpdateMask

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateAlertRequestPb struct {
	Alert                  *UpdateAlertRequestAlert `json:"alert,omitempty"`
	AutoResolveDisplayName bool                     `json:"auto_resolve_display_name,omitempty"`
	Id                     string                   `json:"-" url:"-"`
	UpdateMask             string                   `json:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateAlertRequestFromPb(pb *updateAlertRequestPb) (*UpdateAlertRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertRequest{}
	st.Alert = pb.Alert
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName
	st.Id = pb.Id
	st.UpdateMask = pb.UpdateMask

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateAlertRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateAlertRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateAlertRequestAlertToPb(st *UpdateAlertRequestAlert) (*updateAlertRequestAlertPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAlertRequestAlertPb{}
	pb.Condition = st.Condition
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

type updateAlertRequestAlertPb struct {
	Condition          *AlertCondition `json:"condition,omitempty"`
	CustomBody         string          `json:"custom_body,omitempty"`
	CustomSubject      string          `json:"custom_subject,omitempty"`
	DisplayName        string          `json:"display_name,omitempty"`
	NotifyOnOk         bool            `json:"notify_on_ok,omitempty"`
	OwnerUserName      string          `json:"owner_user_name,omitempty"`
	QueryId            string          `json:"query_id,omitempty"`
	SecondsToRetrigger int             `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateAlertRequestAlertFromPb(pb *updateAlertRequestAlertPb) (*UpdateAlertRequestAlert, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertRequestAlert{}
	st.Condition = pb.Condition
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

func updateAlertV2RequestToPb(st *UpdateAlertV2Request) (*updateAlertV2RequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAlertV2RequestPb{}
	pb.Alert = st.Alert
	pb.Id = st.Id
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

type updateAlertV2RequestPb struct {
	Alert      AlertV2 `json:"alert"`
	Id         string  `json:"-" url:"-"`
	UpdateMask string  `json:"-" url:"update_mask"`
}

func updateAlertV2RequestFromPb(pb *updateAlertV2RequestPb) (*UpdateAlertV2Request, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAlertV2Request{}
	st.Alert = pb.Alert
	st.Id = pb.Id
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

func updateQueryRequestToPb(st *UpdateQueryRequest) (*updateQueryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateQueryRequestPb{}
	pb.AutoResolveDisplayName = st.AutoResolveDisplayName
	pb.Id = st.Id
	pb.Query = st.Query
	pb.UpdateMask = st.UpdateMask

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateQueryRequestPb struct {
	AutoResolveDisplayName bool                     `json:"auto_resolve_display_name,omitempty"`
	Id                     string                   `json:"-" url:"-"`
	Query                  *UpdateQueryRequestQuery `json:"query,omitempty"`
	UpdateMask             string                   `json:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateQueryRequestFromPb(pb *updateQueryRequestPb) (*UpdateQueryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateQueryRequest{}
	st.AutoResolveDisplayName = pb.AutoResolveDisplayName
	st.Id = pb.Id
	st.Query = pb.Query
	st.UpdateMask = pb.UpdateMask

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	pb.Parameters = st.Parameters
	pb.QueryText = st.QueryText
	pb.RunAsMode = st.RunAsMode
	pb.Schema = st.Schema
	pb.Tags = st.Tags
	pb.WarehouseId = st.WarehouseId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateQueryRequestQueryPb struct {
	ApplyAutoLimit bool             `json:"apply_auto_limit,omitempty"`
	Catalog        string           `json:"catalog,omitempty"`
	Description    string           `json:"description,omitempty"`
	DisplayName    string           `json:"display_name,omitempty"`
	OwnerUserName  string           `json:"owner_user_name,omitempty"`
	Parameters     []QueryParameter `json:"parameters,omitempty"`
	QueryText      string           `json:"query_text,omitempty"`
	RunAsMode      RunAsMode        `json:"run_as_mode,omitempty"`
	Schema         string           `json:"schema,omitempty"`
	Tags           []string         `json:"tags,omitempty"`
	WarehouseId    string           `json:"warehouse_id,omitempty"`

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
	st.Parameters = pb.Parameters
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

func updateVisualizationRequestToPb(st *UpdateVisualizationRequest) (*updateVisualizationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateVisualizationRequestPb{}
	pb.Id = st.Id
	pb.UpdateMask = st.UpdateMask
	pb.Visualization = st.Visualization

	return pb, nil
}

type updateVisualizationRequestPb struct {
	Id            string                                   `json:"-" url:"-"`
	UpdateMask    string                                   `json:"update_mask"`
	Visualization *UpdateVisualizationRequestVisualization `json:"visualization,omitempty"`
}

func updateVisualizationRequestFromPb(pb *updateVisualizationRequestPb) (*UpdateVisualizationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateVisualizationRequest{}
	st.Id = pb.Id
	st.UpdateMask = pb.UpdateMask
	st.Visualization = pb.Visualization

	return st, nil
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

type updateVisualizationRequestVisualizationPb struct {
	DisplayName         string `json:"display_name,omitempty"`
	SerializedOptions   string `json:"serialized_options,omitempty"`
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	Type                string `json:"type,omitempty"`

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

func updateWidgetRequestToPb(st *UpdateWidgetRequest) (*updateWidgetRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWidgetRequestPb{}
	pb.DashboardId = st.DashboardId
	pb.Id = st.Id
	pb.Options = st.Options
	pb.Text = st.Text
	pb.VisualizationId = st.VisualizationId
	pb.Width = st.Width

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateWidgetRequestPb struct {
	DashboardId     string        `json:"dashboard_id"`
	Id              string        `json:"-" url:"-"`
	Options         WidgetOptions `json:"options"`
	Text            string        `json:"text,omitempty"`
	VisualizationId string        `json:"visualization_id,omitempty"`
	Width           int           `json:"width"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateWidgetRequestFromPb(pb *updateWidgetRequestPb) (*UpdateWidgetRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWidgetRequest{}
	st.DashboardId = pb.DashboardId
	st.Id = pb.Id
	st.Options = pb.Options
	st.Text = pb.Text
	st.VisualizationId = pb.VisualizationId
	st.Width = pb.Width

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateWidgetRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateWidgetRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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

type userPb struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`

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

type visualizationPb struct {
	CreateTime          string `json:"create_time,omitempty"`
	DisplayName         string `json:"display_name,omitempty"`
	Id                  string `json:"id,omitempty"`
	QueryId             string `json:"query_id,omitempty"`
	SerializedOptions   string `json:"serialized_options,omitempty"`
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	Type                string `json:"type,omitempty"`
	UpdateTime          string `json:"update_time,omitempty"`

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

type warehouseAccessControlRequestPb struct {
	GroupName            string                   `json:"group_name,omitempty"`
	PermissionLevel      WarehousePermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                   `json:"service_principal_name,omitempty"`
	UserName             string                   `json:"user_name,omitempty"`

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

func warehouseAccessControlResponseToPb(st *WarehouseAccessControlResponse) (*warehouseAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehouseAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type warehouseAccessControlResponsePb struct {
	AllPermissions       []WarehousePermission `json:"all_permissions,omitempty"`
	DisplayName          string                `json:"display_name,omitempty"`
	GroupName            string                `json:"group_name,omitempty"`
	ServicePrincipalName string                `json:"service_principal_name,omitempty"`
	UserName             string                `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func warehouseAccessControlResponseFromPb(pb *warehouseAccessControlResponsePb) (*WarehouseAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehouseAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
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

type warehousePermissionPb struct {
	Inherited           bool                     `json:"inherited,omitempty"`
	InheritedFromObject []string                 `json:"inherited_from_object,omitempty"`
	PermissionLevel     WarehousePermissionLevel `json:"permission_level,omitempty"`

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

func warehousePermissionsToPb(st *WarehousePermissions) (*warehousePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type warehousePermissionsPb struct {
	AccessControlList []WarehouseAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                           `json:"object_id,omitempty"`
	ObjectType        string                           `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func warehousePermissionsFromPb(pb *warehousePermissionsPb) (*WarehousePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissions{}
	st.AccessControlList = pb.AccessControlList
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

type warehousePermissionsDescriptionPb struct {
	Description     string                   `json:"description,omitempty"`
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

func warehousePermissionsRequestToPb(st *WarehousePermissionsRequest) (*warehousePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &warehousePermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.WarehouseId = st.WarehouseId

	return pb, nil
}

type warehousePermissionsRequestPb struct {
	AccessControlList []WarehouseAccessControlRequest `json:"access_control_list,omitempty"`
	WarehouseId       string                          `json:"-" url:"-"`
}

func warehousePermissionsRequestFromPb(pb *warehousePermissionsRequestPb) (*WarehousePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WarehousePermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.WarehouseId = pb.WarehouseId

	return st, nil
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

type warehouseTypePairPb struct {
	Enabled       bool                           `json:"enabled,omitempty"`
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

func widgetToPb(st *Widget) (*widgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetPb{}
	pb.Id = st.Id
	pb.Options = st.Options
	pb.Visualization = st.Visualization
	pb.Width = st.Width

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type widgetPb struct {
	Id            string               `json:"id,omitempty"`
	Options       *WidgetOptions       `json:"options,omitempty"`
	Visualization *LegacyVisualization `json:"visualization,omitempty"`
	Width         int                  `json:"width,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func widgetFromPb(pb *widgetPb) (*Widget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Widget{}
	st.Id = pb.Id
	st.Options = pb.Options
	st.Visualization = pb.Visualization
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

func widgetOptionsToPb(st *WidgetOptions) (*widgetOptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &widgetOptionsPb{}
	pb.CreatedAt = st.CreatedAt
	pb.Description = st.Description
	pb.IsHidden = st.IsHidden
	pb.ParameterMappings = st.ParameterMappings
	pb.Position = st.Position
	pb.Title = st.Title
	pb.UpdatedAt = st.UpdatedAt

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type widgetOptionsPb struct {
	CreatedAt         string          `json:"created_at,omitempty"`
	Description       string          `json:"description,omitempty"`
	IsHidden          bool            `json:"isHidden,omitempty"`
	ParameterMappings any             `json:"parameterMappings,omitempty"`
	Position          *WidgetPosition `json:"position,omitempty"`
	Title             string          `json:"title,omitempty"`
	UpdatedAt         string          `json:"updated_at,omitempty"`

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
	st.Position = pb.Position
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

type widgetPositionPb struct {
	AutoHeight bool `json:"autoHeight,omitempty"`
	Col        int  `json:"col,omitempty"`
	Row        int  `json:"row,omitempty"`
	SizeX      int  `json:"sizeX,omitempty"`
	SizeY      int  `json:"sizeY,omitempty"`

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
	s := t.Format(time.RFC3339Nano)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339Nano, *s)
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
	if *s == "" {
		return &[]string{}, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
