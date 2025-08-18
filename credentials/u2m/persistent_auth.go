package u2m

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	cache "github.com/databricks/databricks-sdk-go/credentials/u2m/cache"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/authhandler"
)

const (
	// appClientId is the default client ID used by the SDK for U2M OAuth.
	appClientID = "databricks-cli"

	// defaultPort is the default port for the OAuth2 callback server. If the
	// port is already in use, the next port is tried (8021, 8022, etc.).
	defaultPort = 8020

	// maxPortFallback is the maximum port to try when using the fallback
	// mechanism.
	maxPortFallback = 8040

	// listenerTimeout is the maximum duration spent trying to acquire a
	// listener (including port selection).
	listenerTimeout = 45 * time.Second
)

var (
	// Internal errors used for testing.
	errListenerTimeout = errors.New("failed to listen on any port: timeout")
	errNoPortAvailable = errors.New("no port available to listen on")
)

// PersistentAuth is an OAuth manager that handles the U2M OAuth flow. Tokens
// are stored in and looked up from the provided cache. Tokens include the
// refresh token. On load, if the access token is expired, it is refreshed
// using the refresh token.
//
// The PersistentAuth is safe for concurrent use. The token cache is locked
// during token retrieval, refresh and storage.
type PersistentAuth struct {
	// cache is the token cache to store and lookup tokens.
	cache cache.TokenCache

	// client is the HTTP client to use for OAuth2 requests.
	client *http.Client

	// endpointSupplier is the HTTP endpointSupplier to use for OAuth2 requests.
	endpointSupplier OAuthEndpointSupplier

	// oAuthArgument defines the workspace or account to authenticate to and the
	// cache key for the token.
	oAuthArgument OAuthArgument

	// browser is the function to open a URL in the default browser.
	browser func(url string) error

	// ln is the listener for the OAuth2 callback server.
	ln net.Listener

	// ctx is the context to use for underlying operations. This is needed in
	// order to implement the oauth2.TokenSource interface.
	ctx context.Context

	// redirectAddr is the redirect address for OAuth2 callbacks. The value is
	// set to localhost:PORT by startListener which will dynamically assign a
	// random port. If a value is already provided, it will be used instead
	// (e.g. for testing).
	redirectAddr string

	// Optional port to use for the OAuth2 callback server. If set to 0, the
	// default port with fallback is used. This means that setting a port will
	// disable the fallback mechanism.
	port int

	// netListen is an optional function to listen on a TCP address. If not set,
	// it will use net.Listen by default. This is useful for testing.
	netListen func(network, address string) (net.Listener, error)
}

type PersistentAuthOption func(*PersistentAuth)

// WithTokenCache sets the token cache for the PersistentAuth.
func WithTokenCache(c cache.TokenCache) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.cache = c
	}
}

// WithHttpClient sets the HTTP client for the PersistentAuth.
func WithHttpClient(c *http.Client) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.client = c
	}
}

// WithOAuthEndpointSupplier sets the OAuth endpoint supplier for the
// PersistentAuth.
func WithOAuthEndpointSupplier(c OAuthEndpointSupplier) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.endpointSupplier = c
	}
}

// WithOAuthArgument sets the OAuthArgument for the PersistentAuth.
func WithOAuthArgument(arg OAuthArgument) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.oAuthArgument = arg
	}
}

// WithBrowser sets the browser function for the PersistentAuth.
func WithBrowser(b func(url string) error) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.browser = b
	}
}

// WithPort sets the port for the PersistentAuth.
func WithPort(port int) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.port = port
	}
}

// NewPersistentAuth creates a new PersistentAuth with the provided options.
func NewPersistentAuth(ctx context.Context, opts ...PersistentAuthOption) (*PersistentAuth, error) {
	p := &PersistentAuth{}
	for _, opt := range opts {
		opt(p)
	}
	// By default, PersistentAuth uses the default ApiClient to make HTTP
	// requests. Furthermore, if the endpointSupplier is not provided, it uses
	// this same client to fetch the OAuth endpoints. If the HTTP client is
	// provided but the endpointSupplier is not, we construct a default
	// ApiClient for use with BasicOAuthClient.
	apiClient := httpclient.NewApiClient(httpclient.ClientConfig{})
	if p.client == nil {
		p.client = &http.Client{
			Transport: apiClient,
			// 30 seconds matches the default timeout of the ApiClient
			Timeout: 30 * time.Second,
		}
	}
	if p.endpointSupplier == nil {
		p.endpointSupplier = &BasicOAuthEndpointSupplier{
			Client: apiClient,
		}
	}
	if p.cache == nil {
		var err error
		p.cache, err = cache.NewFileTokenCache()
		if err != nil {
			return nil, fmt.Errorf("cache: %w", err)
		}
	}
	if err := p.validateArg(); err != nil {
		return nil, err
	}
	if p.browser == nil {
		p.browser = browser.OpenURL
	}
	p.ctx = ctx
	return p, nil
}

