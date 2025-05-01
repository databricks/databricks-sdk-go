// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

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

// Create account federation policy
type CreateAccountFederationPolicyRequest struct {
	Policy FederationPolicy
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId string

	ForceSendFields []string
}

func createAccountFederationPolicyRequestToPb(st *CreateAccountFederationPolicyRequest) (*createAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAccountFederationPolicyRequestPb{}
	policyPb, err := federationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}

	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createAccountFederationPolicyRequestPb struct {
	Policy federationPolicyPb `json:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId string `json:"-" url:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAccountFederationPolicyRequestFromPb(pb *createAccountFederationPolicyRequestPb) (*CreateAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAccountFederationPolicyRequest{}
	policyField, err := federationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAccountFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAccountFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCustomAppIntegration struct {
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential bool
	// Name of the custom OAuth app
	Name string
	// List of OAuth redirect urls
	RedirectUrls []string
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	Scopes []string
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string

	ForceSendFields []string
}

func createCustomAppIntegrationToPb(st *CreateCustomAppIntegration) (*createCustomAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCustomAppIntegrationPb{}
	confidentialPb := &st.Confidential
	if confidentialPb != nil {
		pb.Confidential = *confidentialPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	var redirectUrlsPb []string
	for _, item := range st.RedirectUrls {
		itemPb := &item
		if itemPb != nil {
			redirectUrlsPb = append(redirectUrlsPb, *itemPb)
		}
	}
	pb.RedirectUrls = redirectUrlsPb

	var scopesPb []string
	for _, item := range st.Scopes {
		itemPb := &item
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

	tokenAccessPolicyPb, err := tokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	var userAuthorizedScopesPb []string
	for _, item := range st.UserAuthorizedScopes {
		itemPb := &item
		if itemPb != nil {
			userAuthorizedScopesPb = append(userAuthorizedScopesPb, *itemPb)
		}
	}
	pb.UserAuthorizedScopes = userAuthorizedScopesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCustomAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCustomAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := createCustomAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCustomAppIntegrationPb struct {
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
	TokenAccessPolicy *tokenAccessPolicyPb `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCustomAppIntegrationFromPb(pb *createCustomAppIntegrationPb) (*CreateCustomAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomAppIntegration{}
	confidentialField := &pb.Confidential
	if confidentialField != nil {
		st.Confidential = *confidentialField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	var redirectUrlsField []string
	for _, item := range pb.RedirectUrls {
		itemField := &item
		if itemField != nil {
			redirectUrlsField = append(redirectUrlsField, *itemField)
		}
	}
	st.RedirectUrls = redirectUrlsField

	var scopesField []string
	for _, item := range pb.Scopes {
		itemField := &item
		if itemField != nil {
			scopesField = append(scopesField, *itemField)
		}
	}
	st.Scopes = scopesField
	tokenAccessPolicyField, err := tokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	var userAuthorizedScopesField []string
	for _, item := range pb.UserAuthorizedScopes {
		itemField := &item
		if itemField != nil {
			userAuthorizedScopesField = append(userAuthorizedScopesField, *itemField)
		}
	}
	st.UserAuthorizedScopes = userAuthorizedScopesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCustomAppIntegrationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCustomAppIntegrationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateCustomAppIntegrationOutput struct {
	// OAuth client-id generated by the Databricks
	ClientId string
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret string
	// Unique integration id for the custom OAuth app
	IntegrationId string

	ForceSendFields []string
}

func createCustomAppIntegrationOutputToPb(st *CreateCustomAppIntegrationOutput) (*createCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCustomAppIntegrationOutputPb{}
	clientIdPb := &st.ClientId
	if clientIdPb != nil {
		pb.ClientId = *clientIdPb
	}

	clientSecretPb := &st.ClientSecret
	if clientSecretPb != nil {
		pb.ClientSecret = *clientSecretPb
	}

	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCustomAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCustomAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := createCustomAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCustomAppIntegrationOutputPb struct {
	// OAuth client-id generated by the Databricks
	ClientId string `json:"client_id,omitempty"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	ClientSecret string `json:"client_secret,omitempty"`
	// Unique integration id for the custom OAuth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCustomAppIntegrationOutputFromPb(pb *createCustomAppIntegrationOutputPb) (*CreateCustomAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomAppIntegrationOutput{}
	clientIdField := &pb.ClientId
	if clientIdField != nil {
		st.ClientId = *clientIdField
	}
	clientSecretField := &pb.ClientSecret
	if clientSecretField != nil {
		st.ClientSecret = *clientSecretField
	}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCustomAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCustomAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePublishedAppIntegration struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId string
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy

	ForceSendFields []string
}

func createPublishedAppIntegrationToPb(st *CreatePublishedAppIntegration) (*createPublishedAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPublishedAppIntegrationPb{}
	appIdPb := &st.AppId
	if appIdPb != nil {
		pb.AppId = *appIdPb
	}

	tokenAccessPolicyPb, err := tokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPublishedAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPublishedAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := createPublishedAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createPublishedAppIntegrationPb struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	AppId string `json:"app_id,omitempty"`
	// Token access policy
	TokenAccessPolicy *tokenAccessPolicyPb `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPublishedAppIntegrationFromPb(pb *createPublishedAppIntegrationPb) (*CreatePublishedAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePublishedAppIntegration{}
	appIdField := &pb.AppId
	if appIdField != nil {
		st.AppId = *appIdField
	}
	tokenAccessPolicyField, err := tokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPublishedAppIntegrationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPublishedAppIntegrationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	IntegrationId string

	ForceSendFields []string
}

func createPublishedAppIntegrationOutputToPb(st *CreatePublishedAppIntegrationOutput) (*createPublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPublishedAppIntegrationOutputPb{}
	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreatePublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPublishedAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPublishedAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := createPublishedAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createPublishedAppIntegrationOutputPb struct {
	// Unique integration id for the published OAuth app
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPublishedAppIntegrationOutputFromPb(pb *createPublishedAppIntegrationOutputPb) (*CreatePublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePublishedAppIntegrationOutput{}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPublishedAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPublishedAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Create service principal federation policy
type CreateServicePrincipalFederationPolicyRequest struct {
	Policy FederationPolicy
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId string
	// The service principal id for the federation policy.
	ServicePrincipalId int64

	ForceSendFields []string
}

func createServicePrincipalFederationPolicyRequestToPb(st *CreateServicePrincipalFederationPolicyRequest) (*createServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServicePrincipalFederationPolicyRequestPb{}
	policyPb, err := federationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}

	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := createServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createServicePrincipalFederationPolicyRequestPb struct {
	Policy federationPolicyPb `json:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId string `json:"-" url:"policy_id,omitempty"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createServicePrincipalFederationPolicyRequestFromPb(pb *createServicePrincipalFederationPolicyRequestPb) (*CreateServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalFederationPolicyRequest{}
	policyField, err := federationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createServicePrincipalFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServicePrincipalFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateServicePrincipalSecretRequest struct {
	// The lifetime of the secret in seconds. If this parameter is not provided,
	// the secret will have a default lifetime of 730 days (63072000s).
	Lifetime string
	// The service principal ID.
	ServicePrincipalId int64

	ForceSendFields []string
}

func createServicePrincipalSecretRequestToPb(st *CreateServicePrincipalSecretRequest) (*createServicePrincipalSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServicePrincipalSecretRequestPb{}
	lifetimePb := &st.Lifetime
	if lifetimePb != nil {
		pb.Lifetime = *lifetimePb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateServicePrincipalSecretRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createServicePrincipalSecretRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createServicePrincipalSecretRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateServicePrincipalSecretRequest) MarshalJSON() ([]byte, error) {
	pb, err := createServicePrincipalSecretRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createServicePrincipalSecretRequestPb struct {
	// The lifetime of the secret in seconds. If this parameter is not provided,
	// the secret will have a default lifetime of 730 days (63072000s).
	Lifetime string `json:"lifetime,omitempty"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createServicePrincipalSecretRequestFromPb(pb *createServicePrincipalSecretRequestPb) (*CreateServicePrincipalSecretRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalSecretRequest{}
	lifetimeField := &pb.Lifetime
	if lifetimeField != nil {
		st.Lifetime = *lifetimeField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createServicePrincipalSecretRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServicePrincipalSecretRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateServicePrincipalSecretResponse struct {
	// UTC time when the secret was created
	CreateTime string
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	ExpireTime string
	// ID of the secret
	Id string
	// Secret Value
	Secret string
	// Secret Hash
	SecretHash string
	// Status of the secret
	Status string
	// UTC time when the secret was updated
	UpdateTime string

	ForceSendFields []string
}

func createServicePrincipalSecretResponseToPb(st *CreateServicePrincipalSecretResponse) (*createServicePrincipalSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServicePrincipalSecretResponsePb{}
	createTimePb := &st.CreateTime
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	expireTimePb := &st.ExpireTime
	if expireTimePb != nil {
		pb.ExpireTime = *expireTimePb
	}

	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

	secretPb := &st.Secret
	if secretPb != nil {
		pb.Secret = *secretPb
	}

	secretHashPb := &st.SecretHash
	if secretHashPb != nil {
		pb.SecretHash = *secretHashPb
	}

	statusPb := &st.Status
	if statusPb != nil {
		pb.Status = *statusPb
	}

	updateTimePb := &st.UpdateTime
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateServicePrincipalSecretResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createServicePrincipalSecretResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createServicePrincipalSecretResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateServicePrincipalSecretResponse) MarshalJSON() ([]byte, error) {
	pb, err := createServicePrincipalSecretResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createServicePrincipalSecretResponsePb struct {
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

func createServicePrincipalSecretResponseFromPb(pb *createServicePrincipalSecretResponsePb) (*CreateServicePrincipalSecretResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalSecretResponse{}
	createTimeField := &pb.CreateTime
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	expireTimeField := &pb.ExpireTime
	if expireTimeField != nil {
		st.ExpireTime = *expireTimeField
	}
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}
	secretField := &pb.Secret
	if secretField != nil {
		st.Secret = *secretField
	}
	secretHashField := &pb.SecretHash
	if secretHashField != nil {
		st.SecretHash = *secretHashField
	}
	statusField := &pb.Status
	if statusField != nil {
		st.Status = *statusField
	}
	updateTimeField := &pb.UpdateTime
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createServicePrincipalSecretResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServicePrincipalSecretResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete account federation policy
type DeleteAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string
}

func deleteAccountFederationPolicyRequestToPb(st *DeleteAccountFederationPolicyRequest) (*deleteAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountFederationPolicyRequestPb{}
	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	return pb, nil
}

func (st *DeleteAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteAccountFederationPolicyRequestPb struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
}

func deleteAccountFederationPolicyRequestFromPb(pb *deleteAccountFederationPolicyRequestPb) (*DeleteAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountFederationPolicyRequest{}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	return st, nil
}

type DeleteCustomAppIntegrationOutput struct {
}

func deleteCustomAppIntegrationOutputToPb(st *DeleteCustomAppIntegrationOutput) (*deleteCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCustomAppIntegrationOutputPb{}

	return pb, nil
}

func (st *DeleteCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCustomAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCustomAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := deleteCustomAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCustomAppIntegrationOutputPb struct {
}

func deleteCustomAppIntegrationOutputFromPb(pb *deleteCustomAppIntegrationOutputPb) (*DeleteCustomAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCustomAppIntegrationOutput{}

	return st, nil
}

// Delete Custom OAuth App Integration
type DeleteCustomAppIntegrationRequest struct {
	IntegrationId string
}

func deleteCustomAppIntegrationRequestToPb(st *DeleteCustomAppIntegrationRequest) (*deleteCustomAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCustomAppIntegrationRequestPb{}
	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	return pb, nil
}

func (st *DeleteCustomAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCustomAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCustomAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCustomAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCustomAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCustomAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

func deleteCustomAppIntegrationRequestFromPb(pb *deleteCustomAppIntegrationRequestPb) (*DeleteCustomAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCustomAppIntegrationRequest{}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}

	return st, nil
}

type DeletePublishedAppIntegrationOutput struct {
}

func deletePublishedAppIntegrationOutputToPb(st *DeletePublishedAppIntegrationOutput) (*deletePublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePublishedAppIntegrationOutputPb{}

	return pb, nil
}

func (st *DeletePublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePublishedAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePublishedAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := deletePublishedAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deletePublishedAppIntegrationOutputPb struct {
}

func deletePublishedAppIntegrationOutputFromPb(pb *deletePublishedAppIntegrationOutputPb) (*DeletePublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePublishedAppIntegrationOutput{}

	return st, nil
}

// Delete Published OAuth App Integration
type DeletePublishedAppIntegrationRequest struct {
	IntegrationId string
}

func deletePublishedAppIntegrationRequestToPb(st *DeletePublishedAppIntegrationRequest) (*deletePublishedAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePublishedAppIntegrationRequestPb{}
	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	return pb, nil
}

func (st *DeletePublishedAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePublishedAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deletePublishedAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeletePublishedAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := deletePublishedAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deletePublishedAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

func deletePublishedAppIntegrationRequestFromPb(pb *deletePublishedAppIntegrationRequestPb) (*DeletePublishedAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePublishedAppIntegrationRequest{}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
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

// Delete service principal federation policy
type DeleteServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string
	// The service principal id for the federation policy.
	ServicePrincipalId int64
}

func deleteServicePrincipalFederationPolicyRequestToPb(st *DeleteServicePrincipalFederationPolicyRequest) (*deleteServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServicePrincipalFederationPolicyRequestPb{}
	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	return pb, nil
}

func (st *DeleteServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteServicePrincipalFederationPolicyRequestPb struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

func deleteServicePrincipalFederationPolicyRequestFromPb(pb *deleteServicePrincipalFederationPolicyRequestPb) (*DeleteServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalFederationPolicyRequest{}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}

	return st, nil
}

// Delete service principal secret
type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId string
	// The service principal ID.
	ServicePrincipalId int64
}

func deleteServicePrincipalSecretRequestToPb(st *DeleteServicePrincipalSecretRequest) (*deleteServicePrincipalSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServicePrincipalSecretRequestPb{}
	secretIdPb := &st.SecretId
	if secretIdPb != nil {
		pb.SecretId = *secretIdPb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	return pb, nil
}

func (st *DeleteServicePrincipalSecretRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteServicePrincipalSecretRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteServicePrincipalSecretRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteServicePrincipalSecretRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteServicePrincipalSecretRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteServicePrincipalSecretRequestPb struct {
	// The secret ID.
	SecretId string `json:"-" url:"-"`
	// The service principal ID.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

func deleteServicePrincipalSecretRequestFromPb(pb *deleteServicePrincipalSecretRequestPb) (*DeleteServicePrincipalSecretRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalSecretRequest{}
	secretIdField := &pb.SecretId
	if secretIdField != nil {
		st.SecretId = *secretIdField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}

	return st, nil
}

type FederationPolicy struct {
	// Creation time of the federation policy.
	CreateTime string
	// Description of the federation policy.
	Description string
	// Resource name for the federation policy. Example values include
	// `accounts/<account-id>/federationPolicies/my-federation-policy` for
	// Account Federation Policies, and
	// `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
	// for Service Principal Federation Policies. Typically an output parameter,
	// which does not need to be specified in create or update requests. If
	// specified in a request, must match the value in the request URL.
	Name string
	// Specifies the policy to use for validating OIDC claims in your federated
	// tokens.
	OidcPolicy *OidcFederationPolicy
	// Unique, immutable id of the federation policy.
	Uid string
	// Last update time of the federation policy.
	UpdateTime string

	ForceSendFields []string
}

func federationPolicyToPb(st *FederationPolicy) (*federationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &federationPolicyPb{}
	createTimePb := &st.CreateTime
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	descriptionPb := &st.Description
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	oidcPolicyPb, err := oidcFederationPolicyToPb(st.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyPb != nil {
		pb.OidcPolicy = oidcPolicyPb
	}

	uidPb := &st.Uid
	if uidPb != nil {
		pb.Uid = *uidPb
	}

	updateTimePb := &st.UpdateTime
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &federationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := federationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := federationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type federationPolicyPb struct {
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
	OidcPolicy *oidcFederationPolicyPb `json:"oidc_policy,omitempty"`
	// Unique, immutable id of the federation policy.
	Uid string `json:"uid,omitempty"`
	// Last update time of the federation policy.
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func federationPolicyFromPb(pb *federationPolicyPb) (*FederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FederationPolicy{}
	createTimeField := &pb.CreateTime
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	descriptionField := &pb.Description
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	oidcPolicyField, err := oidcFederationPolicyFromPb(pb.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyField != nil {
		st.OidcPolicy = oidcPolicyField
	}
	uidField := &pb.Uid
	if uidField != nil {
		st.Uid = *uidField
	}
	updateTimeField := &pb.UpdateTime
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *federationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st federationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get account federation policy
type GetAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string
}

func getAccountFederationPolicyRequestToPb(st *GetAccountFederationPolicyRequest) (*getAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountFederationPolicyRequestPb{}
	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	return pb, nil
}

func (st *GetAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getAccountFederationPolicyRequestPb struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
}

func getAccountFederationPolicyRequestFromPb(pb *getAccountFederationPolicyRequestPb) (*GetAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountFederationPolicyRequest{}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}

	return st, nil
}

type GetCustomAppIntegrationOutput struct {
	// The client id of the custom OAuth app
	ClientId string
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	Confidential bool

	CreateTime string

	CreatedBy int64

	CreatorUsername string
	// ID of this custom app
	IntegrationId string
	// The display name of the custom OAuth app
	Name string
	// List of OAuth redirect urls
	RedirectUrls []string

	Scopes []string
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string

	ForceSendFields []string
}

func getCustomAppIntegrationOutputToPb(st *GetCustomAppIntegrationOutput) (*getCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCustomAppIntegrationOutputPb{}
	clientIdPb := &st.ClientId
	if clientIdPb != nil {
		pb.ClientId = *clientIdPb
	}

	confidentialPb := &st.Confidential
	if confidentialPb != nil {
		pb.Confidential = *confidentialPb
	}

	createTimePb := &st.CreateTime
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	creatorUsernamePb := &st.CreatorUsername
	if creatorUsernamePb != nil {
		pb.CreatorUsername = *creatorUsernamePb
	}

	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	var redirectUrlsPb []string
	for _, item := range st.RedirectUrls {
		itemPb := &item
		if itemPb != nil {
			redirectUrlsPb = append(redirectUrlsPb, *itemPb)
		}
	}
	pb.RedirectUrls = redirectUrlsPb

	var scopesPb []string
	for _, item := range st.Scopes {
		itemPb := &item
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

	tokenAccessPolicyPb, err := tokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	var userAuthorizedScopesPb []string
	for _, item := range st.UserAuthorizedScopes {
		itemPb := &item
		if itemPb != nil {
			userAuthorizedScopesPb = append(userAuthorizedScopesPb, *itemPb)
		}
	}
	pb.UserAuthorizedScopes = userAuthorizedScopesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCustomAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCustomAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := getCustomAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCustomAppIntegrationOutputPb struct {
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
	TokenAccessPolicy *tokenAccessPolicyPb `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCustomAppIntegrationOutputFromPb(pb *getCustomAppIntegrationOutputPb) (*GetCustomAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationOutput{}
	clientIdField := &pb.ClientId
	if clientIdField != nil {
		st.ClientId = *clientIdField
	}
	confidentialField := &pb.Confidential
	if confidentialField != nil {
		st.Confidential = *confidentialField
	}
	createTimeField := &pb.CreateTime
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	creatorUsernameField := &pb.CreatorUsername
	if creatorUsernameField != nil {
		st.CreatorUsername = *creatorUsernameField
	}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	var redirectUrlsField []string
	for _, item := range pb.RedirectUrls {
		itemField := &item
		if itemField != nil {
			redirectUrlsField = append(redirectUrlsField, *itemField)
		}
	}
	st.RedirectUrls = redirectUrlsField

	var scopesField []string
	for _, item := range pb.Scopes {
		itemField := &item
		if itemField != nil {
			scopesField = append(scopesField, *itemField)
		}
	}
	st.Scopes = scopesField
	tokenAccessPolicyField, err := tokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	var userAuthorizedScopesField []string
	for _, item := range pb.UserAuthorizedScopes {
		itemField := &item
		if itemField != nil {
			userAuthorizedScopesField = append(userAuthorizedScopesField, *itemField)
		}
	}
	st.UserAuthorizedScopes = userAuthorizedScopesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCustomAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCustomAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get OAuth Custom App Integration
type GetCustomAppIntegrationRequest struct {
	// The OAuth app integration ID.
	IntegrationId string
}

func getCustomAppIntegrationRequestToPb(st *GetCustomAppIntegrationRequest) (*getCustomAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCustomAppIntegrationRequestPb{}
	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	return pb, nil
}

func (st *GetCustomAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCustomAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCustomAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCustomAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCustomAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCustomAppIntegrationRequestPb struct {
	// The OAuth app integration ID.
	IntegrationId string `json:"-" url:"-"`
}

func getCustomAppIntegrationRequestFromPb(pb *getCustomAppIntegrationRequestPb) (*GetCustomAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationRequest{}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}

	return st, nil
}

type GetCustomAppIntegrationsOutput struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps []GetCustomAppIntegrationOutput

	NextPageToken string

	ForceSendFields []string
}

func getCustomAppIntegrationsOutputToPb(st *GetCustomAppIntegrationsOutput) (*getCustomAppIntegrationsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCustomAppIntegrationsOutputPb{}

	var appsPb []getCustomAppIntegrationOutputPb
	for _, item := range st.Apps {
		itemPb, err := getCustomAppIntegrationOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetCustomAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCustomAppIntegrationsOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCustomAppIntegrationsOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCustomAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	pb, err := getCustomAppIntegrationsOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCustomAppIntegrationsOutputPb struct {
	// List of Custom OAuth App Integrations defined for the account.
	Apps []getCustomAppIntegrationOutputPb `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCustomAppIntegrationsOutputFromPb(pb *getCustomAppIntegrationsOutputPb) (*GetCustomAppIntegrationsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationsOutput{}

	var appsField []GetCustomAppIntegrationOutput
	for _, item := range pb.Apps {
		itemField, err := getCustomAppIntegrationOutputFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			appsField = append(appsField, *itemField)
		}
	}
	st.Apps = appsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCustomAppIntegrationsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCustomAppIntegrationsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPublishedAppIntegrationOutput struct {
	// App-id of the published app integration
	AppId string

	CreateTime string

	CreatedBy int64
	// Unique integration id for the published OAuth app
	IntegrationId string
	// Display name of the published OAuth app
	Name string
	// Token access policy
	TokenAccessPolicy *TokenAccessPolicy

	ForceSendFields []string
}

func getPublishedAppIntegrationOutputToPb(st *GetPublishedAppIntegrationOutput) (*getPublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppIntegrationOutputPb{}
	appIdPb := &st.AppId
	if appIdPb != nil {
		pb.AppId = *appIdPb
	}

	createTimePb := &st.CreateTime
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	createdByPb := &st.CreatedBy
	if createdByPb != nil {
		pb.CreatedBy = *createdByPb
	}

	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	tokenAccessPolicyPb, err := tokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetPublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getPublishedAppIntegrationOutputPb struct {
	// App-id of the published app integration
	AppId string `json:"app_id,omitempty"`

	CreateTime string `json:"create_time,omitempty"`

	CreatedBy int64 `json:"created_by,omitempty"`
	// Unique integration id for the published OAuth app
	IntegrationId string `json:"integration_id,omitempty"`
	// Display name of the published OAuth app
	Name string `json:"name,omitempty"`
	// Token access policy
	TokenAccessPolicy *tokenAccessPolicyPb `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedAppIntegrationOutputFromPb(pb *getPublishedAppIntegrationOutputPb) (*GetPublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationOutput{}
	appIdField := &pb.AppId
	if appIdField != nil {
		st.AppId = *appIdField
	}
	createTimeField := &pb.CreateTime
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	createdByField := &pb.CreatedBy
	if createdByField != nil {
		st.CreatedBy = *createdByField
	}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}
	tokenAccessPolicyField, err := tokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get OAuth Published App Integration
type GetPublishedAppIntegrationRequest struct {
	IntegrationId string
}

func getPublishedAppIntegrationRequestToPb(st *GetPublishedAppIntegrationRequest) (*getPublishedAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppIntegrationRequestPb{}
	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	return pb, nil
}

func (st *GetPublishedAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getPublishedAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

func getPublishedAppIntegrationRequestFromPb(pb *getPublishedAppIntegrationRequestPb) (*GetPublishedAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationRequest{}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}

	return st, nil
}

type GetPublishedAppIntegrationsOutput struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps []GetPublishedAppIntegrationOutput

	NextPageToken string

	ForceSendFields []string
}

func getPublishedAppIntegrationsOutputToPb(st *GetPublishedAppIntegrationsOutput) (*getPublishedAppIntegrationsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppIntegrationsOutputPb{}

	var appsPb []getPublishedAppIntegrationOutputPb
	for _, item := range st.Apps {
		itemPb, err := getPublishedAppIntegrationOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetPublishedAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedAppIntegrationsOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedAppIntegrationsOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedAppIntegrationsOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getPublishedAppIntegrationsOutputPb struct {
	// List of Published OAuth App Integrations defined for the account.
	Apps []getPublishedAppIntegrationOutputPb `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedAppIntegrationsOutputFromPb(pb *getPublishedAppIntegrationsOutputPb) (*GetPublishedAppIntegrationsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationsOutput{}

	var appsField []GetPublishedAppIntegrationOutput
	for _, item := range pb.Apps {
		itemField, err := getPublishedAppIntegrationOutputFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			appsField = append(appsField, *itemField)
		}
	}
	st.Apps = appsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedAppIntegrationsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedAppIntegrationsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	Apps []PublishedAppOutput
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string

	ForceSendFields []string
}

func getPublishedAppsOutputToPb(st *GetPublishedAppsOutput) (*getPublishedAppsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppsOutputPb{}

	var appsPb []publishedAppOutputPb
	for _, item := range st.Apps {
		itemPb, err := publishedAppOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb

	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetPublishedAppsOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getPublishedAppsOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getPublishedAppsOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetPublishedAppsOutput) MarshalJSON() ([]byte, error) {
	pb, err := getPublishedAppsOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getPublishedAppsOutputPb struct {
	// List of Published OAuth Apps.
	Apps []publishedAppOutputPb `json:"apps,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedAppsOutputFromPb(pb *getPublishedAppsOutputPb) (*GetPublishedAppsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppsOutput{}

	var appsField []PublishedAppOutput
	for _, item := range pb.Apps {
		itemField, err := publishedAppOutputFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			appsField = append(appsField, *itemField)
		}
	}
	st.Apps = appsField
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedAppsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedAppsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get service principal federation policy
type GetServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string
	// The service principal id for the federation policy.
	ServicePrincipalId int64
}

func getServicePrincipalFederationPolicyRequestToPb(st *GetServicePrincipalFederationPolicyRequest) (*getServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServicePrincipalFederationPolicyRequestPb{}
	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	return pb, nil
}

func (st *GetServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := getServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getServicePrincipalFederationPolicyRequestPb struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" url:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`
}

func getServicePrincipalFederationPolicyRequestFromPb(pb *getServicePrincipalFederationPolicyRequestPb) (*GetServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServicePrincipalFederationPolicyRequest{}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}

	return st, nil
}

// List account federation policies
type ListAccountFederationPoliciesRequest struct {
	PageSize int

	PageToken string

	ForceSendFields []string
}

func listAccountFederationPoliciesRequestToPb(st *ListAccountFederationPoliciesRequest) (*listAccountFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountFederationPoliciesRequestPb{}
	pageSizePb := &st.PageSize
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListAccountFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listAccountFederationPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listAccountFederationPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListAccountFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listAccountFederationPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listAccountFederationPoliciesRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAccountFederationPoliciesRequestFromPb(pb *listAccountFederationPoliciesRequestPb) (*ListAccountFederationPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountFederationPoliciesRequest{}
	pageSizeField := &pb.PageSize
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAccountFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAccountFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get custom oauth app integrations
type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername bool

	PageSize int

	PageToken string

	ForceSendFields []string
}

func listCustomAppIntegrationsRequestToPb(st *ListCustomAppIntegrationsRequest) (*listCustomAppIntegrationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCustomAppIntegrationsRequestPb{}
	includeCreatorUsernamePb := &st.IncludeCreatorUsername
	if includeCreatorUsernamePb != nil {
		pb.IncludeCreatorUsername = *includeCreatorUsernamePb
	}

	pageSizePb := &st.PageSize
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListCustomAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listCustomAppIntegrationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listCustomAppIntegrationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListCustomAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listCustomAppIntegrationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listCustomAppIntegrationsRequestPb struct {
	IncludeCreatorUsername bool `json:"-" url:"include_creator_username,omitempty"`

	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCustomAppIntegrationsRequestFromPb(pb *listCustomAppIntegrationsRequestPb) (*ListCustomAppIntegrationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCustomAppIntegrationsRequest{}
	includeCreatorUsernameField := &pb.IncludeCreatorUsername
	if includeCreatorUsernameField != nil {
		st.IncludeCreatorUsername = *includeCreatorUsernameField
	}
	pageSizeField := &pb.PageSize
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCustomAppIntegrationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCustomAppIntegrationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListFederationPoliciesResponse struct {
	NextPageToken string

	Policies []FederationPolicy

	ForceSendFields []string
}

func listFederationPoliciesResponseToPb(st *ListFederationPoliciesResponse) (*listFederationPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFederationPoliciesResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	var policiesPb []federationPolicyPb
	for _, item := range st.Policies {
		itemPb, err := federationPolicyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policiesPb = append(policiesPb, *itemPb)
		}
	}
	pb.Policies = policiesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListFederationPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listFederationPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listFederationPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListFederationPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := listFederationPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listFederationPoliciesResponsePb struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Policies []federationPolicyPb `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFederationPoliciesResponseFromPb(pb *listFederationPoliciesResponsePb) (*ListFederationPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFederationPoliciesResponse{}
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var policiesField []FederationPolicy
	for _, item := range pb.Policies {
		itemField, err := federationPolicyFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			policiesField = append(policiesField, *itemField)
		}
	}
	st.Policies = policiesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFederationPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFederationPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get all the published OAuth apps
type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize int
	// A token that can be used to get the next page of results.
	PageToken string

	ForceSendFields []string
}

func listOAuthPublishedAppsRequestToPb(st *ListOAuthPublishedAppsRequest) (*listOAuthPublishedAppsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listOAuthPublishedAppsRequestPb{}
	pageSizePb := &st.PageSize
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListOAuthPublishedAppsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listOAuthPublishedAppsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listOAuthPublishedAppsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListOAuthPublishedAppsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listOAuthPublishedAppsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listOAuthPublishedAppsRequestPb struct {
	// The max number of OAuth published apps to return in one page.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A token that can be used to get the next page of results.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listOAuthPublishedAppsRequestFromPb(pb *listOAuthPublishedAppsRequestPb) (*ListOAuthPublishedAppsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListOAuthPublishedAppsRequest{}
	pageSizeField := &pb.PageSize
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listOAuthPublishedAppsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listOAuthPublishedAppsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get published oauth app integrations
type ListPublishedAppIntegrationsRequest struct {
	PageSize int

	PageToken string

	ForceSendFields []string
}

func listPublishedAppIntegrationsRequestToPb(st *ListPublishedAppIntegrationsRequest) (*listPublishedAppIntegrationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPublishedAppIntegrationsRequestPb{}
	pageSizePb := &st.PageSize
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListPublishedAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listPublishedAppIntegrationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listPublishedAppIntegrationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListPublishedAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listPublishedAppIntegrationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listPublishedAppIntegrationsRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPublishedAppIntegrationsRequestFromPb(pb *listPublishedAppIntegrationsRequestPb) (*ListPublishedAppIntegrationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPublishedAppIntegrationsRequest{}
	pageSizeField := &pb.PageSize
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPublishedAppIntegrationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPublishedAppIntegrationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List service principal federation policies
type ListServicePrincipalFederationPoliciesRequest struct {
	PageSize int

	PageToken string
	// The service principal id for the federation policy.
	ServicePrincipalId int64

	ForceSendFields []string
}

func listServicePrincipalFederationPoliciesRequestToPb(st *ListServicePrincipalFederationPoliciesRequest) (*listServicePrincipalFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalFederationPoliciesRequestPb{}
	pageSizePb := &st.PageSize
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListServicePrincipalFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listServicePrincipalFederationPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listServicePrincipalFederationPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListServicePrincipalFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := listServicePrincipalFederationPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listServicePrincipalFederationPoliciesRequestPb struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listServicePrincipalFederationPoliciesRequestFromPb(pb *listServicePrincipalFederationPoliciesRequestPb) (*ListServicePrincipalFederationPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalFederationPoliciesRequest{}
	pageSizeField := &pb.PageSize
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	PageToken string
	// The service principal ID.
	ServicePrincipalId int64

	ForceSendFields []string
}

func listServicePrincipalSecretsRequestToPb(st *ListServicePrincipalSecretsRequest) (*listServicePrincipalSecretsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalSecretsRequestPb{}
	pageTokenPb := &st.PageToken
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListServicePrincipalSecretsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listServicePrincipalSecretsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listServicePrincipalSecretsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListServicePrincipalSecretsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listServicePrincipalSecretsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listServicePrincipalSecretsRequestPb struct {
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

func listServicePrincipalSecretsRequestFromPb(pb *listServicePrincipalSecretsRequestPb) (*ListServicePrincipalSecretsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalSecretsRequest{}
	pageTokenField := &pb.PageToken
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalSecretsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalSecretsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListServicePrincipalSecretsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken string
	// List of the secrets
	Secrets []SecretInfo

	ForceSendFields []string
}

func listServicePrincipalSecretsResponseToPb(st *ListServicePrincipalSecretsResponse) (*listServicePrincipalSecretsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalSecretsResponsePb{}
	nextPageTokenPb := &st.NextPageToken
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	var secretsPb []secretInfoPb
	for _, item := range st.Secrets {
		itemPb, err := secretInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			secretsPb = append(secretsPb, *itemPb)
		}
	}
	pb.Secrets = secretsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListServicePrincipalSecretsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listServicePrincipalSecretsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listServicePrincipalSecretsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListServicePrincipalSecretsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listServicePrincipalSecretsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listServicePrincipalSecretsResponsePb struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of the secrets
	Secrets []secretInfoPb `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listServicePrincipalSecretsResponseFromPb(pb *listServicePrincipalSecretsResponsePb) (*ListServicePrincipalSecretsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalSecretsResponse{}
	nextPageTokenField := &pb.NextPageToken
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	var secretsField []SecretInfo
	for _, item := range pb.Secrets {
		itemField, err := secretInfoFromPb(&item)
		if err != nil {
			return nil, err
		}
		if itemField != nil {
			secretsField = append(secretsField, *itemField)
		}
	}
	st.Secrets = secretsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalSecretsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalSecretsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
	Audiences []string
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	Issuer string
	// The public keys used to validate the signature of federated tokens, in
	// JWKS format. Most use cases should not need to specify this field. If
	// jwks_uri and jwks_json are both unspecified (recommended), Databricks
	// automatically fetches the public keys from your issuer’s well known
	// endpoint. Databricks strongly recommends relying on your issuer’s well
	// known endpoint for discovering public keys.
	JwksJson string
	// URL of the public keys used to validate the signature of federated
	// tokens, in JWKS format. Most use cases should not need to specify this
	// field. If jwks_uri and jwks_json are both unspecified (recommended),
	// Databricks automatically fetches the public keys from your issuer’s
	// well known endpoint. Databricks strongly recommends relying on your
	// issuer’s well known endpoint for discovering public keys.
	JwksUri string
	// The required token subject, as specified in the subject claim of
	// federated tokens. Must be specified for service principal federation
	// policies. Must not be specified for account federation policies.
	Subject string
	// The claim that contains the subject of the token. If unspecified, the
	// default value is 'sub'.
	SubjectClaim string

	ForceSendFields []string
}

func oidcFederationPolicyToPb(st *OidcFederationPolicy) (*oidcFederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oidcFederationPolicyPb{}

	var audiencesPb []string
	for _, item := range st.Audiences {
		itemPb := &item
		if itemPb != nil {
			audiencesPb = append(audiencesPb, *itemPb)
		}
	}
	pb.Audiences = audiencesPb

	issuerPb := &st.Issuer
	if issuerPb != nil {
		pb.Issuer = *issuerPb
	}

	jwksJsonPb := &st.JwksJson
	if jwksJsonPb != nil {
		pb.JwksJson = *jwksJsonPb
	}

	jwksUriPb := &st.JwksUri
	if jwksUriPb != nil {
		pb.JwksUri = *jwksUriPb
	}

	subjectPb := &st.Subject
	if subjectPb != nil {
		pb.Subject = *subjectPb
	}

	subjectClaimPb := &st.SubjectClaim
	if subjectClaimPb != nil {
		pb.SubjectClaim = *subjectClaimPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *OidcFederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oidcFederationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := oidcFederationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OidcFederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := oidcFederationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type oidcFederationPolicyPb struct {
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

func oidcFederationPolicyFromPb(pb *oidcFederationPolicyPb) (*OidcFederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OidcFederationPolicy{}

	var audiencesField []string
	for _, item := range pb.Audiences {
		itemField := &item
		if itemField != nil {
			audiencesField = append(audiencesField, *itemField)
		}
	}
	st.Audiences = audiencesField
	issuerField := &pb.Issuer
	if issuerField != nil {
		st.Issuer = *issuerField
	}
	jwksJsonField := &pb.JwksJson
	if jwksJsonField != nil {
		st.JwksJson = *jwksJsonField
	}
	jwksUriField := &pb.JwksUri
	if jwksUriField != nil {
		st.JwksUri = *jwksUriField
	}
	subjectField := &pb.Subject
	if subjectField != nil {
		st.Subject = *subjectField
	}
	subjectClaimField := &pb.SubjectClaim
	if subjectClaimField != nil {
		st.SubjectClaim = *subjectClaimField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *oidcFederationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st oidcFederationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PublishedAppOutput struct {
	// Unique ID of the published OAuth app.
	AppId string
	// Client ID of the published OAuth app. It is the client_id in the OAuth
	// flow
	ClientId string
	// Description of the published OAuth app.
	Description string
	// Whether the published OAuth app is a confidential client. It is always
	// false for published OAuth apps.
	IsConfidentialClient bool
	// The display name of the published OAuth app.
	Name string
	// Redirect URLs of the published OAuth app.
	RedirectUrls []string
	// Required scopes for the published OAuth app.
	Scopes []string

	ForceSendFields []string
}

func publishedAppOutputToPb(st *PublishedAppOutput) (*publishedAppOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishedAppOutputPb{}
	appIdPb := &st.AppId
	if appIdPb != nil {
		pb.AppId = *appIdPb
	}

	clientIdPb := &st.ClientId
	if clientIdPb != nil {
		pb.ClientId = *clientIdPb
	}

	descriptionPb := &st.Description
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	isConfidentialClientPb := &st.IsConfidentialClient
	if isConfidentialClientPb != nil {
		pb.IsConfidentialClient = *isConfidentialClientPb
	}

	namePb := &st.Name
	if namePb != nil {
		pb.Name = *namePb
	}

	var redirectUrlsPb []string
	for _, item := range st.RedirectUrls {
		itemPb := &item
		if itemPb != nil {
			redirectUrlsPb = append(redirectUrlsPb, *itemPb)
		}
	}
	pb.RedirectUrls = redirectUrlsPb

	var scopesPb []string
	for _, item := range st.Scopes {
		itemPb := &item
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *PublishedAppOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &publishedAppOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := publishedAppOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PublishedAppOutput) MarshalJSON() ([]byte, error) {
	pb, err := publishedAppOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type publishedAppOutputPb struct {
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

func publishedAppOutputFromPb(pb *publishedAppOutputPb) (*PublishedAppOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishedAppOutput{}
	appIdField := &pb.AppId
	if appIdField != nil {
		st.AppId = *appIdField
	}
	clientIdField := &pb.ClientId
	if clientIdField != nil {
		st.ClientId = *clientIdField
	}
	descriptionField := &pb.Description
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	isConfidentialClientField := &pb.IsConfidentialClient
	if isConfidentialClientField != nil {
		st.IsConfidentialClient = *isConfidentialClientField
	}
	nameField := &pb.Name
	if nameField != nil {
		st.Name = *nameField
	}

	var redirectUrlsField []string
	for _, item := range pb.RedirectUrls {
		itemField := &item
		if itemField != nil {
			redirectUrlsField = append(redirectUrlsField, *itemField)
		}
	}
	st.RedirectUrls = redirectUrlsField

	var scopesField []string
	for _, item := range pb.Scopes {
		itemField := &item
		if itemField != nil {
			scopesField = append(scopesField, *itemField)
		}
	}
	st.Scopes = scopesField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publishedAppOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishedAppOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SecretInfo struct {
	// UTC time when the secret was created
	CreateTime string
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	ExpireTime string
	// ID of the secret
	Id string
	// Secret Hash
	SecretHash string
	// Status of the secret
	Status string
	// UTC time when the secret was updated
	UpdateTime string

	ForceSendFields []string
}

func secretInfoToPb(st *SecretInfo) (*secretInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &secretInfoPb{}
	createTimePb := &st.CreateTime
	if createTimePb != nil {
		pb.CreateTime = *createTimePb
	}

	expireTimePb := &st.ExpireTime
	if expireTimePb != nil {
		pb.ExpireTime = *expireTimePb
	}

	idPb := &st.Id
	if idPb != nil {
		pb.Id = *idPb
	}

	secretHashPb := &st.SecretHash
	if secretHashPb != nil {
		pb.SecretHash = *secretHashPb
	}

	statusPb := &st.Status
	if statusPb != nil {
		pb.Status = *statusPb
	}

	updateTimePb := &st.UpdateTime
	if updateTimePb != nil {
		pb.UpdateTime = *updateTimePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SecretInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &secretInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := secretInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SecretInfo) MarshalJSON() ([]byte, error) {
	pb, err := secretInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type secretInfoPb struct {
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

func secretInfoFromPb(pb *secretInfoPb) (*SecretInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretInfo{}
	createTimeField := &pb.CreateTime
	if createTimeField != nil {
		st.CreateTime = *createTimeField
	}
	expireTimeField := &pb.ExpireTime
	if expireTimeField != nil {
		st.ExpireTime = *expireTimeField
	}
	idField := &pb.Id
	if idField != nil {
		st.Id = *idField
	}
	secretHashField := &pb.SecretHash
	if secretHashField != nil {
		st.SecretHash = *secretHashField
	}
	statusField := &pb.Status
	if statusField != nil {
		st.Status = *statusField
	}
	updateTimeField := &pb.UpdateTime
	if updateTimeField != nil {
		st.UpdateTime = *updateTimeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *secretInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st secretInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes int
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes int

	ForceSendFields []string
}

func tokenAccessPolicyToPb(st *TokenAccessPolicy) (*tokenAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenAccessPolicyPb{}
	accessTokenTtlInMinutesPb := &st.AccessTokenTtlInMinutes
	if accessTokenTtlInMinutesPb != nil {
		pb.AccessTokenTtlInMinutes = *accessTokenTtlInMinutesPb
	}

	refreshTokenTtlInMinutesPb := &st.RefreshTokenTtlInMinutes
	if refreshTokenTtlInMinutesPb != nil {
		pb.RefreshTokenTtlInMinutes = *refreshTokenTtlInMinutesPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TokenAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tokenAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tokenAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TokenAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := tokenAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type tokenAccessPolicyPb struct {
	// access token time to live in minutes
	AccessTokenTtlInMinutes int `json:"access_token_ttl_in_minutes,omitempty"`
	// refresh token time to live in minutes
	RefreshTokenTtlInMinutes int `json:"refresh_token_ttl_in_minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenAccessPolicyFromPb(pb *tokenAccessPolicyPb) (*TokenAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessPolicy{}
	accessTokenTtlInMinutesField := &pb.AccessTokenTtlInMinutes
	if accessTokenTtlInMinutesField != nil {
		st.AccessTokenTtlInMinutes = *accessTokenTtlInMinutesField
	}
	refreshTokenTtlInMinutesField := &pb.RefreshTokenTtlInMinutes
	if refreshTokenTtlInMinutesField != nil {
		st.RefreshTokenTtlInMinutes = *refreshTokenTtlInMinutesField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenAccessPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenAccessPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Update account federation policy
type UpdateAccountFederationPolicyRequest struct {
	Policy FederationPolicy
	// The identifier for the federation policy.
	PolicyId string
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask string

	ForceSendFields []string
}

func updateAccountFederationPolicyRequestToPb(st *UpdateAccountFederationPolicyRequest) (*updateAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAccountFederationPolicyRequestPb{}
	policyPb, err := federationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}

	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	updateMaskPb := &st.UpdateMask
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateAccountFederationPolicyRequestPb struct {
	Policy federationPolicyPb `json:"policy"`
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

func updateAccountFederationPolicyRequestFromPb(pb *updateAccountFederationPolicyRequestPb) (*UpdateAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAccountFederationPolicyRequest{}
	policyField, err := federationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	updateMaskField := &pb.UpdateMask
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateAccountFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateAccountFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCustomAppIntegration struct {
	IntegrationId string
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls []string
	// List of OAuth scopes to be updated in the custom OAuth app integration,
	// similar to redirect URIs this will fully replace the existing values
	// instead of appending
	Scopes []string
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy *TokenAccessPolicy
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string
}

func updateCustomAppIntegrationToPb(st *UpdateCustomAppIntegration) (*updateCustomAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCustomAppIntegrationPb{}
	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	var redirectUrlsPb []string
	for _, item := range st.RedirectUrls {
		itemPb := &item
		if itemPb != nil {
			redirectUrlsPb = append(redirectUrlsPb, *itemPb)
		}
	}
	pb.RedirectUrls = redirectUrlsPb

	var scopesPb []string
	for _, item := range st.Scopes {
		itemPb := &item
		if itemPb != nil {
			scopesPb = append(scopesPb, *itemPb)
		}
	}
	pb.Scopes = scopesPb

	tokenAccessPolicyPb, err := tokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	var userAuthorizedScopesPb []string
	for _, item := range st.UserAuthorizedScopes {
		itemPb := &item
		if itemPb != nil {
			userAuthorizedScopesPb = append(userAuthorizedScopesPb, *itemPb)
		}
	}
	pb.UserAuthorizedScopes = userAuthorizedScopesPb

	return pb, nil
}

func (st *UpdateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCustomAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCustomAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := updateCustomAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCustomAppIntegrationPb struct {
	IntegrationId string `json:"-" url:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// List of OAuth scopes to be updated in the custom OAuth app integration,
	// similar to redirect URIs this will fully replace the existing values
	// instead of appending
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy to be updated in the custom OAuth app integration
	TokenAccessPolicy *tokenAccessPolicyPb `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`
}

func updateCustomAppIntegrationFromPb(pb *updateCustomAppIntegrationPb) (*UpdateCustomAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCustomAppIntegration{}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}

	var redirectUrlsField []string
	for _, item := range pb.RedirectUrls {
		itemField := &item
		if itemField != nil {
			redirectUrlsField = append(redirectUrlsField, *itemField)
		}
	}
	st.RedirectUrls = redirectUrlsField

	var scopesField []string
	for _, item := range pb.Scopes {
		itemField := &item
		if itemField != nil {
			scopesField = append(scopesField, *itemField)
		}
	}
	st.Scopes = scopesField
	tokenAccessPolicyField, err := tokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	var userAuthorizedScopesField []string
	for _, item := range pb.UserAuthorizedScopes {
		itemField := &item
		if itemField != nil {
			userAuthorizedScopesField = append(userAuthorizedScopesField, *itemField)
		}
	}
	st.UserAuthorizedScopes = userAuthorizedScopesField

	return st, nil
}

type UpdateCustomAppIntegrationOutput struct {
}

func updateCustomAppIntegrationOutputToPb(st *UpdateCustomAppIntegrationOutput) (*updateCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCustomAppIntegrationOutputPb{}

	return pb, nil
}

func (st *UpdateCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCustomAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCustomAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := updateCustomAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCustomAppIntegrationOutputPb struct {
}

func updateCustomAppIntegrationOutputFromPb(pb *updateCustomAppIntegrationOutputPb) (*UpdateCustomAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCustomAppIntegrationOutput{}

	return st, nil
}

type UpdatePublishedAppIntegration struct {
	IntegrationId string
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy *TokenAccessPolicy
}

func updatePublishedAppIntegrationToPb(st *UpdatePublishedAppIntegration) (*updatePublishedAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePublishedAppIntegrationPb{}
	integrationIdPb := &st.IntegrationId
	if integrationIdPb != nil {
		pb.IntegrationId = *integrationIdPb
	}

	tokenAccessPolicyPb, err := tokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	return pb, nil
}

func (st *UpdatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePublishedAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePublishedAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := updatePublishedAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updatePublishedAppIntegrationPb struct {
	IntegrationId string `json:"-" url:"-"`
	// Token access policy to be updated in the published OAuth app integration
	TokenAccessPolicy *tokenAccessPolicyPb `json:"token_access_policy,omitempty"`
}

func updatePublishedAppIntegrationFromPb(pb *updatePublishedAppIntegrationPb) (*UpdatePublishedAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePublishedAppIntegration{}
	integrationIdField := &pb.IntegrationId
	if integrationIdField != nil {
		st.IntegrationId = *integrationIdField
	}
	tokenAccessPolicyField, err := tokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	return st, nil
}

type UpdatePublishedAppIntegrationOutput struct {
}

func updatePublishedAppIntegrationOutputToPb(st *UpdatePublishedAppIntegrationOutput) (*updatePublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePublishedAppIntegrationOutputPb{}

	return pb, nil
}

func (st *UpdatePublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updatePublishedAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updatePublishedAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdatePublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := updatePublishedAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updatePublishedAppIntegrationOutputPb struct {
}

func updatePublishedAppIntegrationOutputFromPb(pb *updatePublishedAppIntegrationOutputPb) (*UpdatePublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePublishedAppIntegrationOutput{}

	return st, nil
}

// Update service principal federation policy
type UpdateServicePrincipalFederationPolicyRequest struct {
	Policy FederationPolicy
	// The identifier for the federation policy.
	PolicyId string
	// The service principal id for the federation policy.
	ServicePrincipalId int64
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask string

	ForceSendFields []string
}

func updateServicePrincipalFederationPolicyRequestToPb(st *UpdateServicePrincipalFederationPolicyRequest) (*updateServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateServicePrincipalFederationPolicyRequestPb{}
	policyPb, err := federationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}

	policyIdPb := &st.PolicyId
	if policyIdPb != nil {
		pb.PolicyId = *policyIdPb
	}

	servicePrincipalIdPb := &st.ServicePrincipalId
	if servicePrincipalIdPb != nil {
		pb.ServicePrincipalId = *servicePrincipalIdPb
	}

	updateMaskPb := &st.UpdateMask
	if updateMaskPb != nil {
		pb.UpdateMask = *updateMaskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateServicePrincipalFederationPolicyRequestPb struct {
	Policy federationPolicyPb `json:"policy"`
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

func updateServicePrincipalFederationPolicyRequestFromPb(pb *updateServicePrincipalFederationPolicyRequestPb) (*UpdateServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateServicePrincipalFederationPolicyRequest{}
	policyField, err := federationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	policyIdField := &pb.PolicyId
	if policyIdField != nil {
		st.PolicyId = *policyIdField
	}
	servicePrincipalIdField := &pb.ServicePrincipalId
	if servicePrincipalIdField != nil {
		st.ServicePrincipalId = *servicePrincipalIdField
	}
	updateMaskField := &pb.UpdateMask
	if updateMaskField != nil {
		st.UpdateMask = *updateMaskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateServicePrincipalFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateServicePrincipalFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
