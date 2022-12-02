package config

import (
	"net/http"
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/internal/env"
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
	env.CleanupEnvironment(t)
	c, err := cf.configureProviderAndReturnConfig(t)
	if cf.assertError != "" {
		require.NotNilf(t, err, "Expected to have %s error", cf.assertError)
		require.Equal(t, cf.assertError, err.Error())
		return
	}
	if err != nil {
		require.NoError(t, err)
		return
	}
	assert.Equal(t, cf.assertAzure, c.IsAzure(), "Auth is not for Azure")
	assert.Equal(t, cf.assertAuth, c.AuthType)
	assert.Equal(t, cf.assertHost, c.Host)
}

func (cf configFixture) configureProviderAndReturnConfig(t *testing.T) (*Config, error) {
	for k, v := range cf.env {
		os.Setenv(k, v)
	}
	client := &Config{
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
	err := client.Authenticate(&http.Request{Header: http.Header{}})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func TestConfig_NoParams(t *testing.T) {
	configFixture{
		assertError: "default auth: cannot configure default credentials",
	}.apply(t)
}

func TestConfig_HostEnv(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_HOST": "x",
		},
		assertError: "default auth: cannot configure default credentials. Config: host=https://x. Env: DATABRICKS_HOST",
	}.apply(t)
}

func TestConfig_TokenEnv(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN": "x",
		},
		assertError: "default auth: cannot configure default credentials. Config: token=***. Env: DATABRICKS_TOKEN",
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
		assertError: "default auth: cannot configure default credentials. Config: username=x, password=***. Env: DATABRICKS_USERNAME, DATABRICKS_PASSWORD",
		assertHost:  "https://x",
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

func TestConfig_ConflictingEnvs(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_HOST":     "x",
			"DATABRICKS_TOKEN":    "x",
			"DATABRICKS_USERNAME": "x",
			"DATABRICKS_PASSWORD": "x",
		},
		assertError: "validate: more than one authorization method configured: basic and pat. Config: host=x, token=***, username=x, password=***. Env: DATABRICKS_HOST, DATABRICKS_TOKEN, DATABRICKS_USERNAME, DATABRICKS_PASSWORD",
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
			"DATABRICKS_CONFIG_FILE": "x",
		},
		assertError: "default auth: cannot configure default credentials. Config: config_file=x. Env: DATABRICKS_CONFIG_FILE",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg(t *testing.T) {
	configFixture{
		// loading with DEFAULT profile in databrickscfs
		env: map[string]string{
			"HOME": "testdata",
		},
		assertHost: "https://dbc-XXXXXXXX-YYYY.cloud.databricks.com",
		assertAuth: "pat",
	}.apply(t)
}

