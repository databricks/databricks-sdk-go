package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"github.com/databricks/databricks-sdk-go/config/experimental/auth/authconv"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

var (
	errInvalidToken       = errors.New("invalid token")
	errInvalidTokenExpiry = errors.New("invalid token expiry")
)

const (
	azureIMDSAuthority = "http://169.254.169.254"
	azureIMDSTokenPath = "/metadata/identity/oauth2/token"

	azureIdentityEndpointEnv               = "IDENTITY_ENDPOINT"
	azureIdentityHeaderEnv                 = "IDENTITY_HEADER"
	azureIdentityServerThumbprintEnv       = "IDENTITY_SERVER_THUMBPRINT"
	azureIMDSEndpointEnv                   = "IMDS_ENDPOINT"
	azureMSIEndpointEnv                    = "MSI_ENDPOINT"
	azureMSISecretEnv                      = "MSI_SECRET"
	azurePodIdentityAuthorityHostEnv       = "AZURE_POD_IDENTITY_AUTHORITY_HOST"
	azureAuthorityHostEnv                  = "AZURE_AUTHORITY_HOST"
	azureClientIDEnv                       = "AZURE_CLIENT_ID"
	azureTenantIDEnv                       = "AZURE_TENANT_ID"
	azureFederatedTokenFileEnv             = "AZURE_FEDERATED_TOKEN_FILE"
	azureArcLinuxTokenDirectory            = "/var/opt/azcmagent/tokens"
	azureArcMaximumSecretFileSize    int64 = 4096

	// Managed identity endpoints are local services and should fail promptly
	// when the SDK isn't running in the expected Azure environment.
	azureMSITimeout = 10 * time.Second
)

type AzureMsiCredentials struct{}

func (c AzureMsiCredentials) Name() string {
	return "azure-msi"
}

func (c AzureMsiCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if !cfg.IsAzure() || !cfg.AzureUseMSI || (cfg.AzureResourceID == "" && cfg.Host == "") {
		return nil, nil
	}
	env := cfg.Environment()
	// If the host is not set, we need to resolve it from the Azure Resource ID.
	// This is only needed for Workspaces, because Accounts always have a host.
	if cfg.Host == "" {
		err := cfg.azureEnsureWorkspaceUrl(ctx, c)
		if err != nil {
			return nil, fmt.Errorf("resolve host: %w", err)
		}
	}
	logger.Debugf(ctx, "Generating AAD token via Azure MSI")
	opts := cacheOptions(cfg)
	inner := azureReuseTokenSource(nil, c.tokenSourceFor(ctx, cfg, "", env.AzureApplicationID), opts...)
	management := azureReuseTokenSource(nil, c.tokenSourceFor(ctx, cfg, "", env.AzureServiceManagementEndpoint()), opts...)
	visitor := azureVisitor(cfg, serviceToServiceVisitor(inner, management, xDatabricksAzureSpManagementToken, false, opts...))
	return newVisitorOAuthCredentials(visitor, inner), nil
}

// tokenSourceFor implements azureHostResolver so MSI can resolve a workspace host.
func (c AzureMsiCredentials) tokenSourceFor(_ context.Context, cfg *Config, _, resource string) auth.TokenSource {
	return newAzureMSITokenSource(cfg.refreshClient, resource, cfg.AzureClientID)
}

// NewAzureMsiTokenSource returns [oauth2.TokenSource] for passwordless authentication via Azure Managed Identity.
func NewAzureMsiTokenSource(client *httpclient.ApiClient, resource, clientID string) oauth2.TokenSource {
	return authconv.OAuth2TokenSource(newAzureMSITokenSource(client, resource, clientID))
}

type azureManagedIdentityType string

const (
	azureManagedIdentityServiceFabric azureManagedIdentityType = "Service Fabric"
	azureManagedIdentityAppService    azureManagedIdentityType = "App Service"
	azureManagedIdentityArc           azureManagedIdentityType = "Azure Arc"
	azureManagedIdentityML            azureManagedIdentityType = "Azure ML"
	azureManagedIdentityCloudShell    azureManagedIdentityType = "Cloud Shell"
	azureManagedIdentityWorkload      azureManagedIdentityType = "workload identity"
	azureManagedIdentityIMDS          azureManagedIdentityType = "IMDS"
)

type azureManagedIdentityEndpoint struct {
	identityType azureManagedIdentityType
	url          string
	secret       string
	tenantID     string
	tokenFile    string
}

