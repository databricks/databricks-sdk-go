// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/settings/settingspb"
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AccountIpAccessEnable) MarshalJSON() ([]byte, error) {
	pb, err := AccountIpAccessEnableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AccountIpAccessEnable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.AccountIpAccessEnablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AccountIpAccessEnableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AccountIpAccessEnableToPb(st *AccountIpAccessEnable) (*settingspb.AccountIpAccessEnablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.AccountIpAccessEnablePb{}
	acctIpAclEnablePb, err := BooleanMessageToPb(&st.AcctIpAclEnable)
	if err != nil {
		return nil, err
	}
	if acctIpAclEnablePb != nil {
		pb.AcctIpAclEnable = *acctIpAclEnablePb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AccountIpAccessEnableFromPb(pb *settingspb.AccountIpAccessEnablePb) (*AccountIpAccessEnable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountIpAccessEnable{}
	acctIpAclEnableField, err := BooleanMessageFromPb(&pb.AcctIpAclEnable)
	if err != nil {
		return nil, err
	}
	if acctIpAclEnableField != nil {
		st.AcctIpAclEnable = *acctIpAclEnableField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	NetworkPolicyId string   `json:"network_policy_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AccountNetworkPolicy) MarshalJSON() ([]byte, error) {
	pb, err := AccountNetworkPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AccountNetworkPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.AccountNetworkPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AccountNetworkPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AccountNetworkPolicyToPb(st *AccountNetworkPolicy) (*settingspb.AccountNetworkPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.AccountNetworkPolicyPb{}
	pb.AccountId = st.AccountId
	egressPb, err := NetworkPolicyEgressToPb(st.Egress)
	if err != nil {
		return nil, err
	}
	if egressPb != nil {
		pb.Egress = egressPb
	}
	pb.NetworkPolicyId = st.NetworkPolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AccountNetworkPolicyFromPb(pb *settingspb.AccountNetworkPolicyPb) (*AccountNetworkPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountNetworkPolicy{}
	st.AccountId = pb.AccountId
	egressField, err := NetworkPolicyEgressFromPb(pb.Egress)
	if err != nil {
		return nil, err
	}
	if egressField != nil {
		st.Egress = egressField
	}
	st.NetworkPolicyId = pb.NetworkPolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AibiDashboardEmbeddingAccessPolicy struct {

	// Wire name: 'access_policy_type'
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType `json:"access_policy_type"`
}

func (st AibiDashboardEmbeddingAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := AibiDashboardEmbeddingAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AibiDashboardEmbeddingAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.AibiDashboardEmbeddingAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AibiDashboardEmbeddingAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AibiDashboardEmbeddingAccessPolicyToPb(st *AibiDashboardEmbeddingAccessPolicy) (*settingspb.AibiDashboardEmbeddingAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.AibiDashboardEmbeddingAccessPolicyPb{}
	accessPolicyTypePb, err := AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeToPb(&st.AccessPolicyType)
	if err != nil {
		return nil, err
	}
	if accessPolicyTypePb != nil {
		pb.AccessPolicyType = *accessPolicyTypePb
	}

	return pb, nil
}

func AibiDashboardEmbeddingAccessPolicyFromPb(pb *settingspb.AibiDashboardEmbeddingAccessPolicyPb) (*AibiDashboardEmbeddingAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingAccessPolicy{}
	accessPolicyTypeField, err := AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeFromPb(&pb.AccessPolicyType)
	if err != nil {
		return nil, err
	}
	if accessPolicyTypeField != nil {
		st.AccessPolicyType = *accessPolicyTypeField
	}

	return st, nil
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

func AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeToPb(st *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) (*settingspb.AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb(*st)
	return &pb, nil
}

func AibiDashboardEmbeddingAccessPolicyAccessPolicyTypeFromPb(pb *settingspb.AibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb) (*AibiDashboardEmbeddingAccessPolicyAccessPolicyType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AibiDashboardEmbeddingAccessPolicyAccessPolicyType(*pb)
	return &st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AibiDashboardEmbeddingAccessPolicySetting) MarshalJSON() ([]byte, error) {
	pb, err := AibiDashboardEmbeddingAccessPolicySettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AibiDashboardEmbeddingAccessPolicySetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.AibiDashboardEmbeddingAccessPolicySettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AibiDashboardEmbeddingAccessPolicySettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AibiDashboardEmbeddingAccessPolicySettingToPb(st *AibiDashboardEmbeddingAccessPolicySetting) (*settingspb.AibiDashboardEmbeddingAccessPolicySettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.AibiDashboardEmbeddingAccessPolicySettingPb{}
	aibiDashboardEmbeddingAccessPolicyPb, err := AibiDashboardEmbeddingAccessPolicyToPb(&st.AibiDashboardEmbeddingAccessPolicy)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingAccessPolicyPb != nil {
		pb.AibiDashboardEmbeddingAccessPolicy = *aibiDashboardEmbeddingAccessPolicyPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AibiDashboardEmbeddingAccessPolicySettingFromPb(pb *settingspb.AibiDashboardEmbeddingAccessPolicySettingPb) (*AibiDashboardEmbeddingAccessPolicySetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingAccessPolicySetting{}
	aibiDashboardEmbeddingAccessPolicyField, err := AibiDashboardEmbeddingAccessPolicyFromPb(&pb.AibiDashboardEmbeddingAccessPolicy)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingAccessPolicyField != nil {
		st.AibiDashboardEmbeddingAccessPolicy = *aibiDashboardEmbeddingAccessPolicyField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type AibiDashboardEmbeddingApprovedDomains struct {

	// Wire name: 'approved_domains'
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

func (st AibiDashboardEmbeddingApprovedDomains) MarshalJSON() ([]byte, error) {
	pb, err := AibiDashboardEmbeddingApprovedDomainsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AibiDashboardEmbeddingApprovedDomains) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.AibiDashboardEmbeddingApprovedDomainsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AibiDashboardEmbeddingApprovedDomainsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AibiDashboardEmbeddingApprovedDomainsToPb(st *AibiDashboardEmbeddingApprovedDomains) (*settingspb.AibiDashboardEmbeddingApprovedDomainsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.AibiDashboardEmbeddingApprovedDomainsPb{}
	pb.ApprovedDomains = st.ApprovedDomains

	return pb, nil
}

func AibiDashboardEmbeddingApprovedDomainsFromPb(pb *settingspb.AibiDashboardEmbeddingApprovedDomainsPb) (*AibiDashboardEmbeddingApprovedDomains, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingApprovedDomains{}
	st.ApprovedDomains = pb.ApprovedDomains

	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AibiDashboardEmbeddingApprovedDomainsSetting) MarshalJSON() ([]byte, error) {
	pb, err := AibiDashboardEmbeddingApprovedDomainsSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AibiDashboardEmbeddingApprovedDomainsSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.AibiDashboardEmbeddingApprovedDomainsSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AibiDashboardEmbeddingApprovedDomainsSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AibiDashboardEmbeddingApprovedDomainsSettingToPb(st *AibiDashboardEmbeddingApprovedDomainsSetting) (*settingspb.AibiDashboardEmbeddingApprovedDomainsSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.AibiDashboardEmbeddingApprovedDomainsSettingPb{}
	aibiDashboardEmbeddingApprovedDomainsPb, err := AibiDashboardEmbeddingApprovedDomainsToPb(&st.AibiDashboardEmbeddingApprovedDomains)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingApprovedDomainsPb != nil {
		pb.AibiDashboardEmbeddingApprovedDomains = *aibiDashboardEmbeddingApprovedDomainsPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AibiDashboardEmbeddingApprovedDomainsSettingFromPb(pb *settingspb.AibiDashboardEmbeddingApprovedDomainsSettingPb) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingApprovedDomainsSetting{}
	aibiDashboardEmbeddingApprovedDomainsField, err := AibiDashboardEmbeddingApprovedDomainsFromPb(&pb.AibiDashboardEmbeddingApprovedDomains)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingApprovedDomainsField != nil {
		st.AibiDashboardEmbeddingApprovedDomains = *aibiDashboardEmbeddingApprovedDomainsField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AutomaticClusterUpdateSetting) MarshalJSON() ([]byte, error) {
	pb, err := AutomaticClusterUpdateSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AutomaticClusterUpdateSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.AutomaticClusterUpdateSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AutomaticClusterUpdateSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AutomaticClusterUpdateSettingToPb(st *AutomaticClusterUpdateSetting) (*settingspb.AutomaticClusterUpdateSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.AutomaticClusterUpdateSettingPb{}
	automaticClusterUpdateWorkspacePb, err := ClusterAutoRestartMessageToPb(&st.AutomaticClusterUpdateWorkspace)
	if err != nil {
		return nil, err
	}
	if automaticClusterUpdateWorkspacePb != nil {
		pb.AutomaticClusterUpdateWorkspace = *automaticClusterUpdateWorkspacePb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AutomaticClusterUpdateSettingFromPb(pb *settingspb.AutomaticClusterUpdateSettingPb) (*AutomaticClusterUpdateSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutomaticClusterUpdateSetting{}
	automaticClusterUpdateWorkspaceField, err := ClusterAutoRestartMessageFromPb(&pb.AutomaticClusterUpdateWorkspace)
	if err != nil {
		return nil, err
	}
	if automaticClusterUpdateWorkspaceField != nil {
		st.AutomaticClusterUpdateWorkspace = *automaticClusterUpdateWorkspaceField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type BooleanMessage struct {

	// Wire name: 'value'
	Value           bool     `json:"value,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st BooleanMessage) MarshalJSON() ([]byte, error) {
	pb, err := BooleanMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *BooleanMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.BooleanMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := BooleanMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func BooleanMessageToPb(st *BooleanMessage) (*settingspb.BooleanMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.BooleanMessagePb{}
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func BooleanMessageFromPb(pb *settingspb.BooleanMessagePb) (*BooleanMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BooleanMessage{}
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	RestartEvenIfNoUpdatesAvailable bool     `json:"restart_even_if_no_updates_available,omitempty"`
	ForceSendFields                 []string `json:"-" tf:"-"`
}

func (st ClusterAutoRestartMessage) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAutoRestartMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAutoRestartMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ClusterAutoRestartMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAutoRestartMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAutoRestartMessageToPb(st *ClusterAutoRestartMessage) (*settingspb.ClusterAutoRestartMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ClusterAutoRestartMessagePb{}
	pb.CanToggle = st.CanToggle
	pb.Enabled = st.Enabled
	enablementDetailsPb, err := ClusterAutoRestartMessageEnablementDetailsToPb(st.EnablementDetails)
	if err != nil {
		return nil, err
	}
	if enablementDetailsPb != nil {
		pb.EnablementDetails = enablementDetailsPb
	}
	maintenanceWindowPb, err := ClusterAutoRestartMessageMaintenanceWindowToPb(st.MaintenanceWindow)
	if err != nil {
		return nil, err
	}
	if maintenanceWindowPb != nil {
		pb.MaintenanceWindow = maintenanceWindowPb
	}
	pb.RestartEvenIfNoUpdatesAvailable = st.RestartEvenIfNoUpdatesAvailable

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterAutoRestartMessageFromPb(pb *settingspb.ClusterAutoRestartMessagePb) (*ClusterAutoRestartMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessage{}
	st.CanToggle = pb.CanToggle
	st.Enabled = pb.Enabled
	enablementDetailsField, err := ClusterAutoRestartMessageEnablementDetailsFromPb(pb.EnablementDetails)
	if err != nil {
		return nil, err
	}
	if enablementDetailsField != nil {
		st.EnablementDetails = enablementDetailsField
	}
	maintenanceWindowField, err := ClusterAutoRestartMessageMaintenanceWindowFromPb(pb.MaintenanceWindow)
	if err != nil {
		return nil, err
	}
	if maintenanceWindowField != nil {
		st.MaintenanceWindow = maintenanceWindowField
	}
	st.RestartEvenIfNoUpdatesAvailable = pb.RestartEvenIfNoUpdatesAvailable

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	UnavailableForNonEnterpriseTier bool     `json:"unavailable_for_non_enterprise_tier,omitempty"`
	ForceSendFields                 []string `json:"-" tf:"-"`
}

func (st ClusterAutoRestartMessageEnablementDetails) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAutoRestartMessageEnablementDetailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAutoRestartMessageEnablementDetails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ClusterAutoRestartMessageEnablementDetailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAutoRestartMessageEnablementDetailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAutoRestartMessageEnablementDetailsToPb(st *ClusterAutoRestartMessageEnablementDetails) (*settingspb.ClusterAutoRestartMessageEnablementDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ClusterAutoRestartMessageEnablementDetailsPb{}
	pb.ForcedForComplianceMode = st.ForcedForComplianceMode
	pb.UnavailableForDisabledEntitlement = st.UnavailableForDisabledEntitlement
	pb.UnavailableForNonEnterpriseTier = st.UnavailableForNonEnterpriseTier

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterAutoRestartMessageEnablementDetailsFromPb(pb *settingspb.ClusterAutoRestartMessageEnablementDetailsPb) (*ClusterAutoRestartMessageEnablementDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageEnablementDetails{}
	st.ForcedForComplianceMode = pb.ForcedForComplianceMode
	st.UnavailableForDisabledEntitlement = pb.UnavailableForDisabledEntitlement
	st.UnavailableForNonEnterpriseTier = pb.UnavailableForNonEnterpriseTier

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ClusterAutoRestartMessageMaintenanceWindow struct {

	// Wire name: 'week_day_based_schedule'
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule `json:"week_day_based_schedule,omitempty"`
}

func (st ClusterAutoRestartMessageMaintenanceWindow) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAutoRestartMessageMaintenanceWindowToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAutoRestartMessageMaintenanceWindow) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ClusterAutoRestartMessageMaintenanceWindowPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAutoRestartMessageMaintenanceWindowFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAutoRestartMessageMaintenanceWindowToPb(st *ClusterAutoRestartMessageMaintenanceWindow) (*settingspb.ClusterAutoRestartMessageMaintenanceWindowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ClusterAutoRestartMessageMaintenanceWindowPb{}
	weekDayBasedSchedulePb, err := ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(st.WeekDayBasedSchedule)
	if err != nil {
		return nil, err
	}
	if weekDayBasedSchedulePb != nil {
		pb.WeekDayBasedSchedule = weekDayBasedSchedulePb
	}

	return pb, nil
}

func ClusterAutoRestartMessageMaintenanceWindowFromPb(pb *settingspb.ClusterAutoRestartMessageMaintenanceWindowPb) (*ClusterAutoRestartMessageMaintenanceWindow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindow{}
	weekDayBasedScheduleField, err := ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleFromPb(pb.WeekDayBasedSchedule)
	if err != nil {
		return nil, err
	}
	if weekDayBasedScheduleField != nil {
		st.WeekDayBasedSchedule = weekDayBasedScheduleField
	}

	return st, nil
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

func ClusterAutoRestartMessageMaintenanceWindowDayOfWeekToPb(st *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) (*settingspb.ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb(*st)
	return &pb, nil
}

func ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFromPb(pb *settingspb.ClusterAutoRestartMessageMaintenanceWindowDayOfWeekPb) (*ClusterAutoRestartMessageMaintenanceWindowDayOfWeek, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterAutoRestartMessageMaintenanceWindowDayOfWeek(*pb)
	return &st, nil
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {

	// Wire name: 'day_of_week'
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek `json:"day_of_week,omitempty"`

	// Wire name: 'frequency'
	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `json:"frequency,omitempty"`

	// Wire name: 'window_start_time'
	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime `json:"window_start_time,omitempty"`
}

func (st ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(st *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) (*settingspb.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb{}
	dayOfWeekPb, err := ClusterAutoRestartMessageMaintenanceWindowDayOfWeekToPb(&st.DayOfWeek)
	if err != nil {
		return nil, err
	}
	if dayOfWeekPb != nil {
		pb.DayOfWeek = *dayOfWeekPb
	}
	frequencyPb, err := ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyToPb(&st.Frequency)
	if err != nil {
		return nil, err
	}
	if frequencyPb != nil {
		pb.Frequency = *frequencyPb
	}
	windowStartTimePb, err := ClusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(st.WindowStartTime)
	if err != nil {
		return nil, err
	}
	if windowStartTimePb != nil {
		pb.WindowStartTime = windowStartTimePb
	}

	return pb, nil
}

func ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleFromPb(pb *settingspb.ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb) (*ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}
	dayOfWeekField, err := ClusterAutoRestartMessageMaintenanceWindowDayOfWeekFromPb(&pb.DayOfWeek)
	if err != nil {
		return nil, err
	}
	if dayOfWeekField != nil {
		st.DayOfWeek = *dayOfWeekField
	}
	frequencyField, err := ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFromPb(&pb.Frequency)
	if err != nil {
		return nil, err
	}
	if frequencyField != nil {
		st.Frequency = *frequencyField
	}
	windowStartTimeField, err := ClusterAutoRestartMessageMaintenanceWindowWindowStartTimeFromPb(pb.WindowStartTime)
	if err != nil {
		return nil, err
	}
	if windowStartTimeField != nil {
		st.WindowStartTime = windowStartTimeField
	}

	return st, nil
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

func ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyToPb(st *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) (*settingspb.ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb(*st)
	return &pb, nil
}

func ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFromPb(pb *settingspb.ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb) (*ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency(*pb)
	return &st, nil
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {

	// Wire name: 'hours'
	Hours int `json:"hours,omitempty"`

	// Wire name: 'minutes'
	Minutes         int      `json:"minutes,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) MarshalJSON() ([]byte, error) {
	pb, err := ClusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ClusterAutoRestartMessageMaintenanceWindowWindowStartTimeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ClusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(st *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) (*settingspb.ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb{}
	pb.Hours = st.Hours
	pb.Minutes = st.Minutes

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ClusterAutoRestartMessageMaintenanceWindowWindowStartTimeFromPb(pb *settingspb.ClusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) (*ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}
	st.Hours = pb.Hours
	st.Minutes = pb.Minutes

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Wire name: 'compliance_standards'
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`

	// Wire name: 'is_enabled'
	IsEnabled       bool     `json:"is_enabled,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ComplianceSecurityProfile) MarshalJSON() ([]byte, error) {
	pb, err := ComplianceSecurityProfileToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ComplianceSecurityProfile) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ComplianceSecurityProfilePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ComplianceSecurityProfileFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ComplianceSecurityProfileToPb(st *ComplianceSecurityProfile) (*settingspb.ComplianceSecurityProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ComplianceSecurityProfilePb{}

	var complianceStandardsPb []settingspb.ComplianceStandardPb
	for _, item := range st.ComplianceStandards {
		itemPb, err := ComplianceStandardToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			complianceStandardsPb = append(complianceStandardsPb, *itemPb)
		}
	}
	pb.ComplianceStandards = complianceStandardsPb
	pb.IsEnabled = st.IsEnabled

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ComplianceSecurityProfileFromPb(pb *settingspb.ComplianceSecurityProfilePb) (*ComplianceSecurityProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfile{}

	var complianceStandardsField []ComplianceStandard
	for _, itemPb := range pb.ComplianceStandards {
		item, err := ComplianceStandardFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			complianceStandardsField = append(complianceStandardsField, *item)
		}
	}
	st.ComplianceStandards = complianceStandardsField
	st.IsEnabled = pb.IsEnabled

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ComplianceSecurityProfileSetting) MarshalJSON() ([]byte, error) {
	pb, err := ComplianceSecurityProfileSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ComplianceSecurityProfileSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ComplianceSecurityProfileSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ComplianceSecurityProfileSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ComplianceSecurityProfileSettingToPb(st *ComplianceSecurityProfileSetting) (*settingspb.ComplianceSecurityProfileSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ComplianceSecurityProfileSettingPb{}
	complianceSecurityProfileWorkspacePb, err := ComplianceSecurityProfileToPb(&st.ComplianceSecurityProfileWorkspace)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfileWorkspacePb != nil {
		pb.ComplianceSecurityProfileWorkspace = *complianceSecurityProfileWorkspacePb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ComplianceSecurityProfileSettingFromPb(pb *settingspb.ComplianceSecurityProfileSettingPb) (*ComplianceSecurityProfileSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfileSetting{}
	complianceSecurityProfileWorkspaceField, err := ComplianceSecurityProfileFromPb(&pb.ComplianceSecurityProfileWorkspace)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfileWorkspaceField != nil {
		st.ComplianceSecurityProfileWorkspace = *complianceSecurityProfileWorkspaceField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Compliance stardard for SHIELD customers
type ComplianceStandard string

const ComplianceStandardCanadaProtectedB ComplianceStandard = `CANADA_PROTECTED_B`

const ComplianceStandardCyberEssentialPlus ComplianceStandard = `CYBER_ESSENTIAL_PLUS`

const ComplianceStandardFedrampHigh ComplianceStandard = `FEDRAMP_HIGH`

const ComplianceStandardFedrampIl5 ComplianceStandard = `FEDRAMP_IL5`

const ComplianceStandardFedrampModerate ComplianceStandard = `FEDRAMP_MODERATE`

const ComplianceStandardGermanyC5 ComplianceStandard = `GERMANY_C5`

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
	case `CANADA_PROTECTED_B`, `CYBER_ESSENTIAL_PLUS`, `FEDRAMP_HIGH`, `FEDRAMP_IL5`, `FEDRAMP_MODERATE`, `GERMANY_C5`, `HIPAA`, `HITRUST`, `IRAP_PROTECTED`, `ISMAP`, `ITAR_EAR`, `K_FSI`, `NONE`, `PCI_DSS`:
		*f = ComplianceStandard(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANADA_PROTECTED_B", "CYBER_ESSENTIAL_PLUS", "FEDRAMP_HIGH", "FEDRAMP_IL5", "FEDRAMP_MODERATE", "GERMANY_C5", "HIPAA", "HITRUST", "IRAP_PROTECTED", "ISMAP", "ITAR_EAR", "K_FSI", "NONE", "PCI_DSS"`, v)
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
		ComplianceStandardGermanyC5,
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

func ComplianceStandardToPb(st *ComplianceStandard) (*settingspb.ComplianceStandardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.ComplianceStandardPb(*st)
	return &pb, nil
}

func ComplianceStandardFromPb(pb *settingspb.ComplianceStandardPb) (*ComplianceStandard, error) {
	if pb == nil {
		return nil, nil
	}
	st := ComplianceStandard(*pb)
	return &st, nil
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

func (st Config) MarshalJSON() ([]byte, error) {
	pb, err := ConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Config) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ConfigToPb(st *Config) (*settingspb.ConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ConfigPb{}
	emailPb, err := EmailConfigToPb(st.Email)
	if err != nil {
		return nil, err
	}
	if emailPb != nil {
		pb.Email = emailPb
	}
	genericWebhookPb, err := GenericWebhookConfigToPb(st.GenericWebhook)
	if err != nil {
		return nil, err
	}
	if genericWebhookPb != nil {
		pb.GenericWebhook = genericWebhookPb
	}
	microsoftTeamsPb, err := MicrosoftTeamsConfigToPb(st.MicrosoftTeams)
	if err != nil {
		return nil, err
	}
	if microsoftTeamsPb != nil {
		pb.MicrosoftTeams = microsoftTeamsPb
	}
	pagerdutyPb, err := PagerdutyConfigToPb(st.Pagerduty)
	if err != nil {
		return nil, err
	}
	if pagerdutyPb != nil {
		pb.Pagerduty = pagerdutyPb
	}
	slackPb, err := SlackConfigToPb(st.Slack)
	if err != nil {
		return nil, err
	}
	if slackPb != nil {
		pb.Slack = slackPb
	}

	return pb, nil
}

func ConfigFromPb(pb *settingspb.ConfigPb) (*Config, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Config{}
	emailField, err := EmailConfigFromPb(pb.Email)
	if err != nil {
		return nil, err
	}
	if emailField != nil {
		st.Email = emailField
	}
	genericWebhookField, err := GenericWebhookConfigFromPb(pb.GenericWebhook)
	if err != nil {
		return nil, err
	}
	if genericWebhookField != nil {
		st.GenericWebhook = genericWebhookField
	}
	microsoftTeamsField, err := MicrosoftTeamsConfigFromPb(pb.MicrosoftTeams)
	if err != nil {
		return nil, err
	}
	if microsoftTeamsField != nil {
		st.MicrosoftTeams = microsoftTeamsField
	}
	pagerdutyField, err := PagerdutyConfigFromPb(pb.Pagerduty)
	if err != nil {
		return nil, err
	}
	if pagerdutyField != nil {
		st.Pagerduty = pagerdutyField
	}
	slackField, err := SlackConfigFromPb(pb.Slack)
	if err != nil {
		return nil, err
	}
	if slackField != nil {
		st.Slack = slackField
	}

	return st, nil
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

func (st CreateIpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := CreateIpAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateIpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateIpAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateIpAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateIpAccessListToPb(st *CreateIpAccessList) (*settingspb.CreateIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateIpAccessListPb{}
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	listTypePb, err := ListTypeToPb(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}

	return pb, nil
}

func CreateIpAccessListFromPb(pb *settingspb.CreateIpAccessListPb) (*CreateIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateIpAccessList{}
	st.IpAddresses = pb.IpAddresses
	st.Label = pb.Label
	listTypeField, err := ListTypeFromPb(&pb.ListType)
	if err != nil {
		return nil, err
	}
	if listTypeField != nil {
		st.ListType = *listTypeField
	}

	return st, nil
}

// An IP access list was successfully created.
type CreateIpAccessListResponse struct {

	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (st CreateIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateIpAccessListResponseToPb(st *CreateIpAccessListResponse) (*settingspb.CreateIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateIpAccessListResponsePb{}
	ipAccessListPb, err := IpAccessListInfoToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	return pb, nil
}

func CreateIpAccessListResponseFromPb(pb *settingspb.CreateIpAccessListResponsePb) (*CreateIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateIpAccessListResponse{}
	ipAccessListField, err := IpAccessListInfoFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}

	return st, nil
}

type CreateNetworkConnectivityConfigRequest struct {

	// Wire name: 'network_connectivity_config'
	NetworkConnectivityConfig CreateNetworkConnectivityConfiguration `json:"network_connectivity_config"`
}

func (st CreateNetworkConnectivityConfigRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateNetworkConnectivityConfigRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateNetworkConnectivityConfigRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateNetworkConnectivityConfigRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateNetworkConnectivityConfigRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateNetworkConnectivityConfigRequestToPb(st *CreateNetworkConnectivityConfigRequest) (*settingspb.CreateNetworkConnectivityConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateNetworkConnectivityConfigRequestPb{}
	networkConnectivityConfigPb, err := CreateNetworkConnectivityConfigurationToPb(&st.NetworkConnectivityConfig)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigPb != nil {
		pb.NetworkConnectivityConfig = *networkConnectivityConfigPb
	}

	return pb, nil
}

func CreateNetworkConnectivityConfigRequestFromPb(pb *settingspb.CreateNetworkConnectivityConfigRequestPb) (*CreateNetworkConnectivityConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkConnectivityConfigRequest{}
	networkConnectivityConfigField, err := CreateNetworkConnectivityConfigurationFromPb(&pb.NetworkConnectivityConfig)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigField != nil {
		st.NetworkConnectivityConfig = *networkConnectivityConfigField
	}

	return st, nil
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

func (st CreateNetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := CreateNetworkConnectivityConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateNetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateNetworkConnectivityConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateNetworkConnectivityConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateNetworkConnectivityConfigurationToPb(st *CreateNetworkConnectivityConfiguration) (*settingspb.CreateNetworkConnectivityConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateNetworkConnectivityConfigurationPb{}
	pb.Name = st.Name
	pb.Region = st.Region

	return pb, nil
}

func CreateNetworkConnectivityConfigurationFromPb(pb *settingspb.CreateNetworkConnectivityConfigurationPb) (*CreateNetworkConnectivityConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkConnectivityConfiguration{}
	st.Name = pb.Name
	st.Region = pb.Region

	return st, nil
}

type CreateNetworkPolicyRequest struct {
	// Network policy configuration details.
	// Wire name: 'network_policy'
	NetworkPolicy AccountNetworkPolicy `json:"network_policy"`
}

func (st CreateNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateNetworkPolicyRequestToPb(st *CreateNetworkPolicyRequest) (*settingspb.CreateNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateNetworkPolicyRequestPb{}
	networkPolicyPb, err := AccountNetworkPolicyToPb(&st.NetworkPolicy)
	if err != nil {
		return nil, err
	}
	if networkPolicyPb != nil {
		pb.NetworkPolicy = *networkPolicyPb
	}

	return pb, nil
}

func CreateNetworkPolicyRequestFromPb(pb *settingspb.CreateNetworkPolicyRequestPb) (*CreateNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkPolicyRequest{}
	networkPolicyField, err := AccountNetworkPolicyFromPb(&pb.NetworkPolicy)
	if err != nil {
		return nil, err
	}
	if networkPolicyField != nil {
		st.NetworkPolicy = *networkPolicyField
	}

	return st, nil
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	// Wire name: 'config'
	Config *Config `json:"config,omitempty"`
	// The display name for the notification destination.
	// Wire name: 'display_name'
	DisplayName     string   `json:"display_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateNotificationDestinationRequestToPb(st *CreateNotificationDestinationRequest) (*settingspb.CreateNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateNotificationDestinationRequestPb{}
	configPb, err := ConfigToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}
	pb.DisplayName = st.DisplayName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateNotificationDestinationRequestFromPb(pb *settingspb.CreateNotificationDestinationRequestPb) (*CreateNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNotificationDestinationRequest{}
	configField, err := ConfigFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	st.DisplayName = pb.DisplayName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	LifetimeSeconds int64    `json:"lifetime_seconds,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateOboTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateOboTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateOboTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateOboTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateOboTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateOboTokenRequestToPb(st *CreateOboTokenRequest) (*settingspb.CreateOboTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateOboTokenRequestPb{}
	pb.ApplicationId = st.ApplicationId
	pb.Comment = st.Comment
	pb.LifetimeSeconds = st.LifetimeSeconds

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateOboTokenRequestFromPb(pb *settingspb.CreateOboTokenRequestPb) (*CreateOboTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOboTokenRequest{}
	st.ApplicationId = pb.ApplicationId
	st.Comment = pb.Comment
	st.LifetimeSeconds = pb.LifetimeSeconds

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {

	// Wire name: 'token_info'
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
	// Value of the token.
	// Wire name: 'token_value'
	TokenValue      string   `json:"token_value,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateOboTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateOboTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateOboTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateOboTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateOboTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateOboTokenResponseToPb(st *CreateOboTokenResponse) (*settingspb.CreateOboTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateOboTokenResponsePb{}
	tokenInfoPb, err := TokenInfoToPb(st.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoPb != nil {
		pb.TokenInfo = tokenInfoPb
	}
	pb.TokenValue = st.TokenValue

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateOboTokenResponseFromPb(pb *settingspb.CreateOboTokenResponsePb) (*CreateOboTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOboTokenResponse{}
	tokenInfoField, err := TokenInfoFromPb(pb.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoField != nil {
		st.TokenInfo = tokenInfoField
	}
	st.TokenValue = pb.TokenValue

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	ResourceNames   []string `json:"resource_names,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreatePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := CreatePrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreatePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreatePrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreatePrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreatePrivateEndpointRuleToPb(st *CreatePrivateEndpointRule) (*settingspb.CreatePrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreatePrivateEndpointRulePb{}
	pb.DomainNames = st.DomainNames
	pb.EndpointService = st.EndpointService
	pb.GroupId = st.GroupId
	pb.ResourceId = st.ResourceId
	pb.ResourceNames = st.ResourceNames

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreatePrivateEndpointRuleFromPb(pb *settingspb.CreatePrivateEndpointRulePb) (*CreatePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePrivateEndpointRule{}
	st.DomainNames = pb.DomainNames
	st.EndpointService = pb.EndpointService
	st.GroupId = pb.GroupId
	st.ResourceId = pb.ResourceId
	st.ResourceNames = pb.ResourceNames

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreatePrivateEndpointRuleRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`

	// Wire name: 'private_endpoint_rule'
	PrivateEndpointRule CreatePrivateEndpointRule `json:"private_endpoint_rule"`
}

func (st CreatePrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreatePrivateEndpointRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreatePrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreatePrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreatePrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreatePrivateEndpointRuleRequestToPb(st *CreatePrivateEndpointRuleRequest) (*settingspb.CreatePrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreatePrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	privateEndpointRulePb, err := CreatePrivateEndpointRuleToPb(&st.PrivateEndpointRule)
	if err != nil {
		return nil, err
	}
	if privateEndpointRulePb != nil {
		pb.PrivateEndpointRule = *privateEndpointRulePb
	}

	return pb, nil
}

func CreatePrivateEndpointRuleRequestFromPb(pb *settingspb.CreatePrivateEndpointRuleRequestPb) (*CreatePrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	privateEndpointRuleField, err := CreatePrivateEndpointRuleFromPb(&pb.PrivateEndpointRule)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleField != nil {
		st.PrivateEndpointRule = *privateEndpointRuleField
	}

	return st, nil
}

type CreateTokenRequest struct {
	// Optional description to attach to the token.
	// Wire name: 'comment'
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	// Wire name: 'lifetime_seconds'
	LifetimeSeconds int64    `json:"lifetime_seconds,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateTokenRequestToPb(st *CreateTokenRequest) (*settingspb.CreateTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateTokenRequestPb{}
	pb.Comment = st.Comment
	pb.LifetimeSeconds = st.LifetimeSeconds

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateTokenRequestFromPb(pb *settingspb.CreateTokenRequestPb) (*CreateTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTokenRequest{}
	st.Comment = pb.Comment
	st.LifetimeSeconds = pb.LifetimeSeconds

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateTokenResponse struct {
	// The information for the new token.
	// Wire name: 'token_info'
	TokenInfo *PublicTokenInfo `json:"token_info,omitempty"`
	// The value of the new token.
	// Wire name: 'token_value'
	TokenValue      string   `json:"token_value,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CreateTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateTokenResponseToPb(st *CreateTokenResponse) (*settingspb.CreateTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CreateTokenResponsePb{}
	tokenInfoPb, err := PublicTokenInfoToPb(st.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoPb != nil {
		pb.TokenInfo = tokenInfoPb
	}
	pb.TokenValue = st.TokenValue

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateTokenResponseFromPb(pb *settingspb.CreateTokenResponsePb) (*CreateTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTokenResponse{}
	tokenInfoField, err := PublicTokenInfoFromPb(pb.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoField != nil {
		st.TokenInfo = tokenInfoField
	}
	st.TokenValue = pb.TokenValue

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	// Wire name: 'compliance_standards'
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	// Enforced = it cannot be overriden at workspace level.
	// Wire name: 'is_enforced'
	IsEnforced      bool     `json:"is_enforced,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CspEnablementAccount) MarshalJSON() ([]byte, error) {
	pb, err := CspEnablementAccountToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CspEnablementAccount) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CspEnablementAccountPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CspEnablementAccountFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CspEnablementAccountToPb(st *CspEnablementAccount) (*settingspb.CspEnablementAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CspEnablementAccountPb{}

	var complianceStandardsPb []settingspb.ComplianceStandardPb
	for _, item := range st.ComplianceStandards {
		itemPb, err := ComplianceStandardToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			complianceStandardsPb = append(complianceStandardsPb, *itemPb)
		}
	}
	pb.ComplianceStandards = complianceStandardsPb
	pb.IsEnforced = st.IsEnforced

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CspEnablementAccountFromPb(pb *settingspb.CspEnablementAccountPb) (*CspEnablementAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CspEnablementAccount{}

	var complianceStandardsField []ComplianceStandard
	for _, itemPb := range pb.ComplianceStandards {
		item, err := ComplianceStandardFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			complianceStandardsField = append(complianceStandardsField, *item)
		}
	}
	st.ComplianceStandards = complianceStandardsField
	st.IsEnforced = pb.IsEnforced

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CspEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	pb, err := CspEnablementAccountSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CspEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CspEnablementAccountSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CspEnablementAccountSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CspEnablementAccountSettingToPb(st *CspEnablementAccountSetting) (*settingspb.CspEnablementAccountSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CspEnablementAccountSettingPb{}
	cspEnablementAccountPb, err := CspEnablementAccountToPb(&st.CspEnablementAccount)
	if err != nil {
		return nil, err
	}
	if cspEnablementAccountPb != nil {
		pb.CspEnablementAccount = *cspEnablementAccountPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CspEnablementAccountSettingFromPb(pb *settingspb.CspEnablementAccountSettingPb) (*CspEnablementAccountSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CspEnablementAccountSetting{}
	cspEnablementAccountField, err := CspEnablementAccountFromPb(&pb.CspEnablementAccount)
	if err != nil {
		return nil, err
	}
	if cspEnablementAccountField != nil {
		st.CspEnablementAccount = *cspEnablementAccountField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	VpcEndpointId   string   `json:"vpc_endpoint_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleToPb(st *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule) (*settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb{}
	pb.AccountId = st.AccountId
	connectionStatePb, err := CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateToPb(&st.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStatePb != nil {
		pb.ConnectionState = *connectionStatePb
	}
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleFromPb(pb *settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb) (*CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule{}
	st.AccountId = pb.AccountId
	connectionStateField, err := CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateFromPb(&pb.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStateField != nil {
		st.ConnectionState = *connectionStateField
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateToPb(st *CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState) (*settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb(*st)
	return &pb, nil
}

func CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStateFromPb(pb *settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionStatePb) (*CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState, error) {
	if pb == nil {
		return nil, nil
	}
	st := CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePrivateLinkConnectionState(*pb)
	return &st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DashboardEmailSubscriptions) MarshalJSON() ([]byte, error) {
	pb, err := DashboardEmailSubscriptionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DashboardEmailSubscriptions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DashboardEmailSubscriptionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DashboardEmailSubscriptionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DashboardEmailSubscriptionsToPb(st *DashboardEmailSubscriptions) (*settingspb.DashboardEmailSubscriptionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DashboardEmailSubscriptionsPb{}
	booleanValPb, err := BooleanMessageToPb(&st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = *booleanValPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DashboardEmailSubscriptionsFromPb(pb *settingspb.DashboardEmailSubscriptionsPb) (*DashboardEmailSubscriptions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DashboardEmailSubscriptions{}
	booleanValField, err := BooleanMessageFromPb(&pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = *booleanValField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DefaultNamespaceSetting) MarshalJSON() ([]byte, error) {
	pb, err := DefaultNamespaceSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DefaultNamespaceSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DefaultNamespaceSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DefaultNamespaceSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DefaultNamespaceSettingToPb(st *DefaultNamespaceSetting) (*settingspb.DefaultNamespaceSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DefaultNamespaceSettingPb{}
	pb.Etag = st.Etag
	namespacePb, err := StringMessageToPb(&st.Namespace)
	if err != nil {
		return nil, err
	}
	if namespacePb != nil {
		pb.Namespace = *namespacePb
	}
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DefaultNamespaceSettingFromPb(pb *settingspb.DefaultNamespaceSettingPb) (*DefaultNamespaceSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DefaultNamespaceSetting{}
	st.Etag = pb.Etag
	namespaceField, err := StringMessageFromPb(&pb.Namespace)
	if err != nil {
		return nil, err
	}
	if namespaceField != nil {
		st.Namespace = *namespaceField
	}
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DefaultWarehouseId struct {
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

	// Wire name: 'string_val'
	StringVal       StringMessage `json:"string_val"`
	ForceSendFields []string      `json:"-" tf:"-"`
}

func (st DefaultWarehouseId) MarshalJSON() ([]byte, error) {
	pb, err := DefaultWarehouseIdToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DefaultWarehouseId) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DefaultWarehouseIdPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DefaultWarehouseIdFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DefaultWarehouseIdToPb(st *DefaultWarehouseId) (*settingspb.DefaultWarehouseIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DefaultWarehouseIdPb{}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName
	stringValPb, err := StringMessageToPb(&st.StringVal)
	if err != nil {
		return nil, err
	}
	if stringValPb != nil {
		pb.StringVal = *stringValPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DefaultWarehouseIdFromPb(pb *settingspb.DefaultWarehouseIdPb) (*DefaultWarehouseId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DefaultWarehouseId{}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName
	stringValField, err := StringMessageFromPb(&pb.StringVal)
	if err != nil {
		return nil, err
	}
	if stringValField != nil {
		st.StringVal = *stringValField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAccountIpAccessEnableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteAccountIpAccessEnableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAccountIpAccessEnableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAccountIpAccessEnableRequestToPb(st *DeleteAccountIpAccessEnableRequest) (*settingspb.DeleteAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteAccountIpAccessEnableRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteAccountIpAccessEnableRequestFromPb(pb *settingspb.DeleteAccountIpAccessEnableRequestPb) (*DeleteAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessEnableRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteAccountIpAccessEnableResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAccountIpAccessEnableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAccountIpAccessEnableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteAccountIpAccessEnableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAccountIpAccessEnableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAccountIpAccessEnableResponseToPb(st *DeleteAccountIpAccessEnableResponse) (*settingspb.DeleteAccountIpAccessEnableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteAccountIpAccessEnableResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteAccountIpAccessEnableResponseFromPb(pb *settingspb.DeleteAccountIpAccessEnableResponsePb) (*DeleteAccountIpAccessEnableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessEnableResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st DeleteAccountIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAccountIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteAccountIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAccountIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAccountIpAccessListRequestToPb(st *DeleteAccountIpAccessListRequest) (*settingspb.DeleteAccountIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteAccountIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

func DeleteAccountIpAccessListRequestFromPb(pb *settingspb.DeleteAccountIpAccessListRequestPb) (*DeleteAccountIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

type DeleteAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*DeleteAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingAccessPolicySettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAibiDashboardEmbeddingAccessPolicySettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAibiDashboardEmbeddingAccessPolicySettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAibiDashboardEmbeddingAccessPolicySettingResponseToPb(st *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) (*settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteAibiDashboardEmbeddingAccessPolicySettingResponseFromPb(pb *settingspb.DeleteAibiDashboardEmbeddingAccessPolicySettingResponsePb) (*DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingAccessPolicySettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponseToPb(st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) (*settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponseFromPb(pb *settingspb.DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDashboardEmailSubscriptionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDashboardEmailSubscriptionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDashboardEmailSubscriptionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDashboardEmailSubscriptionsRequestToPb(st *DeleteDashboardEmailSubscriptionsRequest) (*settingspb.DeleteDashboardEmailSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDashboardEmailSubscriptionsRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteDashboardEmailSubscriptionsRequestFromPb(pb *settingspb.DeleteDashboardEmailSubscriptionsRequestPb) (*DeleteDashboardEmailSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardEmailSubscriptionsRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteDashboardEmailSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDashboardEmailSubscriptionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDashboardEmailSubscriptionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDashboardEmailSubscriptionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDashboardEmailSubscriptionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDashboardEmailSubscriptionsResponseToPb(st *DeleteDashboardEmailSubscriptionsResponse) (*settingspb.DeleteDashboardEmailSubscriptionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDashboardEmailSubscriptionsResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteDashboardEmailSubscriptionsResponseFromPb(pb *settingspb.DeleteDashboardEmailSubscriptionsResponsePb) (*DeleteDashboardEmailSubscriptionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDashboardEmailSubscriptionsResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDefaultNamespaceSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDefaultNamespaceSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDefaultNamespaceSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDefaultNamespaceSettingRequestToPb(st *DeleteDefaultNamespaceSettingRequest) (*settingspb.DeleteDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDefaultNamespaceSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteDefaultNamespaceSettingRequestFromPb(pb *settingspb.DeleteDefaultNamespaceSettingRequestPb) (*DeleteDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultNamespaceSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteDefaultNamespaceSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDefaultNamespaceSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDefaultNamespaceSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDefaultNamespaceSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDefaultNamespaceSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDefaultNamespaceSettingResponseToPb(st *DeleteDefaultNamespaceSettingResponse) (*settingspb.DeleteDefaultNamespaceSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDefaultNamespaceSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteDefaultNamespaceSettingResponseFromPb(pb *settingspb.DeleteDefaultNamespaceSettingResponsePb) (*DeleteDefaultNamespaceSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultNamespaceSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteDefaultWarehouseIdRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDefaultWarehouseIdRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDefaultWarehouseIdRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDefaultWarehouseIdRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDefaultWarehouseIdRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDefaultWarehouseIdRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDefaultWarehouseIdRequestToPb(st *DeleteDefaultWarehouseIdRequest) (*settingspb.DeleteDefaultWarehouseIdRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDefaultWarehouseIdRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteDefaultWarehouseIdRequestFromPb(pb *settingspb.DeleteDefaultWarehouseIdRequestPb) (*DeleteDefaultWarehouseIdRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultWarehouseIdRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The etag is returned.
type DeleteDefaultWarehouseIdResponse struct {
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

func (st DeleteDefaultWarehouseIdResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDefaultWarehouseIdResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDefaultWarehouseIdResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDefaultWarehouseIdResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDefaultWarehouseIdResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDefaultWarehouseIdResponseToPb(st *DeleteDefaultWarehouseIdResponse) (*settingspb.DeleteDefaultWarehouseIdResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDefaultWarehouseIdResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteDefaultWarehouseIdResponseFromPb(pb *settingspb.DeleteDefaultWarehouseIdResponsePb) (*DeleteDefaultWarehouseIdResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultWarehouseIdResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDisableLegacyAccessRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDisableLegacyAccessRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDisableLegacyAccessRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDisableLegacyAccessRequestToPb(st *DeleteDisableLegacyAccessRequest) (*settingspb.DeleteDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDisableLegacyAccessRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteDisableLegacyAccessRequestFromPb(pb *settingspb.DeleteDisableLegacyAccessRequestPb) (*DeleteDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyAccessRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteDisableLegacyAccessResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDisableLegacyAccessResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDisableLegacyAccessResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDisableLegacyAccessResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDisableLegacyAccessResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDisableLegacyAccessResponseToPb(st *DeleteDisableLegacyAccessResponse) (*settingspb.DeleteDisableLegacyAccessResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDisableLegacyAccessResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteDisableLegacyAccessResponseFromPb(pb *settingspb.DeleteDisableLegacyAccessResponsePb) (*DeleteDisableLegacyAccessResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyAccessResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDisableLegacyDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDisableLegacyDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDisableLegacyDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDisableLegacyDbfsRequestToPb(st *DeleteDisableLegacyDbfsRequest) (*settingspb.DeleteDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDisableLegacyDbfsRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteDisableLegacyDbfsRequestFromPb(pb *settingspb.DeleteDisableLegacyDbfsRequestPb) (*DeleteDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyDbfsRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteDisableLegacyDbfsResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDisableLegacyDbfsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDisableLegacyDbfsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDisableLegacyDbfsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDisableLegacyDbfsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDisableLegacyDbfsResponseToPb(st *DeleteDisableLegacyDbfsResponse) (*settingspb.DeleteDisableLegacyDbfsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDisableLegacyDbfsResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteDisableLegacyDbfsResponseFromPb(pb *settingspb.DeleteDisableLegacyDbfsResponsePb) (*DeleteDisableLegacyDbfsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyDbfsResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDisableLegacyFeaturesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDisableLegacyFeaturesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDisableLegacyFeaturesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDisableLegacyFeaturesRequestToPb(st *DeleteDisableLegacyFeaturesRequest) (*settingspb.DeleteDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDisableLegacyFeaturesRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteDisableLegacyFeaturesRequestFromPb(pb *settingspb.DeleteDisableLegacyFeaturesRequestPb) (*DeleteDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyFeaturesRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteDisableLegacyFeaturesResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDisableLegacyFeaturesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDisableLegacyFeaturesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteDisableLegacyFeaturesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDisableLegacyFeaturesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDisableLegacyFeaturesResponseToPb(st *DeleteDisableLegacyFeaturesResponse) (*settingspb.DeleteDisableLegacyFeaturesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteDisableLegacyFeaturesResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteDisableLegacyFeaturesResponseFromPb(pb *settingspb.DeleteDisableLegacyFeaturesResponsePb) (*DeleteDisableLegacyFeaturesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyFeaturesResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st DeleteIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteIpAccessListRequestToPb(st *DeleteIpAccessListRequest) (*settingspb.DeleteIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

func DeleteIpAccessListRequestFromPb(pb *settingspb.DeleteIpAccessListRequestPb) (*DeleteIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

type DeleteLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteLlmProxyPartnerPoweredWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteLlmProxyPartnerPoweredWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteLlmProxyPartnerPoweredWorkspaceRequestToPb(st *DeleteLlmProxyPartnerPoweredWorkspaceRequest) (*settingspb.DeleteLlmProxyPartnerPoweredWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteLlmProxyPartnerPoweredWorkspaceRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb *settingspb.DeleteLlmProxyPartnerPoweredWorkspaceRequestPb) (*DeleteLlmProxyPartnerPoweredWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLlmProxyPartnerPoweredWorkspaceRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteLlmProxyPartnerPoweredWorkspaceResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteLlmProxyPartnerPoweredWorkspaceResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteLlmProxyPartnerPoweredWorkspaceResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteLlmProxyPartnerPoweredWorkspaceResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteLlmProxyPartnerPoweredWorkspaceResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteLlmProxyPartnerPoweredWorkspaceResponseToPb(st *DeleteLlmProxyPartnerPoweredWorkspaceResponse) (*settingspb.DeleteLlmProxyPartnerPoweredWorkspaceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteLlmProxyPartnerPoweredWorkspaceResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteLlmProxyPartnerPoweredWorkspaceResponseFromPb(pb *settingspb.DeleteLlmProxyPartnerPoweredWorkspaceResponsePb) (*DeleteLlmProxyPartnerPoweredWorkspaceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLlmProxyPartnerPoweredWorkspaceResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
}

func (st DeleteNetworkConnectivityConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteNetworkConnectivityConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteNetworkConnectivityConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteNetworkConnectivityConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteNetworkConnectivityConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteNetworkConnectivityConfigurationRequestToPb(st *DeleteNetworkConnectivityConfigurationRequest) (*settingspb.DeleteNetworkConnectivityConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteNetworkConnectivityConfigurationRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId

	return pb, nil
}

func DeleteNetworkConnectivityConfigurationRequestFromPb(pb *settingspb.DeleteNetworkConnectivityConfigurationRequestPb) (*DeleteNetworkConnectivityConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkConnectivityConfigurationRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId

	return st, nil
}

type DeleteNetworkPolicyRequest struct {
	// The unique identifier of the network policy to delete.
	NetworkPolicyId string `json:"-" tf:"-"`
}

func (st DeleteNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteNetworkPolicyRequestToPb(st *DeleteNetworkPolicyRequest) (*settingspb.DeleteNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteNetworkPolicyRequestPb{}
	pb.NetworkPolicyId = st.NetworkPolicyId

	return pb, nil
}

func DeleteNetworkPolicyRequestFromPb(pb *settingspb.DeleteNetworkPolicyRequestPb) (*DeleteNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkPolicyRequest{}
	st.NetworkPolicyId = pb.NetworkPolicyId

	return st, nil
}

type DeleteNotificationDestinationRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st DeleteNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteNotificationDestinationRequestToPb(st *DeleteNotificationDestinationRequest) (*settingspb.DeleteNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteNotificationDestinationRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteNotificationDestinationRequestFromPb(pb *settingspb.DeleteNotificationDestinationRequestPb) (*DeleteNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNotificationDestinationRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeletePersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeletePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeletePersonalComputeSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeletePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeletePersonalComputeSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeletePersonalComputeSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeletePersonalComputeSettingRequestToPb(st *DeletePersonalComputeSettingRequest) (*settingspb.DeletePersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeletePersonalComputeSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeletePersonalComputeSettingRequestFromPb(pb *settingspb.DeletePersonalComputeSettingRequestPb) (*DeletePersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePersonalComputeSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeletePersonalComputeSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeletePersonalComputeSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeletePersonalComputeSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeletePersonalComputeSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeletePersonalComputeSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeletePersonalComputeSettingResponseToPb(st *DeletePersonalComputeSettingResponse) (*settingspb.DeletePersonalComputeSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeletePersonalComputeSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeletePersonalComputeSettingResponseFromPb(pb *settingspb.DeletePersonalComputeSettingResponsePb) (*DeletePersonalComputeSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePersonalComputeSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" tf:"-"`
}

func (st DeletePrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeletePrivateEndpointRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeletePrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeletePrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeletePrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeletePrivateEndpointRuleRequestToPb(st *DeletePrivateEndpointRuleRequest) (*settingspb.DeletePrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeletePrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PrivateEndpointRuleId = st.PrivateEndpointRuleId

	return pb, nil
}

func DeletePrivateEndpointRuleRequestFromPb(pb *settingspb.DeletePrivateEndpointRuleRequestPb) (*DeletePrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PrivateEndpointRuleId = pb.PrivateEndpointRuleId

	return st, nil
}

type DeleteRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRestrictWorkspaceAdminsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteRestrictWorkspaceAdminsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRestrictWorkspaceAdminsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteRestrictWorkspaceAdminsSettingRequestToPb(st *DeleteRestrictWorkspaceAdminsSettingRequest) (*settingspb.DeleteRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteRestrictWorkspaceAdminsSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteRestrictWorkspaceAdminsSettingRequestFromPb(pb *settingspb.DeleteRestrictWorkspaceAdminsSettingRequestPb) (*DeleteRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRestrictWorkspaceAdminsSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteRestrictWorkspaceAdminsSettingResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRestrictWorkspaceAdminsSettingResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRestrictWorkspaceAdminsSettingResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteRestrictWorkspaceAdminsSettingResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRestrictWorkspaceAdminsSettingResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteRestrictWorkspaceAdminsSettingResponseToPb(st *DeleteRestrictWorkspaceAdminsSettingResponse) (*settingspb.DeleteRestrictWorkspaceAdminsSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteRestrictWorkspaceAdminsSettingResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteRestrictWorkspaceAdminsSettingResponseFromPb(pb *settingspb.DeleteRestrictWorkspaceAdminsSettingResponsePb) (*DeleteRestrictWorkspaceAdminsSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRestrictWorkspaceAdminsSettingResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DeleteSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteSqlResultsDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteSqlResultsDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteSqlResultsDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteSqlResultsDownloadRequestToPb(st *DeleteSqlResultsDownloadRequest) (*settingspb.DeleteSqlResultsDownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteSqlResultsDownloadRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteSqlResultsDownloadRequestFromPb(pb *settingspb.DeleteSqlResultsDownloadRequestPb) (*DeleteSqlResultsDownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSqlResultsDownloadRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st DeleteSqlResultsDownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteSqlResultsDownloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteSqlResultsDownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteSqlResultsDownloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteSqlResultsDownloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteSqlResultsDownloadResponseToPb(st *DeleteSqlResultsDownloadResponse) (*settingspb.DeleteSqlResultsDownloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteSqlResultsDownloadResponsePb{}
	pb.Etag = st.Etag

	return pb, nil
}

func DeleteSqlResultsDownloadResponseFromPb(pb *settingspb.DeleteSqlResultsDownloadResponsePb) (*DeleteSqlResultsDownloadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteSqlResultsDownloadResponse{}
	st.Etag = pb.Etag

	return st, nil
}

type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId string `json:"-" tf:"-"`
}

func (st DeleteTokenManagementRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteTokenManagementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteTokenManagementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DeleteTokenManagementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteTokenManagementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteTokenManagementRequestToPb(st *DeleteTokenManagementRequest) (*settingspb.DeleteTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DeleteTokenManagementRequestPb{}
	pb.TokenId = st.TokenId

	return pb, nil
}

func DeleteTokenManagementRequestFromPb(pb *settingspb.DeleteTokenManagementRequestPb) (*DeleteTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTokenManagementRequest{}
	st.TokenId = pb.TokenId

	return st, nil
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

func DestinationTypeToPb(st *DestinationType) (*settingspb.DestinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.DestinationTypePb(*st)
	return &pb, nil
}

func DestinationTypeFromPb(pb *settingspb.DestinationTypePb) (*DestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DestinationType(*pb)
	return &st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DisableLegacyAccess) MarshalJSON() ([]byte, error) {
	pb, err := DisableLegacyAccessToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DisableLegacyAccess) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DisableLegacyAccessPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DisableLegacyAccessFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DisableLegacyAccessToPb(st *DisableLegacyAccess) (*settingspb.DisableLegacyAccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DisableLegacyAccessPb{}
	disableLegacyAccessPb, err := BooleanMessageToPb(&st.DisableLegacyAccess)
	if err != nil {
		return nil, err
	}
	if disableLegacyAccessPb != nil {
		pb.DisableLegacyAccess = *disableLegacyAccessPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DisableLegacyAccessFromPb(pb *settingspb.DisableLegacyAccessPb) (*DisableLegacyAccess, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyAccess{}
	disableLegacyAccessField, err := BooleanMessageFromPb(&pb.DisableLegacyAccess)
	if err != nil {
		return nil, err
	}
	if disableLegacyAccessField != nil {
		st.DisableLegacyAccess = *disableLegacyAccessField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DisableLegacyDbfs) MarshalJSON() ([]byte, error) {
	pb, err := DisableLegacyDbfsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DisableLegacyDbfs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DisableLegacyDbfsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DisableLegacyDbfsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DisableLegacyDbfsToPb(st *DisableLegacyDbfs) (*settingspb.DisableLegacyDbfsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DisableLegacyDbfsPb{}
	disableLegacyDbfsPb, err := BooleanMessageToPb(&st.DisableLegacyDbfs)
	if err != nil {
		return nil, err
	}
	if disableLegacyDbfsPb != nil {
		pb.DisableLegacyDbfs = *disableLegacyDbfsPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DisableLegacyDbfsFromPb(pb *settingspb.DisableLegacyDbfsPb) (*DisableLegacyDbfs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyDbfs{}
	disableLegacyDbfsField, err := BooleanMessageFromPb(&pb.DisableLegacyDbfs)
	if err != nil {
		return nil, err
	}
	if disableLegacyDbfsField != nil {
		st.DisableLegacyDbfs = *disableLegacyDbfsField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DisableLegacyFeatures) MarshalJSON() ([]byte, error) {
	pb, err := DisableLegacyFeaturesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DisableLegacyFeatures) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.DisableLegacyFeaturesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DisableLegacyFeaturesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DisableLegacyFeaturesToPb(st *DisableLegacyFeatures) (*settingspb.DisableLegacyFeaturesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.DisableLegacyFeaturesPb{}
	disableLegacyFeaturesPb, err := BooleanMessageToPb(&st.DisableLegacyFeatures)
	if err != nil {
		return nil, err
	}
	if disableLegacyFeaturesPb != nil {
		pb.DisableLegacyFeatures = *disableLegacyFeaturesPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DisableLegacyFeaturesFromPb(pb *settingspb.DisableLegacyFeaturesPb) (*DisableLegacyFeatures, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyFeatures{}
	disableLegacyFeaturesField, err := BooleanMessageFromPb(&pb.DisableLegacyFeatures)
	if err != nil {
		return nil, err
	}
	if disableLegacyFeaturesField != nil {
		st.DisableLegacyFeatures = *disableLegacyFeaturesField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st EgressNetworkPolicy) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyToPb(st *EgressNetworkPolicy) (*settingspb.EgressNetworkPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyPb{}
	internetAccessPb, err := EgressNetworkPolicyInternetAccessPolicyToPb(st.InternetAccess)
	if err != nil {
		return nil, err
	}
	if internetAccessPb != nil {
		pb.InternetAccess = internetAccessPb
	}

	return pb, nil
}

func EgressNetworkPolicyFromPb(pb *settingspb.EgressNetworkPolicyPb) (*EgressNetworkPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicy{}
	internetAccessField, err := EgressNetworkPolicyInternetAccessPolicyFromPb(pb.InternetAccess)
	if err != nil {
		return nil, err
	}
	if internetAccessField != nil {
		st.InternetAccess = internetAccessField
	}

	return st, nil
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

func (st EgressNetworkPolicyInternetAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyInternetAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyInternetAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyInternetAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyInternetAccessPolicyToPb(st *EgressNetworkPolicyInternetAccessPolicy) (*settingspb.EgressNetworkPolicyInternetAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyPb{}

	var allowedInternetDestinationsPb []settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb
	for _, item := range st.AllowedInternetDestinations {
		itemPb, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedInternetDestinationsPb = append(allowedInternetDestinationsPb, *itemPb)
		}
	}
	pb.AllowedInternetDestinations = allowedInternetDestinationsPb

	var allowedStorageDestinationsPb []settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb
	for _, item := range st.AllowedStorageDestinations {
		itemPb, err := EgressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedStorageDestinationsPb = append(allowedStorageDestinationsPb, *itemPb)
		}
	}
	pb.AllowedStorageDestinations = allowedStorageDestinationsPb
	logOnlyModePb, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(st.LogOnlyMode)
	if err != nil {
		return nil, err
	}
	if logOnlyModePb != nil {
		pb.LogOnlyMode = logOnlyModePb
	}
	restrictionModePb, err := EgressNetworkPolicyInternetAccessPolicyRestrictionModeToPb(&st.RestrictionMode)
	if err != nil {
		return nil, err
	}
	if restrictionModePb != nil {
		pb.RestrictionMode = *restrictionModePb
	}

	return pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyPb) (*EgressNetworkPolicyInternetAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicy{}

	var allowedInternetDestinationsField []EgressNetworkPolicyInternetAccessPolicyInternetDestination
	for _, itemPb := range pb.AllowedInternetDestinations {
		item, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allowedInternetDestinationsField = append(allowedInternetDestinationsField, *item)
		}
	}
	st.AllowedInternetDestinations = allowedInternetDestinationsField

	var allowedStorageDestinationsField []EgressNetworkPolicyInternetAccessPolicyStorageDestination
	for _, itemPb := range pb.AllowedStorageDestinations {
		item, err := EgressNetworkPolicyInternetAccessPolicyStorageDestinationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allowedStorageDestinationsField = append(allowedStorageDestinationsField, *item)
		}
	}
	st.AllowedStorageDestinations = allowedStorageDestinationsField
	logOnlyModeField, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeFromPb(pb.LogOnlyMode)
	if err != nil {
		return nil, err
	}
	if logOnlyModeField != nil {
		st.LogOnlyMode = logOnlyModeField
	}
	restrictionModeField, err := EgressNetworkPolicyInternetAccessPolicyRestrictionModeFromPb(&pb.RestrictionMode)
	if err != nil {
		return nil, err
	}
	if restrictionModeField != nil {
		st.RestrictionMode = *restrictionModeField
	}

	return st, nil
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
	Type            EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType `json:"type,omitempty"`
	ForceSendFields []string                                                                          `json:"-" tf:"-"`
}

func (st EgressNetworkPolicyInternetAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyInternetAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(st *EgressNetworkPolicyInternetAccessPolicyInternetDestination) (*settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb{}
	pb.Destination = st.Destination
	protocolPb, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolToPb(&st.Protocol)
	if err != nil {
		return nil, err
	}
	if protocolPb != nil {
		pb.Protocol = *protocolPb
	}
	typePb, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyInternetDestinationFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationPb) (*EgressNetworkPolicyInternetAccessPolicyInternetDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyInternetDestination{}
	st.Destination = pb.Destination
	protocolField, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolFromPb(&pb.Protocol)
	if err != nil {
		return nil, err
	}
	if protocolField != nil {
		st.Protocol = *protocolField
	}
	typeField, err := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolToPb(st *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) (*settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb(*st)
	return &pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb) (*EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol(*pb)
	return &st, nil
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

func EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) (*settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb) (*EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType(*pb)
	return &st, nil
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode struct {

	// Wire name: 'log_only_mode_type'
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType `json:"log_only_mode_type,omitempty"`

	// Wire name: 'workloads'
	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType `json:"workloads,omitempty"`
}

func (st EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(st *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) (*settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModePb{}
	logOnlyModeTypePb, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeToPb(&st.LogOnlyModeType)
	if err != nil {
		return nil, err
	}
	if logOnlyModeTypePb != nil {
		pb.LogOnlyModeType = *logOnlyModeTypePb
	}

	var workloadsPb []settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb
	for _, item := range st.Workloads {
		itemPb, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			workloadsPb = append(workloadsPb, *itemPb)
		}
	}
	pb.Workloads = workloadsPb

	return pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyLogOnlyModeFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModePb) (*EgressNetworkPolicyInternetAccessPolicyLogOnlyMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyLogOnlyMode{}
	logOnlyModeTypeField, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeFromPb(&pb.LogOnlyModeType)
	if err != nil {
		return nil, err
	}
	if logOnlyModeTypeField != nil {
		st.LogOnlyModeType = *logOnlyModeTypeField
	}

	var workloadsField []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType
	for _, itemPb := range pb.Workloads {
		item, err := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			workloadsField = append(workloadsField, *item)
		}
	}
	st.Workloads = workloadsField

	return st, nil
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

func EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) (*settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb) (*EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType(*pb)
	return &st, nil
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

func EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) (*settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb) (*EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType(*pb)
	return &st, nil
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

func EgressNetworkPolicyInternetAccessPolicyRestrictionModeToPb(st *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) (*settingspb.EgressNetworkPolicyInternetAccessPolicyRestrictionModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyInternetAccessPolicyRestrictionModePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyRestrictionModeFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyRestrictionModePb) (*EgressNetworkPolicyInternetAccessPolicyRestrictionMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyRestrictionMode(*pb)
	return &st, nil
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
	Type            EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType `json:"type,omitempty"`
	ForceSendFields []string                                                                        `json:"-" tf:"-"`
}

func (st EgressNetworkPolicyInternetAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyInternetAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyInternetAccessPolicyStorageDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(st *EgressNetworkPolicyInternetAccessPolicyStorageDestination) (*settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb{}
	pb.AllowedPaths = st.AllowedPaths
	pb.AzureContainer = st.AzureContainer
	pb.AzureDnsZone = st.AzureDnsZone
	pb.AzureStorageAccount = st.AzureStorageAccount
	pb.AzureStorageService = st.AzureStorageService
	pb.BucketName = st.BucketName
	pb.Region = st.Region
	typePb, err := EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeToPb(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyStorageDestinationFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationPb) (*EgressNetworkPolicyInternetAccessPolicyStorageDestination, error) {
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
	typeField, err := EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeFromPb(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) (*settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeFromPb(pb *settingspb.EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb) (*EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType(*pb)
	return &st, nil
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

func (st EgressNetworkPolicyNetworkAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyNetworkAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyNetworkAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyNetworkAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyNetworkAccessPolicyToPb(st *EgressNetworkPolicyNetworkAccessPolicy) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyPb{}

	var allowedInternetDestinationsPb []settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb
	for _, item := range st.AllowedInternetDestinations {
		itemPb, err := EgressNetworkPolicyNetworkAccessPolicyInternetDestinationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedInternetDestinationsPb = append(allowedInternetDestinationsPb, *itemPb)
		}
	}
	pb.AllowedInternetDestinations = allowedInternetDestinationsPb

	var allowedStorageDestinationsPb []settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb
	for _, item := range st.AllowedStorageDestinations {
		itemPb, err := EgressNetworkPolicyNetworkAccessPolicyStorageDestinationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedStorageDestinationsPb = append(allowedStorageDestinationsPb, *itemPb)
		}
	}
	pb.AllowedStorageDestinations = allowedStorageDestinationsPb
	policyEnforcementPb, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementToPb(st.PolicyEnforcement)
	if err != nil {
		return nil, err
	}
	if policyEnforcementPb != nil {
		pb.PolicyEnforcement = policyEnforcementPb
	}
	restrictionModePb, err := EgressNetworkPolicyNetworkAccessPolicyRestrictionModeToPb(&st.RestrictionMode)
	if err != nil {
		return nil, err
	}
	if restrictionModePb != nil {
		pb.RestrictionMode = *restrictionModePb
	}

	return pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyPb) (*EgressNetworkPolicyNetworkAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicy{}

	var allowedInternetDestinationsField []EgressNetworkPolicyNetworkAccessPolicyInternetDestination
	for _, itemPb := range pb.AllowedInternetDestinations {
		item, err := EgressNetworkPolicyNetworkAccessPolicyInternetDestinationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allowedInternetDestinationsField = append(allowedInternetDestinationsField, *item)
		}
	}
	st.AllowedInternetDestinations = allowedInternetDestinationsField

	var allowedStorageDestinationsField []EgressNetworkPolicyNetworkAccessPolicyStorageDestination
	for _, itemPb := range pb.AllowedStorageDestinations {
		item, err := EgressNetworkPolicyNetworkAccessPolicyStorageDestinationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allowedStorageDestinationsField = append(allowedStorageDestinationsField, *item)
		}
	}
	st.AllowedStorageDestinations = allowedStorageDestinationsField
	policyEnforcementField, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementFromPb(pb.PolicyEnforcement)
	if err != nil {
		return nil, err
	}
	if policyEnforcementField != nil {
		st.PolicyEnforcement = policyEnforcementField
	}
	restrictionModeField, err := EgressNetworkPolicyNetworkAccessPolicyRestrictionModeFromPb(&pb.RestrictionMode)
	if err != nil {
		return nil, err
	}
	if restrictionModeField != nil {
		st.RestrictionMode = *restrictionModeField
	}

	return st, nil
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
	ForceSendFields         []string                                                                         `json:"-" tf:"-"`
}

func (st EgressNetworkPolicyNetworkAccessPolicyInternetDestination) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyNetworkAccessPolicyInternetDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyNetworkAccessPolicyInternetDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyNetworkAccessPolicyInternetDestinationToPb(st *EgressNetworkPolicyNetworkAccessPolicyInternetDestination) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb{}
	pb.Destination = st.Destination
	internetDestinationTypePb, err := EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeToPb(&st.InternetDestinationType)
	if err != nil {
		return nil, err
	}
	if internetDestinationTypePb != nil {
		pb.InternetDestinationType = *internetDestinationTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyInternetDestinationFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationPb) (*EgressNetworkPolicyNetworkAccessPolicyInternetDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicyInternetDestination{}
	st.Destination = pb.Destination
	internetDestinationTypeField, err := EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeFromPb(&pb.InternetDestinationType)
	if err != nil {
		return nil, err
	}
	if internetDestinationTypeField != nil {
		st.InternetDestinationType = *internetDestinationTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeToPb(st *EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypeFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationTypePb) (*EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyNetworkAccessPolicyInternetDestinationInternetDestinationType(*pb)
	return &st, nil
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

func (st EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementToPb(st *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb{}

	var dryRunModeProductFilterPb []settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb
	for _, item := range st.DryRunModeProductFilter {
		itemPb, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dryRunModeProductFilterPb = append(dryRunModeProductFilterPb, *itemPb)
		}
	}
	pb.DryRunModeProductFilter = dryRunModeProductFilterPb
	enforcementModePb, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeToPb(&st.EnforcementMode)
	if err != nil {
		return nil, err
	}
	if enforcementModePb != nil {
		pb.EnforcementMode = *enforcementModePb
	}

	return pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementPb) (*EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcement{}

	var dryRunModeProductFilterField []EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter
	for _, itemPb := range pb.DryRunModeProductFilter {
		item, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dryRunModeProductFilterField = append(dryRunModeProductFilterField, *item)
		}
	}
	st.DryRunModeProductFilter = dryRunModeProductFilterField
	enforcementModeField, err := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeFromPb(&pb.EnforcementMode)
	if err != nil {
		return nil, err
	}
	if enforcementModeField != nil {
		st.EnforcementMode = *enforcementModeField
	}

	return st, nil
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

func EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterToPb(st *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb(*st)
	return &pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilterPb) (*EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementDryRunModeProductFilter(*pb)
	return &st, nil
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

func EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeToPb(st *EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModeFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementModePb) (*EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyNetworkAccessPolicyPolicyEnforcementEnforcementMode(*pb)
	return &st, nil
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

func EgressNetworkPolicyNetworkAccessPolicyRestrictionModeToPb(st *EgressNetworkPolicyNetworkAccessPolicyRestrictionMode) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyRestrictionModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyNetworkAccessPolicyRestrictionModePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyRestrictionModeFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyRestrictionModePb) (*EgressNetworkPolicyNetworkAccessPolicyRestrictionMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyNetworkAccessPolicyRestrictionMode(*pb)
	return &st, nil
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
	ForceSendFields        []string                                                                       `json:"-" tf:"-"`
}

func (st EgressNetworkPolicyNetworkAccessPolicyStorageDestination) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyNetworkAccessPolicyStorageDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EgressNetworkPolicyNetworkAccessPolicyStorageDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EgressNetworkPolicyNetworkAccessPolicyStorageDestinationToPb(st *EgressNetworkPolicyNetworkAccessPolicyStorageDestination) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb{}
	pb.AzureStorageAccount = st.AzureStorageAccount
	pb.AzureStorageService = st.AzureStorageService
	pb.BucketName = st.BucketName
	pb.Region = st.Region
	storageDestinationTypePb, err := EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeToPb(&st.StorageDestinationType)
	if err != nil {
		return nil, err
	}
	if storageDestinationTypePb != nil {
		pb.StorageDestinationType = *storageDestinationTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyStorageDestinationFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationPb) (*EgressNetworkPolicyNetworkAccessPolicyStorageDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyNetworkAccessPolicyStorageDestination{}
	st.AzureStorageAccount = pb.AzureStorageAccount
	st.AzureStorageService = pb.AzureStorageService
	st.BucketName = pb.BucketName
	st.Region = pb.Region
	storageDestinationTypeField, err := EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeFromPb(&pb.StorageDestinationType)
	if err != nil {
		return nil, err
	}
	if storageDestinationTypeField != nil {
		st.StorageDestinationType = *storageDestinationTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeToPb(st *EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType) (*settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb(*st)
	return &pb, nil
}

func EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypeFromPb(pb *settingspb.EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationTypePb) (*EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyNetworkAccessPolicyStorageDestinationStorageDestinationType(*pb)
	return &st, nil
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

func EgressResourceTypeToPb(st *EgressResourceType) (*settingspb.EgressResourceTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.EgressResourceTypePb(*st)
	return &pb, nil
}

func EgressResourceTypeFromPb(pb *settingspb.EgressResourceTypePb) (*EgressResourceType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressResourceType(*pb)
	return &st, nil
}

type EmailConfig struct {
	// Email addresses to notify.
	// Wire name: 'addresses'
	Addresses []string `json:"addresses,omitempty"`
}

func (st EmailConfig) MarshalJSON() ([]byte, error) {
	pb, err := EmailConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EmailConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EmailConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EmailConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EmailConfigToPb(st *EmailConfig) (*settingspb.EmailConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EmailConfigPb{}
	pb.Addresses = st.Addresses

	return pb, nil
}

func EmailConfigFromPb(pb *settingspb.EmailConfigPb) (*EmailConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmailConfig{}
	st.Addresses = pb.Addresses

	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EnableExportNotebook) MarshalJSON() ([]byte, error) {
	pb, err := EnableExportNotebookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnableExportNotebook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EnableExportNotebookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnableExportNotebookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnableExportNotebookToPb(st *EnableExportNotebook) (*settingspb.EnableExportNotebookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EnableExportNotebookPb{}
	booleanValPb, err := BooleanMessageToPb(st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = booleanValPb
	}
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnableExportNotebookFromPb(pb *settingspb.EnableExportNotebookPb) (*EnableExportNotebook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableExportNotebook{}
	booleanValField, err := BooleanMessageFromPb(pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = booleanValField
	}
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EnableNotebookTableClipboard) MarshalJSON() ([]byte, error) {
	pb, err := EnableNotebookTableClipboardToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnableNotebookTableClipboard) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EnableNotebookTableClipboardPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnableNotebookTableClipboardFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnableNotebookTableClipboardToPb(st *EnableNotebookTableClipboard) (*settingspb.EnableNotebookTableClipboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EnableNotebookTableClipboardPb{}
	booleanValPb, err := BooleanMessageToPb(st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = booleanValPb
	}
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnableNotebookTableClipboardFromPb(pb *settingspb.EnableNotebookTableClipboardPb) (*EnableNotebookTableClipboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableNotebookTableClipboard{}
	booleanValField, err := BooleanMessageFromPb(pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = booleanValField
	}
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EnableResultsDownloading) MarshalJSON() ([]byte, error) {
	pb, err := EnableResultsDownloadingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnableResultsDownloading) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EnableResultsDownloadingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnableResultsDownloadingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnableResultsDownloadingToPb(st *EnableResultsDownloading) (*settingspb.EnableResultsDownloadingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EnableResultsDownloadingPb{}
	booleanValPb, err := BooleanMessageToPb(st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = booleanValPb
	}
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnableResultsDownloadingFromPb(pb *settingspb.EnableResultsDownloadingPb) (*EnableResultsDownloading, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableResultsDownloading{}
	booleanValField, err := BooleanMessageFromPb(pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = booleanValField
	}
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {

	// Wire name: 'is_enabled'
	IsEnabled       bool     `json:"is_enabled,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EnhancedSecurityMonitoring) MarshalJSON() ([]byte, error) {
	pb, err := EnhancedSecurityMonitoringToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnhancedSecurityMonitoring) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EnhancedSecurityMonitoringPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnhancedSecurityMonitoringFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnhancedSecurityMonitoringToPb(st *EnhancedSecurityMonitoring) (*settingspb.EnhancedSecurityMonitoringPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EnhancedSecurityMonitoringPb{}
	pb.IsEnabled = st.IsEnabled

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnhancedSecurityMonitoringFromPb(pb *settingspb.EnhancedSecurityMonitoringPb) (*EnhancedSecurityMonitoring, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnhancedSecurityMonitoring{}
	st.IsEnabled = pb.IsEnabled

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EnhancedSecurityMonitoringSetting) MarshalJSON() ([]byte, error) {
	pb, err := EnhancedSecurityMonitoringSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EnhancedSecurityMonitoringSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EnhancedSecurityMonitoringSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EnhancedSecurityMonitoringSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EnhancedSecurityMonitoringSettingToPb(st *EnhancedSecurityMonitoringSetting) (*settingspb.EnhancedSecurityMonitoringSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EnhancedSecurityMonitoringSettingPb{}
	enhancedSecurityMonitoringWorkspacePb, err := EnhancedSecurityMonitoringToPb(&st.EnhancedSecurityMonitoringWorkspace)
	if err != nil {
		return nil, err
	}
	if enhancedSecurityMonitoringWorkspacePb != nil {
		pb.EnhancedSecurityMonitoringWorkspace = *enhancedSecurityMonitoringWorkspacePb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EnhancedSecurityMonitoringSettingFromPb(pb *settingspb.EnhancedSecurityMonitoringSettingPb) (*EnhancedSecurityMonitoringSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnhancedSecurityMonitoringSetting{}
	enhancedSecurityMonitoringWorkspaceField, err := EnhancedSecurityMonitoringFromPb(&pb.EnhancedSecurityMonitoringWorkspace)
	if err != nil {
		return nil, err
	}
	if enhancedSecurityMonitoringWorkspaceField != nil {
		st.EnhancedSecurityMonitoringWorkspace = *enhancedSecurityMonitoringWorkspaceField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Account level policy for ESM
type EsmEnablementAccount struct {

	// Wire name: 'is_enforced'
	IsEnforced      bool     `json:"is_enforced,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EsmEnablementAccount) MarshalJSON() ([]byte, error) {
	pb, err := EsmEnablementAccountToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EsmEnablementAccount) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EsmEnablementAccountPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EsmEnablementAccountFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EsmEnablementAccountToPb(st *EsmEnablementAccount) (*settingspb.EsmEnablementAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EsmEnablementAccountPb{}
	pb.IsEnforced = st.IsEnforced

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EsmEnablementAccountFromPb(pb *settingspb.EsmEnablementAccountPb) (*EsmEnablementAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EsmEnablementAccount{}
	st.IsEnforced = pb.IsEnforced

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EsmEnablementAccountSetting) MarshalJSON() ([]byte, error) {
	pb, err := EsmEnablementAccountSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *EsmEnablementAccountSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.EsmEnablementAccountSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := EsmEnablementAccountSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func EsmEnablementAccountSettingToPb(st *EsmEnablementAccountSetting) (*settingspb.EsmEnablementAccountSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.EsmEnablementAccountSettingPb{}
	esmEnablementAccountPb, err := EsmEnablementAccountToPb(&st.EsmEnablementAccount)
	if err != nil {
		return nil, err
	}
	if esmEnablementAccountPb != nil {
		pb.EsmEnablementAccount = *esmEnablementAccountPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func EsmEnablementAccountSettingFromPb(pb *settingspb.EsmEnablementAccountSettingPb) (*EsmEnablementAccountSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EsmEnablementAccountSetting{}
	esmEnablementAccountField, err := EsmEnablementAccountFromPb(&pb.EsmEnablementAccount)
	if err != nil {
		return nil, err
	}
	if esmEnablementAccountField != nil {
		st.EsmEnablementAccount = *esmEnablementAccountField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	TokenType       TokenType `json:"tokenType,omitempty"`
	ForceSendFields []string  `json:"-" tf:"-"`
}

func (st ExchangeToken) MarshalJSON() ([]byte, error) {
	pb, err := ExchangeTokenToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExchangeToken) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ExchangeTokenPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExchangeTokenFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExchangeTokenToPb(st *ExchangeToken) (*settingspb.ExchangeTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ExchangeTokenPb{}
	pb.Credential = st.Credential
	pb.CredentialEolTime = st.CredentialEolTime
	pb.OwnerId = st.OwnerId
	pb.Scopes = st.Scopes
	tokenTypePb, err := TokenTypeToPb(&st.TokenType)
	if err != nil {
		return nil, err
	}
	if tokenTypePb != nil {
		pb.TokenType = *tokenTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExchangeTokenFromPb(pb *settingspb.ExchangeTokenPb) (*ExchangeToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeToken{}
	st.Credential = pb.Credential
	st.CredentialEolTime = pb.CredentialEolTime
	st.OwnerId = pb.OwnerId
	st.Scopes = pb.Scopes
	tokenTypeField, err := TokenTypeFromPb(&pb.TokenType)
	if err != nil {
		return nil, err
	}
	if tokenTypeField != nil {
		st.TokenType = *tokenTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st ExchangeTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := ExchangeTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExchangeTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ExchangeTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExchangeTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExchangeTokenRequestToPb(st *ExchangeTokenRequest) (*settingspb.ExchangeTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ExchangeTokenRequestPb{}
	partitionIdPb, err := PartitionIdToPb(&st.PartitionId)
	if err != nil {
		return nil, err
	}
	if partitionIdPb != nil {
		pb.PartitionId = *partitionIdPb
	}
	pb.Scopes = st.Scopes

	var tokenTypePb []settingspb.TokenTypePb
	for _, item := range st.TokenType {
		itemPb, err := TokenTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tokenTypePb = append(tokenTypePb, *itemPb)
		}
	}
	pb.TokenType = tokenTypePb

	return pb, nil
}

func ExchangeTokenRequestFromPb(pb *settingspb.ExchangeTokenRequestPb) (*ExchangeTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeTokenRequest{}
	partitionIdField, err := PartitionIdFromPb(&pb.PartitionId)
	if err != nil {
		return nil, err
	}
	if partitionIdField != nil {
		st.PartitionId = *partitionIdField
	}
	st.Scopes = pb.Scopes

	var tokenTypeField []TokenType
	for _, itemPb := range pb.TokenType {
		item, err := TokenTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tokenTypeField = append(tokenTypeField, *item)
		}
	}
	st.TokenType = tokenTypeField

	return st, nil
}

// Exhanged tokens were successfully returned.
type ExchangeTokenResponse struct {

	// Wire name: 'values'
	Values []ExchangeToken `json:"values,omitempty"`
}

func (st ExchangeTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := ExchangeTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExchangeTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ExchangeTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExchangeTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExchangeTokenResponseToPb(st *ExchangeTokenResponse) (*settingspb.ExchangeTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ExchangeTokenResponsePb{}

	var valuesPb []settingspb.ExchangeTokenPb
	for _, item := range st.Values {
		itemPb, err := ExchangeTokenToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			valuesPb = append(valuesPb, *itemPb)
		}
	}
	pb.Values = valuesPb

	return pb, nil
}

func ExchangeTokenResponseFromPb(pb *settingspb.ExchangeTokenResponsePb) (*ExchangeTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeTokenResponse{}

	var valuesField []ExchangeToken
	for _, itemPb := range pb.Values {
		item, err := ExchangeTokenFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			valuesField = append(valuesField, *item)
		}
	}
	st.Values = valuesField

	return st, nil
}

// An IP access list was successfully returned.
type FetchIpAccessListResponse struct {

	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (st FetchIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := FetchIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FetchIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.FetchIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FetchIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FetchIpAccessListResponseToPb(st *FetchIpAccessListResponse) (*settingspb.FetchIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.FetchIpAccessListResponsePb{}
	ipAccessListPb, err := IpAccessListInfoToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	return pb, nil
}

func FetchIpAccessListResponseFromPb(pb *settingspb.FetchIpAccessListResponsePb) (*FetchIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FetchIpAccessListResponse{}
	ipAccessListField, err := IpAccessListInfoFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}

	return st, nil
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
	UsernameSet     bool     `json:"username_set,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GenericWebhookConfig) MarshalJSON() ([]byte, error) {
	pb, err := GenericWebhookConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GenericWebhookConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GenericWebhookConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GenericWebhookConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GenericWebhookConfigToPb(st *GenericWebhookConfig) (*settingspb.GenericWebhookConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GenericWebhookConfigPb{}
	pb.Password = st.Password
	pb.PasswordSet = st.PasswordSet
	pb.Url = st.Url
	pb.UrlSet = st.UrlSet
	pb.Username = st.Username
	pb.UsernameSet = st.UsernameSet

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GenericWebhookConfigFromPb(pb *settingspb.GenericWebhookConfigPb) (*GenericWebhookConfig, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetAccountIpAccessEnableRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAccountIpAccessEnableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetAccountIpAccessEnableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAccountIpAccessEnableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAccountIpAccessEnableRequestToPb(st *GetAccountIpAccessEnableRequest) (*settingspb.GetAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetAccountIpAccessEnableRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetAccountIpAccessEnableRequestFromPb(pb *settingspb.GetAccountIpAccessEnableRequestPb) (*GetAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountIpAccessEnableRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st GetAccountIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAccountIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAccountIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetAccountIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAccountIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAccountIpAccessListRequestToPb(st *GetAccountIpAccessListRequest) (*settingspb.GetAccountIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetAccountIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

func GetAccountIpAccessListRequestFromPb(pb *settingspb.GetAccountIpAccessListRequestPb) (*GetAccountIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

type GetAibiDashboardEmbeddingAccessPolicySettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*settingspb.GetAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *settingspb.GetAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*GetAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAibiDashboardEmbeddingAccessPolicySettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetAibiDashboardEmbeddingApprovedDomainsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settingspb.GetAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *settingspb.GetAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*GetAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetAutomaticClusterUpdateSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetAutomaticClusterUpdateSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAutomaticClusterUpdateSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetAutomaticClusterUpdateSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAutomaticClusterUpdateSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAutomaticClusterUpdateSettingRequestToPb(st *GetAutomaticClusterUpdateSettingRequest) (*settingspb.GetAutomaticClusterUpdateSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetAutomaticClusterUpdateSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetAutomaticClusterUpdateSettingRequestFromPb(pb *settingspb.GetAutomaticClusterUpdateSettingRequestPb) (*GetAutomaticClusterUpdateSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAutomaticClusterUpdateSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetComplianceSecurityProfileSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetComplianceSecurityProfileSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetComplianceSecurityProfileSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetComplianceSecurityProfileSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetComplianceSecurityProfileSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetComplianceSecurityProfileSettingRequestToPb(st *GetComplianceSecurityProfileSettingRequest) (*settingspb.GetComplianceSecurityProfileSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetComplianceSecurityProfileSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetComplianceSecurityProfileSettingRequestFromPb(pb *settingspb.GetComplianceSecurityProfileSettingRequestPb) (*GetComplianceSecurityProfileSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetComplianceSecurityProfileSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetCspEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetCspEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCspEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetCspEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCspEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCspEnablementAccountSettingRequestToPb(st *GetCspEnablementAccountSettingRequest) (*settingspb.GetCspEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetCspEnablementAccountSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetCspEnablementAccountSettingRequestFromPb(pb *settingspb.GetCspEnablementAccountSettingRequestPb) (*GetCspEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCspEnablementAccountSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDashboardEmailSubscriptionsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDashboardEmailSubscriptionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetDashboardEmailSubscriptionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDashboardEmailSubscriptionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetDashboardEmailSubscriptionsRequestToPb(st *GetDashboardEmailSubscriptionsRequest) (*settingspb.GetDashboardEmailSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetDashboardEmailSubscriptionsRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetDashboardEmailSubscriptionsRequestFromPb(pb *settingspb.GetDashboardEmailSubscriptionsRequestPb) (*GetDashboardEmailSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDashboardEmailSubscriptionsRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDefaultNamespaceSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDefaultNamespaceSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetDefaultNamespaceSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDefaultNamespaceSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetDefaultNamespaceSettingRequestToPb(st *GetDefaultNamespaceSettingRequest) (*settingspb.GetDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetDefaultNamespaceSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetDefaultNamespaceSettingRequestFromPb(pb *settingspb.GetDefaultNamespaceSettingRequestPb) (*GetDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDefaultNamespaceSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDefaultWarehouseIdRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetDefaultWarehouseIdRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDefaultWarehouseIdRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDefaultWarehouseIdRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetDefaultWarehouseIdRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDefaultWarehouseIdRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetDefaultWarehouseIdRequestToPb(st *GetDefaultWarehouseIdRequest) (*settingspb.GetDefaultWarehouseIdRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetDefaultWarehouseIdRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetDefaultWarehouseIdRequestFromPb(pb *settingspb.GetDefaultWarehouseIdRequestPb) (*GetDefaultWarehouseIdRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDefaultWarehouseIdRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDisableLegacyAccessRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDisableLegacyAccessRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetDisableLegacyAccessRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDisableLegacyAccessRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetDisableLegacyAccessRequestToPb(st *GetDisableLegacyAccessRequest) (*settingspb.GetDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetDisableLegacyAccessRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetDisableLegacyAccessRequestFromPb(pb *settingspb.GetDisableLegacyAccessRequestPb) (*GetDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyAccessRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDisableLegacyDbfsRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDisableLegacyDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetDisableLegacyDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDisableLegacyDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetDisableLegacyDbfsRequestToPb(st *GetDisableLegacyDbfsRequest) (*settingspb.GetDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetDisableLegacyDbfsRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetDisableLegacyDbfsRequestFromPb(pb *settingspb.GetDisableLegacyDbfsRequestPb) (*GetDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyDbfsRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDisableLegacyFeaturesRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDisableLegacyFeaturesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetDisableLegacyFeaturesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDisableLegacyFeaturesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetDisableLegacyFeaturesRequestToPb(st *GetDisableLegacyFeaturesRequest) (*settingspb.GetDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetDisableLegacyFeaturesRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetDisableLegacyFeaturesRequestFromPb(pb *settingspb.GetDisableLegacyFeaturesRequestPb) (*GetDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyFeaturesRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetEnhancedSecurityMonitoringSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetEnhancedSecurityMonitoringSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetEnhancedSecurityMonitoringSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetEnhancedSecurityMonitoringSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetEnhancedSecurityMonitoringSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetEnhancedSecurityMonitoringSettingRequestToPb(st *GetEnhancedSecurityMonitoringSettingRequest) (*settingspb.GetEnhancedSecurityMonitoringSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetEnhancedSecurityMonitoringSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetEnhancedSecurityMonitoringSettingRequestFromPb(pb *settingspb.GetEnhancedSecurityMonitoringSettingRequestPb) (*GetEnhancedSecurityMonitoringSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEnhancedSecurityMonitoringSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetEsmEnablementAccountSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetEsmEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetEsmEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetEsmEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetEsmEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetEsmEnablementAccountSettingRequestToPb(st *GetEsmEnablementAccountSettingRequest) (*settingspb.GetEsmEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetEsmEnablementAccountSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetEsmEnablementAccountSettingRequestFromPb(pb *settingspb.GetEsmEnablementAccountSettingRequestPb) (*GetEsmEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEsmEnablementAccountSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" tf:"-"`
}

func (st GetIpAccessListRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetIpAccessListRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetIpAccessListRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetIpAccessListRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetIpAccessListRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetIpAccessListRequestToPb(st *GetIpAccessListRequest) (*settingspb.GetIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetIpAccessListRequestPb{}
	pb.IpAccessListId = st.IpAccessListId

	return pb, nil
}

func GetIpAccessListRequestFromPb(pb *settingspb.GetIpAccessListRequestPb) (*GetIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListRequest{}
	st.IpAccessListId = pb.IpAccessListId

	return st, nil
}

type GetIpAccessListResponse struct {

	// Wire name: 'ip_access_list'
	IpAccessList *IpAccessListInfo `json:"ip_access_list,omitempty"`
}

func (st GetIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetIpAccessListResponseToPb(st *GetIpAccessListResponse) (*settingspb.GetIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetIpAccessListResponsePb{}
	ipAccessListPb, err := IpAccessListInfoToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	return pb, nil
}

func GetIpAccessListResponseFromPb(pb *settingspb.GetIpAccessListResponsePb) (*GetIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListResponse{}
	ipAccessListField, err := IpAccessListInfoFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}

	return st, nil
}

// IP access lists were successfully returned.
type GetIpAccessListsResponse struct {

	// Wire name: 'ip_access_lists'
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func (st GetIpAccessListsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetIpAccessListsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetIpAccessListsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetIpAccessListsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetIpAccessListsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetIpAccessListsResponseToPb(st *GetIpAccessListsResponse) (*settingspb.GetIpAccessListsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetIpAccessListsResponsePb{}

	var ipAccessListsPb []settingspb.IpAccessListInfoPb
	for _, item := range st.IpAccessLists {
		itemPb, err := IpAccessListInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			ipAccessListsPb = append(ipAccessListsPb, *itemPb)
		}
	}
	pb.IpAccessLists = ipAccessListsPb

	return pb, nil
}

func GetIpAccessListsResponseFromPb(pb *settingspb.GetIpAccessListsResponsePb) (*GetIpAccessListsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListsResponse{}

	var ipAccessListsField []IpAccessListInfo
	for _, itemPb := range pb.IpAccessLists {
		item, err := IpAccessListInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			ipAccessListsField = append(ipAccessListsField, *item)
		}
	}
	st.IpAccessLists = ipAccessListsField

	return st, nil
}

type GetLlmProxyPartnerPoweredAccountRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetLlmProxyPartnerPoweredAccountRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetLlmProxyPartnerPoweredAccountRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLlmProxyPartnerPoweredAccountRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetLlmProxyPartnerPoweredAccountRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLlmProxyPartnerPoweredAccountRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLlmProxyPartnerPoweredAccountRequestToPb(st *GetLlmProxyPartnerPoweredAccountRequest) (*settingspb.GetLlmProxyPartnerPoweredAccountRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetLlmProxyPartnerPoweredAccountRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetLlmProxyPartnerPoweredAccountRequestFromPb(pb *settingspb.GetLlmProxyPartnerPoweredAccountRequestPb) (*GetLlmProxyPartnerPoweredAccountRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLlmProxyPartnerPoweredAccountRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetLlmProxyPartnerPoweredEnforceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetLlmProxyPartnerPoweredEnforceRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetLlmProxyPartnerPoweredEnforceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLlmProxyPartnerPoweredEnforceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetLlmProxyPartnerPoweredEnforceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLlmProxyPartnerPoweredEnforceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLlmProxyPartnerPoweredEnforceRequestToPb(st *GetLlmProxyPartnerPoweredEnforceRequest) (*settingspb.GetLlmProxyPartnerPoweredEnforceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetLlmProxyPartnerPoweredEnforceRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetLlmProxyPartnerPoweredEnforceRequestFromPb(pb *settingspb.GetLlmProxyPartnerPoweredEnforceRequestPb) (*GetLlmProxyPartnerPoweredEnforceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLlmProxyPartnerPoweredEnforceRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetLlmProxyPartnerPoweredWorkspaceRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetLlmProxyPartnerPoweredWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetLlmProxyPartnerPoweredWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLlmProxyPartnerPoweredWorkspaceRequestToPb(st *GetLlmProxyPartnerPoweredWorkspaceRequest) (*settingspb.GetLlmProxyPartnerPoweredWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetLlmProxyPartnerPoweredWorkspaceRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb *settingspb.GetLlmProxyPartnerPoweredWorkspaceRequestPb) (*GetLlmProxyPartnerPoweredWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLlmProxyPartnerPoweredWorkspaceRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
}

func (st GetNetworkConnectivityConfigurationRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetNetworkConnectivityConfigurationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetNetworkConnectivityConfigurationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetNetworkConnectivityConfigurationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetNetworkConnectivityConfigurationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetNetworkConnectivityConfigurationRequestToPb(st *GetNetworkConnectivityConfigurationRequest) (*settingspb.GetNetworkConnectivityConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetNetworkConnectivityConfigurationRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId

	return pb, nil
}

func GetNetworkConnectivityConfigurationRequestFromPb(pb *settingspb.GetNetworkConnectivityConfigurationRequestPb) (*GetNetworkConnectivityConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkConnectivityConfigurationRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId

	return st, nil
}

type GetNetworkPolicyRequest struct {
	// The unique identifier of the network policy to retrieve.
	NetworkPolicyId string `json:"-" tf:"-"`
}

func (st GetNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetNetworkPolicyRequestToPb(st *GetNetworkPolicyRequest) (*settingspb.GetNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetNetworkPolicyRequestPb{}
	pb.NetworkPolicyId = st.NetworkPolicyId

	return pb, nil
}

func GetNetworkPolicyRequestFromPb(pb *settingspb.GetNetworkPolicyRequestPb) (*GetNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkPolicyRequest{}
	st.NetworkPolicyId = pb.NetworkPolicyId

	return st, nil
}

type GetNotificationDestinationRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st GetNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetNotificationDestinationRequestToPb(st *GetNotificationDestinationRequest) (*settingspb.GetNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetNotificationDestinationRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetNotificationDestinationRequestFromPb(pb *settingspb.GetNotificationDestinationRequestPb) (*GetNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNotificationDestinationRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetPersonalComputeSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetPersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetPersonalComputeSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetPersonalComputeSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPersonalComputeSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPersonalComputeSettingRequestToPb(st *GetPersonalComputeSettingRequest) (*settingspb.GetPersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetPersonalComputeSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetPersonalComputeSettingRequestFromPb(pb *settingspb.GetPersonalComputeSettingRequestPb) (*GetPersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalComputeSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" tf:"-"`
}

func (st GetPrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetPrivateEndpointRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetPrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPrivateEndpointRuleRequestToPb(st *GetPrivateEndpointRuleRequest) (*settingspb.GetPrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetPrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PrivateEndpointRuleId = st.PrivateEndpointRuleId

	return pb, nil
}

func GetPrivateEndpointRuleRequestFromPb(pb *settingspb.GetPrivateEndpointRuleRequestPb) (*GetPrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PrivateEndpointRuleId = pb.PrivateEndpointRuleId

	return st, nil
}

type GetRestrictWorkspaceAdminsSettingRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRestrictWorkspaceAdminsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetRestrictWorkspaceAdminsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRestrictWorkspaceAdminsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRestrictWorkspaceAdminsSettingRequestToPb(st *GetRestrictWorkspaceAdminsSettingRequest) (*settingspb.GetRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetRestrictWorkspaceAdminsSettingRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetRestrictWorkspaceAdminsSettingRequestFromPb(pb *settingspb.GetRestrictWorkspaceAdminsSettingRequestPb) (*GetRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRestrictWorkspaceAdminsSettingRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetSqlResultsDownloadRequest struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag            string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetSqlResultsDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetSqlResultsDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetSqlResultsDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetSqlResultsDownloadRequestToPb(st *GetSqlResultsDownloadRequest) (*settingspb.GetSqlResultsDownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetSqlResultsDownloadRequestPb{}
	pb.Etag = st.Etag

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetSqlResultsDownloadRequestFromPb(pb *settingspb.GetSqlResultsDownloadRequestPb) (*GetSqlResultsDownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetSqlResultsDownloadRequest{}
	st.Etag = pb.Etag

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetStatusRequest struct {
	Keys string `json:"-" tf:"-"`
}

func (st GetStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetStatusRequestToPb(st *GetStatusRequest) (*settingspb.GetStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetStatusRequestPb{}
	pb.Keys = st.Keys

	return pb, nil
}

func GetStatusRequestFromPb(pb *settingspb.GetStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	st.Keys = pb.Keys

	return st, nil
}

type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string `json:"-" tf:"-"`
}

func (st GetTokenManagementRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetTokenManagementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetTokenManagementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetTokenManagementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetTokenManagementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetTokenManagementRequestToPb(st *GetTokenManagementRequest) (*settingspb.GetTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetTokenManagementRequestPb{}
	pb.TokenId = st.TokenId

	return pb, nil
}

func GetTokenManagementRequestFromPb(pb *settingspb.GetTokenManagementRequestPb) (*GetTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenManagementRequest{}
	st.TokenId = pb.TokenId

	return st, nil
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []TokenPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetTokenPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetTokenPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetTokenPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetTokenPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetTokenPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetTokenPermissionLevelsResponseToPb(st *GetTokenPermissionLevelsResponse) (*settingspb.GetTokenPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetTokenPermissionLevelsResponsePb{}

	var permissionLevelsPb []settingspb.TokenPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := TokenPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetTokenPermissionLevelsResponseFromPb(pb *settingspb.GetTokenPermissionLevelsResponsePb) (*GetTokenPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenPermissionLevelsResponse{}

	var permissionLevelsField []TokenPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := TokenPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

// Token with specified Token ID was successfully returned.
type GetTokenResponse struct {

	// Wire name: 'token_info'
	TokenInfo *TokenInfo `json:"token_info,omitempty"`
}

func (st GetTokenResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetTokenResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetTokenResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetTokenResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetTokenResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetTokenResponseToPb(st *GetTokenResponse) (*settingspb.GetTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetTokenResponsePb{}
	tokenInfoPb, err := TokenInfoToPb(st.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoPb != nil {
		pb.TokenInfo = tokenInfoPb
	}

	return pb, nil
}

func GetTokenResponseFromPb(pb *settingspb.GetTokenResponsePb) (*GetTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenResponse{}
	tokenInfoField, err := TokenInfoFromPb(pb.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoField != nil {
		st.TokenInfo = tokenInfoField
	}

	return st, nil
}

type GetWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
}

func (st GetWorkspaceNetworkOptionRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetWorkspaceNetworkOptionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetWorkspaceNetworkOptionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.GetWorkspaceNetworkOptionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetWorkspaceNetworkOptionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetWorkspaceNetworkOptionRequestToPb(st *GetWorkspaceNetworkOptionRequest) (*settingspb.GetWorkspaceNetworkOptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.GetWorkspaceNetworkOptionRequestPb{}
	pb.WorkspaceId = st.WorkspaceId

	return pb, nil
}

func GetWorkspaceNetworkOptionRequestFromPb(pb *settingspb.GetWorkspaceNetworkOptionRequestPb) (*GetWorkspaceNetworkOptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetWorkspaceNetworkOptionRequest{}
	st.WorkspaceId = pb.WorkspaceId

	return st, nil
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
	UpdatedBy       int64    `json:"updated_by,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st IpAccessListInfo) MarshalJSON() ([]byte, error) {
	pb, err := IpAccessListInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *IpAccessListInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.IpAccessListInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := IpAccessListInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func IpAccessListInfoToPb(st *IpAccessListInfo) (*settingspb.IpAccessListInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.IpAccessListInfoPb{}
	pb.AddressCount = st.AddressCount
	pb.CreatedAt = st.CreatedAt
	pb.CreatedBy = st.CreatedBy
	pb.Enabled = st.Enabled
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	pb.ListId = st.ListId
	listTypePb, err := ListTypeToPb(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}
	pb.UpdatedAt = st.UpdatedAt
	pb.UpdatedBy = st.UpdatedBy

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func IpAccessListInfoFromPb(pb *settingspb.IpAccessListInfoPb) (*IpAccessListInfo, error) {
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
	listTypeField, err := ListTypeFromPb(&pb.ListType)
	if err != nil {
		return nil, err
	}
	if listTypeField != nil {
		st.ListType = *listTypeField
	}
	st.UpdatedAt = pb.UpdatedAt
	st.UpdatedBy = pb.UpdatedBy

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {

	// Wire name: 'ip_access_lists'
	IpAccessLists []IpAccessListInfo `json:"ip_access_lists,omitempty"`
}

func (st ListIpAccessListResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListIpAccessListResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListIpAccessListResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListIpAccessListResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListIpAccessListResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListIpAccessListResponseToPb(st *ListIpAccessListResponse) (*settingspb.ListIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListIpAccessListResponsePb{}

	var ipAccessListsPb []settingspb.IpAccessListInfoPb
	for _, item := range st.IpAccessLists {
		itemPb, err := IpAccessListInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			ipAccessListsPb = append(ipAccessListsPb, *itemPb)
		}
	}
	pb.IpAccessLists = ipAccessListsPb

	return pb, nil
}

func ListIpAccessListResponseFromPb(pb *settingspb.ListIpAccessListResponsePb) (*ListIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListIpAccessListResponse{}

	var ipAccessListsField []IpAccessListInfo
	for _, itemPb := range pb.IpAccessLists {
		item, err := IpAccessListInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			ipAccessListsField = append(ipAccessListsField, *item)
		}
	}
	st.IpAccessLists = ipAccessListsField

	return st, nil
}

type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListNetworkConnectivityConfigurationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListNetworkConnectivityConfigurationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNetworkConnectivityConfigurationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListNetworkConnectivityConfigurationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNetworkConnectivityConfigurationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNetworkConnectivityConfigurationsRequestToPb(st *ListNetworkConnectivityConfigurationsRequest) (*settingspb.ListNetworkConnectivityConfigurationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListNetworkConnectivityConfigurationsRequestPb{}
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListNetworkConnectivityConfigurationsRequestFromPb(pb *settingspb.ListNetworkConnectivityConfigurationsRequestPb) (*ListNetworkConnectivityConfigurationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkConnectivityConfigurationsRequest{}
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The network connectivity configuration list was successfully retrieved.
type ListNetworkConnectivityConfigurationsResponse struct {

	// Wire name: 'items'
	Items []NetworkConnectivityConfiguration `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListNetworkConnectivityConfigurationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListNetworkConnectivityConfigurationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNetworkConnectivityConfigurationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListNetworkConnectivityConfigurationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNetworkConnectivityConfigurationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNetworkConnectivityConfigurationsResponseToPb(st *ListNetworkConnectivityConfigurationsResponse) (*settingspb.ListNetworkConnectivityConfigurationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListNetworkConnectivityConfigurationsResponsePb{}

	var itemsPb []settingspb.NetworkConnectivityConfigurationPb
	for _, item := range st.Items {
		itemPb, err := NetworkConnectivityConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			itemsPb = append(itemsPb, *itemPb)
		}
	}
	pb.Items = itemsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListNetworkConnectivityConfigurationsResponseFromPb(pb *settingspb.ListNetworkConnectivityConfigurationsResponsePb) (*ListNetworkConnectivityConfigurationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkConnectivityConfigurationsResponse{}

	var itemsField []NetworkConnectivityConfiguration
	for _, itemPb := range pb.Items {
		item, err := NetworkConnectivityConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			itemsField = append(itemsField, *item)
		}
	}
	st.Items = itemsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListNetworkPoliciesRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListNetworkPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListNetworkPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNetworkPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListNetworkPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNetworkPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNetworkPoliciesRequestToPb(st *ListNetworkPoliciesRequest) (*settingspb.ListNetworkPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListNetworkPoliciesRequestPb{}
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListNetworkPoliciesRequestFromPb(pb *settingspb.ListNetworkPoliciesRequestPb) (*ListNetworkPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkPoliciesRequest{}
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListNetworkPoliciesResponse struct {
	// List of network policies.
	// Wire name: 'items'
	Items []AccountNetworkPolicy `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListNetworkPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListNetworkPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNetworkPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListNetworkPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNetworkPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNetworkPoliciesResponseToPb(st *ListNetworkPoliciesResponse) (*settingspb.ListNetworkPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListNetworkPoliciesResponsePb{}

	var itemsPb []settingspb.AccountNetworkPolicyPb
	for _, item := range st.Items {
		itemPb, err := AccountNetworkPolicyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			itemsPb = append(itemsPb, *itemPb)
		}
	}
	pb.Items = itemsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListNetworkPoliciesResponseFromPb(pb *settingspb.ListNetworkPoliciesResponsePb) (*ListNetworkPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkPoliciesResponse{}

	var itemsField []AccountNetworkPolicy
	for _, itemPb := range pb.Items {
		item, err := AccountNetworkPolicyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			itemsField = append(itemsField, *item)
		}
	}
	st.Items = itemsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListNotificationDestinationsRequest struct {
	PageSize int64 `json:"-" tf:"-"`

	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListNotificationDestinationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListNotificationDestinationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNotificationDestinationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListNotificationDestinationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNotificationDestinationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNotificationDestinationsRequestToPb(st *ListNotificationDestinationsRequest) (*settingspb.ListNotificationDestinationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListNotificationDestinationsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListNotificationDestinationsRequestFromPb(pb *settingspb.ListNotificationDestinationsRequestPb) (*ListNotificationDestinationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListNotificationDestinationsResponse struct {
	// Page token for next of results.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'results'
	Results         []ListNotificationDestinationsResult `json:"results,omitempty"`
	ForceSendFields []string                             `json:"-" tf:"-"`
}

func (st ListNotificationDestinationsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListNotificationDestinationsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNotificationDestinationsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListNotificationDestinationsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNotificationDestinationsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNotificationDestinationsResponseToPb(st *ListNotificationDestinationsResponse) (*settingspb.ListNotificationDestinationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListNotificationDestinationsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var resultsPb []settingspb.ListNotificationDestinationsResultPb
	for _, item := range st.Results {
		itemPb, err := ListNotificationDestinationsResultToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListNotificationDestinationsResponseFromPb(pb *settingspb.ListNotificationDestinationsResponsePb) (*ListNotificationDestinationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsResponse{}
	st.NextPageToken = pb.NextPageToken

	var resultsField []ListNotificationDestinationsResult
	for _, itemPb := range pb.Results {
		item, err := ListNotificationDestinationsResultFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	Id              string   `json:"id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListNotificationDestinationsResult) MarshalJSON() ([]byte, error) {
	pb, err := ListNotificationDestinationsResultToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListNotificationDestinationsResult) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListNotificationDestinationsResultPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListNotificationDestinationsResultFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListNotificationDestinationsResultToPb(st *ListNotificationDestinationsResult) (*settingspb.ListNotificationDestinationsResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListNotificationDestinationsResultPb{}
	destinationTypePb, err := DestinationTypeToPb(&st.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypePb != nil {
		pb.DestinationType = *destinationTypePb
	}
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListNotificationDestinationsResultFromPb(pb *settingspb.ListNotificationDestinationsResultPb) (*ListNotificationDestinationsResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsResult{}
	destinationTypeField, err := DestinationTypeFromPb(&pb.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypeField != nil {
		st.DestinationType = *destinationTypeField
	}
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" tf:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListPrivateEndpointRulesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListPrivateEndpointRulesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListPrivateEndpointRulesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListPrivateEndpointRulesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListPrivateEndpointRulesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListPrivateEndpointRulesRequestToPb(st *ListPrivateEndpointRulesRequest) (*settingspb.ListPrivateEndpointRulesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListPrivateEndpointRulesRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListPrivateEndpointRulesRequestFromPb(pb *settingspb.ListPrivateEndpointRulesRequestPb) (*ListPrivateEndpointRulesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPrivateEndpointRulesRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The private endpoint rule list was successfully retrieved.
type ListPrivateEndpointRulesResponse struct {

	// Wire name: 'items'
	Items []NccPrivateEndpointRule `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListPrivateEndpointRulesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListPrivateEndpointRulesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListPrivateEndpointRulesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListPrivateEndpointRulesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListPrivateEndpointRulesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListPrivateEndpointRulesResponseToPb(st *ListPrivateEndpointRulesResponse) (*settingspb.ListPrivateEndpointRulesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListPrivateEndpointRulesResponsePb{}

	var itemsPb []settingspb.NccPrivateEndpointRulePb
	for _, item := range st.Items {
		itemPb, err := NccPrivateEndpointRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			itemsPb = append(itemsPb, *itemPb)
		}
	}
	pb.Items = itemsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListPrivateEndpointRulesResponseFromPb(pb *settingspb.ListPrivateEndpointRulesResponsePb) (*ListPrivateEndpointRulesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPrivateEndpointRulesResponse{}

	var itemsField []NccPrivateEndpointRule
	for _, itemPb := range pb.Items {
		item, err := NccPrivateEndpointRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			itemsField = append(itemsField, *item)
		}
	}
	st.Items = itemsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListPublicTokensResponse struct {
	// The information for each token.
	// Wire name: 'token_infos'
	TokenInfos []PublicTokenInfo `json:"token_infos,omitempty"`
}

func (st ListPublicTokensResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListPublicTokensResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListPublicTokensResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListPublicTokensResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListPublicTokensResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListPublicTokensResponseToPb(st *ListPublicTokensResponse) (*settingspb.ListPublicTokensResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListPublicTokensResponsePb{}

	var tokenInfosPb []settingspb.PublicTokenInfoPb
	for _, item := range st.TokenInfos {
		itemPb, err := PublicTokenInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tokenInfosPb = append(tokenInfosPb, *itemPb)
		}
	}
	pb.TokenInfos = tokenInfosPb

	return pb, nil
}

func ListPublicTokensResponseFromPb(pb *settingspb.ListPublicTokensResponsePb) (*ListPublicTokensResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPublicTokensResponse{}

	var tokenInfosField []PublicTokenInfo
	for _, itemPb := range pb.TokenInfos {
		item, err := PublicTokenInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tokenInfosField = append(tokenInfosField, *item)
		}
	}
	st.TokenInfos = tokenInfosField

	return st, nil
}

type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById int64 `json:"-" tf:"-"`
	// Username of the user that created the token.
	CreatedByUsername string   `json:"-" tf:"-"`
	ForceSendFields   []string `json:"-" tf:"-"`
}

func (st ListTokenManagementRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListTokenManagementRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListTokenManagementRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListTokenManagementRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListTokenManagementRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListTokenManagementRequestToPb(st *ListTokenManagementRequest) (*settingspb.ListTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListTokenManagementRequestPb{}
	pb.CreatedById = st.CreatedById
	pb.CreatedByUsername = st.CreatedByUsername

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListTokenManagementRequestFromPb(pb *settingspb.ListTokenManagementRequestPb) (*ListTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTokenManagementRequest{}
	st.CreatedById = pb.CreatedById
	st.CreatedByUsername = pb.CreatedByUsername

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	// Wire name: 'token_infos'
	TokenInfos []TokenInfo `json:"token_infos,omitempty"`
}

func (st ListTokensResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListTokensResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListTokensResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ListTokensResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListTokensResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListTokensResponseToPb(st *ListTokensResponse) (*settingspb.ListTokensResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ListTokensResponsePb{}

	var tokenInfosPb []settingspb.TokenInfoPb
	for _, item := range st.TokenInfos {
		itemPb, err := TokenInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tokenInfosPb = append(tokenInfosPb, *itemPb)
		}
	}
	pb.TokenInfos = tokenInfosPb

	return pb, nil
}

func ListTokensResponseFromPb(pb *settingspb.ListTokensResponsePb) (*ListTokensResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTokensResponse{}

	var tokenInfosField []TokenInfo
	for _, itemPb := range pb.TokenInfos {
		item, err := TokenInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tokenInfosField = append(tokenInfosField, *item)
		}
	}
	st.TokenInfos = tokenInfosField

	return st, nil
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

func ListTypeToPb(st *ListType) (*settingspb.ListTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.ListTypePb(*st)
	return &pb, nil
}

func ListTypeFromPb(pb *settingspb.ListTypePb) (*ListType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListType(*pb)
	return &st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LlmProxyPartnerPoweredAccount) MarshalJSON() ([]byte, error) {
	pb, err := LlmProxyPartnerPoweredAccountToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LlmProxyPartnerPoweredAccount) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.LlmProxyPartnerPoweredAccountPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LlmProxyPartnerPoweredAccountFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LlmProxyPartnerPoweredAccountToPb(st *LlmProxyPartnerPoweredAccount) (*settingspb.LlmProxyPartnerPoweredAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.LlmProxyPartnerPoweredAccountPb{}
	booleanValPb, err := BooleanMessageToPb(&st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = *booleanValPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LlmProxyPartnerPoweredAccountFromPb(pb *settingspb.LlmProxyPartnerPoweredAccountPb) (*LlmProxyPartnerPoweredAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LlmProxyPartnerPoweredAccount{}
	booleanValField, err := BooleanMessageFromPb(&pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = *booleanValField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LlmProxyPartnerPoweredEnforce) MarshalJSON() ([]byte, error) {
	pb, err := LlmProxyPartnerPoweredEnforceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LlmProxyPartnerPoweredEnforce) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.LlmProxyPartnerPoweredEnforcePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LlmProxyPartnerPoweredEnforceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LlmProxyPartnerPoweredEnforceToPb(st *LlmProxyPartnerPoweredEnforce) (*settingspb.LlmProxyPartnerPoweredEnforcePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.LlmProxyPartnerPoweredEnforcePb{}
	booleanValPb, err := BooleanMessageToPb(&st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = *booleanValPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LlmProxyPartnerPoweredEnforceFromPb(pb *settingspb.LlmProxyPartnerPoweredEnforcePb) (*LlmProxyPartnerPoweredEnforce, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LlmProxyPartnerPoweredEnforce{}
	booleanValField, err := BooleanMessageFromPb(&pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = *booleanValField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st LlmProxyPartnerPoweredWorkspace) MarshalJSON() ([]byte, error) {
	pb, err := LlmProxyPartnerPoweredWorkspaceToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LlmProxyPartnerPoweredWorkspace) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.LlmProxyPartnerPoweredWorkspacePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LlmProxyPartnerPoweredWorkspaceFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LlmProxyPartnerPoweredWorkspaceToPb(st *LlmProxyPartnerPoweredWorkspace) (*settingspb.LlmProxyPartnerPoweredWorkspacePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.LlmProxyPartnerPoweredWorkspacePb{}
	booleanValPb, err := BooleanMessageToPb(&st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = *booleanValPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LlmProxyPartnerPoweredWorkspaceFromPb(pb *settingspb.LlmProxyPartnerPoweredWorkspacePb) (*LlmProxyPartnerPoweredWorkspace, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LlmProxyPartnerPoweredWorkspace{}
	booleanValField, err := BooleanMessageFromPb(&pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = *booleanValField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] URL for Microsoft Teams.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet          bool     `json:"url_set,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st MicrosoftTeamsConfig) MarshalJSON() ([]byte, error) {
	pb, err := MicrosoftTeamsConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *MicrosoftTeamsConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.MicrosoftTeamsConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MicrosoftTeamsConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func MicrosoftTeamsConfigToPb(st *MicrosoftTeamsConfig) (*settingspb.MicrosoftTeamsConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.MicrosoftTeamsConfigPb{}
	pb.Url = st.Url
	pb.UrlSet = st.UrlSet

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func MicrosoftTeamsConfigFromPb(pb *settingspb.MicrosoftTeamsConfigPb) (*MicrosoftTeamsConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MicrosoftTeamsConfig{}
	st.Url = pb.Url
	st.UrlSet = pb.UrlSet

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	// Wire name: 'cidr_blocks'
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

func (st NccAwsStableIpRule) MarshalJSON() ([]byte, error) {
	pb, err := NccAwsStableIpRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NccAwsStableIpRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NccAwsStableIpRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NccAwsStableIpRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NccAwsStableIpRuleToPb(st *NccAwsStableIpRule) (*settingspb.NccAwsStableIpRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NccAwsStableIpRulePb{}
	pb.CidrBlocks = st.CidrBlocks

	return pb, nil
}

func NccAwsStableIpRuleFromPb(pb *settingspb.NccAwsStableIpRulePb) (*NccAwsStableIpRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAwsStableIpRule{}
	st.CidrBlocks = pb.CidrBlocks

	return st, nil
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
	UpdatedTime     int64    `json:"updated_time,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NccAzurePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := NccAzurePrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NccAzurePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NccAzurePrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NccAzurePrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NccAzurePrivateEndpointRuleToPb(st *NccAzurePrivateEndpointRule) (*settingspb.NccAzurePrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NccAzurePrivateEndpointRulePb{}
	connectionStatePb, err := NccAzurePrivateEndpointRuleConnectionStateToPb(&st.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStatePb != nil {
		pb.ConnectionState = *connectionStatePb
	}
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NccAzurePrivateEndpointRuleFromPb(pb *settingspb.NccAzurePrivateEndpointRulePb) (*NccAzurePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAzurePrivateEndpointRule{}
	connectionStateField, err := NccAzurePrivateEndpointRuleConnectionStateFromPb(&pb.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStateField != nil {
		st.ConnectionState = *connectionStateField
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func NccAzurePrivateEndpointRuleConnectionStateToPb(st *NccAzurePrivateEndpointRuleConnectionState) (*settingspb.NccAzurePrivateEndpointRuleConnectionStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.NccAzurePrivateEndpointRuleConnectionStatePb(*st)
	return &pb, nil
}

func NccAzurePrivateEndpointRuleConnectionStateFromPb(pb *settingspb.NccAzurePrivateEndpointRuleConnectionStatePb) (*NccAzurePrivateEndpointRuleConnectionState, error) {
	if pb == nil {
		return nil, nil
	}
	st := NccAzurePrivateEndpointRuleConnectionState(*pb)
	return &st, nil
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
	TargetServices  []EgressResourceType `json:"target_services,omitempty"`
	ForceSendFields []string             `json:"-" tf:"-"`
}

func (st NccAzureServiceEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := NccAzureServiceEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NccAzureServiceEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NccAzureServiceEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NccAzureServiceEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NccAzureServiceEndpointRuleToPb(st *NccAzureServiceEndpointRule) (*settingspb.NccAzureServiceEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NccAzureServiceEndpointRulePb{}
	pb.Subnets = st.Subnets
	pb.TargetRegion = st.TargetRegion

	var targetServicesPb []settingspb.EgressResourceTypePb
	for _, item := range st.TargetServices {
		itemPb, err := EgressResourceTypeToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			targetServicesPb = append(targetServicesPb, *itemPb)
		}
	}
	pb.TargetServices = targetServicesPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NccAzureServiceEndpointRuleFromPb(pb *settingspb.NccAzureServiceEndpointRulePb) (*NccAzureServiceEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAzureServiceEndpointRule{}
	st.Subnets = pb.Subnets
	st.TargetRegion = pb.TargetRegion

	var targetServicesField []EgressResourceType
	for _, itemPb := range pb.TargetServices {
		item, err := EgressResourceTypeFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			targetServicesField = append(targetServicesField, *item)
		}
	}
	st.TargetServices = targetServicesField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st NccEgressConfig) MarshalJSON() ([]byte, error) {
	pb, err := NccEgressConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NccEgressConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NccEgressConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NccEgressConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NccEgressConfigToPb(st *NccEgressConfig) (*settingspb.NccEgressConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NccEgressConfigPb{}
	defaultRulesPb, err := NccEgressDefaultRulesToPb(st.DefaultRules)
	if err != nil {
		return nil, err
	}
	if defaultRulesPb != nil {
		pb.DefaultRules = defaultRulesPb
	}
	targetRulesPb, err := NccEgressTargetRulesToPb(st.TargetRules)
	if err != nil {
		return nil, err
	}
	if targetRulesPb != nil {
		pb.TargetRules = targetRulesPb
	}

	return pb, nil
}

func NccEgressConfigFromPb(pb *settingspb.NccEgressConfigPb) (*NccEgressConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressConfig{}
	defaultRulesField, err := NccEgressDefaultRulesFromPb(pb.DefaultRules)
	if err != nil {
		return nil, err
	}
	if defaultRulesField != nil {
		st.DefaultRules = defaultRulesField
	}
	targetRulesField, err := NccEgressTargetRulesFromPb(pb.TargetRules)
	if err != nil {
		return nil, err
	}
	if targetRulesField != nil {
		st.TargetRules = targetRulesField
	}

	return st, nil
}

// Default rules don't have specific targets.
type NccEgressDefaultRules struct {

	// Wire name: 'aws_stable_ip_rule'
	AwsStableIpRule *NccAwsStableIpRule `json:"aws_stable_ip_rule,omitempty"`

	// Wire name: 'azure_service_endpoint_rule'
	AzureServiceEndpointRule *NccAzureServiceEndpointRule `json:"azure_service_endpoint_rule,omitempty"`
}

func (st NccEgressDefaultRules) MarshalJSON() ([]byte, error) {
	pb, err := NccEgressDefaultRulesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NccEgressDefaultRules) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NccEgressDefaultRulesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NccEgressDefaultRulesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NccEgressDefaultRulesToPb(st *NccEgressDefaultRules) (*settingspb.NccEgressDefaultRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NccEgressDefaultRulesPb{}
	awsStableIpRulePb, err := NccAwsStableIpRuleToPb(st.AwsStableIpRule)
	if err != nil {
		return nil, err
	}
	if awsStableIpRulePb != nil {
		pb.AwsStableIpRule = awsStableIpRulePb
	}
	azureServiceEndpointRulePb, err := NccAzureServiceEndpointRuleToPb(st.AzureServiceEndpointRule)
	if err != nil {
		return nil, err
	}
	if azureServiceEndpointRulePb != nil {
		pb.AzureServiceEndpointRule = azureServiceEndpointRulePb
	}

	return pb, nil
}

func NccEgressDefaultRulesFromPb(pb *settingspb.NccEgressDefaultRulesPb) (*NccEgressDefaultRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressDefaultRules{}
	awsStableIpRuleField, err := NccAwsStableIpRuleFromPb(pb.AwsStableIpRule)
	if err != nil {
		return nil, err
	}
	if awsStableIpRuleField != nil {
		st.AwsStableIpRule = awsStableIpRuleField
	}
	azureServiceEndpointRuleField, err := NccAzureServiceEndpointRuleFromPb(pb.AzureServiceEndpointRule)
	if err != nil {
		return nil, err
	}
	if azureServiceEndpointRuleField != nil {
		st.AzureServiceEndpointRule = azureServiceEndpointRuleField
	}

	return st, nil
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

func (st NccEgressTargetRules) MarshalJSON() ([]byte, error) {
	pb, err := NccEgressTargetRulesToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NccEgressTargetRules) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NccEgressTargetRulesPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NccEgressTargetRulesFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NccEgressTargetRulesToPb(st *NccEgressTargetRules) (*settingspb.NccEgressTargetRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NccEgressTargetRulesPb{}

	var awsPrivateEndpointRulesPb []settingspb.CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRulePb
	for _, item := range st.AwsPrivateEndpointRules {
		itemPb, err := CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			awsPrivateEndpointRulesPb = append(awsPrivateEndpointRulesPb, *itemPb)
		}
	}
	pb.AwsPrivateEndpointRules = awsPrivateEndpointRulesPb

	var azurePrivateEndpointRulesPb []settingspb.NccAzurePrivateEndpointRulePb
	for _, item := range st.AzurePrivateEndpointRules {
		itemPb, err := NccAzurePrivateEndpointRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			azurePrivateEndpointRulesPb = append(azurePrivateEndpointRulesPb, *itemPb)
		}
	}
	pb.AzurePrivateEndpointRules = azurePrivateEndpointRulesPb

	return pb, nil
}

func NccEgressTargetRulesFromPb(pb *settingspb.NccEgressTargetRulesPb) (*NccEgressTargetRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressTargetRules{}

	var awsPrivateEndpointRulesField []CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRule
	for _, itemPb := range pb.AwsPrivateEndpointRules {
		item, err := CustomerFacingNetworkConnectivityConfigAwsPrivateEndpointRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			awsPrivateEndpointRulesField = append(awsPrivateEndpointRulesField, *item)
		}
	}
	st.AwsPrivateEndpointRules = awsPrivateEndpointRulesField

	var azurePrivateEndpointRulesField []NccAzurePrivateEndpointRule
	for _, itemPb := range pb.AzurePrivateEndpointRules {
		item, err := NccAzurePrivateEndpointRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			azurePrivateEndpointRulesField = append(azurePrivateEndpointRulesField, *item)
		}
	}
	st.AzurePrivateEndpointRules = azurePrivateEndpointRulesField

	return st, nil
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
	VpcEndpointId   string   `json:"vpc_endpoint_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NccPrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := NccPrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NccPrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NccPrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NccPrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NccPrivateEndpointRuleToPb(st *NccPrivateEndpointRule) (*settingspb.NccPrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NccPrivateEndpointRulePb{}
	pb.AccountId = st.AccountId
	connectionStatePb, err := NccPrivateEndpointRulePrivateLinkConnectionStateToPb(&st.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStatePb != nil {
		pb.ConnectionState = *connectionStatePb
	}
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NccPrivateEndpointRuleFromPb(pb *settingspb.NccPrivateEndpointRulePb) (*NccPrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccPrivateEndpointRule{}
	st.AccountId = pb.AccountId
	connectionStateField, err := NccPrivateEndpointRulePrivateLinkConnectionStateFromPb(&pb.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStateField != nil {
		st.ConnectionState = *connectionStateField
	}
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func NccPrivateEndpointRulePrivateLinkConnectionStateToPb(st *NccPrivateEndpointRulePrivateLinkConnectionState) (*settingspb.NccPrivateEndpointRulePrivateLinkConnectionStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.NccPrivateEndpointRulePrivateLinkConnectionStatePb(*st)
	return &pb, nil
}

func NccPrivateEndpointRulePrivateLinkConnectionStateFromPb(pb *settingspb.NccPrivateEndpointRulePrivateLinkConnectionStatePb) (*NccPrivateEndpointRulePrivateLinkConnectionState, error) {
	if pb == nil {
		return nil, nil
	}
	st := NccPrivateEndpointRulePrivateLinkConnectionState(*pb)
	return &st, nil
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
	UpdatedTime     int64    `json:"updated_time,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NetworkConnectivityConfiguration) MarshalJSON() ([]byte, error) {
	pb, err := NetworkConnectivityConfigurationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NetworkConnectivityConfiguration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NetworkConnectivityConfigurationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NetworkConnectivityConfigurationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NetworkConnectivityConfigurationToPb(st *NetworkConnectivityConfiguration) (*settingspb.NetworkConnectivityConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NetworkConnectivityConfigurationPb{}
	pb.AccountId = st.AccountId
	pb.CreationTime = st.CreationTime
	egressConfigPb, err := NccEgressConfigToPb(st.EgressConfig)
	if err != nil {
		return nil, err
	}
	if egressConfigPb != nil {
		pb.EgressConfig = egressConfigPb
	}
	pb.Name = st.Name
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	pb.Region = st.Region
	pb.UpdatedTime = st.UpdatedTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NetworkConnectivityConfigurationFromPb(pb *settingspb.NetworkConnectivityConfigurationPb) (*NetworkConnectivityConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkConnectivityConfiguration{}
	st.AccountId = pb.AccountId
	st.CreationTime = pb.CreationTime
	egressConfigField, err := NccEgressConfigFromPb(pb.EgressConfig)
	if err != nil {
		return nil, err
	}
	if egressConfigField != nil {
		st.EgressConfig = egressConfigField
	}
	st.Name = pb.Name
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	st.Region = pb.Region
	st.UpdatedTime = pb.UpdatedTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st NetworkPolicyEgress) MarshalJSON() ([]byte, error) {
	pb, err := NetworkPolicyEgressToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NetworkPolicyEgress) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NetworkPolicyEgressPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NetworkPolicyEgressFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NetworkPolicyEgressToPb(st *NetworkPolicyEgress) (*settingspb.NetworkPolicyEgressPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NetworkPolicyEgressPb{}
	networkAccessPb, err := EgressNetworkPolicyNetworkAccessPolicyToPb(st.NetworkAccess)
	if err != nil {
		return nil, err
	}
	if networkAccessPb != nil {
		pb.NetworkAccess = networkAccessPb
	}

	return pb, nil
}

func NetworkPolicyEgressFromPb(pb *settingspb.NetworkPolicyEgressPb) (*NetworkPolicyEgress, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkPolicyEgress{}
	networkAccessField, err := EgressNetworkPolicyNetworkAccessPolicyFromPb(pb.NetworkAccess)
	if err != nil {
		return nil, err
	}
	if networkAccessField != nil {
		st.NetworkAccess = networkAccessField
	}

	return st, nil
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
	Id              string   `json:"id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st NotificationDestination) MarshalJSON() ([]byte, error) {
	pb, err := NotificationDestinationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *NotificationDestination) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.NotificationDestinationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := NotificationDestinationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func NotificationDestinationToPb(st *NotificationDestination) (*settingspb.NotificationDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.NotificationDestinationPb{}
	configPb, err := ConfigToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}
	destinationTypePb, err := DestinationTypeToPb(&st.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypePb != nil {
		pb.DestinationType = *destinationTypePb
	}
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func NotificationDestinationFromPb(pb *settingspb.NotificationDestinationPb) (*NotificationDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotificationDestination{}
	configField, err := ConfigFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	destinationTypeField, err := DestinationTypeFromPb(&pb.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypeField != nil {
		st.DestinationType = *destinationTypeField
	}
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	// Wire name: 'integration_key'
	IntegrationKey string `json:"integration_key,omitempty"`
	// [Output-Only] Whether integration key is set.
	// Wire name: 'integration_key_set'
	IntegrationKeySet bool     `json:"integration_key_set,omitempty"`
	ForceSendFields   []string `json:"-" tf:"-"`
}

func (st PagerdutyConfig) MarshalJSON() ([]byte, error) {
	pb, err := PagerdutyConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PagerdutyConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.PagerdutyConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PagerdutyConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PagerdutyConfigToPb(st *PagerdutyConfig) (*settingspb.PagerdutyConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.PagerdutyConfigPb{}
	pb.IntegrationKey = st.IntegrationKey
	pb.IntegrationKeySet = st.IntegrationKeySet

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PagerdutyConfigFromPb(pb *settingspb.PagerdutyConfigPb) (*PagerdutyConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PagerdutyConfig{}
	st.IntegrationKey = pb.IntegrationKey
	st.IntegrationKeySet = pb.IntegrationKeySet

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	// Wire name: 'workspaceId'
	WorkspaceId     int64    `json:"workspaceId,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PartitionId) MarshalJSON() ([]byte, error) {
	pb, err := PartitionIdToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PartitionId) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.PartitionIdPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PartitionIdFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PartitionIdToPb(st *PartitionId) (*settingspb.PartitionIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.PartitionIdPb{}
	pb.WorkspaceId = st.WorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PartitionIdFromPb(pb *settingspb.PartitionIdPb) (*PartitionId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartitionId{}
	st.WorkspaceId = pb.WorkspaceId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PersonalComputeMessage struct {

	// Wire name: 'value'
	Value PersonalComputeMessageEnum `json:"value"`
}

func (st PersonalComputeMessage) MarshalJSON() ([]byte, error) {
	pb, err := PersonalComputeMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PersonalComputeMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.PersonalComputeMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PersonalComputeMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PersonalComputeMessageToPb(st *PersonalComputeMessage) (*settingspb.PersonalComputeMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.PersonalComputeMessagePb{}
	valuePb, err := PersonalComputeMessageEnumToPb(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	return pb, nil
}

func PersonalComputeMessageFromPb(pb *settingspb.PersonalComputeMessagePb) (*PersonalComputeMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalComputeMessage{}
	valueField, err := PersonalComputeMessageEnumFromPb(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

	return st, nil
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

func PersonalComputeMessageEnumToPb(st *PersonalComputeMessageEnum) (*settingspb.PersonalComputeMessageEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.PersonalComputeMessageEnumPb(*st)
	return &pb, nil
}

func PersonalComputeMessageEnumFromPb(pb *settingspb.PersonalComputeMessageEnumPb) (*PersonalComputeMessageEnum, error) {
	if pb == nil {
		return nil, nil
	}
	st := PersonalComputeMessageEnum(*pb)
	return &st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PersonalComputeSetting) MarshalJSON() ([]byte, error) {
	pb, err := PersonalComputeSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PersonalComputeSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.PersonalComputeSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PersonalComputeSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PersonalComputeSettingToPb(st *PersonalComputeSetting) (*settingspb.PersonalComputeSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.PersonalComputeSettingPb{}
	pb.Etag = st.Etag
	personalComputePb, err := PersonalComputeMessageToPb(&st.PersonalCompute)
	if err != nil {
		return nil, err
	}
	if personalComputePb != nil {
		pb.PersonalCompute = *personalComputePb
	}
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PersonalComputeSettingFromPb(pb *settingspb.PersonalComputeSettingPb) (*PersonalComputeSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalComputeSetting{}
	st.Etag = pb.Etag
	personalComputeField, err := PersonalComputeMessageFromPb(&pb.PersonalCompute)
	if err != nil {
		return nil, err
	}
	if personalComputeField != nil {
		st.PersonalCompute = *personalComputeField
	}
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	TokenId         string   `json:"token_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PublicTokenInfo) MarshalJSON() ([]byte, error) {
	pb, err := PublicTokenInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PublicTokenInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.PublicTokenInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PublicTokenInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PublicTokenInfoToPb(st *PublicTokenInfo) (*settingspb.PublicTokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.PublicTokenInfoPb{}
	pb.Comment = st.Comment
	pb.CreationTime = st.CreationTime
	pb.ExpiryTime = st.ExpiryTime
	pb.TokenId = st.TokenId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PublicTokenInfoFromPb(pb *settingspb.PublicTokenInfoPb) (*PublicTokenInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublicTokenInfo{}
	st.Comment = pb.Comment
	st.CreationTime = pb.CreationTime
	st.ExpiryTime = pb.ExpiryTime
	st.TokenId = pb.TokenId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func (st ReplaceIpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := ReplaceIpAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ReplaceIpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.ReplaceIpAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ReplaceIpAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ReplaceIpAccessListToPb(st *ReplaceIpAccessList) (*settingspb.ReplaceIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.ReplaceIpAccessListPb{}
	pb.Enabled = st.Enabled
	pb.IpAccessListId = st.IpAccessListId
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	listTypePb, err := ListTypeToPb(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}

	return pb, nil
}

func ReplaceIpAccessListFromPb(pb *settingspb.ReplaceIpAccessListPb) (*ReplaceIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReplaceIpAccessList{}
	st.Enabled = pb.Enabled
	st.IpAccessListId = pb.IpAccessListId
	st.IpAddresses = pb.IpAddresses
	st.Label = pb.Label
	listTypeField, err := ListTypeFromPb(&pb.ListType)
	if err != nil {
		return nil, err
	}
	if listTypeField != nil {
		st.ListType = *listTypeField
	}

	return st, nil
}

type RestrictWorkspaceAdminsMessage struct {

	// Wire name: 'status'
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`
}

func (st RestrictWorkspaceAdminsMessage) MarshalJSON() ([]byte, error) {
	pb, err := RestrictWorkspaceAdminsMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RestrictWorkspaceAdminsMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.RestrictWorkspaceAdminsMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RestrictWorkspaceAdminsMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RestrictWorkspaceAdminsMessageToPb(st *RestrictWorkspaceAdminsMessage) (*settingspb.RestrictWorkspaceAdminsMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.RestrictWorkspaceAdminsMessagePb{}
	statusPb, err := RestrictWorkspaceAdminsMessageStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	return pb, nil
}

func RestrictWorkspaceAdminsMessageFromPb(pb *settingspb.RestrictWorkspaceAdminsMessagePb) (*RestrictWorkspaceAdminsMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestrictWorkspaceAdminsMessage{}
	statusField, err := RestrictWorkspaceAdminsMessageStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
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

func RestrictWorkspaceAdminsMessageStatusToPb(st *RestrictWorkspaceAdminsMessageStatus) (*settingspb.RestrictWorkspaceAdminsMessageStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.RestrictWorkspaceAdminsMessageStatusPb(*st)
	return &pb, nil
}

func RestrictWorkspaceAdminsMessageStatusFromPb(pb *settingspb.RestrictWorkspaceAdminsMessageStatusPb) (*RestrictWorkspaceAdminsMessageStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := RestrictWorkspaceAdminsMessageStatus(*pb)
	return &st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st RestrictWorkspaceAdminsSetting) MarshalJSON() ([]byte, error) {
	pb, err := RestrictWorkspaceAdminsSettingToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RestrictWorkspaceAdminsSetting) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.RestrictWorkspaceAdminsSettingPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RestrictWorkspaceAdminsSettingFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RestrictWorkspaceAdminsSettingToPb(st *RestrictWorkspaceAdminsSetting) (*settingspb.RestrictWorkspaceAdminsSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.RestrictWorkspaceAdminsSettingPb{}
	pb.Etag = st.Etag
	restrictWorkspaceAdminsPb, err := RestrictWorkspaceAdminsMessageToPb(&st.RestrictWorkspaceAdmins)
	if err != nil {
		return nil, err
	}
	if restrictWorkspaceAdminsPb != nil {
		pb.RestrictWorkspaceAdmins = *restrictWorkspaceAdminsPb
	}
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RestrictWorkspaceAdminsSettingFromPb(pb *settingspb.RestrictWorkspaceAdminsSettingPb) (*RestrictWorkspaceAdminsSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestrictWorkspaceAdminsSetting{}
	st.Etag = pb.Etag
	restrictWorkspaceAdminsField, err := RestrictWorkspaceAdminsMessageFromPb(&pb.RestrictWorkspaceAdmins)
	if err != nil {
		return nil, err
	}
	if restrictWorkspaceAdminsField != nil {
		st.RestrictWorkspaceAdmins = *restrictWorkspaceAdminsField
	}
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	// Wire name: 'token_id'
	TokenId string `json:"token_id"`
}

func (st RevokeTokenRequest) MarshalJSON() ([]byte, error) {
	pb, err := RevokeTokenRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RevokeTokenRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.RevokeTokenRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RevokeTokenRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RevokeTokenRequestToPb(st *RevokeTokenRequest) (*settingspb.RevokeTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.RevokeTokenRequestPb{}
	pb.TokenId = st.TokenId

	return pb, nil
}

func RevokeTokenRequestFromPb(pb *settingspb.RevokeTokenRequestPb) (*RevokeTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RevokeTokenRequest{}
	st.TokenId = pb.TokenId

	return st, nil
}

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	// Wire name: 'url'
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	// Wire name: 'url_set'
	UrlSet          bool     `json:"url_set,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SlackConfig) MarshalJSON() ([]byte, error) {
	pb, err := SlackConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SlackConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.SlackConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SlackConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SlackConfigToPb(st *SlackConfig) (*settingspb.SlackConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.SlackConfigPb{}
	pb.Url = st.Url
	pb.UrlSet = st.UrlSet

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SlackConfigFromPb(pb *settingspb.SlackConfigPb) (*SlackConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SlackConfig{}
	st.Url = pb.Url
	st.UrlSet = pb.UrlSet

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	SettingName     string   `json:"setting_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SqlResultsDownload) MarshalJSON() ([]byte, error) {
	pb, err := SqlResultsDownloadToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SqlResultsDownload) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.SqlResultsDownloadPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SqlResultsDownloadFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SqlResultsDownloadToPb(st *SqlResultsDownload) (*settingspb.SqlResultsDownloadPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.SqlResultsDownloadPb{}
	booleanValPb, err := BooleanMessageToPb(&st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = *booleanValPb
	}
	pb.Etag = st.Etag
	pb.SettingName = st.SettingName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SqlResultsDownloadFromPb(pb *settingspb.SqlResultsDownloadPb) (*SqlResultsDownload, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SqlResultsDownload{}
	booleanValField, err := BooleanMessageFromPb(&pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = *booleanValField
	}
	st.Etag = pb.Etag
	st.SettingName = pb.SettingName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type StringMessage struct {
	// Represents a generic string value.
	// Wire name: 'value'
	Value           string   `json:"value,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st StringMessage) MarshalJSON() ([]byte, error) {
	pb, err := StringMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *StringMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.StringMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := StringMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func StringMessageToPb(st *StringMessage) (*settingspb.StringMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.StringMessagePb{}
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func StringMessageFromPb(pb *settingspb.StringMessagePb) (*StringMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StringMessage{}
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TokenAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := TokenAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.TokenAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenAccessControlRequestToPb(st *TokenAccessControlRequest) (*settingspb.TokenAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.TokenAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := TokenPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TokenAccessControlRequestFromPb(pb *settingspb.TokenAccessControlRequestPb) (*TokenAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := TokenPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	UserName        string   `json:"user_name,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TokenAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := TokenAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.TokenAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenAccessControlResponseToPb(st *TokenAccessControlResponse) (*settingspb.TokenAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.TokenAccessControlResponsePb{}

	var allPermissionsPb []settingspb.TokenPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := TokenPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TokenAccessControlResponseFromPb(pb *settingspb.TokenAccessControlResponsePb) (*TokenAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessControlResponse{}

	var allPermissionsField []TokenPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := TokenPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	WorkspaceId     int64    `json:"workspace_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TokenInfo) MarshalJSON() ([]byte, error) {
	pb, err := TokenInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.TokenInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenInfoToPb(st *TokenInfo) (*settingspb.TokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.TokenInfoPb{}
	pb.Comment = st.Comment
	pb.CreatedById = st.CreatedById
	pb.CreatedByUsername = st.CreatedByUsername
	pb.CreationTime = st.CreationTime
	pb.ExpiryTime = st.ExpiryTime
	pb.LastUsedDay = st.LastUsedDay
	pb.OwnerId = st.OwnerId
	pb.TokenId = st.TokenId
	pb.WorkspaceId = st.WorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TokenInfoFromPb(pb *settingspb.TokenInfoPb) (*TokenInfo, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TokenPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string             `json:"-" tf:"-"`
}

func (st TokenPermission) MarshalJSON() ([]byte, error) {
	pb, err := TokenPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.TokenPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenPermissionToPb(st *TokenPermission) (*settingspb.TokenPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.TokenPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := TokenPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TokenPermissionFromPb(pb *settingspb.TokenPermissionPb) (*TokenPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := TokenPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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

func TokenPermissionLevelToPb(st *TokenPermissionLevel) (*settingspb.TokenPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.TokenPermissionLevelPb(*st)
	return &pb, nil
}

func TokenPermissionLevelFromPb(pb *settingspb.TokenPermissionLevelPb) (*TokenPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := TokenPermissionLevel(*pb)
	return &st, nil
}

type TokenPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []TokenAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType      string   `json:"object_type,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st TokenPermissions) MarshalJSON() ([]byte, error) {
	pb, err := TokenPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.TokenPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenPermissionsToPb(st *TokenPermissions) (*settingspb.TokenPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.TokenPermissionsPb{}

	var accessControlListPb []settingspb.TokenAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := TokenAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TokenPermissionsFromPb(pb *settingspb.TokenPermissionsPb) (*TokenPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissions{}

	var accessControlListField []TokenAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := TokenAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TokenPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`
	ForceSendFields []string             `json:"-" tf:"-"`
}

func (st TokenPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := TokenPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.TokenPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenPermissionsDescriptionToPb(st *TokenPermissionsDescription) (*settingspb.TokenPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.TokenPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := TokenPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TokenPermissionsDescriptionFromPb(pb *settingspb.TokenPermissionsDescriptionPb) (*TokenPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := TokenPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TokenPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []TokenAccessControlRequest `json:"access_control_list,omitempty"`
}

func (st TokenPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := TokenPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.TokenPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenPermissionsRequestToPb(st *TokenPermissionsRequest) (*settingspb.TokenPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.TokenPermissionsRequestPb{}

	var accessControlListPb []settingspb.TokenAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := TokenAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	return pb, nil
}

func TokenPermissionsRequestFromPb(pb *settingspb.TokenPermissionsRequestPb) (*TokenPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissionsRequest{}

	var accessControlListField []TokenAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := TokenAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField

	return st, nil
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

func TokenTypeToPb(st *TokenType) (*settingspb.TokenTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := settingspb.TokenTypePb(*st)
	return &pb, nil
}

func TokenTypeFromPb(pb *settingspb.TokenTypePb) (*TokenType, error) {
	if pb == nil {
		return nil, nil
	}
	st := TokenType(*pb)
	return &st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting AccountIpAccessEnable `json:"setting"`
}

func (st UpdateAccountIpAccessEnableRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateAccountIpAccessEnableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateAccountIpAccessEnableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateAccountIpAccessEnableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateAccountIpAccessEnableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateAccountIpAccessEnableRequestToPb(st *UpdateAccountIpAccessEnableRequest) (*settingspb.UpdateAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateAccountIpAccessEnableRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := AccountIpAccessEnableToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateAccountIpAccessEnableRequestFromPb(pb *settingspb.UpdateAccountIpAccessEnableRequestPb) (*UpdateAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAccountIpAccessEnableRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := AccountIpAccessEnableFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting AibiDashboardEmbeddingAccessPolicySetting `json:"setting"`
}

func (st UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*settingspb.UpdateAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := AibiDashboardEmbeddingAccessPolicySettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *settingspb.UpdateAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*UpdateAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := AibiDashboardEmbeddingAccessPolicySettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting AibiDashboardEmbeddingApprovedDomainsSetting `json:"setting"`
}

func (st UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*settingspb.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := AibiDashboardEmbeddingApprovedDomainsSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *settingspb.UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := AibiDashboardEmbeddingApprovedDomainsSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting AutomaticClusterUpdateSetting `json:"setting"`
}

func (st UpdateAutomaticClusterUpdateSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateAutomaticClusterUpdateSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateAutomaticClusterUpdateSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateAutomaticClusterUpdateSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateAutomaticClusterUpdateSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateAutomaticClusterUpdateSettingRequestToPb(st *UpdateAutomaticClusterUpdateSettingRequest) (*settingspb.UpdateAutomaticClusterUpdateSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateAutomaticClusterUpdateSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := AutomaticClusterUpdateSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateAutomaticClusterUpdateSettingRequestFromPb(pb *settingspb.UpdateAutomaticClusterUpdateSettingRequestPb) (*UpdateAutomaticClusterUpdateSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAutomaticClusterUpdateSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := AutomaticClusterUpdateSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting ComplianceSecurityProfileSetting `json:"setting"`
}

func (st UpdateComplianceSecurityProfileSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateComplianceSecurityProfileSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateComplianceSecurityProfileSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateComplianceSecurityProfileSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateComplianceSecurityProfileSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateComplianceSecurityProfileSettingRequestToPb(st *UpdateComplianceSecurityProfileSettingRequest) (*settingspb.UpdateComplianceSecurityProfileSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateComplianceSecurityProfileSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := ComplianceSecurityProfileSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateComplianceSecurityProfileSettingRequestFromPb(pb *settingspb.UpdateComplianceSecurityProfileSettingRequestPb) (*UpdateComplianceSecurityProfileSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateComplianceSecurityProfileSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := ComplianceSecurityProfileSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting CspEnablementAccountSetting `json:"setting"`
}

func (st UpdateCspEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCspEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCspEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateCspEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCspEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCspEnablementAccountSettingRequestToPb(st *UpdateCspEnablementAccountSettingRequest) (*settingspb.UpdateCspEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateCspEnablementAccountSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := CspEnablementAccountSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateCspEnablementAccountSettingRequestFromPb(pb *settingspb.UpdateCspEnablementAccountSettingRequestPb) (*UpdateCspEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCspEnablementAccountSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := CspEnablementAccountSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting DashboardEmailSubscriptions `json:"setting"`
}

func (st UpdateDashboardEmailSubscriptionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateDashboardEmailSubscriptionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateDashboardEmailSubscriptionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateDashboardEmailSubscriptionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateDashboardEmailSubscriptionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateDashboardEmailSubscriptionsRequestToPb(st *UpdateDashboardEmailSubscriptionsRequest) (*settingspb.UpdateDashboardEmailSubscriptionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateDashboardEmailSubscriptionsRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := DashboardEmailSubscriptionsToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateDashboardEmailSubscriptionsRequestFromPb(pb *settingspb.UpdateDashboardEmailSubscriptionsRequestPb) (*UpdateDashboardEmailSubscriptionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDashboardEmailSubscriptionsRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := DashboardEmailSubscriptionsFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting DefaultNamespaceSetting `json:"setting"`
}

func (st UpdateDefaultNamespaceSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateDefaultNamespaceSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateDefaultNamespaceSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateDefaultNamespaceSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateDefaultNamespaceSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateDefaultNamespaceSettingRequestToPb(st *UpdateDefaultNamespaceSettingRequest) (*settingspb.UpdateDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateDefaultNamespaceSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := DefaultNamespaceSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateDefaultNamespaceSettingRequestFromPb(pb *settingspb.UpdateDefaultNamespaceSettingRequestPb) (*UpdateDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDefaultNamespaceSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := DefaultNamespaceSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
}

// Details required to update a setting.
type UpdateDefaultWarehouseIdRequest struct {
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting DefaultWarehouseId `json:"setting"`
}

func (st UpdateDefaultWarehouseIdRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateDefaultWarehouseIdRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateDefaultWarehouseIdRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateDefaultWarehouseIdRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateDefaultWarehouseIdRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateDefaultWarehouseIdRequestToPb(st *UpdateDefaultWarehouseIdRequest) (*settingspb.UpdateDefaultWarehouseIdRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateDefaultWarehouseIdRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := DefaultWarehouseIdToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateDefaultWarehouseIdRequestFromPb(pb *settingspb.UpdateDefaultWarehouseIdRequestPb) (*UpdateDefaultWarehouseIdRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDefaultWarehouseIdRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := DefaultWarehouseIdFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting DisableLegacyAccess `json:"setting"`
}

func (st UpdateDisableLegacyAccessRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateDisableLegacyAccessRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateDisableLegacyAccessRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateDisableLegacyAccessRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateDisableLegacyAccessRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateDisableLegacyAccessRequestToPb(st *UpdateDisableLegacyAccessRequest) (*settingspb.UpdateDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateDisableLegacyAccessRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := DisableLegacyAccessToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateDisableLegacyAccessRequestFromPb(pb *settingspb.UpdateDisableLegacyAccessRequestPb) (*UpdateDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyAccessRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := DisableLegacyAccessFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting DisableLegacyDbfs `json:"setting"`
}

func (st UpdateDisableLegacyDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateDisableLegacyDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateDisableLegacyDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateDisableLegacyDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateDisableLegacyDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateDisableLegacyDbfsRequestToPb(st *UpdateDisableLegacyDbfsRequest) (*settingspb.UpdateDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateDisableLegacyDbfsRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := DisableLegacyDbfsToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateDisableLegacyDbfsRequestFromPb(pb *settingspb.UpdateDisableLegacyDbfsRequestPb) (*UpdateDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyDbfsRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := DisableLegacyDbfsFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting DisableLegacyFeatures `json:"setting"`
}

func (st UpdateDisableLegacyFeaturesRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateDisableLegacyFeaturesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateDisableLegacyFeaturesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateDisableLegacyFeaturesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateDisableLegacyFeaturesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateDisableLegacyFeaturesRequestToPb(st *UpdateDisableLegacyFeaturesRequest) (*settingspb.UpdateDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateDisableLegacyFeaturesRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := DisableLegacyFeaturesToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateDisableLegacyFeaturesRequestFromPb(pb *settingspb.UpdateDisableLegacyFeaturesRequestPb) (*UpdateDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyFeaturesRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := DisableLegacyFeaturesFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting EnableExportNotebook `json:"setting"`
}

func (st UpdateEnableExportNotebookRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateEnableExportNotebookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateEnableExportNotebookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateEnableExportNotebookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateEnableExportNotebookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateEnableExportNotebookRequestToPb(st *UpdateEnableExportNotebookRequest) (*settingspb.UpdateEnableExportNotebookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateEnableExportNotebookRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := EnableExportNotebookToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateEnableExportNotebookRequestFromPb(pb *settingspb.UpdateEnableExportNotebookRequestPb) (*UpdateEnableExportNotebookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableExportNotebookRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := EnableExportNotebookFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting EnableNotebookTableClipboard `json:"setting"`
}

func (st UpdateEnableNotebookTableClipboardRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateEnableNotebookTableClipboardRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateEnableNotebookTableClipboardRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateEnableNotebookTableClipboardRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateEnableNotebookTableClipboardRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateEnableNotebookTableClipboardRequestToPb(st *UpdateEnableNotebookTableClipboardRequest) (*settingspb.UpdateEnableNotebookTableClipboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateEnableNotebookTableClipboardRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := EnableNotebookTableClipboardToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateEnableNotebookTableClipboardRequestFromPb(pb *settingspb.UpdateEnableNotebookTableClipboardRequestPb) (*UpdateEnableNotebookTableClipboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableNotebookTableClipboardRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := EnableNotebookTableClipboardFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting EnableResultsDownloading `json:"setting"`
}

func (st UpdateEnableResultsDownloadingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateEnableResultsDownloadingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateEnableResultsDownloadingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateEnableResultsDownloadingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateEnableResultsDownloadingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateEnableResultsDownloadingRequestToPb(st *UpdateEnableResultsDownloadingRequest) (*settingspb.UpdateEnableResultsDownloadingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateEnableResultsDownloadingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := EnableResultsDownloadingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateEnableResultsDownloadingRequestFromPb(pb *settingspb.UpdateEnableResultsDownloadingRequestPb) (*UpdateEnableResultsDownloadingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableResultsDownloadingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := EnableResultsDownloadingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting EnhancedSecurityMonitoringSetting `json:"setting"`
}

func (st UpdateEnhancedSecurityMonitoringSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateEnhancedSecurityMonitoringSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateEnhancedSecurityMonitoringSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateEnhancedSecurityMonitoringSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateEnhancedSecurityMonitoringSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateEnhancedSecurityMonitoringSettingRequestToPb(st *UpdateEnhancedSecurityMonitoringSettingRequest) (*settingspb.UpdateEnhancedSecurityMonitoringSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateEnhancedSecurityMonitoringSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := EnhancedSecurityMonitoringSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateEnhancedSecurityMonitoringSettingRequestFromPb(pb *settingspb.UpdateEnhancedSecurityMonitoringSettingRequestPb) (*UpdateEnhancedSecurityMonitoringSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnhancedSecurityMonitoringSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := EnhancedSecurityMonitoringSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting EsmEnablementAccountSetting `json:"setting"`
}

func (st UpdateEsmEnablementAccountSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateEsmEnablementAccountSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateEsmEnablementAccountSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateEsmEnablementAccountSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateEsmEnablementAccountSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateEsmEnablementAccountSettingRequestToPb(st *UpdateEsmEnablementAccountSettingRequest) (*settingspb.UpdateEsmEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateEsmEnablementAccountSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := EsmEnablementAccountSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateEsmEnablementAccountSettingRequestFromPb(pb *settingspb.UpdateEsmEnablementAccountSettingRequestPb) (*UpdateEsmEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEsmEnablementAccountSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := EsmEnablementAccountSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	ListType        ListType `json:"list_type,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateIpAccessList) MarshalJSON() ([]byte, error) {
	pb, err := UpdateIpAccessListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateIpAccessList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateIpAccessListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateIpAccessListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateIpAccessListToPb(st *UpdateIpAccessList) (*settingspb.UpdateIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateIpAccessListPb{}
	pb.Enabled = st.Enabled
	pb.IpAccessListId = st.IpAccessListId
	pb.IpAddresses = st.IpAddresses
	pb.Label = st.Label
	listTypePb, err := ListTypeToPb(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateIpAccessListFromPb(pb *settingspb.UpdateIpAccessListPb) (*UpdateIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateIpAccessList{}
	st.Enabled = pb.Enabled
	st.IpAccessListId = pb.IpAccessListId
	st.IpAddresses = pb.IpAddresses
	st.Label = pb.Label
	listTypeField, err := ListTypeFromPb(&pb.ListType)
	if err != nil {
		return nil, err
	}
	if listTypeField != nil {
		st.ListType = *listTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredAccount `json:"setting"`
}

func (st UpdateLlmProxyPartnerPoweredAccountRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateLlmProxyPartnerPoweredAccountRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateLlmProxyPartnerPoweredAccountRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateLlmProxyPartnerPoweredAccountRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateLlmProxyPartnerPoweredAccountRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateLlmProxyPartnerPoweredAccountRequestToPb(st *UpdateLlmProxyPartnerPoweredAccountRequest) (*settingspb.UpdateLlmProxyPartnerPoweredAccountRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateLlmProxyPartnerPoweredAccountRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := LlmProxyPartnerPoweredAccountToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateLlmProxyPartnerPoweredAccountRequestFromPb(pb *settingspb.UpdateLlmProxyPartnerPoweredAccountRequestPb) (*UpdateLlmProxyPartnerPoweredAccountRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLlmProxyPartnerPoweredAccountRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := LlmProxyPartnerPoweredAccountFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredEnforce `json:"setting"`
}

func (st UpdateLlmProxyPartnerPoweredEnforceRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateLlmProxyPartnerPoweredEnforceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateLlmProxyPartnerPoweredEnforceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateLlmProxyPartnerPoweredEnforceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateLlmProxyPartnerPoweredEnforceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateLlmProxyPartnerPoweredEnforceRequestToPb(st *UpdateLlmProxyPartnerPoweredEnforceRequest) (*settingspb.UpdateLlmProxyPartnerPoweredEnforceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateLlmProxyPartnerPoweredEnforceRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := LlmProxyPartnerPoweredEnforceToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateLlmProxyPartnerPoweredEnforceRequestFromPb(pb *settingspb.UpdateLlmProxyPartnerPoweredEnforceRequestPb) (*UpdateLlmProxyPartnerPoweredEnforceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLlmProxyPartnerPoweredEnforceRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := LlmProxyPartnerPoweredEnforceFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting LlmProxyPartnerPoweredWorkspace `json:"setting"`
}

func (st UpdateLlmProxyPartnerPoweredWorkspaceRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateLlmProxyPartnerPoweredWorkspaceRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateLlmProxyPartnerPoweredWorkspaceRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateLlmProxyPartnerPoweredWorkspaceRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateLlmProxyPartnerPoweredWorkspaceRequestToPb(st *UpdateLlmProxyPartnerPoweredWorkspaceRequest) (*settingspb.UpdateLlmProxyPartnerPoweredWorkspaceRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateLlmProxyPartnerPoweredWorkspaceRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := LlmProxyPartnerPoweredWorkspaceToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateLlmProxyPartnerPoweredWorkspaceRequestFromPb(pb *settingspb.UpdateLlmProxyPartnerPoweredWorkspaceRequestPb) (*UpdateLlmProxyPartnerPoweredWorkspaceRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateLlmProxyPartnerPoweredWorkspaceRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := LlmProxyPartnerPoweredWorkspaceFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	UpdateMask string `json:"-" tf:"-"` //legacy

}

func (st UpdateNccPrivateEndpointRuleRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateNccPrivateEndpointRuleRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateNccPrivateEndpointRuleRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateNccPrivateEndpointRuleRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateNccPrivateEndpointRuleRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateNccPrivateEndpointRuleRequestToPb(st *UpdateNccPrivateEndpointRuleRequest) (*settingspb.UpdateNccPrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateNccPrivateEndpointRuleRequestPb{}
	pb.NetworkConnectivityConfigId = st.NetworkConnectivityConfigId
	privateEndpointRulePb, err := UpdatePrivateEndpointRuleToPb(&st.PrivateEndpointRule)
	if err != nil {
		return nil, err
	}
	if privateEndpointRulePb != nil {
		pb.PrivateEndpointRule = *privateEndpointRulePb
	}
	pb.PrivateEndpointRuleId = st.PrivateEndpointRuleId
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

func UpdateNccPrivateEndpointRuleRequestFromPb(pb *settingspb.UpdateNccPrivateEndpointRuleRequestPb) (*UpdateNccPrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNccPrivateEndpointRuleRequest{}
	st.NetworkConnectivityConfigId = pb.NetworkConnectivityConfigId
	privateEndpointRuleField, err := UpdatePrivateEndpointRuleFromPb(&pb.PrivateEndpointRule)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleField != nil {
		st.PrivateEndpointRule = *privateEndpointRuleField
	}
	st.PrivateEndpointRuleId = pb.PrivateEndpointRuleId
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

type UpdateNetworkPolicyRequest struct {
	// Updated network policy configuration details.
	// Wire name: 'network_policy'
	NetworkPolicy AccountNetworkPolicy `json:"network_policy"`
	// The unique identifier for the network policy.
	NetworkPolicyId string `json:"-" tf:"-"`
}

func (st UpdateNetworkPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateNetworkPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateNetworkPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateNetworkPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateNetworkPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateNetworkPolicyRequestToPb(st *UpdateNetworkPolicyRequest) (*settingspb.UpdateNetworkPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateNetworkPolicyRequestPb{}
	networkPolicyPb, err := AccountNetworkPolicyToPb(&st.NetworkPolicy)
	if err != nil {
		return nil, err
	}
	if networkPolicyPb != nil {
		pb.NetworkPolicy = *networkPolicyPb
	}
	pb.NetworkPolicyId = st.NetworkPolicyId

	return pb, nil
}

func UpdateNetworkPolicyRequestFromPb(pb *settingspb.UpdateNetworkPolicyRequestPb) (*UpdateNetworkPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNetworkPolicyRequest{}
	networkPolicyField, err := AccountNetworkPolicyFromPb(&pb.NetworkPolicy)
	if err != nil {
		return nil, err
	}
	if networkPolicyField != nil {
		st.NetworkPolicy = *networkPolicyField
	}
	st.NetworkPolicyId = pb.NetworkPolicyId

	return st, nil
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
	Id              string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateNotificationDestinationRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateNotificationDestinationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateNotificationDestinationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateNotificationDestinationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateNotificationDestinationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateNotificationDestinationRequestToPb(st *UpdateNotificationDestinationRequest) (*settingspb.UpdateNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateNotificationDestinationRequestPb{}
	configPb, err := ConfigToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}
	pb.DisplayName = st.DisplayName
	pb.Id = st.Id

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateNotificationDestinationRequestFromPb(pb *settingspb.UpdateNotificationDestinationRequestPb) (*UpdateNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNotificationDestinationRequest{}
	configField, err := ConfigFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	st.DisplayName = pb.DisplayName
	st.Id = pb.Id

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting PersonalComputeSetting `json:"setting"`
}

func (st UpdatePersonalComputeSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdatePersonalComputeSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdatePersonalComputeSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdatePersonalComputeSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdatePersonalComputeSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdatePersonalComputeSettingRequestToPb(st *UpdatePersonalComputeSettingRequest) (*settingspb.UpdatePersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdatePersonalComputeSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := PersonalComputeSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdatePersonalComputeSettingRequestFromPb(pb *settingspb.UpdatePersonalComputeSettingRequestPb) (*UpdatePersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalComputeSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := PersonalComputeSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	ResourceNames   []string `json:"resource_names,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdatePrivateEndpointRule) MarshalJSON() ([]byte, error) {
	pb, err := UpdatePrivateEndpointRuleToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdatePrivateEndpointRule) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdatePrivateEndpointRulePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdatePrivateEndpointRuleFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdatePrivateEndpointRuleToPb(st *UpdatePrivateEndpointRule) (*settingspb.UpdatePrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdatePrivateEndpointRulePb{}
	pb.DomainNames = st.DomainNames
	pb.Enabled = st.Enabled
	pb.ResourceNames = st.ResourceNames

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdatePrivateEndpointRuleFromPb(pb *settingspb.UpdatePrivateEndpointRulePb) (*UpdatePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePrivateEndpointRule{}
	st.DomainNames = pb.DomainNames
	st.Enabled = pb.Enabled
	st.ResourceNames = pb.ResourceNames

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting RestrictWorkspaceAdminsSetting `json:"setting"`
}

func (st UpdateRestrictWorkspaceAdminsSettingRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateRestrictWorkspaceAdminsSettingRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateRestrictWorkspaceAdminsSettingRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateRestrictWorkspaceAdminsSettingRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateRestrictWorkspaceAdminsSettingRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateRestrictWorkspaceAdminsSettingRequestToPb(st *UpdateRestrictWorkspaceAdminsSettingRequest) (*settingspb.UpdateRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateRestrictWorkspaceAdminsSettingRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := RestrictWorkspaceAdminsSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateRestrictWorkspaceAdminsSettingRequestFromPb(pb *settingspb.UpdateRestrictWorkspaceAdminsSettingRequestPb) (*UpdateRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRestrictWorkspaceAdminsSettingRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := RestrictWorkspaceAdminsSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
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
	FieldMask string `json:"field_mask"` //legacy

	// Wire name: 'setting'
	Setting SqlResultsDownload `json:"setting"`
}

func (st UpdateSqlResultsDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateSqlResultsDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateSqlResultsDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateSqlResultsDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateSqlResultsDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateSqlResultsDownloadRequestToPb(st *UpdateSqlResultsDownloadRequest) (*settingspb.UpdateSqlResultsDownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateSqlResultsDownloadRequestPb{}
	pb.AllowMissing = st.AllowMissing
	pb.FieldMask = st.FieldMask
	settingPb, err := SqlResultsDownloadToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
}

func UpdateSqlResultsDownloadRequestFromPb(pb *settingspb.UpdateSqlResultsDownloadRequestPb) (*UpdateSqlResultsDownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateSqlResultsDownloadRequest{}
	st.AllowMissing = pb.AllowMissing
	st.FieldMask = pb.FieldMask
	settingField, err := SqlResultsDownloadFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
}

type UpdateWorkspaceNetworkOptionRequest struct {
	// The workspace ID.
	WorkspaceId int64 `json:"-" tf:"-"`
	// The network option details for the workspace.
	// Wire name: 'workspace_network_option'
	WorkspaceNetworkOption WorkspaceNetworkOption `json:"workspace_network_option"`
}

func (st UpdateWorkspaceNetworkOptionRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateWorkspaceNetworkOptionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateWorkspaceNetworkOptionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.UpdateWorkspaceNetworkOptionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateWorkspaceNetworkOptionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateWorkspaceNetworkOptionRequestToPb(st *UpdateWorkspaceNetworkOptionRequest) (*settingspb.UpdateWorkspaceNetworkOptionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.UpdateWorkspaceNetworkOptionRequestPb{}
	pb.WorkspaceId = st.WorkspaceId
	workspaceNetworkOptionPb, err := WorkspaceNetworkOptionToPb(&st.WorkspaceNetworkOption)
	if err != nil {
		return nil, err
	}
	if workspaceNetworkOptionPb != nil {
		pb.WorkspaceNetworkOption = *workspaceNetworkOptionPb
	}

	return pb, nil
}

func UpdateWorkspaceNetworkOptionRequestFromPb(pb *settingspb.UpdateWorkspaceNetworkOptionRequestPb) (*UpdateWorkspaceNetworkOptionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWorkspaceNetworkOptionRequest{}
	st.WorkspaceId = pb.WorkspaceId
	workspaceNetworkOptionField, err := WorkspaceNetworkOptionFromPb(&pb.WorkspaceNetworkOption)
	if err != nil {
		return nil, err
	}
	if workspaceNetworkOptionField != nil {
		st.WorkspaceNetworkOption = *workspaceNetworkOptionField
	}

	return st, nil
}

type WorkspaceConf map[string]string

func WorkspaceConfToPb(st *WorkspaceConf) (*settingspb.WorkspaceConfPb, error) {
	if st == nil {
		return nil, nil
	}
	stPb := settingspb.WorkspaceConfPb(*st)
	return &stPb, nil
}

func WorkspaceConfFromPb(stPb *settingspb.WorkspaceConfPb) (*WorkspaceConf, error) {
	if stPb == nil {
		return nil, nil
	}
	st := WorkspaceConf(*stPb)
	return &st, nil
}

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
	WorkspaceId     int64    `json:"workspace_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st WorkspaceNetworkOption) MarshalJSON() ([]byte, error) {
	pb, err := WorkspaceNetworkOptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *WorkspaceNetworkOption) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &settingspb.WorkspaceNetworkOptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := WorkspaceNetworkOptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func WorkspaceNetworkOptionToPb(st *WorkspaceNetworkOption) (*settingspb.WorkspaceNetworkOptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &settingspb.WorkspaceNetworkOptionPb{}
	pb.NetworkPolicyId = st.NetworkPolicyId
	pb.WorkspaceId = st.WorkspaceId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func WorkspaceNetworkOptionFromPb(pb *settingspb.WorkspaceNetworkOptionPb) (*WorkspaceNetworkOption, error) {
	if pb == nil {
		return nil, nil
	}
	st := &WorkspaceNetworkOption{}
	st.NetworkPolicyId = pb.NetworkPolicyId
	st.WorkspaceId = pb.WorkspaceId

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
