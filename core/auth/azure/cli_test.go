package azure

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/databricks/sdk-go/core"
	"github.com/databricks/sdk-go/core/client"
	"github.com/databricks/sdk-go/core/config"
)

var azDummy = &client.ApiClient{Config: &config.Config{Host: "https://adb-xyz.c.azuredatabricks.net/"}}

func TestAzureCliCredentials_SkipAws(t *testing.T) {
	aa := AzureCliCredentials{}
	x, err := aa.Configure(context.Background(), 
		&client.ApiClient{Config: &config.Config{Host: "https://xyz.cloud.databricks.com/"}})
	assert.Nil(t, x)
	assert.NoError(t, err)
}

func TestAzureCliCredentials_NotInstalled(t *testing.T) {
	defer core.CleanupEnvironment()()
	os.Setenv("PATH", "whatever")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	require.Error(t, err)
	require.EqualError(t, err, "most likely Azure CLI is not installed. "+
		"See https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest for details")
}

func TestAzureCliCredentials_NotLoggedIn(t *testing.T) {
	defer core.CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("FAIL", "logout")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	assert.NoError(t, err)
}

func TestAzureCliCredentials_Valid(t *testing.T) {
	defer core.CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	aa := AzureCliCredentials{}
	visitor, err := aa.Configure(context.Background(), azDummy)
	assert.NoError(t, err)

	r := &http.Request{Header: http.Header{}}
	err = visitor(r)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer ...", r.Header.Get("Authorization"))
}

func TestAzureCliCredentials_AlwaysExpired(t *testing.T) {
	defer core.CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("EXPIRE", "10S")
	aa := AzureCliCredentials{}
	visitor, err := aa.Configure(context.Background(), azDummy)
	assert.NoError(t, err)

	r := &http.Request{Header: http.Header{}}
	err = visitor(r)

	assert.NoError(t, err)
}

func TestAzureCliCredentials_ExitError(t *testing.T) {
	defer core.CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("FAIL", "yes")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	assert.EqualError(t, err, "cannot get access token: This is just a failing script.\n")
}

func TestAzureCliCredentials_Corrupt(t *testing.T) {
	defer core.CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("FAIL", "corrupt")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	assert.EqualError(t, err, "cannot unmarshal CLI result: invalid character 'a' looking for beginning of object key string")
}

func TestAzureCliCredentials_CorruptExpire(t *testing.T) {
	defer core.CleanupEnvironment()()
	os.Setenv("PATH", "testdata:/bin")
	os.Setenv("EXPIRE", "corrupt")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	assert.EqualError(t, err, "cannot parse expiry: parsing time \"\" as \"2006-01-02 15:04:05.999999\": cannot parse \"\" as \"2006\"")
}

// TODO: this test should rather be on sequencing
// func TestConfigureWithAzureCLI_SP(t *testing.T) {
// 	aa := DatabricksClient{
// 		AzureClientID:     "a",
// 		AzureClientSecret: "b",
// 		AzureTenantID:     "c",
// 		AzureResourceID:   "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c",
// 	}
// 	ctx := context.Background()
// 	auth, err := aa.configureWithAzureCLI(ctx)
// 	assert.NoError(t, err)
// 	assert.Nil(t, auth)
// }
