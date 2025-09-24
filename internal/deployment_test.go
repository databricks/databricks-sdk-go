package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccStorage(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: &provisioning.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	require.NoError(t, err)

	defer func() {
		_, err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		require.NoError(t, err)
	}()

	_, err = a.Storage.GetByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	require.NoError(t, err)

	configs, err := a.Storage.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccNetworks(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
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

	_, err = a.Networks.GetByNetworkId(ctx, netw.NetworkId)
	require.NoError(t, err)

	configs, err := a.Networks.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccCredentials(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}
	role, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("sdk-"),
		AwsCredentials: &provisioning.CreateCredentialAwsCredentials{
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

	_, err = a.Credentials.GetByCredentialsId(ctx, role.CredentialsId)
	require.NoError(t, err)

	configs, err := a.Credentials.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccEncryptionKeys(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
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
	if !a.Config.IsAws() {
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

	_, err = a.PrivateAccess.GetByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	require.NoError(t, err)

	_, err = a.PrivateAccess.List(ctx)
	require.NoError(t, err)
}

func TestMwsAccVpcEndpoints(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
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
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("go-sdk-"),
		RootBucketInfo: &provisioning.RootBucketInfo{
			BucketName: GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		require.NoError(t, err)
	})

	// TODO: OpenAPI: Document retry protocol on AWS IAM registration errors
	// See https://github.com/databricks/terraform-provider-databricks/issues/1424
	role, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("go-sdk-"),
		AwsCredentials: &provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := retries.New[struct{}](retries.OnErrors(apierr.ErrResourceConflict)).Wait(ctx, func(ctx context.Context) error {
			_, err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
			return err
		})
		require.NoError(t, err)
	})

	updateRole, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("go-sdk-"),
		AwsCredentials: &provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := retries.New[struct{}](retries.OnErrors(apierr.ErrResourceConflict)).Wait(ctx, func(ctx context.Context) error {
			_, err := a.Credentials.DeleteByCredentialsId(ctx, updateRole.CredentialsId)
			return err
		})
		require.NoError(t, err)
	})

	// TODO: Add DNS reachability utility
	// Do not use CreateAndWait. If the workspaces is created but does not reach running state,
	// the cleanup step won't be executed since the test will fail at the `require.NoError(t, err)` line.
	waiter, err := a.Workspaces.Create(ctx, provisioning.CreateWorkspaceRequest{
		WorkspaceName:          RandomName("go-sdk-"),
		AwsRegion:              GetEnvOrSkipTest(t, "AWS_REGION"),
		CredentialsId:          role.CredentialsId,
		StorageConfigurationId: storage.StorageConfigurationId,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := a.Workspaces.DeleteByWorkspaceId(ctx, waiter.WorkspaceId)
		require.NoError(t, err)
	})
	created, err := waiter.Get()
	require.NoError(t, err)

	// this also takes a while
	_, err = a.Workspaces.UpdateAndWait(ctx, provisioning.UpdateWorkspaceRequest{
		WorkspaceId: created.WorkspaceId,
		CustomerFacingWorkspace: provisioning.Workspace{
			CredentialsId: updateRole.CredentialsId,
		},
	})
	require.NoError(t, err)

	_, err = a.Workspaces.GetByWorkspaceId(ctx, created.WorkspaceId)
	require.NoError(t, err)

	_, err = a.Workspaces.List(ctx)
	require.NoError(t, err)
}
