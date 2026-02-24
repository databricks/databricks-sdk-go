package config

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

func TestParseExpiry(t *testing.T) {
	testCases := []struct {
		name    string
		expiry  string
		wantErr bool
	}{
		{"RFC3339Nano", "2024-03-20T10:30:00.123456789Z", false},
		{"RFC3339", "2024-03-20T10:30:00Z", false},
		{"RFC3339 with offset", "2024-03-20T10:30:00+01:00", false},
		{"space-separated", "2024-03-20 10:30:00", false},
		{"space-separated with 1 digit fractional", "2024-03-20 10:30:00.1", false},
		{"space-separated with 3 digits fractional", "2024-03-20 10:30:00.123", false},
		{"space-separated with 6 digits fractional", "2024-03-20 10:30:00.123456", false},
		{"space-separated with 9 digits fractional", "2024-03-20 10:30:00.123456789", false},
		{"unsupported format", "March 20, 2024", true},
		{"empty string", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parseExpiry(tc.expiry)
			if (err != nil) != tc.wantErr {
				t.Errorf("parseExpiry(%q) error = %v, wantErr %v", tc.expiry, err, tc.wantErr)
			}
		})
	}
}

func TestValidateCliPath(t *testing.T) {
	tempDir := t.TempDir()

	smallFile := filepath.Join(tempDir, "small_cli")
	if err := os.WriteFile(smallFile, make([]byte, 100), 0755); err != nil {
		t.Fatalf("failed to create small file: %v", err)
	}

	largeFile := filepath.Join(tempDir, "large_cli")
	if err := os.WriteFile(largeFile, make([]byte, databricksCliMinSize+1), 0755); err != nil {
		t.Fatalf("failed to create large file: %v", err)
	}

	dirPath := filepath.Join(tempDir, "cli_dir")
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		t.Fatalf("failed to create directory: %v", err)
	}

	testCases := []struct {
		name    string
		path    string
		wantErr error
	}{
		{"valid CLI", largeFile, nil},
		{"legacy CLI (too small)", smallFile, ErrLegacyCliDetected},
		{"non-existent file", filepath.Join(tempDir, "nonexistent"), ErrCliNotFound},
		{"directory", dirPath, ErrCliNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := validateCliPath(tc.path)
			if tc.wantErr == nil && err != nil {
				t.Errorf("validateCliPath() unexpected error: %v", err)
			} else if tc.wantErr != nil && !errors.Is(err, tc.wantErr) {
				t.Errorf("validateCliPath() error = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestFindDatabricksCli(t *testing.T) {
	tempDir := t.TempDir()

	cliName := "databricks"
	if runtime.GOOS == "windows" {
		cliName = "databricks.exe"
	}
	validCliPath := filepath.Join(tempDir, cliName)
	if err := os.WriteFile(validCliPath, make([]byte, databricksCliMinSize+1), 0755); err != nil {
		t.Fatalf("failed to create mock CLI: %v", err)
	}

	testCases := []struct {
		name    string
		cliPath string
		pathEnv string
		wantErr error
	}{
		{"explicit full path", validCliPath, "", nil},
		{"explicit non-existent path", filepath.Join(tempDir, "nonexistent"), "", ErrCliNotFound},
		{"found in PATH", "", tempDir, nil},
		{"not found in PATH", "", filepath.Join(tempDir, "empty"), ErrCliNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.pathEnv != "" {
				oldPath := os.Getenv("PATH")
				defer os.Setenv("PATH", oldPath)
				os.Setenv("PATH", tc.pathEnv)
			}

			_, err := findDatabricksCli(tc.cliPath)
			if tc.wantErr == nil && err != nil {
				t.Errorf("findDatabricksCli() unexpected error: %v", err)
			} else if tc.wantErr != nil && !errors.Is(err, tc.wantErr) {
				t.Errorf("findDatabricksCli() error = %v, want %v", err, tc.wantErr)
			}
		})
	}
}

func TestBuildCliCommands(t *testing.T) {
	const (
		cliPath     = "/path/to/databricks"
		host        = "https://workspace.cloud.databricks.com"
		accountHost = "https://accounts.cloud.databricks.com"
		unifiedHost = "https://unified.cloud.databricks.com"
		accountID   = "abc-123"
		workspaceID = "456"
	)

	testCases := []struct {
		name        string
		cfg         *Config
		wantCmd     []string
		wantHostCmd []string
	}{
		{
			name:    "workspace host",
			cfg:     &Config{Host: host},
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "account host",
			cfg:     &Config{Host: accountHost, AccountID: accountID},
			wantCmd: []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID},
		},
		{
			name: "unified host with account ID and workspace ID",
			cfg: &Config{
				Host:                       unifiedHost,
				Experimental_IsUnifiedHost: true,
				AccountID:                  accountID,
				WorkspaceID:                workspaceID,
			},
			wantCmd: []string{cliPath, "auth", "token", "--host", unifiedHost, "--experimental-is-unified-host", "--account-id", accountID, "--workspace-id", workspaceID},
		},
		{
			name: "unified host with account ID only",
			cfg: &Config{
				Host:                       unifiedHost,
				Experimental_IsUnifiedHost: true,
				AccountID:                  accountID,
			},
			wantCmd: []string{cliPath, "auth", "token", "--host", unifiedHost, "--experimental-is-unified-host", "--account-id", accountID},
		},
		{
			name: "unified host with workspace ID only",
			cfg: &Config{
				Host:                       unifiedHost,
				Experimental_IsUnifiedHost: true,
				WorkspaceID:                workspaceID,
			},
			wantCmd: []string{cliPath, "auth", "token", "--host", unifiedHost, "--experimental-is-unified-host", "--workspace-id", workspaceID},
		},
		{
			name: "unified host with no account ID or workspace ID",
			cfg: &Config{
				Host:                       unifiedHost,
				Experimental_IsUnifiedHost: true,
			},
			wantCmd: []string{cliPath, "auth", "token", "--host", unifiedHost, "--experimental-is-unified-host"},
		},
		{
			// Explicit false should fall back to the standard host type detection
			name: "unified host false with account host",
			cfg: &Config{
				Host:                       accountHost,
				Experimental_IsUnifiedHost: false,
				AccountID:                  accountID,
			},
			wantCmd: []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID},
		},
		{
			name:        "profile uses --profile with --host fallback",
			cfg:         &Config{Profile: "my-profile", Host: host},
			wantCmd:     []string{cliPath, "auth", "token", "--profile", "my-profile"},
			wantHostCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "profile without host — no fallback",
			cfg:     &Config{Profile: "my-profile"},
			wantCmd: []string{cliPath, "auth", "token", "--profile", "my-profile"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotCmd, gotHostCmd := buildCliCommands(cliPath, tc.cfg)
			if !slices.Equal(gotCmd, tc.wantCmd) {
				t.Errorf("primary cmd = %v, want %v", gotCmd, tc.wantCmd)
			}
			if !slices.Equal(gotHostCmd, tc.wantHostCmd) {
				t.Errorf("host cmd = %v, want %v", gotHostCmd, tc.wantHostCmd)
			}
		})
	}
}

