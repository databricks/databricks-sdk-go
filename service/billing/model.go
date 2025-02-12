// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package billing

import (
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

type ActionConfiguration struct {
	// Databricks action configuration ID.
	ActionConfigurationId string `json:"action_configuration_id,omitempty"`
	// The type of the action.
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target for the action. For example, an email address.
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ActionConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ActionConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns ActionConfigurationType to satisfy [pflag.Value] interface
func (f *ActionConfigurationType) Type() string {
	return "ActionConfigurationType"
}

type AlertConfiguration struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations []ActionConfiguration `json:"action_configurations,omitempty"`
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

func (s *AlertConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AlertConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns AlertConfigurationTriggerType to satisfy [pflag.Value] interface
func (f *AlertConfigurationTriggerType) Type() string {
	return "AlertConfigurationTriggerType"
}

type BudgetConfiguration struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations []AlertConfiguration `json:"alert_configurations,omitempty"`
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
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`
	// Update time of this budget configuration.
	UpdateTime int64 `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BudgetConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BudgetConfigurationFilter struct {
	// A list of tag keys and values that will limit the budget to usage that
	// includes those specific custom tags. Tags are case-sensitive and should
	// be entered exactly as they appear in your usage data.
	Tags []BudgetConfigurationFilterTagClause `json:"tags,omitempty"`
	// If provided, usage must match with the provided Databricks workspace IDs.
	WorkspaceId *BudgetConfigurationFilterWorkspaceIdClause `json:"workspace_id,omitempty"`
}

type BudgetConfigurationFilterClause struct {
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	Values []string `json:"values,omitempty"`
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

// Type always returns BudgetConfigurationFilterOperator to satisfy [pflag.Value] interface
func (f *BudgetConfigurationFilterOperator) Type() string {
	return "BudgetConfigurationFilterOperator"
}

type BudgetConfigurationFilterTagClause struct {
	Key string `json:"key,omitempty"`

	Value *BudgetConfigurationFilterClause `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BudgetConfigurationFilterTagClause) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetConfigurationFilterTagClause) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BudgetConfigurationFilterWorkspaceIdClause struct {
	Operator BudgetConfigurationFilterOperator `json:"operator,omitempty"`

	Values []int64 `json:"values,omitempty"`
}

// Contains the BudgetPolicy details.
type BudgetPolicy struct {
	// A list of tags defined by the customer. At most 20 entries are allowed
	// per policy.
	CustomTags []compute.CustomPolicyTag `json:"custom_tags,omitempty"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"policy_id"`
	// The name of the policy. - Must be unique among active policies. - Can
	// contain only characters from the ISO 8859-1 (latin1) set.
	PolicyName string `json:"policy_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BudgetPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BudgetPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType UsageDashboardType `json:"dashboard_type,omitempty"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBudgetConfigurationBudget struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations []CreateBudgetConfigurationBudgetAlertConfigurations `json:"alert_configurations,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBudgetConfigurationBudgetActionConfigurations struct {
	// The type of the action.
	ActionType ActionConfigurationType `json:"action_type,omitempty"`
	// Target for the action. For example, an email address.
	Target string `json:"target,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBudgetConfigurationBudgetActionConfigurations) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetConfigurationBudgetActionConfigurations) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBudgetConfigurationBudgetAlertConfigurations struct {
	// Configured actions for this alert. These define what happens when an
	// alert enters a triggered state.
	ActionConfigurations []CreateBudgetConfigurationBudgetActionConfigurations `json:"action_configurations,omitempty"`
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

func (s *CreateBudgetConfigurationBudgetAlertConfigurations) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetConfigurationBudgetAlertConfigurations) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateBudgetConfigurationRequest struct {
	// Properties of the new budget configuration.
	Budget CreateBudgetConfigurationBudget `json:"budget"`
}

type CreateBudgetConfigurationResponse struct {
	// The created budget configuration.
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

// A request to create a BudgetPolicy.
type CreateBudgetPolicyRequest struct {
	// A list of tags defined by the customer. At most 40 entries are allowed
	// per policy.
	CustomTags []compute.CustomPolicyTag `json:"custom_tags,omitempty"`
	// The name of the policy. - Must be unique among active policies. - Can
	// contain only characters of 0-9, a-z, A-Z, -, =, ., :, /, @, _, +,
	// whitespace.
	PolicyName string `json:"policy_name,omitempty"`
	// A unique identifier for this request. Restricted to 36 ASCII characters.
	// A random UUID is recommended. This request is only idempotent if a
	// `request_id` is provided.
	RequestId string `json:"request_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateBudgetPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateBudgetPolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateLogDeliveryConfigurationParams struct {
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

func (s *CreateLogDeliveryConfigurationParams) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateLogDeliveryConfigurationParams) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete budget
type DeleteBudgetConfigurationRequest struct {
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" url:"-"`
}

type DeleteBudgetConfigurationResponse struct {
}

// Delete a budget policy
type DeleteBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" url:"-"`
}

type DeleteResponse struct {
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

// Return billable usage logs
type DownloadRequest struct {
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

func (s *DownloadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DownloadRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DownloadResponse struct {
	Contents io.ReadCloser `json:"-"`
}

// Structured representation of a filter to be applied to a list of policies.
// All specified filters will be applied in conjunction.
type Filter struct {
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

func (s *Filter) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Filter) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get usage dashboard
type GetBillingUsageDashboardRequest struct {
	// Workspace level usage dashboard shows usage data for the specified
	// workspace ID. Global level usage dashboard shows usage data for all
	// workspaces in the account.
	DashboardType UsageDashboardType `json:"-" url:"dashboard_type,omitempty"`
	// The workspace ID of the workspace in which the usage dashboard is
	// created.
	WorkspaceId int64 `json:"-" url:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetBillingUsageDashboardRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetBillingUsageDashboardRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetBillingUsageDashboardResponse struct {
	// The unique id of the usage dashboard.
	DashboardId string `json:"dashboard_id,omitempty"`
	// The URL of the usage dashboard.
	DashboardUrl string `json:"dashboard_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetBillingUsageDashboardResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetBillingUsageDashboardResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get budget
type GetBudgetConfigurationRequest struct {
	// The budget configuration ID
	BudgetId string `json:"-" url:"-"`
}

type GetBudgetConfigurationResponse struct {
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

// Get a budget policy
type GetBudgetPolicyRequest struct {
	// The Id of the policy.
	PolicyId string `json:"-" url:"-"`
}

// Get log delivery configuration
type GetLogDeliveryRequest struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId string `json:"-" url:"-"`
}

// The limit configuration of the policy. Limit configuration provide a budget
// policy level cost control by enforcing the limit.
type LimitConfig struct {
}

// Get all budgets
type ListBudgetConfigurationsRequest struct {
	// A page token received from a previous get all budget configurations call.
	// This token can be used to retrieve the subsequent page. Requests first
	// page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetConfigurationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetConfigurationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListBudgetConfigurationsResponse struct {
	Budgets []BudgetConfiguration `json:"budgets,omitempty"`
	// Token which can be sent as `page_token` to retrieve the next page of
	// results. If this field is omitted, there are no subsequent budgets.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetConfigurationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List policies
type ListBudgetPoliciesRequest struct {
	// A filter to apply to the list of policies.
	FilterBy *Filter `json:"-" url:"filter_by,omitempty"`
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
	SortSpec *SortSpec `json:"-" url:"sort_spec,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of policies.
type ListBudgetPoliciesResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	Policies []BudgetPolicy `json:"policies,omitempty"`
	// A token that can be sent as `page_token` to retrieve the previous page.
	// In this field is omitted, there are no previous pages.
	PreviousPageToken string `json:"previous_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListBudgetPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListBudgetPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get all log delivery configurations
type ListLogDeliveryRequest struct {
	// Filter by credential configuration ID.
	CredentialsId string `json:"-" url:"credentials_id,omitempty"`
	// Filter by status `ENABLED` or `DISABLED`.
	Status LogDeliveryConfigStatus `json:"-" url:"status,omitempty"`
	// Filter by storage configuration ID.
	StorageConfigurationId string `json:"-" url:"storage_configuration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListLogDeliveryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListLogDeliveryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Status of log delivery configuration. Set to `ENABLED` (enabled) or
// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable the
// configuration](#operation/patch-log-delivery-config-status) later. Deletion
// of a configuration is not supported, so disable a log delivery configuration
// that is no longer needed.
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

// Type always returns LogDeliveryConfigStatus to satisfy [pflag.Value] interface
func (f *LogDeliveryConfigStatus) Type() string {
	return "LogDeliveryConfigStatus"
}

type LogDeliveryConfiguration struct {
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
	LogDeliveryStatus *LogDeliveryStatus `json:"log_delivery_status,omitempty"`
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

func (s *LogDeliveryConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogDeliveryConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks log delivery status.
type LogDeliveryStatus struct {
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

func (s *LogDeliveryStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s LogDeliveryStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

type PatchStatusResponse struct {
}

type SortSpec struct {
	// Whether to sort in descending order.
	Descending bool `json:"descending,omitempty" url:"descending,omitempty"`
	// The filed to sort by
	Field SortSpecField `json:"field,omitempty" url:"field,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SortSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SortSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

// Type always returns SortSpecField to satisfy [pflag.Value] interface
func (f *SortSpecField) Type() string {
	return "SortSpecField"
}

type UpdateBudgetConfigurationBudget struct {
	// Databricks account ID.
	AccountId string `json:"account_id,omitempty"`
	// Alerts to configure when this budget is in a triggered state. Budgets
	// must have exactly one alert configuration.
	AlertConfigurations []AlertConfiguration `json:"alert_configurations,omitempty"`
	// Databricks budget configuration ID.
	BudgetConfigurationId string `json:"budget_configuration_id,omitempty"`
	// Human-readable name of budget configuration. Max Length: 128
	DisplayName string `json:"display_name,omitempty"`
	// Configured filters for this budget. These are applied to your account's
	// usage to limit the scope of what is considered for this budget. Leave
	// empty to include all usage for this account. All provided filters must be
	// matched for usage to be included.
	Filter *BudgetConfigurationFilter `json:"filter,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateBudgetConfigurationBudget) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateBudgetConfigurationBudget) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateBudgetConfigurationRequest struct {
	// The updated budget. This will overwrite the budget specified by the
	// budget ID.
	Budget UpdateBudgetConfigurationBudget `json:"budget"`
	// The Databricks budget configuration ID.
	BudgetId string `json:"-" url:"-"`
}

type UpdateBudgetConfigurationResponse struct {
	// The updated budget.
	Budget *BudgetConfiguration `json:"budget,omitempty"`
}

// Update a budget policy
type UpdateBudgetPolicyRequest struct {
	// DEPRECATED. This is redundant field as LimitConfig is part of the
	// BudgetPolicy
	LimitConfig *LimitConfig `json:"-" url:"limit_config,omitempty"`
	// Contains the BudgetPolicy details.
	Policy *BudgetPolicy `json:"policy,omitempty"`
	// The Id of the policy. This field is generated by Databricks and globally
	// unique.
	PolicyId string `json:"-" url:"-"`
}

type UpdateLogDeliveryConfigurationStatusRequest struct {
	// Databricks log delivery configuration ID
	LogDeliveryConfigurationId string `json:"-" url:"-"`
	// Status of log delivery configuration. Set to `ENABLED` (enabled) or
	// `DISABLED` (disabled). Defaults to `ENABLED`. You can [enable or disable
	// the configuration](#operation/patch-log-delivery-config-status) later.
	// Deletion of a configuration is not supported, so disable a log delivery
	// configuration that is no longer needed.
	Status LogDeliveryConfigStatus `json:"status"`
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

// Type always returns UsageDashboardType to satisfy [pflag.Value] interface
func (f *UsageDashboardType) Type() string {
	return "UsageDashboardType"
}

type WrappedCreateLogDeliveryConfiguration struct {
	LogDeliveryConfiguration *CreateLogDeliveryConfigurationParams `json:"log_delivery_configuration,omitempty"`
}

type WrappedLogDeliveryConfiguration struct {
	LogDeliveryConfiguration *LogDeliveryConfiguration `json:"log_delivery_configuration,omitempty"`
}

type WrappedLogDeliveryConfigurations struct {
	LogDeliveryConfigurations []LogDeliveryConfiguration `json:"log_delivery_configurations,omitempty"`
}
