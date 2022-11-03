package databricks_test

import (
	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/workspaces"
)

func ExampleConfig_pat() {
	workspaces.New(&databricks.Config{
		Host:  "https://abc.cloud.databricks.com", // env: DATABRICKS_HOST
		Token: "dapi0c2a3f4e...",                  // env: DATABRICKS_TOKEN
	})
}

func ExampleConfig_basic() {
	workspaces.New(&databricks.Config{
		Host:     "https://abc.cloud.databricks.com", // env: DATABRICKS_HOST
		Username: "me@example.com",                   // env: DATABRICKS_USERNAME
		Password: "som3thing!S@cret",                 // env: DATABRICKS_PASSWORD
	})
}

func ExampleConfig_accounts() {
	workspaces.New(&databricks.Config{
		Host:      "https://accounts.cloud.databricks.com", // env: DATABRICKS_HOST
		AccountID: "00000000-0000-0000-0000-111122223333",  // env: DATABRICKS_ACCOUNT_ID
		Username:  "me@example.com",                        // env: DATABRICKS_USERNAME
		Password:  "som3thing!S@cret",                      // env: DATABRICKS_PASSWORD
	})
}

func ExampleConfig_customProfile() {
	workspaces.New(&databricks.Config{
		Profile: "production", // env: DATABRICKS_CONFIG_PROFILE
	})
}

func ExampleConfig_customConfigFile() {
	workspaces.New(&databricks.Config{
		ConfigFile: "/path/to/.databrickscfg", // env: DATABRICKS_CONFIG_FILE
	})
}

func ExampleConfig_azureActiveDirectoryServicePrincipal() {
	workspaces.New(&databricks.Config{
		Host:              "https://adb-123.4.azuredatabricks.net", // env: DATABRICKS_HOST
		AzureResourceID:   "/subscriptions/../resourceGroups/...",  // env: DATABRICKS_AZURE_RESOURCE_ID
		AzureTenantID:     "00000000-0000-0000-0000-111122223334",  // env: ARM_TENANT_ID
		AzureClientID:     "00000000-0000-0000-0000-111122223335",  // env: ARM_CLIENT_ID
		AzureClientSecret: "som3thing!S@cret",                      // env: ARM_CLIENT_SECRET
	})
}

func ExampleConfig_forceAzureActiveDirectoryServicePrincipal() {
	workspaces.New(&databricks.Config{
		Host:              "https://adb-123.4.azuredatabricks.net", // env: DATABRICKS_HOST
		AzureResourceID:   "/subscriptions/../resourceGroups/...",  // env: DATABRICKS_AZURE_RESOURCE_ID
		AzureTenantID:     "00000000-0000-0000-0000-111122223334",  // env: ARM_TENANT_ID
		AzureClientID:     "00000000-0000-0000-0000-111122223335",  // env: ARM_CLIENT_ID
		AzureClientSecret: "som3thing!S@cret",                      // env: ARM_CLIENT_SECRET
		Credentials:       databricks.AzureClientSecretCredentials{},
	})
}

func ExampleConfig_debugging() {
	workspaces.New(&databricks.Config{
		HTTPTimeoutSeconds:  60,
		DebugTruncateBytes:  96,    // env: DATABRICKS_DEBUG_TRUNCATE_BYTES
		DebugHeaders:        false, // env: DATABRICKS_DEBUG_HEADERS
		RateLimitPerSecond:  15,    // env: DATABRICKS_RATE_LIMIT
		RetryTimeoutSeconds: 300,
	})
}
