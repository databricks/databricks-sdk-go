// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobspreview

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type JobsPreviewClient struct {
	JobsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewJobsPreviewClient(cfg *config.Config) (*JobsPreviewClient, error) {
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

	return &JobsPreviewClient{
		Config:               cfg,
		apiClient:            apiClient,
		JobsPreviewInterface: NewJobsPreview(databricksClient),
	}, nil
}

type PolicyComplianceForJobsPreviewClient struct {
	PolicyComplianceForJobsPreviewInterface
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

func NewPolicyComplianceForJobsPreviewClient(cfg *config.Config) (*PolicyComplianceForJobsPreviewClient, error) {
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

	return &PolicyComplianceForJobsPreviewClient{
		Config:                                  cfg,
		apiClient:                               apiClient,
		PolicyComplianceForJobsPreviewInterface: NewPolicyComplianceForJobsPreview(databricksClient),
	}, nil
}
