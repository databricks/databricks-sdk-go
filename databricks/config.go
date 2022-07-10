package databricks

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type Credentials interface {
	Name() string
	Configure(context.Context, *Config) (func(*http.Request) error, error)
}

// Config represents configuration for Databricks Connectivity
type Config struct {
	Credentials Credentials

	// Databricks host (either of workspace endpoint or Accounts API endpoint)
	Host string `name:"host" env:"DATABRICKS_HOST"`

	// Databricks Account ID for Accounts API. This field is used in dependencies.
	AccountID string `name:"account_id" env:"DATABRICKS_ACCOUNT_ID"`

	// When multiple auth attributes are available in the environment, use the auth type
	// specified by this argument. This argument also holds currently selected auth.
	AuthType string `name:"auth_type" auth:"-"`

	// Skip SSL certificate verification for HTTP calls.
	// Use at your own risk or for unit testing purposes.
	InsecureSkipVerify bool `name:"skip_verify" auth:"-"`

	// Number of seconds for HTTP timeout
	HTTPTimeoutSeconds int `name:"http_timeout_seconds" auth:"-"`

	// Truncate JSON fields in JSON above this limit. Default is 96.
	DebugTruncateBytes int `name:"debug_truncate_bytes" env:"DATABRICKS_DEBUG_TRUNCATE_BYTES" auth:"-"`

	// Debug HTTP headers of requests made by the provider. Default is false.
	DebugHeaders bool `name:"debug_headers" env:"DATABRICKS_DEBUG_HEADERS" auth:"-"`

	// Maximum number of requests per second made to Databricks REST API.
	RateLimitPerSecond int `name:"rate_limit" env:"DATABRICKS_RATE_LIMIT" auth:"-"`

	// Azure Resource Manager ID for Azure Databricks workspace, which is exhanged for a Host
	AzureResourceID string `name:"azure_workspace_resource_id" env:"DATABRICKS_AZURE_RESOURCE_ID"`

	// Azure Environment (Public, UsGov, China, Germany) has specific set of API endpoints
	Environment string `name:"azure_environment" env:"ARM_ENVIRONMENT"` // TODO: rename to AzureEnvironment

	RetryWaitMin     time.Duration
	RetryWaitMax     time.Duration
	MaxRetryAttempts int

	// Mutex used by Authenticate method to guard `auth`, which
	// has to be lazily created on the first request to Databricks API.
	// It is done because databricks host and token may often be available
	// only in the middle of Terraform DAG execution.
	mu sync.Mutex

	// HTTP request interceptor, that assigns Authorization header
	auth func(r *http.Request) error
}

// Authenticate adds special headers to HTTP request to authorize it to work with Databricks REST API
func (c *Config) Authenticate(r *http.Request) error {
	err := c.authenticateIfNeeded(r.Context())
	if err != nil {
		return err
	}
	return c.auth(r)
}

// IsAzure returns true if client is configured for Azure Databricks
func (c *Config) IsAzure() bool {
	return strings.Contains(c.Host, ".azuredatabricks.net")
}

// IsGcp returns true if client is configured for GCP
func (c *Config) IsGcp() bool {
	return strings.Contains(c.Host, ".gcp.databricks.com")
}

// IsAws returns true if client is configured for AWS
func (c *Config) IsAws() bool {
	return !c.IsAzure() && !c.IsGcp()
}

// IsAws returns true if client is configured for Accounts API
func (c *Config) IsAccountsClient() bool {
	return strings.HasPrefix(c.Host, "https://accounts.")
}

// authenticateIfNeeded lazily authenticates across authorizers or returns error
func (c *Config) authenticateIfNeeded(ctx context.Context) error {
	if c.auth != nil {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.auth != nil {
		return nil
	}
	if c.Credentials == nil {
		c.Credentials = DefaultCredentials{}
	}
	c.fixHostIfNeeded()
	visitor, err := c.Credentials.Configure(ctx, c)
	if err != nil {
		return fmt.Errorf("cannot configure %s auth: %w", c.Credentials.Name(), err)
	}
	c.auth = visitor
	c.AuthType = c.Credentials.Name()
	c.fixHostIfNeeded()
	// TODO: error customization
	return nil
}

func (c *Config) fixHostIfNeeded() error {
	// Nothing to fix if the host isn't set.
	if c.Host == "" {
		return nil
	}
	parsedHost, err := url.Parse(c.Host)
	if err != nil {
		return err
	}
	// If the host is empty, assume the scheme wasn't included.
	if parsedHost.Host == "" {
		parsedHost, err = url.Parse("https://" + c.Host)
		if err != nil {
			return err
		}
	}
	// Create new instance to ensure other fields are initialized as empty.
	parsedHost = &url.URL{
		Scheme: parsedHost.Scheme,
		Host:   parsedHost.Host,
	}
	// Store sanitized version of c.Host.
	c.Host = parsedHost.String()
	return nil
}
