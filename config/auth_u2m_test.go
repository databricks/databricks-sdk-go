package config

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"golang.org/x/oauth2"
)

type testTokenSource struct {
	token *oauth2.Token
	err   error
}

func (m *testTokenSource) Token() (*oauth2.Token, error) {
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

// Sentinel errors for testing.
var (
	errNetwork             = errors.New("network timeout")
	errAuthentication      = errors.New("authentication failed")
	errInvalidRefreshToken = &u2m.InvalidRefreshTokenError{}
)

func TestWrapWithCLIErrorHandling(t *testing.T) {
	testCases := []struct {
		desc             string
		cfg              *Config
		arg              u2m.OAuthArgument
		ts               oauth2.TokenSource
		wantToken        *oauth2.Token
		wantErr          bool
		wantLoginCommand string // if set, check CLI error message
	}{
		{
			desc: "successful token retrieval with workspace arg",
			cfg:  &Config{Profile: "default"},
			arg:  mustCreateWorkspaceArg(t, "https://example.cloud.databricks.com"),
			ts: &testTokenSource{
				token: testValidToken,
			},
			wantToken: testValidToken,
		},
		{
			desc: "successful token retrieval with account arg",
			cfg:  &Config{Profile: "prod"},
			arg:  mustCreateAccountArg(t, "https://accounts.cloud.databricks.com", "abc-123"),
			ts: &testTokenSource{
				token: testValidToken,
			},
			wantToken: testValidToken,
		},
		{
			desc: "successful token retrieval with expired token",
			cfg:  &Config{},
			arg:  mustCreateWorkspaceArg(t, "https://workspace.databricks.com"),
			ts: &testTokenSource{
				token: testExpiredToken,
			},
			wantToken: testExpiredToken,
		},
		{
			desc: "invalid refresh token error - workspace with profile",
			cfg:  &Config{Profile: "my-workspace"},
			arg:  mustCreateWorkspaceArg(t, "https://workspace.databricks.com"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --profile my-workspace",
		},
		{
			desc: "invalid refresh token error - workspace without profile",
			cfg:  &Config{},
			arg:  mustCreateWorkspaceArg(t, "https://workspace.databricks.com"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --host https://workspace.databricks.com",
		},
		{
			desc: "invalid refresh token error - account with profile",
			cfg:  &Config{Profile: "prod-account"},
			arg:  mustCreateAccountArg(t, "https://accounts.cloud.databricks.com", "xyz-789"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --profile prod-account",
		},
		{
			desc: "invalid refresh token error - account without profile",
			cfg:  &Config{},
			arg:  mustCreateAccountArg(t, "https://accounts.azure.databricks.net", "abc-456"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --host https://accounts.azure.databricks.net --account-id abc-456",
		},
		{
			desc: "wrapped invalid refresh token error",
			cfg:  &Config{Profile: "test"},
			arg:  mustCreateWorkspaceArg(t, "https://test.databricks.com"),
			ts: &testTokenSource{
				err: fmt.Errorf("oauth2: %w", errInvalidRefreshToken),
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --profile test",
		},
		{
			desc: "other error is passed through unchanged - network error",
			cfg:  &Config{Profile: "default"},
			arg:  mustCreateWorkspaceArg(t, "https://workspace.databricks.com"),
			ts: &testTokenSource{
				err: errNetwork,
			},
			wantErr: true,
		},
		{
			desc: "other error is passed through unchanged - authentication error",
			cfg:  &Config{},
			arg:  mustCreateWorkspaceArg(t, "https://workspace.databricks.com"),
			ts: &testTokenSource{
				err: errAuthentication,
			},
			wantErr: true,
		},
		{
			desc: "unified auth argument with profile",
			cfg:  &Config{Profile: "unified-profile"},
			arg:  mustCreateUnifiedArg(t, "https://unified.databricks.com", "unified-123"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --profile unified-profile",
		},
		{
			desc: "empty profile with workspace",
			cfg:  &Config{Profile: ""},
			arg:  mustCreateWorkspaceArg(t, "https://my-workspace.cloud.databricks.com"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --host https://my-workspace.cloud.databricks.com",
		},
		{
			desc: "nil config with workspace arg",
			cfg:  nil,
			arg:  mustCreateWorkspaceArg(t, "https://workspace.databricks.com"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --host https://workspace.databricks.com",
		},
		{
			desc: "nil config with account arg",
			cfg:  nil,
			arg:  mustCreateAccountArg(t, "https://accounts.cloud.databricks.com", "test-123"),
			ts: &testTokenSource{
				err: errInvalidRefreshToken,
			},
			wantErr:          true,
			wantLoginCommand: "databricks auth login --host https://accounts.cloud.databricks.com --account-id test-123",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cfg := tc.cfg
			if cfg == nil {
				cfg = &Config{}
			}

			wrapped := wrapWithCLIErrorHandling(cfg, tc.arg, tc.ts)
			if wrapped == nil {
				t.Fatal("wrapWithCLIErrorHandling returned nil")
			}

			gotToken, gotErr := wrapped.Token(context.Background())

			if tc.wantErr && gotErr == nil {
				t.Fatal("want error but got none")
			}
			if !tc.wantErr && gotErr != nil {
				t.Fatalf("want no error but got: %v", gotErr)
			}
			if tc.wantLoginCommand != "" {
				var cliErr *CliInvalidRefreshTokenError
				if !errors.As(gotErr, &cliErr) {
					t.Fatalf("want CliInvalidRefreshTokenError but got: %T: %v", gotErr, gotErr)
				}
				if cliErr.loginCommand != tc.wantLoginCommand {
					t.Errorf("login command mismatch:\n  got:  %q\n  want: %q", cliErr.loginCommand, tc.wantLoginCommand)
				}
			}
			if tc.wantToken != gotToken {
				t.Errorf("want token %v, got %v", tc.wantToken, gotToken)
			}
		})
	}
}

func mustCreateWorkspaceArg(t *testing.T, host string) u2m.OAuthArgument {
	t.Helper()
	arg, err := u2m.NewBasicWorkspaceOAuthArgument(host)
	if err != nil {
		t.Fatalf("failed to create workspace arg: %v", err)
	}
	return arg
}

func mustCreateAccountArg(t *testing.T, host, accountID string) u2m.OAuthArgument {
	t.Helper()
	arg, err := u2m.NewBasicAccountOAuthArgument(host, accountID)
	if err != nil {
		t.Fatalf("failed to create account arg: %v", err)
	}
	return arg
}

func mustCreateUnifiedArg(t *testing.T, host, accountID string) u2m.OAuthArgument {
	t.Helper()
	arg, err := u2m.NewBasicUnifiedOAuthArgument(host, accountID)
	if err != nil {
		t.Fatalf("failed to create unified arg: %v", err)
	}
	return arg
}
