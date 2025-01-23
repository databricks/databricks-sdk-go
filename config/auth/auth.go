package auth

import (
	"sync"
	"time"

	"golang.org/x/oauth2"
)

const (
	// Default duration for the stale period. The number as been set arbitrarily
	// and might be changed in the future.
	defaultStaleDuration = 3 * time.Minute

	// Disable the asynchronous token refresh by default. This is meant to
	// change in the future once the feature is stable.
	defaultDisableAsyncRefresh = true
)

type Option func(*cachedTokenSource)

// WithCachedToken sets the initial token to be used by a cached token source.
func WithCachedToken(t *oauth2.Token) Option {
	return func(cts *cachedTokenSource) {
		cts.cachedToken = t
	}
}

// WithStaleDuration sets the duration for which a token is considered stale.
// Stale tokens are still valid but will trigger an asynchronous refresh if
// async refresh is enabled. The default value is 3 minutes.
func WithStaleDuration(d time.Duration) Option {
	return func(cts *cachedTokenSource) {
		cts.staleDuration = d
	}
}

// WithAsyncRefresh enables or disables the asynchronous token refresh.
func WithAsyncRefresh(b bool) Option {
	return func(cts *cachedTokenSource) {
		cts.disableAsync = !b
	}
}

// NewCachedTokenProvider returns a new token source that caches the token. The
// token is refreshed when it is expired or about to expire. The token is
// refreshed asynchronously if the async refresh is enabled.
//
// The token cache does not take care of retries in case the token source
// returns and error; it is the responsibility of the provided token source to
// handle retries appropriately.
//
// If the TokenSource is already a cached token source, it is returned as is.
func NewCachedTokenSource(ts oauth2.TokenSource, opts ...Option) oauth2.TokenSource {
	if cts, ok := ts.(*cachedTokenSource); ok {
		return cts
	}

	cts := &cachedTokenSource{
		tokenSource:   ts,
		staleDuration: defaultStaleDuration,
		disableAsync:  defaultDisableAsyncRefresh,
		cachedToken:   nil,
		timeNow:       time.Now,
	}

	for _, opt := range opts {
		opt(cts)
	}

	return cts
}

type cachedTokenSource struct {
	tokenSource   oauth2.TokenSource
	staleDuration time.Duration
	disableAsync  bool

	mu          sync.Mutex
	cachedToken *oauth2.Token

	// Indicates that an async refresh is in progress. This is used to prevent
	// multiple async refreshes from being triggered at the same time.
	isRefreshing bool

	// Error returned by the last async refresh. This is used to prevent
	// multiple async refreshes from being triggered.
	refreshErr error

	timeNow func() time.Time // for testing
}

// Token returns a token from the cache or fetches a new one if the current
// token is expired.
func (cts *cachedTokenSource) Token() (*oauth2.Token, error) {
	if cts.disableAsync {
		return cts.blockingToken()
	}
	return cts.asyncToken()
}

// tokenState represents the state of the token.
type tokenState int

const (
	fresh   tokenState = iota // The token is valid.
	stale                     // The token is valid but will expire soon.
	expired                   // The token has expired and cannot be used.
)

// tokenState returns the state of the token. The function is not thread-safe
// and should be called with the lock held.
func (c *cachedTokenSource) tokenState() tokenState {
	if c.cachedToken == nil {
		return expired
	}
	switch lifeSpan := c.cachedToken.Expiry.Sub(c.timeNow()); {
	case lifeSpan <= 0:
		return expired
	case lifeSpan <= c.staleDuration:
		return stale
	default:
		return fresh
	}
}

func (cts *cachedTokenSource) asyncToken() (*oauth2.Token, error) {
	cts.mu.Lock()
	ts := cts.tokenState()
	cts.mu.Unlock()

	switch ts {
	case fresh:
		cts.mu.Lock()
		defer cts.mu.Unlock()
		return cts.cachedToken, nil
	case stale:
		cts.triggerAsyncRefresh()
		cts.mu.Lock()
		defer cts.mu.Unlock()
		return cts.cachedToken, nil
	default: // expired
		return cts.blockingToken()
	}
}

func (cts *cachedTokenSource) blockingToken() (*oauth2.Token, error) {
	cts.mu.Lock()

	// The lock is kept for the entire operation to ensure that only one
	// blockingToken operation is running at a time.
	defer cts.mu.Unlock()

	cts.isRefreshing = false
	if ts := cts.tokenState(); ts != expired { // fresh or stale
		return cts.cachedToken, nil
	}

	t, err := cts.tokenSource.Token()
	if err != nil {
		return nil, err
	}
	cts.cachedToken = t
	return t, nil
}

func (cts *cachedTokenSource) triggerAsyncRefresh() {
	cts.mu.Lock()
	defer cts.mu.Unlock()
	if !cts.isRefreshing && cts.refreshErr == nil {
		cts.isRefreshing = true
		go cts.asyncRefresh()
	}
}

func (cts *cachedTokenSource) asyncRefresh() {
	t, err := cts.tokenSource.Token()

	cts.mu.Lock()
	defer cts.mu.Unlock()
	cts.isRefreshing = false
	if err != nil {
		cts.refreshErr = err
		return
	}
	cts.cachedToken = t
}