func TestNewCliTokenSource(t *testing.T) {
	tempDir := t.TempDir()

	cliName := "databricks"
	if runtime.GOOS == "windows" {
		cliName = "databricks.exe"
	}
	validCliPath := filepath.Join(tempDir, cliName)
	if err := os.WriteFile(validCliPath, make([]byte, databricksCliMinSize+1), 0755); err != nil {
		t.Fatalf("failed to create mock CLI: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		cfg := &Config{DatabricksCliPath: validCliPath, Host: "https://example.databricks.com"}
		ts, err := NewCliTokenSource(cfg)
		if err != nil {
			t.Fatalf("NewCliTokenSource() unexpected error: %v", err)
		}
		// Verify CLI path was resolved and used
		if ts.cmd[0] != validCliPath {
			t.Errorf("cmd[0] = %q, want %q", ts.cmd[0], validCliPath)
		}
	})

	t.Run("CLI not found", func(t *testing.T) {
		cfg := &Config{DatabricksCliPath: filepath.Join(tempDir, "nonexistent"), Host: "https://example.databricks.com"}
		_, err := NewCliTokenSource(cfg)
		if !errors.Is(err, ErrCliNotFound) {
			t.Errorf("NewCliTokenSource() error = %v, want %v", err, ErrCliNotFound)
		}
	})
}

