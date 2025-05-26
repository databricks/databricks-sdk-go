// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func createAccountFederationPolicyRequestToPb(st *CreateAccountFederationPolicyRequest) (*createAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createAccountFederationPolicyRequestPb{}
	pb.Policy = st.Policy
	pb.PolicyId = st.PolicyId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createAccountFederationPolicyRequestPb struct {
	Policy   FederationPolicy `json:"policy"`
	PolicyId string           `json:"-" url:"policy_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createAccountFederationPolicyRequestFromPb(pb *createAccountFederationPolicyRequestPb) (*CreateAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateAccountFederationPolicyRequest{}
	st.Policy = pb.Policy
	st.PolicyId = pb.PolicyId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createAccountFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createAccountFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createCustomAppIntegrationToPb(st *CreateCustomAppIntegration) (*createCustomAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCustomAppIntegrationPb{}
	pb.Confidential = st.Confidential
	pb.Name = st.Name
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes
	pb.TokenAccessPolicy = st.TokenAccessPolicy
	pb.UserAuthorizedScopes = st.UserAuthorizedScopes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createCustomAppIntegrationPb struct {
	Confidential         bool               `json:"confidential,omitempty"`
	Name                 string             `json:"name,omitempty"`
	RedirectUrls         []string           `json:"redirect_urls,omitempty"`
	Scopes               []string           `json:"scopes,omitempty"`
	TokenAccessPolicy    *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	UserAuthorizedScopes []string           `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCustomAppIntegrationFromPb(pb *createCustomAppIntegrationPb) (*CreateCustomAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomAppIntegration{}
	st.Confidential = pb.Confidential
	st.Name = pb.Name
	st.RedirectUrls = pb.RedirectUrls
	st.Scopes = pb.Scopes
	st.TokenAccessPolicy = pb.TokenAccessPolicy
	st.UserAuthorizedScopes = pb.UserAuthorizedScopes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCustomAppIntegrationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCustomAppIntegrationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createCustomAppIntegrationOutputToPb(st *CreateCustomAppIntegrationOutput) (*createCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCustomAppIntegrationOutputPb{}
	pb.ClientId = st.ClientId
	pb.ClientSecret = st.ClientSecret
	pb.IntegrationId = st.IntegrationId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createCustomAppIntegrationOutputPb struct {
	ClientId      string `json:"client_id,omitempty"`
	ClientSecret  string `json:"client_secret,omitempty"`
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCustomAppIntegrationOutputFromPb(pb *createCustomAppIntegrationOutputPb) (*CreateCustomAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomAppIntegrationOutput{}
	st.ClientId = pb.ClientId
	st.ClientSecret = pb.ClientSecret
	st.IntegrationId = pb.IntegrationId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCustomAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCustomAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPublishedAppIntegrationToPb(st *CreatePublishedAppIntegration) (*createPublishedAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPublishedAppIntegrationPb{}
	pb.AppId = st.AppId
	pb.TokenAccessPolicy = st.TokenAccessPolicy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPublishedAppIntegrationPb struct {
	AppId             string             `json:"app_id,omitempty"`
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPublishedAppIntegrationFromPb(pb *createPublishedAppIntegrationPb) (*CreatePublishedAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePublishedAppIntegration{}
	st.AppId = pb.AppId
	st.TokenAccessPolicy = pb.TokenAccessPolicy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPublishedAppIntegrationPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPublishedAppIntegrationPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPublishedAppIntegrationOutputToPb(st *CreatePublishedAppIntegrationOutput) (*createPublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPublishedAppIntegrationOutputPb{}
	pb.IntegrationId = st.IntegrationId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPublishedAppIntegrationOutputPb struct {
	IntegrationId string `json:"integration_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPublishedAppIntegrationOutputFromPb(pb *createPublishedAppIntegrationOutputPb) (*CreatePublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePublishedAppIntegrationOutput{}
	st.IntegrationId = pb.IntegrationId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPublishedAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPublishedAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createServicePrincipalFederationPolicyRequestToPb(st *CreateServicePrincipalFederationPolicyRequest) (*createServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServicePrincipalFederationPolicyRequestPb{}
	pb.Policy = st.Policy
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createServicePrincipalFederationPolicyRequestPb struct {
	Policy             FederationPolicy `json:"policy"`
	PolicyId           string           `json:"-" url:"policy_id,omitempty"`
	ServicePrincipalId int64            `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createServicePrincipalFederationPolicyRequestFromPb(pb *createServicePrincipalFederationPolicyRequestPb) (*CreateServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalFederationPolicyRequest{}
	st.Policy = pb.Policy
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createServicePrincipalFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServicePrincipalFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createServicePrincipalSecretRequestToPb(st *CreateServicePrincipalSecretRequest) (*createServicePrincipalSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServicePrincipalSecretRequestPb{}
	pb.Lifetime = st.Lifetime
	pb.ServicePrincipalId = st.ServicePrincipalId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createServicePrincipalSecretRequestPb struct {
	Lifetime           string `json:"lifetime,omitempty"`
	ServicePrincipalId int64  `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createServicePrincipalSecretRequestFromPb(pb *createServicePrincipalSecretRequestPb) (*CreateServicePrincipalSecretRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServicePrincipalSecretRequest{}
	st.Lifetime = pb.Lifetime
	st.ServicePrincipalId = pb.ServicePrincipalId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createServicePrincipalSecretRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServicePrincipalSecretRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createServicePrincipalSecretResponseToPb(st *CreateServicePrincipalSecretResponse) (*createServicePrincipalSecretResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServicePrincipalSecretResponsePb{}
	pb.CreateTime = st.CreateTime
	pb.ExpireTime = st.ExpireTime
	pb.Id = st.Id
	pb.Secret = st.Secret
	pb.SecretHash = st.SecretHash
	pb.Status = st.Status
	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createServicePrincipalSecretResponsePb struct {
	CreateTime string `json:"create_time,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	Id         string `json:"id,omitempty"`
	Secret     string `json:"secret,omitempty"`
	SecretHash string `json:"secret_hash,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createServicePrincipalSecretResponseFromPb(pb *createServicePrincipalSecretResponsePb) (*CreateServicePrincipalSecretResponse, error) {
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createServicePrincipalSecretResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServicePrincipalSecretResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteAccountFederationPolicyRequestToPb(st *DeleteAccountFederationPolicyRequest) (*deleteAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteAccountFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

type deleteAccountFederationPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

func deleteAccountFederationPolicyRequestFromPb(pb *deleteAccountFederationPolicyRequestPb) (*DeleteAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteAccountFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

func deleteCustomAppIntegrationOutputToPb(st *DeleteCustomAppIntegrationOutput) (*deleteCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCustomAppIntegrationOutputPb{}

	return pb, nil
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

func deleteCustomAppIntegrationRequestToPb(st *DeleteCustomAppIntegrationRequest) (*deleteCustomAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCustomAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

type deleteCustomAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

func deleteCustomAppIntegrationRequestFromPb(pb *deleteCustomAppIntegrationRequestPb) (*DeleteCustomAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCustomAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

func deletePublishedAppIntegrationOutputToPb(st *DeletePublishedAppIntegrationOutput) (*deletePublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePublishedAppIntegrationOutputPb{}

	return pb, nil
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

func deletePublishedAppIntegrationRequestToPb(st *DeletePublishedAppIntegrationRequest) (*deletePublishedAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePublishedAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

type deletePublishedAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

func deletePublishedAppIntegrationRequestFromPb(pb *deletePublishedAppIntegrationRequestPb) (*DeletePublishedAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePublishedAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
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

func deleteServicePrincipalFederationPolicyRequestToPb(st *DeleteServicePrincipalFederationPolicyRequest) (*deleteServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServicePrincipalFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId

	return pb, nil
}

type deleteServicePrincipalFederationPolicyRequestPb struct {
	PolicyId           string `json:"-" url:"-"`
	ServicePrincipalId int64  `json:"-" url:"-"`
}

func deleteServicePrincipalFederationPolicyRequestFromPb(pb *deleteServicePrincipalFederationPolicyRequestPb) (*DeleteServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId

	return st, nil
}

func deleteServicePrincipalSecretRequestToPb(st *DeleteServicePrincipalSecretRequest) (*deleteServicePrincipalSecretRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServicePrincipalSecretRequestPb{}
	pb.SecretId = st.SecretId
	pb.ServicePrincipalId = st.ServicePrincipalId

	return pb, nil
}

type deleteServicePrincipalSecretRequestPb struct {
	SecretId           string `json:"-" url:"-"`
	ServicePrincipalId int64  `json:"-" url:"-"`
}

func deleteServicePrincipalSecretRequestFromPb(pb *deleteServicePrincipalSecretRequestPb) (*DeleteServicePrincipalSecretRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServicePrincipalSecretRequest{}
	st.SecretId = pb.SecretId
	st.ServicePrincipalId = pb.ServicePrincipalId

	return st, nil
}

func federationPolicyToPb(st *FederationPolicy) (*federationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &federationPolicyPb{}
	pb.CreateTime = st.CreateTime
	pb.Description = st.Description
	pb.Name = st.Name
	pb.OidcPolicy = st.OidcPolicy
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId
	pb.Uid = st.Uid
	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type federationPolicyPb struct {
	CreateTime         string                `json:"create_time,omitempty"`
	Description        string                `json:"description,omitempty"`
	Name               string                `json:"name,omitempty"`
	OidcPolicy         *OidcFederationPolicy `json:"oidc_policy,omitempty"`
	PolicyId           string                `json:"policy_id,omitempty"`
	ServicePrincipalId int64                 `json:"service_principal_id,omitempty"`
	Uid                string                `json:"uid,omitempty"`
	UpdateTime         string                `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func federationPolicyFromPb(pb *federationPolicyPb) (*FederationPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FederationPolicy{}
	st.CreateTime = pb.CreateTime
	st.Description = pb.Description
	st.Name = pb.Name
	st.OidcPolicy = pb.OidcPolicy
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId
	st.Uid = pb.Uid
	st.UpdateTime = pb.UpdateTime

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *federationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st federationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getAccountFederationPolicyRequestToPb(st *GetAccountFederationPolicyRequest) (*getAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getAccountFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId

	return pb, nil
}

type getAccountFederationPolicyRequestPb struct {
	PolicyId string `json:"-" url:"-"`
}

func getAccountFederationPolicyRequestFromPb(pb *getAccountFederationPolicyRequestPb) (*GetAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetAccountFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId

	return st, nil
}

func getCustomAppIntegrationOutputToPb(st *GetCustomAppIntegrationOutput) (*getCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCustomAppIntegrationOutputPb{}
	pb.ClientId = st.ClientId
	pb.Confidential = st.Confidential
	pb.CreateTime = st.CreateTime
	pb.CreatedBy = st.CreatedBy
	pb.CreatorUsername = st.CreatorUsername
	pb.IntegrationId = st.IntegrationId
	pb.Name = st.Name
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes
	pb.TokenAccessPolicy = st.TokenAccessPolicy
	pb.UserAuthorizedScopes = st.UserAuthorizedScopes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getCustomAppIntegrationOutputPb struct {
	ClientId             string             `json:"client_id,omitempty"`
	Confidential         bool               `json:"confidential,omitempty"`
	CreateTime           string             `json:"create_time,omitempty"`
	CreatedBy            int64              `json:"created_by,omitempty"`
	CreatorUsername      string             `json:"creator_username,omitempty"`
	IntegrationId        string             `json:"integration_id,omitempty"`
	Name                 string             `json:"name,omitempty"`
	RedirectUrls         []string           `json:"redirect_urls,omitempty"`
	Scopes               []string           `json:"scopes,omitempty"`
	TokenAccessPolicy    *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	UserAuthorizedScopes []string           `json:"user_authorized_scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCustomAppIntegrationOutputFromPb(pb *getCustomAppIntegrationOutputPb) (*GetCustomAppIntegrationOutput, error) {
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
	st.TokenAccessPolicy = pb.TokenAccessPolicy
	st.UserAuthorizedScopes = pb.UserAuthorizedScopes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCustomAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCustomAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getCustomAppIntegrationRequestToPb(st *GetCustomAppIntegrationRequest) (*getCustomAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCustomAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

type getCustomAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

func getCustomAppIntegrationRequestFromPb(pb *getCustomAppIntegrationRequestPb) (*GetCustomAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

func getCustomAppIntegrationsOutputToPb(st *GetCustomAppIntegrationsOutput) (*getCustomAppIntegrationsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCustomAppIntegrationsOutputPb{}
	pb.Apps = st.Apps
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getCustomAppIntegrationsOutputPb struct {
	Apps          []GetCustomAppIntegrationOutput `json:"apps,omitempty"`
	NextPageToken string                          `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getCustomAppIntegrationsOutputFromPb(pb *getCustomAppIntegrationsOutputPb) (*GetCustomAppIntegrationsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomAppIntegrationsOutput{}
	st.Apps = pb.Apps
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getCustomAppIntegrationsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getCustomAppIntegrationsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getPublishedAppIntegrationOutputToPb(st *GetPublishedAppIntegrationOutput) (*getPublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppIntegrationOutputPb{}
	pb.AppId = st.AppId
	pb.CreateTime = st.CreateTime
	pb.CreatedBy = st.CreatedBy
	pb.IntegrationId = st.IntegrationId
	pb.Name = st.Name
	pb.TokenAccessPolicy = st.TokenAccessPolicy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPublishedAppIntegrationOutputPb struct {
	AppId             string             `json:"app_id,omitempty"`
	CreateTime        string             `json:"create_time,omitempty"`
	CreatedBy         int64              `json:"created_by,omitempty"`
	IntegrationId     string             `json:"integration_id,omitempty"`
	Name              string             `json:"name,omitempty"`
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedAppIntegrationOutputFromPb(pb *getPublishedAppIntegrationOutputPb) (*GetPublishedAppIntegrationOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationOutput{}
	st.AppId = pb.AppId
	st.CreateTime = pb.CreateTime
	st.CreatedBy = pb.CreatedBy
	st.IntegrationId = pb.IntegrationId
	st.Name = pb.Name
	st.TokenAccessPolicy = pb.TokenAccessPolicy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedAppIntegrationOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedAppIntegrationOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getPublishedAppIntegrationRequestToPb(st *GetPublishedAppIntegrationRequest) (*getPublishedAppIntegrationRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppIntegrationRequestPb{}
	pb.IntegrationId = st.IntegrationId

	return pb, nil
}

type getPublishedAppIntegrationRequestPb struct {
	IntegrationId string `json:"-" url:"-"`
}

func getPublishedAppIntegrationRequestFromPb(pb *getPublishedAppIntegrationRequestPb) (*GetPublishedAppIntegrationRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationRequest{}
	st.IntegrationId = pb.IntegrationId

	return st, nil
}

func getPublishedAppIntegrationsOutputToPb(st *GetPublishedAppIntegrationsOutput) (*getPublishedAppIntegrationsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppIntegrationsOutputPb{}
	pb.Apps = st.Apps
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPublishedAppIntegrationsOutputPb struct {
	Apps          []GetPublishedAppIntegrationOutput `json:"apps,omitempty"`
	NextPageToken string                             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedAppIntegrationsOutputFromPb(pb *getPublishedAppIntegrationsOutputPb) (*GetPublishedAppIntegrationsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppIntegrationsOutput{}
	st.Apps = pb.Apps
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedAppIntegrationsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedAppIntegrationsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getPublishedAppsOutputToPb(st *GetPublishedAppsOutput) (*getPublishedAppsOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPublishedAppsOutputPb{}
	pb.Apps = st.Apps
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPublishedAppsOutputPb struct {
	Apps          []PublishedAppOutput `json:"apps,omitempty"`
	NextPageToken string               `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPublishedAppsOutputFromPb(pb *getPublishedAppsOutputPb) (*GetPublishedAppsOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPublishedAppsOutput{}
	st.Apps = pb.Apps
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPublishedAppsOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPublishedAppsOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getServicePrincipalFederationPolicyRequestToPb(st *GetServicePrincipalFederationPolicyRequest) (*getServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServicePrincipalFederationPolicyRequestPb{}
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId

	return pb, nil
}

type getServicePrincipalFederationPolicyRequestPb struct {
	PolicyId           string `json:"-" url:"-"`
	ServicePrincipalId int64  `json:"-" url:"-"`
}

func getServicePrincipalFederationPolicyRequestFromPb(pb *getServicePrincipalFederationPolicyRequestPb) (*GetServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServicePrincipalFederationPolicyRequest{}
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId

	return st, nil
}

func listAccountFederationPoliciesRequestToPb(st *ListAccountFederationPoliciesRequest) (*listAccountFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listAccountFederationPoliciesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listAccountFederationPoliciesRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listAccountFederationPoliciesRequestFromPb(pb *listAccountFederationPoliciesRequestPb) (*ListAccountFederationPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListAccountFederationPoliciesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listAccountFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listAccountFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listCustomAppIntegrationsRequestToPb(st *ListCustomAppIntegrationsRequest) (*listCustomAppIntegrationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listCustomAppIntegrationsRequestPb{}
	pb.IncludeCreatorUsername = st.IncludeCreatorUsername
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listCustomAppIntegrationsRequestPb struct {
	IncludeCreatorUsername bool   `json:"-" url:"include_creator_username,omitempty"`
	PageSize               int    `json:"-" url:"page_size,omitempty"`
	PageToken              string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listCustomAppIntegrationsRequestFromPb(pb *listCustomAppIntegrationsRequestPb) (*ListCustomAppIntegrationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListCustomAppIntegrationsRequest{}
	st.IncludeCreatorUsername = pb.IncludeCreatorUsername
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listCustomAppIntegrationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listCustomAppIntegrationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listFederationPoliciesResponseToPb(st *ListFederationPoliciesResponse) (*listFederationPoliciesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listFederationPoliciesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Policies = st.Policies

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listFederationPoliciesResponsePb struct {
	NextPageToken string             `json:"next_page_token,omitempty"`
	Policies      []FederationPolicy `json:"policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listFederationPoliciesResponseFromPb(pb *listFederationPoliciesResponsePb) (*ListFederationPoliciesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFederationPoliciesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Policies = pb.Policies

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listFederationPoliciesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listFederationPoliciesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listOAuthPublishedAppsRequestToPb(st *ListOAuthPublishedAppsRequest) (*listOAuthPublishedAppsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listOAuthPublishedAppsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listOAuthPublishedAppsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listOAuthPublishedAppsRequestFromPb(pb *listOAuthPublishedAppsRequestPb) (*ListOAuthPublishedAppsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListOAuthPublishedAppsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listOAuthPublishedAppsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listOAuthPublishedAppsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPublishedAppIntegrationsRequestToPb(st *ListPublishedAppIntegrationsRequest) (*listPublishedAppIntegrationsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPublishedAppIntegrationsRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPublishedAppIntegrationsRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPublishedAppIntegrationsRequestFromPb(pb *listPublishedAppIntegrationsRequestPb) (*ListPublishedAppIntegrationsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPublishedAppIntegrationsRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPublishedAppIntegrationsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPublishedAppIntegrationsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listServicePrincipalFederationPoliciesRequestToPb(st *ListServicePrincipalFederationPoliciesRequest) (*listServicePrincipalFederationPoliciesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalFederationPoliciesRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.ServicePrincipalId = st.ServicePrincipalId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listServicePrincipalFederationPoliciesRequestPb struct {
	PageSize           int    `json:"-" url:"page_size,omitempty"`
	PageToken          string `json:"-" url:"page_token,omitempty"`
	ServicePrincipalId int64  `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listServicePrincipalFederationPoliciesRequestFromPb(pb *listServicePrincipalFederationPoliciesRequestPb) (*ListServicePrincipalFederationPoliciesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalFederationPoliciesRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.ServicePrincipalId = pb.ServicePrincipalId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalFederationPoliciesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalFederationPoliciesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listServicePrincipalSecretsRequestToPb(st *ListServicePrincipalSecretsRequest) (*listServicePrincipalSecretsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalSecretsRequestPb{}
	pb.PageToken = st.PageToken
	pb.ServicePrincipalId = st.ServicePrincipalId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listServicePrincipalSecretsRequestPb struct {
	PageToken          string `json:"-" url:"page_token,omitempty"`
	ServicePrincipalId int64  `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listServicePrincipalSecretsRequestFromPb(pb *listServicePrincipalSecretsRequestPb) (*ListServicePrincipalSecretsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalSecretsRequest{}
	st.PageToken = pb.PageToken
	st.ServicePrincipalId = pb.ServicePrincipalId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalSecretsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalSecretsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listServicePrincipalSecretsResponseToPb(st *ListServicePrincipalSecretsResponse) (*listServicePrincipalSecretsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listServicePrincipalSecretsResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Secrets = st.Secrets

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listServicePrincipalSecretsResponsePb struct {
	NextPageToken string       `json:"next_page_token,omitempty"`
	Secrets       []SecretInfo `json:"secrets,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listServicePrincipalSecretsResponseFromPb(pb *listServicePrincipalSecretsResponsePb) (*ListServicePrincipalSecretsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListServicePrincipalSecretsResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Secrets = pb.Secrets

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listServicePrincipalSecretsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listServicePrincipalSecretsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func oidcFederationPolicyToPb(st *OidcFederationPolicy) (*oidcFederationPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &oidcFederationPolicyPb{}
	pb.Audiences = st.Audiences
	pb.Issuer = st.Issuer
	pb.JwksJson = st.JwksJson
	pb.JwksUri = st.JwksUri
	pb.Subject = st.Subject
	pb.SubjectClaim = st.SubjectClaim

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type oidcFederationPolicyPb struct {
	Audiences    []string `json:"audiences,omitempty"`
	Issuer       string   `json:"issuer,omitempty"`
	JwksJson     string   `json:"jwks_json,omitempty"`
	JwksUri      string   `json:"jwks_uri,omitempty"`
	Subject      string   `json:"subject,omitempty"`
	SubjectClaim string   `json:"subject_claim,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func oidcFederationPolicyFromPb(pb *oidcFederationPolicyPb) (*OidcFederationPolicy, error) {
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *oidcFederationPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st oidcFederationPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func publishedAppOutputToPb(st *PublishedAppOutput) (*publishedAppOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &publishedAppOutputPb{}
	pb.AppId = st.AppId
	pb.ClientId = st.ClientId
	pb.Description = st.Description
	pb.IsConfidentialClient = st.IsConfidentialClient
	pb.Name = st.Name
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type publishedAppOutputPb struct {
	AppId                string   `json:"app_id,omitempty"`
	ClientId             string   `json:"client_id,omitempty"`
	Description          string   `json:"description,omitempty"`
	IsConfidentialClient bool     `json:"is_confidential_client,omitempty"`
	Name                 string   `json:"name,omitempty"`
	RedirectUrls         []string `json:"redirect_urls,omitempty"`
	Scopes               []string `json:"scopes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func publishedAppOutputFromPb(pb *publishedAppOutputPb) (*PublishedAppOutput, error) {
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *publishedAppOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st publishedAppOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func secretInfoToPb(st *SecretInfo) (*secretInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &secretInfoPb{}
	pb.CreateTime = st.CreateTime
	pb.ExpireTime = st.ExpireTime
	pb.Id = st.Id
	pb.SecretHash = st.SecretHash
	pb.Status = st.Status
	pb.UpdateTime = st.UpdateTime

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type secretInfoPb struct {
	CreateTime string `json:"create_time,omitempty"`
	ExpireTime string `json:"expire_time,omitempty"`
	Id         string `json:"id,omitempty"`
	SecretHash string `json:"secret_hash,omitempty"`
	Status     string `json:"status,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func secretInfoFromPb(pb *secretInfoPb) (*SecretInfo, error) {
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

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *secretInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st secretInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tokenAccessPolicyToPb(st *TokenAccessPolicy) (*tokenAccessPolicyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tokenAccessPolicyPb{}
	pb.AccessTokenTtlInMinutes = st.AccessTokenTtlInMinutes
	pb.RefreshTokenTtlInMinutes = st.RefreshTokenTtlInMinutes

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tokenAccessPolicyPb struct {
	AccessTokenTtlInMinutes  int `json:"access_token_ttl_in_minutes,omitempty"`
	RefreshTokenTtlInMinutes int `json:"refresh_token_ttl_in_minutes,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tokenAccessPolicyFromPb(pb *tokenAccessPolicyPb) (*TokenAccessPolicy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TokenAccessPolicy{}
	st.AccessTokenTtlInMinutes = pb.AccessTokenTtlInMinutes
	st.RefreshTokenTtlInMinutes = pb.RefreshTokenTtlInMinutes

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tokenAccessPolicyPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tokenAccessPolicyPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateAccountFederationPolicyRequestToPb(st *UpdateAccountFederationPolicyRequest) (*updateAccountFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateAccountFederationPolicyRequestPb{}
	pb.Policy = st.Policy
	pb.PolicyId = st.PolicyId
	pb.UpdateMask = st.UpdateMask

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateAccountFederationPolicyRequestPb struct {
	Policy     FederationPolicy `json:"policy"`
	PolicyId   string           `json:"-" url:"-"`
	UpdateMask string           `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateAccountFederationPolicyRequestFromPb(pb *updateAccountFederationPolicyRequestPb) (*UpdateAccountFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateAccountFederationPolicyRequest{}
	st.Policy = pb.Policy
	st.PolicyId = pb.PolicyId
	st.UpdateMask = pb.UpdateMask

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateAccountFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateAccountFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateCustomAppIntegrationToPb(st *UpdateCustomAppIntegration) (*updateCustomAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCustomAppIntegrationPb{}
	pb.IntegrationId = st.IntegrationId
	pb.RedirectUrls = st.RedirectUrls
	pb.Scopes = st.Scopes
	pb.TokenAccessPolicy = st.TokenAccessPolicy
	pb.UserAuthorizedScopes = st.UserAuthorizedScopes

	return pb, nil
}

type updateCustomAppIntegrationPb struct {
	IntegrationId        string             `json:"-" url:"-"`
	RedirectUrls         []string           `json:"redirect_urls,omitempty"`
	Scopes               []string           `json:"scopes,omitempty"`
	TokenAccessPolicy    *TokenAccessPolicy `json:"token_access_policy,omitempty"`
	UserAuthorizedScopes []string           `json:"user_authorized_scopes,omitempty"`
}

func updateCustomAppIntegrationFromPb(pb *updateCustomAppIntegrationPb) (*UpdateCustomAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCustomAppIntegration{}
	st.IntegrationId = pb.IntegrationId
	st.RedirectUrls = pb.RedirectUrls
	st.Scopes = pb.Scopes
	st.TokenAccessPolicy = pb.TokenAccessPolicy
	st.UserAuthorizedScopes = pb.UserAuthorizedScopes

	return st, nil
}

func updateCustomAppIntegrationOutputToPb(st *UpdateCustomAppIntegrationOutput) (*updateCustomAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCustomAppIntegrationOutputPb{}

	return pb, nil
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

func updatePublishedAppIntegrationToPb(st *UpdatePublishedAppIntegration) (*updatePublishedAppIntegrationPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePublishedAppIntegrationPb{}
	pb.IntegrationId = st.IntegrationId
	pb.TokenAccessPolicy = st.TokenAccessPolicy

	return pb, nil
}

type updatePublishedAppIntegrationPb struct {
	IntegrationId     string             `json:"-" url:"-"`
	TokenAccessPolicy *TokenAccessPolicy `json:"token_access_policy,omitempty"`
}

func updatePublishedAppIntegrationFromPb(pb *updatePublishedAppIntegrationPb) (*UpdatePublishedAppIntegration, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdatePublishedAppIntegration{}
	st.IntegrationId = pb.IntegrationId
	st.TokenAccessPolicy = pb.TokenAccessPolicy

	return st, nil
}

func updatePublishedAppIntegrationOutputToPb(st *UpdatePublishedAppIntegrationOutput) (*updatePublishedAppIntegrationOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updatePublishedAppIntegrationOutputPb{}

	return pb, nil
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

func updateServicePrincipalFederationPolicyRequestToPb(st *UpdateServicePrincipalFederationPolicyRequest) (*updateServicePrincipalFederationPolicyRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateServicePrincipalFederationPolicyRequestPb{}
	pb.Policy = st.Policy
	pb.PolicyId = st.PolicyId
	pb.ServicePrincipalId = st.ServicePrincipalId
	pb.UpdateMask = st.UpdateMask

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateServicePrincipalFederationPolicyRequestPb struct {
	Policy             FederationPolicy `json:"policy"`
	PolicyId           string           `json:"-" url:"-"`
	ServicePrincipalId int64            `json:"-" url:"-"`
	UpdateMask         string           `json:"-" url:"update_mask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateServicePrincipalFederationPolicyRequestFromPb(pb *updateServicePrincipalFederationPolicyRequestPb) (*UpdateServicePrincipalFederationPolicyRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateServicePrincipalFederationPolicyRequest{}
	st.Policy = pb.Policy
	st.PolicyId = pb.PolicyId
	st.ServicePrincipalId = pb.ServicePrincipalId
	st.UpdateMask = pb.UpdateMask

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateServicePrincipalFederationPolicyRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateServicePrincipalFederationPolicyRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
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
