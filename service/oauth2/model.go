// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

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
}

type CreateCustomAppIntegrationOutput struct {
	// oauth client-id generated by the Databricks
	ClientId string `json:"client_id,omitempty"`
	// oauth client-secret generated by the Databricks if this is a confidential
	// oauth app client-secret will be generated.
	ClientSecret string `json:"client_secret,omitempty"`
	// unique integration id for the custom oauth app
	IntegrationId string `json:"integration_id,omitempty"`
}

type CreateOAuthEnrollment struct {
	// If true, enable OAuth for all the published applications in the account.
	EnableAllPublishedApps bool `json:"enable_all_published_apps,omitempty"`
}

type CreatePublishedAppIntegration struct {
	// app_id of the oauth published app integration. For example power-bi,
	// tableau-deskop
	AppId string `json:"app_id,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}

type CreatePublishedAppIntegrationOutput struct {
	// unique integration id for the published oauth app
	IntegrationId string `json:"integration_id,omitempty"`
}

// Create service principal secret
type CreateServicePrincipalSecretRequest struct {
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
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
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId string `json:"-" url:"-"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
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
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
}

type GetCustomAppIntegrationsOutput struct {
	// Array of Custom OAuth App Integrations defined for the account.
	Apps []GetCustomAppIntegrationOutput `json:"apps,omitempty"`
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
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
}

type GetPublishedAppIntegrationsOutput struct {
	// Array of Published OAuth App Integrations defined for the account.
	Apps []GetPublishedAppIntegrationOutput `json:"apps,omitempty"`
}

type GetPublishedAppsOutput struct {
	// Array of Published OAuth Apps.
	Apps []PublishedAppOutput `json:"apps,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return.
	PageSize int64 `json:"-" url:"page_size,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`
}

// List service principal secrets
type ListServicePrincipalSecretsRequest struct {
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

type ListServicePrincipalSecretsResponse struct {
	// List of the secrets
	Secrets []SecretInfo `json:"secrets,omitempty"`
}

type OAuthEnrollmentStatus struct {
	// Is OAuth enrolled for the account.
	IsEnabled bool `json:"is_enabled,omitempty"`
}

type PublishedAppOutput struct {
	// Unique ID of the published OAuth app.
	AppId string `json:"app_id,omitempty"`
	// Client ID of the published OAuth app. It is the client_id in the OAuth
	// flow
	ClientId string `json:"client_id,omitempty"`
	// Description of the published OAuth app.
	Description string `json:"description,omitempty"`
	// Whether the published OAuth app is a confidential client. It is always
	// false for published OAuth apps.
	IsConfidentialClient bool `json:"is_confidential_client,omitempty"`
	// Name of the published OAuth app.
	Name string `json:"name,omitempty"`
	// Redirect URLs of the published OAuth app.
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Required scopes for the published OAuth app.
	Scopes []string `json:"scopes,omitempty"`
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
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes int `json:"access_token_ttl_in_minutes,omitempty"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes int `json:"refresh_token_ttl_in_minutes,omitempty"`
}

type UpdateCustomAppIntegration struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
	// List of oauth redirect urls to be updated in the custom oauth app
	// integration
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Token access policy to be updated in the custom oauth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}

type UpdatePublishedAppIntegration struct {
	// The oauth app integration ID.
	IntegrationId string `json:"-" url:"-"`
	// Token access policy to be updated in the published oauth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}
