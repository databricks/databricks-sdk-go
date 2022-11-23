// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package deployment

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/useragent"
)

func NewCredentialConfigurations(client databricksClient) *CredentialConfigurationsAPI {
	return &CredentialConfigurationsAPI{
		impl: &credentialConfigurationsImpl{
			client: client,
		},
	}
}

// These APIs manage credential configurations for this workspace. Databricks
// needs access to a cross-account service IAM role in your AWS account so that
// Databricks can deploy clusters in the appropriate VPC for the new workspace.
// A credential configuration encapsulates this role information, and its ID is
// used when creating a new workspace.
type CredentialConfigurationsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(CredentialConfigurationsService)
	impl CredentialConfigurationsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *CredentialConfigurationsAPI) WithImpl(impl CredentialConfigurationsService) *CredentialConfigurationsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level CredentialConfigurations API implementation
func (a *CredentialConfigurationsAPI) Impl() CredentialConfigurationsService {
	return a.impl
}

// Create credential configuration
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
// [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
func (a *CredentialConfigurationsAPI) CreateCredentialConfig(ctx context.Context, request CreateCredentialRequest) (*Credential, error) {
	return a.impl.CreateCredentialConfig(ctx, request)
}

// Delete credential configuration
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialConfigurationsAPI) DeleteCredentialConfig(ctx context.Context, request DeleteCredentialConfigRequest) error {
	return a.impl.DeleteCredentialConfig(ctx, request)
}

// Delete credential configuration
//
// Deletes a Databricks credential configuration object for an account, both
// specified by ID. You cannot delete a credential that is associated with any
// workspace.
func (a *CredentialConfigurationsAPI) DeleteCredentialConfigByCredentialsId(ctx context.Context, credentialsId string) error {
	return a.impl.DeleteCredentialConfig(ctx, DeleteCredentialConfigRequest{
		CredentialsId: credentialsId,
	})
}

// Get credential configuration
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialConfigurationsAPI) GetCredentialConfig(ctx context.Context, request GetCredentialConfigRequest) (*Credential, error) {
	return a.impl.GetCredentialConfig(ctx, request)
}

// Get credential configuration
//
// Gets a Databricks credential configuration object for an account, both
// specified by ID.
func (a *CredentialConfigurationsAPI) GetCredentialConfigByCredentialsId(ctx context.Context, credentialsId string) (*Credential, error) {
	return a.impl.GetCredentialConfig(ctx, GetCredentialConfigRequest{
		CredentialsId: credentialsId,
	})
}

// Get all credential configurations
//
// Gets all Databricks credential configurations associated with an account
// specified by ID.
func (a *CredentialConfigurationsAPI) ListCredentials(ctx context.Context) ([]Credential, error) {
	return a.impl.ListCredentials(ctx)
}

func NewKeyConfigurations(client databricksClient) *KeyConfigurationsAPI {
	return &KeyConfigurationsAPI{
		impl: &keyConfigurationsImpl{
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
// platform. If you are not sure, contact your Databricks reprsentative.
type KeyConfigurationsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(KeyConfigurationsService)
	impl KeyConfigurationsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *KeyConfigurationsAPI) WithImpl(impl KeyConfigurationsService) *KeyConfigurationsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level KeyConfigurations API implementation
func (a *KeyConfigurationsAPI) Impl() KeyConfigurationsService {
	return a.impl
}

// Create encryption key configuration
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
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *KeyConfigurationsAPI) CreateKeyConfig(ctx context.Context, request CreateCustomerManagedKeyRequest) (*CustomerManagedKey, error) {
	return a.impl.CreateKeyConfig(ctx, request)
}

// Delete encryption key configuration
//
// Deletes a customer-managed key configuration object for an account. You
// cannot delete a configuration that is associated with a running workspace.
func (a *KeyConfigurationsAPI) DeleteKeyConfig(ctx context.Context, request DeleteKeyConfigRequest) error {
	return a.impl.DeleteKeyConfig(ctx, request)
}

// Delete encryption key configuration
//
// Deletes a customer-managed key configuration object for an account. You
// cannot delete a configuration that is associated with a running workspace.
func (a *KeyConfigurationsAPI) DeleteKeyConfigByCustomerManagedKeyId(ctx context.Context, customerManagedKeyId string) error {
	return a.impl.DeleteKeyConfig(ctx, DeleteKeyConfigRequest{
		CustomerManagedKeyId: customerManagedKeyId,
	})
}

// Get encryption key configuration
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
// platform.
func (a *KeyConfigurationsAPI) GetKeyConfig(ctx context.Context, request GetKeyConfigRequest) (*CustomerManagedKey, error) {
	return a.impl.GetKeyConfig(ctx, request)
}

// Get encryption key configuration
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
// platform.
func (a *KeyConfigurationsAPI) GetKeyConfigByCustomerManagedKeyId(ctx context.Context, customerManagedKeyId string) (*CustomerManagedKey, error) {
	return a.impl.GetKeyConfig(ctx, GetKeyConfigRequest{
		CustomerManagedKeyId: customerManagedKeyId,
	})
}

// Get history of a key's associations with workspaces
//
// Gets a list of records that show how key configurations are associated with
// workspaces.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types, subscription types, and AWS regions.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *KeyConfigurationsAPI) GetKeyWorkspaceHistory(ctx context.Context) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	return a.impl.GetKeyWorkspaceHistory(ctx)
}

