package databricks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/oauth2"
	
	"github.com/databricks/sdk-go/databricks/internal"
)

// List of management information
const armDatabricksResourceID string = "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d"

type AzureCliCredentials struct{
}

func (c AzureCliCredentials) Name() string {
	return "azure-cli"
}

func (c AzureCliCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if !cfg.IsAzure() {
		return nil, nil
	}
	ts := azureCliTokenSource{armDatabricksResourceID}
	_, err := ts.Token()
	if err != nil {
		if strings.Contains(err.Error(), "No subscription found") {
			// auth is not configured
			return nil, nil
		}
		if strings.Contains(err.Error(), "executable file not found") {
			// , fmt.Errorf("most likely Azure CLI is not installed. " +
			// 	"See https://docs.microsoft.com/en-us/cli/azure/?view=azure-cli-latest for details")
			return nil, nil
		}
		return nil, err
	}
	log.Printf("[INFO] Using Azure CLI authentication with AAD tokens")
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
	log.Printf("[INFO] Refreshed OAuth token for %s from Azure CLI, which expires on %s",
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
