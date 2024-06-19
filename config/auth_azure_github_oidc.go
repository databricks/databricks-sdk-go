package config

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"golang.org/x/oauth2"
)

// AzureGithubOIDCCredentials provides credentials for GitHub Actions that use
// an Azure Active Directory Federated Identity to authenticate with Azure.
type AzureGithubOIDCCredentials struct{}

// Name implements [CredentialsStrategy.Name].
func (c AzureGithubOIDCCredentials) Name() string {
	return "github-oidc-azure"
}

// Configure implements [CredentialsStrategy.Configure].
func (c AzureGithubOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	// Sanity check that the config is configured for Azure Databricks.
	if !cfg.IsAzure() || cfg.AzureClientID == "" || cfg.Host == "" || cfg.AzureTenantID == "" {
		return nil, nil
	}

	idToken, err := requestIDToken(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if idToken == "" {
		return nil, nil
	}

	ts := &azureOICDCredentialsProvider{
		aadEndpoint:   fmt.Sprintf("%s%s", cfg.DatabricksEnvironment.AzureActiveDirectoryEndpoint(), cfg.AzureTenantID),
		clientID:      cfg.AzureClientID,
		applicationID: cfg.DatabricksEnvironment.AzureApplicationID,
		idToken:       idToken,
		httpClient:    cfg.refreshClient,
	}

	return credentials.NewOAuthCredentialsProvider(refreshableVisitor(ts), ts.Token), nil
}

// requestIDToken requests an ID token from the Github Action.
func requestIDToken(ctx context.Context, cfg *Config) (string, error) {
	tokenRequestURL := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	if tokenRequestURL == "" {
		return "", nil
	}
	token := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	if token == "" {
		return "", nil
	}

	resp := struct { // anonymous struct to parse the response
		Value string `json:"value"`
	}{}
	err := cfg.refreshClient.Do(ctx, "GET", tokenRequestURL,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", token)),
		httpclient.WithResponseUnmarshal(&resp),
		httpclient.WithRequestData(map[string]string{
			"audience": "api://AzureADTokenExchange",
		}),
	)
	if err != nil {
		return "", err
	}

	return resp.Value, nil
}

// azureOIDCTokenSource implements [oauth2.TokenSource] to obtain Azure auth
// tokens from an ID token.
type azureOICDCredentialsProvider struct {
	aadEndpoint   string
	clientID      string
	applicationID string
	idToken       string

	httpClient *httpclient.ApiClient
}

const azureOICDTimeout = 10 * time.Second

func (cp *azureOICDCredentialsProvider) Token() (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), azureOICDTimeout)
	defer cancel()

	resp := struct { // anonymous struct to parse the response
		TokenType    string      `json:"token_type"`
		AccessToken  string      `json:"access_token"`
		RefreshToken string      `json:"refresh_token"`
		ExpiresOn    json.Number `json:"expires_on"`
	}{}
	err := cp.httpClient.Do(ctx, "POST", cp.aadEndpoint,
		httpclient.WithRequestData(map[string]string{
			"grant_type": "client_credentials",
			"resource":   cp.applicationID,
			"client_id":  cp.clientID,
			// Use the ID token instead of a client_secret.
			"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
			"client_assertion":      cp.idToken,
		}),
		httpclient.WithResponseUnmarshal(&resp),
	)
	if err != nil {
		return nil, err
	}

	epochInSec, err := resp.ExpiresOn.Int64()
	if err != nil {
		return nil, fmt.Errorf("invalid token: cannot parse token expiry: %w", err)
	}
	return &oauth2.Token{
		TokenType:    resp.TokenType,
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		Expiry:       time.Unix(epochInSec, 0),
	}, nil
}