// azureManagedIdentityEndpointFromEnvironment follows the selection order used
// by Azure Identity's ManagedIdentityCredential. The first matching Azure host
// environment wins, and IMDS is used only when no host-specific environment is
// configured.
func azureManagedIdentityEndpointFromEnvironment(clientID string) (azureManagedIdentityEndpoint, string, error) {
	identityEndpoint := getenv(azureIdentityEndpointEnv)
	if identityEndpoint != "" {
		identityHeader := getenv(azureIdentityHeaderEnv)
		if identityHeader != "" {
			if getenv(azureIdentityServerThumbprintEnv) != "" {
				endpoint := azureManagedIdentityEndpoint{
					identityType: azureManagedIdentityServiceFabric,
					url:          identityEndpoint,
					secret:       identityHeader,
				}
				if clientID != "" {
					return endpoint, clientID, fmt.Errorf("azure_client_id is not supported by %q managed identity", endpoint.identityType)
				}
				return endpoint, clientID, nil
			}
			return azureManagedIdentityEndpoint{
				identityType: azureManagedIdentityAppService,
				url:          identityEndpoint,
				secret:       identityHeader,
			}, clientID, nil
		}
		if getenv(azureIMDSEndpointEnv) != "" {
			endpoint := azureManagedIdentityEndpoint{
				identityType: azureManagedIdentityArc,
				url:          identityEndpoint,
			}
			if clientID != "" {
				return endpoint, clientID, fmt.Errorf("azure_client_id is not supported by %q managed identity", endpoint.identityType)
			}
			return endpoint, clientID, nil
		}
		return azureManagedIdentityEndpoint{}, clientID, fmt.Errorf(
			"no managed identity endpoint found: %s requires %s or %s",
			azureIdentityEndpointEnv,
			azureIdentityHeaderEnv,
			azureIMDSEndpointEnv,
		)
	}

	if msiEndpoint := getenv(azureMSIEndpointEnv); msiEndpoint != "" {
		if secret := getenv(azureMSISecretEnv); secret != "" {
			return azureManagedIdentityEndpoint{
				identityType: azureManagedIdentityML,
				url:          msiEndpoint,
				secret:       secret,
			}, clientID, nil
		}
		endpoint := azureManagedIdentityEndpoint{
			identityType: azureManagedIdentityCloudShell,
			url:          msiEndpoint,
		}
		if clientID != "" {
			return endpoint, clientID, fmt.Errorf("azure_client_id is not supported by %q managed identity", endpoint.identityType)
		}
		return endpoint, clientID, nil
	}

	authorityHost := getenv(azureAuthorityHostEnv)
	tenantID := getenv(azureTenantIDEnv)
	tokenFile := getenv(azureFederatedTokenFileEnv)
	if authorityHost != "" && tenantID != "" && tokenFile != "" {
		if clientID == "" {
			clientID = getenv(azureClientIDEnv)
		}
		endpoint := azureManagedIdentityEndpoint{
			identityType: azureManagedIdentityWorkload,
			url:          authorityHost,
			tenantID:     tenantID,
			tokenFile:    tokenFile,
		}
		if clientID == "" {
			return endpoint, clientID, errors.New("azure_client_id is required for workload identity")
		}
		return endpoint, clientID, nil
	}

	imdsAuthority := getenv(azurePodIdentityAuthorityHostEnv)
	if imdsAuthority == "" {
		imdsAuthority = azureIMDSAuthority
	}
	return azureManagedIdentityEndpoint{
		identityType: azureManagedIdentityIMDS,
		url:          strings.TrimRight(imdsAuthority, "/") + azureIMDSTokenPath,
	}, clientID, nil
}

type azureMSITokenSource struct {
	client             *httpclient.ApiClient
	resource           string
	clientID           string
	endpoint           azureManagedIdentityEndpoint
	endpointErr        error
	now                func() time.Time
	readAzureArcSecret func(string) (string, error)
}

func newAzureMSITokenSource(client *httpclient.ApiClient, resource, clientID string) *azureMSITokenSource {
	endpoint, clientID, err := azureManagedIdentityEndpointFromEnvironment(clientID)
	return &azureMSITokenSource{
		client:             client,
		resource:           resource,
		clientID:           clientID,
		endpoint:           endpoint,
		endpointErr:        err,
		now:                time.Now,
		readAzureArcSecret: readAzureArcSecret,
	}
}

