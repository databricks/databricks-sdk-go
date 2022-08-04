package workspaces

import (
	"github.com/databricks/sdk-go/databricks"
	"github.com/databricks/sdk-go/service/clusters"
	"github.com/databricks/sdk-go/service/commands"
)

type WorkspacesClient struct {
	Config *databricks.Config

	Clusters *clusters.ClustersAPI
	Commands *commands.CommandsAPI
}

func New(c ...*databricks.Config) *WorkspacesClient {
	var cfg *databricks.Config
	if len(c) == 1 {
		// first config
		cfg = c[0]
	} else {
		// default config
		cfg = &databricks.Config{}
	}
	return &WorkspacesClient{
		Config:   cfg,
		Clusters: clusters.New(cfg),
		Commands: commands.New(cfg),
	}
}
