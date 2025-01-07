package oauth

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
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials/cache"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/authhandler"
)

const (
	// these values are predefined by Databricks as a public client
	// and is specific to this application only. Using these values
	// for other applications is not allowed.
	appClientID     = "databricks-cli"
	appRedirectAddr = "localhost:8020"

	// lockfile location
	lockFilePath = ".databricks/token-cache.lock"

	// maximum amount of time to acquire listener on appRedirectAddr
	listenerTimeout = 45 * time.Second
)

// OAuthClient provides the http functionality needed for interacting with the
// Databricks OAuth APIs.
type OAuthClient interface {
	// GetHttpClient returns an HTTP client for OAuth2 requests.
	GetHttpClient(context.Context) *http.Client

	// GetWorkspaceOAuthEndpoints returns the OAuth2 endpoints for the workspace.
	GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error)

	// GetAccountOAuthEndpoints returns the OAuth2 endpoints for the account.
	GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error)
}

type BasicOAuthClient struct {
	client *httpclient.ApiClient
}

func (c *BasicOAuthClient) GetHttpClient(_ context.Context) *http.Client {
	return c.client.ToHttpClient()
}

func (c *BasicOAuthClient) GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error) {
	return GetWorkspaceOAuthEndpoints(ctx, c.client, workspaceHost)
}

func (c *BasicOAuthClient) GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error) {
	return GetAccountOAuthEndpoints(ctx, accountHost, accountId)
}

// PersistentAuth is an OAuth manager that handles the U2M OAuth flow. Tokens
// are stored in and looked up from the provided cache. Tokens include the
// refresh token. On load, if the access token is expired, it is refreshed
// using the refresh token.
//
// The PersistentAuth is safe for concurrent use. The token cache is locked
// during token retrieval, refresh and storage.
type PersistentAuth struct {
	// Cache is the token cache to store and lookup tokens.
	cache cache.TokenCache
	// Locker is the lock to synchronize token cache access.
	locker sync.Locker
	// Client is the HTTP client to use for OAuth2 requests.
	client OAuthClient
	// Browser is the function to open a URL in the default browser.
	browser func(url string) error
	// ln is the listener for the OAuth2 callback server.
	ln net.Listener
}

type PersistentAuthOption func(*PersistentAuth)

// WithTokenCache sets the token cache for the PersistentAuth.
func WithTokenCache(c cache.TokenCache) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.cache = c
	}
}

// WithLocker sets the locker for the PersistentAuth.
func WithLocker(l sync.Locker) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.locker = l
	}
}

// WithApiClient sets the HTTP client for the PersistentAuth.
func WithOAuthClient(c OAuthClient) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.client = c
	}
}

// WithBrowser sets the browser function for the PersistentAuth.
func WithBrowser(b func(url string) error) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.browser = b
	}
}

// NewPersistentAuth creates a new PersistentAuth with the provided options.
func NewPersistentAuth(ctx context.Context, opts ...PersistentAuthOption) (*PersistentAuth, error) {
	p := &PersistentAuth{}
	for _, opt := range opts {
		opt(p)
	}
	if p.client == nil {
		p.client = &BasicOAuthClient{
			client: httpclient.NewApiClient(httpclient.ClientConfig{}),
		}
	}
	if p.cache == nil {
		var err error
		p.cache, err = cache.NewFileTokenCache()
		if err != nil {
			return nil, fmt.Errorf("cache: %w", err)
		}
	}
	if p.locker == nil {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("home: %w", err)
		}

		p.locker, err = newLocker(filepath.Join(home, lockFilePath))
		if err != nil {
			return nil, fmt.Errorf("locker: %w", err)
		}
	}
	if p.browser == nil {
		p.browser = browser.OpenURL
	}
	return p, nil
}

// tokenErrorResponse is the response from the OAuth2 token endpoint when an
// error occurs.
type tokenErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// Load loads the OAuth2 token for the given OAuthArgument from the cache. If
// the token is expired, it is refreshed using the refresh token.
func (a *PersistentAuth) Load(ctx context.Context, arg OAuthArgument) (t *oauth2.Token, err error) {
	a.locker.Lock()
	defer a.locker.Unlock()

	if err := a.validateArg(arg); err != nil {
		return nil, err
	}
	// TODO: remove this listener after several releases.
	err = a.startListener(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting listener: %w", err)
	}
	defer a.Close()

	key := arg.GetCacheKey(ctx)
	t, err = a.cache.Lookup(key)
	if err != nil {
		return nil, fmt.Errorf("cache: %w", err)
	}
	// refresh if invalid
	if !t.Valid() {
		t, err = a.refresh(ctx, arg, t)
		if err != nil {
			return nil, fmt.Errorf("token refresh: %w", err)
		}
	}
	// do not print refresh token to end-user
	t.RefreshToken = ""
	return t, nil
}

// refresh refreshes the token for the given OAuthArgument, storing the new
// token in the cache.
func (a *PersistentAuth) refresh(ctx context.Context, arg OAuthArgument, oldToken *oauth2.Token) (*oauth2.Token, error) {
	// OAuth2 config is invoked only for expired tokens to speed up
	// the happy path in the token retrieval
	cfg, err := a.oauth2Config(ctx, arg)
	if err != nil {
		return nil, err
	}
	// make OAuth2 library use our client
	ctx = a.setOAuthContext(ctx)
	// eagerly refresh token
	t, err := cfg.TokenSource(ctx, oldToken).Token()
	if err != nil {
		var httpErr *httpclient.HttpError
		if errors.As(err, &httpErr) {
			resp := &tokenErrorResponse{}
			err = json.Unmarshal([]byte(httpErr.Message), resp)
			if err != nil {
				return nil, fmt.Errorf("unexpected parsing token response: %w", err)
			}
			// Invalid refresh tokens get their own error type so they can be
			// better presented to users.
			if resp.ErrorDescription == "Refresh token is invalid" {
				return nil, &InvalidRefreshTokenError{err}
			} else {
				return nil, fmt.Errorf("unexpected error refreshing token: %s", resp.ErrorDescription)
			}
		}
		return nil, fmt.Errorf("token refresh: %w", err)
	}
	err = a.cache.Store(arg.GetCacheKey(ctx), t)
	if err != nil {
		return nil, fmt.Errorf("cache refresh: %w", err)
	}
	return t, nil
}

