package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"golang.org/x/oauth2"
)

// TODO: this file is a reduced copy from github.com/Azure/go-autorest/autorest/azure/environments.go
var environments = map[string]Environment{
	"CHINA":        chinaCloud,
	"GERMAN":       germanCloud,
	"PUBLIC":       publicCloud,
	"USGOVERNMENT": usGovernmentCloud,
}

type azureCommon struct {
	AzureResourceID string `name:"azure_workspace_resource_id" env:"DATABRICKS_AZURE_RESOURCE_ID"`
	Environment     string `name:"azure_environment" env:"ARM_ENVIRONMENT"`
}

func (ac azureCommon) getEnvironment() (Environment, error) {
	if ac.Environment == "" {
		ac.Environment = "public"
	}
	env, ok := environments[strings.ToUpper(ac.Environment)]
	if !ok {
		return env, fmt.Errorf("azure envionment not found: %s", ac.Environment)
	}
	return env, nil
}

func (az azureCommon) getWorkspaceUrl(ctx context.Context, management oauth2.TokenSource) (string, error) {
	env, err := az.getEnvironment()
	if err != nil {
		return "", err
	}
	resourceManager := oauth2.NewClient(ctx, management)
	resp, err := resourceManager.Get(env.ResourceManagerEndpoint + az.AzureResourceID + "?api-version=2018-04-01")
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

type Environment struct {
	Name                      string `json:"name"`
	ServiceManagementEndpoint string `json:"serviceManagementEndpoint"`
	ResourceManagerEndpoint   string `json:"resourceManagerEndpoint"`
	ActiveDirectoryEndpoint   string `json:"activeDirectoryEndpoint"`
}

var (
	// publicCloud is the default public Azure cloud environment
	publicCloud = Environment{
		Name:                      "AzurePublicCloud",
		ServiceManagementEndpoint: "https://management.core.windows.net/",
		ResourceManagerEndpoint:   "https://management.azure.com/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.com/",
	}

	// usGovernmentCloud is the cloud environment for the US Government
	usGovernmentCloud = Environment{
		Name:                      "AzureUSGovernmentCloud",
		ServiceManagementEndpoint: "https://management.core.usgovcloudapi.net/",
		ResourceManagerEndpoint:   "https://management.usgovcloudapi.net/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.us/",
	}

	// chinaCloud is the cloud environment operated in China
	chinaCloud = Environment{
		Name:                      "AzureChinaCloud",
		ServiceManagementEndpoint: "https://management.core.chinacloudapi.cn/",
		ResourceManagerEndpoint:   "https://management.chinacloudapi.cn/",
		ActiveDirectoryEndpoint:   "https://login.chinacloudapi.cn/",
	}

	// germanCloud is the cloud environment operated in Germany
	germanCloud = Environment{
		Name:                      "AzureGermanCloud",
		ServiceManagementEndpoint: "https://management.core.cloudapi.de/",
		ResourceManagerEndpoint:   "https://management.microsoftazure.de/",
		ActiveDirectoryEndpoint:   "https://login.microsoftonline.de/",
	}
)
