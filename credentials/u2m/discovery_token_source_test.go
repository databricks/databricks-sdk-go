package u2m

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestDeriveHostFromIssuer(t *testing.T) {
	tests := []struct {
		name    string
		issuer  string
		want    string
		wantErr bool
	}{
		{
			name:   "workspace oidc",
			issuer: "https://adb-xxx.azuredatabricks.net/oidc",
			want:   "https://adb-xxx.azuredatabricks.net",
		},
		{
			name:   "workspace cloud",
			issuer: "https://workspace.cloud.databricks.com/oidc",
			want:   "https://workspace.cloud.databricks.com",
		},
		{
			name:   "spog with account path",
			issuer: "https://nike.databricks.com/oidc/accounts/xxx",
			want:   "https://nike.databricks.com",
		},
		{
			name:    "empty issuer",
			issuer:  "",
			wantErr: true,
		},
		{
			name:    "http scheme",
			issuer:  "http://insecure.net/oidc",
			wantErr: true,
		},
		{
			name:    "no host",
			issuer:  "https:///oidc",
			wantErr: true,
		},
		{
			name:   "http localhost allowed",
			issuer: "http://127.0.0.1:12345/oidc",
			want:   "http://127.0.0.1:12345",
		},
		{
			name:    "http non-local host rejected",
			issuer:  "http://127.0.0.1.attacker.tld/oidc",
			wantErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := DeriveHostFromIssuer(tc.issuer)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("DeriveHostFromIssuer(%q): want error, got nil", tc.issuer)
				}
				return
			}
			if err != nil {
				t.Fatalf("DeriveHostFromIssuer(%q): unexpected error: %v", tc.issuer, err)
			}
			if got != tc.want {
				t.Errorf("DeriveHostFromIssuer(%q) = %q, want %q", tc.issuer, got, tc.want)
			}
		})
	}
}

func TestDeriveTokenEndpoint(t *testing.T) {
	tests := []struct {
		name   string
		issuer string
		want   string
	}{
		{
			name:   "standard",
			issuer: "https://adb-xxx.net/oidc",
			want:   "https://adb-xxx.net/oidc/v1/token",
		},
		{
			name:   "trailing slash",
			issuer: "https://adb-xxx.net/oidc/",
			want:   "https://adb-xxx.net/oidc/v1/token",
		},
		{
			name:   "account path",
			issuer: "https://nike.databricks.com/oidc/accounts/abc123",
			want:   "https://nike.databricks.com/oidc/accounts/abc123/v1/token",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DeriveTokenEndpoint(tc.issuer)
			if got != tc.want {
				t.Errorf("DeriveTokenEndpoint(%q) = %q, want %q", tc.issuer, got, tc.want)
			}
		})
	}
}

func TestBuildDiscoveryAuthorizeURL(t *testing.T) {
	pkce := PKCEParams{
		Challenge:       "test-challenge",
		ChallengeMethod: "S256",
		Verifier:        "test-verifier",
	}
	scopes := []string{"offline_access", "all-apis"}
	redirectAddr := "localhost:8020"
	state := "test-state"

	got := BuildDiscoveryAuthorizeURL(redirectAddr, state, pkce, scopes)

	// Parse the top-level URL.
	u, err := url.Parse(got)
	if err != nil {
		t.Fatalf("parsing URL: %v", err)
	}
	if u.Scheme != "https" || u.Host != "login.databricks.com" {
		t.Errorf("want host https://login.databricks.com, got %s://%s", u.Scheme, u.Host)
	}

	// Extract and decode destination_url.
	destURL := u.Query().Get("destination_url")
	if destURL == "" {
		t.Fatal("destination_url query param is empty")
	}

	// Parse the destination_url as a relative URL with query params.
	destParsed, err := url.Parse(destURL)
	if err != nil {
		t.Fatalf("parsing destination_url: %v", err)
	}
	if destParsed.Path != "/oidc/v1/authorize" {
		t.Errorf("destination_url path = %q, want %q", destParsed.Path, "/oidc/v1/authorize")
	}

	// Verify all expected OAuth params.
	q := destParsed.Query()
	expectations := map[string]string{
		"client_id":             appClientID,
		"redirect_uri":          "http://localhost:8020",
		"response_type":         "code",
		"scope":                 "offline_access all-apis",
		"state":                 "test-state",
		"code_challenge":        "test-challenge",
		"code_challenge_method": "S256",
	}
	for key, want := range expectations {
		if got := q.Get(key); got != want {
			t.Errorf("destination_url param %q = %q, want %q", key, got, want)
		}
	}
}

