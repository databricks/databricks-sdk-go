package databricks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"golang.org/x/oauth2"
)

type azureEnvironment struct {
	Name                      string `json:"name"`
	ServiceManagementEndpoint string `json:"serviceManagementEndpoint"`
	ResourceManagerEndpoint   string `json:"resourceManagerEndpoint"`
	ActiveDirectoryEndpoint   string `json:"activeDirectoryEndpoint"`
}

// based on github.com/Azure/go-autorest/autorest/azure/azureEnvironments.go
var (
	publicCloud = azureEnvironment{
		Name:                      "AzurePublicCloud",
		ServiceManagementEndpoint: "https://management.core.windows.net/",
		ResourceManagerEndpoint:   "https://management.azure.com/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.com/",
	}

	usGovernmentCloud = azureEnvironment{
		Name:                      "AzureUSGovernmentCloud",
		ServiceManagementEndpoint: "https://management.core.usgovcloudapi.net/",
		ResourceManagerEndpoint:   "https://management.usgovcloudapi.net/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.us/",
	}

	chinaCloud = azureEnvironment{
		Name:                      "AzureChinaCloud",
		ServiceManagementEndpoint: "https://management.core.chinacloudapi.cn/",
		ResourceManagerEndpoint:   "https://management.chinacloudapi.cn/",
		ActiveDirectoryEndpoint:   "https://login.chinacloudapi.cn/",
	}

	germanCloud = azureEnvironment{
		Name:                      "AzureGermanCloud",
		ServiceManagementEndpoint: "https://management.core.cloudapi.de/",
		ResourceManagerEndpoint:   "https://management.microsoftazure.de/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.de/",
	}

	azureEnvironments = map[string]azureEnvironment{
		"CHINA":        chinaCloud,
		"GERMAN":       germanCloud,
		"PUBLIC":       publicCloud,
		"USGOVERNMENT": usGovernmentCloud,
	}
)

func (c *Config) GetAzureEnvironment() (azureEnvironment, error) {
	if c.Environment == "" {
		c.Environment = "public"
	}
	env, ok := azureEnvironments[strings.ToUpper(c.Environment)]
	if !ok {
		return env, fmt.Errorf("azure envionment not found: %s", c.Environment)
	}
	return env, nil
}

func (c *Config) getWorkspaceUrl(ctx context.Context, management oauth2.TokenSource) (string, error) {
	env, err := c.GetAzureEnvironment()
	if err != nil {
		return "", err
	}
	resourceManager := oauth2.NewClient(ctx, management)
	resp, err := resourceManager.Get(env.ResourceManagerEndpoint + c.AzureResourceID + "?api-version=2018-04-01")
	if err != nil {
		return "", fmt.Errorf("cannot resolve workspace: %w", err)
	}
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("cannot read: %w", err)
	}
	var workspaceMetadata struct {
		Properties struct {
			WorkspaceURL string `json:"workspaceUrl"`
		} `json:"properties"`
	}
	err = json.Unmarshal(raw, &workspaceMetadata)
	if err != nil {
		return "", fmt.Errorf("cannot unmarshal: %w", err)
	}
	return fmt.Sprintf("https://%s/", workspaceMetadata.Properties.WorkspaceURL), nil
}
