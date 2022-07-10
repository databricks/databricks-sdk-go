package databricks

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/sdk-go/databricks/internal"
)

type AzureClientSecretCredentials struct {
	TenantID     string `name:"azure_tenant_id" env:"ARM_TENANT_ID"`
	ClientID     string `name:"azure_client_id" env:"ARM_CLIENT_ID"`
	ClientSecret string `name:"azure_client_secret" env:"ARM_CLIENT_SECRET" auth:"sensitive"`
}

func (c AzureClientSecretCredentials) Name() string {
	return "azure-client-secret"
}

func (c AzureClientSecretCredentials) tokenSourceFor(ctx context.Context, aadEndpoint, resource string) oauth2.TokenSource {
	return (&clientcredentials.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		TokenURL:     fmt.Sprintf("%s%s/oauth2/token", aadEndpoint, c.TenantID),
		EndpointParams: url.Values{
			"resource": []string{resource},
		},
	}).TokenSource(ctx)
}

func (c AzureClientSecretCredentials) ensureWorkspaceUrl(ctx context.Context, env azureEnvironment, cfg *Config) error {
	if cfg.Host == "" {
		// azure resource ID can also be used in lieu of host by some of the clients, like Terraform
		management := c.tokenSourceFor(ctx, env.ActiveDirectoryEndpoint, env.ResourceManagerEndpoint)
		host, err := cfg.getWorkspaceUrl(ctx, management)
		if err != nil {
			return fmt.Errorf("cannot get workspace url: %w", err)
		}
		log.Printf("[INFO] Discovered workspace url: %s", host)
		cfg.Host = host
	}
	return nil
}

func (c AzureClientSecretCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	// TODO: this has to be exposed for Terraform, as we cannot create AKV scopes by SPNs
	if c.ClientID == "" || c.ClientSecret == "" || c.TenantID == "" || cfg.AzureResourceID == "" {
		return nil, nil
	}
	if !cfg.IsAzure() {
		return nil, nil
	}
	env, err := cfg.GetAzureEnvironment()
	if err != nil {
		return nil, err
	}

	err = c.ensureWorkspaceUrl(ctx, env, cfg)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Generating AAD token for Service Principal (%s)", c.ClientID)
	return func(r *http.Request) error {
		ctx := r.Context()
		r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", cfg.AzureResourceID)
		return internal.ServiceToServiceVisitor(
			c.tokenSourceFor(ctx, env.ActiveDirectoryEndpoint, armDatabricksResourceID),
			c.tokenSourceFor(ctx, env.ActiveDirectoryEndpoint, env.ServiceManagementEndpoint),
			"X-Databricks-Azure-SP-Management-Token")(r)
	}, nil
}
