// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

type ActionConfiguration struct {
	// Databricks action configuration ID.
	// Wire name: 'action_configuration_id'
	ActionConfigurationId string
	// The type of the action.
	// Wire name: 'action_type'
	ActionType ActionConfigurationType
	// Target for the action. For example, an email address.
	// Wire name: 'target'
	Target string

	ForceSendFields []string `tf:"-"`
}

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

func (st *ActionConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &actionConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := actionConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ActionConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := actionConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type actionConfigurationPb struct {
	// Databricks action configuration ID.
	ActionConfigurationId string `json:"action_configuration_id,omitempty"`
	// The type of the action.
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target for the action. For example, an email address.
	Target string `json:"target,omitempty"`

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

type ActionConfigurationType string
type actionConfigurationTypePb string

const ActionConfigurationTypeEmailNotification ActionConfigurationType = `EMAIL_NOTIFICATION`

// String representation for [fmt.Print]
func (f *ActionConfigurationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActionConfigurationType) Set(v string) error {
	switch v {
	case `EMAIL_NOTIFICATION`:
		*f = ActionConfigurationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMAIL_NOTIFICATION"`, v)
	}
}

// Type always returns ActionConfigurationType to satisfy [pflag.Value] interface
func (f *ActionConfigurationType) Type() string {
	return "ActionConfigurationType"
}

func actionConfigurationTypeToPb(st *ActionConfigurationType) (*actionConfigurationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := actionConfigurationTypePb(*st)
	return &pb, nil
}

func actionConfigurationTypeFromPb(pb *actionConfigurationTypePb) (*ActionConfigurationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ActionConfigurationType(*pb)
	return &st, nil
}

type AlertConfiguration struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	// Wire name: 'action_configurations'
	ActionConfigurations []ActionConfiguration
	// Databricks alert configuration ID.
	// Wire name: 'alert_configuration_id'
	AlertConfigurationId string
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	// Wire name: 'quantity_threshold'
	QuantityThreshold string
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	// Wire name: 'quantity_type'
	QuantityType AlertConfigurationQuantityType
	// The time window of usage data for the budget.
	// Wire name: 'time_period'
	TimePeriod AlertConfigurationTimePeriod
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	// Wire name: 'trigger_type'
	TriggerType AlertConfigurationTriggerType

	ForceSendFields []string `tf:"-"`
}

func alertConfigurationToPb(st *AlertConfiguration) (*alertConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &alertConfigurationPb{}

	var actionConfigurationsPb []actionConfigurationPb
	for _, item := range st.ActionConfigurations {
		itemPb, err := actionConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			actionConfigurationsPb = append(actionConfigurationsPb, *itemPb)
		}
	}
	pb.ActionConfigurations = actionConfigurationsPb

	pb.AlertConfigurationId = st.AlertConfigurationId

	pb.QuantityThreshold = st.QuantityThreshold

	pb.QuantityType = st.QuantityType

	pb.TimePeriod = st.TimePeriod

	pb.TriggerType = st.TriggerType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *AlertConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &alertConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := alertConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AlertConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := alertConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type alertConfigurationPb struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations []actionConfigurationPb `json:"action_configurations,omitempty"`
	// Databricks alert configuration ID.
	AlertConfigurationId string `json:"alert_configuration_id,omitempty"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	QuantityThreshold string `json:"quantity_threshold,omitempty"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
	// The time window of usage data for the budget.
	TimePeriod AlertConfigurationTimePeriod `json:"time_period,omitempty"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType AlertConfigurationTriggerType `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func alertConfigurationFromPb(pb *alertConfigurationPb) (*AlertConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConfiguration{}

	var actionConfigurationsField []ActionConfiguration
	for _, itemPb := range pb.ActionConfigurations {
		item, err := actionConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			actionConfigurationsField = append(actionConfigurationsField, *item)
		}
	}
	st.ActionConfigurations = actionConfigurationsField
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

type AlertConfigurationQuantityType string
type alertConfigurationQuantityTypePb string

const AlertConfigurationQuantityTypeListPriceDollarsUsd AlertConfigurationQuantityType = `LIST_PRICE_DOLLARS_USD`

// String representation for [fmt.Print]
func (f *AlertConfigurationQuantityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertConfigurationQuantityType) Set(v string) error {
	switch v {
	case `LIST_PRICE_DOLLARS_USD`:
		*f = AlertConfigurationQuantityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LIST_PRICE_DOLLARS_USD"`, v)
	}
}

// Type always returns AlertConfigurationQuantityType to satisfy [pflag.Value] interface
func (f *AlertConfigurationQuantityType) Type() string {
	return "AlertConfigurationQuantityType"
}

func alertConfigurationQuantityTypeToPb(st *AlertConfigurationQuantityType) (*alertConfigurationQuantityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := alertConfigurationQuantityTypePb(*st)
	return &pb, nil
}

func alertConfigurationQuantityTypeFromPb(pb *alertConfigurationQuantityTypePb) (*AlertConfigurationQuantityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertConfigurationQuantityType(*pb)
	return &st, nil
}

type AlertConfigurationTimePeriod string
type alertConfigurationTimePeriodPb string

const AlertConfigurationTimePeriodMonth AlertConfigurationTimePeriod = `MONTH`

// String representation for [fmt.Print]
func (f *AlertConfigurationTimePeriod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertConfigurationTimePeriod) Set(v string) error {
	switch v {
	case `MONTH`:
		*f = AlertConfigurationTimePeriod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "MONTH"`, v)
	}
}

// Type always returns AlertConfigurationTimePeriod to satisfy [pflag.Value] interface
func (f *AlertConfigurationTimePeriod) Type() string {
	return "AlertConfigurationTimePeriod"
}

func alertConfigurationTimePeriodToPb(st *AlertConfigurationTimePeriod) (*alertConfigurationTimePeriodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := alertConfigurationTimePeriodPb(*st)
	return &pb, nil
}

func alertConfigurationTimePeriodFromPb(pb *alertConfigurationTimePeriodPb) (*AlertConfigurationTimePeriod, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertConfigurationTimePeriod(*pb)
	return &st, nil
}

type AlertConfigurationTriggerType string
type alertConfigurationTriggerTypePb string

const AlertConfigurationTriggerTypeCumulativeSpendingExceeded AlertConfigurationTriggerType = `CUMULATIVE_SPENDING_EXCEEDED`

// String representation for [fmt.Print]
func (f *AlertConfigurationTriggerType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AlertConfigurationTriggerType) Set(v string) error {
	switch v {
	case `CUMULATIVE_SPENDING_EXCEEDED`:
		*f = AlertConfigurationTriggerType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CUMULATIVE_SPENDING_EXCEEDED"`, v)
	}
}

// Type always returns AlertConfigurationTriggerType to satisfy [pflag.Value] interface
func (f *AlertConfigurationTriggerType) Type() string {
	return "AlertConfigurationTriggerType"
}

func alertConfigurationTriggerTypeToPb(st *AlertConfigurationTriggerType) (*alertConfigurationTriggerTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := alertConfigurationTriggerTypePb(*st)
	return &pb, nil
}

func alertConfigurationTriggerTypeFromPb(pb *alertConfigurationTriggerTypePb) (*AlertConfigurationTriggerType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertConfigurationTriggerType(*pb)
	return &st, nil
}

type BudgetConfiguration struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	// Wire name: 'alert_configurations'
	AlertConfigurations []AlertConfiguration
	// Databricks budget configuration ID.
	// Wire name: 'budget_configuration_id'
	BudgetConfigurationId string
	// Creation time of this budget configuration.
	// Wire name: 'create_time'
	CreateTime int64
	// Human-readable name of budget configuration. Max Length: 128
	// Wire name: 'display_name'
	DisplayName string
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	// Wire name: 'filter'
	Filter *BudgetConfigurationFilter
	// Update time of this budget configuration.
	// Wire name: 'update_time'
	UpdateTime int64

	ForceSendFields []string `tf:"-"`
}

