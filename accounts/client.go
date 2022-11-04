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
	Credentials                  credentials.CredentialConfigurationsService
	CustomerManagedKeys          customermanagedkeys.KeyConfigurationsService
	LogDelivery                  logdelivery.LogDeliveryService
	Networks                     networks.NetworkConfigurationsService
	PrivateAccessSettings        privateaccesssettings.PrivateAccessSettingsService
	StorageConfigurations        storageconfigurations.StorageConfigurationsService
	UsageDownload                usagedownload.BillableUsageDownloadService
	VpcEndpoints                 vpcendpoints.VpcEndpointsService
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
	apiClient, err := client.New(cfg)
	if err != nil {
		panic(err)
	}
	return &AccountsClient{
		Config:                       cfg,
		AccountPermissionAssignments: accountpermissionassignments.NewAccountPermissionAssignments(apiClient),
		Budgets:                      budgets.NewBudgets(apiClient),
		Credentials:                  credentials.NewCredentialConfigurations(apiClient),
		CustomerManagedKeys:          customermanagedkeys.NewKeyConfigurations(apiClient),
		LogDelivery:                  logdelivery.NewLogDelivery(apiClient),
		Networks:                     networks.NewNetworkConfigurations(apiClient),
		PrivateAccessSettings:        privateaccesssettings.NewPrivateAccessSettings(apiClient),
		StorageConfigurations:        storageconfigurations.NewStorageConfigurations(apiClient),
		UsageDownload:                usagedownload.NewBillableUsageDownload(apiClient),
		VpcEndpoints:                 vpcendpoints.NewVpcEndpoints(apiClient),
		Workspaces:                   workspaces.NewWorkspaces(apiClient),
	}
}
