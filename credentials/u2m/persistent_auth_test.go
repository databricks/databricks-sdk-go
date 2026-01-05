package u2m

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"golang.org/x/oauth2"
)

type tokenCacheMock struct {
	store  func(key string, t *oauth2.Token) error
	lookup func(key string) (*oauth2.Token, error)
}

func (m *tokenCacheMock) Store(key string, t *oauth2.Token) error {
	return m.store(key, t)
}

func (m *tokenCacheMock) Lookup(key string) (*oauth2.Token, error) {
	return m.lookup(key)
}

func TestToken(t *testing.T) {
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != "https://abc/oidc/accounts/xyz" {
				t.Fatalf("lookup(): want key 'https://abc/oidc/accounts/xyz', got %s", key)
			}
			return &oauth2.Token{
				AccessToken: "bcd",
				Expiry:      time.Now().Add(1 * time.Minute),
			}, nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://abc", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(context.Background(), WithTokenCache(cache), WithOAuthArgument(arg))
	if err != nil {
		t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()

	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "bcd" {
		t.Errorf("p.Token(): want access token 'bcd', got %s", tok.AccessToken)
	}
	if tok.RefreshToken != "" {
		t.Errorf("p.Token(): want refresh token '', got %s", tok.RefreshToken)
	}
}

type MockOAuthEndpointSupplier struct{}

func (m MockOAuthEndpointSupplier) GetAccountOAuthEndpoints(ctx context.Context, accountHost string, accountId string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/accounts/%s/v1/authorize", accountHost, accountId),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/accounts/%s/v1/token", accountHost, accountId),
	}, nil
}

func (m MockOAuthEndpointSupplier) GetWorkspaceOAuthEndpoints(ctx context.Context, workspaceHost string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/v1/authorize", workspaceHost),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/v1/token", workspaceHost),
	}, nil
}

func (m MockOAuthEndpointSupplier) GetUnifiedOAuthEndpoints(ctx context.Context, host string, accountId string) (*OAuthAuthorizationServer, error) {
	return &OAuthAuthorizationServer{
		AuthorizationEndpoint: fmt.Sprintf("%s/oidc/accounts/%s/v1/authorize", host, accountId),
		TokenEndpoint:         fmt.Sprintf("%s/oidc/accounts/%s/v1/token", host, accountId),
	}, nil
}

func TestToken_RefreshesExpiredAccessToken(t *testing.T) {
	ctx := context.Background()
	expectedKey := "https://accounts.cloud.databricks.com/oidc/accounts/xyz"
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != expectedKey {
				t.Fatalf("lookup(): want key %s, got %s", expectedKey, key)
			}
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
		store: func(key string, tok *oauth2.Token) error {
			if key != expectedKey {
				t.Fatalf("store(): want key %s, got %s", expectedKey, key)
			}
			if tok.RefreshToken != "def" {
				t.Fatalf("store(): want refresh token 'def', got %s", tok.RefreshToken)
			}
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=refreshed&refresh_token=def`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	tok, err := p.Token()
	if err != nil {
		t.Fatalf("p.Token(): want no error, got %v", err)
	}
	if tok.AccessToken != "refreshed" {
		t.Errorf("p.Token(): want access token 'refreshed', got %s", tok.AccessToken)
	}
	if tok.RefreshToken != "" {
		t.Errorf("p.Token(): want refresh token '', got %s", tok.RefreshToken)
	}
}

func TestToken_ReturnsError(t *testing.T) {
	ctx := context.Background()
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != "https://accounts.cloud.databricks.com/oidc/accounts/xyz" {
				t.Fatalf("lookup(): want key 'https://accounts.cloud.databricks.com/oidc/accounts/xyz', got %s", key)
			}
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `{"error": "invalid_grant", "error_description": "Invalid Client"}`,
					Status:   401,
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()
	tok, err := p.Token()

	if tok != nil {
		t.Errorf("p.Token(): want nil, got %v", tok)
	}
	if !strings.Contains(err.Error(), "Invalid Client (error code: invalid_grant)") {
		t.Errorf("p.Token(): want error containing 'Invalid Client (error code: invalid_grant)', got %v", err)
	}
}

func TestToken_ReturnsInvalidRefreshTokenError(t *testing.T) {
	ctx := context.Background()
	cache := &tokenCacheMock{
		lookup: func(key string) (*oauth2.Token, error) {
			if key != "https://accounts.cloud.databricks.com/oidc/accounts/xyz" {
				t.Fatalf("lookup(): want key 'https://accounts.cloud.databricks.com/oidc/accounts/xyz', got %s", key)
			}
			return &oauth2.Token{
				AccessToken:  "expired",
				RefreshToken: "cde",
				Expiry:       time.Now().Add(-1 * time.Minute),
			}, nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}
	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `{"error": "invalid_grant", "error_description": "Refresh token is invalid"}`,
					Status:   401,
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()
	tok, err := p.Token()
	if tok != nil {
		t.Fatalf("p.Token(): want nil, got %v", tok)
	}
	target := &InvalidRefreshTokenError{}
	if !errors.As(err, &target) {
		t.Fatalf("p.Token(): want error of type InvalidRefreshTokenError, got %v", err)
	}
}

func TestChallenge(t *testing.T) {
	ctx := context.Background()

	browserOpened := make(chan string)
	browser := func(redirect string) error {
		u, err := url.ParseRequestURI(redirect)
		if err != nil {
			return err
		}
		if u.Path != "/oidc/accounts/xyz/v1/authorize" {
			t.Fatalf("browser(): want path '/oidc/accounts/xyz/v1/authorize', got %s", u.Path)
		}
		// for now we're ignoring asserting the fields of the redirect
		query := u.Query()
		browserOpened <- query.Get("state")
		return nil
	}
	cache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			if key != "https://accounts.cloud.databricks.com/oidc/accounts/xyz" {
				t.Fatalf("store(): want key 'https://accounts.cloud.databricks.com/oidc/accounts/xyz', got %s", key)
			}
			if tok.AccessToken != "__THAT__" {
				t.Fatalf("store(): want access token '__THAT__', got %s", tok.AccessToken)
			}
			if tok.RefreshToken != "__SOMETHING__" {
				t.Fatalf("store(): want refresh token '__SOMETHING__', got %s", tok.RefreshToken)
			}
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(cache),
		WithBrowser(browser),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=__THAT__&refresh_token=__SOMETHING__`,
					ResponseHeaders: map[string][]string{
						"Content-Type": {"application/x-www-form-urlencoded"},
					},
				},
			},
		}),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	errc := make(chan error)
	go func() {
		err := p.Challenge()
		errc <- err
		close(errc)
	}()

	state := <-browserOpened
	resp, err := http.Get(fmt.Sprintf("http://localhost:8020?code=__THIS__&state=%s", state))
	if err != nil {
		t.Fatalf("http.Get(): want no error, got %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("http.Get(): want status code 200, got %d", resp.StatusCode)
	}

	err = <-errc
	if err != nil {
		t.Fatalf("p.Challenge(): want no error, got %v", err)
	}
}