func budgetConfigurationToPb(st *BudgetConfiguration) (*budgetConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationPb{}
	pb.AccountId = st.AccountId

	var alertConfigurationsPb []alertConfigurationPb
	for _, item := range st.AlertConfigurations {
		itemPb, err := alertConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			alertConfigurationsPb = append(alertConfigurationsPb, *itemPb)
		}
	}
	pb.AlertConfigurations = alertConfigurationsPb

	pb.BudgetConfigurationId = st.BudgetConfigurationId

	pb.CreateTime = st.CreateTime

	pb.DisplayName = st.DisplayName

	filterPb, err := budgetConfigurationFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}

	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *BudgetConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &budgetConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := budgetConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BudgetConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := budgetConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type budgetConfigurationPb struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations []alertConfigurationPb `json:"alert_configurations,omitempty"`
	// Databricks budget configuration ID.
	BudgetConfigurationId string `json:"budget_configuration_id,omitempty"`
	// Creation time of this budget configuration.
	CreateTime int64 `json:"create_time,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter *budgetConfigurationFilterPb `json:"filter,omitempty"`
	// Update time of this budget configuration.
	UpdateTime int64 `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func budgetConfigurationFromPb(pb *budgetConfigurationPb) (*BudgetConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfiguration{}
	st.AccountId = pb.AccountId

	var alertConfigurationsField []AlertConfiguration
	for _, itemPb := range pb.AlertConfigurations {
		item, err := alertConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			alertConfigurationsField = append(alertConfigurationsField, *item)
		}
	}
	st.AlertConfigurations = alertConfigurationsField
	st.BudgetConfigurationId = pb.BudgetConfigurationId
	st.CreateTime = pb.CreateTime
	st.DisplayName = pb.DisplayName
	filterField, err := budgetConfigurationFilterFromPb(pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = filterField
	}
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

type BudgetConfigurationFilter struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	// Wire name: 'tags'
	Tags []BudgetConfigurationFilterTagClause
	// If provided, usage must match with the provided Databricks workspace IDs.
	// Wire name: 'workspace_id'
	WorkspaceId *BudgetConfigurationFilterWorkspaceIdClause
}

func budgetConfigurationFilterToPb(st *BudgetConfigurationFilter) (*budgetConfigurationFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationFilterPb{}

	var tagsPb []budgetConfigurationFilterTagClausePb
	for _, item := range st.Tags {
		itemPb, err := budgetConfigurationFilterTagClauseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	workspaceIdPb, err := budgetConfigurationFilterWorkspaceIdClauseToPb(st.WorkspaceId)
	if err != nil {
		return nil, err
	}
	if workspaceIdPb != nil {
		pb.WorkspaceId = workspaceIdPb
	}

	return pb, nil
}

func (st *BudgetConfigurationFilter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &budgetConfigurationFilterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := budgetConfigurationFilterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BudgetConfigurationFilter) MarshalJSON() ([]byte, error) {
	pb, err := budgetConfigurationFilterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type budgetConfigurationFilterPb struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	Tags []budgetConfigurationFilterTagClausePb `json:"tags,omitempty"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	WorkspaceId *budgetConfigurationFilterWorkspaceIdClausePb `json:"workspace_id,omitempty"`
}

func budgetConfigurationFilterFromPb(pb *budgetConfigurationFilterPb) (*BudgetConfigurationFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilter{}

	var tagsField []BudgetConfigurationFilterTagClause
	for _, itemPb := range pb.Tags {
		item, err := budgetConfigurationFilterTagClauseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	workspaceIdField, err := budgetConfigurationFilterWorkspaceIdClauseFromPb(pb.WorkspaceId)
	if err != nil {
		return nil, err
	}
	if workspaceIdField != nil {
		st.WorkspaceId = workspaceIdField
	}

	return st, nil
}

type BudgetConfigurationFilterClause struct {

	// Wire name: 'operator'
	Operator BudgetConfigurationFilterOperator

	// Wire name: 'values'
	Values []string
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

func (st *BudgetConfigurationFilterClause) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &budgetConfigurationFilterClausePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := budgetConfigurationFilterClauseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BudgetConfigurationFilterClause) MarshalJSON() ([]byte, error) {
	pb, err := budgetConfigurationFilterClauseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type budgetConfigurationFilterClausePb struct {
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	Values []string `json:"values,omitempty"`
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

type BudgetConfigurationFilterOperator string
type budgetConfigurationFilterOperatorPb string

const BudgetConfigurationFilterOperatorIn BudgetConfigurationFilterOperator = `IN`

// String representation for [fmt.Print]
func (f *BudgetConfigurationFilterOperator) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *BudgetConfigurationFilterOperator) Set(v string) error {
	switch v {
	case `IN`:
		*f = BudgetConfigurationFilterOperator(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN"`, v)
	}
}

// Type always returns BudgetConfigurationFilterOperator to satisfy [pflag.Value] interface
func (f *BudgetConfigurationFilterOperator) Type() string {
	return "BudgetConfigurationFilterOperator"
}

func budgetConfigurationFilterOperatorToPb(st *BudgetConfigurationFilterOperator) (*budgetConfigurationFilterOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := budgetConfigurationFilterOperatorPb(*st)
	return &pb, nil
}

func budgetConfigurationFilterOperatorFromPb(pb *budgetConfigurationFilterOperatorPb) (*BudgetConfigurationFilterOperator, error) {
	if pb == nil {
		return nil, nil
	}
	st := BudgetConfigurationFilterOperator(*pb)
	return &st, nil
}

type BudgetConfigurationFilterTagClause struct {

	// Wire name: 'key'
	Key string

	// Wire name: 'value'
	Value *BudgetConfigurationFilterClause

	ForceSendFields []string `tf:"-"`
}

