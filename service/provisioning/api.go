// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

// These APIs allow you to manage Credentials, Encryption Keys, Networks, Private Access, Storage, Vpc Endpoints, Workspaces, etc.
package provisioning

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewCredentials(client *client.DatabricksClient) *CredentialsAPI {
	return &CredentialsAPI{
		impl: &credentialsImpl{
			client: client,
		},
	}
}

// These APIs manage credential configurations for this workspace. Databricks
// needs access to a cross-account service IAM role in your AWS account so that
// Databricks can deploy clusters in the appropriate VPC for the new workspace.
// A credential configuration encapsulates this role information, and its ID is
// used when creating a new workspace.
type CredentialsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CredentialsService)
	impl CredentialsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CredentialsAPI) WithImpl(impl CredentialsService) *CredentialsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Credentials API implementation
func (a *CredentialsAPI) Impl() CredentialsService {
	return a.impl
}

// Create credential configuration.
//
// Creates a Databricks credential configuration that represents cloud
// cross-account credentials for a specified account. Databricks uses this to
// set up network infrastructure properly to host Databricks clusters. For your
// AWS IAM role, you need to trust the External ID (the Databricks Account API
// account ID) in the returned credential object, and configure the required
// access policy.
//
// Save the response's `credentials_id` field, which is the ID for your new
// credential configuration object.
//
// For information about how to create a new workspace with this API, see
// [Create a new workspace using the Account API]
//
// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
func (a *CredentialsAPI) Create(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	return a.impl.Create(ctx, request)
}

// Delete credential configuration.
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialsAPI) Delete(ctx context.Context, request DeleteCredentialRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete credential configuration.
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialsAPI) DeleteByCredentialsId(ctx context.Context, credentialsId string) error {
	return a.impl.Delete(ctx, DeleteCredentialRequest{
		CredentialsId: credentialsId,
	})
}

// Get credential configuration.
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialsAPI) Get(ctx context.Context, request GetCredentialRequest) (*Credential, error) {
	return a.impl.Get(ctx, request)
}

// Get credential configuration.
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialsAPI) GetByCredentialsId(ctx context.Context, credentialsId string) (*Credential, error) {
	return a.impl.Get(ctx, GetCredentialRequest{
		CredentialsId: credentialsId,
	})
}

// Get all credential configurations.
//
// Gets all Databricks credential configurations associated with an account
// specified by ID.
func (a *CredentialsAPI) List(ctx context.Context) ([]Credential, error) {
	return a.impl.List(ctx)
}

func NewEncryptionKeys(client *client.DatabricksClient) *EncryptionKeysAPI {
	return &EncryptionKeysAPI{
		impl: &encryptionKeysImpl{
			client: client,
		},
	}
}

// These APIs manage encryption key configurations for this workspace
// (optional). A key configuration encapsulates the AWS KMS key information and
// some information about how the key configuration can be used. There are two
// possible uses for key configurations:
//
// * Managed services: A key configuration can be used to encrypt a workspace's
// notebook and secret data in the control plane, as well as Databricks SQL
// queries and query history. * Storage: A key configuration can be used to
// encrypt a workspace's DBFS and EBS data in the data plane.
//
// In both of these cases, the key configuration's ID is used when creating a
// new workspace. This Preview feature is available if your account is on the E2
// version of the platform. Updating a running workspace with workspace storage
// encryption requires that the workspace is on the E2 version of the platform.
// If you have an older workspace, it might not be on the E2 version of the
// platform. If you are not sure, contact your Databricks representative.
type EncryptionKeysAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(EncryptionKeysService)
	impl EncryptionKeysService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *EncryptionKeysAPI) WithImpl(impl EncryptionKeysService) *EncryptionKeysAPI {
	a.impl = impl
	return a
}

// Impl returns low-level EncryptionKeys API implementation
func (a *EncryptionKeysAPI) Impl() EncryptionKeysService {
	return a.impl
}

