package workspaces

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/commands"
	"github.com/databricks/databricks-sdk-go/service/groups"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/serviceprincipals"
	"github.com/databricks/databricks-sdk-go/service/users"
)

type WorkspacesClient struct {
	Config *databricks.Config

	Clusters          *clusters.ClustersAPI
	Commands          *commands.CommandsAPI
	Groups            groups.GroupsService
	Jobs              jobs.JobsService
	ServicePrincipals serviceprincipals.ServiceprincipalsService
	Users             users.UsersService
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
	// TODO: migrate other signatures here
	apiClient := client.New(cfg)
	return &WorkspacesClient{
		Config:            cfg,
		Clusters:          clusters.New(cfg),
		Commands:          commands.New(cfg),
		Groups:            groups.New(apiClient),
		Jobs:              jobs.New(apiClient),
		ServicePrincipals: serviceprincipals.New(apiClient),
		Users:             users.New(apiClient),
	}
}
