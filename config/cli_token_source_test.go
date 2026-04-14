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

func TestParseCliVersion(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		want    cliVersion
		wantErr bool
	}{
		{
			name:  "standard version",
			input: "Databricks CLI v0.295.0",
			want:  cliVersion{0, 295, 0},
		},
		{
			name:  "patch version",
			input: "Databricks CLI v0.207.1",
			want:  cliVersion{0, 207, 1},
		},
		{
			name:  "major version",
			input: "Databricks CLI v1.0.0",
			want:  cliVersion{1, 0, 0},
		},
		{
			name:    "empty string",
			input:   "",
			wantErr: true,
		},
		{
			name:    "malformed",
			input:   "not a version",
			wantErr: true,
		},
		{
			name:    "missing prefix",
			input:   "v0.207.1",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseCliVersion(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Errorf("parseCliVersion(%q) = %v, want error", tc.input, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseCliVersion(%q) error = %v", tc.input, err)
			}
			if got != tc.want {
				t.Errorf("parseCliVersion(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestCliVersion_AtLeast(t *testing.T) {
	testCases := []struct {
		name  string
		v     cliVersion
		other cliVersion
		want  bool
	}{
		{"equal", cliVersion{0, 207, 1}, cliVersion{0, 207, 1}, true},
		{"higher patch", cliVersion{0, 207, 2}, cliVersion{0, 207, 1}, true},
		{"lower patch", cliVersion{0, 207, 0}, cliVersion{0, 207, 1}, false},
		{"higher minor", cliVersion{0, 208, 0}, cliVersion{0, 207, 1}, true},
		{"lower minor", cliVersion{0, 206, 9}, cliVersion{0, 207, 1}, false},
		{"higher major", cliVersion{1, 0, 0}, cliVersion{0, 207, 1}, true},
		{"zero vs zero", cliVersion{0, 0, 0}, cliVersion{0, 0, 0}, true},
		{"zero vs nonzero", cliVersion{0, 0, 0}, cliVersion{0, 207, 1}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.v.AtLeast(tc.other); got != tc.want {
				t.Errorf("%v.AtLeast(%v) = %v, want %v", tc.v, tc.other, got, tc.want)
			}
		})
	}
}

func TestBuildCliCommand(t *testing.T) {
	const (
		cliPath     = "/path/to/databricks"
		host        = "https://workspace.cloud.databricks.com"
		accountHost = "https://accounts.cloud.databricks.com"
		unifiedHost = "https://unified.cloud.databricks.com"
		accountID   = "abc-123"
		workspaceID = "456"
	)

	testCases := []struct {
		name    string
		cfg     *Config
		ver     cliVersion
		wantCmd []string
	}{
		{
			name:    "host only — old CLI, no force-refresh",
			cfg:     &Config{Host: host},
			ver:     cliVersion{0, 200, 0},
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "host only — new CLI, with force-refresh",
			cfg:     &Config{Host: host},
			ver:     cliVersion{0, 296, 0},
			wantCmd: []string{cliPath, "auth", "token", "--host", host, "--force-refresh"},
		},
		{
			name:    "account host — with force-refresh",
			cfg:     &Config{Host: accountHost, AccountID: accountID},
			ver:     cliVersion{0, 296, 0},
			wantCmd: []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID, "--force-refresh"},
		},
		{
			name:    "account host — without force-refresh",
			cfg:     &Config{Host: accountHost, AccountID: accountID},
			ver:     cliVersion{0, 200, 0},
			wantCmd: []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID},
		},
		{
			name: "former unified host treated as workspace",
			cfg: &Config{
				Host:        unifiedHost,
				AccountID:   accountID,
				WorkspaceID: workspaceID,
			},
			ver:     cliVersion{0, 295, 0},
			wantCmd: []string{cliPath, "auth", "token", "--host", unifiedHost},
		},
		{
			name:    "profile with new CLI — uses --profile, no force-refresh",
			cfg:     &Config{Profile: "my-profile", Host: host},
			ver:     cliVersion{0, 207, 1},
			wantCmd: []string{cliPath, "auth", "token", "--profile", "my-profile"},
		},
		{
			name:    "profile with newest CLI — uses --profile and --force-refresh",
			cfg:     &Config{Profile: "my-profile", Host: host},
			ver:     cliVersion{0, 296, 0},
			wantCmd: []string{cliPath, "auth", "token", "--profile", "my-profile", "--force-refresh"},
		},
		{
			name:    "profile with old CLI — falls back to --host, no force-refresh",
			cfg:     &Config{Profile: "my-profile", Host: host},
			ver:     cliVersion{0, 207, 0},
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "profile without host and old CLI — nil",
			cfg:     &Config{Profile: "my-profile"},
			ver:     cliVersion{0, 207, 0},
			wantCmd: nil,
		},
		{
			name:    "profile without host and new CLI — uses --profile",
			cfg:     &Config{Profile: "my-profile"},
			ver:     cliVersion{0, 207, 1},
			wantCmd: []string{cliPath, "auth", "token", "--profile", "my-profile"},
		},
		{
			name:    "profile without host and newest CLI — --profile and --force-refresh",
			cfg:     &Config{Profile: "my-profile"},
			ver:     cliVersion{0, 296, 0},
			wantCmd: []string{cliPath, "auth", "token", "--profile", "my-profile", "--force-refresh"},
		},
		{
			name:    "zero version (detection failed) — falls back to --host, no force-refresh",
			cfg:     &Config{Profile: "my-profile", Host: host},
			ver:     cliVersion{},
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "neither profile nor host — nil",
			cfg:     &Config{},
			ver:     cliVersion{0, 296, 0},
			wantCmd: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := buildCliCommand(context.Background(), cliPath, tc.cfg, tc.ver)
			if !slices.Equal(got, tc.wantCmd) {
				t.Errorf("buildCliCommand() = %v, want %v", got, tc.wantCmd)
			}
		})
	}
}