func TestChallenge_ReturnsErrorOnFailure(t *testing.T) {
	ctx := context.Background()
	browserOpened := make(chan string)
	browser := func(redirect string) error {
		u, err := url.ParseRequestURI(redirect)
		if err != nil {
			return err
		}
		if u.Path != "/oidc/accounts/xyz/v1/authorize" {
			t.Fatalf("browser(): want path '/oidc/accounts/xyz/v1/authorize', got %s", u.Path)
		}
		// for now we're ignoring asserting the fields of the redirect
		query := u.Query()
		browserOpened <- query.Get("state")
		return nil
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): want no error, got %v", err)
	}

	p, err := NewPersistentAuth(ctx, WithBrowser(browser), WithOAuthArgument(arg))
	if err != nil {
		t.Errorf("NewPersistentAuth(): want no error, got %v", err)
	}
	defer p.Close()

	errc := make(chan error)
	go func() {
		err := p.Challenge()
		errc <- err
		close(errc)
	}()

	<-browserOpened
	resp, err := http.Get("http://localhost:8020?error=access_denied&error_description=Policy%20evaluation%20failed%20for%20this%20request")
	if err != nil {
		t.Fatalf("http.Get(): want no error, got %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 400 {
		t.Fatalf("http.Get(): want status code 400, got %d", resp.StatusCode)
	}

	err = <-errc
	if err == nil {
		t.Fatalf("p.Challenge(): want error, got nil")
	}
	if !strings.Contains(err.Error(), "authorize: access_denied: Policy evaluation failed for this request") {
		t.Fatalf("p.Challenge(): want error containing 'authorize: access_denied: Policy evaluation failed for this request', got %v", err)
	}
}

