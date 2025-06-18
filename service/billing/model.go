// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/service/compute"
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
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	TriggerType AlertConfigurationTriggerType `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	UpdateTime int64 `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type BudgetConfigurationFilterClause struct {

	// Wire name: 'operator'
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	// Wire name: 'values'
	Values []string `json:"values,omitempty"`
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

type BudgetConfigurationFilterTagClause struct {

	// Wire name: 'key'
	Key string `json:"key,omitempty"`

	// Wire name: 'value'
	Value *BudgetConfigurationFilterClause `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type BudgetConfigurationFilterWorkspaceIdClause struct {

	// Wire name: 'operator'
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	// Wire name: 'values'
	Values []int64 `json:"values,omitempty"`
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
	PolicyName string `json:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type CreateBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	// Wire name: 'dashboard_type'
	DashboardType UsageDashboardType `json:"dashboard_type,omitempty"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type CreateBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type CreateBudgetConfigurationBudgetActionConfigurations struct {
	// The type of the action.
	// Wire name: 'action_type'
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target for the action. For example, an email address.
	// Wire name: 'target'
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	TriggerType AlertConfigurationTriggerType `json:"trigger_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type CreateBudgetConfigurationRequest struct {
	// Properties of the new budget configuration.
	// Wire name: 'budget'
	Budget CreateBudgetConfigurationBudget `json:"budget"`
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

type CreateBudgetConfigurationResponse struct {
	// The created budget configuration.
	// Wire name: 'budget'
	Budget *BudgetConfiguration `json:"budget,omitempty"`
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
	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Delete budget
type DeleteBudgetConfigurationRequest struct {
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" tf:"-"`
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

type DeleteBudgetConfigurationResponse struct {
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

// Delete a budget policy
type DeleteBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" tf:"-"`
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

type DeleteResponse struct {
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

// Return billable usage logs
type DownloadRequest struct {
	// Format: `YYYY-MM`. Last month to return billable usage logs for. This
	// field is required.
	EndMonth string `json:"-" tf:"-"`
	// Specify whether to include personally identifiable information in the
	// billable usage logs, for example the email addresses of cluster creators.
	// Handle this information with care. Defaults to false.
	PersonalData bool `json:"-" tf:"-"`
	// Format: `YYYY-MM`. First month to return billable usage logs for. This
	// field is required.
	StartMonth string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type DownloadResponse struct {
	Contents io.ReadCloser `json:"-" tf:"-"`
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
	PolicyName string `json:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Get usage dashboard
type GetBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType UsageDashboardType `json:"-" tf:"-"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId int64 `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type GetBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	// Wire name: 'dashboard_id'
	DashboardId string `json:"dashboard_id,omitempty"`
	// The URL of the usage dashboard.
	// Wire name: 'dashboard_url'
	DashboardUrl string `json:"dashboard_url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Get budget
type GetBudgetConfigurationRequest struct {
	// The budget configuration ID
	BudgetId string `json:"-" tf:"-"`
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

type GetBudgetConfigurationResponse struct {

	// Wire name: 'budget'
	Budget *BudgetConfiguration `json:"budget,omitempty"`
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

// Get a budget policy
type GetBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" tf:"-"`
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

type GetLogDeliveryConfigurationResponse struct {
	// The fetched log delivery configuration
	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

func (st *GetLogDeliveryConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLogDeliveryConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLogDeliveryConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLogDeliveryConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := getLogDeliveryConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get log delivery configuration
type GetLogDeliveryRequest struct {
	// The log delivery configuration id of customer
	LogDeliveryConfigurationId string `json:"-" tf:"-"`
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

// The limit configuration of the policy. Limit configuration provide a budget
// policy level cost control by enforcing the limit.
type LimitConfig struct {
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

// Get all budgets
type ListBudgetConfigurationsRequest struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type ListBudgetConfigurationsResponse struct {

	// Wire name: 'budgets'
	Budgets []BudgetConfiguration `json:"budgets,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// List policies
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
	SortSpec *SortSpec `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	PreviousPageToken string `json:"previous_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

// Get all log delivery configurations
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
	StorageConfigurationId string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	WorkspaceIdsFilter []int64 `json:"workspace_ids_filter,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Status DeliveryStatus `json:"status"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type PatchStatusResponse struct {
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

type SortSpec struct {
	// Whether to sort in descending order.
	// Wire name: 'descending'
	Descending bool `json:"descending,omitempty"`
	// The filed to sort by
	// Wire name: 'field'
	Field SortSpecField `json:"field,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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

type UpdateBudgetConfigurationRequest struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	// Wire name: 'budget'
	Budget UpdateBudgetConfigurationBudget `json:"budget"`
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" tf:"-"`
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

type UpdateBudgetConfigurationResponse struct {
	// The updated budget.
	// Wire name: 'budget'
	Budget *BudgetConfiguration `json:"budget,omitempty"`
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

// Update a budget policy
type UpdateBudgetPolicyRequest struct {
	// DEPRECATED. This is redundant field as LimitConfig is part of the
	// BudgetPolicy
	LimitConfig *LimitConfig `json:"-" tf:"-"`
	// Contains the BudgetPolicy details.
	// Wire name: 'policy'
	Policy BudgetPolicy `json:"policy"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"-" tf:"-"`
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

// * Properties of the new log delivery configuration.
type WrappedCreateLogDeliveryConfiguration struct {
	// * Log Delivery Configuration
	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration CreateLogDeliveryConfigurationParams `json:"log_delivery_configuration"`
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

type WrappedLogDeliveryConfiguration struct {
	// The created log delivery configuration
	// Wire name: 'log_delivery_configuration'
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
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

type WrappedLogDeliveryConfigurations struct {
	// Log delivery configurations were returned successfully.
	// Wire name: 'log_delivery_configurations'
	LogDeliveryConfigurations []LogDeliveryConfiguration `json:"log_delivery_configurations,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
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