func budgetConfigurationFilterTagClauseToPb(st *BudgetConfigurationFilterTagClause) (*budgetConfigurationFilterTagClausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetConfigurationFilterTagClausePb{}
	pb.Key = st.Key

	valuePb, err := budgetConfigurationFilterClauseToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *BudgetConfigurationFilterTagClause) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &budgetConfigurationFilterTagClausePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := budgetConfigurationFilterTagClauseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BudgetConfigurationFilterTagClause) MarshalJSON() ([]byte, error) {
	pb, err := budgetConfigurationFilterTagClauseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type budgetConfigurationFilterTagClausePb struct {
	Key string `json:"key,omitempty"`

	Value *budgetConfigurationFilterClausePb `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func budgetConfigurationFilterTagClauseFromPb(pb *budgetConfigurationFilterTagClausePb) (*BudgetConfigurationFilterTagClause, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilterTagClause{}
	st.Key = pb.Key
	valueField, err := budgetConfigurationFilterClauseFromPb(pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *budgetConfigurationFilterTagClausePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st budgetConfigurationFilterTagClausePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BudgetConfigurationFilterWorkspaceIdClause struct {

	// Wire name: 'operator'
	Operator BudgetConfigurationFilterOperator

	// Wire name: 'values'
	Values []int64
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

func (st *BudgetConfigurationFilterWorkspaceIdClause) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &budgetConfigurationFilterWorkspaceIdClausePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := budgetConfigurationFilterWorkspaceIdClauseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BudgetConfigurationFilterWorkspaceIdClause) MarshalJSON() ([]byte, error) {
	pb, err := budgetConfigurationFilterWorkspaceIdClauseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type budgetConfigurationFilterWorkspaceIdClausePb struct {
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	Values []int64 `json:"values,omitempty"`
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

// Contains the BudgetPolicy details.
type BudgetPolicy struct {
	// List of workspaces that this budget policy will be exclusively bound to.
	// An empty binding implies that this budget policy is open to any workspace
	// in the account.
	// Wire name: 'binding_workspace_ids'
	BindingWorkspaceIds []int64
	// A list of tags defined by the customer. At most 20 entries are allowed
	// per policy.
	// Wire name: 'custom_tags'
	CustomTags []compute.CustomPolicyTag
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	// Wire name: 'policy_id'
	PolicyId string
	// The name of the policy. - Must be unique among active policies. - Can
	// contain only characters from the ISO 8859-1 (latin1) set. - Can't start
	// with reserved keywords such as `databricks:default-policy`.
	// Wire name: 'policy_name'
	PolicyName string

	ForceSendFields []string `tf:"-"`
}

func budgetPolicyToPb(st *BudgetPolicy) (*budgetPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &budgetPolicyPb{}
	pb.BindingWorkspaceIds = st.BindingWorkspaceIds

	var customTagsPb []compute.CustomPolicyTagPb
	for _, item := range st.CustomTags {
		itemPb, err := compute.CustomPolicyTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			customTagsPb = append(customTagsPb, *itemPb)
		}
	}
	pb.CustomTags = customTagsPb

	pb.PolicyId = st.PolicyId

	pb.PolicyName = st.PolicyName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *BudgetPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &budgetPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := budgetPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BudgetPolicy) MarshalJSON() ([]byte, error) {
	pb, err := budgetPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type budgetPolicyPb struct {
	// List of workspaces that this budget policy will be exclusively bound to.
	// An empty binding implies that this budget policy is open to any workspace
	// in the account.
	BindingWorkspaceIds []int64 `json:"binding_workspace_ids,omitempty"`
	// A list of tags defined by the customer. At most 20 entries are allowed
	// per policy.
	CustomTags []compute.CustomPolicyTagPb `json:"custom_tags,omitempty"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"policy_id,omitempty"`
	// The name of the policy. - Must be unique among active policies. - Can
	// contain only characters from the ISO 8859-1 (latin1) set. - Can't start
	// with reserved keywords such as `databricks:default-policy`.
	PolicyName string `json:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func budgetPolicyFromPb(pb *budgetPolicyPb) (*BudgetPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetPolicy{}
	st.BindingWorkspaceIds = pb.BindingWorkspaceIds

	var customTagsField []compute.CustomPolicyTag
	for _, itemPb := range pb.CustomTags {
		item, err := compute.CustomPolicyTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			customTagsField = append(customTagsField, *item)
		}
	}
	st.CustomTags = customTagsField
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

type CreateBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	// Wire name: 'dashboard_type'
	DashboardType UsageDashboardType
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
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

func (st *CreateBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBillingUsageDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBillingUsageDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := createBillingUsageDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBillingUsageDashboardRequestPb struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType UsageDashboardType `json:"dashboard_type,omitempty"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

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

type CreateBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBillingUsageDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBillingUsageDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := createBillingUsageDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBillingUsageDashboardResponsePb struct {
	// The unique id of the usage dashboard.
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

type CreateBudgetConfigurationBudget struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	// Wire name: 'alert_configurations'
	AlertConfigurations []CreateBudgetConfigurationBudgetAlertConfigurations
	// Human-readable name of budget configuration. Max Length: 128
	// Wire name: 'display_name'
	DisplayName string
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	// Wire name: 'filter'
	Filter *BudgetConfigurationFilter

	ForceSendFields []string `tf:"-"`
}

func createBudgetConfigurationBudgetToPb(st *CreateBudgetConfigurationBudget) (*createBudgetConfigurationBudgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationBudgetPb{}
	pb.AccountId = st.AccountId

	var alertConfigurationsPb []createBudgetConfigurationBudgetAlertConfigurationsPb
	for _, item := range st.AlertConfigurations {
		itemPb, err := createBudgetConfigurationBudgetAlertConfigurationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			alertConfigurationsPb = append(alertConfigurationsPb, *itemPb)
		}
	}
	pb.AlertConfigurations = alertConfigurationsPb

	pb.DisplayName = st.DisplayName

	filterPb, err := budgetConfigurationFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBudgetConfigurationBudgetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBudgetConfigurationBudgetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	pb, err := createBudgetConfigurationBudgetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBudgetConfigurationBudgetPb struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations []createBudgetConfigurationBudgetAlertConfigurationsPb `json:"alert_configurations,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter *budgetConfigurationFilterPb `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBudgetConfigurationBudgetFromPb(pb *createBudgetConfigurationBudgetPb) (*CreateBudgetConfigurationBudget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudget{}
	st.AccountId = pb.AccountId

	var alertConfigurationsField []CreateBudgetConfigurationBudgetAlertConfigurations
	for _, itemPb := range pb.AlertConfigurations {
		item, err := createBudgetConfigurationBudgetAlertConfigurationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			alertConfigurationsField = append(alertConfigurationsField, *item)
		}
	}
	st.AlertConfigurations = alertConfigurationsField
	st.DisplayName = pb.DisplayName
	filterField, err := budgetConfigurationFilterFromPb(pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = filterField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createBudgetConfigurationBudgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createBudgetConfigurationBudgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateBudgetConfigurationBudgetActionConfigurations struct {
	// The type of the action.
	// Wire name: 'action_type'
	ActionType ActionConfigurationType
	// Target for the action. For example, an email address.
	// Wire name: 'target'
	Target string

	ForceSendFields []string `tf:"-"`
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

func (st *CreateBudgetConfigurationBudgetActionConfigurations) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBudgetConfigurationBudgetActionConfigurationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBudgetConfigurationBudgetActionConfigurationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBudgetConfigurationBudgetActionConfigurations) MarshalJSON() ([]byte, error) {
	pb, err := createBudgetConfigurationBudgetActionConfigurationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBudgetConfigurationBudgetActionConfigurationsPb struct {
	// The type of the action.
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target for the action. For example, an email address.
	Target string `json:"target,omitempty"`

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

type CreateBudgetConfigurationBudgetAlertConfigurations struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	// Wire name: 'action_configurations'
	ActionConfigurations []CreateBudgetConfigurationBudgetActionConfigurations
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	// Wire name: 'quantity_threshold'
	QuantityThreshold string
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	// Wire name: 'quantity_type'
	QuantityType AlertConfigurationQuantityType
	// The time window of usage data for the budget.
	// Wire name: 'time_period'
	TimePeriod AlertConfigurationTimePeriod
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	// Wire name: 'trigger_type'
	TriggerType AlertConfigurationTriggerType

	ForceSendFields []string `tf:"-"`
}

func createBudgetConfigurationBudgetAlertConfigurationsToPb(st *CreateBudgetConfigurationBudgetAlertConfigurations) (*createBudgetConfigurationBudgetAlertConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationBudgetAlertConfigurationsPb{}

	var actionConfigurationsPb []createBudgetConfigurationBudgetActionConfigurationsPb
	for _, item := range st.ActionConfigurations {
		itemPb, err := createBudgetConfigurationBudgetActionConfigurationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			actionConfigurationsPb = append(actionConfigurationsPb, *itemPb)
		}
	}
	pb.ActionConfigurations = actionConfigurationsPb

	pb.QuantityThreshold = st.QuantityThreshold

	pb.QuantityType = st.QuantityType

	pb.TimePeriod = st.TimePeriod

	pb.TriggerType = st.TriggerType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateBudgetConfigurationBudgetAlertConfigurations) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBudgetConfigurationBudgetAlertConfigurationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBudgetConfigurationBudgetAlertConfigurationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBudgetConfigurationBudgetAlertConfigurations) MarshalJSON() ([]byte, error) {
	pb, err := createBudgetConfigurationBudgetAlertConfigurationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBudgetConfigurationBudgetAlertConfigurationsPb struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations []createBudgetConfigurationBudgetActionConfigurationsPb `json:"action_configurations,omitempty"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	QuantityThreshold string `json:"quantity_threshold,omitempty"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	QuantityType AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
	// The time window of usage data for the budget.
	TimePeriod AlertConfigurationTimePeriod `json:"time_period,omitempty"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	TriggerType AlertConfigurationTriggerType `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBudgetConfigurationBudgetAlertConfigurationsFromPb(pb *createBudgetConfigurationBudgetAlertConfigurationsPb) (*CreateBudgetConfigurationBudgetAlertConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudgetAlertConfigurations{}

	var actionConfigurationsField []CreateBudgetConfigurationBudgetActionConfigurations
	for _, itemPb := range pb.ActionConfigurations {
		item, err := createBudgetConfigurationBudgetActionConfigurationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			actionConfigurationsField = append(actionConfigurationsField, *item)
		}
	}
	st.ActionConfigurations = actionConfigurationsField
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

type CreateBudgetConfigurationRequest struct {
	// Properties of the new budget configuration.
	// Wire name: 'budget'
	Budget CreateBudgetConfigurationBudget
}

func createBudgetConfigurationRequestToPb(st *CreateBudgetConfigurationRequest) (*createBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationRequestPb{}
	budgetPb, err := createBudgetConfigurationBudgetToPb(&st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = *budgetPb
	}

	return pb, nil
}

func (st *CreateBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := createBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBudgetConfigurationRequestPb struct {
	// Properties of the new budget configuration.
	Budget createBudgetConfigurationBudgetPb `json:"budget"`
}

func createBudgetConfigurationRequestFromPb(pb *createBudgetConfigurationRequestPb) (*CreateBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationRequest{}
	budgetField, err := createBudgetConfigurationBudgetFromPb(&pb.Budget)
	if err != nil {
		return nil, err
	}
	if budgetField != nil {
		st.Budget = *budgetField
	}

	return st, nil
}

type CreateBudgetConfigurationResponse struct {
	// The created budget configuration.
	// Wire name: 'budget'
	Budget *BudgetConfiguration
}

func createBudgetConfigurationResponseToPb(st *CreateBudgetConfigurationResponse) (*createBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetConfigurationResponsePb{}
	budgetPb, err := budgetConfigurationToPb(st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = budgetPb
	}

	return pb, nil
}

func (st *CreateBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBudgetConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBudgetConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBudgetConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := createBudgetConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBudgetConfigurationResponsePb struct {
	// The created budget configuration.
	Budget *budgetConfigurationPb `json:"budget,omitempty"`
}

func createBudgetConfigurationResponseFromPb(pb *createBudgetConfigurationResponsePb) (*CreateBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationResponse{}
	budgetField, err := budgetConfigurationFromPb(pb.Budget)
	if err != nil {
		return nil, err
	}
	if budgetField != nil {
		st.Budget = budgetField
	}

	return st, nil
}

// A request to create a BudgetPolicy.
type CreateBudgetPolicyRequest struct {
	// The policy to create. `policy_id` needs to be empty as it will be
	// generated `policy_name` must be provided, custom_tags may need to be
	// provided depending on the cloud provider. All other fields are optional.
	// Wire name: 'policy'
	Policy *BudgetPolicy
	// A unique identifier for this request. Restricted to 36 ASCII characters.
	// A random UUID is recommended. This request is only idempotent if a
	// `request_id` is provided.
	// Wire name: 'request_id'
	RequestId string

	ForceSendFields []string `tf:"-"`
}

func createBudgetPolicyRequestToPb(st *CreateBudgetPolicyRequest) (*createBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createBudgetPolicyRequestPb{}
	policyPb, err := budgetPolicyToPb(st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = policyPb
	}

	pb.RequestId = st.RequestId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createBudgetPolicyRequestPb struct {
	// The policy to create. `policy_id` needs to be empty as it will be
	// generated `policy_name` must be provided, custom_tags may need to be
	// provided depending on the cloud provider. All other fields are optional.
	Policy *budgetPolicyPb `json:"policy,omitempty"`
	// A unique identifier for this request. Restricted to 36 ASCII characters.
	// A random UUID is recommended. This request is only idempotent if a
	// `request_id` is provided.
	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createBudgetPolicyRequestFromPb(pb *createBudgetPolicyRequestPb) (*CreateBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetPolicyRequest{}
	policyField, err := budgetPolicyFromPb(pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = policyField
	}
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

type CreateLogDeliveryConfigurationParams struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	// Wire name: 'config_name'
	ConfigName string
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'credentials_id'
	CredentialsId string
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	// Wire name: 'delivery_path_prefix'
	DeliveryPathPrefix string
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	// Wire name: 'delivery_start_time'
	DeliveryStartTime string
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'log_type'
	LogType LogType
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// Wire name: 'output_format'
	OutputFormat OutputFormat
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	// Wire name: 'status'
	Status LogDeliveryConfigStatus
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	// Wire name: 'workspace_ids_filter'
	WorkspaceIdsFilter []int64

	ForceSendFields []string `tf:"-"`
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

func (st *CreateLogDeliveryConfigurationParams) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createLogDeliveryConfigurationParamsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createLogDeliveryConfigurationParamsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateLogDeliveryConfigurationParams) MarshalJSON() ([]byte, error) {
	pb, err := createLogDeliveryConfigurationParamsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createLogDeliveryConfigurationParamsPb struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `json:"config_name,omitempty"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId string `json:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType LogType `json:"log_type"`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat OutputFormat `json:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId string `json:"storage_configuration_id"`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`

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

// Delete budget
type DeleteBudgetConfigurationRequest struct {
	// The Databricks budget configuration ID.
	// Wire name: 'budget_id'
	BudgetId string `tf:"-"`
}

func deleteBudgetConfigurationRequestToPb(st *DeleteBudgetConfigurationRequest) (*deleteBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteBudgetConfigurationRequestPb{}
	pb.BudgetId = st.BudgetId

	return pb, nil
}

func (st *DeleteBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteBudgetConfigurationRequestPb struct {
	// The Databricks budget configuration ID.
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

type DeleteBudgetConfigurationResponse struct {
}

func deleteBudgetConfigurationResponseToPb(st *DeleteBudgetConfigurationResponse) (*deleteBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteBudgetConfigurationResponsePb{}

	return pb, nil
}

func (st *DeleteBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteBudgetConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteBudgetConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteBudgetConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteBudgetConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteBudgetConfigurationResponsePb struct {
}

func deleteBudgetConfigurationResponseFromPb(pb *deleteBudgetConfigurationResponsePb) (*DeleteBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteBudgetConfigurationResponse{}

	return st, nil
}

// Delete a budget policy
type DeleteBudgetPolicyRequest struct {
	// The Id of the policy.
	// Wire name: 'policy_id'
	PolicyId string `tf:"-"`
}

func deleteBudgetPolicyRequestToPb(st *DeleteBudgetPolicyRequest) (*deleteBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteBudgetPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func (st *DeleteBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteBudgetPolicyRequestPb struct {
	// The Id of the policy.
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

// The status string for log delivery. Possible values are: * `CREATED`: There
// were no log delivery attempts since the config was created. * `SUCCEEDED`:
// The latest attempt of log delivery has succeeded completely. *
// `USER_FAILURE`: The latest attempt of log delivery failed because of
// misconfiguration of customer provided permissions on role or storage. *
// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
// Databricks internal error. Contact support if it doesn't go away soon. *
// `NOT_FOUND`: The log delivery status as the configuration has been disabled
// since the release of this feature or there are no workspaces in the account.
type DeliveryStatus string
type deliveryStatusPb string

// There were no log delivery attempts since the config was created.
const DeliveryStatusCreated DeliveryStatus = `CREATED`

// The log delivery status as the configuration has been disabled since the
// release of this feature or there are no workspaces in the account.
const DeliveryStatusNotFound DeliveryStatus = `NOT_FOUND`

// The latest attempt of log delivery has succeeded completely.
const DeliveryStatusSucceeded DeliveryStatus = `SUCCEEDED`

// The latest attempt of log delivery failed because of an <Databricks> internal
// error. Contact support if it doesn't go away soon.
const DeliveryStatusSystemFailure DeliveryStatus = `SYSTEM_FAILURE`

// The latest attempt of log delivery failed because of misconfiguration of
// customer provided permissions on role or storage.
const DeliveryStatusUserFailure DeliveryStatus = `USER_FAILURE`

// String representation for [fmt.Print]
func (f *DeliveryStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeliveryStatus) Set(v string) error {
	switch v {
	case `CREATED`, `NOT_FOUND`, `SUCCEEDED`, `SYSTEM_FAILURE`, `USER_FAILURE`:
		*f = DeliveryStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CREATED", "NOT_FOUND", "SUCCEEDED", "SYSTEM_FAILURE", "USER_FAILURE"`, v)
	}
}

// Type always returns DeliveryStatus to satisfy [pflag.Value] interface
func (f *DeliveryStatus) Type() string {
	return "DeliveryStatus"
}

func deliveryStatusToPb(st *DeliveryStatus) (*deliveryStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := deliveryStatusPb(*st)
	return &pb, nil
}

func deliveryStatusFromPb(pb *deliveryStatusPb) (*DeliveryStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeliveryStatus(*pb)
	return &st, nil
}

// Return billable usage logs
type DownloadRequest struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	// Wire name: 'end_month'
	EndMonth string `tf:"-"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	// Wire name: 'personal_data'
	PersonalData bool `tf:"-"`
	// Format: `YYYY-MM`. First month to return billable usage logs for. This
	// field is required.
	// Wire name: 'start_month'
	StartMonth string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *DownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &downloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := downloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := downloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type downloadRequestPb struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth string `json:"-" url:"end_month"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData bool `json:"-" url:"personal_data,omitempty"`
	// Format: `YYYY-MM`. First month to return billable usage logs for. This
	// field is required.
	StartMonth string `json:"-" url:"start_month"`

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

type DownloadResponse struct {

	// Wire name: 'contents'
	Contents io.ReadCloser `tf:"-"`
}

func downloadResponseToPb(st *DownloadResponse) (*downloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &downloadResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

func (st *DownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &downloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := downloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := downloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Structured representation of a filter to be applied to a list of policies.
// All specified filters will be applied in conjunction.
type Filter struct {
	// The policy creator user id to be filtered on. If unspecified, all
	// policies will be returned.
	// Wire name: 'creator_user_id'
	CreatorUserId int64
	// The policy creator user name to be filtered on. If unspecified, all
	// policies will be returned.
	// Wire name: 'creator_user_name'
	CreatorUserName string
	// The partial name of policies to be filtered on. If unspecified, all
	// policies will be returned.
	// Wire name: 'policy_name'
	PolicyName string

	ForceSendFields []string `tf:"-"`
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

func (st *Filter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := filterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Filter) MarshalJSON() ([]byte, error) {
	pb, err := filterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type filterPb struct {
	// The policy creator user id to be filtered on. If unspecified, all
	// policies will be returned.
	CreatorUserId int64 `json:"creator_user_id,omitempty" url:"creator_user_id,omitempty"`
	// The policy creator user name to be filtered on. If unspecified, all
	// policies will be returned.
	CreatorUserName string `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	// The partial name of policies to be filtered on. If unspecified, all
	// policies will be returned.
	PolicyName string `json:"policy_name,omitempty" url:"policy_name,omitempty"`

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

// Get usage dashboard
type GetBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	// Wire name: 'dashboard_type'
	DashboardType UsageDashboardType `tf:"-"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *GetBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getBillingUsageDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getBillingUsageDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := getBillingUsageDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getBillingUsageDashboardRequestPb struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType UsageDashboardType `json:"-" url:"dashboard_type,omitempty"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId int64 `json:"-" url:"workspace_id,omitempty"`

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

type GetBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string
	// The URL of the usage dashboard.
	// Wire name: 'dashboard_url'
	DashboardUrl string

	ForceSendFields []string `tf:"-"`
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

func (st *GetBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getBillingUsageDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getBillingUsageDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := getBillingUsageDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getBillingUsageDashboardResponsePb struct {
	// The unique id of the usage dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The URL of the usage dashboard.
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

// Get budget
type GetBudgetConfigurationRequest struct {
	// The budget configuration ID
	// Wire name: 'budget_id'
	BudgetId string `tf:"-"`
}

func getBudgetConfigurationRequestToPb(st *GetBudgetConfigurationRequest) (*getBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBudgetConfigurationRequestPb{}
	pb.BudgetId = st.BudgetId

	return pb, nil
}

func (st *GetBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := getBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getBudgetConfigurationRequestPb struct {
	// The budget configuration ID
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

type GetBudgetConfigurationResponse struct {

	// Wire name: 'budget'
	Budget *BudgetConfiguration
}

func getBudgetConfigurationResponseToPb(st *GetBudgetConfigurationResponse) (*getBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBudgetConfigurationResponsePb{}
	budgetPb, err := budgetConfigurationToPb(st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = budgetPb
	}

	return pb, nil
}

func (st *GetBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getBudgetConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getBudgetConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetBudgetConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := getBudgetConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getBudgetConfigurationResponsePb struct {
	Budget *budgetConfigurationPb `json:"budget,omitempty"`
}

func getBudgetConfigurationResponseFromPb(pb *getBudgetConfigurationResponsePb) (*GetBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBudgetConfigurationResponse{}
	budgetField, err := budgetConfigurationFromPb(pb.Budget)
	if err != nil {
		return nil, err
	}
	if budgetField != nil {
		st.Budget = budgetField
	}

	return st, nil
}

// Get a budget policy
type GetBudgetPolicyRequest struct {
	// The Id of the policy.
	// Wire name: 'policy_id'
	PolicyId string `tf:"-"`
}

func getBudgetPolicyRequestToPb(st *GetBudgetPolicyRequest) (*getBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getBudgetPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func (st *GetBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getBudgetPolicyRequestPb struct {
	// The Id of the policy.
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

// Get log delivery configuration
type GetLogDeliveryRequest struct {
	// Databricks log delivery configuration ID
	// Wire name: 'log_delivery_configuration_id'
	LogDeliveryConfigurationId string `tf:"-"`
}

func getLogDeliveryRequestToPb(st *GetLogDeliveryRequest) (*getLogDeliveryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLogDeliveryRequestPb{}
	pb.LogDeliveryConfigurationId = st.LogDeliveryConfigurationId

	return pb, nil
}

func (st *GetLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLogDeliveryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLogDeliveryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLogDeliveryRequest) MarshalJSON() ([]byte, error) {
	pb, err := getLogDeliveryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getLogDeliveryRequestPb struct {
	// Databricks log delivery configuration ID
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

// The limit configuration of the policy. Limit configuration provide a budget
// policy level cost control by enforcing the limit.
type LimitConfig struct {
}

func limitConfigToPb(st *LimitConfig) (*limitConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &limitConfigPb{}

	return pb, nil
}

func (st *LimitConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &limitConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := limitConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LimitConfig) MarshalJSON() ([]byte, error) {
	pb, err := limitConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Get all budgets
type ListBudgetConfigurationsRequest struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
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

func (st *ListBudgetConfigurationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listBudgetConfigurationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listBudgetConfigurationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListBudgetConfigurationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listBudgetConfigurationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listBudgetConfigurationsRequestPb struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
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

type ListBudgetConfigurationsResponse struct {

	// Wire name: 'budgets'
	Budgets []BudgetConfiguration
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listBudgetConfigurationsResponseToPb(st *ListBudgetConfigurationsResponse) (*listBudgetConfigurationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listBudgetConfigurationsResponsePb{}

	var budgetsPb []budgetConfigurationPb
	for _, item := range st.Budgets {
		itemPb, err := budgetConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			budgetsPb = append(budgetsPb, *itemPb)
		}
	}
	pb.Budgets = budgetsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListBudgetConfigurationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listBudgetConfigurationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listBudgetConfigurationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListBudgetConfigurationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listBudgetConfigurationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listBudgetConfigurationsResponsePb struct {
	Budgets []budgetConfigurationPb `json:"budgets,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listBudgetConfigurationsResponseFromPb(pb *listBudgetConfigurationsResponsePb) (*ListBudgetConfigurationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetConfigurationsResponse{}

	var budgetsField []BudgetConfiguration
	for _, itemPb := range pb.Budgets {
		item, err := budgetConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			budgetsField = append(budgetsField, *item)
		}
	}
	st.Budgets = budgetsField
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

// List policies
type ListBudgetPoliciesRequest struct {
	// A filter to apply to the list of policies.
	// Wire name: 'filter_by'
	FilterBy *Filter `tf:"-"`
	// The maximum number of budget policies to return. If unspecified, at most
	// 100 budget policies will be returned. The maximum value is 1000; values
	// above 1000 will be coerced to 1000.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// A page token, received from a previous `ListServerlessPolicies` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	//
	// When paginating, all other parameters provided to
	// `ListServerlessPoliciesRequest` must match the call that provided the
	// page token.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// The sort specification.
	// Wire name: 'sort_spec'
	SortSpec *SortSpec `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listBudgetPoliciesRequestToPb(st *ListBudgetPoliciesRequest) (*listBudgetPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listBudgetPoliciesRequestPb{}
	filterByPb, err := filterToPb(st.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByPb != nil {
		pb.FilterBy = filterByPb
	}

	pb.PageSize = st.PageSize

	pb.PageToken = st.PageToken

	sortSpecPb, err := sortSpecToPb(st.SortSpec)
	if err != nil {
		return nil, err
	}
	if sortSpecPb != nil {
		pb.SortSpec = sortSpecPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListBudgetPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listBudgetPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listBudgetPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListBudgetPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listBudgetPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listBudgetPoliciesRequestPb struct {
	// A filter to apply to the list of policies.
	FilterBy *filterPb `json:"-" url:"filter_by,omitempty"`
	// The maximum number of budget policies to return. If unspecified, at most
	// 100 budget policies will be returned. The maximum value is 1000; values
	// above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListServerlessPolicies` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	//
	// When paginating, all other parameters provided to
	// `ListServerlessPoliciesRequest` must match the call that provided the
	// page token.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The sort specification.
	SortSpec *sortSpecPb `json:"-" url:"sort_spec,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listBudgetPoliciesRequestFromPb(pb *listBudgetPoliciesRequestPb) (*ListBudgetPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetPoliciesRequest{}
	filterByField, err := filterFromPb(pb.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByField != nil {
		st.FilterBy = filterByField
	}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	sortSpecField, err := sortSpecFromPb(pb.SortSpec)
	if err != nil {
		return nil, err
	}
	if sortSpecField != nil {
		st.SortSpec = sortSpecField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listBudgetPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listBudgetPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A list of policies.
type ListBudgetPoliciesResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'policies'
	Policies []BudgetPolicy
	// A token that can be sent as `page_token` to retrieve the previous page.
	// In this field is omitted, there are no previous pages.
	// Wire name: 'previous_page_token'
	PreviousPageToken string

	ForceSendFields []string `tf:"-"`
}

func listBudgetPoliciesResponseToPb(st *ListBudgetPoliciesResponse) (*listBudgetPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listBudgetPoliciesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var policiesPb []budgetPolicyPb
	for _, item := range st.Policies {
		itemPb, err := budgetPolicyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policiesPb = append(policiesPb, *itemPb)
		}
	}
	pb.Policies = policiesPb

	pb.PreviousPageToken = st.PreviousPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListBudgetPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listBudgetPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listBudgetPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListBudgetPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listBudgetPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listBudgetPoliciesResponsePb struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	Policies []budgetPolicyPb `json:"policies,omitempty"`
	// A token that can be sent as `page_token` to retrieve the previous page.
	// In this field is omitted, there are no previous pages.
	PreviousPageToken string `json:"previous_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listBudgetPoliciesResponseFromPb(pb *listBudgetPoliciesResponsePb) (*ListBudgetPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetPoliciesResponse{}
	st.NextPageToken = pb.NextPageToken

	var policiesField []BudgetPolicy
	for _, itemPb := range pb.Policies {
		item, err := budgetPolicyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			policiesField = append(policiesField, *item)
		}
	}
	st.Policies = policiesField
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

// Get all log delivery configurations
type ListLogDeliveryRequest struct {
	// Filter by credential configuration ID.
	// Wire name: 'credentials_id'
	CredentialsId string `tf:"-"`
	// Filter by status `ENABLED` or `DISABLED`.
	// Wire name: 'status'
	Status LogDeliveryConfigStatus `tf:"-"`
	// Filter by storage configuration ID.
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listLogDeliveryRequestToPb(st *ListLogDeliveryRequest) (*listLogDeliveryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listLogDeliveryRequestPb{}
	pb.CredentialsId = st.CredentialsId

	pb.Status = st.Status

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listLogDeliveryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listLogDeliveryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListLogDeliveryRequest) MarshalJSON() ([]byte, error) {
	pb, err := listLogDeliveryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listLogDeliveryRequestPb struct {
	// Filter by credential configuration ID.
	CredentialsId string `json:"-" url:"credentials_id,omitempty"`
	// Filter by status `ENABLED` or `DISABLED`.
	Status LogDeliveryConfigStatus `json:"-" url:"status,omitempty"`
	// Filter by storage configuration ID.
	StorageConfigurationId string `json:"-" url:"storage_configuration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listLogDeliveryRequestFromPb(pb *listLogDeliveryRequestPb) (*ListLogDeliveryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListLogDeliveryRequest{}
	st.CredentialsId = pb.CredentialsId
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

// Status of log delivery configuration. Set to `ENABLED` (enabled) or
// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable the
// configuration](#operation/patch-log-delivery-config-status) later. Deletion
// of a configuration is not supported, so disable a log delivery configuration
// that is no longer needed.
type LogDeliveryConfigStatus string
type logDeliveryConfigStatusPb string

const LogDeliveryConfigStatusDisabled LogDeliveryConfigStatus = `DISABLED`

const LogDeliveryConfigStatusEnabled LogDeliveryConfigStatus = `ENABLED`

// String representation for [fmt.Print]
func (f *LogDeliveryConfigStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LogDeliveryConfigStatus) Set(v string) error {
	switch v {
	case `DISABLED`, `ENABLED`:
		*f = LogDeliveryConfigStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISABLED", "ENABLED"`, v)
	}
}

// Type always returns LogDeliveryConfigStatus to satisfy [pflag.Value] interface
func (f *LogDeliveryConfigStatus) Type() string {
	return "LogDeliveryConfigStatus"
}

func logDeliveryConfigStatusToPb(st *LogDeliveryConfigStatus) (*logDeliveryConfigStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := logDeliveryConfigStatusPb(*st)
	return &pb, nil
}

func logDeliveryConfigStatusFromPb(pb *logDeliveryConfigStatusPb) (*LogDeliveryConfigStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := LogDeliveryConfigStatus(*pb)
	return &st, nil
}

type LogDeliveryConfiguration struct {
	// The Databricks account ID that hosts the log delivery configuration.
	// Wire name: 'account_id'
	AccountId string
	// Databricks log delivery configuration ID.
	// Wire name: 'config_id'
	ConfigId string
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	// Wire name: 'config_name'
	ConfigName string
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	// Wire name: 'creation_time'
	CreationTime int64
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'credentials_id'
	CredentialsId string
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	// Wire name: 'delivery_path_prefix'
	DeliveryPathPrefix string
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	// Wire name: 'delivery_start_time'
	DeliveryStartTime string
	// Databricks log delivery status.
	// Wire name: 'log_delivery_status'
	LogDeliveryStatus *LogDeliveryStatus
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'log_type'
	LogType LogType
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// Wire name: 'output_format'
	OutputFormat OutputFormat
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	// Wire name: 'status'
	Status LogDeliveryConfigStatus
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	// Wire name: 'update_time'
	UpdateTime int64
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	// Wire name: 'workspace_ids_filter'
	WorkspaceIdsFilter []int64

	ForceSendFields []string `tf:"-"`
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

	logDeliveryStatusPb, err := logDeliveryStatusToPb(st.LogDeliveryStatus)
	if err != nil {
		return nil, err
	}
	if logDeliveryStatusPb != nil {
		pb.LogDeliveryStatus = logDeliveryStatusPb
	}

	pb.LogType = st.LogType

	pb.OutputFormat = st.OutputFormat

	pb.Status = st.Status

	pb.StorageConfigurationId = st.StorageConfigurationId

	pb.UpdateTime = st.UpdateTime

	pb.WorkspaceIdsFilter = st.WorkspaceIdsFilter

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *LogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logDeliveryConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logDeliveryConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := logDeliveryConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logDeliveryConfigurationPb struct {
	// The Databricks account ID that hosts the log delivery configuration.
	AccountId string `json:"account_id,omitempty"`
	// Databricks log delivery configuration ID.
	ConfigId string `json:"config_id,omitempty"`
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	ConfigName string `json:"config_name,omitempty"`
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	CredentialsId string `json:"credentials_id,omitempty"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if `log_type` is `BILLABLE_USAGE`. This is the
	// optional start month and year for delivery, specified in `YYYY-MM`
	// format. Defaults to current year and month. `BILLABLE_USAGE` logs are not
	// available for usage before March 2019 (`2019-03`).
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// Databricks log delivery status.
	LogDeliveryStatus *logDeliveryStatusPb `json:"log_delivery_status,omitempty"`
	// Log delivery type. Supported values are:
	//
	// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the
	// CSV schema, see the [View billable usage].
	//
	// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema,
	// see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	LogType LogType `json:"log_type,omitempty"`
	// The file type of log delivery.
	//
	// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the
	// CSV (comma-separated values) format is supported. For the schema, see the
	// [View billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be
	// `JSON`. Only the JSON (JavaScript Object Notation) format is supported.
	// For the schema, see the [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	OutputFormat OutputFormat `json:"output_format,omitempty"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	StorageConfigurationId string `json:"storage_configuration_id,omitempty"`
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	UpdateTime int64 `json:"update_time,omitempty"`
	// Optional filter that specifies workspace IDs to deliver logs for. By
	// default the workspace filter is empty and log delivery applies at the
	// account level, delivering workspace-level logs for all workspaces in your
	// account, plus account level logs. You can optionally set this field to an
	// array of workspace IDs (each one is an `int64`) to which log delivery
	// should apply, in which case only workspace-level logs relating to the
	// specified workspaces are delivered. If you plan to use different log
	// delivery configurations for different workspaces, set this field
	// explicitly. Be aware that delivery configurations mentioning specific
	// workspaces won't apply to new workspaces created in the future, and
	// delivery won't include account level logs. For some types of Databricks
	// deployments there is only one workspace per account ID, so this field is
	// unnecessary.
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`

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
	logDeliveryStatusField, err := logDeliveryStatusFromPb(pb.LogDeliveryStatus)
	if err != nil {
		return nil, err
	}
	if logDeliveryStatusField != nil {
		st.LogDeliveryStatus = logDeliveryStatusField
	}
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

// Databricks log delivery status.
type LogDeliveryStatus struct {
	// The UTC time for the latest log delivery attempt.
	// Wire name: 'last_attempt_time'
	LastAttemptTime string
	// The UTC time for the latest successful log delivery.
	// Wire name: 'last_successful_attempt_time'
	LastSuccessfulAttemptTime string
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	// Wire name: 'message'
	Message string
	// The status string for log delivery. Possible values are: * `CREATED`:
	// There were no log delivery attempts since the config was created. *
	// `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
	// * `USER_FAILURE`: The latest attempt of log delivery failed because of
	// misconfiguration of customer provided permissions on role or storage. *
	// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
	// Databricks internal error. Contact support if it doesn't go away soon. *
	// `NOT_FOUND`: The log delivery status as the configuration has been
	// disabled since the release of this feature or there are no workspaces in
	// the account.
	// Wire name: 'status'
	Status DeliveryStatus

	ForceSendFields []string `tf:"-"`
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

func (st *LogDeliveryStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logDeliveryStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logDeliveryStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogDeliveryStatus) MarshalJSON() ([]byte, error) {
	pb, err := logDeliveryStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logDeliveryStatusPb struct {
	// The UTC time for the latest log delivery attempt.
	LastAttemptTime string `json:"last_attempt_time,omitempty"`
	// The UTC time for the latest successful log delivery.
	LastSuccessfulAttemptTime string `json:"last_successful_attempt_time,omitempty"`
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	Message string `json:"message,omitempty"`
	// The status string for log delivery. Possible values are: * `CREATED`:
	// There were no log delivery attempts since the config was created. *
	// `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
	// * `USER_FAILURE`: The latest attempt of log delivery failed because of
	// misconfiguration of customer provided permissions on role or storage. *
	// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
	// Databricks internal error. Contact support if it doesn't go away soon. *
	// `NOT_FOUND`: The log delivery status as the configuration has been
	// disabled since the release of this feature or there are no workspaces in
	// the account.
	Status DeliveryStatus `json:"status,omitempty"`

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

// Log delivery type. Supported values are:
//
// * `BILLABLE_USAGE` — Configure [billable usage log delivery]. For the CSV
// schema, see the [View billable usage].
//
// * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON schema, see
// [Configure audit logging]
//
// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
type LogType string
type logTypePb string

const LogTypeAuditLogs LogType = `AUDIT_LOGS`

const LogTypeBillableUsage LogType = `BILLABLE_USAGE`

// String representation for [fmt.Print]
func (f *LogType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LogType) Set(v string) error {
	switch v {
	case `AUDIT_LOGS`, `BILLABLE_USAGE`:
		*f = LogType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AUDIT_LOGS", "BILLABLE_USAGE"`, v)
	}
}

// Type always returns LogType to satisfy [pflag.Value] interface
func (f *LogType) Type() string {
	return "LogType"
}

func logTypeToPb(st *LogType) (*logTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := logTypePb(*st)
	return &pb, nil
}

func logTypeFromPb(pb *logTypePb) (*LogType, error) {
	if pb == nil {
		return nil, nil
	}
	st := LogType(*pb)
	return &st, nil
}

// The file type of log delivery.
//
// * If `log_type` is `BILLABLE_USAGE`, this value must be `CSV`. Only the CSV
// (comma-separated values) format is supported. For the schema, see the [View
// billable usage] * If `log_type` is `AUDIT_LOGS`, this value must be `JSON`.
// Only the JSON (JavaScript Object Notation) format is supported. For the
// schema, see the [Configuring audit logs].
//
// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
type OutputFormat string
type outputFormatPb string

const OutputFormatCsv OutputFormat = `CSV`

const OutputFormatJson OutputFormat = `JSON`

// String representation for [fmt.Print]
func (f *OutputFormat) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OutputFormat) Set(v string) error {
	switch v {
	case `CSV`, `JSON`:
		*f = OutputFormat(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CSV", "JSON"`, v)
	}
}

// Type always returns OutputFormat to satisfy [pflag.Value] interface
func (f *OutputFormat) Type() string {
	return "OutputFormat"
}

func outputFormatToPb(st *OutputFormat) (*outputFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := outputFormatPb(*st)
	return &pb, nil
}

func outputFormatFromPb(pb *outputFormatPb) (*OutputFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := OutputFormat(*pb)
	return &st, nil
}

type PatchStatusResponse struct {
}

func patchStatusResponseToPb(st *PatchStatusResponse) (*patchStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchStatusResponsePb{}

	return pb, nil
}

func (st *PatchStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &patchStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := patchStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PatchStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := patchStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type patchStatusResponsePb struct {
}

func patchStatusResponseFromPb(pb *patchStatusResponsePb) (*PatchStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchStatusResponse{}

	return st, nil
}

type SortSpec struct {
	// Whether to sort in descending order.
	// Wire name: 'descending'
	Descending bool
	// The filed to sort by
	// Wire name: 'field'
	Field SortSpecField

	ForceSendFields []string `tf:"-"`
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

func (st *SortSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sortSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sortSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SortSpec) MarshalJSON() ([]byte, error) {
	pb, err := sortSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type sortSpecPb struct {
	// Whether to sort in descending order.
	Descending bool `json:"descending,omitempty" url:"descending,omitempty"`
	// The filed to sort by
	Field SortSpecField `json:"field,omitempty" url:"field,omitempty"`

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

type SortSpecField string
type sortSpecFieldPb string

const SortSpecFieldPolicyName SortSpecField = `POLICY_NAME`

// String representation for [fmt.Print]
func (f *SortSpecField) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SortSpecField) Set(v string) error {
	switch v {
	case `POLICY_NAME`:
		*f = SortSpecField(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "POLICY_NAME"`, v)
	}
}

// Type always returns SortSpecField to satisfy [pflag.Value] interface
func (f *SortSpecField) Type() string {
	return "SortSpecField"
}

func sortSpecFieldToPb(st *SortSpecField) (*sortSpecFieldPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := sortSpecFieldPb(*st)
	return &pb, nil
}

func sortSpecFieldFromPb(pb *sortSpecFieldPb) (*SortSpecField, error) {
	if pb == nil {
		return nil, nil
	}
	st := SortSpecField(*pb)
	return &st, nil
}

type UpdateBudgetConfigurationBudget struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	// Wire name: 'alert_configurations'
	AlertConfigurations []AlertConfiguration
	// Databricks budget configuration ID.
	// Wire name: 'budget_configuration_id'
	BudgetConfigurationId string
	// Human-readable name of budget configuration. Max Length: 128
	// Wire name: 'display_name'
	DisplayName string
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	// Wire name: 'filter'
	Filter *BudgetConfigurationFilter

	ForceSendFields []string `tf:"-"`
}

func updateBudgetConfigurationBudgetToPb(st *UpdateBudgetConfigurationBudget) (*updateBudgetConfigurationBudgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetConfigurationBudgetPb{}
	pb.AccountId = st.AccountId

	var alertConfigurationsPb []alertConfigurationPb
	for _, item := range st.AlertConfigurations {
		itemPb, err := alertConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			alertConfigurationsPb = append(alertConfigurationsPb, *itemPb)
		}
	}
	pb.AlertConfigurations = alertConfigurationsPb

	pb.BudgetConfigurationId = st.BudgetConfigurationId

	pb.DisplayName = st.DisplayName

	filterPb, err := budgetConfigurationFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateBudgetConfigurationBudgetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateBudgetConfigurationBudgetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	pb, err := updateBudgetConfigurationBudgetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateBudgetConfigurationBudgetPb struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations []alertConfigurationPb `json:"alert_configurations,omitempty"`
	// Databricks budget configuration ID.
	BudgetConfigurationId string `json:"budget_configuration_id,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter *budgetConfigurationFilterPb `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateBudgetConfigurationBudgetFromPb(pb *updateBudgetConfigurationBudgetPb) (*UpdateBudgetConfigurationBudget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationBudget{}
	st.AccountId = pb.AccountId

	var alertConfigurationsField []AlertConfiguration
	for _, itemPb := range pb.AlertConfigurations {
		item, err := alertConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			alertConfigurationsField = append(alertConfigurationsField, *item)
		}
	}
	st.AlertConfigurations = alertConfigurationsField
	st.BudgetConfigurationId = pb.BudgetConfigurationId
	st.DisplayName = pb.DisplayName
	filterField, err := budgetConfigurationFilterFromPb(pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = filterField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateBudgetConfigurationBudgetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateBudgetConfigurationBudgetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateBudgetConfigurationRequest struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	// Wire name: 'budget'
	Budget UpdateBudgetConfigurationBudget
	// The Databricks budget configuration ID.
	// Wire name: 'budget_id'
	BudgetId string `tf:"-"`
}

func updateBudgetConfigurationRequestToPb(st *UpdateBudgetConfigurationRequest) (*updateBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetConfigurationRequestPb{}
	budgetPb, err := updateBudgetConfigurationBudgetToPb(&st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = *budgetPb
	}

	pb.BudgetId = st.BudgetId

	return pb, nil
}

func (st *UpdateBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateBudgetConfigurationRequestPb struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	Budget updateBudgetConfigurationBudgetPb `json:"budget"`
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" url:"-"`
}

func updateBudgetConfigurationRequestFromPb(pb *updateBudgetConfigurationRequestPb) (*UpdateBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationRequest{}
	budgetField, err := updateBudgetConfigurationBudgetFromPb(&pb.Budget)
	if err != nil {
		return nil, err
	}
	if budgetField != nil {
		st.Budget = *budgetField
	}
	st.BudgetId = pb.BudgetId

	return st, nil
}

type UpdateBudgetConfigurationResponse struct {
	// The updated budget.
	// Wire name: 'budget'
	Budget *BudgetConfiguration
}

func updateBudgetConfigurationResponseToPb(st *UpdateBudgetConfigurationResponse) (*updateBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetConfigurationResponsePb{}
	budgetPb, err := budgetConfigurationToPb(st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = budgetPb
	}

	return pb, nil
}

func (st *UpdateBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateBudgetConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateBudgetConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateBudgetConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateBudgetConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateBudgetConfigurationResponsePb struct {
	// The updated budget.
	Budget *budgetConfigurationPb `json:"budget,omitempty"`
}

func updateBudgetConfigurationResponseFromPb(pb *updateBudgetConfigurationResponsePb) (*UpdateBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationResponse{}
	budgetField, err := budgetConfigurationFromPb(pb.Budget)
	if err != nil {
		return nil, err
	}
	if budgetField != nil {
		st.Budget = budgetField
	}

	return st, nil
}

// Update a budget policy
type UpdateBudgetPolicyRequest struct {
	// DEPRECATED. This is redundant field as LimitConfig is part of the
	// BudgetPolicy
	// Wire name: 'limit_config'
	LimitConfig *LimitConfig `tf:"-"`
	// Contains the BudgetPolicy details.
	// Wire name: 'policy'
	Policy BudgetPolicy
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	// Wire name: 'policy_id'
	PolicyId string `tf:"-"`
}

func updateBudgetPolicyRequestToPb(st *UpdateBudgetPolicyRequest) (*updateBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateBudgetPolicyRequestPb{}
	limitConfigPb, err := limitConfigToPb(st.LimitConfig)
	if err != nil {
		return nil, err
	}
	if limitConfigPb != nil {
		pb.LimitConfig = limitConfigPb
	}

	policyPb, err := budgetPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}

	pb.PolicyId = st.PolicyId

	return pb, nil
}

func (st *UpdateBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateBudgetPolicyRequestPb struct {
	// DEPRECATED. This is redundant field as LimitConfig is part of the
	// BudgetPolicy
	LimitConfig *limitConfigPb `json:"-" url:"limit_config,omitempty"`
	// Contains the BudgetPolicy details.
	Policy budgetPolicyPb `json:"policy"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"-" url:"-"`
}

func updateBudgetPolicyRequestFromPb(pb *updateBudgetPolicyRequestPb) (*UpdateBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetPolicyRequest{}
	limitConfigField, err := limitConfigFromPb(pb.LimitConfig)
	if err != nil {
		return nil, err
	}
	if limitConfigField != nil {
		st.LimitConfig = limitConfigField
	}
	policyField, err := budgetPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.PolicyId = pb.PolicyId

	return st, nil
}

type UpdateLogDeliveryConfigurationStatusRequest struct {
	// Databricks log delivery configuration ID
	// Wire name: 'log_delivery_configuration_id'
	LogDeliveryConfigurationId string `tf:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	// Wire name: 'status'
	Status LogDeliveryConfigStatus
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

func (st *UpdateLogDeliveryConfigurationStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateLogDeliveryConfigurationStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateLogDeliveryConfigurationStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateLogDeliveryConfigurationStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateLogDeliveryConfigurationStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateLogDeliveryConfigurationStatusRequestPb struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId string `json:"-" url:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status"`
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

type UsageDashboardType string
type usageDashboardTypePb string

const UsageDashboardTypeUsageDashboardTypeGlobal UsageDashboardType = `USAGE_DASHBOARD_TYPE_GLOBAL`

const UsageDashboardTypeUsageDashboardTypeWorkspace UsageDashboardType = `USAGE_DASHBOARD_TYPE_WORKSPACE`

// String representation for [fmt.Print]
func (f *UsageDashboardType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UsageDashboardType) Set(v string) error {
	switch v {
	case `USAGE_DASHBOARD_TYPE_GLOBAL`, `USAGE_DASHBOARD_TYPE_WORKSPACE`:
		*f = UsageDashboardType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "USAGE_DASHBOARD_TYPE_GLOBAL", "USAGE_DASHBOARD_TYPE_WORKSPACE"`, v)
	}
}

// Type always returns UsageDashboardType to satisfy [pflag.Value] interface
func (f *UsageDashboardType) Type() string {
	return "UsageDashboardType"
}

func usageDashboardTypeToPb(st *UsageDashboardType) (*usageDashboardTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := usageDashboardTypePb(*st)
	return &pb, nil
}

func usageDashboardTypeFromPb(pb *usageDashboardTypePb) (*UsageDashboardType, error) {
	if pb == nil {
		return nil, nil
	}
	st := UsageDashboardType(*pb)
	return &st, nil
}

type WrappedCreateLogDeliveryConfiguration struct {

	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration *CreateLogDeliveryConfigurationParams
}

func wrappedCreateLogDeliveryConfigurationToPb(st *WrappedCreateLogDeliveryConfiguration) (*wrappedCreateLogDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &wrappedCreateLogDeliveryConfigurationPb{}
	logDeliveryConfigurationPb, err := createLogDeliveryConfigurationParamsToPb(st.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationPb != nil {
		pb.LogDeliveryConfiguration = logDeliveryConfigurationPb
	}

	return pb, nil
}

func (st *WrappedCreateLogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &wrappedCreateLogDeliveryConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := wrappedCreateLogDeliveryConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WrappedCreateLogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := wrappedCreateLogDeliveryConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type wrappedCreateLogDeliveryConfigurationPb struct {
	LogDeliveryConfiguration *createLogDeliveryConfigurationParamsPb `json:"log_delivery_configuration,omitempty"`
}

func wrappedCreateLogDeliveryConfigurationFromPb(pb *wrappedCreateLogDeliveryConfigurationPb) (*WrappedCreateLogDeliveryConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedCreateLogDeliveryConfiguration{}
	logDeliveryConfigurationField, err := createLogDeliveryConfigurationParamsFromPb(pb.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationField != nil {
		st.LogDeliveryConfiguration = logDeliveryConfigurationField
	}

	return st, nil
}

type WrappedLogDeliveryConfiguration struct {

	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration *LogDeliveryConfiguration
}

func wrappedLogDeliveryConfigurationToPb(st *WrappedLogDeliveryConfiguration) (*wrappedLogDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &wrappedLogDeliveryConfigurationPb{}
	logDeliveryConfigurationPb, err := logDeliveryConfigurationToPb(st.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationPb != nil {
		pb.LogDeliveryConfiguration = logDeliveryConfigurationPb
	}

	return pb, nil
}

func (st *WrappedLogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &wrappedLogDeliveryConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := wrappedLogDeliveryConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WrappedLogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := wrappedLogDeliveryConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type wrappedLogDeliveryConfigurationPb struct {
	LogDeliveryConfiguration *logDeliveryConfigurationPb `json:"log_delivery_configuration,omitempty"`
}

func wrappedLogDeliveryConfigurationFromPb(pb *wrappedLogDeliveryConfigurationPb) (*WrappedLogDeliveryConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedLogDeliveryConfiguration{}
	logDeliveryConfigurationField, err := logDeliveryConfigurationFromPb(pb.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationField != nil {
		st.LogDeliveryConfiguration = logDeliveryConfigurationField
	}

	return st, nil
}

type WrappedLogDeliveryConfigurations struct {

	// Wire name: 'log_delivery_configurations'
	LogDeliveryConfigurations []LogDeliveryConfiguration
}

func wrappedLogDeliveryConfigurationsToPb(st *WrappedLogDeliveryConfigurations) (*wrappedLogDeliveryConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &wrappedLogDeliveryConfigurationsPb{}

	var logDeliveryConfigurationsPb []logDeliveryConfigurationPb
	for _, item := range st.LogDeliveryConfigurations {
		itemPb, err := logDeliveryConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			logDeliveryConfigurationsPb = append(logDeliveryConfigurationsPb, *itemPb)
		}
	}
	pb.LogDeliveryConfigurations = logDeliveryConfigurationsPb

	return pb, nil
}

func (st *WrappedLogDeliveryConfigurations) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &wrappedLogDeliveryConfigurationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := wrappedLogDeliveryConfigurationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WrappedLogDeliveryConfigurations) MarshalJSON() ([]byte, error) {
	pb, err := wrappedLogDeliveryConfigurationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type wrappedLogDeliveryConfigurationsPb struct {
	LogDeliveryConfigurations []logDeliveryConfigurationPb `json:"log_delivery_configurations,omitempty"`
}

func wrappedLogDeliveryConfigurationsFromPb(pb *wrappedLogDeliveryConfigurationsPb) (*WrappedLogDeliveryConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedLogDeliveryConfigurations{}

	var logDeliveryConfigurationsField []LogDeliveryConfiguration
	for _, itemPb := range pb.LogDeliveryConfigurations {
		item, err := logDeliveryConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			logDeliveryConfigurationsField = append(logDeliveryConfigurationsField, *item)
		}
	}
	st.LogDeliveryConfigurations = logDeliveryConfigurationsField

	return st, nil
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