// Get all encryption key configurations
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
func (a *KeyConfigurationsAPI) ListKeyConfigs(ctx context.Context) ([]CustomerManagedKey, error) {
	return a.impl.ListKeyConfigs(ctx)
}

func NewNetworkConfigurations(client databricksClient) *NetworkConfigurationsAPI {
	return &NetworkConfigurationsAPI{
		impl: &networkConfigurationsImpl{
			client: client,
		},
	}
}

// These APIs manage network configurations for customer-managed VPCs
// (optional). A network configuration encapsulates the IDs for AWS VPCs,
// subnets, and security groups. Its ID is used when creating a new workspace if
// you use customer-managed VPCs.
type NetworkConfigurationsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(NetworkConfigurationsService)
	impl NetworkConfigurationsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *NetworkConfigurationsAPI) WithImpl(impl NetworkConfigurationsService) *NetworkConfigurationsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level NetworkConfigurations API implementation
func (a *NetworkConfigurationsAPI) Impl() NetworkConfigurationsService {
	return a.impl
}

// Create network configuration
//
// Creates a Databricks network configuration that represents an AWS VPC and its
// resources. The VPC will be used for new Databricks clusters. This requires a
// pre-existing VPC and subnets. For VPC requirements, see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
//
// **Important**: You can share one customer-managed VPC with multiple
// workspaces in a single account. Therefore, you can share one VPC across
// multiple Account API network configurations. However, you **cannot** reuse
// subnets or Security Groups between workspaces. Because a Databricks Account
// API network configuration encapsulates this information, you cannot reuse a
// Databricks Account API network configuration across workspaces. If you plan
// to share one VPC with multiple workspaces, make sure you size your VPC and
// subnets accordingly. For information about how to create a new workspace with
// this API, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) CreateNetworkConfig(ctx context.Context, request CreateNetworkRequest) (*Network, error) {
	return a.impl.CreateNetworkConfig(ctx, request)
}

// Delete network configuration
//
// Deletes a Databricks network configuration, which represents an AWS VPC and
// its resources. You cannot delete a network that is associated with a
// workspace.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) DeleteNetworkConfig(ctx context.Context, request DeleteNetworkConfigRequest) error {
	return a.impl.DeleteNetworkConfig(ctx, request)
}

// Delete network configuration
//
// Deletes a Databricks network configuration, which represents an AWS VPC and
// its resources. You cannot delete a network that is associated with a
// workspace.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) DeleteNetworkConfigByNetworkId(ctx context.Context, networkId string) error {
	return a.impl.DeleteNetworkConfig(ctx, DeleteNetworkConfigRequest{
		NetworkId: networkId,
	})
}

// Get a network configuration
//
// Gets a Databricks network configuration, which represents an AWS VPC and its
// resources. This requires a pre-existing VPC and subnets. For VPC
// requirements, see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetNetworkConfig(ctx context.Context, request GetNetworkConfigRequest) (*Network, error) {
	return a.impl.GetNetworkConfig(ctx, request)
}

// Get a network configuration
//
// Gets a Databricks network configuration, which represents an AWS VPC and its
// resources. This requires a pre-existing VPC and subnets. For VPC
// requirements, see [Customer-managed
// VPC](http://docs.databricks.com/administration-guide/cloud-configurations/aws/customer-managed-vpc.html).
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) GetNetworkConfigByNetworkId(ctx context.Context, networkId string) (*Network, error) {
	return a.impl.GetNetworkConfig(ctx, GetNetworkConfigRequest{
		NetworkId: networkId,
	})
}

