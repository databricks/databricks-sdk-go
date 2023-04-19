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
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &credential)
	return &credential, err
}

func (a *credentialsImpl) Delete(ctx context.Context, request DeleteCredentialRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), request.CredentialsId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *credentialsImpl) Get(ctx context.Context, request GetCredentialRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), request.CredentialsId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &credential)
	return &credential, err
}

func (a *credentialsImpl) List(ctx context.Context) ([]Credential, error) {
	var credentialList []Credential
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &credentialList)
	return credentialList, err
}

// unexported type that holds implementations of just EncryptionKeys API methods
type encryptionKeysImpl struct {
	client *client.DatabricksClient
}

func (a *encryptionKeysImpl) Create(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

func (a *encryptionKeysImpl) Delete(ctx context.Context, request DeleteEncryptionKeyRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), request.CustomerManagedKeyId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *encryptionKeysImpl) Get(ctx context.Context, request GetEncryptionKeyRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), request.CustomerManagedKeyId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &customerManagedKey)
	return &customerManagedKey, err
}

func (a *encryptionKeysImpl) List(ctx context.Context) ([]CustomerManagedKey, error) {
	var customerManagedKeyList []CustomerManagedKey
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &customerManagedKeyList)
	return customerManagedKeyList, err
}

// unexported type that holds implementations of just Networks API methods
type networksImpl struct {
	client *client.DatabricksClient
}

func (a *networksImpl) Create(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &network)
	return &network, err
}

func (a *networksImpl) Delete(ctx context.Context, request DeleteNetworkRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), request.NetworkId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *networksImpl) Get(ctx context.Context, request GetNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), request.NetworkId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &network)
	return &network, err
}

func (a *networksImpl) List(ctx context.Context) ([]Network, error) {
	var networkList []Network
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &networkList)
	return networkList, err
}

// unexported type that holds implementations of just PrivateAccess API methods
type privateAccessImpl struct {
	client *client.DatabricksClient
}

func (a *privateAccessImpl) Create(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

func (a *privateAccessImpl) Delete(ctx context.Context, request DeletePrivateAccesRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *privateAccessImpl) Get(ctx context.Context, request GetPrivateAccesRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

func (a *privateAccessImpl) List(ctx context.Context) ([]PrivateAccessSettings, error) {
	var privateAccessSettingsList []PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &privateAccessSettingsList)
	return privateAccessSettingsList, err
}

func (a *privateAccessImpl) Replace(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	err := a.client.Do(ctx, http.MethodPut, path, request, nil)
	return err
}

// unexported type that holds implementations of just Storage API methods
type storageImpl struct {
	client *client.DatabricksClient
}

func (a *storageImpl) Create(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &storageConfiguration)
	return &storageConfiguration, err
}

func (a *storageImpl) Delete(ctx context.Context, request DeleteStorageRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), request.StorageConfigurationId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *storageImpl) Get(ctx context.Context, request GetStorageRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), request.StorageConfigurationId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &storageConfiguration)
	return &storageConfiguration, err
}

func (a *storageImpl) List(ctx context.Context) ([]StorageConfiguration, error) {
	var storageConfigurationList []StorageConfiguration
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &storageConfigurationList)
	return storageConfigurationList, err
}

// unexported type that holds implementations of just VpcEndpoints API methods
type vpcEndpointsImpl struct {
	client *client.DatabricksClient
}

func (a *vpcEndpointsImpl) Create(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

func (a *vpcEndpointsImpl) Delete(ctx context.Context, request DeleteVpcEndpointRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), request.VpcEndpointId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *vpcEndpointsImpl) Get(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), request.VpcEndpointId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

func (a *vpcEndpointsImpl) List(ctx context.Context) ([]VpcEndpoint, error) {
	var vpcEndpointList []VpcEndpoint
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &vpcEndpointList)
	return vpcEndpointList, err
}

// unexported type that holds implementations of just Workspaces API methods
type workspacesImpl struct {
	client *client.DatabricksClient
}

func (a *workspacesImpl) Create(ctx context.Context, request CreateWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &workspace)
	return &workspace, err
}

func (a *workspacesImpl) Delete(ctx context.Context, request DeleteWorkspaceRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *workspacesImpl) Get(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &workspace)
	return &workspace, err
}

func (a *workspacesImpl) List(ctx context.Context) ([]Workspace, error) {
	var workspaceList []Workspace
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &workspaceList)
	return workspaceList, err
}

func (a *workspacesImpl) Update(ctx context.Context, request UpdateWorkspaceRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
