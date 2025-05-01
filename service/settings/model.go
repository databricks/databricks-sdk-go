// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package settings

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type AccountIpAccessEnable struct {
	AcctIpAclEnable BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func accountIpAccessEnableToPb(st *AccountIpAccessEnable) (*accountIpAccessEnablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &accountIpAccessEnablePb{}
	acctIpAclEnablePb, err := booleanMessageToPb(&st.AcctIpAclEnable)
	if err != nil {
		return nil, err
	}
	if acctIpAclEnablePb != nil {
		pb.AcctIpAclEnable = *acctIpAclEnablePb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type accountIpAccessEnablePb struct {
	AcctIpAclEnable booleanMessagePb `json:"acct_ip_acl_enable"`
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

func accountIpAccessEnableFromPb(pb *accountIpAccessEnablePb) (*AccountIpAccessEnable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AccountIpAccessEnable{}
	acctIpAclEnableField, err := booleanMessageFromPb(&pb.AcctIpAclEnable)
	if err != nil {
		return nil, err
	}
	if acctIpAclEnableField != nil {
		st.AcctIpAclEnable = *acctIpAclEnableField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *accountIpAccessEnablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st accountIpAccessEnablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AibiDashboardEmbeddingAccessPolicy struct {
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType
}

func aibiDashboardEmbeddingAccessPolicyToPb(st *AibiDashboardEmbeddingAccessPolicy) (*aibiDashboardEmbeddingAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingAccessPolicyPb{}
	accessPolicyTypePb, err := identity(&st.AccessPolicyType)
	if err != nil {
		return nil, err
	}
	if accessPolicyTypePb != nil {
		pb.AccessPolicyType = *accessPolicyTypePb
	}

	return pb, nil
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

type aibiDashboardEmbeddingAccessPolicyPb struct {
	AccessPolicyType AibiDashboardEmbeddingAccessPolicyAccessPolicyType `json:"access_policy_type"`
}

func aibiDashboardEmbeddingAccessPolicyFromPb(pb *aibiDashboardEmbeddingAccessPolicyPb) (*AibiDashboardEmbeddingAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingAccessPolicy{}
	accessPolicyTypeField, err := identity(&pb.AccessPolicyType)
	if err != nil {
		return nil, err
	}
	if accessPolicyTypeField != nil {
		st.AccessPolicyType = *accessPolicyTypeField
	}

	return st, nil
}

type AibiDashboardEmbeddingAccessPolicyAccessPolicyType string
type aibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb string

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

func aibiDashboardEmbeddingAccessPolicyAccessPolicyTypeToPb(st *AibiDashboardEmbeddingAccessPolicyAccessPolicyType) (*aibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := aibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb(*st)
	return &pb, nil
}

func aibiDashboardEmbeddingAccessPolicyAccessPolicyTypeFromPb(pb *aibiDashboardEmbeddingAccessPolicyAccessPolicyTypePb) (*AibiDashboardEmbeddingAccessPolicyAccessPolicyType, error) {
	if pb == nil {
		return nil, nil
	}
	st := AibiDashboardEmbeddingAccessPolicyAccessPolicyType(*pb)
	return &st, nil
}

type AibiDashboardEmbeddingAccessPolicySetting struct {
	AibiDashboardEmbeddingAccessPolicy AibiDashboardEmbeddingAccessPolicy
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func aibiDashboardEmbeddingAccessPolicySettingToPb(st *AibiDashboardEmbeddingAccessPolicySetting) (*aibiDashboardEmbeddingAccessPolicySettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingAccessPolicySettingPb{}
	aibiDashboardEmbeddingAccessPolicyPb, err := aibiDashboardEmbeddingAccessPolicyToPb(&st.AibiDashboardEmbeddingAccessPolicy)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingAccessPolicyPb != nil {
		pb.AibiDashboardEmbeddingAccessPolicy = *aibiDashboardEmbeddingAccessPolicyPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type aibiDashboardEmbeddingAccessPolicySettingPb struct {
	AibiDashboardEmbeddingAccessPolicy aibiDashboardEmbeddingAccessPolicyPb `json:"aibi_dashboard_embedding_access_policy"`
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

func aibiDashboardEmbeddingAccessPolicySettingFromPb(pb *aibiDashboardEmbeddingAccessPolicySettingPb) (*AibiDashboardEmbeddingAccessPolicySetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingAccessPolicySetting{}
	aibiDashboardEmbeddingAccessPolicyField, err := aibiDashboardEmbeddingAccessPolicyFromPb(&pb.AibiDashboardEmbeddingAccessPolicy)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingAccessPolicyField != nil {
		st.AibiDashboardEmbeddingAccessPolicy = *aibiDashboardEmbeddingAccessPolicyField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aibiDashboardEmbeddingAccessPolicySettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aibiDashboardEmbeddingAccessPolicySettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AibiDashboardEmbeddingApprovedDomains struct {
	ApprovedDomains []string
}

func aibiDashboardEmbeddingApprovedDomainsToPb(st *AibiDashboardEmbeddingApprovedDomains) (*aibiDashboardEmbeddingApprovedDomainsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingApprovedDomainsPb{}

	var approvedDomainsPb []string
	for _, item := range st.ApprovedDomains {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			approvedDomainsPb = append(approvedDomainsPb, *itemPb)
		}
	}
	pb.ApprovedDomains = approvedDomainsPb

	return pb, nil
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

type aibiDashboardEmbeddingApprovedDomainsPb struct {
	ApprovedDomains []string `json:"approved_domains,omitempty"`
}

func aibiDashboardEmbeddingApprovedDomainsFromPb(pb *aibiDashboardEmbeddingApprovedDomainsPb) (*AibiDashboardEmbeddingApprovedDomains, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingApprovedDomains{}

	var approvedDomainsField []string
	for _, itemPb := range pb.ApprovedDomains {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			approvedDomainsField = append(approvedDomainsField, *item)
		}
	}
	st.ApprovedDomains = approvedDomainsField

	return st, nil
}

type AibiDashboardEmbeddingApprovedDomainsSetting struct {
	AibiDashboardEmbeddingApprovedDomains AibiDashboardEmbeddingApprovedDomains
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func aibiDashboardEmbeddingApprovedDomainsSettingToPb(st *AibiDashboardEmbeddingApprovedDomainsSetting) (*aibiDashboardEmbeddingApprovedDomainsSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aibiDashboardEmbeddingApprovedDomainsSettingPb{}
	aibiDashboardEmbeddingApprovedDomainsPb, err := aibiDashboardEmbeddingApprovedDomainsToPb(&st.AibiDashboardEmbeddingApprovedDomains)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingApprovedDomainsPb != nil {
		pb.AibiDashboardEmbeddingApprovedDomains = *aibiDashboardEmbeddingApprovedDomainsPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type aibiDashboardEmbeddingApprovedDomainsSettingPb struct {
	AibiDashboardEmbeddingApprovedDomains aibiDashboardEmbeddingApprovedDomainsPb `json:"aibi_dashboard_embedding_approved_domains"`
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

func aibiDashboardEmbeddingApprovedDomainsSettingFromPb(pb *aibiDashboardEmbeddingApprovedDomainsSettingPb) (*AibiDashboardEmbeddingApprovedDomainsSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AibiDashboardEmbeddingApprovedDomainsSetting{}
	aibiDashboardEmbeddingApprovedDomainsField, err := aibiDashboardEmbeddingApprovedDomainsFromPb(&pb.AibiDashboardEmbeddingApprovedDomains)
	if err != nil {
		return nil, err
	}
	if aibiDashboardEmbeddingApprovedDomainsField != nil {
		st.AibiDashboardEmbeddingApprovedDomains = *aibiDashboardEmbeddingApprovedDomainsField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aibiDashboardEmbeddingApprovedDomainsSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aibiDashboardEmbeddingApprovedDomainsSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutomaticClusterUpdateSetting struct {
	AutomaticClusterUpdateWorkspace ClusterAutoRestartMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func automaticClusterUpdateSettingToPb(st *AutomaticClusterUpdateSetting) (*automaticClusterUpdateSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &automaticClusterUpdateSettingPb{}
	automaticClusterUpdateWorkspacePb, err := clusterAutoRestartMessageToPb(&st.AutomaticClusterUpdateWorkspace)
	if err != nil {
		return nil, err
	}
	if automaticClusterUpdateWorkspacePb != nil {
		pb.AutomaticClusterUpdateWorkspace = *automaticClusterUpdateWorkspacePb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type automaticClusterUpdateSettingPb struct {
	AutomaticClusterUpdateWorkspace clusterAutoRestartMessagePb `json:"automatic_cluster_update_workspace"`
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

func automaticClusterUpdateSettingFromPb(pb *automaticClusterUpdateSettingPb) (*AutomaticClusterUpdateSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutomaticClusterUpdateSetting{}
	automaticClusterUpdateWorkspaceField, err := clusterAutoRestartMessageFromPb(&pb.AutomaticClusterUpdateWorkspace)
	if err != nil {
		return nil, err
	}
	if automaticClusterUpdateWorkspaceField != nil {
		st.AutomaticClusterUpdateWorkspace = *automaticClusterUpdateWorkspaceField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *automaticClusterUpdateSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st automaticClusterUpdateSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BooleanMessage struct {
	Value bool

	ForceSendFields []string
}

func booleanMessageToPb(st *BooleanMessage) (*booleanMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &booleanMessagePb{}
	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type booleanMessagePb struct {
	Value bool `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func booleanMessageFromPb(pb *booleanMessagePb) (*BooleanMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BooleanMessage{}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *booleanMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st booleanMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAutoRestartMessage struct {
	CanToggle bool

	Enabled bool
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails *ClusterAutoRestartMessageEnablementDetails

	MaintenanceWindow *ClusterAutoRestartMessageMaintenanceWindow

	RestartEvenIfNoUpdatesAvailable bool

	ForceSendFields []string
}

func clusterAutoRestartMessageToPb(st *ClusterAutoRestartMessage) (*clusterAutoRestartMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessagePb{}
	canTogglePb, err := identity(&st.CanToggle)
	if err != nil {
		return nil, err
	}
	if canTogglePb != nil {
		pb.CanToggle = *canTogglePb
	}

	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	enablementDetailsPb, err := clusterAutoRestartMessageEnablementDetailsToPb(st.EnablementDetails)
	if err != nil {
		return nil, err
	}
	if enablementDetailsPb != nil {
		pb.EnablementDetails = enablementDetailsPb
	}

	maintenanceWindowPb, err := clusterAutoRestartMessageMaintenanceWindowToPb(st.MaintenanceWindow)
	if err != nil {
		return nil, err
	}
	if maintenanceWindowPb != nil {
		pb.MaintenanceWindow = maintenanceWindowPb
	}

	restartEvenIfNoUpdatesAvailablePb, err := identity(&st.RestartEvenIfNoUpdatesAvailable)
	if err != nil {
		return nil, err
	}
	if restartEvenIfNoUpdatesAvailablePb != nil {
		pb.RestartEvenIfNoUpdatesAvailable = *restartEvenIfNoUpdatesAvailablePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterAutoRestartMessagePb struct {
	CanToggle bool `json:"can_toggle,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
	// Contains an information about the enablement status judging (e.g. whether
	// the enterprise tier is enabled) This is only additional information that
	// MUST NOT be used to decide whether the setting is enabled or not. This is
	// intended to use only for purposes like showing an error message to the
	// customer with the additional details. For example, using these details we
	// can check why exactly the feature is disabled for this customer.
	EnablementDetails *clusterAutoRestartMessageEnablementDetailsPb `json:"enablement_details,omitempty"`

	MaintenanceWindow *clusterAutoRestartMessageMaintenanceWindowPb `json:"maintenance_window,omitempty"`

	RestartEvenIfNoUpdatesAvailable bool `json:"restart_even_if_no_updates_available,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAutoRestartMessageFromPb(pb *clusterAutoRestartMessagePb) (*ClusterAutoRestartMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessage{}
	canToggleField, err := identity(&pb.CanToggle)
	if err != nil {
		return nil, err
	}
	if canToggleField != nil {
		st.CanToggle = *canToggleField
	}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	enablementDetailsField, err := clusterAutoRestartMessageEnablementDetailsFromPb(pb.EnablementDetails)
	if err != nil {
		return nil, err
	}
	if enablementDetailsField != nil {
		st.EnablementDetails = enablementDetailsField
	}
	maintenanceWindowField, err := clusterAutoRestartMessageMaintenanceWindowFromPb(pb.MaintenanceWindow)
	if err != nil {
		return nil, err
	}
	if maintenanceWindowField != nil {
		st.MaintenanceWindow = maintenanceWindowField
	}
	restartEvenIfNoUpdatesAvailableField, err := identity(&pb.RestartEvenIfNoUpdatesAvailable)
	if err != nil {
		return nil, err
	}
	if restartEvenIfNoUpdatesAvailableField != nil {
		st.RestartEvenIfNoUpdatesAvailable = *restartEvenIfNoUpdatesAvailableField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAutoRestartMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAutoRestartMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Contains an information about the enablement status judging (e.g. whether the
// enterprise tier is enabled) This is only additional information that MUST NOT
// be used to decide whether the setting is enabled or not. This is intended to
// use only for purposes like showing an error message to the customer with the
// additional details. For example, using these details we can check why exactly
// the feature is disabled for this customer.
type ClusterAutoRestartMessageEnablementDetails struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode bool
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement bool
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier bool

	ForceSendFields []string
}

func clusterAutoRestartMessageEnablementDetailsToPb(st *ClusterAutoRestartMessageEnablementDetails) (*clusterAutoRestartMessageEnablementDetailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageEnablementDetailsPb{}
	forcedForComplianceModePb, err := identity(&st.ForcedForComplianceMode)
	if err != nil {
		return nil, err
	}
	if forcedForComplianceModePb != nil {
		pb.ForcedForComplianceMode = *forcedForComplianceModePb
	}

	unavailableForDisabledEntitlementPb, err := identity(&st.UnavailableForDisabledEntitlement)
	if err != nil {
		return nil, err
	}
	if unavailableForDisabledEntitlementPb != nil {
		pb.UnavailableForDisabledEntitlement = *unavailableForDisabledEntitlementPb
	}

	unavailableForNonEnterpriseTierPb, err := identity(&st.UnavailableForNonEnterpriseTier)
	if err != nil {
		return nil, err
	}
	if unavailableForNonEnterpriseTierPb != nil {
		pb.UnavailableForNonEnterpriseTier = *unavailableForNonEnterpriseTierPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterAutoRestartMessageEnablementDetailsPb struct {
	// The feature is force enabled if compliance mode is active
	ForcedForComplianceMode bool `json:"forced_for_compliance_mode,omitempty"`
	// The feature is unavailable if the corresponding entitlement disabled (see
	// getShieldEntitlementEnable)
	UnavailableForDisabledEntitlement bool `json:"unavailable_for_disabled_entitlement,omitempty"`
	// The feature is unavailable if the customer doesn't have enterprise tier
	UnavailableForNonEnterpriseTier bool `json:"unavailable_for_non_enterprise_tier,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAutoRestartMessageEnablementDetailsFromPb(pb *clusterAutoRestartMessageEnablementDetailsPb) (*ClusterAutoRestartMessageEnablementDetails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageEnablementDetails{}
	forcedForComplianceModeField, err := identity(&pb.ForcedForComplianceMode)
	if err != nil {
		return nil, err
	}
	if forcedForComplianceModeField != nil {
		st.ForcedForComplianceMode = *forcedForComplianceModeField
	}
	unavailableForDisabledEntitlementField, err := identity(&pb.UnavailableForDisabledEntitlement)
	if err != nil {
		return nil, err
	}
	if unavailableForDisabledEntitlementField != nil {
		st.UnavailableForDisabledEntitlement = *unavailableForDisabledEntitlementField
	}
	unavailableForNonEnterpriseTierField, err := identity(&pb.UnavailableForNonEnterpriseTier)
	if err != nil {
		return nil, err
	}
	if unavailableForNonEnterpriseTierField != nil {
		st.UnavailableForNonEnterpriseTier = *unavailableForNonEnterpriseTierField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAutoRestartMessageEnablementDetailsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAutoRestartMessageEnablementDetailsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ClusterAutoRestartMessageMaintenanceWindow struct {
	WeekDayBasedSchedule *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule
}

func clusterAutoRestartMessageMaintenanceWindowToPb(st *ClusterAutoRestartMessageMaintenanceWindow) (*clusterAutoRestartMessageMaintenanceWindowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowPb{}
	weekDayBasedSchedulePb, err := clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(st.WeekDayBasedSchedule)
	if err != nil {
		return nil, err
	}
	if weekDayBasedSchedulePb != nil {
		pb.WeekDayBasedSchedule = weekDayBasedSchedulePb
	}

	return pb, nil
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

type clusterAutoRestartMessageMaintenanceWindowPb struct {
	WeekDayBasedSchedule *clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb `json:"week_day_based_schedule,omitempty"`
}

func clusterAutoRestartMessageMaintenanceWindowFromPb(pb *clusterAutoRestartMessageMaintenanceWindowPb) (*ClusterAutoRestartMessageMaintenanceWindow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindow{}
	weekDayBasedScheduleField, err := clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleFromPb(pb.WeekDayBasedSchedule)
	if err != nil {
		return nil, err
	}
	if weekDayBasedScheduleField != nil {
		st.WeekDayBasedSchedule = weekDayBasedScheduleField
	}

	return st, nil
}

type ClusterAutoRestartMessageMaintenanceWindowDayOfWeek string
type clusterAutoRestartMessageMaintenanceWindowDayOfWeekPb string

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

func clusterAutoRestartMessageMaintenanceWindowDayOfWeekToPb(st *ClusterAutoRestartMessageMaintenanceWindowDayOfWeek) (*clusterAutoRestartMessageMaintenanceWindowDayOfWeekPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := clusterAutoRestartMessageMaintenanceWindowDayOfWeekPb(*st)
	return &pb, nil
}

func clusterAutoRestartMessageMaintenanceWindowDayOfWeekFromPb(pb *clusterAutoRestartMessageMaintenanceWindowDayOfWeekPb) (*ClusterAutoRestartMessageMaintenanceWindowDayOfWeek, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterAutoRestartMessageMaintenanceWindowDayOfWeek(*pb)
	return &st, nil
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule struct {
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek

	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency

	WindowStartTime *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime
}

func clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleToPb(st *ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule) (*clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb{}
	dayOfWeekPb, err := identity(&st.DayOfWeek)
	if err != nil {
		return nil, err
	}
	if dayOfWeekPb != nil {
		pb.DayOfWeek = *dayOfWeekPb
	}

	frequencyPb, err := identity(&st.Frequency)
	if err != nil {
		return nil, err
	}
	if frequencyPb != nil {
		pb.Frequency = *frequencyPb
	}

	windowStartTimePb, err := clusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(st.WindowStartTime)
	if err != nil {
		return nil, err
	}
	if windowStartTimePb != nil {
		pb.WindowStartTime = windowStartTimePb
	}

	return pb, nil
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

type clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb struct {
	DayOfWeek ClusterAutoRestartMessageMaintenanceWindowDayOfWeek `json:"day_of_week,omitempty"`

	Frequency ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency `json:"frequency,omitempty"`

	WindowStartTime *clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb `json:"window_start_time,omitempty"`
}

func clusterAutoRestartMessageMaintenanceWindowWeekDayBasedScheduleFromPb(pb *clusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedulePb) (*ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindowWeekDayBasedSchedule{}
	dayOfWeekField, err := identity(&pb.DayOfWeek)
	if err != nil {
		return nil, err
	}
	if dayOfWeekField != nil {
		st.DayOfWeek = *dayOfWeekField
	}
	frequencyField, err := identity(&pb.Frequency)
	if err != nil {
		return nil, err
	}
	if frequencyField != nil {
		st.Frequency = *frequencyField
	}
	windowStartTimeField, err := clusterAutoRestartMessageMaintenanceWindowWindowStartTimeFromPb(pb.WindowStartTime)
	if err != nil {
		return nil, err
	}
	if windowStartTimeField != nil {
		st.WindowStartTime = windowStartTimeField
	}

	return st, nil
}

type ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency string
type clusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb string

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

func clusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyToPb(st *ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency) (*clusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := clusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb(*st)
	return &pb, nil
}

func clusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyFromPb(pb *clusterAutoRestartMessageMaintenanceWindowWeekDayFrequencyPb) (*ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency, error) {
	if pb == nil {
		return nil, nil
	}
	st := ClusterAutoRestartMessageMaintenanceWindowWeekDayFrequency(*pb)
	return &st, nil
}

type ClusterAutoRestartMessageMaintenanceWindowWindowStartTime struct {
	Hours int

	Minutes int

	ForceSendFields []string
}

func clusterAutoRestartMessageMaintenanceWindowWindowStartTimeToPb(st *ClusterAutoRestartMessageMaintenanceWindowWindowStartTime) (*clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb{}
	hoursPb, err := identity(&st.Hours)
	if err != nil {
		return nil, err
	}
	if hoursPb != nil {
		pb.Hours = *hoursPb
	}

	minutesPb, err := identity(&st.Minutes)
	if err != nil {
		return nil, err
	}
	if minutesPb != nil {
		pb.Minutes = *minutesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb struct {
	Hours int `json:"hours,omitempty"`

	Minutes int `json:"minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func clusterAutoRestartMessageMaintenanceWindowWindowStartTimeFromPb(pb *clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) (*ClusterAutoRestartMessageMaintenanceWindowWindowStartTime, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ClusterAutoRestartMessageMaintenanceWindowWindowStartTime{}
	hoursField, err := identity(&pb.Hours)
	if err != nil {
		return nil, err
	}
	if hoursField != nil {
		st.Hours = *hoursField
	}
	minutesField, err := identity(&pb.Minutes)
	if err != nil {
		return nil, err
	}
	if minutesField != nil {
		st.Minutes = *minutesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st clusterAutoRestartMessageMaintenanceWindowWindowStartTimePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// SHIELD feature: CSP
type ComplianceSecurityProfile struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards []ComplianceStandard

	IsEnabled bool

	ForceSendFields []string
}

func complianceSecurityProfileToPb(st *ComplianceSecurityProfile) (*complianceSecurityProfilePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &complianceSecurityProfilePb{}

	var complianceStandardsPb []ComplianceStandard
	for _, item := range st.ComplianceStandards {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			complianceStandardsPb = append(complianceStandardsPb, *itemPb)
		}
	}
	pb.ComplianceStandards = complianceStandardsPb

	isEnabledPb, err := identity(&st.IsEnabled)
	if err != nil {
		return nil, err
	}
	if isEnabledPb != nil {
		pb.IsEnabled = *isEnabledPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type complianceSecurityProfilePb struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`

	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func complianceSecurityProfileFromPb(pb *complianceSecurityProfilePb) (*ComplianceSecurityProfile, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfile{}

	var complianceStandardsField []ComplianceStandard
	for _, itemPb := range pb.ComplianceStandards {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			complianceStandardsField = append(complianceStandardsField, *item)
		}
	}
	st.ComplianceStandards = complianceStandardsField
	isEnabledField, err := identity(&pb.IsEnabled)
	if err != nil {
		return nil, err
	}
	if isEnabledField != nil {
		st.IsEnabled = *isEnabledField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *complianceSecurityProfilePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st complianceSecurityProfilePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ComplianceSecurityProfileSetting struct {
	// SHIELD feature: CSP
	ComplianceSecurityProfileWorkspace ComplianceSecurityProfile
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func complianceSecurityProfileSettingToPb(st *ComplianceSecurityProfileSetting) (*complianceSecurityProfileSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &complianceSecurityProfileSettingPb{}
	complianceSecurityProfileWorkspacePb, err := complianceSecurityProfileToPb(&st.ComplianceSecurityProfileWorkspace)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfileWorkspacePb != nil {
		pb.ComplianceSecurityProfileWorkspace = *complianceSecurityProfileWorkspacePb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type complianceSecurityProfileSettingPb struct {
	// SHIELD feature: CSP
	ComplianceSecurityProfileWorkspace complianceSecurityProfilePb `json:"compliance_security_profile_workspace"`
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

func complianceSecurityProfileSettingFromPb(pb *complianceSecurityProfileSettingPb) (*ComplianceSecurityProfileSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ComplianceSecurityProfileSetting{}
	complianceSecurityProfileWorkspaceField, err := complianceSecurityProfileFromPb(&pb.ComplianceSecurityProfileWorkspace)
	if err != nil {
		return nil, err
	}
	if complianceSecurityProfileWorkspaceField != nil {
		st.ComplianceSecurityProfileWorkspace = *complianceSecurityProfileWorkspaceField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *complianceSecurityProfileSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st complianceSecurityProfileSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Compliance stardard for SHIELD customers
type ComplianceStandard string
type ComplianceStandardPb string

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

func ComplianceStandardToPb(st *ComplianceStandard) (*ComplianceStandardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := ComplianceStandardPb(*st)
	return &pb, nil
}

func ComplianceStandardFromPb(pb *ComplianceStandardPb) (*ComplianceStandard, error) {
	if pb == nil {
		return nil, nil
	}
	st := ComplianceStandard(*pb)
	return &st, nil
}

type Config struct {
	Email *EmailConfig

	GenericWebhook *GenericWebhookConfig

	MicrosoftTeams *MicrosoftTeamsConfig

	Pagerduty *PagerdutyConfig

	Slack *SlackConfig
}

func configToPb(st *Config) (*configPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &configPb{}
	emailPb, err := emailConfigToPb(st.Email)
	if err != nil {
		return nil, err
	}
	if emailPb != nil {
		pb.Email = emailPb
	}

	genericWebhookPb, err := genericWebhookConfigToPb(st.GenericWebhook)
	if err != nil {
		return nil, err
	}
	if genericWebhookPb != nil {
		pb.GenericWebhook = genericWebhookPb
	}

	microsoftTeamsPb, err := microsoftTeamsConfigToPb(st.MicrosoftTeams)
	if err != nil {
		return nil, err
	}
	if microsoftTeamsPb != nil {
		pb.MicrosoftTeams = microsoftTeamsPb
	}

	pagerdutyPb, err := pagerdutyConfigToPb(st.Pagerduty)
	if err != nil {
		return nil, err
	}
	if pagerdutyPb != nil {
		pb.Pagerduty = pagerdutyPb
	}

	slackPb, err := slackConfigToPb(st.Slack)
	if err != nil {
		return nil, err
	}
	if slackPb != nil {
		pb.Slack = slackPb
	}

	return pb, nil
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

type configPb struct {
	Email *emailConfigPb `json:"email,omitempty"`

	GenericWebhook *genericWebhookConfigPb `json:"generic_webhook,omitempty"`

	MicrosoftTeams *microsoftTeamsConfigPb `json:"microsoft_teams,omitempty"`

	Pagerduty *pagerdutyConfigPb `json:"pagerduty,omitempty"`

	Slack *slackConfigPb `json:"slack,omitempty"`
}

func configFromPb(pb *configPb) (*Config, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Config{}
	emailField, err := emailConfigFromPb(pb.Email)
	if err != nil {
		return nil, err
	}
	if emailField != nil {
		st.Email = emailField
	}
	genericWebhookField, err := genericWebhookConfigFromPb(pb.GenericWebhook)
	if err != nil {
		return nil, err
	}
	if genericWebhookField != nil {
		st.GenericWebhook = genericWebhookField
	}
	microsoftTeamsField, err := microsoftTeamsConfigFromPb(pb.MicrosoftTeams)
	if err != nil {
		return nil, err
	}
	if microsoftTeamsField != nil {
		st.MicrosoftTeams = microsoftTeamsField
	}
	pagerdutyField, err := pagerdutyConfigFromPb(pb.Pagerduty)
	if err != nil {
		return nil, err
	}
	if pagerdutyField != nil {
		st.Pagerduty = pagerdutyField
	}
	slackField, err := slackConfigFromPb(pb.Slack)
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
	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	Label string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType
}

func createIpAccessListToPb(st *CreateIpAccessList) (*createIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createIpAccessListPb{}

	var ipAddressesPb []string
	for _, item := range st.IpAddresses {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			ipAddressesPb = append(ipAddressesPb, *itemPb)
		}
	}
	pb.IpAddresses = ipAddressesPb

	labelPb, err := identity(&st.Label)
	if err != nil {
		return nil, err
	}
	if labelPb != nil {
		pb.Label = *labelPb
	}

	listTypePb, err := identity(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}

	return pb, nil
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

type createIpAccessListPb struct {
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

func createIpAccessListFromPb(pb *createIpAccessListPb) (*CreateIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateIpAccessList{}

	var ipAddressesField []string
	for _, itemPb := range pb.IpAddresses {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			ipAddressesField = append(ipAddressesField, *item)
		}
	}
	st.IpAddresses = ipAddressesField
	labelField, err := identity(&pb.Label)
	if err != nil {
		return nil, err
	}
	if labelField != nil {
		st.Label = *labelField
	}
	listTypeField, err := identity(&pb.ListType)
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
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo
}

func createIpAccessListResponseToPb(st *CreateIpAccessListResponse) (*createIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createIpAccessListResponsePb{}
	ipAccessListPb, err := ipAccessListInfoToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	return pb, nil
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

type createIpAccessListResponsePb struct {
	// Definition of an IP Access list
	IpAccessList *ipAccessListInfoPb `json:"ip_access_list,omitempty"`
}

func createIpAccessListResponseFromPb(pb *createIpAccessListResponsePb) (*CreateIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateIpAccessListResponse{}
	ipAccessListField, err := ipAccessListInfoFromPb(pb.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListField != nil {
		st.IpAccessList = ipAccessListField
	}

	return st, nil
}

// Create a network connectivity configuration
type CreateNetworkConnectivityConfigRequest struct {
	// Properties of the new network connectivity configuration.
	NetworkConnectivityConfig CreateNetworkConnectivityConfiguration
}

func createNetworkConnectivityConfigRequestToPb(st *CreateNetworkConnectivityConfigRequest) (*createNetworkConnectivityConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNetworkConnectivityConfigRequestPb{}
	networkConnectivityConfigPb, err := createNetworkConnectivityConfigurationToPb(&st.NetworkConnectivityConfig)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigPb != nil {
		pb.NetworkConnectivityConfig = *networkConnectivityConfigPb
	}

	return pb, nil
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

type createNetworkConnectivityConfigRequestPb struct {
	// Properties of the new network connectivity configuration.
	NetworkConnectivityConfig createNetworkConnectivityConfigurationPb `json:"network_connectivity_config"`
}

func createNetworkConnectivityConfigRequestFromPb(pb *createNetworkConnectivityConfigRequestPb) (*CreateNetworkConnectivityConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkConnectivityConfigRequest{}
	networkConnectivityConfigField, err := createNetworkConnectivityConfigurationFromPb(&pb.NetworkConnectivityConfig)
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
	Name string
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string
}

func createNetworkConnectivityConfigurationToPb(st *CreateNetworkConnectivityConfiguration) (*createNetworkConnectivityConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNetworkConnectivityConfigurationPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	return pb, nil
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

type createNetworkConnectivityConfigurationPb struct {
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	Name string `json:"name"`
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string `json:"region"`
}

func createNetworkConnectivityConfigurationFromPb(pb *createNetworkConnectivityConfigurationPb) (*CreateNetworkConnectivityConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNetworkConnectivityConfiguration{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}

	return st, nil
}

type CreateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *Config
	// The display name for the notification destination.
	DisplayName string

	ForceSendFields []string
}

func createNotificationDestinationRequestToPb(st *CreateNotificationDestinationRequest) (*createNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createNotificationDestinationRequestPb{}
	configPb, err := configToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createNotificationDestinationRequestPb struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *configPb `json:"config,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createNotificationDestinationRequestFromPb(pb *createNotificationDestinationRequestPb) (*CreateNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateNotificationDestinationRequest{}
	configField, err := configFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createNotificationDestinationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createNotificationDestinationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Configuration details for creating on-behalf tokens.
type CreateOboTokenRequest struct {
	// Application ID of the service principal.
	ApplicationId string
	// Comment that describes the purpose of the token.
	Comment string
	// The number of seconds before the token expires.
	LifetimeSeconds int64

	ForceSendFields []string
}

func createOboTokenRequestToPb(st *CreateOboTokenRequest) (*createOboTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createOboTokenRequestPb{}
	applicationIdPb, err := identity(&st.ApplicationId)
	if err != nil {
		return nil, err
	}
	if applicationIdPb != nil {
		pb.ApplicationId = *applicationIdPb
	}

	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	lifetimeSecondsPb, err := identity(&st.LifetimeSeconds)
	if err != nil {
		return nil, err
	}
	if lifetimeSecondsPb != nil {
		pb.LifetimeSeconds = *lifetimeSecondsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createOboTokenRequestPb struct {
	// Application ID of the service principal.
	ApplicationId string `json:"application_id"`
	// Comment that describes the purpose of the token.
	Comment string `json:"comment,omitempty"`
	// The number of seconds before the token expires.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createOboTokenRequestFromPb(pb *createOboTokenRequestPb) (*CreateOboTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOboTokenRequest{}
	applicationIdField, err := identity(&pb.ApplicationId)
	if err != nil {
		return nil, err
	}
	if applicationIdField != nil {
		st.ApplicationId = *applicationIdField
	}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	lifetimeSecondsField, err := identity(&pb.LifetimeSeconds)
	if err != nil {
		return nil, err
	}
	if lifetimeSecondsField != nil {
		st.LifetimeSeconds = *lifetimeSecondsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createOboTokenRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createOboTokenRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// An on-behalf token was successfully created for the service principal.
type CreateOboTokenResponse struct {
	TokenInfo *TokenInfo
	// Value of the token.
	TokenValue string

	ForceSendFields []string
}

func createOboTokenResponseToPb(st *CreateOboTokenResponse) (*createOboTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createOboTokenResponsePb{}
	tokenInfoPb, err := tokenInfoToPb(st.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoPb != nil {
		pb.TokenInfo = tokenInfoPb
	}

	tokenValuePb, err := identity(&st.TokenValue)
	if err != nil {
		return nil, err
	}
	if tokenValuePb != nil {
		pb.TokenValue = *tokenValuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createOboTokenResponsePb struct {
	TokenInfo *tokenInfoPb `json:"token_info,omitempty"`
	// Value of the token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createOboTokenResponseFromPb(pb *createOboTokenResponsePb) (*CreateOboTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOboTokenResponse{}
	tokenInfoField, err := tokenInfoFromPb(pb.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoField != nil {
		st.TokenInfo = tokenInfoField
	}
	tokenValueField, err := identity(&pb.TokenValue)
	if err != nil {
		return nil, err
	}
	if tokenValueField != nil {
		st.TokenValue = *tokenValueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createOboTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createOboTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Properties of the new private endpoint rule. Note that you must approve the
// endpoint in Azure portal after initialization.
type CreatePrivateEndpointRule struct {
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string
	// Only used by private endpoints to Azure first-party services. Enum: blob
	// | dfs | sqlServer | mysqlServer
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId string
	// The Azure resource ID of the target resource.
	ResourceId string

	ForceSendFields []string
}

func createPrivateEndpointRuleToPb(st *CreatePrivateEndpointRule) (*createPrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPrivateEndpointRulePb{}

	var domainNamesPb []string
	for _, item := range st.DomainNames {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			domainNamesPb = append(domainNamesPb, *itemPb)
		}
	}
	pb.DomainNames = domainNamesPb

	groupIdPb, err := identity(&st.GroupId)
	if err != nil {
		return nil, err
	}
	if groupIdPb != nil {
		pb.GroupId = *groupIdPb
	}

	resourceIdPb, err := identity(&st.ResourceId)
	if err != nil {
		return nil, err
	}
	if resourceIdPb != nil {
		pb.ResourceId = *resourceIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createPrivateEndpointRulePb struct {
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string `json:"domain_names,omitempty"`
	// Only used by private endpoints to Azure first-party services. Enum: blob
	// | dfs | sqlServer | mysqlServer
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId string `json:"group_id,omitempty"`
	// The Azure resource ID of the target resource.
	ResourceId string `json:"resource_id"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPrivateEndpointRuleFromPb(pb *createPrivateEndpointRulePb) (*CreatePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePrivateEndpointRule{}

	var domainNamesField []string
	for _, itemPb := range pb.DomainNames {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			domainNamesField = append(domainNamesField, *item)
		}
	}
	st.DomainNames = domainNamesField
	groupIdField, err := identity(&pb.GroupId)
	if err != nil {
		return nil, err
	}
	if groupIdField != nil {
		st.GroupId = *groupIdField
	}
	resourceIdField, err := identity(&pb.ResourceId)
	if err != nil {
		return nil, err
	}
	if resourceIdField != nil {
		st.ResourceId = *resourceIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Create a private endpoint rule
type CreatePrivateEndpointRuleRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
	PrivateEndpointRule CreatePrivateEndpointRule
}

func createPrivateEndpointRuleRequestToPb(st *CreatePrivateEndpointRuleRequest) (*createPrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPrivateEndpointRuleRequestPb{}
	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	privateEndpointRulePb, err := createPrivateEndpointRuleToPb(&st.PrivateEndpointRule)
	if err != nil {
		return nil, err
	}
	if privateEndpointRulePb != nil {
		pb.PrivateEndpointRule = *privateEndpointRulePb
	}

	return pb, nil
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

type createPrivateEndpointRuleRequestPb struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
	PrivateEndpointRule createPrivateEndpointRulePb `json:"private_endpoint_rule"`
}

func createPrivateEndpointRuleRequestFromPb(pb *createPrivateEndpointRuleRequestPb) (*CreatePrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePrivateEndpointRuleRequest{}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	privateEndpointRuleField, err := createPrivateEndpointRuleFromPb(&pb.PrivateEndpointRule)
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
	Comment string
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds int64

	ForceSendFields []string
}

func createTokenRequestToPb(st *CreateTokenRequest) (*createTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTokenRequestPb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	lifetimeSecondsPb, err := identity(&st.LifetimeSeconds)
	if err != nil {
		return nil, err
	}
	if lifetimeSecondsPb != nil {
		pb.LifetimeSeconds = *lifetimeSecondsPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createTokenRequestPb struct {
	// Optional description to attach to the token.
	Comment string `json:"comment,omitempty"`
	// The lifetime of the token, in seconds.
	//
	// If the lifetime is not specified, this token remains valid indefinitely.
	LifetimeSeconds int64 `json:"lifetime_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createTokenRequestFromPb(pb *createTokenRequestPb) (*CreateTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTokenRequest{}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	lifetimeSecondsField, err := identity(&pb.LifetimeSeconds)
	if err != nil {
		return nil, err
	}
	if lifetimeSecondsField != nil {
		st.LifetimeSeconds = *lifetimeSecondsField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createTokenRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createTokenRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateTokenResponse struct {
	// The information for the new token.
	TokenInfo *PublicTokenInfo
	// The value of the new token.
	TokenValue string

	ForceSendFields []string
}

func createTokenResponseToPb(st *CreateTokenResponse) (*createTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTokenResponsePb{}
	tokenInfoPb, err := publicTokenInfoToPb(st.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoPb != nil {
		pb.TokenInfo = tokenInfoPb
	}

	tokenValuePb, err := identity(&st.TokenValue)
	if err != nil {
		return nil, err
	}
	if tokenValuePb != nil {
		pb.TokenValue = *tokenValuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type createTokenResponsePb struct {
	// The information for the new token.
	TokenInfo *publicTokenInfoPb `json:"token_info,omitempty"`
	// The value of the new token.
	TokenValue string `json:"token_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createTokenResponseFromPb(pb *createTokenResponsePb) (*CreateTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTokenResponse{}
	tokenInfoField, err := publicTokenInfoFromPb(pb.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoField != nil {
		st.TokenInfo = tokenInfoField
	}
	tokenValueField, err := identity(&pb.TokenValue)
	if err != nil {
		return nil, err
	}
	if tokenValueField != nil {
		st.TokenValue = *tokenValueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createTokenResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createTokenResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Account level policy for CSP
type CspEnablementAccount struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards []ComplianceStandard
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced bool

	ForceSendFields []string
}

func cspEnablementAccountToPb(st *CspEnablementAccount) (*cspEnablementAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cspEnablementAccountPb{}

	var complianceStandardsPb []ComplianceStandard
	for _, item := range st.ComplianceStandards {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			complianceStandardsPb = append(complianceStandardsPb, *itemPb)
		}
	}
	pb.ComplianceStandards = complianceStandardsPb

	isEnforcedPb, err := identity(&st.IsEnforced)
	if err != nil {
		return nil, err
	}
	if isEnforcedPb != nil {
		pb.IsEnforced = *isEnforcedPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type cspEnablementAccountPb struct {
	// Set by customers when they request Compliance Security Profile (CSP)
	// Invariants are enforced in Settings policy.
	ComplianceStandards []ComplianceStandard `json:"compliance_standards,omitempty"`
	// Enforced = it cannot be overriden at workspace level.
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cspEnablementAccountFromPb(pb *cspEnablementAccountPb) (*CspEnablementAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CspEnablementAccount{}

	var complianceStandardsField []ComplianceStandard
	for _, itemPb := range pb.ComplianceStandards {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			complianceStandardsField = append(complianceStandardsField, *item)
		}
	}
	st.ComplianceStandards = complianceStandardsField
	isEnforcedField, err := identity(&pb.IsEnforced)
	if err != nil {
		return nil, err
	}
	if isEnforcedField != nil {
		st.IsEnforced = *isEnforcedField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cspEnablementAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cspEnablementAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CspEnablementAccountSetting struct {
	// Account level policy for CSP
	CspEnablementAccount CspEnablementAccount
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func cspEnablementAccountSettingToPb(st *CspEnablementAccountSetting) (*cspEnablementAccountSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cspEnablementAccountSettingPb{}
	cspEnablementAccountPb, err := cspEnablementAccountToPb(&st.CspEnablementAccount)
	if err != nil {
		return nil, err
	}
	if cspEnablementAccountPb != nil {
		pb.CspEnablementAccount = *cspEnablementAccountPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type cspEnablementAccountSettingPb struct {
	// Account level policy for CSP
	CspEnablementAccount cspEnablementAccountPb `json:"csp_enablement_account"`
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

func cspEnablementAccountSettingFromPb(pb *cspEnablementAccountSettingPb) (*CspEnablementAccountSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CspEnablementAccountSetting{}
	cspEnablementAccountField, err := cspEnablementAccountFromPb(&pb.CspEnablementAccount)
	if err != nil {
		return nil, err
	}
	if cspEnablementAccountField != nil {
		st.CspEnablementAccount = *cspEnablementAccountField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cspEnablementAccountSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cspEnablementAccountSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	Namespace StringMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func defaultNamespaceSettingToPb(st *DefaultNamespaceSetting) (*defaultNamespaceSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &defaultNamespaceSettingPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	namespacePb, err := stringMessageToPb(&st.Namespace)
	if err != nil {
		return nil, err
	}
	if namespacePb != nil {
		pb.Namespace = *namespacePb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type defaultNamespaceSettingPb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	Namespace stringMessagePb `json:"namespace"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func defaultNamespaceSettingFromPb(pb *defaultNamespaceSettingPb) (*DefaultNamespaceSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DefaultNamespaceSetting{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	namespaceField, err := stringMessageFromPb(&pb.Namespace)
	if err != nil {
		return nil, err
	}
	if namespaceField != nil {
		st.Namespace = *namespaceField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *defaultNamespaceSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st defaultNamespaceSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func deleteAccountIpAccessEnableRequestToPb(st *DeleteAccountIpAccessEnableRequest) (*deleteAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountIpAccessEnableRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteAccountIpAccessEnableRequestPb struct {
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

func deleteAccountIpAccessEnableRequestFromPb(pb *deleteAccountIpAccessEnableRequestPb) (*DeleteAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessEnableRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAccountIpAccessEnableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAccountIpAccessEnableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteAccountIpAccessEnableResponseToPb(st *DeleteAccountIpAccessEnableResponse) (*deleteAccountIpAccessEnableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountIpAccessEnableResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteAccountIpAccessEnableResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteAccountIpAccessEnableResponseFromPb(pb *deleteAccountIpAccessEnableResponsePb) (*DeleteAccountIpAccessEnableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessEnableResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
}

// Delete access list
type DeleteAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string
}

func deleteAccountIpAccessListRequestToPb(st *DeleteAccountIpAccessListRequest) (*deleteAccountIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountIpAccessListRequestPb{}
	ipAccessListIdPb, err := identity(&st.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdPb != nil {
		pb.IpAccessListId = *ipAccessListIdPb
	}

	return pb, nil
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

type deleteAccountIpAccessListRequestPb struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func deleteAccountIpAccessListRequestFromPb(pb *deleteAccountIpAccessListRequestPb) (*DeleteAccountIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountIpAccessListRequest{}
	ipAccessListIdField, err := identity(&pb.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdField != nil {
		st.IpAccessListId = *ipAccessListIdField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deleteAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *DeleteAibiDashboardEmbeddingAccessPolicySettingRequest) (*deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
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

func deleteAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*DeleteAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingAccessPolicySettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAibiDashboardEmbeddingAccessPolicySettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteAibiDashboardEmbeddingAccessPolicySettingResponseToPb(st *DeleteAibiDashboardEmbeddingAccessPolicySettingResponse) (*deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteAibiDashboardEmbeddingAccessPolicySettingResponseFromPb(pb *deleteAibiDashboardEmbeddingAccessPolicySettingResponsePb) (*DeleteAibiDashboardEmbeddingAccessPolicySettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingAccessPolicySettingResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
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

func deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseToPb(st *DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse) (*deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteAibiDashboardEmbeddingApprovedDomainsSettingResponseFromPb(pb *deleteAibiDashboardEmbeddingApprovedDomainsSettingResponsePb) (*DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAibiDashboardEmbeddingApprovedDomainsSettingResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deleteDefaultNamespaceSettingRequestToPb(st *DeleteDefaultNamespaceSettingRequest) (*deleteDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDefaultNamespaceSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteDefaultNamespaceSettingRequestPb struct {
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

func deleteDefaultNamespaceSettingRequestFromPb(pb *deleteDefaultNamespaceSettingRequestPb) (*DeleteDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultNamespaceSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDefaultNamespaceSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDefaultNamespaceSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteDefaultNamespaceSettingResponseToPb(st *DeleteDefaultNamespaceSettingResponse) (*deleteDefaultNamespaceSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDefaultNamespaceSettingResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteDefaultNamespaceSettingResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteDefaultNamespaceSettingResponseFromPb(pb *deleteDefaultNamespaceSettingResponsePb) (*DeleteDefaultNamespaceSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDefaultNamespaceSettingResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deleteDisableLegacyAccessRequestToPb(st *DeleteDisableLegacyAccessRequest) (*deleteDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyAccessRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteDisableLegacyAccessRequestPb struct {
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

func deleteDisableLegacyAccessRequestFromPb(pb *deleteDisableLegacyAccessRequestPb) (*DeleteDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyAccessRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDisableLegacyAccessRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDisableLegacyAccessRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteDisableLegacyAccessResponseToPb(st *DeleteDisableLegacyAccessResponse) (*deleteDisableLegacyAccessResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyAccessResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteDisableLegacyAccessResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteDisableLegacyAccessResponseFromPb(pb *deleteDisableLegacyAccessResponsePb) (*DeleteDisableLegacyAccessResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyAccessResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deleteDisableLegacyDbfsRequestToPb(st *DeleteDisableLegacyDbfsRequest) (*deleteDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyDbfsRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteDisableLegacyDbfsRequestPb struct {
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

func deleteDisableLegacyDbfsRequestFromPb(pb *deleteDisableLegacyDbfsRequestPb) (*DeleteDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyDbfsRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDisableLegacyDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDisableLegacyDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteDisableLegacyDbfsResponseToPb(st *DeleteDisableLegacyDbfsResponse) (*deleteDisableLegacyDbfsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyDbfsResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteDisableLegacyDbfsResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteDisableLegacyDbfsResponseFromPb(pb *deleteDisableLegacyDbfsResponsePb) (*DeleteDisableLegacyDbfsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyDbfsResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deleteDisableLegacyFeaturesRequestToPb(st *DeleteDisableLegacyFeaturesRequest) (*deleteDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyFeaturesRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteDisableLegacyFeaturesRequestPb struct {
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

func deleteDisableLegacyFeaturesRequestFromPb(pb *deleteDisableLegacyFeaturesRequestPb) (*DeleteDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyFeaturesRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteDisableLegacyFeaturesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteDisableLegacyFeaturesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteDisableLegacyFeaturesResponseToPb(st *DeleteDisableLegacyFeaturesResponse) (*deleteDisableLegacyFeaturesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDisableLegacyFeaturesResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteDisableLegacyFeaturesResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteDisableLegacyFeaturesResponseFromPb(pb *deleteDisableLegacyFeaturesResponsePb) (*DeleteDisableLegacyFeaturesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDisableLegacyFeaturesResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
}

// Delete access list
type DeleteIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string
}

func deleteIpAccessListRequestToPb(st *DeleteIpAccessListRequest) (*deleteIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteIpAccessListRequestPb{}
	ipAccessListIdPb, err := identity(&st.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdPb != nil {
		pb.IpAccessListId = *ipAccessListIdPb
	}

	return pb, nil
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

type deleteIpAccessListRequestPb struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func deleteIpAccessListRequestFromPb(pb *deleteIpAccessListRequestPb) (*DeleteIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteIpAccessListRequest{}
	ipAccessListIdField, err := identity(&pb.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdField != nil {
		st.IpAccessListId = *ipAccessListIdField
	}

	return st, nil
}

// Delete a network connectivity configuration
type DeleteNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string
}

func deleteNetworkConnectivityConfigurationRequestToPb(st *DeleteNetworkConnectivityConfigurationRequest) (*deleteNetworkConnectivityConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkConnectivityConfigurationRequestPb{}
	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	return pb, nil
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

type deleteNetworkConnectivityConfigurationRequestPb struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

func deleteNetworkConnectivityConfigurationRequestFromPb(pb *deleteNetworkConnectivityConfigurationRequestPb) (*DeleteNetworkConnectivityConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkConnectivityConfigurationRequest{}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}

	return st, nil
}

type DeleteNetworkConnectivityConfigurationResponse struct {
}

func deleteNetworkConnectivityConfigurationResponseToPb(st *DeleteNetworkConnectivityConfigurationResponse) (*deleteNetworkConnectivityConfigurationResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNetworkConnectivityConfigurationResponsePb{}

	return pb, nil
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

type deleteNetworkConnectivityConfigurationResponsePb struct {
}

func deleteNetworkConnectivityConfigurationResponseFromPb(pb *deleteNetworkConnectivityConfigurationResponsePb) (*DeleteNetworkConnectivityConfigurationResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNetworkConnectivityConfigurationResponse{}

	return st, nil
}

// Delete a notification destination
type DeleteNotificationDestinationRequest struct {
	Id string
}

func deleteNotificationDestinationRequestToPb(st *DeleteNotificationDestinationRequest) (*deleteNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteNotificationDestinationRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	return pb, nil
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

type deleteNotificationDestinationRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteNotificationDestinationRequestFromPb(pb *deleteNotificationDestinationRequestPb) (*DeleteNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteNotificationDestinationRequest{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deletePersonalComputeSettingRequestToPb(st *DeletePersonalComputeSettingRequest) (*deletePersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePersonalComputeSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deletePersonalComputeSettingRequestPb struct {
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

func deletePersonalComputeSettingRequestFromPb(pb *deletePersonalComputeSettingRequestPb) (*DeletePersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePersonalComputeSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deletePersonalComputeSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deletePersonalComputeSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deletePersonalComputeSettingResponseToPb(st *DeletePersonalComputeSettingResponse) (*deletePersonalComputeSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePersonalComputeSettingResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deletePersonalComputeSettingResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deletePersonalComputeSettingResponseFromPb(pb *deletePersonalComputeSettingResponsePb) (*DeletePersonalComputeSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePersonalComputeSettingResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
}

// Delete a private endpoint rule
type DeletePrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string
}

func deletePrivateEndpointRuleRequestToPb(st *DeletePrivateEndpointRuleRequest) (*deletePrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePrivateEndpointRuleRequestPb{}
	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	privateEndpointRuleIdPb, err := identity(&st.PrivateEndpointRuleId)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleIdPb != nil {
		pb.PrivateEndpointRuleId = *privateEndpointRuleIdPb
	}

	return pb, nil
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

type deletePrivateEndpointRuleRequestPb struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

func deletePrivateEndpointRuleRequestFromPb(pb *deletePrivateEndpointRuleRequestPb) (*DeletePrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePrivateEndpointRuleRequest{}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	privateEndpointRuleIdField, err := identity(&pb.PrivateEndpointRuleId)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleIdField != nil {
		st.PrivateEndpointRuleId = *privateEndpointRuleIdField
	}

	return st, nil
}

type DeleteResponse struct {
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func deleteRestrictWorkspaceAdminsSettingRequestToPb(st *DeleteRestrictWorkspaceAdminsSettingRequest) (*deleteRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRestrictWorkspaceAdminsSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type deleteRestrictWorkspaceAdminsSettingRequestPb struct {
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

func deleteRestrictWorkspaceAdminsSettingRequestFromPb(pb *deleteRestrictWorkspaceAdminsSettingRequestPb) (*DeleteRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRestrictWorkspaceAdminsSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteRestrictWorkspaceAdminsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteRestrictWorkspaceAdminsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string
}

func deleteRestrictWorkspaceAdminsSettingResponseToPb(st *DeleteRestrictWorkspaceAdminsSettingResponse) (*deleteRestrictWorkspaceAdminsSettingResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRestrictWorkspaceAdminsSettingResponsePb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	return pb, nil
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

type deleteRestrictWorkspaceAdminsSettingResponsePb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// delete pattern to perform setting deletions in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// DELETE request to identify the rule set version you are deleting.
	Etag string `json:"etag"`
}

func deleteRestrictWorkspaceAdminsSettingResponseFromPb(pb *deleteRestrictWorkspaceAdminsSettingResponsePb) (*DeleteRestrictWorkspaceAdminsSettingResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRestrictWorkspaceAdminsSettingResponse{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	return st, nil
}

// Delete a token
type DeleteTokenManagementRequest struct {
	// The ID of the token to revoke.
	TokenId string
}

func deleteTokenManagementRequestToPb(st *DeleteTokenManagementRequest) (*deleteTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTokenManagementRequestPb{}
	tokenIdPb, err := identity(&st.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdPb != nil {
		pb.TokenId = *tokenIdPb
	}

	return pb, nil
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

type deleteTokenManagementRequestPb struct {
	// The ID of the token to revoke.
	TokenId string `json:"-" url:"-"`
}

func deleteTokenManagementRequestFromPb(pb *deleteTokenManagementRequestPb) (*DeleteTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTokenManagementRequest{}
	tokenIdField, err := identity(&pb.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdField != nil {
		st.TokenId = *tokenIdField
	}

	return st, nil
}

type DestinationType string
type destinationTypePb string

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

func destinationTypeToPb(st *DestinationType) (*destinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := destinationTypePb(*st)
	return &pb, nil
}

func destinationTypeFromPb(pb *destinationTypePb) (*DestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := DestinationType(*pb)
	return &st, nil
}

type DisableLegacyAccess struct {
	DisableLegacyAccess BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func disableLegacyAccessToPb(st *DisableLegacyAccess) (*disableLegacyAccessPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableLegacyAccessPb{}
	disableLegacyAccessPb, err := booleanMessageToPb(&st.DisableLegacyAccess)
	if err != nil {
		return nil, err
	}
	if disableLegacyAccessPb != nil {
		pb.DisableLegacyAccess = *disableLegacyAccessPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type disableLegacyAccessPb struct {
	DisableLegacyAccess booleanMessagePb `json:"disable_legacy_access"`
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

func disableLegacyAccessFromPb(pb *disableLegacyAccessPb) (*DisableLegacyAccess, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyAccess{}
	disableLegacyAccessField, err := booleanMessageFromPb(&pb.DisableLegacyAccess)
	if err != nil {
		return nil, err
	}
	if disableLegacyAccessField != nil {
		st.DisableLegacyAccess = *disableLegacyAccessField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *disableLegacyAccessPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st disableLegacyAccessPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DisableLegacyDbfs struct {
	DisableLegacyDbfs BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func disableLegacyDbfsToPb(st *DisableLegacyDbfs) (*disableLegacyDbfsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableLegacyDbfsPb{}
	disableLegacyDbfsPb, err := booleanMessageToPb(&st.DisableLegacyDbfs)
	if err != nil {
		return nil, err
	}
	if disableLegacyDbfsPb != nil {
		pb.DisableLegacyDbfs = *disableLegacyDbfsPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type disableLegacyDbfsPb struct {
	DisableLegacyDbfs booleanMessagePb `json:"disable_legacy_dbfs"`
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

func disableLegacyDbfsFromPb(pb *disableLegacyDbfsPb) (*DisableLegacyDbfs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyDbfs{}
	disableLegacyDbfsField, err := booleanMessageFromPb(&pb.DisableLegacyDbfs)
	if err != nil {
		return nil, err
	}
	if disableLegacyDbfsField != nil {
		st.DisableLegacyDbfs = *disableLegacyDbfsField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *disableLegacyDbfsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st disableLegacyDbfsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DisableLegacyFeatures struct {
	DisableLegacyFeatures BooleanMessage
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func disableLegacyFeaturesToPb(st *DisableLegacyFeatures) (*disableLegacyFeaturesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &disableLegacyFeaturesPb{}
	disableLegacyFeaturesPb, err := booleanMessageToPb(&st.DisableLegacyFeatures)
	if err != nil {
		return nil, err
	}
	if disableLegacyFeaturesPb != nil {
		pb.DisableLegacyFeatures = *disableLegacyFeaturesPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type disableLegacyFeaturesPb struct {
	DisableLegacyFeatures booleanMessagePb `json:"disable_legacy_features"`
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

func disableLegacyFeaturesFromPb(pb *disableLegacyFeaturesPb) (*DisableLegacyFeatures, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DisableLegacyFeatures{}
	disableLegacyFeaturesField, err := booleanMessageFromPb(&pb.DisableLegacyFeatures)
	if err != nil {
		return nil, err
	}
	if disableLegacyFeaturesField != nil {
		st.DisableLegacyFeatures = *disableLegacyFeaturesField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *disableLegacyFeaturesPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st disableLegacyFeaturesPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The network policies applying for egress traffic. This message is used by the
// UI/REST API. We translate this message to the format expected by the
// dataplane in Lakehouse Network Manager (for the format expected by the
// dataplane, see networkconfig.textproto).
type EgressNetworkPolicy struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess *EgressNetworkPolicyInternetAccessPolicy
}

func EgressNetworkPolicyToPb(st *EgressNetworkPolicy) (*EgressNetworkPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &EgressNetworkPolicyPb{}
	internetAccessPb, err := egressNetworkPolicyInternetAccessPolicyToPb(st.InternetAccess)
	if err != nil {
		return nil, err
	}
	if internetAccessPb != nil {
		pb.InternetAccess = internetAccessPb
	}

	return pb, nil
}

func (st *EgressNetworkPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &EgressNetworkPolicyPb{}
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

func (st EgressNetworkPolicy) MarshalJSON() ([]byte, error) {
	pb, err := EgressNetworkPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EgressNetworkPolicyPb struct {
	// The access policy enforced for egress traffic to the internet.
	InternetAccess *egressNetworkPolicyInternetAccessPolicyPb `json:"internet_access,omitempty"`
}

func EgressNetworkPolicyFromPb(pb *EgressNetworkPolicyPb) (*EgressNetworkPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicy{}
	internetAccessField, err := egressNetworkPolicyInternetAccessPolicyFromPb(pb.InternetAccess)
	if err != nil {
		return nil, err
	}
	if internetAccessField != nil {
		st.InternetAccess = internetAccessField
	}

	return st, nil
}

type EgressNetworkPolicyInternetAccessPolicy struct {
	AllowedInternetDestinations []EgressNetworkPolicyInternetAccessPolicyInternetDestination

	AllowedStorageDestinations []EgressNetworkPolicyInternetAccessPolicyStorageDestination
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode
	// At which level can Databricks and Databricks managed compute access
	// Internet. FULL_ACCESS: Databricks can access Internet. No blocking rules
	// will apply. RESTRICTED_ACCESS: Databricks can only access explicitly
	// allowed internet and storage destinations, as well as UC connections and
	// external locations. PRIVATE_ACCESS_ONLY (not used): Databricks can only
	// access destinations via private link.
	RestrictionMode EgressNetworkPolicyInternetAccessPolicyRestrictionMode
}

func egressNetworkPolicyInternetAccessPolicyToPb(st *EgressNetworkPolicyInternetAccessPolicy) (*egressNetworkPolicyInternetAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyPb{}

	var allowedInternetDestinationsPb []egressNetworkPolicyInternetAccessPolicyInternetDestinationPb
	for _, item := range st.AllowedInternetDestinations {
		itemPb, err := egressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedInternetDestinationsPb = append(allowedInternetDestinationsPb, *itemPb)
		}
	}
	pb.AllowedInternetDestinations = allowedInternetDestinationsPb

	var allowedStorageDestinationsPb []egressNetworkPolicyInternetAccessPolicyStorageDestinationPb
	for _, item := range st.AllowedStorageDestinations {
		itemPb, err := egressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedStorageDestinationsPb = append(allowedStorageDestinationsPb, *itemPb)
		}
	}
	pb.AllowedStorageDestinations = allowedStorageDestinationsPb

	logOnlyModePb, err := egressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(st.LogOnlyMode)
	if err != nil {
		return nil, err
	}
	if logOnlyModePb != nil {
		pb.LogOnlyMode = logOnlyModePb
	}

	restrictionModePb, err := identity(&st.RestrictionMode)
	if err != nil {
		return nil, err
	}
	if restrictionModePb != nil {
		pb.RestrictionMode = *restrictionModePb
	}

	return pb, nil
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

type egressNetworkPolicyInternetAccessPolicyPb struct {
	AllowedInternetDestinations []egressNetworkPolicyInternetAccessPolicyInternetDestinationPb `json:"allowed_internet_destinations,omitempty"`

	AllowedStorageDestinations []egressNetworkPolicyInternetAccessPolicyStorageDestinationPb `json:"allowed_storage_destinations,omitempty"`
	// Optional. If not specified, assume the policy is enforced for all
	// workloads.
	LogOnlyMode *egressNetworkPolicyInternetAccessPolicyLogOnlyModePb `json:"log_only_mode,omitempty"`
	// At which level can Databricks and Databricks managed compute access
	// Internet. FULL_ACCESS: Databricks can access Internet. No blocking rules
	// will apply. RESTRICTED_ACCESS: Databricks can only access explicitly
	// allowed internet and storage destinations, as well as UC connections and
	// external locations. PRIVATE_ACCESS_ONLY (not used): Databricks can only
	// access destinations via private link.
	RestrictionMode EgressNetworkPolicyInternetAccessPolicyRestrictionMode `json:"restriction_mode,omitempty"`
}

func egressNetworkPolicyInternetAccessPolicyFromPb(pb *egressNetworkPolicyInternetAccessPolicyPb) (*EgressNetworkPolicyInternetAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicy{}

	var allowedInternetDestinationsField []EgressNetworkPolicyInternetAccessPolicyInternetDestination
	for _, itemPb := range pb.AllowedInternetDestinations {
		item, err := egressNetworkPolicyInternetAccessPolicyInternetDestinationFromPb(&itemPb)
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
		item, err := egressNetworkPolicyInternetAccessPolicyStorageDestinationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allowedStorageDestinationsField = append(allowedStorageDestinationsField, *item)
		}
	}
	st.AllowedStorageDestinations = allowedStorageDestinationsField
	logOnlyModeField, err := egressNetworkPolicyInternetAccessPolicyLogOnlyModeFromPb(pb.LogOnlyMode)
	if err != nil {
		return nil, err
	}
	if logOnlyModeField != nil {
		st.LogOnlyMode = logOnlyModeField
	}
	restrictionModeField, err := identity(&pb.RestrictionMode)
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
	Destination string
	// The filtering protocol used by the DP. For private and public preview,
	// SEG will only support TCP filtering (i.e. DNS based filtering, filtering
	// by destination IP address), so protocol will be set to TCP by default and
	// hidden from the user. In the future, users may be able to select HTTP
	// filtering (i.e. SNI based filtering, filtering by FQDN).
	Protocol EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol

	Type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType

	ForceSendFields []string
}

func egressNetworkPolicyInternetAccessPolicyInternetDestinationToPb(st *EgressNetworkPolicyInternetAccessPolicyInternetDestination) (*egressNetworkPolicyInternetAccessPolicyInternetDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyInternetDestinationPb{}
	destinationPb, err := identity(&st.Destination)
	if err != nil {
		return nil, err
	}
	if destinationPb != nil {
		pb.Destination = *destinationPb
	}

	protocolPb, err := identity(&st.Protocol)
	if err != nil {
		return nil, err
	}
	if protocolPb != nil {
		pb.Protocol = *protocolPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type egressNetworkPolicyInternetAccessPolicyInternetDestinationPb struct {
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

func egressNetworkPolicyInternetAccessPolicyInternetDestinationFromPb(pb *egressNetworkPolicyInternetAccessPolicyInternetDestinationPb) (*EgressNetworkPolicyInternetAccessPolicyInternetDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyInternetDestination{}
	destinationField, err := identity(&pb.Destination)
	if err != nil {
		return nil, err
	}
	if destinationField != nil {
		st.Destination = *destinationField
	}
	protocolField, err := identity(&pb.Protocol)
	if err != nil {
		return nil, err
	}
	if protocolField != nil {
		st.Protocol = *protocolField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *egressNetworkPolicyInternetAccessPolicyInternetDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st egressNetworkPolicyInternetAccessPolicyInternetDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The filtering protocol used by the DP. For private and public preview, SEG
// will only support TCP filtering (i.e. DNS based filtering, filtering by
// destination IP address), so protocol will be set to TCP by default and hidden
// from the user. In the future, users may be able to select HTTP filtering
// (i.e. SNI based filtering, filtering by FQDN).
type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol string
type egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb string

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

func egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolToPb(st *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol) (*egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb(*st)
	return &pb, nil
}

func egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolFromPb(pb *egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocolPb) (*EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationFilteringProtocol(*pb)
	return &st, nil
}

type EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType string
type egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb string

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

func egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType) (*egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb(*st)
	return &pb, nil
}

func egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypeFromPb(pb *egressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationTypePb) (*EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyInternetDestinationInternetDestinationType(*pb)
	return &st, nil
}

type EgressNetworkPolicyInternetAccessPolicyLogOnlyMode struct {
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType

	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType
}

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeToPb(st *EgressNetworkPolicyInternetAccessPolicyLogOnlyMode) (*egressNetworkPolicyInternetAccessPolicyLogOnlyModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyLogOnlyModePb{}
	logOnlyModeTypePb, err := identity(&st.LogOnlyModeType)
	if err != nil {
		return nil, err
	}
	if logOnlyModeTypePb != nil {
		pb.LogOnlyModeType = *logOnlyModeTypePb
	}

	var workloadsPb []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType
	for _, item := range st.Workloads {
		itemPb, err := identity(&item)
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

type egressNetworkPolicyInternetAccessPolicyLogOnlyModePb struct {
	LogOnlyModeType EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType `json:"log_only_mode_type,omitempty"`

	Workloads []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType `json:"workloads,omitempty"`
}

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeFromPb(pb *egressNetworkPolicyInternetAccessPolicyLogOnlyModePb) (*EgressNetworkPolicyInternetAccessPolicyLogOnlyMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyLogOnlyMode{}
	logOnlyModeTypeField, err := identity(&pb.LogOnlyModeType)
	if err != nil {
		return nil, err
	}
	if logOnlyModeTypeField != nil {
		st.LogOnlyModeType = *logOnlyModeTypeField
	}

	var workloadsField []EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType
	for _, itemPb := range pb.Workloads {
		item, err := identity(&itemPb)
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
type egressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb string

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

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType) (*egressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := egressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb(*st)
	return &pb, nil
}

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypeFromPb(pb *egressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeTypePb) (*EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyLogOnlyModeLogOnlyModeType(*pb)
	return &st, nil
}

// The values should match the list of workloads used in networkconfig.proto
type EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType string
type egressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb string

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

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType) (*egressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := egressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb(*st)
	return &pb, nil
}

func egressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypeFromPb(pb *egressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadTypePb) (*EgressNetworkPolicyInternetAccessPolicyLogOnlyModeWorkloadType, error) {
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
type egressNetworkPolicyInternetAccessPolicyRestrictionModePb string

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

func egressNetworkPolicyInternetAccessPolicyRestrictionModeToPb(st *EgressNetworkPolicyInternetAccessPolicyRestrictionMode) (*egressNetworkPolicyInternetAccessPolicyRestrictionModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := egressNetworkPolicyInternetAccessPolicyRestrictionModePb(*st)
	return &pb, nil
}

func egressNetworkPolicyInternetAccessPolicyRestrictionModeFromPb(pb *egressNetworkPolicyInternetAccessPolicyRestrictionModePb) (*EgressNetworkPolicyInternetAccessPolicyRestrictionMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyRestrictionMode(*pb)
	return &st, nil
}

// Users can specify accessible storage destinations.
type EgressNetworkPolicyInternetAccessPolicyStorageDestination struct {
	AllowedPaths []string

	AzureContainer string

	AzureDnsZone string

	AzureStorageAccount string

	AzureStorageService string

	BucketName string

	Region string

	Type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType

	ForceSendFields []string
}

func egressNetworkPolicyInternetAccessPolicyStorageDestinationToPb(st *EgressNetworkPolicyInternetAccessPolicyStorageDestination) (*egressNetworkPolicyInternetAccessPolicyStorageDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &egressNetworkPolicyInternetAccessPolicyStorageDestinationPb{}

	var allowedPathsPb []string
	for _, item := range st.AllowedPaths {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allowedPathsPb = append(allowedPathsPb, *itemPb)
		}
	}
	pb.AllowedPaths = allowedPathsPb

	azureContainerPb, err := identity(&st.AzureContainer)
	if err != nil {
		return nil, err
	}
	if azureContainerPb != nil {
		pb.AzureContainer = *azureContainerPb
	}

	azureDnsZonePb, err := identity(&st.AzureDnsZone)
	if err != nil {
		return nil, err
	}
	if azureDnsZonePb != nil {
		pb.AzureDnsZone = *azureDnsZonePb
	}

	azureStorageAccountPb, err := identity(&st.AzureStorageAccount)
	if err != nil {
		return nil, err
	}
	if azureStorageAccountPb != nil {
		pb.AzureStorageAccount = *azureStorageAccountPb
	}

	azureStorageServicePb, err := identity(&st.AzureStorageService)
	if err != nil {
		return nil, err
	}
	if azureStorageServicePb != nil {
		pb.AzureStorageService = *azureStorageServicePb
	}

	bucketNamePb, err := identity(&st.BucketName)
	if err != nil {
		return nil, err
	}
	if bucketNamePb != nil {
		pb.BucketName = *bucketNamePb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	typePb, err := identity(&st.Type)
	if err != nil {
		return nil, err
	}
	if typePb != nil {
		pb.Type = *typePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type egressNetworkPolicyInternetAccessPolicyStorageDestinationPb struct {
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

func egressNetworkPolicyInternetAccessPolicyStorageDestinationFromPb(pb *egressNetworkPolicyInternetAccessPolicyStorageDestinationPb) (*EgressNetworkPolicyInternetAccessPolicyStorageDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EgressNetworkPolicyInternetAccessPolicyStorageDestination{}

	var allowedPathsField []string
	for _, itemPb := range pb.AllowedPaths {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allowedPathsField = append(allowedPathsField, *item)
		}
	}
	st.AllowedPaths = allowedPathsField
	azureContainerField, err := identity(&pb.AzureContainer)
	if err != nil {
		return nil, err
	}
	if azureContainerField != nil {
		st.AzureContainer = *azureContainerField
	}
	azureDnsZoneField, err := identity(&pb.AzureDnsZone)
	if err != nil {
		return nil, err
	}
	if azureDnsZoneField != nil {
		st.AzureDnsZone = *azureDnsZoneField
	}
	azureStorageAccountField, err := identity(&pb.AzureStorageAccount)
	if err != nil {
		return nil, err
	}
	if azureStorageAccountField != nil {
		st.AzureStorageAccount = *azureStorageAccountField
	}
	azureStorageServiceField, err := identity(&pb.AzureStorageService)
	if err != nil {
		return nil, err
	}
	if azureStorageServiceField != nil {
		st.AzureStorageService = *azureStorageServiceField
	}
	bucketNameField, err := identity(&pb.BucketName)
	if err != nil {
		return nil, err
	}
	if bucketNameField != nil {
		st.BucketName = *bucketNameField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}
	typeField, err := identity(&pb.Type)
	if err != nil {
		return nil, err
	}
	if typeField != nil {
		st.Type = *typeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *egressNetworkPolicyInternetAccessPolicyStorageDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st egressNetworkPolicyInternetAccessPolicyStorageDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType string
type egressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb string

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

func egressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeToPb(st *EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType) (*egressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := egressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb(*st)
	return &pb, nil
}

func egressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypeFromPb(pb *egressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationTypePb) (*EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressNetworkPolicyInternetAccessPolicyStorageDestinationStorageDestinationType(*pb)
	return &st, nil
}

// The target resources that are supported by Network Connectivity Config. Note:
// some egress types can support general types that are not defined in
// EgressResourceType. E.g.: Azure private endpoint supports private link
// enabled Azure services.
type EgressResourceType string
type egressResourceTypePb string

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

func egressResourceTypeToPb(st *EgressResourceType) (*egressResourceTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := egressResourceTypePb(*st)
	return &pb, nil
}

func egressResourceTypeFromPb(pb *egressResourceTypePb) (*EgressResourceType, error) {
	if pb == nil {
		return nil, nil
	}
	st := EgressResourceType(*pb)
	return &st, nil
}

type EmailConfig struct {
	// Email addresses to notify.
	Addresses []string
}

func emailConfigToPb(st *EmailConfig) (*emailConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &emailConfigPb{}

	var addressesPb []string
	for _, item := range st.Addresses {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			addressesPb = append(addressesPb, *itemPb)
		}
	}
	pb.Addresses = addressesPb

	return pb, nil
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

type emailConfigPb struct {
	// Email addresses to notify.
	Addresses []string `json:"addresses,omitempty"`
}

func emailConfigFromPb(pb *emailConfigPb) (*EmailConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmailConfig{}

	var addressesField []string
	for _, itemPb := range pb.Addresses {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			addressesField = append(addressesField, *item)
		}
	}
	st.Addresses = addressesField

	return st, nil
}

type Empty struct {
}

func emptyToPb(st *Empty) (*emptyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &emptyPb{}

	return pb, nil
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

type emptyPb struct {
}

func emptyFromPb(pb *emptyPb) (*Empty, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Empty{}

	return st, nil
}

type EnableExportNotebook struct {
	BooleanVal *BooleanMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func enableExportNotebookToPb(st *EnableExportNotebook) (*enableExportNotebookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableExportNotebookPb{}
	booleanValPb, err := booleanMessageToPb(st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = booleanValPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enableExportNotebookPb struct {
	BooleanVal *booleanMessagePb `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enableExportNotebookFromPb(pb *enableExportNotebookPb) (*EnableExportNotebook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableExportNotebook{}
	booleanValField, err := booleanMessageFromPb(pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = booleanValField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enableExportNotebookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enableExportNotebookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnableNotebookTableClipboard struct {
	BooleanVal *BooleanMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func enableNotebookTableClipboardToPb(st *EnableNotebookTableClipboard) (*enableNotebookTableClipboardPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableNotebookTableClipboardPb{}
	booleanValPb, err := booleanMessageToPb(st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = booleanValPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enableNotebookTableClipboardPb struct {
	BooleanVal *booleanMessagePb `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enableNotebookTableClipboardFromPb(pb *enableNotebookTableClipboardPb) (*EnableNotebookTableClipboard, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableNotebookTableClipboard{}
	booleanValField, err := booleanMessageFromPb(pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = booleanValField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enableNotebookTableClipboardPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enableNotebookTableClipboardPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnableResultsDownloading struct {
	BooleanVal *BooleanMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func enableResultsDownloadingToPb(st *EnableResultsDownloading) (*enableResultsDownloadingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enableResultsDownloadingPb{}
	booleanValPb, err := booleanMessageToPb(st.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValPb != nil {
		pb.BooleanVal = booleanValPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enableResultsDownloadingPb struct {
	BooleanVal *booleanMessagePb `json:"boolean_val,omitempty"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enableResultsDownloadingFromPb(pb *enableResultsDownloadingPb) (*EnableResultsDownloading, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnableResultsDownloading{}
	booleanValField, err := booleanMessageFromPb(pb.BooleanVal)
	if err != nil {
		return nil, err
	}
	if booleanValField != nil {
		st.BooleanVal = booleanValField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enableResultsDownloadingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enableResultsDownloadingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// SHIELD feature: ESM
type EnhancedSecurityMonitoring struct {
	IsEnabled bool

	ForceSendFields []string
}

func enhancedSecurityMonitoringToPb(st *EnhancedSecurityMonitoring) (*enhancedSecurityMonitoringPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enhancedSecurityMonitoringPb{}
	isEnabledPb, err := identity(&st.IsEnabled)
	if err != nil {
		return nil, err
	}
	if isEnabledPb != nil {
		pb.IsEnabled = *isEnabledPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enhancedSecurityMonitoringPb struct {
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func enhancedSecurityMonitoringFromPb(pb *enhancedSecurityMonitoringPb) (*EnhancedSecurityMonitoring, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnhancedSecurityMonitoring{}
	isEnabledField, err := identity(&pb.IsEnabled)
	if err != nil {
		return nil, err
	}
	if isEnabledField != nil {
		st.IsEnabled = *isEnabledField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enhancedSecurityMonitoringPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enhancedSecurityMonitoringPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EnhancedSecurityMonitoringSetting struct {
	// SHIELD feature: ESM
	EnhancedSecurityMonitoringWorkspace EnhancedSecurityMonitoring
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func enhancedSecurityMonitoringSettingToPb(st *EnhancedSecurityMonitoringSetting) (*enhancedSecurityMonitoringSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &enhancedSecurityMonitoringSettingPb{}
	enhancedSecurityMonitoringWorkspacePb, err := enhancedSecurityMonitoringToPb(&st.EnhancedSecurityMonitoringWorkspace)
	if err != nil {
		return nil, err
	}
	if enhancedSecurityMonitoringWorkspacePb != nil {
		pb.EnhancedSecurityMonitoringWorkspace = *enhancedSecurityMonitoringWorkspacePb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type enhancedSecurityMonitoringSettingPb struct {
	// SHIELD feature: ESM
	EnhancedSecurityMonitoringWorkspace enhancedSecurityMonitoringPb `json:"enhanced_security_monitoring_workspace"`
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

func enhancedSecurityMonitoringSettingFromPb(pb *enhancedSecurityMonitoringSettingPb) (*EnhancedSecurityMonitoringSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EnhancedSecurityMonitoringSetting{}
	enhancedSecurityMonitoringWorkspaceField, err := enhancedSecurityMonitoringFromPb(&pb.EnhancedSecurityMonitoringWorkspace)
	if err != nil {
		return nil, err
	}
	if enhancedSecurityMonitoringWorkspaceField != nil {
		st.EnhancedSecurityMonitoringWorkspace = *enhancedSecurityMonitoringWorkspaceField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *enhancedSecurityMonitoringSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st enhancedSecurityMonitoringSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Account level policy for ESM
type EsmEnablementAccount struct {
	IsEnforced bool

	ForceSendFields []string
}

func esmEnablementAccountToPb(st *EsmEnablementAccount) (*esmEnablementAccountPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &esmEnablementAccountPb{}
	isEnforcedPb, err := identity(&st.IsEnforced)
	if err != nil {
		return nil, err
	}
	if isEnforcedPb != nil {
		pb.IsEnforced = *isEnforcedPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type esmEnablementAccountPb struct {
	IsEnforced bool `json:"is_enforced,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func esmEnablementAccountFromPb(pb *esmEnablementAccountPb) (*EsmEnablementAccount, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EsmEnablementAccount{}
	isEnforcedField, err := identity(&pb.IsEnforced)
	if err != nil {
		return nil, err
	}
	if isEnforcedField != nil {
		st.IsEnforced = *isEnforcedField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *esmEnablementAccountPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st esmEnablementAccountPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EsmEnablementAccountSetting struct {
	// Account level policy for ESM
	EsmEnablementAccount EsmEnablementAccount
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func esmEnablementAccountSettingToPb(st *EsmEnablementAccountSetting) (*esmEnablementAccountSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &esmEnablementAccountSettingPb{}
	esmEnablementAccountPb, err := esmEnablementAccountToPb(&st.EsmEnablementAccount)
	if err != nil {
		return nil, err
	}
	if esmEnablementAccountPb != nil {
		pb.EsmEnablementAccount = *esmEnablementAccountPb
	}

	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type esmEnablementAccountSettingPb struct {
	// Account level policy for ESM
	EsmEnablementAccount esmEnablementAccountPb `json:"esm_enablement_account"`
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

func esmEnablementAccountSettingFromPb(pb *esmEnablementAccountSettingPb) (*EsmEnablementAccountSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EsmEnablementAccountSetting{}
	esmEnablementAccountField, err := esmEnablementAccountFromPb(&pb.EsmEnablementAccount)
	if err != nil {
		return nil, err
	}
	if esmEnablementAccountField != nil {
		st.EsmEnablementAccount = *esmEnablementAccountField
	}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *esmEnablementAccountSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st esmEnablementAccountSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The exchange token is the result of the token exchange with the IdP
type ExchangeToken struct {
	// The requested token.
	Credential string
	// The end-of-life timestamp of the token. The value is in milliseconds
	// since the Unix epoch.
	CredentialEolTime int64
	// User ID of the user that owns this token.
	OwnerId int64
	// The scopes of access granted in the token.
	Scopes []string
	// The type of this exchange token
	TokenType TokenType

	ForceSendFields []string
}

func exchangeTokenToPb(st *ExchangeToken) (*exchangeTokenPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeTokenPb{}
	credentialPb, err := identity(&st.Credential)
	if err != nil {
		return nil, err
	}
	if credentialPb != nil {
		pb.Credential = *credentialPb
	}

	credentialEolTimePb, err := identity(&st.CredentialEolTime)
	if err != nil {
		return nil, err
	}
	if credentialEolTimePb != nil {
		pb.CredentialEolTime = *credentialEolTimePb
	}

	ownerIdPb, err := identity(&st.OwnerId)
	if err != nil {
		return nil, err
	}
	if ownerIdPb != nil {
		pb.OwnerId = *ownerIdPb
	}

	var scopesPb []string
	for _, item := range st.Scopes {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

	tokenTypePb, err := identity(&st.TokenType)
	if err != nil {
		return nil, err
	}
	if tokenTypePb != nil {
		pb.TokenType = *tokenTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type exchangeTokenPb struct {
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

func exchangeTokenFromPb(pb *exchangeTokenPb) (*ExchangeToken, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeToken{}
	credentialField, err := identity(&pb.Credential)
	if err != nil {
		return nil, err
	}
	if credentialField != nil {
		st.Credential = *credentialField
	}
	credentialEolTimeField, err := identity(&pb.CredentialEolTime)
	if err != nil {
		return nil, err
	}
	if credentialEolTimeField != nil {
		st.CredentialEolTime = *credentialEolTimeField
	}
	ownerIdField, err := identity(&pb.OwnerId)
	if err != nil {
		return nil, err
	}
	if ownerIdField != nil {
		st.OwnerId = *ownerIdField
	}

	var scopesField []string
	for _, itemPb := range pb.Scopes {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			scopesField = append(scopesField, *item)
		}
	}
	st.Scopes = scopesField
	tokenTypeField, err := identity(&pb.TokenType)
	if err != nil {
		return nil, err
	}
	if tokenTypeField != nil {
		st.TokenType = *tokenTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *exchangeTokenPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st exchangeTokenPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Exchange a token with the IdP
type ExchangeTokenRequest struct {
	// The partition of Credentials store
	PartitionId PartitionId
	// Array of scopes for the token request.
	Scopes []string
	// A list of token types being requested
	TokenType []TokenType
}

func exchangeTokenRequestToPb(st *ExchangeTokenRequest) (*exchangeTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeTokenRequestPb{}
	partitionIdPb, err := partitionIdToPb(&st.PartitionId)
	if err != nil {
		return nil, err
	}
	if partitionIdPb != nil {
		pb.PartitionId = *partitionIdPb
	}

	var scopesPb []string
	for _, item := range st.Scopes {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

	var tokenTypePb []TokenType
	for _, item := range st.TokenType {
		itemPb, err := identity(&item)
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

type exchangeTokenRequestPb struct {
	// The partition of Credentials store
	PartitionId partitionIdPb `json:"partitionId"`
	// Array of scopes for the token request.
	Scopes []string `json:"scopes"`
	// A list of token types being requested
	TokenType []TokenType `json:"tokenType"`
}

func exchangeTokenRequestFromPb(pb *exchangeTokenRequestPb) (*ExchangeTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeTokenRequest{}
	partitionIdField, err := partitionIdFromPb(&pb.PartitionId)
	if err != nil {
		return nil, err
	}
	if partitionIdField != nil {
		st.PartitionId = *partitionIdField
	}

	var scopesField []string
	for _, itemPb := range pb.Scopes {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			scopesField = append(scopesField, *item)
		}
	}
	st.Scopes = scopesField

	var tokenTypeField []TokenType
	for _, itemPb := range pb.TokenType {
		item, err := identity(&itemPb)
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
	Values []ExchangeToken
}

func exchangeTokenResponseToPb(st *ExchangeTokenResponse) (*exchangeTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exchangeTokenResponsePb{}

	var valuesPb []exchangeTokenPb
	for _, item := range st.Values {
		itemPb, err := exchangeTokenToPb(&item)
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

type exchangeTokenResponsePb struct {
	Values []exchangeTokenPb `json:"values,omitempty"`
}

func exchangeTokenResponseFromPb(pb *exchangeTokenResponsePb) (*ExchangeTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExchangeTokenResponse{}

	var valuesField []ExchangeToken
	for _, itemPb := range pb.Values {
		item, err := exchangeTokenFromPb(&itemPb)
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
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo
}

func fetchIpAccessListResponseToPb(st *FetchIpAccessListResponse) (*fetchIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fetchIpAccessListResponsePb{}
	ipAccessListPb, err := ipAccessListInfoToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	return pb, nil
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

type fetchIpAccessListResponsePb struct {
	// Definition of an IP Access list
	IpAccessList *ipAccessListInfoPb `json:"ip_access_list,omitempty"`
}

func fetchIpAccessListResponseFromPb(pb *fetchIpAccessListResponsePb) (*FetchIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FetchIpAccessListResponse{}
	ipAccessListField, err := ipAccessListInfoFromPb(pb.IpAccessList)
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
	Password string
	// [Output-Only] Whether password is set.
	PasswordSet bool
	// [Input-Only] URL for webhook.
	Url string
	// [Output-Only] Whether URL is set.
	UrlSet bool
	// [Input-Only][Optional] Username for webhook.
	Username string
	// [Output-Only] Whether username is set.
	UsernameSet bool

	ForceSendFields []string
}

func genericWebhookConfigToPb(st *GenericWebhookConfig) (*genericWebhookConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &genericWebhookConfigPb{}
	passwordPb, err := identity(&st.Password)
	if err != nil {
		return nil, err
	}
	if passwordPb != nil {
		pb.Password = *passwordPb
	}

	passwordSetPb, err := identity(&st.PasswordSet)
	if err != nil {
		return nil, err
	}
	if passwordSetPb != nil {
		pb.PasswordSet = *passwordSetPb
	}

	urlPb, err := identity(&st.Url)
	if err != nil {
		return nil, err
	}
	if urlPb != nil {
		pb.Url = *urlPb
	}

	urlSetPb, err := identity(&st.UrlSet)
	if err != nil {
		return nil, err
	}
	if urlSetPb != nil {
		pb.UrlSet = *urlSetPb
	}

	usernamePb, err := identity(&st.Username)
	if err != nil {
		return nil, err
	}
	if usernamePb != nil {
		pb.Username = *usernamePb
	}

	usernameSetPb, err := identity(&st.UsernameSet)
	if err != nil {
		return nil, err
	}
	if usernameSetPb != nil {
		pb.UsernameSet = *usernameSetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type genericWebhookConfigPb struct {
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

func genericWebhookConfigFromPb(pb *genericWebhookConfigPb) (*GenericWebhookConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GenericWebhookConfig{}
	passwordField, err := identity(&pb.Password)
	if err != nil {
		return nil, err
	}
	if passwordField != nil {
		st.Password = *passwordField
	}
	passwordSetField, err := identity(&pb.PasswordSet)
	if err != nil {
		return nil, err
	}
	if passwordSetField != nil {
		st.PasswordSet = *passwordSetField
	}
	urlField, err := identity(&pb.Url)
	if err != nil {
		return nil, err
	}
	if urlField != nil {
		st.Url = *urlField
	}
	urlSetField, err := identity(&pb.UrlSet)
	if err != nil {
		return nil, err
	}
	if urlSetField != nil {
		st.UrlSet = *urlSetField
	}
	usernameField, err := identity(&pb.Username)
	if err != nil {
		return nil, err
	}
	if usernameField != nil {
		st.Username = *usernameField
	}
	usernameSetField, err := identity(&pb.UsernameSet)
	if err != nil {
		return nil, err
	}
	if usernameSetField != nil {
		st.UsernameSet = *usernameSetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *genericWebhookConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st genericWebhookConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getAccountIpAccessEnableRequestToPb(st *GetAccountIpAccessEnableRequest) (*getAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountIpAccessEnableRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getAccountIpAccessEnableRequestPb struct {
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

func getAccountIpAccessEnableRequestFromPb(pb *getAccountIpAccessEnableRequestPb) (*GetAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountIpAccessEnableRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAccountIpAccessEnableRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAccountIpAccessEnableRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get IP access list
type GetAccountIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string
}

func getAccountIpAccessListRequestToPb(st *GetAccountIpAccessListRequest) (*getAccountIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountIpAccessListRequestPb{}
	ipAccessListIdPb, err := identity(&st.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdPb != nil {
		pb.IpAccessListId = *ipAccessListIdPb
	}

	return pb, nil
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

type getAccountIpAccessListRequestPb struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func getAccountIpAccessListRequestFromPb(pb *getAccountIpAccessListRequestPb) (*GetAccountIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountIpAccessListRequest{}
	ipAccessListIdField, err := identity(&pb.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdField != nil {
		st.IpAccessListId = *ipAccessListIdField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func getAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *GetAibiDashboardEmbeddingAccessPolicySettingRequest) (*getAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
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

func getAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *getAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*GetAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAibiDashboardEmbeddingAccessPolicySettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAibiDashboardEmbeddingAccessPolicySettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAibiDashboardEmbeddingAccessPolicySettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *GetAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
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

func getAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*GetAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getAutomaticClusterUpdateSettingRequestToPb(st *GetAutomaticClusterUpdateSettingRequest) (*getAutomaticClusterUpdateSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAutomaticClusterUpdateSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getAutomaticClusterUpdateSettingRequestPb struct {
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

func getAutomaticClusterUpdateSettingRequestFromPb(pb *getAutomaticClusterUpdateSettingRequestPb) (*GetAutomaticClusterUpdateSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAutomaticClusterUpdateSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getAutomaticClusterUpdateSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getAutomaticClusterUpdateSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getComplianceSecurityProfileSettingRequestToPb(st *GetComplianceSecurityProfileSettingRequest) (*getComplianceSecurityProfileSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getComplianceSecurityProfileSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getComplianceSecurityProfileSettingRequestPb struct {
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

func getComplianceSecurityProfileSettingRequestFromPb(pb *getComplianceSecurityProfileSettingRequestPb) (*GetComplianceSecurityProfileSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetComplianceSecurityProfileSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getComplianceSecurityProfileSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getComplianceSecurityProfileSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getCspEnablementAccountSettingRequestToPb(st *GetCspEnablementAccountSettingRequest) (*getCspEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCspEnablementAccountSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getCspEnablementAccountSettingRequestPb struct {
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

func getCspEnablementAccountSettingRequestFromPb(pb *getCspEnablementAccountSettingRequestPb) (*GetCspEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCspEnablementAccountSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCspEnablementAccountSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCspEnablementAccountSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getDefaultNamespaceSettingRequestToPb(st *GetDefaultNamespaceSettingRequest) (*getDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDefaultNamespaceSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getDefaultNamespaceSettingRequestPb struct {
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

func getDefaultNamespaceSettingRequestFromPb(pb *getDefaultNamespaceSettingRequestPb) (*GetDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDefaultNamespaceSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDefaultNamespaceSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDefaultNamespaceSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getDisableLegacyAccessRequestToPb(st *GetDisableLegacyAccessRequest) (*getDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDisableLegacyAccessRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getDisableLegacyAccessRequestPb struct {
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

func getDisableLegacyAccessRequestFromPb(pb *getDisableLegacyAccessRequestPb) (*GetDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyAccessRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDisableLegacyAccessRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDisableLegacyAccessRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getDisableLegacyDbfsRequestToPb(st *GetDisableLegacyDbfsRequest) (*getDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDisableLegacyDbfsRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getDisableLegacyDbfsRequestPb struct {
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

func getDisableLegacyDbfsRequestFromPb(pb *getDisableLegacyDbfsRequestPb) (*GetDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyDbfsRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDisableLegacyDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDisableLegacyDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getDisableLegacyFeaturesRequestToPb(st *GetDisableLegacyFeaturesRequest) (*getDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDisableLegacyFeaturesRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getDisableLegacyFeaturesRequestPb struct {
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

func getDisableLegacyFeaturesRequestFromPb(pb *getDisableLegacyFeaturesRequestPb) (*GetDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDisableLegacyFeaturesRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getDisableLegacyFeaturesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getDisableLegacyFeaturesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getEnhancedSecurityMonitoringSettingRequestToPb(st *GetEnhancedSecurityMonitoringSettingRequest) (*getEnhancedSecurityMonitoringSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEnhancedSecurityMonitoringSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getEnhancedSecurityMonitoringSettingRequestPb struct {
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

func getEnhancedSecurityMonitoringSettingRequestFromPb(pb *getEnhancedSecurityMonitoringSettingRequestPb) (*GetEnhancedSecurityMonitoringSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEnhancedSecurityMonitoringSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEnhancedSecurityMonitoringSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEnhancedSecurityMonitoringSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Etag string

	ForceSendFields []string
}

func getEsmEnablementAccountSettingRequestToPb(st *GetEsmEnablementAccountSettingRequest) (*getEsmEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getEsmEnablementAccountSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getEsmEnablementAccountSettingRequestPb struct {
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

func getEsmEnablementAccountSettingRequestFromPb(pb *getEsmEnablementAccountSettingRequestPb) (*GetEsmEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetEsmEnablementAccountSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getEsmEnablementAccountSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getEsmEnablementAccountSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get access list
type GetIpAccessListRequest struct {
	// The ID for the corresponding IP access list
	IpAccessListId string
}

func getIpAccessListRequestToPb(st *GetIpAccessListRequest) (*getIpAccessListRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIpAccessListRequestPb{}
	ipAccessListIdPb, err := identity(&st.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdPb != nil {
		pb.IpAccessListId = *ipAccessListIdPb
	}

	return pb, nil
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

type getIpAccessListRequestPb struct {
	// The ID for the corresponding IP access list
	IpAccessListId string `json:"-" url:"-"`
}

func getIpAccessListRequestFromPb(pb *getIpAccessListRequestPb) (*GetIpAccessListRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListRequest{}
	ipAccessListIdField, err := identity(&pb.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdField != nil {
		st.IpAccessListId = *ipAccessListIdField
	}

	return st, nil
}

type GetIpAccessListResponse struct {
	// Definition of an IP Access list
	IpAccessList *IpAccessListInfo
}

func getIpAccessListResponseToPb(st *GetIpAccessListResponse) (*getIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIpAccessListResponsePb{}
	ipAccessListPb, err := ipAccessListInfoToPb(st.IpAccessList)
	if err != nil {
		return nil, err
	}
	if ipAccessListPb != nil {
		pb.IpAccessList = ipAccessListPb
	}

	return pb, nil
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

type getIpAccessListResponsePb struct {
	// Definition of an IP Access list
	IpAccessList *ipAccessListInfoPb `json:"ip_access_list,omitempty"`
}

func getIpAccessListResponseFromPb(pb *getIpAccessListResponsePb) (*GetIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListResponse{}
	ipAccessListField, err := ipAccessListInfoFromPb(pb.IpAccessList)
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
	IpAccessLists []IpAccessListInfo
}

func getIpAccessListsResponseToPb(st *GetIpAccessListsResponse) (*getIpAccessListsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getIpAccessListsResponsePb{}

	var ipAccessListsPb []ipAccessListInfoPb
	for _, item := range st.IpAccessLists {
		itemPb, err := ipAccessListInfoToPb(&item)
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

type getIpAccessListsResponsePb struct {
	IpAccessLists []ipAccessListInfoPb `json:"ip_access_lists,omitempty"`
}

func getIpAccessListsResponseFromPb(pb *getIpAccessListsResponsePb) (*GetIpAccessListsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetIpAccessListsResponse{}

	var ipAccessListsField []IpAccessListInfo
	for _, itemPb := range pb.IpAccessLists {
		item, err := ipAccessListInfoFromPb(&itemPb)
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

// Get a network connectivity configuration
type GetNetworkConnectivityConfigurationRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string
}

func getNetworkConnectivityConfigurationRequestToPb(st *GetNetworkConnectivityConfigurationRequest) (*getNetworkConnectivityConfigurationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNetworkConnectivityConfigurationRequestPb{}
	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	return pb, nil
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

type getNetworkConnectivityConfigurationRequestPb struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
}

func getNetworkConnectivityConfigurationRequestFromPb(pb *getNetworkConnectivityConfigurationRequestPb) (*GetNetworkConnectivityConfigurationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNetworkConnectivityConfigurationRequest{}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}

	return st, nil
}

// Get a notification destination
type GetNotificationDestinationRequest struct {
	Id string
}

func getNotificationDestinationRequestToPb(st *GetNotificationDestinationRequest) (*getNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getNotificationDestinationRequestPb{}
	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	return pb, nil
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

type getNotificationDestinationRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getNotificationDestinationRequestFromPb(pb *getNotificationDestinationRequestPb) (*GetNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetNotificationDestinationRequest{}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func getPersonalComputeSettingRequestToPb(st *GetPersonalComputeSettingRequest) (*getPersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPersonalComputeSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getPersonalComputeSettingRequestPb struct {
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

func getPersonalComputeSettingRequestFromPb(pb *getPersonalComputeSettingRequestPb) (*GetPersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPersonalComputeSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPersonalComputeSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPersonalComputeSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Gets a private endpoint rule
type GetPrivateEndpointRuleRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string
}

func getPrivateEndpointRuleRequestToPb(st *GetPrivateEndpointRuleRequest) (*getPrivateEndpointRuleRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPrivateEndpointRuleRequestPb{}
	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	privateEndpointRuleIdPb, err := identity(&st.PrivateEndpointRuleId)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleIdPb != nil {
		pb.PrivateEndpointRuleId = *privateEndpointRuleIdPb
	}

	return pb, nil
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

type getPrivateEndpointRuleRequestPb struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
}

func getPrivateEndpointRuleRequestFromPb(pb *getPrivateEndpointRuleRequestPb) (*GetPrivateEndpointRuleRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPrivateEndpointRuleRequest{}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	privateEndpointRuleIdField, err := identity(&pb.PrivateEndpointRuleId)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleIdField != nil {
		st.PrivateEndpointRuleId = *privateEndpointRuleIdField
	}

	return st, nil
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
	Etag string

	ForceSendFields []string
}

func getRestrictWorkspaceAdminsSettingRequestToPb(st *GetRestrictWorkspaceAdminsSettingRequest) (*getRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRestrictWorkspaceAdminsSettingRequestPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type getRestrictWorkspaceAdminsSettingRequestPb struct {
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

func getRestrictWorkspaceAdminsSettingRequestFromPb(pb *getRestrictWorkspaceAdminsSettingRequestPb) (*GetRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRestrictWorkspaceAdminsSettingRequest{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRestrictWorkspaceAdminsSettingRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRestrictWorkspaceAdminsSettingRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Check configuration status
type GetStatusRequest struct {
	Keys string
}

func getStatusRequestToPb(st *GetStatusRequest) (*getStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatusRequestPb{}
	keysPb, err := identity(&st.Keys)
	if err != nil {
		return nil, err
	}
	if keysPb != nil {
		pb.Keys = *keysPb
	}

	return pb, nil
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

type getStatusRequestPb struct {
	Keys string `json:"-" url:"keys"`
}

func getStatusRequestFromPb(pb *getStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	keysField, err := identity(&pb.Keys)
	if err != nil {
		return nil, err
	}
	if keysField != nil {
		st.Keys = *keysField
	}

	return st, nil
}

// Get token info
type GetTokenManagementRequest struct {
	// The ID of the token to get.
	TokenId string
}

func getTokenManagementRequestToPb(st *GetTokenManagementRequest) (*getTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTokenManagementRequestPb{}
	tokenIdPb, err := identity(&st.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdPb != nil {
		pb.TokenId = *tokenIdPb
	}

	return pb, nil
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

type getTokenManagementRequestPb struct {
	// The ID of the token to get.
	TokenId string `json:"-" url:"-"`
}

func getTokenManagementRequestFromPb(pb *getTokenManagementRequestPb) (*GetTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenManagementRequest{}
	tokenIdField, err := identity(&pb.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdField != nil {
		st.TokenId = *tokenIdField
	}

	return st, nil
}

type GetTokenPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []TokenPermissionsDescription
}

func getTokenPermissionLevelsResponseToPb(st *GetTokenPermissionLevelsResponse) (*getTokenPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTokenPermissionLevelsResponsePb{}

	var permissionLevelsPb []tokenPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := tokenPermissionsDescriptionToPb(&item)
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

type getTokenPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []tokenPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getTokenPermissionLevelsResponseFromPb(pb *getTokenPermissionLevelsResponsePb) (*GetTokenPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenPermissionLevelsResponse{}

	var permissionLevelsField []TokenPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := tokenPermissionsDescriptionFromPb(&itemPb)
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
	TokenInfo *TokenInfo
}

func getTokenResponseToPb(st *GetTokenResponse) (*getTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getTokenResponsePb{}
	tokenInfoPb, err := tokenInfoToPb(st.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoPb != nil {
		pb.TokenInfo = tokenInfoPb
	}

	return pb, nil
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

type getTokenResponsePb struct {
	TokenInfo *tokenInfoPb `json:"token_info,omitempty"`
}

func getTokenResponseFromPb(pb *getTokenResponsePb) (*GetTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetTokenResponse{}
	tokenInfoField, err := tokenInfoFromPb(pb.TokenInfo)
	if err != nil {
		return nil, err
	}
	if tokenInfoField != nil {
		st.TokenInfo = tokenInfoField
	}

	return st, nil
}

// Definition of an IP Access list
type IpAccessListInfo struct {
	// Total number of IP or CIDR values.
	AddressCount int
	// Creation timestamp in milliseconds.
	CreatedAt int64
	// User ID of the user who created this list.
	CreatedBy int64
	// Specifies whether this IP access list is enabled.
	Enabled bool

	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	Label string
	// Universally unique identifier (UUID) of the IP access list.
	ListId string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType
	// Update timestamp in milliseconds.
	UpdatedAt int64
	// User ID of the user who updated this list.
	UpdatedBy int64

	ForceSendFields []string
}

func ipAccessListInfoToPb(st *IpAccessListInfo) (*ipAccessListInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ipAccessListInfoPb{}
	addressCountPb, err := identity(&st.AddressCount)
	if err != nil {
		return nil, err
	}
	if addressCountPb != nil {
		pb.AddressCount = *addressCountPb
	}

	createdAtPb, err := identity(&st.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtPb != nil {
		pb.CreatedAt = *createdAtPb
	}

	createdByPb, err := identity(&st.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	var ipAddressesPb []string
	for _, item := range st.IpAddresses {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			ipAddressesPb = append(ipAddressesPb, *itemPb)
		}
	}
	pb.IpAddresses = ipAddressesPb

	labelPb, err := identity(&st.Label)
	if err != nil {
		return nil, err
	}
	if labelPb != nil {
		pb.Label = *labelPb
	}

	listIdPb, err := identity(&st.ListId)
	if err != nil {
		return nil, err
	}
	if listIdPb != nil {
		pb.ListId = *listIdPb
	}

	listTypePb, err := identity(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}

	updatedAtPb, err := identity(&st.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtPb != nil {
		pb.UpdatedAt = *updatedAtPb
	}

	updatedByPb, err := identity(&st.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByPb != nil {
		pb.UpdatedBy = *updatedByPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type ipAccessListInfoPb struct {
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

func ipAccessListInfoFromPb(pb *ipAccessListInfoPb) (*IpAccessListInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IpAccessListInfo{}
	addressCountField, err := identity(&pb.AddressCount)
	if err != nil {
		return nil, err
	}
	if addressCountField != nil {
		st.AddressCount = *addressCountField
	}
	createdAtField, err := identity(&pb.CreatedAt)
	if err != nil {
		return nil, err
	}
	if createdAtField != nil {
		st.CreatedAt = *createdAtField
	}
	createdByField, err := identity(&pb.CreatedBy)
	if err != nil {
		return nil, err
	}
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}

	var ipAddressesField []string
	for _, itemPb := range pb.IpAddresses {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			ipAddressesField = append(ipAddressesField, *item)
		}
	}
	st.IpAddresses = ipAddressesField
	labelField, err := identity(&pb.Label)
	if err != nil {
		return nil, err
	}
	if labelField != nil {
		st.Label = *labelField
	}
	listIdField, err := identity(&pb.ListId)
	if err != nil {
		return nil, err
	}
	if listIdField != nil {
		st.ListId = *listIdField
	}
	listTypeField, err := identity(&pb.ListType)
	if err != nil {
		return nil, err
	}
	if listTypeField != nil {
		st.ListType = *listTypeField
	}
	updatedAtField, err := identity(&pb.UpdatedAt)
	if err != nil {
		return nil, err
	}
	if updatedAtField != nil {
		st.UpdatedAt = *updatedAtField
	}
	updatedByField, err := identity(&pb.UpdatedBy)
	if err != nil {
		return nil, err
	}
	if updatedByField != nil {
		st.UpdatedBy = *updatedByField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ipAccessListInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ipAccessListInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// IP access lists were successfully returned.
type ListIpAccessListResponse struct {
	IpAccessLists []IpAccessListInfo
}

func listIpAccessListResponseToPb(st *ListIpAccessListResponse) (*listIpAccessListResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listIpAccessListResponsePb{}

	var ipAccessListsPb []ipAccessListInfoPb
	for _, item := range st.IpAccessLists {
		itemPb, err := ipAccessListInfoToPb(&item)
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

type listIpAccessListResponsePb struct {
	IpAccessLists []ipAccessListInfoPb `json:"ip_access_lists,omitempty"`
}

func listIpAccessListResponseFromPb(pb *listIpAccessListResponsePb) (*ListIpAccessListResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListIpAccessListResponse{}

	var ipAccessListsField []IpAccessListInfo
	for _, itemPb := range pb.IpAccessLists {
		item, err := ipAccessListInfoFromPb(&itemPb)
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

// The private endpoint rule list was successfully retrieved.
type ListNccAzurePrivateEndpointRulesResponse struct {
	Items []NccAzurePrivateEndpointRule
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string

	ForceSendFields []string
}

func listNccAzurePrivateEndpointRulesResponseToPb(st *ListNccAzurePrivateEndpointRulesResponse) (*listNccAzurePrivateEndpointRulesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNccAzurePrivateEndpointRulesResponsePb{}

	var itemsPb []nccAzurePrivateEndpointRulePb
	for _, item := range st.Items {
		itemPb, err := nccAzurePrivateEndpointRuleToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			itemsPb = append(itemsPb, *itemPb)
		}
	}
	pb.Items = itemsPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listNccAzurePrivateEndpointRulesResponsePb struct {
	Items []nccAzurePrivateEndpointRulePb `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNccAzurePrivateEndpointRulesResponseFromPb(pb *listNccAzurePrivateEndpointRulesResponsePb) (*ListNccAzurePrivateEndpointRulesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNccAzurePrivateEndpointRulesResponse{}

	var itemsField []NccAzurePrivateEndpointRule
	for _, itemPb := range pb.Items {
		item, err := nccAzurePrivateEndpointRuleFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			itemsField = append(itemsField, *item)
		}
	}
	st.Items = itemsField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNccAzurePrivateEndpointRulesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNccAzurePrivateEndpointRulesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List network connectivity configurations
type ListNetworkConnectivityConfigurationsRequest struct {
	// Pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listNetworkConnectivityConfigurationsRequestToPb(st *ListNetworkConnectivityConfigurationsRequest) (*listNetworkConnectivityConfigurationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNetworkConnectivityConfigurationsRequestPb{}
	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listNetworkConnectivityConfigurationsRequestPb struct {
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNetworkConnectivityConfigurationsRequestFromPb(pb *listNetworkConnectivityConfigurationsRequestPb) (*ListNetworkConnectivityConfigurationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkConnectivityConfigurationsRequest{}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNetworkConnectivityConfigurationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNetworkConnectivityConfigurationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The network connectivity configuration list was successfully retrieved.
type ListNetworkConnectivityConfigurationsResponse struct {
	Items []NetworkConnectivityConfiguration
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string

	ForceSendFields []string
}

func listNetworkConnectivityConfigurationsResponseToPb(st *ListNetworkConnectivityConfigurationsResponse) (*listNetworkConnectivityConfigurationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNetworkConnectivityConfigurationsResponsePb{}

	var itemsPb []networkConnectivityConfigurationPb
	for _, item := range st.Items {
		itemPb, err := networkConnectivityConfigurationToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			itemsPb = append(itemsPb, *itemPb)
		}
	}
	pb.Items = itemsPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listNetworkConnectivityConfigurationsResponsePb struct {
	Items []networkConnectivityConfigurationPb `json:"items,omitempty"`
	// A token that can be used to get the next page of results. If null, there
	// are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNetworkConnectivityConfigurationsResponseFromPb(pb *listNetworkConnectivityConfigurationsResponsePb) (*ListNetworkConnectivityConfigurationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNetworkConnectivityConfigurationsResponse{}

	var itemsField []NetworkConnectivityConfiguration
	for _, itemPb := range pb.Items {
		item, err := networkConnectivityConfigurationFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			itemsField = append(itemsField, *item)
		}
	}
	st.Items = itemsField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNetworkConnectivityConfigurationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNetworkConnectivityConfigurationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List notification destinations
type ListNotificationDestinationsRequest struct {
	PageSize int64

	PageToken string

	ForceSendFields []string
}

func listNotificationDestinationsRequestToPb(st *ListNotificationDestinationsRequest) (*listNotificationDestinationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNotificationDestinationsRequestPb{}
	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listNotificationDestinationsRequestPb struct {
	PageSize int64 `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNotificationDestinationsRequestFromPb(pb *listNotificationDestinationsRequestPb) (*ListNotificationDestinationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsRequest{}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNotificationDestinationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNotificationDestinationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNotificationDestinationsResponse struct {
	// Page token for next of results.
	NextPageToken string

	Results []ListNotificationDestinationsResult

	ForceSendFields []string
}

func listNotificationDestinationsResponseToPb(st *ListNotificationDestinationsResponse) (*listNotificationDestinationsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNotificationDestinationsResponsePb{}
	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	var resultsPb []listNotificationDestinationsResultPb
	for _, item := range st.Results {
		itemPb, err := listNotificationDestinationsResultToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			resultsPb = append(resultsPb, *itemPb)
		}
	}
	pb.Results = resultsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listNotificationDestinationsResponsePb struct {
	// Page token for next of results.
	NextPageToken string `json:"next_page_token,omitempty"`

	Results []listNotificationDestinationsResultPb `json:"results,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNotificationDestinationsResponseFromPb(pb *listNotificationDestinationsResponsePb) (*ListNotificationDestinationsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsResponse{}
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var resultsField []ListNotificationDestinationsResult
	for _, itemPb := range pb.Results {
		item, err := listNotificationDestinationsResultFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			resultsField = append(resultsField, *item)
		}
	}
	st.Results = resultsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNotificationDestinationsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNotificationDestinationsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListNotificationDestinationsResult struct {
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType DestinationType
	// The display name for the notification destination.
	DisplayName string
	// UUID identifying notification destination.
	Id string

	ForceSendFields []string
}

func listNotificationDestinationsResultToPb(st *ListNotificationDestinationsResult) (*listNotificationDestinationsResultPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listNotificationDestinationsResultPb{}
	destinationTypePb, err := identity(&st.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypePb != nil {
		pb.DestinationType = *destinationTypePb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listNotificationDestinationsResultPb struct {
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listNotificationDestinationsResultFromPb(pb *listNotificationDestinationsResultPb) (*ListNotificationDestinationsResult, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListNotificationDestinationsResult{}
	destinationTypeField, err := identity(&pb.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypeField != nil {
		st.DestinationType = *destinationTypeField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listNotificationDestinationsResultPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listNotificationDestinationsResultPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List private endpoint rules
type ListPrivateEndpointRulesRequest struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string
	// Pagination token to go to next page based on previous query.
	PageToken string

	ForceSendFields []string
}

func listPrivateEndpointRulesRequestToPb(st *ListPrivateEndpointRulesRequest) (*listPrivateEndpointRulesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPrivateEndpointRulesRequestPb{}
	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listPrivateEndpointRulesRequestPb struct {
	// Your Network Connectvity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Pagination token to go to next page based on previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPrivateEndpointRulesRequestFromPb(pb *listPrivateEndpointRulesRequestPb) (*ListPrivateEndpointRulesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPrivateEndpointRulesRequest{}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPrivateEndpointRulesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPrivateEndpointRulesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPublicTokensResponse struct {
	// The information for each token.
	TokenInfos []PublicTokenInfo
}

func listPublicTokensResponseToPb(st *ListPublicTokensResponse) (*listPublicTokensResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPublicTokensResponsePb{}

	var tokenInfosPb []publicTokenInfoPb
	for _, item := range st.TokenInfos {
		itemPb, err := publicTokenInfoToPb(&item)
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

type listPublicTokensResponsePb struct {
	// The information for each token.
	TokenInfos []publicTokenInfoPb `json:"token_infos,omitempty"`
}

func listPublicTokensResponseFromPb(pb *listPublicTokensResponsePb) (*ListPublicTokensResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPublicTokensResponse{}

	var tokenInfosField []PublicTokenInfo
	for _, itemPb := range pb.TokenInfos {
		item, err := publicTokenInfoFromPb(&itemPb)
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

// List all tokens
type ListTokenManagementRequest struct {
	// User ID of the user that created the token.
	CreatedById int64
	// Username of the user that created the token.
	CreatedByUsername string

	ForceSendFields []string
}

func listTokenManagementRequestToPb(st *ListTokenManagementRequest) (*listTokenManagementRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTokenManagementRequestPb{}
	createdByIdPb, err := identity(&st.CreatedById)
	if err != nil {
		return nil, err
	}
	if createdByIdPb != nil {
		pb.CreatedById = *createdByIdPb
	}

	createdByUsernamePb, err := identity(&st.CreatedByUsername)
	if err != nil {
		return nil, err
	}
	if createdByUsernamePb != nil {
		pb.CreatedByUsername = *createdByUsernamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type listTokenManagementRequestPb struct {
	// User ID of the user that created the token.
	CreatedById int64 `json:"-" url:"created_by_id,omitempty"`
	// Username of the user that created the token.
	CreatedByUsername string `json:"-" url:"created_by_username,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listTokenManagementRequestFromPb(pb *listTokenManagementRequestPb) (*ListTokenManagementRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTokenManagementRequest{}
	createdByIdField, err := identity(&pb.CreatedById)
	if err != nil {
		return nil, err
	}
	if createdByIdField != nil {
		st.CreatedById = *createdByIdField
	}
	createdByUsernameField, err := identity(&pb.CreatedByUsername)
	if err != nil {
		return nil, err
	}
	if createdByUsernameField != nil {
		st.CreatedByUsername = *createdByUsernameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listTokenManagementRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listTokenManagementRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Tokens were successfully returned.
type ListTokensResponse struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos []TokenInfo
}

func listTokensResponseToPb(st *ListTokensResponse) (*listTokensResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTokensResponsePb{}

	var tokenInfosPb []tokenInfoPb
	for _, item := range st.TokenInfos {
		itemPb, err := tokenInfoToPb(&item)
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

type listTokensResponsePb struct {
	// Token metadata of each user-created token in the workspace
	TokenInfos []tokenInfoPb `json:"token_infos,omitempty"`
}

func listTokensResponseFromPb(pb *listTokensResponsePb) (*ListTokensResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTokensResponse{}

	var tokenInfosField []TokenInfo
	for _, itemPb := range pb.TokenInfos {
		item, err := tokenInfoFromPb(&itemPb)
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
type listTypePb string

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

func listTypeToPb(st *ListType) (*listTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := listTypePb(*st)
	return &pb, nil
}

func listTypeFromPb(pb *listTypePb) (*ListType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ListType(*pb)
	return &st, nil
}

type MicrosoftTeamsConfig struct {
	// [Input-Only] URL for Microsoft Teams.
	Url string
	// [Output-Only] Whether URL is set.
	UrlSet bool

	ForceSendFields []string
}

func microsoftTeamsConfigToPb(st *MicrosoftTeamsConfig) (*microsoftTeamsConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &microsoftTeamsConfigPb{}
	urlPb, err := identity(&st.Url)
	if err != nil {
		return nil, err
	}
	if urlPb != nil {
		pb.Url = *urlPb
	}

	urlSetPb, err := identity(&st.UrlSet)
	if err != nil {
		return nil, err
	}
	if urlSetPb != nil {
		pb.UrlSet = *urlSetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type microsoftTeamsConfigPb struct {
	// [Input-Only] URL for Microsoft Teams.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func microsoftTeamsConfigFromPb(pb *microsoftTeamsConfigPb) (*MicrosoftTeamsConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MicrosoftTeamsConfig{}
	urlField, err := identity(&pb.Url)
	if err != nil {
		return nil, err
	}
	if urlField != nil {
		st.Url = *urlField
	}
	urlSetField, err := identity(&pb.UrlSet)
	if err != nil {
		return nil, err
	}
	if urlSetField != nil {
		st.UrlSet = *urlSetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *microsoftTeamsConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st microsoftTeamsConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The stable AWS IP CIDR blocks. You can use these to configure the firewall of
// your resources to allow traffic from your Databricks workspace.
type NccAwsStableIpRule struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks []string
}

func nccAwsStableIpRuleToPb(st *NccAwsStableIpRule) (*nccAwsStableIpRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccAwsStableIpRulePb{}

	var cidrBlocksPb []string
	for _, item := range st.CidrBlocks {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			cidrBlocksPb = append(cidrBlocksPb, *itemPb)
		}
	}
	pb.CidrBlocks = cidrBlocksPb

	return pb, nil
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

type nccAwsStableIpRulePb struct {
	// The list of stable IP CIDR blocks from which Databricks network traffic
	// originates when accessing your resources.
	CidrBlocks []string `json:"cidr_blocks,omitempty"`
}

func nccAwsStableIpRuleFromPb(pb *nccAwsStableIpRulePb) (*NccAwsStableIpRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAwsStableIpRule{}

	var cidrBlocksField []string
	for _, itemPb := range pb.CidrBlocks {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			cidrBlocksField = append(cidrBlocksField, *item)
		}
	}
	st.CidrBlocks = cidrBlocksField

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
	ConnectionState NccAzurePrivateEndpointRuleConnectionState
	// Time in epoch milliseconds when this object was created.
	CreationTime int64
	// Whether this private endpoint is deactivated.
	Deactivated bool
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string
	// The name of the Azure private endpoint resource.
	EndpointName string
	// Only used by private endpoints to Azure first-party services. Enum: blob
	// | dfs | sqlServer | mysqlServer
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId string
	// The ID of a network connectivity configuration, which is the parent
	// resource of this private endpoint rule object.
	NetworkConnectivityConfigId string
	// The Azure resource ID of the target resource.
	ResourceId string
	// The ID of a private endpoint rule.
	RuleId string
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64

	ForceSendFields []string
}

func nccAzurePrivateEndpointRuleToPb(st *NccAzurePrivateEndpointRule) (*nccAzurePrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccAzurePrivateEndpointRulePb{}
	connectionStatePb, err := identity(&st.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStatePb != nil {
		pb.ConnectionState = *connectionStatePb
	}

	creationTimePb, err := identity(&st.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	deactivatedPb, err := identity(&st.Deactivated)
	if err != nil {
		return nil, err
	}
	if deactivatedPb != nil {
		pb.Deactivated = *deactivatedPb
	}

	deactivatedAtPb, err := identity(&st.DeactivatedAt)
	if err != nil {
		return nil, err
	}
	if deactivatedAtPb != nil {
		pb.DeactivatedAt = *deactivatedAtPb
	}

	var domainNamesPb []string
	for _, item := range st.DomainNames {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			domainNamesPb = append(domainNamesPb, *itemPb)
		}
	}
	pb.DomainNames = domainNamesPb

	endpointNamePb, err := identity(&st.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNamePb != nil {
		pb.EndpointName = *endpointNamePb
	}

	groupIdPb, err := identity(&st.GroupId)
	if err != nil {
		return nil, err
	}
	if groupIdPb != nil {
		pb.GroupId = *groupIdPb
	}

	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	resourceIdPb, err := identity(&st.ResourceId)
	if err != nil {
		return nil, err
	}
	if resourceIdPb != nil {
		pb.ResourceId = *resourceIdPb
	}

	ruleIdPb, err := identity(&st.RuleId)
	if err != nil {
		return nil, err
	}
	if ruleIdPb != nil {
		pb.RuleId = *ruleIdPb
	}

	updatedTimePb, err := identity(&st.UpdatedTime)
	if err != nil {
		return nil, err
	}
	if updatedTimePb != nil {
		pb.UpdatedTime = *updatedTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type nccAzurePrivateEndpointRulePb struct {
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
	ConnectionState NccAzurePrivateEndpointRuleConnectionState `json:"connection_state,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// Whether this private endpoint is deactivated.
	Deactivated bool `json:"deactivated,omitempty"`
	// Time in epoch milliseconds when this object was deactivated.
	DeactivatedAt int64 `json:"deactivated_at,omitempty"`
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string `json:"domain_names,omitempty"`
	// The name of the Azure private endpoint resource.
	EndpointName string `json:"endpoint_name,omitempty"`
	// Only used by private endpoints to Azure first-party services. Enum: blob
	// | dfs | sqlServer | mysqlServer
	//
	// The sub-resource type (group ID) of the target resource. Note that to
	// connect to workspace root storage (root DBFS), you need two endpoints,
	// one for blob and one for dfs.
	GroupId string `json:"group_id,omitempty"`
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

func nccAzurePrivateEndpointRuleFromPb(pb *nccAzurePrivateEndpointRulePb) (*NccAzurePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAzurePrivateEndpointRule{}
	connectionStateField, err := identity(&pb.ConnectionState)
	if err != nil {
		return nil, err
	}
	if connectionStateField != nil {
		st.ConnectionState = *connectionStateField
	}
	creationTimeField, err := identity(&pb.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	deactivatedField, err := identity(&pb.Deactivated)
	if err != nil {
		return nil, err
	}
	if deactivatedField != nil {
		st.Deactivated = *deactivatedField
	}
	deactivatedAtField, err := identity(&pb.DeactivatedAt)
	if err != nil {
		return nil, err
	}
	if deactivatedAtField != nil {
		st.DeactivatedAt = *deactivatedAtField
	}

	var domainNamesField []string
	for _, itemPb := range pb.DomainNames {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			domainNamesField = append(domainNamesField, *item)
		}
	}
	st.DomainNames = domainNamesField
	endpointNameField, err := identity(&pb.EndpointName)
	if err != nil {
		return nil, err
	}
	if endpointNameField != nil {
		st.EndpointName = *endpointNameField
	}
	groupIdField, err := identity(&pb.GroupId)
	if err != nil {
		return nil, err
	}
	if groupIdField != nil {
		st.GroupId = *groupIdField
	}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	resourceIdField, err := identity(&pb.ResourceId)
	if err != nil {
		return nil, err
	}
	if resourceIdField != nil {
		st.ResourceId = *resourceIdField
	}
	ruleIdField, err := identity(&pb.RuleId)
	if err != nil {
		return nil, err
	}
	if ruleIdField != nil {
		st.RuleId = *ruleIdField
	}
	updatedTimeField, err := identity(&pb.UpdatedTime)
	if err != nil {
		return nil, err
	}
	if updatedTimeField != nil {
		st.UpdatedTime = *updatedTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nccAzurePrivateEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nccAzurePrivateEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NccAzurePrivateEndpointRuleConnectionState string
type nccAzurePrivateEndpointRuleConnectionStatePb string

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

func nccAzurePrivateEndpointRuleConnectionStateToPb(st *NccAzurePrivateEndpointRuleConnectionState) (*nccAzurePrivateEndpointRuleConnectionStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := nccAzurePrivateEndpointRuleConnectionStatePb(*st)
	return &pb, nil
}

func nccAzurePrivateEndpointRuleConnectionStateFromPb(pb *nccAzurePrivateEndpointRuleConnectionStatePb) (*NccAzurePrivateEndpointRuleConnectionState, error) {
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
	Subnets []string
	// The Azure region in which this service endpoint rule applies..
	TargetRegion string
	// The Azure services to which this service endpoint rule applies to.
	TargetServices []EgressResourceType

	ForceSendFields []string
}

func nccAzureServiceEndpointRuleToPb(st *NccAzureServiceEndpointRule) (*nccAzureServiceEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccAzureServiceEndpointRulePb{}

	var subnetsPb []string
	for _, item := range st.Subnets {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			subnetsPb = append(subnetsPb, *itemPb)
		}
	}
	pb.Subnets = subnetsPb

	targetRegionPb, err := identity(&st.TargetRegion)
	if err != nil {
		return nil, err
	}
	if targetRegionPb != nil {
		pb.TargetRegion = *targetRegionPb
	}

	var targetServicesPb []EgressResourceType
	for _, item := range st.TargetServices {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			targetServicesPb = append(targetServicesPb, *itemPb)
		}
	}
	pb.TargetServices = targetServicesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type nccAzureServiceEndpointRulePb struct {
	// The list of subnets from which Databricks network traffic originates when
	// accessing your Azure resources.
	Subnets []string `json:"subnets,omitempty"`
	// The Azure region in which this service endpoint rule applies..
	TargetRegion string `json:"target_region,omitempty"`
	// The Azure services to which this service endpoint rule applies to.
	TargetServices []EgressResourceType `json:"target_services,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func nccAzureServiceEndpointRuleFromPb(pb *nccAzureServiceEndpointRulePb) (*NccAzureServiceEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccAzureServiceEndpointRule{}

	var subnetsField []string
	for _, itemPb := range pb.Subnets {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			subnetsField = append(subnetsField, *item)
		}
	}
	st.Subnets = subnetsField
	targetRegionField, err := identity(&pb.TargetRegion)
	if err != nil {
		return nil, err
	}
	if targetRegionField != nil {
		st.TargetRegion = *targetRegionField
	}

	var targetServicesField []EgressResourceType
	for _, itemPb := range pb.TargetServices {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			targetServicesField = append(targetServicesField, *item)
		}
	}
	st.TargetServices = targetServicesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *nccAzureServiceEndpointRulePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st nccAzureServiceEndpointRulePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NccEgressConfig struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules *NccEgressDefaultRules
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules *NccEgressTargetRules
}

func nccEgressConfigToPb(st *NccEgressConfig) (*nccEgressConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccEgressConfigPb{}
	defaultRulesPb, err := nccEgressDefaultRulesToPb(st.DefaultRules)
	if err != nil {
		return nil, err
	}
	if defaultRulesPb != nil {
		pb.DefaultRules = defaultRulesPb
	}

	targetRulesPb, err := nccEgressTargetRulesToPb(st.TargetRules)
	if err != nil {
		return nil, err
	}
	if targetRulesPb != nil {
		pb.TargetRules = targetRulesPb
	}

	return pb, nil
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

type nccEgressConfigPb struct {
	// The network connectivity rules that are applied by default without
	// resource specific configurations. You can find the stable network
	// information of your serverless compute resources here.
	DefaultRules *nccEgressDefaultRulesPb `json:"default_rules,omitempty"`
	// The network connectivity rules that configured for each destinations.
	// These rules override default rules.
	TargetRules *nccEgressTargetRulesPb `json:"target_rules,omitempty"`
}

func nccEgressConfigFromPb(pb *nccEgressConfigPb) (*NccEgressConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressConfig{}
	defaultRulesField, err := nccEgressDefaultRulesFromPb(pb.DefaultRules)
	if err != nil {
		return nil, err
	}
	if defaultRulesField != nil {
		st.DefaultRules = defaultRulesField
	}
	targetRulesField, err := nccEgressTargetRulesFromPb(pb.TargetRules)
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
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule *NccAwsStableIpRule
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule *NccAzureServiceEndpointRule
}

func nccEgressDefaultRulesToPb(st *NccEgressDefaultRules) (*nccEgressDefaultRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccEgressDefaultRulesPb{}
	awsStableIpRulePb, err := nccAwsStableIpRuleToPb(st.AwsStableIpRule)
	if err != nil {
		return nil, err
	}
	if awsStableIpRulePb != nil {
		pb.AwsStableIpRule = awsStableIpRulePb
	}

	azureServiceEndpointRulePb, err := nccAzureServiceEndpointRuleToPb(st.AzureServiceEndpointRule)
	if err != nil {
		return nil, err
	}
	if azureServiceEndpointRulePb != nil {
		pb.AzureServiceEndpointRule = azureServiceEndpointRulePb
	}

	return pb, nil
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

type nccEgressDefaultRulesPb struct {
	// The stable AWS IP CIDR blocks. You can use these to configure the
	// firewall of your resources to allow traffic from your Databricks
	// workspace.
	AwsStableIpRule *nccAwsStableIpRulePb `json:"aws_stable_ip_rule,omitempty"`
	// The stable Azure service endpoints. You can configure the firewall of
	// your Azure resources to allow traffic from your Databricks serverless
	// compute resources.
	AzureServiceEndpointRule *nccAzureServiceEndpointRulePb `json:"azure_service_endpoint_rule,omitempty"`
}

func nccEgressDefaultRulesFromPb(pb *nccEgressDefaultRulesPb) (*NccEgressDefaultRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressDefaultRules{}
	awsStableIpRuleField, err := nccAwsStableIpRuleFromPb(pb.AwsStableIpRule)
	if err != nil {
		return nil, err
	}
	if awsStableIpRuleField != nil {
		st.AwsStableIpRule = awsStableIpRuleField
	}
	azureServiceEndpointRuleField, err := nccAzureServiceEndpointRuleFromPb(pb.AzureServiceEndpointRule)
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
	AzurePrivateEndpointRules []NccAzurePrivateEndpointRule
}

func nccEgressTargetRulesToPb(st *NccEgressTargetRules) (*nccEgressTargetRulesPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &nccEgressTargetRulesPb{}

	var azurePrivateEndpointRulesPb []nccAzurePrivateEndpointRulePb
	for _, item := range st.AzurePrivateEndpointRules {
		itemPb, err := nccAzurePrivateEndpointRuleToPb(&item)
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

type nccEgressTargetRulesPb struct {
	AzurePrivateEndpointRules []nccAzurePrivateEndpointRulePb `json:"azure_private_endpoint_rules,omitempty"`
}

func nccEgressTargetRulesFromPb(pb *nccEgressTargetRulesPb) (*NccEgressTargetRules, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NccEgressTargetRules{}

	var azurePrivateEndpointRulesField []NccAzurePrivateEndpointRule
	for _, itemPb := range pb.AzurePrivateEndpointRules {
		item, err := nccAzurePrivateEndpointRuleFromPb(&itemPb)
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

// Properties of the new network connectivity configuration.
type NetworkConnectivityConfiguration struct {
	// The Databricks account ID that hosts the credential.
	AccountId string
	// Time in epoch milliseconds when this object was created.
	CreationTime int64
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig *NccEgressConfig
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
	Name string
	// Databricks network connectivity configuration ID.
	NetworkConnectivityConfigId string
	// The region for the network connectivity configuration. Only workspaces in
	// the same region can be attached to the network connectivity
	// configuration.
	Region string
	// Time in epoch milliseconds when this object was updated.
	UpdatedTime int64

	ForceSendFields []string
}

func networkConnectivityConfigurationToPb(st *NetworkConnectivityConfiguration) (*networkConnectivityConfigurationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &networkConnectivityConfigurationPb{}
	accountIdPb, err := identity(&st.AccountId)
	if err != nil {
		return nil, err
	}
	if accountIdPb != nil {
		pb.AccountId = *accountIdPb
	}

	creationTimePb, err := identity(&st.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	egressConfigPb, err := nccEgressConfigToPb(st.EgressConfig)
	if err != nil {
		return nil, err
	}
	if egressConfigPb != nil {
		pb.EgressConfig = egressConfigPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	updatedTimePb, err := identity(&st.UpdatedTime)
	if err != nil {
		return nil, err
	}
	if updatedTimePb != nil {
		pb.UpdatedTime = *updatedTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type networkConnectivityConfigurationPb struct {
	// The Databricks account ID that hosts the credential.
	AccountId string `json:"account_id,omitempty"`
	// Time in epoch milliseconds when this object was created.
	CreationTime int64 `json:"creation_time,omitempty"`
	// The network connectivity rules that apply to network traffic from your
	// serverless compute resources.
	EgressConfig *nccEgressConfigPb `json:"egress_config,omitempty"`
	// The name of the network connectivity configuration. The name can contain
	// alphanumeric characters, hyphens, and underscores. The length must be
	// between 3 and 30 characters. The name must match the regular expression
	// ^[0-9a-zA-Z-_]{3,30}$
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

func networkConnectivityConfigurationFromPb(pb *networkConnectivityConfigurationPb) (*NetworkConnectivityConfiguration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NetworkConnectivityConfiguration{}
	accountIdField, err := identity(&pb.AccountId)
	if err != nil {
		return nil, err
	}
	if accountIdField != nil {
		st.AccountId = *accountIdField
	}
	creationTimeField, err := identity(&pb.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	egressConfigField, err := nccEgressConfigFromPb(pb.EgressConfig)
	if err != nil {
		return nil, err
	}
	if egressConfigField != nil {
		st.EgressConfig = egressConfigField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}
	updatedTimeField, err := identity(&pb.UpdatedTime)
	if err != nil {
		return nil, err
	}
	if updatedTimeField != nil {
		st.UpdatedTime = *updatedTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *networkConnectivityConfigurationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st networkConnectivityConfigurationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotificationDestination struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	Config *Config
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType DestinationType
	// The display name for the notification destination.
	DisplayName string
	// UUID identifying notification destination.
	Id string

	ForceSendFields []string
}

func notificationDestinationToPb(st *NotificationDestination) (*notificationDestinationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notificationDestinationPb{}
	configPb, err := configToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	destinationTypePb, err := identity(&st.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypePb != nil {
		pb.DestinationType = *destinationTypePb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type notificationDestinationPb struct {
	// The configuration for the notification destination. Will be exactly one
	// of the nested configs. Only returns for users with workspace admin
	// permissions.
	Config *configPb `json:"config,omitempty"`
	// [Output-only] The type of the notification destination. The type can not
	// be changed once set.
	DestinationType DestinationType `json:"destination_type,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	Id string `json:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notificationDestinationFromPb(pb *notificationDestinationPb) (*NotificationDestination, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotificationDestination{}
	configField, err := configFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	destinationTypeField, err := identity(&pb.DestinationType)
	if err != nil {
		return nil, err
	}
	if destinationTypeField != nil {
		st.DestinationType = *destinationTypeField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notificationDestinationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notificationDestinationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PagerdutyConfig struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey string
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet bool

	ForceSendFields []string
}

func pagerdutyConfigToPb(st *PagerdutyConfig) (*pagerdutyConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pagerdutyConfigPb{}
	integrationKeyPb, err := identity(&st.IntegrationKey)
	if err != nil {
		return nil, err
	}
	if integrationKeyPb != nil {
		pb.IntegrationKey = *integrationKeyPb
	}

	integrationKeySetPb, err := identity(&st.IntegrationKeySet)
	if err != nil {
		return nil, err
	}
	if integrationKeySetPb != nil {
		pb.IntegrationKeySet = *integrationKeySetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type pagerdutyConfigPb struct {
	// [Input-Only] Integration key for PagerDuty.
	IntegrationKey string `json:"integration_key,omitempty"`
	// [Output-Only] Whether integration key is set.
	IntegrationKeySet bool `json:"integration_key_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pagerdutyConfigFromPb(pb *pagerdutyConfigPb) (*PagerdutyConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PagerdutyConfig{}
	integrationKeyField, err := identity(&pb.IntegrationKey)
	if err != nil {
		return nil, err
	}
	if integrationKeyField != nil {
		st.IntegrationKey = *integrationKeyField
	}
	integrationKeySetField, err := identity(&pb.IntegrationKeySet)
	if err != nil {
		return nil, err
	}
	if integrationKeySetField != nil {
		st.IntegrationKeySet = *integrationKeySetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pagerdutyConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pagerdutyConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Partition by workspace or account
type PartitionId struct {
	// The ID of the workspace.
	WorkspaceId int64

	ForceSendFields []string
}

func partitionIdToPb(st *PartitionId) (*partitionIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &partitionIdPb{}
	workspaceIdPb, err := identity(&st.WorkspaceId)
	if err != nil {
		return nil, err
	}
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type partitionIdPb struct {
	// The ID of the workspace.
	WorkspaceId int64 `json:"workspaceId,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func partitionIdFromPb(pb *partitionIdPb) (*PartitionId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PartitionId{}
	workspaceIdField, err := identity(&pb.WorkspaceId)
	if err != nil {
		return nil, err
	}
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *partitionIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st partitionIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PersonalComputeMessage struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value PersonalComputeMessageEnum
}

func personalComputeMessageToPb(st *PersonalComputeMessage) (*personalComputeMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &personalComputeMessagePb{}
	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	return pb, nil
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

type personalComputeMessagePb struct {
	// ON: Grants all users in all workspaces access to the Personal Compute
	// default policy, allowing all users to create single-machine compute
	// resources. DELEGATE: Moves access control for the Personal Compute
	// default policy to individual workspaces and requires a workspace’s
	// users or groups to be added to the ACLs of that workspace’s Personal
	// Compute default policy before they will be able to create compute
	// resources through that policy.
	Value PersonalComputeMessageEnum `json:"value"`
}

func personalComputeMessageFromPb(pb *personalComputeMessagePb) (*PersonalComputeMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalComputeMessage{}
	valueField, err := identity(&pb.Value)
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
type personalComputeMessageEnumPb string

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

func personalComputeMessageEnumToPb(st *PersonalComputeMessageEnum) (*personalComputeMessageEnumPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := personalComputeMessageEnumPb(*st)
	return &pb, nil
}

func personalComputeMessageEnumFromPb(pb *personalComputeMessageEnumPb) (*PersonalComputeMessageEnum, error) {
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
	Etag string

	PersonalCompute PersonalComputeMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func personalComputeSettingToPb(st *PersonalComputeSetting) (*personalComputeSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &personalComputeSettingPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	personalComputePb, err := personalComputeMessageToPb(&st.PersonalCompute)
	if err != nil {
		return nil, err
	}
	if personalComputePb != nil {
		pb.PersonalCompute = *personalComputePb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type personalComputeSettingPb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	PersonalCompute personalComputeMessagePb `json:"personal_compute"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func personalComputeSettingFromPb(pb *personalComputeSettingPb) (*PersonalComputeSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PersonalComputeSetting{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	personalComputeField, err := personalComputeMessageFromPb(&pb.PersonalCompute)
	if err != nil {
		return nil, err
	}
	if personalComputeField != nil {
		st.PersonalCompute = *personalComputeField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *personalComputeSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st personalComputeSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PublicTokenInfo struct {
	// Comment the token was created with, if applicable.
	Comment string
	// Server time (in epoch milliseconds) when the token was created.
	CreationTime int64
	// Server time (in epoch milliseconds) when the token will expire, or -1 if
	// not applicable.
	ExpiryTime int64
	// The ID of this token.
	TokenId string

	ForceSendFields []string
}

func publicTokenInfoToPb(st *PublicTokenInfo) (*publicTokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publicTokenInfoPb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	creationTimePb, err := identity(&st.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	expiryTimePb, err := identity(&st.ExpiryTime)
	if err != nil {
		return nil, err
	}
	if expiryTimePb != nil {
		pb.ExpiryTime = *expiryTimePb
	}

	tokenIdPb, err := identity(&st.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdPb != nil {
		pb.TokenId = *tokenIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type publicTokenInfoPb struct {
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

func publicTokenInfoFromPb(pb *publicTokenInfoPb) (*PublicTokenInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublicTokenInfo{}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	creationTimeField, err := identity(&pb.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	expiryTimeField, err := identity(&pb.ExpiryTime)
	if err != nil {
		return nil, err
	}
	if expiryTimeField != nil {
		st.ExpiryTime = *expiryTimeField
	}
	tokenIdField, err := identity(&pb.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdField != nil {
		st.TokenId = *tokenIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publicTokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publicTokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Details required to replace an IP access list.
type ReplaceIpAccessList struct {
	// Specifies whether this IP access list is enabled.
	Enabled bool
	// The ID for the corresponding IP access list
	IpAccessListId string

	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	Label string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType
}

func replaceIpAccessListToPb(st *ReplaceIpAccessList) (*replaceIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &replaceIpAccessListPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	ipAccessListIdPb, err := identity(&st.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdPb != nil {
		pb.IpAccessListId = *ipAccessListIdPb
	}

	var ipAddressesPb []string
	for _, item := range st.IpAddresses {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			ipAddressesPb = append(ipAddressesPb, *itemPb)
		}
	}
	pb.IpAddresses = ipAddressesPb

	labelPb, err := identity(&st.Label)
	if err != nil {
		return nil, err
	}
	if labelPb != nil {
		pb.Label = *labelPb
	}

	listTypePb, err := identity(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}

	return pb, nil
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

type replaceIpAccessListPb struct {
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

func replaceIpAccessListFromPb(pb *replaceIpAccessListPb) (*ReplaceIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReplaceIpAccessList{}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	ipAccessListIdField, err := identity(&pb.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdField != nil {
		st.IpAccessListId = *ipAccessListIdField
	}

	var ipAddressesField []string
	for _, itemPb := range pb.IpAddresses {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			ipAddressesField = append(ipAddressesField, *item)
		}
	}
	st.IpAddresses = ipAddressesField
	labelField, err := identity(&pb.Label)
	if err != nil {
		return nil, err
	}
	if labelField != nil {
		st.Label = *labelField
	}
	listTypeField, err := identity(&pb.ListType)
	if err != nil {
		return nil, err
	}
	if listTypeField != nil {
		st.ListType = *listTypeField
	}

	return st, nil
}

type ReplaceResponse struct {
}

func replaceResponseToPb(st *ReplaceResponse) (*replaceResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &replaceResponsePb{}

	return pb, nil
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

type replaceResponsePb struct {
}

func replaceResponseFromPb(pb *replaceResponsePb) (*ReplaceResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReplaceResponse{}

	return st, nil
}

type RestrictWorkspaceAdminsMessage struct {
	Status RestrictWorkspaceAdminsMessageStatus
}

func restrictWorkspaceAdminsMessageToPb(st *RestrictWorkspaceAdminsMessage) (*restrictWorkspaceAdminsMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restrictWorkspaceAdminsMessagePb{}
	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	return pb, nil
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

type restrictWorkspaceAdminsMessagePb struct {
	Status RestrictWorkspaceAdminsMessageStatus `json:"status"`
}

func restrictWorkspaceAdminsMessageFromPb(pb *restrictWorkspaceAdminsMessagePb) (*RestrictWorkspaceAdminsMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestrictWorkspaceAdminsMessage{}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

type RestrictWorkspaceAdminsMessageStatus string
type restrictWorkspaceAdminsMessageStatusPb string

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

func restrictWorkspaceAdminsMessageStatusToPb(st *RestrictWorkspaceAdminsMessageStatus) (*restrictWorkspaceAdminsMessageStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := restrictWorkspaceAdminsMessageStatusPb(*st)
	return &pb, nil
}

func restrictWorkspaceAdminsMessageStatusFromPb(pb *restrictWorkspaceAdminsMessageStatusPb) (*RestrictWorkspaceAdminsMessageStatus, error) {
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
	Etag string

	RestrictWorkspaceAdmins RestrictWorkspaceAdminsMessage
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string

	ForceSendFields []string
}

func restrictWorkspaceAdminsSettingToPb(st *RestrictWorkspaceAdminsSetting) (*restrictWorkspaceAdminsSettingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restrictWorkspaceAdminsSettingPb{}
	etagPb, err := identity(&st.Etag)
	if err != nil {
		return nil, err
	}
	if etagPb != nil {
		pb.Etag = *etagPb
	}

	restrictWorkspaceAdminsPb, err := restrictWorkspaceAdminsMessageToPb(&st.RestrictWorkspaceAdmins)
	if err != nil {
		return nil, err
	}
	if restrictWorkspaceAdminsPb != nil {
		pb.RestrictWorkspaceAdmins = *restrictWorkspaceAdminsPb
	}

	settingNamePb, err := identity(&st.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNamePb != nil {
		pb.SettingName = *settingNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type restrictWorkspaceAdminsSettingPb struct {
	// etag used for versioning. The response is at least as fresh as the eTag
	// provided. This is used for optimistic concurrency control as a way to
	// help prevent simultaneous writes of a setting overwriting each other. It
	// is strongly suggested that systems make use of the etag in the read ->
	// update pattern to perform setting updates in order to avoid race
	// conditions. That is, get an etag from a GET request, and pass it with the
	// PATCH request to identify the setting version you are updating.
	Etag string `json:"etag,omitempty"`

	RestrictWorkspaceAdmins restrictWorkspaceAdminsMessagePb `json:"restrict_workspace_admins"`
	// Name of the corresponding setting. This field is populated in the
	// response, but it will not be respected even if it's set in the request
	// body. The setting name in the path parameter will be respected instead.
	// Setting name is required to be 'default' if the setting only has one
	// instance per workspace.
	SettingName string `json:"setting_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restrictWorkspaceAdminsSettingFromPb(pb *restrictWorkspaceAdminsSettingPb) (*RestrictWorkspaceAdminsSetting, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestrictWorkspaceAdminsSetting{}
	etagField, err := identity(&pb.Etag)
	if err != nil {
		return nil, err
	}
	if etagField != nil {
		st.Etag = *etagField
	}
	restrictWorkspaceAdminsField, err := restrictWorkspaceAdminsMessageFromPb(&pb.RestrictWorkspaceAdmins)
	if err != nil {
		return nil, err
	}
	if restrictWorkspaceAdminsField != nil {
		st.RestrictWorkspaceAdmins = *restrictWorkspaceAdminsField
	}
	settingNameField, err := identity(&pb.SettingName)
	if err != nil {
		return nil, err
	}
	if settingNameField != nil {
		st.SettingName = *settingNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restrictWorkspaceAdminsSettingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restrictWorkspaceAdminsSettingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RevokeTokenRequest struct {
	// The ID of the token to be revoked.
	TokenId string
}

func revokeTokenRequestToPb(st *RevokeTokenRequest) (*revokeTokenRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &revokeTokenRequestPb{}
	tokenIdPb, err := identity(&st.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdPb != nil {
		pb.TokenId = *tokenIdPb
	}

	return pb, nil
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

type revokeTokenRequestPb struct {
	// The ID of the token to be revoked.
	TokenId string `json:"token_id"`
}

func revokeTokenRequestFromPb(pb *revokeTokenRequestPb) (*RevokeTokenRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RevokeTokenRequest{}
	tokenIdField, err := identity(&pb.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdField != nil {
		st.TokenId = *tokenIdField
	}

	return st, nil
}

type RevokeTokenResponse struct {
}

func revokeTokenResponseToPb(st *RevokeTokenResponse) (*revokeTokenResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &revokeTokenResponsePb{}

	return pb, nil
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

type revokeTokenResponsePb struct {
}

func revokeTokenResponseFromPb(pb *revokeTokenResponsePb) (*RevokeTokenResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RevokeTokenResponse{}

	return st, nil
}

type SetStatusResponse struct {
}

func setStatusResponseToPb(st *SetStatusResponse) (*setStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setStatusResponsePb{}

	return pb, nil
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

type setStatusResponsePb struct {
}

func setStatusResponseFromPb(pb *setStatusResponsePb) (*SetStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetStatusResponse{}

	return st, nil
}

type SlackConfig struct {
	// [Input-Only] URL for Slack destination.
	Url string
	// [Output-Only] Whether URL is set.
	UrlSet bool

	ForceSendFields []string
}

func slackConfigToPb(st *SlackConfig) (*slackConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &slackConfigPb{}
	urlPb, err := identity(&st.Url)
	if err != nil {
		return nil, err
	}
	if urlPb != nil {
		pb.Url = *urlPb
	}

	urlSetPb, err := identity(&st.UrlSet)
	if err != nil {
		return nil, err
	}
	if urlSetPb != nil {
		pb.UrlSet = *urlSetPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type slackConfigPb struct {
	// [Input-Only] URL for Slack destination.
	Url string `json:"url,omitempty"`
	// [Output-Only] Whether URL is set.
	UrlSet bool `json:"url_set,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func slackConfigFromPb(pb *slackConfigPb) (*SlackConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SlackConfig{}
	urlField, err := identity(&pb.Url)
	if err != nil {
		return nil, err
	}
	if urlField != nil {
		st.Url = *urlField
	}
	urlSetField, err := identity(&pb.UrlSet)
	if err != nil {
		return nil, err
	}
	if urlSetField != nil {
		st.UrlSet = *urlSetField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *slackConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st slackConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StringMessage struct {
	// Represents a generic string value.
	Value string

	ForceSendFields []string
}

func stringMessageToPb(st *StringMessage) (*stringMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stringMessagePb{}
	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type stringMessagePb struct {
	// Represents a generic string value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func stringMessageFromPb(pb *stringMessagePb) (*StringMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StringMessage{}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *stringMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st stringMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenAccessControlRequest struct {
	// name of the group
	GroupName string
	// Permission level
	PermissionLevel TokenPermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func tokenAccessControlRequestToPb(st *TokenAccessControlRequest) (*tokenAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenAccessControlRequestPb{}
	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tokenAccessControlRequestPb struct {
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

func tokenAccessControlRequestFromPb(pb *tokenAccessControlRequestPb) (*TokenAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessControlRequest{}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenAccessControlResponse struct {
	// All permissions.
	AllPermissions []TokenPermission
	// Display name of the user or service principal.
	DisplayName string
	// name of the group
	GroupName string
	// Name of the service principal.
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func tokenAccessControlResponseToPb(st *TokenAccessControlResponse) (*tokenAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenAccessControlResponsePb{}

	var allPermissionsPb []tokenPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := tokenPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tokenAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []tokenPermissionPb `json:"all_permissions,omitempty"`
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

func tokenAccessControlResponseFromPb(pb *tokenAccessControlResponsePb) (*TokenAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessControlResponse{}

	var allPermissionsField []TokenPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := tokenPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenInfo struct {
	// Comment that describes the purpose of the token, specified by the token
	// creator.
	Comment string
	// User ID of the user that created the token.
	CreatedById int64
	// Username of the user that created the token.
	CreatedByUsername string
	// Timestamp when the token was created.
	CreationTime int64
	// Timestamp when the token expires.
	ExpiryTime int64
	// Approximate timestamp for the day the token was last used. Accurate up to
	// 1 day.
	LastUsedDay int64
	// User ID of the user that owns the token.
	OwnerId int64
	// ID of the token.
	TokenId string
	// If applicable, the ID of the workspace that the token was created in.
	WorkspaceId int64

	ForceSendFields []string
}

func tokenInfoToPb(st *TokenInfo) (*tokenInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenInfoPb{}
	commentPb, err := identity(&st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = *commentPb
	}

	createdByIdPb, err := identity(&st.CreatedById)
	if err != nil {
		return nil, err
	}
	if createdByIdPb != nil {
		pb.CreatedById = *createdByIdPb
	}

	createdByUsernamePb, err := identity(&st.CreatedByUsername)
	if err != nil {
		return nil, err
	}
	if createdByUsernamePb != nil {
		pb.CreatedByUsername = *createdByUsernamePb
	}

	creationTimePb, err := identity(&st.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimePb != nil {
		pb.CreationTime = *creationTimePb
	}

	expiryTimePb, err := identity(&st.ExpiryTime)
	if err != nil {
		return nil, err
	}
	if expiryTimePb != nil {
		pb.ExpiryTime = *expiryTimePb
	}

	lastUsedDayPb, err := identity(&st.LastUsedDay)
	if err != nil {
		return nil, err
	}
	if lastUsedDayPb != nil {
		pb.LastUsedDay = *lastUsedDayPb
	}

	ownerIdPb, err := identity(&st.OwnerId)
	if err != nil {
		return nil, err
	}
	if ownerIdPb != nil {
		pb.OwnerId = *ownerIdPb
	}

	tokenIdPb, err := identity(&st.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdPb != nil {
		pb.TokenId = *tokenIdPb
	}

	workspaceIdPb, err := identity(&st.WorkspaceId)
	if err != nil {
		return nil, err
	}
	if workspaceIdPb != nil {
		pb.WorkspaceId = *workspaceIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tokenInfoPb struct {
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

func tokenInfoFromPb(pb *tokenInfoPb) (*TokenInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenInfo{}
	commentField, err := identity(&pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = *commentField
	}
	createdByIdField, err := identity(&pb.CreatedById)
	if err != nil {
		return nil, err
	}
	if createdByIdField != nil {
		st.CreatedById = *createdByIdField
	}
	createdByUsernameField, err := identity(&pb.CreatedByUsername)
	if err != nil {
		return nil, err
	}
	if createdByUsernameField != nil {
		st.CreatedByUsername = *createdByUsernameField
	}
	creationTimeField, err := identity(&pb.CreationTime)
	if err != nil {
		return nil, err
	}
	if creationTimeField != nil {
		st.CreationTime = *creationTimeField
	}
	expiryTimeField, err := identity(&pb.ExpiryTime)
	if err != nil {
		return nil, err
	}
	if expiryTimeField != nil {
		st.ExpiryTime = *expiryTimeField
	}
	lastUsedDayField, err := identity(&pb.LastUsedDay)
	if err != nil {
		return nil, err
	}
	if lastUsedDayField != nil {
		st.LastUsedDay = *lastUsedDayField
	}
	ownerIdField, err := identity(&pb.OwnerId)
	if err != nil {
		return nil, err
	}
	if ownerIdField != nil {
		st.OwnerId = *ownerIdField
	}
	tokenIdField, err := identity(&pb.TokenId)
	if err != nil {
		return nil, err
	}
	if tokenIdField != nil {
		st.TokenId = *tokenIdField
	}
	workspaceIdField, err := identity(&pb.WorkspaceId)
	if err != nil {
		return nil, err
	}
	if workspaceIdField != nil {
		st.WorkspaceId = *workspaceIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenPermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel TokenPermissionLevel

	ForceSendFields []string
}

func tokenPermissionToPb(st *TokenPermission) (*tokenPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionPb{}
	inheritedPb, err := identity(&st.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedPb != nil {
		pb.Inherited = *inheritedPb
	}

	var inheritedFromObjectPb []string
	for _, item := range st.InheritedFromObject {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			inheritedFromObjectPb = append(inheritedFromObjectPb, *itemPb)
		}
	}
	pb.InheritedFromObject = inheritedFromObjectPb

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tokenPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenPermissionFromPb(pb *tokenPermissionPb) (*TokenPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermission{}
	inheritedField, err := identity(&pb.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedField != nil {
		st.Inherited = *inheritedField
	}

	var inheritedFromObjectField []string
	for _, itemPb := range pb.InheritedFromObject {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			inheritedFromObjectField = append(inheritedFromObjectField, *item)
		}
	}
	st.InheritedFromObject = inheritedFromObjectField
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type TokenPermissionLevel string
type tokenPermissionLevelPb string

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

func tokenPermissionLevelToPb(st *TokenPermissionLevel) (*tokenPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := tokenPermissionLevelPb(*st)
	return &pb, nil
}

func tokenPermissionLevelFromPb(pb *tokenPermissionLevelPb) (*TokenPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := TokenPermissionLevel(*pb)
	return &st, nil
}

type TokenPermissions struct {
	AccessControlList []TokenAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
}

func tokenPermissionsToPb(st *TokenPermissions) (*tokenPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionsPb{}

	var accessControlListPb []tokenAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := tokenAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	objectIdPb, err := identity(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	objectTypePb, err := identity(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tokenPermissionsPb struct {
	AccessControlList []tokenAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenPermissionsFromPb(pb *tokenPermissionsPb) (*TokenPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissions{}

	var accessControlListField []TokenAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := tokenAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	objectIdField, err := identity(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := identity(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenPermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel TokenPermissionLevel

	ForceSendFields []string
}

func tokenPermissionsDescriptionToPb(st *TokenPermissionsDescription) (*tokenPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionsDescriptionPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type tokenPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel TokenPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenPermissionsDescriptionFromPb(pb *tokenPermissionsDescriptionPb) (*TokenPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissionsDescription{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenPermissionsRequest struct {
	AccessControlList []TokenAccessControlRequest
}

func tokenPermissionsRequestToPb(st *TokenPermissionsRequest) (*tokenPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenPermissionsRequestPb{}

	var accessControlListPb []tokenAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := tokenAccessControlRequestToPb(&item)
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

type tokenPermissionsRequestPb struct {
	AccessControlList []tokenAccessControlRequestPb `json:"access_control_list,omitempty"`
}

func tokenPermissionsRequestFromPb(pb *tokenPermissionsRequestPb) (*TokenPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenPermissionsRequest{}

	var accessControlListField []TokenAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := tokenAccessControlRequestFromPb(&itemPb)
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
type tokenTypePb string

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

func tokenTypeToPb(st *TokenType) (*tokenTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := tokenTypePb(*st)
	return &pb, nil
}

func tokenTypeFromPb(pb *tokenTypePb) (*TokenType, error) {
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
	FieldMask string

	Setting AccountIpAccessEnable
}

func updateAccountIpAccessEnableRequestToPb(st *UpdateAccountIpAccessEnableRequest) (*updateAccountIpAccessEnableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAccountIpAccessEnableRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := accountIpAccessEnableToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateAccountIpAccessEnableRequestPb struct {
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

	Setting accountIpAccessEnablePb `json:"setting"`
}

func updateAccountIpAccessEnableRequestFromPb(pb *updateAccountIpAccessEnableRequestPb) (*UpdateAccountIpAccessEnableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAccountIpAccessEnableRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := accountIpAccessEnableFromPb(&pb.Setting)
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
	FieldMask string

	Setting AibiDashboardEmbeddingAccessPolicySetting
}

func updateAibiDashboardEmbeddingAccessPolicySettingRequestToPb(st *UpdateAibiDashboardEmbeddingAccessPolicySettingRequest) (*updateAibiDashboardEmbeddingAccessPolicySettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAibiDashboardEmbeddingAccessPolicySettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := aibiDashboardEmbeddingAccessPolicySettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateAibiDashboardEmbeddingAccessPolicySettingRequestPb struct {
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

	Setting aibiDashboardEmbeddingAccessPolicySettingPb `json:"setting"`
}

func updateAibiDashboardEmbeddingAccessPolicySettingRequestFromPb(pb *updateAibiDashboardEmbeddingAccessPolicySettingRequestPb) (*UpdateAibiDashboardEmbeddingAccessPolicySettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAibiDashboardEmbeddingAccessPolicySettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := aibiDashboardEmbeddingAccessPolicySettingFromPb(&pb.Setting)
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
	FieldMask string

	Setting AibiDashboardEmbeddingApprovedDomainsSetting
}

func updateAibiDashboardEmbeddingApprovedDomainsSettingRequestToPb(st *UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest) (*updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := aibiDashboardEmbeddingApprovedDomainsSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb struct {
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

	Setting aibiDashboardEmbeddingApprovedDomainsSettingPb `json:"setting"`
}

func updateAibiDashboardEmbeddingApprovedDomainsSettingRequestFromPb(pb *updateAibiDashboardEmbeddingApprovedDomainsSettingRequestPb) (*UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAibiDashboardEmbeddingApprovedDomainsSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := aibiDashboardEmbeddingApprovedDomainsSettingFromPb(&pb.Setting)
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
	FieldMask string

	Setting AutomaticClusterUpdateSetting
}

func updateAutomaticClusterUpdateSettingRequestToPb(st *UpdateAutomaticClusterUpdateSettingRequest) (*updateAutomaticClusterUpdateSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAutomaticClusterUpdateSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := automaticClusterUpdateSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateAutomaticClusterUpdateSettingRequestPb struct {
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

	Setting automaticClusterUpdateSettingPb `json:"setting"`
}

func updateAutomaticClusterUpdateSettingRequestFromPb(pb *updateAutomaticClusterUpdateSettingRequestPb) (*UpdateAutomaticClusterUpdateSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAutomaticClusterUpdateSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := automaticClusterUpdateSettingFromPb(&pb.Setting)
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
	FieldMask string

	Setting ComplianceSecurityProfileSetting
}

func updateComplianceSecurityProfileSettingRequestToPb(st *UpdateComplianceSecurityProfileSettingRequest) (*updateComplianceSecurityProfileSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateComplianceSecurityProfileSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := complianceSecurityProfileSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateComplianceSecurityProfileSettingRequestPb struct {
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

	Setting complianceSecurityProfileSettingPb `json:"setting"`
}

func updateComplianceSecurityProfileSettingRequestFromPb(pb *updateComplianceSecurityProfileSettingRequestPb) (*UpdateComplianceSecurityProfileSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateComplianceSecurityProfileSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := complianceSecurityProfileSettingFromPb(&pb.Setting)
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
	FieldMask string

	Setting CspEnablementAccountSetting
}

func updateCspEnablementAccountSettingRequestToPb(st *UpdateCspEnablementAccountSettingRequest) (*updateCspEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCspEnablementAccountSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := cspEnablementAccountSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateCspEnablementAccountSettingRequestPb struct {
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

	Setting cspEnablementAccountSettingPb `json:"setting"`
}

func updateCspEnablementAccountSettingRequestFromPb(pb *updateCspEnablementAccountSettingRequestPb) (*UpdateCspEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCspEnablementAccountSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := cspEnablementAccountSettingFromPb(&pb.Setting)
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
	Setting DefaultNamespaceSetting
}

func updateDefaultNamespaceSettingRequestToPb(st *UpdateDefaultNamespaceSettingRequest) (*updateDefaultNamespaceSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDefaultNamespaceSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := defaultNamespaceSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateDefaultNamespaceSettingRequestPb struct {
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
	Setting defaultNamespaceSettingPb `json:"setting"`
}

func updateDefaultNamespaceSettingRequestFromPb(pb *updateDefaultNamespaceSettingRequestPb) (*UpdateDefaultNamespaceSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDefaultNamespaceSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := defaultNamespaceSettingFromPb(&pb.Setting)
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
	FieldMask string

	Setting DisableLegacyAccess
}

func updateDisableLegacyAccessRequestToPb(st *UpdateDisableLegacyAccessRequest) (*updateDisableLegacyAccessRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDisableLegacyAccessRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := disableLegacyAccessToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateDisableLegacyAccessRequestPb struct {
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

	Setting disableLegacyAccessPb `json:"setting"`
}

func updateDisableLegacyAccessRequestFromPb(pb *updateDisableLegacyAccessRequestPb) (*UpdateDisableLegacyAccessRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyAccessRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := disableLegacyAccessFromPb(&pb.Setting)
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
	FieldMask string

	Setting DisableLegacyDbfs
}

func updateDisableLegacyDbfsRequestToPb(st *UpdateDisableLegacyDbfsRequest) (*updateDisableLegacyDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDisableLegacyDbfsRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := disableLegacyDbfsToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateDisableLegacyDbfsRequestPb struct {
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

	Setting disableLegacyDbfsPb `json:"setting"`
}

func updateDisableLegacyDbfsRequestFromPb(pb *updateDisableLegacyDbfsRequestPb) (*UpdateDisableLegacyDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyDbfsRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := disableLegacyDbfsFromPb(&pb.Setting)
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
	FieldMask string

	Setting DisableLegacyFeatures
}

func updateDisableLegacyFeaturesRequestToPb(st *UpdateDisableLegacyFeaturesRequest) (*updateDisableLegacyFeaturesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateDisableLegacyFeaturesRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := disableLegacyFeaturesToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateDisableLegacyFeaturesRequestPb struct {
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

	Setting disableLegacyFeaturesPb `json:"setting"`
}

func updateDisableLegacyFeaturesRequestFromPb(pb *updateDisableLegacyFeaturesRequestPb) (*UpdateDisableLegacyFeaturesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateDisableLegacyFeaturesRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := disableLegacyFeaturesFromPb(&pb.Setting)
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
	FieldMask []string

	Setting EnableExportNotebook
}

func updateEnableExportNotebookRequestToPb(st *UpdateEnableExportNotebookRequest) (*updateEnableExportNotebookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnableExportNotebookRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := fieldMaskToPb(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := enableExportNotebookToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateEnableExportNotebookRequestPb struct {
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

	Setting enableExportNotebookPb `json:"setting"`
}

func updateEnableExportNotebookRequestFromPb(pb *updateEnableExportNotebookRequestPb) (*UpdateEnableExportNotebookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableExportNotebookRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := fieldMaskFromPb(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := enableExportNotebookFromPb(&pb.Setting)
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
	FieldMask []string

	Setting EnableNotebookTableClipboard
}

func updateEnableNotebookTableClipboardRequestToPb(st *UpdateEnableNotebookTableClipboardRequest) (*updateEnableNotebookTableClipboardRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnableNotebookTableClipboardRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := fieldMaskToPb(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := enableNotebookTableClipboardToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateEnableNotebookTableClipboardRequestPb struct {
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

	Setting enableNotebookTableClipboardPb `json:"setting"`
}

func updateEnableNotebookTableClipboardRequestFromPb(pb *updateEnableNotebookTableClipboardRequestPb) (*UpdateEnableNotebookTableClipboardRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableNotebookTableClipboardRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := fieldMaskFromPb(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := enableNotebookTableClipboardFromPb(&pb.Setting)
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
	FieldMask string

	Setting EnableResultsDownloading
}

func updateEnableResultsDownloadingRequestToPb(st *UpdateEnableResultsDownloadingRequest) (*updateEnableResultsDownloadingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnableResultsDownloadingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := enableResultsDownloadingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateEnableResultsDownloadingRequestPb struct {
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

	Setting enableResultsDownloadingPb `json:"setting"`
}

func updateEnableResultsDownloadingRequestFromPb(pb *updateEnableResultsDownloadingRequestPb) (*UpdateEnableResultsDownloadingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnableResultsDownloadingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := enableResultsDownloadingFromPb(&pb.Setting)
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
	FieldMask string

	Setting EnhancedSecurityMonitoringSetting
}

func updateEnhancedSecurityMonitoringSettingRequestToPb(st *UpdateEnhancedSecurityMonitoringSettingRequest) (*updateEnhancedSecurityMonitoringSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEnhancedSecurityMonitoringSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := enhancedSecurityMonitoringSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateEnhancedSecurityMonitoringSettingRequestPb struct {
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

	Setting enhancedSecurityMonitoringSettingPb `json:"setting"`
}

func updateEnhancedSecurityMonitoringSettingRequestFromPb(pb *updateEnhancedSecurityMonitoringSettingRequestPb) (*UpdateEnhancedSecurityMonitoringSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEnhancedSecurityMonitoringSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := enhancedSecurityMonitoringSettingFromPb(&pb.Setting)
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
	FieldMask string

	Setting EsmEnablementAccountSetting
}

func updateEsmEnablementAccountSettingRequestToPb(st *UpdateEsmEnablementAccountSettingRequest) (*updateEsmEnablementAccountSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateEsmEnablementAccountSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := esmEnablementAccountSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateEsmEnablementAccountSettingRequestPb struct {
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

	Setting esmEnablementAccountSettingPb `json:"setting"`
}

func updateEsmEnablementAccountSettingRequestFromPb(pb *updateEsmEnablementAccountSettingRequestPb) (*UpdateEsmEnablementAccountSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateEsmEnablementAccountSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := esmEnablementAccountSettingFromPb(&pb.Setting)
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
	Enabled bool
	// The ID for the corresponding IP access list
	IpAccessListId string

	IpAddresses []string
	// Label for the IP access list. This **cannot** be empty.
	Label string
	// Type of IP access list. Valid values are as follows and are
	// case-sensitive:
	//
	// * `ALLOW`: An allow list. Include this IP or range. * `BLOCK`: A block
	// list. Exclude this IP or range. IP addresses in the block list are
	// excluded even if they are included in an allow list.
	ListType ListType

	ForceSendFields []string
}

func updateIpAccessListToPb(st *UpdateIpAccessList) (*updateIpAccessListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateIpAccessListPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	ipAccessListIdPb, err := identity(&st.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdPb != nil {
		pb.IpAccessListId = *ipAccessListIdPb
	}

	var ipAddressesPb []string
	for _, item := range st.IpAddresses {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			ipAddressesPb = append(ipAddressesPb, *itemPb)
		}
	}
	pb.IpAddresses = ipAddressesPb

	labelPb, err := identity(&st.Label)
	if err != nil {
		return nil, err
	}
	if labelPb != nil {
		pb.Label = *labelPb
	}

	listTypePb, err := identity(&st.ListType)
	if err != nil {
		return nil, err
	}
	if listTypePb != nil {
		pb.ListType = *listTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateIpAccessListPb struct {
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

func updateIpAccessListFromPb(pb *updateIpAccessListPb) (*UpdateIpAccessList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateIpAccessList{}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	ipAccessListIdField, err := identity(&pb.IpAccessListId)
	if err != nil {
		return nil, err
	}
	if ipAccessListIdField != nil {
		st.IpAccessListId = *ipAccessListIdField
	}

	var ipAddressesField []string
	for _, itemPb := range pb.IpAddresses {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			ipAddressesField = append(ipAddressesField, *item)
		}
	}
	st.IpAddresses = ipAddressesField
	labelField, err := identity(&pb.Label)
	if err != nil {
		return nil, err
	}
	if labelField != nil {
		st.Label = *labelField
	}
	listTypeField, err := identity(&pb.ListType)
	if err != nil {
		return nil, err
	}
	if listTypeField != nil {
		st.ListType = *listTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateIpAccessListPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateIpAccessListPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Update a private endpoint rule
type UpdateNccAzurePrivateEndpointRulePublicRequest struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
	PrivateEndpointRule UpdatePrivateEndpointRule
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	UpdateMask []string
}

func updateNccAzurePrivateEndpointRulePublicRequestToPb(st *UpdateNccAzurePrivateEndpointRulePublicRequest) (*updateNccAzurePrivateEndpointRulePublicRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateNccAzurePrivateEndpointRulePublicRequestPb{}
	networkConnectivityConfigIdPb, err := identity(&st.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdPb != nil {
		pb.NetworkConnectivityConfigId = *networkConnectivityConfigIdPb
	}

	privateEndpointRulePb, err := updatePrivateEndpointRuleToPb(&st.PrivateEndpointRule)
	if err != nil {
		return nil, err
	}
	if privateEndpointRulePb != nil {
		pb.PrivateEndpointRule = *privateEndpointRulePb
	}

	privateEndpointRuleIdPb, err := identity(&st.PrivateEndpointRuleId)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleIdPb != nil {
		pb.PrivateEndpointRuleId = *privateEndpointRuleIdPb
	}

	updateMaskPb, err := fieldMaskToPb(&st.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	return pb, nil
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

type updateNccAzurePrivateEndpointRulePublicRequestPb struct {
	// Your Network Connectivity Configuration ID.
	NetworkConnectivityConfigId string `json:"-" url:"-"`
	// Properties of the new private endpoint rule. Note that you must approve
	// the endpoint in Azure portal after initialization.
	PrivateEndpointRule updatePrivateEndpointRulePb `json:"private_endpoint_rule"`
	// Your private endpoint rule ID.
	PrivateEndpointRuleId string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	UpdateMask string `json:"-" url:"update_mask"`
}

func updateNccAzurePrivateEndpointRulePublicRequestFromPb(pb *updateNccAzurePrivateEndpointRulePublicRequestPb) (*UpdateNccAzurePrivateEndpointRulePublicRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNccAzurePrivateEndpointRulePublicRequest{}
	networkConnectivityConfigIdField, err := identity(&pb.NetworkConnectivityConfigId)
	if err != nil {
		return nil, err
	}
	if networkConnectivityConfigIdField != nil {
		st.NetworkConnectivityConfigId = *networkConnectivityConfigIdField
	}
	privateEndpointRuleField, err := updatePrivateEndpointRuleFromPb(&pb.PrivateEndpointRule)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleField != nil {
		st.PrivateEndpointRule = *privateEndpointRuleField
	}
	privateEndpointRuleIdField, err := identity(&pb.PrivateEndpointRuleId)
	if err != nil {
		return nil, err
	}
	if privateEndpointRuleIdField != nil {
		st.PrivateEndpointRuleId = *privateEndpointRuleIdField
	}
	updateMaskField, err := fieldMaskFromPb(&pb.UpdateMask)
	if err != nil {
		return nil, err
	}
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	return st, nil
}

type UpdateNotificationDestinationRequest struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *Config
	// The display name for the notification destination.
	DisplayName string
	// UUID identifying notification destination.
	Id string

	ForceSendFields []string
}

func updateNotificationDestinationRequestToPb(st *UpdateNotificationDestinationRequest) (*updateNotificationDestinationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateNotificationDestinationRequestPb{}
	configPb, err := configToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type updateNotificationDestinationRequestPb struct {
	// The configuration for the notification destination. Must wrap EXACTLY one
	// of the nested configs.
	Config *configPb `json:"config,omitempty"`
	// The display name for the notification destination.
	DisplayName string `json:"display_name,omitempty"`
	// UUID identifying notification destination.
	Id string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateNotificationDestinationRequestFromPb(pb *updateNotificationDestinationRequestPb) (*UpdateNotificationDestinationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateNotificationDestinationRequest{}
	configField, err := configFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateNotificationDestinationRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateNotificationDestinationRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Details required to update a setting.
type UpdatePersonalComputeSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
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
	FieldMask string

	Setting PersonalComputeSetting
}

func updatePersonalComputeSettingRequestToPb(st *UpdatePersonalComputeSettingRequest) (*updatePersonalComputeSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePersonalComputeSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := personalComputeSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updatePersonalComputeSettingRequestPb struct {
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

	Setting personalComputeSettingPb `json:"setting"`
}

func updatePersonalComputeSettingRequestFromPb(pb *updatePersonalComputeSettingRequestPb) (*UpdatePersonalComputeSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePersonalComputeSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := personalComputeSettingFromPb(&pb.Setting)
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
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string
}

func updatePrivateEndpointRuleToPb(st *UpdatePrivateEndpointRule) (*updatePrivateEndpointRulePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePrivateEndpointRulePb{}

	var domainNamesPb []string
	for _, item := range st.DomainNames {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			domainNamesPb = append(domainNamesPb, *itemPb)
		}
	}
	pb.DomainNames = domainNamesPb

	return pb, nil
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

type updatePrivateEndpointRulePb struct {
	// Only used by private endpoints to customer-managed resources.
	//
	// Domain names of target private link service. When updating this field,
	// the full list of target domain_names must be specified.
	DomainNames []string `json:"domain_names,omitempty"`
}

func updatePrivateEndpointRuleFromPb(pb *updatePrivateEndpointRulePb) (*UpdatePrivateEndpointRule, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePrivateEndpointRule{}

	var domainNamesField []string
	for _, itemPb := range pb.DomainNames {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			domainNamesField = append(domainNamesField, *item)
		}
	}
	st.DomainNames = domainNamesField

	return st, nil
}

type UpdateResponse struct {
}

func updateResponseToPb(st *UpdateResponse) (*updateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateResponsePb{}

	return pb, nil
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

type updateResponsePb struct {
}

func updateResponseFromPb(pb *updateResponsePb) (*UpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateResponse{}

	return st, nil
}

// Details required to update a setting.
type UpdateRestrictWorkspaceAdminsSettingRequest struct {
	// This should always be set to true for Settings API. Added for AIP
	// compliance.
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
	FieldMask string

	Setting RestrictWorkspaceAdminsSetting
}

func updateRestrictWorkspaceAdminsSettingRequestToPb(st *UpdateRestrictWorkspaceAdminsSettingRequest) (*updateRestrictWorkspaceAdminsSettingRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRestrictWorkspaceAdminsSettingRequestPb{}
	allowMissingPb, err := identity(&st.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingPb != nil {
		pb.AllowMissing = *allowMissingPb
	}

	fieldMaskPb, err := identity(&st.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskPb != nil {
		pb.FieldMask = *fieldMaskPb
	}

	settingPb, err := restrictWorkspaceAdminsSettingToPb(&st.Setting)
	if err != nil {
		return nil, err
	}
	if settingPb != nil {
		pb.Setting = *settingPb
	}

	return pb, nil
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

type updateRestrictWorkspaceAdminsSettingRequestPb struct {
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

	Setting restrictWorkspaceAdminsSettingPb `json:"setting"`
}

func updateRestrictWorkspaceAdminsSettingRequestFromPb(pb *updateRestrictWorkspaceAdminsSettingRequestPb) (*UpdateRestrictWorkspaceAdminsSettingRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRestrictWorkspaceAdminsSettingRequest{}
	allowMissingField, err := identity(&pb.AllowMissing)
	if err != nil {
		return nil, err
	}
	if allowMissingField != nil {
		st.AllowMissing = *allowMissingField
	}
	fieldMaskField, err := identity(&pb.FieldMask)
	if err != nil {
		return nil, err
	}
	if fieldMaskField != nil {
		st.FieldMask = *fieldMaskField
	}
	settingField, err := restrictWorkspaceAdminsSettingFromPb(&pb.Setting)
	if err != nil {
		return nil, err
	}
	if settingField != nil {
		st.Setting = *settingField
	}

	return st, nil
}

type WorkspaceConf map[string]string
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
