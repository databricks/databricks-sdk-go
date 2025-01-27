// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ClusterPoliciesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	ClusterPolicies ClusterPoliciesInterface
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
		cfg:             cfg,
		apiClient:       apiClient,
		ClusterPolicies: NewClusterPolicies(databricksClient),
	}, nil
}

type ClustersClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Clusters ClustersInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Clusters:  NewClusters(databricksClient),
	}, nil
}

type CommandExecutionClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	CommandExecution CommandExecutionInterface
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
		cfg:              cfg,
		apiClient:        apiClient,
		CommandExecution: NewCommandExecution(databricksClient),
	}, nil
}

type GlobalInitScriptsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	GlobalInitScripts GlobalInitScriptsInterface
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
		cfg:               cfg,
		apiClient:         apiClient,
		GlobalInitScripts: NewGlobalInitScripts(databricksClient),
	}, nil
}

type InstancePoolsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	InstancePools InstancePoolsInterface
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
		cfg:           cfg,
		apiClient:     apiClient,
		InstancePools: NewInstancePools(databricksClient),
	}, nil
}

type InstanceProfilesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	InstanceProfiles InstanceProfilesInterface
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
		cfg:              cfg,
		apiClient:        apiClient,
		InstanceProfiles: NewInstanceProfiles(databricksClient),
	}, nil
}

type LibrariesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Libraries LibrariesInterface
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
		cfg:       cfg,
		apiClient: apiClient,
		Libraries: NewLibraries(databricksClient),
	}, nil
}

type PolicyComplianceForClustersClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	PolicyComplianceForClusters PolicyComplianceForClustersInterface
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
		cfg:                         cfg,
		apiClient:                   apiClient,
		PolicyComplianceForClusters: NewPolicyComplianceForClusters(databricksClient),
	}, nil
}

type PolicyFamiliesClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	PolicyFamilies PolicyFamiliesInterface
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
		cfg:            cfg,
		apiClient:      apiClient,
		PolicyFamilies: NewPolicyFamilies(databricksClient),
	}, nil
}
