package workspaces

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/commands"
	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/service/deltapipelines"
	"github.com/databricks/databricks-sdk-go/service/datasharing"
	"github.com/databricks/databricks-sdk-go/service/groups"
	"github.com/databricks/databricks-sdk-go/service/instancepools"
	"github.com/databricks/databricks-sdk-go/service/instanceprofiles"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/libraries"
	"github.com/databricks/databricks-sdk-go/service/serviceprincipals"
	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/databricks-sdk-go/service/users"
)

type WorkspacesClient struct {
	Config *databricks.Config

	Clusters          clusters.ClustersService
	Commands          *commands.CommandsAPI
	Dbfs              dbfs.DbfsService
	DeltaPipelines    deltapipelines.DeltapipelinesService
	DeltaSharing      datasharing.DatasharingService
	Groups            groups.GroupsService
	Jobs              jobs.JobsService
	InstancePools     instancepools.InstancepoolsService
	InstanceProfiles  instanceprofiles.InstanceprofilesService
	Libraries         libraries.LibrariesService
	ServicePrincipals serviceprincipals.ServiceprincipalsService
	UnityCatalog      unitycatalog.UnitycatalogService
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
		Clusters:          clusters.New(apiClient),
		Commands:          commands.New(cfg),
		Dbfs:              dbfs.New(apiClient),
		DeltaSharing:      datasharing.New(apiClient),
		DeltaPipelines:    deltapipelines.New(apiClient),
		Groups:            groups.New(apiClient),
		Jobs:              jobs.New(apiClient),
		InstancePools:     instancepools.New(apiClient),
		InstanceProfiles:  instanceprofiles.New(apiClient),
		Libraries:         libraries.New(apiClient),
		ServicePrincipals: serviceprincipals.New(apiClient),
		UnityCatalog:      unitycatalog.New(apiClient),
		Users:             users.New(apiClient),
	}
}
