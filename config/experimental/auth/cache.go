package auth

import (
	"context"
	"sync"
	"time"

	"golang.org/x/oauth2"
)

// Maximum duration for the stale period. This value is chosen to provide
// robustness for standard OAuth tokens, supporting 99.95% availability,
// adding breathing room over the existing 99.99% SLA.
// For tokens with shorter lifetimes (e.g., FastPath tokens with 10-minute TTL),
// the stale period is computed as min(TTL × 0.5, maxStaleDuration) to adapt to
// the token's actual lifetime.
const maxStaleDuration = 20 * time.Minute

// computeStalePeriod calculates the stale period for a token based on its TTL.
// The stale period is the time before expiry when a token is considered stale
// and should be refreshed asynchronously.
//
// Formula: min(TTL × 0.5, maxStaleDuration)
//
// Edge cases:
//   - TTL <= 0 (expired or no expiry): returns maxStaleDuration.
//   - Very short TTL (e.g., 2 seconds): returns TTL / 2 without minimum enforcement.
//   - Standard OAuth (60 minutes): returns 20 minutes (capped at max).
//   - FastPath (10 minutes): returns 5 minutes.
func computeStalePeriod(ttl time.Duration) time.Duration {
	if ttl <= 0 {
		// Token is already expired or has no expiry (Expiry.IsZero()).
		// Use max stale duration as a safe default.
		return maxStaleDuration
	}
	stalePeriod := ttl / 2
	// Cap the stale period at the maximum stale duration.
	if stalePeriod > maxStaleDuration {
		return maxStaleDuration
	}
	return stalePeriod
}

type Option func(*cachedTokenSource)

// WithCachedToken sets the initial token to be used by a cached token source.
func WithCachedToken(t *oauth2.Token) Option {
	return func(cts *cachedTokenSource) {
		cts.cachedToken = t
	}
}

// WithAsyncRefresh enables or disables the asynchronous token refresh.
func WithAsyncRefresh(b bool) Option {
	return func(cts *cachedTokenSource) {
		cts.disableAsync = !b
	}
}

// NewCachedTokenSource wraps a [TokenSource] to cache the tokens it returns.
// By default, the cache will refresh tokens asynchronously a few minutes before
// they expire.
//
// The token cache is safe for concurrent use by multiple goroutines and will
// guarantee that only one token refresh is triggered at a time.
//
// The token cache does not take care of retries in case the token source
// returns and error; it is the responsibility of the provided token source to
// handle retries appropriately.
//
// If the TokenSource is already a cached token source (obtained by calling this
// function), it is returned as is.
func NewCachedTokenSource(ts TokenSource, opts ...Option) TokenSource {
	// This is meant as a niche optimization to avoid double caching of the
	// token source in situations where the user calls needs caching guarantees
	// but does not know if the token source is already cached.
	if cts, ok := ts.(*cachedTokenSource); ok {
		return cts
	}

	cts := &cachedTokenSource{
		tokenSource:   ts,
		staleDuration: maxStaleDuration,
		timeNow:       time.Now,
	}

	for _, opt := range opts {
		opt(cts)
	}

	// If an initial token was provided via WithCachedToken, compute its stale
	// period based on its TTL. This ensures the first call to tokenState() uses
	// the correct stale period instead of the default maximum.
	if cts.cachedToken != nil {
		ttl := cts.cachedToken.Expiry.Sub(cts.timeNow())
		cts.staleDuration = computeStalePeriod(ttl)
	}

	return cts
}

type cachedTokenSource struct {
	// The token source to obtain tokens from.
	tokenSource TokenSource

	// If true, only refresh the token with a blocking call when it is expired.
	disableAsync bool

	// Duration during which a token is considered stale, see tokenState.
	// This value is computed dynamically for each token based on its TTL at
	// acquisition time using the formula: min(TTL × 0.5, maxStaleDuration).
	// The value remains fixed for the lifetime of each token.
	staleDuration time.Duration

	mu          sync.Mutex
	cachedToken *oauth2.Token

	// Indicates that an async refresh is in progress. This is used to prevent
	// multiple async refreshes from being triggered at the same time.
	isRefreshing bool

	// Error returned by the last refresh. Async refreshes are disabled if this
	// value is not nil so that the cache does not continue sending request to
	// a potentially failing server. The next blocking call will re-enable async
	// refreshes by setting this value to nil if it succeeds, or return the
	// error if it fails.
	refreshErr error

	timeNow func() time.Time // for testing
}

// Token returns a token from the cache or fetches a new one if the current
// token is expired.
func (cts *cachedTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	if cts.disableAsync {
		return cts.blockingToken(ctx)
	}
	return cts.asyncToken(ctx)
}

// tokenState represents the state of the token. Each token can be in one of
// the following three states:
//   - fresh: The token is valid.
//   - stale: The token is valid but will expire soon.
//   - expired: The token has expired and cannot be used.
//
// Token state through time:
//
//	issue time     expiry time
//	    v               v
//	    | fresh | stale | expired -> time
//	    |     valid     |
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

func (cts *cachedTokenSource) asyncToken(ctx context.Context) (*oauth2.Token, error) {
	cts.mu.Lock()
	ts := cts.tokenState()
	t := cts.cachedToken
	cts.mu.Unlock()

	switch ts {
	case fresh:
		return t, nil
	case stale:
		cts.triggerAsyncRefresh(ctx)
		return t, nil
	default: // expired
		return cts.blockingToken(ctx)
	}
}

func (cts *cachedTokenSource) blockingToken(ctx context.Context) (*oauth2.Token, error) {
	cts.mu.Lock()

	// The lock is kept for the entire operation to ensure that only one
	// blockingToken operation is running at a time.
	defer cts.mu.Unlock()

	// This is important to recover from potential previous failed attempts
	// to refresh the token asynchronously, see declaration of refreshErr for
	// more information.
	cts.isRefreshing = false
	cts.refreshErr = nil

	// It's possible that the token got refreshed (either by a blockingToken or
	// an asyncRefresh call) while this particular call was waiting to acquire
	// the mutex. This check avoids refreshing the token again in such cases.
	if ts := cts.tokenState(); ts != expired { // fresh or stale
		return cts.cachedToken, nil
	}

	t, err := cts.tokenSource.Token(ctx)
	if err != nil {
		return nil, err
	}
	cts.cachedToken = t
	// Compute and store the stale period for this specific token based on its
	// TTL at acquisition time. This value remains fixed for the token's lifetime.
	ttl := t.Expiry.Sub(cts.timeNow())
	cts.staleDuration = computeStalePeriod(ttl)
	return t, nil
}

func (cts *cachedTokenSource) triggerAsyncRefresh(ctx context.Context) {
	cts.mu.Lock()
	defer cts.mu.Unlock()
	if !cts.isRefreshing && cts.refreshErr == nil {
		cts.isRefreshing = true

		go func() {
			t, err := cts.tokenSource.Token(ctx)

			cts.mu.Lock()
			defer cts.mu.Unlock()
			cts.isRefreshing = false
			if err != nil {
				cts.refreshErr = err
				return
			}
			cts.cachedToken = t
			// Compute and store the stale period for this specific token based on
			// its TTL at acquisition time. This value remains fixed for the token's
			// lifetime.
			ttl := t.Expiry.Sub(cts.timeNow())
			cts.staleDuration = computeStalePeriod(ttl)
		}()
	}
}
