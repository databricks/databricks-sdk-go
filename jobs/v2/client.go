// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package jobs

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
)

// The Jobs API allows you to create, edit, and delete jobs.
//
// You can use a Databricks job to run a data processing or data analysis task
// in a Databricks cluster with scalable resources. Your job can consist of a
// single task or can be a large, multi-task workflow with complex dependencies.
// Databricks manages the task orchestration, cluster management, monitoring,
// and error reporting for all of your jobs. You can run your jobs immediately
// or periodically through an easy-to-use scheduling system. You can implement
// job tasks using notebooks, JARS, Delta Live Tables pipelines, or Python,
// Scala, Spark submit, and Java applications.
//
// You should never hard code secrets or store them in plain text. Use the
// [Secrets CLI] to manage secrets in the [Databricks CLI]. Use the [Secrets
// utility] to reference secrets in notebooks and jobs.
//
// [Databricks CLI]: https://docs.databricks.com/dev-tools/cli/index.html
// [Secrets CLI]: https://docs.databricks.com/dev-tools/cli/secrets-cli.html
// [Secrets utility]: https://docs.databricks.com/dev-tools/databricks-utils.html#dbutils-secrets
type JobsClient struct {
	JobsAPI
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
		JobsAPI: JobsAPI{
			jobsImpl: jobsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}

// The compliance APIs allow you to view and manage the policy compliance status
// of jobs in your workspace. This API currently only supports compliance
// controls for cluster policies.
//
// A job is in compliance if its cluster configurations satisfy the rules of all
// their respective cluster policies. A job could be out of compliance if a
// cluster policy it uses was updated after the job was last edited. The job is
// considered out of compliance if any of its clusters no longer comply with
// their updated policies.
//
// The get and list compliance APIs allow you to view the policy compliance
// status of a job. The enforce compliance API allows you to update a job so
// that it becomes compliant with all of its policies.
type PolicyComplianceForJobsClient struct {
	PolicyComplianceForJobsAPI
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
		PolicyComplianceForJobsAPI: PolicyComplianceForJobsAPI{
			policyComplianceForJobsImpl: policyComplianceForJobsImpl{
				client: apiClient.ApiClient(),
			},
		},
	}, nil
}
