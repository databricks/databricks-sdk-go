// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package sqlpb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccessControlPb struct {
	GroupName       string            `json:"group_name,omitempty"`
	PermissionLevel PermissionLevelPb `json:"permission_level,omitempty"`
	UserName        string            `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AccessControlPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AccessControlPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AggregationPb string

const AggregationAvg AggregationPb = `AVG`

const AggregationCount AggregationPb = `COUNT`

const AggregationCountDistinct AggregationPb = `COUNT_DISTINCT`

const AggregationMax AggregationPb = `MAX`

const AggregationMedian AggregationPb = `MEDIAN`

const AggregationMin AggregationPb = `MIN`

const AggregationStddev AggregationPb = `STDDEV`

const AggregationSum AggregationPb = `SUM`

type AlertPb struct {
	Condition          *AlertConditionPb `json:"condition,omitempty"`
	CreateTime         string            `json:"create_time,omitempty"`
	CustomBody         string            `json:"custom_body,omitempty"`
	CustomSubject      string            `json:"custom_subject,omitempty"`
	DisplayName        string            `json:"display_name,omitempty"`
	Id                 string            `json:"id,omitempty"`
	LifecycleState     LifecycleStatePb  `json:"lifecycle_state,omitempty"`
	NotifyOnOk         bool              `json:"notify_on_ok,omitempty"`
	OwnerUserName      string            `json:"owner_user_name,omitempty"`
	ParentPath         string            `json:"parent_path,omitempty"`
	QueryId            string            `json:"query_id,omitempty"`
	SecondsToRetrigger int               `json:"seconds_to_retrigger,omitempty"`
	State              AlertStatePb      `json:"state,omitempty"`
	TriggerTime        string            `json:"trigger_time,omitempty"`
	UpdateTime         string            `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertConditionPb struct {
	EmptyResultState AlertStatePb               `json:"empty_result_state,omitempty"`
	Op               AlertOperatorPb            `json:"op,omitempty"`
	Operand          *AlertConditionOperandPb   `json:"operand,omitempty"`
	Threshold        *AlertConditionThresholdPb `json:"threshold,omitempty"`
}

type AlertConditionOperandPb struct {
	Column *AlertOperandColumnPb `json:"column,omitempty"`
}

type AlertConditionThresholdPb struct {
	Value *AlertOperandValuePb `json:"value,omitempty"`
}

type AlertEvaluationStatePb string

const AlertEvaluationStateError AlertEvaluationStatePb = `ERROR`

const AlertEvaluationStateOk AlertEvaluationStatePb = `OK`

const AlertEvaluationStateTriggered AlertEvaluationStatePb = `TRIGGERED`

const AlertEvaluationStateUnknown AlertEvaluationStatePb = `UNKNOWN`

