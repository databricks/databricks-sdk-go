package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/databricks/databricks-sdk-go/service/provisioning"
	"github.com/stretchr/testify/assert"
)

func TestMwsAccStorage(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	assert.NoError(t, err)

	defer func() {
		err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		assert.NoError(t, err)
	}()

	byId, err := a.Storage.GetByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	assert.NoError(t, err)

	byName, err := a.Storage.GetByStorageConfigurationName(ctx, byId.StorageConfigurationName)
	assert.NoError(t, err)
	assert.Equal(t, byId.StorageConfigurationId, byName.StorageConfigurationId)

	configs, err := a.Storage.List(ctx)
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	defer func() {
		err = a.Networks.DeleteByNetworkId(ctx, netw.NetworkId)
		assert.NoError(t, err)
	}()

	byId, err := a.Networks.GetByNetworkId(ctx, netw.NetworkId)
	assert.NoError(t, err)

	byName, err := a.Networks.GetByNetworkName(ctx, byId.NetworkName)
	assert.NoError(t, err)
	assert.Equal(t, byId.NetworkId, byName.NetworkId)

	configs, err := a.Networks.List(ctx)
	assert.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccCredentials(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
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
	assert.NoError(t, err)

	t.Cleanup(func() {
		err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
		assert.NoError(t, err)
	})

	byId, err := a.Credentials.GetByCredentialsId(ctx, role.CredentialsId)
	assert.NoError(t, err)

	byName, err := a.Credentials.GetByCredentialsName(ctx, byId.CredentialsName)
	assert.NoError(t, err)
	assert.Equal(t, byId.CredentialsId, byName.CredentialsId)

	configs, err := a.Credentials.List(ctx)
	assert.NoError(t, err)
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
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := a.EncryptionKeys.DeleteByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
		assert.NoError(t, err)
	})

	byId, err := a.EncryptionKeys.GetByCustomerManagedKeyId(ctx, created.CustomerManagedKeyId)
	assert.NoError(t, err)
	assert.Equal(t, provisioning.KeyUseCaseManagedServices, byId.UseCases[0])

	all, err := a.EncryptionKeys.List(ctx)
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestMwsAccPrivateAccess(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsName: RandomName("go-sdk-"),
		Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
	})
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
		assert.NoError(t, err)
	})
	err = a.PrivateAccess.Replace(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsId:   created.PrivateAccessSettingsId,
		PrivateAccessSettingsName: RandomName("go-sdk-"),
		Region:                    GetEnvOrSkipTest(t, "AWS_REGION"),
	})
	assert.NoError(t, err)

	byId, err := a.PrivateAccess.GetByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	assert.NoError(t, err)

	byName, err := a.PrivateAccess.GetByPrivateAccessSettingsName(ctx, byId.PrivateAccessSettingsName)
	assert.NoError(t, err)
	assert.Equal(t, byId.PrivateAccessSettingsId, byName.PrivateAccessSettingsId)

	all, err := a.PrivateAccess.List(ctx)
	assert.NoError(t, err)

	names, err := a.PrivateAccess.PrivateAccessSettingsPrivateAccessSettingsNameToPrivateAccessSettingsIdMap(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.PrivateAccessSettingsId, names[byId.PrivateAccessSettingsName])
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
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := a.VpcEndpoints.DeleteByVpcEndpointId(ctx, created.VpcEndpointId)
		assert.NoError(t, err)
	})

	byId, err := a.VpcEndpoints.GetByVpcEndpointId(ctx, created.VpcEndpointId)
	assert.NoError(t, err)
	assert.Equal(t, provisioning.EndpointUseCaseDataplaneRelayAccess, byId.UseCase)

	all, err := a.VpcEndpoints.List(ctx)
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)
}

func TestMwsAccWorkspaces(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("go-sdk-"),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: GetEnvOrSkipTest(t, "TEST_ROOT_BUCKET"),
		},
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		assert.NoError(t, err)
	})

	// TODO: OpenAPI: Document retry protocol on AWS IAM registration errors
	// See https://github.com/databricks/terraform-provider-databricks/issues/1424
	role, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("go-sdk-"),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := retries.New[struct{}](retries.OnErrors(apierr.ErrResourceConflict)).Wait(ctx, func(ctx context.Context) error {
			return a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
		})
		assert.NoError(t, err)
	})

	updateRole, err := a.Credentials.Create(ctx, provisioning.CreateCredentialRequest{
		CredentialsName: RandomName("go-sdk-"),
		AwsCredentials: provisioning.CreateCredentialAwsCredentials{
			StsRole: &provisioning.CreateCredentialStsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := retries.New[struct{}](retries.OnErrors(apierr.ErrResourceConflict)).Wait(ctx, func(ctx context.Context) error {
			return a.Credentials.DeleteByCredentialsId(ctx, updateRole.CredentialsId)
		})
		assert.NoError(t, err)
	})

	// TODO: Add DNS reachability utility
	created, err := a.Workspaces.CreateAndWait(ctx, provisioning.CreateWorkspaceRequest{
		WorkspaceName:          RandomName("go-sdk-"),
		AwsRegion:              GetEnvOrSkipTest(t, "AWS_REGION"),
		CredentialsId:          role.CredentialsId,
		StorageConfigurationId: storage.StorageConfigurationId,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		err := a.Workspaces.DeleteByWorkspaceId(ctx, created.WorkspaceId)
		assert.NoError(t, err)
	})

	// this also takes a while
	_, err = a.Workspaces.UpdateAndWait(ctx, provisioning.UpdateWorkspaceRequest{
		WorkspaceId:   created.WorkspaceId,
		CredentialsId: updateRole.CredentialsId,
	})
	assert.NoError(t, err)

	byId, err := a.Workspaces.GetByWorkspaceId(ctx, created.WorkspaceId)
	assert.NoError(t, err)

	byName, err := a.Workspaces.GetByWorkspaceName(ctx, byId.WorkspaceName)
	assert.NoError(t, err)
	assert.Equal(t, byId.WorkspaceId, byName.WorkspaceId)

	all, err := a.Workspaces.List(ctx)
	assert.NoError(t, err)

	names, err := a.Workspaces.WorkspaceWorkspaceNameToWorkspaceIdMap(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(names), len(all))
	assert.Equal(t, byId.WorkspaceId, names[byId.WorkspaceName])
}