// Create encryption key configuration.
//
// Creates a customer-managed key configuration object for an account, specified
// by ID. This operation uploads a reference to a customer-managed key to
// Databricks. If the key is assigned as a workspace's customer-managed key for
// managed services, Databricks uses the key to encrypt the workspaces notebooks
// and secrets in the control plane, in addition to Databricks SQL queries and
// query history. If it is specified as a workspace's customer-managed key for
// workspace storage, the key encrypts the workspace's root S3 bucket (which
// contains the workspace's root DBFS and system data) and, optionally, cluster
// EBS volume data.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions that currently support creation of
// Databricks workspaces.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *EncryptionKeysAPI) Create(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {
	return a.impl.Create(ctx, request)
}

// Delete encryption key configuration.
//
// Deletes a customer-managed key configuration object for an account. You
// cannot delete a configuration that is associated with a running workspace.
func (a *EncryptionKeysAPI) Delete(ctx context.Context, request DeleteEncryptionKeyRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete encryption key configuration.
//
// Deletes a customer-managed key configuration object for an account. You
// cannot delete a configuration that is associated with a running workspace.
func (a *EncryptionKeysAPI) DeleteByCustomerManagedKeyId(ctx context.Context, customerManagedKeyId string) error {
	return a.impl.Delete(ctx, DeleteEncryptionKeyRequest{
		CustomerManagedKeyId: customerManagedKeyId,
	})
}

// Get encryption key configuration.
//
// Gets a customer-managed key configuration object for an account, specified by
// ID. This operation uploads a reference to a customer-managed key to
// Databricks. If assigned as a workspace's customer-managed key for managed
// services, Databricks uses the key to encrypt the workspaces notebooks and
// secrets in the control plane, in addition to Databricks SQL queries and query
// history. If it is specified as a workspace's customer-managed key for
// storage, the key encrypts the workspace's root S3 bucket (which contains the
// workspace's root DBFS and system data) and, optionally, cluster EBS volume
// data.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.",
func (a *EncryptionKeysAPI) Get(ctx context.Context, request GetEncryptionKeyRequest) (*CustomerManagedKey, error) {
	return a.impl.Get(ctx, request)
}

// Get encryption key configuration.
//
// Gets a customer-managed key configuration object for an account, specified by
// ID. This operation uploads a reference to a customer-managed key to
// Databricks. If assigned as a workspace's customer-managed key for managed
// services, Databricks uses the key to encrypt the workspaces notebooks and
// secrets in the control plane, in addition to Databricks SQL queries and query
// history. If it is specified as a workspace's customer-managed key for
// storage, the key encrypts the workspace's root S3 bucket (which contains the
// workspace's root DBFS and system data) and, optionally, cluster EBS volume
// data.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.",
func (a *EncryptionKeysAPI) GetByCustomerManagedKeyId(ctx context.Context, customerManagedKeyId string) (*CustomerManagedKey, error) {
	return a.impl.Get(ctx, GetEncryptionKeyRequest{
		CustomerManagedKeyId: customerManagedKeyId,
	})
}

// Get all encryption key configurations.
//
// Gets all customer-managed key configuration objects for an account. If the
// key is specified as a workspace's managed services customer-managed key,
// Databricks uses the key to encrypt the workspace's notebooks and secrets in
// the control plane, in addition to Databricks SQL queries and query history.
// If the key is specified as a workspace's storage customer-managed key, the
// key is used to encrypt the workspace's root S3 bucket and optionally can
// encrypt cluster EBS volumes data in the data plane.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *EncryptionKeysAPI) List(ctx context.Context) ([]CustomerManagedKey, error) {
	return a.impl.List(ctx)
}

func NewNetworks(client *client.DatabricksClient) *NetworksAPI {
	return &NetworksAPI{
		impl: &networksImpl{
			client: client,
		},
	}
}

// These APIs manage network configurations for customer-managed VPCs
// (optional). Its ID is used when creating a new workspace if you use
// customer-managed VPCs.
type NetworksAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(NetworksService)
	impl NetworksService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *NetworksAPI) WithImpl(impl NetworksService) *NetworksAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Networks API implementation
