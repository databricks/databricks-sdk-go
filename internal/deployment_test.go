package internal

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/deployment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMwsAccStorage(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	a := databricks.Must(databricks.NewAccountClient(&databricks.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	}))
	if !a.Config.IsAccountsClient() || !a.Config.IsAws() {
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
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	a := databricks.Must(databricks.NewAccountClient(&databricks.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	}))
	if !a.Config.IsAccountsClient() || !a.Config.IsAws() {
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
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	ctx := context.Background()
	a := databricks.Must(databricks.NewAccountClient(&databricks.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	}))
	if !a.Config.IsAccountsClient() || !a.Config.IsAws() {
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

	defer func() {
		err = a.Credentials.DeleteByCredentialsId(ctx, role.CredentialsId)
		require.NoError(t, err)
	}()

	_, err = a.Credentials.GetByCredentialsId(ctx, role.CredentialsId)
	require.NoError(t, err)

	configs, err := a.Credentials.List(ctx)
	require.NoError(t, err)
	assert.True(t, len(configs) > 0)
}
