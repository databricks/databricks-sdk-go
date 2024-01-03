// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package mocks

import (
	"fmt"

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
		panic(fmt.Sprintf("expected AccessControl to be *iam.MockAccountAccessControlInterface, actual was %T", m.AccountClient.AccessControl))
	}
	return api
}

func (m *MockAccountClient) GetMockBillableUsageAPI() *billing.MockBillableUsageInterface {
	api, ok := m.AccountClient.BillableUsage.(*billing.MockBillableUsageInterface)
	if !ok {
		panic(fmt.Sprintf("expected BillableUsage to be *billing.MockBillableUsageInterface, actual was %T", m.AccountClient.BillableUsage))
	}
	return api
}

func (m *MockAccountClient) GetMockBudgetsAPI() *billing.MockBudgetsInterface {
	api, ok := m.AccountClient.Budgets.(*billing.MockBudgetsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Budgets to be *billing.MockBudgetsInterface, actual was %T", m.AccountClient.Budgets))
	}
	return api
}

func (m *MockAccountClient) GetMockCredentialsAPI() *provisioning.MockCredentialsInterface {
	api, ok := m.AccountClient.Credentials.(*provisioning.MockCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Credentials to be *provisioning.MockCredentialsInterface, actual was %T", m.AccountClient.Credentials))
	}
	return api
}

func (m *MockAccountClient) GetMockCustomAppIntegrationAPI() *oauth2.MockCustomAppIntegrationInterface {
	api, ok := m.AccountClient.CustomAppIntegration.(*oauth2.MockCustomAppIntegrationInterface)
	if !ok {
		panic(fmt.Sprintf("expected CustomAppIntegration to be *oauth2.MockCustomAppIntegrationInterface, actual was %T", m.AccountClient.CustomAppIntegration))
	}
	return api
}

