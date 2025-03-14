// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ClusterPoliciesClient struct {
	ClusterPoliciesInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ClusterPoliciesClient{
		apiClient:                databricksClient.ApiClient(),
		ClusterPoliciesInterface: NewClusterPolicies(databricksClient),
	}, nil
}

type ClustersClient struct {
	ClustersInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &ClustersClient{
		apiClient:         databricksClient.ApiClient(),
		ClustersInterface: NewClusters(databricksClient),
	}, nil
}

type CommandExecutionClient struct {
	CommandExecutionInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &CommandExecutionClient{
		apiClient:                 databricksClient.ApiClient(),
		CommandExecutionInterface: NewCommandExecution(databricksClient),
	}, nil
}

type GlobalInitScriptsClient struct {
	GlobalInitScriptsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &GlobalInitScriptsClient{
		apiClient:                  databricksClient.ApiClient(),
		GlobalInitScriptsInterface: NewGlobalInitScripts(databricksClient),
	}, nil
}

type InstancePoolsClient struct {
	InstancePoolsInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &InstancePoolsClient{
		apiClient:              databricksClient.ApiClient(),
		InstancePoolsInterface: NewInstancePools(databricksClient),
	}, nil
}

type InstanceProfilesClient struct {
	InstanceProfilesInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &InstanceProfilesClient{
		apiClient:                 databricksClient.ApiClient(),
		InstanceProfilesInterface: NewInstanceProfiles(databricksClient),
	}, nil
}

type LibrariesClient struct {
	LibrariesInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &LibrariesClient{
		apiClient:          databricksClient.ApiClient(),
		LibrariesInterface: NewLibraries(databricksClient),
	}, nil
}

type PolicyComplianceForClustersClient struct {
	PolicyComplianceForClustersInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PolicyComplianceForClustersClient{
		apiClient:                            databricksClient.ApiClient(),
		PolicyComplianceForClustersInterface: NewPolicyComplianceForClusters(databricksClient),
	}, nil
}

type PolicyFamiliesClient struct {
	PolicyFamiliesInterface
	apiClient *httpclient.ApiClient
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
	databricksClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PolicyFamiliesClient{
		apiClient:               databricksClient.ApiClient(),
		PolicyFamiliesInterface: NewPolicyFamilies(databricksClient),
	}, nil
}