// Get all network configurations
//
// Gets a list of all Databricks network configurations for an account,
// specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *NetworkConfigurationsAPI) ListNetworkConfigs(ctx context.Context) ([]Network, error) {
	return a.impl.ListNetworkConfigs(ctx)
}

func NewPrivateAccessSettings(client databricksClient) *PrivateAccessSettingsAPI {
	return &PrivateAccessSettingsAPI{
		impl: &privateAccessSettingsImpl{
			client: client,
		},
	}
}

// These APIs manage private access settings for this account. A private access
// settings object specifies how your workspace is accessed using AWS
// PrivateLink. Each workspace that has any PrivateLink connections must include
// the ID for a private access settings object is in its workspace configuration
// object. Your account must be enabled for PrivateLink to use these APIs.
// Before configuring PrivateLink, it is important to read the [Databricks
// article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
type PrivateAccessSettingsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(PrivateAccessSettingsService)
	impl PrivateAccessSettingsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *PrivateAccessSettingsAPI) WithImpl(impl PrivateAccessSettingsService) *PrivateAccessSettingsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level PrivateAccessSettings API implementation
func (a *PrivateAccessSettingsAPI) Impl() PrivateAccessSettingsService {
	return a.impl
}

// Create private access settings
//
// Creates a private access settings object, which specifies how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink). To
// use AWS PrivateLink, a workspace must have a private access settings object
// referenced by ID in the workspace's `private_access_settings_id` property.
//
// You can share one private access settings with multiple workspaces in a
// single account. However, private access settings are specific to AWS regions,
// so only workspaces in the same AWS region can use a given private access
// settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) CreatePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	return a.impl.CreatePrivateAccessSettings(ctx, request)
}

// Delete a private access settings object
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) DeletePrivateAccessSettings(ctx context.Context, request DeletePrivateAccessSettingsRequest) error {
	return a.impl.DeletePrivateAccessSettings(ctx, request)
}

// Delete a private access settings object
//
// Deletes a private access settings object, which determines how your workspace
// is accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) DeletePrivateAccessSettingsByPrivateAccessSettingsId(ctx context.Context, privateAccessSettingsId string) error {
	return a.impl.DeletePrivateAccessSettings(ctx, DeletePrivateAccessSettingsRequest{
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Get a private access settings object
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettings(ctx context.Context, request GetPrivateAccessSettingsRequest) (*PrivateAccessSettings, error) {
	return a.impl.GetPrivateAccessSettings(ctx, request)
}

// Get a private access settings object
//
// Gets a private access settings object, which specifies how your workspace is
// accessed over [AWS PrivateLink](https://aws.amazon.com/privatelink).
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) GetPrivateAccessSettingsByPrivateAccessSettingsId(ctx context.Context, privateAccessSettingsId string) (*PrivateAccessSettings, error) {
	return a.impl.GetPrivateAccessSettings(ctx, GetPrivateAccessSettingsRequest{
		PrivateAccessSettingsId: privateAccessSettingsId,
	})
}

// Get all private access settings objects
//
// Gets a list of all private access settings objects for an account, specified
// by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for AWS PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) ListPrivateAccessSettings(ctx context.Context) ([]PrivateAccessSettings, error) {
	return a.impl.ListPrivateAccessSettings(ctx)
}

// Replace private access settings
//
// Updates an existing private access settings object, which specifies how your
// workspace is accessed over [AWS
// PrivateLink](https://aws.amazon.com/privatelink). To use AWS PrivateLink, a
// workspace must have a private access settings object referenced by ID in the
// workspace's `private_access_settings_id` property.
//
// This operation completely overwrites your existing private access settings
// object attached to your workspaces. All workspaces attached to the private
// access settings are affected by any change. If `public_access_enabled`,
// `private_access_level`, or `allowed_vpc_endpoint_ids` are updated, effects of
// these changes might take several minutes to propagate to the workspace API.
// You can share one private access settings object with multiple workspaces in
// a single account. However, private access settings are specific to AWS
// regions, so only workspaces in the same AWS region can use a given private
// access settings object.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *PrivateAccessSettingsAPI) ReplacePrivateAccessSettings(ctx context.Context, request UpsertPrivateAccessSettingsRequest) error {
	return a.impl.ReplacePrivateAccessSettings(ctx, request)
}

