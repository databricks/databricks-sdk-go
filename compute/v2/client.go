// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ClusterPoliciesClient struct {
	ClusterPoliciesAPI
}

func NewClusterPoliciesClient(cfg *config.Config) (*ClusterPoliciesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ClusterPoliciesClient{
		ClusterPoliciesAPI: ClusterPoliciesAPI{
			clusterPoliciesImpl: clusterPoliciesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type ClustersClient struct {
	ClustersAPI
}

func NewClustersClient(cfg *config.Config) (*ClustersClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ClustersClient{
		ClustersAPI: ClustersAPI{
			clustersImpl: clustersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type CommandExecutionClient struct {
	CommandExecutionAPI
}

func NewCommandExecutionClient(cfg *config.Config) (*CommandExecutionClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CommandExecutionClient{
		CommandExecutionAPI: CommandExecutionAPI{
			commandExecutionImpl: commandExecutionImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type GlobalInitScriptsClient struct {
	GlobalInitScriptsAPI
}

func NewGlobalInitScriptsClient(cfg *config.Config) (*GlobalInitScriptsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GlobalInitScriptsClient{
		GlobalInitScriptsAPI: GlobalInitScriptsAPI{
			globalInitScriptsImpl: globalInitScriptsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type InstancePoolsClient struct {
	InstancePoolsAPI
}

func NewInstancePoolsClient(cfg *config.Config) (*InstancePoolsClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &InstancePoolsClient{
		InstancePoolsAPI: InstancePoolsAPI{
			instancePoolsImpl: instancePoolsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type InstanceProfilesClient struct {
	InstanceProfilesAPI
}

func NewInstanceProfilesClient(cfg *config.Config) (*InstanceProfilesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &InstanceProfilesClient{
		InstanceProfilesAPI: InstanceProfilesAPI{
			instanceProfilesImpl: instanceProfilesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type LibrariesClient struct {
	LibrariesAPI
}

func NewLibrariesClient(cfg *config.Config) (*LibrariesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LibrariesClient{
		LibrariesAPI: LibrariesAPI{
			librariesImpl: librariesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type PolicyComplianceForClustersClient struct {
	PolicyComplianceForClustersAPI
}

func NewPolicyComplianceForClustersClient(cfg *config.Config) (*PolicyComplianceForClustersClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PolicyComplianceForClustersClient{
		PolicyComplianceForClustersAPI: PolicyComplianceForClustersAPI{
			policyComplianceForClustersImpl: policyComplianceForClustersImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

type PolicyFamiliesClient struct {
	PolicyFamiliesAPI
}

func NewPolicyFamiliesClient(cfg *config.Config) (*PolicyFamiliesClient, error) {
	if cfg == nil {
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}

	if cfg.IsAccountClient() {
		return nil, errors.New("invalid configuration: please provide a valid workspace config for the requested workspace service client")
	}

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PolicyFamiliesClient{
		PolicyFamiliesAPI: PolicyFamiliesAPI{
			policyFamiliesImpl: policyFamiliesImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
