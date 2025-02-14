// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccountIpAccessEnable struct {
	AcctIpAclEnable BooleanMessage `json:"acct_ip_acl_enable"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AccountIpAccessEnable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AccountIpAccessEnable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AibiDashboardEmbeddingAccessPolicy struct {
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType `json:"access_policy_type"`
}

type AibiDashboardEmbeddingAccessPolicyAccessPolicyType string

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `ALLOW_ALL_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `ALLOW_APPROVED_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeDenyAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyType = `DENY_ALL_DOMAINS`

// String representation for [fmt.Print]
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Set(v string) error {
	switch v {
	case `ALLOW_ALL_DOMAINS`, `ALLOW_APPROVED_DOMAINS`, `DENY_ALL_DOMAINS`:
		*f = AibiDashboardEmbeddingAccessPolicyAccessPolicyType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL_DOMAINS", "ALLOW_APPROVED_DOMAINS", "DENY_ALL_DOMAINS"`, v)
	}
}

// Type always returns AibiDashboardEmbeddingAccessPolicyAccessPolicyType to satisfy [pflag.Value] interface
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Type() string {
	return "AibiDashboardEmbeddingAccessPolicyAccessPolicyType"
}

type AibiDashboardEmbeddingAccessPolicySetting struct {
	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicy `json:"aibi_dashboard_embedding_access_policy"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AibiDashboardEmbeddingAccessPolicySetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AibiDashboardEmbeddingAccessPolicySetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AibiDashboardEmbeddingApprovedDomains struct {
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

type AibiDashboardEmbeddingApprovedDomainsSetting struct {
	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomains `json:"aibi_dashboard_embedding_approved_domains"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AibiDashboardEmbeddingApprovedDomainsSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AibiDashboardEmbeddingApprovedDomainsSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutomaticClusterUpdateSetting struct {
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessage `json:"automatic_cluster_update_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AutomaticClusterUpdateSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutomaticClusterUpdateSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type BooleanMessage struct {
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *BooleanMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BooleanMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAutoRestartMessage struct {
	CanToggle bool `json:"can_toggle,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails `json:"enablement_details,omitempty"`

	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow `json:"maintenance_window,omitempty"`

	RestartEvenIfNoUpdatesAvailable bool `json:"restart_even_if_no_updates_available,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode bool `json:"forced_for_compliance_mode,omitempty"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement bool `json:"unavailable_for_disabled_entitlement,omitempty"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier bool `json:"unavailable_for_non_enterprise_tier,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessageEnablementDetails) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageEnablementDetails) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `json:"week_day_based_schedule,omitempty"`
}

type ClusterAutoRestartMessageMaintenanceWindowDayOfWeek string

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFriday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `FRIDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekMonday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `MONDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSaturday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `SATURDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `SUNDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekThursday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `THURSDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekTuesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `TUESDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekWednesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeek = `WEDNESDAY`

// String representation for [fmt.Print]
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Set(v string) error {
	switch v {
	case `FRIDAY`, `MONDAY`, `SATURDAY`, `SUNDAY`, `THURSDAY`, `TUESDAY`, `WEDNESDAY`:
		*f = ClusterAutoRestartMessageMaintenanceWindowDayOfWeek(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FRIDAY", "MONDAY", "SATURDAY", "SUNDAY", "THURSDAY", "TUESDAY", "WEDNESDAY"`, v)
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowDayOfWeek to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowDayOfWeek"
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek `json:"day_of_week,omitempty"`

	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `json:"frequency,omitempty"`

	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `json:"window_start_time,omitempty"`
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency string

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `EVERY_WEEK`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_AND_THIRD_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FIRST_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_AND_FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `SECOND_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency = `THIRD_OF_MONTH`

// String representation for [fmt.Print]
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Set(v string) error {
	switch v {
	case `EVERY_WEEK`, `FIRST_AND_THIRD_OF_MONTH`, `FIRST_OF_MONTH`, `FOURTH_OF_MONTH`, `SECOND_AND_FOURTH_OF_MONTH`, `SECOND_OF_MONTH`, `THIRD_OF_MONTH`:
		*f = ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EVERY_WEEK", "FIRST_AND_THIRD_OF_MONTH", "FIRST_OF_MONTH", "FOURTH_OF_MONTH", "SECOND_AND_FOURTH_OF_MONTH", "SECOND_OF_MONTH", "THIRD_OF_MONTH"`, v)
	}
}

// Type always returns ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency to satisfy [pflag.Value] interface
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Type() string {
	return "ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency"
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours int `json:"hours,omitempty"`

	Minutes int `json:"minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`

	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ComplianceSecurityProfileSetting struct {
	// SHIELD feature: CSP
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfile `json:"compliance_security_profile_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ComplianceSecurityProfileSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ComplianceSecurityProfileSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Compliance stardard for SHIELD customers
type ComplianceStandard string

const ComplianceStandardCanadaProtectedB ComplianceStandard = `CANADA_PROTECTED_B`

const ComplianceStandardCyberEssentialPlus ComplianceStandard = `CYBER_ESSENTIAL_PLUS`

const ComplianceStandardFedrampHigh ComplianceStandard = `FEDRAMP_HIGH`

const ComplianceStandardFedrampIl5 ComplianceStandard = `FEDRAMP_IL5`

const ComplianceStandardFedrampModerate ComplianceStandard = `FEDRAMP_MODERATE`

const ComplianceStandardHipaa ComplianceStandard = `HIPAA`

const ComplianceStandardHitrust ComplianceStandard = `HITRUST`

const ComplianceStandardIrapProtected ComplianceStandard = `IRAP_PROTECTED`

const ComplianceStandardIsmap ComplianceStandard = `ISMAP`

const ComplianceStandardItarEar ComplianceStandard = `ITAR_EAR`

const ComplianceStandardNone ComplianceStandard = `NONE`

const ComplianceStandardPciDss ComplianceStandard = `PCI_DSS`

// String representation for [fmt.Print]
func (f *ComplianceStandard) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComplianceStandard) Set(v string) error {
	switch v {
	case `CANADA_PROTECTED_B`, `CYBER_ESSENTIAL_PLUS`, `FEDRAMP_HIGH`, `FEDRAMP_IL5`, `FEDRAMP_MODERATE`, `HIPAA`, `HITRUST`, `IRAP_PROTECTED`, `ISMAP`, `ITAR_EAR`, `NONE`, `PCI_DSS`:
		*f = ComplianceStandard(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANADA_PROTECTED_B", "CYBER_ESSENTIAL_PLUS", "FEDRAMP_HIGH", "FEDRAMP_IL5", "FEDRAMP_MODERATE", "HIPAA", "HITRUST", "IRAP_PROTECTED", "ISMAP", "ITAR_EAR", "NONE", "PCI_DSS"`, v)
	}
}

