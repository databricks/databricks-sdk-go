// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just Credentials API methods
type credentialsImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsImpl) Create(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {

	requestPb, pbErr := createCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var credentialPb credentialPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&credentialPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := credentialFromPb(&credentialPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) Delete(ctx context.Context, request DeleteCredentialRequest) error {

	requestPb, pbErr := deleteCredentialRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), requestPb.CredentialsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *credentialsImpl) Get(ctx context.Context, request GetCredentialRequest) (*Credential, error) {

	requestPb, pbErr := getCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var credentialPb credentialPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), requestPb.CredentialsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&credentialPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := credentialFromPb(&credentialPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) List(ctx context.Context) ([]Credential, error) {

	var credentialListPb []credentialPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&credentialListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []Credential
	for _, item := range credentialListPb {
		itemResp, err := credentialFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

// unexported type that holds implementations of just EncryptionKeys API methods
type encryptionKeysImpl struct {
	client *client.DatabricksClient
}

func (a *encryptionKeysImpl) Create(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {

	requestPb, pbErr := createCustomerManagedKeyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customerManagedKeyPb customerManagedKeyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&customerManagedKeyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := customerManagedKeyFromPb(&customerManagedKeyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *encryptionKeysImpl) Delete(ctx context.Context, request DeleteEncryptionKeyRequest) error {

	requestPb, pbErr := deleteEncryptionKeyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), requestPb.CustomerManagedKeyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *encryptionKeysImpl) Get(ctx context.Context, request GetEncryptionKeyRequest) (*CustomerManagedKey, error) {

	requestPb, pbErr := getEncryptionKeyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customerManagedKeyPb customerManagedKeyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), requestPb.CustomerManagedKeyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&customerManagedKeyPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := customerManagedKeyFromPb(&customerManagedKeyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *encryptionKeysImpl) List(ctx context.Context) ([]CustomerManagedKey, error) {

	var customerManagedKeyListPb []customerManagedKeyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&customerManagedKeyListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []CustomerManagedKey
	for _, item := range customerManagedKeyListPb {
		itemResp, err := customerManagedKeyFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

// unexported type that holds implementations of just Networks API methods
type networksImpl struct {
	client *client.DatabricksClient
}

func (a *networksImpl) Create(ctx context.Context, request CreateNetworkRequest) (*Network, error) {

	requestPb, pbErr := createNetworkRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkPb networkPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&networkPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := networkFromPb(&networkPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networksImpl) Delete(ctx context.Context, request DeleteNetworkRequest) error {

	requestPb, pbErr := deleteNetworkRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), requestPb.NetworkId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *networksImpl) Get(ctx context.Context, request GetNetworkRequest) (*Network, error) {

	requestPb, pbErr := getNetworkRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkPb networkPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), requestPb.NetworkId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&networkPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := networkFromPb(&networkPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networksImpl) List(ctx context.Context) ([]Network, error) {

	var networkListPb []networkPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&networkListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []Network
	for _, item := range networkListPb {
		itemResp, err := networkFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

// unexported type that holds implementations of just PrivateAccess API methods
type privateAccessImpl struct {
	client *client.DatabricksClient
}

func (a *privateAccessImpl) Create(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {

	requestPb, pbErr := upsertPrivateAccessSettingsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var privateAccessSettingsPb privateAccessSettingsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&privateAccessSettingsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := privateAccessSettingsFromPb(&privateAccessSettingsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *privateAccessImpl) Delete(ctx context.Context, request DeletePrivateAccesRequest) error {

	requestPb, pbErr := deletePrivateAccesRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), requestPb.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *privateAccessImpl) Get(ctx context.Context, request GetPrivateAccesRequest) (*PrivateAccessSettings, error) {

	requestPb, pbErr := getPrivateAccesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var privateAccessSettingsPb privateAccessSettingsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), requestPb.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&privateAccessSettingsPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := privateAccessSettingsFromPb(&privateAccessSettingsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *privateAccessImpl) List(ctx context.Context) ([]PrivateAccessSettings, error) {

	var privateAccessSettingsListPb []privateAccessSettingsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&privateAccessSettingsListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []PrivateAccessSettings
	for _, item := range privateAccessSettingsListPb {
		itemResp, err := privateAccessSettingsFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

func (a *privateAccessImpl) Replace(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {

	requestPb, pbErr := upsertPrivateAccessSettingsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var replaceResponsePb replaceResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), requestPb.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		(*requestPb),
		&replaceResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

// unexported type that holds implementations of just Storage API methods
type storageImpl struct {
	client *client.DatabricksClient
}

func (a *storageImpl) Create(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {

	requestPb, pbErr := createStorageConfigurationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var storageConfigurationPb storageConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&storageConfigurationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := storageConfigurationFromPb(&storageConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *storageImpl) Delete(ctx context.Context, request DeleteStorageRequest) error {

	requestPb, pbErr := deleteStorageRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), requestPb.StorageConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *storageImpl) Get(ctx context.Context, request GetStorageRequest) (*StorageConfiguration, error) {

	requestPb, pbErr := getStorageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var storageConfigurationPb storageConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), requestPb.StorageConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&storageConfigurationPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := storageConfigurationFromPb(&storageConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *storageImpl) List(ctx context.Context) ([]StorageConfiguration, error) {

	var storageConfigurationListPb []storageConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&storageConfigurationListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []StorageConfiguration
	for _, item := range storageConfigurationListPb {
		itemResp, err := storageConfigurationFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

// unexported type that holds implementations of just VpcEndpoints API methods
type vpcEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *vpcEndpointsImpl) Create(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error) {

	requestPb, pbErr := createVpcEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vpcEndpointPb vpcEndpointPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&vpcEndpointPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := vpcEndpointFromPb(&vpcEndpointPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vpcEndpointsImpl) Delete(ctx context.Context, request DeleteVpcEndpointRequest) error {

	requestPb, pbErr := deleteVpcEndpointRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), requestPb.VpcEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *vpcEndpointsImpl) Get(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {

	requestPb, pbErr := getVpcEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vpcEndpointPb vpcEndpointPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), requestPb.VpcEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&vpcEndpointPb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := vpcEndpointFromPb(&vpcEndpointPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vpcEndpointsImpl) List(ctx context.Context) ([]VpcEndpoint, error) {

	var vpcEndpointListPb []vpcEndpointPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&vpcEndpointListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []VpcEndpoint
	for _, item := range vpcEndpointListPb {
		itemResp, err := vpcEndpointFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

// unexported type that holds implementations of just Workspaces API methods
type workspacesImpl struct {
	client *client.DatabricksClient
}

func (a *workspacesImpl) Create(ctx context.Context, request CreateWorkspaceRequest) (*Workspace, error) {

	requestPb, pbErr := createWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspacePb workspacePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		(*requestPb),
		&workspacePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := workspaceFromPb(&workspacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspacesImpl) Delete(ctx context.Context, request DeleteWorkspaceRequest) error {

	requestPb, pbErr := deleteWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var deleteResponsePb deleteResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), requestPb.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		(*requestPb),
		&deleteResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}

func (a *workspacesImpl) Get(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {

	requestPb, pbErr := getWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspacePb workspacePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), requestPb.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		(*requestPb),
		&workspacePb,
	)
	if err != nil {
		return nil, err
	}
	resp, err := workspaceFromPb(&workspacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspacesImpl) List(ctx context.Context) ([]Workspace, error) {

	var workspaceListPb []workspacePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&workspaceListPb,
	)
	if err != nil {
		return nil, err
	}
	var resp []Workspace
	for _, item := range workspaceListPb {
		itemResp, err := workspaceFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

func (a *workspacesImpl) Update(ctx context.Context, request UpdateWorkspaceRequest) error {

	requestPb, pbErr := updateWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	var updateResponsePb updateResponsePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), requestPb.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(
		ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		(*requestPb),
		&updateResponsePb,
	)
	if err != nil {
		return err
	}

	return err
}
