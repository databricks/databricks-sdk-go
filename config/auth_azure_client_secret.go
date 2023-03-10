package config

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/logger"
)

type AzureClientSecretCredentials struct {
}

func (c AzureClientSecretCredentials) Name() string {
	return "azure-client-secret"
}

func (c AzureClientSecretCredentials) tokenSourceFor(
	ctx context.Context, cfg *Config, env azureEnvironment, resource string) oauth2.TokenSource {
	return (&clientcredentials.Config{
		ClientID:     cfg.AzureClientID,
		ClientSecret: cfg.AzureClientSecret,
		TokenURL:     fmt.Sprintf("%s%s/oauth2/token", env.ActiveDirectoryEndpoint, cfg.AzureTenantID),
		EndpointParams: url.Values{
			"resource": []string{resource},
		},
	}).TokenSource(ctx)
}

// TODO: We need to expose which authentication mechanism is used to Terraform,
// as we cannot create AKV backed secret scopes when authenticated as SP.
// If we are authenticated as SP and wish to create one we want to fail early.
// Also see https://github.com/databricks/terraform-provider-databricks/issues/1490.
func (c AzureClientSecretCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.AzureClientID == "" || cfg.AzureClientSecret == "" || cfg.AzureTenantID == "" {
		return nil, nil
	}
	if !cfg.IsAzure() {
		return nil, nil
	}
	env, err := cfg.GetAzureEnvironment()
	if err != nil {
		return nil, err
	}
	err = cfg.azureEnsureWorkspaceUrl(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("resolve host: %w", err)
	}
	logger.Infof(ctx, "Generating AAD token for Service Principal (%s)", cfg.AzureClientID)
	refreshCtx := context.Background()
	inner := c.tokenSourceFor(refreshCtx, cfg, env, cfg.getAzureLoginAppID())
	platform := c.tokenSourceFor(refreshCtx, cfg, env, env.ServiceManagementEndpoint)
	return func(r *http.Request) error {
		if cfg.AzureResourceID != "" {
			r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", cfg.AzureResourceID)
		}
		return serviceToServiceVisitor(inner, platform,
			"X-Databricks-Azure-SP-Management-Token")(r)
	}, nil
}
