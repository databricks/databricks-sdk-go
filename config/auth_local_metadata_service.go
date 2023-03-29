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

// Credentials provider that fetches a token from a locally running HTTP server
//
// The credentials provider will perform a GET request to the configured URL.
//
// The header "Metadata: true" will be added to the request, which must be
// verified by server to prevent SSRF attacks
//
// The server is expected to return a JSON response with the following fields:
//
//	host: URL of the Databricks host to connect to.
//	access_token: The requested access token.
//	token_type: The type of token, which is a "Bearer" access token.
//	expires_on: The timespan when the access token expires.
type LocalMetadataServiceCredentials struct{}

func (c LocalMetadataServiceCredentials) Name() string {
	return "local-metadata-service"
}

func (c LocalMetadataServiceCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.LocalMetadataServiceUrl == "" {
		return nil, nil
	}

	parsedMetadataServiceUrl, err := url.Parse(cfg.LocalMetadataServiceUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid auth server URL: %w", err)
	}

	// only allow localhost URLs
	if parsedMetadataServiceUrl.Hostname() != "localhost" && parsedMetadataServiceUrl.Hostname() != "127.0.0.1" {
		return nil, fmt.Errorf("invalid auth server URL: %s", cfg.LocalMetadataServiceUrl)
	}

	resp, err := makeRequest(parsedMetadataServiceUrl)
	if err != nil {
		return nil, nil
	}

	if resp == nil {
		return nil, nil
	}

	if resp.Host == "" {
		return nil, nil
	}

	if cfg.Host != "" && cfg.Host != resp.Host {
		return nil, nil
	}

	cfg.Host = resp.Host

	ts := metadataServiceTokenSource{
		metadataServiceURL: parsedMetadataServiceUrl,
		databricksHost:     cfg.Host,
	}
	return refreshableVisitor(&ts), nil
}

// makes a request to the server and returns the token
func makeRequest(serverUrl *url.URL) (*serverResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), metadataServiceTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}
	req.Header.Add("Metadata", "true")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("token response: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return nil, nil
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("token read: %w", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("token error: %s", body)
	}

	var response serverResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("token parse: %w", err)
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
	metadataServiceURL *url.URL
	databricksHost     string
}

func (t metadataServiceTokenSource) Token() (*oauth2.Token, error) {
	token, err := makeRequest(t.metadataServiceURL)

	if err != nil {
		return nil, err
	}

	if token.Host != t.databricksHost {
		return nil, fmt.Errorf("host in token (%s) doesn't match configured host (%s)", token.Host, t.databricksHost)
	}

	epoch, err := token.ExpiresOn.Int64()
	if err != nil {
		return nil, fmt.Errorf("token expires on: %w", err)
	}

	expiry := time.Unix(epoch, 0)

	logger.Infof(context.Background(), "Refreshed access token for %s from local credentials server, which expires on %s",
		t.metadataServiceURL.String(), expiry)

	return &oauth2.Token{
		AccessToken: token.AccessToken,
		TokenType:   token.TokenType,
		Expiry:      expiry,
	}, nil
}

type serverResponse struct {
	// URL of the Databricks host to connect to.
	Host string `json:"host"`

	// The requested access token. When you call a secured REST API,
	// the token is embedded in the Authorization request header field as
	// a "bearer" token, allowing the API to authenticate the caller.
	AccessToken string `json:"access_token"`

	// The type of token, which is a "Bearer" access token, which means the
	// resource can give access to the bearer of this token.
	TokenType string `json:"token_type"`

	// The timespan when the access token expires.
	// The date is represented as the number of seconds from "1970-01-01T0:0:0Z UTC"
	// If zero the token is valid forever.
	ExpiresOn json.Number `json:"expires_on"`
}
