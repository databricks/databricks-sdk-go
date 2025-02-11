// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioningpreview

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

// unexported type that holds implementations of just CredentialsPreview API methods
type credentialsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsPreviewImpl) Create(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/credentials", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &credential)
	return &credential, err
}

func (a *credentialsPreviewImpl) Delete(ctx context.Context, request DeleteCredentialRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), request.CredentialsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *credentialsPreviewImpl) Get(ctx context.Context, request GetCredentialRequest) (*Credential, error) {
	var credential Credential
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), request.CredentialsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &credential)
	return &credential, err
}

func (a *credentialsPreviewImpl) List(ctx context.Context) ([]Credential, error) {
	var credentialList []Credential
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/credentials", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &credentialList)
	return credentialList, err
}

// unexported type that holds implementations of just EncryptionKeysPreview API methods
type encryptionKeysPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *encryptionKeysPreviewImpl) Create(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &customerManagedKey)
	return &customerManagedKey, err
}

func (a *encryptionKeysPreviewImpl) Delete(ctx context.Context, request DeleteEncryptionKeyRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), request.CustomerManagedKeyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *encryptionKeysPreviewImpl) Get(ctx context.Context, request GetEncryptionKeyRequest) (*CustomerManagedKey, error) {
	var customerManagedKey CustomerManagedKey
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), request.CustomerManagedKeyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &customerManagedKey)
	return &customerManagedKey, err
}

func (a *encryptionKeysPreviewImpl) List(ctx context.Context) ([]CustomerManagedKey, error) {
	var customerManagedKeyList []CustomerManagedKey
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &customerManagedKeyList)
	return customerManagedKeyList, err
}

// unexported type that holds implementations of just NetworksPreview API methods
type networksPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *networksPreviewImpl) Create(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/networks", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &network)
	return &network, err
}

func (a *networksPreviewImpl) Delete(ctx context.Context, request DeleteNetworkRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), request.NetworkId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *networksPreviewImpl) Get(ctx context.Context, request GetNetworkRequest) (*Network, error) {
	var network Network
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), request.NetworkId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &network)
	return &network, err
}

func (a *networksPreviewImpl) List(ctx context.Context) ([]Network, error) {
	var networkList []Network
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/networks", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &networkList)
	return networkList, err
}

// unexported type that holds implementations of just PrivateAccessPreview API methods
type privateAccessPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *privateAccessPreviewImpl) Create(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

func (a *privateAccessPreviewImpl) Delete(ctx context.Context, request DeletePrivateAccesRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *privateAccessPreviewImpl) Get(ctx context.Context, request GetPrivateAccesRequest) (*PrivateAccessSettings, error) {
	var privateAccessSettings PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &privateAccessSettings)
	return &privateAccessSettings, err
}

func (a *privateAccessPreviewImpl) List(ctx context.Context) ([]PrivateAccessSettings, error) {
	var privateAccessSettingsList []PrivateAccessSettings
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &privateAccessSettingsList)
	return privateAccessSettingsList, err
}

func (a *privateAccessPreviewImpl) Replace(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	var replaceResponse ReplaceResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPut, path, headers, queryParams, request, &replaceResponse)
	return err
}

// unexported type that holds implementations of just StoragePreview API methods
type storagePreviewImpl struct {
	client *client.DatabricksClient
}

func (a *storagePreviewImpl) Create(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &storageConfiguration)
	return &storageConfiguration, err
}

func (a *storagePreviewImpl) Delete(ctx context.Context, request DeleteStorageRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), request.StorageConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *storagePreviewImpl) Get(ctx context.Context, request GetStorageRequest) (*StorageConfiguration, error) {
	var storageConfiguration StorageConfiguration
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), request.StorageConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &storageConfiguration)
	return &storageConfiguration, err
}

func (a *storagePreviewImpl) List(ctx context.Context) ([]StorageConfiguration, error) {
	var storageConfigurationList []StorageConfiguration
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &storageConfigurationList)
	return storageConfigurationList, err
}

// unexported type that holds implementations of just VpcEndpointsPreview API methods
type vpcEndpointsPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *vpcEndpointsPreviewImpl) Create(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

func (a *vpcEndpointsPreviewImpl) Delete(ctx context.Context, request DeleteVpcEndpointRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), request.VpcEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *vpcEndpointsPreviewImpl) Get(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {
	var vpcEndpoint VpcEndpoint
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), request.VpcEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &vpcEndpoint)
	return &vpcEndpoint, err
}

func (a *vpcEndpointsPreviewImpl) List(ctx context.Context) ([]VpcEndpoint, error) {
	var vpcEndpointList []VpcEndpoint
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &vpcEndpointList)
	return vpcEndpointList, err
}

// unexported type that holds implementations of just WorkspacesPreview API methods
type workspacesPreviewImpl struct {
	client *client.DatabricksClient
}

func (a *workspacesPreviewImpl) Create(ctx context.Context, request CreateWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPost, path, headers, queryParams, request, &workspace)
	return &workspace, err
}

func (a *workspacesPreviewImpl) Delete(ctx context.Context, request DeleteWorkspaceRequest) error {
	var deleteResponse DeleteResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodDelete, path, headers, queryParams, request, &deleteResponse)
	return err
}

func (a *workspacesPreviewImpl) Get(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {
	var workspace Workspace
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, queryParams, request, &workspace)
	return &workspace, err
}

func (a *workspacesPreviewImpl) List(ctx context.Context) ([]Workspace, error) {
	var workspaceList []Workspace
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx, http.MethodGet, path, headers, nil, nil, &workspaceList)
	return workspaceList, err
}

func (a *workspacesPreviewImpl) Update(ctx context.Context, request UpdateWorkspaceRequest) error {
	var updateResponse UpdateResponse
	path := fmt.Sprintf("/api/2.0preview/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx, http.MethodPatch, path, headers, queryParams, request, &updateResponse)
	return err
}
