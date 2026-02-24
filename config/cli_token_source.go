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
	"time"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

// databricksCliMinSize distinguishes the new Go CLI (>= 0.100.0) from the
// legacy Python CLI by file size.
const databricksCliMinSize = 1024 * 1024

var (
	ErrCliNotFound       = errors.New("databricks CLI not found")
	ErrLegacyCliDetected = errors.New("legacy databricks CLI detected; upgrade to >= 0.100.0")
)

type cliTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiry      string `json:"expiry"`
}

type CliTokenSource struct {
	// cmd is the primary command to execute (--profile when available, --host otherwise).
	cmd []string

	// hostCmd is a fallback command using --host, used when the primary --profile
	// command fails because the CLI is too old to support --profile.
	hostCmd []string
}

func NewCliTokenSource(cfg *Config) (*CliTokenSource, error) {
	cliPath, err := findDatabricksCli(cfg.DatabricksCliPath)
	if err != nil {
		return nil, err
	}
	profileCmd, hostCmd := buildCliCommands(cliPath, cfg)
	return &CliTokenSource{cmd: profileCmd, hostCmd: hostCmd}, nil
}

// buildCliCommands constructs the CLI commands for fetching an auth token.
// When cfg.Profile is set, the primary command uses --profile and a fallback
// --host command is also returned for compatibility with older CLIs.
// When cfg.Profile is empty, the primary command uses --host and no fallback
// is needed.
func buildCliCommands(cliPath string, cfg *Config) (primaryCmd []string, hostCmd []string) {
	if cfg.Profile != "" {
		primary := []string{cliPath, "auth", "token", "--profile", cfg.Profile}
		if cfg.Host != "" {
			// Build a --host fallback for old CLIs that don't support --profile.
			return primary, buildHostCommand(cliPath, cfg)
		}
		return primary, nil
	}
	return buildHostCommand(cliPath, cfg), nil
}

// buildHostCommand constructs the legacy --host based CLI command.
func buildHostCommand(cliPath string, cfg *Config) []string {
	cmd := []string{cliPath, "auth", "token", "--host", cfg.Host}
	switch cfg.HostType() {
	case UnifiedHost:
		cmd = append(cmd, "--experimental-is-unified-host")
		if cfg.AccountID != "" {
			cmd = append(cmd, "--account-id", cfg.AccountID)
		}
		if cfg.WorkspaceID != "" {
			cmd = append(cmd, "--workspace-id", cfg.WorkspaceID)
		}
	case AccountHost:
		cmd = append(cmd, "--account-id", cfg.AccountID)
	}
	return cmd
}

// Token fetches an OAuth token by shelling out to the Databricks CLI.
// When a --profile command is configured, it is tried first. If the CLI
// returns "unknown flag: --profile" (indicating an older CLI version),
// the fallback --host command is used instead.
func (c *CliTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	tok, err := c.execCliCommand(ctx, c.cmd)
	if err != nil && c.hostCmd != nil && isUnknownFlagError(err) {
		logger.Warnf(ctx, "Databricks CLI does not support --profile flag. Falling back to --host. Please upgrade your CLI to the latest version.")
		return c.execCliCommand(ctx, c.hostCmd)
	}
	return tok, err
}

func (c *CliTokenSource) execCliCommand(ctx context.Context, args []string) (*oauth2.Token, error) {
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return nil, fmt.Errorf("cannot get access token: %s", strings.TrimSpace(string(exitErr.Stderr)))
		}
		return nil, fmt.Errorf("cannot get access token: %w", err)
	}
	var resp cliTokenResponse
	if err := json.Unmarshal(stdout, &resp); err != nil {
		return nil, fmt.Errorf("cannot parse CLI response: %w", err)
	}
	expiry, err := parseExpiry(resp.Expiry)
	if err != nil {
		return nil, fmt.Errorf("cannot parse token expiry: %w", err)
	}
	return &oauth2.Token{
		AccessToken: resp.AccessToken,
		TokenType:   resp.TokenType,
		Expiry:      expiry,
	}, nil
}

// isUnknownFlagError returns true if the error indicates the CLI does not
// recognize the --profile flag. This happens with older CLI versions that
// predate profile-based token lookup.
func isUnknownFlagError(err error) bool {
	return strings.Contains(err.Error(), "unknown flag: --profile")
}

// parseExpiry parses an expiry time string in multiple formats for cross-SDK compatibility.
//
// Supported formats:
//  1. RFC 3339 with nanoseconds (e.g., "2024-03-20T10:30:00.123456789Z")
//  2. RFC 3339 (e.g., "2024-03-20T10:30:00Z" or "2024-03-20T10:30:00+01:00")
//  3. Space-separated without fractional seconds (e.g., "2024-03-20 10:30:00")
//  4. Space-separated with fractional seconds (e.g., "2024-03-20 10:30:00.123")
//
// The Databricks CLI outputs time.RFC3339Nano format.
//
// Formats 3 and 4 are supported for parity with the Python and Java SDKs, which accept
// space-separated datetime formats and assume the local timezone.
func parseExpiry(expiry string) (time.Time, error) {
	if t, err := time.Parse(time.RFC3339Nano, expiry); err == nil {
		return t, nil
	}
	if t, err := time.Parse(time.RFC3339, expiry); err == nil {
		return t, nil
	}

	if t, err := time.Parse("2006-01-02 15:04:05", expiry); err == nil {
		return t, nil
	}

	layouts := []string{
		"2006-01-02 15:04:05.9",
		"2006-01-02 15:04:05.99",
		"2006-01-02 15:04:05.999",
		"2006-01-02 15:04:05.9999",
		"2006-01-02 15:04:05.99999",
		"2006-01-02 15:04:05.999999",
		"2006-01-02 15:04:05.9999999",
		"2006-01-02 15:04:05.99999999",
		"2006-01-02 15:04:05.999999999",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, expiry); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("cannot parse expiry %q", expiry)
}

func findDatabricksCli(cliPath string) (string, error) {
	if cliPath == "" {
		path, err := findInPath("databricks")
		if err != nil && runtime.GOOS == "windows" {
			return findInPath("databricks.exe")
		}
		return path, err
	}
	if strings.Contains(cliPath, string(filepath.Separator)) || strings.Contains(cliPath, "/") {
		return validateCliPath(cliPath)
	}
	return findInPath(cliPath)
}

func findInPath(name string) (string, error) {
	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return "", ErrCliNotFound
	}
	lastErr := ErrCliNotFound
	for _, dir := range filepath.SplitList(pathEnv) {
		validPath, err := validateCliPath(filepath.Join(dir, name))
		if err == nil {
			return validPath, nil
		}
		if errors.Is(err, ErrLegacyCliDetected) {
			lastErr = err
		}
	}
	return "", lastErr
}

func validateCliPath(path string) (string, error) {
	info, err := os.Stat(path)
	if os.IsNotExist(err) || (err == nil && info.IsDir()) {
		return "", ErrCliNotFound
	}
	if err != nil {
		return "", fmt.Errorf("cannot stat CLI path: %w", err)
	}
	if info.Size() < databricksCliMinSize {
		return "", ErrLegacyCliDetected
	}
	return path, nil
}
