package accounts

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/accountpermissionassignments"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/deployment"
)

type AccountsClient struct {
	Config                       *databricks.Config
	AccountPermissionAssignments accountpermissionassignments.AccountPermissionAssignmentsService
	Budgets                      billing.BudgetsService
	Credentials                  deployment.CredentialConfigurationsService
	CustomerManagedKeys          deployment.KeyConfigurationsService
	LogDelivery                  billing.LogDeliveryService
	Networks                     deployment.NetworkConfigurationsService
	PrivateAccessSettings        deployment.PrivateAccessSettingsService
	StorageConfigurations        deployment.StorageConfigurationsService
	UsageDownload                billing.BillableUsageDownloadService
	VpcEndpoints                 deployment.VpcEndpointsService
	Workspaces                   deployment.WorkspacesService
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
		Budgets:                      billing.NewBudgets(apiClient),
		Credentials:                  deployment.NewCredentialConfigurations(apiClient),
		CustomerManagedKeys:          deployment.NewKeyConfigurations(apiClient),
		LogDelivery:                  billing.NewLogDelivery(apiClient),
		Networks:                     deployment.NewNetworkConfigurations(apiClient),
		PrivateAccessSettings:        deployment.NewPrivateAccessSettings(apiClient),
		StorageConfigurations:        deployment.NewStorageConfigurations(apiClient),
		UsageDownload:                billing.NewBillableUsageDownload(apiClient),
		VpcEndpoints:                 deployment.NewVpcEndpoints(apiClient),
		Workspaces:                   deployment.NewWorkspaces(apiClient),
	}
}
