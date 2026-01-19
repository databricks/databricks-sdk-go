// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingsv2

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

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

// Values returns all possible values for AibiDashboardEmbeddingAccessPolicyAccessPolicyType.
//
// There is no guarantee on the order of the values in the slice.
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Values() []AibiDashboardEmbeddingAccessPolicyAccessPolicyType {
	return []AibiDashboardEmbeddingAccessPolicyAccessPolicyType{
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowAllDomains,
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains,
		AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeDenyAllDomains,
	}
}

// Type always returns AibiDashboardEmbeddingAccessPolicyAccessPolicyType to satisfy [pflag.Value] interface
func (f *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) Type() string {
	return "AibiDashboardEmbeddingAccessPolicyAccessPolicyType"
}

type AibiDashboardEmbeddingApprovedDomains struct {
	ApprovedDomains []string `json:"approved_domains,omitempty"`
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

// Values returns all possible values for ClusterAutoRestartMessageMaintenanceWindowDayOfWeek.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) Values() []ClusterAutoRestartMessageMaintenanceWindowDayOfWeek {
	return []ClusterAutoRestartMessageMaintenanceWindowDayOfWeek{
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFriday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekMonday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSaturday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekThursday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekTuesday,
		ClusterAutoRestartMessageMaintenanceWindowDayOfWeekWednesday,
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

// Values returns all possible values for ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency.
//
// There is no guarantee on the order of the values in the slice.
func (f *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) Values() []ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency {
	return []ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency{
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth,
		ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth,
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

type GetPublicAccountSettingRequest struct {
	Name string `json:"-" url:"-"`
}

type GetPublicAccountUserPreferenceRequest struct {
	// User Setting name.
	Name string `json:"-" url:"-"`
	// User ID of the user whose setting is being retrieved.
	UserId string `json:"-" url:"-"`
}

type GetPublicWorkspaceSettingRequest struct {
	// Name of the setting
	Name string `json:"-" url:"-"`
}

type IntegerMessage struct {
	Value int `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *IntegerMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s IntegerMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountSettingsMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListAccountSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListAccountSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountSettingsMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountSettingsMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountSettingsMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountSettingsMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountUserPreferencesMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListAccountUserPreferencesMetadataRequest` call. Provide this to
	// retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListAccountUserPreferencesMetadataRequest` must match the call that
	// provided the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// User ID of the user whose settings metadata is being retrieved.
	UserId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountUserPreferencesMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountUserPreferencesMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListAccountUserPreferencesMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountUserPreferencesMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountUserPreferencesMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceSettingsMetadataRequest struct {
	// The maximum number of settings to return. The service may return fewer
	// than this value. If unspecified, at most 200 settings will be returned.
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous
	// `ListWorkspaceSettingsMetadataRequest` call. Provide this to retrieve the
	// subsequent page.
	//
	// When paginating, all other parameters provided to
	// `ListWorkspaceSettingsMetadataRequest` must match the call that provided
	// the page token.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceSettingsMetadataRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceSettingsMetadataRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListWorkspaceSettingsMetadataResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of all settings available via public APIs and their metadata
	SettingsMetadata []SettingsMetadata `json:"settings_metadata,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListWorkspaceSettingsMetadataResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListWorkspaceSettingsMetadataResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PatchPublicAccountSettingRequest struct {
	Name string `json:"-" url:"-"`

	Setting Setting `json:"setting"`
}

type PatchPublicAccountUserPreferenceRequest struct {
	Name string `json:"-" url:"-"`

	Setting UserPreference `json:"setting"`
	// User ID of the user whose setting is being updated.
	UserId string `json:"-" url:"-"`
}

type PatchPublicWorkspaceSettingRequest struct {
	// Name of the setting
	Name string `json:"-" url:"-"`

	Setting Setting `json:"setting"`
}

type PersonalComputeMessage struct {
	Value PersonalComputeMessagePersonalComputeMessageEnum `json:"value,omitempty"`
}

// ON: Grants all users in all workspaces access to the Personal Compute default
// policy, allowing all users to create single-machine compute resources.
// DELEGATE: Moves access control for the Personal Compute default policy to
// individual workspaces and requires a workspace’s users or groups to be
// added to the ACLs of that workspace’s Personal Compute default policy
// before they will be able to create compute resources through that policy.
type PersonalComputeMessagePersonalComputeMessageEnum string

const PersonalComputeMessagePersonalComputeMessageEnumDelegate PersonalComputeMessagePersonalComputeMessageEnum = `DELEGATE`

const PersonalComputeMessagePersonalComputeMessageEnumOn PersonalComputeMessagePersonalComputeMessageEnum = `ON`

// String representation for [fmt.Print]
func (f *PersonalComputeMessagePersonalComputeMessageEnum) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PersonalComputeMessagePersonalComputeMessageEnum) Set(v string) error {
	switch v {
	case `DELEGATE`, `ON`:
		*f = PersonalComputeMessagePersonalComputeMessageEnum(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELEGATE", "ON"`, v)
	}
}

// Values returns all possible values for PersonalComputeMessagePersonalComputeMessageEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *PersonalComputeMessagePersonalComputeMessageEnum) Values() []PersonalComputeMessagePersonalComputeMessageEnum {
	return []PersonalComputeMessagePersonalComputeMessageEnum{
		PersonalComputeMessagePersonalComputeMessageEnumDelegate,
		PersonalComputeMessagePersonalComputeMessageEnumOn,
	}
}

// Type always returns PersonalComputeMessagePersonalComputeMessageEnum to satisfy [pflag.Value] interface
func (f *PersonalComputeMessagePersonalComputeMessageEnum) Type() string {
	return "PersonalComputeMessagePersonalComputeMessageEnum"
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

// Values returns all possible values for RestrictWorkspaceAdminsMessageStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RestrictWorkspaceAdminsMessageStatus) Values() []RestrictWorkspaceAdminsMessageStatus {
	return []RestrictWorkspaceAdminsMessageStatus{
		RestrictWorkspaceAdminsMessageStatusAllowAll,
		RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs,
	}
}

// Type always returns RestrictWorkspaceAdminsMessageStatus to satisfy [pflag.Value] interface
func (f *RestrictWorkspaceAdminsMessageStatus) Type() string {
	return "RestrictWorkspaceAdminsMessageStatus"
}

type Setting struct {
	// Setting value for aibi_dashboard_embedding_access_policy setting. This is
	// the setting value set by consumers, check
	// effective_aibi_dashboard_embedding_access_policy for final setting value.
	AibiDashboardEmbeddingAccessPolicy *AibiDashboardEmbeddingAccessPolicy `json:"aibi_dashboard_embedding_access_policy,omitempty"`
	// Setting value for aibi_dashboard_embedding_approved_domains setting. This
	// is the setting value set by consumers, check
	// effective_aibi_dashboard_embedding_approved_domains for final setting
	// value.
	AibiDashboardEmbeddingApprovedDomains *AibiDashboardEmbeddingApprovedDomains `json:"aibi_dashboard_embedding_approved_domains,omitempty"`
	// Setting value for automatic_cluster_update_workspace setting. This is the
	// setting value set by consumers, check
	// effective_automatic_cluster_update_workspace for final setting value.
	AutomaticClusterUpdateWorkspace *ClusterAutoRestartMessage `json:"automatic_cluster_update_workspace,omitempty"`
	// Setting value for boolean type setting. This is the setting value set by
	// consumers, check effective_boolean_val for final setting value.
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Effective setting value for aibi_dashboard_embedding_access_policy
	// setting. This is the final effective value of setting. To set a value use
	// aibi_dashboard_embedding_access_policy.
	EffectiveAibiDashboardEmbeddingAccessPolicy *AibiDashboardEmbeddingAccessPolicy `json:"effective_aibi_dashboard_embedding_access_policy,omitempty"`
	// Effective setting value for aibi_dashboard_embedding_approved_domains
	// setting. This is the final effective value of setting. To set a value use
	// aibi_dashboard_embedding_approved_domains.
	EffectiveAibiDashboardEmbeddingApprovedDomains *AibiDashboardEmbeddingApprovedDomains `json:"effective_aibi_dashboard_embedding_approved_domains,omitempty"`
	// Effective setting value for automatic_cluster_update_workspace setting.
	// This is the final effective value of setting. To set a value use
	// automatic_cluster_update_workspace.
	EffectiveAutomaticClusterUpdateWorkspace *ClusterAutoRestartMessage `json:"effective_automatic_cluster_update_workspace,omitempty"`
	// Effective setting value for boolean type setting. This is the final
	// effective value of setting. To set a value use boolean_val.
	EffectiveBooleanVal *BooleanMessage `json:"effective_boolean_val,omitempty"`
	// Effective setting value for integer type setting. This is the final
	// effective value of setting. To set a value use integer_val.
	EffectiveIntegerVal *IntegerMessage `json:"effective_integer_val,omitempty"`
	// Effective setting value for personal_compute setting. This is the final
	// effective value of setting. To set a value use personal_compute.
	EffectivePersonalCompute *PersonalComputeMessage `json:"effective_personal_compute,omitempty"`
	// Effective setting value for restrict_workspace_admins setting. This is
	// the final effective value of setting. To set a value use
	// restrict_workspace_admins.
	EffectiveRestrictWorkspaceAdmins *RestrictWorkspaceAdminsMessage `json:"effective_restrict_workspace_admins,omitempty"`
	// Effective setting value for string type setting. This is the final
	// effective value of setting. To set a value use string_val.
	EffectiveStringVal *StringMessage `json:"effective_string_val,omitempty"`
	// Setting value for integer type setting. This is the setting value set by
	// consumers, check effective_integer_val for final setting value.
	IntegerVal *IntegerMessage `json:"integer_val,omitempty"`
	// Name of the setting.
	Name string `json:"name,omitempty"`
	// Setting value for personal_compute setting. This is the setting value set
	// by consumers, check effective_personal_compute for final setting value.
	PersonalCompute *PersonalComputeMessage `json:"personal_compute,omitempty"`
	// Setting value for restrict_workspace_admins setting. This is the setting
	// value set by consumers, check effective_restrict_workspace_admins for
	// final setting value.
	RestrictWorkspaceAdmins *RestrictWorkspaceAdminsMessage `json:"restrict_workspace_admins,omitempty"`
	// Setting value for string type setting. This is the setting value set by
	// consumers, check effective_string_val for final setting value.
	StringVal *StringMessage `json:"string_val,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Setting) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Setting) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SettingsMetadata struct {
	// Setting description for what this setting controls
	Description string `json:"description,omitempty"`
	// Link to databricks documentation for the setting
	DocsLink string `json:"docs_link,omitempty"`
	// Name of the setting.
	Name string `json:"name,omitempty"`
	// Sample message depicting the type of the setting. To set this setting,
	// the value sent must match this type.
	Type string `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SettingsMetadata) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SettingsMetadata) MarshalJSON() ([]byte, error) {
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

// User Preference represents a user-specific setting scoped to an individual
// user within an account. Unlike workspace or account settings that apply to
// all users, user preferences allow personal customization (e.g., UI theme,
// editor preferences) without affecting other users.
type UserPreference struct {
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`

	EffectiveBooleanVal *BooleanMessage `json:"effective_boolean_val,omitempty"`

	EffectiveStringVal *StringMessage `json:"effective_string_val,omitempty"`
	// Name of the setting.
	Name string `json:"name,omitempty"`

	StringVal *StringMessage `json:"string_val,omitempty"`
	// User ID of the user.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UserPreference) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UserPreference) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