func (a *NetworksAPI) Impl() NetworksService {
	return a.impl
}

// Create network configuration.
//
// Creates a Databricks network configuration that represents an VPC and its
// resources. The VPC will be used for new Databricks clusters. This requires a
// pre-existing VPC and subnets.
func (a *NetworksAPI) Create(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	return a.impl.Create(ctx, request)
}

// Delete a network configuration.
//
// Deletes a Databricks network configuration, which represents a cloud VPC and
// its resources. You cannot delete a network that is associated with a
// workspace.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworksAPI) Delete(ctx context.Context, request DeleteNetworkRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a network configuration.
//
// Deletes a Databricks network configuration, which represents a cloud VPC and
// its resources. You cannot delete a network that is associated with a
// workspace.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworksAPI) DeleteByNetworkId(ctx context.Context, networkId string) error {
	return a.impl.Delete(ctx, DeleteNetworkRequest{
		NetworkId: networkId,
	})
}

// Get a network configuration.
//
// Gets a Databricks network configuration, which represents a cloud VPC and its
// resources.
func (a *NetworksAPI) Get(ctx context.Context, request GetNetworkRequest) (*Network, error) {
	return a.impl.Get(ctx, request)
}

// Get a network configuration.
//
// Gets a Databricks network configuration, which represents a cloud VPC and its
// resources.
func (a *NetworksAPI) GetByNetworkId(ctx context.Context, networkId string) (*Network, error) {
	return a.impl.Get(ctx, GetNetworkRequest{
		NetworkId: networkId,
	})
}

// Get all network configurations.
//
// Gets a list of all Databricks network configurations for an account,
// specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworksAPI) List(ctx context.Context) ([]Network, error) {
	return a.impl.List(ctx)
}

func NewPrivateAccess(client *client.DatabricksClient) *PrivateAccessAPI {
	return &PrivateAccessAPI{
		impl: &privateAccessImpl{
			client: client,
		},
	}
}

// These APIs manage private access settings for this account.
type PrivateAccessAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PrivateAccessService)
	impl PrivateAccessService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PrivateAccessAPI) WithImpl(impl PrivateAccessService) *PrivateAccessAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PrivateAccess API implementation
func (a *PrivateAccessAPI) Impl() PrivateAccessService {
	return a.impl
}

// Create private access settings.
//
// Creates a private access settings object, which specifies how your workspace
// is accessed over [AWS PrivateLink]. To use AWS PrivateLink, a workspace must
// have a private access settings object referenced by ID in the workspace's
// `private_access_settings_id` property.
//
// You can share one private access settings with multiple workspaces in a
// single account. However, private access settings are specific to AWS regions,
// so only workspaces in the same AWS region can use a given private access
// settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *PrivateAccessAPI) Create(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	return a.impl.Create(ctx, request)
}

