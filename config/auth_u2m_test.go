package config

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestU2MCredentials_Name(t *testing.T) {
	if got := (u2mCredentials{}).Name(); got != "databricks-cli" {
		t.Errorf("Name() = %q, want %q", got, "databricks-cli")
	}
}

func TestU2MCredentials_Configure(t *testing.T) {
	testCases := []struct {
		name string
		// skip is a function that returns true if the test should be skipped.
		skip func() bool
		// setup prepares the test environment and returns the config to use.
		// The cleanup function is called after the test.
		setup func(t *testing.T) (cfg *Config, cleanup func())
		// wantErr is a substring expected in the error message, or empty if no error expected.
		wantErr string
		// wantSkip is true if Configure should return nil, nil (skip this auth method).
		wantSkip bool
		// wantProvider is true if Configure should return a non-nil CredentialsProvider.
		wantProvider bool
	}{
		{
			name: "missing host returns error",
			setup: func(t *testing.T) (*Config, func()) {
				return &Config{}, func() {}
			},
			wantErr: "host is required",
		},
		{
			name: "CLI not found skips auth",
			setup: func(t *testing.T) (*Config, func()) {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: "/nonexistent/path/to/databricks",
				}, func() {}
			},
			wantSkip: true,
		},
		{
			name: "CLI not found returns error when auth type explicit",
			setup: func(t *testing.T) (*Config, func()) {
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: "/nonexistent/path/to/databricks",
					AuthType:          "databricks-cli",
				}, func() {}
			},
			wantErr: "databricks CLI not found",
		},
		{
			name: "legacy CLI detected skips auth",
			setup: func(t *testing.T) (*Config, func()) {
				tempDir := t.TempDir()
				cliName := "databricks"
				if runtime.GOOS == "windows" {
					cliName = "databricks.exe"
				}
				legacyCliPath := filepath.Join(tempDir, cliName)
				if err := os.WriteFile(legacyCliPath, make([]byte, 100), 0755); err != nil {
					t.Fatalf("failed to create legacy CLI file: %v", err)
				}
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: legacyCliPath,
				}, func() {}
			},
			wantSkip: true,
		},
		{
			name: "OAuth not configured skips auth",
			skip: func() bool { return runtime.GOOS == "windows" },
			setup: func(t *testing.T) (*Config, func()) {
				mockCli := createMockCli(t, `#!/bin/sh
echo "Error: databricks OAuth is not configured for this host" >&2
exit 1`)
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: mockCli,
				}, func() {}
			},
			wantSkip: true,
		},
		{
			name: "OAuth not configured returns error when auth type explicit",
			skip: func() bool { return runtime.GOOS == "windows" },
			setup: func(t *testing.T) (*Config, func()) {
				mockCli := createMockCli(t, `#!/bin/sh
echo "Error: databricks OAuth is not configured for this host" >&2
exit 1`)
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: mockCli,
					AuthType:          "databricks-cli",
				}, func() {}
			},
			wantErr: "databricks OAuth is not",
		},
		{
			name: "token error passes through CLI error",
			skip: func() bool { return runtime.GOOS == "windows" },
			setup: func(t *testing.T) (*Config, func()) {
				mockCli := createMockCli(t, `#!/bin/sh
cat >&2 << 'EOF'
A new access token could not be retrieved because the refresh token is invalid. To reauthenticate, run the following command:
  $ databricks auth login --host https://workspace.cloud.databricks.com
EOF
exit 1`)
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: mockCli,
				}, func() {}
			},
			wantErr: "refresh token is invalid",
		},
		{
			name: "successful authentication",
			skip: func() bool { return runtime.GOOS == "windows" },
			setup: func(t *testing.T) (*Config, func()) {
				expiry := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
				response := cliTokenResponse{
					AccessToken: "test-access-token",
					TokenType:   "Bearer",
					Expiry:      expiry,
				}
				jsonResponse, _ := json.Marshal(response)
				mockCli := createMockCli(t, "#!/bin/sh\necho '"+string(jsonResponse)+"'")
				return &Config{
					Host:              "https://workspace.cloud.databricks.com",
					DatabricksCliPath: mockCli,
				}, func() {}
			},
			wantProvider: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.skip != nil && tc.skip() {
				t.Skip("Skipping test on this platform")
			}

			cfg, cleanup := tc.setup(t)
			defer cleanup()

			cp, err := (u2mCredentials{}).Configure(context.Background(), cfg)

			switch {
			case tc.wantErr != "":
				if err == nil {
					t.Fatalf("Configure() error = nil, want error containing %q", tc.wantErr)
				}
				if !strings.Contains(err.Error(), tc.wantErr) {
					t.Errorf("Configure() error = %q, want error containing %q", err, tc.wantErr)
				}
			case tc.wantSkip:
				if err != nil {
					t.Errorf("Configure() error = %v, want nil (should skip auth)", err)
				}
				if cp != nil {
					t.Errorf("Configure() = %v, want nil (should skip auth)", cp)
				}
			case tc.wantProvider:
				if err != nil {
					t.Fatalf("Configure() error = %v, want nil", err)
				}
				if cp == nil {
					t.Fatal("Configure() = nil, want CredentialsProvider")
				}
			}
		})
	}
}

// createMockCli creates a mock CLI script in a temp directory and returns its path.
// The script content should be a valid shell script starting with #!/bin/sh.
func createMockCli(t *testing.T, scriptContent string) string {
	t.Helper()
	tempDir := t.TempDir()
	mockCli := filepath.Join(tempDir, "databricks")
	if err := os.WriteFile(mockCli, []byte(scriptContent), 0755); err != nil {
		t.Fatalf("failed to create mock CLI: %v", err)
	}
	if err := padFile(mockCli, databricksCliMinSize+1); err != nil {
		t.Fatalf("failed to pad mock CLI: %v", err)
	}
	return mockCli
}

// padFile appends null bytes to make the file at least minSize bytes.
// This is needed because the CLI detection checks file size to distinguish
// the new Go-based CLI from the legacy Python CLI.
func padFile(path string, minSize int64) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.Size() >= minSize {
		return nil
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	padding := make([]byte, minSize-info.Size())
	_, err = f.Write(padding)
	return err
}