// Challenge initiates the OAuth2 login flow for the given OAuthArgument. The
// OAuth2 flow is started by opening the browser to the OAuth2 authorization
// URL. The user is redirected to the callback server on appRedirectAddr. The
// callback server listens for the redirect from the identity provider and
// exchanges the authorization code for an access token. It returns the OAuth2
// token on success.
func (a *PersistentAuth) Challenge(ctx context.Context, arg OAuthArgument) (*oauth2.Token, error) {
	a.locker.Lock()
	defer a.locker.Unlock()

	if err := a.validateArg(arg); err != nil {
		return nil, err
	}
	err := a.startListener(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting listener: %w", err)
	}
	// The listener will be closed by the callback server automatically, but if
	// the callback server is not created, we need to close the listener manually.
	defer a.Close()

	cfg, err := a.oauth2Config(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("fetching oauth config: %w", err)
	}
	cb, err := a.newCallbackServer(ctx, arg)
	if err != nil {
		return nil, fmt.Errorf("callback server: %w", err)
	}
	defer cb.Close()

	state, pkce := a.stateAndPKCE()
	// make OAuth2 library use our client
	ctx = a.setOAuthContext(ctx)
	ts := authhandler.TokenSourceWithPKCE(ctx, cfg, state, cb.Handler, pkce)
	t, err := ts.Token()
	if err != nil {
		return nil, fmt.Errorf("authorize: %w", err)
	}
	// cache token identified by host (and possibly the account id)
	err = a.cache.Store(arg.GetCacheKey(ctx), t)
	if err != nil {
		return nil, fmt.Errorf("store: %w", err)
	}
	return t, nil
}

// startListener starts a listener on appRedirectAddr, retrying if the address
// is already in use.
func (a *PersistentAuth) startListener(ctx context.Context) error {
	listener, err := retries.Poll(ctx, listenerTimeout,
		func() (*net.Listener, *retries.Err) {
			var lc net.ListenConfig
			l, err := lc.Listen(ctx, "tcp", appRedirectAddr)
			if err != nil {
				logger.Debugf(ctx, "failed to listen on %s: %v, retrying", appRedirectAddr, err)
				return nil, retries.Continue(err)
			}
			return &l, nil
		})
	if err != nil {
		return fmt.Errorf("listener: %w", err)
	}
	a.ln = *listener
	return nil
}

func (a *PersistentAuth) Close() error {
	if a.ln == nil {
		return nil
	}
	return a.ln.Close()
}

// validateArg ensures that the OAuthArgument is either a WorkspaceOAuthArgument
// or an AccountOAuthArgument.
func (a *PersistentAuth) validateArg(arg OAuthArgument) error {
	_, isWorkspaceArg := arg.(WorkspaceOAuthArgument)
	_, isAccountArg := arg.(AccountOAuthArgument)
	if !isWorkspaceArg && !isAccountArg {
		return fmt.Errorf("unsupported OAuthArgument type: %T, must implement either WorkspaceOAuthArgument or AccountOAuthArgument interface", arg)
	}
	return nil
}

// oauth2Config returns the OAuth2 configuration for the given OAuthArgument.
func (a *PersistentAuth) oauth2Config(ctx context.Context, arg OAuthArgument) (*oauth2.Config, error) {
	scopes := []string{
		"offline_access", // ensures OAuth token includes refresh token
		"all-apis",       // ensures OAuth token has access to all control-plane APIs
	}
	var endpoints *OAuthAuthorizationServer
	var err error
	switch argg := arg.(type) {
	case WorkspaceOAuthArgument:
		endpoints, err = a.client.GetWorkspaceOAuthEndpoints(ctx, argg.GetWorkspaceHost(ctx))
	case AccountOAuthArgument:
		endpoints, err = a.client.GetAccountOAuthEndpoints(
			ctx, argg.GetAccountHost(ctx), argg.GetAccountId(ctx))
	default:
		return nil, fmt.Errorf("unsupported OAuthArgument type: %T, must implement either WorkspaceOAuthArgument or AccountOAuthArgument interface", arg)
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
		RedirectURL: fmt.Sprintf("http://%s", appRedirectAddr),
		Scopes:      scopes,
	}, nil
}

func (a *PersistentAuth) stateAndPKCE() (string, *authhandler.PKCEParams) {
	verifier := a.randomString(64)
	verifierSha256 := sha256.Sum256([]byte(verifier))
	challenge := base64.RawURLEncoding.EncodeToString(verifierSha256[:])
	return a.randomString(16), &authhandler.PKCEParams{
		Challenge:       challenge,
		ChallengeMethod: "S256",
		Verifier:        verifier,
	}
}

func (a *PersistentAuth) randomString(size int) string {
	raw := make([]byte, size)
	_, _ = rand.Read(raw)
	return base64.RawURLEncoding.EncodeToString(raw)
}

func (a *PersistentAuth) setOAuthContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, oauth2.HTTPClient, a.client.GetHttpClient(ctx))
}
