// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func accountIpAccessEnableToPb(st *AccountIpAccessEnable) (*accountIpAccessEnablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountIpAccessEnablePb{}
	pb.AcctIpAclEnable = st.AcctIpAclEnable
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type accountIpAccessEnablePb struct {
	AcctIpAclEnable BooleanMessage `json:"acct_ip_acl_enable"`
	Etag            string         `json:"etag,omitempty"`
	SettingName     string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func accountIpAccessEnableFromPb(pb *accountIpAccessEnablePb) (*AccountIpAccessEnable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountIpAccessEnable{}
	st.AcctIpAclEnable = pb.AcctIpAclEnable
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *accountIpAccessEnablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st accountIpAccessEnablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func accountNetworkPolicyToPb(st *AccountNetworkPolicy) (*accountNetworkPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountNetworkPolicyPb{}
	pb.AccountId = st.AccountId
	pb.Egress = st.Egress
	pb.NetworkPolicyId = st.NetworkPolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type accountNetworkPolicyPb struct {
	AccountId       string               `json:"account_id,omitempty"`
	Egress          *NetworkPolicyEgress `json:"egress,omitempty"`
	NetworkPolicyId string               `json:"network_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func accountNetworkPolicyFromPb(pb *accountNetworkPolicyPb) (*AccountNetworkPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountNetworkPolicy{}
	st.AccountId = pb.AccountId
	st.Egress = pb.Egress
	st.NetworkPolicyId = pb.NetworkPolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *accountNetworkPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st accountNetworkPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func aibiDashboardEmbeddingAccessPolicyToPb(st *AibiDashboardEmbeddingAccessPolicy) (*aibiDashboardEmbeddingAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingAccessPolicyPb{}
	pb.AccessPolicyType = st.AccessPolicyType

	return pb, nil
}

type aibiDashboardEmbeddingAccessPolicyPb struct {
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType `json:"access_policy_type"`
}

func aibiDashboardEmbeddingAccessPolicyFromPb(pb *aibiDashboardEmbeddingAccessPolicyPb) (*AibiDashboardEmbeddingAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingAccessPolicy{}
	st.AccessPolicyType = pb.AccessPolicyType

	return st, nil
}

func aibiDashboardEmbeddingAccessPolicySettingToPb(st *AibiDashboardEmbeddingAccessPolicySetting) (*aibiDashboardEmbeddingAccessPolicySettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingAccessPolicySettingPb{}
	pb.AibiDashboardEmbeddingAccessPolicy = st.AibiDashboardEmbeddingAccessPolicy
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type aibiDashboardEmbeddingAccessPolicySettingPb struct {
	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicy `json:"aibi_dashboard_embedding_access_policy"`
	Etag                               string                             `json:"etag,omitempty"`
	SettingName                        string                             `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aibiDashboardEmbeddingAccessPolicySettingFromPb(pb *aibiDashboardEmbeddingAccessPolicySettingPb) (*AibiDashboardEmbeddingAccessPolicySetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingAccessPolicySetting{}
	st.AibiDashboardEmbeddingAccessPolicy = pb.AibiDashboardEmbeddingAccessPolicy
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aibiDashboardEmbeddingAccessPolicySettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aibiDashboardEmbeddingAccessPolicySettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func aibiDashboardEmbeddingApprovedDomainsToPb(st *AibiDashboardEmbeddingApprovedDomains) (*aibiDashboardEmbeddingApprovedDomainsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingApprovedDomainsPb{}
	pb.ApprovedDomains = st.ApprovedDomains

	return pb, nil
}

type aibiDashboardEmbeddingApprovedDomainsPb struct {
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

func aibiDashboardEmbeddingApprovedDomainsFromPb(pb *aibiDashboardEmbeddingApprovedDomainsPb) (*AibiDashboardEmbeddingApprovedDomains, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingApprovedDomains{}
	st.ApprovedDomains = pb.ApprovedDomains

	return st, nil
}

func aibiDashboardEmbeddingApprovedDomainsSettingToPb(st *AibiDashboardEmbeddingApprovedDomainsSetting) (*aibiDashboardEmbeddingApprovedDomainsSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingApprovedDomainsSettingPb{}
	pb.AibiDashboardEmbeddingApprovedDomains = st.AibiDashboardEmbeddingApprovedDomains
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type aibiDashboardEmbeddingApprovedDomainsSettingPb struct {
	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomains `json:"aibi_dashboard_embedding_approved_domains"`
	Etag                                  string                                `json:"etag,omitempty"`
	SettingName                           string                                `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aibiDashboardEmbeddingApprovedDomainsSettingFromPb(pb *aibiDashboardEmbeddingApprovedDomainsSettingPb) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingApprovedDomainsSetting{}
	st.AibiDashboardEmbeddingApprovedDomains = pb.AibiDashboardEmbeddingApprovedDomains
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aibiDashboardEmbeddingApprovedDomainsSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aibiDashboardEmbeddingApprovedDomainsSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func automaticClusterUpdateSettingToPb(st *AutomaticClusterUpdateSetting) (*automaticClusterUpdateSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &automaticClusterUpdateSettingPb{}
	pb.AutomaticClusterUpdateWorkspace = st.AutomaticClusterUpdateWorkspace
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type automaticClusterUpdateSettingPb struct {
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessage `json:"automatic_cluster_update_workspace"`
	Etag                            string                    `json:"etag,omitempty"`
	SettingName                     string                    `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func automaticClusterUpdateSettingFromPb(pb *automaticClusterUpdateSettingPb) (*AutomaticClusterUpdateSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutomaticClusterUpdateSetting{}
	st.AutomaticClusterUpdateWorkspace = pb.AutomaticClusterUpdateWorkspace
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *automaticClusterUpdateSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st automaticClusterUpdateSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func booleanMessageToPb(st *BooleanMessage) (*booleanMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &booleanMessagePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type booleanMessagePb struct {
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func booleanMessageFromPb(pb *booleanMessagePb) (*BooleanMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BooleanMessage{}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *booleanMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st booleanMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterAutoRestartMessageToPb(st *ClusterAutoRestartMessage) (*clusterAutoRestartMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessagePb{}
	pb.CanToggle = st.CanToggle
	pb.Enabled = st.Enabled
	pb.EnablementDetails = st.EnablementDetails
	pb.MaintenanceWindow = st.MaintenanceWindow
	pb.RestartEvenIfNoUpdatesAvailable = st.RestartEvenIfNoUpdatesAvailable

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterAutoRestartMessagePb struct {
	CanToggle                       bool                                        `json:"can_toggle,omitempty"`
	Enabled                         bool                                        `json:"enabled,omitempty"`
	EnablementDetails               *ClusterAutoRestartMessageEnablementDetails `json:"enablement_details,omitempty"`
	MaintenanceWindow               *ClusterAutoRestartMessageMaintenanceWindow `json:"maintenance_window,omitempty"`
	RestartEvenIfNoUpdatesAvailable bool                                        `json:"restart_even_if_no_updates_available,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAutoRestartMessageFromPb(pb *clusterAutoRestartMessagePb) (*ClusterAutoRestartMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessage{}
	st.CanToggle = pb.CanToggle
	st.Enabled = pb.Enabled
	st.EnablementDetails = pb.EnablementDetails
	st.MaintenanceWindow = pb.MaintenanceWindow
	st.RestartEvenIfNoUpdatesAvailable = pb.RestartEvenIfNoUpdatesAvailable

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAutoRestartMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAutoRestartMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterAutoRestartMessageEnablementDetailsToPb(st *ClusterAutoRestartMessageEnablementDetails) (*clusterAutoRestartMessageEnablementDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageEnablementDetailsPb{}
	pb.ForcedForComplianceMode = st.ForcedForComplianceMode
	pb.UnavailableForDisabledEntitlement = st.UnavailableForDisabledEntitlement
	pb.UnavailableForNonEnterpriseTier = st.UnavailableForNonEnterpriseTier

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterAutoRestartMessageEnablementDetailsPb struct {
	ForcedForComplianceMode           bool `json:"forced_for_compliance_mode,omitempty"`
	UnavailableForDisabledEntitlement bool `json:"unavailable_for_disabled_entitlement,omitempty"`
	UnavailableForNonEnterpriseTier   bool `json:"unavailable_for_non_enterprise_tier,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAutoRestartMessageEnablementDetailsFromPb(pb *clusterAutoRestartMessageEnablementDetailsPb) (*ClusterAutoRestartMessageEnablementDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageEnablementDetails{}
	st.ForcedForComplianceMode = pb.ForcedForComplianceMode
	st.UnavailableForDisabledEntitlement = pb.UnavailableForDisabledEntitlement
	st.UnavailableForNonEnterpriseTier = pb.UnavailableForNonEnterpriseTier

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAutoRestartMessageEnablementDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAutoRestartMessageEnablementDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func clusterAutoRestartMessageMaintenanceWindowToPb(st *ClusterAutoRestartMessageMaintenanceWindow) (*clusterAutoRestartMessageMaintenanceWindowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowPb{}
	pb.WeekDayBasedSchedule = st.WeekDayBasedSchedule

	return pb, nil
}

type clusterAutoRestartMessageMaintenanceWindowPb struct {
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `json:"week_day_based_schedule,omitempty"`
}

func clusterAutoRestartMessageMaintenanceWindowFromPb(pb *clusterAutoRestartMessageMaintenanceWindowPb) (*ClusterAutoRestartMessageMaintenanceWindow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindow{}
	st.WeekDayBasedSchedule = pb.WeekDayBasedSchedule

	return st, nil
}

func clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(st *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) (*clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb{}
	pb.DayOfWeek = st.DayOfWeek
	pb.Frequency = st.Frequency
	pb.WindowStartTime = st.WindowStartTime

	return pb, nil
}

type clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb struct {
	DayOfWeek       ClusterAutoRestartMessageMaintenanceWindowDayOfWeek        `json:"day_of_week,omitempty"`
	Frequency       ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `json:"frequency,omitempty"`
	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `json:"window_start_time,omitempty"`
}

func clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleFromPb(pb *clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb) (*ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}
	st.DayOfWeek = pb.DayOfWeek
	st.Frequency = pb.Frequency
	st.WindowStartTime = pb.WindowStartTime

	return st, nil
}

func clusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(st *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) (*clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb{}
	pb.Hours = st.Hours
	pb.Minutes = st.Minutes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb struct {
	Hours   int `json:"hours,omitempty"`
	Minutes int `json:"minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAutoRestartMessageMaintenanceWindowWindowStartTimeFromPb(pb *clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) (*ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}
	st.Hours = pb.Hours
	st.Minutes = pb.Minutes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func complianceSecurityProfileToPb(st *ComplianceSecurityProfile) (*complianceSecurityProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &complianceSecurityProfilePb{}
	pb.ComplianceStandards = st.ComplianceStandards
	pb.IsEnabled = st.IsEnabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type complianceSecurityProfilePb struct {
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	IsEnabled           bool                 `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func complianceSecurityProfileFromPb(pb *complianceSecurityProfilePb) (*ComplianceSecurityProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfile{}
	st.ComplianceStandards = pb.ComplianceStandards
	st.IsEnabled = pb.IsEnabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *complianceSecurityProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st complianceSecurityProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func complianceSecurityProfileSettingToPb(st *ComplianceSecurityProfileSetting) (*complianceSecurityProfileSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &complianceSecurityProfileSettingPb{}
	pb.ComplianceSecurityProfileWorkspace = st.ComplianceSecurityProfileWorkspace
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type complianceSecurityProfileSettingPb struct {
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfile `json:"compliance_security_profile_workspace"`
	Etag                               string                    `json:"etag,omitempty"`
	SettingName                        string                    `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func complianceSecurityProfileSettingFromPb(pb *complianceSecurityProfileSettingPb) (*ComplianceSecurityProfileSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfileSetting{}
	st.ComplianceSecurityProfileWorkspace = pb.ComplianceSecurityProfileWorkspace
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *complianceSecurityProfileSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st complianceSecurityProfileSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func configToPb(st *Config) (*configPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &configPb{}
	pb.Email = st.Email
	pb.GenericWebhook = st.GenericWebhook
	pb.MicrosoftTeams = st.MicrosoftTeams
	pb.Pagerduty = st.Pagerduty
	pb.Slack = st.Slack

	return pb, nil
}

type configPb struct {
	Email          *EmailConfig          `json:"email,omitempty"`
	GenericWebhook *GenericWebhookConfig `json:"generic_webhook,omitempty"`
	MicrosoftTeams *MicrosoftTeamsConfig `json:"microsoft_teams,omitempty"`
	Pagerduty      *PagerdutyConfig      `json:"pagerduty,omitempty"`
	Slack          *SlackConfig          `json:"slack,omitempty"`
}

func configFromPb(pb *configPb) (*Config, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Config{}
	st.Email = pb.Email
	st.GenericWebhook = pb.GenericWebhook
	st.MicrosoftTeams = pb.MicrosoftTeams
	st.Pagerduty = pb.Pagerduty
	st.Slack = pb.Slack

	return st, nil
}

func createIpAccessListToPb(st *CreateIpAccessList) (*createIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createIpAccessListPb{}
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	pb.ListType = st.ListType

	return pb, nil
}

type createIpAccessListPb struct {
	IpAddresses []string `json:"ip_addresses,omitempty"`
	Label       string   `json:"label"`
	ListType    ListType `json:"list_type"`
}

func createIpAccessListFromPb(pb *createIpAccessListPb) (*CreateIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateIpAccessList{}
	st.IpAddresses = pb.IpAddresses
	st.Label = pb.Label
	st.ListType = pb.ListType

	return st, nil
}

func createIpAccessListResponseToPb(st *CreateIpAccessListResponse) (*createIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createIpAccessListResponsePb{}
	pb.IpAccessList = st.IpAccessList

	return pb, nil
}

type createIpAccessListResponsePb struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func createIpAccessListResponseFromPb(pb *createIpAccessListResponsePb) (*CreateIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateIpAccessListResponse{}
	st.IpAccessList = pb.IpAccessList

	return st, nil
}

func createNetworkConnectivityConfigRequestToPb(st *CreateNetworkConnectivityConfigRequest) (*createNetworkConnectivityConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNetworkConnectivityConfigRequestPb{}
	pb.NetworkConnectivityConfig = st.NetworkConnectivityConfig

	return pb, nil
}

type createNetworkConnectivityConfigRequestPb struct {
	NetworkConnectivityConfig CreateNetworkConnectivityConfiguration `json:"network_connectivity_config"`
}

func createNetworkConnectivityConfigRequestFromPb(pb *createNetworkConnectivityConfigRequestPb) (*CreateNetworkConnectivityConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkConnectivityConfigRequest{}
	st.NetworkConnectivityConfig = pb.NetworkConnectivityConfig

	return st, nil
}

func createNetworkConnectivityConfigurationToPb(st *CreateNetworkConnectivityConfiguration) (*createNetworkConnectivityConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNetworkConnectivityConfigurationPb{}
	pb.Name = st.Name
	pb.Region = st.Region

	return pb, nil
}

type createNetworkConnectivityConfigurationPb struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

func createNetworkConnectivityConfigurationFromPb(pb *createNetworkConnectivityConfigurationPb) (*CreateNetworkConnectivityConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkConnectivityConfiguration{}
	st.Name = pb.Name
	st.Region = pb.Region

	return st, nil
}

func createNetworkPolicyRequestToPb(st *CreateNetworkPolicyRequest) (*createNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNetworkPolicyRequestPb{}
	pb.NetworkPolicy = st.NetworkPolicy

	return pb, nil
}

type createNetworkPolicyRequestPb struct {
	NetworkPolicy AccountNetworkPolicy `json:"network_policy"`
}

func createNetworkPolicyRequestFromPb(pb *createNetworkPolicyRequestPb) (*CreateNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkPolicyRequest{}
	st.NetworkPolicy = pb.NetworkPolicy

	return st, nil
}

func createNotificationDestinationRequestToPb(st *CreateNotificationDestinationRequest) (*createNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNotificationDestinationRequestPb{}
	pb.Config = st.Config
	pb.DisplayName = st.DisplayName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createNotificationDestinationRequestPb struct {
	Config      *Config `json:"config,omitempty"`
	DisplayName string  `json:"display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createNotificationDestinationRequestFromPb(pb *createNotificationDestinationRequestPb) (*CreateNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNotificationDestinationRequest{}
	st.Config = pb.Config
	st.DisplayName = pb.DisplayName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createNotificationDestinationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createNotificationDestinationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createOboTokenRequestToPb(st *CreateOboTokenRequest) (*createOboTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createOboTokenRequestPb{}
	pb.ApplicationId = st.ApplicationId
	pb.Comment = st.Comment
	pb.LifetimeSeconds = st.LifetimeSeconds

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createOboTokenRequestPb struct {
	ApplicationId   string `json:"application_id"`
	Comment         string `json:"comment,omitempty"`
	LifetimeSeconds int64  `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createOboTokenRequestFromPb(pb *createOboTokenRequestPb) (*CreateOboTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOboTokenRequest{}
	st.ApplicationId = pb.ApplicationId
	st.Comment = pb.Comment
	st.LifetimeSeconds = pb.LifetimeSeconds

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createOboTokenRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createOboTokenRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createOboTokenResponseToPb(st *CreateOboTokenResponse) (*createOboTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createOboTokenResponsePb{}
	pb.TokenInfo = st.TokenInfo
	pb.TokenValue = st.TokenValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createOboTokenResponsePb struct {
	TokenInfo  *TokenInfo `json:"token_info,omitempty"`
	TokenValue string     `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createOboTokenResponseFromPb(pb *createOboTokenResponsePb) (*CreateOboTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOboTokenResponse{}
	st.TokenInfo = pb.TokenInfo
	st.TokenValue = pb.TokenValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createOboTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createOboTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPrivateEndpointRuleToPb(st *CreatePrivateEndpointRule) (*createPrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPrivateEndpointRulePb{}
	pb.DomainNames = st.DomainNames
	pb.EndpointService = st.EndpointService
	pb.GroupId = st.GroupId
	pb.ResourceId = st.ResourceId
	pb.ResourceNames = st.ResourceNames

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPrivateEndpointRulePb struct {
	DomainNames     []string `json:"domain_names,omitempty"`
	EndpointService string   `json:"endpoint_service,omitempty"`
	GroupId         string   `json:"group_id,omitempty"`
	ResourceId      string   `json:"resource_id,omitempty"`
	ResourceNames   []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPrivateEndpointRuleFromPb(pb *createPrivateEndpointRulePb) (*CreatePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePrivateEndpointRule{}
	st.DomainNames = pb.DomainNames
	st.EndpointService = pb.EndpointService
	st.GroupId = pb.GroupId
	st.ResourceId = pb.ResourceId
	st.ResourceNames = pb.ResourceNames

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPrivateEndpointRuleRequestToPb(st *CreatePrivateEndpointRuleRequest) (*createPrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PrivateEndpointRule = st.PrivateEndpointRule

	return pb, nil
}

type createPrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string                    `json:"-" url:"-"`
	PrivateEndpointRule         CreatePrivateEndpointRule `json:"private_endpoint_rule"`
}

func createPrivateEndpointRuleRequestFromPb(pb *createPrivateEndpointRuleRequestPb) (*CreatePrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PrivateEndpointRule = pb.PrivateEndpointRule

	return st, nil
}

func createTokenRequestToPb(st *CreateTokenRequest) (*createTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTokenRequestPb{}
	pb.Comment = st.Comment
	pb.LifetimeSeconds = st.LifetimeSeconds

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createTokenRequestPb struct {
	Comment         string `json:"comment,omitempty"`
	LifetimeSeconds int64  `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createTokenRequestFromPb(pb *createTokenRequestPb) (*CreateTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTokenRequest{}
	st.Comment = pb.Comment
	st.LifetimeSeconds = pb.LifetimeSeconds

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createTokenRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createTokenRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createTokenResponseToPb(st *CreateTokenResponse) (*createTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTokenResponsePb{}
	pb.TokenInfo = st.TokenInfo
	pb.TokenValue = st.TokenValue

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createTokenResponsePb struct {
	TokenInfo  *PublicTokenInfo `json:"token_info,omitempty"`
	TokenValue string           `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createTokenResponseFromPb(pb *createTokenResponsePb) (*CreateTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTokenResponse{}
	st.TokenInfo = pb.TokenInfo
	st.TokenValue = pb.TokenValue

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cspEnablementAccountToPb(st *CspEnablementAccount) (*cspEnablementAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cspEnablementAccountPb{}
	pb.ComplianceStandards = st.ComplianceStandards
	pb.IsEnforced = st.IsEnforced

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cspEnablementAccountPb struct {
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	IsEnforced          bool                 `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cspEnablementAccountFromPb(pb *cspEnablementAccountPb) (*CspEnablementAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CspEnablementAccount{}
	st.ComplianceStandards = pb.ComplianceStandards
	st.IsEnforced = pb.IsEnforced

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cspEnablementAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cspEnablementAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cspEnablementAccountSettingToPb(st *CspEnablementAccountSetting) (*cspEnablementAccountSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cspEnablementAccountSettingPb{}
	pb.CspEnablementAccount = st.CspEnablementAccount
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cspEnablementAccountSettingPb struct {
	CspEnablementAccount CspEnablementAccount `json:"csp_enablement_account"`
	Etag                 string               `json:"etag,omitempty"`
	SettingName          string               `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cspEnablementAccountSettingFromPb(pb *cspEnablementAccountSettingPb) (*CspEnablementAccountSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CspEnablementAccountSetting{}
	st.CspEnablementAccount = pb.CspEnablementAccount
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cspEnablementAccountSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cspEnablementAccountSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func customerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleToPb(st *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) (*customerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb{}
	pb.AccountId = st.AccountId
	pb.ConnectionState = st.ConnectionState
	pb.CreationTime = st.CreationTime
	pb.Deactivated = st.Deactivated
	pb.DeactivatedAt = st.DeactivatedAt
	pb.DomainNames = st.DomainNames
	pb.Enabled = st.Enabled
	pb.EndpointService = st.EndpointService
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.ResourceNames = st.ResourceNames
	pb.RuleId = st.RuleId
	pb.UpdatedTime = st.UpdatedTime
	pb.VpcEndpointId = st.VpcEndpointId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type customerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb struct {
	AccountId                   string                                                                                  `json:"account_id,omitempty"`
	ConnectionState             CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState `json:"connection_state,omitempty"`
	CreationTime                int64                                                                                   `json:"creation_time,omitempty"`
	Deactivated                 bool                                                                                    `json:"deactivated,omitempty"`
	DeactivatedAt               int64                                                                                   `json:"deactivated_at,omitempty"`
	DomainNames                 []string                                                                                `json:"domain_names,omitempty"`
	Enabled                     bool                                                                                    `json:"enabled,omitempty"`
	EndpointService             string                                                                                  `json:"endpoint_service,omitempty"`
	NetworkConnectivityConfigId string                                                                                  `json:"network_connectivity_config_id,omitempty"`
	ResourceNames               []string                                                                                `json:"resource_names,omitempty"`
	RuleId                      string                                                                                  `json:"rule_id,omitempty"`
	UpdatedTime                 int64                                                                                   `json:"updated_time,omitempty"`
	VpcEndpointId               string                                                                                  `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleFromPb(pb *customerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb) (*CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule{}
	st.AccountId = pb.AccountId
	st.ConnectionState = pb.ConnectionState
	st.CreationTime = pb.CreationTime
	st.Deactivated = pb.Deactivated
	st.DeactivatedAt = pb.DeactivatedAt
	st.DomainNames = pb.DomainNames
	st.Enabled = pb.Enabled
	st.EndpointService = pb.EndpointService
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.ResourceNames = pb.ResourceNames
	st.RuleId = pb.RuleId
	st.UpdatedTime = pb.UpdatedTime
	st.VpcEndpointId = pb.VpcEndpointId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dashboardEmailSubscriptionsToPb(st *DashboardEmailSubscriptions) (*dashboardEmailSubscriptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dashboardEmailSubscriptionsPb{}
	pb.BooleanVal = st.BooleanVal
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dashboardEmailSubscriptionsPb struct {
	BooleanVal  BooleanMessage `json:"boolean_val"`
	Etag        string         `json:"etag,omitempty"`
	SettingName string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dashboardEmailSubscriptionsFromPb(pb *dashboardEmailSubscriptionsPb) (*DashboardEmailSubscriptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardEmailSubscriptions{}
	st.BooleanVal = pb.BooleanVal
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dashboardEmailSubscriptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dashboardEmailSubscriptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func defaultNamespaceSettingToPb(st *DefaultNamespaceSetting) (*defaultNamespaceSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &defaultNamespaceSettingPb{}
	pb.Etag = st.Etag
	pb.Namespace = st.Namespace
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type defaultNamespaceSettingPb struct {
	Etag        string        `json:"etag,omitempty"`
	Namespace   StringMessage `json:"namespace"`
	SettingName string        `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func defaultNamespaceSettingFromPb(pb *defaultNamespaceSettingPb) (*DefaultNamespaceSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DefaultNamespaceSetting{}
	st.Etag = pb.Etag
	st.Namespace = pb.Namespace
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *defaultNamespaceSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st defaultNamespaceSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAccountIpAccessEnableRequestToPb(st *DeleteAccountIpAccessEnableRequest) (*deleteAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountIpAccessEnableRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteAccountIpAccessEnableRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteAccountIpAccessEnableRequestFromPb(pb *deleteAccountIpAccessEnableRequestPb) (*DeleteAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessEnableRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAccountIpAccessEnableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAccountIpAccessEnableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAccountIpAccessEnableResponseToPb(st *DeleteAccountIpAccessEnableResponse) (*deleteAccountIpAccessEnableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountIpAccessEnableResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteAccountIpAccessEnableResponsePb struct {
	Etag string `json:"etag"`
}

func deleteAccountIpAccessEnableResponseFromPb(pb *deleteAccountIpAccessEnableResponsePb) (*DeleteAccountIpAccessEnableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessEnableResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteAccountIpAccessListRequestToPb(st *DeleteAccountIpAccessListRequest) (*deleteAccountIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

type deleteAccountIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

func deleteAccountIpAccessListRequestFromPb(pb *deleteAccountIpAccessListRequestPb) (*DeleteAccountIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

func deleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*DeleteAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingAccessPolicySettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAibiDashboardEmbeddingAccessPolicySettingResponseToPb(st *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) (*deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb struct {
	Etag string `json:"etag"`
}

func deleteAibiDashboardEmbeddingAccessPolicySettingResponseFromPb(pb *deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb) (*DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingAccessPolicySettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseToPb(st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) (*deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb struct {
	Etag string `json:"etag"`
}

func deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseFromPb(pb *deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteDashboardEmailSubscriptionsRequestToPb(st *DeleteDashboardEmailSubscriptionsRequest) (*deleteDashboardEmailSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardEmailSubscriptionsRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteDashboardEmailSubscriptionsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDashboardEmailSubscriptionsRequestFromPb(pb *deleteDashboardEmailSubscriptionsRequestPb) (*DeleteDashboardEmailSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardEmailSubscriptionsRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDashboardEmailSubscriptionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDashboardEmailSubscriptionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDashboardEmailSubscriptionsResponseToPb(st *DeleteDashboardEmailSubscriptionsResponse) (*deleteDashboardEmailSubscriptionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDashboardEmailSubscriptionsResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteDashboardEmailSubscriptionsResponsePb struct {
	Etag string `json:"etag"`
}

func deleteDashboardEmailSubscriptionsResponseFromPb(pb *deleteDashboardEmailSubscriptionsResponsePb) (*DeleteDashboardEmailSubscriptionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardEmailSubscriptionsResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteDefaultNamespaceSettingRequestToPb(st *DeleteDefaultNamespaceSettingRequest) (*deleteDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDefaultNamespaceSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteDefaultNamespaceSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDefaultNamespaceSettingRequestFromPb(pb *deleteDefaultNamespaceSettingRequestPb) (*DeleteDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultNamespaceSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDefaultNamespaceSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDefaultNamespaceSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDefaultNamespaceSettingResponseToPb(st *DeleteDefaultNamespaceSettingResponse) (*deleteDefaultNamespaceSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDefaultNamespaceSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteDefaultNamespaceSettingResponsePb struct {
	Etag string `json:"etag"`
}

func deleteDefaultNamespaceSettingResponseFromPb(pb *deleteDefaultNamespaceSettingResponsePb) (*DeleteDefaultNamespaceSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultNamespaceSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteDisableLegacyAccessRequestToPb(st *DeleteDisableLegacyAccessRequest) (*deleteDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyAccessRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteDisableLegacyAccessRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDisableLegacyAccessRequestFromPb(pb *deleteDisableLegacyAccessRequestPb) (*DeleteDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyAccessRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDisableLegacyAccessRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDisableLegacyAccessRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDisableLegacyAccessResponseToPb(st *DeleteDisableLegacyAccessResponse) (*deleteDisableLegacyAccessResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyAccessResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteDisableLegacyAccessResponsePb struct {
	Etag string `json:"etag"`
}

func deleteDisableLegacyAccessResponseFromPb(pb *deleteDisableLegacyAccessResponsePb) (*DeleteDisableLegacyAccessResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyAccessResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteDisableLegacyDbfsRequestToPb(st *DeleteDisableLegacyDbfsRequest) (*deleteDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyDbfsRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteDisableLegacyDbfsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDisableLegacyDbfsRequestFromPb(pb *deleteDisableLegacyDbfsRequestPb) (*DeleteDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyDbfsRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDisableLegacyDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDisableLegacyDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDisableLegacyDbfsResponseToPb(st *DeleteDisableLegacyDbfsResponse) (*deleteDisableLegacyDbfsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyDbfsResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteDisableLegacyDbfsResponsePb struct {
	Etag string `json:"etag"`
}

func deleteDisableLegacyDbfsResponseFromPb(pb *deleteDisableLegacyDbfsResponsePb) (*DeleteDisableLegacyDbfsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyDbfsResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteDisableLegacyFeaturesRequestToPb(st *DeleteDisableLegacyFeaturesRequest) (*deleteDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyFeaturesRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteDisableLegacyFeaturesRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteDisableLegacyFeaturesRequestFromPb(pb *deleteDisableLegacyFeaturesRequestPb) (*DeleteDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyFeaturesRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDisableLegacyFeaturesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDisableLegacyFeaturesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDisableLegacyFeaturesResponseToPb(st *DeleteDisableLegacyFeaturesResponse) (*deleteDisableLegacyFeaturesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyFeaturesResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteDisableLegacyFeaturesResponsePb struct {
	Etag string `json:"etag"`
}

func deleteDisableLegacyFeaturesResponseFromPb(pb *deleteDisableLegacyFeaturesResponsePb) (*DeleteDisableLegacyFeaturesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyFeaturesResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteIpAccessListRequestToPb(st *DeleteIpAccessListRequest) (*deleteIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

type deleteIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

func deleteIpAccessListRequestFromPb(pb *deleteIpAccessListRequestPb) (*DeleteIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

func deleteLlmProxyPartnerPoweredWorkspaceRequestToPb(st *DeleteLlmProxyPartnerPoweredWorkspaceRequest) (*deleteLlmProxyPartnerPoweredWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteLlmProxyPartnerPoweredWorkspaceRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteLlmProxyPartnerPoweredWorkspaceRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb *deleteLlmProxyPartnerPoweredWorkspaceRequestPb) (*DeleteLlmProxyPartnerPoweredWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLlmProxyPartnerPoweredWorkspaceRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteLlmProxyPartnerPoweredWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteLlmProxyPartnerPoweredWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteLlmProxyPartnerPoweredWorkspaceResponseToPb(st *DeleteLlmProxyPartnerPoweredWorkspaceResponse) (*deleteLlmProxyPartnerPoweredWorkspaceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteLlmProxyPartnerPoweredWorkspaceResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteLlmProxyPartnerPoweredWorkspaceResponsePb struct {
	Etag string `json:"etag"`
}

func deleteLlmProxyPartnerPoweredWorkspaceResponseFromPb(pb *deleteLlmProxyPartnerPoweredWorkspaceResponsePb) (*DeleteLlmProxyPartnerPoweredWorkspaceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLlmProxyPartnerPoweredWorkspaceResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteNetworkConnectivityConfigurationRequestToPb(st *DeleteNetworkConnectivityConfigurationRequest) (*deleteNetworkConnectivityConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkConnectivityConfigurationRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId

	return pb, nil
}

type deleteNetworkConnectivityConfigurationRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

func deleteNetworkConnectivityConfigurationRequestFromPb(pb *deleteNetworkConnectivityConfigurationRequestPb) (*DeleteNetworkConnectivityConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkConnectivityConfigurationRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId

	return st, nil
}

func deleteNetworkConnectivityConfigurationResponseToPb(st *DeleteNetworkConnectivityConfigurationResponse) (*deleteNetworkConnectivityConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkConnectivityConfigurationResponsePb{}

	return pb, nil
}

type deleteNetworkConnectivityConfigurationResponsePb struct {
}

func deleteNetworkConnectivityConfigurationResponseFromPb(pb *deleteNetworkConnectivityConfigurationResponsePb) (*DeleteNetworkConnectivityConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkConnectivityConfigurationResponse{}

	return st, nil
}

func deleteNetworkPolicyRequestToPb(st *DeleteNetworkPolicyRequest) (*deleteNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkPolicyRequestPb{}
	pb.NetworkPolicyId = st.NetworkPolicyId

	return pb, nil
}

type deleteNetworkPolicyRequestPb struct {
	NetworkPolicyId string `json:"-" url:"-"`
}

func deleteNetworkPolicyRequestFromPb(pb *deleteNetworkPolicyRequestPb) (*DeleteNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkPolicyRequest{}
	st.NetworkPolicyId = pb.NetworkPolicyId

	return st, nil
}

func deleteNetworkPolicyRpcResponseToPb(st *DeleteNetworkPolicyRpcResponse) (*deleteNetworkPolicyRpcResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkPolicyRpcResponsePb{}

	return pb, nil
}

type deleteNetworkPolicyRpcResponsePb struct {
}

func deleteNetworkPolicyRpcResponseFromPb(pb *deleteNetworkPolicyRpcResponsePb) (*DeleteNetworkPolicyRpcResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkPolicyRpcResponse{}

	return st, nil
}

func deleteNotificationDestinationRequestToPb(st *DeleteNotificationDestinationRequest) (*deleteNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNotificationDestinationRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteNotificationDestinationRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteNotificationDestinationRequestFromPb(pb *deleteNotificationDestinationRequestPb) (*DeleteNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNotificationDestinationRequest{}
	st.Id = pb.Id

	return st, nil
}

func deletePersonalComputeSettingRequestToPb(st *DeletePersonalComputeSettingRequest) (*deletePersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePersonalComputeSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deletePersonalComputeSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deletePersonalComputeSettingRequestFromPb(pb *deletePersonalComputeSettingRequestPb) (*DeletePersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePersonalComputeSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deletePersonalComputeSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deletePersonalComputeSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deletePersonalComputeSettingResponseToPb(st *DeletePersonalComputeSettingResponse) (*deletePersonalComputeSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePersonalComputeSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deletePersonalComputeSettingResponsePb struct {
	Etag string `json:"etag"`
}

func deletePersonalComputeSettingResponseFromPb(pb *deletePersonalComputeSettingResponsePb) (*DeletePersonalComputeSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePersonalComputeSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deletePrivateEndpointRuleRequestToPb(st *DeletePrivateEndpointRuleRequest) (*deletePrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PrivateEndpointRuleId = st.PrivateEndpointRuleId

	return pb, nil
}

type deletePrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	PrivateEndpointRuleId       string `json:"-" url:"-"`
}

func deletePrivateEndpointRuleRequestFromPb(pb *deletePrivateEndpointRuleRequestPb) (*DeletePrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PrivateEndpointRuleId = pb.PrivateEndpointRuleId

	return st, nil
}

func deleteRestrictWorkspaceAdminsSettingRequestToPb(st *DeleteRestrictWorkspaceAdminsSettingRequest) (*deleteRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRestrictWorkspaceAdminsSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteRestrictWorkspaceAdminsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteRestrictWorkspaceAdminsSettingRequestFromPb(pb *deleteRestrictWorkspaceAdminsSettingRequestPb) (*DeleteRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRestrictWorkspaceAdminsSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteRestrictWorkspaceAdminsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteRestrictWorkspaceAdminsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteRestrictWorkspaceAdminsSettingResponseToPb(st *DeleteRestrictWorkspaceAdminsSettingResponse) (*deleteRestrictWorkspaceAdminsSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRestrictWorkspaceAdminsSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteRestrictWorkspaceAdminsSettingResponsePb struct {
	Etag string `json:"etag"`
}

func deleteRestrictWorkspaceAdminsSettingResponseFromPb(pb *deleteRestrictWorkspaceAdminsSettingResponsePb) (*DeleteRestrictWorkspaceAdminsSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRestrictWorkspaceAdminsSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteSqlResultsDownloadRequestToPb(st *DeleteSqlResultsDownloadRequest) (*deleteSqlResultsDownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSqlResultsDownloadRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deleteSqlResultsDownloadRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteSqlResultsDownloadRequestFromPb(pb *deleteSqlResultsDownloadRequestPb) (*DeleteSqlResultsDownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSqlResultsDownloadRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteSqlResultsDownloadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteSqlResultsDownloadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteSqlResultsDownloadResponseToPb(st *DeleteSqlResultsDownloadResponse) (*deleteSqlResultsDownloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteSqlResultsDownloadResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

type deleteSqlResultsDownloadResponsePb struct {
	Etag string `json:"etag"`
}

func deleteSqlResultsDownloadResponseFromPb(pb *deleteSqlResultsDownloadResponsePb) (*DeleteSqlResultsDownloadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSqlResultsDownloadResponse{}
	st.Etag = pb.Etag

	return st, nil
}

func deleteTokenManagementRequestToPb(st *DeleteTokenManagementRequest) (*deleteTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTokenManagementRequestPb{}
	pb.TokenId = st.TokenId

	return pb, nil
}

type deleteTokenManagementRequestPb struct {
	TokenId string `json:"-" url:"-"`
}

func deleteTokenManagementRequestFromPb(pb *deleteTokenManagementRequestPb) (*DeleteTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTokenManagementRequest{}
	st.TokenId = pb.TokenId

	return st, nil
}

func disableLegacyAccessToPb(st *DisableLegacyAccess) (*disableLegacyAccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableLegacyAccessPb{}
	pb.DisableLegacyAccess = st.DisableLegacyAccess
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type disableLegacyAccessPb struct {
	DisableLegacyAccess BooleanMessage `json:"disable_legacy_access"`
	Etag                string         `json:"etag,omitempty"`
	SettingName         string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func disableLegacyAccessFromPb(pb *disableLegacyAccessPb) (*DisableLegacyAccess, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyAccess{}
	st.DisableLegacyAccess = pb.DisableLegacyAccess
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *disableLegacyAccessPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st disableLegacyAccessPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func disableLegacyDbfsToPb(st *DisableLegacyDbfs) (*disableLegacyDbfsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableLegacyDbfsPb{}
	pb.DisableLegacyDbfs = st.DisableLegacyDbfs
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type disableLegacyDbfsPb struct {
	DisableLegacyDbfs BooleanMessage `json:"disable_legacy_dbfs"`
	Etag              string         `json:"etag,omitempty"`
	SettingName       string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func disableLegacyDbfsFromPb(pb *disableLegacyDbfsPb) (*DisableLegacyDbfs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyDbfs{}
	st.DisableLegacyDbfs = pb.DisableLegacyDbfs
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *disableLegacyDbfsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st disableLegacyDbfsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func disableLegacyFeaturesToPb(st *DisableLegacyFeatures) (*disableLegacyFeaturesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableLegacyFeaturesPb{}
	pb.DisableLegacyFeatures = st.DisableLegacyFeatures
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type disableLegacyFeaturesPb struct {
	DisableLegacyFeatures BooleanMessage `json:"disable_legacy_features"`
	Etag                  string         `json:"etag,omitempty"`
	SettingName           string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func disableLegacyFeaturesFromPb(pb *disableLegacyFeaturesPb) (*DisableLegacyFeatures, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyFeatures{}
	st.DisableLegacyFeatures = pb.DisableLegacyFeatures
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *disableLegacyFeaturesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st disableLegacyFeaturesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func egressNetworkPolicyToPb(st *EgressNetworkPolicy) (*egressNetworkPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyPb{}
	pb.InternetAccess = st.InternetAccess

	return pb, nil
}

type egressNetworkPolicyPb struct {
	InternetAccess *EgressNetworkPolicyInternetAccessPolicy `json:"internet_access,omitempty"`
}

func egressNetworkPolicyFromPb(pb *egressNetworkPolicyPb) (*EgressNetworkPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicy{}
	st.InternetAccess = pb.InternetAccess

	return st, nil
}

func egressNetworkPolicyInternetAccessPolicyToPb(st *EgressNetworkPolicyInternetAccessPolicy) (*egressNetworkPolicyInternetAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyPb{}
	pb.AllowedInternetDestinations = st.AllowedInternetDestinations
	pb.AllowedStorageDestinations = st.AllowedStorageDestinations
	pb.LogOnlyMode = st.LogOnlyMode
	pb.RestrictionMode = st.RestrictionMode

	return pb, nil
}

type egressNetworkPolicyInternetAccessPolicyPb struct {
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`
	AllowedStorageDestinations  []EgressNetworkPolicyInternetAccessPolicyStorageDestination  `json:"allowed_storage_destinations,omitempty"`
	LogOnlyMode                 *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode          `json:"log_only_mode,omitempty"`
	RestrictionMode             EgressNetworkPolicyInternetAccessPolicyRestrictionMode       `json:"restriction_mode,omitempty"`
}

func egressNetworkPolicyInternetAccessPolicyFromPb(pb *egressNetworkPolicyInternetAccessPolicyPb) (*EgressNetworkPolicyInternetAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicy{}
	st.AllowedInternetDestinations = pb.AllowedInternetDestinations
	st.AllowedStorageDestinations = pb.AllowedStorageDestinations
	st.LogOnlyMode = pb.LogOnlyMode
	st.RestrictionMode = pb.RestrictionMode

	return st, nil
}

func egressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(st *EgressNetworkPolicyInternetAccessPolicyInternetDestination) (*egressNetworkPolicyInternetAccessPolicyInternetDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyInternetDestinationPb{}
	pb.Destination = st.Destination
	pb.Protocol = st.Protocol
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type egressNetworkPolicyInternetAccessPolicyInternetDestinationPb struct {
	Destination string                                                                                         `json:"destination,omitempty"`
	Protocol    EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol `json:"protocol,omitempty"`
	Type        EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType              `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func egressNetworkPolicyInternetAccessPolicyInternetDestinationFromPb(pb *egressNetworkPolicyInternetAccessPolicyInternetDestinationPb) (*EgressNetworkPolicyInternetAccessPolicyInternetDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyInternetDestination{}
	st.Destination = pb.Destination
	st.Protocol = pb.Protocol
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *egressNetworkPolicyInternetAccessPolicyInternetDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st egressNetworkPolicyInternetAccessPolicyInternetDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(st *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) (*egressNetworkPolicyInternetAccessPolicyLogOnlyModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyLogOnlyModePb{}
	pb.LogOnlyModeType = st.LogOnlyModeType
	pb.Workloads = st.Workloads

	return pb, nil
}

type egressNetworkPolicyInternetAccessPolicyLogOnlyModePb struct {
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType `json:"log_only_mode_type,omitempty"`
	Workloads       []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType  `json:"workloads,omitempty"`
}

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeFromPb(pb *egressNetworkPolicyInternetAccessPolicyLogOnlyModePb) (*EgressNetworkPolicyInternetAccessPolicyLogOnlyMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyLogOnlyMode{}
	st.LogOnlyModeType = pb.LogOnlyModeType
	st.Workloads = pb.Workloads

	return st, nil
}

func egressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(st *EgressNetworkPolicyInternetAccessPolicyStorageDestination) (*egressNetworkPolicyInternetAccessPolicyStorageDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyStorageDestinationPb{}
	pb.AllowedPaths = st.AllowedPaths
	pb.AzureContainer = st.AzureContainer
	pb.AzureDnsZone = st.AzureDnsZone
	pb.AzureStorageAccount = st.AzureStorageAccount
	pb.AzureStorageService = st.AzureStorageService
	pb.BucketName = st.BucketName
	pb.Region = st.Region
	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type egressNetworkPolicyInternetAccessPolicyStorageDestinationPb struct {
	AllowedPaths        []string                                                                        `json:"allowed_paths,omitempty"`
	AzureContainer      string                                                                          `json:"azure_container,omitempty"`
	AzureDnsZone        string                                                                          `json:"azure_dns_zone,omitempty"`
	AzureStorageAccount string                                                                          `json:"azure_storage_account,omitempty"`
	AzureStorageService string                                                                          `json:"azure_storage_service,omitempty"`
	BucketName          string                                                                          `json:"bucket_name,omitempty"`
	Region              string                                                                          `json:"region,omitempty"`
	Type                EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func egressNetworkPolicyInternetAccessPolicyStorageDestinationFromPb(pb *egressNetworkPolicyInternetAccessPolicyStorageDestinationPb) (*EgressNetworkPolicyInternetAccessPolicyStorageDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyStorageDestination{}
	st.AllowedPaths = pb.AllowedPaths
	st.AzureContainer = pb.AzureContainer
	st.AzureDnsZone = pb.AzureDnsZone
	st.AzureStorageAccount = pb.AzureStorageAccount
	st.AzureStorageService = pb.AzureStorageService
	st.BucketName = pb.BucketName
	st.Region = pb.Region
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *egressNetworkPolicyInternetAccessPolicyStorageDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st egressNetworkPolicyInternetAccessPolicyStorageDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func egressNetworkPolicyNetworkAccessPolicyToPb(st *EgressNetworkPolicyNetworkAccessPolicy) (*egressNetworkPolicyNetworkAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyPb{}
	pb.AllowedInternetDestinations = st.AllowedInternetDestinations
	pb.AllowedStorageDestinations = st.AllowedStorageDestinations
	pb.PolicyEnforcement = st.PolicyEnforcement
	pb.RestrictionMode = st.RestrictionMode

	return pb, nil
}

type egressNetworkPolicyNetworkAccessPolicyPb struct {
	AllowedInternetDestinations []EgressNetworkPolicyNetworkAccessPolicyInternetDestination `json:"allowed_internet_destinations,omitempty"`
	AllowedStorageDestinations  []EgressNetworkPolicyNetworkAccessPolicyStorageDestination  `json:"allowed_storage_destinations,omitempty"`
	PolicyEnforcement           *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement    `json:"policy_enforcement,omitempty"`
	RestrictionMode             EgressNetworkPolicyNetworkAccessPolicyRestrictionMode       `json:"restriction_mode"`
}

func egressNetworkPolicyNetworkAccessPolicyFromPb(pb *egressNetworkPolicyNetworkAccessPolicyPb) (*EgressNetworkPolicyNetworkAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicy{}
	st.AllowedInternetDestinations = pb.AllowedInternetDestinations
	st.AllowedStorageDestinations = pb.AllowedStorageDestinations
	st.PolicyEnforcement = pb.PolicyEnforcement
	st.RestrictionMode = pb.RestrictionMode

	return st, nil
}

func egressNetworkPolicyNetworkAccessPolicyInternetDestinationToPb(st *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) (*egressNetworkPolicyNetworkAccessPolicyInternetDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyInternetDestinationPb{}
	pb.Destination = st.Destination
	pb.InternetDestinationType = st.InternetDestinationType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type egressNetworkPolicyNetworkAccessPolicyInternetDestinationPb struct {
	Destination             string                                                                           `json:"destination,omitempty"`
	InternetDestinationType EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType `json:"internet_destination_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func egressNetworkPolicyNetworkAccessPolicyInternetDestinationFromPb(pb *egressNetworkPolicyNetworkAccessPolicyInternetDestinationPb) (*EgressNetworkPolicyNetworkAccessPolicyInternetDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicyInternetDestination{}
	st.Destination = pb.Destination
	st.InternetDestinationType = pb.InternetDestinationType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *egressNetworkPolicyNetworkAccessPolicyInternetDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st egressNetworkPolicyNetworkAccessPolicyInternetDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementToPb(st *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) (*egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb{}
	pb.DryRunModeProductFilter = st.DryRunModeProductFilter
	pb.EnforcementMode = st.EnforcementMode

	return pb, nil
}

type egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb struct {
	DryRunModeProductFilter []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter `json:"dry_run_mode_product_filter,omitempty"`
	EnforcementMode         EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode           `json:"enforcement_mode,omitempty"`
}

func egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementFromPb(pb *egressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb) (*EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement{}
	st.DryRunModeProductFilter = pb.DryRunModeProductFilter
	st.EnforcementMode = pb.EnforcementMode

	return st, nil
}

func egressNetworkPolicyNetworkAccessPolicyStorageDestinationToPb(st *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) (*egressNetworkPolicyNetworkAccessPolicyStorageDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyNetworkAccessPolicyStorageDestinationPb{}
	pb.AzureStorageAccount = st.AzureStorageAccount
	pb.AzureStorageService = st.AzureStorageService
	pb.BucketName = st.BucketName
	pb.Region = st.Region
	pb.StorageDestinationType = st.StorageDestinationType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type egressNetworkPolicyNetworkAccessPolicyStorageDestinationPb struct {
	AzureStorageAccount    string                                                                         `json:"azure_storage_account,omitempty"`
	AzureStorageService    string                                                                         `json:"azure_storage_service,omitempty"`
	BucketName             string                                                                         `json:"bucket_name,omitempty"`
	Region                 string                                                                         `json:"region,omitempty"`
	StorageDestinationType EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType `json:"storage_destination_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func egressNetworkPolicyNetworkAccessPolicyStorageDestinationFromPb(pb *egressNetworkPolicyNetworkAccessPolicyStorageDestinationPb) (*EgressNetworkPolicyNetworkAccessPolicyStorageDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicyStorageDestination{}
	st.AzureStorageAccount = pb.AzureStorageAccount
	st.AzureStorageService = pb.AzureStorageService
	st.BucketName = pb.BucketName
	st.Region = pb.Region
	st.StorageDestinationType = pb.StorageDestinationType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *egressNetworkPolicyNetworkAccessPolicyStorageDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st egressNetworkPolicyNetworkAccessPolicyStorageDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func emailConfigToPb(st *EmailConfig) (*emailConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &emailConfigPb{}
	pb.Addresses = st.Addresses

	return pb, nil
}

type emailConfigPb struct {
	Addresses []string `json:"addresses,omitempty"`
}

func emailConfigFromPb(pb *emailConfigPb) (*EmailConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmailConfig{}
	st.Addresses = pb.Addresses

	return st, nil
}

func enableExportNotebookToPb(st *EnableExportNotebook) (*enableExportNotebookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableExportNotebookPb{}
	pb.BooleanVal = st.BooleanVal
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enableExportNotebookPb struct {
	BooleanVal  *BooleanMessage `json:"boolean_val,omitempty"`
	SettingName string          `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enableExportNotebookFromPb(pb *enableExportNotebookPb) (*EnableExportNotebook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableExportNotebook{}
	st.BooleanVal = pb.BooleanVal
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enableExportNotebookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enableExportNotebookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enableNotebookTableClipboardToPb(st *EnableNotebookTableClipboard) (*enableNotebookTableClipboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableNotebookTableClipboardPb{}
	pb.BooleanVal = st.BooleanVal
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enableNotebookTableClipboardPb struct {
	BooleanVal  *BooleanMessage `json:"boolean_val,omitempty"`
	SettingName string          `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enableNotebookTableClipboardFromPb(pb *enableNotebookTableClipboardPb) (*EnableNotebookTableClipboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableNotebookTableClipboard{}
	st.BooleanVal = pb.BooleanVal
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enableNotebookTableClipboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enableNotebookTableClipboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enableResultsDownloadingToPb(st *EnableResultsDownloading) (*enableResultsDownloadingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableResultsDownloadingPb{}
	pb.BooleanVal = st.BooleanVal
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enableResultsDownloadingPb struct {
	BooleanVal  *BooleanMessage `json:"boolean_val,omitempty"`
	SettingName string          `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enableResultsDownloadingFromPb(pb *enableResultsDownloadingPb) (*EnableResultsDownloading, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableResultsDownloading{}
	st.BooleanVal = pb.BooleanVal
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enableResultsDownloadingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enableResultsDownloadingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enhancedSecurityMonitoringToPb(st *EnhancedSecurityMonitoring) (*enhancedSecurityMonitoringPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enhancedSecurityMonitoringPb{}
	pb.IsEnabled = st.IsEnabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enhancedSecurityMonitoringPb struct {
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enhancedSecurityMonitoringFromPb(pb *enhancedSecurityMonitoringPb) (*EnhancedSecurityMonitoring, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnhancedSecurityMonitoring{}
	st.IsEnabled = pb.IsEnabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enhancedSecurityMonitoringPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enhancedSecurityMonitoringPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func enhancedSecurityMonitoringSettingToPb(st *EnhancedSecurityMonitoringSetting) (*enhancedSecurityMonitoringSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enhancedSecurityMonitoringSettingPb{}
	pb.EnhancedSecurityMonitoringWorkspace = st.EnhancedSecurityMonitoringWorkspace
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type enhancedSecurityMonitoringSettingPb struct {
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoring `json:"enhanced_security_monitoring_workspace"`
	Etag                                string                     `json:"etag,omitempty"`
	SettingName                         string                     `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enhancedSecurityMonitoringSettingFromPb(pb *enhancedSecurityMonitoringSettingPb) (*EnhancedSecurityMonitoringSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnhancedSecurityMonitoringSetting{}
	st.EnhancedSecurityMonitoringWorkspace = pb.EnhancedSecurityMonitoringWorkspace
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enhancedSecurityMonitoringSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enhancedSecurityMonitoringSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func esmEnablementAccountToPb(st *EsmEnablementAccount) (*esmEnablementAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &esmEnablementAccountPb{}
	pb.IsEnforced = st.IsEnforced

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type esmEnablementAccountPb struct {
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func esmEnablementAccountFromPb(pb *esmEnablementAccountPb) (*EsmEnablementAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EsmEnablementAccount{}
	st.IsEnforced = pb.IsEnforced

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *esmEnablementAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st esmEnablementAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func esmEnablementAccountSettingToPb(st *EsmEnablementAccountSetting) (*esmEnablementAccountSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &esmEnablementAccountSettingPb{}
	pb.EsmEnablementAccount = st.EsmEnablementAccount
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type esmEnablementAccountSettingPb struct {
	EsmEnablementAccount EsmEnablementAccount `json:"esm_enablement_account"`
	Etag                 string               `json:"etag,omitempty"`
	SettingName          string               `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func esmEnablementAccountSettingFromPb(pb *esmEnablementAccountSettingPb) (*EsmEnablementAccountSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EsmEnablementAccountSetting{}
	st.EsmEnablementAccount = pb.EsmEnablementAccount
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *esmEnablementAccountSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st esmEnablementAccountSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func exchangeTokenToPb(st *ExchangeToken) (*exchangeTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeTokenPb{}
	pb.Credential = st.Credential
	pb.CredentialEolTime = st.CredentialEolTime
	pb.OwnerId = st.OwnerId
	pb.Scopes = st.Scopes
	pb.TokenType = st.TokenType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type exchangeTokenPb struct {
	Credential        string    `json:"credential,omitempty"`
	CredentialEolTime int64     `json:"credentialEolTime,omitempty"`
	OwnerId           int64     `json:"ownerId,omitempty"`
	Scopes            []string  `json:"scopes,omitempty"`
	TokenType         TokenType `json:"tokenType,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func exchangeTokenFromPb(pb *exchangeTokenPb) (*ExchangeToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeToken{}
	st.Credential = pb.Credential
	st.CredentialEolTime = pb.CredentialEolTime
	st.OwnerId = pb.OwnerId
	st.Scopes = pb.Scopes
	st.TokenType = pb.TokenType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *exchangeTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st exchangeTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func exchangeTokenRequestToPb(st *ExchangeTokenRequest) (*exchangeTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeTokenRequestPb{}
	pb.PartitionId = st.PartitionId
	pb.Scopes = st.Scopes
	pb.TokenType = st.TokenType

	return pb, nil
}

type exchangeTokenRequestPb struct {
	PartitionId PartitionId `json:"partitionId"`
	Scopes      []string    `json:"scopes"`
	TokenType   []TokenType `json:"tokenType"`
}

func exchangeTokenRequestFromPb(pb *exchangeTokenRequestPb) (*ExchangeTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeTokenRequest{}
	st.PartitionId = pb.PartitionId
	st.Scopes = pb.Scopes
	st.TokenType = pb.TokenType

	return st, nil
}

func exchangeTokenResponseToPb(st *ExchangeTokenResponse) (*exchangeTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeTokenResponsePb{}
	pb.Values = st.Values

	return pb, nil
}

type exchangeTokenResponsePb struct {
	Values []ExchangeToken `json:"values,omitempty"`
}

func exchangeTokenResponseFromPb(pb *exchangeTokenResponsePb) (*ExchangeTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeTokenResponse{}
	st.Values = pb.Values

	return st, nil
}

func fetchIpAccessListResponseToPb(st *FetchIpAccessListResponse) (*fetchIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fetchIpAccessListResponsePb{}
	pb.IpAccessList = st.IpAccessList

	return pb, nil
}

type fetchIpAccessListResponsePb struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func fetchIpAccessListResponseFromPb(pb *fetchIpAccessListResponsePb) (*FetchIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FetchIpAccessListResponse{}
	st.IpAccessList = pb.IpAccessList

	return st, nil
}

func genericWebhookConfigToPb(st *GenericWebhookConfig) (*genericWebhookConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genericWebhookConfigPb{}
	pb.Password = st.Password
	pb.PasswordSet = st.PasswordSet
	pb.Url = st.Url
	pb.UrlSet = st.UrlSet
	pb.Username = st.Username
	pb.UsernameSet = st.UsernameSet

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type genericWebhookConfigPb struct {
	Password    string `json:"password,omitempty"`
	PasswordSet bool   `json:"password_set,omitempty"`
	Url         string `json:"url,omitempty"`
	UrlSet      bool   `json:"url_set,omitempty"`
	Username    string `json:"username,omitempty"`
	UsernameSet bool   `json:"username_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func genericWebhookConfigFromPb(pb *genericWebhookConfigPb) (*GenericWebhookConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenericWebhookConfig{}
	st.Password = pb.Password
	st.PasswordSet = pb.PasswordSet
	st.Url = pb.Url
	st.UrlSet = pb.UrlSet
	st.Username = pb.Username
	st.UsernameSet = pb.UsernameSet

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genericWebhookConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genericWebhookConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getAccountIpAccessEnableRequestToPb(st *GetAccountIpAccessEnableRequest) (*getAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountIpAccessEnableRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getAccountIpAccessEnableRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getAccountIpAccessEnableRequestFromPb(pb *getAccountIpAccessEnableRequestPb) (*GetAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountIpAccessEnableRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAccountIpAccessEnableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAccountIpAccessEnableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getAccountIpAccessListRequestToPb(st *GetAccountIpAccessListRequest) (*getAccountIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

type getAccountIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

func getAccountIpAccessListRequestFromPb(pb *getAccountIpAccessListRequestPb) (*GetAccountIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

func getAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*getAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *getAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*GetAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAibiDashboardEmbeddingAccessPolicySettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAibiDashboardEmbeddingAccessPolicySettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAibiDashboardEmbeddingAccessPolicySettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*GetAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getAutomaticClusterUpdateSettingRequestToPb(st *GetAutomaticClusterUpdateSettingRequest) (*getAutomaticClusterUpdateSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAutomaticClusterUpdateSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getAutomaticClusterUpdateSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getAutomaticClusterUpdateSettingRequestFromPb(pb *getAutomaticClusterUpdateSettingRequestPb) (*GetAutomaticClusterUpdateSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAutomaticClusterUpdateSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAutomaticClusterUpdateSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAutomaticClusterUpdateSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getComplianceSecurityProfileSettingRequestToPb(st *GetComplianceSecurityProfileSettingRequest) (*getComplianceSecurityProfileSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getComplianceSecurityProfileSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getComplianceSecurityProfileSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getComplianceSecurityProfileSettingRequestFromPb(pb *getComplianceSecurityProfileSettingRequestPb) (*GetComplianceSecurityProfileSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetComplianceSecurityProfileSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getComplianceSecurityProfileSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getComplianceSecurityProfileSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getCspEnablementAccountSettingRequestToPb(st *GetCspEnablementAccountSettingRequest) (*getCspEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCspEnablementAccountSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getCspEnablementAccountSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCspEnablementAccountSettingRequestFromPb(pb *getCspEnablementAccountSettingRequestPb) (*GetCspEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCspEnablementAccountSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCspEnablementAccountSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCspEnablementAccountSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getDashboardEmailSubscriptionsRequestToPb(st *GetDashboardEmailSubscriptionsRequest) (*getDashboardEmailSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDashboardEmailSubscriptionsRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getDashboardEmailSubscriptionsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getDashboardEmailSubscriptionsRequestFromPb(pb *getDashboardEmailSubscriptionsRequestPb) (*GetDashboardEmailSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDashboardEmailSubscriptionsRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDashboardEmailSubscriptionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDashboardEmailSubscriptionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getDefaultNamespaceSettingRequestToPb(st *GetDefaultNamespaceSettingRequest) (*getDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDefaultNamespaceSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getDefaultNamespaceSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getDefaultNamespaceSettingRequestFromPb(pb *getDefaultNamespaceSettingRequestPb) (*GetDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDefaultNamespaceSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDefaultNamespaceSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDefaultNamespaceSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getDisableLegacyAccessRequestToPb(st *GetDisableLegacyAccessRequest) (*getDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDisableLegacyAccessRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getDisableLegacyAccessRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getDisableLegacyAccessRequestFromPb(pb *getDisableLegacyAccessRequestPb) (*GetDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyAccessRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDisableLegacyAccessRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDisableLegacyAccessRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getDisableLegacyDbfsRequestToPb(st *GetDisableLegacyDbfsRequest) (*getDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDisableLegacyDbfsRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getDisableLegacyDbfsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getDisableLegacyDbfsRequestFromPb(pb *getDisableLegacyDbfsRequestPb) (*GetDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyDbfsRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDisableLegacyDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDisableLegacyDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getDisableLegacyFeaturesRequestToPb(st *GetDisableLegacyFeaturesRequest) (*getDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDisableLegacyFeaturesRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getDisableLegacyFeaturesRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getDisableLegacyFeaturesRequestFromPb(pb *getDisableLegacyFeaturesRequestPb) (*GetDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyFeaturesRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDisableLegacyFeaturesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDisableLegacyFeaturesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getEnhancedSecurityMonitoringSettingRequestToPb(st *GetEnhancedSecurityMonitoringSettingRequest) (*getEnhancedSecurityMonitoringSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEnhancedSecurityMonitoringSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getEnhancedSecurityMonitoringSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getEnhancedSecurityMonitoringSettingRequestFromPb(pb *getEnhancedSecurityMonitoringSettingRequestPb) (*GetEnhancedSecurityMonitoringSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEnhancedSecurityMonitoringSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEnhancedSecurityMonitoringSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEnhancedSecurityMonitoringSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getEsmEnablementAccountSettingRequestToPb(st *GetEsmEnablementAccountSettingRequest) (*getEsmEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEsmEnablementAccountSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getEsmEnablementAccountSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getEsmEnablementAccountSettingRequestFromPb(pb *getEsmEnablementAccountSettingRequestPb) (*GetEsmEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEsmEnablementAccountSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEsmEnablementAccountSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEsmEnablementAccountSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getIpAccessListRequestToPb(st *GetIpAccessListRequest) (*getIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

type getIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

func getIpAccessListRequestFromPb(pb *getIpAccessListRequestPb) (*GetIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

func getIpAccessListResponseToPb(st *GetIpAccessListResponse) (*getIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIpAccessListResponsePb{}
	pb.IpAccessList = st.IpAccessList

	return pb, nil
}

type getIpAccessListResponsePb struct {
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func getIpAccessListResponseFromPb(pb *getIpAccessListResponsePb) (*GetIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListResponse{}
	st.IpAccessList = pb.IpAccessList

	return st, nil
}

func getIpAccessListsResponseToPb(st *GetIpAccessListsResponse) (*getIpAccessListsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIpAccessListsResponsePb{}
	pb.IpAccessLists = st.IpAccessLists

	return pb, nil
}

type getIpAccessListsResponsePb struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func getIpAccessListsResponseFromPb(pb *getIpAccessListsResponsePb) (*GetIpAccessListsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListsResponse{}
	st.IpAccessLists = pb.IpAccessLists

	return st, nil
}

func getLlmProxyPartnerPoweredAccountRequestToPb(st *GetLlmProxyPartnerPoweredAccountRequest) (*getLlmProxyPartnerPoweredAccountRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLlmProxyPartnerPoweredAccountRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getLlmProxyPartnerPoweredAccountRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getLlmProxyPartnerPoweredAccountRequestFromPb(pb *getLlmProxyPartnerPoweredAccountRequestPb) (*GetLlmProxyPartnerPoweredAccountRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLlmProxyPartnerPoweredAccountRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getLlmProxyPartnerPoweredAccountRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getLlmProxyPartnerPoweredAccountRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getLlmProxyPartnerPoweredEnforceRequestToPb(st *GetLlmProxyPartnerPoweredEnforceRequest) (*getLlmProxyPartnerPoweredEnforceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLlmProxyPartnerPoweredEnforceRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getLlmProxyPartnerPoweredEnforceRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getLlmProxyPartnerPoweredEnforceRequestFromPb(pb *getLlmProxyPartnerPoweredEnforceRequestPb) (*GetLlmProxyPartnerPoweredEnforceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLlmProxyPartnerPoweredEnforceRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getLlmProxyPartnerPoweredEnforceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getLlmProxyPartnerPoweredEnforceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getLlmProxyPartnerPoweredWorkspaceRequestToPb(st *GetLlmProxyPartnerPoweredWorkspaceRequest) (*getLlmProxyPartnerPoweredWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLlmProxyPartnerPoweredWorkspaceRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getLlmProxyPartnerPoweredWorkspaceRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb *getLlmProxyPartnerPoweredWorkspaceRequestPb) (*GetLlmProxyPartnerPoweredWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLlmProxyPartnerPoweredWorkspaceRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getLlmProxyPartnerPoweredWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getLlmProxyPartnerPoweredWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getNetworkConnectivityConfigurationRequestToPb(st *GetNetworkConnectivityConfigurationRequest) (*getNetworkConnectivityConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNetworkConnectivityConfigurationRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId

	return pb, nil
}

type getNetworkConnectivityConfigurationRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

func getNetworkConnectivityConfigurationRequestFromPb(pb *getNetworkConnectivityConfigurationRequestPb) (*GetNetworkConnectivityConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkConnectivityConfigurationRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId

	return st, nil
}

func getNetworkPolicyRequestToPb(st *GetNetworkPolicyRequest) (*getNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNetworkPolicyRequestPb{}
	pb.NetworkPolicyId = st.NetworkPolicyId

	return pb, nil
}

type getNetworkPolicyRequestPb struct {
	NetworkPolicyId string `json:"-" url:"-"`
}

func getNetworkPolicyRequestFromPb(pb *getNetworkPolicyRequestPb) (*GetNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkPolicyRequest{}
	st.NetworkPolicyId = pb.NetworkPolicyId

	return st, nil
}

func getNotificationDestinationRequestToPb(st *GetNotificationDestinationRequest) (*getNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNotificationDestinationRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getNotificationDestinationRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getNotificationDestinationRequestFromPb(pb *getNotificationDestinationRequestPb) (*GetNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNotificationDestinationRequest{}
	st.Id = pb.Id

	return st, nil
}

func getPersonalComputeSettingRequestToPb(st *GetPersonalComputeSettingRequest) (*getPersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPersonalComputeSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPersonalComputeSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPersonalComputeSettingRequestFromPb(pb *getPersonalComputeSettingRequestPb) (*GetPersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalComputeSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPersonalComputeSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPersonalComputeSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getPrivateEndpointRuleRequestToPb(st *GetPrivateEndpointRuleRequest) (*getPrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PrivateEndpointRuleId = st.PrivateEndpointRuleId

	return pb, nil
}

type getPrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	PrivateEndpointRuleId       string `json:"-" url:"-"`
}

func getPrivateEndpointRuleRequestFromPb(pb *getPrivateEndpointRuleRequestPb) (*GetPrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PrivateEndpointRuleId = pb.PrivateEndpointRuleId

	return st, nil
}

func getRestrictWorkspaceAdminsSettingRequestToPb(st *GetRestrictWorkspaceAdminsSettingRequest) (*getRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRestrictWorkspaceAdminsSettingRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getRestrictWorkspaceAdminsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRestrictWorkspaceAdminsSettingRequestFromPb(pb *getRestrictWorkspaceAdminsSettingRequestPb) (*GetRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRestrictWorkspaceAdminsSettingRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRestrictWorkspaceAdminsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRestrictWorkspaceAdminsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getSqlResultsDownloadRequestToPb(st *GetSqlResultsDownloadRequest) (*getSqlResultsDownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getSqlResultsDownloadRequestPb{}
	pb.Etag = st.Etag

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getSqlResultsDownloadRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getSqlResultsDownloadRequestFromPb(pb *getSqlResultsDownloadRequestPb) (*GetSqlResultsDownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSqlResultsDownloadRequest{}
	st.Etag = pb.Etag

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getSqlResultsDownloadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getSqlResultsDownloadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getStatusRequestToPb(st *GetStatusRequest) (*getStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatusRequestPb{}
	pb.Keys = st.Keys

	return pb, nil
}

type getStatusRequestPb struct {
	Keys string `json:"-" url:"keys"`
}

func getStatusRequestFromPb(pb *getStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	st.Keys = pb.Keys

	return st, nil
}

func getTokenManagementRequestToPb(st *GetTokenManagementRequest) (*getTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTokenManagementRequestPb{}
	pb.TokenId = st.TokenId

	return pb, nil
}

type getTokenManagementRequestPb struct {
	TokenId string `json:"-" url:"-"`
}

func getTokenManagementRequestFromPb(pb *getTokenManagementRequestPb) (*GetTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenManagementRequest{}
	st.TokenId = pb.TokenId

	return st, nil
}

func getTokenPermissionLevelsResponseToPb(st *GetTokenPermissionLevelsResponse) (*getTokenPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTokenPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getTokenPermissionLevelsResponsePb struct {
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

func getTokenPermissionLevelsResponseFromPb(pb *getTokenPermissionLevelsResponsePb) (*GetTokenPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getTokenResponseToPb(st *GetTokenResponse) (*getTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTokenResponsePb{}
	pb.TokenInfo = st.TokenInfo

	return pb, nil
}

type getTokenResponsePb struct {
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
}

func getTokenResponseFromPb(pb *getTokenResponsePb) (*GetTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenResponse{}
	st.TokenInfo = pb.TokenInfo

	return st, nil
}

func getWorkspaceNetworkOptionRequestToPb(st *GetWorkspaceNetworkOptionRequest) (*getWorkspaceNetworkOptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getWorkspaceNetworkOptionRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

type getWorkspaceNetworkOptionRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

func getWorkspaceNetworkOptionRequestFromPb(pb *getWorkspaceNetworkOptionRequestPb) (*GetWorkspaceNetworkOptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceNetworkOptionRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
}

func ipAccessListInfoToPb(st *IpAccessListInfo) (*ipAccessListInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ipAccessListInfoPb{}
	pb.AddressCount = st.AddressCount
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Enabled = st.Enabled
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	pb.ListId = st.ListId
	pb.ListType = st.ListType
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type ipAccessListInfoPb struct {
	AddressCount int      `json:"address_count,omitempty"`
	CreatedAt    int64    `json:"created_at,omitempty"`
	CreatedBy    int64    `json:"created_by,omitempty"`
	Enabled      bool     `json:"enabled,omitempty"`
	IpAddresses  []string `json:"ip_addresses,omitempty"`
	Label        string   `json:"label,omitempty"`
	ListId       string   `json:"list_id,omitempty"`
	ListType     ListType `json:"list_type,omitempty"`
	UpdatedAt    int64    `json:"updated_at,omitempty"`
	UpdatedBy    int64    `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ipAccessListInfoFromPb(pb *ipAccessListInfoPb) (*IpAccessListInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IpAccessListInfo{}
	st.AddressCount = pb.AddressCount
	st.CreatedAt = pb.CreatedAt
	st.CreatedBy = pb.CreatedBy
	st.Enabled = pb.Enabled
	st.IpAddresses = pb.IpAddresses
	st.Label = pb.Label
	st.ListId = pb.ListId
	st.ListType = pb.ListType
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ipAccessListInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ipAccessListInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listIpAccessListResponseToPb(st *ListIpAccessListResponse) (*listIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listIpAccessListResponsePb{}
	pb.IpAccessLists = st.IpAccessLists

	return pb, nil
}

type listIpAccessListResponsePb struct {
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func listIpAccessListResponseFromPb(pb *listIpAccessListResponsePb) (*ListIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListIpAccessListResponse{}
	st.IpAccessLists = pb.IpAccessLists

	return st, nil
}

func listNetworkConnectivityConfigurationsRequestToPb(st *ListNetworkConnectivityConfigurationsRequest) (*listNetworkConnectivityConfigurationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNetworkConnectivityConfigurationsRequestPb{}
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listNetworkConnectivityConfigurationsRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNetworkConnectivityConfigurationsRequestFromPb(pb *listNetworkConnectivityConfigurationsRequestPb) (*ListNetworkConnectivityConfigurationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkConnectivityConfigurationsRequest{}
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNetworkConnectivityConfigurationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNetworkConnectivityConfigurationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listNetworkConnectivityConfigurationsResponseToPb(st *ListNetworkConnectivityConfigurationsResponse) (*listNetworkConnectivityConfigurationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNetworkConnectivityConfigurationsResponsePb{}
	pb.Items = st.Items
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listNetworkConnectivityConfigurationsResponsePb struct {
	Items         []NetworkConnectivityConfiguration `json:"items,omitempty"`
	NextPageToken string                             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNetworkConnectivityConfigurationsResponseFromPb(pb *listNetworkConnectivityConfigurationsResponsePb) (*ListNetworkConnectivityConfigurationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkConnectivityConfigurationsResponse{}
	st.Items = pb.Items
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNetworkConnectivityConfigurationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNetworkConnectivityConfigurationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listNetworkPoliciesRequestToPb(st *ListNetworkPoliciesRequest) (*listNetworkPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNetworkPoliciesRequestPb{}
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listNetworkPoliciesRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNetworkPoliciesRequestFromPb(pb *listNetworkPoliciesRequestPb) (*ListNetworkPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkPoliciesRequest{}
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNetworkPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNetworkPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listNetworkPoliciesResponseToPb(st *ListNetworkPoliciesResponse) (*listNetworkPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNetworkPoliciesResponsePb{}
	pb.Items = st.Items
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listNetworkPoliciesResponsePb struct {
	Items         []AccountNetworkPolicy `json:"items,omitempty"`
	NextPageToken string                 `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNetworkPoliciesResponseFromPb(pb *listNetworkPoliciesResponsePb) (*ListNetworkPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkPoliciesResponse{}
	st.Items = pb.Items
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNetworkPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNetworkPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listNotificationDestinationsRequestToPb(st *ListNotificationDestinationsRequest) (*listNotificationDestinationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNotificationDestinationsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listNotificationDestinationsRequestPb struct {
	PageSize  int64  `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNotificationDestinationsRequestFromPb(pb *listNotificationDestinationsRequestPb) (*ListNotificationDestinationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNotificationDestinationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNotificationDestinationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listNotificationDestinationsResponseToPb(st *ListNotificationDestinationsResponse) (*listNotificationDestinationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNotificationDestinationsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Results = st.Results

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listNotificationDestinationsResponsePb struct {
	NextPageToken string                               `json:"next_page_token,omitempty"`
	Results       []ListNotificationDestinationsResult `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNotificationDestinationsResponseFromPb(pb *listNotificationDestinationsResponsePb) (*ListNotificationDestinationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Results = pb.Results

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNotificationDestinationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNotificationDestinationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listNotificationDestinationsResultToPb(st *ListNotificationDestinationsResult) (*listNotificationDestinationsResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNotificationDestinationsResultPb{}
	pb.DestinationType = st.DestinationType
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listNotificationDestinationsResultPb struct {
	DestinationType DestinationType `json:"destination_type,omitempty"`
	DisplayName     string          `json:"display_name,omitempty"`
	Id              string          `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNotificationDestinationsResultFromPb(pb *listNotificationDestinationsResultPb) (*ListNotificationDestinationsResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsResult{}
	st.DestinationType = pb.DestinationType
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNotificationDestinationsResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNotificationDestinationsResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPrivateEndpointRulesRequestToPb(st *ListPrivateEndpointRulesRequest) (*listPrivateEndpointRulesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPrivateEndpointRulesRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPrivateEndpointRulesRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	PageToken                   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPrivateEndpointRulesRequestFromPb(pb *listPrivateEndpointRulesRequestPb) (*ListPrivateEndpointRulesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPrivateEndpointRulesRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPrivateEndpointRulesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPrivateEndpointRulesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPrivateEndpointRulesResponseToPb(st *ListPrivateEndpointRulesResponse) (*listPrivateEndpointRulesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPrivateEndpointRulesResponsePb{}
	pb.Items = st.Items
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPrivateEndpointRulesResponsePb struct {
	Items         []NccPrivateEndpointRule `json:"items,omitempty"`
	NextPageToken string                   `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPrivateEndpointRulesResponseFromPb(pb *listPrivateEndpointRulesResponsePb) (*ListPrivateEndpointRulesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPrivateEndpointRulesResponse{}
	st.Items = pb.Items
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPrivateEndpointRulesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPrivateEndpointRulesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPublicTokensResponseToPb(st *ListPublicTokensResponse) (*listPublicTokensResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPublicTokensResponsePb{}
	pb.TokenInfos = st.TokenInfos

	return pb, nil
}

type listPublicTokensResponsePb struct {
	TokenInfos []PublicTokenInfo `json:"token_infos,omitempty"`
}

func listPublicTokensResponseFromPb(pb *listPublicTokensResponsePb) (*ListPublicTokensResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPublicTokensResponse{}
	st.TokenInfos = pb.TokenInfos

	return st, nil
}

func listTokenManagementRequestToPb(st *ListTokenManagementRequest) (*listTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTokenManagementRequestPb{}
	pb.CreatedById = st.CreatedById
	pb.CreatedByUsername = st.CreatedByUsername

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listTokenManagementRequestPb struct {
	CreatedById       int64  `json:"-" url:"created_by_id,omitempty"`
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTokenManagementRequestFromPb(pb *listTokenManagementRequestPb) (*ListTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTokenManagementRequest{}
	st.CreatedById = pb.CreatedById
	st.CreatedByUsername = pb.CreatedByUsername

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTokenManagementRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTokenManagementRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listTokensResponseToPb(st *ListTokensResponse) (*listTokensResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTokensResponsePb{}
	pb.TokenInfos = st.TokenInfos

	return pb, nil
}

type listTokensResponsePb struct {
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

func listTokensResponseFromPb(pb *listTokensResponsePb) (*ListTokensResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTokensResponse{}
	st.TokenInfos = pb.TokenInfos

	return st, nil
}

func llmProxyPartnerPoweredAccountToPb(st *LlmProxyPartnerPoweredAccount) (*llmProxyPartnerPoweredAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &llmProxyPartnerPoweredAccountPb{}
	pb.BooleanVal = st.BooleanVal
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type llmProxyPartnerPoweredAccountPb struct {
	BooleanVal  BooleanMessage `json:"boolean_val"`
	Etag        string         `json:"etag,omitempty"`
	SettingName string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func llmProxyPartnerPoweredAccountFromPb(pb *llmProxyPartnerPoweredAccountPb) (*LlmProxyPartnerPoweredAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LlmProxyPartnerPoweredAccount{}
	st.BooleanVal = pb.BooleanVal
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *llmProxyPartnerPoweredAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st llmProxyPartnerPoweredAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func llmProxyPartnerPoweredEnforceToPb(st *LlmProxyPartnerPoweredEnforce) (*llmProxyPartnerPoweredEnforcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &llmProxyPartnerPoweredEnforcePb{}
	pb.BooleanVal = st.BooleanVal
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type llmProxyPartnerPoweredEnforcePb struct {
	BooleanVal  BooleanMessage `json:"boolean_val"`
	Etag        string         `json:"etag,omitempty"`
	SettingName string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func llmProxyPartnerPoweredEnforceFromPb(pb *llmProxyPartnerPoweredEnforcePb) (*LlmProxyPartnerPoweredEnforce, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LlmProxyPartnerPoweredEnforce{}
	st.BooleanVal = pb.BooleanVal
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *llmProxyPartnerPoweredEnforcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st llmProxyPartnerPoweredEnforcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func llmProxyPartnerPoweredWorkspaceToPb(st *LlmProxyPartnerPoweredWorkspace) (*llmProxyPartnerPoweredWorkspacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &llmProxyPartnerPoweredWorkspacePb{}
	pb.BooleanVal = st.BooleanVal
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type llmProxyPartnerPoweredWorkspacePb struct {
	BooleanVal  BooleanMessage `json:"boolean_val"`
	Etag        string         `json:"etag,omitempty"`
	SettingName string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func llmProxyPartnerPoweredWorkspaceFromPb(pb *llmProxyPartnerPoweredWorkspacePb) (*LlmProxyPartnerPoweredWorkspace, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LlmProxyPartnerPoweredWorkspace{}
	st.BooleanVal = pb.BooleanVal
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *llmProxyPartnerPoweredWorkspacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st llmProxyPartnerPoweredWorkspacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func microsoftTeamsConfigToPb(st *MicrosoftTeamsConfig) (*microsoftTeamsConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &microsoftTeamsConfigPb{}
	pb.Url = st.Url
	pb.UrlSet = st.UrlSet

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type microsoftTeamsConfigPb struct {
	Url    string `json:"url,omitempty"`
	UrlSet bool   `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func microsoftTeamsConfigFromPb(pb *microsoftTeamsConfigPb) (*MicrosoftTeamsConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MicrosoftTeamsConfig{}
	st.Url = pb.Url
	st.UrlSet = pb.UrlSet

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *microsoftTeamsConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st microsoftTeamsConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func nccAwsStableIpRuleToPb(st *NccAwsStableIpRule) (*nccAwsStableIpRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccAwsStableIpRulePb{}
	pb.CidrBlocks = st.CidrBlocks

	return pb, nil
}

type nccAwsStableIpRulePb struct {
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

func nccAwsStableIpRuleFromPb(pb *nccAwsStableIpRulePb) (*NccAwsStableIpRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAwsStableIpRule{}
	st.CidrBlocks = pb.CidrBlocks

	return st, nil
}

func nccAzurePrivateEndpointRuleToPb(st *NccAzurePrivateEndpointRule) (*nccAzurePrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccAzurePrivateEndpointRulePb{}
	pb.ConnectionState = st.ConnectionState
	pb.CreationTime = st.CreationTime
	pb.Deactivated = st.Deactivated
	pb.DeactivatedAt = st.DeactivatedAt
	pb.DomainNames = st.DomainNames
	pb.EndpointName = st.EndpointName
	pb.GroupId = st.GroupId
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.ResourceId = st.ResourceId
	pb.RuleId = st.RuleId
	pb.UpdatedTime = st.UpdatedTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type nccAzurePrivateEndpointRulePb struct {
	ConnectionState             NccAzurePrivateEndpointRuleConnectionState `json:"connection_state,omitempty"`
	CreationTime                int64                                      `json:"creation_time,omitempty"`
	Deactivated                 bool                                       `json:"deactivated,omitempty"`
	DeactivatedAt               int64                                      `json:"deactivated_at,omitempty"`
	DomainNames                 []string                                   `json:"domain_names,omitempty"`
	EndpointName                string                                     `json:"endpoint_name,omitempty"`
	GroupId                     string                                     `json:"group_id,omitempty"`
	NetworkConnectivityConfigId string                                     `json:"network_connectivity_config_id,omitempty"`
	ResourceId                  string                                     `json:"resource_id,omitempty"`
	RuleId                      string                                     `json:"rule_id,omitempty"`
	UpdatedTime                 int64                                      `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func nccAzurePrivateEndpointRuleFromPb(pb *nccAzurePrivateEndpointRulePb) (*NccAzurePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAzurePrivateEndpointRule{}
	st.ConnectionState = pb.ConnectionState
	st.CreationTime = pb.CreationTime
	st.Deactivated = pb.Deactivated
	st.DeactivatedAt = pb.DeactivatedAt
	st.DomainNames = pb.DomainNames
	st.EndpointName = pb.EndpointName
	st.GroupId = pb.GroupId
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.ResourceId = pb.ResourceId
	st.RuleId = pb.RuleId
	st.UpdatedTime = pb.UpdatedTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nccAzurePrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nccAzurePrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func nccAzureServiceEndpointRuleToPb(st *NccAzureServiceEndpointRule) (*nccAzureServiceEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccAzureServiceEndpointRulePb{}
	pb.Subnets = st.Subnets
	pb.TargetRegion = st.TargetRegion
	pb.TargetServices = st.TargetServices

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type nccAzureServiceEndpointRulePb struct {
	Subnets        []string             `json:"subnets,omitempty"`
	TargetRegion   string               `json:"target_region,omitempty"`
	TargetServices []EgressResourceType `json:"target_services,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func nccAzureServiceEndpointRuleFromPb(pb *nccAzureServiceEndpointRulePb) (*NccAzureServiceEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAzureServiceEndpointRule{}
	st.Subnets = pb.Subnets
	st.TargetRegion = pb.TargetRegion
	st.TargetServices = pb.TargetServices

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nccAzureServiceEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nccAzureServiceEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func nccEgressConfigToPb(st *NccEgressConfig) (*nccEgressConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccEgressConfigPb{}
	pb.DefaultRules = st.DefaultRules
	pb.TargetRules = st.TargetRules

	return pb, nil
}

type nccEgressConfigPb struct {
	DefaultRules *NccEgressDefaultRules `json:"default_rules,omitempty"`
	TargetRules  *NccEgressTargetRules  `json:"target_rules,omitempty"`
}

func nccEgressConfigFromPb(pb *nccEgressConfigPb) (*NccEgressConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressConfig{}
	st.DefaultRules = pb.DefaultRules
	st.TargetRules = pb.TargetRules

	return st, nil
}

func nccEgressDefaultRulesToPb(st *NccEgressDefaultRules) (*nccEgressDefaultRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccEgressDefaultRulesPb{}
	pb.AwsStableIpRule = st.AwsStableIpRule
	pb.AzureServiceEndpointRule = st.AzureServiceEndpointRule

	return pb, nil
}

type nccEgressDefaultRulesPb struct {
	AwsStableIpRule          *NccAwsStableIpRule          `json:"aws_stable_ip_rule,omitempty"`
	AzureServiceEndpointRule *NccAzureServiceEndpointRule `json:"azure_service_endpoint_rule,omitempty"`
}

func nccEgressDefaultRulesFromPb(pb *nccEgressDefaultRulesPb) (*NccEgressDefaultRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressDefaultRules{}
	st.AwsStableIpRule = pb.AwsStableIpRule
	st.AzureServiceEndpointRule = pb.AzureServiceEndpointRule

	return st, nil
}

func nccEgressTargetRulesToPb(st *NccEgressTargetRules) (*nccEgressTargetRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccEgressTargetRulesPb{}
	pb.AwsPrivateEndpointRules = st.AwsPrivateEndpointRules
	pb.AzurePrivateEndpointRules = st.AzurePrivateEndpointRules

	return pb, nil
}

type nccEgressTargetRulesPb struct {
	AwsPrivateEndpointRules   []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule `json:"aws_private_endpoint_rules,omitempty"`
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule                                   `json:"azure_private_endpoint_rules,omitempty"`
}

func nccEgressTargetRulesFromPb(pb *nccEgressTargetRulesPb) (*NccEgressTargetRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressTargetRules{}
	st.AwsPrivateEndpointRules = pb.AwsPrivateEndpointRules
	st.AzurePrivateEndpointRules = pb.AzurePrivateEndpointRules

	return st, nil
}

func nccPrivateEndpointRuleToPb(st *NccPrivateEndpointRule) (*nccPrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccPrivateEndpointRulePb{}
	pb.AccountId = st.AccountId
	pb.ConnectionState = st.ConnectionState
	pb.CreationTime = st.CreationTime
	pb.Deactivated = st.Deactivated
	pb.DeactivatedAt = st.DeactivatedAt
	pb.DomainNames = st.DomainNames
	pb.Enabled = st.Enabled
	pb.EndpointName = st.EndpointName
	pb.EndpointService = st.EndpointService
	pb.GroupId = st.GroupId
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.ResourceId = st.ResourceId
	pb.ResourceNames = st.ResourceNames
	pb.RuleId = st.RuleId
	pb.UpdatedTime = st.UpdatedTime
	pb.VpcEndpointId = st.VpcEndpointId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type nccPrivateEndpointRulePb struct {
	AccountId                   string                                           `json:"account_id,omitempty"`
	ConnectionState             NccPrivateEndpointRulePrivateLinkConnectionState `json:"connection_state,omitempty"`
	CreationTime                int64                                            `json:"creation_time,omitempty"`
	Deactivated                 bool                                             `json:"deactivated,omitempty"`
	DeactivatedAt               int64                                            `json:"deactivated_at,omitempty"`
	DomainNames                 []string                                         `json:"domain_names,omitempty"`
	Enabled                     bool                                             `json:"enabled,omitempty"`
	EndpointName                string                                           `json:"endpoint_name,omitempty"`
	EndpointService             string                                           `json:"endpoint_service,omitempty"`
	GroupId                     string                                           `json:"group_id,omitempty"`
	NetworkConnectivityConfigId string                                           `json:"network_connectivity_config_id,omitempty"`
	ResourceId                  string                                           `json:"resource_id,omitempty"`
	ResourceNames               []string                                         `json:"resource_names,omitempty"`
	RuleId                      string                                           `json:"rule_id,omitempty"`
	UpdatedTime                 int64                                            `json:"updated_time,omitempty"`
	VpcEndpointId               string                                           `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func nccPrivateEndpointRuleFromPb(pb *nccPrivateEndpointRulePb) (*NccPrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccPrivateEndpointRule{}
	st.AccountId = pb.AccountId
	st.ConnectionState = pb.ConnectionState
	st.CreationTime = pb.CreationTime
	st.Deactivated = pb.Deactivated
	st.DeactivatedAt = pb.DeactivatedAt
	st.DomainNames = pb.DomainNames
	st.Enabled = pb.Enabled
	st.EndpointName = pb.EndpointName
	st.EndpointService = pb.EndpointService
	st.GroupId = pb.GroupId
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.ResourceId = pb.ResourceId
	st.ResourceNames = pb.ResourceNames
	st.RuleId = pb.RuleId
	st.UpdatedTime = pb.UpdatedTime
	st.VpcEndpointId = pb.VpcEndpointId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nccPrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nccPrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func networkConnectivityConfigurationToPb(st *NetworkConnectivityConfiguration) (*networkConnectivityConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkConnectivityConfigurationPb{}
	pb.AccountId = st.AccountId
	pb.CreationTime = st.CreationTime
	pb.EgressConfig = st.EgressConfig
	pb.Name = st.Name
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.Region = st.Region
	pb.UpdatedTime = st.UpdatedTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type networkConnectivityConfigurationPb struct {
	AccountId                   string           `json:"account_id,omitempty"`
	CreationTime                int64            `json:"creation_time,omitempty"`
	EgressConfig                *NccEgressConfig `json:"egress_config,omitempty"`
	Name                        string           `json:"name,omitempty"`
	NetworkConnectivityConfigId string           `json:"network_connectivity_config_id,omitempty"`
	Region                      string           `json:"region,omitempty"`
	UpdatedTime                 int64            `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func networkConnectivityConfigurationFromPb(pb *networkConnectivityConfigurationPb) (*NetworkConnectivityConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkConnectivityConfiguration{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime
	st.EgressConfig = pb.EgressConfig
	st.Name = pb.Name
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.Region = pb.Region
	st.UpdatedTime = pb.UpdatedTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkConnectivityConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkConnectivityConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func networkPolicyEgressToPb(st *NetworkPolicyEgress) (*networkPolicyEgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkPolicyEgressPb{}
	pb.NetworkAccess = st.NetworkAccess

	return pb, nil
}

type networkPolicyEgressPb struct {
	NetworkAccess *EgressNetworkPolicyNetworkAccessPolicy `json:"network_access,omitempty"`
}

func networkPolicyEgressFromPb(pb *networkPolicyEgressPb) (*NetworkPolicyEgress, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkPolicyEgress{}
	st.NetworkAccess = pb.NetworkAccess

	return st, nil
}

func notificationDestinationToPb(st *NotificationDestination) (*notificationDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notificationDestinationPb{}
	pb.Config = st.Config
	pb.DestinationType = st.DestinationType
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type notificationDestinationPb struct {
	Config          *Config         `json:"config,omitempty"`
	DestinationType DestinationType `json:"destination_type,omitempty"`
	DisplayName     string          `json:"display_name,omitempty"`
	Id              string          `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notificationDestinationFromPb(pb *notificationDestinationPb) (*NotificationDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotificationDestination{}
	st.Config = pb.Config
	st.DestinationType = pb.DestinationType
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notificationDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notificationDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pagerdutyConfigToPb(st *PagerdutyConfig) (*pagerdutyConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pagerdutyConfigPb{}
	pb.IntegrationKey = st.IntegrationKey
	pb.IntegrationKeySet = st.IntegrationKeySet

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pagerdutyConfigPb struct {
	IntegrationKey    string `json:"integration_key,omitempty"`
	IntegrationKeySet bool   `json:"integration_key_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pagerdutyConfigFromPb(pb *pagerdutyConfigPb) (*PagerdutyConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PagerdutyConfig{}
	st.IntegrationKey = pb.IntegrationKey
	st.IntegrationKeySet = pb.IntegrationKeySet

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pagerdutyConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pagerdutyConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func partitionIdToPb(st *PartitionId) (*partitionIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &partitionIdPb{}
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type partitionIdPb struct {
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func partitionIdFromPb(pb *partitionIdPb) (*PartitionId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartitionId{}
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *partitionIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st partitionIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func personalComputeMessageToPb(st *PersonalComputeMessage) (*personalComputeMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &personalComputeMessagePb{}
	pb.Value = st.Value

	return pb, nil
}

type personalComputeMessagePb struct {
	Value PersonalComputeMessageEnum `json:"value"`
}

func personalComputeMessageFromPb(pb *personalComputeMessagePb) (*PersonalComputeMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalComputeMessage{}
	st.Value = pb.Value

	return st, nil
}

func personalComputeSettingToPb(st *PersonalComputeSetting) (*personalComputeSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &personalComputeSettingPb{}
	pb.Etag = st.Etag
	pb.PersonalCompute = st.PersonalCompute
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type personalComputeSettingPb struct {
	Etag            string                 `json:"etag,omitempty"`
	PersonalCompute PersonalComputeMessage `json:"personal_compute"`
	SettingName     string                 `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func personalComputeSettingFromPb(pb *personalComputeSettingPb) (*PersonalComputeSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalComputeSetting{}
	st.Etag = pb.Etag
	st.PersonalCompute = pb.PersonalCompute
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *personalComputeSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st personalComputeSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func publicTokenInfoToPb(st *PublicTokenInfo) (*publicTokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publicTokenInfoPb{}
	pb.Comment = st.Comment
	pb.CreationTime = st.CreationTime
	pb.ExpiryTime = st.ExpiryTime
	pb.TokenId = st.TokenId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type publicTokenInfoPb struct {
	Comment      string `json:"comment,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	TokenId      string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func publicTokenInfoFromPb(pb *publicTokenInfoPb) (*PublicTokenInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublicTokenInfo{}
	st.Comment = pb.Comment
	st.CreationTime = pb.CreationTime
	st.ExpiryTime = pb.ExpiryTime
	st.TokenId = pb.TokenId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publicTokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publicTokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func replaceIpAccessListToPb(st *ReplaceIpAccessList) (*replaceIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &replaceIpAccessListPb{}
	pb.Enabled = st.Enabled
	pb.IpAccessListId = st.IpAccessListId
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	pb.ListType = st.ListType

	return pb, nil
}

type replaceIpAccessListPb struct {
	Enabled        bool     `json:"enabled"`
	IpAccessListId string   `json:"-" url:"-"`
	IpAddresses    []string `json:"ip_addresses,omitempty"`
	Label          string   `json:"label"`
	ListType       ListType `json:"list_type"`
}

func replaceIpAccessListFromPb(pb *replaceIpAccessListPb) (*ReplaceIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReplaceIpAccessList{}
	st.Enabled = pb.Enabled
	st.IpAccessListId = pb.IpAccessListId
	st.IpAddresses = pb.IpAddresses
	st.Label = pb.Label
	st.ListType = pb.ListType

	return st, nil
}

func restrictWorkspaceAdminsMessageToPb(st *RestrictWorkspaceAdminsMessage) (*restrictWorkspaceAdminsMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restrictWorkspaceAdminsMessagePb{}
	pb.Status = st.Status

	return pb, nil
}

type restrictWorkspaceAdminsMessagePb struct {
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`
}

func restrictWorkspaceAdminsMessageFromPb(pb *restrictWorkspaceAdminsMessagePb) (*RestrictWorkspaceAdminsMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestrictWorkspaceAdminsMessage{}
	st.Status = pb.Status

	return st, nil
}

func restrictWorkspaceAdminsSettingToPb(st *RestrictWorkspaceAdminsSetting) (*restrictWorkspaceAdminsSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restrictWorkspaceAdminsSettingPb{}
	pb.Etag = st.Etag
	pb.RestrictWorkspaceAdmins = st.RestrictWorkspaceAdmins
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type restrictWorkspaceAdminsSettingPb struct {
	Etag                    string                         `json:"etag,omitempty"`
	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage `json:"restrict_workspace_admins"`
	SettingName             string                         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restrictWorkspaceAdminsSettingFromPb(pb *restrictWorkspaceAdminsSettingPb) (*RestrictWorkspaceAdminsSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestrictWorkspaceAdminsSetting{}
	st.Etag = pb.Etag
	st.RestrictWorkspaceAdmins = pb.RestrictWorkspaceAdmins
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restrictWorkspaceAdminsSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restrictWorkspaceAdminsSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func revokeTokenRequestToPb(st *RevokeTokenRequest) (*revokeTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &revokeTokenRequestPb{}
	pb.TokenId = st.TokenId

	return pb, nil
}

type revokeTokenRequestPb struct {
	TokenId string `json:"token_id"`
}

func revokeTokenRequestFromPb(pb *revokeTokenRequestPb) (*RevokeTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RevokeTokenRequest{}
	st.TokenId = pb.TokenId

	return st, nil
}

func slackConfigToPb(st *SlackConfig) (*slackConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &slackConfigPb{}
	pb.Url = st.Url
	pb.UrlSet = st.UrlSet

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type slackConfigPb struct {
	Url    string `json:"url,omitempty"`
	UrlSet bool   `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func slackConfigFromPb(pb *slackConfigPb) (*SlackConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SlackConfig{}
	st.Url = pb.Url
	st.UrlSet = pb.UrlSet

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *slackConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st slackConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sqlResultsDownloadToPb(st *SqlResultsDownload) (*sqlResultsDownloadPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sqlResultsDownloadPb{}
	pb.BooleanVal = st.BooleanVal
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sqlResultsDownloadPb struct {
	BooleanVal  BooleanMessage `json:"boolean_val"`
	Etag        string         `json:"etag,omitempty"`
	SettingName string         `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sqlResultsDownloadFromPb(pb *sqlResultsDownloadPb) (*SqlResultsDownload, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlResultsDownload{}
	st.BooleanVal = pb.BooleanVal
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sqlResultsDownloadPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sqlResultsDownloadPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func stringMessageToPb(st *StringMessage) (*stringMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stringMessagePb{}
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type stringMessagePb struct {
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func stringMessageFromPb(pb *stringMessagePb) (*StringMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StringMessage{}
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *stringMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st stringMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenAccessControlRequestToPb(st *TokenAccessControlRequest) (*tokenAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenAccessControlRequestPb struct {
	GroupName            string               `json:"group_name,omitempty"`
	PermissionLevel      TokenPermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string               `json:"service_principal_name,omitempty"`
	UserName             string               `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenAccessControlRequestFromPb(pb *tokenAccessControlRequestPb) (*TokenAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenAccessControlResponseToPb(st *TokenAccessControlResponse) (*tokenAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenAccessControlResponsePb struct {
	AllPermissions       []TokenPermission `json:"all_permissions,omitempty"`
	DisplayName          string            `json:"display_name,omitempty"`
	GroupName            string            `json:"group_name,omitempty"`
	ServicePrincipalName string            `json:"service_principal_name,omitempty"`
	UserName             string            `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenAccessControlResponseFromPb(pb *tokenAccessControlResponsePb) (*TokenAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenInfoToPb(st *TokenInfo) (*tokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenInfoPb{}
	pb.Comment = st.Comment
	pb.CreatedById = st.CreatedById
	pb.CreatedByUsername = st.CreatedByUsername
	pb.CreationTime = st.CreationTime
	pb.ExpiryTime = st.ExpiryTime
	pb.LastUsedDay = st.LastUsedDay
	pb.OwnerId = st.OwnerId
	pb.TokenId = st.TokenId
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenInfoPb struct {
	Comment           string `json:"comment,omitempty"`
	CreatedById       int64  `json:"created_by_id,omitempty"`
	CreatedByUsername string `json:"created_by_username,omitempty"`
	CreationTime      int64  `json:"creation_time,omitempty"`
	ExpiryTime        int64  `json:"expiry_time,omitempty"`
	LastUsedDay       int64  `json:"last_used_day,omitempty"`
	OwnerId           int64  `json:"owner_id,omitempty"`
	TokenId           string `json:"token_id,omitempty"`
	WorkspaceId       int64  `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenInfoFromPb(pb *tokenInfoPb) (*TokenInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenInfo{}
	st.Comment = pb.Comment
	st.CreatedById = pb.CreatedById
	st.CreatedByUsername = pb.CreatedByUsername
	st.CreationTime = pb.CreationTime
	st.ExpiryTime = pb.ExpiryTime
	st.LastUsedDay = pb.LastUsedDay
	st.OwnerId = pb.OwnerId
	st.TokenId = pb.TokenId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenPermissionToPb(st *TokenPermission) (*tokenPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenPermissionPb struct {
	Inherited           bool                 `json:"inherited,omitempty"`
	InheritedFromObject []string             `json:"inherited_from_object,omitempty"`
	PermissionLevel     TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenPermissionFromPb(pb *tokenPermissionPb) (*TokenPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenPermissionsToPb(st *TokenPermissions) (*tokenPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenPermissionsPb struct {
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                       `json:"object_id,omitempty"`
	ObjectType        string                       `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenPermissionsFromPb(pb *tokenPermissionsPb) (*TokenPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenPermissionsDescriptionToPb(st *TokenPermissionsDescription) (*tokenPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenPermissionsDescriptionPb struct {
	Description     string               `json:"description,omitempty"`
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenPermissionsDescriptionFromPb(pb *tokenPermissionsDescriptionPb) (*TokenPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenPermissionsRequestToPb(st *TokenPermissionsRequest) (*tokenPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList

	return pb, nil
}

type tokenPermissionsRequestPb struct {
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

func tokenPermissionsRequestFromPb(pb *tokenPermissionsRequestPb) (*TokenPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList

	return st, nil
}

func updateAccountIpAccessEnableRequestToPb(st *UpdateAccountIpAccessEnableRequest) (*updateAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAccountIpAccessEnableRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateAccountIpAccessEnableRequestPb struct {
	AllowMissing bool                  `json:"allow_missing"`
	FieldMask    string                `json:"field_mask"`
	Setting      AccountIpAccessEnable `json:"setting"`
}

func updateAccountIpAccessEnableRequestFromPb(pb *updateAccountIpAccessEnableRequestPb) (*UpdateAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAccountIpAccessEnableRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*updateAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
	AllowMissing bool                                      `json:"allow_missing"`
	FieldMask    string                                    `json:"field_mask"`
	Setting      AibiDashboardEmbeddingAccessPolicySetting `json:"setting"`
}

func updateAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *updateAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*UpdateAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
	AllowMissing bool                                         `json:"allow_missing"`
	FieldMask    string                                       `json:"field_mask"`
	Setting      AibiDashboardEmbeddingApprovedDomainsSetting `json:"setting"`
}

func updateAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateAutomaticClusterUpdateSettingRequestToPb(st *UpdateAutomaticClusterUpdateSettingRequest) (*updateAutomaticClusterUpdateSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAutomaticClusterUpdateSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateAutomaticClusterUpdateSettingRequestPb struct {
	AllowMissing bool                          `json:"allow_missing"`
	FieldMask    string                        `json:"field_mask"`
	Setting      AutomaticClusterUpdateSetting `json:"setting"`
}

func updateAutomaticClusterUpdateSettingRequestFromPb(pb *updateAutomaticClusterUpdateSettingRequestPb) (*UpdateAutomaticClusterUpdateSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAutomaticClusterUpdateSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateComplianceSecurityProfileSettingRequestToPb(st *UpdateComplianceSecurityProfileSettingRequest) (*updateComplianceSecurityProfileSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateComplianceSecurityProfileSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateComplianceSecurityProfileSettingRequestPb struct {
	AllowMissing bool                             `json:"allow_missing"`
	FieldMask    string                           `json:"field_mask"`
	Setting      ComplianceSecurityProfileSetting `json:"setting"`
}

func updateComplianceSecurityProfileSettingRequestFromPb(pb *updateComplianceSecurityProfileSettingRequestPb) (*UpdateComplianceSecurityProfileSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateComplianceSecurityProfileSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateCspEnablementAccountSettingRequestToPb(st *UpdateCspEnablementAccountSettingRequest) (*updateCspEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCspEnablementAccountSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateCspEnablementAccountSettingRequestPb struct {
	AllowMissing bool                        `json:"allow_missing"`
	FieldMask    string                      `json:"field_mask"`
	Setting      CspEnablementAccountSetting `json:"setting"`
}

func updateCspEnablementAccountSettingRequestFromPb(pb *updateCspEnablementAccountSettingRequestPb) (*UpdateCspEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCspEnablementAccountSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateDashboardEmailSubscriptionsRequestToPb(st *UpdateDashboardEmailSubscriptionsRequest) (*updateDashboardEmailSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDashboardEmailSubscriptionsRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateDashboardEmailSubscriptionsRequestPb struct {
	AllowMissing bool                        `json:"allow_missing"`
	FieldMask    string                      `json:"field_mask"`
	Setting      DashboardEmailSubscriptions `json:"setting"`
}

func updateDashboardEmailSubscriptionsRequestFromPb(pb *updateDashboardEmailSubscriptionsRequestPb) (*UpdateDashboardEmailSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDashboardEmailSubscriptionsRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateDefaultNamespaceSettingRequestToPb(st *UpdateDefaultNamespaceSettingRequest) (*updateDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDefaultNamespaceSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateDefaultNamespaceSettingRequestPb struct {
	AllowMissing bool                    `json:"allow_missing"`
	FieldMask    string                  `json:"field_mask"`
	Setting      DefaultNamespaceSetting `json:"setting"`
}

func updateDefaultNamespaceSettingRequestFromPb(pb *updateDefaultNamespaceSettingRequestPb) (*UpdateDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDefaultNamespaceSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateDisableLegacyAccessRequestToPb(st *UpdateDisableLegacyAccessRequest) (*updateDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDisableLegacyAccessRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateDisableLegacyAccessRequestPb struct {
	AllowMissing bool                `json:"allow_missing"`
	FieldMask    string              `json:"field_mask"`
	Setting      DisableLegacyAccess `json:"setting"`
}

func updateDisableLegacyAccessRequestFromPb(pb *updateDisableLegacyAccessRequestPb) (*UpdateDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyAccessRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateDisableLegacyDbfsRequestToPb(st *UpdateDisableLegacyDbfsRequest) (*updateDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDisableLegacyDbfsRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateDisableLegacyDbfsRequestPb struct {
	AllowMissing bool              `json:"allow_missing"`
	FieldMask    string            `json:"field_mask"`
	Setting      DisableLegacyDbfs `json:"setting"`
}

func updateDisableLegacyDbfsRequestFromPb(pb *updateDisableLegacyDbfsRequestPb) (*UpdateDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyDbfsRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateDisableLegacyFeaturesRequestToPb(st *UpdateDisableLegacyFeaturesRequest) (*updateDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDisableLegacyFeaturesRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateDisableLegacyFeaturesRequestPb struct {
	AllowMissing bool                  `json:"allow_missing"`
	FieldMask    string                `json:"field_mask"`
	Setting      DisableLegacyFeatures `json:"setting"`
}

func updateDisableLegacyFeaturesRequestFromPb(pb *updateDisableLegacyFeaturesRequestPb) (*UpdateDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyFeaturesRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateEnableExportNotebookRequestToPb(st *UpdateEnableExportNotebookRequest) (*updateEnableExportNotebookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnableExportNotebookRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateEnableExportNotebookRequestPb struct {
	AllowMissing bool                 `json:"allow_missing"`
	FieldMask    string               `json:"field_mask"`
	Setting      EnableExportNotebook `json:"setting"`
}

func updateEnableExportNotebookRequestFromPb(pb *updateEnableExportNotebookRequestPb) (*UpdateEnableExportNotebookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableExportNotebookRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateEnableNotebookTableClipboardRequestToPb(st *UpdateEnableNotebookTableClipboardRequest) (*updateEnableNotebookTableClipboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnableNotebookTableClipboardRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateEnableNotebookTableClipboardRequestPb struct {
	AllowMissing bool                         `json:"allow_missing"`
	FieldMask    string                       `json:"field_mask"`
	Setting      EnableNotebookTableClipboard `json:"setting"`
}

func updateEnableNotebookTableClipboardRequestFromPb(pb *updateEnableNotebookTableClipboardRequestPb) (*UpdateEnableNotebookTableClipboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableNotebookTableClipboardRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateEnableResultsDownloadingRequestToPb(st *UpdateEnableResultsDownloadingRequest) (*updateEnableResultsDownloadingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnableResultsDownloadingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateEnableResultsDownloadingRequestPb struct {
	AllowMissing bool                     `json:"allow_missing"`
	FieldMask    string                   `json:"field_mask"`
	Setting      EnableResultsDownloading `json:"setting"`
}

func updateEnableResultsDownloadingRequestFromPb(pb *updateEnableResultsDownloadingRequestPb) (*UpdateEnableResultsDownloadingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableResultsDownloadingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateEnhancedSecurityMonitoringSettingRequestToPb(st *UpdateEnhancedSecurityMonitoringSettingRequest) (*updateEnhancedSecurityMonitoringSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnhancedSecurityMonitoringSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateEnhancedSecurityMonitoringSettingRequestPb struct {
	AllowMissing bool                              `json:"allow_missing"`
	FieldMask    string                            `json:"field_mask"`
	Setting      EnhancedSecurityMonitoringSetting `json:"setting"`
}

func updateEnhancedSecurityMonitoringSettingRequestFromPb(pb *updateEnhancedSecurityMonitoringSettingRequestPb) (*UpdateEnhancedSecurityMonitoringSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnhancedSecurityMonitoringSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateEsmEnablementAccountSettingRequestToPb(st *UpdateEsmEnablementAccountSettingRequest) (*updateEsmEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEsmEnablementAccountSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateEsmEnablementAccountSettingRequestPb struct {
	AllowMissing bool                        `json:"allow_missing"`
	FieldMask    string                      `json:"field_mask"`
	Setting      EsmEnablementAccountSetting `json:"setting"`
}

func updateEsmEnablementAccountSettingRequestFromPb(pb *updateEsmEnablementAccountSettingRequestPb) (*UpdateEsmEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEsmEnablementAccountSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateIpAccessListToPb(st *UpdateIpAccessList) (*updateIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateIpAccessListPb{}
	pb.Enabled = st.Enabled
	pb.IpAccessListId = st.IpAccessListId
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	pb.ListType = st.ListType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateIpAccessListPb struct {
	Enabled        bool     `json:"enabled,omitempty"`
	IpAccessListId string   `json:"-" url:"-"`
	IpAddresses    []string `json:"ip_addresses,omitempty"`
	Label          string   `json:"label,omitempty"`
	ListType       ListType `json:"list_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateIpAccessListFromPb(pb *updateIpAccessListPb) (*UpdateIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateIpAccessList{}
	st.Enabled = pb.Enabled
	st.IpAccessListId = pb.IpAccessListId
	st.IpAddresses = pb.IpAddresses
	st.Label = pb.Label
	st.ListType = pb.ListType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateIpAccessListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateIpAccessListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateLlmProxyPartnerPoweredAccountRequestToPb(st *UpdateLlmProxyPartnerPoweredAccountRequest) (*updateLlmProxyPartnerPoweredAccountRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateLlmProxyPartnerPoweredAccountRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateLlmProxyPartnerPoweredAccountRequestPb struct {
	AllowMissing bool                          `json:"allow_missing"`
	FieldMask    string                        `json:"field_mask"`
	Setting      LlmProxyPartnerPoweredAccount `json:"setting"`
}

func updateLlmProxyPartnerPoweredAccountRequestFromPb(pb *updateLlmProxyPartnerPoweredAccountRequestPb) (*UpdateLlmProxyPartnerPoweredAccountRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLlmProxyPartnerPoweredAccountRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateLlmProxyPartnerPoweredEnforceRequestToPb(st *UpdateLlmProxyPartnerPoweredEnforceRequest) (*updateLlmProxyPartnerPoweredEnforceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateLlmProxyPartnerPoweredEnforceRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateLlmProxyPartnerPoweredEnforceRequestPb struct {
	AllowMissing bool                          `json:"allow_missing"`
	FieldMask    string                        `json:"field_mask"`
	Setting      LlmProxyPartnerPoweredEnforce `json:"setting"`
}

func updateLlmProxyPartnerPoweredEnforceRequestFromPb(pb *updateLlmProxyPartnerPoweredEnforceRequestPb) (*UpdateLlmProxyPartnerPoweredEnforceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLlmProxyPartnerPoweredEnforceRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateLlmProxyPartnerPoweredWorkspaceRequestToPb(st *UpdateLlmProxyPartnerPoweredWorkspaceRequest) (*updateLlmProxyPartnerPoweredWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateLlmProxyPartnerPoweredWorkspaceRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateLlmProxyPartnerPoweredWorkspaceRequestPb struct {
	AllowMissing bool                            `json:"allow_missing"`
	FieldMask    string                          `json:"field_mask"`
	Setting      LlmProxyPartnerPoweredWorkspace `json:"setting"`
}

func updateLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb *updateLlmProxyPartnerPoweredWorkspaceRequestPb) (*UpdateLlmProxyPartnerPoweredWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLlmProxyPartnerPoweredWorkspaceRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateNccPrivateEndpointRuleRequestToPb(st *UpdateNccPrivateEndpointRuleRequest) (*updateNccPrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateNccPrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PrivateEndpointRule = st.PrivateEndpointRule
	pb.PrivateEndpointRuleId = st.PrivateEndpointRuleId
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

type updateNccPrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string                    `json:"-" url:"-"`
	PrivateEndpointRule         UpdatePrivateEndpointRule `json:"private_endpoint_rule"`
	PrivateEndpointRuleId       string                    `json:"-" url:"-"`
	UpdateMask                  string                    `json:"-" url:"update_mask"`
}

func updateNccPrivateEndpointRuleRequestFromPb(pb *updateNccPrivateEndpointRuleRequestPb) (*UpdateNccPrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNccPrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PrivateEndpointRule = pb.PrivateEndpointRule
	st.PrivateEndpointRuleId = pb.PrivateEndpointRuleId
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

func updateNetworkPolicyRequestToPb(st *UpdateNetworkPolicyRequest) (*updateNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateNetworkPolicyRequestPb{}
	pb.NetworkPolicy = st.NetworkPolicy
	pb.NetworkPolicyId = st.NetworkPolicyId

	return pb, nil
}

type updateNetworkPolicyRequestPb struct {
	NetworkPolicy   AccountNetworkPolicy `json:"network_policy"`
	NetworkPolicyId string               `json:"-" url:"-"`
}

func updateNetworkPolicyRequestFromPb(pb *updateNetworkPolicyRequestPb) (*UpdateNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNetworkPolicyRequest{}
	st.NetworkPolicy = pb.NetworkPolicy
	st.NetworkPolicyId = pb.NetworkPolicyId

	return st, nil
}

func updateNotificationDestinationRequestToPb(st *UpdateNotificationDestinationRequest) (*updateNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateNotificationDestinationRequestPb{}
	pb.Config = st.Config
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateNotificationDestinationRequestPb struct {
	Config      *Config `json:"config,omitempty"`
	DisplayName string  `json:"display_name,omitempty"`
	Id          string  `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateNotificationDestinationRequestFromPb(pb *updateNotificationDestinationRequestPb) (*UpdateNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNotificationDestinationRequest{}
	st.Config = pb.Config
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateNotificationDestinationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateNotificationDestinationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updatePersonalComputeSettingRequestToPb(st *UpdatePersonalComputeSettingRequest) (*updatePersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePersonalComputeSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updatePersonalComputeSettingRequestPb struct {
	AllowMissing bool                   `json:"allow_missing"`
	FieldMask    string                 `json:"field_mask"`
	Setting      PersonalComputeSetting `json:"setting"`
}

func updatePersonalComputeSettingRequestFromPb(pb *updatePersonalComputeSettingRequestPb) (*UpdatePersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalComputeSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updatePrivateEndpointRuleToPb(st *UpdatePrivateEndpointRule) (*updatePrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePrivateEndpointRulePb{}
	pb.DomainNames = st.DomainNames
	pb.Enabled = st.Enabled
	pb.ResourceNames = st.ResourceNames

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updatePrivateEndpointRulePb struct {
	DomainNames   []string `json:"domain_names,omitempty"`
	Enabled       bool     `json:"enabled,omitempty"`
	ResourceNames []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updatePrivateEndpointRuleFromPb(pb *updatePrivateEndpointRulePb) (*UpdatePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePrivateEndpointRule{}
	st.DomainNames = pb.DomainNames
	st.Enabled = pb.Enabled
	st.ResourceNames = pb.ResourceNames

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updatePrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updatePrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateRestrictWorkspaceAdminsSettingRequestToPb(st *UpdateRestrictWorkspaceAdminsSettingRequest) (*updateRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRestrictWorkspaceAdminsSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateRestrictWorkspaceAdminsSettingRequestPb struct {
	AllowMissing bool                           `json:"allow_missing"`
	FieldMask    string                         `json:"field_mask"`
	Setting      RestrictWorkspaceAdminsSetting `json:"setting"`
}

func updateRestrictWorkspaceAdminsSettingRequestFromPb(pb *updateRestrictWorkspaceAdminsSettingRequestPb) (*UpdateRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRestrictWorkspaceAdminsSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateSqlResultsDownloadRequestToPb(st *UpdateSqlResultsDownloadRequest) (*updateSqlResultsDownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateSqlResultsDownloadRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	pb.Setting = st.Setting

	return pb, nil
}

type updateSqlResultsDownloadRequestPb struct {
	AllowMissing bool               `json:"allow_missing"`
	FieldMask    string             `json:"field_mask"`
	Setting      SqlResultsDownload `json:"setting"`
}

func updateSqlResultsDownloadRequestFromPb(pb *updateSqlResultsDownloadRequestPb) (*UpdateSqlResultsDownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSqlResultsDownloadRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	st.Setting = pb.Setting

	return st, nil
}

func updateWorkspaceNetworkOptionRequestToPb(st *UpdateWorkspaceNetworkOptionRequest) (*updateWorkspaceNetworkOptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWorkspaceNetworkOptionRequestPb{}
	pb.WorkspaceId = st.WorkspaceId
	pb.WorkspaceNetworkOption = st.WorkspaceNetworkOption

	return pb, nil
}

type updateWorkspaceNetworkOptionRequestPb struct {
	WorkspaceId            int64                  `json:"-" url:"-"`
	WorkspaceNetworkOption WorkspaceNetworkOption `json:"workspace_network_option"`
}

func updateWorkspaceNetworkOptionRequestFromPb(pb *updateWorkspaceNetworkOptionRequestPb) (*UpdateWorkspaceNetworkOptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceNetworkOptionRequest{}
	st.WorkspaceId = pb.WorkspaceId
	st.WorkspaceNetworkOption = pb.WorkspaceNetworkOption

	return st, nil
}

type workspaceConfPb WorkspaceConf

func workspaceConfToPb(st *WorkspaceConf) (*workspaceConfPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := workspaceConfPb(*st)
	return &stPb, nil
}

func workspaceConfFromPb(stPb *workspaceConfPb) (*WorkspaceConf, error) {
	if stPb == nil {
		return nil, nil
	}
	st := WorkspaceConf(*stPb)
	return &st, nil
}

func workspaceNetworkOptionToPb(st *WorkspaceNetworkOption) (*workspaceNetworkOptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &workspaceNetworkOptionPb{}
	pb.NetworkPolicyId = st.NetworkPolicyId
	pb.WorkspaceId = st.WorkspaceId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type workspaceNetworkOptionPb struct {
	NetworkPolicyId string `json:"network_policy_id,omitempty"`
	WorkspaceId     int64  `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func workspaceNetworkOptionFromPb(pb *workspaceNetworkOptionPb) (*WorkspaceNetworkOption, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceNetworkOption{}
	st.NetworkPolicyId = pb.NetworkPolicyId
	st.WorkspaceId = pb.WorkspaceId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *workspaceNetworkOptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st workspaceNetworkOptionPb) MarshalJSON() ([]byte, error) {
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