// Type always returns ComplianceStandard to satisfy [pflag.Value] interface
func (f *ComplianceStandard) Type() string {
	return "ComplianceStandard"
}

type Config struct {
	Email *EmailConfig `json:"email,omitempty"`

	GenericWebhook *GenericWebhookConfig `json:"generic_webhook,omitempty"`

	MicrosoftTeams *MicrosoftTeamsConfig `json:"microsoft_teams,omitempty"`

	Pagerduty *PagerdutyConfig `json:"pagerduty,omitempty"`

	Slack *SlackConfig `json:"slack,omitempty"`
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type"`
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

type CreateNetworkConnectivityConfigRequest struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name string `json:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string `json:"region"`
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *Config `json:"config,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId string `json:"application_id"`
	// Comment that describes the purpose of the token.
	Comment string `json:"comment,omitempty"`
	// The number of seconds before the token expires.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateOboTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePrivateEndpointRuleRequest struct {
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId CreatePrivateEndpointRuleRequestGroupId `json:"group_id"`
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id"`
}

// The sub-resource type (group ID) of the target resource. Note that to connect
// to workspace root storage (root DBFS), you need two endpoints, one for `blob`
// and one for `dfs`.
type CreatePrivateEndpointRuleRequestGroupId string

const CreatePrivateEndpointRuleRequestGroupIdBlob CreatePrivateEndpointRuleRequestGroupId = `blob`

const CreatePrivateEndpointRuleRequestGroupIdDfs CreatePrivateEndpointRuleRequestGroupId = `dfs`

const CreatePrivateEndpointRuleRequestGroupIdMysqlServer CreatePrivateEndpointRuleRequestGroupId = `mysqlServer`

const CreatePrivateEndpointRuleRequestGroupIdSqlServer CreatePrivateEndpointRuleRequestGroupId = `sqlServer`