func (m *MockAccountClient) GetMockEncryptionKeysAPI() *provisioning.MockEncryptionKeysInterface {
	api, ok := m.AccountClient.EncryptionKeys.(*provisioning.MockEncryptionKeysInterface)
	if !ok {
		panic(fmt.Sprintf("expected EncryptionKeys to be *provisioning.MockEncryptionKeysInterface, actual was %T", m.AccountClient.EncryptionKeys))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountGroupsAPI() *iam.MockAccountGroupsInterface {
	api, ok := m.AccountClient.Groups.(*iam.MockAccountGroupsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Groups to be *iam.MockAccountGroupsInterface, actual was %T", m.AccountClient.Groups))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountIpAccessListsAPI() *settings.MockAccountIpAccessListsInterface {
	api, ok := m.AccountClient.IpAccessLists.(*settings.MockAccountIpAccessListsInterface)
	if !ok {
		panic(fmt.Sprintf("expected IpAccessLists to be *settings.MockAccountIpAccessListsInterface, actual was %T", m.AccountClient.IpAccessLists))
	}
	return api
}

func (m *MockAccountClient) GetMockLogDeliveryAPI() *billing.MockLogDeliveryInterface {
	api, ok := m.AccountClient.LogDelivery.(*billing.MockLogDeliveryInterface)
	if !ok {
		panic(fmt.Sprintf("expected LogDelivery to be *billing.MockLogDeliveryInterface, actual was %T", m.AccountClient.LogDelivery))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountMetastoreAssignmentsAPI() *catalog.MockAccountMetastoreAssignmentsInterface {
	api, ok := m.AccountClient.MetastoreAssignments.(*catalog.MockAccountMetastoreAssignmentsInterface)
	if !ok {
		panic(fmt.Sprintf("expected MetastoreAssignments to be *catalog.MockAccountMetastoreAssignmentsInterface, actual was %T", m.AccountClient.MetastoreAssignments))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountMetastoresAPI() *catalog.MockAccountMetastoresInterface {
	api, ok := m.AccountClient.Metastores.(*catalog.MockAccountMetastoresInterface)
	if !ok {
		panic(fmt.Sprintf("expected Metastores to be *catalog.MockAccountMetastoresInterface, actual was %T", m.AccountClient.Metastores))
	}
	return api
}

func (m *MockAccountClient) GetMockNetworkConnectivityAPI() *settings.MockNetworkConnectivityInterface {
	api, ok := m.AccountClient.NetworkConnectivity.(*settings.MockNetworkConnectivityInterface)
	if !ok {
		panic(fmt.Sprintf("expected NetworkConnectivity to be *settings.MockNetworkConnectivityInterface, actual was %T", m.AccountClient.NetworkConnectivity))
	}
	return api
}

func (m *MockAccountClient) GetMockNetworksAPI() *provisioning.MockNetworksInterface {
	api, ok := m.AccountClient.Networks.(*provisioning.MockNetworksInterface)
	if !ok {
		panic(fmt.Sprintf("expected Networks to be *provisioning.MockNetworksInterface, actual was %T", m.AccountClient.Networks))
	}
	return api
}

func (m *MockAccountClient) GetMockOAuthPublishedAppsAPI() *oauth2.MockOAuthPublishedAppsInterface {
	api, ok := m.AccountClient.OAuthPublishedApps.(*oauth2.MockOAuthPublishedAppsInterface)
	if !ok {
		panic(fmt.Sprintf("expected OAuthPublishedApps to be *oauth2.MockOAuthPublishedAppsInterface, actual was %T", m.AccountClient.OAuthPublishedApps))
	}
	return api
}

func (m *MockAccountClient) GetMockPrivateAccessAPI() *provisioning.MockPrivateAccessInterface {
	api, ok := m.AccountClient.PrivateAccess.(*provisioning.MockPrivateAccessInterface)
	if !ok {
		panic(fmt.Sprintf("expected PrivateAccess to be *provisioning.MockPrivateAccessInterface, actual was %T", m.AccountClient.PrivateAccess))
	}
	return api
}

func (m *MockAccountClient) GetMockPublishedAppIntegrationAPI() *oauth2.MockPublishedAppIntegrationInterface {
	api, ok := m.AccountClient.PublishedAppIntegration.(*oauth2.MockPublishedAppIntegrationInterface)
	if !ok {
		panic(fmt.Sprintf("expected PublishedAppIntegration to be *oauth2.MockPublishedAppIntegrationInterface, actual was %T", m.AccountClient.PublishedAppIntegration))
	}
	return api
}

func (m *MockAccountClient) GetMockServicePrincipalSecretsAPI() *oauth2.MockServicePrincipalSecretsInterface {
	api, ok := m.AccountClient.ServicePrincipalSecrets.(*oauth2.MockServicePrincipalSecretsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipalSecrets to be *oauth2.MockServicePrincipalSecretsInterface, actual was %T", m.AccountClient.ServicePrincipalSecrets))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountServicePrincipalsAPI() *iam.MockAccountServicePrincipalsInterface {
	api, ok := m.AccountClient.ServicePrincipals.(*iam.MockAccountServicePrincipalsInterface)
	if !ok {
		panic(fmt.Sprintf("expected ServicePrincipals to be *iam.MockAccountServicePrincipalsInterface, actual was %T", m.AccountClient.ServicePrincipals))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountSettingsAPI() *settings.MockAccountSettingsInterface {
	api, ok := m.AccountClient.Settings.(*settings.MockAccountSettingsInterface)
	if !ok {
		panic(fmt.Sprintf("expected Settings to be *settings.MockAccountSettingsInterface, actual was %T", m.AccountClient.Settings))
	}
	return api
}

func (m *MockAccountClient) GetMockStorageAPI() *provisioning.MockStorageInterface {
	api, ok := m.AccountClient.Storage.(*provisioning.MockStorageInterface)
	if !ok {
		panic(fmt.Sprintf("expected Storage to be *provisioning.MockStorageInterface, actual was %T", m.AccountClient.Storage))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountStorageCredentialsAPI() *catalog.MockAccountStorageCredentialsInterface {
	api, ok := m.AccountClient.StorageCredentials.(*catalog.MockAccountStorageCredentialsInterface)
	if !ok {
		panic(fmt.Sprintf("expected StorageCredentials to be *catalog.MockAccountStorageCredentialsInterface, actual was %T", m.AccountClient.StorageCredentials))
	}
	return api
}

func (m *MockAccountClient) GetMockAccountUsersAPI() *iam.MockAccountUsersInterface {
	api, ok := m.AccountClient.Users.(*iam.MockAccountUsersInterface)
	if !ok {
		panic(fmt.Sprintf("expected Users to be *iam.MockAccountUsersInterface, actual was %T", m.AccountClient.Users))
	}
	return api
}

func (m *MockAccountClient) GetMockVpcEndpointsAPI() *provisioning.MockVpcEndpointsInterface {
	api, ok := m.AccountClient.VpcEndpoints.(*provisioning.MockVpcEndpointsInterface)
	if !ok {
		panic(fmt.Sprintf("expected VpcEndpoints to be *provisioning.MockVpcEndpointsInterface, actual was %T", m.AccountClient.VpcEndpoints))
	}
	return api
}

func (m *MockAccountClient) GetMockWorkspaceAssignmentAPI() *iam.MockWorkspaceAssignmentInterface {
	api, ok := m.AccountClient.WorkspaceAssignment.(*iam.MockWorkspaceAssignmentInterface)
	if !ok {
		panic(fmt.Sprintf("expected WorkspaceAssignment to be *iam.MockWorkspaceAssignmentInterface, actual was %T", m.AccountClient.WorkspaceAssignment))
	}
	return api
}

func (m *MockAccountClient) GetMockWorkspacesAPI() *provisioning.MockWorkspacesInterface {
	api, ok := m.AccountClient.Workspaces.(*provisioning.MockWorkspacesInterface)
	if !ok {
		panic(fmt.Sprintf("expected Workspaces to be *provisioning.MockWorkspacesInterface, actual was %T", m.AccountClient.Workspaces))
	}
	return api
}
