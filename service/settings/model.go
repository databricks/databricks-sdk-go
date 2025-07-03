// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AccountIpAccessEnable struct {

	// Wire name: 'acct_ip_acl_enable'
	AcctIpAclEnable BooleanMessage `json:"acct_ip_acl_enable"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AccountIpAccessEnable) EncodeValues(key string, v *url.Values) error {
	pb, err := accountIpAccessEnableToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AccountId string `json:"account_id,omitempty"`
	// The network policies applying for egress traffic.
	// Wire name: 'egress'
	Egress *NetworkPolicyEgress `json:"egress,omitempty"`
	// The unique identifier for the network policy.
	// Wire name: 'network_policy_id'
	NetworkPolicyId string `json:"network_policy_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AccountNetworkPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := accountNetworkPolicyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType `json:"access_policy_type"`
}

func (st *AibiDashboardEmbeddingAccessPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := aibiDashboardEmbeddingAccessPolicyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type AibiDashboardEmbeddingAccessPolicySetting struct {

	// Wire name: 'aibi_dashboard_embedding_access_policy'
	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicy `json:"aibi_dashboard_embedding_access_policy"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AibiDashboardEmbeddingAccessPolicySetting) EncodeValues(key string, v *url.Values) error {
	pb, err := aibiDashboardEmbeddingAccessPolicySettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

func (st *AibiDashboardEmbeddingApprovedDomains) EncodeValues(key string, v *url.Values) error {
	pb, err := aibiDashboardEmbeddingApprovedDomainsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomains `json:"aibi_dashboard_embedding_approved_domains"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AibiDashboardEmbeddingApprovedDomainsSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := aibiDashboardEmbeddingApprovedDomainsSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessage `json:"automatic_cluster_update_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AutomaticClusterUpdateSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := automaticClusterUpdateSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *BooleanMessage) EncodeValues(key string, v *url.Values) error {
	pb, err := booleanMessageToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	CanToggle bool `json:"can_toggle,omitempty"`

	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`

	// Wire name: 'enablement_details'
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails `json:"enablement_details,omitempty"`

	// Wire name: 'maintenance_window'
	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow `json:"maintenance_window,omitempty"`

	// Wire name: 'restart_even_if_no_updates_available'
	RestartEvenIfNoUpdatesAvailable bool `json:"restart_even_if_no_updates_available,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterAutoRestartMessage) EncodeValues(key string, v *url.Values) error {
	pb, err := clusterAutoRestartMessageToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ForcedForComplianceMode bool `json:"forced_for_compliance_mode,omitempty"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	// Wire name: 'unavailable_for_disabled_entitlement'
	UnavailableForDisabledEntitlement bool `json:"unavailable_for_disabled_entitlement,omitempty"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	// Wire name: 'unavailable_for_non_enterprise_tier'
	UnavailableForNonEnterpriseTier bool `json:"unavailable_for_non_enterprise_tier,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterAutoRestartMessageEnablementDetails) EncodeValues(key string, v *url.Values) error {
	pb, err := clusterAutoRestartMessageEnablementDetailsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `json:"week_day_based_schedule,omitempty"`
}

func (st *ClusterAutoRestartMessageMaintenanceWindow) EncodeValues(key string, v *url.Values) error {
	pb, err := clusterAutoRestartMessageMaintenanceWindowToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'day_of_week'
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek `json:"day_of_week,omitempty"`

	// Wire name: 'frequency'
	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `json:"frequency,omitempty"`

	// Wire name: 'window_start_time'
	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `json:"window_start_time,omitempty"`
}

func (st *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) EncodeValues(key string, v *url.Values) error {
	pb, err := clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'hours'
	Hours int `json:"hours,omitempty"`

	// Wire name: 'minutes'
	Minutes int `json:"minutes,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) EncodeValues(key string, v *url.Values) error {
	pb, err := clusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`

	// Wire name: 'is_enabled'
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ComplianceSecurityProfile) EncodeValues(key string, v *url.Values) error {
	pb, err := complianceSecurityProfileToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'compliance_security_profile_workspace'
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfile `json:"compliance_security_profile_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ComplianceSecurityProfileSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := complianceSecurityProfileSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for ComplianceStandard.
//
// There is no guarantee on the order of the values in the slice.
func (f *ComplianceStandard) Values() []ComplianceStandard {
	return []ComplianceStandard{
		ComplianceStandardCanadaProtectedB,
		ComplianceStandardCyberEssentialPlus,
		ComplianceStandardFedrampHigh,
		ComplianceStandardFedrampIl5,
		ComplianceStandardFedrampModerate,
		ComplianceStandardHipaa,
		ComplianceStandardHitrust,
		ComplianceStandardIrapProtected,
		ComplianceStandardIsmap,
		ComplianceStandardItarEar,
		ComplianceStandardKFsi,
		ComplianceStandardNone,
		ComplianceStandardPciDss,
	}
}

// Type always returns ComplianceStandard to satisfy [pflag.Value] interface
func (f *ComplianceStandard) Type() string {
	return "ComplianceStandard"
}

type Config struct {

	// Wire name: 'email'
	Email *EmailConfig `json:"email,omitempty"`

	// Wire name: 'generic_webhook'
	GenericWebhook *GenericWebhookConfig `json:"generic_webhook,omitempty"`

	// Wire name: 'microsoft_teams'
	MicrosoftTeams *MicrosoftTeamsConfig `json:"microsoft_teams,omitempty"`

	// Wire name: 'pagerduty'
	Pagerduty *PagerdutyConfig `json:"pagerduty,omitempty"`

	// Wire name: 'slack'
	Slack *SlackConfig `json:"slack,omitempty"`
}

func (st *Config) EncodeValues(key string, v *url.Values) error {
	pb, err := configToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string `json:"label"`

	// Wire name: 'list_type'
	ListType ListType `json:"list_type"`
}

func (st *CreateIpAccessList) EncodeValues(key string, v *url.Values) error {
	pb, err := createIpAccessListToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (st *CreateIpAccessListResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createIpAccessListResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateNetworkConnectivityConfigRequest struct {

	// Wire name: 'network_connectivity_config'
	NetworkConnectivityConfig CreateNetworkConnectivityConfiguration `json:"network_connectivity_config"`
}

func (st *CreateNetworkConnectivityConfigRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createNetworkConnectivityConfigRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Name string `json:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	// Wire name: 'region'
	Region string `json:"region"`
}

func (st *CreateNetworkConnectivityConfiguration) EncodeValues(key string, v *url.Values) error {
	pb, err := createNetworkConnectivityConfigurationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateNetworkPolicyRequest struct {
	// Network policy configuration details.
	// Wire name: 'network_policy'
	NetworkPolicy AccountNetworkPolicy `json:"network_policy"`
}

func (st *CreateNetworkPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createNetworkPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Config *Config `json:"config,omitempty"`
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateNotificationDestinationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createNotificationDestinationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ApplicationId string `json:"application_id"`
	// Comment that describes the purpose of the token.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The number of seconds before the token expires.
	// Wire name: 'lifetime_seconds'
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateOboTokenRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createOboTokenRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	// Wire name: 'token_value'
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateOboTokenResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createOboTokenResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	// Wire name: 'domain_names'
	DomainNames []string `json:"domain_names,omitempty"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	// Wire name: 'endpoint_service'
	EndpointService string `json:"endpoint_service,omitempty"`
	// Not used by customer-managed private endpoint services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	// Wire name: 'group_id'
	GroupId string `json:"group_id,omitempty"`
	// The Azure resource ID of the target resource.
	// Wire name: 'resource_id'
	ResourceId string `json:"resource_id,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	// Wire name: 'resource_names'
	ResourceNames []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreatePrivateEndpointRule) EncodeValues(key string, v *url.Values) error {
	pb, err := createPrivateEndpointRuleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreatePrivateEndpointRuleRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`

	// Wire name: 'private_endpoint_rule'
	PrivateEndpointRule CreatePrivateEndpointRule `json:"private_endpoint_rule"`
}

func (st *CreatePrivateEndpointRuleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createPrivateEndpointRuleRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	// Wire name: 'lifetime_seconds'
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateTokenRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createTokenRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	// Wire name: 'token_value'
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateTokenResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createTokenResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	// Enforced = it cannot be overriden at workspace level.
	// Wire name: 'is_enforced'
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CspEnablementAccount) EncodeValues(key string, v *url.Values) error {
	pb, err := cspEnablementAccountToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'csp_enablement_account'
	CspEnablementAccount CspEnablementAccount `json:"csp_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CspEnablementAccountSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := cspEnablementAccountSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Properties of the new private endpoint rule. Note that for private endpoints
// towards a VPC endpoint service behind a customer-managed NLB, you must
// approve the endpoint in AWS console after initialization.
type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule struct {
	// Databricks account ID. You can find your account ID from the Accounts
	// Console.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the AWS console
	// before they take effect. The possible values are: - PENDING: The endpoint
	// has been created and pending approval. - ESTABLISHED: The endpoint has
	// been approved and is ready to use in your serverless compute resources. -
	// REJECTED: Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up. - EXPIRED: If the endpoint is created but not approved in 14
	// days, it is EXPIRED.
	// Wire name: 'connection_state'
	ConnectionState CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	// Wire name: 'deactivated'
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	// Wire name: 'deactivated_at'
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// Only used by private endpoints towards a VPC endpoint service for
	// customer-managed VPC endpoint service.
	//
	// The target AWS resource FQDNs accessible via the VPC endpoint service.
	// When updating this field, we perform full update on this field. Please
	// ensure a full list of desired domain_names is provided.
	// Wire name: 'domain_names'
	DomainNames []string `json:"domain_names,omitempty"`
	// Only used by private endpoints towards an AWS S3 service.
	//
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	// Wire name: 'endpoint_service'
	EndpointService string `json:"endpoint_service,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	// Wire name: 'resource_names'
	ResourceNames []string `json:"resource_names,omitempty"`
	// The ID of a private endpoint rule.
	// Wire name: 'rule_id'
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	// Wire name: 'updated_time'
	UpdatedTime int64 `json:"updated_time,omitempty"`
	// The AWS VPC endpoint ID. You can use this ID to identify VPC endpoint
	// created by Databricks.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) EncodeValues(key string, v *url.Values) error {
	pb, err := customerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := customerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState string

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateDisconnected CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `DISCONNECTED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateEstablished CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `ESTABLISHED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateExpired CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `EXPIRED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePending CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `PENDING`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateRejected CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState = `REJECTED`

// String representation for [fmt.Print]
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `ESTABLISHED`, `EXPIRED`, `PENDING`, `REJECTED`:
		*f = CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "ESTABLISHED", "EXPIRED", "PENDING", "REJECTED"`, v)
	}
}

// Values returns all possible values for CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState.
//
// There is no guarantee on the order of the values in the slice.
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) Values() []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState {
	return []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState{
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateDisconnected,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateEstablished,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateExpired,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePending,
		CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateRejected,
	}
}

// Type always returns CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState to satisfy [pflag.Value] interface
func (f *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) Type() string {
	return "CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState"
}

type DashboardEmailSubscriptions struct {

	// Wire name: 'boolean_val'
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DashboardEmailSubscriptions) EncodeValues(key string, v *url.Values) error {
	pb, err := dashboardEmailSubscriptionsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DashboardEmailSubscriptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dashboardEmailSubscriptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dashboardEmailSubscriptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DashboardEmailSubscriptions) MarshalJSON() ([]byte, error) {
	pb, err := dashboardEmailSubscriptionsToPb(&st)
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
	Etag string `json:"etag,omitempty"`

	// Wire name: 'namespace'
	Namespace StringMessage `json:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DefaultNamespaceSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := defaultNamespaceSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteAccountIpAccessEnableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAccountIpAccessEnableRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteAccountIpAccessEnableResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAccountIpAccessEnableResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st *DeleteAccountIpAccessListRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAccountIpAccessListRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAibiDashboardEmbeddingAccessPolicySettingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteDashboardEmailSubscriptionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDashboardEmailSubscriptionsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDashboardEmailSubscriptionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDashboardEmailSubscriptionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDashboardEmailSubscriptionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The etag is returned.
type DeleteDashboardEmailSubscriptionsResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string `json:"etag"`
}

func (st *DeleteDashboardEmailSubscriptionsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDashboardEmailSubscriptionsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteDashboardEmailSubscriptionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDashboardEmailSubscriptionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDashboardEmailSubscriptionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDashboardEmailSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDashboardEmailSubscriptionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteDefaultNamespaceSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDefaultNamespaceSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteDefaultNamespaceSettingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDefaultNamespaceSettingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteDisableLegacyAccessRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDisableLegacyAccessRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteDisableLegacyAccessResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDisableLegacyAccessResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteDisableLegacyDbfsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDisableLegacyDbfsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteDisableLegacyDbfsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDisableLegacyDbfsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteDisableLegacyFeaturesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDisableLegacyFeaturesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteDisableLegacyFeaturesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDisableLegacyFeaturesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st *DeleteIpAccessListRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteIpAccessListRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteLlmProxyPartnerPoweredWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteLlmProxyPartnerPoweredWorkspaceRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteLlmProxyPartnerPoweredWorkspaceResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteLlmProxyPartnerPoweredWorkspaceResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
}

func (st *DeleteNetworkConnectivityConfigurationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteNetworkConnectivityConfigurationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

func (st *DeleteNetworkConnectivityConfigurationResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteNetworkConnectivityConfigurationResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteNetworkPolicyRequest struct {
	// The unique identifier of the network policy to delete.
	NetworkPolicyId string `json:"-" tf:"-"`
}

func (st *DeleteNetworkPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteNetworkPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

func (st *DeleteNetworkPolicyRpcResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteNetworkPolicyRpcResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteNotificationDestinationRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *DeleteNotificationDestinationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteNotificationDestinationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeletePersonalComputeSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deletePersonalComputeSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeletePersonalComputeSettingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deletePersonalComputeSettingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" tf:"-"`
}

func (st *DeletePrivateEndpointRuleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deletePrivateEndpointRuleRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteRestrictWorkspaceAdminsSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRestrictWorkspaceAdminsSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Etag string `json:"etag"`
}

func (st *DeleteRestrictWorkspaceAdminsSettingResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteRestrictWorkspaceAdminsSettingResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DeleteSqlResultsDownloadRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteSqlResultsDownloadRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSqlResultsDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSqlResultsDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteSqlResultsDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// The etag is returned.
type DeleteSqlResultsDownloadResponse struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	// Wire name: 'etag'
	Etag string `json:"etag"`
}

func (st *DeleteSqlResultsDownloadResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteSqlResultsDownloadResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteSqlResultsDownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteSqlResultsDownloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteSqlResultsDownloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteSqlResultsDownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteSqlResultsDownloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId string `json:"-" tf:"-"`
}

func (st *DeleteTokenManagementRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteTokenManagementRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for DestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *DestinationType) Values() []DestinationType {
	return []DestinationType{
		DestinationTypeEmail,
		DestinationTypeMicrosoftTeams,
		DestinationTypePagerduty,
		DestinationTypeSlack,
		DestinationTypeWebhook,
	}
}

// Type always returns DestinationType to satisfy [pflag.Value] interface
func (f *DestinationType) Type() string {
	return "DestinationType"
}

type DisableLegacyAccess struct {

	// Wire name: 'disable_legacy_access'
	DisableLegacyAccess BooleanMessage `json:"disable_legacy_access"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DisableLegacyAccess) EncodeValues(key string, v *url.Values) error {
	pb, err := disableLegacyAccessToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	DisableLegacyDbfs BooleanMessage `json:"disable_legacy_dbfs"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DisableLegacyDbfs) EncodeValues(key string, v *url.Values) error {
	pb, err := disableLegacyDbfsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	DisableLegacyFeatures BooleanMessage `json:"disable_legacy_features"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DisableLegacyFeatures) EncodeValues(key string, v *url.Values) error {
	pb, err := disableLegacyFeaturesToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	InternetAccess *EgressNetworkPolicyInternetAccessPolicy `json:"internet_access,omitempty"`
}

func (st *EgressNetworkPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`

	// Wire name: 'allowed_storage_destinations'
	AllowedStorageDestinations []EgressNetworkPolicyInternetAccessPolicyStorageDestination `json:"allowed_storage_destinations,omitempty"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	// Wire name: 'log_only_mode'
	LogOnlyMode *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode `json:"log_only_mode,omitempty"`

	// Wire name: 'restriction_mode'
	RestrictionMode EgressNetworkPolicyInternetAccessPolicyRestrictionMode `json:"restriction_mode,omitempty"`
}

func (st *EgressNetworkPolicyInternetAccessPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyInternetAccessPolicyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Destination string `json:"destination,omitempty"`

	// Wire name: 'protocol'
	Protocol EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol `json:"protocol,omitempty"`

	// Wire name: 'type'
	Type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EgressNetworkPolicyInternetAccessPolicyInternetDestination) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) Values() []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol {
	return []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol{
		EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolTcp,
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

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Values() []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType {
	return []EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType{
		EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFqdn,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType"
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode struct {

	// Wire name: 'log_only_mode_type'
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType `json:"log_only_mode_type,omitempty"`

	// Wire name: 'workloads'
	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType `json:"workloads,omitempty"`
}

func (st *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) Values() []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType {
	return []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType{
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeAllServices,
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeSelectedServices,
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

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) Values() []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType {
	return []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType{
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeDbsql,
		EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeMlServing,
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

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyRestrictionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Values() []EgressNetworkPolicyInternetAccessPolicyRestrictionMode {
	return []EgressNetworkPolicyInternetAccessPolicyRestrictionMode{
		EgressNetworkPolicyInternetAccessPolicyRestrictionModeFullAccess,
		EgressNetworkPolicyInternetAccessPolicyRestrictionModePrivateAccessOnly,
		EgressNetworkPolicyInternetAccessPolicyRestrictionModeRestrictedAccess,
	}
}

// Type always returns EgressNetworkPolicyInternetAccessPolicyRestrictionMode to satisfy [pflag.Value] interface
func (f *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) Type() string {
	return "EgressNetworkPolicyInternetAccessPolicyRestrictionMode"
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyInternetAccessPolicyStorageDestination struct {

	// Wire name: 'allowed_paths'
	AllowedPaths []string `json:"allowed_paths,omitempty"`

	// Wire name: 'azure_container'
	AzureContainer string `json:"azure_container,omitempty"`

	// Wire name: 'azure_dns_zone'
	AzureDnsZone string `json:"azure_dns_zone,omitempty"`

	// Wire name: 'azure_storage_account'
	AzureStorageAccount string `json:"azure_storage_account,omitempty"`

	// Wire name: 'azure_storage_service'
	AzureStorageService string `json:"azure_storage_service,omitempty"`

	// Wire name: 'bucket_name'
	BucketName string `json:"bucket_name,omitempty"`

	// Wire name: 'region'
	Region string `json:"region,omitempty"`

	// Wire name: 'type'
	Type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EgressNetworkPolicyInternetAccessPolicyStorageDestination) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) Values() []EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType {
	return []EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType{
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAwsS3,
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage,
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeCloudflareR2,
		EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage,
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
	AllowedInternetDestinations []EgressNetworkPolicyNetworkAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`
	// List of storage destinations that serverless workloads are allowed to
	// access when in RESTRICTED_ACCESS mode.
	// Wire name: 'allowed_storage_destinations'
	AllowedStorageDestinations []EgressNetworkPolicyNetworkAccessPolicyStorageDestination `json:"allowed_storage_destinations,omitempty"`
	// Optional. When policy_enforcement is not provided, we default to
	// ENFORCE_MODE_ALL_SERVICES
	// Wire name: 'policy_enforcement'
	PolicyEnforcement *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement `json:"policy_enforcement,omitempty"`
	// The restriction mode that controls how serverless workloads can access
	// the internet.
	// Wire name: 'restriction_mode'
	RestrictionMode EgressNetworkPolicyNetworkAccessPolicyRestrictionMode `json:"restriction_mode"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicy) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyNetworkAccessPolicyToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Destination string `json:"destination,omitempty"`
	// The type of internet destination. Currently only DNS_NAME is supported.
	// Wire name: 'internet_destination_type'
	InternetDestinationType EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType `json:"internet_destination_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyNetworkAccessPolicyInternetDestinationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) Values() []EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType {
	return []EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType{
		EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeDnsName,
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
	DryRunModeProductFilter []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter `json:"dry_run_mode_product_filter,omitempty"`
	// The mode of policy enforcement. ENFORCED blocks traffic that violates
	// policy, while DRY_RUN only logs violations without blocking. When not
	// specified, defaults to ENFORCED.
	// Wire name: 'enforcement_mode'
	EnforcementMode EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode `json:"enforcement_mode,omitempty"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) Values() []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter {
	return []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter{
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterDbsql,
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterMlServing,
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

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) Values() []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode {
	return []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode{
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeDryRun,
		EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeEnforced,
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

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyRestrictionMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) Values() []EgressNetworkPolicyNetworkAccessPolicyRestrictionMode {
	return []EgressNetworkPolicyNetworkAccessPolicyRestrictionMode{
		EgressNetworkPolicyNetworkAccessPolicyRestrictionModeFullAccess,
		EgressNetworkPolicyNetworkAccessPolicyRestrictionModeRestrictedAccess,
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
	AzureStorageAccount string `json:"azure_storage_account,omitempty"`
	// The Azure storage service type (blob, dfs, etc.).
	// Wire name: 'azure_storage_service'
	AzureStorageService string `json:"azure_storage_service,omitempty"`

	// Wire name: 'bucket_name'
	BucketName string `json:"bucket_name,omitempty"`

	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// The type of storage destination.
	// Wire name: 'storage_destination_type'
	StorageDestinationType EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType `json:"storage_destination_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) EncodeValues(key string, v *url.Values) error {
	pb, err := egressNetworkPolicyNetworkAccessPolicyStorageDestinationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) Values() []EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType {
	return []EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType{
		EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAwsS3,
		EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage,
		EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage,
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

// Values returns all possible values for EgressResourceType.
//
// There is no guarantee on the order of the values in the slice.
func (f *EgressResourceType) Values() []EgressResourceType {
	return []EgressResourceType{
		EgressResourceTypeAzureBlobStorage,
	}
}

// Type always returns EgressResourceType to satisfy [pflag.Value] interface
func (f *EgressResourceType) Type() string {
	return "EgressResourceType"
}

type EmailConfig struct {
	// Email addresses to notify.
	// Wire name: 'addresses'
	Addresses []string `json:"addresses,omitempty"`
}

func (st *EmailConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := emailConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type EnableExportNotebook struct {

	// Wire name: 'boolean_val'
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EnableExportNotebook) EncodeValues(key string, v *url.Values) error {
	pb, err := enableExportNotebookToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EnableNotebookTableClipboard) EncodeValues(key string, v *url.Values) error {
	pb, err := enableNotebookTableClipboardToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	BooleanVal *BooleanMessage `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EnableResultsDownloading) EncodeValues(key string, v *url.Values) error {
	pb, err := enableResultsDownloadingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EnhancedSecurityMonitoring) EncodeValues(key string, v *url.Values) error {
	pb, err := enhancedSecurityMonitoringToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'enhanced_security_monitoring_workspace'
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoring `json:"enhanced_security_monitoring_workspace"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EnhancedSecurityMonitoringSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := enhancedSecurityMonitoringSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EsmEnablementAccount) EncodeValues(key string, v *url.Values) error {
	pb, err := esmEnablementAccountToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'esm_enablement_account'
	EsmEnablementAccount EsmEnablementAccount `json:"esm_enablement_account"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *EsmEnablementAccountSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := esmEnablementAccountSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Credential string `json:"credential,omitempty"`
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	// Wire name: 'credentialEolTime'
	CredentialEolTime int64 `json:"credentialEolTime,omitempty"`
	// User ID of the user that owns this token.
	// Wire name: 'ownerId'
	OwnerId int64 `json:"ownerId,omitempty"`
	// The scopes of access granted in the token.
	// Wire name: 'scopes'
	Scopes []string `json:"scopes,omitempty"`
	// The type of this exchange token
	// Wire name: 'tokenType'
	TokenType TokenType `json:"tokenType,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ExchangeToken) EncodeValues(key string, v *url.Values) error {
	pb, err := exchangeTokenToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	PartitionId PartitionId `json:"partitionId"`
	// Array of scopes for the token request.
	// Wire name: 'scopes'
	Scopes []string `json:"scopes"`
	// A list of token types being requested
	// Wire name: 'tokenType'
	TokenType []TokenType `json:"tokenType"`
}

func (st *ExchangeTokenRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := exchangeTokenRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Values []ExchangeToken `json:"values,omitempty"`
}

func (st *ExchangeTokenResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := exchangeTokenResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (st *FetchIpAccessListResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := fetchIpAccessListResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Password string `json:"password,omitempty"`
	// [Output-Only] Whether password is set.
	// Wire name: 'password_set'
	PasswordSet bool `json:"password_set,omitempty"`
	// [Input-Only] URL for webhook.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet bool `json:"url_set,omitempty"`
	// [Input-Only][Optional] Username for webhook.
	// Wire name: 'username'
	Username string `json:"username,omitempty"`
	// [Output-Only] Whether username is set.
	// Wire name: 'username_set'
	UsernameSet bool `json:"username_set,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GenericWebhookConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := genericWebhookConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetAccountIpAccessEnableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAccountIpAccessEnableRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st *GetAccountIpAccessListRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAccountIpAccessListRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetAibiDashboardEmbeddingAccessPolicySettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetAutomaticClusterUpdateSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetAutomaticClusterUpdateSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getAutomaticClusterUpdateSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetComplianceSecurityProfileSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetComplianceSecurityProfileSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getComplianceSecurityProfileSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetCspEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetCspEnablementAccountSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getCspEnablementAccountSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetDashboardEmailSubscriptionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDashboardEmailSubscriptionsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDashboardEmailSubscriptionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDashboardEmailSubscriptionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDashboardEmailSubscriptionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetDefaultNamespaceSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDefaultNamespaceSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetDisableLegacyAccessRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDisableLegacyAccessRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetDisableLegacyDbfsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDisableLegacyDbfsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetDisableLegacyFeaturesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDisableLegacyFeaturesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetEnhancedSecurityMonitoringSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetEnhancedSecurityMonitoringSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getEnhancedSecurityMonitoringSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetEsmEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetEsmEnablementAccountSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getEsmEnablementAccountSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st *GetIpAccessListRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getIpAccessListRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (st *GetIpAccessListResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getIpAccessListResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func (st *GetIpAccessListsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getIpAccessListsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetLlmProxyPartnerPoweredAccountRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetLlmProxyPartnerPoweredAccountRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getLlmProxyPartnerPoweredAccountRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetLlmProxyPartnerPoweredEnforceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetLlmProxyPartnerPoweredEnforceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getLlmProxyPartnerPoweredEnforceRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetLlmProxyPartnerPoweredWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getLlmProxyPartnerPoweredWorkspaceRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
}

func (st *GetNetworkConnectivityConfigurationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getNetworkConnectivityConfigurationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetNetworkPolicyRequest struct {
	// The unique identifier of the network policy to retrieve.
	NetworkPolicyId string `json:"-" tf:"-"`
}

func (st *GetNetworkPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getNetworkPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetNotificationDestinationRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st *GetNotificationDestinationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getNotificationDestinationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetPersonalComputeSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getPersonalComputeSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" tf:"-"`
}

func (st *GetPrivateEndpointRuleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getPrivateEndpointRuleRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetRestrictWorkspaceAdminsSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getRestrictWorkspaceAdminsSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetSqlResultsDownloadRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getSqlResultsDownloadRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getSqlResultsDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getSqlResultsDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := getSqlResultsDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetStatusRequest struct {
	Keys string `json:"-" tf:"-"`
}

func (st *GetStatusRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getStatusRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" tf:"-"`
}

func (st *GetTokenManagementRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getTokenManagementRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st *GetTokenPermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getTokenPermissionLevelsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
}

func (st *GetTokenResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getTokenResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st *GetWorkspaceNetworkOptionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getWorkspaceNetworkOptionRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AddressCount int `json:"address_count,omitempty"`
	// Creation timestamp in milliseconds.
	// Wire name: 'created_at'
	CreatedAt int64 `json:"created_at,omitempty"`
	// User ID of the user who created this list.
	// Wire name: 'created_by'
	CreatedBy int64 `json:"created_by,omitempty"`
	// Specifies whether this IP access list is enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`

	// Wire name: 'ip_addresses'
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string `json:"label,omitempty"`
	// Universally unique identifier (UUID) of the IP access list.
	// Wire name: 'list_id'
	ListId string `json:"list_id,omitempty"`

	// Wire name: 'list_type'
	ListType ListType `json:"list_type,omitempty"`
	// Update timestamp in milliseconds.
	// Wire name: 'updated_at'
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// User ID of the user who updated this list.
	// Wire name: 'updated_by'
	UpdatedBy int64 `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *IpAccessListInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := ipAccessListInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func (st *ListIpAccessListResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listIpAccessListResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListNetworkConnectivityConfigurationsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listNetworkConnectivityConfigurationsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Items []NetworkConnectivityConfiguration `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListNetworkConnectivityConfigurationsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listNetworkConnectivityConfigurationsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListNetworkPoliciesRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListNetworkPoliciesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listNetworkPoliciesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Items []AccountNetworkPolicy `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListNetworkPoliciesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listNetworkPoliciesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListNotificationDestinationsRequest struct {
	PageSize int64 `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListNotificationDestinationsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listNotificationDestinationsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'results'
	Results []ListNotificationDestinationsResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListNotificationDestinationsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listNotificationDestinationsResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListNotificationDestinationsResult) EncodeValues(key string, v *url.Values) error {
	pb, err := listNotificationDestinationsResultToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPrivateEndpointRulesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listPrivateEndpointRulesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// The private endpoint rule list was successfully retrieved.
type ListPrivateEndpointRulesResponse struct {

	// Wire name: 'items'
	Items []NccPrivateEndpointRule `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListPrivateEndpointRulesResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listPrivateEndpointRulesResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListPrivateEndpointRulesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPrivateEndpointRulesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPrivateEndpointRulesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPrivateEndpointRulesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listPrivateEndpointRulesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	// Wire name: 'token_infos'
	TokenInfos []PublicTokenInfo `json:"token_infos,omitempty"`
}

func (st *ListPublicTokensResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listPublicTokensResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById int64 `json:"-" tf:"-"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListTokenManagementRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listTokenManagementRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

func (st *ListTokensResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listTokensResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for ListType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ListType) Values() []ListType {
	return []ListType{
		ListTypeAllow,
		ListTypeBlock,
	}
}

// Type always returns ListType to satisfy [pflag.Value] interface
func (f *ListType) Type() string {
	return "ListType"
}

type LlmProxyPartnerPoweredAccount struct {

	// Wire name: 'boolean_val'
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *LlmProxyPartnerPoweredAccount) EncodeValues(key string, v *url.Values) error {
	pb, err := llmProxyPartnerPoweredAccountToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *LlmProxyPartnerPoweredEnforce) EncodeValues(key string, v *url.Values) error {
	pb, err := llmProxyPartnerPoweredEnforceToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *LlmProxyPartnerPoweredWorkspace) EncodeValues(key string, v *url.Values) error {
	pb, err := llmProxyPartnerPoweredWorkspaceToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *MicrosoftTeamsConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := microsoftTeamsConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

func (st *NccAwsStableIpRule) EncodeValues(key string, v *url.Values) error {
	pb, err := nccAwsStableIpRuleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	ConnectionState NccAzurePrivateEndpointRuleConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	// Wire name: 'deactivated'
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	// Wire name: 'deactivated_at'
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// Not used by customer-managed private endpoint services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	// Wire name: 'domain_names'
	DomainNames []string `json:"domain_names,omitempty"`
	// The name of the Azure private endpoint resource.
	// Wire name: 'endpoint_name'
	EndpointName string `json:"endpoint_name,omitempty"`
	// Only used by private endpoints to Azure first-party services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	// Wire name: 'group_id'
	GroupId string `json:"group_id,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The Azure resource ID of the target resource.
	// Wire name: 'resource_id'
	ResourceId string `json:"resource_id,omitempty"`
	// The ID of a private endpoint rule.
	// Wire name: 'rule_id'
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	// Wire name: 'updated_time'
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NccAzurePrivateEndpointRule) EncodeValues(key string, v *url.Values) error {
	pb, err := nccAzurePrivateEndpointRuleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for NccAzurePrivateEndpointRuleConnectionState.
//
// There is no guarantee on the order of the values in the slice.
func (f *NccAzurePrivateEndpointRuleConnectionState) Values() []NccAzurePrivateEndpointRuleConnectionState {
	return []NccAzurePrivateEndpointRuleConnectionState{
		NccAzurePrivateEndpointRuleConnectionStateDisconnected,
		NccAzurePrivateEndpointRuleConnectionStateEstablished,
		NccAzurePrivateEndpointRuleConnectionStateExpired,
		NccAzurePrivateEndpointRuleConnectionStateInit,
		NccAzurePrivateEndpointRuleConnectionStatePending,
		NccAzurePrivateEndpointRuleConnectionStateRejected,
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
	Subnets []string `json:"subnets,omitempty"`
	// The Azure region in which this service endpoint rule applies..
	// Wire name: 'target_region'
	TargetRegion string `json:"target_region,omitempty"`
	// The Azure services to which this service endpoint rule applies to.
	// Wire name: 'target_services'
	TargetServices []EgressResourceType `json:"target_services,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NccAzureServiceEndpointRule) EncodeValues(key string, v *url.Values) error {
	pb, err := nccAzureServiceEndpointRuleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	DefaultRules *NccEgressDefaultRules `json:"default_rules,omitempty"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	// Wire name: 'target_rules'
	TargetRules *NccEgressTargetRules `json:"target_rules,omitempty"`
}

func (st *NccEgressConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := nccEgressConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'aws_stable_ip_rule'
	AwsStableIpRule *NccAwsStableIpRule `json:"aws_stable_ip_rule,omitempty"`

	// Wire name: 'azure_service_endpoint_rule'
	AzureServiceEndpointRule *NccAzureServiceEndpointRule `json:"azure_service_endpoint_rule,omitempty"`
}

func (st *NccEgressDefaultRules) EncodeValues(key string, v *url.Values) error {
	pb, err := nccEgressDefaultRulesToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// AWS private endpoint rule controls the AWS private endpoint based egress
	// rules.
	// Wire name: 'aws_private_endpoint_rules'
	AwsPrivateEndpointRules []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule `json:"aws_private_endpoint_rules,omitempty"`

	// Wire name: 'azure_private_endpoint_rules'
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule `json:"azure_private_endpoint_rules,omitempty"`
}

func (st *NccEgressTargetRules) EncodeValues(key string, v *url.Values) error {
	pb, err := nccEgressTargetRulesToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type NccPrivateEndpointRule struct {
	// Databricks account ID. You can find your account ID from the Accounts
	// Console.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// The current status of this private endpoint. The private endpoint rules
	// are effective only if the connection state is ESTABLISHED. Remember that
	// you must approve new endpoints on your resources in the Cloud console
	// before they take effect. The possible values are: - PENDING: The endpoint
	// has been created and pending approval. - ESTABLISHED: The endpoint has
	// been approved and is ready to use in your serverless compute resources. -
	// REJECTED: Connection was rejected by the private link resource owner. -
	// DISCONNECTED: Connection was removed by the private link resource owner,
	// the private endpoint becomes informative and should be deleted for
	// clean-up. - EXPIRED: If the endpoint was created but not approved in 14
	// days, it will be EXPIRED.
	// Wire name: 'connection_state'
	ConnectionState NccPrivateEndpointRulePrivateLinkConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	// Wire name: 'deactivated'
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	// Wire name: 'deactivated_at'
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	// Wire name: 'domain_names'
	DomainNames []string `json:"domain_names,omitempty"`
	// Only used by private endpoints towards an AWS S3 service.
	//
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the Azure private endpoint resource.
	// Wire name: 'endpoint_name'
	EndpointName string `json:"endpoint_name,omitempty"`
	// The full target AWS endpoint service name that connects to the
	// destination resources of the private endpoint.
	// Wire name: 'endpoint_service'
	EndpointService string `json:"endpoint_service,omitempty"`
	// Not used by customer-managed private endpoint services.
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	// Wire name: 'group_id'
	GroupId string `json:"group_id,omitempty"`
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The Azure resource ID of the target resource.
	// Wire name: 'resource_id'
	ResourceId string `json:"resource_id,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	// Wire name: 'resource_names'
	ResourceNames []string `json:"resource_names,omitempty"`
	// The ID of a private endpoint rule.
	// Wire name: 'rule_id'
	RuleId string `json:"rule_id,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	// Wire name: 'updated_time'
	UpdatedTime int64 `json:"updated_time,omitempty"`
	// The AWS VPC endpoint ID. You can use this ID to identify the VPC endpoint
	// created by Databricks.
	// Wire name: 'vpc_endpoint_id'
	VpcEndpointId string `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NccPrivateEndpointRule) EncodeValues(key string, v *url.Values) error {
	pb, err := nccPrivateEndpointRuleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *NccPrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &nccPrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := nccPrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st NccPrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := nccPrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type NccPrivateEndpointRulePrivateLinkConnectionState string

const NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected NccPrivateEndpointRulePrivateLinkConnectionState = `DISCONNECTED`

const NccPrivateEndpointRulePrivateLinkConnectionStateEstablished NccPrivateEndpointRulePrivateLinkConnectionState = `ESTABLISHED`

const NccPrivateEndpointRulePrivateLinkConnectionStateExpired NccPrivateEndpointRulePrivateLinkConnectionState = `EXPIRED`

const NccPrivateEndpointRulePrivateLinkConnectionStatePending NccPrivateEndpointRulePrivateLinkConnectionState = `PENDING`

const NccPrivateEndpointRulePrivateLinkConnectionStateRejected NccPrivateEndpointRulePrivateLinkConnectionState = `REJECTED`

// String representation for [fmt.Print]
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) Set(v string) error {
	switch v {
	case `DISCONNECTED`, `ESTABLISHED`, `EXPIRED`, `PENDING`, `REJECTED`:
		*f = NccPrivateEndpointRulePrivateLinkConnectionState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DISCONNECTED", "ESTABLISHED", "EXPIRED", "PENDING", "REJECTED"`, v)
	}
}

// Values returns all possible values for NccPrivateEndpointRulePrivateLinkConnectionState.
//
// There is no guarantee on the order of the values in the slice.
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) Values() []NccPrivateEndpointRulePrivateLinkConnectionState {
	return []NccPrivateEndpointRulePrivateLinkConnectionState{
		NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected,
		NccPrivateEndpointRulePrivateLinkConnectionStateEstablished,
		NccPrivateEndpointRulePrivateLinkConnectionStateExpired,
		NccPrivateEndpointRulePrivateLinkConnectionStatePending,
		NccPrivateEndpointRulePrivateLinkConnectionStateRejected,
	}
}

// Type always returns NccPrivateEndpointRulePrivateLinkConnectionState to satisfy [pflag.Value] interface
func (f *NccPrivateEndpointRulePrivateLinkConnectionState) Type() string {
	return "NccPrivateEndpointRulePrivateLinkConnectionState"
}

// Properties of the new network connectivity configuration.
type NetworkConnectivityConfiguration struct {
	// Your Databricks account ID. You can find your account ID in your
	// Databricks accounts console.
	// Wire name: 'account_id'
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when this object was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	// Wire name: 'egress_config'
	EgressConfig *NccEgressConfig `json:"egress_config,omitempty"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Databricks network connectivity configuration ID.
	// Wire name: 'network_connectivity_config_id'
	NetworkConnectivityConfigId string `json:"network_connectivity_config_id,omitempty"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	// Wire name: 'region'
	Region string `json:"region,omitempty"`
	// Time in epoch milliseconds when this object was updated.
	// Wire name: 'updated_time'
	UpdatedTime int64 `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NetworkConnectivityConfiguration) EncodeValues(key string, v *url.Values) error {
	pb, err := networkConnectivityConfigurationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	NetworkAccess *EgressNetworkPolicyNetworkAccessPolicy `json:"network_access,omitempty"`
}

func (st *NetworkPolicyEgress) EncodeValues(key string, v *url.Values) error {
	pb, err := networkPolicyEgressToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Config *Config `json:"config,omitempty"`
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	// Wire name: 'destination_type'
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *NotificationDestination) EncodeValues(key string, v *url.Values) error {
	pb, err := notificationDestinationToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	IntegrationKey string `json:"integration_key,omitempty"`
	// [Output-Only] Whether integration key is set.
	// Wire name: 'integration_key_set'
	IntegrationKeySet bool `json:"integration_key_set,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PagerdutyConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := pagerdutyConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PartitionId) EncodeValues(key string, v *url.Values) error {
	pb, err := partitionIdToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

	// Wire name: 'value'
	Value PersonalComputeMessageEnum `json:"value"`
}

func (st *PersonalComputeMessage) EncodeValues(key string, v *url.Values) error {
	pb, err := personalComputeMessageToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for PersonalComputeMessageEnum.
//
// There is no guarantee on the order of the values in the slice.
func (f *PersonalComputeMessageEnum) Values() []PersonalComputeMessageEnum {
	return []PersonalComputeMessageEnum{
		PersonalComputeMessageEnumDelegate,
		PersonalComputeMessageEnumOn,
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
	Etag string `json:"etag,omitempty"`

	// Wire name: 'personal_compute'
	PersonalCompute PersonalComputeMessage `json:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PersonalComputeSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := personalComputeSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Comment string `json:"comment,omitempty"`
	// Server time (in epoch milliseconds) when the token was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	// Wire name: 'expiry_time'
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// The ID of this token.
	// Wire name: 'token_id'
	TokenId string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *PublicTokenInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := publicTokenInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Enabled bool `json:"enabled"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`

	// Wire name: 'ip_addresses'
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string `json:"label"`

	// Wire name: 'list_type'
	ListType ListType `json:"list_type"`
}

func (st *ReplaceIpAccessList) EncodeValues(key string, v *url.Values) error {
	pb, err := replaceIpAccessListToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RestrictWorkspaceAdminsMessage struct {

	// Wire name: 'status'
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`
}

func (st *RestrictWorkspaceAdminsMessage) EncodeValues(key string, v *url.Values) error {
	pb, err := restrictWorkspaceAdminsMessageToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type RestrictWorkspaceAdminsSetting struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`

	// Wire name: 'restrict_workspace_admins'
	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage `json:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *RestrictWorkspaceAdminsSetting) EncodeValues(key string, v *url.Values) error {
	pb, err := restrictWorkspaceAdminsSettingToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	TokenId string `json:"token_id"`
}

func (st *RevokeTokenRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := revokeTokenRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SlackConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := slackConfigToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type SqlResultsDownload struct {

	// Wire name: 'boolean_val'
	BooleanVal BooleanMessage `json:"boolean_val"`
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	// Wire name: 'etag'
	Etag string `json:"etag,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	// Wire name: 'setting_name'
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *SqlResultsDownload) EncodeValues(key string, v *url.Values) error {
	pb, err := sqlResultsDownloadToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *SqlResultsDownload) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &sqlResultsDownloadPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := sqlResultsDownloadFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SqlResultsDownload) MarshalJSON() ([]byte, error) {
	pb, err := sqlResultsDownloadToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StringMessage struct {
	// Represents a generic string value.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *StringMessage) EncodeValues(key string, v *url.Values) error {
	pb, err := stringMessageToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenAccessControlRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AllPermissions []TokenPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenAccessControlResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Comment string `json:"comment,omitempty"`
	// User ID of the user that created the token.
	// Wire name: 'created_by_id'
	CreatedById int64 `json:"created_by_id,omitempty"`
	// Username of the user that created the token.
	// Wire name: 'created_by_username'
	CreatedByUsername string `json:"created_by_username,omitempty"`
	// Timestamp when the token was created.
	// Wire name: 'creation_time'
	CreationTime int64 `json:"creation_time,omitempty"`
	// Timestamp when the token expires.
	// Wire name: 'expiry_time'
	ExpiryTime int64 `json:"expiry_time,omitempty"`
	// Approximate timestamp for the day the token was last used. Accurate up to
	// 1 day.
	// Wire name: 'last_used_day'
	LastUsedDay int64 `json:"last_used_day,omitempty"`
	// User ID of the user that owns the token.
	// Wire name: 'owner_id'
	OwnerId int64 `json:"owner_id,omitempty"`
	// ID of the token.
	// Wire name: 'token_id'
	TokenId string `json:"token_id,omitempty"`
	// If applicable, the ID of the workspace that the token was created in.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenPermission) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenPermissionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for TokenPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *TokenPermissionLevel) Values() []TokenPermissionLevel {
	return []TokenPermissionLevel{
		TokenPermissionLevelCanUse,
	}
}

// Type always returns TokenPermissionLevel to satisfy [pflag.Value] interface
func (f *TokenPermissionLevel) Type() string {
	return "TokenPermissionLevel"
}

type TokenPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenPermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenPermissionsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *TokenPermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenPermissionsDescriptionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

func (st *TokenPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := tokenPermissionsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Values returns all possible values for TokenType.
//
// There is no guarantee on the order of the values in the slice.
func (f *TokenType) Values() []TokenType {
	return []TokenType{
		TokenTypeArclightAzureExchangeToken,
		TokenTypeArclightAzureExchangeTokenWithUserDelegationKey,
		TokenTypeArclightMultiTenantAzureExchangeToken,
		TokenTypeArclightMultiTenantAzureExchangeTokenWithUserDelegationKey,
		TokenTypeAzureActiveDirectoryToken,
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting AccountIpAccessEnable `json:"setting"`
}

func (st *UpdateAccountIpAccessEnableRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateAccountIpAccessEnableRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting AibiDashboardEmbeddingAccessPolicySetting `json:"setting"`
}

func (st *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting AibiDashboardEmbeddingApprovedDomainsSetting `json:"setting"`
}

func (st *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting AutomaticClusterUpdateSetting `json:"setting"`
}

func (st *UpdateAutomaticClusterUpdateSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateAutomaticClusterUpdateSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting ComplianceSecurityProfileSetting `json:"setting"`
}

func (st *UpdateComplianceSecurityProfileSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateComplianceSecurityProfileSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting CspEnablementAccountSetting `json:"setting"`
}

func (st *UpdateCspEnablementAccountSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCspEnablementAccountSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
type UpdateDashboardEmailSubscriptionsRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting DashboardEmailSubscriptions `json:"setting"`
}

func (st *UpdateDashboardEmailSubscriptionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateDashboardEmailSubscriptionsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateDashboardEmailSubscriptionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateDashboardEmailSubscriptionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateDashboardEmailSubscriptionsRequestToPb(&st)
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting DefaultNamespaceSetting `json:"setting"`
}

func (st *UpdateDefaultNamespaceSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateDefaultNamespaceSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting DisableLegacyAccess `json:"setting"`
}

func (st *UpdateDisableLegacyAccessRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateDisableLegacyAccessRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting DisableLegacyDbfs `json:"setting"`
}

func (st *UpdateDisableLegacyDbfsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateDisableLegacyDbfsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting DisableLegacyFeatures `json:"setting"`
}

func (st *UpdateDisableLegacyFeaturesRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateDisableLegacyFeaturesRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting EnableExportNotebook `json:"setting"`
}

func (st *UpdateEnableExportNotebookRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateEnableExportNotebookRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting EnableNotebookTableClipboard `json:"setting"`
}

func (st *UpdateEnableNotebookTableClipboardRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateEnableNotebookTableClipboardRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting EnableResultsDownloading `json:"setting"`
}

func (st *UpdateEnableResultsDownloadingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateEnableResultsDownloadingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting EnhancedSecurityMonitoringSetting `json:"setting"`
}

func (st *UpdateEnhancedSecurityMonitoringSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateEnhancedSecurityMonitoringSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting EsmEnablementAccountSetting `json:"setting"`
}

func (st *UpdateEsmEnablementAccountSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateEsmEnablementAccountSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Enabled bool `json:"enabled,omitempty"`
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`

	// Wire name: 'ip_addresses'
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// Label for the IP access list. This **cannot** be empty.
	// Wire name: 'label'
	Label string `json:"label,omitempty"`

	// Wire name: 'list_type'
	ListType ListType `json:"list_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateIpAccessList) EncodeValues(key string, v *url.Values) error {
	pb, err := updateIpAccessListToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredAccount `json:"setting"`
}

func (st *UpdateLlmProxyPartnerPoweredAccountRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateLlmProxyPartnerPoweredAccountRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredEnforce `json:"setting"`
}

func (st *UpdateLlmProxyPartnerPoweredEnforceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateLlmProxyPartnerPoweredEnforceRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredWorkspace `json:"setting"`
}

func (st *UpdateLlmProxyPartnerPoweredWorkspaceRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateLlmProxyPartnerPoweredWorkspaceRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UpdateNccPrivateEndpointRuleRequest struct {
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`

	// Wire name: 'private_endpoint_rule'
	PrivateEndpointRule UpdatePrivateEndpointRule `json:"private_endpoint_rule"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" tf:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	UpdateMask string `json:"-" tf:"-"`
}

func (st *UpdateNccPrivateEndpointRuleRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateNccPrivateEndpointRuleRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateNccPrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateNccPrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateNccPrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateNccPrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateNccPrivateEndpointRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateNetworkPolicyRequest struct {
	// Updated network policy configuration details.
	// Wire name: 'network_policy'
	NetworkPolicy AccountNetworkPolicy `json:"network_policy"`
	// The unique identifier for the network policy.
	NetworkPolicyId string `json:"-" tf:"-"`
}

func (st *UpdateNetworkPolicyRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateNetworkPolicyRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	Config *Config `json:"config,omitempty"`
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	Id string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdateNotificationDestinationRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateNotificationDestinationRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting PersonalComputeSetting `json:"setting"`
}

func (st *UpdatePersonalComputeSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updatePersonalComputeSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	// Only used by private endpoints to customer-managed private endpoint
	// services.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	// Wire name: 'domain_names'
	DomainNames []string `json:"domain_names,omitempty"`
	// Only used by private endpoints towards an AWS S3 service.
	//
	// Update this field to activate/deactivate this private endpoint to allow
	// egress access from serverless compute resources.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// Only used by private endpoints towards AWS S3 service.
	//
	// The globally unique S3 bucket names that will be accessed via the VPC
	// endpoint. The bucket names must be in the same region as the NCC/endpoint
	// service. When updating this field, we perform full update on this field.
	// Please ensure a full list of desired resource_names is provided.
	// Wire name: 'resource_names'
	ResourceNames []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UpdatePrivateEndpointRule) EncodeValues(key string, v *url.Values) error {
	pb, err := updatePrivateEndpointRuleToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting RestrictWorkspaceAdminsSetting `json:"setting"`
}

func (st *UpdateRestrictWorkspaceAdminsSettingRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateRestrictWorkspaceAdminsSettingRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

// Details required to update a setting.
type UpdateSqlResultsDownloadRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
	// Wire name: 'allow_missing'
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
	// Wire name: 'field_mask'
	FieldMask string `json:"field_mask"`

	// Wire name: 'setting'
	Setting SqlResultsDownload `json:"setting"`
}

func (st *UpdateSqlResultsDownloadRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateSqlResultsDownloadRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateSqlResultsDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateSqlResultsDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateSqlResultsDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
	// The network option details for the workspace.
	// Wire name: 'workspace_network_option'
	WorkspaceNetworkOption WorkspaceNetworkOption `json:"workspace_network_option"`
}

func (st *UpdateWorkspaceNetworkOptionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateWorkspaceNetworkOptionRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
	NetworkPolicyId string `json:"network_policy_id,omitempty"`
	// The workspace ID.
	// Wire name: 'workspace_id'
	WorkspaceId int64 `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *WorkspaceNetworkOption) EncodeValues(key string, v *url.Values) error {
	pb, err := workspaceNetworkOptionToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