func TestConfig_PatFromDatabricksCfg_NohostProfile(t *testing.T) {
	configFixture{
		// loading with nohost profile in databrickscfs
		env: map[string]string{
			"HOME":                      "testdata",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
		},
		assertError: "default auth: cannot configure default credentials. " +
			"Config: token=***, profile=nohost. Env: DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_Implicit_DatabricksCfg_Profile_ThroughHost(t *testing.T) {
	configFixture{
		env: map[string]string{
			"HOME":                       "testdata",
			"DATABRICKS_PROFILE_RESOLVE": "true",
			"DATABRICKS_HOST":            "https://dbc-XXXXXXXX-ABCD.cloud.databricks.com",
		},
		assertHost: "https://dbc-XXXXXXXX-ABCD.cloud.databricks.com",
		assertAuth: "basic",
	}.apply(t)
}

func TestConfig_Implicit_DatabricksCfg_Profile_ThroughHost_WithToken(t *testing.T) {
	configFixture{
		env: map[string]string{
			"HOME":                       "testdata",
			"DATABRICKS_PROFILE_RESOLVE": "true",
			"DATABRICKS_HOST":            "https://dbc-XXXXXXXX-ABCD.cloud.databricks.com",
			"DATABRICKS_TOKEN":           "abcd",
		},
		assertError: "validate: more than one authorization method configured: basic and pat. Config: host=https://dbc-XXXXXXXX-ABCD.cloud.databricks.com, token=***, username=abc, password=***, resolve_profile=true. Env: DATABRICKS_HOST, DATABRICKS_TOKEN, DATABRICKS_PROFILE_RESOLVE",
	}.apply(t)
}

func TestConfig_Implicit_DatabricksCfg_Profile_Conflicts(t *testing.T) {
	configFixture{
		env: map[string]string{
			"HOME":                       "testdata",
			"DATABRICKS_PROFILE_RESOLVE": "true",
			"DATABRICKS_HOST":            "https://adb-123.4.azuredatabricks.net",
		},
		assertError: "resolve: azure-justhost and azure-pat and azure-spn match https://adb-123.4.azuredatabricks.net in testdata/.databrickscfg. Config: host=https://adb-123.4.azuredatabricks.net, resolve_profile=true. Env: DATABRICKS_HOST, DATABRICKS_PROFILE_RESOLVE",
	}.apply(t)
}

func TestConfig_ConfigProfileAndToken(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_TOKEN":          "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "testdata",
		},
		assertError: "default auth: cannot configure default credentials. " +
			"Config: token=***, profile=nohost. Env: DATABRICKS_TOKEN, DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_ConfigProfileAndPassword(t *testing.T) {
	configFixture{
		env: map[string]string{
			"DATABRICKS_USERNAME":       "x",
			"DATABRICKS_CONFIG_PROFILE": "nohost",
			"HOME":                      "testdata",
		},
		assertError: "validate: more than one authorization method configured: basic and pat. Config: token=***, username=x, profile=nohost. Env: DATABRICKS_USERNAME, DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

var azResourceID = "/sub/rg/ws"

func TestConfig_AzurePAT(t *testing.T) {
	configFixture{
		// Azure hostnames can support host+token auth, as usual
		host:        "https://adb-xxx.y.azuredatabricks.net/",
		token:       "y",
		assertAzure: true,
		assertHost:  "https://adb-xxx.y.azuredatabricks.net",
		assertAuth:  "pat",
	}.apply(t)
}

func TestConfig_AzureCliHost(t *testing.T) {
	configFixture{
		host:            "x",          // adb-123.4.azuredatabricks.net
		azureResourceID: azResourceID, // skips ensureWorkspaceUrl
		env: map[string]string{
			"PATH": testdataPath(),
			"HOME": "testdata/azure",
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
			"PATH": testdataPath(),
			"HOME": "testdata/azure",
			"FAIL": "yes",
		},
		assertError: "default auth: azure-cli: cannot get access token: This is just a failing script.\n. Config: azure_workspace_resource_id=/sub/rg/ws",
	}.apply(t)
}

func TestConfig_AzureCliHost_AzNotInstalled(t *testing.T) {
	configFixture{
		// `az` not installed, which is expected for deployers on other clouds...
		azureResourceID: azResourceID,
		env: map[string]string{
			"PATH": "whatever",
			"HOME": "testdata/azure",
		},
		assertError: "default auth: cannot configure default credentials. Config: azure_workspace_resource_id=/sub/rg/ws",
	}.apply(t)
}

func TestConfig_AzureCliHost_PatConflict_WithConfigFilePresentWithoutDefaultProfile(t *testing.T) {
	configFixture{
		azureResourceID: azResourceID,
		token:           "x",
		env: map[string]string{
			"PATH": testdataPath(),
			"HOME": "testdata/azure",
		},
		assertError: "validate: more than one authorization method configured: azure and pat. Config: token=***, azure_workspace_resource_id=/sub/rg/ws",
	}.apply(t)
}

func TestConfig_AzureCliHostAndResourceID(t *testing.T) {
	configFixture{
		// omit request to management endpoint to get workspace properties
		azureResourceID: azResourceID,
		host:            "x",
		env: map[string]string{
			"PATH": testdataPath(),
			"HOME": "testdata/azure",
		},
		assertAzure: true,
		assertHost:  "https://x",
		assertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureCliHostAndResourceID_ConfigurationPrecedence(t *testing.T) {
	configFixture{
		// omit request to management endpoint to get workspace properties
		azureResourceID: azResourceID,
		host:            "x",
		env: map[string]string{
			"PATH":                      testdataPath(),
			"HOME":                      "testdata/azure",
			"DATABRICKS_CONFIG_PROFILE": "justhost",
		},
		assertAzure: true,
		assertHost:  "https://x",
		assertAuth:  "azure-cli",
	}.apply(t)
}

func TestConfig_AzureAndPasswordConflict(t *testing.T) { // TODO: this breaks
	configFixture{
		host:            "x",
		azureResourceID: azResourceID,
		env: map[string]string{
			"PATH":                testdataPath(),
			"HOME":                "testdata/azure",
			"DATABRICKS_USERNAME": "x",
		},
		assertError: "validate: more than one authorization method configured: azure and basic. Config: host=x, username=x, azure_workspace_resource_id=/sub/rg/ws. Env: DATABRICKS_USERNAME",
	}.apply(t)
}

func TestConfig_CorruptConfig(t *testing.T) {
	configFixture{
		env: map[string]string{
			"HOME":                      "testdata/corrupt",
			"DATABRICKS_CONFIG_PROFILE": "DEFAULT",
		},
		assertError: "resolve: testdata/corrupt/.databrickscfg has no DEFAULT profile configured. Config: profile=DEFAULT. Env: DATABRICKS_CONFIG_PROFILE",
	}.apply(t)
}

func TestConfig_AuthTypeFromEnv(t *testing.T) {
	configFixture{
		host: "x",
		env: map[string]string{
			"DATABRICKS_TOKEN":     "token",
			"DATABRICKS_USERNAME":  "user",
			"DATABRICKS_PASSWORD":  "password",
			"DATABRICKS_AUTH_TYPE": "basic",
		},
		assertAuth: "basic",
		assertHost: "https://x",
	}.apply(t)
}
