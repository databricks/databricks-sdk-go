package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"github.com/databricks/databricks-sdk-go/logger"
)

// The header used to pass the service management token to the Databricks backend.
const xDatabricksAzureSpManagementToken = "X-Databricks-Azure-SP-Management-Token"

// The header used to pass the workspace resource ID to the Databricks backend.
const xDatabricksAzureWorkspaceResourceId = "X-Databricks-Azure-Workspace-Resource-Id"

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

// There are three scenarios:
//
//  1. The user has logged in with the Azure CLI as a user and has access to the service management endpoint.
//  2. The user has logged in with the Azure CLI as a user and does not have access to the service management endpoint.
//  3. The user has logged in with the Azure CLI as a service principal, and must have access to the service management
//     endpoint to authenticate.
//
// If the user can't access the service management endpoint, we assume they are in case 2 and do not pass the service
// management token. Otherwise, we always pass the service management token.
func (c AzureCliCredentials) getVisitor(ctx context.Context, cfg *Config, innerTokenSource oauth2.TokenSource) (func(*http.Request) error, error) {
	env, err := cfg.GetAzureEnvironment()
	if err != nil {
		return nil, err
	}
	managementTs := &azureCliTokenSource{env.ServiceManagementEndpoint, ""}
	_, err = managementTs.Token()
	if err != nil {
		logger.Debugf(ctx, "Not including service management token in headers: %v", err)
		return azureVisitor(cfg, refreshableVisitor(innerTokenSource)), nil
	}
	return azureVisitor(cfg, serviceToServiceVisitor(innerTokenSource, managementTs, xDatabricksAzureSpManagementToken)), nil
}

func (c AzureCliCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if !cfg.IsAzure() {
		return nil, nil
	}
	// Eagerly get a token to fail fast in case the user is not logged in with the Azure CLI.
	ts := &azureCliTokenSource{cfg.getAzureLoginAppID(), cfg.AzureResourceID}
	_, err := ts.Token()
	if err != nil {
		if strings.Contains(err.Error(), "No subscription found") {
			// auth is not configured
			return nil, nil
		}
		if strings.Contains(err.Error(), "executable file not found") {
			logger.Debugf(ctx, "Most likely Azure CLI is not installed. "+
				"See https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest for details")
			return nil, nil
		}
		return nil, err
	}
	err = cfg.azureEnsureWorkspaceUrl(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("resolve host: %w", err)
	}
	logger.Infof(ctx, "Using Azure CLI authentication with AAD tokens")
	return c.getVisitor(ctx, cfg, ts)
}

type azureCliTokenSource struct {
	resource            string
	workspaceResourceId string
}

type internalCliToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	TokenType    string `json:"tokenType"`
	ExpiresOn    string `json:"expiresOn"`
}

func (ts *azureCliTokenSource) getSubscription() string {
	if ts.workspaceResourceId == "" {
		return ""
	}
	components := strings.Split(ts.workspaceResourceId, "/")
	if len(components) < 3 {
		logger.Warnf(context.Background(), "Invalid azure workspace resource ID")
		return ""
	}
	return components[2]
}

func (ts *azureCliTokenSource) Token() (*oauth2.Token, error) {
	tokenBytes, err := ts.getTokenBytes()
	if err != nil {
		return nil, err
	}
	var it internalCliToken
	err = json.Unmarshal(tokenBytes, &it)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal CLI result: %w", err)
	}
	expiresOn, err := time.ParseInLocation("2006-01-02 15:04:05.999999", it.ExpiresOn, time.Local)
	if err != nil {
		return nil, fmt.Errorf("cannot parse expiry: %w", err)
	}
	logger.Infof(context.Background(), "Refreshed OAuth token for %s from Azure CLI, which expires on %s",
		ts.resource, it.ExpiresOn)

	var extra map[string]interface{}
	err = json.Unmarshal(tokenBytes, &extra)
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

func (ts *azureCliTokenSource) getTokenBytes() ([]byte, error) {
	subscription := ts.getSubscription()
	args := []string{"account", "get-access-token", "--resource",
		ts.resource, "--output", "json"}
	if subscription != "" {
		extendedArgs := make([]string, len(args))
		copy(extendedArgs, args)
		extendedArgs = append(extendedArgs, "--subscription", subscription)
		// This will fail if the user has access to the workspace, but not to the subscription
		// itself.
		// In such case, we fall back to not using the subscription.
		result, err := exec.Command("az", extendedArgs...).Output()
		if err == nil {
			return result, nil
		}
		logger.Warnf(context.Background(), "Failed to get token for subscription. Using resource only token.")
	}
	result, err := exec.Command("az", args...).Output()
	if ee, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("cannot get access token: %s", string(ee.Stderr))
	}
	if err != nil {
		return nil, fmt.Errorf("cannot get access token: %w", err)
	}
	return result, nil
}
