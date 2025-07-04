// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func actionConfigurationToPb(st *ActionConfiguration) (*actionConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &actionConfigurationPb{}
	pb.ActionConfigurationId = st.ActionConfigurationId
	pb.ActionType = st.ActionType
	pb.Target = st.Target

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type actionConfigurationPb struct {
	ActionConfigurationId string                  `json:"action_configuration_id,omitempty"`
	ActionType            ActionConfigurationType `json:"action_type,omitempty"`
	Target                string                  `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func actionConfigurationFromPb(pb *actionConfigurationPb) (*ActionConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ActionConfiguration{}
	st.ActionConfigurationId = pb.ActionConfigurationId
	st.ActionType = pb.ActionType
	st.Target = pb.Target

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *actionConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st actionConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func alertConfigurationToPb(st *AlertConfiguration) (*alertConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConfigurationPb{}
	pb.ActionConfigurations = st.ActionConfigurations
	pb.AlertConfigurationId = st.AlertConfigurationId
	pb.QuantityThreshold = st.QuantityThreshold
	pb.QuantityType = st.QuantityType
	pb.TimePeriod = st.TimePeriod
	pb.TriggerType = st.TriggerType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type alertConfigurationPb struct {
	ActionConfigurations []ActionConfiguration          `json:"action_configurations,omitempty"`
	AlertConfigurationId string                         `json:"alert_configuration_id,omitempty"`
	QuantityThreshold    string                         `json:"quantity_threshold,omitempty"`
	QuantityType         AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
	TimePeriod           AlertConfigurationTimePeriod   `json:"time_period,omitempty"`
	TriggerType          AlertConfigurationTriggerType  `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertConfigurationFromPb(pb *alertConfigurationPb) (*AlertConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConfiguration{}
	st.ActionConfigurations = pb.ActionConfigurations
	st.AlertConfigurationId = pb.AlertConfigurationId
	st.QuantityThreshold = pb.QuantityThreshold
	st.QuantityType = pb.QuantityType
	st.TimePeriod = pb.TimePeriod
	st.TriggerType = pb.TriggerType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *alertConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st alertConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func budgetConfigurationToPb(st *BudgetConfiguration) (*budgetConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationPb{}
	pb.AccountId = st.AccountId
	pb.AlertConfigurations = st.AlertConfigurations
	pb.BudgetConfigurationId = st.BudgetConfigurationId
	pb.CreateTime = st.CreateTime
	pb.DisplayName = st.DisplayName
	pb.Filter = st.Filter
	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type budgetConfigurationPb struct {
	AccountId             string                     `json:"account_id,omitempty"`
	AlertConfigurations   []AlertConfiguration       `json:"alert_configurations,omitempty"`
	BudgetConfigurationId string                     `json:"budget_configuration_id,omitempty"`
	CreateTime            int64                      `json:"create_time,omitempty"`
	DisplayName           string                     `json:"display_name,omitempty"`
	Filter                *BudgetConfigurationFilter `json:"filter,omitempty"`
	UpdateTime            int64                      `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func budgetConfigurationFromPb(pb *budgetConfigurationPb) (*BudgetConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfiguration{}
	st.AccountId = pb.AccountId
	st.AlertConfigurations = pb.AlertConfigurations
	st.BudgetConfigurationId = pb.BudgetConfigurationId
	st.CreateTime = pb.CreateTime
	st.DisplayName = pb.DisplayName
	st.Filter = pb.Filter
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *budgetConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st budgetConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func budgetConfigurationFilterToPb(st *BudgetConfigurationFilter) (*budgetConfigurationFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationFilterPb{}
	pb.Tags = st.Tags
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type budgetConfigurationFilterPb struct {
	Tags        []BudgetConfigurationFilterTagClause        `json:"tags,omitempty"`
	WorkspaceId *BudgetConfigurationFilterWorkspaceIdClause `json:"workspace_id,omitempty"`
}

func budgetConfigurationFilterFromPb(pb *budgetConfigurationFilterPb) (*BudgetConfigurationFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilter{}
	st.Tags = pb.Tags
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func budgetConfigurationFilterClauseToPb(st *BudgetConfigurationFilterClause) (*budgetConfigurationFilterClausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationFilterClausePb{}
	pb.Operator = st.Operator
	pb.Values = st.Values

	return pb, nil
}

type budgetConfigurationFilterClausePb struct {
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`
	Values   []string                          `json:"values,omitempty"`
}

func budgetConfigurationFilterClauseFromPb(pb *budgetConfigurationFilterClausePb) (*BudgetConfigurationFilterClause, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilterClause{}
	st.Operator = pb.Operator
	st.Values = pb.Values

	return st, nil
}

func budgetConfigurationFilterTagClauseToPb(st *BudgetConfigurationFilterTagClause) (*budgetConfigurationFilterTagClausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationFilterTagClausePb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type budgetConfigurationFilterTagClausePb struct {
	Key   string                           `json:"key,omitempty"`
	Value *BudgetConfigurationFilterClause `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func budgetConfigurationFilterTagClauseFromPb(pb *budgetConfigurationFilterTagClausePb) (*BudgetConfigurationFilterTagClause, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilterTagClause{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *budgetConfigurationFilterTagClausePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st budgetConfigurationFilterTagClausePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func budgetConfigurationFilterWorkspaceIdClauseToPb(st *BudgetConfigurationFilterWorkspaceIdClause) (*budgetConfigurationFilterWorkspaceIdClausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationFilterWorkspaceIdClausePb{}
	pb.Operator = st.Operator
	pb.Values = st.Values

	return pb, nil
}

type budgetConfigurationFilterWorkspaceIdClausePb struct {
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`
	Values   []int64                           `json:"values,omitempty"`
}

func budgetConfigurationFilterWorkspaceIdClauseFromPb(pb *budgetConfigurationFilterWorkspaceIdClausePb) (*BudgetConfigurationFilterWorkspaceIdClause, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilterWorkspaceIdClause{}
	st.Operator = pb.Operator
	st.Values = pb.Values

	return st, nil
}

func budgetPolicyToPb(st *BudgetPolicy) (*budgetPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetPolicyPb{}
	pb.BindingWorkspaceIds = st.BindingWorkspaceIds
	pb.CustomTags = st.CustomTags
	pb.PolicyId = st.PolicyId
	pb.PolicyName = st.PolicyName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type budgetPolicyPb struct {
	BindingWorkspaceIds []int64                   `json:"binding_workspace_ids,omitempty"`
	CustomTags          []compute.CustomPolicyTag `json:"custom_tags,omitempty"`
	PolicyId            string                    `json:"policy_id,omitempty"`
	PolicyName          string                    `json:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func budgetPolicyFromPb(pb *budgetPolicyPb) (*BudgetPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetPolicy{}
	st.BindingWorkspaceIds = pb.BindingWorkspaceIds
	st.CustomTags = pb.CustomTags
	st.PolicyId = pb.PolicyId
	st.PolicyName = pb.PolicyName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *budgetPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st budgetPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createBillingUsageDashboardRequestToPb(st *CreateBillingUsageDashboardRequest) (*createBillingUsageDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBillingUsageDashboardRequestPb{}
	pb.DashboardType = st.DashboardType
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createBillingUsageDashboardRequestPb struct {
	DashboardType UsageDashboardType `json:"dashboard_type,omitempty"`
	WorkspaceId   int64              `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBillingUsageDashboardRequestFromPb(pb *createBillingUsageDashboardRequestPb) (*CreateBillingUsageDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBillingUsageDashboardRequest{}
	st.DashboardType = pb.DashboardType
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createBillingUsageDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createBillingUsageDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createBillingUsageDashboardResponseToPb(st *CreateBillingUsageDashboardResponse) (*createBillingUsageDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBillingUsageDashboardResponsePb{}
	pb.DashboardId = st.DashboardId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createBillingUsageDashboardResponsePb struct {
	DashboardId string `json:"dashboard_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBillingUsageDashboardResponseFromPb(pb *createBillingUsageDashboardResponsePb) (*CreateBillingUsageDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBillingUsageDashboardResponse{}
	st.DashboardId = pb.DashboardId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createBillingUsageDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createBillingUsageDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createBudgetConfigurationBudgetToPb(st *CreateBudgetConfigurationBudget) (*createBudgetConfigurationBudgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationBudgetPb{}
	pb.AccountId = st.AccountId
	pb.AlertConfigurations = st.AlertConfigurations
	pb.DisplayName = st.DisplayName
	pb.Filter = st.Filter

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createBudgetConfigurationBudgetPb struct {
	AccountId           string                                               `json:"account_id,omitempty"`
	AlertConfigurations []CreateBudgetConfigurationBudgetAlertConfigurations `json:"alert_configurations,omitempty"`
	DisplayName         string                                               `json:"display_name,omitempty"`
	Filter              *BudgetConfigurationFilter                           `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBudgetConfigurationBudgetFromPb(pb *createBudgetConfigurationBudgetPb) (*CreateBudgetConfigurationBudget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudget{}
	st.AccountId = pb.AccountId
	st.AlertConfigurations = pb.AlertConfigurations
	st.DisplayName = pb.DisplayName
	st.Filter = pb.Filter

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createBudgetConfigurationBudgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createBudgetConfigurationBudgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createBudgetConfigurationBudgetActionConfigurationsToPb(st *CreateBudgetConfigurationBudgetActionConfigurations) (*createBudgetConfigurationBudgetActionConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationBudgetActionConfigurationsPb{}
	pb.ActionType = st.ActionType
	pb.Target = st.Target

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createBudgetConfigurationBudgetActionConfigurationsPb struct {
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	Target     string                  `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBudgetConfigurationBudgetActionConfigurationsFromPb(pb *createBudgetConfigurationBudgetActionConfigurationsPb) (*CreateBudgetConfigurationBudgetActionConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudgetActionConfigurations{}
	st.ActionType = pb.ActionType
	st.Target = pb.Target

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createBudgetConfigurationBudgetActionConfigurationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createBudgetConfigurationBudgetActionConfigurationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createBudgetConfigurationBudgetAlertConfigurationsToPb(st *CreateBudgetConfigurationBudgetAlertConfigurations) (*createBudgetConfigurationBudgetAlertConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationBudgetAlertConfigurationsPb{}
	pb.ActionConfigurations = st.ActionConfigurations
	pb.QuantityThreshold = st.QuantityThreshold
	pb.QuantityType = st.QuantityType
	pb.TimePeriod = st.TimePeriod
	pb.TriggerType = st.TriggerType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createBudgetConfigurationBudgetAlertConfigurationsPb struct {
	ActionConfigurations []CreateBudgetConfigurationBudgetActionConfigurations `json:"action_configurations,omitempty"`
	QuantityThreshold    string                                                `json:"quantity_threshold,omitempty"`
	QuantityType         AlertConfigurationQuantityType                        `json:"quantity_type,omitempty"`
	TimePeriod           AlertConfigurationTimePeriod                          `json:"time_period,omitempty"`
	TriggerType          AlertConfigurationTriggerType                         `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBudgetConfigurationBudgetAlertConfigurationsFromPb(pb *createBudgetConfigurationBudgetAlertConfigurationsPb) (*CreateBudgetConfigurationBudgetAlertConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudgetAlertConfigurations{}
	st.ActionConfigurations = pb.ActionConfigurations
	st.QuantityThreshold = pb.QuantityThreshold
	st.QuantityType = pb.QuantityType
	st.TimePeriod = pb.TimePeriod
	st.TriggerType = pb.TriggerType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createBudgetConfigurationBudgetAlertConfigurationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createBudgetConfigurationBudgetAlertConfigurationsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createBudgetConfigurationRequestToPb(st *CreateBudgetConfigurationRequest) (*createBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationRequestPb{}
	pb.Budget = st.Budget

	return pb, nil
}

type createBudgetConfigurationRequestPb struct {
	Budget CreateBudgetConfigurationBudget `json:"budget"`
}

func createBudgetConfigurationRequestFromPb(pb *createBudgetConfigurationRequestPb) (*CreateBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationRequest{}
	st.Budget = pb.Budget

	return st, nil
}

func createBudgetConfigurationResponseToPb(st *CreateBudgetConfigurationResponse) (*createBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationResponsePb{}
	pb.Budget = st.Budget

	return pb, nil
}

type createBudgetConfigurationResponsePb struct {
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func createBudgetConfigurationResponseFromPb(pb *createBudgetConfigurationResponsePb) (*CreateBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationResponse{}
	st.Budget = pb.Budget

	return st, nil
}

func createBudgetPolicyRequestToPb(st *CreateBudgetPolicyRequest) (*createBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetPolicyRequestPb{}
	pb.Policy = st.Policy
	pb.RequestId = st.RequestId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createBudgetPolicyRequestPb struct {
	Policy    *BudgetPolicy `json:"policy,omitempty"`
	RequestId string        `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBudgetPolicyRequestFromPb(pb *createBudgetPolicyRequestPb) (*CreateBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetPolicyRequest{}
	st.Policy = pb.Policy
	st.RequestId = pb.RequestId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createBudgetPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createBudgetPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createLogDeliveryConfigurationParamsToPb(st *CreateLogDeliveryConfigurationParams) (*createLogDeliveryConfigurationParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createLogDeliveryConfigurationParamsPb{}
	pb.ConfigName = st.ConfigName
	pb.CredentialsId = st.CredentialsId
	pb.DeliveryPathPrefix = st.DeliveryPathPrefix
	pb.DeliveryStartTime = st.DeliveryStartTime
	pb.LogType = st.LogType
	pb.OutputFormat = st.OutputFormat
	pb.Status = st.Status
	pb.StorageConfigurationId = st.StorageConfigurationId
	pb.WorkspaceIdsFilter = st.WorkspaceIdsFilter

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createLogDeliveryConfigurationParamsPb struct {
	ConfigName             string                  `json:"config_name,omitempty"`
	CredentialsId          string                  `json:"credentials_id"`
	DeliveryPathPrefix     string                  `json:"delivery_path_prefix,omitempty"`
	DeliveryStartTime      string                  `json:"delivery_start_time,omitempty"`
	LogType                LogType                 `json:"log_type"`
	OutputFormat           OutputFormat            `json:"output_format"`
	Status                 LogDeliveryConfigStatus `json:"status,omitempty"`
	StorageConfigurationId string                  `json:"storage_configuration_id"`
	WorkspaceIdsFilter     []int64                 `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createLogDeliveryConfigurationParamsFromPb(pb *createLogDeliveryConfigurationParamsPb) (*CreateLogDeliveryConfigurationParams, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateLogDeliveryConfigurationParams{}
	st.ConfigName = pb.ConfigName
	st.CredentialsId = pb.CredentialsId
	st.DeliveryPathPrefix = pb.DeliveryPathPrefix
	st.DeliveryStartTime = pb.DeliveryStartTime
	st.LogType = pb.LogType
	st.OutputFormat = pb.OutputFormat
	st.Status = pb.Status
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.WorkspaceIdsFilter = pb.WorkspaceIdsFilter

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createLogDeliveryConfigurationParamsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createLogDeliveryConfigurationParamsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteBudgetConfigurationRequestToPb(st *DeleteBudgetConfigurationRequest) (*deleteBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteBudgetConfigurationRequestPb{}
	pb.BudgetId = st.BudgetId

	return pb, nil
}

type deleteBudgetConfigurationRequestPb struct {
	BudgetId string `json:"-" url:"-"`
}

func deleteBudgetConfigurationRequestFromPb(pb *deleteBudgetConfigurationRequestPb) (*DeleteBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteBudgetConfigurationRequest{}
	st.BudgetId = pb.BudgetId

	return st, nil
}

func deleteBudgetPolicyRequestToPb(st *DeleteBudgetPolicyRequest) (*deleteBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteBudgetPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

type deleteBudgetPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

func deleteBudgetPolicyRequestFromPb(pb *deleteBudgetPolicyRequestPb) (*DeleteBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteBudgetPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

func downloadRequestToPb(st *DownloadRequest) (*downloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &downloadRequestPb{}
	pb.EndMonth = st.EndMonth
	pb.PersonalData = st.PersonalData
	pb.StartMonth = st.StartMonth

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type downloadRequestPb struct {
	EndMonth     string `json:"-" url:"end_month"`
	PersonalData bool   `json:"-" url:"personal_data,omitempty"`
	StartMonth   string `json:"-" url:"start_month"`

	ForceSendFields []string `json:"-" url:"-"`
}

func downloadRequestFromPb(pb *downloadRequestPb) (*DownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadRequest{}
	st.EndMonth = pb.EndMonth
	st.PersonalData = pb.PersonalData
	st.StartMonth = pb.StartMonth

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *downloadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st downloadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func downloadResponseToPb(st *DownloadResponse) (*downloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &downloadResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

type downloadResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

func downloadResponseFromPb(pb *downloadResponsePb) (*DownloadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadResponse{}
	st.Contents = pb.Contents

	return st, nil
}

func filterToPb(st *Filter) (*filterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filterPb{}
	pb.CreatorUserId = st.CreatorUserId
	pb.CreatorUserName = st.CreatorUserName
	pb.PolicyName = st.PolicyName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type filterPb struct {
	CreatorUserId   int64  `json:"creator_user_id,omitempty" url:"creator_user_id,omitempty"`
	CreatorUserName string `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	PolicyName      string `json:"policy_name,omitempty" url:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func filterFromPb(pb *filterPb) (*Filter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Filter{}
	st.CreatorUserId = pb.CreatorUserId
	st.CreatorUserName = pb.CreatorUserName
	st.PolicyName = pb.PolicyName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *filterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st filterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getBillingUsageDashboardRequestToPb(st *GetBillingUsageDashboardRequest) (*getBillingUsageDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBillingUsageDashboardRequestPb{}
	pb.DashboardType = st.DashboardType
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getBillingUsageDashboardRequestPb struct {
	DashboardType UsageDashboardType `json:"-" url:"dashboard_type,omitempty"`
	WorkspaceId   int64              `json:"-" url:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getBillingUsageDashboardRequestFromPb(pb *getBillingUsageDashboardRequestPb) (*GetBillingUsageDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBillingUsageDashboardRequest{}
	st.DashboardType = pb.DashboardType
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getBillingUsageDashboardRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getBillingUsageDashboardRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getBillingUsageDashboardResponseToPb(st *GetBillingUsageDashboardResponse) (*getBillingUsageDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBillingUsageDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.DashboardUrl = st.DashboardUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getBillingUsageDashboardResponsePb struct {
	DashboardId  string `json:"dashboard_id,omitempty"`
	DashboardUrl string `json:"dashboard_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getBillingUsageDashboardResponseFromPb(pb *getBillingUsageDashboardResponsePb) (*GetBillingUsageDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBillingUsageDashboardResponse{}
	st.DashboardId = pb.DashboardId
	st.DashboardUrl = pb.DashboardUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getBillingUsageDashboardResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getBillingUsageDashboardResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getBudgetConfigurationRequestToPb(st *GetBudgetConfigurationRequest) (*getBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBudgetConfigurationRequestPb{}
	pb.BudgetId = st.BudgetId

	return pb, nil
}

type getBudgetConfigurationRequestPb struct {
	BudgetId string `json:"-" url:"-"`
}

func getBudgetConfigurationRequestFromPb(pb *getBudgetConfigurationRequestPb) (*GetBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBudgetConfigurationRequest{}
	st.BudgetId = pb.BudgetId

	return st, nil
}

func getBudgetConfigurationResponseToPb(st *GetBudgetConfigurationResponse) (*getBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBudgetConfigurationResponsePb{}
	pb.Budget = st.Budget

	return pb, nil
}

type getBudgetConfigurationResponsePb struct {
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func getBudgetConfigurationResponseFromPb(pb *getBudgetConfigurationResponsePb) (*GetBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBudgetConfigurationResponse{}
	st.Budget = pb.Budget

	return st, nil
}

func getBudgetPolicyRequestToPb(st *GetBudgetPolicyRequest) (*getBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBudgetPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

type getBudgetPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

func getBudgetPolicyRequestFromPb(pb *getBudgetPolicyRequestPb) (*GetBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBudgetPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

func getLogDeliveryConfigurationResponseToPb(st *GetLogDeliveryConfigurationResponse) (*getLogDeliveryConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLogDeliveryConfigurationResponsePb{}
	pb.LogDeliveryConfiguration = st.LogDeliveryConfiguration

	return pb, nil
}

type getLogDeliveryConfigurationResponsePb struct {
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

func getLogDeliveryConfigurationResponseFromPb(pb *getLogDeliveryConfigurationResponsePb) (*GetLogDeliveryConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLogDeliveryConfigurationResponse{}
	st.LogDeliveryConfiguration = pb.LogDeliveryConfiguration

	return st, nil
}

func getLogDeliveryRequestToPb(st *GetLogDeliveryRequest) (*getLogDeliveryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLogDeliveryRequestPb{}
	pb.LogDeliveryConfigurationId = st.LogDeliveryConfigurationId

	return pb, nil
}

type getLogDeliveryRequestPb struct {
	LogDeliveryConfigurationId string `json:"-" url:"-"`
}

func getLogDeliveryRequestFromPb(pb *getLogDeliveryRequestPb) (*GetLogDeliveryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLogDeliveryRequest{}
	st.LogDeliveryConfigurationId = pb.LogDeliveryConfigurationId

	return st, nil
}

func limitConfigToPb(st *LimitConfig) (*limitConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &limitConfigPb{}

	return pb, nil
}

type limitConfigPb struct {
}

func limitConfigFromPb(pb *limitConfigPb) (*LimitConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LimitConfig{}

	return st, nil
}

func listBudgetConfigurationsRequestToPb(st *ListBudgetConfigurationsRequest) (*listBudgetConfigurationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listBudgetConfigurationsRequestPb{}
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listBudgetConfigurationsRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listBudgetConfigurationsRequestFromPb(pb *listBudgetConfigurationsRequestPb) (*ListBudgetConfigurationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetConfigurationsRequest{}
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listBudgetConfigurationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listBudgetConfigurationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listBudgetConfigurationsResponseToPb(st *ListBudgetConfigurationsResponse) (*listBudgetConfigurationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listBudgetConfigurationsResponsePb{}
	pb.Budgets = st.Budgets
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listBudgetConfigurationsResponsePb struct {
	Budgets       []BudgetConfiguration `json:"budgets,omitempty"`
	NextPageToken string                `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listBudgetConfigurationsResponseFromPb(pb *listBudgetConfigurationsResponsePb) (*ListBudgetConfigurationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetConfigurationsResponse{}
	st.Budgets = pb.Budgets
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listBudgetConfigurationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listBudgetConfigurationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listBudgetPoliciesRequestToPb(st *ListBudgetPoliciesRequest) (*listBudgetPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listBudgetPoliciesRequestPb{}
	pb.FilterBy = st.FilterBy
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.SortSpec = st.SortSpec

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listBudgetPoliciesRequestPb struct {
	FilterBy  *Filter   `json:"-" url:"filter_by,omitempty"`
	PageSize  int       `json:"-" url:"page_size,omitempty"`
	PageToken string    `json:"-" url:"page_token,omitempty"`
	SortSpec  *SortSpec `json:"-" url:"sort_spec,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listBudgetPoliciesRequestFromPb(pb *listBudgetPoliciesRequestPb) (*ListBudgetPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetPoliciesRequest{}
	st.FilterBy = pb.FilterBy
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.SortSpec = pb.SortSpec

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listBudgetPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listBudgetPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listBudgetPoliciesResponseToPb(st *ListBudgetPoliciesResponse) (*listBudgetPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listBudgetPoliciesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Policies = st.Policies
	pb.PreviousPageToken = st.PreviousPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listBudgetPoliciesResponsePb struct {
	NextPageToken     string         `json:"next_page_token,omitempty"`
	Policies          []BudgetPolicy `json:"policies,omitempty"`
	PreviousPageToken string         `json:"previous_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listBudgetPoliciesResponseFromPb(pb *listBudgetPoliciesResponsePb) (*ListBudgetPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetPoliciesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Policies = pb.Policies
	st.PreviousPageToken = pb.PreviousPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listBudgetPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listBudgetPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listLogDeliveryRequestToPb(st *ListLogDeliveryRequest) (*listLogDeliveryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listLogDeliveryRequestPb{}
	pb.CredentialsId = st.CredentialsId
	pb.PageToken = st.PageToken
	pb.Status = st.Status
	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listLogDeliveryRequestPb struct {
	CredentialsId          string                  `json:"-" url:"credentials_id,omitempty"`
	PageToken              string                  `json:"-" url:"page_token,omitempty"`
	Status                 LogDeliveryConfigStatus `json:"-" url:"status,omitempty"`
	StorageConfigurationId string                  `json:"-" url:"storage_configuration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listLogDeliveryRequestFromPb(pb *listLogDeliveryRequestPb) (*ListLogDeliveryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListLogDeliveryRequest{}
	st.CredentialsId = pb.CredentialsId
	st.PageToken = pb.PageToken
	st.Status = pb.Status
	st.StorageConfigurationId = pb.StorageConfigurationId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listLogDeliveryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listLogDeliveryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logDeliveryConfigurationToPb(st *LogDeliveryConfiguration) (*logDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logDeliveryConfigurationPb{}
	pb.AccountId = st.AccountId
	pb.ConfigId = st.ConfigId
	pb.ConfigName = st.ConfigName
	pb.CreationTime = st.CreationTime
	pb.CredentialsId = st.CredentialsId
	pb.DeliveryPathPrefix = st.DeliveryPathPrefix
	pb.DeliveryStartTime = st.DeliveryStartTime
	pb.LogDeliveryStatus = st.LogDeliveryStatus
	pb.LogType = st.LogType
	pb.OutputFormat = st.OutputFormat
	pb.Status = st.Status
	pb.StorageConfigurationId = st.StorageConfigurationId
	pb.UpdateTime = st.UpdateTime
	pb.WorkspaceIdsFilter = st.WorkspaceIdsFilter

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logDeliveryConfigurationPb struct {
	AccountId              string                  `json:"account_id"`
	ConfigId               string                  `json:"config_id,omitempty"`
	ConfigName             string                  `json:"config_name,omitempty"`
	CreationTime           int64                   `json:"creation_time,omitempty"`
	CredentialsId          string                  `json:"credentials_id"`
	DeliveryPathPrefix     string                  `json:"delivery_path_prefix,omitempty"`
	DeliveryStartTime      string                  `json:"delivery_start_time,omitempty"`
	LogDeliveryStatus      *LogDeliveryStatus      `json:"log_delivery_status,omitempty"`
	LogType                LogType                 `json:"log_type"`
	OutputFormat           OutputFormat            `json:"output_format"`
	Status                 LogDeliveryConfigStatus `json:"status,omitempty"`
	StorageConfigurationId string                  `json:"storage_configuration_id"`
	UpdateTime             int64                   `json:"update_time,omitempty"`
	WorkspaceIdsFilter     []int64                 `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logDeliveryConfigurationFromPb(pb *logDeliveryConfigurationPb) (*LogDeliveryConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogDeliveryConfiguration{}
	st.AccountId = pb.AccountId
	st.ConfigId = pb.ConfigId
	st.ConfigName = pb.ConfigName
	st.CreationTime = pb.CreationTime
	st.CredentialsId = pb.CredentialsId
	st.DeliveryPathPrefix = pb.DeliveryPathPrefix
	st.DeliveryStartTime = pb.DeliveryStartTime
	st.LogDeliveryStatus = pb.LogDeliveryStatus
	st.LogType = pb.LogType
	st.OutputFormat = pb.OutputFormat
	st.Status = pb.Status
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.UpdateTime = pb.UpdateTime
	st.WorkspaceIdsFilter = pb.WorkspaceIdsFilter

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logDeliveryConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logDeliveryConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func logDeliveryStatusToPb(st *LogDeliveryStatus) (*logDeliveryStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logDeliveryStatusPb{}
	pb.LastAttemptTime = st.LastAttemptTime
	pb.LastSuccessfulAttemptTime = st.LastSuccessfulAttemptTime
	pb.Message = st.Message
	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type logDeliveryStatusPb struct {
	LastAttemptTime           string         `json:"last_attempt_time,omitempty"`
	LastSuccessfulAttemptTime string         `json:"last_successful_attempt_time,omitempty"`
	Message                   string         `json:"message"`
	Status                    DeliveryStatus `json:"status"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logDeliveryStatusFromPb(pb *logDeliveryStatusPb) (*LogDeliveryStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogDeliveryStatus{}
	st.LastAttemptTime = pb.LastAttemptTime
	st.LastSuccessfulAttemptTime = pb.LastSuccessfulAttemptTime
	st.Message = pb.Message
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logDeliveryStatusPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logDeliveryStatusPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sortSpecToPb(st *SortSpec) (*sortSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sortSpecPb{}
	pb.Descending = st.Descending
	pb.Field = st.Field

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sortSpecPb struct {
	Descending bool          `json:"descending,omitempty" url:"descending,omitempty"`
	Field      SortSpecField `json:"field,omitempty" url:"field,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sortSpecFromPb(pb *sortSpecPb) (*SortSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SortSpec{}
	st.Descending = pb.Descending
	st.Field = pb.Field

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sortSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sortSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateBudgetConfigurationBudgetToPb(st *UpdateBudgetConfigurationBudget) (*updateBudgetConfigurationBudgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetConfigurationBudgetPb{}
	pb.AccountId = st.AccountId
	pb.AlertConfigurations = st.AlertConfigurations
	pb.BudgetConfigurationId = st.BudgetConfigurationId
	pb.DisplayName = st.DisplayName
	pb.Filter = st.Filter

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateBudgetConfigurationBudgetPb struct {
	AccountId             string                     `json:"account_id,omitempty"`
	AlertConfigurations   []AlertConfiguration       `json:"alert_configurations,omitempty"`
	BudgetConfigurationId string                     `json:"budget_configuration_id,omitempty"`
	DisplayName           string                     `json:"display_name,omitempty"`
	Filter                *BudgetConfigurationFilter `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateBudgetConfigurationBudgetFromPb(pb *updateBudgetConfigurationBudgetPb) (*UpdateBudgetConfigurationBudget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationBudget{}
	st.AccountId = pb.AccountId
	st.AlertConfigurations = pb.AlertConfigurations
	st.BudgetConfigurationId = pb.BudgetConfigurationId
	st.DisplayName = pb.DisplayName
	st.Filter = pb.Filter

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateBudgetConfigurationBudgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateBudgetConfigurationBudgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateBudgetConfigurationRequestToPb(st *UpdateBudgetConfigurationRequest) (*updateBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetConfigurationRequestPb{}
	pb.Budget = st.Budget
	pb.BudgetId = st.BudgetId

	return pb, nil
}

type updateBudgetConfigurationRequestPb struct {
	Budget   UpdateBudgetConfigurationBudget `json:"budget"`
	BudgetId string                          `json:"-" url:"-"`
}

func updateBudgetConfigurationRequestFromPb(pb *updateBudgetConfigurationRequestPb) (*UpdateBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationRequest{}
	st.Budget = pb.Budget
	st.BudgetId = pb.BudgetId

	return st, nil
}

func updateBudgetConfigurationResponseToPb(st *UpdateBudgetConfigurationResponse) (*updateBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetConfigurationResponsePb{}
	pb.Budget = st.Budget

	return pb, nil
}

type updateBudgetConfigurationResponsePb struct {
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func updateBudgetConfigurationResponseFromPb(pb *updateBudgetConfigurationResponsePb) (*UpdateBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationResponse{}
	st.Budget = pb.Budget

	return st, nil
}

func updateBudgetPolicyRequestToPb(st *UpdateBudgetPolicyRequest) (*updateBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetPolicyRequestPb{}
	pb.LimitConfig = st.LimitConfig
	pb.Policy = st.Policy
	pb.PolicyId = st.PolicyId

	return pb, nil
}

type updateBudgetPolicyRequestPb struct {
	LimitConfig *LimitConfig `json:"-" url:"limit_config,omitempty"`
	Policy      BudgetPolicy `json:"policy"`
	PolicyId    string       `json:"-" url:"-"`
}

func updateBudgetPolicyRequestFromPb(pb *updateBudgetPolicyRequestPb) (*UpdateBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetPolicyRequest{}
	st.LimitConfig = pb.LimitConfig
	st.Policy = pb.Policy
	st.PolicyId = pb.PolicyId

	return st, nil
}

func updateLogDeliveryConfigurationStatusRequestToPb(st *UpdateLogDeliveryConfigurationStatusRequest) (*updateLogDeliveryConfigurationStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateLogDeliveryConfigurationStatusRequestPb{}
	pb.LogDeliveryConfigurationId = st.LogDeliveryConfigurationId
	pb.Status = st.Status

	return pb, nil
}

type updateLogDeliveryConfigurationStatusRequestPb struct {
	LogDeliveryConfigurationId string                  `json:"-" url:"-"`
	Status                     LogDeliveryConfigStatus `json:"status"`
}

func updateLogDeliveryConfigurationStatusRequestFromPb(pb *updateLogDeliveryConfigurationStatusRequestPb) (*UpdateLogDeliveryConfigurationStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLogDeliveryConfigurationStatusRequest{}
	st.LogDeliveryConfigurationId = pb.LogDeliveryConfigurationId
	st.Status = pb.Status

	return st, nil
}

func wrappedCreateLogDeliveryConfigurationToPb(st *WrappedCreateLogDeliveryConfiguration) (*wrappedCreateLogDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &wrappedCreateLogDeliveryConfigurationPb{}
	pb.LogDeliveryConfiguration = st.LogDeliveryConfiguration

	return pb, nil
}

type wrappedCreateLogDeliveryConfigurationPb struct {
	LogDeliveryConfiguration CreateLogDeliveryConfigurationParams `json:"log_delivery_configuration"`
}

func wrappedCreateLogDeliveryConfigurationFromPb(pb *wrappedCreateLogDeliveryConfigurationPb) (*WrappedCreateLogDeliveryConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedCreateLogDeliveryConfiguration{}
	st.LogDeliveryConfiguration = pb.LogDeliveryConfiguration

	return st, nil
}

func wrappedLogDeliveryConfigurationToPb(st *WrappedLogDeliveryConfiguration) (*wrappedLogDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &wrappedLogDeliveryConfigurationPb{}
	pb.LogDeliveryConfiguration = st.LogDeliveryConfiguration

	return pb, nil
}

type wrappedLogDeliveryConfigurationPb struct {
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

func wrappedLogDeliveryConfigurationFromPb(pb *wrappedLogDeliveryConfigurationPb) (*WrappedLogDeliveryConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedLogDeliveryConfiguration{}
	st.LogDeliveryConfiguration = pb.LogDeliveryConfiguration

	return st, nil
}

func wrappedLogDeliveryConfigurationsToPb(st *WrappedLogDeliveryConfigurations) (*wrappedLogDeliveryConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &wrappedLogDeliveryConfigurationsPb{}
	pb.LogDeliveryConfigurations = st.LogDeliveryConfigurations
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type wrappedLogDeliveryConfigurationsPb struct {
	LogDeliveryConfigurations []LogDeliveryConfiguration `json:"log_delivery_configurations,omitempty"`
	NextPageToken             string                     `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func wrappedLogDeliveryConfigurationsFromPb(pb *wrappedLogDeliveryConfigurationsPb) (*WrappedLogDeliveryConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedLogDeliveryConfigurations{}
	st.LogDeliveryConfigurations = pb.LogDeliveryConfigurations
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *wrappedLogDeliveryConfigurationsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st wrappedLogDeliveryConfigurationsPb) MarshalJSON() ([]byte, error) {
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
