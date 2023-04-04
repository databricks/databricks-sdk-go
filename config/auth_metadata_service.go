package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

// timeout to wait for metadata service to respond
const metadataServiceTimeout = 10 * time.Second

const MetadataServiceVersion = "1"
const MetadataServiceVersionHeader = "X-Databricks-Metadata-Version"
const MetadataServiceHostHeader = "X-Databricks-Host"

// response expected from the metadata service
type serverResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`

	// epoch in seconds since 1970-01-01T00:00:00Z
	ExpiresOn json.Number `json:"expires_on"`
}

// Credentials provider that fetches a token from a locally running HTTP server
//
// The credentials provider will perform a GET request to the configured URL.
//
// The MUST return 4xx response if the "X-Databricks-Metadata-Version" header
// is not set or set to a version that the server doesn't support.
//
// The server MUST guarantee stable sessions per URL path. That is, if the
// server returns a token for a Host on a given URL path, it MUST continue to return
// tokens for the same Host.
//
// The server MUST return a 4xx response if the Host passed in the "X-Databricks-Host"
// header doesn't match the token.
//
// The server is expected to return a JSON response with the following fields:
//
// - access_token: The requested access token.
// - token_type: The type of token, which is a "Bearer" access token.
// - expires_on: Unix timestamp when the access token expires.
type MetadataServiceCredentials struct{}

func (c MetadataServiceCredentials) Name() string {
	return "metadata-service"
}

func (c MetadataServiceCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.MetadataServiceURL == "" || cfg.Host == "" {
		return nil, nil
	}

	parsedMetadataServiceURL, err := url.Parse(cfg.MetadataServiceURL)
	if err != nil {
		return nil, fmt.Errorf("invalid auth server URL: %w", err)
	}

	// only allow localhost URLs
	if parsedMetadataServiceURL.Hostname() != "localhost" && parsedMetadataServiceURL.Hostname() != "127.0.0.1" {
		return nil, fmt.Errorf("invalid auth server URL: %s", cfg.MetadataServiceURL)
	}

	ms := metadataService{
		metadataServiceURL: parsedMetadataServiceURL,
		config:             cfg,
	}

	response, err := ms.Get()
	if err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	ts := metadataServiceTokenSource{
		metadataService: &ms,
		databricksHost:  cfg.Host,
	}
	return refreshableVisitor(&ts), nil
}

type metadataService struct {
	metadataServiceURL *url.URL
	config             *Config
}

// performs a request to the metadata service and returns the token
func (s metadataService) Get() (*serverResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), metadataServiceTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.metadataServiceURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}

	req.Header.Add(MetadataServiceVersionHeader, MetadataServiceVersion)
	req.Header.Add(MetadataServiceHostHeader, s.config.Host)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token response: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("token read: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token error: %s", body)
	}

	var response serverResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("token parse: %w", err)
	}

	if response.AccessToken == "" {
		return nil, fmt.Errorf("token parse: invalid token")
	}

	return &response, nil
}

type metadataServiceTokenSource struct {
	metadataService *metadataService
	databricksHost  string
}

func (t metadataServiceTokenSource) Token() (*oauth2.Token, error) {
	token, err := t.metadataService.Get()

	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, fmt.Errorf("no token returned from metadata service")
	}

	epoch, err := token.ExpiresOn.Int64()
	if err != nil {
		return nil, fmt.Errorf("token expires on: %w", err)
	}

	expiry := time.Unix(epoch, 0)

	logger.Infof(context.Background(), "Refreshed access token from local metadata service, which expires on %s", expiry.Format(time.RFC3339))

	return &oauth2.Token{
		TokenType:   token.TokenType,
		AccessToken: token.AccessToken,
		Expiry:      expiry,
	}, nil
}