// Token loads the OAuth2 token for the given OAuthArgument from the cache. If
// the token is expired, it is refreshed using the refresh token.
func (a *PersistentAuth) Token() (t *oauth2.Token, err error) {
	err = a.startListener(a.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting listener: %w", err)
	}
	defer a.Close()

	key := a.oAuthArgument.GetCacheKey()
	t, err = a.cache.Lookup(key)
	if err != nil {
		return nil, fmt.Errorf("cache: %w", err)
	}
	// refresh if invalid
	if !t.Valid() {
		t, err = a.refresh(t)
		if err != nil {
			return nil, fmt.Errorf("token refresh: %w", err)
		}
	}

	// Do not include the refresh token for security reasons. Refresh tokens are
	// long-lived credentials that we do not want to expose unnecessarily.
	t.RefreshToken = ""
	return t, nil
}

// refresh refreshes the token for the given OAuthArgument, storing the new
// token in the cache.
func (a *PersistentAuth) refresh(oldToken *oauth2.Token) (*oauth2.Token, error) {
	// OAuth2 config is invoked only for expired tokens to speed up
	// the happy path in the token retrieval
	cfg, err := a.oauth2Config()
	if err != nil {
		return nil, err
	}
	// make OAuth2 library use our client
	ctx := a.setOAuthContext(a.ctx)
	// eagerly refresh token
	t, err := cfg.TokenSource(ctx, oldToken).Token()
	if err != nil {
		// The default RoundTripper of our httpclient.ApiClient returns an error
		// if the response status code is not 2xx. This isn't compliant with the
		// RoundTripper interface, so this error isn't handled by the oauth2
		// library. We need to handle it here.
		var internalHttpError *httpclient.HttpError
		if errors.As(err, &internalHttpError) {
			// error fields
			// https://datatracker.ietf.org/doc/html/rfc6749#section-5.2
			var errResponse struct {
				Error            string `json:"error"`
				ErrorDescription string `json:"error_description"`
			}
			if unmarshalErr := json.Unmarshal([]byte(internalHttpError.Message), &errResponse); unmarshalErr != nil {
				return nil, fmt.Errorf("unmarshal: %w", unmarshalErr)
			}
			// Invalid refresh tokens get their own error type so they can be
			// better presented to users.
			if errResponse.ErrorDescription == "Refresh token is invalid" {
				return nil, &InvalidRefreshTokenError{err}
			}
			return nil, fmt.Errorf("%s (error code: %s)", errResponse.ErrorDescription, errResponse.Error)
		}

		// Handle responses from well-behaved *http.Client implementations.
		var httpErr *oauth2.RetrieveError
		if errors.As(err, &httpErr) {
			// Invalid refresh tokens get their own error type so they can be
			// better presented to users.
			if httpErr.ErrorDescription == "Refresh token is invalid" {
				return nil, &InvalidRefreshTokenError{err}
			}
			return nil, fmt.Errorf("%s (error code: %s)", httpErr.ErrorDescription, httpErr.ErrorCode)
		}
		return nil, err
	}
	err = a.cache.Store(a.oAuthArgument.GetCacheKey(), t)
	if err != nil {
		return nil, fmt.Errorf("cache update: %w", err)
	}
	return t, nil
}

// Challenge initiates the OAuth2 login flow for the given OAuthArgument. The
// OAuth2 flow is started by opening the browser to the OAuth2 authorization
// URL. The user is redirected to the callback server on appRedirectAddr. The
// callback server listens for the redirect from the identity provider and
// exchanges the authorization code for an access token.
func (a *PersistentAuth) Challenge() error {
	err := a.startListener(a.ctx)
	if err != nil {
		return fmt.Errorf("starting listener: %w", err)
	}
	// The listener will be closed by the callback server automatically, but if
	// the callback server is not created, we need to close the listener manually.
	defer a.Close()

	cfg, err := a.oauth2Config()
	if err != nil {
		return fmt.Errorf("fetching oauth config: %w", err)
	}
	cb, err := a.newCallbackServer()
	if err != nil {
		return fmt.Errorf("callback server: %w", err)
	}
	defer cb.Close()

	state, pkce, err := a.stateAndPKCE()
	if err != nil {
		return fmt.Errorf("state and pkce: %w", err)
	}
	// make OAuth2 library use our client
	ctx := a.setOAuthContext(a.ctx)
	ts := authhandler.TokenSourceWithPKCE(ctx, cfg, state, cb.Handler, pkce)
	t, err := ts.Token()
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	// cache token identified by host (and possibly the account id)
	err = a.cache.Store(a.oAuthArgument.GetCacheKey(), t)
	if err != nil {
		return fmt.Errorf("store: %w", err)
	}
	return nil
}