func TestBuildDiscoveryAuthorizeURL_HostOverride(t *testing.T) {
	pkce := PKCEParams{
		Challenge:       "c",
		ChallengeMethod: "S256",
		Verifier:        "v",
	}
	scopes := []string{"offline_access", "all-apis"}
	tests := []struct {
		name string
		host string
	}{
		{
			name: "custom host",
			host: "https://login.dev.databricks.com",
		},
		{
			name: "custom host with trailing slash",
			host: "https://login.dev.databricks.com/",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := buildDiscoveryAuthorizeURL(tc.host, "localhost:8020", "s", pkce, scopes)
			u, err := url.Parse(got)
			if err != nil {
				t.Fatalf("parsing URL: %v", err)
			}
			if u.Host != "login.dev.databricks.com" {
				t.Errorf("host = %q, want login.dev.databricks.com", u.Host)
			}
		})
	}
}

func TestDiscoveryTokenSource_Challenge(t *testing.T) {
	// Create a mock token server that responds to POST /oidc/v1/token.
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("token server: want POST, got %s", r.Method)
		}
		if r.URL.Path != "/oidc/v1/token" {
			t.Errorf("token server: want path /oidc/v1/token, got %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"access_token":"test-access-token","refresh_token":"test-refresh-token","token_type":"Bearer","expires_in":3600}`)
	}))
	defer tokenServer.Close()

	// The issuer will be the mock server URL + /oidc.
	issuer := tokenServer.URL + "/oidc"

	browserOpened := make(chan string, 1)
	browserMock := func(u string) error {
		browserOpened <- u
		return nil
	}

	storedTokens := map[string]*oauth2.Token{}
	cacheMock := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			storedTokens[key] = tok
			return nil
		},
	}

	arg, err := NewBasicDiscoveryOAuthArgument("test-profile")
	if err != nil {
		t.Fatalf("NewBasicDiscoveryOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		context.Background(),
		WithTokenCache(cacheMock),
		WithBrowser(browserMock),
		WithHttpClient(tokenServer.Client()),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
		WithDiscoveryLogin(),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}

	// Start the listener so redirectAddr is set.
	err = p.startListener(context.Background())
	if err != nil {
		t.Fatalf("startListener(): %v", err)
	}
	defer p.Close()

	dts := &discoveryTokenSource{pa: p}

	errc := make(chan error, 1)
	go func() {
		errc <- dts.challenge()
	}()

	// Wait for browser to be called and extract state from the URL.
	var state string
	select {
	case authURL := <-browserOpened:
		u, err := url.Parse(authURL)
		if err != nil {
			t.Fatalf("parsing auth URL: %v", err)
		}
		destURL := u.Query().Get("destination_url")
		dest, err := url.Parse(destURL)
		if err != nil {
			t.Fatalf("parsing destination_url: %v", err)
		}
		state = dest.Query().Get("state")
		if state == "" {
			t.Fatal("state is empty in authorize URL")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for browser to be called")
	}

	// Fire the callback with code, state, and iss.
	callbackURL := fmt.Sprintf("http://%s?code=test-code&state=%s&iss=%s",
		p.redirectAddr, url.QueryEscape(state), url.QueryEscape(issuer))
	resp, err := http.Get(callbackURL)
	if err != nil {
		t.Fatalf("callback GET: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		t.Fatalf("callback: want status 200, got %d", resp.StatusCode)
	}

	// Wait for challenge to complete.
	select {
	case err := <-errc:
		if err != nil {
			t.Fatalf("challenge(): %v", err)
		}
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for challenge to complete")
	}

	// Verify discovered host was set.
	expectedHost, err := DeriveHostFromIssuer(issuer)
	if err != nil {
		t.Fatalf("DeriveHostFromIssuer(%q): %v", issuer, err)
	}
	if arg.GetDiscoveredHost() != expectedHost {
		t.Errorf("discovered host = %q, want %q", arg.GetDiscoveredHost(), expectedHost)
	}
	if len(storedTokens) != 2 {
		t.Fatalf("store count: want 2 keys (profile and host), got %d", len(storedTokens))
	}
	for _, key := range []string{"test-profile", expectedHost} {
		storedToken := storedTokens[key]
		if storedToken == nil {
			t.Fatalf("stored token for key %q is nil", key)
		}
		if storedToken.AccessToken != "test-access-token" {
			t.Errorf("access token for key %q = %q, want %q", key, storedToken.AccessToken, "test-access-token")
		}
		if storedToken.RefreshToken != "test-refresh-token" {
			t.Errorf("refresh token for key %q = %q, want %q", key, storedToken.RefreshToken, "test-refresh-token")
		}
	}
}
