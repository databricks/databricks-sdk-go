package internal

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/service/deployment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccStorage(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	storage, err := a.Storage.Create(ctx, deployment.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: deployment.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	require.NoError(t, err)

	defer func() {
		err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
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
	netw, err := a.Networks.Create(ctx, deployment.CreateNetworkRequest{
		NetworkName:      RandomName("sdk-"),
		VpcId:            RandomHex("vpc-", 17),
		SubnetIds:        []string{RandomHex("subnet-", 17), RandomHex("subnet-", 17)},
		SecurityGroupIds: []string{RandomHex("sg-", 17)},
	})
	require.NoError(t, err)
	defer func() {
		err = a.Networks.DeleteByNetworkId(ctx, netw.NetworkId)
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
	role, err := a.Credentials.Create(ctx, deployment.CreateCredentialRequest{
		CredentialsName: RandomName("sdk-"),
		AwsCredentials: deployment.AwsCredentials{
			StsRole: &deployment.StsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
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

	created, err := a.EncryptionKeys.Create(ctx, deployment.CreateCustomerManagedKeyRequest{
		AwsKeyInfo: deployment.CreateAwsKeyInfo{
			KeyArn:   GetEnvOrSkipTest(t, "TEST_MANAGED_KMS_KEY_ARN"),
			KeyAlias: GetEnvOrSkipTest(t, "TEST_STORAGE_KMS_KEY_ALIAS"),
		},
		UseCases: []string{"MANAGED_SERVICES"}, // TODO: OpenAPI: add enum
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := a.EncryptionKeys.DeleteByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
		require.NoError(t, err)
	})

	byId, err := a.EncryptionKeys.GetByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
	require.NoError(t, err)
	assert.Equal(t, "MANAGED_SERVICES", byId.UseCases[0])

	all, err := a.EncryptionKeys.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestMwsAccPrivateAccess(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	// TODO: OpenAPI: rename entities
	created, err := a.PrivateAccess.Create(ctx, deployment.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsName: RandomName("go-sdk-"),
		Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
		require.NoError(t, err)
	})
	err = a.PrivateAccess.Replace(ctx, deployment.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsId:   created.PrivateAccessSettingsId,
		PrivateAccessSettingsName: RandomName("go-sdk-"),
		Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
	})
	require.NoError(t, err)

	byId, err := a.PrivateAccess.GetByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	require.NoError(t, err)

	all, err := a.PrivateAccess.List(ctx)
	require.NoError(t, err)

	names, err := a.PrivateAccess.PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.PrivateAccessSettingsId, names[byId.PrivateAccessSettingsName])
}

func TestMwsAccVpcEndpoints(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	created, err := a.VpcEndpoints.Create(ctx, deployment.CreateVpcEndpointRequest{
		AwsVpcEndpointId: GetEnvOrSkipTest(t, "TEST_RELAY_VPC_ENDPOINT"),
		Region:           GetEnvOrSkipTest(t, "AWS_REGION"),
		VpcEndpointName:  RandomName("go-sdk-"),
	})
	require.NoError(t, err)

	t.Cleanup(func() {
		err := a.VpcEndpoints.DeleteByVpcEndpointId(ctx, created.VpcEndpointId)
		require.NoError(t, err)
	})

	byId, err := a.VpcEndpoints.GetByVpcEndpointId(ctx, created.VpcEndpointId)
	require.NoError(t, err)
	assert.Equal(t, deployment.EndpointUseCaseDataplaneRelayAccess, byId.UseCase)

	all, err := a.VpcEndpoints.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestMwsAccWorkspaces(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	storage, err := a.Storage.Create(ctx, deployment.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("go-sdk-"),
		RootBucketInfo: deployment.RootBucketInfo{
			BucketName: GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET"),
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		require.NoError(t, err)
	})

	// TODO: OpenAPI: Document retry protocol on AWS IAM registration errors
	// See https://github.com/databricks/terraform-provider-databricks/issues/1424
	role, err := a.Credentials.Create(ctx, deployment.CreateCredentialRequest{
		CredentialsName: RandomName("go-sdk-"),
		AwsCredentials: deployment.AwsCredentials{
			StsRole: &deployment.StsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
		require.NoError(t, err)
	})

	// TODO: Add DNS reachability utility
	created, err := a.Workspaces.CreateAndWait(ctx, deployment.CreateWorkspaceRequest{
		WorkspaceName:          RandomName("go-sdk-"),
		AwsRegion:              GetEnvOrSkipTest(t, "AWS_REGION"),
		CredentialsId:          role.CredentialsId,
		StorageConfigurationId: storage.StorageConfigurationId,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		err := a.Workspaces.DeleteByWorkspaceId(ctx, created.WorkspaceId)
		require.NoError(t, err)
	})

	updateRole, err := a.Credentials.Create(ctx, deployment.CreateCredentialRequest{
		CredentialsName: RandomName("go-sdk-"),
		AwsCredentials: deployment.AwsCredentials{
			StsRole: &deployment.StsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)
	// TODO: OpenAPI: add error retry to properly wait for things like
	// 409 Conflict - INVALID_STATE: Unable to delete, credential is being used by an active workspace
	defer a.Credentials.DeleteByCredentialsId(ctx, updateRole.CredentialsId)
	// t.Cleanup(func() {
	// 	err := a.Credentials.DeleteByCredentialsId(ctx, updateRole.CredentialsId)
	// 	require.NoError(t, err)
	// })

	// this also takes a while
	_, err = a.Workspaces.UpdateAndWait(ctx, deployment.UpdateWorkspaceRequest{
		WorkspaceId:   created.WorkspaceId,
		CredentialsId: updateRole.CredentialsId,
	})
	require.NoError(t, err)

	byId, err := a.Workspaces.GetByWorkspaceId(ctx, created.WorkspaceId)
	require.NoError(t, err)

	all, err := a.Workspaces.List(ctx)
	require.NoError(t, err)

	names, err := a.Workspaces.WorkspaceWorkspaceNameToWorkspaceIdMap(ctx)
	require.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.WorkspaceId, names[byId.WorkspaceName])
}
