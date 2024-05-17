package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

var errInvalidToken = errors.New("invalid token")
var errInvalidTokenExpiry = errors.New("invalid token expiry")

// well-known URL for Azure Instance Metadata Service (IMDS)
// https://learn.microsoft.com/en-us/azure-stack/user/instance-metadata-service
var instanceMetadataPrefix = "http://169.254.169.254/metadata"

// timeout to wait for IMDS response
const azureMsiTimeout = 10 * time.Second

type AzureMsiCredentials struct {
}

func (c AzureMsiCredentials) Name() string {
	return "azure-msi"
}

func (c AzureMsiCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if !cfg.IsAzure() || !cfg.AzureUseMSI || (cfg.AzureResourceID == "" && !cfg.IsAccountClient()) {
		return nil, nil
	}
	env := cfg.Environment()
	if !cfg.IsAccountClient() {
		err := cfg.azureEnsureWorkspaceUrl(ctx, c)
		if err != nil {
			return nil, fmt.Errorf("resolve host: %w", err)
		}
	}
	logger.Debugf(ctx, "Generating AAD token via Azure MSI")
	inner := azureReuseTokenSource(nil, c.tokenSourceFor(ctx, cfg, "", env.AzureApplicationID))
	management := azureReuseTokenSource(nil, c.tokenSourceFor(ctx, cfg, "", env.AzureServiceManagementEndpoint()))
	visitor := azureVisitor(cfg, serviceToServiceVisitor(inner, management, xDatabricksAzureSpManagementToken))
	return credentials.NewOAuthCredentialsProvider(visitor, inner.Token), nil
}

// implementing azureHostResolver for ensureWorkspaceUrl to work
func (c AzureMsiCredentials) tokenSourceFor(_ context.Context, cfg *Config, _, resource string) oauth2.TokenSource {
	return NewAzureMsiTokenSource(cfg.refreshClient, resource, cfg.AzureClientID)
}

// NewAzureMsiTokenSource returns [oauth2.TokenSource] for a passwordless authentication via Azure Managed identity
func NewAzureMsiTokenSource(client *httpclient.ApiClient, resource, clientId string) oauth2.TokenSource {
	return &azureMsiTokenSource{
		client:   client,
		resource: resource,
		clientId: clientId,
	}
}

type azureMsiTokenSource struct {
	client   *httpclient.ApiClient
	resource string
	clientId string
}

func (s azureMsiTokenSource) Token() (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), azureMsiTimeout)
	defer cancel()
	query := map[string]string{
		"api-version": "2018-02-01",
		"resource":    s.resource,
	}
	if s.clientId != "" {
		query["client_id"] = s.clientId
	}
	var inner msiToken
	err := s.client.Do(ctx, http.MethodGet,
		fmt.Sprintf("%s/identity/oauth2/token", instanceMetadataPrefix),
		httpclient.WithRequestHeader("Metadata", "true"),
		httpclient.WithRequestData(query),
		httpclient.WithResponseUnmarshal(&inner),
	)
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}
	return inner.Token()
}

type msiToken struct {
	TokenType    string      `json:"token_type"`
	AccessToken  string      `json:"access_token,omitempty"`
	RefreshToken string      `json:"refresh_token,omitempty"`
	ExpiresOn    json.Number `json:"expires_on"`
}

func (token msiToken) Token() (*oauth2.Token, error) {
	if token.AccessToken == "" {
		return nil, fmt.Errorf("token parse: %w", errInvalidToken)
	}
	epoch, err := token.ExpiresOn.Int64()
	if err != nil {
		// go 1.19 doesn't support multiple error unwraps
		return nil, fmt.Errorf("%w: %s", errInvalidTokenExpiry, err)
	}
	return &oauth2.Token{
		TokenType:    token.TokenType,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       time.Unix(epoch, 0),
	}, nil
}