// String representation for [fmt.Print]
func (f *CreatePrivateEndpointRuleRequestGroupId) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CreatePrivateEndpointRuleRequestGroupId) Set(v string) error {
	switch v {
	case `blob`, `dfs`, `mysqlServer`, `sqlServer`:
		*f = CreatePrivateEndpointRuleRequestGroupId(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "blob", "dfs", "mysqlServer", "sqlServer"`, v)
	}
}

// Type always returns CreatePrivateEndpointRuleRequestGroupId to satisfy [pflag.Value] interface
func (f *CreatePrivateEndpointRuleRequestGroupId) Type() string {
	return "CreatePrivateEndpointRuleRequestGroupId"
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateTokenResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CspEnablementAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CspEnablementAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CspEnablementAccountSetting struct {
	// Account level policy for CSP
	CspEnablementAccount CspEnablementAccount `json:"csp_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CspEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CspEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This represents the setting configuration for the default namespace in the
// Databricks workspace. Setting the default catalog for the workspace
// determines the catalog that is used when queries do not reference a fully
// qualified 3 level name. For example, if the default catalog is set to
// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the object
// 'retail_prod.default.myTable' (the schema 'default' is always assumed). This
// setting requires a restart of clusters and SQL warehouses to take effect.
// Additionally, the default namespace only applies when using Unity
// Catalog-enabled compute.
type DefaultNamespaceSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	Namespace StringMessage `json:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DefaultNamespaceSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DefaultNamespaceSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete the account IP access toggle setting
type DeleteAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteAccountIpAccessEnableResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

// Delete the AI/BI dashboard embedding access policy
type DeleteAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingAccessPolicySettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete AI/BI dashboard embedding approved domains
type DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete the default namespace setting
type DeleteDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDefaultNamespaceSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete Legacy Access Disablement Status
type DeleteDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDisableLegacyAccessResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete the disable legacy DBFS setting
type DeleteDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDisableLegacyDbfsResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete the disable legacy features setting
type DeleteDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteDisableLegacyFeaturesResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

type DeleteNetworkConnectivityConfigurationResponse struct {
}

// Delete a notification destination
type DeleteNotificationDestinationRequest struct {
	Id string `json:"-" url:"-"`
}

// Delete Personal Compute setting
type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeletePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeletePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeletePersonalComputeSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete the restrict workspace admins setting
type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The etag is returned.
type DeleteRestrictWorkspaceAdminsSettingResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId string `json:"-" url:"-"`
}

type DestinationType string

const DestinationTypeEmail DestinationType = `EMAIL`

const DestinationTypeMicrosoftTeams DestinationType = `MICROSOFT_TEAMS`

const DestinationTypePagerduty DestinationType = `PAGERDUTY`

const DestinationTypeSlack DestinationType = `SLACK`

const DestinationTypeWebhook DestinationType = `WEBHOOK`

// String representation for [fmt.Print]
func (f *DestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DestinationType) Set(v string) error {
	switch v {
	case `EMAIL`, `MICROSOFT_TEAMS`, `PAGERDUTY`, `SLACK`, `WEBHOOK`:
		*f = DestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "EMAIL", "MICROSOFT_TEAMS", "PAGERDUTY", "SLACK", "WEBHOOK"`, v)
	}
}

// Type always returns DestinationType to satisfy [pflag.Value] interface
func (f *DestinationType) Type() string {
	return "DestinationType"
}

