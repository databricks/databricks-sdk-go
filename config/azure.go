package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
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
	if c.AzureEnvironment == "" {
		c.AzureEnvironment = "public"
	}
	env, ok := azureEnvironments[strings.ToUpper(c.AzureEnvironment)]
	if !ok {
		return env, fmt.Errorf("azure envionment not found: %s", c.AzureEnvironment)
	}
	return env, nil
}

type azureHostResolver interface {
	tokenSourceFor(ctx context.Context, cfg *Config, env azureEnvironment, resource string) oauth2.TokenSource
}

func (c *Config) azureEnsureWorkspaceUrl(ctx context.Context, ahr azureHostResolver) error {
	if c.AzureResourceID == "" || c.Host != "" {
		return nil
	}
	env, err := c.GetAzureEnvironment()
	if err != nil {
		return err
	}
	// azure resource ID can also be used in lieu of host by some of the clients, like Terraform
	management := ahr.tokenSourceFor(ctx, c, env, env.ResourceManagerEndpoint)
	resourceManager := oauth2.NewClient(ctx, management)
	resp, err := resourceManager.Get(env.ResourceManagerEndpoint + c.AzureResourceID + "?api-version=2018-04-01")
	if err != nil {
		return fmt.Errorf("cannot resolve workspace: %w", err)
	}
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("cannot read: %w", err)
	}
	var workspaceMetadata struct {
		Properties struct {
			WorkspaceURL string `json:"workspaceUrl"`
		} `json:"properties"`
	}
	err = json.Unmarshal(raw, &workspaceMetadata)
	if err != nil {
		return fmt.Errorf("cannot unmarshal: %w", err)
	}
	c.Host = fmt.Sprintf("https://%s", workspaceMetadata.Properties.WorkspaceURL)
	logger.Debugf(ctx, "Discovered workspace url: %s", c.Host)
	return nil
}

// Resource ID of the Azure application we need to log in.
const azureDatabricksLoginAppID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"

func (c *Config) getAzureLoginAppID() string {
	if c.AzureLoginAppID != "" {
		return c.AzureLoginAppID
	}
	return azureDatabricksLoginAppID
}
