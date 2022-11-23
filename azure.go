package databricks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os/exec"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/internal"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
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
	logger.Infof("Discovered workspace url: %s", c.Host)
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

type AzureCliCredentials struct {
}

func (c AzureCliCredentials) Name() string {
	return "azure-cli"
}

// implementing azureHostResolver for ensureWorkspaceUrl to work
func (c AzureCliCredentials) tokenSourceFor(
	ctx context.Context, cfg *Config, env azureEnvironment, resource string) oauth2.TokenSource {
	return &azureCliTokenSource{resource: resource}
}

func (c AzureCliCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if !cfg.IsAzure() {
		return nil, nil
	}
	ts := azureCliTokenSource{cfg.getAzureLoginAppID()}
	_, err := ts.Token()
	if err != nil {
		if strings.Contains(err.Error(), "No subscription found") {
			// auth is not configured
			return nil, nil
		}
		if strings.Contains(err.Error(), "executable file not found") {
			logger.Debugf("Most likely Azure CLI is not installed. " +
				"See https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest for details")
			return nil, nil
		}
		return nil, err
	}
	err = cfg.azureEnsureWorkspaceUrl(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("resolve host: %w", err)
	}
	logger.Infof("Using Azure CLI authentication with AAD tokens")
	return internal.RefreshableVisitor(&ts), nil
}

type azureCliTokenSource struct {
	resource string
}

type internalCliToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	ExpiresOn    string `json:"expiresOn"`
}

func (ts *azureCliTokenSource) Token() (*oauth2.Token, error) {
	out, err := exec.Command("az", "account", "get-access-token", "--resource",
		ts.resource, "--output", "json").Output()
	if ee, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("cannot get access token: %s", string(ee.Stderr))
	}
	if err != nil {
		return nil, fmt.Errorf("cannot get access token: %v", err)
	}
	var it internalCliToken
	err = json.Unmarshal(out, &it)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal CLI result: %w", err)
	}
	expiresOn, err := time.ParseInLocation("2006-01-02 15:04:05.999999", it.ExpiresOn, time.Local)
	if err != nil {
		return nil, fmt.Errorf("cannot parse expiry: %w", err)
	}
	logger.Infof("Refreshed OAuth token for %s from Azure CLI, which expires on %s",
		ts.resource, it.ExpiresOn)

	var extra map[string]interface{}
	err = json.Unmarshal(out, &extra)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal extra: %w", err)
	}
	return (&oauth2.Token{
		AccessToken:  it.AccessToken,
		RefreshToken: it.RefreshToken,
		TokenType:    it.TokenType,
		Expiry:       expiresOn,
	}).WithExtra(extra), nil
}

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
	refreshCtx := context.Background()
	inner := c.tokenSourceFor(refreshCtx, cfg, env, cfg.getAzureLoginAppID())
	platform := c.tokenSourceFor(refreshCtx, cfg, env, env.ServiceManagementEndpoint)
	return func(r *http.Request) error {
		r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", cfg.AzureResourceID)
		return internal.ServiceToServiceVisitor(inner, platform,
			"X-Databricks-Azure-SP-Management-Token")(r)
	}, nil
}
