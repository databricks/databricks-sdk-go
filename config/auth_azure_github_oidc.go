package config

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
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

	ts := &azureOIDCTokenSource{
		aadEndpoint:   fmt.Sprintf("%s%s/oauth2/token", cfg.Environment().AzureActiveDirectoryEndpoint(), cfg.AzureTenantID),
		clientID:      cfg.AzureClientID,
		applicationID: cfg.Environment().AzureApplicationID,
		idToken:       idToken,
		httpClient:    cfg.refreshClient,
	}

	return credentials.NewOAuthCredentialsProvider(refreshableVisitor(ts), ts.Token), nil
}

// requestIDToken requests an ID token from the Github Action.
func requestIDToken(ctx context.Context, cfg *Config) (string, error) {
	if cfg.ActionsIDTokenRequestURL == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestURL, likely not calling from a Github action")
		return "", nil
	}
	if cfg.ActionsIDTokenRequestToken == "" {
		logger.Debugf(ctx, "Missing cfg.ActionsIDTokenRequestToken, likely not calling from a Github action")
		return "", nil
	}

	resp := struct { // anonymous struct to parse the response
		Value string `json:"value"`
	}{}
	err := cfg.refreshClient.Do(ctx, "GET", fmt.Sprintf("%s&audience=api://AzureADTokenExchange", cfg.ActionsIDTokenRequestURL),
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.ActionsIDTokenRequestToken)),
		httpclient.WithResponseUnmarshal(&resp),
	)
	if err != nil {
		return "", fmt.Errorf("failed to request ID token from %s: %w", cfg.ActionsIDTokenRequestURL, err)
	}

	return resp.Value, nil
}

// azureOIDCTokenSource implements [oauth2.TokenSource] to obtain Azure auth
// tokens from an ID token.
type azureOIDCTokenSource struct {
	aadEndpoint   string
	clientID      string
	applicationID string
	idToken       string

	httpClient *httpclient.ApiClient
}

const azureOICDTimeout = 10 * time.Second

func (ts *azureOIDCTokenSource) Token() (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), azureOICDTimeout)
	defer cancel()

	resp := struct { // anonymous struct to parse the response
		TokenType    string      `json:"token_type"`
		AccessToken  string      `json:"access_token"`
		RefreshToken string      `json:"refresh_token"`
		ExpiresOn    json.Number `json:"expires_on"`
	}{}
	err := ts.httpClient.Do(ctx, "POST", ts.aadEndpoint,
		httpclient.WithUrlEncodedData(map[string]string{
			"grant_type": "client_credentials",
			"resource":   ts.applicationID,
			"client_id":  ts.clientID,
			// Use the ID token instead of a client_secret.
			"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
			"client_assertion":      ts.idToken,
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
