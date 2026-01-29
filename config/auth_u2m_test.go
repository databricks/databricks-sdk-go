package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/oauth2"
)

type testTokenSource struct {
	token  *oauth2.Token
	err    error
	counts int // number of times Token() is called
}

func (m *testTokenSource) Token() (*oauth2.Token, error) {
	m.counts++
	return m.token, m.err
}

var testValidToken = &oauth2.Token{
	AccessToken: "valid-access-token",
	TokenType:   "Bearer",
	Expiry:      time.Now().Add(1 * time.Hour),
}

var testExpiredToken = &oauth2.Token{
	AccessToken:  "expired-access-token",
	RefreshToken: "refresh-token",
	TokenType:    "Bearer",
	Expiry:       time.Now().Add(-1 * time.Hour),
}

func getHeader(cp credentials.CredentialsProvider) (http.Header, error) {
	req := &http.Request{Header: make(http.Header)} // dummy request
	if err := cp.SetHeaders(req); err != nil {
		return nil, err
	}
	return req.Header, nil
}

// Sentinel errors for testing.
var (
	errNetwork             = errors.New("network timeout")
	errAuthentication      = errors.New("authentication failed")
	errInvalidRefreshToken = &u2m.InvalidRefreshTokenError{}
)

// mockPersistentAuthFactory creates a test factory for bypassing real auth setup.
// Use this when tests only need to control token behavior without caring about auth configuration.
func mockPersistentAuthFactory(ts oauth2.TokenSource) persistentAuthFactory {
	return func(ctx context.Context, opts ...u2m.PersistentAuthOption) (oauth2.TokenSource, error) {
		return ts, nil
	}
}

