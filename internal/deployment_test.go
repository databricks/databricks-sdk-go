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

	storage, err := a.StorageConfigurations.CreateStorageConfig(ctx, deployment.CreateStorageConfigurationRequest{
		StorageConfigurationName: RandomName("sdk-"),
		RootBucketInfo: deployment.RootBucketInfo{
			BucketName: RandomName("sdk-bucket-"),
		},
	})
	require.NoError(t, err)

	defer func() {
		err = a.StorageConfigurations.DeleteStorageConfigByStorageConfigurationId(ctx, storage.StorageConfigurationId)
		require.NoError(t, err)
	}()

	_, err = a.StorageConfigurations.GetStorageConfigByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	require.NoError(t, err)

	configs, err := a.StorageConfigurations.ListStorageConfigs(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccNetworks(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}
	netw, err := a.NetworkConfigurations.CreateNetworkConfig(ctx, deployment.CreateNetworkRequest{
		NetworkName:      RandomName("sdk-"),
		VpcId:            RandomHex("vpc-", 17),
		SubnetIds:        []string{RandomHex("subnet-", 17), RandomHex("subnet-", 17)},
		SecurityGroupIds: []string{RandomHex("sg-", 17)},
	})
	require.NoError(t, err)
	defer func() {
		err = a.NetworkConfigurations.DeleteNetworkConfigByNetworkId(ctx, netw.NetworkId)
		require.NoError(t, err)
	}()

	_, err = a.NetworkConfigurations.GetNetworkConfigByNetworkId(ctx, netw.NetworkId)
	require.NoError(t, err)

	configs, err := a.NetworkConfigurations.ListNetworkConfigs(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}

func TestMwsAccCredentials(t *testing.T) {
	ctx, a := accountTest(t)
	if !a.Config.IsAws() {
		t.SkipNow()
	}
	role, err := a.CredentialConfigurations.CreateCredentialConfig(ctx, deployment.CreateCredentialRequest{
		CredentialsName: RandomName("sdk-"),
		AwsCredentials: deployment.AwsCredentials{
			StsRole: &deployment.StsRole{
				RoleArn: GetEnvOrSkipTest(t, "TEST_CROSSACCOUNT_ARN"),
			},
		},
	})
	require.NoError(t, err)

	defer func() {
		err = a.CredentialConfigurations.DeleteCredentialConfigByCredentialsId(ctx, role.CredentialsId)
		require.NoError(t, err)
	}()

	_, err = a.CredentialConfigurations.GetCredentialConfigByCredentialsId(ctx, role.CredentialsId)
	require.NoError(t, err)

	configs, err := a.CredentialConfigurations.ListCredentials(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}