type AlertOperandColumnPb struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertOperandColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertOperandColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertOperandValuePb struct {
	BoolValue   bool    `json:"bool_value,omitempty"`
	DoubleValue float64 `json:"double_value,omitempty"`
	StringValue string  `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertOperandValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertOperandValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertOperatorPb string

const AlertOperatorEqual AlertOperatorPb = `EQUAL`

const AlertOperatorGreaterThan AlertOperatorPb = `GREATER_THAN`

const AlertOperatorGreaterThanOrEqual AlertOperatorPb = `GREATER_THAN_OR_EQUAL`

const AlertOperatorIsNull AlertOperatorPb = `IS_NULL`

const AlertOperatorLessThan AlertOperatorPb = `LESS_THAN`

const AlertOperatorLessThanOrEqual AlertOperatorPb = `LESS_THAN_OR_EQUAL`

const AlertOperatorNotEqual AlertOperatorPb = `NOT_EQUAL`

type AlertOptionsPb struct {
	Column           string                         `json:"column"`
	CustomBody       string                         `json:"custom_body,omitempty"`
	CustomSubject    string                         `json:"custom_subject,omitempty"`
	EmptyResultState AlertOptionsEmptyResultStatePb `json:"empty_result_state,omitempty"`
	Muted            bool                           `json:"muted,omitempty"`
	Op               string                         `json:"op"`
	Value            any                            `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertOptionsEmptyResultStatePb string

const AlertOptionsEmptyResultStateOk AlertOptionsEmptyResultStatePb = `ok`

const AlertOptionsEmptyResultStateTriggered AlertOptionsEmptyResultStatePb = `triggered`

const AlertOptionsEmptyResultStateUnknown AlertOptionsEmptyResultStatePb = `unknown`

type AlertQueryPb struct {
	CreatedAt    string          `json:"created_at,omitempty"`
	DataSourceId string          `json:"data_source_id,omitempty"`
	Description  string          `json:"description,omitempty"`
	Id           string          `json:"id,omitempty"`
	IsArchived   bool            `json:"is_archived,omitempty"`
	IsDraft      bool            `json:"is_draft,omitempty"`
	IsSafe       bool            `json:"is_safe,omitempty"`
	Name         string          `json:"name,omitempty"`
	Options      *QueryOptionsPb `json:"options,omitempty"`
	Query        string          `json:"query,omitempty"`
	Tags         []string        `json:"tags,omitempty"`
	UpdatedAt    string          `json:"updated_at,omitempty"`
	UserId       int             `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertStatePb string

const AlertStateOk AlertStatePb = `OK`

const AlertStateTriggered AlertStatePb = `TRIGGERED`

const AlertStateUnknown AlertStatePb = `UNKNOWN`

type AlertV2Pb struct {
	CreateTime        string               `json:"create_time,omitempty"`
	CustomDescription string               `json:"custom_description,omitempty"`
	CustomSummary     string               `json:"custom_summary,omitempty"`
	DisplayName       string               `json:"display_name,omitempty"`
	Evaluation        *AlertV2EvaluationPb `json:"evaluation,omitempty"`
	Id                string               `json:"id,omitempty"`
	LifecycleState    LifecycleStatePb     `json:"lifecycle_state,omitempty"`
	OwnerUserName     string               `json:"owner_user_name,omitempty"`
	ParentPath        string               `json:"parent_path,omitempty"`
	QueryText         string               `json:"query_text,omitempty"`
	RunAsUserName     string               `json:"run_as_user_name,omitempty"`
	Schedule          *CronSchedulePb      `json:"schedule,omitempty"`
	UpdateTime        string               `json:"update_time,omitempty"`
	WarehouseId       string               `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertV2Pb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertV2Pb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2EvaluationPb struct {
	ComparisonOperator ComparisonOperatorPb    `json:"comparison_operator,omitempty"`
	EmptyResultState   AlertEvaluationStatePb  `json:"empty_result_state,omitempty"`
	LastEvaluatedAt    string                  `json:"last_evaluated_at,omitempty"`
	Notification       *AlertV2NotificationPb  `json:"notification,omitempty"`
	Source             *AlertV2OperandColumnPb `json:"source,omitempty"`
	State              AlertEvaluationStatePb  `json:"state,omitempty"`
	Threshold          *AlertV2OperandPb       `json:"threshold,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertV2EvaluationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertV2EvaluationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2NotificationPb struct {
	NotifyOnOk       bool                    `json:"notify_on_ok,omitempty"`
	RetriggerSeconds int                     `json:"retrigger_seconds,omitempty"`
	Subscriptions    []AlertV2SubscriptionPb `json:"subscriptions,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertV2NotificationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertV2NotificationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2OperandPb struct {
	Column *AlertV2OperandColumnPb `json:"column,omitempty"`
	Value  *AlertV2OperandValuePb  `json:"value,omitempty"`
}

type AlertV2OperandColumnPb struct {
	Aggregation AggregationPb `json:"aggregation,omitempty"`
	Display     string        `json:"display,omitempty"`
	Name        string        `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertV2OperandColumnPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertV2OperandColumnPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2OperandValuePb struct {
	BoolValue   bool    `json:"bool_value,omitempty"`
	DoubleValue float64 `json:"double_value,omitempty"`
	StringValue string  `json:"string_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertV2OperandValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertV2OperandValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertV2SubscriptionPb struct {
	DestinationId string `json:"destination_id,omitempty"`
	UserEmail     string `json:"user_email,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertV2SubscriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertV2SubscriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BaseChunkInfoPb struct {
	ByteCount  int64 `json:"byte_count,omitempty"`
	ChunkIndex int   `json:"chunk_index,omitempty"`
	RowCount   int64 `json:"row_count,omitempty"`
	RowOffset  int64 `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BaseChunkInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BaseChunkInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CancelExecutionRequestPb struct {
	StatementId string `json:"-" url:"-"`
}

type ChannelPb struct {
	DbsqlVersion string        `json:"dbsql_version,omitempty"`
	Name         ChannelNamePb `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ChannelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ChannelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ChannelInfoPb struct {
	DbsqlVersion string        `json:"dbsql_version,omitempty"`
	Name         ChannelNamePb `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ChannelInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ChannelInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ChannelNamePb string

const ChannelNameChannelNameCurrent ChannelNamePb = `CHANNEL_NAME_CURRENT`

const ChannelNameChannelNameCustom ChannelNamePb = `CHANNEL_NAME_CUSTOM`

const ChannelNameChannelNamePreview ChannelNamePb = `CHANNEL_NAME_PREVIEW`

const ChannelNameChannelNamePrevious ChannelNamePb = `CHANNEL_NAME_PREVIOUS`

type ClientConfigPb struct {
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

func (st *ClientConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClientConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnInfoPb struct {
	Name             string               `json:"name,omitempty"`
	Position         int                  `json:"position,omitempty"`
	TypeIntervalType string               `json:"type_interval_type,omitempty"`
	TypeName         ColumnInfoTypeNamePb `json:"type_name,omitempty"`
	TypePrecision    int                  `json:"type_precision,omitempty"`
	TypeScale        int                  `json:"type_scale,omitempty"`
	TypeText         string               `json:"type_text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ColumnInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ColumnInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ColumnInfoTypeNamePb string

const ColumnInfoTypeNameArray ColumnInfoTypeNamePb = `ARRAY`

const ColumnInfoTypeNameBinary ColumnInfoTypeNamePb = `BINARY`

const ColumnInfoTypeNameBoolean ColumnInfoTypeNamePb = `BOOLEAN`

const ColumnInfoTypeNameByte ColumnInfoTypeNamePb = `BYTE`

const ColumnInfoTypeNameChar ColumnInfoTypeNamePb = `CHAR`

const ColumnInfoTypeNameDate ColumnInfoTypeNamePb = `DATE`

const ColumnInfoTypeNameDecimal ColumnInfoTypeNamePb = `DECIMAL`

const ColumnInfoTypeNameDouble ColumnInfoTypeNamePb = `DOUBLE`

const ColumnInfoTypeNameFloat ColumnInfoTypeNamePb = `FLOAT`

const ColumnInfoTypeNameInt ColumnInfoTypeNamePb = `INT`

const ColumnInfoTypeNameInterval ColumnInfoTypeNamePb = `INTERVAL`

const ColumnInfoTypeNameLong ColumnInfoTypeNamePb = `LONG`

const ColumnInfoTypeNameMap ColumnInfoTypeNamePb = `MAP`

const ColumnInfoTypeNameNull ColumnInfoTypeNamePb = `NULL`

const ColumnInfoTypeNameShort ColumnInfoTypeNamePb = `SHORT`

const ColumnInfoTypeNameString ColumnInfoTypeNamePb = `STRING`

const ColumnInfoTypeNameStruct ColumnInfoTypeNamePb = `STRUCT`

const ColumnInfoTypeNameTimestamp ColumnInfoTypeNamePb = `TIMESTAMP`

const ColumnInfoTypeNameUserDefinedType ColumnInfoTypeNamePb = `USER_DEFINED_TYPE`

type ComparisonOperatorPb string

const ComparisonOperatorEqual ComparisonOperatorPb = `EQUAL`

const ComparisonOperatorGreaterThan ComparisonOperatorPb = `GREATER_THAN`

const ComparisonOperatorGreaterThanOrEqual ComparisonOperatorPb = `GREATER_THAN_OR_EQUAL`

const ComparisonOperatorIsNotNull ComparisonOperatorPb = `IS_NOT_NULL`

const ComparisonOperatorIsNull ComparisonOperatorPb = `IS_NULL`

const ComparisonOperatorLessThan ComparisonOperatorPb = `LESS_THAN`

const ComparisonOperatorLessThanOrEqual ComparisonOperatorPb = `LESS_THAN_OR_EQUAL`

const ComparisonOperatorNotEqual ComparisonOperatorPb = `NOT_EQUAL`

type CreateAlertPb struct {
	Name    string         `json:"name"`
	Options AlertOptionsPb `json:"options"`
	Parent  string         `json:"parent,omitempty"`
	QueryId string         `json:"query_id"`
	Rearm   int            `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateAlertRequestPb struct {
	Alert                  *CreateAlertRequestAlertPb `json:"alert,omitempty"`
	AutoResolveDisplayName bool                       `json:"auto_resolve_display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateAlertRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateAlertRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateAlertRequestAlertPb struct {
	Condition          *AlertConditionPb `json:"condition,omitempty"`
	CustomBody         string            `json:"custom_body,omitempty"`
	CustomSubject      string            `json:"custom_subject,omitempty"`
	DisplayName        string            `json:"display_name,omitempty"`
	NotifyOnOk         bool              `json:"notify_on_ok,omitempty"`
	ParentPath         string            `json:"parent_path,omitempty"`
	QueryId            string            `json:"query_id,omitempty"`
	SecondsToRetrigger int               `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateAlertRequestAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateAlertRequestAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateAlertV2RequestPb struct {
	Alert AlertV2Pb `json:"alert"`
}

type CreateQueryRequestPb struct {
	AutoResolveDisplayName bool                       `json:"auto_resolve_display_name,omitempty"`
	Query                  *CreateQueryRequestQueryPb `json:"query,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateQueryRequestQueryPb struct {
	ApplyAutoLimit bool               `json:"apply_auto_limit,omitempty"`
	Catalog        string             `json:"catalog,omitempty"`
	Description    string             `json:"description,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	Parameters     []QueryParameterPb `json:"parameters,omitempty"`
	ParentPath     string             `json:"parent_path,omitempty"`
	QueryText      string             `json:"query_text,omitempty"`
	RunAsMode      RunAsModePb        `json:"run_as_mode,omitempty"`
	Schema         string             `json:"schema,omitempty"`
	Tags           []string           `json:"tags,omitempty"`
	WarehouseId    string             `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateQueryRequestQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateQueryRequestQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateQueryVisualizationsLegacyRequestPb struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Options     any    `json:"options"`
	QueryId     string `json:"query_id"`
	Type        string `json:"type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateQueryVisualizationsLegacyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateQueryVisualizationsLegacyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateVisualizationRequestPb struct {
	Visualization *CreateVisualizationRequestVisualizationPb `json:"visualization,omitempty"`
}

type CreateVisualizationRequestVisualizationPb struct {
	DisplayName         string `json:"display_name,omitempty"`
	QueryId             string `json:"query_id,omitempty"`
	SerializedOptions   string `json:"serialized_options,omitempty"`
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	Type                string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateVisualizationRequestVisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateVisualizationRequestVisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateWarehouseRequestPb struct {
	AutoStopMins            int                                   `json:"auto_stop_mins,omitempty"`
	Channel                 *ChannelPb                            `json:"channel,omitempty"`
	ClusterSize             string                                `json:"cluster_size,omitempty"`
	CreatorName             string                                `json:"creator_name,omitempty"`
	EnablePhoton            bool                                  `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                                  `json:"enable_serverless_compute,omitempty"`
	InstanceProfileArn      string                                `json:"instance_profile_arn,omitempty"`
	MaxNumClusters          int                                   `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                                   `json:"min_num_clusters,omitempty"`
	Name                    string                                `json:"name,omitempty"`
	SpotInstancePolicy      SpotInstancePolicyPb                  `json:"spot_instance_policy,omitempty"`
	Tags                    *EndpointTagsPb                       `json:"tags,omitempty"`
	WarehouseType           CreateWarehouseRequestWarehouseTypePb `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateWarehouseRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateWarehouseRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateWarehouseRequestWarehouseTypePb string

const CreateWarehouseRequestWarehouseTypeClassic CreateWarehouseRequestWarehouseTypePb = `CLASSIC`

const CreateWarehouseRequestWarehouseTypePro CreateWarehouseRequestWarehouseTypePb = `PRO`

const CreateWarehouseRequestWarehouseTypeTypeUnspecified CreateWarehouseRequestWarehouseTypePb = `TYPE_UNSPECIFIED`

type CreateWarehouseResponsePb struct {
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateWarehouseResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateWarehouseResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateWidgetPb struct {
	DashboardId     string          `json:"dashboard_id"`
	Options         WidgetOptionsPb `json:"options"`
	Text            string          `json:"text,omitempty"`
	VisualizationId string          `json:"visualization_id,omitempty"`
	Width           int             `json:"width"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateWidgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateWidgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CronSchedulePb struct {
	PauseStatus        SchedulePauseStatusPb `json:"pause_status,omitempty"`
	QuartzCronSchedule string                `json:"quartz_cron_schedule,omitempty"`
	TimezoneId         string                `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CronSchedulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CronSchedulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardPb struct {
	CanEdit                 bool                `json:"can_edit,omitempty"`
	CreatedAt               string              `json:"created_at,omitempty"`
	DashboardFiltersEnabled bool                `json:"dashboard_filters_enabled,omitempty"`
	Id                      string              `json:"id,omitempty"`
	IsArchived              bool                `json:"is_archived,omitempty"`
	IsDraft                 bool                `json:"is_draft,omitempty"`
	IsFavorite              bool                `json:"is_favorite,omitempty"`
	Name                    string              `json:"name,omitempty"`
	Options                 *DashboardOptionsPb `json:"options,omitempty"`
	Parent                  string              `json:"parent,omitempty"`
	PermissionTier          PermissionLevelPb   `json:"permission_tier,omitempty"`
	Slug                    string              `json:"slug,omitempty"`
	Tags                    []string            `json:"tags,omitempty"`
	UpdatedAt               string              `json:"updated_at,omitempty"`
	User                    *UserPb             `json:"user,omitempty"`
	UserId                  int                 `json:"user_id,omitempty"`
	Widgets                 []WidgetPb          `json:"widgets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DashboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DashboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardEditContentPb struct {
	DashboardId string      `json:"-" url:"-"`
	Name        string      `json:"name,omitempty"`
	RunAsRole   RunAsRolePb `json:"run_as_role,omitempty"`
	Tags        []string    `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DashboardEditContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DashboardEditContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DashboardOptionsPb struct {
	MovedToTrashAt string `json:"moved_to_trash_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DashboardOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DashboardOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataSourcePb struct {
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

func (st *DataSourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DataSourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatePrecisionPb string

const DatePrecisionDayPrecision DatePrecisionPb = `DAY_PRECISION`

const DatePrecisionMinutePrecision DatePrecisionPb = `MINUTE_PRECISION`

const DatePrecisionSecondPrecision DatePrecisionPb = `SECOND_PRECISION`

type DateRangePb struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

type DateRangeValuePb struct {
	DateRangeValue        *DateRangePb                     `json:"date_range_value,omitempty"`
	DynamicDateRangeValue DateRangeValueDynamicDateRangePb `json:"dynamic_date_range_value,omitempty"`
	Precision             DatePrecisionPb                  `json:"precision,omitempty"`
	StartDayOfWeek        int                              `json:"start_day_of_week,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DateRangeValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DateRangeValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DateRangeValueDynamicDateRangePb string

const DateRangeValueDynamicDateRangeLast12Months DateRangeValueDynamicDateRangePb = `LAST_12_MONTHS`

const DateRangeValueDynamicDateRangeLast14Days DateRangeValueDynamicDateRangePb = `LAST_14_DAYS`

const DateRangeValueDynamicDateRangeLast24Hours DateRangeValueDynamicDateRangePb = `LAST_24_HOURS`

const DateRangeValueDynamicDateRangeLast30Days DateRangeValueDynamicDateRangePb = `LAST_30_DAYS`

const DateRangeValueDynamicDateRangeLast60Days DateRangeValueDynamicDateRangePb = `LAST_60_DAYS`

const DateRangeValueDynamicDateRangeLast7Days DateRangeValueDynamicDateRangePb = `LAST_7_DAYS`

const DateRangeValueDynamicDateRangeLast8Hours DateRangeValueDynamicDateRangePb = `LAST_8_HOURS`

const DateRangeValueDynamicDateRangeLast90Days DateRangeValueDynamicDateRangePb = `LAST_90_DAYS`

const DateRangeValueDynamicDateRangeLastHour DateRangeValueDynamicDateRangePb = `LAST_HOUR`

const DateRangeValueDynamicDateRangeLastMonth DateRangeValueDynamicDateRangePb = `LAST_MONTH`

const DateRangeValueDynamicDateRangeLastWeek DateRangeValueDynamicDateRangePb = `LAST_WEEK`

const DateRangeValueDynamicDateRangeLastYear DateRangeValueDynamicDateRangePb = `LAST_YEAR`

const DateRangeValueDynamicDateRangeThisMonth DateRangeValueDynamicDateRangePb = `THIS_MONTH`

const DateRangeValueDynamicDateRangeThisWeek DateRangeValueDynamicDateRangePb = `THIS_WEEK`

const DateRangeValueDynamicDateRangeThisYear DateRangeValueDynamicDateRangePb = `THIS_YEAR`

const DateRangeValueDynamicDateRangeToday DateRangeValueDynamicDateRangePb = `TODAY`

const DateRangeValueDynamicDateRangeYesterday DateRangeValueDynamicDateRangePb = `YESTERDAY`

type DateValuePb struct {
	DateValue        string                 `json:"date_value,omitempty"`
	DynamicDateValue DateValueDynamicDatePb `json:"dynamic_date_value,omitempty"`
	Precision        DatePrecisionPb        `json:"precision,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DateValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DateValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DateValueDynamicDatePb string

const DateValueDynamicDateNow DateValueDynamicDatePb = `NOW`

const DateValueDynamicDateYesterday DateValueDynamicDatePb = `YESTERDAY`

type DeleteAlertsLegacyRequestPb struct {
	AlertId string `json:"-" url:"-"`
}

type DeleteDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

type DeleteDashboardWidgetRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteQueriesLegacyRequestPb struct {
	QueryId string `json:"-" url:"-"`
}

type DeleteQueryVisualizationsLegacyRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteResponsePb struct {
}

type DeleteVisualizationRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteWarehouseRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeleteWarehouseResponsePb struct {
}

type DispositionPb string

const DispositionExternalLinks DispositionPb = `EXTERNAL_LINKS`

const DispositionInline DispositionPb = `INLINE`

type EditAlertPb struct {
	AlertId string         `json:"-" url:"-"`
	Name    string         `json:"name"`
	Options AlertOptionsPb `json:"options"`
	QueryId string         `json:"query_id"`
	Rearm   int            `json:"rearm,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EditAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EditAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditWarehouseRequestPb struct {
	AutoStopMins            int                                 `json:"auto_stop_mins,omitempty"`
	Channel                 *ChannelPb                          `json:"channel,omitempty"`
	ClusterSize             string                              `json:"cluster_size,omitempty"`
	CreatorName             string                              `json:"creator_name,omitempty"`
	EnablePhoton            bool                                `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                                `json:"enable_serverless_compute,omitempty"`
	Id                      string                              `json:"-" url:"-"`
	InstanceProfileArn      string                              `json:"instance_profile_arn,omitempty"`
	MaxNumClusters          int                                 `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                                 `json:"min_num_clusters,omitempty"`
	Name                    string                              `json:"name,omitempty"`
	SpotInstancePolicy      SpotInstancePolicyPb                `json:"spot_instance_policy,omitempty"`
	Tags                    *EndpointTagsPb                     `json:"tags,omitempty"`
	WarehouseType           EditWarehouseRequestWarehouseTypePb `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EditWarehouseRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EditWarehouseRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditWarehouseRequestWarehouseTypePb string

const EditWarehouseRequestWarehouseTypeClassic EditWarehouseRequestWarehouseTypePb = `CLASSIC`

const EditWarehouseRequestWarehouseTypePro EditWarehouseRequestWarehouseTypePb = `PRO`

const EditWarehouseRequestWarehouseTypeTypeUnspecified EditWarehouseRequestWarehouseTypePb = `TYPE_UNSPECIFIED`

type EditWarehouseResponsePb struct {
}

type EmptyPb struct {
}

type EndpointConfPairPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointConfPairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointConfPairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointHealthPb struct {
	Details       string               `json:"details,omitempty"`
	FailureReason *TerminationReasonPb `json:"failure_reason,omitempty"`
	Message       string               `json:"message,omitempty"`
	Status        StatusPb             `json:"status,omitempty"`
	Summary       string               `json:"summary,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointHealthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointHealthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointInfoPb struct {
	AutoStopMins            int                         `json:"auto_stop_mins,omitempty"`
	Channel                 *ChannelPb                  `json:"channel,omitempty"`
	ClusterSize             string                      `json:"cluster_size,omitempty"`
	CreatorName             string                      `json:"creator_name,omitempty"`
	EnablePhoton            bool                        `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                        `json:"enable_serverless_compute,omitempty"`
	Health                  *EndpointHealthPb           `json:"health,omitempty"`
	Id                      string                      `json:"id,omitempty"`
	InstanceProfileArn      string                      `json:"instance_profile_arn,omitempty"`
	JdbcUrl                 string                      `json:"jdbc_url,omitempty"`
	MaxNumClusters          int                         `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                         `json:"min_num_clusters,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	NumActiveSessions       int64                       `json:"num_active_sessions,omitempty"`
	NumClusters             int                         `json:"num_clusters,omitempty"`
	OdbcParams              *OdbcParamsPb               `json:"odbc_params,omitempty"`
	SpotInstancePolicy      SpotInstancePolicyPb        `json:"spot_instance_policy,omitempty"`
	State                   StatePb                     `json:"state,omitempty"`
	Tags                    *EndpointTagsPb             `json:"tags,omitempty"`
	WarehouseType           EndpointInfoWarehouseTypePb `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointInfoWarehouseTypePb string

const EndpointInfoWarehouseTypeClassic EndpointInfoWarehouseTypePb = `CLASSIC`

const EndpointInfoWarehouseTypePro EndpointInfoWarehouseTypePb = `PRO`

const EndpointInfoWarehouseTypeTypeUnspecified EndpointInfoWarehouseTypePb = `TYPE_UNSPECIFIED`

type EndpointTagPairPb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointTagPairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointTagPairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointTagsPb struct {
	CustomTags []EndpointTagPairPb `json:"custom_tags,omitempty"`
}

type EnumValuePb struct {
	EnumOptions        string                `json:"enum_options,omitempty"`
	MultiValuesOptions *MultiValuesOptionsPb `json:"multi_values_options,omitempty"`
	Values             []string              `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnumValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnumValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExecuteStatementRequestPb struct {
	ByteLimit     int64                                  `json:"byte_limit,omitempty"`
	Catalog       string                                 `json:"catalog,omitempty"`
	Disposition   DispositionPb                          `json:"disposition,omitempty"`
	Format        FormatPb                               `json:"format,omitempty"`
	OnWaitTimeout ExecuteStatementRequestOnWaitTimeoutPb `json:"on_wait_timeout,omitempty"`
	Parameters    []StatementParameterListItemPb         `json:"parameters,omitempty"`
	RowLimit      int64                                  `json:"row_limit,omitempty"`
	Schema        string                                 `json:"schema,omitempty"`
	Statement     string                                 `json:"statement"`
	WaitTimeout   string                                 `json:"wait_timeout,omitempty"`
	WarehouseId   string                                 `json:"warehouse_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExecuteStatementRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExecuteStatementRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExecuteStatementRequestOnWaitTimeoutPb string

const ExecuteStatementRequestOnWaitTimeoutCancel ExecuteStatementRequestOnWaitTimeoutPb = `CANCEL`

const ExecuteStatementRequestOnWaitTimeoutContinue ExecuteStatementRequestOnWaitTimeoutPb = `CONTINUE`

type ExternalLinkPb struct {
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

func (st *ExternalLinkPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalLinkPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalQuerySourcePb struct {
	AlertId           string                        `json:"alert_id,omitempty"`
	DashboardId       string                        `json:"dashboard_id,omitempty"`
	GenieSpaceId      string                        `json:"genie_space_id,omitempty"`
	JobInfo           *ExternalQuerySourceJobInfoPb `json:"job_info,omitempty"`
	LegacyDashboardId string                        `json:"legacy_dashboard_id,omitempty"`
	NotebookId        string                        `json:"notebook_id,omitempty"`
	SqlQueryId        string                        `json:"sql_query_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalQuerySourcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalQuerySourcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalQuerySourceJobInfoPb struct {
	JobId        string `json:"job_id,omitempty"`
	JobRunId     string `json:"job_run_id,omitempty"`
	JobTaskRunId string `json:"job_task_run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalQuerySourceJobInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalQuerySourceJobInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FormatPb string

const FormatArrowStream FormatPb = `ARROW_STREAM`

const FormatCsv FormatPb = `CSV`

const FormatJsonArray FormatPb = `JSON_ARRAY`

type GetAlertRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetAlertV2RequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetAlertsLegacyRequestPb struct {
	AlertId string `json:"-" url:"-"`
}

type GetConfigRequestPb struct {
}

type GetDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

type GetDbsqlPermissionRequestPb struct {
	ObjectId   string             `json:"-" url:"-"`
	ObjectType ObjectTypePluralPb `json:"-" url:"-"`
}

type GetQueriesLegacyRequestPb struct {
	QueryId string `json:"-" url:"-"`
}

type GetQueryRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetResponsePb struct {
	AccessControlList []AccessControlPb `json:"access_control_list,omitempty"`
	ObjectId          string            `json:"object_id,omitempty"`
	ObjectType        ObjectTypePb      `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetStatementRequestPb struct {
	StatementId string `json:"-" url:"-"`
}

type GetStatementResultChunkNRequestPb struct {
	ChunkIndex  int    `json:"-" url:"-"`
	StatementId string `json:"-" url:"-"`
}

type GetWarehousePermissionLevelsRequestPb struct {
	WarehouseId string `json:"-" url:"-"`
}

type GetWarehousePermissionLevelsResponsePb struct {
	PermissionLevels []WarehousePermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetWarehousePermissionsRequestPb struct {
	WarehouseId string `json:"-" url:"-"`
}

type GetWarehouseRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetWarehouseResponsePb struct {
	AutoStopMins            int                                 `json:"auto_stop_mins,omitempty"`
	Channel                 *ChannelPb                          `json:"channel,omitempty"`
	ClusterSize             string                              `json:"cluster_size,omitempty"`
	CreatorName             string                              `json:"creator_name,omitempty"`
	EnablePhoton            bool                                `json:"enable_photon,omitempty"`
	EnableServerlessCompute bool                                `json:"enable_serverless_compute,omitempty"`
	Health                  *EndpointHealthPb                   `json:"health,omitempty"`
	Id                      string                              `json:"id,omitempty"`
	InstanceProfileArn      string                              `json:"instance_profile_arn,omitempty"`
	JdbcUrl                 string                              `json:"jdbc_url,omitempty"`
	MaxNumClusters          int                                 `json:"max_num_clusters,omitempty"`
	MinNumClusters          int                                 `json:"min_num_clusters,omitempty"`
	Name                    string                              `json:"name,omitempty"`
	NumActiveSessions       int64                               `json:"num_active_sessions,omitempty"`
	NumClusters             int                                 `json:"num_clusters,omitempty"`
	OdbcParams              *OdbcParamsPb                       `json:"odbc_params,omitempty"`
	SpotInstancePolicy      SpotInstancePolicyPb                `json:"spot_instance_policy,omitempty"`
	State                   StatePb                             `json:"state,omitempty"`
	Tags                    *EndpointTagsPb                     `json:"tags,omitempty"`
	WarehouseType           GetWarehouseResponseWarehouseTypePb `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetWarehouseResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetWarehouseResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetWarehouseResponseWarehouseTypePb string

const GetWarehouseResponseWarehouseTypeClassic GetWarehouseResponseWarehouseTypePb = `CLASSIC`

const GetWarehouseResponseWarehouseTypePro GetWarehouseResponseWarehouseTypePb = `PRO`

const GetWarehouseResponseWarehouseTypeTypeUnspecified GetWarehouseResponseWarehouseTypePb = `TYPE_UNSPECIFIED`

type GetWorkspaceWarehouseConfigRequestPb struct {
}

type GetWorkspaceWarehouseConfigResponsePb struct {
	Channel                    *ChannelPb                                          `json:"channel,omitempty"`
	ConfigParam                *RepeatedEndpointConfPairsPb                        `json:"config_param,omitempty"`
	DataAccessConfig           []EndpointConfPairPb                                `json:"data_access_config,omitempty"`
	EnabledWarehouseTypes      []WarehouseTypePairPb                               `json:"enabled_warehouse_types,omitempty"`
	GlobalParam                *RepeatedEndpointConfPairsPb                        `json:"global_param,omitempty"`
	GoogleServiceAccount       string                                              `json:"google_service_account,omitempty"`
	InstanceProfileArn         string                                              `json:"instance_profile_arn,omitempty"`
	SecurityPolicy             GetWorkspaceWarehouseConfigResponseSecurityPolicyPb `json:"security_policy,omitempty"`
	SqlConfigurationParameters *RepeatedEndpointConfPairsPb                        `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetWorkspaceWarehouseConfigResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetWorkspaceWarehouseConfigResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetWorkspaceWarehouseConfigResponseSecurityPolicyPb string

const GetWorkspaceWarehouseConfigResponseSecurityPolicyDataAccessControl GetWorkspaceWarehouseConfigResponseSecurityPolicyPb = `DATA_ACCESS_CONTROL`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyNone GetWorkspaceWarehouseConfigResponseSecurityPolicyPb = `NONE`

const GetWorkspaceWarehouseConfigResponseSecurityPolicyPassthrough GetWorkspaceWarehouseConfigResponseSecurityPolicyPb = `PASSTHROUGH`

type LegacyAlertPb struct {
	CreatedAt       string             `json:"created_at,omitempty"`
	Id              string             `json:"id,omitempty"`
	LastTriggeredAt string             `json:"last_triggered_at,omitempty"`
	Name            string             `json:"name,omitempty"`
	Options         *AlertOptionsPb    `json:"options,omitempty"`
	Parent          string             `json:"parent,omitempty"`
	Query           *AlertQueryPb      `json:"query,omitempty"`
	Rearm           int                `json:"rearm,omitempty"`
	State           LegacyAlertStatePb `json:"state,omitempty"`
	UpdatedAt       string             `json:"updated_at,omitempty"`
	User            *UserPb            `json:"user,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LegacyAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LegacyAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LegacyAlertStatePb string

const LegacyAlertStateOk LegacyAlertStatePb = `ok`

const LegacyAlertStateTriggered LegacyAlertStatePb = `triggered`

const LegacyAlertStateUnknown LegacyAlertStatePb = `unknown`

type LegacyQueryPb struct {
	CanEdit           bool                    `json:"can_edit,omitempty"`
	CreatedAt         string                  `json:"created_at,omitempty"`
	DataSourceId      string                  `json:"data_source_id,omitempty"`
	Description       string                  `json:"description,omitempty"`
	Id                string                  `json:"id,omitempty"`
	IsArchived        bool                    `json:"is_archived,omitempty"`
	IsDraft           bool                    `json:"is_draft,omitempty"`
	IsFavorite        bool                    `json:"is_favorite,omitempty"`
	IsSafe            bool                    `json:"is_safe,omitempty"`
	LastModifiedBy    *UserPb                 `json:"last_modified_by,omitempty"`
	LastModifiedById  int                     `json:"last_modified_by_id,omitempty"`
	LatestQueryDataId string                  `json:"latest_query_data_id,omitempty"`
	Name              string                  `json:"name,omitempty"`
	Options           *QueryOptionsPb         `json:"options,omitempty"`
	Parent            string                  `json:"parent,omitempty"`
	PermissionTier    PermissionLevelPb       `json:"permission_tier,omitempty"`
	Query             string                  `json:"query,omitempty"`
	QueryHash         string                  `json:"query_hash,omitempty"`
	RunAsRole         RunAsRolePb             `json:"run_as_role,omitempty"`
	Tags              []string                `json:"tags,omitempty"`
	UpdatedAt         string                  `json:"updated_at,omitempty"`
	User              *UserPb                 `json:"user,omitempty"`
	UserId            int                     `json:"user_id,omitempty"`
	Visualizations    []LegacyVisualizationPb `json:"visualizations,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LegacyQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LegacyQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LegacyVisualizationPb struct {
	CreatedAt   string         `json:"created_at,omitempty"`
	Description string         `json:"description,omitempty"`
	Id          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Options     any            `json:"options,omitempty"`
	Query       *LegacyQueryPb `json:"query,omitempty"`
	Type        string         `json:"type,omitempty"`
	UpdatedAt   string         `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LegacyVisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LegacyVisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LifecycleStatePb string

const LifecycleStateActive LifecycleStatePb = `ACTIVE`

const LifecycleStateTrashed LifecycleStatePb = `TRASHED`

type ListAlertsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAlertsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAlertsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsResponsePb struct {
	NextPageToken string                      `json:"next_page_token,omitempty"`
	Results       []ListAlertsResponseAlertPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAlertsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAlertsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsResponseAlertPb struct {
	Condition          *AlertConditionPb `json:"condition,omitempty"`
	CreateTime         string            `json:"create_time,omitempty"`
	CustomBody         string            `json:"custom_body,omitempty"`
	CustomSubject      string            `json:"custom_subject,omitempty"`
	DisplayName        string            `json:"display_name,omitempty"`
	Id                 string            `json:"id,omitempty"`
	LifecycleState     LifecycleStatePb  `json:"lifecycle_state,omitempty"`
	NotifyOnOk         bool              `json:"notify_on_ok,omitempty"`
	OwnerUserName      string            `json:"owner_user_name,omitempty"`
	QueryId            string            `json:"query_id,omitempty"`
	SecondsToRetrigger int               `json:"seconds_to_retrigger,omitempty"`
	State              AlertStatePb      `json:"state,omitempty"`
	TriggerTime        string            `json:"trigger_time,omitempty"`
	UpdateTime         string            `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAlertsResponseAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAlertsResponseAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsV2RequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAlertsV2RequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAlertsV2RequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListAlertsV2ResponsePb struct {
	NextPageToken string      `json:"next_page_token,omitempty"`
	Results       []AlertV2Pb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAlertsV2ResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAlertsV2ResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDashboardsRequestPb struct {
	Order    ListOrderPb `json:"-" url:"order,omitempty"`
	Page     int         `json:"-" url:"page,omitempty"`
	PageSize int         `json:"-" url:"page_size,omitempty"`
	Q        string      `json:"-" url:"q,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDashboardsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDashboardsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListOrderPb string

const ListOrderCreatedAt ListOrderPb = `created_at`

const ListOrderName ListOrderPb = `name`

type ListQueriesLegacyRequestPb struct {
	Order    string `json:"-" url:"order,omitempty"`
	Page     int    `json:"-" url:"page,omitempty"`
	PageSize int    `json:"-" url:"page_size,omitempty"`
	Q        string `json:"-" url:"q,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQueriesLegacyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQueriesLegacyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueriesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQueriesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQueriesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueriesResponsePb struct {
	HasNextPage   bool          `json:"has_next_page,omitempty"`
	NextPageToken string        `json:"next_page_token,omitempty"`
	Res           []QueryInfoPb `json:"res,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQueriesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQueriesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueryHistoryRequestPb struct {
	FilterBy       *QueryFilterPb `json:"-" url:"filter_by,omitempty"`
	IncludeMetrics bool           `json:"-" url:"include_metrics,omitempty"`
	MaxResults     int            `json:"-" url:"max_results,omitempty"`
	PageToken      string         `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQueryHistoryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQueryHistoryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueryObjectsResponsePb struct {
	NextPageToken string                            `json:"next_page_token,omitempty"`
	Results       []ListQueryObjectsResponseQueryPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQueryObjectsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQueryObjectsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQueryObjectsResponseQueryPb struct {
	ApplyAutoLimit       bool               `json:"apply_auto_limit,omitempty"`
	Catalog              string             `json:"catalog,omitempty"`
	CreateTime           string             `json:"create_time,omitempty"`
	Description          string             `json:"description,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
	Id                   string             `json:"id,omitempty"`
	LastModifierUserName string             `json:"last_modifier_user_name,omitempty"`
	LifecycleState       LifecycleStatePb   `json:"lifecycle_state,omitempty"`
	OwnerUserName        string             `json:"owner_user_name,omitempty"`
	Parameters           []QueryParameterPb `json:"parameters,omitempty"`
	QueryText            string             `json:"query_text,omitempty"`
	RunAsMode            RunAsModePb        `json:"run_as_mode,omitempty"`
	Schema               string             `json:"schema,omitempty"`
	Tags                 []string           `json:"tags,omitempty"`
	UpdateTime           string             `json:"update_time,omitempty"`
	WarehouseId          string             `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQueryObjectsResponseQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQueryObjectsResponseQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRequestPb struct {
}

type ListResponsePb struct {
	Count    int           `json:"count,omitempty"`
	Page     int           `json:"page,omitempty"`
	PageSize int           `json:"page_size,omitempty"`
	Results  []DashboardPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListVisualizationsForQueryRequestPb struct {
	Id        string `json:"-" url:"-"`
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListVisualizationsForQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListVisualizationsForQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListVisualizationsForQueryResponsePb struct {
	NextPageToken string            `json:"next_page_token,omitempty"`
	Results       []VisualizationPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListVisualizationsForQueryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListVisualizationsForQueryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListWarehousesRequestPb struct {
	RunAsUserId int `json:"-" url:"run_as_user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListWarehousesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListWarehousesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListWarehousesResponsePb struct {
	Warehouses []EndpointInfoPb `json:"warehouses,omitempty"`
}

type MultiValuesOptionsPb struct {
	Prefix    string `json:"prefix,omitempty"`
	Separator string `json:"separator,omitempty"`
	Suffix    string `json:"suffix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MultiValuesOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MultiValuesOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NumericValuePb struct {
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NumericValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NumericValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ObjectTypePb string

const ObjectTypeAlert ObjectTypePb = `alert`

const ObjectTypeDashboard ObjectTypePb = `dashboard`

const ObjectTypeDataSource ObjectTypePb = `data_source`

const ObjectTypeQuery ObjectTypePb = `query`

type ObjectTypePluralPb string

const ObjectTypePluralAlerts ObjectTypePluralPb = `alerts`

const ObjectTypePluralDashboards ObjectTypePluralPb = `dashboards`

const ObjectTypePluralDataSources ObjectTypePluralPb = `data_sources`

const ObjectTypePluralQueries ObjectTypePluralPb = `queries`

type OdbcParamsPb struct {
	Hostname string `json:"hostname,omitempty"`
	Path     string `json:"path,omitempty"`
	Port     int    `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OdbcParamsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OdbcParamsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OwnableObjectTypePb string

const OwnableObjectTypeAlert OwnableObjectTypePb = `alert`

const OwnableObjectTypeDashboard OwnableObjectTypePb = `dashboard`

const OwnableObjectTypeQuery OwnableObjectTypePb = `query`

type ParameterPb struct {
	EnumOptions        string                `json:"enumOptions,omitempty"`
	MultiValuesOptions *MultiValuesOptionsPb `json:"multiValuesOptions,omitempty"`
	Name               string                `json:"name,omitempty"`
	QueryId            string                `json:"queryId,omitempty"`
	Title              string                `json:"title,omitempty"`
	Type               ParameterTypePb       `json:"type,omitempty"`
	Value              any                   `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ParameterTypePb string

const ParameterTypeDatetime ParameterTypePb = `datetime`

const ParameterTypeEnum ParameterTypePb = `enum`

const ParameterTypeNumber ParameterTypePb = `number`

const ParameterTypeQuery ParameterTypePb = `query`

const ParameterTypeText ParameterTypePb = `text`

type PermissionLevelPb string

// Can edit the query
const PermissionLevelCanEdit PermissionLevelPb = `CAN_EDIT`

// Can manage the query
const PermissionLevelCanManage PermissionLevelPb = `CAN_MANAGE`

// Can run the query
const PermissionLevelCanRun PermissionLevelPb = `CAN_RUN`

// Can view the query
const PermissionLevelCanView PermissionLevelPb = `CAN_VIEW`

type PlansStatePb string

const PlansStateEmpty PlansStatePb = `EMPTY`

const PlansStateExists PlansStatePb = `EXISTS`

const PlansStateIgnoredLargePlansSize PlansStatePb = `IGNORED_LARGE_PLANS_SIZE`

const PlansStateIgnoredSmallDuration PlansStatePb = `IGNORED_SMALL_DURATION`

const PlansStateIgnoredSparkPlanType PlansStatePb = `IGNORED_SPARK_PLAN_TYPE`

const PlansStateUnknown PlansStatePb = `UNKNOWN`

type QueryPb struct {
	ApplyAutoLimit       bool               `json:"apply_auto_limit,omitempty"`
	Catalog              string             `json:"catalog,omitempty"`
	CreateTime           string             `json:"create_time,omitempty"`
	Description          string             `json:"description,omitempty"`
	DisplayName          string             `json:"display_name,omitempty"`
	Id                   string             `json:"id,omitempty"`
	LastModifierUserName string             `json:"last_modifier_user_name,omitempty"`
	LifecycleState       LifecycleStatePb   `json:"lifecycle_state,omitempty"`
	OwnerUserName        string             `json:"owner_user_name,omitempty"`
	Parameters           []QueryParameterPb `json:"parameters,omitempty"`
	ParentPath           string             `json:"parent_path,omitempty"`
	QueryText            string             `json:"query_text,omitempty"`
	RunAsMode            RunAsModePb        `json:"run_as_mode,omitempty"`
	Schema               string             `json:"schema,omitempty"`
	Tags                 []string           `json:"tags,omitempty"`
	UpdateTime           string             `json:"update_time,omitempty"`
	WarehouseId          string             `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryBackedValuePb struct {
	MultiValuesOptions *MultiValuesOptionsPb `json:"multi_values_options,omitempty"`
	QueryId            string                `json:"query_id,omitempty"`
	Values             []string              `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryBackedValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryBackedValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryEditContentPb struct {
	DataSourceId string      `json:"data_source_id,omitempty"`
	Description  string      `json:"description,omitempty"`
	Name         string      `json:"name,omitempty"`
	Options      any         `json:"options,omitempty"`
	Query        string      `json:"query,omitempty"`
	QueryId      string      `json:"-" url:"-"`
	RunAsRole    RunAsRolePb `json:"run_as_role,omitempty"`
	Tags         []string    `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryEditContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryEditContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryFilterPb struct {
	QueryStartTimeRange *TimeRangePb    `json:"query_start_time_range,omitempty" url:"query_start_time_range,omitempty"`
	StatementIds        []string        `json:"statement_ids,omitempty" url:"statement_ids,omitempty"`
	Statuses            []QueryStatusPb `json:"statuses,omitempty" url:"statuses,omitempty"`
	UserIds             []int64         `json:"user_ids,omitempty" url:"user_ids,omitempty"`
	WarehouseIds        []string        `json:"warehouse_ids,omitempty" url:"warehouse_ids,omitempty"`
}

type QueryInfoPb struct {
	ChannelUsed        *ChannelInfoPb         `json:"channel_used,omitempty"`
	ClientApplication  string                 `json:"client_application,omitempty"`
	Duration           int64                  `json:"duration,omitempty"`
	EndpointId         string                 `json:"endpoint_id,omitempty"`
	ErrorMessage       string                 `json:"error_message,omitempty"`
	ExecutedAsUserId   int64                  `json:"executed_as_user_id,omitempty"`
	ExecutedAsUserName string                 `json:"executed_as_user_name,omitempty"`
	ExecutionEndTimeMs int64                  `json:"execution_end_time_ms,omitempty"`
	IsFinal            bool                   `json:"is_final,omitempty"`
	LookupKey          string                 `json:"lookup_key,omitempty"`
	Metrics            *QueryMetricsPb        `json:"metrics,omitempty"`
	PlansState         PlansStatePb           `json:"plans_state,omitempty"`
	QueryEndTimeMs     int64                  `json:"query_end_time_ms,omitempty"`
	QueryId            string                 `json:"query_id,omitempty"`
	QuerySource        *ExternalQuerySourcePb `json:"query_source,omitempty"`
	QueryStartTimeMs   int64                  `json:"query_start_time_ms,omitempty"`
	QueryText          string                 `json:"query_text,omitempty"`
	RowsProduced       int64                  `json:"rows_produced,omitempty"`
	SparkUiUrl         string                 `json:"spark_ui_url,omitempty"`
	StatementType      QueryStatementTypePb   `json:"statement_type,omitempty"`
	Status             QueryStatusPb          `json:"status,omitempty"`
	UserId             int64                  `json:"user_id,omitempty"`
	UserName           string                 `json:"user_name,omitempty"`
	WarehouseId        string                 `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryListPb struct {
	Count    int             `json:"count,omitempty"`
	Page     int             `json:"page,omitempty"`
	PageSize int             `json:"page_size,omitempty"`
	Results  []LegacyQueryPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryMetricsPb struct {
	CompilationTimeMs                 int64                `json:"compilation_time_ms,omitempty"`
	ExecutionTimeMs                   int64                `json:"execution_time_ms,omitempty"`
	NetworkSentBytes                  int64                `json:"network_sent_bytes,omitempty"`
	OverloadingQueueStartTimestamp    int64                `json:"overloading_queue_start_timestamp,omitempty"`
	PhotonTotalTimeMs                 int64                `json:"photon_total_time_ms,omitempty"`
	ProjectedRemainingTaskTotalTimeMs int64                `json:"projected_remaining_task_total_time_ms,omitempty"`
	ProvisioningQueueStartTimestamp   int64                `json:"provisioning_queue_start_timestamp,omitempty"`
	PrunedBytes                       int64                `json:"pruned_bytes,omitempty"`
	PrunedFilesCount                  int64                `json:"pruned_files_count,omitempty"`
	QueryCompilationStartTimestamp    int64                `json:"query_compilation_start_timestamp,omitempty"`
	ReadBytes                         int64                `json:"read_bytes,omitempty"`
	ReadCacheBytes                    int64                `json:"read_cache_bytes,omitempty"`
	ReadFilesCount                    int64                `json:"read_files_count,omitempty"`
	ReadPartitionsCount               int64                `json:"read_partitions_count,omitempty"`
	ReadRemoteBytes                   int64                `json:"read_remote_bytes,omitempty"`
	RemainingTaskCount                int64                `json:"remaining_task_count,omitempty"`
	ResultFetchTimeMs                 int64                `json:"result_fetch_time_ms,omitempty"`
	ResultFromCache                   bool                 `json:"result_from_cache,omitempty"`
	RowsProducedCount                 int64                `json:"rows_produced_count,omitempty"`
	RowsReadCount                     int64                `json:"rows_read_count,omitempty"`
	RunnableTasks                     int64                `json:"runnable_tasks,omitempty"`
	SpillToDiskBytes                  int64                `json:"spill_to_disk_bytes,omitempty"`
	TaskTimeOverTimeRange             *TaskTimeOverRangePb `json:"task_time_over_time_range,omitempty"`
	TaskTotalTimeMs                   int64                `json:"task_total_time_ms,omitempty"`
	TotalTimeMs                       int64                `json:"total_time_ms,omitempty"`
	WorkToBeDone                      int64                `json:"work_to_be_done,omitempty"`
	WriteRemoteBytes                  int64                `json:"write_remote_bytes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryMetricsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryMetricsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryOptionsPb struct {
	Catalog        string        `json:"catalog,omitempty"`
	MovedToTrashAt string        `json:"moved_to_trash_at,omitempty"`
	Parameters     []ParameterPb `json:"parameters,omitempty"`
	Schema         string        `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryParameterPb struct {
	DateRangeValue   *DateRangeValuePb   `json:"date_range_value,omitempty"`
	DateValue        *DateValuePb        `json:"date_value,omitempty"`
	EnumValue        *EnumValuePb        `json:"enum_value,omitempty"`
	Name             string              `json:"name,omitempty"`
	NumericValue     *NumericValuePb     `json:"numeric_value,omitempty"`
	QueryBackedValue *QueryBackedValuePb `json:"query_backed_value,omitempty"`
	TextValue        *TextValuePb        `json:"text_value,omitempty"`
	Title            string              `json:"title,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryParameterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryParameterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryPostContentPb struct {
	DataSourceId string      `json:"data_source_id,omitempty"`
	Description  string      `json:"description,omitempty"`
	Name         string      `json:"name,omitempty"`
	Options      any         `json:"options,omitempty"`
	Parent       string      `json:"parent,omitempty"`
	Query        string      `json:"query,omitempty"`
	RunAsRole    RunAsRolePb `json:"run_as_role,omitempty"`
	Tags         []string    `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryPostContentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryPostContentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryStatementTypePb string

const QueryStatementTypeAlter QueryStatementTypePb = `ALTER`

const QueryStatementTypeAnalyze QueryStatementTypePb = `ANALYZE`

const QueryStatementTypeCopy QueryStatementTypePb = `COPY`

const QueryStatementTypeCreate QueryStatementTypePb = `CREATE`

const QueryStatementTypeDelete QueryStatementTypePb = `DELETE`

const QueryStatementTypeDescribe QueryStatementTypePb = `DESCRIBE`

const QueryStatementTypeDrop QueryStatementTypePb = `DROP`

const QueryStatementTypeExplain QueryStatementTypePb = `EXPLAIN`

const QueryStatementTypeGrant QueryStatementTypePb = `GRANT`

const QueryStatementTypeInsert QueryStatementTypePb = `INSERT`

const QueryStatementTypeMerge QueryStatementTypePb = `MERGE`

const QueryStatementTypeOptimize QueryStatementTypePb = `OPTIMIZE`

const QueryStatementTypeOther QueryStatementTypePb = `OTHER`

const QueryStatementTypeRefresh QueryStatementTypePb = `REFRESH`

const QueryStatementTypeReplace QueryStatementTypePb = `REPLACE`

const QueryStatementTypeRevoke QueryStatementTypePb = `REVOKE`

const QueryStatementTypeSelect QueryStatementTypePb = `SELECT`

const QueryStatementTypeSet QueryStatementTypePb = `SET`

const QueryStatementTypeShow QueryStatementTypePb = `SHOW`

const QueryStatementTypeTruncate QueryStatementTypePb = `TRUNCATE`

const QueryStatementTypeUpdate QueryStatementTypePb = `UPDATE`

const QueryStatementTypeUse QueryStatementTypePb = `USE`

type QueryStatusPb string

const QueryStatusCanceled QueryStatusPb = `CANCELED`

const QueryStatusCompiled QueryStatusPb = `COMPILED`

const QueryStatusCompiling QueryStatusPb = `COMPILING`

const QueryStatusFailed QueryStatusPb = `FAILED`

const QueryStatusFinished QueryStatusPb = `FINISHED`

const QueryStatusQueued QueryStatusPb = `QUEUED`

const QueryStatusRunning QueryStatusPb = `RUNNING`

const QueryStatusStarted QueryStatusPb = `STARTED`

type RepeatedEndpointConfPairsPb struct {
	ConfigPair         []EndpointConfPairPb `json:"config_pair,omitempty"`
	ConfigurationPairs []EndpointConfPairPb `json:"configuration_pairs,omitempty"`
}

type RestoreDashboardRequestPb struct {
	DashboardId string `json:"-" url:"-"`
}

type RestoreQueriesLegacyRequestPb struct {
	QueryId string `json:"-" url:"-"`
}

type RestoreResponsePb struct {
}

type ResultDataPb struct {
	ByteCount             int64            `json:"byte_count,omitempty"`
	ChunkIndex            int              `json:"chunk_index,omitempty"`
	DataArray             [][]string       `json:"data_array,omitempty"`
	ExternalLinks         []ExternalLinkPb `json:"external_links,omitempty"`
	NextChunkIndex        int              `json:"next_chunk_index,omitempty"`
	NextChunkInternalLink string           `json:"next_chunk_internal_link,omitempty"`
	RowCount              int64            `json:"row_count,omitempty"`
	RowOffset             int64            `json:"row_offset,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResultDataPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResultDataPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResultManifestPb struct {
	Chunks          []BaseChunkInfoPb `json:"chunks,omitempty"`
	Format          FormatPb          `json:"format,omitempty"`
	Schema          *ResultSchemaPb   `json:"schema,omitempty"`
	TotalByteCount  int64             `json:"total_byte_count,omitempty"`
	TotalChunkCount int               `json:"total_chunk_count,omitempty"`
	TotalRowCount   int64             `json:"total_row_count,omitempty"`
	Truncated       bool              `json:"truncated,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResultManifestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResultManifestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ResultSchemaPb struct {
	ColumnCount int            `json:"column_count,omitempty"`
	Columns     []ColumnInfoPb `json:"columns,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ResultSchemaPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ResultSchemaPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunAsModePb string

const RunAsModeOwner RunAsModePb = `OWNER`

const RunAsModeViewer RunAsModePb = `VIEWER`

type RunAsRolePb string

const RunAsRoleOwner RunAsRolePb = `owner`

const RunAsRoleViewer RunAsRolePb = `viewer`

type SchedulePauseStatusPb string

const SchedulePauseStatusPaused SchedulePauseStatusPb = `PAUSED`

const SchedulePauseStatusUnpaused SchedulePauseStatusPb = `UNPAUSED`

type ServiceErrorPb struct {
	ErrorCode ServiceErrorCodePb `json:"error_code,omitempty"`
	Message   string             `json:"message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServiceErrorPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServiceErrorPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServiceErrorCodePb string

const ServiceErrorCodeAborted ServiceErrorCodePb = `ABORTED`

const ServiceErrorCodeAlreadyExists ServiceErrorCodePb = `ALREADY_EXISTS`

const ServiceErrorCodeBadRequest ServiceErrorCodePb = `BAD_REQUEST`

const ServiceErrorCodeCancelled ServiceErrorCodePb = `CANCELLED`

const ServiceErrorCodeDeadlineExceeded ServiceErrorCodePb = `DEADLINE_EXCEEDED`

const ServiceErrorCodeInternalError ServiceErrorCodePb = `INTERNAL_ERROR`

const ServiceErrorCodeIoError ServiceErrorCodePb = `IO_ERROR`

const ServiceErrorCodeNotFound ServiceErrorCodePb = `NOT_FOUND`

const ServiceErrorCodeResourceExhausted ServiceErrorCodePb = `RESOURCE_EXHAUSTED`

const ServiceErrorCodeServiceUnderMaintenance ServiceErrorCodePb = `SERVICE_UNDER_MAINTENANCE`

const ServiceErrorCodeTemporarilyUnavailable ServiceErrorCodePb = `TEMPORARILY_UNAVAILABLE`

const ServiceErrorCodeUnauthenticated ServiceErrorCodePb = `UNAUTHENTICATED`

const ServiceErrorCodeUnknown ServiceErrorCodePb = `UNKNOWN`

const ServiceErrorCodeWorkspaceTemporarilyUnavailable ServiceErrorCodePb = `WORKSPACE_TEMPORARILY_UNAVAILABLE`

type SetRequestPb struct {
	AccessControlList []AccessControlPb  `json:"access_control_list,omitempty"`
	ObjectId          string             `json:"-" url:"-"`
	ObjectType        ObjectTypePluralPb `json:"-" url:"-"`
}

type SetResponsePb struct {
	AccessControlList []AccessControlPb `json:"access_control_list,omitempty"`
	ObjectId          string            `json:"object_id,omitempty"`
	ObjectType        ObjectTypePb      `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SetResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SetResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetWorkspaceWarehouseConfigRequestPb struct {
	Channel                    *ChannelPb                                         `json:"channel,omitempty"`
	ConfigParam                *RepeatedEndpointConfPairsPb                       `json:"config_param,omitempty"`
	DataAccessConfig           []EndpointConfPairPb                               `json:"data_access_config,omitempty"`
	EnabledWarehouseTypes      []WarehouseTypePairPb                              `json:"enabled_warehouse_types,omitempty"`
	GlobalParam                *RepeatedEndpointConfPairsPb                       `json:"global_param,omitempty"`
	GoogleServiceAccount       string                                             `json:"google_service_account,omitempty"`
	InstanceProfileArn         string                                             `json:"instance_profile_arn,omitempty"`
	SecurityPolicy             SetWorkspaceWarehouseConfigRequestSecurityPolicyPb `json:"security_policy,omitempty"`
	SqlConfigurationParameters *RepeatedEndpointConfPairsPb                       `json:"sql_configuration_parameters,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SetWorkspaceWarehouseConfigRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SetWorkspaceWarehouseConfigRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetWorkspaceWarehouseConfigRequestSecurityPolicyPb string

const SetWorkspaceWarehouseConfigRequestSecurityPolicyDataAccessControl SetWorkspaceWarehouseConfigRequestSecurityPolicyPb = `DATA_ACCESS_CONTROL`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyNone SetWorkspaceWarehouseConfigRequestSecurityPolicyPb = `NONE`

const SetWorkspaceWarehouseConfigRequestSecurityPolicyPassthrough SetWorkspaceWarehouseConfigRequestSecurityPolicyPb = `PASSTHROUGH`

type SetWorkspaceWarehouseConfigResponsePb struct {
}

type SpotInstancePolicyPb string

const SpotInstancePolicyCostOptimized SpotInstancePolicyPb = `COST_OPTIMIZED`

const SpotInstancePolicyPolicyUnspecified SpotInstancePolicyPb = `POLICY_UNSPECIFIED`

const SpotInstancePolicyReliabilityOptimized SpotInstancePolicyPb = `RELIABILITY_OPTIMIZED`

type StartRequestPb struct {
	Id string `json:"-" url:"-"`
}

type StartWarehouseResponsePb struct {
}

type StatePb string

const StateDeleted StatePb = `DELETED`

const StateDeleting StatePb = `DELETING`

const StateRunning StatePb = `RUNNING`

const StateStarting StatePb = `STARTING`

const StateStopped StatePb = `STOPPED`

const StateStopping StatePb = `STOPPING`

type StatementParameterListItemPb struct {
	Name  string `json:"name"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StatementParameterListItemPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StatementParameterListItemPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StatementResponsePb struct {
	Manifest    *ResultManifestPb  `json:"manifest,omitempty"`
	Result      *ResultDataPb      `json:"result,omitempty"`
	StatementId string             `json:"statement_id,omitempty"`
	Status      *StatementStatusPb `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StatementResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StatementResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StatementStatePb string

// user canceled; can come from explicit cancel call, or timeout with
// `on_wait_timeout=CANCEL`
const StatementStateCanceled StatementStatePb = `CANCELED`

// execution successful, and statement closed; result no longer available for
// fetch
const StatementStateClosed StatementStatePb = `CLOSED`

// execution failed; reason for failure described in accomanying error message
const StatementStateFailed StatementStatePb = `FAILED`

// waiting for warehouse
const StatementStatePending StatementStatePb = `PENDING`

// running
const StatementStateRunning StatementStatePb = `RUNNING`

// execution was successful, result data available for fetch
const StatementStateSucceeded StatementStatePb = `SUCCEEDED`

type StatementStatusPb struct {
	Error *ServiceErrorPb  `json:"error,omitempty"`
	State StatementStatePb `json:"state,omitempty"`
}

type StatusPb string

const StatusDegraded StatusPb = `DEGRADED`

const StatusFailed StatusPb = `FAILED`

const StatusHealthy StatusPb = `HEALTHY`

const StatusStatusUnspecified StatusPb = `STATUS_UNSPECIFIED`

type StopRequestPb struct {
	Id string `json:"-" url:"-"`
}

type StopWarehouseResponsePb struct {
}

type SuccessPb struct {
	Message SuccessMessagePb `json:"message,omitempty"`
}

type SuccessMessagePb string

const SuccessMessageSuccess SuccessMessagePb = `Success`

type TaskTimeOverRangePb struct {
	Entries  []TaskTimeOverRangeEntryPb `json:"entries,omitempty"`
	Interval int64                      `json:"interval,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TaskTimeOverRangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TaskTimeOverRangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TaskTimeOverRangeEntryPb struct {
	TaskCompletedTimeMs int64 `json:"task_completed_time_ms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TaskTimeOverRangeEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TaskTimeOverRangeEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TerminationReasonPb struct {
	Code       TerminationReasonCodePb `json:"code,omitempty"`
	Parameters map[string]string       `json:"parameters,omitempty"`
	Type       TerminationReasonTypePb `json:"type,omitempty"`
}

type TerminationReasonCodePb string

const TerminationReasonCodeAbuseDetected TerminationReasonCodePb = `ABUSE_DETECTED`

const TerminationReasonCodeAttachProjectFailure TerminationReasonCodePb = `ATTACH_PROJECT_FAILURE`

const TerminationReasonCodeAwsAuthorizationFailure TerminationReasonCodePb = `AWS_AUTHORIZATION_FAILURE`

const TerminationReasonCodeAwsInsufficientFreeAddressesInSubnetFailure TerminationReasonCodePb = `AWS_INSUFFICIENT_FREE_ADDRESSES_IN_SUBNET_FAILURE`

const TerminationReasonCodeAwsInsufficientInstanceCapacityFailure TerminationReasonCodePb = `AWS_INSUFFICIENT_INSTANCE_CAPACITY_FAILURE`

const TerminationReasonCodeAwsMaxSpotInstanceCountExceededFailure TerminationReasonCodePb = `AWS_MAX_SPOT_INSTANCE_COUNT_EXCEEDED_FAILURE`

const TerminationReasonCodeAwsRequestLimitExceeded TerminationReasonCodePb = `AWS_REQUEST_LIMIT_EXCEEDED`

const TerminationReasonCodeAwsUnsupportedFailure TerminationReasonCodePb = `AWS_UNSUPPORTED_FAILURE`

const TerminationReasonCodeAzureByokKeyPermissionFailure TerminationReasonCodePb = `AZURE_BYOK_KEY_PERMISSION_FAILURE`

const TerminationReasonCodeAzureEphemeralDiskFailure TerminationReasonCodePb = `AZURE_EPHEMERAL_DISK_FAILURE`

const TerminationReasonCodeAzureInvalidDeploymentTemplate TerminationReasonCodePb = `AZURE_INVALID_DEPLOYMENT_TEMPLATE`

const TerminationReasonCodeAzureOperationNotAllowedException TerminationReasonCodePb = `AZURE_OPERATION_NOT_ALLOWED_EXCEPTION`

const TerminationReasonCodeAzureQuotaExceededException TerminationReasonCodePb = `AZURE_QUOTA_EXCEEDED_EXCEPTION`

const TerminationReasonCodeAzureResourceManagerThrottling TerminationReasonCodePb = `AZURE_RESOURCE_MANAGER_THROTTLING`

const TerminationReasonCodeAzureResourceProviderThrottling TerminationReasonCodePb = `AZURE_RESOURCE_PROVIDER_THROTTLING`

const TerminationReasonCodeAzureUnexpectedDeploymentTemplateFailure TerminationReasonCodePb = `AZURE_UNEXPECTED_DEPLOYMENT_TEMPLATE_FAILURE`

const TerminationReasonCodeAzureVmExtensionFailure TerminationReasonCodePb = `AZURE_VM_EXTENSION_FAILURE`

const TerminationReasonCodeAzureVnetConfigurationFailure TerminationReasonCodePb = `AZURE_VNET_CONFIGURATION_FAILURE`

const TerminationReasonCodeBootstrapTimeout TerminationReasonCodePb = `BOOTSTRAP_TIMEOUT`

const TerminationReasonCodeBootstrapTimeoutCloudProviderException TerminationReasonCodePb = `BOOTSTRAP_TIMEOUT_CLOUD_PROVIDER_EXCEPTION`

const TerminationReasonCodeCloudProviderDiskSetupFailure TerminationReasonCodePb = `CLOUD_PROVIDER_DISK_SETUP_FAILURE`

const TerminationReasonCodeCloudProviderLaunchFailure TerminationReasonCodePb = `CLOUD_PROVIDER_LAUNCH_FAILURE`

const TerminationReasonCodeCloudProviderResourceStockout TerminationReasonCodePb = `CLOUD_PROVIDER_RESOURCE_STOCKOUT`

const TerminationReasonCodeCloudProviderShutdown TerminationReasonCodePb = `CLOUD_PROVIDER_SHUTDOWN`

const TerminationReasonCodeCommunicationLost TerminationReasonCodePb = `COMMUNICATION_LOST`

const TerminationReasonCodeContainerLaunchFailure TerminationReasonCodePb = `CONTAINER_LAUNCH_FAILURE`

const TerminationReasonCodeControlPlaneRequestFailure TerminationReasonCodePb = `CONTROL_PLANE_REQUEST_FAILURE`

const TerminationReasonCodeDatabaseConnectionFailure TerminationReasonCodePb = `DATABASE_CONNECTION_FAILURE`

const TerminationReasonCodeDbfsComponentUnhealthy TerminationReasonCodePb = `DBFS_COMPONENT_UNHEALTHY`

const TerminationReasonCodeDockerImagePullFailure TerminationReasonCodePb = `DOCKER_IMAGE_PULL_FAILURE`

const TerminationReasonCodeDriverUnreachable TerminationReasonCodePb = `DRIVER_UNREACHABLE`

const TerminationReasonCodeDriverUnresponsive TerminationReasonCodePb = `DRIVER_UNRESPONSIVE`

const TerminationReasonCodeExecutionComponentUnhealthy TerminationReasonCodePb = `EXECUTION_COMPONENT_UNHEALTHY`

const TerminationReasonCodeGcpQuotaExceeded TerminationReasonCodePb = `GCP_QUOTA_EXCEEDED`

const TerminationReasonCodeGcpServiceAccountDeleted TerminationReasonCodePb = `GCP_SERVICE_ACCOUNT_DELETED`

const TerminationReasonCodeGlobalInitScriptFailure TerminationReasonCodePb = `GLOBAL_INIT_SCRIPT_FAILURE`

const TerminationReasonCodeHiveMetastoreProvisioningFailure TerminationReasonCodePb = `HIVE_METASTORE_PROVISIONING_FAILURE`

const TerminationReasonCodeImagePullPermissionDenied TerminationReasonCodePb = `IMAGE_PULL_PERMISSION_DENIED`

const TerminationReasonCodeInactivity TerminationReasonCodePb = `INACTIVITY`

const TerminationReasonCodeInitScriptFailure TerminationReasonCodePb = `INIT_SCRIPT_FAILURE`

const TerminationReasonCodeInstancePoolClusterFailure TerminationReasonCodePb = `INSTANCE_POOL_CLUSTER_FAILURE`

const TerminationReasonCodeInstanceUnreachable TerminationReasonCodePb = `INSTANCE_UNREACHABLE`

const TerminationReasonCodeInternalError TerminationReasonCodePb = `INTERNAL_ERROR`

const TerminationReasonCodeInvalidArgument TerminationReasonCodePb = `INVALID_ARGUMENT`

const TerminationReasonCodeInvalidSparkImage TerminationReasonCodePb = `INVALID_SPARK_IMAGE`

const TerminationReasonCodeIpExhaustionFailure TerminationReasonCodePb = `IP_EXHAUSTION_FAILURE`

const TerminationReasonCodeJobFinished TerminationReasonCodePb = `JOB_FINISHED`

const TerminationReasonCodeK8sAutoscalingFailure TerminationReasonCodePb = `K8S_AUTOSCALING_FAILURE`

const TerminationReasonCodeK8sDbrClusterLaunchTimeout TerminationReasonCodePb = `K8S_DBR_CLUSTER_LAUNCH_TIMEOUT`

const TerminationReasonCodeMetastoreComponentUnhealthy TerminationReasonCodePb = `METASTORE_COMPONENT_UNHEALTHY`

const TerminationReasonCodeNephosResourceManagement TerminationReasonCodePb = `NEPHOS_RESOURCE_MANAGEMENT`

const TerminationReasonCodeNetworkConfigurationFailure TerminationReasonCodePb = `NETWORK_CONFIGURATION_FAILURE`

const TerminationReasonCodeNfsMountFailure TerminationReasonCodePb = `NFS_MOUNT_FAILURE`

const TerminationReasonCodeNpipTunnelSetupFailure TerminationReasonCodePb = `NPIP_TUNNEL_SETUP_FAILURE`

const TerminationReasonCodeNpipTunnelTokenFailure TerminationReasonCodePb = `NPIP_TUNNEL_TOKEN_FAILURE`

const TerminationReasonCodeRequestRejected TerminationReasonCodePb = `REQUEST_REJECTED`

const TerminationReasonCodeRequestThrottled TerminationReasonCodePb = `REQUEST_THROTTLED`

const TerminationReasonCodeSecretResolutionError TerminationReasonCodePb = `SECRET_RESOLUTION_ERROR`

const TerminationReasonCodeSecurityDaemonRegistrationException TerminationReasonCodePb = `SECURITY_DAEMON_REGISTRATION_EXCEPTION`

const TerminationReasonCodeSelfBootstrapFailure TerminationReasonCodePb = `SELF_BOOTSTRAP_FAILURE`

const TerminationReasonCodeSkippedSlowNodes TerminationReasonCodePb = `SKIPPED_SLOW_NODES`

const TerminationReasonCodeSlowImageDownload TerminationReasonCodePb = `SLOW_IMAGE_DOWNLOAD`

const TerminationReasonCodeSparkError TerminationReasonCodePb = `SPARK_ERROR`

const TerminationReasonCodeSparkImageDownloadFailure TerminationReasonCodePb = `SPARK_IMAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeSparkStartupFailure TerminationReasonCodePb = `SPARK_STARTUP_FAILURE`

const TerminationReasonCodeSpotInstanceTermination TerminationReasonCodePb = `SPOT_INSTANCE_TERMINATION`

const TerminationReasonCodeStorageDownloadFailure TerminationReasonCodePb = `STORAGE_DOWNLOAD_FAILURE`

const TerminationReasonCodeStsClientSetupFailure TerminationReasonCodePb = `STS_CLIENT_SETUP_FAILURE`

const TerminationReasonCodeSubnetExhaustedFailure TerminationReasonCodePb = `SUBNET_EXHAUSTED_FAILURE`

const TerminationReasonCodeTemporarilyUnavailable TerminationReasonCodePb = `TEMPORARILY_UNAVAILABLE`

const TerminationReasonCodeTrialExpired TerminationReasonCodePb = `TRIAL_EXPIRED`

const TerminationReasonCodeUnexpectedLaunchFailure TerminationReasonCodePb = `UNEXPECTED_LAUNCH_FAILURE`

const TerminationReasonCodeUnknown TerminationReasonCodePb = `UNKNOWN`

const TerminationReasonCodeUnsupportedInstanceType TerminationReasonCodePb = `UNSUPPORTED_INSTANCE_TYPE`

const TerminationReasonCodeUpdateInstanceProfileFailure TerminationReasonCodePb = `UPDATE_INSTANCE_PROFILE_FAILURE`

const TerminationReasonCodeUserRequest TerminationReasonCodePb = `USER_REQUEST`

const TerminationReasonCodeWorkerSetupFailure TerminationReasonCodePb = `WORKER_SETUP_FAILURE`

const TerminationReasonCodeWorkspaceCancelledError TerminationReasonCodePb = `WORKSPACE_CANCELLED_ERROR`

const TerminationReasonCodeWorkspaceConfigurationError TerminationReasonCodePb = `WORKSPACE_CONFIGURATION_ERROR`

type TerminationReasonTypePb string

const TerminationReasonTypeClientError TerminationReasonTypePb = `CLIENT_ERROR`

const TerminationReasonTypeCloudFailure TerminationReasonTypePb = `CLOUD_FAILURE`

const TerminationReasonTypeServiceFault TerminationReasonTypePb = `SERVICE_FAULT`

const TerminationReasonTypeSuccess TerminationReasonTypePb = `SUCCESS`

type TextValuePb struct {
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TextValuePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TextValuePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TimeRangePb struct {
	EndTimeMs   int64 `json:"end_time_ms,omitempty" url:"end_time_ms,omitempty"`
	StartTimeMs int64 `json:"start_time_ms,omitempty" url:"start_time_ms,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TimeRangePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TimeRangePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TransferOwnershipObjectIdPb struct {
	NewOwner string `json:"new_owner,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TransferOwnershipObjectIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TransferOwnershipObjectIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TransferOwnershipRequestPb struct {
	NewOwner   string                      `json:"new_owner,omitempty"`
	ObjectId   TransferOwnershipObjectIdPb `json:"-" url:"-"`
	ObjectType OwnableObjectTypePb         `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TransferOwnershipRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TransferOwnershipRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TrashAlertRequestPb struct {
	Id string `json:"-" url:"-"`
}

type TrashAlertV2RequestPb struct {
	Id string `json:"-" url:"-"`
}

type TrashQueryRequestPb struct {
	Id string `json:"-" url:"-"`
}

type UpdateAlertRequestPb struct {
	Alert                  *UpdateAlertRequestAlertPb `json:"alert,omitempty"`
	AutoResolveDisplayName bool                       `json:"auto_resolve_display_name,omitempty"`
	Id                     string                     `json:"-" url:"-"`
	UpdateMask             string                     `json:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateAlertRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateAlertRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateAlertRequestAlertPb struct {
	Condition          *AlertConditionPb `json:"condition,omitempty"`
	CustomBody         string            `json:"custom_body,omitempty"`
	CustomSubject      string            `json:"custom_subject,omitempty"`
	DisplayName        string            `json:"display_name,omitempty"`
	NotifyOnOk         bool              `json:"notify_on_ok,omitempty"`
	OwnerUserName      string            `json:"owner_user_name,omitempty"`
	QueryId            string            `json:"query_id,omitempty"`
	SecondsToRetrigger int               `json:"seconds_to_retrigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateAlertRequestAlertPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateAlertRequestAlertPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateAlertV2RequestPb struct {
	Alert      AlertV2Pb `json:"alert"`
	Id         string    `json:"-" url:"-"`
	UpdateMask string    `json:"-" url:"update_mask"`
}

type UpdateQueryRequestPb struct {
	AutoResolveDisplayName bool                       `json:"auto_resolve_display_name,omitempty"`
	Id                     string                     `json:"-" url:"-"`
	Query                  *UpdateQueryRequestQueryPb `json:"query,omitempty"`
	UpdateMask             string                     `json:"update_mask"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateQueryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateQueryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateQueryRequestQueryPb struct {
	ApplyAutoLimit bool               `json:"apply_auto_limit,omitempty"`
	Catalog        string             `json:"catalog,omitempty"`
	Description    string             `json:"description,omitempty"`
	DisplayName    string             `json:"display_name,omitempty"`
	OwnerUserName  string             `json:"owner_user_name,omitempty"`
	Parameters     []QueryParameterPb `json:"parameters,omitempty"`
	QueryText      string             `json:"query_text,omitempty"`
	RunAsMode      RunAsModePb        `json:"run_as_mode,omitempty"`
	Schema         string             `json:"schema,omitempty"`
	Tags           []string           `json:"tags,omitempty"`
	WarehouseId    string             `json:"warehouse_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateQueryRequestQueryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateQueryRequestQueryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateResponsePb struct {
}

type UpdateVisualizationRequestPb struct {
	Id            string                                     `json:"-" url:"-"`
	UpdateMask    string                                     `json:"update_mask"`
	Visualization *UpdateVisualizationRequestVisualizationPb `json:"visualization,omitempty"`
}

type UpdateVisualizationRequestVisualizationPb struct {
	DisplayName         string `json:"display_name,omitempty"`
	SerializedOptions   string `json:"serialized_options,omitempty"`
	SerializedQueryPlan string `json:"serialized_query_plan,omitempty"`
	Type                string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateVisualizationRequestVisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateVisualizationRequestVisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateWidgetRequestPb struct {
	DashboardId     string          `json:"dashboard_id"`
	Id              string          `json:"-" url:"-"`
	Options         WidgetOptionsPb `json:"options"`
	Text            string          `json:"text,omitempty"`
	VisualizationId string          `json:"visualization_id,omitempty"`
	Width           int             `json:"width"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateWidgetRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateWidgetRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UserPb struct {
	Email string `json:"email,omitempty"`
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UserPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UserPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type VisualizationPb struct {
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

func (st *VisualizationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st VisualizationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehouseAccessControlRequestPb struct {
	GroupName            string                     `json:"group_name,omitempty"`
	PermissionLevel      WarehousePermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                     `json:"service_principal_name,omitempty"`
	UserName             string                     `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WarehouseAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WarehouseAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehouseAccessControlResponsePb struct {
	AllPermissions       []WarehousePermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                  `json:"display_name,omitempty"`
	GroupName            string                  `json:"group_name,omitempty"`
	ServicePrincipalName string                  `json:"service_principal_name,omitempty"`
	UserName             string                  `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WarehouseAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WarehouseAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionPb struct {
	Inherited           bool                       `json:"inherited,omitempty"`
	InheritedFromObject []string                   `json:"inherited_from_object,omitempty"`
	PermissionLevel     WarehousePermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WarehousePermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WarehousePermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionLevelPb string

const WarehousePermissionLevelCanManage WarehousePermissionLevelPb = `CAN_MANAGE`

const WarehousePermissionLevelCanMonitor WarehousePermissionLevelPb = `CAN_MONITOR`

const WarehousePermissionLevelCanUse WarehousePermissionLevelPb = `CAN_USE`

const WarehousePermissionLevelCanView WarehousePermissionLevelPb = `CAN_VIEW`

const WarehousePermissionLevelIsOwner WarehousePermissionLevelPb = `IS_OWNER`

type WarehousePermissionsPb struct {
	AccessControlList []WarehouseAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                             `json:"object_id,omitempty"`
	ObjectType        string                             `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WarehousePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WarehousePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionsDescriptionPb struct {
	Description     string                     `json:"description,omitempty"`
	PermissionLevel WarehousePermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WarehousePermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WarehousePermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehousePermissionsRequestPb struct {
	AccessControlList []WarehouseAccessControlRequestPb `json:"access_control_list,omitempty"`
	WarehouseId       string                            `json:"-" url:"-"`
}

type WarehouseTypePairPb struct {
	Enabled       bool                             `json:"enabled,omitempty"`
	WarehouseType WarehouseTypePairWarehouseTypePb `json:"warehouse_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WarehouseTypePairPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WarehouseTypePairPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WarehouseTypePairWarehouseTypePb string

const WarehouseTypePairWarehouseTypeClassic WarehouseTypePairWarehouseTypePb = `CLASSIC`

const WarehouseTypePairWarehouseTypePro WarehouseTypePairWarehouseTypePb = `PRO`

const WarehouseTypePairWarehouseTypeTypeUnspecified WarehouseTypePairWarehouseTypePb = `TYPE_UNSPECIFIED`

type WidgetPb struct {
	Id            string                 `json:"id,omitempty"`
	Options       *WidgetOptionsPb       `json:"options,omitempty"`
	Visualization *LegacyVisualizationPb `json:"visualization,omitempty"`
	Width         int                    `json:"width,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WidgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WidgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WidgetOptionsPb struct {
	CreatedAt         string            `json:"created_at,omitempty"`
	Description       string            `json:"description,omitempty"`
	IsHidden          bool              `json:"isHidden,omitempty"`
	ParameterMappings any               `json:"parameterMappings,omitempty"`
	Position          *WidgetPositionPb `json:"position,omitempty"`
	Title             string            `json:"title,omitempty"`
	UpdatedAt         string            `json:"updated_at,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WidgetOptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WidgetOptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type WidgetPositionPb struct {
	AutoHeight bool `json:"autoHeight,omitempty"`
	Col        int  `json:"col,omitempty"`
	Row        int  `json:"row,omitempty"`
	SizeX      int  `json:"sizeX,omitempty"`
	SizeY      int  `json:"sizeY,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WidgetPositionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WidgetPositionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
