// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/oauth2/oauth2pb"
)

type CreateAccountFederationPolicyRequest struct {

	// Wire name: 'policy'
	Policy FederationPolicy `json:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId        string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreateAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateAccountFederationPolicyRequestToPb(st *CreateAccountFederationPolicyRequest) (*oauth2pb.CreateAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreateAccountFederationPolicyRequestPb{}
	policyPb, err := FederationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}
	pb.PolicyId = st.PolicyId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateAccountFederationPolicyRequestFromPb(pb *oauth2pb.CreateAccountFederationPolicyRequestPb) (*CreateAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAccountFederationPolicyRequest{}
	policyField, err := FederationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.PolicyId = pb.PolicyId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateCustomAppIntegration struct {
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	// Wire name: 'confidential'
	Confidential bool `json:"confidential,omitempty"`
	// Name of the custom OAuth app
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// List of OAuth redirect urls
	// Wire name: 'redirect_urls'
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// OAuth scopes granted to the application. Supported scopes: all-apis, sql,
	// offline_access, openid, profile, email.
	// Wire name: 'scopes'
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy
	// Wire name: 'token_access_policy'
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	// Wire name: 'user_authorized_scopes'
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`
	ForceSendFields      []string `json:"-" tf:"-"`
}

func (st CreateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := CreateCustomAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreateCustomAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCustomAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCustomAppIntegrationToPb(st *CreateCustomAppIntegration) (*oauth2pb.CreateCustomAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreateCustomAppIntegrationPb{}
	pb.Confidential = st.Confidential
	pb.Name = st.Name
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes
	tokenAccessPolicyPb, err := TokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}
	pb.UserAuthorizedScopes = st.UserAuthorizedScopes

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateCustomAppIntegrationFromPb(pb *oauth2pb.CreateCustomAppIntegrationPb) (*CreateCustomAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomAppIntegration{}
	st.Confidential = pb.Confidential
	st.Name = pb.Name
	st.RedirectUrls = pb.RedirectUrls
	st.Scopes = pb.Scopes
	tokenAccessPolicyField, err := TokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}
	st.UserAuthorizedScopes = pb.UserAuthorizedScopes

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateCustomAppIntegrationOutput struct {
	// OAuth client-id generated by the Databricks
	// Wire name: 'client_id'
	ClientId string `json:"client_id,omitempty"`
	// OAuth client-secret generated by the Databricks. If this is a
	// confidential OAuth app client-secret will be generated.
	// Wire name: 'client_secret'
	ClientSecret string `json:"client_secret,omitempty"`
	// Unique integration id for the custom OAuth app
	// Wire name: 'integration_id'
	IntegrationId   string   `json:"integration_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := CreateCustomAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreateCustomAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCustomAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCustomAppIntegrationOutputToPb(st *CreateCustomAppIntegrationOutput) (*oauth2pb.CreateCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreateCustomAppIntegrationOutputPb{}
	pb.ClientId = st.ClientId
	pb.ClientSecret = st.ClientSecret
	pb.IntegrationId = st.IntegrationId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateCustomAppIntegrationOutputFromPb(pb *oauth2pb.CreateCustomAppIntegrationOutputPb) (*CreateCustomAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomAppIntegrationOutput{}
	st.ClientId = pb.ClientId
	st.ClientSecret = pb.ClientSecret
	st.IntegrationId = pb.IntegrationId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreatePublishedAppIntegration struct {
	// App id of the OAuth published app integration. For example power-bi,
	// tableau-deskop
	// Wire name: 'app_id'
	AppId string `json:"app_id,omitempty"`
	// Token access policy
	// Wire name: 'token_access_policy'
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	ForceSendFields   []string           `json:"-" tf:"-"`
}

func (st CreatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := CreatePublishedAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreatePublishedAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreatePublishedAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreatePublishedAppIntegrationToPb(st *CreatePublishedAppIntegration) (*oauth2pb.CreatePublishedAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreatePublishedAppIntegrationPb{}
	pb.AppId = st.AppId
	tokenAccessPolicyPb, err := TokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreatePublishedAppIntegrationFromPb(pb *oauth2pb.CreatePublishedAppIntegrationPb) (*CreatePublishedAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePublishedAppIntegration{}
	st.AppId = pb.AppId
	tokenAccessPolicyField, err := TokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreatePublishedAppIntegrationOutput struct {
	// Unique integration id for the published OAuth app
	// Wire name: 'integration_id'
	IntegrationId   string   `json:"integration_id,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreatePublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := CreatePublishedAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreatePublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreatePublishedAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreatePublishedAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreatePublishedAppIntegrationOutputToPb(st *CreatePublishedAppIntegrationOutput) (*oauth2pb.CreatePublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreatePublishedAppIntegrationOutputPb{}
	pb.IntegrationId = st.IntegrationId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreatePublishedAppIntegrationOutputFromPb(pb *oauth2pb.CreatePublishedAppIntegrationOutputPb) (*CreatePublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePublishedAppIntegrationOutput{}
	st.IntegrationId = pb.IntegrationId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateServicePrincipalFederationPolicyRequest struct {

	// Wire name: 'policy'
	Policy FederationPolicy `json:"policy"`
	// The identifier for the federation policy. The identifier must contain
	// only lowercase alphanumeric characters, numbers, hyphens, and slashes. If
	// unspecified, the id will be assigned by Databricks.
	PolicyId string `json:"-" tf:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64    `json:"-" tf:"-"`
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st CreateServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreateServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateServicePrincipalFederationPolicyRequestToPb(st *CreateServicePrincipalFederationPolicyRequest) (*oauth2pb.CreateServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreateServicePrincipalFederationPolicyRequestPb{}
	policyPb, err := FederationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateServicePrincipalFederationPolicyRequestFromPb(pb *oauth2pb.CreateServicePrincipalFederationPolicyRequestPb) (*CreateServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalFederationPolicyRequest{}
	policyField, err := FederationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateServicePrincipalSecretRequest struct {
	// The lifetime of the secret in seconds. If this parameter is not provided,
	// the secret will have a default lifetime of 730 days (63072000s).
	// Wire name: 'lifetime'
	Lifetime string `json:"lifetime,omitempty"` //legacy
	// The service principal ID.
	ServicePrincipalId string   `json:"-" tf:"-"`
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st CreateServicePrincipalSecretRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateServicePrincipalSecretRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateServicePrincipalSecretRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreateServicePrincipalSecretRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateServicePrincipalSecretRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateServicePrincipalSecretRequestToPb(st *CreateServicePrincipalSecretRequest) (*oauth2pb.CreateServicePrincipalSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreateServicePrincipalSecretRequestPb{}
	pb.Lifetime = st.Lifetime
	pb.ServicePrincipalId = st.ServicePrincipalId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateServicePrincipalSecretRequestFromPb(pb *oauth2pb.CreateServicePrincipalSecretRequestPb) (*CreateServicePrincipalSecretRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalSecretRequest{}
	st.Lifetime = pb.Lifetime
	st.ServicePrincipalId = pb.ServicePrincipalId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateServicePrincipalSecretResponse struct {
	// UTC time when the secret was created
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	// Wire name: 'expire_time'
	ExpireTime string `json:"expire_time,omitempty"` //legacy
	// ID of the secret
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Secret Value
	// Wire name: 'secret'
	Secret string `json:"secret,omitempty"`
	// Secret Hash
	// Wire name: 'secret_hash'
	SecretHash string `json:"secret_hash,omitempty"`
	// Status of the secret
	// Wire name: 'status'
	Status string `json:"status,omitempty"`
	// UTC time when the secret was updated
	// Wire name: 'update_time'
	UpdateTime      string   `json:"update_time,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateServicePrincipalSecretResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateServicePrincipalSecretResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateServicePrincipalSecretResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.CreateServicePrincipalSecretResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateServicePrincipalSecretResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateServicePrincipalSecretResponseToPb(st *CreateServicePrincipalSecretResponse) (*oauth2pb.CreateServicePrincipalSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.CreateServicePrincipalSecretResponsePb{}
	pb.CreateTime = st.CreateTime
	pb.ExpireTime = st.ExpireTime
	pb.Id = st.Id
	pb.Secret = st.Secret
	pb.SecretHash = st.SecretHash
	pb.Status = st.Status
	pb.UpdateTime = st.UpdateTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateServicePrincipalSecretResponseFromPb(pb *oauth2pb.CreateServicePrincipalSecretResponsePb) (*CreateServicePrincipalSecretResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalSecretResponse{}
	st.CreateTime = pb.CreateTime
	st.ExpireTime = pb.ExpireTime
	st.Id = pb.Id
	st.Secret = pb.Secret
	st.SecretHash = pb.SecretHash
	st.Status = pb.Status
	st.UpdateTime = pb.UpdateTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" tf:"-"`
}

func (st DeleteAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.DeleteAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteAccountFederationPolicyRequestToPb(st *DeleteAccountFederationPolicyRequest) (*oauth2pb.DeleteAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.DeleteAccountFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func DeleteAccountFederationPolicyRequestFromPb(pb *oauth2pb.DeleteAccountFederationPolicyRequestPb) (*DeleteAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

type DeleteCustomAppIntegrationRequest struct {
	IntegrationId string `json:"-" tf:"-"`
}

func (st DeleteCustomAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteCustomAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCustomAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.DeleteCustomAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteCustomAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteCustomAppIntegrationRequestToPb(st *DeleteCustomAppIntegrationRequest) (*oauth2pb.DeleteCustomAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.DeleteCustomAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

func DeleteCustomAppIntegrationRequestFromPb(pb *oauth2pb.DeleteCustomAppIntegrationRequestPb) (*DeleteCustomAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCustomAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

type DeletePublishedAppIntegrationRequest struct {
	IntegrationId string `json:"-" tf:"-"`
}

func (st DeletePublishedAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeletePublishedAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeletePublishedAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.DeletePublishedAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeletePublishedAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeletePublishedAppIntegrationRequestToPb(st *DeletePublishedAppIntegrationRequest) (*oauth2pb.DeletePublishedAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.DeletePublishedAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

func DeletePublishedAppIntegrationRequestFromPb(pb *oauth2pb.DeletePublishedAppIntegrationRequestPb) (*DeletePublishedAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePublishedAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

type DeleteServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" tf:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" tf:"-"`
}

func (st DeleteServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.DeleteServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteServicePrincipalFederationPolicyRequestToPb(st *DeleteServicePrincipalFederationPolicyRequest) (*oauth2pb.DeleteServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.DeleteServicePrincipalFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId

	return pb, nil
}

func DeleteServicePrincipalFederationPolicyRequestFromPb(pb *oauth2pb.DeleteServicePrincipalFederationPolicyRequestPb) (*DeleteServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId

	return st, nil
}

type DeleteServicePrincipalSecretRequest struct {
	// The secret ID.
	SecretId string `json:"-" tf:"-"`
	// The service principal ID.
	ServicePrincipalId string `json:"-" tf:"-"`
}

func (st DeleteServicePrincipalSecretRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteServicePrincipalSecretRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteServicePrincipalSecretRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.DeleteServicePrincipalSecretRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteServicePrincipalSecretRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteServicePrincipalSecretRequestToPb(st *DeleteServicePrincipalSecretRequest) (*oauth2pb.DeleteServicePrincipalSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.DeleteServicePrincipalSecretRequestPb{}
	pb.SecretId = st.SecretId
	pb.ServicePrincipalId = st.ServicePrincipalId

	return pb, nil
}

func DeleteServicePrincipalSecretRequestFromPb(pb *oauth2pb.DeleteServicePrincipalSecretRequestPb) (*DeleteServicePrincipalSecretRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalSecretRequest{}
	st.SecretId = pb.SecretId
	st.ServicePrincipalId = pb.ServicePrincipalId

	return st, nil
}

type FederationPolicy struct {
	// Creation time of the federation policy.
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"` //legacy
	// Description of the federation policy.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Resource name for the federation policy. Example values include
	// `accounts/<account-id>/federationPolicies/my-federation-policy` for
	// Account Federation Policies, and
	// `accounts/<account-id>/servicePrincipals/<service-principal-id>/federationPolicies/my-federation-policy`
	// for Service Principal Federation Policies. Typically an output parameter,
	// which does not need to be specified in create or update requests. If
	// specified in a request, must match the value in the request URL.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'oidc_policy'
	OidcPolicy *OidcFederationPolicy `json:"oidc_policy,omitempty"`
	// The ID of the federation policy.
	// Wire name: 'policy_id'
	PolicyId string `json:"policy_id,omitempty"`
	// The service principal ID that this federation policy applies to. Only set
	// for service principal federation policies.
	// Wire name: 'service_principal_id'
	ServicePrincipalId int64 `json:"service_principal_id,omitempty"`
	// Unique, immutable id of the federation policy.
	// Wire name: 'uid'
	Uid string `json:"uid,omitempty"`
	// Last update time of the federation policy.
	// Wire name: 'update_time'
	UpdateTime      string   `json:"update_time,omitempty"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := FederationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.FederationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FederationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FederationPolicyToPb(st *FederationPolicy) (*oauth2pb.FederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.FederationPolicyPb{}
	pb.CreateTime = st.CreateTime
	pb.Description = st.Description
	pb.Name = st.Name
	oidcPolicyPb, err := OidcFederationPolicyToPb(st.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyPb != nil {
		pb.OidcPolicy = oidcPolicyPb
	}
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId
	pb.Uid = st.Uid
	pb.UpdateTime = st.UpdateTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FederationPolicyFromPb(pb *oauth2pb.FederationPolicyPb) (*FederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FederationPolicy{}
	st.CreateTime = pb.CreateTime
	st.Description = pb.Description
	st.Name = pb.Name
	oidcPolicyField, err := OidcFederationPolicyFromPb(pb.OidcPolicy)
	if err != nil {
		return nil, err
	}
	if oidcPolicyField != nil {
		st.OidcPolicy = oidcPolicyField
	}
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId
	st.Uid = pb.Uid
	st.UpdateTime = pb.UpdateTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetAccountFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" tf:"-"`
}

func (st GetAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetAccountFederationPolicyRequestToPb(st *GetAccountFederationPolicyRequest) (*oauth2pb.GetAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetAccountFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

func GetAccountFederationPolicyRequestFromPb(pb *oauth2pb.GetAccountFederationPolicyRequestPb) (*GetAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

type GetCustomAppIntegrationOutput struct {
	// The client id of the custom OAuth app
	// Wire name: 'client_id'
	ClientId string `json:"client_id,omitempty"`
	// This field indicates whether an OAuth client secret is required to
	// authenticate this client.
	// Wire name: 'confidential'
	Confidential bool `json:"confidential,omitempty"`

	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`

	// Wire name: 'created_by'
	CreatedBy int64 `json:"created_by,omitempty"`

	// Wire name: 'creator_username'
	CreatorUsername string `json:"creator_username,omitempty"`
	// ID of this custom app
	// Wire name: 'integration_id'
	IntegrationId string `json:"integration_id,omitempty"`
	// The display name of the custom OAuth app
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// List of OAuth redirect urls
	// Wire name: 'redirect_urls'
	RedirectUrls []string `json:"redirect_urls,omitempty"`

	// Wire name: 'scopes'
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy
	// Wire name: 'token_access_policy'
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	// Wire name: 'user_authorized_scopes'
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`
	ForceSendFields      []string `json:"-" tf:"-"`
}

func (st GetCustomAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := GetCustomAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCustomAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetCustomAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCustomAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCustomAppIntegrationOutputToPb(st *GetCustomAppIntegrationOutput) (*oauth2pb.GetCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetCustomAppIntegrationOutputPb{}
	pb.ClientId = st.ClientId
	pb.Confidential = st.Confidential
	pb.CreateTime = st.CreateTime
	pb.CreatedBy = st.CreatedBy
	pb.CreatorUsername = st.CreatorUsername
	pb.IntegrationId = st.IntegrationId
	pb.Name = st.Name
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes
	tokenAccessPolicyPb, err := TokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}
	pb.UserAuthorizedScopes = st.UserAuthorizedScopes

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetCustomAppIntegrationOutputFromPb(pb *oauth2pb.GetCustomAppIntegrationOutputPb) (*GetCustomAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationOutput{}
	st.ClientId = pb.ClientId
	st.Confidential = pb.Confidential
	st.CreateTime = pb.CreateTime
	st.CreatedBy = pb.CreatedBy
	st.CreatorUsername = pb.CreatorUsername
	st.IntegrationId = pb.IntegrationId
	st.Name = pb.Name
	st.RedirectUrls = pb.RedirectUrls
	st.Scopes = pb.Scopes
	tokenAccessPolicyField, err := TokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}
	st.UserAuthorizedScopes = pb.UserAuthorizedScopes

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetCustomAppIntegrationRequest struct {
	// The OAuth app integration ID.
	IntegrationId string `json:"-" tf:"-"`
}

func (st GetCustomAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCustomAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCustomAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetCustomAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCustomAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCustomAppIntegrationRequestToPb(st *GetCustomAppIntegrationRequest) (*oauth2pb.GetCustomAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetCustomAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

func GetCustomAppIntegrationRequestFromPb(pb *oauth2pb.GetCustomAppIntegrationRequestPb) (*GetCustomAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

type GetCustomAppIntegrationsOutput struct {
	// List of Custom OAuth App Integrations defined for the account.
	// Wire name: 'apps'
	Apps []GetCustomAppIntegrationOutput `json:"apps,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetCustomAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	pb, err := GetCustomAppIntegrationsOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCustomAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetCustomAppIntegrationsOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCustomAppIntegrationsOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCustomAppIntegrationsOutputToPb(st *GetCustomAppIntegrationsOutput) (*oauth2pb.GetCustomAppIntegrationsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetCustomAppIntegrationsOutputPb{}

	var appsPb []oauth2pb.GetCustomAppIntegrationOutputPb
	for _, item := range st.Apps {
		itemPb, err := GetCustomAppIntegrationOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetCustomAppIntegrationsOutputFromPb(pb *oauth2pb.GetCustomAppIntegrationsOutputPb) (*GetCustomAppIntegrationsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationsOutput{}

	var appsField []GetCustomAppIntegrationOutput
	for _, itemPb := range pb.Apps {
		item, err := GetCustomAppIntegrationOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			appsField = append(appsField, *item)
		}
	}
	st.Apps = appsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetPublishedAppIntegrationOutput struct {
	// App-id of the published app integration
	// Wire name: 'app_id'
	AppId string `json:"app_id,omitempty"`

	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`

	// Wire name: 'created_by'
	CreatedBy int64 `json:"created_by,omitempty"`
	// Unique integration id for the published OAuth app
	// Wire name: 'integration_id'
	IntegrationId string `json:"integration_id,omitempty"`
	// Display name of the published OAuth app
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Token access policy
	// Wire name: 'token_access_policy'
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	ForceSendFields   []string           `json:"-" tf:"-"`
}

func (st GetPublishedAppIntegrationOutput) MarshalJSON() ([]byte, error) {
	pb, err := GetPublishedAppIntegrationOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPublishedAppIntegrationOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetPublishedAppIntegrationOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPublishedAppIntegrationOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPublishedAppIntegrationOutputToPb(st *GetPublishedAppIntegrationOutput) (*oauth2pb.GetPublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetPublishedAppIntegrationOutputPb{}
	pb.AppId = st.AppId
	pb.CreateTime = st.CreateTime
	pb.CreatedBy = st.CreatedBy
	pb.IntegrationId = st.IntegrationId
	pb.Name = st.Name
	tokenAccessPolicyPb, err := TokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetPublishedAppIntegrationOutputFromPb(pb *oauth2pb.GetPublishedAppIntegrationOutputPb) (*GetPublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationOutput{}
	st.AppId = pb.AppId
	st.CreateTime = pb.CreateTime
	st.CreatedBy = pb.CreatedBy
	st.IntegrationId = pb.IntegrationId
	st.Name = pb.Name
	tokenAccessPolicyField, err := TokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetPublishedAppIntegrationRequest struct {
	IntegrationId string `json:"-" tf:"-"`
}

func (st GetPublishedAppIntegrationRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetPublishedAppIntegrationRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPublishedAppIntegrationRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetPublishedAppIntegrationRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPublishedAppIntegrationRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPublishedAppIntegrationRequestToPb(st *GetPublishedAppIntegrationRequest) (*oauth2pb.GetPublishedAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetPublishedAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

func GetPublishedAppIntegrationRequestFromPb(pb *oauth2pb.GetPublishedAppIntegrationRequestPb) (*GetPublishedAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

type GetPublishedAppIntegrationsOutput struct {
	// List of Published OAuth App Integrations defined for the account.
	// Wire name: 'apps'
	Apps []GetPublishedAppIntegrationOutput `json:"apps,omitempty"`

	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetPublishedAppIntegrationsOutput) MarshalJSON() ([]byte, error) {
	pb, err := GetPublishedAppIntegrationsOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPublishedAppIntegrationsOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetPublishedAppIntegrationsOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPublishedAppIntegrationsOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPublishedAppIntegrationsOutputToPb(st *GetPublishedAppIntegrationsOutput) (*oauth2pb.GetPublishedAppIntegrationsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetPublishedAppIntegrationsOutputPb{}

	var appsPb []oauth2pb.GetPublishedAppIntegrationOutputPb
	for _, item := range st.Apps {
		itemPb, err := GetPublishedAppIntegrationOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetPublishedAppIntegrationsOutputFromPb(pb *oauth2pb.GetPublishedAppIntegrationsOutputPb) (*GetPublishedAppIntegrationsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationsOutput{}

	var appsField []GetPublishedAppIntegrationOutput
	for _, itemPb := range pb.Apps {
		item, err := GetPublishedAppIntegrationOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			appsField = append(appsField, *item)
		}
	}
	st.Apps = appsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetPublishedAppsOutput struct {
	// List of Published OAuth Apps.
	// Wire name: 'apps'
	Apps []PublishedAppOutput `json:"apps,omitempty"`
	// A token that can be used to get the next page of results. If not present,
	// there are no more results to show.
	// Wire name: 'next_page_token'
	NextPageToken   string   `json:"next_page_token,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GetPublishedAppsOutput) MarshalJSON() ([]byte, error) {
	pb, err := GetPublishedAppsOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetPublishedAppsOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetPublishedAppsOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetPublishedAppsOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetPublishedAppsOutputToPb(st *GetPublishedAppsOutput) (*oauth2pb.GetPublishedAppsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetPublishedAppsOutputPb{}

	var appsPb []oauth2pb.PublishedAppOutputPb
	for _, item := range st.Apps {
		itemPb, err := PublishedAppOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			appsPb = append(appsPb, *itemPb)
		}
	}
	pb.Apps = appsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetPublishedAppsOutputFromPb(pb *oauth2pb.GetPublishedAppsOutputPb) (*GetPublishedAppsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppsOutput{}

	var appsField []PublishedAppOutput
	for _, itemPb := range pb.Apps {
		item, err := PublishedAppOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			appsField = append(appsField, *item)
		}
	}
	st.Apps = appsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetServicePrincipalFederationPolicyRequest struct {
	// The identifier for the federation policy.
	PolicyId string `json:"-" tf:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" tf:"-"`
}

func (st GetServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.GetServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetServicePrincipalFederationPolicyRequestToPb(st *GetServicePrincipalFederationPolicyRequest) (*oauth2pb.GetServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.GetServicePrincipalFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId

	return pb, nil
}

func GetServicePrincipalFederationPolicyRequestFromPb(pb *oauth2pb.GetServicePrincipalFederationPolicyRequestPb) (*GetServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServicePrincipalFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId

	return st, nil
}

type ListAccountFederationPoliciesRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListAccountFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListAccountFederationPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListAccountFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListAccountFederationPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListAccountFederationPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListAccountFederationPoliciesRequestToPb(st *ListAccountFederationPoliciesRequest) (*oauth2pb.ListAccountFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListAccountFederationPoliciesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListAccountFederationPoliciesRequestFromPb(pb *oauth2pb.ListAccountFederationPoliciesRequestPb) (*ListAccountFederationPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountFederationPoliciesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListCustomAppIntegrationsRequest struct {
	IncludeCreatorUsername bool `json:"-" tf:"-"`

	PageSize int `json:"-" tf:"-"`

	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListCustomAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListCustomAppIntegrationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListCustomAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListCustomAppIntegrationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListCustomAppIntegrationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListCustomAppIntegrationsRequestToPb(st *ListCustomAppIntegrationsRequest) (*oauth2pb.ListCustomAppIntegrationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListCustomAppIntegrationsRequestPb{}
	pb.IncludeCreatorUsername = st.IncludeCreatorUsername
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListCustomAppIntegrationsRequestFromPb(pb *oauth2pb.ListCustomAppIntegrationsRequestPb) (*ListCustomAppIntegrationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCustomAppIntegrationsRequest{}
	st.IncludeCreatorUsername = pb.IncludeCreatorUsername
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListFederationPoliciesResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'policies'
	Policies        []FederationPolicy `json:"policies,omitempty"`
	ForceSendFields []string           `json:"-" tf:"-"`
}

func (st ListFederationPoliciesResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListFederationPoliciesResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListFederationPoliciesResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListFederationPoliciesResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListFederationPoliciesResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListFederationPoliciesResponseToPb(st *ListFederationPoliciesResponse) (*oauth2pb.ListFederationPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListFederationPoliciesResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var policiesPb []oauth2pb.FederationPolicyPb
	for _, item := range st.Policies {
		itemPb, err := FederationPolicyToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			policiesPb = append(policiesPb, *itemPb)
		}
	}
	pb.Policies = policiesPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListFederationPoliciesResponseFromPb(pb *oauth2pb.ListFederationPoliciesResponsePb) (*ListFederationPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFederationPoliciesResponse{}
	st.NextPageToken = pb.NextPageToken

	var policiesField []FederationPolicy
	for _, itemPb := range pb.Policies {
		item, err := FederationPolicyFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			policiesField = append(policiesField, *item)
		}
	}
	st.Policies = policiesField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListOAuthPublishedAppsRequest struct {
	// The max number of OAuth published apps to return in one page.
	PageSize int `json:"-" tf:"-"`
	// A token that can be used to get the next page of results.
	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListOAuthPublishedAppsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListOAuthPublishedAppsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListOAuthPublishedAppsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListOAuthPublishedAppsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListOAuthPublishedAppsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListOAuthPublishedAppsRequestToPb(st *ListOAuthPublishedAppsRequest) (*oauth2pb.ListOAuthPublishedAppsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListOAuthPublishedAppsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListOAuthPublishedAppsRequestFromPb(pb *oauth2pb.ListOAuthPublishedAppsRequestPb) (*ListOAuthPublishedAppsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListOAuthPublishedAppsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListPublishedAppIntegrationsRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken       string   `json:"-" tf:"-"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ListPublishedAppIntegrationsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListPublishedAppIntegrationsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListPublishedAppIntegrationsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListPublishedAppIntegrationsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListPublishedAppIntegrationsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListPublishedAppIntegrationsRequestToPb(st *ListPublishedAppIntegrationsRequest) (*oauth2pb.ListPublishedAppIntegrationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListPublishedAppIntegrationsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListPublishedAppIntegrationsRequestFromPb(pb *oauth2pb.ListPublishedAppIntegrationsRequestPb) (*ListPublishedAppIntegrationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPublishedAppIntegrationsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListServicePrincipalFederationPoliciesRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64    `json:"-" tf:"-"`
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st ListServicePrincipalFederationPoliciesRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListServicePrincipalFederationPoliciesRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListServicePrincipalFederationPoliciesRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListServicePrincipalFederationPoliciesRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListServicePrincipalFederationPoliciesRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListServicePrincipalFederationPoliciesRequestToPb(st *ListServicePrincipalFederationPoliciesRequest) (*oauth2pb.ListServicePrincipalFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListServicePrincipalFederationPoliciesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ServicePrincipalId = st.ServicePrincipalId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListServicePrincipalFederationPoliciesRequestFromPb(pb *oauth2pb.ListServicePrincipalFederationPoliciesRequestPb) (*ListServicePrincipalFederationPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalFederationPoliciesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ServicePrincipalId = pb.ServicePrincipalId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListServicePrincipalSecretsRequest struct {
	PageSize int `json:"-" tf:"-"`
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the secrets for this service principal.
	// Provide this token to retrieve the next page of secret entries. When
	// providing a `page_token`, all other parameters provided to the request
	// must match the previous request. To list all of the secrets for a service
	// principal, it is necessary to continue requesting pages of entries until
	// the response contains no `next_page_token`. Note that the number of
	// entries returned must not be used to determine when the listing is
	// complete.
	PageToken string `json:"-" tf:"-"`
	// The service principal ID.
	ServicePrincipalId string   `json:"-" tf:"-"`
	ForceSendFields    []string `json:"-" tf:"-"`
}

func (st ListServicePrincipalSecretsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListServicePrincipalSecretsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListServicePrincipalSecretsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListServicePrincipalSecretsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListServicePrincipalSecretsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListServicePrincipalSecretsRequestToPb(st *ListServicePrincipalSecretsRequest) (*oauth2pb.ListServicePrincipalSecretsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListServicePrincipalSecretsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ServicePrincipalId = st.ServicePrincipalId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListServicePrincipalSecretsRequestFromPb(pb *oauth2pb.ListServicePrincipalSecretsRequestPb) (*ListServicePrincipalSecretsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalSecretsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ServicePrincipalId = pb.ServicePrincipalId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListServicePrincipalSecretsResponse struct {
	// A token, which can be sent as `page_token` to retrieve the next page.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`
	// List of the secrets
	// Wire name: 'secrets'
	Secrets         []SecretInfo `json:"secrets,omitempty"`
	ForceSendFields []string     `json:"-" tf:"-"`
}

func (st ListServicePrincipalSecretsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListServicePrincipalSecretsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListServicePrincipalSecretsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.ListServicePrincipalSecretsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListServicePrincipalSecretsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListServicePrincipalSecretsResponseToPb(st *ListServicePrincipalSecretsResponse) (*oauth2pb.ListServicePrincipalSecretsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.ListServicePrincipalSecretsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var secretsPb []oauth2pb.SecretInfoPb
	for _, item := range st.Secrets {
		itemPb, err := SecretInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			secretsPb = append(secretsPb, *itemPb)
		}
	}
	pb.Secrets = secretsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListServicePrincipalSecretsResponseFromPb(pb *oauth2pb.ListServicePrincipalSecretsResponsePb) (*ListServicePrincipalSecretsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalSecretsResponse{}
	st.NextPageToken = pb.NextPageToken

	var secretsField []SecretInfo
	for _, itemPb := range pb.Secrets {
		item, err := SecretInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			secretsField = append(secretsField, *item)
		}
	}
	st.Secrets = secretsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	// Wire name: 'audiences'
	Audiences []string `json:"audiences,omitempty"`
	// The required token issuer, as specified in the 'iss' claim of federated
	// tokens.
	// Wire name: 'issuer'
	Issuer string `json:"issuer,omitempty"`
	// The public keys used to validate the signature of federated tokens, in
	// JWKS format. Most use cases should not need to specify this field. If
	// jwks_uri and jwks_json are both unspecified (recommended), Databricks
	// automatically fetches the public keys from your issuer’s well known
	// endpoint. Databricks strongly recommends relying on your issuer’s well
	// known endpoint for discovering public keys.
	// Wire name: 'jwks_json'
	JwksJson string `json:"jwks_json,omitempty"`
	// URL of the public keys used to validate the signature of federated
	// tokens, in JWKS format. Most use cases should not need to specify this
	// field. If jwks_uri and jwks_json are both unspecified (recommended),
	// Databricks automatically fetches the public keys from your issuer’s
	// well known endpoint. Databricks strongly recommends relying on your
	// issuer’s well known endpoint for discovering public keys.
	// Wire name: 'jwks_uri'
	JwksUri string `json:"jwks_uri,omitempty"`
	// The required token subject, as specified in the subject claim of
	// federated tokens. Must be specified for service principal federation
	// policies. Must not be specified for account federation policies.
	// Wire name: 'subject'
	Subject string `json:"subject,omitempty"`
	// The claim that contains the subject of the token. If unspecified, the
	// default value is 'sub'.
	// Wire name: 'subject_claim'
	SubjectClaim    string   `json:"subject_claim,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OidcFederationPolicy) MarshalJSON() ([]byte, error) {
	pb, err := OidcFederationPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *OidcFederationPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.OidcFederationPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := OidcFederationPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func OidcFederationPolicyToPb(st *OidcFederationPolicy) (*oauth2pb.OidcFederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.OidcFederationPolicyPb{}
	pb.Audiences = st.Audiences
	pb.Issuer = st.Issuer
	pb.JwksJson = st.JwksJson
	pb.JwksUri = st.JwksUri
	pb.Subject = st.Subject
	pb.SubjectClaim = st.SubjectClaim

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func OidcFederationPolicyFromPb(pb *oauth2pb.OidcFederationPolicyPb) (*OidcFederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OidcFederationPolicy{}
	st.Audiences = pb.Audiences
	st.Issuer = pb.Issuer
	st.JwksJson = pb.JwksJson
	st.JwksUri = pb.JwksUri
	st.Subject = pb.Subject
	st.SubjectClaim = pb.SubjectClaim

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type PublishedAppOutput struct {
	// Unique ID of the published OAuth app.
	// Wire name: 'app_id'
	AppId string `json:"app_id,omitempty"`
	// Client ID of the published OAuth app. It is the client_id in the OAuth
	// flow
	// Wire name: 'client_id'
	ClientId string `json:"client_id,omitempty"`
	// Description of the published OAuth app.
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Whether the published OAuth app is a confidential client. It is always
	// false for published OAuth apps.
	// Wire name: 'is_confidential_client'
	IsConfidentialClient bool `json:"is_confidential_client,omitempty"`
	// The display name of the published OAuth app.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Redirect URLs of the published OAuth app.
	// Wire name: 'redirect_urls'
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// Required scopes for the published OAuth app.
	// Wire name: 'scopes'
	Scopes          []string `json:"scopes,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PublishedAppOutput) MarshalJSON() ([]byte, error) {
	pb, err := PublishedAppOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PublishedAppOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.PublishedAppOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PublishedAppOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PublishedAppOutputToPb(st *PublishedAppOutput) (*oauth2pb.PublishedAppOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.PublishedAppOutputPb{}
	pb.AppId = st.AppId
	pb.ClientId = st.ClientId
	pb.Description = st.Description
	pb.IsConfidentialClient = st.IsConfidentialClient
	pb.Name = st.Name
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PublishedAppOutputFromPb(pb *oauth2pb.PublishedAppOutputPb) (*PublishedAppOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishedAppOutput{}
	st.AppId = pb.AppId
	st.ClientId = pb.ClientId
	st.Description = pb.Description
	st.IsConfidentialClient = pb.IsConfidentialClient
	st.Name = pb.Name
	st.RedirectUrls = pb.RedirectUrls
	st.Scopes = pb.Scopes

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SecretInfo struct {
	// UTC time when the secret was created
	// Wire name: 'create_time'
	CreateTime string `json:"create_time,omitempty"`
	// UTC time when the secret will expire. If the field is not present, the
	// secret does not expire.
	// Wire name: 'expire_time'
	ExpireTime string `json:"expire_time,omitempty"` //legacy
	// ID of the secret
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Secret Hash
	// Wire name: 'secret_hash'
	SecretHash string `json:"secret_hash,omitempty"`
	// Status of the secret
	// Wire name: 'status'
	Status string `json:"status,omitempty"`
	// UTC time when the secret was updated
	// Wire name: 'update_time'
	UpdateTime      string   `json:"update_time,omitempty"`
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st SecretInfo) MarshalJSON() ([]byte, error) {
	pb, err := SecretInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SecretInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.SecretInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SecretInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SecretInfoToPb(st *SecretInfo) (*oauth2pb.SecretInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.SecretInfoPb{}
	pb.CreateTime = st.CreateTime
	pb.ExpireTime = st.ExpireTime
	pb.Id = st.Id
	pb.SecretHash = st.SecretHash
	pb.Status = st.Status
	pb.UpdateTime = st.UpdateTime

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SecretInfoFromPb(pb *oauth2pb.SecretInfoPb) (*SecretInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SecretInfo{}
	st.CreateTime = pb.CreateTime
	st.ExpireTime = pb.ExpireTime
	st.Id = pb.Id
	st.SecretHash = pb.SecretHash
	st.Status = pb.Status
	st.UpdateTime = pb.UpdateTime

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TokenAccessPolicy struct {
	// access token time to live in minutes
	// Wire name: 'access_token_ttl_in_minutes'
	AccessTokenTtlInMinutes int `json:"access_token_ttl_in_minutes,omitempty"`
	// refresh token time to live in minutes
	// Wire name: 'refresh_token_ttl_in_minutes'
	RefreshTokenTtlInMinutes int      `json:"refresh_token_ttl_in_minutes,omitempty"`
	ForceSendFields          []string `json:"-" tf:"-"`
}

func (st TokenAccessPolicy) MarshalJSON() ([]byte, error) {
	pb, err := TokenAccessPolicyToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TokenAccessPolicy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.TokenAccessPolicyPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TokenAccessPolicyFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TokenAccessPolicyToPb(st *TokenAccessPolicy) (*oauth2pb.TokenAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.TokenAccessPolicyPb{}
	pb.AccessTokenTtlInMinutes = st.AccessTokenTtlInMinutes
	pb.RefreshTokenTtlInMinutes = st.RefreshTokenTtlInMinutes

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TokenAccessPolicyFromPb(pb *oauth2pb.TokenAccessPolicyPb) (*TokenAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessPolicy{}
	st.AccessTokenTtlInMinutes = pb.AccessTokenTtlInMinutes
	st.RefreshTokenTtlInMinutes = pb.RefreshTokenTtlInMinutes

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateAccountFederationPolicyRequest struct {

	// Wire name: 'policy'
	Policy FederationPolicy `json:"policy"`
	// The identifier for the federation policy.
	PolicyId string `json:"-" tf:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask      string   `json:"-" tf:"-"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateAccountFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateAccountFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateAccountFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.UpdateAccountFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateAccountFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateAccountFederationPolicyRequestToPb(st *UpdateAccountFederationPolicyRequest) (*oauth2pb.UpdateAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.UpdateAccountFederationPolicyRequestPb{}
	policyPb, err := FederationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}
	pb.PolicyId = st.PolicyId
	pb.UpdateMask = st.UpdateMask

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateAccountFederationPolicyRequestFromPb(pb *oauth2pb.UpdateAccountFederationPolicyRequestPb) (*UpdateAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAccountFederationPolicyRequest{}
	policyField, err := FederationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.PolicyId = pb.PolicyId
	st.UpdateMask = pb.UpdateMask

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateCustomAppIntegration struct {
	IntegrationId string `json:"-" tf:"-"`
	// List of OAuth redirect urls to be updated in the custom OAuth app
	// integration
	// Wire name: 'redirect_urls'
	RedirectUrls []string `json:"redirect_urls,omitempty"`
	// List of OAuth scopes to be updated in the custom OAuth app integration,
	// similar to redirect URIs this will fully replace the existing values
	// instead of appending
	// Wire name: 'scopes'
	Scopes []string `json:"scopes,omitempty"`
	// Token access policy to be updated in the custom OAuth app integration
	// Wire name: 'token_access_policy'
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	// Scopes that will need to be consented by end user to mint the access
	// token. If the user does not authorize the access token will not be
	// minted. Must be a subset of scopes.
	// Wire name: 'user_authorized_scopes'
	UserAuthorizedScopes []string `json:"user_authorized_scopes,omitempty"`
}

func (st UpdateCustomAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCustomAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCustomAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.UpdateCustomAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCustomAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCustomAppIntegrationToPb(st *UpdateCustomAppIntegration) (*oauth2pb.UpdateCustomAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.UpdateCustomAppIntegrationPb{}
	pb.IntegrationId = st.IntegrationId
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes
	tokenAccessPolicyPb, err := TokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}
	pb.UserAuthorizedScopes = st.UserAuthorizedScopes

	return pb, nil
}

func UpdateCustomAppIntegrationFromPb(pb *oauth2pb.UpdateCustomAppIntegrationPb) (*UpdateCustomAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCustomAppIntegration{}
	st.IntegrationId = pb.IntegrationId
	st.RedirectUrls = pb.RedirectUrls
	st.Scopes = pb.Scopes
	tokenAccessPolicyField, err := TokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}
	st.UserAuthorizedScopes = pb.UserAuthorizedScopes

	return st, nil
}

type UpdatePublishedAppIntegration struct {
	IntegrationId string `json:"-" tf:"-"`
	// Token access policy to be updated in the published OAuth app integration
	// Wire name: 'token_access_policy'
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}

func (st UpdatePublishedAppIntegration) MarshalJSON() ([]byte, error) {
	pb, err := UpdatePublishedAppIntegrationToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdatePublishedAppIntegration) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.UpdatePublishedAppIntegrationPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdatePublishedAppIntegrationFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdatePublishedAppIntegrationToPb(st *UpdatePublishedAppIntegration) (*oauth2pb.UpdatePublishedAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.UpdatePublishedAppIntegrationPb{}
	pb.IntegrationId = st.IntegrationId
	tokenAccessPolicyPb, err := TokenAccessPolicyToPb(st.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyPb != nil {
		pb.TokenAccessPolicy = tokenAccessPolicyPb
	}

	return pb, nil
}

func UpdatePublishedAppIntegrationFromPb(pb *oauth2pb.UpdatePublishedAppIntegrationPb) (*UpdatePublishedAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePublishedAppIntegration{}
	st.IntegrationId = pb.IntegrationId
	tokenAccessPolicyField, err := TokenAccessPolicyFromPb(pb.TokenAccessPolicy)
	if err != nil {
		return nil, err
	}
	if tokenAccessPolicyField != nil {
		st.TokenAccessPolicy = tokenAccessPolicyField
	}

	return st, nil
}

type UpdateServicePrincipalFederationPolicyRequest struct {

	// Wire name: 'policy'
	Policy FederationPolicy `json:"policy"`
	// The identifier for the federation policy.
	PolicyId string `json:"-" tf:"-"`
	// The service principal id for the federation policy.
	ServicePrincipalId int64 `json:"-" tf:"-"`
	// The field mask specifies which fields of the policy to update. To specify
	// multiple fields in the field mask, use comma as the separator (no space).
	// The special value '*' indicates that all fields should be updated (full
	// replacement). If unspecified, all fields that are set in the policy
	// provided in the update request will overwrite the corresponding fields in
	// the existing policy. Example value: 'description,oidc_policy.audiences'.
	UpdateMask      string   `json:"-" tf:"-"` //legacy
	ForceSendFields []string `json:"-" tf:"-"`
}

func (st UpdateServicePrincipalFederationPolicyRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateServicePrincipalFederationPolicyRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateServicePrincipalFederationPolicyRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &oauth2pb.UpdateServicePrincipalFederationPolicyRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateServicePrincipalFederationPolicyRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateServicePrincipalFederationPolicyRequestToPb(st *UpdateServicePrincipalFederationPolicyRequest) (*oauth2pb.UpdateServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oauth2pb.UpdateServicePrincipalFederationPolicyRequestPb{}
	policyPb, err := FederationPolicyToPb(&st.Policy)
	if err != nil {
		return nil, err
	}
	if policyPb != nil {
		pb.Policy = *policyPb
	}
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId
	pb.UpdateMask = st.UpdateMask

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateServicePrincipalFederationPolicyRequestFromPb(pb *oauth2pb.UpdateServicePrincipalFederationPolicyRequestPb) (*UpdateServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateServicePrincipalFederationPolicyRequest{}
	policyField, err := FederationPolicyFromPb(&pb.Policy)
	if err != nil {
		return nil, err
	}
	if policyField != nil {
		st.Policy = *policyField
	}
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId
	st.UpdateMask = pb.UpdateMask

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
