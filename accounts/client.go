package accounts

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/accountpermissionassignments"
	"github.com/databricks/databricks-sdk-go/service/budgets"
	"github.com/databricks/databricks-sdk-go/service/credentials"
	"github.com/databricks/databricks-sdk-go/service/customermanagedkeys"
	"github.com/databricks/databricks-sdk-go/service/logdelivery"
	"github.com/databricks/databricks-sdk-go/service/networks"
	"github.com/databricks/databricks-sdk-go/service/privateaccesssettings"
	"github.com/databricks/databricks-sdk-go/service/storageconfigurations"
	"github.com/databricks/databricks-sdk-go/service/usagedownload"
	"github.com/databricks/databricks-sdk-go/service/vpcendpoints"
	"github.com/databricks/databricks-sdk-go/service/workspaces"
)

type AccountsClient struct {
	Config                       *databricks.Config
	AccountPermissionAssignments accountpermissionassignments.AccountPermissionAssignmentsService
	Budgets                      budgets.BudgetsService
	Credentials                  credentials.CredentialsService
	CustomerManagedKeys          customermanagedkeys.CustomermanagedkeysService
	LogDelivery                  logdelivery.LogdeliveryService
	Networks                     networks.NetworkConfigurationsService
	PrivateAccessSettings        privateaccesssettings.PrivateaccesssettingsService
	StorageConfigurations        storageconfigurations.StorageconfigurationsService
	UsageDownload                usagedownload.UsagedownloadService
	VpcEndpoints                 vpcendpoints.VpcendpointsService
	Workspaces                   workspaces.WorkspacesService
}

func New(c ...*databricks.Config) *AccountsClient {
	var cfg *databricks.Config
	if len(c) == 1 {
		// first config
		cfg = c[0]
	} else {
		// default config
		cfg = &databricks.Config{}
	}
	apiClient := client.New(cfg)
	return &AccountsClient{
		Config:                       cfg,
		AccountPermissionAssignments: accountpermissionassignments.NewAccountPermissionAssignments(apiClient),
		Budgets:                      budgets.New(apiClient),
		Credentials:                  credentials.New(apiClient),
		CustomerManagedKeys:          customermanagedkeys.New(apiClient),
		LogDelivery:                  logdelivery.New(apiClient),
		Networks:                     networks.NewNetworkConfigurations(apiClient),
		PrivateAccessSettings:        privateaccesssettings.New(apiClient),
		StorageConfigurations:        storageconfigurations.New(apiClient),
		UsageDownload:                usagedownload.New(apiClient),
		VpcEndpoints:                 vpcendpoints.New(apiClient),
		Workspaces:                   workspaces.NewWorkspaces(apiClient),
	}
}
