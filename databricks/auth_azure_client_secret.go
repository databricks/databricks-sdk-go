package databricks

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/databricks/databricks-sdk-go/databricks/internal"
	"github.com/databricks/databricks-sdk-go/databricks/logger"
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

func (c AzureClientSecretCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	// TODO: this has to be exposed for Terraform, as we cannot create AKV scopes by SPNs
	if cfg.AzureClientID == "" || cfg.AzureClientSecret == "" || cfg.AzureTenantID == "" || cfg.AzureResourceID == "" {
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
	logger.Infof("Generating AAD token for Service Principal (%s)", cfg.AzureClientID)
	return func(r *http.Request) error {
		ctx := r.Context()
		r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", cfg.AzureResourceID)
		return internal.ServiceToServiceVisitor(
			c.tokenSourceFor(ctx, cfg, env, armDatabricksResourceID),
			c.tokenSourceFor(ctx, cfg, env, env.ServiceManagementEndpoint),
			"X-Databricks-Azure-SP-Management-Token")(r)
	}, nil
}