func NewStorageConfigurations(client databricksClient) *StorageConfigurationsAPI {
	return &StorageConfigurationsAPI{
		impl: &storageConfigurationsImpl{
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
type StorageConfigurationsAPI struct {
	// impl contains low-level REST API interface, that could be overridden
	// through WithImpl(StorageConfigurationsService)
	impl StorageConfigurationsService
}

// WithImpl could be used to override low-level API implementations for unit
// testing purposes with [github.com/golang/mock] or other mocking frameworks.
func (a *StorageConfigurationsAPI) WithImpl(impl StorageConfigurationsService) *StorageConfigurationsAPI {
	a.impl = impl
	return a
}

// Impl returns low-level StorageConfigurations API implementation
func (a *StorageConfigurationsAPI) Impl() StorageConfigurationsService {
	return a.impl
}

// Create new storage configuration
//
// Creates new storage configuration for an account, specified by ID. Uploads a
// storage configuration object that represents the root AWS S3 bucket in your
// account. Databricks stores related workspace assets including DBFS, cluster
// logs, and job results. For the AWS S3 bucket, you need to configure the
// required bucket policy.
//
// For information about how to create a new workspace with this API, see
// [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html)
func (a *StorageConfigurationsAPI) CreateStorageConfig(ctx context.Context, request CreateStorageConfigurationRequest) (*StorageConfiguration, error) {
	return a.impl.CreateStorageConfig(ctx, request)
}

// Delete storage configuration
//
// Deletes a Databricks storage configuration. You cannot delete a storage
// configuration that is associated with any workspace.
func (a *StorageConfigurationsAPI) DeleteStorageConfig(ctx context.Context, request DeleteStorageConfigRequest) error {
	return a.impl.DeleteStorageConfig(ctx, request)
}

// Delete storage configuration
//
// Deletes a Databricks storage configuration. You cannot delete a storage
// configuration that is associated with any workspace.
func (a *StorageConfigurationsAPI) DeleteStorageConfigByStorageConfigurationId(ctx context.Context, storageConfigurationId string) error {
	return a.impl.DeleteStorageConfig(ctx, DeleteStorageConfigRequest{
		StorageConfigurationId: storageConfigurationId,
	})
}

// Get storage configuration
//
// Gets a Databricks storage configuration for an account, both specified by ID.
func (a *StorageConfigurationsAPI) GetStorageConfig(ctx context.Context, request GetStorageConfigRequest) (*StorageConfiguration, error) {
	return a.impl.GetStorageConfig(ctx, request)
}

// Get storage configuration
//
// Gets a Databricks storage configuration for an account, both specified by ID.
func (a *StorageConfigurationsAPI) GetStorageConfigByStorageConfigurationId(ctx context.Context, storageConfigurationId string) (*StorageConfiguration, error) {
	return a.impl.GetStorageConfig(ctx, GetStorageConfigRequest{
		StorageConfigurationId: storageConfigurationId,
	})
}

// Get all storage configurations
//
// Gets a list of all Databricks storage configurations for your account,
// specified by ID.
func (a *StorageConfigurationsAPI) ListStorageConfigs(ctx context.Context) ([]StorageConfiguration, error) {
	return a.impl.ListStorageConfigs(ctx)
}

func NewVpcEndpoints(client databricksClient) *VpcEndpointsAPI {
	return &VpcEndpointsAPI{
		impl: &vpcEndpointsImpl{
			client: client,
		},
	}
}

// These APIs manage VPC endpoint configurations for this account. This object
// registers an AWS VPC endpoint in your Databricks account so your workspace
// can use it with AWS PrivateLink. Your VPC endpoint connects to one of two VPC
// endpoint services -- one for workspace (both for front-end connection and for
// back-end connection to REST APIs) and one for the back-end secure cluster
// connectivity relay from the data plane. Your account must be enabled for
// PrivateLink to use these APIs. Before configuring PrivateLink, it is
// important to read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
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

// Create VPC endpoint configuration
//
// Creates a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/vpc-endpoints.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// **Important**: When you register a VPC endpoint to the Databricks workspace
// VPC endpoint service for any workspace, **in this release Databricks enables
// front-end (web application and REST API) access from the source network of
// the VPC endpoint to all workspaces in that AWS region in your Databricks
// account if the workspaces have any PrivateLink connections in their workspace
// configuration**. If you have questions about this behavior, contact your
// Databricks representative.
//
// Within AWS, your VPC endpoint stays in `pendingAcceptance` state until you
// register it in a VPC endpoint configuration through the Account API. After
// you register the VPC endpoint configuration, the Databricks [endpoint
// service](https://docs.aws.amazon.com/vpc/latest/privatelink/privatelink-share-your-services.html)
// automatically accepts the VPC endpoint and it eventually transitions to the
// `available` state.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) CreateVpcEndpoint(ctx context.Context, request CreateVpcEndpointRequest) (*VpcEndpoint, error) {
	return a.impl.CreateVpcEndpoint(ctx, request)
}

// Delete VPC endpoint configuration
//
// Deletes a VPC endpoint configuration, which represents an [AWS VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// that can communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// Upon deleting a VPC endpoint configuration, the VPC endpoint in AWS changes
// its state from `accepted` to `rejected`, which means that it is no longer
// usable from your VPC.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) DeleteVpcEndpoint(ctx context.Context, request DeleteVpcEndpointRequest) error {
	return a.impl.DeleteVpcEndpoint(ctx, request)
}

// Delete VPC endpoint configuration
//
// Deletes a VPC endpoint configuration, which represents an [AWS VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// that can communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// Upon deleting a VPC endpoint configuration, the VPC endpoint in AWS changes
// its state from `accepted` to `rejected`, which means that it is no longer
// usable from your VPC.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) DeleteVpcEndpointByVpcEndpointId(ctx context.Context, vpcEndpointId string) error {
	return a.impl.DeleteVpcEndpoint(ctx, DeleteVpcEndpointRequest{
		VpcEndpointId: vpcEndpointId,
	})
}