func TestPersistentAuth_startListener_startFrom8020(t *testing.T) {
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != "localhost:8020" {
		t.Errorf("pa.redirectAddr should be localhost:8020, got %s", pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_incrementalFallBack(t *testing.T) {
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		if address == "localhost:8020" {
			return nil, fmt.Errorf("address already in use")
		}
		if address == "localhost:8021" {
			return nil, fmt.Errorf("address already in use")
		}
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != "localhost:8022" {
		t.Errorf("pa.redirectAddr should be localhost:8022, got %s", pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_noAvailablePort(t *testing.T) {
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		return nil, fmt.Errorf("address already in use")
	}

	gotErr := pa.startListener(context.Background())

	if !errors.Is(gotErr, errNoPortAvailable) {
		t.Fatalf("pa.startListener(): want error %v, got %v", errNoPortAvailable, gotErr)
	}
}

func TestPersistentAuth_startListener_maxPortFallbackIncluded(t *testing.T) {
	maxAddress := fmt.Sprintf("localhost:%d", maxPortFallback)
	pa := &PersistentAuth{}
	pa.netListen = func(_, address string) (net.Listener, error) {
		if address == maxAddress {
			return nil, nil
		}
		return nil, fmt.Errorf("address already in use")
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != maxAddress {
		t.Errorf("pa.redirectAddr should be %s, got %s", maxAddress, pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_explicitPort(t *testing.T) {
	explicitPort := 1337
	pa := &PersistentAuth{port: explicitPort}
	pa.netListen = func(_, address string) (net.Listener, error) {
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if gotErr != nil {
		t.Fatalf("pa.startListener(): want no error, got %v", gotErr)
	}
	if pa.redirectAddr != "localhost:1337" {
		t.Errorf("pa.redirectAddr should be localhost:1337, got %s", pa.redirectAddr)
	}
}

func TestPersistentAuth_startListener_explicitPortNoFallBack(t *testing.T) {
	testError := errors.New("test error")
	explicitPort := 1337
	pa := &PersistentAuth{port: explicitPort}
	pa.netListen = func(_, address string) (net.Listener, error) {
		if address == "localhost:1337" {
			return nil, testError
		}
		return nil, nil
	}

	gotErr := pa.startListener(context.Background())

	if !errors.Is(gotErr, testError) {
		t.Fatalf("pa.startListener(): want error %v, got %v", testError, gotErr)
	}
}

// TestU2M_ScopesAndOfflineAccess verifies that OAuth scopes are correctly configured
// and sent during the authorization flow, and that the disableOfflineAccess flag
// correctly controls whether offline_access is added to the scope.
func TestU2M_ScopesAndOfflineAccess(t *testing.T) {
	tests := []struct {
		name           string
		scopes         []string
		disableOffline bool
		want           string
	}{
		{
			name:           "single scope with offline_access",
			scopes:         []string{"dashboards"},
			disableOffline: false,
			want:           "dashboards offline_access",
		},
		{
			name:           "multiple scopes with offline_access",
			scopes:         []string{"files", "jobs", "mlflow:read"},
			disableOffline: false,
			want:           "files jobs mlflow:read offline_access",
		},
		{
			name:           "disable offline_access",
			scopes:         []string{"files", "jobs"},
			disableOffline: true,
			want:           "files jobs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			scopeReceived := make(chan string, 1)
			browserOpened := make(chan string, 1)
			browser := func(redirect string) error {
				u, err := url.ParseRequestURI(redirect)
				if err != nil {
					return err
				}
				query := u.Query()
				scopeReceived <- query.Get("scope")
				browserOpened <- query.Get("state")
				return nil
			}

			cache := &tokenCacheMock{
				store: func(key string, tok *oauth2.Token) error {
					return nil
				},
			}

			arg, err := NewBasicWorkspaceOAuthArgument("https://workspace.cloud.databricks.com")
			if err != nil {
				t.Fatalf("NewBasicWorkspaceOAuthArgument(): want no error, got %v", err)
			}

			tokenResponse := `access_token=token&refresh_token=refresh`
			if tt.disableOffline {
				tokenResponse = `access_token=token`
			}

			opts := []PersistentAuthOption{
				WithTokenCache(cache),
				WithBrowser(browser),
				WithHttpClient(&http.Client{
					Transport: fixtures.SliceTransport{
						{
							Method:   "POST",
							Resource: "/oidc/v1/token",
							Response: tokenResponse,
							ResponseHeaders: map[string][]string{
								"Content-Type": {"application/x-www-form-urlencoded"},
							},
						},
					},
				}),
				WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
				WithOAuthArgument(arg),
				WithDisableOfflineAccess(tt.disableOffline),
				WithScopes(tt.scopes),
			}

			p, err := NewPersistentAuth(ctx, opts...)
			if err != nil {
				t.Fatalf("NewPersistentAuth(): want no error, got %v", err)
			}
			defer p.Close()

			errc := make(chan error)
			go func() {
				err := p.Challenge()
				errc <- err
				close(errc)
			}()

			scope := <-scopeReceived
			state := <-browserOpened

			if scope != tt.want {
				t.Errorf("scope: want %q, got %q", tt.want, scope)
			}

			resp, err := http.Get(fmt.Sprintf("http://localhost:8020?code=__CODE__&state=%s", state))
			if err != nil {
				t.Fatalf("http.Get(): want no error, got %v", err)
			}
			defer resp.Body.Close()

			err = <-errc
			if err != nil {
				t.Fatalf("p.Challenge(): want no error, got %v", err)
			}
		})
	}
}
