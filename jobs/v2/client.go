package jobs

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/httpclient"
)

type JobsClient struct {
	cfg       *config.Config
	apiClient *httpclient.ApiClient

	Jobs JobsInterface
}

func NewJobsClientFromConfig(c ...*config.Config) (*JobsClient, error) {
	var cfg *config.Config
	if len(c) == 1 {
		// first config
		cfg = (*config.Config)(c[0])
	} else {
		// default config
		cfg = &config.Config{}
	}
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	if cfg.IsAccountClient() {
		return nil, errors.New("account config for a workspace client is not supported")
	}
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}
	databricksClient, err := client.NewWithClient(cfg, apiClient)
	if err != nil {
		return nil, err
	}

	return &JobsClient{
		cfg:       cfg,
		apiClient: apiClient,
		Jobs:      NewJobs(databricksClient),
	}, nil
}
