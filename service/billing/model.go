// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/billing/billingpb"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/service/compute/computepb"
)

type ActionConfiguration struct {
	// Databricks action configuration ID.
	// Wire name: 'action_configuration_id'
	ActionConfigurationId string `json:"action_configuration_id,omitempty"`
	// The type of the action.
	// Wire name: 'action_type'
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target for the action. For example, an email address.
	// Wire name: 'target'
	Target          string   `json:"target,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ActionConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := ActionConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ActionConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.ActionConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ActionConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ActionConfigurationToPb(st *ActionConfiguration) (*billingpb.ActionConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.ActionConfigurationPb{}
	pb.ActionConfigurationId = st.ActionConfigurationId
	actionTypePb, err := ActionConfigurationTypeToPb(&st.ActionType)
	if err != nil {
		return nil, err
	}
	if actionTypePb != nil {
		pb.ActionType = *actionTypePb
	}
	pb.Target = st.Target

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ActionConfigurationFromPb(pb *billingpb.ActionConfigurationPb) (*ActionConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ActionConfiguration{}
	st.ActionConfigurationId = pb.ActionConfigurationId
	actionTypeField, err := ActionConfigurationTypeFromPb(&pb.ActionType)
	if err != nil {
		return nil, err
	}
	if actionTypeField != nil {
		st.ActionType = *actionTypeField
	}
	st.Target = pb.Target

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ActionConfigurationType string

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

// Values returns all possible values for ActionConfigurationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActionConfigurationType) Values() []ActionConfigurationType {
	return []ActionConfigurationType{
		ActionConfigurationTypeEmailNotification,
	}
}

// Type always returns ActionConfigurationType to satisfy [pflag.Value] interface
func (f *ActionConfigurationType) Type() string {
	return "ActionConfigurationType"
}

func ActionConfigurationTypeToPb(st *ActionConfigurationType) (*billingpb.ActionConfigurationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.ActionConfigurationTypePb(*st)
	return &pb, nil
}

func ActionConfigurationTypeFromPb(pb *billingpb.ActionConfigurationTypePb) (*ActionConfigurationType, error) {
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
	ActionConfigurations []ActionConfiguration `json:"action_configurations,omitempty"`
	// Databricks alert configuration ID.
	// Wire name: 'alert_configuration_id'
	AlertConfigurationId string `json:"alert_configuration_id,omitempty"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	// Wire name: 'quantity_threshold'
	QuantityThreshold string `json:"quantity_threshold,omitempty"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	// Wire name: 'quantity_type'
	QuantityType AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
	// The time window of usage data for the budget.
	// Wire name: 'time_period'
	TimePeriod AlertConfigurationTimePeriod `json:"time_period,omitempty"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	// Wire name: 'trigger_type'
	TriggerType     AlertConfigurationTriggerType `json:"trigger_type,omitempty"`
	ForceSendFields []string                      `json:"-" tf:"-"`
}