func (s *azureMSITokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	if s.endpointErr != nil {
		return nil, s.endpointErr
	}
	ctx, cancel := context.WithTimeout(ctx, azureMSITimeout)
	defer cancel()

	now := s.now
	if now == nil {
		now = time.Now
	}
	requestTime := now()

	switch s.endpoint.identityType {
	case azureManagedIdentityWorkload:
		return s.workloadIdentityToken(ctx, requestTime)
	case azureManagedIdentityServiceFabric,
		azureManagedIdentityAppService,
		azureManagedIdentityML,
		azureManagedIdentityCloudShell,
		azureManagedIdentityIMDS:
		request, err := s.endpoint.tokenRequest(s.resource, s.clientID)
		if err != nil {
			return nil, err
		}
		inner, err := s.requestToken(ctx, request)
		if err != nil {
			return nil, err
		}
		return inner.managedIdentityToken(requestTime, s.endpoint.identityType == azureManagedIdentityML)
	case azureManagedIdentityArc:
		return s.azureArcToken(ctx, requestTime)
	default:
		return nil, fmt.Errorf("unknown Azure managed identity type: %q", s.endpoint.identityType)
	}
}

type azureMSITokenRequest struct {
	method      string
	url         string
	headers     map[string]string
	data        map[string]string
	form        bool
	encodedForm string
}

func (e azureManagedIdentityEndpoint) tokenRequest(resource, clientID string) (azureMSITokenRequest, error) {
	request := azureMSITokenRequest{
		method:  http.MethodGet,
		url:     e.url,
		headers: map[string]string{},
		data: map[string]string{
			"resource": resource,
		},
	}

	switch e.identityType {
	case azureManagedIdentityServiceFabric:
		request.headers["Secret"] = e.secret
		request.data["api-version"] = "2019-07-01-preview"
	case azureManagedIdentityAppService:
		request.headers["X-IDENTITY-HEADER"] = e.secret
		request.data["api-version"] = "2019-08-01"
		if clientID != "" {
			request.data["client_id"] = clientID
		}
	case azureManagedIdentityArc:
		request.headers["Metadata"] = "true"
		request.data["api-version"] = "2020-06-01"
	case azureManagedIdentityML:
		request.headers["secret"] = e.secret
		request.data["api-version"] = "2017-09-01"
		if clientID != "" {
			request.data["clientid"] = clientID
		}
	case azureManagedIdentityCloudShell:
		request.method = http.MethodPost
		request.headers["Metadata"] = "true"
		request.form = true
	case azureManagedIdentityIMDS:
		request.headers["Metadata"] = "true"
		request.data["api-version"] = "2018-02-01"
		if clientID != "" {
			request.data["client_id"] = clientID
		}
	case azureManagedIdentityWorkload:
		return azureMSITokenRequest{}, errors.New("workload identity requires a token exchange request")
	default:
		return azureMSITokenRequest{}, fmt.Errorf("unknown Azure managed identity type: %q", e.identityType)
	}

	return request, nil
}

func (s *azureMSITokenSource) requestToken(ctx context.Context, request azureMSITokenRequest) (msiToken, error) {
	var inner msiToken
	requestOptions := []httpclient.DoOption{
		httpclient.WithRequestHeaders(request.headers),
		httpclient.WithResponseUnmarshal(&inner),
	}
	if request.encodedForm != "" {
		requestOptions = append(
			requestOptions,
			httpclient.WithRequestHeader("Content-Type", httpclient.UrlEncodedContentType),
			httpclient.WithRequestData(strings.NewReader(request.encodedForm)),
		)
	} else if request.form {
		requestOptions = append(requestOptions, httpclient.WithUrlEncodedData(request.data))
	} else {
		requestOptions = append(requestOptions, httpclient.WithRequestData(request.data))
	}
	if err := s.client.Do(ctx, request.method, request.url, requestOptions...); err != nil {
		return msiToken{}, fmt.Errorf("request managed identity token from %q: %w", request.url, err)
	}
	return inner, nil
}

