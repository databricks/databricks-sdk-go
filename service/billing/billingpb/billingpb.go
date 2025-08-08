// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billingpb

import (
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute/computepb"
)

type ActionConfigurationPb struct {
	ActionConfigurationId string                    `json:"action_configuration_id,omitempty"`
	ActionType            ActionConfigurationTypePb `json:"action_type,omitempty"`
	Target                string                    `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ActionConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ActionConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ActionConfigurationTypePb string

const ActionConfigurationTypeEmailNotification ActionConfigurationTypePb = `EMAIL_NOTIFICATION`

type AlertConfigurationPb struct {
	ActionConfigurations []ActionConfigurationPb          `json:"action_configurations,omitempty"`
	AlertConfigurationId string                           `json:"alert_configuration_id,omitempty"`
	QuantityThreshold    string                           `json:"quantity_threshold,omitempty"`
	QuantityType         AlertConfigurationQuantityTypePb `json:"quantity_type,omitempty"`
	TimePeriod           AlertConfigurationTimePeriodPb   `json:"time_period,omitempty"`
	TriggerType          AlertConfigurationTriggerTypePb  `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AlertConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AlertConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AlertConfigurationQuantityTypePb string

const AlertConfigurationQuantityTypeListPriceDollarsUsd AlertConfigurationQuantityTypePb = `LIST_PRICE_DOLLARS_USD`

type AlertConfigurationTimePeriodPb string

const AlertConfigurationTimePeriodMonth AlertConfigurationTimePeriodPb = `MONTH`

type AlertConfigurationTriggerTypePb string

const AlertConfigurationTriggerTypeCumulativeSpendingExceeded AlertConfigurationTriggerTypePb = `CUMULATIVE_SPENDING_EXCEEDED`

type BudgetConfigurationPb struct {
	AccountId             string                       `json:"account_id,omitempty"`
	AlertConfigurations   []AlertConfigurationPb       `json:"alert_configurations,omitempty"`
	BudgetConfigurationId string                       `json:"budget_configuration_id,omitempty"`
	CreateTime            int64                        `json:"create_time,omitempty"`
	DisplayName           string                       `json:"display_name,omitempty"`
	Filter                *BudgetConfigurationFilterPb `json:"filter,omitempty"`
	UpdateTime            int64                        `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BudgetConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BudgetConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BudgetConfigurationFilterPb struct {
	Tags        []BudgetConfigurationFilterTagClausePb        `json:"tags,omitempty"`
	WorkspaceId *BudgetConfigurationFilterWorkspaceIdClausePb `json:"workspace_id,omitempty"`
}

type BudgetConfigurationFilterClausePb struct {
	Operator BudgetConfigurationFilterOperatorPb `json:"operator,omitempty"`
	Values   []string                            `json:"values,omitempty"`
}

type BudgetConfigurationFilterOperatorPb string

const BudgetConfigurationFilterOperatorIn BudgetConfigurationFilterOperatorPb = `IN`

type BudgetConfigurationFilterTagClausePb struct {
	Key   string                             `json:"key,omitempty"`
	Value *BudgetConfigurationFilterClausePb `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BudgetConfigurationFilterTagClausePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BudgetConfigurationFilterTagClausePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BudgetConfigurationFilterWorkspaceIdClausePb struct {
	Operator BudgetConfigurationFilterOperatorPb `json:"operator,omitempty"`
	Values   []int64                             `json:"values,omitempty"`
}

type BudgetPolicyPb struct {
	BindingWorkspaceIds []int64                       `json:"binding_workspace_ids,omitempty"`
	CustomTags          []computepb.CustomPolicyTagPb `json:"custom_tags,omitempty"`
	PolicyId            string                        `json:"policy_id,omitempty"`
	PolicyName          string                        `json:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BudgetPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BudgetPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateBillingUsageDashboardRequestPb struct {
	DashboardType UsageDashboardTypePb `json:"dashboard_type,omitempty"`
	WorkspaceId   int64                `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateBillingUsageDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateBillingUsageDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateBillingUsageDashboardResponsePb struct {
	DashboardId string `json:"dashboard_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateBillingUsageDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateBillingUsageDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateBudgetConfigurationBudgetPb struct {
	AccountId           string                                                 `json:"account_id,omitempty"`
	AlertConfigurations []CreateBudgetConfigurationBudgetAlertConfigurationsPb `json:"alert_configurations,omitempty"`
	DisplayName         string                                                 `json:"display_name,omitempty"`
	Filter              *BudgetConfigurationFilterPb                           `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateBudgetConfigurationBudgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateBudgetConfigurationBudgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateBudgetConfigurationBudgetActionConfigurationsPb struct {
	ActionType ActionConfigurationTypePb `json:"action_type,omitempty"`
	Target     string                    `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateBudgetConfigurationBudgetActionConfigurationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateBudgetConfigurationBudgetActionConfigurationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateBudgetConfigurationBudgetAlertConfigurationsPb struct {
	ActionConfigurations []CreateBudgetConfigurationBudgetActionConfigurationsPb `json:"action_configurations,omitempty"`
	QuantityThreshold    string                                                  `json:"quantity_threshold,omitempty"`
	QuantityType         AlertConfigurationQuantityTypePb                        `json:"quantity_type,omitempty"`
	TimePeriod           AlertConfigurationTimePeriodPb                          `json:"time_period,omitempty"`
	TriggerType          AlertConfigurationTriggerTypePb                         `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateBudgetConfigurationBudgetAlertConfigurationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateBudgetConfigurationBudgetAlertConfigurationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateBudgetConfigurationRequestPb struct {
	Budget CreateBudgetConfigurationBudgetPb `json:"budget"`
}

type CreateBudgetConfigurationResponsePb struct {
	Budget *BudgetConfigurationPb `json:"budget,omitempty"`
}

type CreateBudgetPolicyRequestPb struct {
	Policy    *BudgetPolicyPb `json:"policy,omitempty"`
	RequestId string          `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateBudgetPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateBudgetPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateLogDeliveryConfigurationParamsPb struct {
	ConfigName             string                    `json:"config_name,omitempty"`
	CredentialsId          string                    `json:"credentials_id"`
	DeliveryPathPrefix     string                    `json:"delivery_path_prefix,omitempty"`
	DeliveryStartTime      string                    `json:"delivery_start_time,omitempty"`
	LogType                LogTypePb                 `json:"log_type"`
	OutputFormat           OutputFormatPb            `json:"output_format"`
	Status                 LogDeliveryConfigStatusPb `json:"status,omitempty"`
	StorageConfigurationId string                    `json:"storage_configuration_id"`
	WorkspaceIdsFilter     []int64                   `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateLogDeliveryConfigurationParamsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateLogDeliveryConfigurationParamsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteBudgetConfigurationRequestPb struct {
	BudgetId string `json:"-" url:"-"`
}

type DeleteBudgetConfigurationResponsePb struct {
}

type DeleteBudgetPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

type DeliveryStatusPb string

const DeliveryStatusCreated DeliveryStatusPb = `CREATED`

const DeliveryStatusNotFound DeliveryStatusPb = `NOT_FOUND`

const DeliveryStatusSucceeded DeliveryStatusPb = `SUCCEEDED`

const DeliveryStatusSystemFailure DeliveryStatusPb = `SYSTEM_FAILURE`

const DeliveryStatusUserFailure DeliveryStatusPb = `USER_FAILURE`

type DownloadRequestPb struct {
	EndMonth     string `json:"-" url:"end_month"`
	PersonalData bool   `json:"-" url:"personal_data,omitempty"`
	StartMonth   string `json:"-" url:"start_month"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DownloadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DownloadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DownloadResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

type FilterPb struct {
	CreatorUserId   int64  `json:"creator_user_id,omitempty" url:"creator_user_id,omitempty"`
	CreatorUserName string `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	PolicyName      string `json:"policy_name,omitempty" url:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FilterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FilterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetBillingUsageDashboardRequestPb struct {
	DashboardType UsageDashboardTypePb `json:"-" url:"dashboard_type,omitempty"`
	WorkspaceId   int64                `json:"-" url:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetBillingUsageDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetBillingUsageDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetBillingUsageDashboardResponsePb struct {
	DashboardId  string `json:"dashboard_id,omitempty"`
	DashboardUrl string `json:"dashboard_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetBillingUsageDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetBillingUsageDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetBudgetConfigurationRequestPb struct {
	BudgetId string `json:"-" url:"-"`
}

type GetBudgetConfigurationResponsePb struct {
	Budget *BudgetConfigurationPb `json:"budget,omitempty"`
}

type GetBudgetPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

type GetLogDeliveryConfigurationResponsePb struct {
	LogDeliveryConfiguration *LogDeliveryConfigurationPb `json:"log_delivery_configuration,omitempty"`
}

type GetLogDeliveryRequestPb struct {
	LogDeliveryConfigurationId string `json:"-" url:"-"`
}

type LimitConfigPb struct {
}

type ListBudgetConfigurationsRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListBudgetConfigurationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListBudgetConfigurationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListBudgetConfigurationsResponsePb struct {
	Budgets       []BudgetConfigurationPb `json:"budgets,omitempty"`
	NextPageToken string                  `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListBudgetConfigurationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListBudgetConfigurationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListBudgetPoliciesRequestPb struct {
	FilterBy  *FilterPb   `json:"-" url:"filter_by,omitempty"`
	PageSize  int         `json:"-" url:"page_size,omitempty"`
	PageToken string      `json:"-" url:"page_token,omitempty"`
	SortSpec  *SortSpecPb `json:"-" url:"sort_spec,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListBudgetPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListBudgetPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListBudgetPoliciesResponsePb struct {
	NextPageToken     string           `json:"next_page_token,omitempty"`
	Policies          []BudgetPolicyPb `json:"policies,omitempty"`
	PreviousPageToken string           `json:"previous_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListBudgetPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListBudgetPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListLogDeliveryRequestPb struct {
	CredentialsId          string                    `json:"-" url:"credentials_id,omitempty"`
	PageToken              string                    `json:"-" url:"page_token,omitempty"`
	Status                 LogDeliveryConfigStatusPb `json:"-" url:"status,omitempty"`
	StorageConfigurationId string                    `json:"-" url:"storage_configuration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListLogDeliveryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListLogDeliveryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogDeliveryConfigStatusPb string

const LogDeliveryConfigStatusDisabled LogDeliveryConfigStatusPb = `DISABLED`

const LogDeliveryConfigStatusEnabled LogDeliveryConfigStatusPb = `ENABLED`

type LogDeliveryConfigurationPb struct {
	AccountId              string                    `json:"account_id"`
	ConfigId               string                    `json:"config_id,omitempty"`
	ConfigName             string                    `json:"config_name,omitempty"`
	CreationTime           int64                     `json:"creation_time,omitempty"`
	CredentialsId          string                    `json:"credentials_id"`
	DeliveryPathPrefix     string                    `json:"delivery_path_prefix,omitempty"`
	DeliveryStartTime      string                    `json:"delivery_start_time,omitempty"`
	LogDeliveryStatus      *LogDeliveryStatusPb      `json:"log_delivery_status,omitempty"`
	LogType                LogTypePb                 `json:"log_type"`
	OutputFormat           OutputFormatPb            `json:"output_format"`
	Status                 LogDeliveryConfigStatusPb `json:"status,omitempty"`
	StorageConfigurationId string                    `json:"storage_configuration_id"`
	UpdateTime             int64                     `json:"update_time,omitempty"`
	WorkspaceIdsFilter     []int64                   `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogDeliveryConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogDeliveryConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogDeliveryStatusPb struct {
	LastAttemptTime           string           `json:"last_attempt_time,omitempty"`
	LastSuccessfulAttemptTime string           `json:"last_successful_attempt_time,omitempty"`
	Message                   string           `json:"message"`
	Status                    DeliveryStatusPb `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LogDeliveryStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LogDeliveryStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogTypePb string

const LogTypeAuditLogs LogTypePb = `AUDIT_LOGS`

const LogTypeBillableUsage LogTypePb = `BILLABLE_USAGE`

type OutputFormatPb string

const OutputFormatCsv OutputFormatPb = `CSV`

const OutputFormatJson OutputFormatPb = `JSON`

type PatchStatusResponsePb struct {
}

type SortSpecPb struct {
	Descending bool            `json:"descending,omitempty" url:"descending,omitempty"`
	Field      SortSpecFieldPb `json:"field,omitempty" url:"field,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SortSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SortSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SortSpecFieldPb string

const SortSpecFieldPolicyName SortSpecFieldPb = `POLICY_NAME`

type UpdateBudgetConfigurationBudgetPb struct {
	AccountId             string                       `json:"account_id,omitempty"`
	AlertConfigurations   []AlertConfigurationPb       `json:"alert_configurations,omitempty"`
	BudgetConfigurationId string                       `json:"budget_configuration_id,omitempty"`
	DisplayName           string                       `json:"display_name,omitempty"`
	Filter                *BudgetConfigurationFilterPb `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateBudgetConfigurationBudgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateBudgetConfigurationBudgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateBudgetConfigurationRequestPb struct {
	Budget   UpdateBudgetConfigurationBudgetPb `json:"budget"`
	BudgetId string                            `json:"-" url:"-"`
}

type UpdateBudgetConfigurationResponsePb struct {
	Budget *BudgetConfigurationPb `json:"budget,omitempty"`
}

type UpdateBudgetPolicyRequestPb struct {
	LimitConfig *LimitConfigPb `json:"-" url:"limit_config,omitempty"`
	Policy      BudgetPolicyPb `json:"policy"`
	PolicyId    string         `json:"-" url:"-"`
}

type UpdateLogDeliveryConfigurationStatusRequestPb struct {
	LogDeliveryConfigurationId string                    `json:"-" url:"-"`
	Status                     LogDeliveryConfigStatusPb `json:"status"`
}

type UsageDashboardTypePb string

const UsageDashboardTypeUsageDashboardTypeGlobal UsageDashboardTypePb = `USAGE_DASHBOARD_TYPE_GLOBAL`

const UsageDashboardTypeUsageDashboardTypeWorkspace UsageDashboardTypePb = `USAGE_DASHBOARD_TYPE_WORKSPACE`

type WrappedCreateLogDeliveryConfigurationPb struct {
	LogDeliveryConfiguration CreateLogDeliveryConfigurationParamsPb `json:"log_delivery_configuration"`
}

type WrappedLogDeliveryConfigurationPb struct {
	LogDeliveryConfiguration *LogDeliveryConfigurationPb `json:"log_delivery_configuration,omitempty"`
}

type WrappedLogDeliveryConfigurationsPb struct {
	LogDeliveryConfigurations []LogDeliveryConfigurationPb `json:"log_delivery_configurations,omitempty"`
	NextPageToken             string                       `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WrappedLogDeliveryConfigurationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WrappedLogDeliveryConfigurationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
