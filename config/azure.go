package config

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
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
	AzurePublicCloud = azureEnvironment{
		Name:                      "PUBLIC",
		ServiceManagementEndpoint: "https://management.core.windows.net/",
		ResourceManagerEndpoint:   "https://management.azure.com/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.com/",
	}

	AzureUsGovernmentCloud = azureEnvironment{
		Name:                      "USGOVERNMENT",
		ServiceManagementEndpoint: "https://management.core.usgovcloudapi.net/",
		ResourceManagerEndpoint:   "https://management.usgovcloudapi.net/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.us/",
	}

	AzureChinaCloud = azureEnvironment{
		Name:                      "CHINA",
		ServiceManagementEndpoint: "https://management.core.chinacloudapi.cn/",
		ResourceManagerEndpoint:   "https://management.chinacloudapi.cn/",
		ActiveDirectoryEndpoint:   "https://login.chinacloudapi.cn/",
	}
)

type azureHostResolver interface {
	tokenSourceFor(ctx context.Context, cfg *Config, aadEndpoint, resource string) oauth2.TokenSource
}

func (c *Config) azureEnsureWorkspaceUrl(ctx context.Context, ahr azureHostResolver) error {
	if c.AzureResourceID == "" || c.Host != "" {
		return nil
	}
	azureEnv := c.Environment().AzureEnvironment
	// azure resource ID can also be used in lieu of host by some of the clients, like Terraform
	management := ahr.tokenSourceFor(ctx, c, azureEnv.ActiveDirectoryEndpoint, azureEnv.ResourceManagerEndpoint)
	var workspaceMetadata struct {
		Properties struct {
			WorkspaceURL string `json:"workspaceUrl"`
		} `json:"properties"`
	}
	requestURL := azureEnv.ResourceManagerEndpoint + c.AzureResourceID + "?api-version=2018-04-01"
	err := httpclient.DefaultClient.Do(ctx, "GET", requestURL,
		httpclient.WithResponseUnmarshal(&workspaceMetadata),
		httpclient.WithTokenSource(management),
	)
	if err != nil {
		return fmt.Errorf("resolve workspace: %w", err)
	}
	c.Host = fmt.Sprintf("https://%s", workspaceMetadata.Properties.WorkspaceURL)
	logger.Debugf(ctx, "Discovered workspace url: %s", c.Host)
	return nil
}
