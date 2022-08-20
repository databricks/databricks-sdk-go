package workspaces

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/commands"
	"github.com/databricks/databricks-sdk-go/service/groups"
	"github.com/databricks/databricks-sdk-go/service/serviceprincipals"
	"github.com/databricks/databricks-sdk-go/service/users"
)

type WorkspacesClient struct {
	Config *databricks.Config

	Clusters          *clusters.ClustersAPI
	Commands          *commands.CommandsAPI
	Groups            groups.GroupsService
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
	return &WorkspacesClient{
		Config:            cfg,
		Clusters:          clusters.New(cfg),
		Commands:          commands.New(cfg),
		Groups:            groups.New(cfg),
		ServicePrincipals: serviceprincipals.New(cfg),
		Users:             users.New(cfg),
	}
}
