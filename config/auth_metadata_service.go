package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

// timeout to wait for metadata service to respond
const metadataServiceTimeout = 10 * time.Second

const MetadataServiceVersion = "1"
const MetadataServiceVersionHeader = "X-Databricks-Metadata-Version"
const MetadataServiceHostHeader = "X-Databricks-Host"

var errMetadataServiceMalformed = errors.New("invalid auth server URL")
var errMetadataServiceNotLocalhost = errors.New("only localhost URLs are allowed")

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

func (c MetadataServiceCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.MetadataServiceURL == "" || cfg.Host == "" {
		return nil, nil
	}
	parsedMetadataServiceURL, err := url.Parse(cfg.MetadataServiceURL)
	if err != nil {
		// go 1.19 doesn't allow multiple error unwraping
		return nil, fmt.Errorf("%w: %s", errMetadataServiceMalformed, err)
	}
	// only allow localhost URLs
	if parsedMetadataServiceURL.Hostname() != "localhost" && parsedMetadataServiceURL.Hostname() != "127.0.0.1" {
		return nil, fmt.Errorf("%w: %s", errMetadataServiceNotLocalhost, cfg.MetadataServiceURL)
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
	visitor := refreshableVisitor(&ms)
	return credentials.NewCredentialsProvider(visitor), nil
}

type metadataService struct {
	metadataServiceURL *url.URL
	config             *Config
}

// performs a request to the metadata service and returns the token
func (s metadataService) Get() (*oauth2.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), metadataServiceTimeout)
	defer cancel()
	var inner msiToken
	err := s.config.refreshClient.Do(ctx, http.MethodGet,
		s.metadataServiceURL.String(),
		httpclient.WithRequestHeader(MetadataServiceVersionHeader, MetadataServiceVersion),
		httpclient.WithRequestHeader(MetadataServiceHostHeader, s.config.Host),
		httpclient.WithResponseUnmarshal(&inner),
	)
	if err != nil {
		return nil, fmt.Errorf("token request: %w", err)
	}
	return inner.Token()
}

func (t metadataService) Token() (*oauth2.Token, error) {
	token, err := t.Get()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, fmt.Errorf("no token returned from metadata service")
	}
	logger.Debugf(context.Background(),
		"Refreshed access token from local metadata service, which expires on %s",
		token.Expiry.Format(time.RFC3339))
	return token, nil
}
