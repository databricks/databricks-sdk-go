package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func New(cfg *config.Config) (*DatabricksClient, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	clientCfg, err := config.HTTPClientConfigFromConfig(cfg)
	if err != nil {
		return nil, err
	}
	return &DatabricksClient{
		Config: cfg,
		client: httpclient.NewApiClient(clientCfg),
	}, nil
}

func NewWithClient(cfg *config.Config, client *httpclient.ApiClient) (*DatabricksClient, error) {
	err := cfg.EnsureResolved()
	if err != nil {
		return nil, err
	}
	return &DatabricksClient{
		Config: cfg,
		client: client,
	}, nil
}

type DatabricksClient struct {
	Config *config.Config
	client *httpclient.ApiClient
}

// ConfiguredAccountID returns Databricks Account ID if it's provided in config,
// empty string otherwise
func (c *DatabricksClient) ConfiguredAccountID() string {
	return c.Config.AccountID
}

// ApiClient returns the inner Api Client.
func (c *DatabricksClient) ApiClient() *httpclient.ApiClient {
	return c.client
}

// GetOAuthToken returns a new OAuth token using the provided token. The token must be a JWT token.
// The resulting token is scoped to the authorization details provided.
//
// **NOTE:** Experimental: This API may change or be removed in a future release
// without warning.
func (c *DatabricksClient) GetOAuthToken(ctx context.Context, authDetails string, token *oauth2.Token) (*oauth2.Token, error) {
	return c.client.GetOAuthToken(ctx, authDetails, token)
}

// Do sends an HTTP request against path.
func (c *DatabricksClient) Do(ctx context.Context, method, path string,
	headers map[string]string, queryParams map[string]any, request, response any,
	visitors ...func(*http.Request) error) error {
	opts := []httpclient.DoOption{}
	for _, v := range visitors {
		opts = append(opts, httpclient.WithRequestVisitor(v))
	}
	opts = append(opts, httpclient.WithQueryParameters(queryParams))
	opts = append(opts, httpclient.WithRequestHeaders(headers))
	opts = append(opts, httpclient.WithRequestData(request))
	opts = append(opts, httpclient.WithResponseUnmarshal(response))
	// Remove extra `/` from path for files API.
	// Once the OpenAPI spec doesn't include the extra slash, we can remove this
	if strings.HasPrefix(path, "/api/2.0/fs/files//") {
		path = strings.Replace(path, "/api/2.0/fs/files//", "/api/2.0/fs/files/", 1)
	}
	return c.client.Do(ctx, method, path, opts...)
}

// DoProto sends an HTTP request with proto message marshalling/unmarshalling.
// This is a completely independent implementation that doesn't use the Do function.
func (c *DatabricksClient) DoProto(ctx context.Context, method, path string,
	headers map[string]string, queryParams map[string]any,
	request proto.Message, response proto.Message) error {

	// Build the full URL
	baseURL := strings.TrimRight(c.Config.Host, "/")
	fullURL := baseURL + path

	// Add query parameters
	if len(queryParams) > 0 {
		parsedURL, err := url.Parse(fullURL)
		if err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}

		values := parsedURL.Query()
		for key, value := range queryParams {
			values.Set(key, fmt.Sprintf("%v", value))
		}
		parsedURL.RawQuery = values.Encode()
		fullURL = parsedURL.String()
	}

	// Prepare request body
	var body io.Reader
	if request != nil {
		jsonData, err := protojson.Marshal(request)
		if err != nil {
			return fmt.Errorf("failed to marshal proto request: %w", err)
		}
		body = bytes.NewReader(jsonData)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	// Set default headers
	if httpReq.Body != nil {
		httpReq.Header.Set("Content-Type", "application/json")
	}
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("User-Agent", "databricks-sdk-go-proto/1.0.0")

	// Add custom headers
	for key, value := range headers {
		httpReq.Header.Set(key, value)
	}

	// Apply authentication
	if c.client != nil {
		// Get the client configuration to access auth visitor
		clientCfg, err := config.HTTPClientConfigFromConfig(c.Config)
		if err != nil {
			return fmt.Errorf("failed to get client config: %w", err)
		}

		if clientCfg.AuthVisitor != nil {
			if err := clientCfg.AuthVisitor(httpReq); err != nil {
				return fmt.Errorf("authentication failed: %w", err)
			}
		}
	}

	// Create HTTP client with proper transport and timeout
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Use the configured transport if available
	if c.Config.HTTPTransport != nil {
		httpClient.Transport = c.Config.HTTPTransport
	}

	// Make the request
	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer httpResp.Body.Close()

	// Check status code
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(httpResp.Body)
		return fmt.Errorf("HTTP %d: %s", httpResp.StatusCode, string(bodyBytes))
	}

	// Handle response
	if response != nil {
		bodyBytes, err := io.ReadAll(httpResp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %w", err)
		}

		if err := protojson.Unmarshal(bodyBytes, response); err != nil {
			return fmt.Errorf("failed to unmarshal proto response: %w", err)
		}
	}

	return nil
}
