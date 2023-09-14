// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

type CreateCustomAppIntegration struct {
	// indicates if an oauth client-secret should be generated
	Confidential bool `json:"confidential,omitempty"`
	// name of the custom oauth app
	Name string `json:"name"`
	// List of oauth redirect urls
	RedirectUrls []string `json:"redirect_urls"`
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCustomAppIntegrationOutput struct {
	// oauth client-id generated by the Databricks
	ClientId string `json:"client_id,omitempty"`
	// oauth client-secret generated by the Databricks if this is a confidential
	// oauth app client-secret will be generated.
	ClientSecret string `json:"client_secret,omitempty"`
	// unique integration id for the custom oauth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateOAuthEnrollment struct {
	// If true, enable OAuth for all the published applications in the account.
	EnableAllPublishedApps bool `json:"enable_all_published_apps,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateOAuthEnrollment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateOAuthEnrollment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePublishedAppIntegration struct {
	// app_id of the oauth published app integration. For example power-bi,
	// tableau-deskop
	AppId string `json:"app_id,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePublishedAppIntegrationOutput struct {
	// unique integration id for the published oauth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreatePublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Create service principal secret
type CreateServicePrincipalSecretRequest struct {
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateServicePrincipalSecretRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServicePrincipalSecretRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateServicePrincipalSecretResponse struct {
	// UTC time when the secret was created
	CreateTime string `json:"create_time,omitempty"`
	// ID of the secret
	Id string `json:"id,omitempty"`
	// Secret Value
	Secret string `json:"secret,omitempty"`
	// Secret Hash
	SecretHash string `json:"secret_hash,omitempty"`
	// Status of the secret
	Status string `json:"status,omitempty"`
	// UTC time when the secret was updated
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateServicePrincipalSecretResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServicePrincipalSecretResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteCustomAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteCustomAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeletePublishedAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeletePublishedAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId string `json:"-" url:"-"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteServicePrincipalSecretRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteServicePrincipalSecretRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetCustomAppIntegrationOutput struct {
	// oauth client id of the custom oauth app
	ClientId string `json:"client_id,omitempty"`
	// indicates if an oauth client-secret should be generated
	Confidential bool `json:"confidential,omitempty"`
	// ID of this custom app
	IntegrationId string `json:"integration_id,omitempty"`
	// name of the custom oauth app
	Name string `json:"name,omitempty"`
	// List of oauth redirect urls
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetCustomAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCustomAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetCustomAppIntegrationsOutput struct {
	// Array of Custom OAuth App Integrations defined for the account.
	Apps []GetCustomAppIntegrationOutput `json:"apps,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetCustomAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCustomAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedAppIntegrationOutput struct {
	// app-id of the published app integration
	AppId string `json:"app_id,omitempty"`
	// unique integration id for the published oauth app
	IntegrationId string `json:"integration_id,omitempty"`
	// name of the published oauth app
	Name string `json:"name,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPublishedAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedAppIntegrationsOutput struct {
	// Array of Published OAuth App Integrations defined for the account.
	Apps []GetPublishedAppIntegrationOutput `json:"apps,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetPublishedAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List service principal secrets
type ListServicePrincipalSecretsRequest struct {
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-"`
}

func (s *ListServicePrincipalSecretsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalSecretsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListServicePrincipalSecretsResponse struct {
	// List of the secrets
	Secrets []SecretInfo `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListServicePrincipalSecretsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalSecretsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type OAuthEnrollmentStatus struct {
	// Is OAuth enrolled for the account.
	IsEnabled bool `json:"is_enabled,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *OAuthEnrollmentStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OAuthEnrollmentStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SecretInfo struct {
	// UTC time when the secret was created
	CreateTime string `json:"create_time,omitempty"`
	// ID of the secret
	Id string `json:"id,omitempty"`
	// Secret Hash
	SecretHash string `json:"secret_hash,omitempty"`
	// Status of the secret
	Status string `json:"status,omitempty"`
	// UTC time when the secret was updated
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *SecretInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SecretInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes int `json:"access_token_ttl_in_minutes,omitempty"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes int `json:"refresh_token_ttl_in_minutes,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *TokenAccessPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCustomAppIntegration struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
	// List of oauth redirect urls to be updated in the custom oauth app
	// integration
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Token access policy to be updated in the custom oauth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdatePublishedAppIntegration struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
	// Token access policy to be updated in the published oauth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UpdatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
