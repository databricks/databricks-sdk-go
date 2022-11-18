package accounts

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
	"github.com/databricks/databricks-sdk-go/service/billing"
	"github.com/databricks/databricks-sdk-go/service/deployment"
	"github.com/databricks/databricks-sdk-go/service/permissions"
	"github.com/databricks/databricks-sdk-go/service/scim"
)

type AccountsClient struct {
	Config                *databricks.Config
	Budgets               billing.BudgetsService
	Credentials           deployment.CredentialConfigurationsService
	Groups                scim.AccountGroupsService
	CustomerManagedKeys   deployment.KeyConfigurationsService
	LogDelivery           billing.LogDeliveryService
	Networks              deployment.NetworkConfigurationsService
	PrivateAccessSettings deployment.PrivateAccessSettingsService
	ServicePrincipals     scim.AccountServicePrincipalsService
	StorageConfigurations deployment.StorageConfigurationsService
	UsageDownload         billing.BillableUsageDownloadService
	Users                 scim.AccountUsersService
	VpcEndpoints          deployment.VpcEndpointsService
	WorkspaceAssignment   permissions.WorkspaceAssignmentService
	Workspaces            deployment.WorkspacesService
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
	err := cfg.EnsureResolved()
	if err != nil {
		panic(err)
	}
	if cfg.AccountID == "" {
		panic("AccountID is not specified on config")
	}
	apiClient, err := client.New(cfg)
	if err != nil {
		panic(err)
	}
	return &AccountsClient{
		Config:                cfg,
		Budgets:               billing.NewBudgets(apiClient),
		Credentials:           deployment.NewCredentialConfigurations(apiClient),
		CustomerManagedKeys:   deployment.NewKeyConfigurations(apiClient),
		Groups:                scim.NewAccountGroups(apiClient),
		LogDelivery:           billing.NewLogDelivery(apiClient),
		Networks:              deployment.NewNetworkConfigurations(apiClient),
		PrivateAccessSettings: deployment.NewPrivateAccessSettings(apiClient),
		ServicePrincipals:     scim.NewAccountServicePrincipals(apiClient),
		StorageConfigurations: deployment.NewStorageConfigurations(apiClient),
		UsageDownload:         billing.NewBillableUsageDownload(apiClient),
		Users:                 scim.NewAccountUsers(apiClient),
		VpcEndpoints:          deployment.NewVpcEndpoints(apiClient),
		WorkspaceAssignment:   permissions.NewWorkspaceAssignment(apiClient),
		Workspaces:            deployment.NewWorkspaces(apiClient),
	}
}
