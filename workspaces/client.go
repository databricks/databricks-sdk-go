package workspaces

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/clusterpolicies"
	"github.com/databricks/databricks-sdk-go/service/clusters"
	"github.com/databricks/databricks-sdk-go/service/commands"
	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/service/deltapipelines"
	"github.com/databricks/databricks-sdk-go/service/gitcredentials"
	"github.com/databricks/databricks-sdk-go/service/globalinitscripts"
	"github.com/databricks/databricks-sdk-go/service/instancepools"
	"github.com/databricks/databricks-sdk-go/service/instanceprofiles"
	"github.com/databricks/databricks-sdk-go/service/ipaccesslists"
	"github.com/databricks/databricks-sdk-go/service/jobs"
	"github.com/databricks/databricks-sdk-go/service/libraries"
	"github.com/databricks/databricks-sdk-go/service/mlflow"
	"github.com/databricks/databricks-sdk-go/service/permissions"
	"github.com/databricks/databricks-sdk-go/service/repos"
	"github.com/databricks/databricks-sdk-go/service/scim"
	"github.com/databricks/databricks-sdk-go/service/secrets"
	"github.com/databricks/databricks-sdk-go/service/tokenmanagement"
	"github.com/databricks/databricks-sdk-go/service/tokens"
	"github.com/databricks/databricks-sdk-go/service/unitycatalog"
	"github.com/databricks/databricks-sdk-go/service/warehouses"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/databricks-sdk-go/service/workspaceconf"
)

type WorkspacesClient struct {
	Config *databricks.Config

	Catalogs             unitycatalog.CatalogsService
	ClusterPolicies      clusterpolicies.ClusterPoliciesService
	Clusters             clusters.ClustersService
	Commands             *commands.CommandsAPI
	CurrentUser          scim.CurrentUserService
	Dbfs                 dbfs.DbfsService
	DeltaPipelines       deltapipelines.DeltaPipelinesService
	Experiments          mlflow.ExperimentsService
	ExternalLocations    unitycatalog.ExternalLocationsService
	GitCredentials       gitcredentials.GitCredentialsService
	GlobalInitScripts    globalinitscripts.GlobalInitScriptsService
	Grants               unitycatalog.GrantsService
	Groups               scim.GroupsService
	Jobs                 jobs.JobsService
	InstancePools        instancepools.InstancePoolsService
	InstanceProfiles     instanceprofiles.InstanceprofilesService
	IpAccessLists        ipaccesslists.IpAccessListsService
	Libraries            libraries.LibrariesService
	Metastores           unitycatalog.MetastoresService
	MLflowArtifacts      mlflow.MLflowArtifactsService
	MLflowDatabricks     mlflow.MLflowDatabricksService
	MLflowMetrics        mlflow.MLflowMetricsService
	MLflowRuns           mlflow.MLflowRunsService
	ModelVersions        mlflow.ModelVersionsService
	ModelVersionComments mlflow.ModelVersionCommentsService
	Permissions          permissions.PermissionsService
	Providers            unitycatalog.ProvidersService
	RecipientActivation  unitycatalog.RecipientActivationService
	Recipients           unitycatalog.RecipientsService
	RegisteredModels     mlflow.RegisteredModelsService
	RegistryWebhooks     mlflow.RegistryWebhooksService
	Repos                repos.ReposService
	Schemas              unitycatalog.SchemasService
	Secrets              secrets.SecretsService
	Shares               unitycatalog.SharesService
	ServicePrincipals    scim.ServicePrincipalsService
	StorageCredentials   unitycatalog.StorageCredentialsService
	Tables               unitycatalog.TablesService
	Tokens               tokens.TokensService
	TokenManagement      tokenmanagement.TokenManagementService
	TransitionRequests   mlflow.TransitionRequestsService
	UnityFiles           unitycatalog.UnityFilesService
	Users                scim.UsersService
	Warehouses           warehouses.WarehousesService
	Workspace            workspace.WorkspaceService
	WorkspaceConf        workspaceconf.WorkspaceConfService
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
		Config:               cfg,
		Catalogs:             unitycatalog.NewCatalogs(apiClient),
		ClusterPolicies:      clusterpolicies.NewClusterPolicies(apiClient),
		Clusters:             clusters.NewClusters(apiClient),
		Commands:             commands.New(cfg),
		CurrentUser:          scim.NewCurrentUser(apiClient),
		Dbfs:                 dbfs.NewDbfs(apiClient),
		DeltaPipelines:       deltapipelines.NewDeltaPipelines(apiClient),
		Experiments:          mlflow.NewExperiments(apiClient),
		ExternalLocations:    unitycatalog.NewExternalLocations(apiClient),
		GitCredentials:       gitcredentials.NewGitCredentials(apiClient),
		GlobalInitScripts:    globalinitscripts.NewGlobalInitScripts(apiClient),
		Grants:               unitycatalog.NewGrants(apiClient),
		Groups:               scim.NewGroups(apiClient),
		Jobs:                 jobs.NewJobs(apiClient),
		InstancePools:        instancepools.NewInstancePools(apiClient),
		InstanceProfiles:     instanceprofiles.New(apiClient),
		IpAccessLists:        ipaccesslists.NewIpAccessLists(apiClient),
		Libraries:            libraries.NewLibraries(apiClient),
		Metastores:           unitycatalog.NewMetastores(apiClient),
		MLflowArtifacts:      mlflow.NewMLflowArtifacts(apiClient),
		MLflowDatabricks:     mlflow.NewMLflowDatabricks(apiClient),
		MLflowMetrics:        mlflow.NewMLflowMetrics(apiClient),
		MLflowRuns:           mlflow.NewMLflowRuns(apiClient),
		ModelVersions:        mlflow.NewModelVersions(apiClient),
		ModelVersionComments: mlflow.NewModelVersionComments(apiClient),
		Permissions:          permissions.NewPermissions(apiClient),
		Providers:            unitycatalog.NewProviders(apiClient),
		RecipientActivation:  unitycatalog.NewRecipientActivation(apiClient),
		Recipients:           unitycatalog.NewRecipients(apiClient),
		RegisteredModels:     mlflow.NewRegisteredModels(apiClient),
		RegistryWebhooks:     mlflow.NewRegistryWebhooks(apiClient),
		Repos:                repos.NewRepos(apiClient),
		Schemas:              unitycatalog.NewSchemas(apiClient),
		Secrets:              secrets.NewSecrets(apiClient),
		Shares:               unitycatalog.NewShares(apiClient),
		ServicePrincipals:    scim.NewServicePrincipals(apiClient),
		StorageCredentials:   unitycatalog.NewStorageCredentials(apiClient),
		Tables:               unitycatalog.NewTables(apiClient),
		Tokens:               tokens.NewTokens(apiClient),
		TokenManagement:      tokenmanagement.NewTokenManagement(apiClient),
		TransitionRequests:   mlflow.NewTransitionRequests(apiClient),
		UnityFiles:           unitycatalog.NewUnityFiles(apiClient),
		Users:                scim.NewUsers(apiClient),
		Warehouses:           warehouses.NewWarehouses(apiClient),
		Workspace:            workspace.NewWorkspace(apiClient),
		WorkspaceConf:        workspaceconf.NewWorkspaceConf(apiClient),
	}
}
