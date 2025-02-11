// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package computepreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type ClusterPoliciesPreviewClient struct {
	ClusterPoliciesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewClusterPoliciesPreviewClient(cfg *config.Config) (*ClusterPoliciesPreviewClient, error) {
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

	return &ClusterPoliciesPreviewClient{
		Config:                          cfg,
		apiClient:                       apiClient,
		ClusterPoliciesPreviewInterface: NewClusterPoliciesPreview(databricksClient),
	}, nil
}

type ClustersPreviewClient struct {
	ClustersPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewClustersPreviewClient(cfg *config.Config) (*ClustersPreviewClient, error) {
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

	return &ClustersPreviewClient{
		Config:                   cfg,
		apiClient:                apiClient,
		ClustersPreviewInterface: NewClustersPreview(databricksClient),
	}, nil
}

type CommandExecutionPreviewClient struct {
	CommandExecutionPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewCommandExecutionPreviewClient(cfg *config.Config) (*CommandExecutionPreviewClient, error) {
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

	return &CommandExecutionPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		CommandExecutionPreviewInterface: NewCommandExecutionPreview(databricksClient),
	}, nil
}

type GlobalInitScriptsPreviewClient struct {
	GlobalInitScriptsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewGlobalInitScriptsPreviewClient(cfg *config.Config) (*GlobalInitScriptsPreviewClient, error) {
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

	return &GlobalInitScriptsPreviewClient{
		Config:                            cfg,
		apiClient:                         apiClient,
		GlobalInitScriptsPreviewInterface: NewGlobalInitScriptsPreview(databricksClient),
	}, nil
}

type InstancePoolsPreviewClient struct {
	InstancePoolsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewInstancePoolsPreviewClient(cfg *config.Config) (*InstancePoolsPreviewClient, error) {
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

	return &InstancePoolsPreviewClient{
		Config:                        cfg,
		apiClient:                     apiClient,
		InstancePoolsPreviewInterface: NewInstancePoolsPreview(databricksClient),
	}, nil
}

type InstanceProfilesPreviewClient struct {
	InstanceProfilesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewInstanceProfilesPreviewClient(cfg *config.Config) (*InstanceProfilesPreviewClient, error) {
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

	return &InstanceProfilesPreviewClient{
		Config:                           cfg,
		apiClient:                        apiClient,
		InstanceProfilesPreviewInterface: NewInstanceProfilesPreview(databricksClient),
	}, nil
}

type LibrariesPreviewClient struct {
	LibrariesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewLibrariesPreviewClient(cfg *config.Config) (*LibrariesPreviewClient, error) {
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

	return &LibrariesPreviewClient{
		Config:                    cfg,
		apiClient:                 apiClient,
		LibrariesPreviewInterface: NewLibrariesPreview(databricksClient),
	}, nil
}

type PolicyComplianceForClustersPreviewClient struct {
	PolicyComplianceForClustersPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewPolicyComplianceForClustersPreviewClient(cfg *config.Config) (*PolicyComplianceForClustersPreviewClient, error) {
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

	return &PolicyComplianceForClustersPreviewClient{
		Config:    cfg,
		apiClient: apiClient,
		PolicyComplianceForClustersPreviewInterface: NewPolicyComplianceForClustersPreview(databricksClient),
	}, nil
}

type PolicyFamiliesPreviewClient struct {
	PolicyFamiliesPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewPolicyFamiliesPreviewClient(cfg *config.Config) (*PolicyFamiliesPreviewClient, error) {
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

	return &PolicyFamiliesPreviewClient{
		Config:                         cfg,
		apiClient:                      apiClient,
		PolicyFamiliesPreviewInterface: NewPolicyFamiliesPreview(databricksClient),
	}, nil
}
