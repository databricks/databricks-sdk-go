package core

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type configFixture struct {
	host              string
	token             string
	username          string
	password          string
	configFile        string
	profile           string
	azureClientID     string
	azureClientSecret string
	azureTenantID     string
	azureResourceID   string
	authType          string
	env               map[string]string
	assertError       string
	assertAuth        string
	assertHost        string
	assertAzure       bool
}

func (cf configFixture) apply(t *testing.T) {
	c, err := cf.configureProviderAndReturnClient(t)
	if cf.assertError != "" {
		require.NotNilf(t, err, "Expected to have %s error", cf.assertError)
		require.True(t, strings.HasPrefix(err.Error(), cf.assertError),
			"Expected to have '%s' error, but got '%s'", cf.assertError, err)
		return
	}
	if err != nil {
		require.NoError(t, err)
		return
	}
	assert.Equal(t, cf.assertAzure, c.IsAzure())
	assert.Equal(t, cf.assertAuth, c.AuthType)
	assert.Equal(t, cf.assertHost, c.Host)
}

func (cf configFixture) configureProviderAndReturnClient(t *testing.T) (*DatabricksClient, error) {
	defer CleanupEnvironment()()
	for k, v := range cf.env {
		os.Setenv(k, v)
	}
	ctx := context.Background()
	client := &DatabricksClient{
		Host:              cf.host,
		Token:             cf.token,
		Username:          cf.username,
		Password:          cf.password,
		Profile:           cf.profile,
		ConfigFile:        cf.configFile,
		AzureClientID:     cf.azureClientID,
		AzureClientSecret: cf.azureClientSecret,
		AzureTenantID:     cf.azureTenantID,
		AzureResourceID:   cf.azureResourceID,
		AuthType:          cf.authType,
	}
	err := client.Authenticate(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}


func TestConfig_NoParams(t *testing.T) {
	configFixture{
		assertError: "authentication is not configured for provider",
	}.apply(t)
}

func TestConfig_HostEnv(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_HOST": "x",
		},
		assertError: "authentication is not configured for provider",
	}.apply(t)
}

func TestConfig_TokenEnv(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertError: "authentication is not configured for provider. Environment variables used: DATABRICKS_TOKEN",
	}.apply(t)
}