// Delete a private access settings object.
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink].
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *PrivateAccessAPI) Delete(ctx context.Context, request DeletePrivateAccesRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a private access settings object.
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink].
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *PrivateAccessAPI) DeleteByPrivateAccessSettingsId(ctx context.Context, privateAccessSettingsId string) error {
	return a.impl.Delete(ctx, DeletePrivateAccesRequest{
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Get a private access settings object.
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink].
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *PrivateAccessAPI) Get(ctx context.Context, request GetPrivateAccesRequest) (*PrivateAccessSettings, error) {
	return a.impl.Get(ctx, request)
}

// Get a private access settings object.
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink].
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *PrivateAccessAPI) GetByPrivateAccessSettingsId(ctx context.Context, privateAccessSettingsId string) (*PrivateAccessSettings, error) {
	return a.impl.Get(ctx, GetPrivateAccesRequest{
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Get all private access settings objects.
//
// Gets a list of all private access settings objects for an account, specified
// by ID.
func (a *PrivateAccessAPI) List(ctx context.Context) ([]PrivateAccessSettings, error) {
	return a.impl.List(ctx)
}

// Replace private access settings.
//
// Updates an existing private access settings object, which specifies how your
// workspace is accessed over [AWS PrivateLink]. To use AWS PrivateLink, a
// workspace must have a private access settings object referenced by ID in the
// workspace's `private_access_settings_id` property.
//
// This operation completely overwrites your existing private access settings
// object attached to your workspaces. All workspaces attached to the private
// access settings are affected by any change. If `public_access_enabled`,
// `private_access_level`, or `allowed_vpc_endpoint_ids` are updated, effects of
// these changes might take several minutes to propagate to the workspace API.
//
// You can share one private access settings object with multiple workspaces in
// a single account. However, private access settings are specific to AWS
// regions, so only workspaces in the same AWS region can use a given private
// access settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *PrivateAccessAPI) Replace(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	return a.impl.Replace(ctx, request)
}

func NewStorage(client *client.DatabricksClient) *StorageAPI {
	return &StorageAPI{
		impl: &storageImpl{
			client: client,
		},
	}
}

// These APIs manage storage configurations for this workspace. A root storage
// S3 bucket in your account is required to store objects like cluster logs,
// notebook revisions, and job results. You can also use the root storage S3
// bucket for storage of non-production DBFS data. A storage configuration
// encapsulates this bucket information, and its ID is used when creating a new
// workspace.
type StorageAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(StorageService)
	impl StorageService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *StorageAPI) WithImpl(impl StorageService) *StorageAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Storage API implementation
func (a *StorageAPI) Impl() StorageService {
	return a.impl
}

// Create new storage configuration.
//
// Creates new storage configuration for an account, specified by ID. Uploads a
// storage configuration object that represents the root AWS S3 bucket in your
// account. Databricks stores related workspace assets including DBFS, cluster
// logs, and job results. For the AWS S3 bucket, you need to configure the
// required bucket policy.
//
// For information about how to create a new workspace with this API, see
// [Create a new workspace using the Account API]
//
// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
func (a *StorageAPI) Create(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	return a.impl.Create(ctx, request)
}

// Delete storage configuration.
//
// Deletes a Databricks storage configuration. You cannot delete a storage
// configuration that is associated with any workspace.
func (a *StorageAPI) Delete(ctx context.Context, request DeleteStorageRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete storage configuration.
//
// Deletes a Databricks storage configuration. You cannot delete a storage
// configuration that is associated with any workspace.
func (a *StorageAPI) DeleteByStorageConfigurationId(ctx context.Context, storageConfigurationId string) error {
	return a.impl.Delete(ctx, DeleteStorageRequest{
		StorageConfigurationId: storageConfigurationId,
	})
}

// Get storage configuration.
//
// Gets a Databricks storage configuration for an account, both specified by ID.
func (a *StorageAPI) Get(ctx context.Context, request GetStorageRequest) (*StorageConfiguration, error) {
	return a.impl.Get(ctx, request)
}

// Get storage configuration.
//
// Gets a Databricks storage configuration for an account, both specified by ID.
func (a *StorageAPI) GetByStorageConfigurationId(ctx context.Context, storageConfigurationId string) (*StorageConfiguration, error) {
	return a.impl.Get(ctx, GetStorageRequest{
		StorageConfigurationId: storageConfigurationId,
	})
}

// Get all storage configurations.
//
// Gets a list of all Databricks storage configurations for your account,
// specified by ID.
func (a *StorageAPI) List(ctx context.Context) ([]StorageConfiguration, error) {
	return a.impl.List(ctx)
}

func NewVpcEndpoints(client *client.DatabricksClient) *VpcEndpointsAPI {
	return &VpcEndpointsAPI{
		impl: &vpcEndpointsImpl{
			client: client,
		},
	}
}

// These APIs manage VPC endpoint configurations for this account.
type VpcEndpointsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(VpcEndpointsService)
	impl VpcEndpointsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *VpcEndpointsAPI) WithImpl(impl VpcEndpointsService) *VpcEndpointsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level VpcEndpoints API implementation
func (a *VpcEndpointsAPI) Impl() VpcEndpointsService {
	return a.impl
}

// Create VPC endpoint configuration.
//
// Creates a VPC endpoint configuration, which represents a [VPC endpoint]
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink].
//
// After you create the VPC endpoint configuration, the Databricks [endpoint
// service] automatically accepts the VPC endpoint.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
// [VPC endpoint]: https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html
// [endpoint service]: https://docs.aws.amazon.com/vpc/latest/privatelink/privatelink-share-your-services.html
func (a *VpcEndpointsAPI) Create(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error) {
	return a.impl.Create(ctx, request)
}