// writeMockCli creates a shell script that passes the file-size check in
// findDatabricksCli and executes body when run.
func writeMockCli(t *testing.T, dir, body string) string {
	t.Helper()
	path := filepath.Join(dir, "databricks")
	content := append([]byte(body), make([]byte, databricksCliMinSize)...)
	if err := os.WriteFile(path, content, 0755); err != nil {
		t.Fatalf("failed to create mock CLI: %v", err)
	}
	return path
}

func TestNewCliTokenSource(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping shell script test on Windows")
	}

	tempDir := t.TempDir()
	cliScript := writeMockCli(t, tempDir, "#!/bin/sh\necho 'Databricks CLI v0.295.0'")

	t.Run("success with host", func(t *testing.T) {
		cfg := &Config{DatabricksCliPath: cliScript, Host: "https://example.databricks.com"}
		ts, err := NewCliTokenSource(context.Background(), cfg)
		if err != nil {
			t.Fatalf("NewCliTokenSource() unexpected error: %v", err)
		}
		if ts.cmd[0] != cliScript {
			t.Errorf("cmd[0] = %q, want %q", ts.cmd[0], cliScript)
		}
	})

	t.Run("CLI not found", func(t *testing.T) {
		cfg := &Config{DatabricksCliPath: filepath.Join(tempDir, "nonexistent"), Host: "https://example.databricks.com"}
		_, err := NewCliTokenSource(context.Background(), cfg)
		if !errors.Is(err, ErrCliNotFound) {
			t.Errorf("NewCliTokenSource() error = %v, want %v", err, ErrCliNotFound)
		}
	})

	t.Run("neither profile nor host", func(t *testing.T) {
		cfg := &Config{DatabricksCliPath: cliScript}
		_, err := NewCliTokenSource(context.Background(), cfg)
		if err == nil {
			t.Fatal("NewCliTokenSource() error = nil, want error")
		}
		if !strings.Contains(err.Error(), "neither profile nor host") {
			t.Errorf("NewCliTokenSource() error = %v, want error containing %q", err, "neither profile nor host")
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