type DisableLegacyAccess struct {
	DisableLegacyAccess BooleanMessage `json:"disable_legacy_access"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DisableLegacyAccess) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DisableLegacyAccess) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DisableLegacyDbfs struct {
	DisableLegacyDbfs BooleanMessage `json:"disable_legacy_dbfs"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DisableLegacyDbfs) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DisableLegacyDbfs) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DisableLegacyFeatures struct {
	DisableLegacyFeatures BooleanMessage `json:"disable_legacy_features"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DisableLegacyFeatures) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DisableLegacyFeatures) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess *EgressNetworkPolicyInternetAccessPolicy `json:"internet_access,omitempty"`
}

type EgressNetworkPolicyInternetAccessPolicy struct {
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`

	AllowedStorageDestinations []EgressNetworkPolicyInternetAccessPolicyStorageDestination `json:"allowed_storage_destinations,omitempty"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode `json:"log_only_mode,omitempty"`
	// At which level can Databricks and Databricks managed compute access
	// Internet. FULL_ACCESS: Databricks can access Internet. No blocking rules
	// will apply. RESTRICTED_ACCESS: Databricks can only access explicitly
	// allowed internet and storage destinations, as well as UC connections and
	// external locations. PRIVATE_ACCESS_ONLY (not used): Databricks can only
	// access destinations via private link.
	RestrictionMode EgressNetworkPolicyInternetAccessPolicyRestrictionMode `json:"restriction_mode,omitempty"`
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination struct {
	Destination string `json:"destination,omitempty"`
	// The filtering protocol used by the DP. For private and public preview,
	// SEG will only support TCP filtering (i.e. DNS based filtering, filtering
	// by destination IP address), so protocol will be set to TCP by default and
	// hidden from the user. In the future, users may be able to select HTTP
	// filtering (i.e. SNI based filtering, filtering by FQDN).
	Protocol EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol `json:"protocol,omitempty"`

	Type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EgressNetworkPolicyInternetAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyInternetAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The filtering protocol used by the DP. For private and public preview, SEG
// will only support TCP filtering (i.e. DNS based filtering, filtering by
// destination IP address), so protocol will be set to TCP by default and hidden
// from the user. In the future, users may be able to select HTTP filtering
// (i.e. SNI based filtering, filtering by FQDN).
type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolTcp EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol = `TCP`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Set(v string) error {
	switch v {
	case `TCP`:
		*f = EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "TCP"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol"
}

type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFqdn EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType = `FQDN`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Set(v string) error {
	switch v {
	case `FQDN`:
		*f = EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FQDN"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType"
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode struct {
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType `json:"log_only_mode_type,omitempty"`

	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType `json:"workloads,omitempty"`
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeAllServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType = `ALL_SERVICES`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeSelectedServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType = `SELECTED_SERVICES`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Set(v string) error {
	switch v {
	case `ALL_SERVICES`, `SELECTED_SERVICES`:
		*f = EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_SERVICES", "SELECTED_SERVICES"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType"
}

// The values should match the list of workloads used in networkconfig.proto
type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeDbsql EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType = `DBSQL`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeMlServing EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType = `ML_SERVING`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Set(v string) error {
	switch v {
	case `DBSQL`, `ML_SERVING`:
		*f = EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DBSQL", "ML_SERVING"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType"
}

// At which level can Databricks and Databricks managed compute access Internet.
// FULL_ACCESS: Databricks can access Internet. No blocking rules will apply.
// RESTRICTED_ACCESS: Databricks can only access explicitly allowed internet and
// storage destinations, as well as UC connections and external locations.
// PRIVATE_ACCESS_ONLY (not used): Databricks can only access destinations via
// private link.
type EgressNetworkPolicyInternetAccessPolicyRestrictionMode string

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeFullAccess EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `FULL_ACCESS`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModePrivateAccessOnly EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `PRIVATE_ACCESS_ONLY`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeRestrictedAccess EgressNetworkPolicyInternetAccessPolicyRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Set(v string) error {
	switch v {
	case `FULL_ACCESS`, `PRIVATE_ACCESS_ONLY`, `RESTRICTED_ACCESS`:
		*f = EgressNetworkPolicyInternetAccessPolicyRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_ACCESS", "PRIVATE_ACCESS_ONLY", "RESTRICTED_ACCESS"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyRestrictionMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyRestrictionMode"
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyInternetAccessPolicyStorageDestination struct {
	AllowedPaths []string `json:"allowed_paths,omitempty"`

	AzureContainer string `json:"azure_container,omitempty"`

	AzureDnsZone string `json:"azure_dns_zone,omitempty"`

	AzureStorageAccount string `json:"azure_storage_account,omitempty"`

	AzureStorageService string `json:"azure_storage_service,omitempty"`

	BucketName string `json:"bucket_name,omitempty"`

	Region string `json:"region,omitempty"`

	Type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EgressNetworkPolicyInternetAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EgressNetworkPolicyInternetAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType string

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAwsS3 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `AWS_S3`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `AZURE_STORAGE`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeCloudflareR2 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `CLOUDFLARE_R2`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType = `GOOGLE_CLOUD_STORAGE`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Set(v string) error {
	switch v {
	case `AWS_S3`, `AZURE_STORAGE`, `CLOUDFLARE_R2`, `GOOGLE_CLOUD_STORAGE`:
		*f = EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_S3", "AZURE_STORAGE", "CLOUDFLARE_R2", "GOOGLE_CLOUD_STORAGE"`, v)
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType"
}

type EmailConfig struct {
	// Email addresses to notify.
	Addresses []string `json:"addresses,omitempty"`
}

