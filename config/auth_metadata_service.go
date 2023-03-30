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

// timeout to wait for local auth server to respond
const metadataServiceTimeout = 10 * time.Second

// response expected from the metadata service
type serverResponse struct {
	// URL of the metadata service to connect to.
	Host string `json:"host"`

	// The requested access token.
	Token oauth2.Token `json:"token"`
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
// The server is expected to return a JSON response with the following fields:
//
//	host: URL of the Databricks host to connect to.
//	token: object with the following fields
//    access_token: The requested access token.
//	  token_type: The type of token, which is a "Bearer" access token.
//	  expires_on: The timespan when the access token expires.

type MetadataServiceCredentials struct{}

func (c MetadataServiceCredentials) Name() string {
	return "local-metadata-service"
}

func (c MetadataServiceCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.MetadataServiceURL == "" {
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

	ms := metadataService{parsedMetadataServiceURL}

	resp, err := ms.Get()
	if err != nil {
		return nil, nil
	}

	if resp == nil {
		return nil, nil
	}

	if cfg.Host != "" && cfg.Host != resp.Host {
		return nil, nil
	}

	cfg.Host = resp.Host

	ts := metadataServiceTokenSource{
		metadataService: &ms,
		databricksHost:  cfg.Host,
	}
	return refreshableVisitor(&ts), nil
}

type metadataService struct {
	*url.URL
}

// performs a request to the metadata service and returns the token
func (s metadataService) Get() (*serverResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), metadataServiceTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}

	req.Header.Add("X-Databricks-Metadata-Version", "1")

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

	if !response.Token.Valid() || response.Host == "" {
		return nil, fmt.Errorf("token parse: invalid token")
	}

	// parse and normalize url
	parsedHost, err := url.Parse(response.Host)
	if err != nil {
		return nil, fmt.Errorf("token parse: Invalid URL %s", response.Host)
	}

	response.Host = parsedHost.String()

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

	logger.Infof(context.Background(), "Refreshed access token from local credentials server, which expires on %s", token.Token.Expiry)

	return &token.Token, nil
}
