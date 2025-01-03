package oauth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials/cache"
	"github.com/databricks/databricks-sdk-go/httpclient"
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

// PersistentAuth is an OAuth manager that handles the U2M OAuth flow. Tokens
// are stored in and looked up from the provided cache. Tokens include the
// refresh token. On load, if the access token is expired, it is refreshed
// using the refresh token.
type PersistentAuth struct {
	// Cache is the token cache to store and lookup tokens.
	cache cache.TokenCache
	// Locker is the lock to synchronize token cache access.
	locker sync.Locker
	// Client is the HTTP client to use for OAuth2 requests.
	client *httpclient.ApiClient
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
func WithApiClient(c *httpclient.ApiClient) PersistentAuthOption {
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

func NewPersistentAuth(ctx context.Context, opts ...PersistentAuthOption) (*PersistentAuth, error) {
	p := &PersistentAuth{}
	for _, opt := range opts {
		opt(p)
	}
	if p.client == nil {
		p.client = httpclient.NewApiClient(httpclient.ClientConfig{})
	}
	if p.cache == nil {
		p.cache = &cache.FileTokenCache{}
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

func (a *PersistentAuth) Load(ctx context.Context, arg OAuthArgument) (*oauth2.Token, error) {
	key := arg.GetCacheKey(ctx)
	t, err := a.cache.Lookup(key)
	if err != nil {
		return nil, fmt.Errorf("cache: %w", err)
	}
	// refresh if invalid
	if !t.Valid() {
		// OAuth2 config is invoked only for expired tokens to speed up
		// the happy path in the token retrieval
		cfg, err := a.oauth2Config(ctx, arg.GetHost(ctx), arg.GetAccountId(ctx))
		if err != nil {
			return nil, err
		}
		// make OAuth2 library use our client
		ctx = a.client.InContextForOAuth2(ctx)
		// eagerly refresh token
		t, err = cfg.TokenSource(ctx, t).Token()
		if err != nil {
			return nil, fmt.Errorf("token refresh: %w", err)
		}
		err = a.cache.Store(key, t)
		if err != nil {
			return nil, fmt.Errorf("cache refresh: %w", err)
		}
	}
	// do not print refresh token to end-user
	t.RefreshToken = ""
	return t, nil
}

func (a *PersistentAuth) Challenge(ctx context.Context, arg OAuthArgument) error {
	err := a.startListener(ctx)
	if err != nil {
		return fmt.Errorf("starting listener: %w", err)
	}
	cfg, err := a.oauth2Config(ctx, arg.GetHost(ctx), arg.GetAccountId(ctx))
	if err != nil {
		return fmt.Errorf("fetching oauth config: %w", err)
	}
	cb, err := a.newCallback(ctx, arg)
	if err != nil {
		return fmt.Errorf("callback server: %w", err)
	}
	defer cb.Close()
	state, pkce := a.stateAndPKCE()
	// make OAuth2 library use our client
	ctx = a.client.InContextForOAuth2(ctx)
	ts := authhandler.TokenSourceWithPKCE(ctx, cfg, state, cb.Handler, pkce)
	t, err := ts.Token()
	if err != nil {
		return fmt.Errorf("authorize: %w", err)
	}
	// cache token identified by host (and possibly the account id)
	err = a.cache.Store(arg.GetCacheKey(ctx), t)
	if err != nil {
		return fmt.Errorf("store: %w", err)
	}
	return nil
}

func (a *PersistentAuth) startListener(ctx context.Context) error {
	listener, err := retries.Poll(ctx, listenerTimeout,
		func() (*net.Listener, *retries.Err) {
			var lc net.ListenConfig
			l, err := lc.Listen(ctx, "tcp", appRedirectAddr)
			if err != nil {
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

func (a *PersistentAuth) oauth2Config(ctx context.Context, host string, accountId string) (*oauth2.Config, error) {
	// in this iteration of CLI, we're using all scopes by default,
	// because tools like CLI and Terraform do use all apis. This
	// decision may be reconsidered later, once we have a proper
	// taxonomy of all scopes ready and implemented.
	scopes := []string{
		"offline_access",
		"all-apis",
	}
	endpoints, err := a.client.GetOidcEndpoints(ctx, host, accountId)
	if err != nil {
		return nil, fmt.Errorf("oidc: %w", err)
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
