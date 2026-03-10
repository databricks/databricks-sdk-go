package u2m

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"golang.org/x/oauth2"
)

func TestCallbackServer_ExtractsIssuer(t *testing.T) {
	ctx := context.Background()

	browserOpened := make(chan string)
	browser := func(redirect string) error {
		browserOpened <- redirect
		return nil
	}
	tokenCache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(tokenCache),
		WithBrowser(browser),
		WithHttpClient(&http.Client{
			Transport: fixtures.SliceTransport{
				{
					Method:   "POST",
					Resource: "/oidc/accounts/xyz/v1/token",
					Response: `access_token=__ACCESS__&refresh_token=__REFRESH__`,
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
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	// Start a listener so we can create the callback server.
	err = p.startListener(ctx)
	if err != nil {
		t.Fatalf("startListener(): %v", err)
	}

	cb, err := p.newCallbackServer()
	if err != nil {
		t.Fatalf("newCallbackServer(): %v", err)
	}
	defer cb.Close()

	// Verify issuer is empty before any callback.
	if got := cb.Issuer(); got != "" {
		t.Fatalf("Issuer() before callback: want %q, got %q", "", got)
	}

	// Fire a callback with iss parameter.
	issuerURL := "https://adb-123.azuredatabricks.net/oidc"
	resp, err := http.Get(fmt.Sprintf("http://%s?code=xxx&state=yyy&iss=%s", p.redirectAddr, issuerURL))
	if err != nil {
		t.Fatalf("http.Get(): %v", err)
	}
	defer resp.Body.Close()

	// Drain the feedback channel so ServeHTTP completes.
	<-cb.feedbackCh

	if got := cb.Issuer(); got != issuerURL {
		t.Fatalf("Issuer(): want %q, got %q", issuerURL, got)
	}
}

func TestCallbackServer_NoIssuer(t *testing.T) {
	ctx := context.Background()

	tokenCache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(tokenCache),
		WithBrowser(func(string) error { return nil }),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	err = p.startListener(ctx)
	if err != nil {
		t.Fatalf("startListener(): %v", err)
	}

	cb, err := p.newCallbackServer()
	if err != nil {
		t.Fatalf("newCallbackServer(): %v", err)
	}
	defer cb.Close()

	// Fire a callback without iss parameter.
	resp, err := http.Get(fmt.Sprintf("http://%s?code=xxx&state=yyy", p.redirectAddr))
	if err != nil {
		t.Fatalf("http.Get(): %v", err)
	}
	defer resp.Body.Close()

	// Drain the feedback channel so ServeHTTP completes.
	<-cb.feedbackCh

	if got := cb.Issuer(); got != "" {
		t.Fatalf("Issuer(): want %q, got %q", "", got)
	}
}

func TestCallbackServer_IssuerWithAccountPath(t *testing.T) {
	ctx := context.Background()

	tokenCache := &tokenCacheMock{
		store: func(key string, tok *oauth2.Token) error {
			return nil
		},
	}
	arg, err := NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "xyz")
	if err != nil {
		t.Fatalf("NewBasicAccountOAuthArgument(): %v", err)
	}

	p, err := NewPersistentAuth(
		ctx,
		WithTokenCache(tokenCache),
		WithBrowser(func(string) error { return nil }),
		WithOAuthEndpointSupplier(MockOAuthEndpointSupplier{}),
		WithOAuthArgument(arg),
	)
	if err != nil {
		t.Fatalf("NewPersistentAuth(): %v", err)
	}
	defer p.Close()

	err = p.startListener(ctx)
	if err != nil {
		t.Fatalf("startListener(): %v", err)
	}

	cb, err := p.newCallbackServer()
	if err != nil {
		t.Fatalf("newCallbackServer(): %v", err)
	}
	defer cb.Close()

	// Fire a callback with iss containing an account path.
	issuerURL := "https://nike.databricks.com/oidc/accounts/abc-123-def"
	resp, err := http.Get(fmt.Sprintf("http://%s?code=xxx&state=yyy&iss=%s", p.redirectAddr, issuerURL))
	if err != nil {
		t.Fatalf("http.Get(): %v", err)
	}
	defer resp.Body.Close()

	// Drain the feedback channel so ServeHTTP completes.
	<-cb.feedbackCh

	if got := cb.Issuer(); got != issuerURL {
		t.Fatalf("Issuer(): want %q, got %q", issuerURL, got)
	}
}
