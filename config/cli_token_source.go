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

// cliVersion represents a parsed Databricks CLI semver version.
type cliVersion struct {
	Major, Minor, Patch int
}

// AtLeast returns true if v is greater than or equal to other.
func (v cliVersion) AtLeast(other cliVersion) bool {
	if v.Major != other.Major {
		return v.Major > other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor > other.Minor
	}
	return v.Patch >= other.Patch
}

func (v cliVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// Minimum CLI versions for flag support.
var (
	cliVersionForProfile      = cliVersion{0, 207, 1} // https://github.com/databricks/cli/pull/855
	cliVersionForForceRefresh = cliVersion{0, 296, 0} // https://github.com/databricks/cli/pull/4767
)

// getCliVersion runs "databricks version" and parses the output.
func getCliVersion(ctx context.Context, cliPath string) (cliVersion, error) {
	cmd := exec.CommandContext(ctx, cliPath, "version")
	out, err := cmd.Output()
	if err != nil {
		return cliVersion{}, fmt.Errorf("cannot get CLI version: %w", err)
	}
	return parseCliVersion(strings.TrimSpace(string(out)))
}

// parseCliVersion parses a version string like "Databricks CLI v0.207.1".
func parseCliVersion(s string) (cliVersion, error) {
	s = strings.TrimPrefix(s, "Databricks CLI v")
	var v cliVersion
	_, err := fmt.Sscanf(s, "%d.%d.%d", &v.Major, &v.Minor, &v.Patch)
	if err != nil {
		return cliVersion{}, fmt.Errorf("cannot parse CLI version %q: %w", s, err)
	}
	return v, nil
}

// CliTokenSource fetches OAuth tokens by shelling out to the Databricks CLI.
type CliTokenSource struct {
	cmd []string
}

// NewCliTokenSource creates a [CliTokenSource] by detecting the installed CLI
// version and building the appropriate auth token command.
func NewCliTokenSource(ctx context.Context, cfg *Config) (*CliTokenSource, error) {
	cliPath, err := findDatabricksCli(cfg.DatabricksCliPath)
	if err != nil {
		return nil, err
	}
	cmd, err := resolveCliCommand(ctx, cliPath, cfg)
	if err != nil {
		return nil, err
	}
	return &CliTokenSource{cmd: cmd}, nil
}

// resolveCliCommand detects the CLI version and builds the command to execute.
func resolveCliCommand(ctx context.Context, cliPath string, cfg *Config) ([]string, error) {
	ver, err := getCliVersion(ctx, cliPath)
	if err != nil {
		ver = cliVersion{}
	}
	cmd := buildCliCommand(ctx, cliPath, cfg, ver)
	if cmd == nil {
		return nil, fmt.Errorf("cannot configure CLI token source: neither profile nor host is set")
	}
	return cmd, nil
}

// buildCliCommand constructs the CLI command for fetching an auth token.
// The CLI version determines which flags are used.
func buildCliCommand(ctx context.Context, cliPath string, cfg *Config, ver cliVersion) []string {
	var cmd []string

	// --profile is a global Cobra flag — old CLIs accept it silently but
	// fail with "cannot fetch credentials" instead of "unknown flag".
	// We use version detection to decide --profile vs --host.
	if cfg.Profile != "" {
		if ver.AtLeast(cliVersionForProfile) {
			cmd = []string{cliPath, "auth", "token", "--profile", cfg.Profile}
		} else {
			logger.Warnf(ctx, "Profile %q was specified but Databricks CLI v%s does not support --profile (requires >= v%s). Falling back to --host.", cfg.Profile, ver, cliVersionForProfile)
		}
	}

	if cmd == nil && cfg.Host != "" {
		cmd = []string{cliPath, "auth", "token", "--host", cfg.Host}
		if cfg.HostType() == AccountHost {
			cmd = append(cmd, "--account-id", cfg.AccountID)
		}
	}

	if cmd == nil {
		return nil
	}

	if ver.AtLeast(cliVersionForForceRefresh) {
		cmd = append(cmd, "--force-refresh")
	} else {
		logger.Warnf(ctx, "Databricks CLI v%s does not support --force-refresh (requires >= v%s). The CLI's token cache may provide stale tokens.", ver, cliVersionForForceRefresh)
	}

	return cmd
}

// Token fetches an OAuth token by shelling out to the Databricks CLI.
func (c *CliTokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	return c.execCliCommand(ctx, c.cmd)
}

func (c *CliTokenSource) execCliCommand(ctx context.Context, args []string) (*oauth2.Token, error) {
	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	stdout, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			// We intentionally discard exec.ExitError — the stderr text is the
			// CLI's error contract; exit codes and process state are not useful.
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
