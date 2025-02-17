package config

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/credentials/u2m"
	"github.com/databricks/databricks-sdk-go/credentials/u2m/cache"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

type mockU2mTokenSource struct {
	token *oauth2.Token
	err   error
}

func (m mockU2mTokenSource) Token() (*oauth2.Token, error) {
	return m.token, m.err
}

func TestU2MCredentials(t *testing.T) {
	tests := []struct {
		name       string
		cfg        *Config
		auth       oauth2.TokenSource
		expectErr  string
		expectAuth string
	}{
		{
			name: "happy path",
			cfg: &Config{
				Host: "https://myworkspace.cloud.databricks.com",
			},
			auth: mockU2mTokenSource{
				token: &oauth2.Token{
					AccessToken: "dummy_access_token",
					Expiry:      time.Now().Add(1 * time.Hour),
				},
			},
			expectAuth: "Bearer dummy_access_token",
		},
		{
			name: "expired token with invalid refresh token",
			cfg: &Config{
				Host: "https://myworkspace.cloud.databricks.com",
			},
			auth: mockU2mTokenSource{
				err: &u2m.InvalidRefreshTokenError{},
			},
			expectErr: `a new access token could not be retrieved because the refresh token is invalid. If using the CLI, run the following command to reauthenticate:

  $ databricks auth login --host https://myworkspace.cloud.databricks.com`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			strat := u2mCredentials{
				ts: tt.auth,
			}
			provider, err := strat.Configure(ctx, tt.cfg)
			if tt.expectErr != "" {
				require.ErrorContains(t, err, tt.expectErr)
				return
			}
			require.NoError(t, err)

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://databricks.com", nil)
			require.NoError(t, err)

			err = provider.SetHeaders(req)
			require.NoError(t, err)
			require.Equal(t, tt.expectAuth, req.Header.Get("Authorization"))
		})
	}
}

func TestDatabricksCli_ErrorHandler(t *testing.T) {
	invalidRefreshTokenError := fmt.Errorf("refresh: %w", &u2m.InvalidRefreshTokenError{})
	workspaceArg := func() (u2m.OAuthArgument, error) {
		return u2m.NewBasicWorkspaceOAuthArgument("https://myworkspace.cloud.databricks.com")
	}
	accountArg := func() (u2m.OAuthArgument, error) {
		return u2m.NewBasicAccountOAuthArgument("https://accounts.cloud.databricks.com", "abc")
	}
	testCases := []struct {
		name string
		cfg  *Config
		arg  func() (u2m.OAuthArgument, error)
		err  error
		want error
	}{
		{
			name: "not configured is ignored",
			arg:  workspaceArg,
			err:  cache.ErrNotConfigured,
			want: nil,
		},
		{
			name: "other error is ignored",
			arg:  workspaceArg,
			err:  errors.New("foobar"),
			want: nil,
		},
		{
			name: "invalid refresh token is adapted: profile provided",
			arg:  workspaceArg,
			cfg:  &Config{Profile: "my-profile"},
			err:  invalidRefreshTokenError,
			want: &CliInvalidRefreshTokenError{loginCommand: "databricks auth login --profile my-profile", err: invalidRefreshTokenError},
		},
		{
			name: "invalid refresh token is adapted: profile not provided for workspace",
			cfg:  &Config{},
			arg:  workspaceArg,
			err:  invalidRefreshTokenError,
			want: &CliInvalidRefreshTokenError{loginCommand: "databricks auth login --host https://myworkspace.cloud.databricks.com", err: invalidRefreshTokenError},
		},
		{
			name: "invalid refresh token is adapted: profile not provided for account",
			cfg:  &Config{},
			arg:  accountArg,
			err:  invalidRefreshTokenError,
			want: &CliInvalidRefreshTokenError{loginCommand: "databricks auth login --host https://accounts.cloud.databricks.com --account-id abc", err: invalidRefreshTokenError},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			arg, err := tc.arg()
			require.NoError(t, err)
			got := DatabricksCliCredentials.errorHandler(context.Background(), tc.cfg, arg, tc.err)
			require.Equal(t, tc.want, got)
		})
	}
}
