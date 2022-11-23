package databricks

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/databricks/databricks-sdk-go/internal"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/useragent"
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
	AzureResourceID string `name:"azure_workspace_resource_id" env:"DATABRICKS_AZURE_RESOURCE_ID" auth:"azure"`

	AzureUseMSI       bool   `name:"azure_use_msi" env:"ARM_USE_MSI" auth:"azure"`
	AzureClientSecret string `name:"azure_client_secret" env:"ARM_CLIENT_SECRET" auth:"azure,sensitive"`
	AzureClientID     string `name:"azure_client_id" env:"ARM_CLIENT_ID" auth:"azure"`
	AzureTenantID     string `name:"azure_tenant_id" env:"ARM_TENANT_ID" auth:"azure"`
	// AzureEnvironment (Public, UsGov, China, Germany) has specific set of API endpoints.
	AzureEnvironment string `name:"azure_environment" env:"ARM_ENVIRONMENT"`

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

	// Azure Login Application ID. Must be set if authenticating for non-production workspaces.
	AzureLoginAppID string `name:"azure_login_app_id" env:"DATABRICKS_AZURE_LOGIN_APP_ID" auth:"azure"`

	// Number of seconds to keep retrying HTTP requests. Default is 300 (5 minutes)
	RetryTimeoutSeconds int `name:"retry_timeout_seconds" auth:"-"`

	Loaders []Loader

	// marker for configuration resolving
	resolved bool

	// Mutex used by Authenticate method to guard `auth`, which
	// has to be lazily created on the first request to Databricks API.
	// It is done because databricks host and token may often be available
	// only in the middle of Terraform DAG execution.
	// This mutex is also used for config resolution.
	mu sync.Mutex

	// HTTP request interceptor, that assigns Authorization header
	auth func(r *http.Request) error
}

// Authenticate adds special headers to HTTP request to authorize it to work with Databricks REST API
func (c *Config) Authenticate(r *http.Request) error {
	err := c.EnsureResolved()
	if err != nil {
		return err
	}
	err = c.authenticateIfNeeded(r.Context())
	if err != nil {
		return err
	}
	return c.auth(r)
}

// IsAzure returns true if client is configured for Azure Databricks
func (c *Config) IsAzure() bool {
	return strings.Contains(c.Host, ".azuredatabricks.net") || c.AzureResourceID != ""
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
func (c *Config) IsAccountClient() bool {
	return strings.HasPrefix(c.Host, "https://accounts.")
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
			KnownConfigLoader{},
		}
	}
	for _, loader := range c.Loaders {
		logger.Tracef("Loading config via %s", loader.Name())
		err := loader.Configure(c)
		if err != nil {

			return c.wrapDebug(fmt.Errorf("resolve: %w", err))
		}
	}
	err := ConfigAttributes.Validate(c)
	if err != nil {
		return c.wrapDebug(fmt.Errorf("validate: %w", err))
	}
	c.resolved = true
	return nil
}

