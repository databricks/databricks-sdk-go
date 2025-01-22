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

type CachedTokenSourceOptions struct {
	// DisableAsyncRefresh disables the asynchronous token refresh.
	DisableAsyncRefresh bool

	// StaleDuration is the duration before the token expires. If unset, the
	// default duration of 3 minutes is used.
	StaleDuration time.Duration
}

func (ctso *CachedTokenSourceOptions) disableAsyncRefresh() bool {
	if ctso == nil {
		return defaultDisableAsyncRefresh
	}
	return ctso.DisableAsyncRefresh
}

func (ctso *CachedTokenSourceOptions) staleDuration() time.Duration {
	if ctso == nil || ctso.StaleDuration == 0 {
		return defaultStaleDuration
	}
	return ctso.StaleDuration
}

// NewCachedTokenProvider returns a new token source that caches the token.
//
// The TokenSource is expected to take care of potential retries on its own.
//
// If the TokenSource is already a cached token source, it is returned as is.
func NewCachedTokenSource(ts oauth2.TokenSource, opts *CachedTokenSourceOptions) oauth2.TokenSource {
	if cts, ok := ts.(*cachedTokenSource); ok {
		return cts
	}

	return &cachedTokenSource{
		tokenSource:   ts,
		staleDuration: opts.staleDuration(),
		disableAsync:  opts.disableAsyncRefresh(),
		timeNow:       time.Now,
	}
}

type cachedTokenSource struct {
	tokenSource   oauth2.TokenSource
	staleDuration time.Duration
	disableAsync  bool

	mu           sync.Mutex
	cachedToken  *oauth2.Token
	isRefreshing bool
	refreshErr   error

	// timeNow is a function that returns the current time. It is used to
	// determine the current time in tests.
	timeNow func() time.Time
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

// blockingToken returns a token from the cache or fetches a new one if the
// current token is expired. The function guarantees that only one refresh call
// we be made if several goroutines are calling it concurrently.
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

// triggerAsyncRefresh
func (cts *cachedTokenSource) triggerAsyncRefresh() {
	cts.mu.Lock()
	defer cts.mu.Unlock()
	if !cts.isRefreshing && cts.refreshErr == nil {
		go cts.asyncRefresh()
	}
}

func (cts *cachedTokenSource) asyncRefresh() {
	cts.mu.Lock()
	cts.isRefreshing = true
	cts.mu.Unlock()

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
