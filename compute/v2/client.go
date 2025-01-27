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
	cfg       *config.Config
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
		cfg:                      cfg,
		apiClient:                apiClient,
		ClusterPoliciesInterface: NewClusterPolicies(databricksClient),
	}, nil
}

type ClustersClient struct {
	ClustersInterface
	cfg       *config.Config
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
		cfg:               cfg,
		apiClient:         apiClient,
		ClustersInterface: NewClusters(databricksClient),
	}, nil
}

type CommandExecutionClient struct {
	CommandExecutionInterface
	cfg       *config.Config
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		CommandExecutionInterface: NewCommandExecution(databricksClient),
	}, nil
}

type GlobalInitScriptsClient struct {
	GlobalInitScriptsInterface
	cfg       *config.Config
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
		cfg:                        cfg,
		apiClient:                  apiClient,
		GlobalInitScriptsInterface: NewGlobalInitScripts(databricksClient),
	}, nil
}

type InstancePoolsClient struct {
	InstancePoolsInterface
	cfg       *config.Config
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
		cfg:                    cfg,
		apiClient:              apiClient,
		InstancePoolsInterface: NewInstancePools(databricksClient),
	}, nil
}

type InstanceProfilesClient struct {
	InstanceProfilesInterface
	cfg       *config.Config
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
		cfg:                       cfg,
		apiClient:                 apiClient,
		InstanceProfilesInterface: NewInstanceProfiles(databricksClient),
	}, nil
}

type LibrariesClient struct {
	LibrariesInterface
	cfg       *config.Config
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
		cfg:                cfg,
		apiClient:          apiClient,
		LibrariesInterface: NewLibraries(databricksClient),
	}, nil
}

type PolicyComplianceForClustersClient struct {
	PolicyComplianceForClustersInterface
	cfg       *config.Config
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
		cfg:                                  cfg,
		apiClient:                            apiClient,
		PolicyComplianceForClustersInterface: NewPolicyComplianceForClusters(databricksClient),
	}, nil
}

type PolicyFamiliesClient struct {
	PolicyFamiliesInterface
	cfg       *config.Config
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
		cfg:                     cfg,
		apiClient:               apiClient,
		PolicyFamiliesInterface: NewPolicyFamilies(databricksClient),
	}, nil
}