type Empty struct {
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnhancedSecurityMonitoring) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnhancedSecurityMonitoring) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EnhancedSecurityMonitoringSetting struct {
	// SHIELD feature: ESM
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoring `json:"enhanced_security_monitoring_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EnhancedSecurityMonitoringSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EnhancedSecurityMonitoringSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EsmEnablementAccount) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EsmEnablementAccount) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EsmEnablementAccountSetting struct {
	// Account level policy for ESM
	EsmEnablementAccount EsmEnablementAccount `json:"esm_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *EsmEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EsmEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential string `json:"credential,omitempty"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime int64 `json:"credentialEolTime,omitempty"`
	// User ID of the user that owns this token.
	OwnerId int64 `json:"ownerId,omitempty"`
	// The scopes of access granted in the token.
	Scopes []string `json:"scopes,omitempty"`
	// The type of this exchange token
	TokenType TokenType `json:"tokenType,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ExchangeToken) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExchangeToken) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId PartitionId `json:"partitionId"`
	// Array of scopes for the token request.
	Scopes []string `json:"scopes"`
	// A list of token types being requested
	TokenType []TokenType `json:"tokenType"`
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {
	Values []ExchangeToken `json:"values,omitempty"`
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

type GenericWebhookConfig struct {
	// [Input-Only][Optional] Password for webhook.
	Password string `json:"password,omitempty"`
	// [Output-Only] Whether password is set.
	PasswordSet bool `json:"password_set,omitempty"`
	// [Input-Only] URL for webhook.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`
	// [Input-Only][Optional] Username for webhook.
	Username string `json:"username,omitempty"`
	// [Output-Only] Whether username is set.
	UsernameSet bool `json:"username_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GenericWebhookConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GenericWebhookConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the account IP access toggle setting
type GetAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

// Retrieve the AI/BI dashboard embedding access policy
type GetAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Retrieve the list of domains approved to host embedded AI/BI dashboards
type GetAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the automatic cluster update setting
type GetAutomaticClusterUpdateSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAutomaticClusterUpdateSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the compliance security profile setting
type GetComplianceSecurityProfileSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetComplianceSecurityProfileSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the compliance security profile setting for new workspaces
type GetCspEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCspEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the default namespace setting
type GetDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Retrieve Legacy Access Disablement Status
type GetDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the disable legacy DBFS setting
type GetDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the disable legacy features setting
type GetDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the enhanced security monitoring setting
type GetEnhancedSecurityMonitoringSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEnhancedSecurityMonitoringSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the enhanced security monitoring setting for new workspaces
type GetEsmEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetEsmEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

// Get a notification destination
type GetNotificationDestinationRequest struct {
	Id string `json:"-" url:"-"`
}

// Get Personal Compute setting
type GetPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

// Get the restrict workspace admins setting
type GetRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Check configuration status
type GetStatusRequest struct {
	Keys string `json:"-" url:"keys"`
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount int `json:"address_count,omitempty"`
	// Creation timestamp in milliseconds.
	CreatedAt int64 `json:"created_at,omitempty"`
	// User ID of the user who created this list.
	CreatedBy int64 `json:"created_by,omitempty"`
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// Universally unique identifier (UUID) of the IP access list.
	ListId string `json:"list_id,omitempty"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list.
	UpdatedBy int64 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IpAccessListInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

type ListNccAzurePrivateEndpointRulesResponse struct {
	Items []NccAzurePrivateEndpointRule `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNccAzurePrivateEndpointRulesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNccAzurePrivateEndpointRulesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNetworkConnectivityConfigurationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNetworkConnectivityConfigurationsResponse struct {
	Items []NetworkConnectivityConfiguration `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNetworkConnectivityConfigurationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNetworkConnectivityConfigurationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List notification destinations
type ListNotificationDestinationsRequest struct {
	PageSize int64 `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNotificationDestinationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNotificationDestinationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNotificationDestinationsResponse struct {
	// Page token for next of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Results []ListNotificationDestinationsResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNotificationDestinationsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNotificationDestinationsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListNotificationDestinationsResult struct {
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListNotificationDestinationsResult) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListNotificationDestinationsResult) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPrivateEndpointRulesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPrivateEndpointRulesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos []PublicTokenInfo `json:"token_infos,omitempty"`
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById int64 `json:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTokenManagementRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTokenManagementRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

// Type of IP access list. Valid values are as follows and are case-sensitive:
//
// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block list.
// Exclude this IP or range. IP addresses in the block list are excluded even if
// they are included in an allow list.
type ListType string

// An allow list. Include this IP or range.
const ListTypeAllow ListType = `ALLOW`

// A block list. Exclude this IP or range. IP addresses in the block list are
// excluded even if they are included in an allow list.
const ListTypeBlock ListType = `BLOCK`

// String representation for [fmt.Print]
func (f *ListType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ListType) Set(v string) error {
	switch v {
	case `ALLOW`, `BLOCK`:
		*f = ListType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW", "BLOCK"`, v)
	}
}

// Type always returns ListType to satisfy [pflag.Value] interface
func (f *ListType) Type() string {
	return "ListType"
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] URL for Microsoft Teams.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *MicrosoftTeamsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s MicrosoftTeamsConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

type NccAzurePrivateEndpointRule struct {
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is `ESTABLISHED`. Remember
	// that you must approve new endpoints on your resources in the Azure portal
	// before they take effect.
	//
	// The possible values are: - INIT: (deprecated) The endpoint has been
	// created and pending approval. - PENDING: The endpoint has been created
	// and pending approval. - ESTABLISHED: The endpoint has been approved and
	// is ready to use in your serverless compute resources. - REJECTED:
	// Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up.
	ConnectionState NccAzurePrivateEndpointRuleConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// The name of the Azure private endpoint resource.
	EndpointName string `json:"endpoint_name,omitempty"`
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for `blob` and one for `dfs`.
	GroupId NccAzurePrivateEndpointRuleGroupId `json:"group_id,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id,omitempty"`
	// The ID of a private endpoint rule.
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NccAzurePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzurePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The current status of this private endpoint. The private endpoint rules are
// effective only if the connection state is `ESTABLISHED`. Remember that you
// must approve new endpoints on your resources in the Azure portal before they
// take effect.
//
// The possible values are: - INIT: (deprecated) The endpoint has been created
// and pending approval. - PENDING: The endpoint has been created and pending
// approval. - ESTABLISHED: The endpoint has been approved and is ready to use
// in your serverless compute resources. - REJECTED: Connection was rejected by
// the private link resource owner. - DISCONNECTED: Connection was removed by
// the private link resource owner, the private endpoint becomes informative and
// should be deleted for clean-up.
type NccAzurePrivateEndpointRuleConnectionState string

const NccAzurePrivateEndpointRuleConnectionStateDisconnected NccAzurePrivateEndpointRuleConnectionState = `DISCONNECTED`

const NccAzurePrivateEndpointRuleConnectionStateEstablished NccAzurePrivateEndpointRuleConnectionState = `ESTABLISHED`

const NccAzurePrivateEndpointRuleConnectionStateInit NccAzurePrivateEndpointRuleConnectionState = `INIT`

const NccAzurePrivateEndpointRuleConnectionStatePending NccAzurePrivateEndpointRuleConnectionState = `PENDING`

const NccAzurePrivateEndpointRuleConnectionStateRejected NccAzurePrivateEndpointRuleConnectionState = `REJECTED`

// String representation for [fmt.Print]
func (f *NccAzurePrivateEndpointRuleConnectionState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccAzurePrivateEndpointRuleConnectionState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `ESTABLISHED`, `INIT`, `PENDING`, `REJECTED`:
		*f = NccAzurePrivateEndpointRuleConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "ESTABLISHED", "INIT", "PENDING", "REJECTED"`, v)
	}
}

// Type always returns NccAzurePrivateEndpointRuleConnectionState to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleConnectionState) Type() string {
	return "NccAzurePrivateEndpointRuleConnectionState"
}

// The sub-resource type (group ID) of the target resource. Note that to connect
// to workspace root storage (root DBFS), you need two endpoints, one for `blob`
// and one for `dfs`.
type NccAzurePrivateEndpointRuleGroupId string

const NccAzurePrivateEndpointRuleGroupIdBlob NccAzurePrivateEndpointRuleGroupId = `blob`

const NccAzurePrivateEndpointRuleGroupIdDfs NccAzurePrivateEndpointRuleGroupId = `dfs`

const NccAzurePrivateEndpointRuleGroupIdMysqlServer NccAzurePrivateEndpointRuleGroupId = `mysqlServer`

const NccAzurePrivateEndpointRuleGroupIdSqlServer NccAzurePrivateEndpointRuleGroupId = `sqlServer`

// String representation for [fmt.Print]
func (f *NccAzurePrivateEndpointRuleGroupId) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccAzurePrivateEndpointRuleGroupId) Set(v string) error {
	switch v {
	case `blob`, `dfs`, `mysqlServer`, `sqlServer`:
		*f = NccAzurePrivateEndpointRuleGroupId(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "blob", "dfs", "mysqlServer", "sqlServer"`, v)
	}
}

// Type always returns NccAzurePrivateEndpointRuleGroupId to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleGroupId) Type() string {
	return "NccAzurePrivateEndpointRuleGroupId"
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets []string `json:"subnets,omitempty"`
	// The Azure region in which this service endpoint rule applies.
	TargetRegion string `json:"target_region,omitempty"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices []string `json:"target_services,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NccAzureServiceEndpointRule) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NccAzureServiceEndpointRule) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The network connectivity rules that apply to network traffic from your
// serverless compute resources.
type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules *NccEgressDefaultRules `json:"default_rules,omitempty"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules *NccEgressTargetRules `json:"target_rules,omitempty"`
}

// The network connectivity rules that are applied by default without resource
// specific configurations. You can find the stable network information of your
// serverless compute resources here.
type NccEgressDefaultRules struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule *NccAwsStableIpRule `json:"aws_stable_ip_rule,omitempty"`
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule *NccAzureServiceEndpointRule `json:"azure_service_endpoint_rule,omitempty"`
}

// The network connectivity rules that configured for each destinations. These
// rules override default rules.
type NccEgressTargetRules struct {
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule `json:"azure_private_endpoint_rules,omitempty"`
}

type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig *NccEgressConfig `json:"egress_config,omitempty"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// `^[0-9a-zA-Z-_]{3,30}$`.
	Name string `json:"name,omitempty"`
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string `json:"region,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type NotificationDestination struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	Config *Config `json:"config,omitempty"`
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *NotificationDestination) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s NotificationDestination) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey string `json:"integration_key,omitempty"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet bool `json:"integration_key_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PagerdutyConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PagerdutyConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PartitionId) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PartitionId) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value PersonalComputeMessageEnum `json:"value"`
}

// ON: Grants all users in all workspaces access to the Personal Compute default
// policy, allowing all users to create single-machine compute resources.
// DELEGATE: Moves access control for the Personal Compute default policy to
// individual workspaces and requires a workspace’s users or groups to be
// added to the ACLs of that workspace’s Personal Compute default policy
// before they will be able to create compute resources through that policy.
type PersonalComputeMessageEnum string

const PersonalComputeMessageEnumDelegate PersonalComputeMessageEnum = `DELEGATE`

const PersonalComputeMessageEnumOn PersonalComputeMessageEnum = `ON`

// String representation for [fmt.Print]
func (f *PersonalComputeMessageEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalComputeMessageEnum) Set(v string) error {
	switch v {
	case `DELEGATE`, `ON`:
		*f = PersonalComputeMessageEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELEGATE", "ON"`, v)
	}
}