func (s *azureMSITokenSource) workloadIdentityToken(ctx context.Context, requestTime time.Time) (*oauth2.Token, error) {
	assertion, err := os.ReadFile(s.endpoint.tokenFile)
	if err != nil {
		return nil, fmt.Errorf("read federated token file %q: %w", s.endpoint.tokenFile, err)
	}

	tokenEndpoint := strings.TrimRight(s.endpoint.url, "/") + "/" + url.PathEscape(s.endpoint.tenantID) + "/oauth2/v2.0/token"
	scope := s.resource
	if !strings.HasSuffix(scope, "/.default") {
		scope = strings.TrimRight(scope, "/") + "/.default"
	}
	form := url.Values{
		"client_assertion":      []string{string(assertion)},
		"client_assertion_type": []string{"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
		"client_id":             []string{s.clientID},
		"grant_type":            []string{"client_credentials"},
		"scope":                 []string{scope},
	}
	request := azureMSITokenRequest{
		method:      http.MethodPost,
		url:         tokenEndpoint,
		headers:     map[string]string{},
		encodedForm: form.Encode(),
	}
	inner, err := s.requestToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return inner.managedIdentityToken(requestTime, false)
}

func (s *azureMSITokenSource) azureArcToken(ctx context.Context, requestTime time.Time) (*oauth2.Token, error) {
	request, err := s.endpoint.tokenRequest(s.resource, s.clientID)
	if err != nil {
		return nil, err
	}
	inner, err := s.requestToken(ctx, request)
	if err == nil {
		return inner.managedIdentityToken(requestTime, false)
	}

	keyFile, challenged, challengeErr := azureArcChallengeKeyFile(err)
	if challengeErr != nil {
		return nil, fmt.Errorf("handle Azure Arc challenge from %q: %w", request.url, challengeErr)
	}
	if !challenged {
		return nil, err
	}
	secretReader := s.readAzureArcSecret
	if secretReader == nil {
		secretReader = readAzureArcSecret
	}
	secret, err := secretReader(keyFile)
	if err != nil {
		return nil, fmt.Errorf("read Azure Arc managed identity secret %q: %w", keyFile, err)
	}
	request.headers["Authorization"] = "Basic " + secret
	inner, err = s.requestToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return inner.managedIdentityToken(requestTime, false)
}

func azureArcChallengeKeyFile(err error) (string, bool, error) {
	httpError := azureManagedIdentityHTTPError(err)
	if httpError == nil || httpError.StatusCode != http.StatusUnauthorized {
		return "", false, nil
	}
	challenge := httpError.Header().Get("WWW-Authenticate")
	if challenge == "" {
		return "", true, errors.New("Azure Arc managed identity response has no WWW-Authenticate header")
	}
	_, keyFile, ok := strings.Cut(challenge, "=")
	if !ok {
		return "", true, fmt.Errorf("invalid Azure Arc WWW-Authenticate header: %q", challenge)
	}
	keyFile = strings.Trim(strings.TrimSpace(keyFile), `"`)
	if keyFile == "" {
		return "", true, fmt.Errorf("invalid Azure Arc WWW-Authenticate header: %q", challenge)
	}
	return keyFile, true, nil
}

func azureManagedIdentityHTTPError(err error) *httpclient.HttpError {
	var azureError *tokenError
	if errors.As(err, &azureError) {
		return azureError.err
	}
	var httpError *httpclient.HttpError
	if errors.As(err, &httpError) {
		return httpError
	}
	return nil
}

func readAzureArcSecret(keyFile string) (string, error) {
	expectedDirectory, err := azureArcTokenDirectory()
	if err != nil {
		return "", err
	}
	return readAzureArcSecretFromDirectory(keyFile, expectedDirectory)
}

func readAzureArcSecretFromDirectory(keyFile, expectedDirectory string) (string, error) {
	info, err := os.Stat(keyFile)
	if err != nil {
		return "", err
	}
	if filepath.Clean(filepath.Dir(keyFile)) != filepath.Clean(expectedDirectory) {
		return "", fmt.Errorf("unexpected Azure Arc managed identity secret directory %q", filepath.Dir(keyFile))
	}
	if filepath.Ext(keyFile) != ".key" {
		return "", fmt.Errorf("Azure Arc managed identity secret file %q must have a .key extension", keyFile)
	}
	if info.Size() > azureArcMaximumSecretFileSize {
		return "", fmt.Errorf("Azure Arc managed identity secret file %q exceeds %d bytes", keyFile, azureArcMaximumSecretFileSize)
	}
	secret, err := os.ReadFile(keyFile)
	if err != nil {
		return "", err
	}
	return string(secret), nil
}

func azureArcTokenDirectory() (string, error) {
	switch runtime.GOOS {
	case "linux":
		return azureArcLinuxTokenDirectory, nil
	case "windows":
		programData := getenv("PROGRAMDATA")
		if programData == "" {
			return "", errors.New("PROGRAMDATA is required for Azure Arc managed identity on Windows")
		}
		return filepath.Join(programData, "AzureConnectedMachineAgent", "Tokens"), nil
	default:
		return "", fmt.Errorf("Azure Arc managed identity is not supported on %q", runtime.GOOS)
	}
}

type msiToken struct {
	TokenType    string           `json:"token_type"`
	AccessToken  string           `json:"access_token,omitempty"`
	RefreshToken string           `json:"refresh_token,omitempty"`
	ExpiresOn    azureTokenExpiry `json:"expires_on"`
	ExpiresIn    azureTokenExpiry `json:"expires_in"`
}

// azureTokenExpiry accepts the numeric and string expiry representations used
// by the different managed identity endpoints. Azure ML can also return a date
// string, which is parsed only for that environment.
type azureTokenExpiry string

func (e *azureTokenExpiry) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var value string
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		*e = azureTokenExpiry(value)
		return nil
	}
	var value json.Number
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = azureTokenExpiry(value.String())
	return nil
}