func TestCliTokenSource_Token(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping shell script test on Windows")
	}

	expiry := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	validResponse, _ := json.Marshal(cliTokenResponse{
		AccessToken: "test-token",
		TokenType:   "Bearer",
		Expiry:      expiry,
	})

	testCases := []struct {
		name       string
		script     string
		wantToken  string
		wantErrMsg string
	}{
		{
			name:      "success",
			script:    "#!/bin/sh\necho '" + string(validResponse) + "'",
			wantToken: "test-token",
		},
		{
			name:       "CLI error",
			script:     "#!/bin/sh\necho 'auth failed' >&2\nexit 1",
			wantErrMsg: "cannot get access token",
		},
		{
			name:       "invalid JSON",
			script:     "#!/bin/sh\necho 'not json'",
			wantErrMsg: "cannot parse CLI response",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			mockScript := filepath.Join(tempDir, "mock_cli.sh")
			if err := os.WriteFile(mockScript, []byte(tc.script), 0755); err != nil {
				t.Fatalf("failed to create mock script: %v", err)
			}

			ts := &CliTokenSource{cmd: []string{mockScript}}
			token, err := ts.Token(context.Background())

			if tc.wantErrMsg != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErrMsg) {
					t.Errorf("Token() error = %v, want error containing %q", err, tc.wantErrMsg)
				}
				return
			}
			if err != nil {
				t.Fatalf("Token() error = %v", err)
			}
			if token.AccessToken != tc.wantToken {
				t.Errorf("AccessToken = %q, want %q", token.AccessToken, tc.wantToken)
			}
		})
	}
}

func TestCliTokenSource_Token_FallbackOnUnknownFlag(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping shell script test on Windows")
	}

	expiry := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	validResponse, _ := json.Marshal(cliTokenResponse{
		AccessToken: "fallback-token",
		TokenType:   "Bearer",
		Expiry:      expiry,
	})

	tempDir := t.TempDir()

	// Primary script simulates an old CLI that doesn't know --profile.
	profileScript := filepath.Join(tempDir, "profile_cli.sh")
	if err := os.WriteFile(profileScript, []byte("#!/bin/sh\necho 'Error: unknown flag: --profile' >&2\nexit 1"), 0755); err != nil {
		t.Fatalf("failed to create profile script: %v", err)
	}

	// Fallback script succeeds with --host.
	hostScript := filepath.Join(tempDir, "host_cli.sh")
	if err := os.WriteFile(hostScript, []byte("#!/bin/sh\necho '"+string(validResponse)+"'"), 0755); err != nil {
		t.Fatalf("failed to create host script: %v", err)
	}

	ts := &CliTokenSource{
		cmd:     []string{profileScript},
		hostCmd: []string{hostScript},
	}
	token, err := ts.Token(context.Background())
	if err != nil {
		t.Fatalf("Token() error = %v, want fallback to succeed", err)
	}
	if token.AccessToken != "fallback-token" {
		t.Errorf("AccessToken = %q, want %q", token.AccessToken, "fallback-token")
	}
}

func TestCliTokenSource_Token_NoFallbackOnRealError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping shell script test on Windows")
	}

	tempDir := t.TempDir()

	// Primary script fails with a real auth error (not unknown flag).
	profileScript := filepath.Join(tempDir, "profile_cli.sh")
	if err := os.WriteFile(profileScript, []byte("#!/bin/sh\necho 'cache: databricks OAuth is not configured for this host' >&2\nexit 1"), 0755); err != nil {
		t.Fatalf("failed to create profile script: %v", err)
	}

	// Fallback script would succeed, but should not be called.
	hostScript := filepath.Join(tempDir, "host_cli.sh")
	if err := os.WriteFile(hostScript, []byte("#!/bin/sh\necho 'should not reach here' >&2\nexit 1"), 0755); err != nil {
		t.Fatalf("failed to create host script: %v", err)
	}

	ts := &CliTokenSource{
		cmd:     []string{profileScript},
		hostCmd: []string{hostScript},
	}
	_, err := ts.Token(context.Background())
	if err == nil {
		t.Fatal("Token() error = nil, want error")
	}
	if !strings.Contains(err.Error(), "databricks OAuth is not configured") {
		t.Errorf("Token() error = %v, want error containing original auth failure", err)
	}
}
