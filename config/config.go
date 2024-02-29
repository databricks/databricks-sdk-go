package config

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

// CredentialsProvider responsible for configuring static or refreshable
// authentication credentials for Databricks REST APIs
type CredentialsProvider interface {
	// Name returns human-addressable name of this credentials provider name
	Name() string

	// Configure creates HTTP Request Visitor or returns nil if a given credetials
	// are not configured. It returns an error if credentials are misconfigured.
	// Takes a context and a pointer to a Config instance, that holds auth mutex.
	Configure(context.Context, *Config) (func(*http.Request) error, error)
}

type Loader interface {
	// Name is human-addressable representation of this config resolver
	Name() string
	Configure(*Config) error
}

// Config represents configuration for Databricks Connectivity
type Config struct {
	// Credentials holds an instance of Credentials Provider to authenticate with Databricks REST APIs.
	// If no credentials provider is specified, `DefaultCredentials` are implicitly used.
	Credentials CredentialsProvider

	// Databricks host (either of workspace endpoint or Accounts API endpoint)
	Host string `name:"host" env:"DATABRICKS_HOST"`

	ClusterID   string `name:"cluster_id" env:"DATABRICKS_CLUSTER_ID"`
	WarehouseID string `name:"warehouse_id" env:"DATABRICKS_WAREHOUSE_ID"`

	// URL of the metadata service that provides authentication credentials.
	MetadataServiceURL string `name:"metadata_service_url" env:"DATABRICKS_METADATA_SERVICE_URL" auth:"metadata-service,sensitive"`

	// Databricks Account ID for Accounts API. This field is used in dependencies.
	AccountID string `name:"account_id" env:"DATABRICKS_ACCOUNT_ID"`

	Token    string `name:"token" env:"DATABRICKS_TOKEN" auth:"pat,sensitive"`
	Username string `name:"username" env:"DATABRICKS_USERNAME" auth:"basic"`
	Password string `name:"password" env:"DATABRICKS_PASSWORD" auth:"basic,sensitive"`

	// Connection profile specified within ~/.databrickscfg.
	Profile string `name:"profile" env:"DATABRICKS_CONFIG_PROFILE"`

	// Location of the Databricks CLI credentials file, that is created
	// by `databricks configure --token` command. By default, it is located
	// in ~/.databrickscfg.
	ConfigFile string `name:"config_file" env:"DATABRICKS_CONFIG_FILE"`

	GoogleServiceAccount string `name:"google_service_account" env:"DATABRICKS_GOOGLE_SERVICE_ACCOUNT" auth:"google"`
	GoogleCredentials    string `name:"google_credentials" env:"GOOGLE_CREDENTIALS" auth:"google,sensitive"`

	// Azure Resource Manager ID for Azure Databricks workspace, which is exhanged for a Host
	AzureResourceID string `name:"azure_workspace_resource_id" env:"DATABRICKS_AZURE_RESOURCE_ID" auth:"azure-cli,azure-msi" auth_group:"azure"`

	AzureUseMSI       bool   `name:"azure_use_msi" env:"ARM_USE_MSI" auth:"azure-msi" auth_group:"azure"`
	AzureClientSecret string `name:"azure_client_secret" env:"ARM_CLIENT_SECRET" auth:"azure-client-secret,sensitive" auth_group:"azure"`
	AzureClientID     string `name:"azure_client_id" env:"ARM_CLIENT_ID" auth:"azure-client-secret,azure-msi" auth_group:"azure"`
	AzureTenantID     string `name:"azure_tenant_id" env:"ARM_TENANT_ID" auth:"azure-cli,azure-client-secret" auth_group:"azure"`

	// AzureEnvironment (PUBLIC, USGOVERNMENT, CHINA) has specific set of API endpoints. Starting from v0.26.0,
	// the environment is determined based on the workspace hostname, if it's specified.
	AzureEnvironment string `name:"azure_environment" env:"ARM_ENVIRONMENT"`

	// Azure Login Application ID. Must be set if authenticating for non-production workspaces. Starting from v0.26.0,
	// the correct Azure Login App ID is determined based on the Azure Databricks Workspace hostname.
	//
	// Deprecated: this configuration property no longer has any effect and will be removed in the future
	// versions of Go SDK.
	AzureLoginAppID string `name:"azure_login_app_id" env:"DATABRICKS_AZURE_LOGIN_APP_ID" auth:"azure"`

	ClientID     string `name:"client_id" env:"DATABRICKS_CLIENT_ID" auth:"oauth"`
	ClientSecret string `name:"client_secret" env:"DATABRICKS_CLIENT_SECRET" auth:"oauth,sensitive"`

	// Path to the Databricks CLI (version >= 0.100.0).
	DatabricksCliPath string `name:"databricks_cli_path" env:"DATABRICKS_CLI_PATH"`

	// When multiple auth attributes are available in the environment, use the auth type
	// specified by this argument. This argument also holds currently selected auth.
	AuthType string `name:"auth_type" env:"DATABRICKS_AUTH_TYPE" auth:"-"`

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

	// Number of seconds to keep retrying HTTP requests. Default is 300 (5 minutes)
	RetryTimeoutSeconds int `name:"retry_timeout_seconds" auth:"-"`

	// HTTPTransport can be overriden for unit testing and together with tooling like https://github.com/google/go-replayers
	HTTPTransport http.RoundTripper

	// Environment override to return when resolving the current environment.
	DatabricksEnvironment *DatabricksEnvironment

	Loaders []Loader

	// marker for configuration resolving
	resolved bool

	// internal client used in for authentication purposes:
	//  - Databricks Metadata Service: request/refresh tokens from parent processes, like Databricks VSCode extension
	//	- Azure Active Directory (AAD): request/refresh OAuth tokens
	//  - Azure Instance Metadata Service (IMDS): request/refresh OAuth tokens for Azure Managed Identity
	//  - Azure Resource Manager (ARM): resolve host if only Azure Databricks Resource ID provided
	refreshClient *httpclient.ApiClient

	// internal background context used for authentication purposes together with refreshClient
	refreshCtx context.Context

	// marker for testing fixture
	isTesting bool

	// Mutex used by Authenticate method to guard `auth`, which
	// has to be lazily created on the first request to Databricks API.
	// It is done because databricks host and token may often be available
	// only in the middle of Terraform DAG execution.
	// This mutex is also used for config resolution.
	mu sync.Mutex

	// HTTP request interceptor, that assigns Authorization header
	auth func(r *http.Request) error
}

