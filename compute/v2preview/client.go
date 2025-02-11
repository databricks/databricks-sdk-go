// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package computepreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ClusterPoliciesClient struct {
	ClusterPoliciesInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ClusterPoliciesClient{
		Config:                   cfg,
		apiClient:                apiClient,
		ClusterPoliciesInterface: NewClusterPolicies(databricksClient),
	}, nil
}

type ClustersClient struct {
	ClustersInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &ClustersClient{
		Config:            cfg,
		apiClient:         apiClient,
		ClustersInterface: NewClusters(databricksClient),
	}, nil
}

type CommandExecutionClient struct {
	CommandExecutionInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &CommandExecutionClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		CommandExecutionInterface: NewCommandExecution(databricksClient),
	}, nil
}

type GlobalInitScriptsClient struct {
	GlobalInitScriptsInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &GlobalInitScriptsClient{
		Config:                     cfg,
		apiClient:                  apiClient,
		GlobalInitScriptsInterface: NewGlobalInitScripts(databricksClient),
	}, nil
}

type InstancePoolsClient struct {
	InstancePoolsInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &InstancePoolsClient{
		Config:                 cfg,
		apiClient:              apiClient,
		InstancePoolsInterface: NewInstancePools(databricksClient),
	}, nil
}

type InstanceProfilesClient struct {
	InstanceProfilesInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &InstanceProfilesClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		InstanceProfilesInterface: NewInstanceProfiles(databricksClient),
	}, nil
}

type LibrariesClient struct {
	LibrariesInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &LibrariesClient{
		Config:             cfg,
		apiClient:          apiClient,
		LibrariesInterface: NewLibraries(databricksClient),
	}, nil
}

type PolicyComplianceForClustersClient struct {
	PolicyComplianceForClustersInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &PolicyComplianceForClustersClient{
		Config:                               cfg,
		apiClient:                            apiClient,
		PolicyComplianceForClustersInterface: NewPolicyComplianceForClusters(databricksClient),
	}, nil
}

type PolicyFamiliesClient struct {
	PolicyFamiliesInterface
	Config    *config.Config
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
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &PolicyFamiliesClient{
		Config:                  cfg,
		apiClient:               apiClient,
		PolicyFamiliesInterface: NewPolicyFamilies(databricksClient),
	}, nil
}
