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
	cmd []string
}

func NewCliTokenSource(cfg *Config) (*CliTokenSource, error) {
	cliPath, err := findDatabricksCli(cfg.DatabricksCliPath)
	if err != nil {
		return nil, err
	}
	cmd := buildCliCommand(cliPath, cfg)
	return &CliTokenSource{cmd: cmd}, nil
}

// buildCliCommand constructs the CLI command arguments for fetching an auth token.
// It handles unified hosts (with optional account_id and workspace_id), account hosts,
// and workspace hosts.
func buildCliCommand(cliPath string, cfg *Config) []string {
	cmd := []string{cliPath, "auth", "token", "--host", cfg.Host}
	switch cfg.HostType() {
	case UnifiedHost:
		cmd = append(cmd, "--experimental-is-unified-host")
		if cfg.AccountID != "" {
			cmd = append(cmd, "--account-id", cfg.AccountID)
		}
		if cfg.WorkspaceId != "" {
			cmd = append(cmd, "--workspace-id", cfg.WorkspaceId)
		}
	case AccountHost:
		cmd = append(cmd, "--account-id", cfg.AccountID)
	}
	return cmd
}

func (c *CliTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	cmd := exec.CommandContext(ctx, c.cmd[0], c.cmd[1:]...)
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