func (st AlertConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := AlertConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AlertConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.AlertConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AlertConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AlertConfigurationToPb(st *AlertConfiguration) (*billingpb.AlertConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.AlertConfigurationPb{}

	var actionConfigurationsPb []billingpb.ActionConfigurationPb
	for _, item := range st.ActionConfigurations {
		itemPb, err := ActionConfigurationToPb(&item)
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
	quantityTypePb, err := AlertConfigurationQuantityTypeToPb(&st.QuantityType)
	if err != nil {
		return nil, err
	}
	if quantityTypePb != nil {
		pb.QuantityType = *quantityTypePb
	}
	timePeriodPb, err := AlertConfigurationTimePeriodToPb(&st.TimePeriod)
	if err != nil {
		return nil, err
	}
	if timePeriodPb != nil {
		pb.TimePeriod = *timePeriodPb
	}
	triggerTypePb, err := AlertConfigurationTriggerTypeToPb(&st.TriggerType)
	if err != nil {
		return nil, err
	}
	if triggerTypePb != nil {
		pb.TriggerType = *triggerTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AlertConfigurationFromPb(pb *billingpb.AlertConfigurationPb) (*AlertConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AlertConfiguration{}

	var actionConfigurationsField []ActionConfiguration
	for _, itemPb := range pb.ActionConfigurations {
		item, err := ActionConfigurationFromPb(&itemPb)
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
	quantityTypeField, err := AlertConfigurationQuantityTypeFromPb(&pb.QuantityType)
	if err != nil {
		return nil, err
	}
	if quantityTypeField != nil {
		st.QuantityType = *quantityTypeField
	}
	timePeriodField, err := AlertConfigurationTimePeriodFromPb(&pb.TimePeriod)
	if err != nil {
		return nil, err
	}
	if timePeriodField != nil {
		st.TimePeriod = *timePeriodField
	}
	triggerTypeField, err := AlertConfigurationTriggerTypeFromPb(&pb.TriggerType)
	if err != nil {
		return nil, err
	}
	if triggerTypeField != nil {
		st.TriggerType = *triggerTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AlertConfigurationQuantityType string

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

// Values returns all possible values for AlertConfigurationQuantityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertConfigurationQuantityType) Values() []AlertConfigurationQuantityType {
	return []AlertConfigurationQuantityType{
		AlertConfigurationQuantityTypeListPriceDollarsUsd,
	}
}

// Type always returns AlertConfigurationQuantityType to satisfy [pflag.Value] interface
func (f *AlertConfigurationQuantityType) Type() string {
	return "AlertConfigurationQuantityType"
}

func AlertConfigurationQuantityTypeToPb(st *AlertConfigurationQuantityType) (*billingpb.AlertConfigurationQuantityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.AlertConfigurationQuantityTypePb(*st)
	return &pb, nil
}

func AlertConfigurationQuantityTypeFromPb(pb *billingpb.AlertConfigurationQuantityTypePb) (*AlertConfigurationQuantityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertConfigurationQuantityType(*pb)
	return &st, nil
}

type AlertConfigurationTimePeriod string

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

// Values returns all possible values for AlertConfigurationTimePeriod.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertConfigurationTimePeriod) Values() []AlertConfigurationTimePeriod {
	return []AlertConfigurationTimePeriod{
		AlertConfigurationTimePeriodMonth,
	}
}

// Type always returns AlertConfigurationTimePeriod to satisfy [pflag.Value] interface
func (f *AlertConfigurationTimePeriod) Type() string {
	return "AlertConfigurationTimePeriod"
}

func AlertConfigurationTimePeriodToPb(st *AlertConfigurationTimePeriod) (*billingpb.AlertConfigurationTimePeriodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.AlertConfigurationTimePeriodPb(*st)
	return &pb, nil
}

func AlertConfigurationTimePeriodFromPb(pb *billingpb.AlertConfigurationTimePeriodPb) (*AlertConfigurationTimePeriod, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertConfigurationTimePeriod(*pb)
	return &st, nil
}

type AlertConfigurationTriggerType string

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

// Values returns all possible values for AlertConfigurationTriggerType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AlertConfigurationTriggerType) Values() []AlertConfigurationTriggerType {
	return []AlertConfigurationTriggerType{
		AlertConfigurationTriggerTypeCumulativeSpendingExceeded,
	}
}

// Type always returns AlertConfigurationTriggerType to satisfy [pflag.Value] interface
func (f *AlertConfigurationTriggerType) Type() string {
	return "AlertConfigurationTriggerType"
}

func AlertConfigurationTriggerTypeToPb(st *AlertConfigurationTriggerType) (*billingpb.AlertConfigurationTriggerTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.AlertConfigurationTriggerTypePb(*st)
	return &pb, nil
}

func AlertConfigurationTriggerTypeFromPb(pb *billingpb.AlertConfigurationTriggerTypePb) (*AlertConfigurationTriggerType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AlertConfigurationTriggerType(*pb)
	return &st, nil
}

type BudgetConfiguration struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	// Wire name: 'alert_configurations'
	AlertConfigurations []AlertConfiguration `json:"alert_configurations,omitempty"`
	// Databricks budget configuration ID.
	// Wire name: 'budget_configuration_id'
	BudgetConfigurationId string `json:"budget_configuration_id,omitempty"`
	// Creation time of this budget configuration.
	// Wire name: 'create_time'
	CreateTime int64 `json:"create_time,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	// Wire name: 'filter'
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`
	// Update time of this budget configuration.
	// Wire name: 'update_time'
	UpdateTime      int64    `json:"update_time,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st BudgetConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := BudgetConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BudgetConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.BudgetConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BudgetConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BudgetConfigurationToPb(st *BudgetConfiguration) (*billingpb.BudgetConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.BudgetConfigurationPb{}
	pb.AccountId = st.AccountId

	var alertConfigurationsPb []billingpb.AlertConfigurationPb
	for _, item := range st.AlertConfigurations {
		itemPb, err := AlertConfigurationToPb(&item)
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
	filterPb, err := BudgetConfigurationFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}
	pb.UpdateTime = st.UpdateTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func BudgetConfigurationFromPb(pb *billingpb.BudgetConfigurationPb) (*BudgetConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfiguration{}
	st.AccountId = pb.AccountId

	var alertConfigurationsField []AlertConfiguration
	for _, itemPb := range pb.AlertConfigurations {
		item, err := AlertConfigurationFromPb(&itemPb)
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
	filterField, err := BudgetConfigurationFilterFromPb(pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = filterField
	}
	st.UpdateTime = pb.UpdateTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type BudgetConfigurationFilter struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	// Wire name: 'tags'
	Tags []BudgetConfigurationFilterTagClause `json:"tags,omitempty"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	// Wire name: 'workspace_id'
	WorkspaceId *BudgetConfigurationFilterWorkspaceIdClause `json:"workspace_id,omitempty"`
}

func (st BudgetConfigurationFilter) MarshalJSON() ([]byte, error) {
	pb, err := BudgetConfigurationFilterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BudgetConfigurationFilter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.BudgetConfigurationFilterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BudgetConfigurationFilterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BudgetConfigurationFilterToPb(st *BudgetConfigurationFilter) (*billingpb.BudgetConfigurationFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.BudgetConfigurationFilterPb{}

	var tagsPb []billingpb.BudgetConfigurationFilterTagClausePb
	for _, item := range st.Tags {
		itemPb, err := BudgetConfigurationFilterTagClauseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb
	workspaceIdPb, err := BudgetConfigurationFilterWorkspaceIdClauseToPb(st.WorkspaceId)
	if err != nil {
		return nil, err
	}
	if workspaceIdPb != nil {
		pb.WorkspaceId = workspaceIdPb
	}

	return pb, nil
}

func BudgetConfigurationFilterFromPb(pb *billingpb.BudgetConfigurationFilterPb) (*BudgetConfigurationFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilter{}

	var tagsField []BudgetConfigurationFilterTagClause
	for _, itemPb := range pb.Tags {
		item, err := BudgetConfigurationFilterTagClauseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	workspaceIdField, err := BudgetConfigurationFilterWorkspaceIdClauseFromPb(pb.WorkspaceId)
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
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	// Wire name: 'values'
	Values []string `json:"values,omitempty"`
}

func (st BudgetConfigurationFilterClause) MarshalJSON() ([]byte, error) {
	pb, err := BudgetConfigurationFilterClauseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BudgetConfigurationFilterClause) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.BudgetConfigurationFilterClausePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BudgetConfigurationFilterClauseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BudgetConfigurationFilterClauseToPb(st *BudgetConfigurationFilterClause) (*billingpb.BudgetConfigurationFilterClausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.BudgetConfigurationFilterClausePb{}
	operatorPb, err := BudgetConfigurationFilterOperatorToPb(&st.Operator)
	if err != nil {
		return nil, err
	}
	if operatorPb != nil {
		pb.Operator = *operatorPb
	}
	pb.Values = st.Values

	return pb, nil
}

func BudgetConfigurationFilterClauseFromPb(pb *billingpb.BudgetConfigurationFilterClausePb) (*BudgetConfigurationFilterClause, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilterClause{}
	operatorField, err := BudgetConfigurationFilterOperatorFromPb(&pb.Operator)
	if err != nil {
		return nil, err
	}
	if operatorField != nil {
		st.Operator = *operatorField
	}
	st.Values = pb.Values

	return st, nil
}

type BudgetConfigurationFilterOperator string

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

// Values returns all possible values for BudgetConfigurationFilterOperator.
//
// There is no guarantee on the order of the values in the slice.
func (f *BudgetConfigurationFilterOperator) Values() []BudgetConfigurationFilterOperator {
	return []BudgetConfigurationFilterOperator{
		BudgetConfigurationFilterOperatorIn,
	}
}

// Type always returns BudgetConfigurationFilterOperator to satisfy [pflag.Value] interface
func (f *BudgetConfigurationFilterOperator) Type() string {
	return "BudgetConfigurationFilterOperator"
}

func BudgetConfigurationFilterOperatorToPb(st *BudgetConfigurationFilterOperator) (*billingpb.BudgetConfigurationFilterOperatorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.BudgetConfigurationFilterOperatorPb(*st)
	return &pb, nil
}

func BudgetConfigurationFilterOperatorFromPb(pb *billingpb.BudgetConfigurationFilterOperatorPb) (*BudgetConfigurationFilterOperator, error) {
	if pb == nil {
		return nil, nil
	}
	st := BudgetConfigurationFilterOperator(*pb)
	return &st, nil
}

type BudgetConfigurationFilterTagClause struct {

	// Wire name: 'key'
	Key string `json:"key,omitempty"`

	// Wire name: 'value'
	Value           *BudgetConfigurationFilterClause `json:"value,omitempty"`
	ForceSendFields []string                         `json:"-" tf:"-"`
}

func (st BudgetConfigurationFilterTagClause) MarshalJSON() ([]byte, error) {
	pb, err := BudgetConfigurationFilterTagClauseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BudgetConfigurationFilterTagClause) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.BudgetConfigurationFilterTagClausePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BudgetConfigurationFilterTagClauseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BudgetConfigurationFilterTagClauseToPb(st *BudgetConfigurationFilterTagClause) (*billingpb.BudgetConfigurationFilterTagClausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.BudgetConfigurationFilterTagClausePb{}
	pb.Key = st.Key
	valuePb, err := BudgetConfigurationFilterClauseToPb(st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = valuePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func BudgetConfigurationFilterTagClauseFromPb(pb *billingpb.BudgetConfigurationFilterTagClausePb) (*BudgetConfigurationFilterTagClause, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilterTagClause{}
	st.Key = pb.Key
	valueField, err := BudgetConfigurationFilterClauseFromPb(pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = valueField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type BudgetConfigurationFilterWorkspaceIdClause struct {

	// Wire name: 'operator'
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	// Wire name: 'values'
	Values []int64 `json:"values,omitempty"`
}

func (st BudgetConfigurationFilterWorkspaceIdClause) MarshalJSON() ([]byte, error) {
	pb, err := BudgetConfigurationFilterWorkspaceIdClauseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BudgetConfigurationFilterWorkspaceIdClause) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.BudgetConfigurationFilterWorkspaceIdClausePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BudgetConfigurationFilterWorkspaceIdClauseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BudgetConfigurationFilterWorkspaceIdClauseToPb(st *BudgetConfigurationFilterWorkspaceIdClause) (*billingpb.BudgetConfigurationFilterWorkspaceIdClausePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.BudgetConfigurationFilterWorkspaceIdClausePb{}
	operatorPb, err := BudgetConfigurationFilterOperatorToPb(&st.Operator)
	if err != nil {
		return nil, err
	}
	if operatorPb != nil {
		pb.Operator = *operatorPb
	}
	pb.Values = st.Values

	return pb, nil
}

func BudgetConfigurationFilterWorkspaceIdClauseFromPb(pb *billingpb.BudgetConfigurationFilterWorkspaceIdClausePb) (*BudgetConfigurationFilterWorkspaceIdClause, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BudgetConfigurationFilterWorkspaceIdClause{}
	operatorField, err := BudgetConfigurationFilterOperatorFromPb(&pb.Operator)
	if err != nil {
		return nil, err
	}
	if operatorField != nil {
		st.Operator = *operatorField
	}
	st.Values = pb.Values

	return st, nil
}

// Contains the BudgetPolicy details.
type BudgetPolicy struct {
	// List of workspaces that this budget policy will be exclusively bound to.
	// An empty binding implies that this budget policy is open to any workspace
	// in the account.
	// Wire name: 'binding_workspace_ids'
	BindingWorkspaceIds []int64 `json:"binding_workspace_ids,omitempty"`
	// A list of tags defined by the customer. At most 20 entries are allowed
	// per policy.
	// Wire name: 'custom_tags'
	CustomTags []compute.CustomPolicyTag `json:"custom_tags,omitempty"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// The name of the policy. - Must be unique among active policies. - Can
	// contain only characters from the ISO 8859-1 (latin1) set. - Can't start
	// with reserved keywords such as `databricks:default-policy`.
	// Wire name: 'policy_name'
	PolicyName      string   `json:"policy_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st BudgetPolicy) MarshalJSON() ([]byte, error) {
	pb, err := BudgetPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BudgetPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.BudgetPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BudgetPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BudgetPolicyToPb(st *BudgetPolicy) (*billingpb.BudgetPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.BudgetPolicyPb{}
	pb.BindingWorkspaceIds = st.BindingWorkspaceIds

	var customTagsPb []computepb.CustomPolicyTagPb
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func BudgetPolicyFromPb(pb *billingpb.BudgetPolicyPb) (*BudgetPolicy, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	// Wire name: 'dashboard_type'
	DashboardType UsageDashboardType `json:"dashboard_type,omitempty"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	// Wire name: 'workspace_id'
	WorkspaceId     int64    `json:"workspace_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateBillingUsageDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBillingUsageDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBillingUsageDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBillingUsageDashboardRequestToPb(st *CreateBillingUsageDashboardRequest) (*billingpb.CreateBillingUsageDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBillingUsageDashboardRequestPb{}
	dashboardTypePb, err := UsageDashboardTypeToPb(&st.DashboardType)
	if err != nil {
		return nil, err
	}
	if dashboardTypePb != nil {
		pb.DashboardType = *dashboardTypePb
	}
	pb.WorkspaceId = st.WorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateBillingUsageDashboardRequestFromPb(pb *billingpb.CreateBillingUsageDashboardRequestPb) (*CreateBillingUsageDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBillingUsageDashboardRequest{}
	dashboardTypeField, err := UsageDashboardTypeFromPb(&pb.DashboardType)
	if err != nil {
		return nil, err
	}
	if dashboardTypeField != nil {
		st.DashboardType = *dashboardTypeField
	}
	st.WorkspaceId = pb.WorkspaceId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	// Wire name: 'dashboard_id'
	DashboardId     string   `json:"dashboard_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateBillingUsageDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBillingUsageDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBillingUsageDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBillingUsageDashboardResponseToPb(st *CreateBillingUsageDashboardResponse) (*billingpb.CreateBillingUsageDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBillingUsageDashboardResponsePb{}
	pb.DashboardId = st.DashboardId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateBillingUsageDashboardResponseFromPb(pb *billingpb.CreateBillingUsageDashboardResponsePb) (*CreateBillingUsageDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBillingUsageDashboardResponse{}
	st.DashboardId = pb.DashboardId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateBudgetConfigurationBudget struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	// Wire name: 'alert_configurations'
	AlertConfigurations []CreateBudgetConfigurationBudgetAlertConfigurations `json:"alert_configurations,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	// Wire name: 'filter'
	Filter          *BudgetConfigurationFilter `json:"filter,omitempty"`
	ForceSendFields []string                   `json:"-" tf:"-"`
}

func (st CreateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	pb, err := CreateBudgetConfigurationBudgetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBudgetConfigurationBudgetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBudgetConfigurationBudgetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBudgetConfigurationBudgetToPb(st *CreateBudgetConfigurationBudget) (*billingpb.CreateBudgetConfigurationBudgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBudgetConfigurationBudgetPb{}
	pb.AccountId = st.AccountId

	var alertConfigurationsPb []billingpb.CreateBudgetConfigurationBudgetAlertConfigurationsPb
	for _, item := range st.AlertConfigurations {
		itemPb, err := CreateBudgetConfigurationBudgetAlertConfigurationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			alertConfigurationsPb = append(alertConfigurationsPb, *itemPb)
		}
	}
	pb.AlertConfigurations = alertConfigurationsPb
	pb.DisplayName = st.DisplayName
	filterPb, err := BudgetConfigurationFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateBudgetConfigurationBudgetFromPb(pb *billingpb.CreateBudgetConfigurationBudgetPb) (*CreateBudgetConfigurationBudget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudget{}
	st.AccountId = pb.AccountId

	var alertConfigurationsField []CreateBudgetConfigurationBudgetAlertConfigurations
	for _, itemPb := range pb.AlertConfigurations {
		item, err := CreateBudgetConfigurationBudgetAlertConfigurationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			alertConfigurationsField = append(alertConfigurationsField, *item)
		}
	}
	st.AlertConfigurations = alertConfigurationsField
	st.DisplayName = pb.DisplayName
	filterField, err := BudgetConfigurationFilterFromPb(pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = filterField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateBudgetConfigurationBudgetActionConfigurations struct {
	// The type of the action.
	// Wire name: 'action_type'
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target for the action. For example, an email address.
	// Wire name: 'target'
	Target          string   `json:"target,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateBudgetConfigurationBudgetActionConfigurations) MarshalJSON() ([]byte, error) {
	pb, err := CreateBudgetConfigurationBudgetActionConfigurationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBudgetConfigurationBudgetActionConfigurations) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBudgetConfigurationBudgetActionConfigurationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBudgetConfigurationBudgetActionConfigurationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBudgetConfigurationBudgetActionConfigurationsToPb(st *CreateBudgetConfigurationBudgetActionConfigurations) (*billingpb.CreateBudgetConfigurationBudgetActionConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBudgetConfigurationBudgetActionConfigurationsPb{}
	actionTypePb, err := ActionConfigurationTypeToPb(&st.ActionType)
	if err != nil {
		return nil, err
	}
	if actionTypePb != nil {
		pb.ActionType = *actionTypePb
	}
	pb.Target = st.Target

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateBudgetConfigurationBudgetActionConfigurationsFromPb(pb *billingpb.CreateBudgetConfigurationBudgetActionConfigurationsPb) (*CreateBudgetConfigurationBudgetActionConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudgetActionConfigurations{}
	actionTypeField, err := ActionConfigurationTypeFromPb(&pb.ActionType)
	if err != nil {
		return nil, err
	}
	if actionTypeField != nil {
		st.ActionType = *actionTypeField
	}
	st.Target = pb.Target

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateBudgetConfigurationBudgetAlertConfigurations struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	// Wire name: 'action_configurations'
	ActionConfigurations []CreateBudgetConfigurationBudgetActionConfigurations `json:"action_configurations,omitempty"`
	// The threshold for the budget alert to determine if it is in a triggered
	// state. The number is evaluated based on `quantity_type`.
	// Wire name: 'quantity_threshold'
	QuantityThreshold string `json:"quantity_threshold,omitempty"`
	// The way to calculate cost for this budget alert. This is what
	// `quantity_threshold` is measured in.
	// Wire name: 'quantity_type'
	QuantityType AlertConfigurationQuantityType `json:"quantity_type,omitempty"`
	// The time window of usage data for the budget.
	// Wire name: 'time_period'
	TimePeriod AlertConfigurationTimePeriod `json:"time_period,omitempty"`
	// The evaluation method to determine when this budget alert is in a
	// triggered state.
	// Wire name: 'trigger_type'
	TriggerType     AlertConfigurationTriggerType `json:"trigger_type,omitempty"`
	ForceSendFields []string                      `json:"-" tf:"-"`
}

func (st CreateBudgetConfigurationBudgetAlertConfigurations) MarshalJSON() ([]byte, error) {
	pb, err := CreateBudgetConfigurationBudgetAlertConfigurationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBudgetConfigurationBudgetAlertConfigurations) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBudgetConfigurationBudgetAlertConfigurationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBudgetConfigurationBudgetAlertConfigurationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBudgetConfigurationBudgetAlertConfigurationsToPb(st *CreateBudgetConfigurationBudgetAlertConfigurations) (*billingpb.CreateBudgetConfigurationBudgetAlertConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBudgetConfigurationBudgetAlertConfigurationsPb{}

	var actionConfigurationsPb []billingpb.CreateBudgetConfigurationBudgetActionConfigurationsPb
	for _, item := range st.ActionConfigurations {
		itemPb, err := CreateBudgetConfigurationBudgetActionConfigurationsToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			actionConfigurationsPb = append(actionConfigurationsPb, *itemPb)
		}
	}
	pb.ActionConfigurations = actionConfigurationsPb
	pb.QuantityThreshold = st.QuantityThreshold
	quantityTypePb, err := AlertConfigurationQuantityTypeToPb(&st.QuantityType)
	if err != nil {
		return nil, err
	}
	if quantityTypePb != nil {
		pb.QuantityType = *quantityTypePb
	}
	timePeriodPb, err := AlertConfigurationTimePeriodToPb(&st.TimePeriod)
	if err != nil {
		return nil, err
	}
	if timePeriodPb != nil {
		pb.TimePeriod = *timePeriodPb
	}
	triggerTypePb, err := AlertConfigurationTriggerTypeToPb(&st.TriggerType)
	if err != nil {
		return nil, err
	}
	if triggerTypePb != nil {
		pb.TriggerType = *triggerTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateBudgetConfigurationBudgetAlertConfigurationsFromPb(pb *billingpb.CreateBudgetConfigurationBudgetAlertConfigurationsPb) (*CreateBudgetConfigurationBudgetAlertConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationBudgetAlertConfigurations{}

	var actionConfigurationsField []CreateBudgetConfigurationBudgetActionConfigurations
	for _, itemPb := range pb.ActionConfigurations {
		item, err := CreateBudgetConfigurationBudgetActionConfigurationsFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			actionConfigurationsField = append(actionConfigurationsField, *item)
		}
	}
	st.ActionConfigurations = actionConfigurationsField
	st.QuantityThreshold = pb.QuantityThreshold
	quantityTypeField, err := AlertConfigurationQuantityTypeFromPb(&pb.QuantityType)
	if err != nil {
		return nil, err
	}
	if quantityTypeField != nil {
		st.QuantityType = *quantityTypeField
	}
	timePeriodField, err := AlertConfigurationTimePeriodFromPb(&pb.TimePeriod)
	if err != nil {
		return nil, err
	}
	if timePeriodField != nil {
		st.TimePeriod = *timePeriodField
	}
	triggerTypeField, err := AlertConfigurationTriggerTypeFromPb(&pb.TriggerType)
	if err != nil {
		return nil, err
	}
	if triggerTypeField != nil {
		st.TriggerType = *triggerTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateBudgetConfigurationRequest struct {
	// Properties of the new budget configuration.
	// Wire name: 'budget'
	Budget CreateBudgetConfigurationBudget `json:"budget"`
}

func (st CreateBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBudgetConfigurationRequestToPb(st *CreateBudgetConfigurationRequest) (*billingpb.CreateBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBudgetConfigurationRequestPb{}
	budgetPb, err := CreateBudgetConfigurationBudgetToPb(&st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = *budgetPb
	}

	return pb, nil
}

func CreateBudgetConfigurationRequestFromPb(pb *billingpb.CreateBudgetConfigurationRequestPb) (*CreateBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationRequest{}
	budgetField, err := CreateBudgetConfigurationBudgetFromPb(&pb.Budget)
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
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func (st CreateBudgetConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateBudgetConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBudgetConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBudgetConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBudgetConfigurationResponseToPb(st *CreateBudgetConfigurationResponse) (*billingpb.CreateBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBudgetConfigurationResponsePb{}
	budgetPb, err := BudgetConfigurationToPb(st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = budgetPb
	}

	return pb, nil
}

func CreateBudgetConfigurationResponseFromPb(pb *billingpb.CreateBudgetConfigurationResponsePb) (*CreateBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetConfigurationResponse{}
	budgetField, err := BudgetConfigurationFromPb(pb.Budget)
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
	Policy *BudgetPolicy `json:"policy,omitempty"`
	// A unique identifier for this request. Restricted to 36 ASCII characters.
	// A random UUID is recommended. This request is only idempotent if a
	// `request_id` is provided.
	// Wire name: 'request_id'
	RequestId       string   `json:"request_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateBudgetPolicyRequestToPb(st *CreateBudgetPolicyRequest) (*billingpb.CreateBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateBudgetPolicyRequestPb{}
	policyPb, err := BudgetPolicyToPb(st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = policyPb
	}
	pb.RequestId = st.RequestId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateBudgetPolicyRequestFromPb(pb *billingpb.CreateBudgetPolicyRequestPb) (*CreateBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateBudgetPolicyRequest{}
	policyField, err := BudgetPolicyFromPb(pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = policyField
	}
	st.RequestId = pb.RequestId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// * Log Delivery Configuration
type CreateLogDeliveryConfigurationParams struct {
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	// Wire name: 'config_name'
	ConfigName string `json:"config_name,omitempty"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'credentials_id'
	CredentialsId string `json:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	// Wire name: 'delivery_path_prefix'
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if log_type is BILLABLE_USAGE. This is the
	// optional start month and year for delivery, specified in YYYY-MM format.
	// Defaults to current year and month. BILLABLE_USAGE logs are not available
	// for usage before March 2019 (2019-03).
	// Wire name: 'delivery_start_time'
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// Log delivery type. Supported values are: * `BILLABLE_USAGE` — Configure
	// [billable usage log delivery]. For the CSV schema, see the [View billable
	// usage]. * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON
	// schema, see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'log_type'
	LogType LogType `json:"log_type"`
	// The file type of log delivery. * If `log_type` is `BILLABLE_USAGE`, this
	// value must be `CSV`. Only the CSV (comma-separated values) format is
	// supported. For the schema, see the [View billable usage] * If `log_type`
	// is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON (JavaScript
	// Object Notation) format is supported. For the schema, see the
	// [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// Wire name: 'output_format'
	OutputFormat OutputFormat `json:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	// Wire name: 'status'
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'storage_configuration_id'
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
	// Wire name: 'workspace_ids_filter'
	WorkspaceIdsFilter []int64  `json:"workspace_ids_filter,omitempty"`
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st CreateLogDeliveryConfigurationParams) MarshalJSON() ([]byte, error) {
	pb, err := CreateLogDeliveryConfigurationParamsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateLogDeliveryConfigurationParams) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.CreateLogDeliveryConfigurationParamsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateLogDeliveryConfigurationParamsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateLogDeliveryConfigurationParamsToPb(st *CreateLogDeliveryConfigurationParams) (*billingpb.CreateLogDeliveryConfigurationParamsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.CreateLogDeliveryConfigurationParamsPb{}
	pb.ConfigName = st.ConfigName
	pb.CredentialsId = st.CredentialsId
	pb.DeliveryPathPrefix = st.DeliveryPathPrefix
	pb.DeliveryStartTime = st.DeliveryStartTime
	logTypePb, err := LogTypeToPb(&st.LogType)
	if err != nil {
		return nil, err
	}
	if logTypePb != nil {
		pb.LogType = *logTypePb
	}
	outputFormatPb, err := OutputFormatToPb(&st.OutputFormat)
	if err != nil {
		return nil, err
	}
	if outputFormatPb != nil {
		pb.OutputFormat = *outputFormatPb
	}
	statusPb, err := LogDeliveryConfigStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StorageConfigurationId = st.StorageConfigurationId
	pb.WorkspaceIdsFilter = st.WorkspaceIdsFilter

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateLogDeliveryConfigurationParamsFromPb(pb *billingpb.CreateLogDeliveryConfigurationParamsPb) (*CreateLogDeliveryConfigurationParams, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateLogDeliveryConfigurationParams{}
	st.ConfigName = pb.ConfigName
	st.CredentialsId = pb.CredentialsId
	st.DeliveryPathPrefix = pb.DeliveryPathPrefix
	st.DeliveryStartTime = pb.DeliveryStartTime
	logTypeField, err := LogTypeFromPb(&pb.LogType)
	if err != nil {
		return nil, err
	}
	if logTypeField != nil {
		st.LogType = *logTypeField
	}
	outputFormatField, err := OutputFormatFromPb(&pb.OutputFormat)
	if err != nil {
		return nil, err
	}
	if outputFormatField != nil {
		st.OutputFormat = *outputFormatField
	}
	statusField, err := LogDeliveryConfigStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.WorkspaceIdsFilter = pb.WorkspaceIdsFilter

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteBudgetConfigurationRequest struct {
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" tf:"-"`
}

func (st DeleteBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.DeleteBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteBudgetConfigurationRequestToPb(st *DeleteBudgetConfigurationRequest) (*billingpb.DeleteBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.DeleteBudgetConfigurationRequestPb{}
	pb.BudgetId = st.BudgetId

	return pb, nil
}

func DeleteBudgetConfigurationRequestFromPb(pb *billingpb.DeleteBudgetConfigurationRequestPb) (*DeleteBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteBudgetConfigurationRequest{}
	st.BudgetId = pb.BudgetId

	return st, nil
}

type DeleteBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" tf:"-"`
}

func (st DeleteBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.DeleteBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteBudgetPolicyRequestToPb(st *DeleteBudgetPolicyRequest) (*billingpb.DeleteBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.DeleteBudgetPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func DeleteBudgetPolicyRequestFromPb(pb *billingpb.DeleteBudgetPolicyRequestPb) (*DeleteBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteBudgetPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

// * The status string for log delivery. Possible values are: `CREATED`: There
// were no log delivery attempts since the config was created. `SUCCEEDED`: The
// latest attempt of log delivery has succeeded completely. `USER_FAILURE`: The
// latest attempt of log delivery failed because of misconfiguration of customer
// provided permissions on role or storage. `SYSTEM_FAILURE`: The latest attempt
// of log delivery failed because of an Databricks internal error. Contact
// support if it doesn't go away soon. `NOT_FOUND`: The log delivery status as
// the configuration has been disabled since the release of this feature or
// there are no workspaces in the account.
type DeliveryStatus string

const DeliveryStatusCreated DeliveryStatus = `CREATED`

const DeliveryStatusNotFound DeliveryStatus = `NOT_FOUND`

const DeliveryStatusSucceeded DeliveryStatus = `SUCCEEDED`

const DeliveryStatusSystemFailure DeliveryStatus = `SYSTEM_FAILURE`

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

// Values returns all possible values for DeliveryStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *DeliveryStatus) Values() []DeliveryStatus {
	return []DeliveryStatus{
		DeliveryStatusCreated,
		DeliveryStatusNotFound,
		DeliveryStatusSucceeded,
		DeliveryStatusSystemFailure,
		DeliveryStatusUserFailure,
	}
}

// Type always returns DeliveryStatus to satisfy [pflag.Value] interface
func (f *DeliveryStatus) Type() string {
	return "DeliveryStatus"
}

func DeliveryStatusToPb(st *DeliveryStatus) (*billingpb.DeliveryStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.DeliveryStatusPb(*st)
	return &pb, nil
}

func DeliveryStatusFromPb(pb *billingpb.DeliveryStatusPb) (*DeliveryStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeliveryStatus(*pb)
	return &st, nil
}

type DownloadRequest struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth string `json:"-" tf:"-"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData bool `json:"-" tf:"-"`
	// Format specification for month in the format `YYYY-MM`. This is used to
	// specify billable usage `start_month` and `end_month` properties.
	// **Note**: Billable usage logs are unavailable before March 2019
	// (`2019-03`).
	StartMonth      string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := DownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.DownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DownloadRequestToPb(st *DownloadRequest) (*billingpb.DownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.DownloadRequestPb{}
	pb.EndMonth = st.EndMonth
	pb.PersonalData = st.PersonalData
	pb.StartMonth = st.StartMonth

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DownloadRequestFromPb(pb *billingpb.DownloadRequestPb) (*DownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadRequest{}
	st.EndMonth = pb.EndMonth
	st.PersonalData = pb.PersonalData
	st.StartMonth = pb.StartMonth

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DownloadResponse struct {
	Contents io.ReadCloser `json:"-" tf:"-"`
}

func (st DownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := DownloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.DownloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DownloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DownloadResponseToPb(st *DownloadResponse) (*billingpb.DownloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.DownloadResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

func DownloadResponseFromPb(pb *billingpb.DownloadResponsePb) (*DownloadResponse, error) {
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
	CreatorUserId int64 `json:"creator_user_id,omitempty"`
	// The policy creator user name to be filtered on. If unspecified, all
	// policies will be returned.
	// Wire name: 'creator_user_name'
	CreatorUserName string `json:"creator_user_name,omitempty"`
	// The partial name of policies to be filtered on. If unspecified, all
	// policies will be returned.
	// Wire name: 'policy_name'
	PolicyName      string   `json:"policy_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Filter) MarshalJSON() ([]byte, error) {
	pb, err := FilterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Filter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.FilterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FilterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FilterToPb(st *Filter) (*billingpb.FilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.FilterPb{}
	pb.CreatorUserId = st.CreatorUserId
	pb.CreatorUserName = st.CreatorUserName
	pb.PolicyName = st.PolicyName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FilterFromPb(pb *billingpb.FilterPb) (*Filter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Filter{}
	st.CreatorUserId = pb.CreatorUserId
	st.CreatorUserName = pb.CreatorUserName
	st.PolicyName = pb.PolicyName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType UsageDashboardType `json:"-" tf:"-"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId     int64    `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetBillingUsageDashboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.GetBillingUsageDashboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetBillingUsageDashboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetBillingUsageDashboardRequestToPb(st *GetBillingUsageDashboardRequest) (*billingpb.GetBillingUsageDashboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.GetBillingUsageDashboardRequestPb{}
	dashboardTypePb, err := UsageDashboardTypeToPb(&st.DashboardType)
	if err != nil {
		return nil, err
	}
	if dashboardTypePb != nil {
		pb.DashboardType = *dashboardTypePb
	}
	pb.WorkspaceId = st.WorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetBillingUsageDashboardRequestFromPb(pb *billingpb.GetBillingUsageDashboardRequestPb) (*GetBillingUsageDashboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBillingUsageDashboardRequest{}
	dashboardTypeField, err := UsageDashboardTypeFromPb(&pb.DashboardType)
	if err != nil {
		return nil, err
	}
	if dashboardTypeField != nil {
		st.DashboardType = *dashboardTypeField
	}
	st.WorkspaceId = pb.WorkspaceId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The URL of the usage dashboard.
	// Wire name: 'dashboard_url'
	DashboardUrl    string   `json:"dashboard_url,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetBillingUsageDashboardResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.GetBillingUsageDashboardResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetBillingUsageDashboardResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetBillingUsageDashboardResponseToPb(st *GetBillingUsageDashboardResponse) (*billingpb.GetBillingUsageDashboardResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.GetBillingUsageDashboardResponsePb{}
	pb.DashboardId = st.DashboardId
	pb.DashboardUrl = st.DashboardUrl

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetBillingUsageDashboardResponseFromPb(pb *billingpb.GetBillingUsageDashboardResponsePb) (*GetBillingUsageDashboardResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBillingUsageDashboardResponse{}
	st.DashboardId = pb.DashboardId
	st.DashboardUrl = pb.DashboardUrl

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetBudgetConfigurationRequest struct {
	// The budget configuration ID
	BudgetId string `json:"-" tf:"-"`
}

func (st GetBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.GetBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetBudgetConfigurationRequestToPb(st *GetBudgetConfigurationRequest) (*billingpb.GetBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.GetBudgetConfigurationRequestPb{}
	pb.BudgetId = st.BudgetId

	return pb, nil
}

func GetBudgetConfigurationRequestFromPb(pb *billingpb.GetBudgetConfigurationRequestPb) (*GetBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBudgetConfigurationRequest{}
	st.BudgetId = pb.BudgetId

	return st, nil
}

type GetBudgetConfigurationResponse struct {

	// Wire name: 'budget'
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func (st GetBudgetConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetBudgetConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.GetBudgetConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetBudgetConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetBudgetConfigurationResponseToPb(st *GetBudgetConfigurationResponse) (*billingpb.GetBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.GetBudgetConfigurationResponsePb{}
	budgetPb, err := BudgetConfigurationToPb(st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = budgetPb
	}

	return pb, nil
}

func GetBudgetConfigurationResponseFromPb(pb *billingpb.GetBudgetConfigurationResponsePb) (*GetBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBudgetConfigurationResponse{}
	budgetField, err := BudgetConfigurationFromPb(pb.Budget)
	if err != nil {
		return nil, err
	}
	if budgetField != nil {
		st.Budget = budgetField
	}

	return st, nil
}

type GetBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" tf:"-"`
}

func (st GetBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.GetBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetBudgetPolicyRequestToPb(st *GetBudgetPolicyRequest) (*billingpb.GetBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.GetBudgetPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func GetBudgetPolicyRequestFromPb(pb *billingpb.GetBudgetPolicyRequestPb) (*GetBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetBudgetPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

type GetLogDeliveryConfigurationResponse struct {
	// The fetched log delivery configuration
	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

func (st GetLogDeliveryConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetLogDeliveryConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLogDeliveryConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.GetLogDeliveryConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLogDeliveryConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLogDeliveryConfigurationResponseToPb(st *GetLogDeliveryConfigurationResponse) (*billingpb.GetLogDeliveryConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.GetLogDeliveryConfigurationResponsePb{}
	logDeliveryConfigurationPb, err := LogDeliveryConfigurationToPb(st.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationPb != nil {
		pb.LogDeliveryConfiguration = logDeliveryConfigurationPb
	}

	return pb, nil
}

func GetLogDeliveryConfigurationResponseFromPb(pb *billingpb.GetLogDeliveryConfigurationResponsePb) (*GetLogDeliveryConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLogDeliveryConfigurationResponse{}
	logDeliveryConfigurationField, err := LogDeliveryConfigurationFromPb(pb.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationField != nil {
		st.LogDeliveryConfiguration = logDeliveryConfigurationField
	}

	return st, nil
}

type GetLogDeliveryRequest struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId string `json:"-" tf:"-"`
}

func (st GetLogDeliveryRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetLogDeliveryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.GetLogDeliveryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLogDeliveryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLogDeliveryRequestToPb(st *GetLogDeliveryRequest) (*billingpb.GetLogDeliveryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.GetLogDeliveryRequestPb{}
	pb.LogDeliveryConfigurationId = st.LogDeliveryConfigurationId

	return pb, nil
}

func GetLogDeliveryRequestFromPb(pb *billingpb.GetLogDeliveryRequestPb) (*GetLogDeliveryRequest, error) {
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

func (st LimitConfig) MarshalJSON() ([]byte, error) {
	pb, err := LimitConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LimitConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.LimitConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LimitConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LimitConfigToPb(st *LimitConfig) (*billingpb.LimitConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.LimitConfigPb{}

	return pb, nil
}

func LimitConfigFromPb(pb *billingpb.LimitConfigPb) (*LimitConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LimitConfig{}

	return st, nil
}

type ListBudgetConfigurationsRequest struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListBudgetConfigurationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListBudgetConfigurationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListBudgetConfigurationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.ListBudgetConfigurationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListBudgetConfigurationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListBudgetConfigurationsRequestToPb(st *ListBudgetConfigurationsRequest) (*billingpb.ListBudgetConfigurationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.ListBudgetConfigurationsRequestPb{}
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListBudgetConfigurationsRequestFromPb(pb *billingpb.ListBudgetConfigurationsRequestPb) (*ListBudgetConfigurationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetConfigurationsRequest{}
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListBudgetConfigurationsResponse struct {

	// Wire name: 'budgets'
	Budgets []BudgetConfiguration `json:"budgets,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListBudgetConfigurationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListBudgetConfigurationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListBudgetConfigurationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.ListBudgetConfigurationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListBudgetConfigurationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListBudgetConfigurationsResponseToPb(st *ListBudgetConfigurationsResponse) (*billingpb.ListBudgetConfigurationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.ListBudgetConfigurationsResponsePb{}

	var budgetsPb []billingpb.BudgetConfigurationPb
	for _, item := range st.Budgets {
		itemPb, err := BudgetConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			budgetsPb = append(budgetsPb, *itemPb)
		}
	}
	pb.Budgets = budgetsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListBudgetConfigurationsResponseFromPb(pb *billingpb.ListBudgetConfigurationsResponsePb) (*ListBudgetConfigurationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetConfigurationsResponse{}

	var budgetsField []BudgetConfiguration
	for _, itemPb := range pb.Budgets {
		item, err := BudgetConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			budgetsField = append(budgetsField, *item)
		}
	}
	st.Budgets = budgetsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListBudgetPoliciesRequest struct {
	// A filter to apply to the list of policies.
	FilterBy *Filter `json:"-" tf:"-"`
	// The maximum number of budget policies to return. If unspecified, at most
	// 100 budget policies will be returned. The maximum value is 1000; values
	// above 1000 will be coerced to 1000.
	PageSize int `json:"-" tf:"-"`
	// A page token, received from a previous `ListServerlessPolicies` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	//
	// When paginating, all other parameters provided to
	// `ListServerlessPoliciesRequest` must match the call that provided the
	// page token.
	PageToken string `json:"-" tf:"-"`
	// The sort specification.
	SortSpec        *SortSpec `json:"-" tf:"-"`
	ForceSendFields []string  `json:"-" tf:"-"`
}

func (st ListBudgetPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListBudgetPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListBudgetPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.ListBudgetPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListBudgetPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListBudgetPoliciesRequestToPb(st *ListBudgetPoliciesRequest) (*billingpb.ListBudgetPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.ListBudgetPoliciesRequestPb{}
	filterByPb, err := FilterToPb(st.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByPb != nil {
		pb.FilterBy = filterByPb
	}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	sortSpecPb, err := SortSpecToPb(st.SortSpec)
	if err != nil {
		return nil, err
	}
	if sortSpecPb != nil {
		pb.SortSpec = sortSpecPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListBudgetPoliciesRequestFromPb(pb *billingpb.ListBudgetPoliciesRequestPb) (*ListBudgetPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetPoliciesRequest{}
	filterByField, err := FilterFromPb(pb.FilterBy)
	if err != nil {
		return nil, err
	}
	if filterByField != nil {
		st.FilterBy = filterByField
	}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	sortSpecField, err := SortSpecFromPb(pb.SortSpec)
	if err != nil {
		return nil, err
	}
	if sortSpecField != nil {
		st.SortSpec = sortSpecField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A list of policies.
type ListBudgetPoliciesResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'policies'
	Policies []BudgetPolicy `json:"policies,omitempty"`
	// A token that can be sent as `page_token` to retrieve the previous page.
	// In this field is omitted, there are no previous pages.
	// Wire name: 'previous_page_token'
	PreviousPageToken string   `json:"previous_page_token,omitempty"`
	ForceSendFields   []string `json:"-" tf:"-"`
}

func (st ListBudgetPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListBudgetPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListBudgetPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.ListBudgetPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListBudgetPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListBudgetPoliciesResponseToPb(st *ListBudgetPoliciesResponse) (*billingpb.ListBudgetPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.ListBudgetPoliciesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var policiesPb []billingpb.BudgetPolicyPb
	for _, item := range st.Policies {
		itemPb, err := BudgetPolicyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policiesPb = append(policiesPb, *itemPb)
		}
	}
	pb.Policies = policiesPb
	pb.PreviousPageToken = st.PreviousPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListBudgetPoliciesResponseFromPb(pb *billingpb.ListBudgetPoliciesResponsePb) (*ListBudgetPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListBudgetPoliciesResponse{}
	st.NextPageToken = pb.NextPageToken

	var policiesField []BudgetPolicy
	for _, itemPb := range pb.Policies {
		item, err := BudgetPolicyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			policiesField = append(policiesField, *item)
		}
	}
	st.Policies = policiesField
	st.PreviousPageToken = pb.PreviousPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListLogDeliveryRequest struct {
	// The Credentials id to filter the search results with
	CredentialsId string `json:"-" tf:"-"`
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken string `json:"-" tf:"-"`
	// The log delivery status to filter the search results with
	Status LogDeliveryConfigStatus `json:"-" tf:"-"`
	// The Storage Configuration id to filter the search results with
	StorageConfigurationId string   `json:"-" tf:"-"`
	ForceSendFields        []string `json:"-" tf:"-"`
}

func (st ListLogDeliveryRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListLogDeliveryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.ListLogDeliveryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListLogDeliveryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListLogDeliveryRequestToPb(st *ListLogDeliveryRequest) (*billingpb.ListLogDeliveryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.ListLogDeliveryRequestPb{}
	pb.CredentialsId = st.CredentialsId
	pb.PageToken = st.PageToken
	statusPb, err := LogDeliveryConfigStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StorageConfigurationId = st.StorageConfigurationId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListLogDeliveryRequestFromPb(pb *billingpb.ListLogDeliveryRequestPb) (*ListLogDeliveryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListLogDeliveryRequest{}
	st.CredentialsId = pb.CredentialsId
	st.PageToken = pb.PageToken
	statusField, err := LogDeliveryConfigStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StorageConfigurationId = pb.StorageConfigurationId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// * Log Delivery Status
//
// `ENABLED`: All dependencies have executed and succeeded `DISABLED`: At least
// one dependency has succeeded
type LogDeliveryConfigStatus string

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

// Values returns all possible values for LogDeliveryConfigStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *LogDeliveryConfigStatus) Values() []LogDeliveryConfigStatus {
	return []LogDeliveryConfigStatus{
		LogDeliveryConfigStatusDisabled,
		LogDeliveryConfigStatusEnabled,
	}
}

// Type always returns LogDeliveryConfigStatus to satisfy [pflag.Value] interface
func (f *LogDeliveryConfigStatus) Type() string {
	return "LogDeliveryConfigStatus"
}

func LogDeliveryConfigStatusToPb(st *LogDeliveryConfigStatus) (*billingpb.LogDeliveryConfigStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.LogDeliveryConfigStatusPb(*st)
	return &pb, nil
}

func LogDeliveryConfigStatusFromPb(pb *billingpb.LogDeliveryConfigStatusPb) (*LogDeliveryConfigStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := LogDeliveryConfigStatus(*pb)
	return &st, nil
}

// * Log Delivery Configuration
type LogDeliveryConfiguration struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string `json:"account_id"`
	// The unique UUID of log delivery configuration
	// Wire name: 'config_id'
	ConfigId string `json:"config_id,omitempty"`
	// The optional human-readable name of the log delivery configuration.
	// Defaults to empty.
	// Wire name: 'config_name'
	ConfigName string `json:"config_name,omitempty"`
	// Time in epoch milliseconds when the log delivery configuration was
	// created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// The ID for a method:credentials/create that represents the AWS IAM role
	// with policy and trust relationship as described in the main billable
	// usage documentation page. See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'credentials_id'
	CredentialsId string `json:"credentials_id"`
	// The optional delivery path prefix within Amazon S3 storage. Defaults to
	// empty, which means that logs are delivered to the root of the bucket.
	// This must be a valid S3 object key. This must not start or end with a
	// slash character.
	// Wire name: 'delivery_path_prefix'
	DeliveryPathPrefix string `json:"delivery_path_prefix,omitempty"`
	// This field applies only if log_type is BILLABLE_USAGE. This is the
	// optional start month and year for delivery, specified in YYYY-MM format.
	// Defaults to current year and month. BILLABLE_USAGE logs are not available
	// for usage before March 2019 (2019-03).
	// Wire name: 'delivery_start_time'
	DeliveryStartTime string `json:"delivery_start_time,omitempty"`
	// The LogDeliveryStatus of this log delivery configuration
	// Wire name: 'log_delivery_status'
	LogDeliveryStatus *LogDeliveryStatus `json:"log_delivery_status,omitempty"`
	// Log delivery type. Supported values are: * `BILLABLE_USAGE` — Configure
	// [billable usage log delivery]. For the CSV schema, see the [View billable
	// usage]. * `AUDIT_LOGS` — Configure [audit log delivery]. For the JSON
	// schema, see [Configure audit logging]
	//
	// [Configure audit logging]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// [audit log delivery]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [billable usage log delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'log_type'
	LogType LogType `json:"log_type"`
	// The file type of log delivery. * If `log_type` is `BILLABLE_USAGE`, this
	// value must be `CSV`. Only the CSV (comma-separated values) format is
	// supported. For the schema, see the [View billable usage] * If `log_type`
	// is `AUDIT_LOGS`, this value must be `JSON`. Only the JSON (JavaScript
	// Object Notation) format is supported. For the schema, see the
	// [Configuring audit logs].
	//
	// [Configuring audit logs]: https://docs.databricks.com/administration-guide/account-settings/audit-logs.html
	// [View billable usage]: https://docs.databricks.com/administration-guide/account-settings/usage.html
	// Wire name: 'output_format'
	OutputFormat OutputFormat `json:"output_format"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	// Wire name: 'status'
	Status LogDeliveryConfigStatus `json:"status,omitempty"`
	// The ID for a method:storage/create that represents the S3 bucket with
	// bucket policy as described in the main billable usage documentation page.
	// See [Configure billable usage delivery].
	//
	// [Configure billable usage delivery]: https://docs.databricks.com/administration-guide/account-settings/billable-usage-delivery.html
	// Wire name: 'storage_configuration_id'
	StorageConfigurationId string `json:"storage_configuration_id"`
	// Time in epoch milliseconds when the log delivery configuration was
	// updated.
	// Wire name: 'update_time'
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
	// Wire name: 'workspace_ids_filter'
	WorkspaceIdsFilter []int64  `json:"workspace_ids_filter,omitempty"`
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st LogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := LogDeliveryConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.LogDeliveryConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogDeliveryConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogDeliveryConfigurationToPb(st *LogDeliveryConfiguration) (*billingpb.LogDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.LogDeliveryConfigurationPb{}
	pb.AccountId = st.AccountId
	pb.ConfigId = st.ConfigId
	pb.ConfigName = st.ConfigName
	pb.CreationTime = st.CreationTime
	pb.CredentialsId = st.CredentialsId
	pb.DeliveryPathPrefix = st.DeliveryPathPrefix
	pb.DeliveryStartTime = st.DeliveryStartTime
	logDeliveryStatusPb, err := LogDeliveryStatusToPb(st.LogDeliveryStatus)
	if err != nil {
		return nil, err
	}
	if logDeliveryStatusPb != nil {
		pb.LogDeliveryStatus = logDeliveryStatusPb
	}
	logTypePb, err := LogTypeToPb(&st.LogType)
	if err != nil {
		return nil, err
	}
	if logTypePb != nil {
		pb.LogType = *logTypePb
	}
	outputFormatPb, err := OutputFormatToPb(&st.OutputFormat)
	if err != nil {
		return nil, err
	}
	if outputFormatPb != nil {
		pb.OutputFormat = *outputFormatPb
	}
	statusPb, err := LogDeliveryConfigStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StorageConfigurationId = st.StorageConfigurationId
	pb.UpdateTime = st.UpdateTime
	pb.WorkspaceIdsFilter = st.WorkspaceIdsFilter

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogDeliveryConfigurationFromPb(pb *billingpb.LogDeliveryConfigurationPb) (*LogDeliveryConfiguration, error) {
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
	logDeliveryStatusField, err := LogDeliveryStatusFromPb(pb.LogDeliveryStatus)
	if err != nil {
		return nil, err
	}
	if logDeliveryStatusField != nil {
		st.LogDeliveryStatus = logDeliveryStatusField
	}
	logTypeField, err := LogTypeFromPb(&pb.LogType)
	if err != nil {
		return nil, err
	}
	if logTypeField != nil {
		st.LogType = *logTypeField
	}
	outputFormatField, err := OutputFormatFromPb(&pb.OutputFormat)
	if err != nil {
		return nil, err
	}
	if outputFormatField != nil {
		st.OutputFormat = *outputFormatField
	}
	statusField, err := LogDeliveryConfigStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StorageConfigurationId = pb.StorageConfigurationId
	st.UpdateTime = pb.UpdateTime
	st.WorkspaceIdsFilter = pb.WorkspaceIdsFilter

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type LogDeliveryStatus struct {
	// The UTC time for the latest log delivery attempt.
	// Wire name: 'last_attempt_time'
	LastAttemptTime string `json:"last_attempt_time,omitempty"`
	// The UTC time for the latest successful log delivery.
	// Wire name: 'last_successful_attempt_time'
	LastSuccessfulAttemptTime string `json:"last_successful_attempt_time,omitempty"`
	// Informative message about the latest log delivery attempt. If the log
	// delivery fails with USER_FAILURE, error details will be provided for
	// fixing misconfigurations in cloud permissions.
	// Wire name: 'message'
	Message string `json:"message"`
	// Enum that describes the status. Possible values are: * `CREATED`: There
	// were no log delivery attempts since the config was created. *
	// `SUCCEEDED`: The latest attempt of log delivery has succeeded completely.
	// * `USER_FAILURE`: The latest attempt of log delivery failed because of
	// misconfiguration of customer provided permissions on role or storage. *
	// `SYSTEM_FAILURE`: The latest attempt of log delivery failed because of an
	// Databricks internal error. Contact support if it doesn't go away soon. *
	// `NOT_FOUND`: The log delivery status as the configuration has been
	// disabled since the release of this feature or there are no workspaces in
	// the account.
	// Wire name: 'status'
	Status          DeliveryStatus `json:"status"`
	ForceSendFields []string       `json:"-" tf:"-"`
}

func (st LogDeliveryStatus) MarshalJSON() ([]byte, error) {
	pb, err := LogDeliveryStatusToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogDeliveryStatus) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.LogDeliveryStatusPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogDeliveryStatusFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogDeliveryStatusToPb(st *LogDeliveryStatus) (*billingpb.LogDeliveryStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.LogDeliveryStatusPb{}
	pb.LastAttemptTime = st.LastAttemptTime
	pb.LastSuccessfulAttemptTime = st.LastSuccessfulAttemptTime
	pb.Message = st.Message
	statusPb, err := DeliveryStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogDeliveryStatusFromPb(pb *billingpb.LogDeliveryStatusPb) (*LogDeliveryStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogDeliveryStatus{}
	st.LastAttemptTime = pb.LastAttemptTime
	st.LastSuccessfulAttemptTime = pb.LastSuccessfulAttemptTime
	st.Message = pb.Message
	statusField, err := DeliveryStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// * Log Delivery Type
type LogType string

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

// Values returns all possible values for LogType.
//
// There is no guarantee on the order of the values in the slice.
func (f *LogType) Values() []LogType {
	return []LogType{
		LogTypeAuditLogs,
		LogTypeBillableUsage,
	}
}

// Type always returns LogType to satisfy [pflag.Value] interface
func (f *LogType) Type() string {
	return "LogType"
}

func LogTypeToPb(st *LogType) (*billingpb.LogTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.LogTypePb(*st)
	return &pb, nil
}

func LogTypeFromPb(pb *billingpb.LogTypePb) (*LogType, error) {
	if pb == nil {
		return nil, nil
	}
	st := LogType(*pb)
	return &st, nil
}

// * Log Delivery Output Format
type OutputFormat string

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

// Values returns all possible values for OutputFormat.
//
// There is no guarantee on the order of the values in the slice.
func (f *OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		OutputFormatCsv,
		OutputFormatJson,
	}
}

// Type always returns OutputFormat to satisfy [pflag.Value] interface
func (f *OutputFormat) Type() string {
	return "OutputFormat"
}

func OutputFormatToPb(st *OutputFormat) (*billingpb.OutputFormatPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.OutputFormatPb(*st)
	return &pb, nil
}

func OutputFormatFromPb(pb *billingpb.OutputFormatPb) (*OutputFormat, error) {
	if pb == nil {
		return nil, nil
	}
	st := OutputFormat(*pb)
	return &st, nil
}

type SortSpec struct {
	// Whether to sort in descending order.
	// Wire name: 'descending'
	Descending bool `json:"descending,omitempty"`
	// The filed to sort by
	// Wire name: 'field'
	Field           SortSpecField `json:"field,omitempty"`
	ForceSendFields []string      `json:"-" tf:"-"`
}

func (st SortSpec) MarshalJSON() ([]byte, error) {
	pb, err := SortSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SortSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.SortSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SortSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SortSpecToPb(st *SortSpec) (*billingpb.SortSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.SortSpecPb{}
	pb.Descending = st.Descending
	fieldPb, err := SortSpecFieldToPb(&st.Field)
	if err != nil {
		return nil, err
	}
	if fieldPb != nil {
		pb.Field = *fieldPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SortSpecFromPb(pb *billingpb.SortSpecPb) (*SortSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SortSpec{}
	st.Descending = pb.Descending
	fieldField, err := SortSpecFieldFromPb(&pb.Field)
	if err != nil {
		return nil, err
	}
	if fieldField != nil {
		st.Field = *fieldField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SortSpecField string

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

// Values returns all possible values for SortSpecField.
//
// There is no guarantee on the order of the values in the slice.
func (f *SortSpecField) Values() []SortSpecField {
	return []SortSpecField{
		SortSpecFieldPolicyName,
	}
}

// Type always returns SortSpecField to satisfy [pflag.Value] interface
func (f *SortSpecField) Type() string {
	return "SortSpecField"
}

func SortSpecFieldToPb(st *SortSpecField) (*billingpb.SortSpecFieldPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.SortSpecFieldPb(*st)
	return &pb, nil
}

func SortSpecFieldFromPb(pb *billingpb.SortSpecFieldPb) (*SortSpecField, error) {
	if pb == nil {
		return nil, nil
	}
	st := SortSpecField(*pb)
	return &st, nil
}

type UpdateBudgetConfigurationBudget struct {
	// Databricks account ID.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	// Wire name: 'alert_configurations'
	AlertConfigurations []AlertConfiguration `json:"alert_configurations,omitempty"`
	// Databricks budget configuration ID.
	// Wire name: 'budget_configuration_id'
	BudgetConfigurationId string `json:"budget_configuration_id,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	// Wire name: 'filter'
	Filter          *BudgetConfigurationFilter `json:"filter,omitempty"`
	ForceSendFields []string                   `json:"-" tf:"-"`
}

func (st UpdateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	pb, err := UpdateBudgetConfigurationBudgetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.UpdateBudgetConfigurationBudgetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateBudgetConfigurationBudgetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateBudgetConfigurationBudgetToPb(st *UpdateBudgetConfigurationBudget) (*billingpb.UpdateBudgetConfigurationBudgetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.UpdateBudgetConfigurationBudgetPb{}
	pb.AccountId = st.AccountId

	var alertConfigurationsPb []billingpb.AlertConfigurationPb
	for _, item := range st.AlertConfigurations {
		itemPb, err := AlertConfigurationToPb(&item)
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
	filterPb, err := BudgetConfigurationFilterToPb(st.Filter)
	if err != nil {
		return nil, err
	}
	if filterPb != nil {
		pb.Filter = filterPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateBudgetConfigurationBudgetFromPb(pb *billingpb.UpdateBudgetConfigurationBudgetPb) (*UpdateBudgetConfigurationBudget, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationBudget{}
	st.AccountId = pb.AccountId

	var alertConfigurationsField []AlertConfiguration
	for _, itemPb := range pb.AlertConfigurations {
		item, err := AlertConfigurationFromPb(&itemPb)
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
	filterField, err := BudgetConfigurationFilterFromPb(pb.Filter)
	if err != nil {
		return nil, err
	}
	if filterField != nil {
		st.Filter = filterField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateBudgetConfigurationRequest struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	// Wire name: 'budget'
	Budget UpdateBudgetConfigurationBudget `json:"budget"`
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" tf:"-"`
}

func (st UpdateBudgetConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateBudgetConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateBudgetConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.UpdateBudgetConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateBudgetConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateBudgetConfigurationRequestToPb(st *UpdateBudgetConfigurationRequest) (*billingpb.UpdateBudgetConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.UpdateBudgetConfigurationRequestPb{}
	budgetPb, err := UpdateBudgetConfigurationBudgetToPb(&st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = *budgetPb
	}
	pb.BudgetId = st.BudgetId

	return pb, nil
}

func UpdateBudgetConfigurationRequestFromPb(pb *billingpb.UpdateBudgetConfigurationRequestPb) (*UpdateBudgetConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationRequest{}
	budgetField, err := UpdateBudgetConfigurationBudgetFromPb(&pb.Budget)
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
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

func (st UpdateBudgetConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := UpdateBudgetConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateBudgetConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.UpdateBudgetConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateBudgetConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateBudgetConfigurationResponseToPb(st *UpdateBudgetConfigurationResponse) (*billingpb.UpdateBudgetConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.UpdateBudgetConfigurationResponsePb{}
	budgetPb, err := BudgetConfigurationToPb(st.Budget)
	if err != nil {
		return nil, err
	}
	if budgetPb != nil {
		pb.Budget = budgetPb
	}

	return pb, nil
}

func UpdateBudgetConfigurationResponseFromPb(pb *billingpb.UpdateBudgetConfigurationResponsePb) (*UpdateBudgetConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetConfigurationResponse{}
	budgetField, err := BudgetConfigurationFromPb(pb.Budget)
	if err != nil {
		return nil, err
	}
	if budgetField != nil {
		st.Budget = budgetField
	}

	return st, nil
}

type UpdateBudgetPolicyRequest struct {
	// DEPRECATED. This is redundant field as LimitConfig is part of the
	// BudgetPolicy
	LimitConfig *LimitConfig `json:"-" tf:"-"`
	// The policy to update. `creator_user_id` cannot be specified in the
	// request. All other fields must be specified even if not changed. The
	// `policy_id` is used to identify the policy to update.
	// Wire name: 'policy'
	Policy BudgetPolicy `json:"policy"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"-" tf:"-"`
}

func (st UpdateBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateBudgetPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.UpdateBudgetPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateBudgetPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateBudgetPolicyRequestToPb(st *UpdateBudgetPolicyRequest) (*billingpb.UpdateBudgetPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.UpdateBudgetPolicyRequestPb{}
	limitConfigPb, err := LimitConfigToPb(st.LimitConfig)
	if err != nil {
		return nil, err
	}
	if limitConfigPb != nil {
		pb.LimitConfig = limitConfigPb
	}
	policyPb, err := BudgetPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func UpdateBudgetPolicyRequestFromPb(pb *billingpb.UpdateBudgetPolicyRequestPb) (*UpdateBudgetPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateBudgetPolicyRequest{}
	limitConfigField, err := LimitConfigFromPb(pb.LimitConfig)
	if err != nil {
		return nil, err
	}
	if limitConfigField != nil {
		st.LimitConfig = limitConfigField
	}
	policyField, err := BudgetPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.PolicyId = pb.PolicyId

	return st, nil
}

// * Update Log Delivery Configuration
type UpdateLogDeliveryConfigurationStatusRequest struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId string `json:"-" tf:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	// Wire name: 'status'
	Status LogDeliveryConfigStatus `json:"status"`
}

func (st UpdateLogDeliveryConfigurationStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateLogDeliveryConfigurationStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateLogDeliveryConfigurationStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.UpdateLogDeliveryConfigurationStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateLogDeliveryConfigurationStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateLogDeliveryConfigurationStatusRequestToPb(st *UpdateLogDeliveryConfigurationStatusRequest) (*billingpb.UpdateLogDeliveryConfigurationStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.UpdateLogDeliveryConfigurationStatusRequestPb{}
	pb.LogDeliveryConfigurationId = st.LogDeliveryConfigurationId
	statusPb, err := LogDeliveryConfigStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	return pb, nil
}

func UpdateLogDeliveryConfigurationStatusRequestFromPb(pb *billingpb.UpdateLogDeliveryConfigurationStatusRequestPb) (*UpdateLogDeliveryConfigurationStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLogDeliveryConfigurationStatusRequest{}
	st.LogDeliveryConfigurationId = pb.LogDeliveryConfigurationId
	statusField, err := LogDeliveryConfigStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

type UsageDashboardType string

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

// Values returns all possible values for UsageDashboardType.
//
// There is no guarantee on the order of the values in the slice.
func (f *UsageDashboardType) Values() []UsageDashboardType {
	return []UsageDashboardType{
		UsageDashboardTypeUsageDashboardTypeGlobal,
		UsageDashboardTypeUsageDashboardTypeWorkspace,
	}
}

// Type always returns UsageDashboardType to satisfy [pflag.Value] interface
func (f *UsageDashboardType) Type() string {
	return "UsageDashboardType"
}

func UsageDashboardTypeToPb(st *UsageDashboardType) (*billingpb.UsageDashboardTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := billingpb.UsageDashboardTypePb(*st)
	return &pb, nil
}

func UsageDashboardTypeFromPb(pb *billingpb.UsageDashboardTypePb) (*UsageDashboardType, error) {
	if pb == nil {
		return nil, nil
	}
	st := UsageDashboardType(*pb)
	return &st, nil
}

// * Properties of the new log delivery configuration.
type WrappedCreateLogDeliveryConfiguration struct {

	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration CreateLogDeliveryConfigurationParams `json:"log_delivery_configuration"`
}

func (st WrappedCreateLogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := WrappedCreateLogDeliveryConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WrappedCreateLogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.WrappedCreateLogDeliveryConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WrappedCreateLogDeliveryConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WrappedCreateLogDeliveryConfigurationToPb(st *WrappedCreateLogDeliveryConfiguration) (*billingpb.WrappedCreateLogDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.WrappedCreateLogDeliveryConfigurationPb{}
	logDeliveryConfigurationPb, err := CreateLogDeliveryConfigurationParamsToPb(&st.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationPb != nil {
		pb.LogDeliveryConfiguration = *logDeliveryConfigurationPb
	}

	return pb, nil
}

func WrappedCreateLogDeliveryConfigurationFromPb(pb *billingpb.WrappedCreateLogDeliveryConfigurationPb) (*WrappedCreateLogDeliveryConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedCreateLogDeliveryConfiguration{}
	logDeliveryConfigurationField, err := CreateLogDeliveryConfigurationParamsFromPb(&pb.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationField != nil {
		st.LogDeliveryConfiguration = *logDeliveryConfigurationField
	}

	return st, nil
}

type WrappedLogDeliveryConfiguration struct {
	// The created log delivery configuration
	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

func (st WrappedLogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := WrappedLogDeliveryConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WrappedLogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.WrappedLogDeliveryConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WrappedLogDeliveryConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WrappedLogDeliveryConfigurationToPb(st *WrappedLogDeliveryConfiguration) (*billingpb.WrappedLogDeliveryConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.WrappedLogDeliveryConfigurationPb{}
	logDeliveryConfigurationPb, err := LogDeliveryConfigurationToPb(st.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationPb != nil {
		pb.LogDeliveryConfiguration = logDeliveryConfigurationPb
	}

	return pb, nil
}

func WrappedLogDeliveryConfigurationFromPb(pb *billingpb.WrappedLogDeliveryConfigurationPb) (*WrappedLogDeliveryConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedLogDeliveryConfiguration{}
	logDeliveryConfigurationField, err := LogDeliveryConfigurationFromPb(pb.LogDeliveryConfiguration)
	if err != nil {
		return nil, err
	}
	if logDeliveryConfigurationField != nil {
		st.LogDeliveryConfiguration = logDeliveryConfigurationField
	}

	return st, nil
}

type WrappedLogDeliveryConfigurations struct {
	// Log delivery configurations were returned successfully.
	// Wire name: 'log_delivery_configurations'
	LogDeliveryConfigurations []LogDeliveryConfiguration `json:"log_delivery_configurations,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WrappedLogDeliveryConfigurations) MarshalJSON() ([]byte, error) {
	pb, err := WrappedLogDeliveryConfigurationsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WrappedLogDeliveryConfigurations) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &billingpb.WrappedLogDeliveryConfigurationsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WrappedLogDeliveryConfigurationsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WrappedLogDeliveryConfigurationsToPb(st *WrappedLogDeliveryConfigurations) (*billingpb.WrappedLogDeliveryConfigurationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &billingpb.WrappedLogDeliveryConfigurationsPb{}

	var logDeliveryConfigurationsPb []billingpb.LogDeliveryConfigurationPb
	for _, item := range st.LogDeliveryConfigurations {
		itemPb, err := LogDeliveryConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			logDeliveryConfigurationsPb = append(logDeliveryConfigurationsPb, *itemPb)
		}
	}
	pb.LogDeliveryConfigurations = logDeliveryConfigurationsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func WrappedLogDeliveryConfigurationsFromPb(pb *billingpb.WrappedLogDeliveryConfigurationsPb) (*WrappedLogDeliveryConfigurations, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WrappedLogDeliveryConfigurations{}

	var logDeliveryConfigurationsField []LogDeliveryConfiguration
	for _, itemPb := range pb.LogDeliveryConfigurations {
		item, err := LogDeliveryConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			logDeliveryConfigurationsField = append(logDeliveryConfigurationsField, *item)
		}
	}
	st.LogDeliveryConfigurations = logDeliveryConfigurationsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
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