// NewWithWorkspaceHost returns a new instance of the Config with the host set to
// the workspace host. Fields that are not relevant to workspace-level config,
// like account ID, are omitted. Workspace-level attributes that cannot be
// computed from the host alone, like Azure Resource ID, are also omitted.
func (c *Config) NewWithWorkspaceHost(host string) (*Config, error) {
	err := c.EnsureResolved()
	if err != nil {
		return nil, err
	}

	var fieldsToSkip = map[string]struct{}{
		"Host":            {},
		"AzureResourceID": {},
		"AccountID":       {},
	}
	res := &Config{}
	cv := reflect.ValueOf(c).Elem()
	resv := reflect.ValueOf(res).Elem()
	for i := 0; i < resv.NumField(); i++ {
		field := resv.Field(i)
		if !field.CanSet() {
			continue
		}
		if _, ok := fieldsToSkip[resv.Type().Field(i).Name]; ok {
			continue
		}
		field.Set(cv.Field(i))
	}

	res.Host = host
	res.isTesting = c.isTesting
	// We need to reresolve the config with the updated host in general. For
	// example, the audience for OAuth tokens provided by GCP is derived from
	// the host, so account-level tokens cannot be used at workspace-level or
	// vice-versa.
	//
	// In the future, when unified login is widely available, we may be able to
	// reuse the authentication visitor specifically for in-house OAuth.
	return res, nil
}

// Authenticate adds special headers to HTTP request to authorize it to work with Databricks REST API
func (c *Config) Authenticate(r *http.Request) error {
	err := c.EnsureResolved()
	if err != nil {
		return err
	}
	err = c.authenticateIfNeeded()
	if err != nil {
		return err
	}
	return c.auth(r)
}