// Delete VPC endpoint configuration.
//
// Deletes a VPC endpoint configuration, which represents an [AWS VPC endpoint]
// that can communicate privately with Databricks over [AWS PrivateLink].
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [AWS VPC endpoint]: https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *VpcEndpointsAPI) Delete(ctx context.Context, request DeleteVpcEndpointRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete VPC endpoint configuration.
//
// Deletes a VPC endpoint configuration, which represents an [AWS VPC endpoint]
// that can communicate privately with Databricks over [AWS PrivateLink].
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [AWS VPC endpoint]: https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *VpcEndpointsAPI) DeleteByVpcEndpointId(ctx context.Context, vpcEndpointId string) error {
	return a.impl.Delete(ctx, DeleteVpcEndpointRequest{
		VpcEndpointId: vpcEndpointId,
	})
}

// Get a VPC endpoint configuration.
//
// Gets a VPC endpoint configuration, which represents a [VPC endpoint] object
// in AWS used to communicate privately with Databricks over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [VPC endpoint]: https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html
func (a *VpcEndpointsAPI) Get(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {
	return a.impl.Get(ctx, request)
}

// Get a VPC endpoint configuration.
//
// Gets a VPC endpoint configuration, which represents a [VPC endpoint] object
// in AWS used to communicate privately with Databricks over [AWS PrivateLink].
//
// [AWS PrivateLink]: https://aws.amazon.com/privatelink
// [VPC endpoint]: https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html
func (a *VpcEndpointsAPI) GetByVpcEndpointId(ctx context.Context, vpcEndpointId string) (*VpcEndpoint, error) {
	return a.impl.Get(ctx, GetVpcEndpointRequest{
		VpcEndpointId: vpcEndpointId,
	})
}

// Get all VPC endpoint configurations.
//
// Gets a list of all VPC endpoints for an account, specified by ID.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink].
//
// [Databricks article about PrivateLink]: https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html
func (a *VpcEndpointsAPI) List(ctx context.Context) ([]VpcEndpoint, error) {
	return a.impl.List(ctx)
}

func NewWorkspaces(client *client.DatabricksClient) *WorkspacesAPI {
	return &WorkspacesAPI{
		impl: &workspacesImpl{
			client: client,
		},
	}
}

// These APIs manage workspaces for this account. A Databricks workspace is an
// environment for accessing all of your Databricks assets. The workspace
// organizes objects (notebooks, libraries, and experiments) into folders, and
// provides access to data and computational resources such as clusters and
// jobs.
//
// These endpoints are available if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
type WorkspacesAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(WorkspacesService)
	impl WorkspacesService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *WorkspacesAPI) WithImpl(impl WorkspacesService) *WorkspacesAPI {
	a.impl = impl
	return a
}

// Impl returns low-level Workspaces API implementation
func (a *WorkspacesAPI) Impl() WorkspacesService {
	return a.impl
}

