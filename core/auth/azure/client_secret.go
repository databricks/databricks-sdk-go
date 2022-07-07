package azure

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/databricks/sdk-go/core/auth/internal"
	"github.com/databricks/sdk-go/core/client"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type ClientSecretCredentials struct {
	AzureResourceID string `name:"azure_workspace_resource_id" env:"DATABRICKS_AZURE_RESOURCE_ID"`
	Environment     string `name:"azure_environment" env:"ARM_ENVIRONMENT"`
	TenantID        string `name:"azure_tenant_id" env:"ARM_TENANT_ID"`
	ClientID        string `name:"azure_client_id" env:"ARM_CLIENT_ID"`
	ClientSecret    string `name:"azure_client_secret" env:"ARM_CLIENT_SECRET" auth:"sensitive"`
}

func (c ClientSecretCredentials) Name() string {
	return "azure-client-secret"
}

func (c ClientSecretCredentials) tokenSourceFor(ctx context.Context, aadEndpoint, resource string) oauth2.TokenSource {
	return (&clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     fmt.Sprintf("%s%s/oauth2/token", aadEndpoint, c.TenantID),
		EndpointParams: url.Values{
			"resource": []string{resource},
		},
	}).TokenSource(ctx)
}

func (c ClientSecretCredentials) ensureWorkspaceUrl(ctx context.Context, az azureCommon, env Environment, conf *client.ApiClient) error {
	if conf.Config.Host == "" {
		// azure resource ID can also be used in lieu of host by some of the clients, like Terraform
		management := c.tokenSourceFor(ctx, env.ActiveDirectoryEndpoint, env.ResourceManagerEndpoint)
		host, err := az.getWorkspaceUrl(ctx, management)
		if err != nil {
			return fmt.Errorf("cannot get workspace url: %w", err)
		}
		log.Printf("[INFO] Discovered workspace url: %s", host)
		conf.Config.Host = host
	}
	return nil
}

func (c ClientSecretCredentials) Configure(ctx context.Context, conf *client.ApiClient) (func(*http.Request) error, error) {
	// TODO: this has to be exposed for Terraform, as we cannot create AKV scopes by SPNs
	if c.ClientID == "" || c.ClientSecret == "" || c.TenantID == "" || c.AzureResourceID == "" {
		return nil, nil
	}
	az := azureCommon{c.AzureResourceID, c.Environment}
	env, err := az.getEnvironment()
	if err != nil {
		return nil, err
	}
	err = c.ensureWorkspaceUrl(ctx, az, env, conf)
	if err != nil {
		return nil, err
	}
	host := internal.Host(conf.Config.Host)
	if !host.IsAzure() {
		return nil, nil
	}
	log.Printf("[INFO] Generating AAD token for Service Principal (%s)", c.ClientID)
	return func(r *http.Request) error {
		ctx := r.Context()
		r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", c.AzureResourceID)
		return internal.ServiceToServiceVisitor(
			c.tokenSourceFor(ctx, env.ActiveDirectoryEndpoint, armDatabricksResourceID),
			c.tokenSourceFor(ctx, env.ActiveDirectoryEndpoint, env.ServiceManagementEndpoint),
			"X-Databricks-Azure-SP-Management-Token")(r)
	}, nil
}