// startListener starts a listener on appRedirectAddr, retrying if the address
// is already in use.
func (pa *PersistentAuth) startListener(ctx context.Context) error {
	if pa.port != 0 { // if port is set, use it
		return pa.startListenerWithPort(pa.port)
	}
	return pa.startListenerWithFallback(ctx)
}

// startListenerWithFallback starts a listener that will try to find a free
// port to listen on starting from the default port and incrementing by 1 until
// a free port is found.
func (pa *PersistentAuth) startListenerWithFallback(ctx context.Context) error {
	startTime := time.Now()
	for port := defaultPort; port <= maxPortFallback; port++ {
		if time.Since(startTime) > listenerTimeout {
			return errListenerTimeout
		}
		if err := pa.startListenerWithPort(port); err != nil {
			logger.Debugf(ctx, "failed to listen on %d: %v, retrying", port, err)
			continue
		}
		logger.Debugf(ctx, "OAuth callback server listening on %s", pa.redirectAddr)
		return nil
	}
	return errNoPortAvailable
}

func (pa *PersistentAuth) startListenerWithPort(port int) error {
	addr := fmt.Sprintf("localhost:%d", port)
	listener, err := pa.listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", addr, err)
	}
	pa.ln = listener
	pa.redirectAddr = addr
	return nil
}

func (pa *PersistentAuth) listen(network, addr string) (net.Listener, error) {
	if pa.netListen != nil {
		return pa.netListen(network, addr)
	} else {
		return net.Listen(network, addr)
	}
}

func (a *PersistentAuth) Close() error {
	if a.ln == nil {
		return nil
	}
	return a.ln.Close()
}

// validateArg ensures that the OAuthArgument is either a WorkspaceOAuthArgument
// or an AccountOAuthArgument.
func (a *PersistentAuth) validateArg() error {
	if a.oAuthArgument == nil {
		return errors.New("missing OAuthArgument")
	}
	_, isWorkspaceArg := a.oAuthArgument.(WorkspaceOAuthArgument)
	_, isAccountArg := a.oAuthArgument.(AccountOAuthArgument)
	if !isWorkspaceArg && !isAccountArg {
		return fmt.Errorf("unsupported OAuthArgument type: %T, must implement either WorkspaceOAuthArgument or AccountOAuthArgument interface", a.oAuthArgument)
	}
	return nil
}

// oauth2Config returns the OAuth2 configuration for the given OAuthArgument.
func (a *PersistentAuth) oauth2Config() (*oauth2.Config, error) {
	scopes := []string{
		"offline_access", // ensures OAuth token includes refresh token
		"all-apis",       // ensures OAuth token has access to all control-plane APIs
	}
	var endpoints *OAuthAuthorizationServer
	var err error
	switch argg := a.oAuthArgument.(type) {
	case WorkspaceOAuthArgument:
		endpoints, err = a.endpointSupplier.GetWorkspaceOAuthEndpoints(a.ctx, argg.GetWorkspaceHost())
	case AccountOAuthArgument:
		endpoints, err = a.endpointSupplier.GetAccountOAuthEndpoints(
			a.ctx, argg.GetAccountHost(), argg.GetAccountId())
	default:
		return nil, fmt.Errorf("unsupported OAuthArgument type: %T, must implement either WorkspaceOAuthArgument or AccountOAuthArgument interface", a.oAuthArgument)
	}
	if err != nil {
		return nil, fmt.Errorf("fetching OAuth endpoints: %w", err)
	}
	return &oauth2.Config{
		ClientID: appClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:   endpoints.AuthorizationEndpoint,
			TokenURL:  endpoints.TokenEndpoint,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: fmt.Sprintf("http://%s", a.redirectAddr),
		Scopes:      scopes,
	}, nil
}

func (a *PersistentAuth) stateAndPKCE() (string, *authhandler.PKCEParams, error) {
	verifier, err := a.randomString(64)
	if err != nil {
		return "", nil, fmt.Errorf("verifier: %w", err)
	}
	verifierSha256 := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(verifierSha256[:])
	state, err := a.randomString(16)
	if err != nil {
		return "", nil, fmt.Errorf("state: %w", err)
	}
	return state, &authhandler.PKCEParams{
		Challenge:       challenge,
		ChallengeMethod: "S256",
		Verifier:        verifier,
	}, nil
}

func (a *PersistentAuth) randomString(size int) (string, error) {
	raw := make([]byte, size)
	// ignore error as rand.Reader never returns an error
	_, err := rand.Read(raw)
	if err != nil {
		return "", fmt.Errorf("rand.Read: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(raw), nil
}

func (a *PersistentAuth) setOAuthContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, oauth2.HTTPClient, a.client)
}

var _ oauth2.TokenSource = (*PersistentAuth)(nil)
