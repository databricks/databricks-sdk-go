package internal

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/common/environment"
	"github.com/databricks/databricks-sdk-go/qa"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccStorage(t *testing.T) {
	ctx, a := accountTest(t)
	if !IsCloud(environment.CloudAWS) {
		t.SkipNow()
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	require.NoError(t, err)

	defer func() {
		_, err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		require.NoError(t, err)
	}()

	byId, err := a.Storage.GetByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	require.NoError(t, err)

	byName, err := a.Storage.GetByStorageConfigurationName(ctx, byId.StorageConfigurationName)
	require.NoError(t, err)
	assert.Equal(t, byId.StorageConfigurationId, byName.StorageConfigurationId)

	configs, err := a.Storage.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccNetworks(t *testing.T) {
	ctx, a := accountTest(t)
	if !IsCloud(environment.CloudAWS) {
		t.SkipNow()
	}
	netw, err := a.Networks.Create(ctx, provisioning.CreateNetworkRequest{
		NetworkName:      RandomName("sdk-"),
		VpcId:            RandomHex("vpc-", 17),
		SubnetIds:        []string{RandomHex("subnet-", 17), RandomHex("subnet-", 17)},
		SecurityGroupIds: []string{RandomHex("sg-", 17)},
	})
	require.NoError(t, err)
	defer func() {
		_, err = a.Networks.DeleteByNetworkId(ctx, netw.NetworkId)
		require.NoError(t, err)
	}()

	byId, err := a.Networks.GetByNetworkId(ctx, netw.NetworkId)
	require.NoError(t, err)

	byName, err := a.Networks.GetByNetworkName(ctx, byId.NetworkName)
	require.NoError(t, err)
	assert.Equal(t, byId.NetworkId, byName.NetworkId)

	configs, err := a.Networks.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccCredentials(t *testing.T) {
	ctx, a := accountTest(t)
	if !IsCloud(environment.CloudAWS) {
		t.SkipNow()
	}
	role, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("sdk-"),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
		require.NoError(t, err)
	})

	byId, err := a.Credentials.GetByCredentialsId(ctx, role.CredentialsId)
	require.NoError(t, err)

	byName, err := a.Credentials.GetByCredentialsName(ctx, byId.CredentialsName)
	require.NoError(t, err)
	assert.Equal(t, byId.CredentialsId, byName.CredentialsId)

	configs, err := a.Credentials.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccEncryptionKeys(t *testing.T) {
	ctx, a := accountTest(t)
	if !IsCloud(environment.CloudAWS) {
		t.SkipNow()
	}

	created, err := a.EncryptionKeys.Create(ctx, provisioning.CreateCustomerManagedKeyRequest{
		AwsKeyInfo: &provisioning.CreateAwsKeyInfo{
			KeyArn:   GetEnvOrSkipTest(t, "TEST_MANAGED_KMS_KEY_ARN"),
			KeyAlias: GetEnvOrSkipTest(t, "TEST_STORAGE_KMS_KEY_ALIAS"),
		},
		UseCases: []provisioning.KeyUseCase{provisioning.KeyUseCaseManagedServices},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := a.EncryptionKeys.DeleteByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
		require.NoError(t, err)
	})

	byId, err := a.EncryptionKeys.GetByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
	require.NoError(t, err)
	assert.Equal(t, provisioning.KeyUseCaseManagedServices, byId.UseCases[0])

	all, err := a.EncryptionKeys.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestMwsAccPrivateAccess(t *testing.T) {
	ctx, a := accountTest(t)
	if !IsCloud(environment.CloudAWS) {
		t.SkipNow()
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.CreatePrivateAccessSettingsRequest{
		PrivateAccessSettingsName: RandomName("go-sdk-"),
		Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
		require.NoError(t, err)
	})
	_, err = a.PrivateAccess.Replace(ctx, provisioning.ReplacePrivateAccessSettingsRequest{
		PrivateAccessSettingsId: created.PrivateAccessSettingsId,
		CustomerFacingPrivateAccessSettings: provisioning.PrivateAccessSettings{
			PrivateAccessSettingsName: RandomName("go-sdk-"),
			Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
		},
	})
	require.NoError(t, err)

	byId, err := a.PrivateAccess.GetByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	require.NoError(t, err)

	byName, err := a.PrivateAccess.GetByPrivateAccessSettingsName(ctx, byId.PrivateAccessSettingsName)
	require.NoError(t, err)
	assert.Equal(t, byId.PrivateAccessSettingsId, byName.PrivateAccessSettingsId)

	all, err := a.PrivateAccess.List(ctx)
	require.NoError(t, err)

	names, err := a.PrivateAccess.PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.PrivateAccessSettingsId, names[byId.PrivateAccessSettingsName])
}

func TestMwsAccVpcEndpoints(t *testing.T) {
	ctx, a := accountTest(t)
	if !IsCloud(environment.CloudAWS) {
		t.SkipNow()
	}

	created, err := a.VpcEndpoints.Create(ctx, provisioning.CreateVpcEndpointRequest{
		AwsVpcEndpointId: GetEnvOrSkipTest(t, "TEST_RELAY_VPC_ENDPOINT"),
		Region:           GetEnvOrSkipTest(t, "AWS_REGION"),
		VpcEndpointName:  RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := a.VpcEndpoints.DeleteByVpcEndpointId(ctx, created.VpcEndpointId)
		require.NoError(t, err)
	})

	byId, err := a.VpcEndpoints.GetByVpcEndpointId(ctx, created.VpcEndpointId)
	require.NoError(t, err)
	assert.Equal(t, provisioning.EndpointUseCaseDataplaneRelayAccess, byId.UseCase)

	all, err := a.VpcEndpoints.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestMwsAccWorkspaces(t *testing.T) {
	const accountID = "test-account"
	createRequest := provisioning.CreateWorkspaceRequest{
		WorkspaceName:          "test-workspace",
		AwsRegion:              "us-east-1",
		CredentialsId:          "test-credential-original",
		StorageConfigurationId: "test-storage",
	}
	wantCreated := provisioning.Workspace{
		WorkspaceId:            123,
		WorkspaceName:          createRequest.WorkspaceName,
		CredentialsId:          createRequest.CredentialsId,
		StorageConfigurationId: createRequest.StorageConfigurationId,
		WorkspaceStatus:        provisioning.WorkspaceStatusRunning,
	}
	wantUpdated := wantCreated
	wantUpdated.CredentialsId = "test-credential-updated"

	accountPath := "/api/2.0/accounts/" + accountID
	workspacePath := fmt.Sprintf("%s/workspaces/%d", accountPath, wantCreated.WorkspaceId)

	httpFixtures := qa.HTTPFixtures{
		{
			Method:          http.MethodPost,
			Resource:        accountPath + "/workspaces",
			ExpectedRequest: createRequest,
			Response: provisioning.Workspace{
				WorkspaceId: wantCreated.WorkspaceId,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: workspacePath + "?",
			Response: wantCreated,
		},
		{
			Method:   http.MethodPatch,
			Resource: workspacePath,
			ExpectedRequest: provisioning.Workspace{
				CredentialsId: wantUpdated.CredentialsId,
			},
			Response: provisioning.Workspace{
				WorkspaceId: wantUpdated.WorkspaceId,
			},
		},
		{
			Method:   http.MethodGet,
			Resource: workspacePath + "?",
			Response: wantUpdated,
		},
		{
			Method:   http.MethodDelete,
			Resource: workspacePath + "?",
			Response: wantUpdated,
		},
	}

	httpFixtures.ApplyClient(t, func(ctx context.Context, apiClient *client.DatabricksClient) {
		apiClient.Config.AccountID = accountID
		workspaces := provisioning.NewWorkspaces(apiClient)
		workspaceComparison := cmpopts.IgnoreFields(provisioning.Workspace{}, "ForceSendFields")

		waiter, err := workspaces.Create(ctx, createRequest)
		if err != nil {
			t.Fatalf("create workspace: %v", err)
		}

		gotCreated, err := waiter.Get()
		if err != nil {
			t.Fatalf("wait for created workspace: %v", err)
		}
		if diff := cmp.Diff(&wantCreated, gotCreated, workspaceComparison); diff != "" {
			t.Errorf("created workspace mismatch (-want +got):\n%s", diff)
		}

		updateWaiter, err := workspaces.Update(ctx, provisioning.UpdateWorkspaceRequest{
			WorkspaceId: wantUpdated.WorkspaceId,
			CustomerFacingWorkspace: provisioning.Workspace{
				CredentialsId: wantUpdated.CredentialsId,
			},
		})
		if err != nil {
			t.Fatalf("update workspace: %v", err)
		}

		gotUpdated, err := updateWaiter.Get()
		if err != nil {
			t.Fatalf("wait for updated workspace: %v", err)
		}
		if diff := cmp.Diff(&wantUpdated, gotUpdated, workspaceComparison); diff != "" {
			t.Errorf("updated workspace mismatch (-want +got):\n%s", diff)
		}

		gotDeleted, err := workspaces.DeleteByWorkspaceId(ctx, wantUpdated.WorkspaceId)
		if err != nil {
			t.Fatalf("delete workspace: %v", err)
		}
		if diff := cmp.Diff(&wantUpdated, gotDeleted, workspaceComparison); diff != "" {
			t.Errorf("deleted workspace mismatch (-want +got):\n%s", diff)
		}
	})
}