// Get a VPC endpoint configuration
//
// Gets a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) GetVpcEndpoint(ctx context.Context, request GetVpcEndpointRequest) (*VpcEndpoint, error) {
	return a.impl.GetVpcEndpoint(ctx, request)
}

// Get a VPC endpoint configuration
//
// Gets a VPC endpoint configuration, which represents a [VPC
// endpoint](https://docs.aws.amazon.com/vpc/latest/privatelink/concepts.html)
// object in AWS used to communicate privately with Databricks over [AWS
// PrivateLink](https://aws.amazon.com/privatelink).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) GetVpcEndpointByVpcEndpointId(ctx context.Context, vpcEndpointId string) (*VpcEndpoint, error) {
	return a.impl.GetVpcEndpoint(ctx, GetVpcEndpointRequest{
		VpcEndpointId: vpcEndpointId,
	})
}

// Get all VPC endpoint configurations
//
// Gets a list of all VPC endpoints for an account, specified by ID.
//
// Before configuring PrivateLink, read the [Databricks article about
// PrivateLink](https://docs.databricks.com/administration-guide/cloud-configurations/aws/privatelink.html).
//
// This operation is available only if your account is on the E2 version of the
// platform and your Databricks account is enabled for PrivateLink (Public
// Preview). Contact your Databricks representative to enable your account for
// PrivateLink.
func (a *VpcEndpointsAPI) ListVpcEndpoints(ctx context.Context) ([]VpcEndpoint, error) {
	return a.impl.ListVpcEndpoints(ctx)
}

