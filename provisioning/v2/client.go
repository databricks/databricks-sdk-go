// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package provisioning

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

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
