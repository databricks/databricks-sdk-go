// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

type JobsClient struct {
	JobsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &JobsClient{
		JobsInterface: NewJobs(apiClient.ApiClient()),
	}, nil
}

type PolicyComplianceForJobsClient struct {
	PolicyComplianceForJobsInterface
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

	apiClient, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	return &PolicyComplianceForJobsClient{
		PolicyComplianceForJobsInterface: NewPolicyComplianceForJobs(apiClient.ApiClient()),
	}, nil
}
