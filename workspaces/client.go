package workspaces

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/commands"
	"github.com/databricks/databricks-sdk-go/service/datasharing"
	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/service/deltapipelines"
	"github.com/databricks/databricks-sdk-go/service/gitcredentials"
	"github.com/databricks/databricks-sdk-go/service/globalinitscripts"
	"github.com/databricks/databricks-sdk-go/service/groups"
	"github.com/databricks/databricks-sdk-go/service/instancepools"
	"github.com/databricks/databricks-sdk-go/service/instanceprofiles"
	"github.com/databricks/databricks-sdk-go/service/ipaccesslists"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/libraries"
	"github.com/databricks/databricks-sdk-go/service/mlflow"
	"github.com/databricks/databricks-sdk-go/service/mlflowext"
	"github.com/databricks/databricks-sdk-go/service/modelregistry"
	"github.com/databricks/databricks-sdk-go/service/modelregistrywebhooks"
	"github.com/databricks/databricks-sdk-go/service/modelversioncomments"
	"github.com/databricks/databricks-sdk-go/service/permissions"
	"github.com/databricks/databricks-sdk-go/service/repos"
	"github.com/databricks/databricks-sdk-go/service/secrets"
	"github.com/databricks/databricks-sdk-go/service/serviceprincipals"
	"github.com/databricks/databricks-sdk-go/service/tokenmanagement"
	"github.com/databricks/databricks-sdk-go/service/tokens"
	"github.com/databricks/databricks-sdk-go/service/transitionrequests"
	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/databricks-sdk-go/service/users"
	"github.com/databricks/databricks-sdk-go/service/warehouses"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

type WorkspacesClient struct {
	Config *databricks.Config

	Clusters              clusters.ClustersService
	Commands              *commands.CommandsAPI
	Dbfs                  dbfs.DbfsService
	DeltaPipelines        deltapipelines.DeltapipelinesService
	DeltaSharing          datasharing.DatasharingService
	GitCredentials        gitcredentials.GitcredentialsService
	GlobalInitScripts     globalinitscripts.GlobalinitscriptsService
	Groups                groups.GroupsService
	Jobs                  jobs.JobsService
	InstancePools         instancepools.InstancepoolsService
	InstanceProfiles      instanceprofiles.InstanceprofilesService
	IpAccessLists         ipaccesslists.IpaccesslistsService
	Libraries             libraries.LibrariesService
	MLflow                mlflow.MlflowService
	MLflowExt             mlflowext.MlflowextService
	ModelRegistry         modelregistry.ModelregistryService
	ModelRegistryWebhooks modelregistrywebhooks.ModelregistrywebhooksService
	ModelVersionComments  modelversioncomments.ModelversioncommentsService
	Permissions           permissions.PermissionsService
	Repos                 repos.ReposService
	Secrets               secrets.SecretsService
	ServicePrincipals     serviceprincipals.ServiceprincipalsService
	Tokens                tokens.TokensService
	TokenManagement       tokenmanagement.TokenmanagementService
	TransitionRequests    transitionrequests.TransitionrequestsService
	UnityCatalog          unitycatalog.UnitycatalogService
	Users                 users.UsersService
	Warehouses            warehouses.WarehousesService
	Workspace             workspace.WorkspaceService
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
	apiClient := client.New(cfg)
	return &WorkspacesClient{
		Config:                cfg,
		Clusters:              clusters.New(apiClient),
		Commands:              commands.New(cfg),
		Dbfs:                  dbfs.New(apiClient),
		DeltaPipelines:        deltapipelines.New(apiClient),
		DeltaSharing:          datasharing.New(apiClient),
		GitCredentials:        gitcredentials.New(apiClient),
		GlobalInitScripts:     globalinitscripts.New(apiClient),
		Groups:                groups.New(apiClient),
		Jobs:                  jobs.New(apiClient),
		InstancePools:         instancepools.New(apiClient),
		InstanceProfiles:      instanceprofiles.New(apiClient),
		IpAccessLists:         ipaccesslists.New(apiClient),
		Libraries:             libraries.New(apiClient),
		MLflow:                mlflow.New(apiClient),
		MLflowExt:             mlflowext.New(apiClient),
		ModelRegistry:         modelregistry.New(apiClient),
		ModelRegistryWebhooks: modelregistrywebhooks.New(apiClient),
		ModelVersionComments:  modelversioncomments.New(apiClient),
		Permissions:           permissions.New(apiClient),
		Repos:                 repos.New(apiClient),
		Secrets:               secrets.New(apiClient),
		ServicePrincipals:     serviceprincipals.New(apiClient),
		Tokens:                tokens.New(apiClient),
		TokenManagement:       tokenmanagement.New(apiClient),
		TransitionRequests:    transitionrequests.New(apiClient),
		UnityCatalog:          unitycatalog.New(apiClient),
		Users:                 users.New(apiClient),
		Warehouses:            warehouses.New(apiClient),
		Workspace:             workspace.New(apiClient),
	}
}