// Type always returns PersonalComputeMessageEnum to satisfy [pflag.Value] interface
func (f *PersonalComputeMessageEnum) Type() string {
	return "PersonalComputeMessageEnum"
}

type PersonalComputeSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	PersonalCompute PersonalComputeMessage `json:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PersonalComputeSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment string `json:"comment,omitempty"`
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// The ID of this token.
	TokenId string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublicTokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublicTokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type"`
}

type ReplaceResponse struct {
}

type RestrictWorkspaceAdminsMessage struct {
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`
}

type RestrictWorkspaceAdminsMessageStatus string

const RestrictWorkspaceAdminsMessageStatusAllowAll RestrictWorkspaceAdminsMessageStatus = `ALLOW_ALL`

const RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs RestrictWorkspaceAdminsMessageStatus = `RESTRICT_TOKENS_AND_JOB_RUN_AS`

// String representation for [fmt.Print]
func (f *RestrictWorkspaceAdminsMessageStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RestrictWorkspaceAdminsMessageStatus) Set(v string) error {
	switch v {
	case `ALLOW_ALL`, `RESTRICT_TOKENS_AND_JOB_RUN_AS`:
		*f = RestrictWorkspaceAdminsMessageStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALLOW_ALL", "RESTRICT_TOKENS_AND_JOB_RUN_AS"`, v)
	}
}

