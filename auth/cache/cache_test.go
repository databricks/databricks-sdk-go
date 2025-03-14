package cache

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/auth"
	"golang.org/x/oauth2"
)

func TestNewCachedTokenSource_noCaching(t *testing.T) {
	want := &cachedTokenSource{}
	got := NewCachedTokenSource(want, nil)
	if got != want {
		t.Errorf("NewCachedTokenSource() = %v, want %v", got, want)
	}
}

func TestNewCachedTokenSource_default(t *testing.T) {
	ts := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return nil, nil
	})

	got, ok := NewCachedTokenSource(ts).(*cachedTokenSource)
	if !ok {
		t.Fatalf("NewCachedTokenSource() = %T, want *cachedTokenSource", got)
	}

	if got.staleDuration != defaultStaleDuration {
		t.Errorf("NewCachedTokenSource() staleDuration = %v, want %v", got.staleDuration, defaultStaleDuration)
	}
	if got.disableAsync != false {
		t.Errorf("NewCachedTokenSource() disableAsync = %v, want %v", got.disableAsync, false)
	}
	if got.cachedToken != nil {
		t.Errorf("NewCachedTokenSource() cachedToken = %v, want nil", got.cachedToken)
	}
}

func TestNewCachedTokenSource_options(t *testing.T) {
	ts := auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
		return nil, nil
	})

	wantDisableAsync := false
	wantCachedToken := &oauth2.Token{Expiry: time.Unix(42, 0)}

	opts := []Option{
		WithAsyncRefresh(!wantDisableAsync),
		WithCachedToken(wantCachedToken),
	}

	got, ok := NewCachedTokenSource(ts, opts...).(*cachedTokenSource)
	if !ok {
		t.Fatalf("NewCachedTokenSource() = %T, want *cachedTokenSource", got)
	}

	if got.disableAsync != wantDisableAsync {
		t.Errorf("NewCachedTokenSource(): disableAsync = %v, want %v", got.disableAsync, wantDisableAsync)
	}
	if got.cachedToken != wantCachedToken {
		t.Errorf("NewCachedTokenSource(): cachedToken = %v, want %v", got.cachedToken, wantCachedToken)
	}
}

func TestCachedTokenSource_tokenState(t *testing.T) {
	now := time.Unix(1337, 0) // mock value for time.Now()

	testCases := []struct {
		token         *oauth2.Token
		staleDuration time.Duration
		want          tokenState
	}{
		{
			token:         nil,
			staleDuration: 10 * time.Minute,
			want:          expired,
		},
		{
			token: &oauth2.Token{
				Expiry: now.Add(-1 * time.Second),
			},
			staleDuration: 10 * time.Minute,
			want:          expired,
		},
		{
			token: &oauth2.Token{
				Expiry: now.Add(1 * time.Hour),
			},
			staleDuration: 10 * time.Minute,
			want:          fresh,
		},
		{
			token: &oauth2.Token{
				Expiry: now.Add(5 * time.Minute),
			},
			staleDuration: 10 * time.Minute,
			want:          stale,
		},
	}

	for _, tc := range testCases {
		cts := &cachedTokenSource{
			cachedToken:   tc.token,
			staleDuration: tc.staleDuration,
			disableAsync:  false,
			timeNow:       func() time.Time { return now },
		}

		got := cts.tokenState()

		if got != tc.want {
			t.Errorf("tokenState() = %v, want %v", got, tc.want)
		}
	}
}

func TestCachedTokenSource_Token(t *testing.T) {
	now := time.Unix(1337, 0) // mock value for time.Now()
	nTokenCalls := 10         // number of goroutines calling Token()
	testCases := []struct {
		desc         string        // description of the test case
		cachedToken  *oauth2.Token // token cached before calling Token()
		disableAsync bool          // whether are disabled or not
		refreshErr   error         // whether the cache was in error state

		returnedToken *oauth2.Token // token returned by the token source
		returnedError error         // error returned by the token source

		wantCalls int           // expected number of calls to the token source
		wantToken *oauth2.Token // expected token in the cache
	}{
		{
			desc:          "[Blocking] no cached token",
			disableAsync:  true,
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:          "[Blocking] expired cached token",
			disableAsync:  true,
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:         "[Blocking] fresh cached token",
			disableAsync: true,
			cachedToken:  &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:    0,
			wantToken:    &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:         "[Blocking] stale cached token",
			disableAsync: true,
			cachedToken:  &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			wantCalls:    0,
			wantToken:    &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
		},
		{
			desc:          "[Blocking] refresh error",
			disableAsync:  true,
			returnedError: fmt.Errorf("test error"),
			wantCalls:     10,
		},
		{
			desc:          "[Blocking] recover from error",
			disableAsync:  true,
			refreshErr:    fmt.Errorf("refresh error"),
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Minute)},
			returnedToken: &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
			wantCalls:     10,
			wantToken:     &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
		},
		{
			desc:          "[Async] no cached token",
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:          "[Async] no cached token",
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:          "[Async] expired cached token",
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:        "[Async] fresh cached token",
			cachedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:   0,
			wantToken:   &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:          "[Async] stale cached token",
			cachedToken:   &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			returnedToken: &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Hour)},
		},
		{
			desc:          "[Async] refresh error",
			cachedToken:   &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			returnedError: fmt.Errorf("test error"),
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
		},
		{
			desc:          "[Async] stale cached token, expired token returned",
			cachedToken:   &oauth2.Token{Expiry: now.Add(1 * time.Minute)},
			returnedToken: &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
			wantCalls:     1,
			wantToken:     &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
		},
		{
			desc:          "[Async] recover from error",
			refreshErr:    fmt.Errorf("refresh error"),
			cachedToken:   &oauth2.Token{Expiry: now.Add(-1 * time.Minute)},
			returnedToken: &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
			wantCalls:     10,
			wantToken:     &oauth2.Token{Expiry: now.Add(-1 * time.Hour)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			gotCalls := int32(0)
			cts := &cachedTokenSource{
				disableAsync:  tc.disableAsync,
				staleDuration: 10 * time.Minute,
				cachedToken:   tc.cachedToken,
				timeNow:       func() time.Time { return now },
				tokenSource: auth.TokenSourceFn(func(_ context.Context) (*oauth2.Token, error) {
					atomic.AddInt32(&gotCalls, 1)
					time.Sleep(10 * time.Millisecond)
					return tc.returnedToken, tc.returnedError
				}),
			}

			wg := sync.WaitGroup{}
			for i := 0; i < nTokenCalls; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					cts.Token(context.Background())
				}()
			}

			wg.Wait()

			// Wait for async refreshes to finish. This part is a little brittle
			// but necessary to ensure that the async refresh is done before
			// checking the results.
			time.Sleep(20 * time.Millisecond)

			if int(gotCalls) != tc.wantCalls {
				t.Errorf("want %d calls to cts.tokenSource.Token(), got %d", tc.wantCalls, gotCalls)
			}
			if !reflect.DeepEqual(tc.wantToken, cts.cachedToken) {
				t.Errorf("want cached token %v, got %v", tc.wantToken, cts.cachedToken)
			}
		})
	}
}