func NewWorkspaces(client databricksClient) *WorkspacesAPI {
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

// Create a new workspace
//
// Creates a new workspace using a credential configuration and a storage
// configuration, an optional network configuration (if using a customer-managed
// VPC), an optional managed services key configuration (if using
// customer-managed keys for managed services), and an optional storage key
// configuration (if using customer-managed keys for storage). The key
// configurations used for managed services and storage encryption can be the
// same or different.
//
// **Important**: This operation is asynchronous. A response with HTTP status
// code 200 means the request has been accepted and is in progress, but does not
// mean that the workspace deployed successfully and is running. The initial
// workspace status is typically `PROVISIONING`. Use the workspace ID
// (`workspace_id`) field in the response to identify the new workspace and make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// You can share one customer-managed VPC with multiple workspaces in a single
// account. It is not required to create a new VPC for each workspace. However,
// you **cannot** reuse subnets or Security Groups between workspaces. If you
// plan to share one VPC with multiple workspaces, make sure you size your VPC
// and subnets accordingly. Because a Databricks Account API network
// configuration encapsulates this information, you cannot reuse a Databricks
// Account API network configuration across workspaces.\nFor information about
// how to create a new workspace with this API **including error handling**, see
// [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// **Important**: Customer-managed VPCs, PrivateLink, and customer-managed keys
// are supported on a limited set of deployment and subscription types. If you
// have questions about availability, contact your Databricks
// representative.\n\nThis operation is available only if your account is on the
// E2 version of the platform or on a select custom plan that allows multiple
// workspaces per account.
func (a *WorkspacesAPI) CreateWorkspace(ctx context.Context, request CreateWorkspaceRequest) (*Workspace, error) {
	return a.impl.CreateWorkspace(ctx, request)
}

// Calls [WorkspacesAPI.CreateWorkspace] and waits to reach RUNNING state
//
// You can override the default timeout of 20 minutes by calling adding
// retries.Timeout[Workspace](60*time.Minute) functional option.
func (a *WorkspacesAPI) CreateWorkspaceAndWait(ctx context.Context, createWorkspaceRequest CreateWorkspaceRequest, options ...retries.Option[Workspace]) (*Workspace, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "long-running")
	workspace, err := a.CreateWorkspace(ctx, createWorkspaceRequest)
	if err != nil {
		return nil, err
	}
	i := retries.Info[Workspace]{Timeout: 20 * time.Minute}
	for _, o := range options {
		o(&i)
	}
	return retries.Poll[Workspace](ctx, i.Timeout, func() (*Workspace, *retries.Err) {
		workspace, err := a.GetWorkspace(ctx, GetWorkspaceRequest{
			WorkspaceId: workspace.WorkspaceId,
		})
		if err != nil {
			return nil, retries.Halt(err)
		}
		for _, o := range options {
			o(&retries.Info[Workspace]{
				Info:    *workspace,
				Timeout: i.Timeout,
			})
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

// Delete workspace
//
// Terminates and deletes a Databricks workspace. From an API perspective,
// deletion is immediate. However, it might take a few minutes for all
// workspaces resources to be deleted, depending on the size and number of
// workspace resources.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) DeleteWorkspace(ctx context.Context, request DeleteWorkspaceRequest) error {
	return a.impl.DeleteWorkspace(ctx, request)
}

// Delete workspace
//
// Terminates and deletes a Databricks workspace. From an API perspective,
// deletion is immediate. However, it might take a few minutes for all
// workspaces resources to be deleted, depending on the size and number of
// workspace resources.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) DeleteWorkspaceByWorkspaceId(ctx context.Context, workspaceId int64) error {
	return a.impl.DeleteWorkspace(ctx, DeleteWorkspaceRequest{
		WorkspaceId: workspaceId,
	})
}

// Get all workspaces
//
// Gets a list of all workspaces associated with an account, specified by ID.
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) GetAllWorkspaces(ctx context.Context) ([]Workspace, error) {
	return a.impl.GetAllWorkspaces(ctx)
}

// Get workspace
//
// Gets information including status for a Databricks workspace, specified by
// ID. In the response, the `workspace_status` field indicates the current
// status. After initial workspace creation (which is asynchronous), make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) GetWorkspace(ctx context.Context, request GetWorkspaceRequest) (*Workspace, error) {
	return a.impl.GetWorkspace(ctx, request)
}

// Get workspace
//
// Gets information including status for a Databricks workspace, specified by
// ID. In the response, the `workspace_status` field indicates the current
// status. After initial workspace creation (which is asynchronous), make
// repeated `GET` requests with the workspace ID and check its status. The
// workspace becomes available when the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// This operation is available only if your account is on the E2 version of the
// platform or on a select custom plan that allows multiple workspaces per
// account.
func (a *WorkspacesAPI) GetWorkspaceByWorkspaceId(ctx context.Context, workspaceId int64) (*Workspace, error) {
	return a.impl.GetWorkspace(ctx, GetWorkspaceRequest{
		WorkspaceId: workspaceId,
	})
}

// Get the history of a workspace's associations with keys
//
// Gets a list of all associations with key configuration objects for the
// specified workspace that encapsulate customer-managed keys that encrypt
// managed services, workspace storage, or in some cases both.
//
// **Important**: In the current implementation, keys cannot be rotated or
// removed from a workspace. It is possible for a workspace to show a storage
// customer-managed key having been attached and then detached if the workspace
// was updated to use the key and the update operation failed.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types and subscription types.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *WorkspacesAPI) GetWorkspaceKeyHistory(ctx context.Context, request GetWorkspaceKeyHistoryRequest) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	return a.impl.GetWorkspaceKeyHistory(ctx, request)
}

