// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package compute

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type ClusterPoliciesClient struct {
	ClusterPoliciesInterface
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
		ClusterPoliciesInterface: NewClusterPolicies(databricksClient),
	}, nil
}

type ClustersClient struct {
	ClustersInterface
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
		ClustersInterface: NewClusters(databricksClient),
	}, nil
}

type CommandExecutionClient struct {
	CommandExecutionInterface
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
		CommandExecutionInterface: NewCommandExecution(databricksClient),
	}, nil
}

type GlobalInitScriptsClient struct {
	GlobalInitScriptsInterface
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
		GlobalInitScriptsInterface: NewGlobalInitScripts(databricksClient),
	}, nil
}

type InstancePoolsClient struct {
	InstancePoolsInterface
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
		InstancePoolsInterface: NewInstancePools(databricksClient),
	}, nil
}

type InstanceProfilesClient struct {
	InstanceProfilesInterface
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
		InstanceProfilesInterface: NewInstanceProfiles(databricksClient),
	}, nil
}

type LibrariesClient struct {
	LibrariesInterface
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
		LibrariesInterface: NewLibraries(databricksClient),
	}, nil
}

type PolicyComplianceForClustersClient struct {
	PolicyComplianceForClustersInterface
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
		PolicyComplianceForClustersInterface: NewPolicyComplianceForClusters(databricksClient),
	}, nil
}

type PolicyFamiliesClient struct {
	PolicyFamiliesInterface
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
		PolicyFamiliesInterface: NewPolicyFamilies(databricksClient),
	}, nil
}