func (e azureTokenExpiry) Int64() (int64, error) {
	return strconv.ParseInt(string(e), 10, 64)
}

func (e azureTokenExpiry) String() string {
	return string(e)
}

func (token msiToken) Token() (*oauth2.Token, error) {
	return token.oauth2Token(time.Time{}, false, false)
}

func (token msiToken) managedIdentityToken(requestTime time.Time, parseAzureMLExpiry bool) (*oauth2.Token, error) {
	return token.oauth2Token(requestTime, true, parseAzureMLExpiry)
}

func (token msiToken) oauth2Token(requestTime time.Time, allowExpiresIn, parseAzureMLExpiry bool) (*oauth2.Token, error) {
	if token.AccessToken == "" {
		return nil, fmt.Errorf("token parse: %w", errInvalidToken)
	}
	expiry, err := token.expiry(requestTime, allowExpiresIn, parseAzureMLExpiry)
	if err != nil {
		return nil, err
	}
	tokenType := token.TokenType
	if tokenType == "" {
		tokenType = "Bearer"
	}
	return &oauth2.Token{
		TokenType:    tokenType,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expiry:       expiry,
	}, nil
}

func (token msiToken) expiry(requestTime time.Time, allowExpiresIn, parseAzureMLExpiry bool) (time.Time, error) {
	if token.ExpiresOn != "" {
		epoch, err := token.ExpiresOn.Int64()
		if err == nil {
			return time.Unix(epoch, 0), nil
		}
		if parseAzureMLExpiry {
			expiry, parseErr := parseAzureMLTokenExpiry(token.ExpiresOn.String())
			if parseErr == nil {
				return expiry, nil
			}
			return time.Time{}, fmt.Errorf("%w: %s", errInvalidTokenExpiry, parseErr)
		}
		return time.Time{}, fmt.Errorf("%w: %s", errInvalidTokenExpiry, err)
	}
	if allowExpiresIn && token.ExpiresIn != "" {
		seconds, err := token.ExpiresIn.Int64()
		if err != nil {
			return time.Time{}, fmt.Errorf("%w: %s", errInvalidTokenExpiry, err)
		}
		return requestTime.Add(time.Duration(seconds) * time.Second), nil
	}
	return time.Time{}, fmt.Errorf("%w: expires_on is missing", errInvalidTokenExpiry)
}

func parseAzureMLTokenExpiry(value string) (time.Time, error) {
	const utcSuffix = " +00:00"
	if !strings.HasSuffix(value, utcSuffix) {
		return time.Time{}, fmt.Errorf("unsupported Azure ML token expiry %q", value)
	}
	value = strings.TrimSuffix(value, utcSuffix)
	formats := []string{
		"01/02/2006 15:04:05",
		"1/2/2006 15:04:05",
		"01/02/2006 03:04:05 PM",
		"1/2/2006 3:04:05 PM",
	}
	for _, format := range formats {
		if expiry, err := time.ParseInLocation(format, value, time.UTC); err == nil {
			return expiry, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported Azure ML token expiry %q", value+utcSuffix)
}
