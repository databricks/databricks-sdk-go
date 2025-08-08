// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2pb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateAccountFederationPolicyRequestPb struct {
	Policy   FederationPolicyPb `json:"policy"`
	PolicyId string             `json:"-" url:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateAccountFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateAccountFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCustomAppIntegrationPb struct {
	Confidential         bool                 `json:"confidential,omitempty"`
	Name                 string               `json:"name,omitempty"`
	RedirectUrls         []string             `json:"redirect_urls,omitempty"`
	Scopes               []string             `json:"scopes,omitempty"`
	TokenAccessPolicy    *TokenAccessPolicyPb `json:"token_access_policy,omitempty"`
	UserAuthorizedScopes []string             `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCustomAppIntegrationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCustomAppIntegrationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCustomAppIntegrationOutputPb struct {
	ClientId      string `json:"client_id,omitempty"`
	ClientSecret  string `json:"client_secret,omitempty"`
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCustomAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCustomAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePublishedAppIntegrationPb struct {
	AppId             string               `json:"app_id,omitempty"`
	TokenAccessPolicy *TokenAccessPolicyPb `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePublishedAppIntegrationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePublishedAppIntegrationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePublishedAppIntegrationOutputPb struct {
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePublishedAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePublishedAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateServicePrincipalFederationPolicyRequestPb struct {
	Policy             FederationPolicyPb `json:"policy"`
	PolicyId           string             `json:"-" url:"policy_id,omitempty"`
	ServicePrincipalId int64              `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateServicePrincipalFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateServicePrincipalFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateServicePrincipalSecretRequestPb struct {
	Lifetime           string `json:"lifetime,omitempty"`
	ServicePrincipalId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateServicePrincipalSecretRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateServicePrincipalSecretRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateServicePrincipalSecretResponsePb struct {
	CreateTime string `json:"create_time,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	Id         string `json:"id,omitempty"`
	Secret     string `json:"secret,omitempty"`
	SecretHash string `json:"secret_hash,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateServicePrincipalSecretResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateServicePrincipalSecretResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteAccountFederationPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

type DeleteCustomAppIntegrationOutputPb struct {
}

type DeleteCustomAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

type DeletePublishedAppIntegrationOutputPb struct {
}

type DeletePublishedAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

type DeleteServicePrincipalFederationPolicyRequestPb struct {
	PolicyId           string `json:"-" url:"-"`
	ServicePrincipalId int64  `json:"-" url:"-"`
}

type DeleteServicePrincipalSecretRequestPb struct {
	SecretId           string `json:"-" url:"-"`
	ServicePrincipalId string `json:"-" url:"-"`
}

type FederationPolicyPb struct {
	CreateTime         string                  `json:"create_time,omitempty"`
	Description        string                  `json:"description,omitempty"`
	Name               string                  `json:"name,omitempty"`
	OidcPolicy         *OidcFederationPolicyPb `json:"oidc_policy,omitempty"`
	PolicyId           string                  `json:"policy_id,omitempty"`
	ServicePrincipalId int64                   `json:"service_principal_id,omitempty"`
	Uid                string                  `json:"uid,omitempty"`
	UpdateTime         string                  `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FederationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FederationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetAccountFederationPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

type GetCustomAppIntegrationOutputPb struct {
	ClientId             string               `json:"client_id,omitempty"`
	Confidential         bool                 `json:"confidential,omitempty"`
	CreateTime           string               `json:"create_time,omitempty"`
	CreatedBy            int64                `json:"created_by,omitempty"`
	CreatorUsername      string               `json:"creator_username,omitempty"`
	IntegrationId        string               `json:"integration_id,omitempty"`
	Name                 string               `json:"name,omitempty"`
	RedirectUrls         []string             `json:"redirect_urls,omitempty"`
	Scopes               []string             `json:"scopes,omitempty"`
	TokenAccessPolicy    *TokenAccessPolicyPb `json:"token_access_policy,omitempty"`
	UserAuthorizedScopes []string             `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetCustomAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetCustomAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetCustomAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

type GetCustomAppIntegrationsOutputPb struct {
	Apps          []GetCustomAppIntegrationOutputPb `json:"apps,omitempty"`
	NextPageToken string                            `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetCustomAppIntegrationsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetCustomAppIntegrationsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPublishedAppIntegrationOutputPb struct {
	AppId             string               `json:"app_id,omitempty"`
	CreateTime        string               `json:"create_time,omitempty"`
	CreatedBy         int64                `json:"created_by,omitempty"`
	IntegrationId     string               `json:"integration_id,omitempty"`
	Name              string               `json:"name,omitempty"`
	TokenAccessPolicy *TokenAccessPolicyPb `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPublishedAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPublishedAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPublishedAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

type GetPublishedAppIntegrationsOutputPb struct {
	Apps          []GetPublishedAppIntegrationOutputPb `json:"apps,omitempty"`
	NextPageToken string                               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPublishedAppIntegrationsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPublishedAppIntegrationsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPublishedAppsOutputPb struct {
	Apps          []PublishedAppOutputPb `json:"apps,omitempty"`
	NextPageToken string                 `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPublishedAppsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPublishedAppsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetServicePrincipalFederationPolicyRequestPb struct {
	PolicyId           string `json:"-" url:"-"`
	ServicePrincipalId int64  `json:"-" url:"-"`
}

type ListAccountFederationPoliciesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListAccountFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListAccountFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListCustomAppIntegrationsRequestPb struct {
	IncludeCreatorUsername bool   `json:"-" url:"include_creator_username,omitempty"`
	PageSize               int    `json:"-" url:"page_size,omitempty"`
	PageToken              string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListCustomAppIntegrationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListCustomAppIntegrationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFederationPoliciesResponsePb struct {
	NextPageToken string               `json:"next_page_token,omitempty"`
	Policies      []FederationPolicyPb `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListFederationPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListFederationPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListOAuthPublishedAppsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListOAuthPublishedAppsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListOAuthPublishedAppsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPublishedAppIntegrationsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPublishedAppIntegrationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPublishedAppIntegrationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListServicePrincipalFederationPoliciesRequestPb struct {
	PageSize           int    `json:"-" url:"page_size,omitempty"`
	PageToken          string `json:"-" url:"page_token,omitempty"`
	ServicePrincipalId int64  `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListServicePrincipalFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListServicePrincipalFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListServicePrincipalSecretsRequestPb struct {
	PageSize           int    `json:"-" url:"page_size,omitempty"`
	PageToken          string `json:"-" url:"page_token,omitempty"`
	ServicePrincipalId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListServicePrincipalSecretsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListServicePrincipalSecretsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListServicePrincipalSecretsResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	Secrets       []SecretInfoPb `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListServicePrincipalSecretsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListServicePrincipalSecretsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type OidcFederationPolicyPb struct {
	Audiences    []string `json:"audiences,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	JwksJson     string   `json:"jwks_json,omitempty"`
	JwksUri      string   `json:"jwks_uri,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	SubjectClaim string   `json:"subject_claim,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OidcFederationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OidcFederationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PublishedAppOutputPb struct {
	AppId                string   `json:"app_id,omitempty"`
	ClientId             string   `json:"client_id,omitempty"`
	Description          string   `json:"description,omitempty"`
	IsConfidentialClient bool     `json:"is_confidential_client,omitempty"`
	Name                 string   `json:"name,omitempty"`
	RedirectUrls         []string `json:"redirect_urls,omitempty"`
	Scopes               []string `json:"scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PublishedAppOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PublishedAppOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SecretInfoPb struct {
	CreateTime string `json:"create_time,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	Id         string `json:"id,omitempty"`
	SecretHash string `json:"secret_hash,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SecretInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SecretInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenAccessPolicyPb struct {
	AccessTokenTtlInMinutes  int `json:"access_token_ttl_in_minutes,omitempty"`
	RefreshTokenTtlInMinutes int `json:"refresh_token_ttl_in_minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TokenAccessPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TokenAccessPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateAccountFederationPolicyRequestPb struct {
	Policy     FederationPolicyPb `json:"policy"`
	PolicyId   string             `json:"-" url:"-"`
	UpdateMask string             `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateAccountFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateAccountFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCustomAppIntegrationPb struct {
	IntegrationId        string               `json:"-" url:"-"`
	RedirectUrls         []string             `json:"redirect_urls,omitempty"`
	Scopes               []string             `json:"scopes,omitempty"`
	TokenAccessPolicy    *TokenAccessPolicyPb `json:"token_access_policy,omitempty"`
	UserAuthorizedScopes []string             `json:"user_authorized_scopes,omitempty"`
}

type UpdateCustomAppIntegrationOutputPb struct {
}

type UpdatePublishedAppIntegrationPb struct {
	IntegrationId     string               `json:"-" url:"-"`
	TokenAccessPolicy *TokenAccessPolicyPb `json:"token_access_policy,omitempty"`
}

type UpdatePublishedAppIntegrationOutputPb struct {
}

type UpdateServicePrincipalFederationPolicyRequestPb struct {
	Policy             FederationPolicyPb `json:"policy"`
	PolicyId           string             `json:"-" url:"-"`
	ServicePrincipalId int64              `json:"-" url:"-"`
	UpdateMask         string             `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateServicePrincipalFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateServicePrincipalFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