func TestU2MCredentials_Configure(t *testing.T) {
	testCases := []struct {
		desc            string
		cfg             *Config
		testTokenSource *testTokenSource
		wantConfigErr   string // error message from Configure()
		wantHeaderErr   string // error message from SetHeaders()
		wantAuthHeader  string // expected Authorization header
	}{
		{
			desc: "missing host returns error",
			cfg: &Config{
				Host: "",
			},
			wantConfigErr: "host is required",
		},
		{
			desc: "successful configuration with workspace host",
			cfg: &Config{
				Host: "https://workspace.cloud.databricks.com",
			},
			testTokenSource: &testTokenSource{
				token: testValidToken,
			},
			wantAuthHeader: "Bearer valid-access-token",
		},
		{
			desc: "successful configuration with account host and account ID",
			cfg: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "abc-123",
			},
			testTokenSource: &testTokenSource{
				token: testValidToken,
			},
			wantAuthHeader: "Bearer valid-access-token",
		},
		{
			desc: "expired token is accepted",
			cfg: &Config{
				Host: "https://workspace.cloud.databricks.com",
			},
			testTokenSource: &testTokenSource{
				token: testExpiredToken,
			},
			wantAuthHeader: "Bearer expired-access-token",
		},
		{
			desc: "token source returns network error",
			cfg: &Config{
				Host: "https://workspace.cloud.databricks.com",
			},
			testTokenSource: &testTokenSource{
				err: errNetwork,
			},
			wantHeaderErr: "network timeout",
		},
		{
			desc: "token source returns authentication error",
			cfg: &Config{
				Host: "https://workspace.cloud.databricks.com",
			},
			testTokenSource: &testTokenSource{
				err: errAuthentication,
			},
			wantHeaderErr: "authentication failed",
		},
		{
			desc: "invalid refresh token error - workspace with profile",
			cfg: &Config{
				Host:     "https://workspace.cloud.databricks.com",
				Profile:  "my-workspace",
				resolved: true,
			},
			testTokenSource: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantHeaderErr: "databricks auth login --profile my-workspace",
		},
		{
			desc: "invalid refresh token error - workspace without profile",
			cfg: &Config{
				Host:     "https://workspace.cloud.databricks.com",
				resolved: true,
			},
			testTokenSource: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantHeaderErr: "databricks auth login --host https://workspace.cloud.databricks.com",
		},
		{
			desc: "invalid refresh token error - account with profile",
			cfg: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "xyz-789",
				Profile:   "prod-account",
				resolved:  true,
			},
			testTokenSource: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantHeaderErr: "databricks auth login --profile prod-account",
		},
		{
			desc: "invalid refresh token error - account without profile",
			cfg: &Config{
				Host:      "https://accounts.cloud.databricks.com",
				AccountID: "abc-123",
				resolved:  true,
			},
			testTokenSource: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantHeaderErr: "databricks auth login --host https://accounts.cloud.databricks.com --account-id abc-123",
		},
		{
			desc: "wrapped invalid refresh token error",
			cfg: &Config{
				Host:     "https://workspace.cloud.databricks.com",
				Profile:  "test",
				resolved: true,
			},
			testTokenSource: &testTokenSource{
				err: fmt.Errorf("oauth2: %w", errInvalidRefreshToken),
			},
			wantHeaderErr: "databricks auth login --profile test",
		},
		{
			desc: "invalid refresh token error - azure account without profile",
			cfg: &Config{
				Host:      "https://accounts.azure.databricks.net",
				AccountID: "abc-456",
				resolved:  true,
			},
			testTokenSource: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantHeaderErr: "databricks auth login --host https://accounts.azure.databricks.net --account-id abc-456",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			ctx := context.Background()
			u := u2mCredentials{
				newPersistentAuth: mockPersistentAuthFactory(tc.testTokenSource),
			}

			cp, gotConfigErr := u.Configure(ctx, tc.cfg)

			if tc.wantConfigErr != "" {
				if gotConfigErr == nil {
					t.Fatalf("Configure() want error containing %q, got nil", tc.wantConfigErr)
				}
				if !strings.Contains(gotConfigErr.Error(), tc.wantConfigErr) {
					t.Errorf("Configure() error = %q, want error containing %q", gotConfigErr.Error(), tc.wantConfigErr)
				}
				return
			}
			if gotConfigErr != nil {
				t.Fatalf("Configure() unexpected error: %v", gotConfigErr)
			}

			header, gotHeaderErr := getHeader(cp)

			if tc.wantHeaderErr != "" {
				if gotHeaderErr == nil {
					t.Fatalf("SetHeaders() want error containing %q, got nil", tc.wantHeaderErr)
				}
				if !strings.Contains(gotHeaderErr.Error(), tc.wantHeaderErr) {
					t.Errorf("SetHeaders() error = %q, want error containing %q", gotHeaderErr.Error(), tc.wantHeaderErr)
				}
				return
			}
			if gotHeaderErr != nil {
				t.Fatalf("SetHeaders() unexpected error: %v", gotHeaderErr)
			}
			if gotHeader := header.Get("Authorization"); gotHeader != tc.wantAuthHeader {
				t.Errorf("Authorization header = %q, want %q", gotHeader, tc.wantAuthHeader)
			}
		})
	}
}

func TestU2MCredentials_Configure_TokenCaching(t *testing.T) {
	ts := &testTokenSource{token: testValidToken}

	u := u2mCredentials{
		newPersistentAuth: mockPersistentAuthFactory(ts),
	}
	cfg := &Config{
		Host: "https://workspace.cloud.databricks.com",
	}

	cp, err := u.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Configure() error = %v", err)
	}

	// Call the credentials provider twice, only the first call should
	// actually call the underlying token source.
	if _, err := getHeader(cp); err != nil {
		t.Fatalf("First getHeader() error = %v", err)
	}
	if _, err := getHeader(cp); err != nil {
		t.Fatalf("Second getHeader() error = %v", err)
	}

	if ts.counts != 1 {
		t.Errorf("token source call count = %d, want 1 (should use cache)", ts.counts)
	}
}

