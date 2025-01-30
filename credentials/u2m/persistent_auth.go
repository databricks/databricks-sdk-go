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
	"time"

	cache "github.com/databricks/databricks-sdk-go/credentials/u2m/cache"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/retries"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/authhandler"
)

const (
	// appClientId is the default client ID used by the SDK for U2M OAuth.
	appClientID = "databricks-cli"

	// appRedirectAddr is the default address for the OAuth2 callback server.
	appRedirectAddr = "localhost:8020"

	// listenerTimeout is the maximum amount of time to acquire listener on
	// appRedirectAddr.
	listenerTimeout = 45 * time.Second
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
	client OAuthClient
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
}

type PersistentAuthOption func(*PersistentAuth)

// WithTokenCache sets the token cache for the PersistentAuth.
func WithTokenCache(c cache.TokenCache) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.cache = c
	}
}

// WithApiClient sets the HTTP client for the PersistentAuth.
func WithOAuthClient(c OAuthClient) PersistentAuthOption {
	return func(a *PersistentAuth) {
		a.client = c
	}
}

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

// NewPersistentAuth creates a new PersistentAuth with the provided options.
func NewPersistentAuth(ctx context.Context, opts ...PersistentAuthOption) (*PersistentAuth, error) {
	p := &PersistentAuth{}
	for _, opt := range opts {
		opt(p)
	}
	if p.client == nil {
		p.client = &BasicOAuthClient{
			Client: httpclient.NewApiClient(httpclient.ClientConfig{}),
		}
	}
	if p.cache == nil {
		var err error
		p.cache, err = cache.NewFileTokenCache()
		if err != nil {
			return nil, fmt.Errorf("cache: %w", err)
		}
	}
	if p.oAuthArgument == nil {
		return nil, errors.New("missing OAuthArgument")
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
	// do not print refresh token to end-user
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
// exchanges the authorization code for an access token. It returns the OAuth2
// token on success.
func (a *PersistentAuth) Challenge() (*oauth2.Token, error) {
	err := a.startListener(a.ctx)
	if err != nil {
		return nil, fmt.Errorf("starting listener: %w", err)
	}
	// The listener will be closed by the callback server automatically, but if
	// the callback server is not created, we need to close the listener manually.
	defer a.Close()

	cfg, err := a.oauth2Config()
	if err != nil {
		return nil, fmt.Errorf("fetching oauth config: %w", err)
	}
	cb, err := a.newCallbackServer()
	if err != nil {
		return nil, fmt.Errorf("callback server: %w", err)
	}
	defer cb.Close()

	state, pkce, err := a.stateAndPKCE()
	if err != nil {
		return nil, fmt.Errorf("state and pkce: %w", err)
	}
	// make OAuth2 library use our client
	ctx := a.setOAuthContext(a.ctx)
	ts := authhandler.TokenSourceWithPKCE(ctx, cfg, state, cb.Handler, pkce)
	t, err := ts.Token()
	if err != nil {
		return nil, fmt.Errorf("authorize: %w", err)
	}
	// cache token identified by host (and possibly the account id)
	err = a.cache.Store(a.oAuthArgument.GetCacheKey(), t)
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
func (a *PersistentAuth) validateArg() error {
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
		endpoints, err = a.client.GetWorkspaceOAuthEndpoints(a.ctx, argg.GetWorkspaceHost())
	case AccountOAuthArgument:
		endpoints, err = a.client.GetAccountOAuthEndpoints(
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
		RedirectURL: fmt.Sprintf("http://%s", appRedirectAddr),
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
	return context.WithValue(ctx, oauth2.HTTPClient, a.client.GetHttpClient(ctx))
}

var _ oauth2.TokenSource = (*PersistentAuth)(nil)
