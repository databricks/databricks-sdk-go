// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/billing"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/oauth2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/provisioning"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settings"
)

// NewAccountClient creates new mocked version of Databricks SDK client for Accounts
// which can be used for testing.
func NewAccountClient(t interface {
	mock.TestingT
	Cleanup(func())
}, cfg *config.Config) *databricks.AccountClient {
	return &databricks.AccountClient{
		Config: cfg,

		AccessControl:           iam.NewMockAccountAccessControlInterface(t),
		BillableUsage:           billing.NewMockBillableUsageInterface(t),
		Budgets:                 billing.NewMockBudgetsInterface(t),
		Credentials:             provisioning.NewMockCredentialsInterface(t),
		CustomAppIntegration:    oauth2.NewMockCustomAppIntegrationInterface(t),
		EncryptionKeys:          provisioning.NewMockEncryptionKeysInterface(t),
		Groups:                  iam.NewMockAccountGroupsInterface(t),
		IpAccessLists:           settings.NewMockAccountIpAccessListsInterface(t),
		LogDelivery:             billing.NewMockLogDeliveryInterface(t),
		MetastoreAssignments:    catalog.NewMockAccountMetastoreAssignmentsInterface(t),
		Metastores:              catalog.NewMockAccountMetastoresInterface(t),
		NetworkConnectivity:     settings.NewMockNetworkConnectivityInterface(t),
		Networks:                provisioning.NewMockNetworksInterface(t),
		OAuthPublishedApps:      oauth2.NewMockOAuthPublishedAppsInterface(t),
		PrivateAccess:           provisioning.NewMockPrivateAccessInterface(t),
		PublishedAppIntegration: oauth2.NewMockPublishedAppIntegrationInterface(t),
		ServicePrincipalSecrets: oauth2.NewMockServicePrincipalSecretsInterface(t),
		ServicePrincipals:       iam.NewMockAccountServicePrincipalsInterface(t),
		Settings:                settings.NewMockAccountSettingsInterface(t),
		Storage:                 provisioning.NewMockStorageInterface(t),
		StorageCredentials:      catalog.NewMockAccountStorageCredentialsInterface(t),
		Users:                   iam.NewMockAccountUsersInterface(t),
		VpcEndpoints:            provisioning.NewMockVpcEndpointsInterface(t),
		WorkspaceAssignment:     iam.NewMockWorkspaceAssignmentInterface(t),
		Workspaces:              provisioning.NewMockWorkspacesInterface(t),
	}
}

func GetMockAccountAccessControlAPI(c *databricks.AccountClient) *iam.MockAccountAccessControlInterface {
	api, ok := c.AccessControl.(*iam.MockAccountAccessControlInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockBillableUsageAPI(c *databricks.AccountClient) *billing.MockBillableUsageInterface {
	api, ok := c.BillableUsage.(*billing.MockBillableUsageInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockBudgetsAPI(c *databricks.AccountClient) *billing.MockBudgetsInterface {
	api, ok := c.Budgets.(*billing.MockBudgetsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockCredentialsAPI(c *databricks.AccountClient) *provisioning.MockCredentialsInterface {
	api, ok := c.Credentials.(*provisioning.MockCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockCustomAppIntegrationAPI(c *databricks.AccountClient) *oauth2.MockCustomAppIntegrationInterface {
	api, ok := c.CustomAppIntegration.(*oauth2.MockCustomAppIntegrationInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockEncryptionKeysAPI(c *databricks.AccountClient) *provisioning.MockEncryptionKeysInterface {
	api, ok := c.EncryptionKeys.(*provisioning.MockEncryptionKeysInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountGroupsAPI(c *databricks.AccountClient) *iam.MockAccountGroupsInterface {
	api, ok := c.Groups.(*iam.MockAccountGroupsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountIpAccessListsAPI(c *databricks.AccountClient) *settings.MockAccountIpAccessListsInterface {
	api, ok := c.IpAccessLists.(*settings.MockAccountIpAccessListsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockLogDeliveryAPI(c *databricks.AccountClient) *billing.MockLogDeliveryInterface {
	api, ok := c.LogDelivery.(*billing.MockLogDeliveryInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountMetastoreAssignmentsAPI(c *databricks.AccountClient) *catalog.MockAccountMetastoreAssignmentsInterface {
	api, ok := c.MetastoreAssignments.(*catalog.MockAccountMetastoreAssignmentsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountMetastoresAPI(c *databricks.AccountClient) *catalog.MockAccountMetastoresInterface {
	api, ok := c.Metastores.(*catalog.MockAccountMetastoresInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockNetworkConnectivityAPI(c *databricks.AccountClient) *settings.MockNetworkConnectivityInterface {
	api, ok := c.NetworkConnectivity.(*settings.MockNetworkConnectivityInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockNetworksAPI(c *databricks.AccountClient) *provisioning.MockNetworksInterface {
	api, ok := c.Networks.(*provisioning.MockNetworksInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockOAuthPublishedAppsAPI(c *databricks.AccountClient) *oauth2.MockOAuthPublishedAppsInterface {
	api, ok := c.OAuthPublishedApps.(*oauth2.MockOAuthPublishedAppsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockPrivateAccessAPI(c *databricks.AccountClient) *provisioning.MockPrivateAccessInterface {
	api, ok := c.PrivateAccess.(*provisioning.MockPrivateAccessInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockPublishedAppIntegrationAPI(c *databricks.AccountClient) *oauth2.MockPublishedAppIntegrationInterface {
	api, ok := c.PublishedAppIntegration.(*oauth2.MockPublishedAppIntegrationInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockServicePrincipalSecretsAPI(c *databricks.AccountClient) *oauth2.MockServicePrincipalSecretsInterface {
	api, ok := c.ServicePrincipalSecrets.(*oauth2.MockServicePrincipalSecretsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountServicePrincipalsAPI(c *databricks.AccountClient) *iam.MockAccountServicePrincipalsInterface {
	api, ok := c.ServicePrincipals.(*iam.MockAccountServicePrincipalsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountSettingsAPI(c *databricks.AccountClient) *settings.MockAccountSettingsInterface {
	api, ok := c.Settings.(*settings.MockAccountSettingsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockStorageAPI(c *databricks.AccountClient) *provisioning.MockStorageInterface {
	api, ok := c.Storage.(*provisioning.MockStorageInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountStorageCredentialsAPI(c *databricks.AccountClient) *catalog.MockAccountStorageCredentialsInterface {
	api, ok := c.StorageCredentials.(*catalog.MockAccountStorageCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockAccountUsersAPI(c *databricks.AccountClient) *iam.MockAccountUsersInterface {
	api, ok := c.Users.(*iam.MockAccountUsersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockVpcEndpointsAPI(c *databricks.AccountClient) *provisioning.MockVpcEndpointsInterface {
	api, ok := c.VpcEndpoints.(*provisioning.MockVpcEndpointsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockWorkspaceAssignmentAPI(c *databricks.AccountClient) *iam.MockWorkspaceAssignmentInterface {
	api, ok := c.WorkspaceAssignment.(*iam.MockWorkspaceAssignmentInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func GetMockWorkspacesAPI(c *databricks.AccountClient) *provisioning.MockWorkspacesInterface {
	api, ok := c.Workspaces.(*provisioning.MockWorkspacesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
