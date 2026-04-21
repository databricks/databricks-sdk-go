package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/exp/slices"
)

// captureLogger records Info/Warn messages for assertion in tests.
// Trace/Debug/Error are no-ops as they are not exercised by the CLI
// token source logic under test.
type captureLogger struct {
	infos []string
	warns []string
}

func (l *captureLogger) Enabled(context.Context, logger.Level) bool { return true }
func (l *captureLogger) Tracef(context.Context, string, ...any)     {}
func (l *captureLogger) Debugf(context.Context, string, ...any)     {}
func (l *captureLogger) Infof(_ context.Context, format string, v ...any) {
	l.infos = append(l.infos, fmt.Sprintf(format, v...))
}
func (l *captureLogger) Warnf(_ context.Context, format string, v ...any) {
	l.warns = append(l.warns, fmt.Sprintf(format, v...))
}
func (l *captureLogger) Errorf(context.Context, string, ...any) {}

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
		want    string
		wantErr bool
	}{
		{"standard version", "Databricks CLI v0.295.0", "v0.295.0", false},
		{"patch version", "Databricks CLI v0.207.1", "v0.207.1", false},
		{"major version", "Databricks CLI v1.0.0", "v1.0.0", false},
		{"dev build bare", "Databricks CLI v0.0.0-dev", "v0.0.0-dev", false},
		{"dev build with commit", "Databricks CLI v0.0.0-dev+abc123def456", "v0.0.0-dev+abc123def456", false},
		{"empty string", "", "", true},
		{"malformed", "not a version", "", true},
		{"missing v prefix", "Databricks CLI 0.207.1", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseCliVersion(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Errorf("parseCliVersion(%q) = %q, want error", tc.input, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseCliVersion(%q) error = %v", tc.input, err)
			}
			if got != tc.want {
				t.Errorf("parseCliVersion(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}

func TestDisplayVersion(t *testing.T) {
	if got := displayVersion("v0.207.1"); got != "v0.207.1" {
		t.Errorf("displayVersion(%q) = %q, want %q", "v0.207.1", got, "v0.207.1")
	}
	if got := displayVersion(""); got != "unknown" {
		t.Errorf("displayVersion(%q) = %q, want %q", "", got, "unknown")
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
		ver     string
		wantCmd []string
	}{
		{
			name:    "host only",
			cfg:     &Config{Host: host},
			ver:     "v0.200.0",
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "account host",
			cfg:     &Config{Host: accountHost, AccountID: accountID},
			ver:     "v0.200.0",
			wantCmd: []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID},
		},
		{
			name: "former unified host treated as workspace",
			cfg: &Config{
				Host:        unifiedHost,
				AccountID:   accountID,
				WorkspaceID: workspaceID,
			},
			ver:     "v0.295.0",
			wantCmd: []string{cliPath, "auth", "token", "--host", unifiedHost},
		},
		{
			name:    "profile with new CLI — uses --profile",
			cfg:     &Config{Profile: "my-profile", Host: host},
			ver:     "v0.207.1",
			wantCmd: []string{cliPath, "auth", "token", "--profile", "my-profile"},
		},
		{
			name:    "profile with old CLI — falls back to --host",
			cfg:     &Config{Profile: "my-profile", Host: host},
			ver:     "v0.207.0",
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "profile without host and old CLI — nil",
			cfg:     &Config{Profile: "my-profile"},
			ver:     "v0.207.0",
			wantCmd: nil,
		},
		{
			name:    "profile without host and new CLI — uses --profile",
			cfg:     &Config{Profile: "my-profile"},
			ver:     "v0.207.1",
			wantCmd: []string{cliPath, "auth", "token", "--profile", "my-profile"},
		},
		{
			name:    "unknown version (detection failed) — falls back to --host",
			cfg:     &Config{Profile: "my-profile", Host: host},
			ver:     "",
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "neither profile nor host — nil",
			cfg:     &Config{},
			ver:     "v0.295.0",
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

func TestBuildHostCommand(t *testing.T) {
	const (
		cliPath     = "/path/to/databricks"
		host        = "https://workspace.cloud.databricks.com"
		accountHost = "https://accounts.cloud.databricks.com"
		accountID   = "abc-123"
	)

	testCases := []struct {
		name    string
		cfg     *Config
		wantCmd []string
	}{
		{
			name:    "workspace host",
			cfg:     &Config{Host: host},
			wantCmd: []string{cliPath, "auth", "token", "--host", host},
		},
		{
			name:    "account host appends --account-id",
			cfg:     &Config{Host: accountHost, AccountID: accountID},
			wantCmd: []string{cliPath, "auth", "token", "--host", accountHost, "--account-id", accountID},
		},
		{
			name:    "no host — nil",
			cfg:     &Config{},
			wantCmd: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := buildHostCommand(cliPath, tc.cfg)
			if !slices.Equal(got, tc.wantCmd) {
				t.Errorf("buildHostCommand() = %v, want %v", got, tc.wantCmd)
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
	nonexistent := filepath.Join(tempDir, "nonexistent")

	testCases := []struct {
		name             string
		cfg              *Config
		wantErr          error  // for errors.Is comparison
		wantErrSubstring string // for substring match when wantErr is nil
		wantCmdFlag      string // a flag that should appear in ts.cmd
	}{
		{
			name:        "success with host",
			cfg:         &Config{DatabricksCliPath: cliScript, Host: "https://example.databricks.com"},
			wantCmdFlag: "--host",
		},
		{
			name:        "success with profile",
			cfg:         &Config{DatabricksCliPath: cliScript, Profile: "my-profile", Host: "https://example.databricks.com"},
			wantCmdFlag: "--profile",
		},
		{
			name:    "CLI not found",
			cfg:     &Config{DatabricksCliPath: nonexistent, Host: "https://example.databricks.com"},
			wantErr: ErrCliNotFound,
		},
		{
			name:             "neither profile nor host",
			cfg:              &Config{DatabricksCliPath: cliScript},
			wantErrSubstring: "neither profile nor host",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts, err := NewCliTokenSource(context.Background(), tc.cfg)

			if tc.wantErr != nil {
				if !errors.Is(err, tc.wantErr) {
					t.Errorf("NewCliTokenSource() error = %v, want %v", err, tc.wantErr)
				}
				return
			}
			if tc.wantErrSubstring != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErrSubstring) {
					t.Errorf("NewCliTokenSource() error = %v, want error containing %q", err, tc.wantErrSubstring)
				}
				return
			}
			if err != nil {
				t.Fatalf("NewCliTokenSource() unexpected error: %v", err)
			}
			if tc.wantCmdFlag != "" && !slices.Contains(ts.cmd, tc.wantCmdFlag) {
				t.Errorf("cmd = %v, expected %q flag", ts.cmd, tc.wantCmdFlag)
			}
		})
	}
}

func TestNewCliTokenSource_DevBuildInfoLog(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping shell script test on Windows")
	}

	testCases := []struct {
		name             string
		versionOutput    string
		wantInfoContains string // empty => expect no info log
	}{
		{
			name:             "dev build bare",
			versionOutput:    "Databricks CLI v0.0.0-dev",
			wantInfoContains: "v0.0.0-dev",
		},
		{
			name:             "dev build with commit",
			versionOutput:    "Databricks CLI v0.0.0-dev+abc123def456",
			wantInfoContains: "v0.0.0-dev+abc123def456",
		},
		{
			name:          "released version",
			versionOutput: "Databricks CLI v0.295.0",
		},
		{
			name:          "pre-release on recent base",
			versionOutput: "Databricks CLI v0.296.0-rc1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			cliScript := writeMockCli(t, tempDir, "#!/bin/sh\necho '"+tc.versionOutput+"'")

			log := &captureLogger{}
			ctx := logger.NewContext(context.Background(), log)
			cfg := &Config{DatabricksCliPath: cliScript, Host: "https://example.databricks.com"}

			if _, err := NewCliTokenSource(ctx, cfg); err != nil {
				t.Fatalf("NewCliTokenSource() unexpected error: %v", err)
			}

			if tc.wantInfoContains == "" {
				if len(log.infos) != 0 {
					t.Errorf("expected no info log, got: %v", log.infos)
				}
				return
			}
			if len(log.infos) != 1 {
				t.Fatalf("expected exactly one info log, got %d: %v", len(log.infos), log.infos)
			}
			if !strings.Contains(log.infos[0], tc.wantInfoContains) {
				t.Errorf("info log = %q, want it to contain %q", log.infos[0], tc.wantInfoContains)
			}
			if !strings.Contains(log.infos[0], "-ldflags") {
				t.Errorf("info log = %q, want it to mention -ldflags as the resolution", log.infos[0])
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
			wantErrMsg: "cannot get access token: auth failed",
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

func TestCliTokenSource_Token_PreservesExitError(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping shell script test on Windows")
	}

	tempDir := t.TempDir()
	mockScript := filepath.Join(tempDir, "mock_cli.sh")
	if err := os.WriteFile(mockScript, []byte("#!/bin/sh\necho 'auth failed' >&2\nexit 2"), 0755); err != nil {
		t.Fatalf("failed to create mock script: %v", err)
	}

	ts := &CliTokenSource{cmd: []string{mockScript}}
	_, err := ts.Token(context.Background())
	if err == nil {
		t.Fatal("Token() returned no error, want error")
	}

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) {
		t.Fatalf("errors.As(err, &exitErr) = false, want true; err = %v", err)
	}
	if got := exitErr.ExitCode(); got != 2 {
		t.Errorf("exitErr.ExitCode() = %d, want 2", got)
	}
}