// WaitGetWorkspaceRunning repeatedly calls [WorkspacesAPI.Get] and waits to reach RUNNING state
func (a *WorkspacesAPI) WaitGetWorkspaceRunning(ctx context.Context, workspaceId int64,
	timeout time.Duration, callback func(*Workspace)) (*Workspace, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	return retries.Poll[Workspace](ctx, timeout, func() (*Workspace, *retries.Err) {
		workspace, err := a.Get(ctx, GetWorkspaceRequest{
			WorkspaceId: workspaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		if callback != nil {
			callback(workspace)
		}
		status := workspace.WorkspaceStatus
		statusMessage := workspace.WorkspaceStatusMessage
		switch status {
		case WorkspaceStatusRunning: // target state
			return workspace, nil
		case WorkspaceStatusBanned, WorkspaceStatusFailed:
			err := fmt.Errorf("failed to reach %s, got %s: %s",
				WorkspaceStatusRunning, status, statusMessage)
			return nil, retries.Halt(err)
		default:
			return nil, retries.Continues(statusMessage)
		}
	})
}

// WaitGetWorkspaceRunning is a wrapper that calls [WorkspacesAPI.WaitGetWorkspaceRunning] and waits to reach RUNNING state.
type WaitGetWorkspaceRunning[R any] struct {
	Response    *R
	WorkspaceId int64 `json:"workspace_id"`
	poll        func(time.Duration, func(*Workspace)) (*Workspace, error)
	callback    func(*Workspace)
	timeout     time.Duration
}

// OnProgress invokes a callback every time it polls for the status update.
func (w *WaitGetWorkspaceRunning[R]) OnProgress(callback func(*Workspace)) *WaitGetWorkspaceRunning[R] {
	w.callback = callback
	return w
}

// Get the Workspace with the default timeout of 20 minutes.
func (w *WaitGetWorkspaceRunning[R]) Get() (*Workspace, error) {
	return w.poll(w.timeout, w.callback)
}

// Get the Workspace with custom timeout.
func (w *WaitGetWorkspaceRunning[R]) GetWithTimeout(timeout time.Duration) (*Workspace, error) {
	return w.poll(timeout, w.callback)
}

// Create a new workspace.
//
// Creates a new workspace.
//
// **Important**: This operation is asynchronous. A response with HTTP status
// code 200 means the request has been accepted and is in progress, but does not
// mean that the workspace deployed successfully and is running. The initial
// workspace status is typically `PROVISIONING`. Use the workspace ID
// (`workspace_id`) field in the response to identify the new workspace and make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
func (a *WorkspacesAPI) Create(ctx context.Context, createWorkspaceRequest CreateWorkspaceRequest) (*WaitGetWorkspaceRunning[Workspace], error) {
	workspace, err := a.impl.Create(ctx, createWorkspaceRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetWorkspaceRunning[Workspace]{
		Response:    workspace,
		WorkspaceId: workspace.WorkspaceId,
		poll: func(timeout time.Duration, callback func(*Workspace)) (*Workspace, error) {
			return a.WaitGetWorkspaceRunning(ctx, workspace.WorkspaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [WorkspacesAPI.Create] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Workspace](60*time.Minute) functional option.
//
// Deprecated: use [WorkspacesAPI.Create].Get() or [WorkspacesAPI.WaitGetWorkspaceRunning]
func (a *WorkspacesAPI) CreateAndWait(ctx context.Context, createWorkspaceRequest CreateWorkspaceRequest, options ...retries.Option[Workspace]) (*Workspace, error) {
	wait, err := a.Create(ctx, createWorkspaceRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Workspace]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Workspace) {
		for _, o := range options {
			o(&retries.Info[Workspace]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}

// Delete a workspace.
//
// Terminates and deletes a Databricks workspace. From an API perspective,
// deletion is immediate. However, it might take a few minutes for all
// workspaces resources to be deleted, depending on the size and number of
// workspace resources.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) Delete(ctx context.Context, request DeleteWorkspaceRequest) error {
	return a.impl.Delete(ctx, request)
}

// Delete a workspace.
//
// Terminates and deletes a Databricks workspace. From an API perspective,
// deletion is immediate. However, it might take a few minutes for all
// workspaces resources to be deleted, depending on the size and number of
// workspace resources.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) DeleteByWorkspaceId(ctx context.Context, workspaceId int64) error {
	return a.impl.Delete(ctx, DeleteWorkspaceRequest{
		WorkspaceId: workspaceId,
	})
}

// Get a workspace.
//
// Gets information including status for a Databricks workspace, specified by
// ID. In the response, the `workspace_status` field indicates the current
// status. After initial workspace creation (which is asynchronous), make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account API].
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
//
// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
func (a *WorkspacesAPI) Get(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {
	return a.impl.Get(ctx, request)
}

// Get a workspace.
//
// Gets information including status for a Databricks workspace, specified by
// ID. In the response, the `workspace_status` field indicates the current
// status. After initial workspace creation (which is asynchronous), make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account API].
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
//
// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
func (a *WorkspacesAPI) GetByWorkspaceId(ctx context.Context, workspaceId int64) (*Workspace, error) {
	return a.impl.Get(ctx, GetWorkspaceRequest{
		WorkspaceId: workspaceId,
	})
}

// Get all workspaces.
//
// Gets a list of all workspaces associated with an account, specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) List(ctx context.Context) ([]Workspace, error) {
	return a.impl.List(ctx)
}

// Update workspace configuration.
//
// Updates a workspace configuration for either a running workspace or a failed
// workspace. The elements that can be updated varies between these two use
// cases.
//
// ### Update a failed workspace You can update a Databricks workspace
// configuration for failed workspace deployment for some fields, but not all
// fields. For a failed workspace, this request supports updates to the
// following fields only: - Credential configuration ID - Storage configuration
// ID - Network configuration ID. Used only to add or change a network
// configuration for a customer-managed VPC. For a failed workspace only, you
// can convert a workspace with Databricks-managed VPC to use a customer-managed
// VPC by adding this ID. You cannot downgrade a workspace with a
// customer-managed VPC to be a Databricks-managed VPC. You can update the
// network configuration for a failed or running workspace to add PrivateLink
// support, though you must also add a private access settings object. - Key
// configuration ID for managed services (control plane storage, such as
// notebook source and Databricks SQL queries). Used only if you use
// customer-managed keys for managed services. - Key configuration ID for
// workspace storage (root S3 bucket and, optionally, EBS volumes). Used only if
// you use customer-managed keys for workspace storage. **Important**: If the
// workspace was ever in the running state, even if briefly before becoming a
// failed workspace, you cannot add a new key configuration ID for workspace
// storage. - Private access settings ID to add PrivateLink support. You can add
// or update the private access settings ID to upgrade a workspace to add
// support for front-end, back-end, or both types of connectivity. You cannot
// remove (downgrade) any existing front-end or back-end PrivateLink support on
// a workspace.
//
// After calling the `PATCH` operation to update the workspace configuration,
// make repeated `GET` requests with the workspace ID and check the workspace
// status. The workspace is successful if the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account API].
//
// ### Update a running workspace You can update a Databricks workspace
// configuration for running workspaces for some fields, but not all fields. For
// a running workspace, this request supports updating the following fields
// only: - Credential configuration ID
//
// - Network configuration ID. Used only if you already use a customer-managed
// VPC. You cannot convert a running workspace from a Databricks-managed VPC to
// a customer-managed VPC. You can use a network configuration update in this
// API for a failed or running workspace to add support for PrivateLink,
// although you also need to add a private access settings object.
//
// - Key configuration ID for managed services (control plane storage, such as
// notebook source and Databricks SQL queries). Databricks does not directly
// encrypt the data with the customer-managed key (CMK). Databricks uses both
// the CMK and the Databricks managed key (DMK) that is unique to your workspace
// to encrypt the Data Encryption Key (DEK). Databricks uses the DEK to encrypt
// your workspace's managed services persisted data. If the workspace does not
// already have a CMK for managed services, adding this ID enables managed
// services encryption for new or updated data. Existing managed services data
// that existed before adding the key remains not encrypted with the DEK until
// it is modified. If the workspace already has customer-managed keys for
// managed services, this request rotates (changes) the CMK keys and the DEK is
// re-encrypted with the DMK and the new CMK. - Key configuration ID for
// workspace storage (root S3 bucket and, optionally, EBS volumes). You can set
// this only if the workspace does not already have a customer-managed key
// configuration for workspace storage. - Private access settings ID to add
// PrivateLink support. You can add or update the private access settings ID to
// upgrade a workspace to add support for front-end, back-end, or both types of
// connectivity. You cannot remove (downgrade) any existing front-end or
// back-end PrivateLink support on a workspace.
//
// **Important**: To update a running workspace, your workspace must have no
// running compute resources that run in your workspace's VPC in the Classic
// data plane. For example, stop all all-purpose clusters, job clusters, pools
// with running clusters, and Classic SQL warehouses. If you do not terminate
// all cluster instances in the workspace before calling this API, the request
// will fail.
//
// ### Wait until changes take effect. After calling the `PATCH` operation to
// update the workspace configuration, make repeated `GET` requests with the
// workspace ID and check the workspace status and the status of the fields. *
// For workspaces with a Databricks-managed VPC, the workspace status becomes
// `PROVISIONING` temporarily (typically under 20 minutes). If the workspace
// update is successful, the workspace status changes to `RUNNING`. Note that
// you can also check the workspace status in the [Account Console]. However,
// you cannot use or create clusters for another 20 minutes after that status
// change. This results in a total of up to 40 minutes in which you cannot
// create clusters. If you create or use clusters before this time interval
// elapses, clusters do not launch successfully, fail, or could cause other
// unexpected behavior.
//
// * For workspaces with a customer-managed VPC, the workspace status stays at
// status `RUNNING` and the VPC change happens immediately. A change to the
// storage customer-managed key configuration ID might take a few minutes to
// update, so continue to check the workspace until you observe that it has been
// updated. If the update fails, the workspace might revert silently to its
// original configuration. After the workspace has been updated, you cannot use
// or create clusters for another 20 minutes. If you create or use clusters
// before this time interval elapses, clusters do not launch successfully, fail,
// or could cause other unexpected behavior.
//
// If you update the _storage_ customer-managed key configurations, it takes 20
// minutes for the changes to fully take effect. During the 20 minute wait, it
// is important that you stop all REST API calls to the DBFS API. If you are
// modifying _only the managed services key configuration_, you can omit the 20
// minute wait.
//
// **Important**: Customer-managed keys and customer-managed VPCs are supported
// by only some deployment types and subscription types. If you have questions
// about availability, contact your Databricks representative.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
//
// [Account Console]: https://docs.databricks.com/administration-guide/account-settings-e2/account-console-e2.html
// [Create a new workspace using the Account API]: http://docs.databricks.com/administration-guide/account-api/new-workspace.html
func (a *WorkspacesAPI) Update(ctx context.Context, updateWorkspaceRequest UpdateWorkspaceRequest) (*WaitGetWorkspaceRunning[any], error) {
	err := a.impl.Update(ctx, updateWorkspaceRequest)
	if err != nil {
		return nil, err
	}
	return &WaitGetWorkspaceRunning[any]{

		WorkspaceId: updateWorkspaceRequest.WorkspaceId,
		poll: func(timeout time.Duration, callback func(*Workspace)) (*Workspace, error) {
			return a.WaitGetWorkspaceRunning(ctx, updateWorkspaceRequest.WorkspaceId, timeout, callback)
		},
		timeout:  20 * time.Minute,
		callback: nil,
	}, nil
}

// Calls [WorkspacesAPI.Update] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Workspace](60*time.Minute) functional option.
//
// Deprecated: use [WorkspacesAPI.Update].Get() or [WorkspacesAPI.WaitGetWorkspaceRunning]
func (a *WorkspacesAPI) UpdateAndWait(ctx context.Context, updateWorkspaceRequest UpdateWorkspaceRequest, options ...retries.Option[Workspace]) (*Workspace, error) {
	wait, err := a.Update(ctx, updateWorkspaceRequest)
	if err != nil {
		return nil, err
	}
	tmp := &retries.Info[Workspace]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(tmp)
	}
	wait.timeout = tmp.Timeout
	wait.callback = func(info *Workspace) {
		for _, o := range options {
			o(&retries.Info[Workspace]{
				Info:    info,
				Timeout: wait.timeout,
			})
		}
	}
	return wait.Get()
}
