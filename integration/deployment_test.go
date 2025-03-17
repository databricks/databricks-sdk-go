package integration

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/provisioning/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccStorage(t *testing.T) {
	ctx, cfg := accountTest(t)

	StorageAPI, err := provisioning.NewStorageClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.SkipNow()
	}

	storage, err := StorageAPI.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	require.NoError(t, err)

	defer func() {
		err = StorageAPI.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		require.NoError(t, err)
	}()

	byId, err := StorageAPI.GetByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	require.NoError(t, err)

	byName, err := StorageAPI.GetByStorageConfigurationName(ctx, byId.StorageConfigurationName)
	require.NoError(t, err)
	assert.Equal(t, byId.StorageConfigurationId, byName.StorageConfigurationId)

	configs, err := StorageAPI.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccNetworks(t *testing.T) {
	ctx, cfg := accountTest(t)

	NetworksAPI, err := provisioning.NewNetworksClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.SkipNow()
	}
	netw, err := NetworksAPI.Create(ctx, provisioning.CreateNetworkRequest{
		NetworkName:      RandomName("sdk-"),
		VpcId:            RandomHex("vpc-", 17),
		SubnetIds:        []string{RandomHex("subnet-", 17), RandomHex("subnet-", 17)},
		SecurityGroupIds: []string{RandomHex("sg-", 17)},
	})
	require.NoError(t, err)
	defer func() {
		err = NetworksAPI.DeleteByNetworkId(ctx, netw.NetworkId)
		require.NoError(t, err)
	}()

	byId, err := NetworksAPI.GetByNetworkId(ctx, netw.NetworkId)
	require.NoError(t, err)

	byName, err := NetworksAPI.GetByNetworkName(ctx, byId.NetworkName)
	require.NoError(t, err)
	assert.Equal(t, byId.NetworkId, byName.NetworkId)

	configs, err := NetworksAPI.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccCredentials(t *testing.T) {
	ctx, cfg := accountTest(t)
	CredentialsAPI, err := provisioning.NewCredentialsClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.SkipNow()
	}
	role, err := CredentialsAPI.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("sdk-"),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = CredentialsAPI.DeleteByCredentialsId(ctx, role.CredentialsId)
		require.NoError(t, err)
	})

	byId, err := CredentialsAPI.GetByCredentialsId(ctx, role.CredentialsId)
	require.NoError(t, err)

	byName, err := CredentialsAPI.GetByCredentialsName(ctx, byId.CredentialsName)
	require.NoError(t, err)
	assert.Equal(t, byId.CredentialsId, byName.CredentialsId)

	configs, err := CredentialsAPI.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccEncryptionKeys(t *testing.T) {
	ctx, cfg := accountTest(t)
	EncryptionKeysAPI, err := provisioning.NewEncryptionKeysClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.SkipNow()
	}

	created, err := EncryptionKeysAPI.Create(ctx, provisioning.CreateCustomerManagedKeyRequest{
		AwsKeyInfo: &provisioning.CreateAwsKeyInfo{
			KeyArn:   GetEnvOrSkipTest(t, "TEST_MANAGED_KMS_KEY_ARN"),
			KeyAlias: GetEnvOrSkipTest(t, "TEST_STORAGE_KMS_KEY_ALIAS"),
		},
		UseCases: []provisioning.KeyUseCase{provisioning.KeyUseCaseManagedServices},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := EncryptionKeysAPI.DeleteByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
		require.NoError(t, err)
	})

	byId, err := EncryptionKeysAPI.GetByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
	require.NoError(t, err)
	assert.Equal(t, provisioning.KeyUseCaseManagedServices, byId.UseCases[0])

	all, err := EncryptionKeysAPI.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestMwsAccPrivateAccess(t *testing.T) {
	ctx, cfg := accountTest(t)
	PrivateAccessAPI, err := provisioning.NewPrivateAccessClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.SkipNow()
	}

	created, err := PrivateAccessAPI.Create(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsName: RandomName("go-sdk-"),
		Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := PrivateAccessAPI.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
		require.NoError(t, err)
	})
	err = PrivateAccessAPI.Replace(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsId:   created.PrivateAccessSettingsId,
		PrivateAccessSettingsName: RandomName("go-sdk-"),
		Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
	})
	require.NoError(t, err)

	byId, err := PrivateAccessAPI.GetByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	require.NoError(t, err)

	byName, err := PrivateAccessAPI.GetByPrivateAccessSettingsName(ctx, byId.PrivateAccessSettingsName)
	require.NoError(t, err)
	assert.Equal(t, byId.PrivateAccessSettingsId, byName.PrivateAccessSettingsId)

	all, err := PrivateAccessAPI.List(ctx)
	require.NoError(t, err)

	names, err := PrivateAccessAPI.PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.PrivateAccessSettingsId, names[byId.PrivateAccessSettingsName])
}