// Type always returns RestrictWorkspaceAdminsMessageStatus to satisfy [pflag.Value] interface
func (f *RestrictWorkspaceAdminsMessageStatus) Type() string {
	return "RestrictWorkspaceAdminsMessageStatus"
}

type RestrictWorkspaceAdminsSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage `json:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *RestrictWorkspaceAdminsSetting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s RestrictWorkspaceAdminsSetting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string `json:"token_id"`
}

type RevokeTokenResponse struct {
}

type SetStatusResponse struct {
}

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SlackConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SlackConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type StringMessage struct {
	// Represents a generic string value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *StringMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s StringMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions []TokenPermission `json:"all_permissions,omitempty"`
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

func (s *TokenAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment string `json:"comment,omitempty"`
	// User ID of the user that created the token.
	CreatedById int64 `json:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"created_by_username,omitempty"`
	// Timestamp when the token was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Timestamp when the token expires.
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// Approximate timestamp for the day the token was last used. Accurate up to
	// 1 day.
	LastUsedDay int64 `json:"last_used_day,omitempty"`
	// User ID of the user that owns the token.
	OwnerId int64 `json:"owner_id,omitempty"`
	// ID of the token.
	TokenId string `json:"token_id,omitempty"`
	// If applicable, the ID of the workspace that the token was created in.
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type TokenPermissionLevel string

const TokenPermissionLevelCanUse TokenPermissionLevel = `CAN_USE`

// String representation for [fmt.Print]
func (f *TokenPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_USE`:
		*f = TokenPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_USE"`, v)
	}
}

