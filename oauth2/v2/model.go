// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"github.com/databricks/databricks-sdk-go/databricks/marshal"
)

// Create account federation policy
type CreateAccountFederationPolicyRequest struct {
	Policy *FederationPolicy `json:"policy,omitempty"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId string `json:"-" url:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCustomAppIntegration struct {
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential bool `json:"confidential,omitempty"`
	// Name of the custom OAuth app
	Name string `json:"name,omitempty"`
	// List of OAuth redirect urls
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateCustomAppIntegrationOutput struct {
	// OAuth client-id generated by the Databricks
	ClientId string `json:"client_id,omitempty"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret string `json:"client_secret,omitempty"`
	// Unique integration id for the custom OAuth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePublishedAppIntegration struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId string `json:"app_id,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreatePublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Create service principal federation policy
type CreateServicePrincipalFederationPolicyRequest struct {
	Policy *FederationPolicy `json:"policy,omitempty"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId string `json:"-" url:"policy_id,omitempty"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateServicePrincipalSecretRequest struct {
	// The lifetime of the secret in seconds. If this parameter is not provided,
	// the secret will have a default lifetime of 730 days (63072000s).
	Lifetime string `json:"lifetime,omitempty"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
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
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	ExpireTime string `json:"expire_time,omitempty"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateServicePrincipalSecretResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServicePrincipalSecretResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete account federation policy
type DeleteAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
}

type DeleteCustomAppIntegrationOutput struct {
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	IntegrationId string `json:"-" url:"-"`
}

type DeletePublishedAppIntegrationOutput struct {
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	IntegrationId string `json:"-" url:"-"`
}

type DeleteResponse struct {
}

// Delete service principal federation policy
type DeleteServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId string `json:"-" url:"-"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

type FederationPolicy struct {
	// Creation time of the federation policy.
	CreateTime string `json:"create_time,omitempty"`
	// Description of the federation policy.
	Description string `json:"description,omitempty"`
	// Resource name for the federation policy. Example values include
	// `accounts/<account-id>/federationPolicies/my-federation-policy` for
	// Account Federation Policies, and
	// `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
	// for Service Principal Federation Policies. Typically an output parameter,
	// which does not need to be specified in create or update requests. If
	// specified in a request, must match the value in the request URL.
	Name string `json:"name,omitempty"`
	// Specifies the policy to use for validating OIDC claims in your federated
	// tokens.
	OidcPolicy *OidcFederationPolicy `json:"oidc_policy,omitempty"`
	// Unique, immutable id of the federation policy.
	Uid string `json:"uid,omitempty"`
	// Last update time of the federation policy.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *FederationPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FederationPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get account federation policy
type GetAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
}

type GetCustomAppIntegrationOutput struct {
	// The client id of the custom OAuth app
	ClientId string `json:"client_id,omitempty"`
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential bool `json:"confidential,omitempty"`

	CreateTime string `json:"create_time,omitempty"`

	CreatedBy int64 `json:"created_by,omitempty"`

	CreatorUsername string `json:"creator_username,omitempty"`
	// ID of this custom app
	IntegrationId string `json:"integration_id,omitempty"`
	// The display name of the custom OAuth app
	Name string `json:"name,omitempty"`
	// List of OAuth redirect urls
	RedirectUrls []string `json:"redirect_urls,omitempty"`

	Scopes []string `json:"scopes,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	// The OAuth app integration ID.
	IntegrationId string `json:"-" url:"-"`
}

type GetCustomAppIntegrationsOutput struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps []GetCustomAppIntegrationOutput `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetCustomAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetCustomAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedAppIntegrationOutput struct {
	// App-id of the published app integration
	AppId string `json:"app_id,omitempty"`

	CreateTime string `json:"create_time,omitempty"`

	CreatedBy int64 `json:"created_by,omitempty"`
	// Unique integration id for the published OAuth app
	IntegrationId string `json:"integration_id,omitempty"`
	// Display name of the published OAuth app
	Name string `json:"name,omitempty"`
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	IntegrationId string `json:"-" url:"-"`
}

type GetPublishedAppIntegrationsOutput struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps []GetPublishedAppIntegrationOutput `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPublishedAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	Apps []PublishedAppOutput `json:"apps,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetPublishedAppsOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetPublishedAppsOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get service principal federation policy
type GetServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

// List account federation policies
type ListAccountFederationPoliciesRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListAccountFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAccountFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get custom oauth app integrations
type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername bool `json:"-" url:"include_creator_username,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListCustomAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListCustomAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListFederationPoliciesResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Policies []FederationPolicy `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListFederationPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListFederationPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListOAuthPublishedAppsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListOAuthPublishedAppsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get published oauth app integrations
type ListPublishedAppIntegrationsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListPublishedAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListPublishedAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List service principal federation policies
type ListServicePrincipalFederationPoliciesRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// List service principal secrets
type ListServicePrincipalSecretsRequest struct {
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the secrets for this service principal.
	// Provide this token to retrieve the next page of secret entries. When
	// providing a `page_token`, all other parameters provided to the request
	// must match the previous request. To list all of the secrets for a service
	// principal, it is necessary to continue requesting pages of entries until
	// the response contains no `next_page_token`. Note that the number of
	// entries returned must not be used to determine when the listing is
	// complete.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalSecretsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalSecretsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListServicePrincipalSecretsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of the secrets
	Secrets []SecretInfo `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListServicePrincipalSecretsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListServicePrincipalSecretsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Specifies the policy to use for validating OIDC claims in your federated
// tokens.
type OidcFederationPolicy struct {
	// The allowed token audiences, as specified in the 'aud' claim of federated
	// tokens. The audience identifier is intended to represent the recipient of
	// the token. Can be any non-empty string value. As long as the audience in
	// the token matches at least one audience in the policy, the token is
	// considered a match. If audiences is unspecified, defaults to your
	// Databricks account id.
	Audiences []string `json:"audiences,omitempty"`
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	Issuer string `json:"issuer,omitempty"`
	// The public keys used to validate the signature of federated tokens, in
	// JWKS format. Most use cases should not need to specify this field. If
	// jwks_uri and jwks_json are both unspecified (recommended), Databricks
	// automatically fetches the public keys from your issuer’s well known
	// endpoint. Databricks strongly recommends relying on your issuer’s well
	// known endpoint for discovering public keys.
	JwksJson string `json:"jwks_json,omitempty"`
	// URL of the public keys used to validate the signature of federated
	// tokens, in JWKS format. Most use cases should not need to specify this
	// field. If jwks_uri and jwks_json are both unspecified (recommended),
	// Databricks automatically fetches the public keys from your issuer’s
	// well known endpoint. Databricks strongly recommends relying on your
	// issuer’s well known endpoint for discovering public keys.
	JwksUri string `json:"jwks_uri,omitempty"`
	// The required token subject, as specified in the subject claim of
	// federated tokens. Must be specified for service principal federation
	// policies. Must not be specified for account federation policies.
	Subject string `json:"subject,omitempty"`
	// The claim that contains the subject of the token. If unspecified, the
	// default value is 'sub'.
	SubjectClaim string `json:"subject_claim,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *OidcFederationPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OidcFederationPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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
	// The display name of the published OAuth app.
	Name string `json:"name,omitempty"`
	// Redirect URLs of the published OAuth app.
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Required scopes for the published OAuth app.
	Scopes []string `json:"scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *PublishedAppOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PublishedAppOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SecretInfo struct {
	// UTC time when the secret was created
	CreateTime string `json:"create_time,omitempty"`
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	ExpireTime string `json:"expire_time,omitempty"`
	// ID of the secret
	Id string `json:"id,omitempty"`
	// Secret Hash
	SecretHash string `json:"secret_hash,omitempty"`
	// Status of the secret
	Status string `json:"status,omitempty"`
	// UTC time when the secret was updated
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TokenAccessPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TokenAccessPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Update account federation policy
type UpdateAccountFederationPolicyRequest struct {
	Policy *FederationPolicy `json:"policy,omitempty"`
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask string `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateCustomAppIntegration struct {
	IntegrationId string `json:"-" url:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// List of OAuth scopes to be updated in the custom OAuth app integration,
	// similar to redirect URIs this will fully replace the existing values
	// instead of appending
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`
}

type UpdateCustomAppIntegrationOutput struct {
}

type UpdatePublishedAppIntegration struct {
	IntegrationId string `json:"-" url:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}

type UpdatePublishedAppIntegrationOutput struct {
}

// Update service principal federation policy
type UpdateServicePrincipalFederationPolicyRequest struct {
	Policy *FederationPolicy `json:"policy,omitempty"`
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask string `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
