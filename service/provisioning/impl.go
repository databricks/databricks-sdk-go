// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/service/provisioning/provisioningpb"
)

// unexported type that holds implementations of just Credentials API methods
type credentialsImpl struct {
	client *client.DatabricksClient
}

func (a *credentialsImpl) Create(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	requestPb, pbErr := CreateCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var credentialPb provisioningpb.CredentialPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&credentialPb,
	)
	resp, err := CredentialFromPb(&credentialPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) Delete(ctx context.Context, request DeleteCredentialRequest) error {
	requestPb, pbErr := DeleteCredentialRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), request.CredentialsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *credentialsImpl) Get(ctx context.Context, request GetCredentialRequest) (*Credential, error) {
	requestPb, pbErr := GetCredentialRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var credentialPb provisioningpb.CredentialPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials/%v", a.client.ConfiguredAccountID(), request.CredentialsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&credentialPb,
	)
	resp, err := CredentialFromPb(&credentialPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *credentialsImpl) List(ctx context.Context) ([]Credential, error) {

	var credentialListPb []provisioningpb.CredentialPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/credentials", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&credentialListPb,
	)
	var resp []Credential
	for _, item := range credentialListPb {
		itemResp, err := CredentialFromPb(&item)
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
	requestPb, pbErr := CreateCustomerManagedKeyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customerManagedKeyPb provisioningpb.CustomerManagedKeyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&customerManagedKeyPb,
	)
	resp, err := CustomerManagedKeyFromPb(&customerManagedKeyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *encryptionKeysImpl) Delete(ctx context.Context, request DeleteEncryptionKeyRequest) error {
	requestPb, pbErr := DeleteEncryptionKeyRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), request.CustomerManagedKeyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *encryptionKeysImpl) Get(ctx context.Context, request GetEncryptionKeyRequest) (*CustomerManagedKey, error) {
	requestPb, pbErr := GetEncryptionKeyRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var customerManagedKeyPb provisioningpb.CustomerManagedKeyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys/%v", a.client.ConfiguredAccountID(), request.CustomerManagedKeyId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&customerManagedKeyPb,
	)
	resp, err := CustomerManagedKeyFromPb(&customerManagedKeyPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *encryptionKeysImpl) List(ctx context.Context) ([]CustomerManagedKey, error) {

	var customerManagedKeyListPb []provisioningpb.CustomerManagedKeyPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/customer-managed-keys", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&customerManagedKeyListPb,
	)
	var resp []CustomerManagedKey
	for _, item := range customerManagedKeyListPb {
		itemResp, err := CustomerManagedKeyFromPb(&item)
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
	requestPb, pbErr := CreateNetworkRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkPb provisioningpb.NetworkPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&networkPb,
	)
	resp, err := NetworkFromPb(&networkPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networksImpl) Delete(ctx context.Context, request DeleteNetworkRequest) error {
	requestPb, pbErr := DeleteNetworkRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), request.NetworkId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *networksImpl) Get(ctx context.Context, request GetNetworkRequest) (*Network, error) {
	requestPb, pbErr := GetNetworkRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var networkPb provisioningpb.NetworkPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks/%v", a.client.ConfiguredAccountID(), request.NetworkId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&networkPb,
	)
	resp, err := NetworkFromPb(&networkPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *networksImpl) List(ctx context.Context) ([]Network, error) {

	var networkListPb []provisioningpb.NetworkPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/networks", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&networkListPb,
	)
	var resp []Network
	for _, item := range networkListPb {
		itemResp, err := NetworkFromPb(&item)
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
	requestPb, pbErr := UpsertPrivateAccessSettingsRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var privateAccessSettingsPb provisioningpb.PrivateAccessSettingsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&privateAccessSettingsPb,
	)
	resp, err := PrivateAccessSettingsFromPb(&privateAccessSettingsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *privateAccessImpl) Delete(ctx context.Context, request DeletePrivateAccesRequest) error {
	requestPb, pbErr := DeletePrivateAccesRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *privateAccessImpl) Get(ctx context.Context, request GetPrivateAccesRequest) (*PrivateAccessSettings, error) {
	requestPb, pbErr := GetPrivateAccesRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var privateAccessSettingsPb provisioningpb.PrivateAccessSettingsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&privateAccessSettingsPb,
	)
	resp, err := PrivateAccessSettingsFromPb(&privateAccessSettingsPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *privateAccessImpl) List(ctx context.Context) ([]PrivateAccessSettings, error) {

	var privateAccessSettingsListPb []provisioningpb.PrivateAccessSettingsPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&privateAccessSettingsListPb,
	)
	var resp []PrivateAccessSettings
	for _, item := range privateAccessSettingsListPb {
		itemResp, err := PrivateAccessSettingsFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

func (a *privateAccessImpl) Replace(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	requestPb, pbErr := UpsertPrivateAccessSettingsRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/private-access-settings/%v", a.client.ConfiguredAccountID(), request.PrivateAccessSettingsId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPut,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

// unexported type that holds implementations of just Storage API methods
type storageImpl struct {
	client *client.DatabricksClient
}

func (a *storageImpl) Create(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	requestPb, pbErr := CreateStorageConfigurationRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var storageConfigurationPb provisioningpb.StorageConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&storageConfigurationPb,
	)
	resp, err := StorageConfigurationFromPb(&storageConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *storageImpl) Delete(ctx context.Context, request DeleteStorageRequest) error {
	requestPb, pbErr := DeleteStorageRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), request.StorageConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *storageImpl) Get(ctx context.Context, request GetStorageRequest) (*StorageConfiguration, error) {
	requestPb, pbErr := GetStorageRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var storageConfigurationPb provisioningpb.StorageConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations/%v", a.client.ConfiguredAccountID(), request.StorageConfigurationId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&storageConfigurationPb,
	)
	resp, err := StorageConfigurationFromPb(&storageConfigurationPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *storageImpl) List(ctx context.Context) ([]StorageConfiguration, error) {

	var storageConfigurationListPb []provisioningpb.StorageConfigurationPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/storage-configurations", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&storageConfigurationListPb,
	)
	var resp []StorageConfiguration
	for _, item := range storageConfigurationListPb {
		itemResp, err := StorageConfigurationFromPb(&item)
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
	requestPb, pbErr := CreateVpcEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vpcEndpointPb provisioningpb.VpcEndpointPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&vpcEndpointPb,
	)
	resp, err := VpcEndpointFromPb(&vpcEndpointPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vpcEndpointsImpl) Delete(ctx context.Context, request DeleteVpcEndpointRequest) error {
	requestPb, pbErr := DeleteVpcEndpointRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), request.VpcEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *vpcEndpointsImpl) Get(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {
	requestPb, pbErr := GetVpcEndpointRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var vpcEndpointPb provisioningpb.VpcEndpointPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints/%v", a.client.ConfiguredAccountID(), request.VpcEndpointId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&vpcEndpointPb,
	)
	resp, err := VpcEndpointFromPb(&vpcEndpointPb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *vpcEndpointsImpl) List(ctx context.Context) ([]VpcEndpoint, error) {

	var vpcEndpointListPb []provisioningpb.VpcEndpointPb
	path := fmt.Sprintf("/api/2.0/accounts/%v/vpc-endpoints", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&vpcEndpointListPb,
	)
	var resp []VpcEndpoint
	for _, item := range vpcEndpointListPb {
		itemResp, err := VpcEndpointFromPb(&item)
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
	requestPb, pbErr := CreateWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspacePb provisioningpb.WorkspacePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.ConfiguredAccountID())
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPost,
		path,
		headers,
		queryParams,
		requestPb,
		&workspacePb,
	)
	resp, err := WorkspaceFromPb(&workspacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspacesImpl) Delete(ctx context.Context, request DeleteWorkspaceRequest) error {
	requestPb, pbErr := DeleteWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodDelete,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}

func (a *workspacesImpl) Get(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {
	requestPb, pbErr := GetWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return nil, pbErr
	}

	var workspacePb provisioningpb.WorkspacePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		queryParams,
		requestPb,
		&workspacePb,
	)
	resp, err := WorkspaceFromPb(&workspacePb)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (a *workspacesImpl) List(ctx context.Context) ([]Workspace, error) {

	var workspaceListPb []provisioningpb.WorkspacePb
	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces", a.client.ConfiguredAccountID())

	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodGet,
		path,
		headers,
		nil,
		nil,
		&workspaceListPb,
	)
	var resp []Workspace
	for _, item := range workspaceListPb {
		itemResp, err := WorkspaceFromPb(&item)
		if err != nil {
			return nil, err
		}
		resp = append(resp, *itemResp)
	}

	return resp, err
}

func (a *workspacesImpl) Update(ctx context.Context, request UpdateWorkspaceRequest) error {
	requestPb, pbErr := UpdateWorkspaceRequestToPb(&request)
	if pbErr != nil {
		return pbErr
	}

	path := fmt.Sprintf("/api/2.0/accounts/%v/workspaces/%v", a.client.ConfiguredAccountID(), request.WorkspaceId)
	queryParams := make(map[string]any)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	err := a.client.Do(ctx,
		http.MethodPatch,
		path,
		headers,
		queryParams,
		requestPb,
		nil,
	)

	return err
}
