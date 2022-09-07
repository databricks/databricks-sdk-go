package databricks

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var azDummy = &Config{Host: "https://adb-xyz.c.azuredatabricks.net/"}

// testdataPath returns the PATH to use for the duration of a test.
// It must only return absolute directories because Go refuses to run
// exexutables found in a relative directory as of 1.19.
// More information at https://tip.golang.org/doc/go1.19#os-exec-path.
func testdataPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	paths := []string{
		filepath.Join(cwd, "testdata"),
		"/bin",
	}
	return strings.Join(paths, ":")
}

func TestAzureCliCredentials_SkipAws(t *testing.T) {
	aa := AzureCliCredentials{}
	x, err := aa.Configure(context.Background(), &Config{Host: "https://xyz.cloud.databricks.com/"})
	assert.Nil(t, x)
	assert.NoError(t, err)
}

func TestAzureCliCredentials_NotInstalled(t *testing.T) {
	env.CleanupEnvironment(t)
	os.Setenv("PATH", "whatever")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	require.NoError(t, err)
}

func TestAzureCliCredentials_NotLoggedIn(t *testing.T) {
	env.CleanupEnvironment(t)
	os.Setenv("PATH", testdataPath())
	os.Setenv("FAIL", "logout")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	assert.NoError(t, err)
}

func TestAzureCliCredentials_Valid(t *testing.T) {
	env.CleanupEnvironment(t)
	os.Setenv("PATH", testdataPath())
	aa := AzureCliCredentials{}
	visitor, err := aa.Configure(context.Background(), azDummy)
	assert.NoError(t, err)

	r := &http.Request{Header: http.Header{}}
	err = visitor(r)
	assert.NoError(t, err)

	assert.Equal(t, "Bearer ...", r.Header.Get("Authorization"))
}

func TestAzureCliCredentials_AlwaysExpired(t *testing.T) {
	env.CleanupEnvironment(t)
	os.Setenv("PATH", testdataPath())
	os.Setenv("EXPIRE", "10S")
	aa := AzureCliCredentials{}
	visitor, err := aa.Configure(context.Background(), azDummy)
	assert.NoError(t, err)

	r := &http.Request{Header: http.Header{}}
	err = visitor(r)

	assert.NoError(t, err)
}

func TestAzureCliCredentials_ExitError(t *testing.T) {
	env.CleanupEnvironment(t)
	os.Setenv("PATH", testdataPath())
	os.Setenv("FAIL", "yes")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	assert.EqualError(t, err, "cannot get access token: This is just a failing script.\n")
}

func TestAzureCliCredentials_Corrupt(t *testing.T) {
	env.CleanupEnvironment(t)
	os.Setenv("PATH", testdataPath())
	os.Setenv("FAIL", "corrupt")
	aa := AzureCliCredentials{}
	_, err := aa.Configure(context.Background(), azDummy)
	assert.EqualError(t, err, "cannot unmarshal CLI result: invalid character 'a' looking for beginning of object key string")
}

func TestAzureCliCredentials_CorruptExpire(t *testing.T) {
	env.CleanupEnvironment(t)
	os.Setenv("PATH", testdataPath())
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
