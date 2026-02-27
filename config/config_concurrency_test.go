package config

import (
	"context"
	"net/http"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/config/credentials"
)

// TestAuthenticateConcurrency verifies that Config.Authenticate is safe for
// concurrent use and does not trigger race detector warnings.
func TestAuthenticateConcurrency(t *testing.T) {
	cfg := &Config{
		Host:  "https://test.cloud.databricks.com",
		Token: "fake-token",
	}

	// Ensure config is resolved first.
	if err := cfg.EnsureResolved(); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	const numGoroutines = 10

	// Launch multiple goroutines that all try to authenticate concurrently.
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			req, _ := http.NewRequest("GET", "https://test.com", nil)
			// This should not panic or cause data races.
			_ = cfg.Authenticate(req)
		}()
	}

	wg.Wait()

	// Verify authentication happened exactly once.
	if cfg.credentialsProvider == nil {
		t.Error("Expected credentialsProvider to be initialized")
	}
}

// TestAuthenticateIfNeededConcurrency tests authenticateIfNeeded directly.
func TestAuthenticateIfNeededConcurrency(t *testing.T) {
	cfg := &Config{
		Host:  "https://test.cloud.databricks.com",
		Token: "fake-token",
	}

	// Ensure config is resolved first.
	if err := cfg.EnsureResolved(); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	const numGoroutines = 100

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = cfg.authenticateIfNeeded()
		}()
	}

	wg.Wait()

	// The important thing is that we didn't panic or hit data races.
}

// TestAuthenticateOnce verifies that authentication happens exactly once
// even with concurrent calls.
func TestAuthenticateOnce(t *testing.T) {
	callCount := 0
	var mu sync.Mutex

	cfg := &Config{
		Host: "https://test.cloud.databricks.com",
		Credentials: &testStrategy{
			configure: func(ctx context.Context, c *Config) (credentials.CredentialsProvider, error) {
				mu.Lock()
				callCount++
				mu.Unlock()
				return &testProvider{}, nil
			},
		},
	}

	if err := cfg.EnsureResolved(); err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	const numGoroutines = 50

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = cfg.authenticateIfNeeded()
		}()
	}

	wg.Wait()

	mu.Lock()
	count := callCount
	mu.Unlock()

	if count != 1 {
		t.Errorf("Expected Configure to be called exactly once, but was called %d times", count)
	}
}

// Test helpers.
type testStrategy struct {
	configure func(ctx context.Context, c *Config) (credentials.CredentialsProvider, error)
}

func (ts *testStrategy) Name() string {
	return "test"
}

func (ts *testStrategy) Configure(ctx context.Context, c *Config) (credentials.CredentialsProvider, error) {
	if ts.configure != nil {
		return ts.configure(ctx, c)
	}
	return &testProvider{}, nil
}

type testProvider struct{}

func (tp *testProvider) SetHeaders(r *http.Request) error {
	r.Header.Set("Authorization", "Bearer test-token")
	return nil
}
