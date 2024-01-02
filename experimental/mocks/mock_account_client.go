// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"github.com/databricks/databricks-sdk-go"
	"github.com/stretchr/testify/mock"

	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/billing"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/catalog"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/iam"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/oauth2"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/provisioning"
	"github.com/databricks/databricks-sdk-go/experimental/mocks/service/settings"
)

type MockAccountClient struct {
	AccountClient *databricks.AccountClient
}

// NewMockAccountClient creates new mocked version of Databricks SDK client for Accounts
// which can be used for testing.
func NewMockAccountClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAccountClient {
	return &MockAccountClient{
		AccountClient: &databricks.AccountClient{
			Config: nil,

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
		},
	}
}

func (m *MockAccountClient) GetMockAccountAccessControlAPI() *iam.MockAccountAccessControlInterface {
	api, ok := m.AccountClient.AccessControl.(*iam.MockAccountAccessControlInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockBillableUsageAPI() *billing.MockBillableUsageInterface {
	api, ok := m.AccountClient.BillableUsage.(*billing.MockBillableUsageInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockBudgetsAPI() *billing.MockBudgetsInterface {
	api, ok := m.AccountClient.Budgets.(*billing.MockBudgetsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockCredentialsAPI() *provisioning.MockCredentialsInterface {
	api, ok := m.AccountClient.Credentials.(*provisioning.MockCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockCustomAppIntegrationAPI() *oauth2.MockCustomAppIntegrationInterface {
	api, ok := m.AccountClient.CustomAppIntegration.(*oauth2.MockCustomAppIntegrationInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockEncryptionKeysAPI() *provisioning.MockEncryptionKeysInterface {
	api, ok := m.AccountClient.EncryptionKeys.(*provisioning.MockEncryptionKeysInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountGroupsAPI() *iam.MockAccountGroupsInterface {
	api, ok := m.AccountClient.Groups.(*iam.MockAccountGroupsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountIpAccessListsAPI() *settings.MockAccountIpAccessListsInterface {
	api, ok := m.AccountClient.IpAccessLists.(*settings.MockAccountIpAccessListsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockLogDeliveryAPI() *billing.MockLogDeliveryInterface {
	api, ok := m.AccountClient.LogDelivery.(*billing.MockLogDeliveryInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountMetastoreAssignmentsAPI() *catalog.MockAccountMetastoreAssignmentsInterface {
	api, ok := m.AccountClient.MetastoreAssignments.(*catalog.MockAccountMetastoreAssignmentsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountMetastoresAPI() *catalog.MockAccountMetastoresInterface {
	api, ok := m.AccountClient.Metastores.(*catalog.MockAccountMetastoresInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockNetworkConnectivityAPI() *settings.MockNetworkConnectivityInterface {
	api, ok := m.AccountClient.NetworkConnectivity.(*settings.MockNetworkConnectivityInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockNetworksAPI() *provisioning.MockNetworksInterface {
	api, ok := m.AccountClient.Networks.(*provisioning.MockNetworksInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockOAuthPublishedAppsAPI() *oauth2.MockOAuthPublishedAppsInterface {
	api, ok := m.AccountClient.OAuthPublishedApps.(*oauth2.MockOAuthPublishedAppsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockPrivateAccessAPI() *provisioning.MockPrivateAccessInterface {
	api, ok := m.AccountClient.PrivateAccess.(*provisioning.MockPrivateAccessInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockPublishedAppIntegrationAPI() *oauth2.MockPublishedAppIntegrationInterface {
	api, ok := m.AccountClient.PublishedAppIntegration.(*oauth2.MockPublishedAppIntegrationInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockServicePrincipalSecretsAPI() *oauth2.MockServicePrincipalSecretsInterface {
	api, ok := m.AccountClient.ServicePrincipalSecrets.(*oauth2.MockServicePrincipalSecretsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountServicePrincipalsAPI() *iam.MockAccountServicePrincipalsInterface {
	api, ok := m.AccountClient.ServicePrincipals.(*iam.MockAccountServicePrincipalsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountSettingsAPI() *settings.MockAccountSettingsInterface {
	api, ok := m.AccountClient.Settings.(*settings.MockAccountSettingsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockStorageAPI() *provisioning.MockStorageInterface {
	api, ok := m.AccountClient.Storage.(*provisioning.MockStorageInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountStorageCredentialsAPI() *catalog.MockAccountStorageCredentialsInterface {
	api, ok := m.AccountClient.StorageCredentials.(*catalog.MockAccountStorageCredentialsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockAccountUsersAPI() *iam.MockAccountUsersInterface {
	api, ok := m.AccountClient.Users.(*iam.MockAccountUsersInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockVpcEndpointsAPI() *provisioning.MockVpcEndpointsInterface {
	api, ok := m.AccountClient.VpcEndpoints.(*provisioning.MockVpcEndpointsInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockWorkspaceAssignmentAPI() *iam.MockWorkspaceAssignmentInterface {
	api, ok := m.AccountClient.WorkspaceAssignment.(*iam.MockWorkspaceAssignmentInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}

func (m *MockAccountClient) GetMockWorkspacesAPI() *provisioning.MockWorkspacesInterface {
	api, ok := m.AccountClient.Workspaces.(*provisioning.MockWorkspacesInterface)
	if !ok {
		panic("Invalid API interface type")
	}
	return api
}
