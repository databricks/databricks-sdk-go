package config

import (
	"context"
	"fmt"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
)

type AzureClientSecretCredentials struct {
}

func (c AzureClientSecretCredentials) Name() string {
	return "azure-client-secret"
}

func (c AzureClientSecretCredentials) tokenSourceFor(
	ctx context.Context, cfg *Config, aadEndpoint, resource string) oauth2.TokenSource {
	return (&clientcredentials.Config{
		ClientID:     cfg.AzureClientID,
		ClientSecret: cfg.AzureClientSecret,
		TokenURL:     fmt.Sprintf("%s%s/oauth2/token", aadEndpoint, cfg.AzureTenantID),
		EndpointParams: url.Values{
			"resource": []string{resource},
		},
	}).TokenSource(ctx)
}

// TODO: We need to expose which authentication mechanism is used to Terraform,
// as we cannot create AKV backed secret scopes when authenticated as SP.
// If we are authenticated as SP and wish to create one we want to fail early.
// Also see https://github.com/databricks/terraform-provider-databricks/issues/1490.
func (c AzureClientSecretCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.AzureClientID == "" || cfg.AzureClientSecret == "" || cfg.AzureTenantID == "" {
		return nil, nil
	}
	if !cfg.IsAzure() {
		return nil, nil
	}
	err := cfg.loadAzureTenantId(ctx)
	if err != nil {
		return nil, fmt.Errorf("load tenant id: %w", err)
	}
	err = cfg.azureEnsureWorkspaceUrl(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("resolve host: %w", err)
	}
	logger.Infof(ctx, "Generating AAD token for Service Principal (%s)", cfg.AzureClientID)
	env := cfg.Environment()
	aadEndpoint := env.AzureActiveDirectoryEndpoint()
	managementEndpoint := env.AzureServiceManagementEndpoint()
	inner := azureReuseTokenSource(nil, c.tokenSourceFor(ctx, cfg, aadEndpoint, env.AzureApplicationID))
	management := azureReuseTokenSource(nil, c.tokenSourceFor(ctx, cfg, aadEndpoint, managementEndpoint))
	visitor := azureVisitor(cfg, serviceToServiceVisitor(inner, management, xDatabricksAzureSpManagementToken))
	return credentials.NewOAuthCredentialsProvider(visitor, inner.Token), nil
}
