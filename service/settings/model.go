// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"encoding/json"
	"fmt"
)

type AccountIpAccessEnable struct {

	// Wire name: 'acct_ip_acl_enable'
	AcctIpAclEnable BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *AccountIpAccessEnable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountIpAccessEnablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountIpAccessEnableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountIpAccessEnable) MarshalJSON() ([]byte, error) {
	pb, err := accountIpAccessEnableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AccountNetworkPolicy struct {
	// The associated account ID for this Network Policy object.
	// Wire name: 'account_id'
	AccountId string
	// The network policies applying for egress traffic.
	// Wire name: 'egress'
	Egress *NetworkPolicyEgress
	// The unique identifier for the network policy.
	// Wire name: 'network_policy_id'
	NetworkPolicyId string

	ForceSendFields []string `tf:"-"`
}

func (st *AccountNetworkPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &accountNetworkPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := accountNetworkPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AccountNetworkPolicy) MarshalJSON() ([]byte, error) {
	pb, err := accountNetworkPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AibiDashboardEmbeddingAccessPolicy struct {

	// Wire name: 'access_policy_type'
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType
}

func (st *AibiDashboardEmbeddingAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aibiDashboardEmbeddingAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aibiDashboardEmbeddingAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AibiDashboardEmbeddingAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := aibiDashboardEmbeddingAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'aibi_dashboard_embedding_access_policy'
	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicy
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *AibiDashboardEmbeddingAccessPolicySetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aibiDashboardEmbeddingAccessPolicySettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aibiDashboardEmbeddingAccessPolicySettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AibiDashboardEmbeddingAccessPolicySetting) MarshalJSON() ([]byte, error) {
	pb, err := aibiDashboardEmbeddingAccessPolicySettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AibiDashboardEmbeddingApprovedDomains struct {

	// Wire name: 'approved_domains'
	ApprovedDomains []string
}

func (st *AibiDashboardEmbeddingApprovedDomains) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aibiDashboardEmbeddingApprovedDomainsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aibiDashboardEmbeddingApprovedDomainsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AibiDashboardEmbeddingApprovedDomains) MarshalJSON() ([]byte, error) {
	pb, err := aibiDashboardEmbeddingApprovedDomainsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AibiDashboardEmbeddingApprovedDomainsSetting struct {

	// Wire name: 'aibi_dashboard_embedding_approved_domains'
	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomains
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *AibiDashboardEmbeddingApprovedDomainsSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aibiDashboardEmbeddingApprovedDomainsSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aibiDashboardEmbeddingApprovedDomainsSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AibiDashboardEmbeddingApprovedDomainsSetting) MarshalJSON() ([]byte, error) {
	pb, err := aibiDashboardEmbeddingApprovedDomainsSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AutomaticClusterUpdateSetting struct {

	// Wire name: 'automatic_cluster_update_workspace'
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *AutomaticClusterUpdateSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &automaticClusterUpdateSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := automaticClusterUpdateSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AutomaticClusterUpdateSetting) MarshalJSON() ([]byte, error) {
	pb, err := automaticClusterUpdateSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type BooleanMessage struct {

	// Wire name: 'value'
	Value bool

	ForceSendFields []string `tf:"-"`
}

func (st *BooleanMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &booleanMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := booleanMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BooleanMessage) MarshalJSON() ([]byte, error) {
	pb, err := booleanMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterAutoRestartMessage struct {

	// Wire name: 'can_toggle'
	CanToggle bool

	// Wire name: 'enabled'
	Enabled bool
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	// Wire name: 'enablement_details'
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails

	// Wire name: 'maintenance_window'
	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow

	// Wire name: 'restart_even_if_no_updates_available'
	RestartEvenIfNoUpdatesAvailable bool

	ForceSendFields []string `tf:"-"`
}

func (st *ClusterAutoRestartMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAutoRestartMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAutoRestartMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAutoRestartMessage) MarshalJSON() ([]byte, error) {
	pb, err := clusterAutoRestartMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	// Wire name: 'forced_for_compliance_mode'
	ForcedForComplianceMode bool
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	// Wire name: 'unavailable_for_disabled_entitlement'
	UnavailableForDisabledEntitlement bool
	// The feature is unavailable if the customer doesn't have enterprise tier
	// Wire name: 'unavailable_for_non_enterprise_tier'
	UnavailableForNonEnterpriseTier bool

	ForceSendFields []string `tf:"-"`
}

func (st *ClusterAutoRestartMessageEnablementDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAutoRestartMessageEnablementDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAutoRestartMessageEnablementDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAutoRestartMessageEnablementDetails) MarshalJSON() ([]byte, error) {
	pb, err := clusterAutoRestartMessageEnablementDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ClusterAutoRestartMessageMaintenanceWindow struct {

	// Wire name: 'week_day_based_schedule'
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
}

func (st *ClusterAutoRestartMessageMaintenanceWindow) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAutoRestartMessageMaintenanceWindowFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAutoRestartMessageMaintenanceWindow) MarshalJSON() ([]byte, error) {
	pb, err := clusterAutoRestartMessageMaintenanceWindowToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'day_of_week'
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek

	// Wire name: 'frequency'
	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency

	// Wire name: 'window_start_time'
	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
}

func (st *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) MarshalJSON() ([]byte, error) {
	pb, err := clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'hours'
	Hours int

	// Wire name: 'minutes'
	Minutes int

	ForceSendFields []string `tf:"-"`
}

func (st *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := clusterAutoRestartMessageMaintenanceWindowWindowStartTimeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) MarshalJSON() ([]byte, error) {
	pb, err := clusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Wire name: 'compliance_standards'
	ComplianceStandards []ComplianceStandard

	// Wire name: 'is_enabled'
	IsEnabled bool

	ForceSendFields []string `tf:"-"`
}

func (st *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &complianceSecurityProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := complianceSecurityProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	pb, err := complianceSecurityProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ComplianceSecurityProfileSetting struct {
	// SHIELD feature: CSP
	// Wire name: 'compliance_security_profile_workspace'
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfile
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *ComplianceSecurityProfileSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &complianceSecurityProfileSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := complianceSecurityProfileSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ComplianceSecurityProfileSetting) MarshalJSON() ([]byte, error) {
	pb, err := complianceSecurityProfileSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

const ComplianceStandardKFsi ComplianceStandard = `K_FSI`

const ComplianceStandardNone ComplianceStandard = `NONE`

const ComplianceStandardPciDss ComplianceStandard = `PCI_DSS`

// String representation for [fmt.Print]
func (f *ComplianceStandard) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ComplianceStandard) Set(v string) error {
	switch v {
	case `CANADA_PROTECTED_B`, `CYBER_ESSENTIAL_PLUS`, `FEDRAMP_HIGH`, `FEDRAMP_IL5`, `FEDRAMP_MODERATE`, `HIPAA`, `HITRUST`, `IRAP_PROTECTED`, `ISMAP`, `ITAR_EAR`, `K_FSI`, `NONE`, `PCI_DSS`:
		*f = ComplianceStandard(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANADA_PROTECTED_B", "CYBER_ESSENTIAL_PLUS", "FEDRAMP_HIGH", "FEDRAMP_IL5", "FEDRAMP_MODERATE", "HIPAA", "HITRUST", "IRAP_PROTECTED", "ISMAP", "ITAR_EAR", "K_FSI", "NONE", "PCI_DSS"`, v)
	}
}

// Type always returns ComplianceStandard to satisfy [pflag.Value] interface
func (f *ComplianceStandard) Type() string {
	return "ComplianceStandard"
}

type Config struct {

	// Wire name: 'email'
	Email *EmailConfig

	// Wire name: 'generic_webhook'
	GenericWebhook *GenericWebhookConfig

	// Wire name: 'microsoft_teams'
	MicrosoftTeams *MicrosoftTeamsConfig

	// Wire name: 'pagerduty'
	Pagerduty *PagerdutyConfig

	// Wire name: 'slack'
	Slack *SlackConfig
}

func (st *Config) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &configPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := configFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Config) MarshalJSON() ([]byte, error) {
	pb, err := configToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to configure a block list or allow list.
type CreateIpAccessList struct {

	// Wire name: 'ip_addresses'
	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	// Wire name: 'list_type'
	ListType ListType
}

func (st *CreateIpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createIpAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createIpAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateIpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := createIpAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {
	// Definition of an IP Access list
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo
}

func (st *CreateIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := createIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create a network connectivity configuration
type CreateNetworkConnectivityConfigRequest struct {
	// Properties of the new network connectivity configuration.
	// Wire name: 'network_connectivity_config'
	NetworkConnectivityConfig CreateNetworkConnectivityConfiguration
}

func (st *CreateNetworkConnectivityConfigRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createNetworkConnectivityConfigRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createNetworkConnectivityConfigRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateNetworkConnectivityConfigRequest) MarshalJSON() ([]byte, error) {
	pb, err := createNetworkConnectivityConfigRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Properties of the new network connectivity configuration.
type CreateNetworkConnectivityConfiguration struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	// Wire name: 'name'
	Name string
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	// Wire name: 'region'
	Region string
}

func (st *CreateNetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createNetworkConnectivityConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createNetworkConnectivityConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateNetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := createNetworkConnectivityConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create a network policy
type CreateNetworkPolicyRequest struct {

	// Wire name: 'network_policy'
	NetworkPolicy AccountNetworkPolicy
}

func (st *CreateNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	// Wire name: 'config'
	Config *Config
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := createNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	// Wire name: 'application_id'
	ApplicationId string
	// Comment that describes the purpose of the token.
	// Wire name: 'comment'
	Comment string
	// The number of seconds before the token expires.
	// Wire name: 'lifetime_seconds'
	LifetimeSeconds int64

	ForceSendFields []string `tf:"-"`
}

func (st *CreateOboTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createOboTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createOboTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := createOboTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {

	// Wire name: 'token_info'
	TokenInfo *TokenInfo
	// Value of the token.
	// Wire name: 'token_value'
	TokenValue string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createOboTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createOboTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := createOboTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type CreatePrivateEndpointRule struct {
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	// Wire name: 'domain_names'
	DomainNames []string
	// Only used by private endpoints to Azure first-party services. Enum: blob
	// | dfs | sqlServer | mysqlServer
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	// Wire name: 'group_id'
	GroupId string
	// The Azure resource ID of the target resource.
	// Wire name: 'resource_id'
	ResourceId string

	ForceSendFields []string `tf:"-"`
}

func (st *CreatePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := createPrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Create a private endpoint rule
type CreatePrivateEndpointRuleRequest struct {
	// Your Network Connectivity Configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `tf:"-"`
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
	// Wire name: 'private_endpoint_rule'
	PrivateEndpointRule CreatePrivateEndpointRule
}

func (st *CreatePrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := createPrivateEndpointRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	// Wire name: 'comment'
	Comment string
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	// Wire name: 'lifetime_seconds'
	LifetimeSeconds int64

	ForceSendFields []string `tf:"-"`
}

func (st *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := createTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateTokenResponse struct {
	// The information for the new token.
	// Wire name: 'token_info'
	TokenInfo *PublicTokenInfo
	// The value of the new token.
	// Wire name: 'token_value'
	TokenValue string

	ForceSendFields []string `tf:"-"`
}

func (st *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := createTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	// Wire name: 'compliance_standards'
	ComplianceStandards []ComplianceStandard
	// Enforced = it cannot be overriden at workspace level.
	// Wire name: 'is_enforced'
	IsEnforced bool

	ForceSendFields []string `tf:"-"`
}

func (st *CspEnablementAccount) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cspEnablementAccountPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cspEnablementAccountFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CspEnablementAccount) MarshalJSON() ([]byte, error) {
	pb, err := cspEnablementAccountToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CspEnablementAccountSetting struct {
	// Account level policy for CSP
	// Wire name: 'csp_enablement_account'
	CspEnablementAccount CspEnablementAccount
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *CspEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cspEnablementAccountSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cspEnablementAccountSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CspEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	pb, err := cspEnablementAccountSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string

	// Wire name: 'namespace'
	Namespace StringMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *DefaultNamespaceSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &defaultNamespaceSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := defaultNamespaceSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DefaultNamespaceSetting) MarshalJSON() ([]byte, error) {
	pb, err := defaultNamespaceSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountIpAccessEnableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountIpAccessEnableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountIpAccessEnableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteAccountIpAccessEnableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountIpAccessEnableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountIpAccessEnableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountIpAccessEnableResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountIpAccessEnableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	// Wire name: 'ip_access_list_id'
	IpAccessListId string `tf:"-"`
}

func (st *DeleteAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAibiDashboardEmbeddingAccessPolicySettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteAibiDashboardEmbeddingAccessPolicySettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDefaultNamespaceSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDefaultNamespaceSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDefaultNamespaceSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteDefaultNamespaceSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDefaultNamespaceSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDefaultNamespaceSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDefaultNamespaceSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDefaultNamespaceSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDisableLegacyAccessRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDisableLegacyAccessRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDisableLegacyAccessRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteDisableLegacyAccessResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDisableLegacyAccessResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDisableLegacyAccessResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDisableLegacyAccessResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDisableLegacyAccessResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDisableLegacyDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDisableLegacyDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDisableLegacyDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteDisableLegacyDbfsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDisableLegacyDbfsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDisableLegacyDbfsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDisableLegacyDbfsResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDisableLegacyDbfsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDisableLegacyFeaturesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDisableLegacyFeaturesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDisableLegacyFeaturesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteDisableLegacyFeaturesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDisableLegacyFeaturesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDisableLegacyFeaturesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDisableLegacyFeaturesResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDisableLegacyFeaturesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	// Wire name: 'ip_access_list_id'
	IpAccessListId string `tf:"-"`
}

func (st *DeleteIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete the enable partner powered AI features workspace setting
type DeleteLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteLlmProxyPartnerPoweredWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteLlmProxyPartnerPoweredWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The etag is returned.
type DeleteLlmProxyPartnerPoweredWorkspaceResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteLlmProxyPartnerPoweredWorkspaceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteLlmProxyPartnerPoweredWorkspaceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteLlmProxyPartnerPoweredWorkspaceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteLlmProxyPartnerPoweredWorkspaceResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteLlmProxyPartnerPoweredWorkspaceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `tf:"-"`
}

func (st *DeleteNetworkConnectivityConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteNetworkConnectivityConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteNetworkConnectivityConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteNetworkConnectivityConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteNetworkConnectivityConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteNetworkConnectivityConfigurationResponse struct {
}

func (st *DeleteNetworkConnectivityConfigurationResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteNetworkConnectivityConfigurationResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteNetworkConnectivityConfigurationResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteNetworkConnectivityConfigurationResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteNetworkConnectivityConfigurationResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a network policy
type DeleteNetworkPolicyRequest struct {
	// The unique identifier of the network policy to delete.
	// Wire name: 'network_policy_id'
	NetworkPolicyId string `tf:"-"`
}

func (st *DeleteNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteNetworkPolicyRpcResponse struct {
}

func (st *DeleteNetworkPolicyRpcResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteNetworkPolicyRpcResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteNetworkPolicyRpcResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteNetworkPolicyRpcResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteNetworkPolicyRpcResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a notification destination
type DeleteNotificationDestinationRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *DeleteNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeletePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePersonalComputeSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePersonalComputeSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := deletePersonalComputeSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeletePersonalComputeSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePersonalComputeSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePersonalComputeSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePersonalComputeSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := deletePersonalComputeSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `tf:"-"`
	// Your private endpoint rule ID.
	// Wire name: 'private_endpoint_rule_id'
	PrivateEndpointRuleId string `tf:"-"`
}

func (st *DeletePrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := deletePrivateEndpointRuleRequestToPb(&st)
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

// Delete the restrict workspace admins setting
type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *DeleteRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRestrictWorkspaceAdminsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRestrictWorkspaceAdminsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteRestrictWorkspaceAdminsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string
}

func (st *DeleteRestrictWorkspaceAdminsSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRestrictWorkspaceAdminsSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRestrictWorkspaceAdminsSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRestrictWorkspaceAdminsSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRestrictWorkspaceAdminsSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	// Wire name: 'token_id'
	TokenId string `tf:"-"`
}

func (st *DeleteTokenManagementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTokenManagementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTokenManagementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTokenManagementRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteTokenManagementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'disable_legacy_access'
	DisableLegacyAccess BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *DisableLegacyAccess) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &disableLegacyAccessPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := disableLegacyAccessFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DisableLegacyAccess) MarshalJSON() ([]byte, error) {
	pb, err := disableLegacyAccessToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DisableLegacyDbfs struct {

	// Wire name: 'disable_legacy_dbfs'
	DisableLegacyDbfs BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *DisableLegacyDbfs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &disableLegacyDbfsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := disableLegacyDbfsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DisableLegacyDbfs) MarshalJSON() ([]byte, error) {
	pb, err := disableLegacyDbfsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DisableLegacyFeatures struct {

	// Wire name: 'disable_legacy_features'
	DisableLegacyFeatures BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *DisableLegacyFeatures) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &disableLegacyFeaturesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := disableLegacyFeaturesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DisableLegacyFeatures) MarshalJSON() ([]byte, error) {
	pb, err := disableLegacyFeaturesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy struct {
	// The access policy enforced for egress traffic to the internet.
	// Wire name: 'internet_access'
	InternetAccess *EgressNetworkPolicyInternetAccessPolicy
}

func (st *EgressNetworkPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicy) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EgressNetworkPolicyInternetAccessPolicy struct {

	// Wire name: 'allowed_internet_destinations'
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestination

	// Wire name: 'allowed_storage_destinations'
	AllowedStorageDestinations []EgressNetworkPolicyInternetAccessPolicyStorageDestination
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	// Wire name: 'log_only_mode'
	LogOnlyMode *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
	// At which level can Databricks and Databricks managed compute access
	// Internet. FULL_ACCESS: Databricks can access Internet. No blocking rules
	// will apply. RESTRICTED_ACCESS: Databricks can only access explicitly
	// allowed internet and storage destinations, as well as UC connections and
	// external locations. PRIVATE_ACCESS_ONLY (not used): Databricks can only
	// access destinations via private link.
	// Wire name: 'restriction_mode'
	RestrictionMode EgressNetworkPolicyInternetAccessPolicyRestrictionMode
}

func (st *EgressNetworkPolicyInternetAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyInternetAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyInternetAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyInternetAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyInternetAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support domain name (FQDN) destinations for the time
// being, though going forwards we want to support host names and IP addresses.
type EgressNetworkPolicyInternetAccessPolicyInternetDestination struct {

	// Wire name: 'destination'
	Destination string
	// The filtering protocol used by the DP. For private and public preview,
	// SEG will only support TCP filtering (i.e. DNS based filtering, filtering
	// by destination IP address), so protocol will be set to TCP by default and
	// hidden from the user. In the future, users may be able to select HTTP
	// filtering (i.e. SNI based filtering, filtering by FQDN).
	// Wire name: 'protocol'
	Protocol EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol

	// Wire name: 'type'
	Type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType

	ForceSendFields []string `tf:"-"`
}

func (st *EgressNetworkPolicyInternetAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyInternetAccessPolicyInternetDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyInternetAccessPolicyInternetDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyInternetAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'log_only_mode_type'
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType

	// Wire name: 'workloads'
	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType
}

func (st *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyInternetAccessPolicyLogOnlyModePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyInternetAccessPolicyLogOnlyModeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'allowed_paths'
	AllowedPaths []string

	// Wire name: 'azure_container'
	AzureContainer string

	// Wire name: 'azure_dns_zone'
	AzureDnsZone string

	// Wire name: 'azure_storage_account'
	AzureStorageAccount string

	// Wire name: 'azure_storage_service'
	AzureStorageService string

	// Wire name: 'bucket_name'
	BucketName string

	// Wire name: 'region'
	Region string

	// Wire name: 'type'
	Type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType

	ForceSendFields []string `tf:"-"`
}

func (st *EgressNetworkPolicyInternetAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyInternetAccessPolicyStorageDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyInternetAccessPolicyStorageDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyInternetAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type EgressNetworkPolicyNetworkAccessPolicy struct {
	// List of internet destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	// Wire name: 'allowed_internet_destinations'
	AllowedInternetDestinations []EgressNetworkPolicyNetworkAccessPolicyInternetDestination
	// List of storage destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	// Wire name: 'allowed_storage_destinations'
	AllowedStorageDestinations []EgressNetworkPolicyNetworkAccessPolicyStorageDestination
	// Optional. When policy_enforcement is not provided, we default to
	// ENFORCE_MODE_ALL_SERVICES
	// Wire name: 'policy_enforcement'
	PolicyEnforcement *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement
	// The restriction mode that controls how serverless workloads can access
	// the internet.
	// Wire name: 'restriction_mode'
	RestrictionMode EgressNetworkPolicyNetworkAccessPolicyRestrictionMode
}

func (st *EgressNetworkPolicyNetworkAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyNetworkAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyNetworkAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyNetworkAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Users can specify accessible internet destinations when outbound access is
// restricted. We only support DNS_NAME (FQDN format) destinations for the time
// being. Going forward we may extend support to host names and IP addresses.
type EgressNetworkPolicyNetworkAccessPolicyInternetDestination struct {
	// The internet destination to which access will be allowed. Format
	// dependent on the destination type.
	// Wire name: 'destination'
	Destination string
	// The type of internet destination. Currently only DNS_NAME is supported.
	// Wire name: 'internet_destination_type'
	InternetDestinationType EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType

	ForceSendFields []string `tf:"-"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyInternetDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyNetworkAccessPolicyInternetDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyNetworkAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyNetworkAccessPolicyInternetDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType string

const EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeDnsName EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType = `DNS_NAME`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) Set(v string) error {
	switch v {
	case `DNS_NAME`:
		*f = EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DNS_NAME"`, v)
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType"
}

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement struct {
	// When empty, it means dry run for all products. When non-empty, it means
	// dry run for specific products and for the other products, they will run
	// in enforced mode.
	// Wire name: 'dry_run_mode_product_filter'
	DryRunModeProductFilter []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter
	// The mode of policy enforcement. ENFORCED blocks traffic that violates
	// policy, while DRY_RUN only logs violations without blocking. When not
	// specified, defaults to ENFORCED.
	// Wire name: 'enforcement_mode'
	EnforcementMode EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode
}

func (st *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The values should match the list of workloads used in networkconfig.proto
type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter string

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterDbsql EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter = `DBSQL`

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterMlServing EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter = `ML_SERVING`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) Set(v string) error {
	switch v {
	case `DBSQL`, `ML_SERVING`:
		*f = EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DBSQL", "ML_SERVING"`, v)
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter"
}

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode string

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeDryRun EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode = `DRY_RUN`

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeEnforced EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode = `ENFORCED`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) Set(v string) error {
	switch v {
	case `DRY_RUN`, `ENFORCED`:
		*f = EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DRY_RUN", "ENFORCED"`, v)
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode"
}

// At which level can Databricks and Databricks managed compute access Internet.
// FULL_ACCESS: Databricks can access Internet. No blocking rules will apply.
// RESTRICTED_ACCESS: Databricks can only access explicitly allowed internet and
// storage destinations, as well as UC connections and external locations.
type EgressNetworkPolicyNetworkAccessPolicyRestrictionMode string

const EgressNetworkPolicyNetworkAccessPolicyRestrictionModeFullAccess EgressNetworkPolicyNetworkAccessPolicyRestrictionMode = `FULL_ACCESS`

const EgressNetworkPolicyNetworkAccessPolicyRestrictionModeRestrictedAccess EgressNetworkPolicyNetworkAccessPolicyRestrictionMode = `RESTRICTED_ACCESS`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) Set(v string) error {
	switch v {
	case `FULL_ACCESS`, `RESTRICTED_ACCESS`:
		*f = EgressNetworkPolicyNetworkAccessPolicyRestrictionMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FULL_ACCESS", "RESTRICTED_ACCESS"`, v)
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyRestrictionMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyRestrictionMode"
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyNetworkAccessPolicyStorageDestination struct {
	// The Azure storage account name.
	// Wire name: 'azure_storage_account'
	AzureStorageAccount string
	// The Azure storage service type (blob, dfs, etc.).
	// Wire name: 'azure_storage_service'
	AzureStorageService string

	// Wire name: 'bucket_name'
	BucketName string
	// The region of the S3 bucket.
	// Wire name: 'region'
	Region string
	// The type of storage destination.
	// Wire name: 'storage_destination_type'
	StorageDestinationType EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType

	ForceSendFields []string `tf:"-"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyStorageDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := egressNetworkPolicyNetworkAccessPolicyStorageDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EgressNetworkPolicyNetworkAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	pb, err := egressNetworkPolicyNetworkAccessPolicyStorageDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType string

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAwsS3 EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType = `AWS_S3`

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType = `AZURE_STORAGE`

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType = `GOOGLE_CLOUD_STORAGE`

// String representation for [fmt.Print]
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) Set(v string) error {
	switch v {
	case `AWS_S3`, `AZURE_STORAGE`, `GOOGLE_CLOUD_STORAGE`:
		*f = EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_S3", "AZURE_STORAGE", "GOOGLE_CLOUD_STORAGE"`, v)
	}
}

// Type always returns EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) Type() string {
	return "EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType"
}

// The target resources that are supported by Network Connectivity Config. Note:
// some egress types can support general types that are not defined in
// EgressResourceType. E.g.: Azure private endpoint supports private link
// enabled Azure services.
type EgressResourceType string

const EgressResourceTypeAzureBlobStorage EgressResourceType = `AZURE_BLOB_STORAGE`

// String representation for [fmt.Print]
func (f *EgressResourceType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EgressResourceType) Set(v string) error {
	switch v {
	case `AZURE_BLOB_STORAGE`:
		*f = EgressResourceType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AZURE_BLOB_STORAGE"`, v)
	}
}

// Type always returns EgressResourceType to satisfy [pflag.Value] interface
func (f *EgressResourceType) Type() string {
	return "EgressResourceType"
}

type EmailConfig struct {
	// Email addresses to notify.
	// Wire name: 'addresses'
	Addresses []string
}

func (st *EmailConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &emailConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := emailConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EmailConfig) MarshalJSON() ([]byte, error) {
	pb, err := emailConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Empty struct {
}

func (st *Empty) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &emptyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := emptyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Empty) MarshalJSON() ([]byte, error) {
	pb, err := emptyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnableExportNotebook struct {

	// Wire name: 'boolean_val'
	BooleanVal *BooleanMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *EnableExportNotebook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enableExportNotebookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enableExportNotebookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnableExportNotebook) MarshalJSON() ([]byte, error) {
	pb, err := enableExportNotebookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnableNotebookTableClipboard struct {

	// Wire name: 'boolean_val'
	BooleanVal *BooleanMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *EnableNotebookTableClipboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enableNotebookTableClipboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enableNotebookTableClipboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnableNotebookTableClipboard) MarshalJSON() ([]byte, error) {
	pb, err := enableNotebookTableClipboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnableResultsDownloading struct {

	// Wire name: 'boolean_val'
	BooleanVal *BooleanMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *EnableResultsDownloading) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enableResultsDownloadingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enableResultsDownloadingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnableResultsDownloading) MarshalJSON() ([]byte, error) {
	pb, err := enableResultsDownloadingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {

	// Wire name: 'is_enabled'
	IsEnabled bool

	ForceSendFields []string `tf:"-"`
}

func (st *EnhancedSecurityMonitoring) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enhancedSecurityMonitoringPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enhancedSecurityMonitoringFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnhancedSecurityMonitoring) MarshalJSON() ([]byte, error) {
	pb, err := enhancedSecurityMonitoringToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EnhancedSecurityMonitoringSetting struct {
	// SHIELD feature: ESM
	// Wire name: 'enhanced_security_monitoring_workspace'
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoring
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *EnhancedSecurityMonitoringSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &enhancedSecurityMonitoringSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := enhancedSecurityMonitoringSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EnhancedSecurityMonitoringSetting) MarshalJSON() ([]byte, error) {
	pb, err := enhancedSecurityMonitoringSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Account level policy for ESM
type EsmEnablementAccount struct {

	// Wire name: 'is_enforced'
	IsEnforced bool

	ForceSendFields []string `tf:"-"`
}

func (st *EsmEnablementAccount) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &esmEnablementAccountPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := esmEnablementAccountFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EsmEnablementAccount) MarshalJSON() ([]byte, error) {
	pb, err := esmEnablementAccountToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EsmEnablementAccountSetting struct {
	// Account level policy for ESM
	// Wire name: 'esm_enablement_account'
	EsmEnablementAccount EsmEnablementAccount
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *EsmEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &esmEnablementAccountSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := esmEnablementAccountSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EsmEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	pb, err := esmEnablementAccountSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	// Wire name: 'credential'
	Credential string
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	// Wire name: 'credentialEolTime'
	CredentialEolTime int64
	// User ID of the user that owns this token.
	// Wire name: 'ownerId'
	OwnerId int64
	// The scopes of access granted in the token.
	// Wire name: 'scopes'
	Scopes []string
	// The type of this exchange token
	// Wire name: 'tokenType'
	TokenType TokenType

	ForceSendFields []string `tf:"-"`
}

func (st *ExchangeToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exchangeTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exchangeTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExchangeToken) MarshalJSON() ([]byte, error) {
	pb, err := exchangeTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	// Wire name: 'partitionId'
	PartitionId PartitionId
	// Array of scopes for the token request.
	// Wire name: 'scopes'
	Scopes []string
	// A list of token types being requested
	// Wire name: 'tokenType'
	TokenType []TokenType
}

func (st *ExchangeTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exchangeTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exchangeTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExchangeTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := exchangeTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {

	// Wire name: 'values'
	Values []ExchangeToken
}

func (st *ExchangeTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exchangeTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exchangeTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExchangeTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := exchangeTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {
	// Definition of an IP Access list
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo
}

func (st *FetchIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fetchIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fetchIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FetchIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := fetchIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GenericWebhookConfig struct {
	// [Input-Only][Optional] Password for webhook.
	// Wire name: 'password'
	Password string
	// [Output-Only] Whether password is set.
	// Wire name: 'password_set'
	PasswordSet bool
	// [Input-Only] URL for webhook.
	// Wire name: 'url'
	Url string
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet bool
	// [Input-Only][Optional] Username for webhook.
	// Wire name: 'username'
	Username string
	// [Output-Only] Whether username is set.
	// Wire name: 'username_set'
	UsernameSet bool

	ForceSendFields []string `tf:"-"`
}

func (st *GenericWebhookConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &genericWebhookConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := genericWebhookConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GenericWebhookConfig) MarshalJSON() ([]byte, error) {
	pb, err := genericWebhookConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountIpAccessEnableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountIpAccessEnableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountIpAccessEnableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	// Wire name: 'ip_access_list_id'
	IpAccessListId string `tf:"-"`
}

func (st *GetAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAutomaticClusterUpdateSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAutomaticClusterUpdateSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAutomaticClusterUpdateSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAutomaticClusterUpdateSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getComplianceSecurityProfileSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getComplianceSecurityProfileSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetComplianceSecurityProfileSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getComplianceSecurityProfileSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCspEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCspEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCspEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCspEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDefaultNamespaceSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDefaultNamespaceSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDefaultNamespaceSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDisableLegacyAccessRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDisableLegacyAccessRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDisableLegacyAccessRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDisableLegacyDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDisableLegacyDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDisableLegacyDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDisableLegacyFeaturesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDisableLegacyFeaturesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDisableLegacyFeaturesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEnhancedSecurityMonitoringSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEnhancedSecurityMonitoringSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEnhancedSecurityMonitoringSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getEnhancedSecurityMonitoringSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getEsmEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getEsmEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetEsmEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getEsmEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	// Wire name: 'ip_access_list_id'
	IpAccessListId string `tf:"-"`
}

func (st *GetIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := getIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo
}

func (st *GetIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := getIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {

	// Wire name: 'ip_access_lists'
	IpAccessLists []IpAccessListInfo
}

func (st *GetIpAccessListsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getIpAccessListsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getIpAccessListsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetIpAccessListsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getIpAccessListsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get the enable partner powered AI features account setting
type GetLlmProxyPartnerPoweredAccountRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetLlmProxyPartnerPoweredAccountRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLlmProxyPartnerPoweredAccountRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLlmProxyPartnerPoweredAccountRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLlmProxyPartnerPoweredAccountRequest) MarshalJSON() ([]byte, error) {
	pb, err := getLlmProxyPartnerPoweredAccountRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get the enforcement status of partner powered AI features account setting
type GetLlmProxyPartnerPoweredEnforceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetLlmProxyPartnerPoweredEnforceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLlmProxyPartnerPoweredEnforceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLlmProxyPartnerPoweredEnforceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLlmProxyPartnerPoweredEnforceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getLlmProxyPartnerPoweredEnforceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get the enable partner powered AI features workspace setting
type GetLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLlmProxyPartnerPoweredWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := getLlmProxyPartnerPoweredWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `tf:"-"`
}

func (st *GetNetworkConnectivityConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getNetworkConnectivityConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getNetworkConnectivityConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetNetworkConnectivityConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := getNetworkConnectivityConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a network policy
type GetNetworkPolicyRequest struct {
	// The unique identifier of the network policy to retrieve.
	// Wire name: 'network_policy_id'
	NetworkPolicyId string `tf:"-"`
}

func (st *GetNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get a notification destination
type GetNotificationDestinationRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st *GetNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := getNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetPersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPersonalComputeSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPersonalComputeSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPersonalComputeSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Gets a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `tf:"-"`
	// Your private endpoint rule ID.
	// Wire name: 'private_endpoint_rule_id'
	PrivateEndpointRuleId string `tf:"-"`
}

func (st *GetPrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPrivateEndpointRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *GetRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRestrictWorkspaceAdminsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRestrictWorkspaceAdminsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRestrictWorkspaceAdminsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Check configuration status
type GetStatusRequest struct {

	// Wire name: 'keys'
	Keys string `tf:"-"`
}

func (st *GetStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	// Wire name: 'token_id'
	TokenId string `tf:"-"`
}

func (st *GetTokenManagementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getTokenManagementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getTokenManagementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetTokenManagementRequest) MarshalJSON() ([]byte, error) {
	pb, err := getTokenManagementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []TokenPermissionsDescription
}

func (st *GetTokenPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getTokenPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getTokenPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetTokenPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getTokenPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {

	// Wire name: 'token_info'
	TokenInfo *TokenInfo
}

func (st *GetTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := getTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Get workspace network configuration
type GetWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`
}

func (st *GetWorkspaceNetworkOptionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getWorkspaceNetworkOptionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getWorkspaceNetworkOptionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetWorkspaceNetworkOptionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getWorkspaceNetworkOptionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	// Wire name: 'address_count'
	AddressCount int
	// Creation timestamp in milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64
	// User ID of the user who created this list.
	// Wire name: 'created_by'
	CreatedBy int64
	// Specifies whether this IP access list is enabled.
	// Wire name: 'enabled'
	Enabled bool

	// Wire name: 'ip_addresses'
	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string
	// Universally unique identifier (UUID) of the IP access list.
	// Wire name: 'list_id'
	ListId string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	// Wire name: 'list_type'
	ListType ListType
	// Update timestamp in milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64
	// User ID of the user who updated this list.
	// Wire name: 'updated_by'
	UpdatedBy int64

	ForceSendFields []string `tf:"-"`
}

func (st *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ipAccessListInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ipAccessListInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st IpAccessListInfo) MarshalJSON() ([]byte, error) {
	pb, err := ipAccessListInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {

	// Wire name: 'ip_access_lists'
	IpAccessLists []IpAccessListInfo
}

func (st *ListIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := listIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The private endpoint rule list was successfully retrieved.
type ListNccAzurePrivateEndpointRulesResponse struct {

	// Wire name: 'items'
	Items []NccAzurePrivateEndpointRule
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *ListNccAzurePrivateEndpointRulesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNccAzurePrivateEndpointRulesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNccAzurePrivateEndpointRulesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNccAzurePrivateEndpointRulesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listNccAzurePrivateEndpointRulesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListNetworkConnectivityConfigurationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNetworkConnectivityConfigurationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNetworkConnectivityConfigurationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNetworkConnectivityConfigurationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listNetworkConnectivityConfigurationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The network connectivity configuration list was successfully retrieved.
type ListNetworkConnectivityConfigurationsResponse struct {

	// Wire name: 'items'
	Items []NetworkConnectivityConfiguration
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *ListNetworkConnectivityConfigurationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNetworkConnectivityConfigurationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNetworkConnectivityConfigurationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNetworkConnectivityConfigurationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listNetworkConnectivityConfigurationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List network policies
type ListNetworkPoliciesRequest struct {
	// Pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListNetworkPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNetworkPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNetworkPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNetworkPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listNetworkPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListNetworkPoliciesResponse struct {
	// List of network policies.
	// Wire name: 'items'
	Items []AccountNetworkPolicy
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func (st *ListNetworkPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNetworkPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNetworkPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNetworkPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listNetworkPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List notification destinations
type ListNotificationDestinationsRequest struct {

	// Wire name: 'page_size'
	PageSize int64 `tf:"-"`

	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListNotificationDestinationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNotificationDestinationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNotificationDestinationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNotificationDestinationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listNotificationDestinationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListNotificationDestinationsResponse struct {
	// Page token for next of results.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'results'
	Results []ListNotificationDestinationsResult

	ForceSendFields []string `tf:"-"`
}

func (st *ListNotificationDestinationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNotificationDestinationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNotificationDestinationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNotificationDestinationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listNotificationDestinationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListNotificationDestinationsResult struct {
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	// Wire name: 'destination_type'
	DestinationType DestinationType
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying notification destination.
	// Wire name: 'id'
	Id string

	ForceSendFields []string `tf:"-"`
}

func (st *ListNotificationDestinationsResult) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listNotificationDestinationsResultPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listNotificationDestinationsResultFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListNotificationDestinationsResult) MarshalJSON() ([]byte, error) {
	pb, err := listNotificationDestinationsResultToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `tf:"-"`
	// Pagination token to go to next page based on previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListPrivateEndpointRulesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPrivateEndpointRulesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPrivateEndpointRulesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPrivateEndpointRulesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listPrivateEndpointRulesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	// Wire name: 'token_infos'
	TokenInfos []PublicTokenInfo
}

func (st *ListPublicTokensResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPublicTokensResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPublicTokensResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPublicTokensResponse) MarshalJSON() ([]byte, error) {
	pb, err := listPublicTokensResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	// Wire name: 'created_by_id'
	CreatedById int64 `tf:"-"`
	// Username of the user that created the token.
	// Wire name: 'created_by_username'
	CreatedByUsername string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *ListTokenManagementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTokenManagementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTokenManagementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTokenManagementRequest) MarshalJSON() ([]byte, error) {
	pb, err := listTokenManagementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	// Wire name: 'token_infos'
	TokenInfos []TokenInfo
}

func (st *ListTokensResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTokensResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTokensResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTokensResponse) MarshalJSON() ([]byte, error) {
	pb, err := listTokensResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type LlmProxyPartnerPoweredAccount struct {

	// Wire name: 'boolean_val'
	BooleanVal BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *LlmProxyPartnerPoweredAccount) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &llmProxyPartnerPoweredAccountPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := llmProxyPartnerPoweredAccountFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LlmProxyPartnerPoweredAccount) MarshalJSON() ([]byte, error) {
	pb, err := llmProxyPartnerPoweredAccountToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LlmProxyPartnerPoweredEnforce struct {

	// Wire name: 'boolean_val'
	BooleanVal BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *LlmProxyPartnerPoweredEnforce) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &llmProxyPartnerPoweredEnforcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := llmProxyPartnerPoweredEnforceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LlmProxyPartnerPoweredEnforce) MarshalJSON() ([]byte, error) {
	pb, err := llmProxyPartnerPoweredEnforceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LlmProxyPartnerPoweredWorkspace struct {

	// Wire name: 'boolean_val'
	BooleanVal BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *LlmProxyPartnerPoweredWorkspace) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &llmProxyPartnerPoweredWorkspacePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := llmProxyPartnerPoweredWorkspaceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LlmProxyPartnerPoweredWorkspace) MarshalJSON() ([]byte, error) {
	pb, err := llmProxyPartnerPoweredWorkspaceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] URL for Microsoft Teams.
	// Wire name: 'url'
	Url string
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet bool

	ForceSendFields []string `tf:"-"`
}

func (st *MicrosoftTeamsConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &microsoftTeamsConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := microsoftTeamsConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MicrosoftTeamsConfig) MarshalJSON() ([]byte, error) {
	pb, err := microsoftTeamsConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	// Wire name: 'cidr_blocks'
	CidrBlocks []string
}

func (st *NccAwsStableIpRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nccAwsStableIpRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nccAwsStableIpRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NccAwsStableIpRule) MarshalJSON() ([]byte, error) {
	pb, err := nccAwsStableIpRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type NccAzurePrivateEndpointRule struct {
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the Azure portal
	// before they take effect. The possible values are: - INIT: (deprecated)
	// The endpoint has been created and pending approval. - PENDING: The
	// endpoint has been created and pending approval. - ESTABLISHED: The
	// endpoint has been approved and is ready to use in your serverless compute
	// resources. - REJECTED: Connection was rejected by the private link
	// resource owner. - DISCONNECTED: Connection was removed by the private
	// link resource owner, the private endpoint becomes informative and should
	// be deleted for clean-up. - EXPIRED: If the endpoint was created but not
	// approved in 14 days, it will be EXPIRED.
	// Wire name: 'connection_state'
	ConnectionState NccAzurePrivateEndpointRuleConnectionState
	// Time in epoch milliseconds when this object was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Whether this private endpoint is deactivated.
	// Wire name: 'deactivated'
	Deactivated bool
	// Time in epoch milliseconds when this object was deactivated.
	// Wire name: 'deactivated_at'
	DeactivatedAt int64
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	// Wire name: 'domain_names'
	DomainNames []string
	// The name of the Azure private endpoint resource.
	// Wire name: 'endpoint_name'
	EndpointName string
	// Only used by private endpoints to Azure first-party services. Enum: blob
	// | dfs | sqlServer | mysqlServer
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	// Wire name: 'group_id'
	GroupId string
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string
	// The Azure resource ID of the target resource.
	// Wire name: 'resource_id'
	ResourceId string
	// The ID of a private endpoint rule.
	// Wire name: 'rule_id'
	RuleId string
	// Time in epoch milliseconds when this object was updated.
	// Wire name: 'updated_time'
	UpdatedTime int64

	ForceSendFields []string `tf:"-"`
}

func (st *NccAzurePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nccAzurePrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nccAzurePrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NccAzurePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := nccAzurePrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NccAzurePrivateEndpointRuleConnectionState string

const NccAzurePrivateEndpointRuleConnectionStateDisconnected NccAzurePrivateEndpointRuleConnectionState = `DISCONNECTED`

const NccAzurePrivateEndpointRuleConnectionStateEstablished NccAzurePrivateEndpointRuleConnectionState = `ESTABLISHED`

const NccAzurePrivateEndpointRuleConnectionStateExpired NccAzurePrivateEndpointRuleConnectionState = `EXPIRED`

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
	case `DISCONNECTED`, `ESTABLISHED`, `EXPIRED`, `INIT`, `PENDING`, `REJECTED`:
		*f = NccAzurePrivateEndpointRuleConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "ESTABLISHED", "EXPIRED", "INIT", "PENDING", "REJECTED"`, v)
	}
}

// Type always returns NccAzurePrivateEndpointRuleConnectionState to satisfy [pflag.Value] interface
func (f *NccAzurePrivateEndpointRuleConnectionState) Type() string {
	return "NccAzurePrivateEndpointRuleConnectionState"
}

// The stable Azure service endpoints. You can configure the firewall of your
// Azure resources to allow traffic from your Databricks serverless compute
// resources.
type NccAzureServiceEndpointRule struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	// Wire name: 'subnets'
	Subnets []string
	// The Azure region in which this service endpoint rule applies..
	// Wire name: 'target_region'
	TargetRegion string
	// The Azure services to which this service endpoint rule applies to.
	// Wire name: 'target_services'
	TargetServices []EgressResourceType

	ForceSendFields []string `tf:"-"`
}

func (st *NccAzureServiceEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nccAzureServiceEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nccAzureServiceEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NccAzureServiceEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := nccAzureServiceEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	// Wire name: 'default_rules'
	DefaultRules *NccEgressDefaultRules
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	// Wire name: 'target_rules'
	TargetRules *NccEgressTargetRules
}

func (st *NccEgressConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nccEgressConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nccEgressConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NccEgressConfig) MarshalJSON() ([]byte, error) {
	pb, err := nccEgressConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Default rules don't have specific targets.
type NccEgressDefaultRules struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	// Wire name: 'aws_stable_ip_rule'
	AwsStableIpRule *NccAwsStableIpRule
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	// Wire name: 'azure_service_endpoint_rule'
	AzureServiceEndpointRule *NccAzureServiceEndpointRule
}

func (st *NccEgressDefaultRules) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nccEgressDefaultRulesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nccEgressDefaultRulesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NccEgressDefaultRules) MarshalJSON() ([]byte, error) {
	pb, err := nccEgressDefaultRulesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Target rule controls the egress rules that are dedicated to specific
// resources.
type NccEgressTargetRules struct {

	// Wire name: 'azure_private_endpoint_rules'
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule
}

func (st *NccEgressTargetRules) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nccEgressTargetRulesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nccEgressTargetRulesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NccEgressTargetRules) MarshalJSON() ([]byte, error) {
	pb, err := nccEgressTargetRulesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Properties of the new network connectivity configuration.
type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	// Wire name: 'account_id'
	AccountId string
	// Time in epoch milliseconds when this object was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	// Wire name: 'egress_config'
	EgressConfig *NccEgressConfig
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	// Wire name: 'name'
	Name string
	// Databricks network connectivity configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	// Wire name: 'region'
	Region string
	// Time in epoch milliseconds when this object was updated.
	// Wire name: 'updated_time'
	UpdatedTime int64

	ForceSendFields []string `tf:"-"`
}

func (st *NetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkConnectivityConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkConnectivityConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := networkConnectivityConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto). This policy should be consistent
// with [[com.databricks.api.proto.settingspolicy.EgressNetworkPolicy]]. Details
// see API-design:
// https://docs.google.com/document/d/1DKWO_FpZMCY4cF2O62LpwII1lx8gsnDGG-qgE3t3TOA/
type NetworkPolicyEgress struct {
	// The access policy enforced for egress traffic to the internet.
	// Wire name: 'network_access'
	NetworkAccess *EgressNetworkPolicyNetworkAccessPolicy
}

func (st *NetworkPolicyEgress) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &networkPolicyEgressPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := networkPolicyEgressFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NetworkPolicyEgress) MarshalJSON() ([]byte, error) {
	pb, err := networkPolicyEgressToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NotificationDestination struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	// Wire name: 'config'
	Config *Config
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	// Wire name: 'destination_type'
	DestinationType DestinationType
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying notification destination.
	// Wire name: 'id'
	Id string

	ForceSendFields []string `tf:"-"`
}

func (st *NotificationDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &notificationDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := notificationDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NotificationDestination) MarshalJSON() ([]byte, error) {
	pb, err := notificationDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	// Wire name: 'integration_key'
	IntegrationKey string
	// [Output-Only] Whether integration key is set.
	// Wire name: 'integration_key_set'
	IntegrationKeySet bool

	ForceSendFields []string `tf:"-"`
}

func (st *PagerdutyConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &pagerdutyConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := pagerdutyConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PagerdutyConfig) MarshalJSON() ([]byte, error) {
	pb, err := pagerdutyConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	// Wire name: 'workspaceId'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func (st *PartitionId) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &partitionIdPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := partitionIdFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PartitionId) MarshalJSON() ([]byte, error) {
	pb, err := partitionIdToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	// Wire name: 'value'
	Value PersonalComputeMessageEnum
}

func (st *PersonalComputeMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &personalComputeMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := personalComputeMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PersonalComputeMessage) MarshalJSON() ([]byte, error) {
	pb, err := personalComputeMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string

	// Wire name: 'personal_compute'
	PersonalCompute PersonalComputeMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *PersonalComputeSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &personalComputeSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := personalComputeSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	pb, err := personalComputeSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	// Wire name: 'comment'
	Comment string
	// Server time (in epoch milliseconds) when the token was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	// Wire name: 'expiry_time'
	ExpiryTime int64
	// The ID of this token.
	// Wire name: 'token_id'
	TokenId string

	ForceSendFields []string `tf:"-"`
}

func (st *PublicTokenInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &publicTokenInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := publicTokenInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PublicTokenInfo) MarshalJSON() ([]byte, error) {
	pb, err := publicTokenInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	// Wire name: 'enabled'
	Enabled bool
	// The ID for the corresponding IP access list
	// Wire name: 'ip_access_list_id'
	IpAccessListId string `tf:"-"`

	// Wire name: 'ip_addresses'
	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	// Wire name: 'list_type'
	ListType ListType
}

func (st *ReplaceIpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &replaceIpAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := replaceIpAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReplaceIpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := replaceIpAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ReplaceResponse struct {
}

func (st *ReplaceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &replaceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := replaceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReplaceResponse) MarshalJSON() ([]byte, error) {
	pb, err := replaceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RestrictWorkspaceAdminsMessage struct {

	// Wire name: 'status'
	Status RestrictWorkspaceAdminsMessageStatus
}

func (st *RestrictWorkspaceAdminsMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restrictWorkspaceAdminsMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restrictWorkspaceAdminsMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestrictWorkspaceAdminsMessage) MarshalJSON() ([]byte, error) {
	pb, err := restrictWorkspaceAdminsMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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
	// Wire name: 'etag'
	Etag string

	// Wire name: 'restrict_workspace_admins'
	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string

	ForceSendFields []string `tf:"-"`
}

func (st *RestrictWorkspaceAdminsSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restrictWorkspaceAdminsSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restrictWorkspaceAdminsSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestrictWorkspaceAdminsSetting) MarshalJSON() ([]byte, error) {
	pb, err := restrictWorkspaceAdminsSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	// Wire name: 'token_id'
	TokenId string
}

func (st *RevokeTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &revokeTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := revokeTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RevokeTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := revokeTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type RevokeTokenResponse struct {
}

func (st *RevokeTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &revokeTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := revokeTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RevokeTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := revokeTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SetStatusResponse struct {
}

func (st *SetStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := setStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	// Wire name: 'url'
	Url string
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet bool

	ForceSendFields []string `tf:"-"`
}

func (st *SlackConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &slackConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := slackConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SlackConfig) MarshalJSON() ([]byte, error) {
	pb, err := slackConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StringMessage struct {
	// Represents a generic string value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func (st *StringMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &stringMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := stringMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StringMessage) MarshalJSON() ([]byte, error) {
	pb, err := stringMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TokenAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *TokenAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := tokenAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TokenAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []TokenPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func (st *TokenAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := tokenAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	// Wire name: 'comment'
	Comment string
	// User ID of the user that created the token.
	// Wire name: 'created_by_id'
	CreatedById int64
	// Username of the user that created the token.
	// Wire name: 'created_by_username'
	CreatedByUsername string
	// Timestamp when the token was created.
	// Wire name: 'creation_time'
	CreationTime int64
	// Timestamp when the token expires.
	// Wire name: 'expiry_time'
	ExpiryTime int64
	// Approximate timestamp for the day the token was last used. Accurate up to
	// 1 day.
	// Wire name: 'last_used_day'
	LastUsedDay int64
	// User ID of the user that owns the token.
	// Wire name: 'owner_id'
	OwnerId int64
	// ID of the token.
	// Wire name: 'token_id'
	TokenId string
	// If applicable, the ID of the workspace that the token was created in.
	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func (st *TokenInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenInfo) MarshalJSON() ([]byte, error) {
	pb, err := tokenInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TokenPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *TokenPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenPermission) MarshalJSON() ([]byte, error) {
	pb, err := tokenPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

	// Wire name: 'access_control_list'
	AccessControlList []TokenAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func (st *TokenPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenPermissions) MarshalJSON() ([]byte, error) {
	pb, err := tokenPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TokenPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func (st *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := tokenPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type TokenPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []TokenAccessControlRequest
}

func (st *TokenPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := tokenPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The type of token request. As of now, only `AZURE_ACTIVE_DIRECTORY_TOKEN` is
// supported.
type TokenType string

const TokenTypeArclightAzureExchangeToken TokenType = `ARCLIGHT_AZURE_EXCHANGE_TOKEN`

const TokenTypeArclightAzureExchangeTokenWithUserDelegationKey TokenType = `ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`

const TokenTypeArclightMultiTenantAzureExchangeToken TokenType = `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN`

const TokenTypeArclightMultiTenantAzureExchangeTokenWithUserDelegationKey TokenType = `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`

const TokenTypeAzureActiveDirectoryToken TokenType = `AZURE_ACTIVE_DIRECTORY_TOKEN`

// String representation for [fmt.Print]
func (f *TokenType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *TokenType) Set(v string) error {
	switch v {
	case `ARCLIGHT_AZURE_EXCHANGE_TOKEN`, `ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`, `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN`, `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`, `AZURE_ACTIVE_DIRECTORY_TOKEN`:
		*f = TokenType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ARCLIGHT_AZURE_EXCHANGE_TOKEN", "ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY", "ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN", "ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY", "AZURE_ACTIVE_DIRECTORY_TOKEN"`, v)
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
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting AccountIpAccessEnable
}

func (st *UpdateAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAccountIpAccessEnableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAccountIpAccessEnableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateAccountIpAccessEnableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting AibiDashboardEmbeddingAccessPolicySetting
}

func (st *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting AibiDashboardEmbeddingApprovedDomainsSetting
}

func (st *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateAutomaticClusterUpdateSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting AutomaticClusterUpdateSetting
}

func (st *UpdateAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAutomaticClusterUpdateSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAutomaticClusterUpdateSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAutomaticClusterUpdateSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateAutomaticClusterUpdateSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateComplianceSecurityProfileSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting ComplianceSecurityProfileSetting
}

func (st *UpdateComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateComplianceSecurityProfileSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateComplianceSecurityProfileSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateComplianceSecurityProfileSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateComplianceSecurityProfileSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateCspEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting CspEnablementAccountSetting
}

func (st *UpdateCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCspEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCspEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCspEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateCspEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateDefaultNamespaceSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string
	// This represents the setting configuration for the default namespace in
	// the Databricks workspace. Setting the default catalog for the workspace
	// determines the catalog that is used when queries do not reference a fully
	// qualified 3 level name. For example, if the default catalog is set to
	// 'retail_prod' then a query 'SELECT * FROM myTable' would reference the
	// object 'retail_prod.default.myTable' (the schema 'default' is always
	// assumed). This setting requires a restart of clusters and SQL warehouses
	// to take effect. Additionally, the default namespace only applies when
	// using Unity Catalog-enabled compute.
	// Wire name: 'setting'
	Setting DefaultNamespaceSetting
}

func (st *UpdateDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateDefaultNamespaceSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateDefaultNamespaceSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateDefaultNamespaceSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateDisableLegacyAccessRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting DisableLegacyAccess
}

func (st *UpdateDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateDisableLegacyAccessRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateDisableLegacyAccessRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateDisableLegacyAccessRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateDisableLegacyDbfsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting DisableLegacyDbfs
}

func (st *UpdateDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateDisableLegacyDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateDisableLegacyDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateDisableLegacyDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateDisableLegacyFeaturesRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting DisableLegacyFeatures
}

func (st *UpdateDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateDisableLegacyFeaturesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateDisableLegacyFeaturesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateDisableLegacyFeaturesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateEnableExportNotebookRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting EnableExportNotebook
}

func (st *UpdateEnableExportNotebookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateEnableExportNotebookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateEnableExportNotebookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateEnableExportNotebookRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateEnableExportNotebookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateEnableNotebookTableClipboardRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting EnableNotebookTableClipboard
}

func (st *UpdateEnableNotebookTableClipboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateEnableNotebookTableClipboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateEnableNotebookTableClipboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateEnableNotebookTableClipboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateEnableNotebookTableClipboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateEnableResultsDownloadingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting EnableResultsDownloading
}

func (st *UpdateEnableResultsDownloadingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateEnableResultsDownloadingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateEnableResultsDownloadingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateEnableResultsDownloadingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateEnableResultsDownloadingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateEnhancedSecurityMonitoringSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting EnhancedSecurityMonitoringSetting
}

func (st *UpdateEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateEnhancedSecurityMonitoringSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateEnhancedSecurityMonitoringSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateEnhancedSecurityMonitoringSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateEnhancedSecurityMonitoringSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateEsmEnablementAccountSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting EsmEnablementAccountSetting
}

func (st *UpdateEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateEsmEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateEsmEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateEsmEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateEsmEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update an IP access list.
type UpdateIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	// Wire name: 'enabled'
	Enabled bool
	// The ID for the corresponding IP access list
	// Wire name: 'ip_access_list_id'
	IpAccessListId string `tf:"-"`

	// Wire name: 'ip_addresses'
	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	// Wire name: 'list_type'
	ListType ListType

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateIpAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateIpAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := updateIpAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredAccountRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredAccount
}

func (st *UpdateLlmProxyPartnerPoweredAccountRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateLlmProxyPartnerPoweredAccountRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateLlmProxyPartnerPoweredAccountRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateLlmProxyPartnerPoweredAccountRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateLlmProxyPartnerPoweredAccountRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredEnforceRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredEnforce
}

func (st *UpdateLlmProxyPartnerPoweredEnforceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateLlmProxyPartnerPoweredEnforceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateLlmProxyPartnerPoweredEnforceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateLlmProxyPartnerPoweredEnforceRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateLlmProxyPartnerPoweredEnforceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateLlmProxyPartnerPoweredWorkspaceRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredWorkspace
}

func (st *UpdateLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateLlmProxyPartnerPoweredWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateLlmProxyPartnerPoweredWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Update a private endpoint rule
type UpdateNccAzurePrivateEndpointRulePublicRequest struct {
	// Your Network Connectivity Configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `tf:"-"`
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
	// Wire name: 'private_endpoint_rule'
	PrivateEndpointRule UpdatePrivateEndpointRule
	// Your private endpoint rule ID.
	// Wire name: 'private_endpoint_rule_id'
	PrivateEndpointRuleId string `tf:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	// Wire name: 'update_mask'
	UpdateMask string `tf:"-"`
}

func (st *UpdateNccAzurePrivateEndpointRulePublicRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateNccAzurePrivateEndpointRulePublicRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateNccAzurePrivateEndpointRulePublicRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateNccAzurePrivateEndpointRulePublicRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateNccAzurePrivateEndpointRulePublicRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Update a network policy
type UpdateNetworkPolicyRequest struct {

	// Wire name: 'network_policy'
	NetworkPolicy AccountNetworkPolicy
	// The unique identifier for the network policy.
	// Wire name: 'network_policy_id'
	NetworkPolicyId string `tf:"-"`
}

func (st *UpdateNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	// Wire name: 'config'
	Config *Config
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string
	// UUID identifying notification destination.
	// Wire name: 'id'
	Id string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func (st *UpdateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting PersonalComputeSetting
}

func (st *UpdatePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePersonalComputeSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePersonalComputeSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updatePersonalComputeSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type UpdatePrivateEndpointRule struct {
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	// Wire name: 'domain_names'
	DomainNames []string
}

func (st *UpdatePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := updatePrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateResponse struct {
}

func (st *UpdateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
	AllowMissing bool
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
	// Wire name: 'field_mask'
	FieldMask string

	// Wire name: 'setting'
	Setting RestrictWorkspaceAdminsSetting
}

func (st *UpdateRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRestrictWorkspaceAdminsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRestrictWorkspaceAdminsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateRestrictWorkspaceAdminsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Update workspace network configuration
type UpdateWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `tf:"-"`

	// Wire name: 'workspace_network_option'
	WorkspaceNetworkOption WorkspaceNetworkOption
}

func (st *UpdateWorkspaceNetworkOptionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWorkspaceNetworkOptionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWorkspaceNetworkOptionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWorkspaceNetworkOptionRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateWorkspaceNetworkOptionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type WorkspaceConf map[string]string

type WorkspaceNetworkOption struct {
	// The network policy ID to apply to the workspace. This controls the
	// network access rules for all serverless compute resources in the
	// workspace. Each workspace can only be linked to one policy at a time. If
	// no policy is explicitly assigned, the workspace will use
	// 'default-policy'.
	// Wire name: 'network_policy_id'
	NetworkPolicyId string
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64

	ForceSendFields []string `tf:"-"`
}

func (st *WorkspaceNetworkOption) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &workspaceNetworkOptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := workspaceNetworkOptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st WorkspaceNetworkOption) MarshalJSON() ([]byte, error) {
	pb, err := workspaceNetworkOptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
