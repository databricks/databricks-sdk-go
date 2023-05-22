package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

type DatabricksCliCredentials struct {
}

func (c DatabricksCliCredentials) Name() string {
	return "databricks-cli"
}

func (c DatabricksCliCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if !cfg.IsAws() {
		return nil, nil
	}
	if cfg.Host == "" {
		return nil, nil
	}

	ts, err := newDatabricksCliTokenSource(cfg)
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			logger.Debugf(ctx, "Most likely the Databricks CLI is not installed")
			return nil, nil
		}
		if err == errLegacyDatabricksCli {
			logger.Debugf(ctx, "Databricks CLI version <0.100.0 detected")
			return nil, nil
		}
		return nil, err
	}

	_, err = ts.Token()
	if err != nil {
		if strings.Contains(err.Error(), "databricks OAuth is not") {
			// OAuth is not configured or not supported
			return nil, nil
		}
		return nil, err
	}
	logger.Debugf(ctx, "Using Databricks CLI authentication with Databricks OAuth tokens")
	return refreshableVisitor(ts), nil
}

var errLegacyDatabricksCli = errors.New("legacy Databricks CLI detected")

type databricksCliTokenSource struct {
	name string
	args []string
}

func newDatabricksCliTokenSource(cfg *Config) (*databricksCliTokenSource, error) {
	args := []string{"auth", "token", "--host", cfg.Host}

	if cfg.IsAccountClient() {
		args = append(args, "--account-id", cfg.AccountID)
	}

	databricksCliPath := cfg.DatabricksCliPath
	if databricksCliPath == "" {
		databricksCliPath = "databricks"
	}

	// Resolve absolute path to the Databricks CLI executable.
	path, err := exec.LookPath(databricksCliPath)
	if err != nil {
		return nil, err
	}

	// Resolve symlinks in order to figure out executable size.
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		return nil, err
	}

	// Determine executable size as signal to determine old/new Databricks CLI.
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	// The new Databricks CLI is a single binary with size > 1MB.
	// We use the size as a signal to determine which Databricks CLI is installed.
	if stat.Size() < (1024 * 1024) {
		return nil, errLegacyDatabricksCli
	}

	return &databricksCliTokenSource{name: path, args: args}, nil
}

func (ts *databricksCliTokenSource) Token() (*oauth2.Token, error) {
	out, err := exec.Command(ts.name, ts.args...).Output()
	if ee, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("cannot get access token: %s", string(ee.Stderr))
	}
	if err != nil {
		return nil, fmt.Errorf("cannot get access token: %v", err)
	}
	var t oauth2.Token
	err = json.Unmarshal(out, &t)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal Databricks CLI result: %w", err)
	}
	logger.Infof(context.Background(), "Refreshed OAuth token from Databricks CLI, expires on %s", t.Expiry)
	return &t, nil
}
