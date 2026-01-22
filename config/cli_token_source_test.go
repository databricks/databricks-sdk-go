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

	testCases := []struct {
		name    string
		cfg     *Config
		wantCmd []string
		wantErr error
	}{
		{
			name:    "workspace host",
			cfg:     &Config{DatabricksCliPath: validCliPath, Host: "https://workspace.cloud.databricks.com"},
			wantCmd: []string{validCliPath, "auth", "token", "--host", "https://workspace.cloud.databricks.com"},
		},
		{
			name:    "account host",
			cfg:     &Config{DatabricksCliPath: validCliPath, Host: "https://accounts.cloud.databricks.com", AccountID: "abc-123"},
			wantCmd: []string{validCliPath, "auth", "token", "--host", "https://accounts.cloud.databricks.com", "--account-id", "abc-123"},
		},
		{
			name:    "CLI not found",
			cfg:     &Config{DatabricksCliPath: filepath.Join(tempDir, "nonexistent"), Host: "https://workspace.cloud.databricks.com"},
			wantErr: ErrCliNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := NewCliTokenSource(tc.cfg)
			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Errorf("NewCliTokenSource() error = %v, want %v", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewCliTokenSource() unexpected error: %v", err)
			}
			for i := range ts.cmd {
				if ts.cmd[i] != tc.wantCmd[i] {
					t.Errorf("cmd[%d] = %q, want %q", i, ts.cmd[i], tc.wantCmd[i])
				}
			}
		})
	}
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