// IsAzure returns if the client is configured for Azure Databricks.
func (c *Config) IsAzure() bool {
	if c.AzureResourceID != "" {
		return true
	}
	return c.Environment().Cloud == CloudAzure
}

// IsGcp returns if the client is configured for Databricks on Google Cloud.
func (c *Config) IsGcp() bool {
	return c.Environment().Cloud == CloudGCP
}

// IsAws returns if the client is configured for Databricks on AWS.
func (c *Config) IsAws() bool {
	return c.Host != "" && !c.IsAzure() && !c.IsGcp()
}

// IsAccountClient returns true if client is configured for Accounts API
func (c *Config) IsAccountClient() bool {
	if c.AccountID != "" && c.isTesting {
		return true
	}

	accountsPrefixes := []string{
		"https://accounts.",
		"https://accounts-dod.",
	}
	for _, prefix := range accountsPrefixes {
		if strings.HasPrefix(c.Host, prefix) {
			return true
		}
	}
	return false
}

func (c *Config) EnsureResolved() error {
	if c.resolved {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.resolved {
		return nil
	}
	if len(c.Loaders) == 0 {
		c.Loaders = []Loader{
			ConfigAttributes,
			ConfigFile,
		}
	}
	ctx := context.Background()
	for _, loader := range c.Loaders {
		logger.Tracef(ctx, "Loading config via %s", loader.Name())
		err := loader.Configure(c)
		if err != nil {

			return c.wrapDebug(fmt.Errorf("resolve: %w", err))
		}
	}
	err := ConfigAttributes.Validate(c)
	if err != nil {
		return c.wrapDebug(fmt.Errorf("validate: %w", err))
	}
	c.refreshCtx = ctx
	c.refreshClient = httpclient.NewApiClient(httpclient.ClientConfig{
		DebugHeaders:       c.DebugHeaders,
		DebugTruncateBytes: c.DebugTruncateBytes,
		InsecureSkipVerify: c.InsecureSkipVerify,
		RetryTimeout:       time.Duration(c.RetryTimeoutSeconds) * time.Second,
		HTTPTimeout:        time.Duration(c.HTTPTimeoutSeconds) * time.Second,
		Transport:          c.HTTPTransport,
		ErrorMapper:        c.refreshTokenErrorMapper,
		TransientErrors: []string{
			"throttled",
			"too many requests",
			"429",
			"request limit exceeded",
			"rate limit",
		},
	})
	c.resolved = true
	return nil
}

func (c *Config) WithTesting() *Config {
	c.isTesting = true
	return c
}

func (c *Config) CanonicalHostName() string {
	c.fixHostIfNeeded()
	return c.Host
}

func (c *Config) wrapDebug(err error) error {
	debug := ConfigAttributes.DebugString(c)
	if debug == "" {
		return err
	}
	return fmt.Errorf("%w. %s", err, debug)
}

// authenticateIfNeeded lazily authenticates across authorizers or returns error
func (c *Config) authenticateIfNeeded() error {
	if c.auth != nil {
		return nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.auth != nil {
		return nil
	}
	if c.Credentials == nil {
		c.Credentials = &DefaultCredentials{}
	}
	c.fixHostIfNeeded()
	ctx := c.refreshClient.InContextForOAuth2(c.refreshCtx)
	visitor, err := c.Credentials.Configure(ctx, c)
	if err != nil {
		return c.wrapDebug(fmt.Errorf("%s auth: %w", c.Credentials.Name(), err))
	}
	if visitor == nil {
		return c.wrapDebug(fmt.Errorf("%s auth: not configured", c.Credentials.Name()))
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

func (c *Config) refreshTokenErrorMapper(ctx context.Context, resp common.ResponseWrapper) error {
	defaultErr := httpclient.DefaultErrorMapper(ctx, resp)
	if defaultErr == nil {
		return nil
	}
	err, ok := defaultErr.(*httpclient.HttpError)
	if !ok {
		return err
	}
	if c.IsAzure() {
		return c.mapAzureError(err)
	}
	return &tokenError{
		message: err.Message,
		err:     err,
	}
}
