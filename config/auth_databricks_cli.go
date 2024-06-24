package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

type DatabricksCliCredentials struct {
}

func (c DatabricksCliCredentials) Name() string {
	return "databricks-cli"
}

func (c DatabricksCliCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" {
		return nil, nil
	}

	ts, err := newDatabricksCliTokenSource(ctx, cfg)
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
		if strings.Contains(err.Error(), "no configuration file found at") {
			// databricks auth token produced this error message between
			// v0.207.1 and v0.209.1
			return nil, nil
		}
		if strings.Contains(err.Error(), "databricks OAuth is not") {
			// OAuth is not configured or not supported
			return nil, nil
		}
		return nil, err
	}
	logger.Debugf(ctx, "Using Databricks CLI authentication with Databricks OAuth tokens")
	visitor := refreshableVisitor(ts)
	return credentials.NewOAuthCredentialsProvider(visitor, ts.Token), nil
}

var errLegacyDatabricksCli = errors.New("legacy Databricks CLI detected")

type databricksCliTokenSource struct {
	ctx  context.Context
	name string
	cfg  *Config
}

func newDatabricksCliTokenSource(ctx context.Context, cfg *Config) (*databricksCliTokenSource, error) {
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

	return &databricksCliTokenSource{ctx: ctx, name: path, cfg: cfg}, nil
}

func (ts *databricksCliTokenSource) Token() (*oauth2.Token, error) {
	baseArgs := []string{"auth", "token"}
	if ts.cfg.IsAccountClient() {
		args := append(baseArgs, "--host", ts.cfg.Host, "--account-id", ts.cfg.AccountID)
		return ts.tokenInner(args)
	}
	// Try workspace-level auth first, falling back to account-level auth if account ID is available
	args := append(baseArgs, "--host", ts.cfg.Host)
	t, wsErr := ts.tokenInner(args)
	if wsErr == nil {
		return t, nil
	}
	if ts.cfg.AccountID == "" {
		return nil, wsErr
	}
	logger.Debugf(ts.ctx, "account ID available, falling back to account-level authentication")
	args = append(baseArgs, "--host", ts.cfg.Environment().AccountsHost(), "--account-id", ts.cfg.AccountID)
	t, acctErr := ts.tokenInner(args)
	if acctErr == nil {
		return t, nil
	}
	return nil, acctErr
}

func (ts *databricksCliTokenSource) tokenInner(args []string) (*oauth2.Token, error) {
	logger.Debugf(ts.ctx, "running command: '%s %s'", ts.name, strings.Join(args, " "))
	out, err := exec.Command(ts.name, args...).Output()
	if ee, ok := err.(*exec.ExitError); ok {
		logger.Debugf(ts.ctx, "command '%s %s' failed: %s", ts.name, strings.Join(args, " "), string(ee.Stderr))
		return nil, fmt.Errorf("cannot get access token: %s", string(ee.Stderr))
	}
	if err != nil {
		logger.Debugf(ts.ctx, "command '%s %s' failed to run: %w", ts.name, strings.Join(args, " "), err)
		return nil, fmt.Errorf("cannot get access token: %w", err)
	}
	var t oauth2.Token
	err = json.Unmarshal(out, &t)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal Databricks CLI result: %w", err)
	}
	logger.Infof(context.Background(), "Refreshed OAuth token from Databricks CLI, expires on %s", t.Expiry)
	return &t, nil
}