func (c *Config) wrapDebug(err error) error {
	debug := ConfigAttributes.DebugString(c)
	if debug == "" {
		return err
	}
	return fmt.Errorf("%w. %s", err, debug)
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
		c.Credentials = &DefaultCredentials{}
	}
	c.fixHostIfNeeded()
	visitor, err := c.Credentials.Configure(ctx, c)
	if err != nil {
		return c.wrapDebug(fmt.Errorf("%s auth: %w", c.Credentials.Name(), err))
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

func (c *Config) GetAuthType() string         { return c.AuthType }
func (c *Config) GetAccountID() string        { return c.AccountID }
func (c *Config) GetHost() string             { return c.Host }
func (c *Config) GetRetryTimeoutSeconds() int { return c.RetryTimeoutSeconds }
func (c *Config) GetRateLimitPerSecond() int  { return c.RateLimitPerSecond }
func (c *Config) GetHTTPTimeoutSeconds() int  { return c.HTTPTimeoutSeconds }
func (c *Config) GetDebugTruncateBytes() int  { return c.DebugTruncateBytes }
func (c *Config) IsDebugHeaders() bool        { return c.DebugHeaders }
func (c *Config) IsInsecureSkipVerify() bool  { return c.InsecureSkipVerify }

func NewMockConfig(auth func(r *http.Request) error) *Config {
	if auth == nil {
		auth = func(r *http.Request) error {
			return nil
		}
	}
	return &Config{
		AuthType: "mock",
		auth:     auth,
		resolved: true,
	}
}

// Version of this SDK
func Version() string {
	return internal.Version
}

// WithProduct is expected to be set by developers to differentiate their app from others.
//
// Example setting is:
//
//	func init() {
//		databricks.WithProduct("your-product", "0.0.1")
//	}
func WithProduct(name, version string) {
	useragent.WithProduct(name, version)
}

var ConfigAttributes = loadAttrs()

// Attributes holds meta-representation of Config configuration options
type Attributes []ConfigAttribute

func (a Attributes) DebugString(cfg *Config) string {
	buf := []string{}
	attrsUsed := []string{}
	envsUsed := []string{}
	for _, attr := range a {
		if attr.IsZero(cfg) {
			continue
		}
		// Don't include internal fields in debug representation.
		if attr.Internal {
			continue
		}
		v := "***"
		if !attr.Sensitive {
			v = attr.GetString(cfg)
		}
		attrsUsed = append(attrsUsed, fmt.Sprintf("%s=%s", attr.Name, v))
		for _, envName := range attr.EnvVars {
			v := os.Getenv(envName)
			if v == "" {
				continue
			}
			envsUsed = append(envsUsed, envName)
		}
	}
	if len(attrsUsed) > 0 {
		buf = append(buf, fmt.Sprintf("Config: %s", strings.Join(attrsUsed, ", ")))
	}
	if len(envsUsed) > 0 {
		buf = append(buf, fmt.Sprintf("Env: %s", strings.Join(envsUsed, ", ")))
	}
	return strings.Join(buf, ". ")
}

func (a Attributes) Validate(cfg *Config) error {
	authsUsed := map[string]bool{}
	for _, attr := range a {
		if attr.IsZero(cfg) {
			continue
		}
		if attr.Auth == "" {
			continue
		}
		authsUsed[attr.Auth] = true
	}
	if len(authsUsed) <= 1 {
		return nil
	}
	if cfg.AuthType != "" {
		// client has auth preference set
		return nil
	}
	names := []string{}
	for v := range authsUsed {
		names = append(names, v)
	}
	sort.Strings(names)
	return fmt.Errorf("more than one authorization method configured: %s",
		strings.Join(names, " and "))
}

// Name implements Loader interface for environment variables
func (a Attributes) Name() string {
	return "environment"
}

// Configure implements Loader interface for environment variables
func (a Attributes) Configure(cfg *Config) error {
	for _, attr := range a {
		if !attr.IsZero(cfg) {
			// don't overwtite a value previously set
			continue
		}
		v := attr.ReadEnv()
		if v == "" {
			continue
		}
		err := attr.SetS(cfg, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Attributes) ResolveFromStringMap(cfg *Config, m map[string]string) error {
	for _, attr := range a {
		if !attr.IsZero(cfg) {
			// don't overwtite a value previously set
			continue
		}
		v, ok := m[attr.Name]
		if !ok || v == "" {
			continue
		}
		err := attr.SetS(cfg, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Attributes) ResolveFromAnyMap(cfg *Config, m map[string]interface{}) error {
	for _, attr := range a {
		if !attr.IsZero(cfg) {
			// don't overwtite a value previously set
			continue
		}
		v, ok := m[attr.Name]
		if !ok {
			continue
		}
		err := attr.Set(cfg, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadAttrs() (attrs Attributes) {
	t := reflect.TypeOf(Config{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		nameTag := field.Tag.Get("name")
		if nameTag == "" {
			continue
		}
		sensitive := false
		auth := field.Tag.Get("auth")
		authSplit := strings.Split(auth, ",")
		if len(authSplit) == 2 {
			auth = authSplit[0]
			sensitive = authSplit[1] == "sensitive"
		}
		// internal config fields are skipped in debugging
		internal := false
		if auth == "-" {
			auth = ""
			internal = true
		}
		attr := ConfigAttribute{
			Name:      nameTag,
			Auth:      auth,
			Kind:      field.Type.Kind(),
			Sensitive: sensitive,
			Internal:  internal,
			num:       i,
		}
		envTag := field.Tag.Get("env")
		if envTag != "" {
			attr.EnvVars = strings.Split(envTag, ",")
		}
		attrs = append(attrs, attr)
	}
	return
}

// ConfigAttribute provides generic way to work with Config configuration
// attributes and parses `name`, `env`, and `auth` field tags.
type ConfigAttribute struct {
	Name      string
	Kind      reflect.Kind
	EnvVars   []string
	Auth      string
	Sensitive bool
	Internal  bool
	num       int
}

func (a *ConfigAttribute) ReadEnv() string {
	for _, envName := range a.EnvVars {
		v := os.Getenv(envName)
		if v == "" {
			continue
		}
		return v
	}
	return ""
}

func (a *ConfigAttribute) SetS(cfg *Config, v string) error {
	switch a.Kind {
	case reflect.String:
		return a.Set(cfg, v)
	case reflect.Bool:
		vv, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		return a.Set(cfg, vv)
	case reflect.Int:
		vv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		return a.Set(cfg, vv)
	default:
		return fmt.Errorf("cannot set %s of unknown type %s",
			a.Name, reflectKind(a.Kind))
	}
}

func (a *ConfigAttribute) Set(cfg *Config, i interface{}) error {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	switch a.Kind {
	case reflect.String:
		field.SetString(i.(string))
	case reflect.Bool:
		field.SetBool(i.(bool))
	case reflect.Int:
		field.SetInt(int64(i.(int)))
	default:
		// must extensively test with providerFixture to avoid this one
		return fmt.Errorf("cannot set %s of unknown type %s", a.Name, reflectKind(a.Kind))
	}
	return nil
}

func (a *ConfigAttribute) IsZero(cfg *Config) bool {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	return field.IsZero()
}

func (a *ConfigAttribute) GetString(cfg *Config) string {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	return fmt.Sprintf("%v", field.Interface())
}

var kindMap = map[reflect.Kind]string{
	reflect.Bool:          "Bool",
	reflect.Int:           "Int",
	reflect.Int8:          "Int8",
	reflect.Int16:         "Int16",
	reflect.Int32:         "Int32",
	reflect.Int64:         "Int64",
	reflect.Uint:          "Uint",
	reflect.Uint8:         "Uint8",
	reflect.Uint16:        "Uint16",
	reflect.Uint32:        "Uint32",
	reflect.Uint64:        "Uint64",
	reflect.Uintptr:       "Uintptr",
	reflect.Float32:       "Float32",
	reflect.Float64:       "Float64",
	reflect.Complex64:     "Complex64",
	reflect.Complex128:    "Complex128",
	reflect.Array:         "Array",
	reflect.Chan:          "Chan",
	reflect.Func:          "Func",
	reflect.Interface:     "Interface",
	reflect.Ptr:           "Ptr",
	reflect.Slice:         "Slice",
	reflect.String:        "String",
	reflect.Struct:        "Struct",
	reflect.UnsafePointer: "UnsafePointer",
}

func reflectKind(k reflect.Kind) string {
	n, ok := kindMap[k]
	if !ok {
		return "other"
	}
	return n
}