// Type always returns TokenPermissionLevel to satisfy [pflag.Value] interface
func (f *TokenPermissionLevel) Type() string {
	return "TokenPermissionLevel"
}

type TokenPermissions struct {
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

// The type of token request. As of now, only `AZURE_ACTIVE_DIRECTORY_TOKEN` is
// supported.
type TokenType string

const TokenTypeArclightAzureExchangeToken TokenType = `ARCLIGHT_AZURE_EXCHANGE_TOKEN`

const TokenTypeArclightAzureExchangeTokenWithUserDelegationKey TokenType = `ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`

const TokenTypeAzureActiveDirectoryToken TokenType = `AZURE_ACTIVE_DIRECTORY_TOKEN`

// String representation for [fmt.Print]
func (f *TokenType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenType) Set(v string) error {
	switch v {
	case `ARCLIGHT_AZURE_EXCHANGE_TOKEN`, `ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`, `AZURE_ACTIVE_DIRECTORY_TOKEN`:
		*f = TokenType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARCLIGHT_AZURE_EXCHANGE_TOKEN", "ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY", "AZURE_ACTIVE_DIRECTORY_TOKEN"`, v)
	}
}

// Type always returns TokenType to satisfy [pflag.Value] interface
func (f *TokenType) Type() string {
	return "TokenType"
}

// Details required to update a setting.
type UpdateAccountIpAccessEnableRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting AccountIpAccessEnable `json:"setting"`
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting AibiDashboardEmbeddingAccessPolicySetting `json:"setting"`
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting AibiDashboardEmbeddingApprovedDomainsSetting `json:"setting"`
}

// Details required to update a setting.
type UpdateAutomaticClusterUpdateSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting AutomaticClusterUpdateSetting `json:"setting"`
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting ComplianceSecurityProfileSetting `json:"setting"`
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting CspEnablementAccountSetting `json:"setting"`
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`
	// This represents the setting configuration for the default namespace in
	// the Databricks workspace. Setting the default catalog for the workspace
	// determines the catalog that is used when queries do not reference a fully
	// qualified 3 level name. For example, if the default catalog is set to
	// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the
	// object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed). This setting requires a restart of clusters and SQL warehouses
	// to take effect. Additionally, the default namespace only applies when
	// using Unity Catalog-enabled compute.
	Setting DefaultNamespaceSetting `json:"setting"`
}

// Details required to update a setting.
type UpdateDisableLegacyAccessRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting DisableLegacyAccess `json:"setting"`
}

// Details required to update a setting.
type UpdateDisableLegacyDbfsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting DisableLegacyDbfs `json:"setting"`
}

// Details required to update a setting.
type UpdateDisableLegacyFeaturesRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting DisableLegacyFeatures `json:"setting"`
}

// Details required to update a setting.
type UpdateEnhancedSecurityMonitoringSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting EnhancedSecurityMonitoringSetting `json:"setting"`
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting EsmEnablementAccountSetting `json:"setting"`
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`

	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	Label string `json:"label,omitempty"`
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType `json:"list_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *Config `json:"config,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting PersonalComputeSetting `json:"setting"`
}

type UpdateResponse struct {
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	AllowMissing bool `json:"allow_missing"`
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
	FieldMask string `json:"field_mask"`

	Setting RestrictWorkspaceAdminsSetting `json:"setting"`
}

type WorkspaceConf map[string]string
