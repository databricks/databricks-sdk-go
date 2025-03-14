// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type JobsClient struct {
	JobsInterface
	apiClient *httpclient.ApiClient
}

func NewJobsClient(cfg *config.Config) (*JobsClient, error) {
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

	return &JobsClient{
		apiClient:     databricksClient.ApiClient(),
		JobsInterface: NewJobs(databricksClient),
	}, nil
}

type PolicyComplianceForJobsClient struct {
	PolicyComplianceForJobsInterface
	apiClient *httpclient.ApiClient
}

func NewPolicyComplianceForJobsClient(cfg *config.Config) (*PolicyComplianceForJobsClient, error) {
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

	return &PolicyComplianceForJobsClient{
		apiClient:                        databricksClient.ApiClient(),
		PolicyComplianceForJobsInterface: NewPolicyComplianceForJobs(databricksClient),
	}, nil
}
