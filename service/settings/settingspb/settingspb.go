// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settingspb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AccountIpAccessEnablePb struct {
	AcctIpAclEnable BooleanMessagePb `json:"acct_ip_acl_enable"`
	Etag            string           `json:"etag,omitempty"`
	SettingName     string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AccountIpAccessEnablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AccountIpAccessEnablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AccountNetworkPolicyPb struct {
	AccountId       string                 `json:"account_id,omitempty"`
	Egress          *NetworkPolicyEgressPb `json:"egress,omitempty"`
	NetworkPolicyId string                 `json:"network_policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AccountNetworkPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AccountNetworkPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AibiDashboardEmbeddingAccessPolicyPb struct {
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb `json:"access_policy_type"`
}

type AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb string

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb = `ALLOW_ALL_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeAllowApprovedDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb = `ALLOW_APPROVED_DOMAINS`

const AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeDenyAllDomains AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb = `DENY_ALL_DOMAINS`

type AibiDashboardEmbeddingAccessPolicySettingPb struct {
	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicyPb `json:"aibi_dashboard_embedding_access_policy"`
	Etag                               string                               `json:"etag,omitempty"`
	SettingName                        string                               `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AibiDashboardEmbeddingAccessPolicySettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AibiDashboardEmbeddingAccessPolicySettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AibiDashboardEmbeddingApprovedDomainsPb struct {
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

type AibiDashboardEmbeddingApprovedDomainsSettingPb struct {
	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomainsPb `json:"aibi_dashboard_embedding_approved_domains"`
	Etag                                  string                                  `json:"etag,omitempty"`
	SettingName                           string                                  `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AibiDashboardEmbeddingApprovedDomainsSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AibiDashboardEmbeddingApprovedDomainsSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutomaticClusterUpdateSettingPb struct {
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessagePb `json:"automatic_cluster_update_workspace"`
	Etag                            string                      `json:"etag,omitempty"`
	SettingName                     string                      `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AutomaticClusterUpdateSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AutomaticClusterUpdateSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BooleanMessagePb struct {
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BooleanMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BooleanMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAutoRestartMessagePb struct {
	CanToggle                       bool                                          `json:"can_toggle,omitempty"`
	Enabled                         bool                                          `json:"enabled,omitempty"`
	EnablementDetails               *ClusterAutoRestartMessageEnablementDetailsPb `json:"enablement_details,omitempty"`
	MaintenanceWindow               *ClusterAutoRestartMessageMaintenanceWindowPb `json:"maintenance_window,omitempty"`
	RestartEvenIfNoUpdatesAvailable bool                                          `json:"restart_even_if_no_updates_available,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterAutoRestartMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterAutoRestartMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAutoRestartMessageEnablementDetailsPb struct {
	ForcedForComplianceMode           bool `json:"forced_for_compliance_mode,omitempty"`
	UnavailableForDisabledEntitlement bool `json:"unavailable_for_disabled_entitlement,omitempty"`
	UnavailableForNonEnterpriseTier   bool `json:"unavailable_for_non_enterprise_tier,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterAutoRestartMessageEnablementDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterAutoRestartMessageEnablementDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAutoRestartMessageMaintenanceWindowPb struct {
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb `json:"week_day_based_schedule,omitempty"`
}

type ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb string

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFriday ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb = `FRIDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekMonday ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb = `MONDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSaturday ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb = `SATURDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekSunday ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb = `SUNDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekThursday ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb = `THURSDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekTuesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb = `TUESDAY`

const ClusterAutoRestartMessageMaintenanceWindowDayOfWeekWednesday ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb = `WEDNESDAY`

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb struct {
	DayOfWeek       ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb        `json:"day_of_week,omitempty"`
	Frequency       ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb `json:"frequency,omitempty"`
	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb `json:"window_start_time,omitempty"`
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb string

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyEveryWeek ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb = `EVERY_WEEK`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstAndThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb = `FIRST_AND_THIRD_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFirstOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb = `FIRST_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb = `FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondAndFourthOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb = `SECOND_AND_FOURTH_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencySecondOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb = `SECOND_OF_MONTH`

const ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyThirdOfMonth ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb = `THIRD_OF_MONTH`

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb struct {
	Hours   int `json:"hours,omitempty"`
	Minutes int `json:"minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComplianceSecurityProfilePb struct {
	ComplianceStandards []ComplianceStandardPb `json:"compliance_standards,omitempty"`
	IsEnabled           bool                   `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ComplianceSecurityProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ComplianceSecurityProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComplianceSecurityProfileSettingPb struct {
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfilePb `json:"compliance_security_profile_workspace"`
	Etag                               string                      `json:"etag,omitempty"`
	SettingName                        string                      `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ComplianceSecurityProfileSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ComplianceSecurityProfileSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComplianceStandardPb string

const ComplianceStandardCanadaProtectedB ComplianceStandardPb = `CANADA_PROTECTED_B`

const ComplianceStandardCyberEssentialPlus ComplianceStandardPb = `CYBER_ESSENTIAL_PLUS`

const ComplianceStandardFedrampHigh ComplianceStandardPb = `FEDRAMP_HIGH`

const ComplianceStandardFedrampIl5 ComplianceStandardPb = `FEDRAMP_IL5`

const ComplianceStandardFedrampModerate ComplianceStandardPb = `FEDRAMP_MODERATE`

const ComplianceStandardGermanyC5 ComplianceStandardPb = `GERMANY_C5`

const ComplianceStandardHipaa ComplianceStandardPb = `HIPAA`

const ComplianceStandardHitrust ComplianceStandardPb = `HITRUST`

const ComplianceStandardIrapProtected ComplianceStandardPb = `IRAP_PROTECTED`

const ComplianceStandardIsmap ComplianceStandardPb = `ISMAP`

const ComplianceStandardItarEar ComplianceStandardPb = `ITAR_EAR`

const ComplianceStandardKFsi ComplianceStandardPb = `K_FSI`

const ComplianceStandardNone ComplianceStandardPb = `NONE`

const ComplianceStandardPciDss ComplianceStandardPb = `PCI_DSS`

type ConfigPb struct {
	Email          *EmailConfigPb          `json:"email,omitempty"`
	GenericWebhook *GenericWebhookConfigPb `json:"generic_webhook,omitempty"`
	MicrosoftTeams *MicrosoftTeamsConfigPb `json:"microsoft_teams,omitempty"`
	Pagerduty      *PagerdutyConfigPb      `json:"pagerduty,omitempty"`
	Slack          *SlackConfigPb          `json:"slack,omitempty"`
}

type CreateIpAccessListPb struct {
	IpAddresses []string   `json:"ip_addresses,omitempty"`
	Label       string     `json:"label"`
	ListType    ListTypePb `json:"list_type"`
}

type CreateIpAccessListResponsePb struct {
	IpAccessList *IpAccessListInfoPb `json:"ip_access_list,omitempty"`
}

type CreateNetworkConnectivityConfigRequestPb struct {
	NetworkConnectivityConfig CreateNetworkConnectivityConfigurationPb `json:"network_connectivity_config"`
}

type CreateNetworkConnectivityConfigurationPb struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

type CreateNetworkPolicyRequestPb struct {
	NetworkPolicy AccountNetworkPolicyPb `json:"network_policy"`
}

type CreateNotificationDestinationRequestPb struct {
	Config      *ConfigPb `json:"config,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateNotificationDestinationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateNotificationDestinationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateOboTokenRequestPb struct {
	ApplicationId   string `json:"application_id"`
	Comment         string `json:"comment,omitempty"`
	LifetimeSeconds int64  `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateOboTokenRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateOboTokenRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateOboTokenResponsePb struct {
	TokenInfo  *TokenInfoPb `json:"token_info,omitempty"`
	TokenValue string       `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateOboTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateOboTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePrivateEndpointRulePb struct {
	DomainNames     []string `json:"domain_names,omitempty"`
	EndpointService string   `json:"endpoint_service,omitempty"`
	GroupId         string   `json:"group_id,omitempty"`
	ResourceId      string   `json:"resource_id,omitempty"`
	ResourceNames   []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string                      `json:"-" url:"-"`
	PrivateEndpointRule         CreatePrivateEndpointRulePb `json:"private_endpoint_rule"`
}

type CreateTokenRequestPb struct {
	Comment         string `json:"comment,omitempty"`
	LifetimeSeconds int64  `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateTokenRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateTokenRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateTokenResponsePb struct {
	TokenInfo  *PublicTokenInfoPb `json:"token_info,omitempty"`
	TokenValue string             `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CspEnablementAccountPb struct {
	ComplianceStandards []ComplianceStandardPb `json:"compliance_standards,omitempty"`
	IsEnforced          bool                   `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CspEnablementAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CspEnablementAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CspEnablementAccountSettingPb struct {
	CspEnablementAccount CspEnablementAccountPb `json:"csp_enablement_account"`
	Etag                 string                 `json:"etag,omitempty"`
	SettingName          string                 `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CspEnablementAccountSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CspEnablementAccountSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb struct {
	AccountId                   string                                                                                    `json:"account_id,omitempty"`
	ConnectionState             CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb `json:"connection_state,omitempty"`
	CreationTime                int64                                                                                     `json:"creation_time,omitempty"`
	Deactivated                 bool                                                                                      `json:"deactivated,omitempty"`
	DeactivatedAt               int64                                                                                     `json:"deactivated_at,omitempty"`
	DomainNames                 []string                                                                                  `json:"domain_names,omitempty"`
	Enabled                     bool                                                                                      `json:"enabled,omitempty"`
	EndpointService             string                                                                                    `json:"endpoint_service,omitempty"`
	NetworkConnectivityConfigId string                                                                                    `json:"network_connectivity_config_id,omitempty"`
	ResourceNames               []string                                                                                  `json:"resource_names,omitempty"`
	RuleId                      string                                                                                    `json:"rule_id,omitempty"`
	UpdatedTime                 int64                                                                                     `json:"updated_time,omitempty"`
	VpcEndpointId               string                                                                                    `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb string

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateDisconnected CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb = `DISCONNECTED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateEstablished CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb = `ESTABLISHED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateExpired CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb = `EXPIRED`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePending CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb = `PENDING`

const CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateRejected CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb = `REJECTED`

type DashboardEmailSubscriptionsPb struct {
	BooleanVal  BooleanMessagePb `json:"boolean_val"`
	Etag        string           `json:"etag,omitempty"`
	SettingName string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DashboardEmailSubscriptionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DashboardEmailSubscriptionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DefaultNamespaceSettingPb struct {
	Etag        string          `json:"etag,omitempty"`
	Namespace   StringMessagePb `json:"namespace"`
	SettingName string          `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DefaultNamespaceSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DefaultNamespaceSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DefaultWarehouseIdPb struct {
	Etag        string          `json:"etag,omitempty"`
	SettingName string          `json:"setting_name,omitempty"`
	StringVal   StringMessagePb `json:"string_val"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DefaultWarehouseIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DefaultWarehouseIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAccountIpAccessEnableRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteAccountIpAccessEnableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteAccountIpAccessEnableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAccountIpAccessEnableResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteAccountIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

type DeleteAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAibiDashboardEmbeddingAccessPolicySettingResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteDashboardEmailSubscriptionsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDashboardEmailSubscriptionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDashboardEmailSubscriptionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDashboardEmailSubscriptionsResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteDefaultNamespaceSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDefaultNamespaceSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDefaultNamespaceSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDefaultNamespaceSettingResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteDefaultWarehouseIdRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDefaultWarehouseIdRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDefaultWarehouseIdRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDefaultWarehouseIdResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteDisableLegacyAccessRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDisableLegacyAccessRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDisableLegacyAccessRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDisableLegacyAccessResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteDisableLegacyDbfsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDisableLegacyDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDisableLegacyDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDisableLegacyDbfsResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteDisableLegacyFeaturesRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteDisableLegacyFeaturesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteDisableLegacyFeaturesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDisableLegacyFeaturesResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

type DeleteLlmProxyPartnerPoweredWorkspaceRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteLlmProxyPartnerPoweredWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteLlmProxyPartnerPoweredWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteLlmProxyPartnerPoweredWorkspaceResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteNetworkConnectivityConfigurationRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

type DeleteNetworkPolicyRequestPb struct {
	NetworkPolicyId string `json:"-" url:"-"`
}

type DeleteNotificationDestinationRequestPb struct {
	Id string `json:"-" url:"-"`
}

type DeletePersonalComputeSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeletePersonalComputeSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeletePersonalComputeSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeletePersonalComputeSettingResponsePb struct {
	Etag string `json:"etag"`
}

type DeletePrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	PrivateEndpointRuleId       string `json:"-" url:"-"`
}

type DeleteRestrictWorkspaceAdminsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteRestrictWorkspaceAdminsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteRestrictWorkspaceAdminsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteRestrictWorkspaceAdminsSettingResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteSqlResultsDownloadRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeleteSqlResultsDownloadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeleteSqlResultsDownloadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteSqlResultsDownloadResponsePb struct {
	Etag string `json:"etag"`
}

type DeleteTokenManagementRequestPb struct {
	TokenId string `json:"-" url:"-"`
}

type DestinationTypePb string

const DestinationTypeEmail DestinationTypePb = `EMAIL`

const DestinationTypeMicrosoftTeams DestinationTypePb = `MICROSOFT_TEAMS`

const DestinationTypePagerduty DestinationTypePb = `PAGERDUTY`

const DestinationTypeSlack DestinationTypePb = `SLACK`

const DestinationTypeWebhook DestinationTypePb = `WEBHOOK`

type DisableLegacyAccessPb struct {
	DisableLegacyAccess BooleanMessagePb `json:"disable_legacy_access"`
	Etag                string           `json:"etag,omitempty"`
	SettingName         string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DisableLegacyAccessPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DisableLegacyAccessPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DisableLegacyDbfsPb struct {
	DisableLegacyDbfs BooleanMessagePb `json:"disable_legacy_dbfs"`
	Etag              string           `json:"etag,omitempty"`
	SettingName       string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DisableLegacyDbfsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DisableLegacyDbfsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DisableLegacyFeaturesPb struct {
	DisableLegacyFeatures BooleanMessagePb `json:"disable_legacy_features"`
	Etag                  string           `json:"etag,omitempty"`
	SettingName           string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DisableLegacyFeaturesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DisableLegacyFeaturesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EgressNetworkPolicyPb struct {
	InternetAccess *EgressNetworkPolicyInternetAccessPolicyPb `json:"internet_access,omitempty"`
}

type EgressNetworkPolicyInternetAccessPolicyPb struct {
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb `json:"allowed_internet_destinations,omitempty"`
	AllowedStorageDestinations  []EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb  `json:"allowed_storage_destinations,omitempty"`
	LogOnlyMode                 *EgressNetworkPolicyInternetAccessPolicyLogOnlyModePb          `json:"log_only_mode,omitempty"`
	RestrictionMode             EgressNetworkPolicyInternetAccessPolicyRestrictionModePb       `json:"restriction_mode,omitempty"`
}

type EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb struct {
	Destination string                                                                                           `json:"destination,omitempty"`
	Protocol    EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb `json:"protocol,omitempty"`
	Type        EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb              `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolTcp EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb = `TCP`

type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb string

const EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFqdn EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb = `FQDN`

type EgressNetworkPolicyInternetAccessPolicyLogOnlyModePb struct {
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb `json:"log_only_mode_type,omitempty"`
	Workloads       []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb  `json:"workloads,omitempty"`
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeAllServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb = `ALL_SERVICES`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeSelectedServices EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb = `SELECTED_SERVICES`

type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb string

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeDbsql EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb = `DBSQL`

const EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeMlServing EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb = `ML_SERVING`

type EgressNetworkPolicyInternetAccessPolicyRestrictionModePb string

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeFullAccess EgressNetworkPolicyInternetAccessPolicyRestrictionModePb = `FULL_ACCESS`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModePrivateAccessOnly EgressNetworkPolicyInternetAccessPolicyRestrictionModePb = `PRIVATE_ACCESS_ONLY`

const EgressNetworkPolicyInternetAccessPolicyRestrictionModeRestrictedAccess EgressNetworkPolicyInternetAccessPolicyRestrictionModePb = `RESTRICTED_ACCESS`

type EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb struct {
	AllowedPaths        []string                                                                          `json:"allowed_paths,omitempty"`
	AzureContainer      string                                                                            `json:"azure_container,omitempty"`
	AzureDnsZone        string                                                                            `json:"azure_dns_zone,omitempty"`
	AzureStorageAccount string                                                                            `json:"azure_storage_account,omitempty"`
	AzureStorageService string                                                                            `json:"azure_storage_service,omitempty"`
	BucketName          string                                                                            `json:"bucket_name,omitempty"`
	Region              string                                                                            `json:"region,omitempty"`
	Type                EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb string

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAwsS3 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb = `AWS_S3`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb = `AZURE_STORAGE`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeCloudflareR2 EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb = `CLOUDFLARE_R2`

const EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb = `GOOGLE_CLOUD_STORAGE`

type EgressNetworkPolicyNetworkAccessPolicyPb struct {
	AllowedInternetDestinations []EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb `json:"allowed_internet_destinations,omitempty"`
	AllowedStorageDestinations  []EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb  `json:"allowed_storage_destinations,omitempty"`
	PolicyEnforcement           *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb    `json:"policy_enforcement,omitempty"`
	RestrictionMode             EgressNetworkPolicyNetworkAccessPolicyRestrictionModePb       `json:"restriction_mode"`
}

type EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb struct {
	Destination             string                                                                             `json:"destination,omitempty"`
	InternetDestinationType EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypePb `json:"internet_destination_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypePb string

const EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeDnsName EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypePb = `DNS_NAME`

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb struct {
	DryRunModeProductFilter []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb `json:"dry_run_mode_product_filter,omitempty"`
	EnforcementMode         EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModePb           `json:"enforcement_mode,omitempty"`
}

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb string

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterDbsql EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb = `DBSQL`

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterMlServing EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb = `ML_SERVING`

type EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModePb string

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeDryRun EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModePb = `DRY_RUN`

const EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeEnforced EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModePb = `ENFORCED`

type EgressNetworkPolicyNetworkAccessPolicyRestrictionModePb string

const EgressNetworkPolicyNetworkAccessPolicyRestrictionModeFullAccess EgressNetworkPolicyNetworkAccessPolicyRestrictionModePb = `FULL_ACCESS`

const EgressNetworkPolicyNetworkAccessPolicyRestrictionModeRestrictedAccess EgressNetworkPolicyNetworkAccessPolicyRestrictionModePb = `RESTRICTED_ACCESS`

type EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb struct {
	AzureStorageAccount    string                                                                           `json:"azure_storage_account,omitempty"`
	AzureStorageService    string                                                                           `json:"azure_storage_service,omitempty"`
	BucketName             string                                                                           `json:"bucket_name,omitempty"`
	Region                 string                                                                           `json:"region,omitempty"`
	StorageDestinationType EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb `json:"storage_destination_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb string

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAwsS3 EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb = `AWS_S3`

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeAzureStorage EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb = `AZURE_STORAGE`

const EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeGoogleCloudStorage EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb = `GOOGLE_CLOUD_STORAGE`

type EgressResourceTypePb string

const EgressResourceTypeAzureBlobStorage EgressResourceTypePb = `AZURE_BLOB_STORAGE`

type EmailConfigPb struct {
	Addresses []string `json:"addresses,omitempty"`
}

type EmptyPb struct {
}

type EnableExportNotebookPb struct {
	BooleanVal  *BooleanMessagePb `json:"boolean_val,omitempty"`
	SettingName string            `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnableExportNotebookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnableExportNotebookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnableNotebookTableClipboardPb struct {
	BooleanVal  *BooleanMessagePb `json:"boolean_val,omitempty"`
	SettingName string            `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnableNotebookTableClipboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnableNotebookTableClipboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnableResultsDownloadingPb struct {
	BooleanVal  *BooleanMessagePb `json:"boolean_val,omitempty"`
	SettingName string            `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnableResultsDownloadingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnableResultsDownloadingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnhancedSecurityMonitoringPb struct {
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnhancedSecurityMonitoringPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnhancedSecurityMonitoringPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnhancedSecurityMonitoringSettingPb struct {
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoringPb `json:"enhanced_security_monitoring_workspace"`
	Etag                                string                       `json:"etag,omitempty"`
	SettingName                         string                       `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EnhancedSecurityMonitoringSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EnhancedSecurityMonitoringSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EsmEnablementAccountPb struct {
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EsmEnablementAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EsmEnablementAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EsmEnablementAccountSettingPb struct {
	EsmEnablementAccount EsmEnablementAccountPb `json:"esm_enablement_account"`
	Etag                 string                 `json:"etag,omitempty"`
	SettingName          string                 `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EsmEnablementAccountSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EsmEnablementAccountSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExchangeTokenPb struct {
	Credential        string      `json:"credential,omitempty"`
	CredentialEolTime int64       `json:"credentialEolTime,omitempty"`
	OwnerId           int64       `json:"ownerId,omitempty"`
	Scopes            []string    `json:"scopes,omitempty"`
	TokenType         TokenTypePb `json:"tokenType,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExchangeTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExchangeTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExchangeTokenRequestPb struct {
	PartitionId PartitionIdPb `json:"partitionId"`
	Scopes      []string      `json:"scopes"`
	TokenType   []TokenTypePb `json:"tokenType"`
}

type ExchangeTokenResponsePb struct {
	Values []ExchangeTokenPb `json:"values,omitempty"`
}

type FetchIpAccessListResponsePb struct {
	IpAccessList *IpAccessListInfoPb `json:"ip_access_list,omitempty"`
}

type GenericWebhookConfigPb struct {
	Password    string `json:"password,omitempty"`
	PasswordSet bool   `json:"password_set,omitempty"`
	Url         string `json:"url,omitempty"`
	UrlSet      bool   `json:"url_set,omitempty"`
	Username    string `json:"username,omitempty"`
	UsernameSet bool   `json:"username_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GenericWebhookConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GenericWebhookConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAccountIpAccessEnableRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetAccountIpAccessEnableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetAccountIpAccessEnableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAccountIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

type GetAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetAibiDashboardEmbeddingAccessPolicySettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetAibiDashboardEmbeddingAccessPolicySettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAutomaticClusterUpdateSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetAutomaticClusterUpdateSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetAutomaticClusterUpdateSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetComplianceSecurityProfileSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetComplianceSecurityProfileSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetComplianceSecurityProfileSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetCspEnablementAccountSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetCspEnablementAccountSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetCspEnablementAccountSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDashboardEmailSubscriptionsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetDashboardEmailSubscriptionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetDashboardEmailSubscriptionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDefaultNamespaceSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetDefaultNamespaceSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetDefaultNamespaceSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDefaultWarehouseIdRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetDefaultWarehouseIdRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetDefaultWarehouseIdRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDisableLegacyAccessRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetDisableLegacyAccessRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetDisableLegacyAccessRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDisableLegacyDbfsRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetDisableLegacyDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetDisableLegacyDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDisableLegacyFeaturesRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetDisableLegacyFeaturesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetDisableLegacyFeaturesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetEnableExportNotebookRequestPb struct {
}

type GetEnableNotebookTableClipboardRequestPb struct {
}

type GetEnableResultsDownloadingRequestPb struct {
}

type GetEnhancedSecurityMonitoringSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetEnhancedSecurityMonitoringSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetEnhancedSecurityMonitoringSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetEsmEnablementAccountSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetEsmEnablementAccountSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetEsmEnablementAccountSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetIpAccessListRequestPb struct {
	IpAccessListId string `json:"-" url:"-"`
}

type GetIpAccessListResponsePb struct {
	IpAccessList *IpAccessListInfoPb `json:"ip_access_list,omitempty"`
}

type GetIpAccessListsResponsePb struct {
	IpAccessLists []IpAccessListInfoPb `json:"ip_access_lists,omitempty"`
}

type GetLlmProxyPartnerPoweredAccountRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetLlmProxyPartnerPoweredAccountRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetLlmProxyPartnerPoweredAccountRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetLlmProxyPartnerPoweredEnforceRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetLlmProxyPartnerPoweredEnforceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetLlmProxyPartnerPoweredEnforceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetLlmProxyPartnerPoweredWorkspaceRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetLlmProxyPartnerPoweredWorkspaceRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetLlmProxyPartnerPoweredWorkspaceRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetNetworkConnectivityConfigurationRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

type GetNetworkPolicyRequestPb struct {
	NetworkPolicyId string `json:"-" url:"-"`
}

type GetNotificationDestinationRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetPersonalComputeSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPersonalComputeSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPersonalComputeSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	PrivateEndpointRuleId       string `json:"-" url:"-"`
}

type GetRestrictWorkspaceAdminsSettingRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetRestrictWorkspaceAdminsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetRestrictWorkspaceAdminsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetSqlResultsDownloadRequestPb struct {
	Etag string `json:"-" url:"etag,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetSqlResultsDownloadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetSqlResultsDownloadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetStatusRequestPb struct {
	Keys string `json:"-" url:"keys"`
}

type GetTokenManagementRequestPb struct {
	TokenId string `json:"-" url:"-"`
}

type GetTokenPermissionLevelsRequestPb struct {
}

type GetTokenPermissionLevelsResponsePb struct {
	PermissionLevels []TokenPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetTokenPermissionsRequestPb struct {
}

type GetTokenResponsePb struct {
	TokenInfo *TokenInfoPb `json:"token_info,omitempty"`
}

type GetWorkspaceNetworkOptionRequestPb struct {
	WorkspaceId int64 `json:"-" url:"-"`
}

type IpAccessListInfoPb struct {
	AddressCount int        `json:"address_count,omitempty"`
	CreatedAt    int64      `json:"created_at,omitempty"`
	CreatedBy    int64      `json:"created_by,omitempty"`
	Enabled      bool       `json:"enabled,omitempty"`
	IpAddresses  []string   `json:"ip_addresses,omitempty"`
	Label        string     `json:"label,omitempty"`
	ListId       string     `json:"list_id,omitempty"`
	ListType     ListTypePb `json:"list_type,omitempty"`
	UpdatedAt    int64      `json:"updated_at,omitempty"`
	UpdatedBy    int64      `json:"updated_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *IpAccessListInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st IpAccessListInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListIpAccessListResponsePb struct {
	IpAccessLists []IpAccessListInfoPb `json:"ip_access_lists,omitempty"`
}

type ListIpAccessListsPb struct {
}

type ListNetworkConnectivityConfigurationsRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListNetworkConnectivityConfigurationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListNetworkConnectivityConfigurationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNetworkConnectivityConfigurationsResponsePb struct {
	Items         []NetworkConnectivityConfigurationPb `json:"items,omitempty"`
	NextPageToken string                               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListNetworkConnectivityConfigurationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListNetworkConnectivityConfigurationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNetworkPoliciesRequestPb struct {
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListNetworkPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListNetworkPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNetworkPoliciesResponsePb struct {
	Items         []AccountNetworkPolicyPb `json:"items,omitempty"`
	NextPageToken string                   `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListNetworkPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListNetworkPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNotificationDestinationsRequestPb struct {
	PageSize  int64  `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListNotificationDestinationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListNotificationDestinationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNotificationDestinationsResponsePb struct {
	NextPageToken string                                 `json:"next_page_token,omitempty"`
	Results       []ListNotificationDestinationsResultPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListNotificationDestinationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListNotificationDestinationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNotificationDestinationsResultPb struct {
	DestinationType DestinationTypePb `json:"destination_type,omitempty"`
	DisplayName     string            `json:"display_name,omitempty"`
	Id              string            `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListNotificationDestinationsResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListNotificationDestinationsResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPrivateEndpointRulesRequestPb struct {
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	PageToken                   string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPrivateEndpointRulesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPrivateEndpointRulesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPrivateEndpointRulesResponsePb struct {
	Items         []NccPrivateEndpointRulePb `json:"items,omitempty"`
	NextPageToken string                     `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPrivateEndpointRulesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPrivateEndpointRulesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPublicTokensResponsePb struct {
	TokenInfos []PublicTokenInfoPb `json:"token_infos,omitempty"`
}

type ListTokenManagementRequestPb struct {
	CreatedById       int64  `json:"-" url:"created_by_id,omitempty"`
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListTokenManagementRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListTokenManagementRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListTokensPb struct {
}

type ListTokensResponsePb struct {
	TokenInfos []TokenInfoPb `json:"token_infos,omitempty"`
}

type ListTypePb string

// An allow list. Include this IP or range.
const ListTypeAllow ListTypePb = `ALLOW`

// A block list. Exclude this IP or range. IP addresses in the block list are
// excluded even if they are included in an allow list.
const ListTypeBlock ListTypePb = `BLOCK`

type LlmProxyPartnerPoweredAccountPb struct {
	BooleanVal  BooleanMessagePb `json:"boolean_val"`
	Etag        string           `json:"etag,omitempty"`
	SettingName string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LlmProxyPartnerPoweredAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LlmProxyPartnerPoweredAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LlmProxyPartnerPoweredEnforcePb struct {
	BooleanVal  BooleanMessagePb `json:"boolean_val"`
	Etag        string           `json:"etag,omitempty"`
	SettingName string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LlmProxyPartnerPoweredEnforcePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LlmProxyPartnerPoweredEnforcePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LlmProxyPartnerPoweredWorkspacePb struct {
	BooleanVal  BooleanMessagePb `json:"boolean_val"`
	Etag        string           `json:"etag,omitempty"`
	SettingName string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *LlmProxyPartnerPoweredWorkspacePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st LlmProxyPartnerPoweredWorkspacePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type MicrosoftTeamsConfigPb struct {
	Url    string `json:"url,omitempty"`
	UrlSet bool   `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *MicrosoftTeamsConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st MicrosoftTeamsConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NccAwsStableIpRulePb struct {
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

type NccAzurePrivateEndpointRulePb struct {
	ConnectionState             NccAzurePrivateEndpointRuleConnectionStatePb `json:"connection_state,omitempty"`
	CreationTime                int64                                        `json:"creation_time,omitempty"`
	Deactivated                 bool                                         `json:"deactivated,omitempty"`
	DeactivatedAt               int64                                        `json:"deactivated_at,omitempty"`
	DomainNames                 []string                                     `json:"domain_names,omitempty"`
	EndpointName                string                                       `json:"endpoint_name,omitempty"`
	GroupId                     string                                       `json:"group_id,omitempty"`
	NetworkConnectivityConfigId string                                       `json:"network_connectivity_config_id,omitempty"`
	ResourceId                  string                                       `json:"resource_id,omitempty"`
	RuleId                      string                                       `json:"rule_id,omitempty"`
	UpdatedTime                 int64                                        `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NccAzurePrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NccAzurePrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NccAzurePrivateEndpointRuleConnectionStatePb string

const NccAzurePrivateEndpointRuleConnectionStateDisconnected NccAzurePrivateEndpointRuleConnectionStatePb = `DISCONNECTED`

const NccAzurePrivateEndpointRuleConnectionStateEstablished NccAzurePrivateEndpointRuleConnectionStatePb = `ESTABLISHED`

const NccAzurePrivateEndpointRuleConnectionStateExpired NccAzurePrivateEndpointRuleConnectionStatePb = `EXPIRED`

const NccAzurePrivateEndpointRuleConnectionStateInit NccAzurePrivateEndpointRuleConnectionStatePb = `INIT`

const NccAzurePrivateEndpointRuleConnectionStatePending NccAzurePrivateEndpointRuleConnectionStatePb = `PENDING`

const NccAzurePrivateEndpointRuleConnectionStateRejected NccAzurePrivateEndpointRuleConnectionStatePb = `REJECTED`

type NccAzureServiceEndpointRulePb struct {
	Subnets        []string               `json:"subnets,omitempty"`
	TargetRegion   string                 `json:"target_region,omitempty"`
	TargetServices []EgressResourceTypePb `json:"target_services,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NccAzureServiceEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NccAzureServiceEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NccEgressConfigPb struct {
	DefaultRules *NccEgressDefaultRulesPb `json:"default_rules,omitempty"`
	TargetRules  *NccEgressTargetRulesPb  `json:"target_rules,omitempty"`
}

type NccEgressDefaultRulesPb struct {
	AwsStableIpRule          *NccAwsStableIpRulePb          `json:"aws_stable_ip_rule,omitempty"`
	AzureServiceEndpointRule *NccAzureServiceEndpointRulePb `json:"azure_service_endpoint_rule,omitempty"`
}

type NccEgressTargetRulesPb struct {
	AwsPrivateEndpointRules   []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb `json:"aws_private_endpoint_rules,omitempty"`
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRulePb                                   `json:"azure_private_endpoint_rules,omitempty"`
}

type NccPrivateEndpointRulePb struct {
	AccountId                   string                                             `json:"account_id,omitempty"`
	ConnectionState             NccPrivateEndpointRulePrivateLinkConnectionStatePb `json:"connection_state,omitempty"`
	CreationTime                int64                                              `json:"creation_time,omitempty"`
	Deactivated                 bool                                               `json:"deactivated,omitempty"`
	DeactivatedAt               int64                                              `json:"deactivated_at,omitempty"`
	DomainNames                 []string                                           `json:"domain_names,omitempty"`
	Enabled                     bool                                               `json:"enabled,omitempty"`
	EndpointName                string                                             `json:"endpoint_name,omitempty"`
	EndpointService             string                                             `json:"endpoint_service,omitempty"`
	GroupId                     string                                             `json:"group_id,omitempty"`
	NetworkConnectivityConfigId string                                             `json:"network_connectivity_config_id,omitempty"`
	ResourceId                  string                                             `json:"resource_id,omitempty"`
	ResourceNames               []string                                           `json:"resource_names,omitempty"`
	RuleId                      string                                             `json:"rule_id,omitempty"`
	UpdatedTime                 int64                                              `json:"updated_time,omitempty"`
	VpcEndpointId               string                                             `json:"vpc_endpoint_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NccPrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NccPrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NccPrivateEndpointRulePrivateLinkConnectionStatePb string

const NccPrivateEndpointRulePrivateLinkConnectionStateDisconnected NccPrivateEndpointRulePrivateLinkConnectionStatePb = `DISCONNECTED`

const NccPrivateEndpointRulePrivateLinkConnectionStateEstablished NccPrivateEndpointRulePrivateLinkConnectionStatePb = `ESTABLISHED`

const NccPrivateEndpointRulePrivateLinkConnectionStateExpired NccPrivateEndpointRulePrivateLinkConnectionStatePb = `EXPIRED`

const NccPrivateEndpointRulePrivateLinkConnectionStatePending NccPrivateEndpointRulePrivateLinkConnectionStatePb = `PENDING`

const NccPrivateEndpointRulePrivateLinkConnectionStateRejected NccPrivateEndpointRulePrivateLinkConnectionStatePb = `REJECTED`

type NetworkConnectivityConfigurationPb struct {
	AccountId                   string             `json:"account_id,omitempty"`
	CreationTime                int64              `json:"creation_time,omitempty"`
	EgressConfig                *NccEgressConfigPb `json:"egress_config,omitempty"`
	Name                        string             `json:"name,omitempty"`
	NetworkConnectivityConfigId string             `json:"network_connectivity_config_id,omitempty"`
	Region                      string             `json:"region,omitempty"`
	UpdatedTime                 int64              `json:"updated_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NetworkConnectivityConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NetworkConnectivityConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NetworkPolicyEgressPb struct {
	NetworkAccess *EgressNetworkPolicyNetworkAccessPolicyPb `json:"network_access,omitempty"`
}

type NotificationDestinationPb struct {
	Config          *ConfigPb         `json:"config,omitempty"`
	DestinationType DestinationTypePb `json:"destination_type,omitempty"`
	DisplayName     string            `json:"display_name,omitempty"`
	Id              string            `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NotificationDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NotificationDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PagerdutyConfigPb struct {
	IntegrationKey    string `json:"integration_key,omitempty"`
	IntegrationKeySet bool   `json:"integration_key_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PagerdutyConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PagerdutyConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PartitionIdPb struct {
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PartitionIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PartitionIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PersonalComputeMessagePb struct {
	Value PersonalComputeMessageEnumPb `json:"value"`
}

type PersonalComputeMessageEnumPb string

const PersonalComputeMessageEnumDelegate PersonalComputeMessageEnumPb = `DELEGATE`

const PersonalComputeMessageEnumOn PersonalComputeMessageEnumPb = `ON`

type PersonalComputeSettingPb struct {
	Etag            string                   `json:"etag,omitempty"`
	PersonalCompute PersonalComputeMessagePb `json:"personal_compute"`
	SettingName     string                   `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PersonalComputeSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PersonalComputeSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PublicTokenInfoPb struct {
	Comment      string `json:"comment,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty"`
	TokenId      string `json:"token_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PublicTokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PublicTokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ReplaceIpAccessListPb struct {
	Enabled        bool       `json:"enabled"`
	IpAccessListId string     `json:"-" url:"-"`
	IpAddresses    []string   `json:"ip_addresses,omitempty"`
	Label          string     `json:"label"`
	ListType       ListTypePb `json:"list_type"`
}

type RestrictWorkspaceAdminsMessagePb struct {
	Status RestrictWorkspaceAdminsMessageStatusPb `json:"status"`
}

type RestrictWorkspaceAdminsMessageStatusPb string

const RestrictWorkspaceAdminsMessageStatusAllowAll RestrictWorkspaceAdminsMessageStatusPb = `ALLOW_ALL`

const RestrictWorkspaceAdminsMessageStatusRestrictTokensAndJobRunAs RestrictWorkspaceAdminsMessageStatusPb = `RESTRICT_TOKENS_AND_JOB_RUN_AS`

type RestrictWorkspaceAdminsSettingPb struct {
	Etag                    string                           `json:"etag,omitempty"`
	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessagePb `json:"restrict_workspace_admins"`
	SettingName             string                           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RestrictWorkspaceAdminsSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RestrictWorkspaceAdminsSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RevokeTokenRequestPb struct {
	TokenId string `json:"token_id"`
}

type RevokeTokenResponsePb struct {
}

type SlackConfigPb struct {
	Url    string `json:"url,omitempty"`
	UrlSet bool   `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SlackConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SlackConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SqlResultsDownloadPb struct {
	BooleanVal  BooleanMessagePb `json:"boolean_val"`
	Etag        string           `json:"etag,omitempty"`
	SettingName string           `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SqlResultsDownloadPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SqlResultsDownloadPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StringMessagePb struct {
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StringMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StringMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenAccessControlRequestPb struct {
	GroupName            string                 `json:"group_name,omitempty"`
	PermissionLevel      TokenPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                 `json:"service_principal_name,omitempty"`
	UserName             string                 `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenAccessControlResponsePb struct {
	AllPermissions       []TokenPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string              `json:"display_name,omitempty"`
	GroupName            string              `json:"group_name,omitempty"`
	ServicePrincipalName string              `json:"service_principal_name,omitempty"`
	UserName             string              `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenInfoPb struct {
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

func (st *TokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenPermissionPb struct {
	Inherited           bool                   `json:"inherited,omitempty"`
	InheritedFromObject []string               `json:"inherited_from_object,omitempty"`
	PermissionLevel     TokenPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenPermissionLevelPb string

const TokenPermissionLevelCanUse TokenPermissionLevelPb = `CAN_USE`

type TokenPermissionsPb struct {
	AccessControlList []TokenAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                         `json:"object_id,omitempty"`
	ObjectType        string                         `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenPermissionsDescriptionPb struct {
	Description     string                 `json:"description,omitempty"`
	PermissionLevel TokenPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenPermissionsRequestPb struct {
	AccessControlList []TokenAccessControlRequestPb `json:"access_control_list,omitempty"`
}

type TokenTypePb string

const TokenTypeArclightAzureExchangeToken TokenTypePb = `ARCLIGHT_AZURE_EXCHANGE_TOKEN`

const TokenTypeArclightAzureExchangeTokenWithUserDelegationKey TokenTypePb = `ARCLIGHT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`

const TokenTypeArclightMultiTenantAzureExchangeToken TokenTypePb = `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN`

const TokenTypeArclightMultiTenantAzureExchangeTokenWithUserDelegationKey TokenTypePb = `ARCLIGHT_MULTI_TENANT_AZURE_EXCHANGE_TOKEN_WITH_USER_DELEGATION_KEY`

const TokenTypeAzureActiveDirectoryToken TokenTypePb = `AZURE_ACTIVE_DIRECTORY_TOKEN`

type UpdateAccountIpAccessEnableRequestPb struct {
	AllowMissing bool                    `json:"allow_missing"`
	FieldMask    string                  `json:"field_mask"`
	Setting      AccountIpAccessEnablePb `json:"setting"`
}

type UpdateAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
	AllowMissing bool                                        `json:"allow_missing"`
	FieldMask    string                                      `json:"field_mask"`
	Setting      AibiDashboardEmbeddingAccessPolicySettingPb `json:"setting"`
}

type UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
	AllowMissing bool                                           `json:"allow_missing"`
	FieldMask    string                                         `json:"field_mask"`
	Setting      AibiDashboardEmbeddingApprovedDomainsSettingPb `json:"setting"`
}

type UpdateAutomaticClusterUpdateSettingRequestPb struct {
	AllowMissing bool                            `json:"allow_missing"`
	FieldMask    string                          `json:"field_mask"`
	Setting      AutomaticClusterUpdateSettingPb `json:"setting"`
}

type UpdateComplianceSecurityProfileSettingRequestPb struct {
	AllowMissing bool                               `json:"allow_missing"`
	FieldMask    string                             `json:"field_mask"`
	Setting      ComplianceSecurityProfileSettingPb `json:"setting"`
}

type UpdateCspEnablementAccountSettingRequestPb struct {
	AllowMissing bool                          `json:"allow_missing"`
	FieldMask    string                        `json:"field_mask"`
	Setting      CspEnablementAccountSettingPb `json:"setting"`
}

type UpdateDashboardEmailSubscriptionsRequestPb struct {
	AllowMissing bool                          `json:"allow_missing"`
	FieldMask    string                        `json:"field_mask"`
	Setting      DashboardEmailSubscriptionsPb `json:"setting"`
}

type UpdateDefaultNamespaceSettingRequestPb struct {
	AllowMissing bool                      `json:"allow_missing"`
	FieldMask    string                    `json:"field_mask"`
	Setting      DefaultNamespaceSettingPb `json:"setting"`
}

type UpdateDefaultWarehouseIdRequestPb struct {
	AllowMissing bool                 `json:"allow_missing"`
	FieldMask    string               `json:"field_mask"`
	Setting      DefaultWarehouseIdPb `json:"setting"`
}

type UpdateDisableLegacyAccessRequestPb struct {
	AllowMissing bool                  `json:"allow_missing"`
	FieldMask    string                `json:"field_mask"`
	Setting      DisableLegacyAccessPb `json:"setting"`
}

type UpdateDisableLegacyDbfsRequestPb struct {
	AllowMissing bool                `json:"allow_missing"`
	FieldMask    string              `json:"field_mask"`
	Setting      DisableLegacyDbfsPb `json:"setting"`
}

type UpdateDisableLegacyFeaturesRequestPb struct {
	AllowMissing bool                    `json:"allow_missing"`
	FieldMask    string                  `json:"field_mask"`
	Setting      DisableLegacyFeaturesPb `json:"setting"`
}

type UpdateEnableExportNotebookRequestPb struct {
	AllowMissing bool                   `json:"allow_missing"`
	FieldMask    string                 `json:"field_mask"`
	Setting      EnableExportNotebookPb `json:"setting"`
}

type UpdateEnableNotebookTableClipboardRequestPb struct {
	AllowMissing bool                           `json:"allow_missing"`
	FieldMask    string                         `json:"field_mask"`
	Setting      EnableNotebookTableClipboardPb `json:"setting"`
}

type UpdateEnableResultsDownloadingRequestPb struct {
	AllowMissing bool                       `json:"allow_missing"`
	FieldMask    string                     `json:"field_mask"`
	Setting      EnableResultsDownloadingPb `json:"setting"`
}

type UpdateEnhancedSecurityMonitoringSettingRequestPb struct {
	AllowMissing bool                                `json:"allow_missing"`
	FieldMask    string                              `json:"field_mask"`
	Setting      EnhancedSecurityMonitoringSettingPb `json:"setting"`
}

type UpdateEsmEnablementAccountSettingRequestPb struct {
	AllowMissing bool                          `json:"allow_missing"`
	FieldMask    string                        `json:"field_mask"`
	Setting      EsmEnablementAccountSettingPb `json:"setting"`
}

type UpdateIpAccessListPb struct {
	Enabled        bool       `json:"enabled,omitempty"`
	IpAccessListId string     `json:"-" url:"-"`
	IpAddresses    []string   `json:"ip_addresses,omitempty"`
	Label          string     `json:"label,omitempty"`
	ListType       ListTypePb `json:"list_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateIpAccessListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateIpAccessListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateLlmProxyPartnerPoweredAccountRequestPb struct {
	AllowMissing bool                            `json:"allow_missing"`
	FieldMask    string                          `json:"field_mask"`
	Setting      LlmProxyPartnerPoweredAccountPb `json:"setting"`
}

type UpdateLlmProxyPartnerPoweredEnforceRequestPb struct {
	AllowMissing bool                            `json:"allow_missing"`
	FieldMask    string                          `json:"field_mask"`
	Setting      LlmProxyPartnerPoweredEnforcePb `json:"setting"`
}

type UpdateLlmProxyPartnerPoweredWorkspaceRequestPb struct {
	AllowMissing bool                              `json:"allow_missing"`
	FieldMask    string                            `json:"field_mask"`
	Setting      LlmProxyPartnerPoweredWorkspacePb `json:"setting"`
}

type UpdateNccPrivateEndpointRuleRequestPb struct {
	NetworkConnectivityConfigId string                      `json:"-" url:"-"`
	PrivateEndpointRule         UpdatePrivateEndpointRulePb `json:"private_endpoint_rule"`
	PrivateEndpointRuleId       string                      `json:"-" url:"-"`
	UpdateMask                  string                      `json:"-" url:"update_mask"`
}

type UpdateNetworkPolicyRequestPb struct {
	NetworkPolicy   AccountNetworkPolicyPb `json:"network_policy"`
	NetworkPolicyId string                 `json:"-" url:"-"`
}

type UpdateNotificationDestinationRequestPb struct {
	Config      *ConfigPb `json:"config,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	Id          string    `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateNotificationDestinationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateNotificationDestinationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdatePersonalComputeSettingRequestPb struct {
	AllowMissing bool                     `json:"allow_missing"`
	FieldMask    string                   `json:"field_mask"`
	Setting      PersonalComputeSettingPb `json:"setting"`
}

type UpdatePrivateEndpointRulePb struct {
	DomainNames   []string `json:"domain_names,omitempty"`
	Enabled       bool     `json:"enabled,omitempty"`
	ResourceNames []string `json:"resource_names,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdatePrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdatePrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRestrictWorkspaceAdminsSettingRequestPb struct {
	AllowMissing bool                             `json:"allow_missing"`
	FieldMask    string                           `json:"field_mask"`
	Setting      RestrictWorkspaceAdminsSettingPb `json:"setting"`
}

type UpdateSqlResultsDownloadRequestPb struct {
	AllowMissing bool                 `json:"allow_missing"`
	FieldMask    string               `json:"field_mask"`
	Setting      SqlResultsDownloadPb `json:"setting"`
}

type UpdateWorkspaceNetworkOptionRequestPb struct {
	WorkspaceId            int64                    `json:"-" url:"-"`
	WorkspaceNetworkOption WorkspaceNetworkOptionPb `json:"workspace_network_option"`
}

type WorkspaceConfPb map[string]string

type WorkspaceNetworkOptionPb struct {
	NetworkPolicyId string `json:"network_policy_id,omitempty"`
	WorkspaceId     int64  `json:"workspace_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *WorkspaceNetworkOptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st WorkspaceNetworkOptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
