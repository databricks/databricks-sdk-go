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
		name           string
		cfg            *Config
		wantForceCmd   []string
		wantProfileCmd []string
		wantHostCmd    []string
	}{
		{
			name:         "workspace host only — force-refresh based on host",
			cfg:          &Config{Host: host},
			wantForceCmd: []string{cliPath, "auth", "token", "--host", host, "--force-refresh"},
			wantHostCmd:  []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:         "account host only — force-refresh based on host with account-id",
			cfg:          &Config{Host: accountHost, AccountID: accountID},
			wantForceCmd: []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID, "--force-refresh"},
			wantHostCmd:  []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID},
		},
		{
			name: "former unified host treated as workspace",
			cfg: &Config{
				Host:        unifiedHost,
				AccountID:   accountID,
				WorkspaceID: workspaceID,
			},
			wantForceCmd: []string{cliPath, "auth", "token", "--host", unifiedHost, "--force-refresh"},
			wantHostCmd:  []string{cliPath, "auth", "token", "--host", unifiedHost},
		},
		{
			name:           "profile with host — force-refresh based on profile",
			cfg:            &Config{Profile: "my-profile", Host: host},
			wantForceCmd:   []string{cliPath, "auth", "token", "--profile", "my-profile", "--force-refresh"},
			wantProfileCmd: []string{cliPath, "auth", "token", "--profile", "my-profile"},
			wantHostCmd:    []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:           "profile without host — no host fallback",
			cfg:            &Config{Profile: "my-profile"},
			wantForceCmd:   []string{cliPath, "auth", "token", "--profile", "my-profile", "--force-refresh"},
			wantProfileCmd: []string{cliPath, "auth", "token", "--profile", "my-profile"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotForceCmd, gotProfileCmd, gotHostCmd := buildCliCommands(cliPath, tc.cfg)
			if !slices.Equal(gotForceCmd, tc.wantForceCmd) {
				t.Errorf("force cmd = %v, want %v", gotForceCmd, tc.wantForceCmd)
			}
			if !slices.Equal(gotProfileCmd, tc.wantProfileCmd) {
				t.Errorf("profile cmd = %v, want %v", gotProfileCmd, tc.wantProfileCmd)
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
		if ts.hostCmd[0] != validCliPath {
			t.Errorf("hostCmd[0] = %q, want %q", ts.hostCmd[0], validCliPath)
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

			ts := &CliTokenSource{hostCmd: []string{mockScript}}
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

func TestCliTokenSource_Token_Fallback(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping shell script test on Windows")
	}

	expiry := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	successResponse, _ := json.Marshal(cliTokenResponse{
		AccessToken: "test-token",
		TokenType:   "Bearer",
		Expiry:      expiry,
	})

	success := "#!/bin/sh\necho '" + string(successResponse) + "'"
	unknownForceRefresh := "#!/bin/sh\necho 'Error: unknown flag: --force-refresh' >&2\nexit 1"
	unknownProfile := "#!/bin/sh\necho 'Error: unknown flag: --profile' >&2\nexit 1"
	realError := "#!/bin/sh\necho 'cache: databricks OAuth is not configured for this host' >&2\nexit 1"
	unreachable := "#!/bin/sh\necho 'should not reach here' >&2\nexit 1"

	testCases := []struct {
		name          string
		forceScript   string // empty means nil forceCmd
		profileScript string // empty means nil profileCmd
		hostScript    string // empty means nil hostCmd
		wantToken     string
		wantErrMsg    string
	}{
		{
			name:          "force-refresh succeeds",
			forceScript:   success,
			profileScript: unreachable,
			hostScript:    unreachable,
			wantToken:     "test-token",
		},
		{
			name:          "force-refresh falls back to profile",
			forceScript:   unknownForceRefresh,
			profileScript: success,
			wantToken:     "test-token",
		},
		{
			name:        "force-refresh falls back to host (no profile)",
			forceScript: unknownForceRefresh,
			hostScript:  success,
			wantToken:   "test-token",
		},
		{
			name:          "profile falls back to host",
			profileScript: unknownProfile,
			hostScript:    success,
			wantToken:     "test-token",
		},
		{
			name:          "full fallback chain: force -> profile -> host",
			forceScript:   unknownProfile,
			profileScript: unknownProfile,
			hostScript:    success,
			wantToken:     "test-token",
		},
		{
			name:          "real error stops fallback",
			forceScript:   realError,
			profileScript: unreachable,
			hostScript:    unreachable,
			wantErrMsg:    "databricks OAuth is not configured",
		},
		{
			name:          "nil hostCmd after profile failure returns error",
			forceScript:   unknownProfile,
			profileScript: unknownProfile,
			wantErrMsg:    "no CLI commands available",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			var ts CliTokenSource

			if tc.forceScript != "" {
				path := filepath.Join(tempDir, "force_cli.sh")
				if err := os.WriteFile(path, []byte(tc.forceScript), 0755); err != nil {
					t.Fatalf("failed to create force script: %v", err)
				}
				ts.forceCmd = []string{path}
			}
			if tc.profileScript != "" {
				path := filepath.Join(tempDir, "profile_cli.sh")
				if err := os.WriteFile(path, []byte(tc.profileScript), 0755); err != nil {
					t.Fatalf("failed to create profile script: %v", err)
				}
				ts.profileCmd = []string{path}
			}
			if tc.hostScript != "" {
				path := filepath.Join(tempDir, "host_cli.sh")
				if err := os.WriteFile(path, []byte(tc.hostScript), 0755); err != nil {
					t.Fatalf("failed to create host script: %v", err)
				}
				ts.hostCmd = []string{path}
			}

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