// Get the history of a workspace's associations with keys
//
// Gets a list of all associations with key configuration objects for the
// specified workspace that encapsulate customer-managed keys that encrypt
// managed services, workspace storage, or in some cases both.
//
// **Important**: In the current implementation, keys cannot be rotated or
// removed from a workspace. It is possible for a workspace to show a storage
// customer-managed key having been attached and then detached if the workspace
// was updated to use the key and the update operation failed.
//
// **Important**: Customer-managed keys are supported only for some deployment
// types and subscription types.
//
// This operation is available only if your account is on the E2 version of the
// platform.
func (a *WorkspacesAPI) GetWorkspaceKeyHistoryByWorkspaceId(ctx context.Context, workspaceId int64) (*ListWorkspaceEncryptionKeyRecordsResponse, error) {
	return a.impl.GetWorkspaceKeyHistory(ctx, GetWorkspaceKeyHistoryRequest{
		WorkspaceId: workspaceId,
	})
}

// Update workspace configuration
//
// Updates a workspace configuration for either a running workspace or a failed
// workspace. The elements that can be updated varies between these two use
// cases.
//
// ### Update a failed workspace You can update a Databricks workspace
// configuration for failed workspace deployment for some fields, but not all
// fields. For a failed workspace, this request supports updates to the
// following fields only: - Credential configuration ID - Storage configuration
// ID - Network configuration ID. Used only if you use customer-managed VPC. -
// Key configuration ID for managed services (control plane storage, such as
// notebook source and Databricks SQL queries). Used only if you use
// customer-managed keys for managed services. - Key configuration ID for
// workspace storage (root S3 bucket and, optionally, EBS volumes). Used only if
// you use customer-managed keys for workspace storage. **Important**: If the
// workspace was ever in the running state, even if briefly before becoming a
// failed workspace, you cannot add a new key configuration ID for workspace
// storage.
//
// After calling the `PATCH` operation to update the workspace configuration,
// make repeated `GET` requests with the workspace ID and check the workspace
// status. The workspace is successful if the status changes to `RUNNING`.
//
// For information about how to create a new workspace with this API **including
// error handling**, see [Create a new workspace using the Account
// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
//
// ### Update a running workspace You can update a Databricks workspace
// configuration for running workspaces for some fields, but not all fields. For
// a running workspace, this request supports updating the following fields
// only: - Credential configuration ID
//
// - Network configuration ID. Used only if you already use use customer-managed
// VPC. This change is supported only if you specified a network configuration
// ID in your original workspace creation. In other words, you cannot switch
// from a Databricks-managed VPC to a customer-managed VPC. **Note**: You cannot
// use a network configuration update in this API to add support for PrivateLink
// (in Public Preview). To add PrivateLink to an existing workspace, contact
// your Databricks representative.
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
// configuration for workspace storage.
//
// **Important**: For updating running workspaces, this API is unavailable on
// Mondays, Tuesdays, and Thursdays from 4:30pm-7:30pm PST due to routine
// maintenance. Plan your workspace updates accordingly. For questions about
// this schedule, contact your Databricks representative.
//
// **Important**: To update a running workspace, your workspace must have no
// running cluster instances, which includes all-purpose clusters, job clusters,
// and pools that might have running clusters. Terminate all cluster instances
// in the workspace before calling this API.
//
// ### Wait until changes take effect. After calling the `PATCH` operation to
// update the workspace configuration, make repeated `GET` requests with the
// workspace ID and check the workspace status and the status of the fields. *
// For workspaces with a Databricks-managed VPC, the workspace status becomes
// `PROVISIONING` temporarily (typically under 20 minutes). If the workspace
// update is successful, the workspace status changes to `RUNNING`. Note that
// you can also check the workspace status in the [Account
// Console](https://docs.databricks.com/administration-guide/account-settings-e2/account-console-e2.html).
// However, you cannot use or create clusters for another 20 minutes after that
// status change. This results in a total of up to 40 minutes in which you
// cannot create clusters. If you create or use clusters before this time
// interval elapses, clusters do not launch successfully, fail, or could cause
// other unexpected behavior.
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
func (a *WorkspacesAPI) UpdateWorkspace(ctx context.Context, request UpdateWorkspaceRequest) error {
	return a.impl.UpdateWorkspace(ctx, request)
}