func TestU2MCredentials_Configure_Scopes(t *testing.T) {
	const workspaceHost = "https://workspace.cloud.databricks.com"
	testCases := []struct {
		name   string
		scopes []string
		want   []string
	}{
		{
			name:   "nil scopes uses default",
			scopes: nil,
			want:   []string{"all-apis"},
		},
		{
			name:   "empty scopes uses default",
			scopes: []string{},
			want:   []string{"all-apis"},
		},
		{
			name:   "multiple scopes are sorted",
			scopes: []string{"clusters", "jobs", "sql:read"},
			want:   []string{"clusters", "jobs", "sql:read"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := &testTokenSource{token: testValidToken}
			var capturedPA *u2m.PersistentAuth

			u := u2mCredentials{
				newPersistentAuth: func(ctx context.Context, opts ...u2m.PersistentAuthOption) (oauth2.TokenSource, error) {
					pa, err := u2m.NewPersistentAuth(ctx, opts...)
					if err != nil {
						return nil, err
					}
					capturedPA = pa
					return ts, nil
				},
			}
			cfg := &Config{
				Host:   workspaceHost,
				Scopes: tc.scopes,
			}

			// Mark scopes as coming from config file to bypass validation as custom scopes are only
			// allowed from config files with databricks-cli auth.
			// TODO: remove this block once the token store can identify scopes based on their identity
			// and we no longer have to block users from specifying scopes explicitly.
			if len(tc.scopes) > 0 {
				for i := range ConfigAttributes {
					if ConfigAttributes[i].Name == "scopes" {
						cfg.SetAttrSource(&ConfigAttributes[i], Source{Type: SourceFile, Name: "test"})
						break
					}
				}
			}

			_, err := u.Configure(context.Background(), cfg)
			if err != nil {
				t.Fatalf("Configure() error = %v", err)
			}

			arg, err := u2m.NewBasicWorkspaceOAuthArgument(workspaceHost)
			if err != nil {
				t.Fatalf("NewBasicWorkspaceOAuthArgument() error = %v", err)
			}
			wantPA, err := u2m.NewPersistentAuth(context.Background(),
				u2m.WithOAuthArgument(arg),
				u2m.WithScopes(tc.want),
			)
			if err != nil {
				t.Fatalf("NewPersistentAuth() error = %v", err)
			}

			if diff := cmp.Diff(wantPA, capturedPA,
				cmp.AllowUnexported(u2m.PersistentAuth{}, u2m.BasicWorkspaceOAuthArgument{}),
				cmpopts.IgnoreFields(u2m.PersistentAuth{},
					"cache", "client", "endpointSupplier", "browser",
					"ln", "ctx", "redirectAddr", "port", "netListen", "disableOfflineAccess"),
			); diff != "" {
				t.Errorf("PersistentAuth mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestU2MCredentials_Configure_CustomScopesError(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         *Config
		scopeSource *Source
		wantErr     string
	}{
		{
			name: "error when scopes are explicitly set",
			cfg: &Config{
				Host:   "https://workspace.cloud.databricks.com",
				Scopes: []string{"clusters", "jobs"},
			},
			wantErr: ErrCustomScopesNotSupported,
		},
		{
			name: "no error when scopes are empty",
			cfg: &Config{
				Host:   "https://workspace.cloud.databricks.com",
				Scopes: []string{},
			},
		},
		{
			name: "no error when scopes are nil",
			cfg: &Config{
				Host: "https://workspace.cloud.databricks.com",
			},
		},
		{
			name: "no error when scopes come from config file",
			cfg: &Config{
				Host:   "https://workspace.cloud.databricks.com",
				Scopes: []string{"sql"},
			},
			scopeSource: &Source{Type: SourceFile, Name: "~/.databrickscfg"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.scopeSource != nil {
				for i := range ConfigAttributes {
					if ConfigAttributes[i].Name == "scopes" {
						tc.cfg.SetAttrSource(&ConfigAttributes[i], *tc.scopeSource)
						break
					}
				}
			}
			ts := &testTokenSource{token: testValidToken}
			u := u2mCredentials{
				newPersistentAuth: mockPersistentAuthFactory(ts),
			}

			_, err := u.Configure(context.Background(), tc.cfg)
			var gotErr string
			if err != nil {
				gotErr = err.Error()
			}
			if gotErr != tc.wantErr {
				t.Errorf("Configure() error = %q, want %q", gotErr, tc.wantErr)
			}
		})
	}
}
