// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package databricks

import (
	"errors"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
)

type WorkspaceClient struct {
	Config    *config.Config
	apiClient *httpclient.ApiClient
}

var ErrNotWorkspaceClient = errors.New("invalid Databricks Workspace configuration")

// NewWorkspaceClient creates new Databricks SDK client for Workspaces or
// returns error in case configuration is wrong
func NewWorkspaceClient(c ...*Config) (*WorkspaceClient, error) {
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
		return nil, ErrNotWorkspaceClient
	}
	apiClient, err := cfg.NewApiClient()
	if err != nil {
		return nil, err
	}

	return &WorkspaceClient{
		Config:    cfg,
		apiClient: apiClient,
	}, nil
}
