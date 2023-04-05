package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

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

func (c AzureMsiCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if !cfg.IsAzure() || !cfg.AzureUseMSI || cfg.AzureResourceID == "" {
		return nil, nil
	}
	env, err := c.getInstanceEnvironment(ctx)
	if err != nil {
		return nil, err
	}
	err = cfg.azureEnsureWorkspaceUrl(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("resolve host: %w", err)
	}
	logger.Debugf(ctx, "Generating AAD token via Azure MSI")
	inner := azureMsiTokenSource{
		resource: cfg.getAzureLoginAppID(),
		clientId: cfg.AzureClientID,
	}
	platform := azureMsiTokenSource{
		resource: env.ServiceManagementEndpoint,
		clientId: cfg.AzureClientID,
	}
	return func(r *http.Request) error {
		r.Header.Set("X-Databricks-Azure-Workspace-Resource-Id", cfg.AzureResourceID)
		return serviceToServiceVisitor(inner, platform,
			"X-Databricks-Azure-SP-Management-Token")(r)
	}, nil
}

func (c AzureMsiCredentials) getInstanceEnvironment(ctx context.Context) (*azureEnvironment, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/instance", instanceMetadataPrefix), nil)
	if err != nil {
		return nil, fmt.Errorf("metadata request: %w", err)
	}
	query := req.URL.Query()
	query.Add("api-version", "2021-12-13")
	req.URL.RawQuery = query.Encode()
	req.Header.Add("Metadata", "true")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("metadata response: %w", err)
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("metadata read: %w", err)
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("metadata error: %s", raw)
	}
	var metadata struct {
		Compute struct {
			Environment string `json:"azEnvironment"`
		} `json:"compute"`
	}
	err = json.Unmarshal(raw, &metadata)
	if err != nil {
		return nil, fmt.Errorf("metadata parse: %w", err)
	}
	for _, v := range azureEnvironments {
		if v.Name == metadata.Compute.Environment {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("cannot determine environment: %s",
		metadata.Compute.Environment)
}

// implementing azureHostResolver for ensureWorkspaceUrl to work
func (c AzureMsiCredentials) tokenSourceFor(_ context.Context, cfg *Config, _ azureEnvironment, resource string) oauth2.TokenSource {
	return azureMsiTokenSource{
		resource: resource,
		clientId: cfg.AzureClientID,
	}
}

type azureMsiTokenSource struct {
	resource string
	clientId string
}

func (s azureMsiTokenSource) Token() (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), azureMsiTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		fmt.Sprintf("%s/identity/oauth2/token", instanceMetadataPrefix), nil)
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}
	query := req.URL.Query()
	query.Add("api-version", "2018-02-01")
	query.Add("resource", s.resource)
	if s.clientId != "" {
		query.Add("client_id", s.clientId)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Add("Metadata", "true")
	return makeMsiRequest(req)
}

func makeMsiRequest(req *http.Request) (*oauth2.Token, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token response: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("token read: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token error: %s", raw)
	}
	var token azureMsiToken
	err = json.Unmarshal(raw, &token)
	if err != nil {
		return nil, fmt.Errorf("token parse: %w", err)
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("token parse: invalid token")
	}
	epoch, err := token.ExpiresOn.Int64()
	if err != nil {
		return nil, fmt.Errorf("token expires on: %w", err)
	}
	return &oauth2.Token{
		TokenType:   token.TokenType,
		AccessToken: token.AccessToken,
		Expiry:      time.Unix(epoch, 0),
	}, nil
}

type azureMsiToken struct {
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"token_type"`
	ExpiresOn   json.Number `json:"expires_on"`
}
