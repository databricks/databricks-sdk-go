// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// These APIs manage credential configurations for this workspace. Databricks
// needs access to a cross-account service IAM role in your AWS account so that
// Databricks can deploy clusters in the appropriate VPC for the new workspace.
// A credential configuration encapsulates this role information, and its ID is
// used when creating a new workspace.
type CredentialsClient struct {
	CredentialsAPI
}

func NewCredentialsClient(cfg *config.Config) (*CredentialsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CredentialsClient{
		CredentialsAPI: CredentialsAPI{
			credentialsImpl: credentialsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type EncryptionKeysClient struct {
	EncryptionKeysAPI
}

func NewEncryptionKeysClient(cfg *config.Config) (*EncryptionKeysClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &EncryptionKeysClient{
		EncryptionKeysAPI: EncryptionKeysAPI{
			encryptionKeysImpl: encryptionKeysImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage network configurations for customer-managed VPCs
// (optional). Its ID is used when creating a new workspace if you use
// customer-managed VPCs.
type NetworksClient struct {
	NetworksAPI
}

func NewNetworksClient(cfg *config.Config) (*NetworksClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &NetworksClient{
		NetworksAPI: NetworksAPI{
			networksImpl: networksImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage private access settings for this account.
type PrivateAccessClient struct {
	PrivateAccessAPI
}

func NewPrivateAccessClient(cfg *config.Config) (*PrivateAccessClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PrivateAccessClient{
		PrivateAccessAPI: PrivateAccessAPI{
			privateAccessImpl: privateAccessImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage storage configurations for this workspace. A root storage
// S3 bucket in your account is required to store objects like cluster logs,
// notebook revisions, and job results. You can also use the root storage S3
// bucket for storage of non-production DBFS data. A storage configuration
// encapsulates this bucket information, and its ID is used when creating a new
// workspace.
type StorageClient struct {
	StorageAPI
}

func NewStorageClient(cfg *config.Config) (*StorageClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &StorageClient{
		StorageAPI: StorageAPI{
			storageImpl: storageImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// These APIs manage VPC endpoint configurations for this account.
type VpcEndpointsClient struct {
	VpcEndpointsAPI
}

func NewVpcEndpointsClient(cfg *config.Config) (*VpcEndpointsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &VpcEndpointsClient{
		VpcEndpointsAPI: VpcEndpointsAPI{
			vpcEndpointsImpl: vpcEndpointsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
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
type WorkspacesClient struct {
	WorkspacesAPI
}

func NewWorkspacesClient(cfg *config.Config) (*WorkspacesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.AccountID == "" || !cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid account config for the requested account service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &WorkspacesClient{
		WorkspacesAPI: WorkspacesAPI{
			workspacesImpl: workspacesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