func TestMwsAccVpcEndpoints(t *testing.T) {
	ctx, cfg := accountTest(t)
	VpcEndpointsAPI, err := provisioning.NewVpcEndpointsClient(cfg)
	require.NoError(t, err)
	if !cfg.IsAws() {
		t.SkipNow()
	}

	created, err := VpcEndpointsAPI.Create(ctx, provisioning.CreateVpcEndpointRequest{
		AwsVpcEndpointId: GetEnvOrSkipTest(t, "TEST_RELAY_VPC_ENDPOINT"),
		Region:           GetEnvOrSkipTest(t, "AWS_REGION"),
		VpcEndpointName:  RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := VpcEndpointsAPI.DeleteByVpcEndpointId(ctx, created.VpcEndpointId)
		require.NoError(t, err)
	})

	byId, err := VpcEndpointsAPI.GetByVpcEndpointId(ctx, created.VpcEndpointId)
	require.NoError(t, err)
	assert.Equal(t, provisioning.EndpointUseCaseDataplaneRelayAccess, byId.UseCase)

	all, err := VpcEndpointsAPI.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

// TODO: Enable this test when LRO is implemented
//
// func TestMwsAccWorkspaces(t *testing.T) {
// 	ctx, cfg := accountTest(t)
// 	StorageAPI, err := provisioning.NewStorageClient(cfg)
// 	require.NoError(t, err)
// 	if !StorageAPI.Config.IsAws() {
// 		t.SkipNow()
// 	}

// 	storage, err := StorageAPI.Create(ctx, provisioning.CreateStorageConfigurationRequest{
// 		StorageConfigurationName: RandomName("go-sdk-"),
// 		RootBucketInfo: provisioning.RootBucketInfo{
// 			BucketName: GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET"),
// 		},
// 	})
// 	require.NoError(t, err)
// 	t.Cleanup(func() {
// 		err := StorageAPI.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
// 		require.NoError(t, err)
// 	})

// 	// TODO: OpenAPI: Document retry protocol on AWS IAM registration errors
// 	// See https://github.com/databricks/terraform-provider-databricks/issues/1424
// 	CredentialsAPI, err := provisioning.NewCredentialsClient(cfg)
// 	require.NoError(t, err)
// 	role, err := CredentialsAPI.Create(ctx, provisioning.CreateCredentialRequest{
// 		CredentialsName: RandomName("go-sdk-"),
// 		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
// 			StsRole: &provisioning.CreateCredentialStsRole{
// 				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
// 			},
// 		},
// 	})
// 	require.NoError(t, err)
// 	t.Cleanup(func() {
// 		err := retries.New[struct{}](retries.OnErrors(apierr.ErrResourceConflict)).Wait(ctx, func(ctx context.Context) error {
// 			return CredentialsAPI.DeleteByCredentialsId(ctx, role.CredentialsId)
// 		})
// 		require.NoError(t, err)
// 	})

// 	updateRole, err := CredentialsAPI.Create(ctx, provisioning.CreateCredentialRequest{
// 		CredentialsName: RandomName("go-sdk-"),
// 		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
// 			StsRole: &provisioning.CreateCredentialStsRole{
// 				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
// 			},
// 		},
// 	})
// 	require.NoError(t, err)
// 	t.Cleanup(func() {
// 		err := retries.New[struct{}](retries.OnErrors(apierr.ErrResourceConflict)).Wait(ctx, func(ctx context.Context) error {
// 			return CredentialsAPI.DeleteByCredentialsId(ctx, updateRole.CredentialsId)
// 		})
// 		require.NoError(t, err)
// 	})

// 	// TODO: Add DNS reachability utility
// 	// Do not use CreateAndWait. If the workspaces is created but does not reach running state,
// 	// the cleanup step won't be executed since the test will fail at the `require.NoError(t, err)` line.
// 	WorkspacesAPI, err := provisioning.NewWorkspacesClient(cfg)
// 	require.NoError(t, err)
// 	waiter, err := WorkspacesAPI.Create(ctx, provisioning.CreateWorkspaceRequest{
// 		WorkspaceName:          RandomName("go-sdk-"),
// 		AwsRegion:              GetEnvOrSkipTest(t, "AWS_REGION"),
// 		CredentialsId:          role.CredentialsId,
// 		StorageConfigurationId: storage.StorageConfigurationId,
// 	})
// 	require.NoError(t, err)
// 	t.Cleanup(func() {
// 		err := WorkspacesAPI.DeleteByWorkspaceId(ctx, waiter.WorkspaceId)
// 		require.NoError(t, err)
// 	})
// 	created, err := waiter.Get()
// 	require.NoError(t, err)

// 	// this also takes a while
// 	_, err = WorkspacesAPI.UpdateAndWait(ctx, provisioning.UpdateWorkspaceRequest{
// 		WorkspaceId:   created.WorkspaceId,
// 		CredentialsId: updateRole.CredentialsId,
// 	})
// 	require.NoError(t, err)

// 	byId, err := WorkspacesAPI.GetByWorkspaceId(ctx, created.WorkspaceId)
// 	require.NoError(t, err)

// 	byName, err := WorkspacesAPI.GetByWorkspaceName(ctx, byId.WorkspaceName)
// 	require.NoError(t, err)
// 	assert.Equal(t, byId.WorkspaceId, byName.WorkspaceId)

// 	all, err := WorkspacesAPI.List(ctx)
// 	require.NoError(t, err)

// 	names, err := WorkspacesAPI.WorkspaceWorkspaceNameToWorkspaceIdMap(ctx)
// 	require.NoError(t, err)
// 	assert.Equal(t, len(names), len(all))
// 	assert.Equal(t, byId.WorkspaceId, names[byId.WorkspaceName])
// }
