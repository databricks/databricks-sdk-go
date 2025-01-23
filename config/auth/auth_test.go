package auth

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

type mockTokenSource func() (*oauth2.Token, error)

func (m mockTokenSource) Token() (*oauth2.Token, error) {
	return m()
}

func TestNewCachedTokenSource_noCaching(t *testing.T) {
	want := &cachedTokenSource{}
	got := NewCachedTokenSource(want, nil)
	if got != want {
		t.Errorf("NewCachedTokenSource() = %v, want %v", got, want)
	}
}

func TestNewCachedTokenSource(t *testing.T) {
	ts := mockTokenSource(func() (*oauth2.Token, error) {
		return nil, nil
	})

	testCases := []struct {
		options *CachedTokenSourceOptions
		want    *cachedTokenSource
	}{
		{
			options: nil,
			want: &cachedTokenSource{
				tokenSource:   ts,
				staleDuration: defaultStaleDuration,
				disableAsync:  defaultDisableAsyncRefresh,
			},
		},
		{
			options: &CachedTokenSourceOptions{},
			want: &cachedTokenSource{
				tokenSource:   ts,
				staleDuration: defaultStaleDuration,
				disableAsync:  false,
			},
		},
		{
			options: &CachedTokenSourceOptions{
				DisableAsyncRefresh: true,
			},
			want: &cachedTokenSource{
				tokenSource:   ts,
				staleDuration: defaultStaleDuration,
				disableAsync:  true,
			},
		},
		{
			options: &CachedTokenSourceOptions{
				StaleDuration: 5 * time.Minute,
			},
			want: &cachedTokenSource{
				tokenSource:   ts,
				staleDuration: 5 * time.Minute,
				disableAsync:  false,
			},
		},
	}

	for _, tc := range testCases {
		got, ok := NewCachedTokenSource(ts, tc.options).(*cachedTokenSource)
		if !ok {
			t.Fatalf("NewCachedTokenSource() = %T, want *cachedTokenSource", got)
		}

		if got.staleDuration != tc.want.staleDuration {
			t.Errorf("NewCachedTokenSource() staleDuration = %v, want %v", got.staleDuration, tc.want.staleDuration)
		}
		if got.disableAsync != tc.want.disableAsync {
			t.Errorf("NewCachedTokenSource() disableAsync = %v, want %v", got.disableAsync, tc.want.disableAsync)
		}
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
			wantCalls:     10,
			wantToken:     &oauth2.Token{Expiry: now.Add(-1 * time.Second)},
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
				tokenSource: mockTokenSource(func() (*oauth2.Token, error) {
					atomic.AddInt32(&gotCalls, 1)
					return tc.returnedToken, tc.returnedError
				}),
			}

			wg := sync.WaitGroup{}
			for i := 0; i < nTokenCalls; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					cts.Token()
				}()
			}
			wg.Wait()

			if int(gotCalls) != tc.wantCalls {
				t.Errorf("want %d calls to cts.tokenSource.Token(), got %d", tc.wantCalls, gotCalls)
			}
			if !reflect.DeepEqual(tc.wantToken, cts.cachedToken) {
				t.Errorf("want cached token %v, got %v", tc.wantToken, cts.cachedToken)
			}
		})
	}

}