func TestConfig_HostTokenEnv(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_HOST":  "x",
			"DATABRICKS_TOKEN": "x",
		},
		assertAuth: "pat",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_HostParamTokenEnv(t *testing.T) {
	configFixture{
		host: "https://x",
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertAuth: "pat",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_UserPasswordEnv(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertError: "authentication is not configured for provider." +
			" Environment variables used: DATABRICKS_USERNAME, DATABRICKS_PASSWORD",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_BasicAuth(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertAuth: "basic",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_AttributePrecedence(t *testing.T) {
	configFixture{
		host: "y",
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertAuth: "basic",
		assertHost: "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Mix(t *testing.T) {
	configFixture{
		host:     "y",
		username: "x",
		env: map[string]string{
			"DATABRICKS_PASSWORD": "x",
		},
		assertAuth: "basic",
		assertHost: "https://y",
	}.apply(t)
}

func TestConfig_BasicAuth_Attrs(t *testing.T) {
	configFixture{
		host:       "y",
		username:   "x",
		password:   "x",
		assertAuth: "basic",
		assertHost: "https://y",
	}.apply(t)
}

func TestConfig_AzurePAT(t *testing.T) {
	configFixture{
		// Azure hostnames can support host+token auth, as usual
		host:        "https://adb-xxx.y.azuredatabricks.net/",
		token:       "y",
		assertAzure: true,
		assertHost:  "https://adb-xxx.y.azuredatabricks.net/",
		assertAuth:  "pat",
	}.apply(t)
}

func TestConfig_ConflictingEnvs(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_TOKEN":    "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertError: "more than one authorization method configured: password and token",
	}.apply(t)
}

func TestConfig_ConflictingEnvs_AuthType(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_TOKEN":    "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		authType:   "basic",
		assertAuth: "basic",
		assertHost: "https://x",
	}.apply(t)
}

func TestConfig_ConfigFile(t *testing.T) {
	configFixture{
		env: map[string]string{
			"CONFIG_FILE": "x",
		},
		assertError: "authentication is not configured for provider",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg(t *testing.T) {
	configFixture{
		// loading with DEFAULT profile in databrickscfs
		env: map[string]string{
			"HOME": "testdata",
		},
		assertHost: "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com/",
		assertAuth: "databricks-cli",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg_NohostProfile(t *testing.T) {
	configFixture{
		// loading with nohost profile in databrickscfs
		env: map[string]string{
			"HOME":                      "testdata",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
		},
		assertError: "cannot configure databricks-cli auth: config " +
			"file testdata/.databrickscfg is corrupt: cannot find host in nohost profile",
	}.apply(t)
}

func TestConfig_ConfigProfileAndToken(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN":          "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "testdata",
		},
		assertError: "more than one authorization method configured: config profile and token",
	}.apply(t)
}

func TestConfig_ConfigProfileAndPassword(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME":       "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "testdata",
		},
		assertError: "more than one authorization method configured: config profile and password",
	}.apply(t)
}

var azResourceID = "/subscriptions/a/resourceGroups/b/providers/Microsoft.Databricks/workspaces/c"

func TestConfig_AzureCliHost(t *testing.T) {
	configFixture{
		// this test will skip ensureWorkspaceUrl
		host:            "x",
		azureResourceID: azResourceID,
		env: map[string]string{
			// // these may fail on windows. use docker container for testing.
			"PATH": "testdata",
			"HOME": "testdata",
		},
		assertAzure: true,
		assertHost:  "https://x",
		assertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureCliHost_Fail(t *testing.T) {
	configFixture{
		azureResourceID: azResourceID,
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": "testdata",
			"HOME": "testdata",
			"FAIL": "yes",
		},
		assertError: "cannot configure azure-cli auth: Invoking Azure CLI " +
			"failed with the following error: This is just a failing script.",
	}.apply(t)
}

func TestConfig_AzureCliHost_AzNotInstalled(t *testing.T) {
	configFixture{
		// `az` not installed, which is expected for deployers on other clouds...
		azureResourceID: azResourceID,
		env: map[string]string{
			"PATH": "whatever",
			"HOME": "testdata",
		},
		assertError: "cannot configure azure-cli auth: most likely Azure CLI is not installed.",
	}.apply(t)
}

func TestConfig_AzureCliHost_PatConflict(t *testing.T) {
	configFixture{
		azureResourceID: azResourceID,
		token:           "x",
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": "testdata",
			"HOME": "testdata",
		},
		assertError: "more than one authorization method configured: azure and token",
	}.apply(t)
}

func TestConfig_AzureCliHostAndResourceID(t *testing.T) {
	configFixture{
		// omit request to management endpoint to get workspace properties
		azureResourceID: azResourceID,
		host:            "x",
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH": "testdata",
			"HOME": "testdata",
		},
		assertAzure: true,
		assertHost:  "https://x",
		assertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureAndPasswordConflict(t *testing.T) {
	configFixture{
		host:            "x",
		azureResourceID: azResourceID,
		env: map[string]string{
			// these may fail on windows. use docker container for testing.
			"PATH":                "testdata",
			"HOME":                "testdata",
			"DATABRICKS_USERNAME": "x",
		},
		assertError: "more than one authorization method configured: azure and password",
	}.apply(t)
}

func TestConfig_CorruptConfig(t *testing.T) {
	configFixture{
		env: map[string]string{
			"HOME": "testdata/corrupt",
		},
		assertError: "cannot configure databricks-cli auth: " +
			"testdata/corrupt/.databrickscfg has no DEFAULT profile configured",
	}.apply(t)
}
