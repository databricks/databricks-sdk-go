package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
)

type BricksCliCredentials struct {
}

func (c BricksCliCredentials) Name() string {
	return "bricks-cli"
}

func (c BricksCliCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if !cfg.IsAws() {
		return nil, nil
	}
	if cfg.Host == "" {
		return nil, nil
	}
	ts := bricksCliTokenSource{cfg}
	_, err := ts.Token()
	if err != nil {
		if strings.Contains(err.Error(), "databricks OAuth is not") {
			// OAuth is not configured or not supported
			return nil, nil
		}
		if strings.Contains(err.Error(), "executable file not found") {
			logger.Debugf(ctx, "Most likely Bricks CLI is not installed")
			return nil, nil
		}
		return nil, err
	}
	logger.Debugf(ctx, "Using Bricks CLI authentication with Databricks OAuth tokens")
	return refreshableVisitor(&ts), nil
}

type bricksCliTokenSource struct {
	cfg *Config
}

func (ts *bricksCliTokenSource) Token() (*oauth2.Token, error) {
	what := []string{"auth", "token", "--host", ts.cfg.Host}

	if ts.cfg.IsAccountClient() {
		what = append(what, "--account-id", ts.cfg.AccountID)
	}

	bricksCli := ts.cfg.BricksCliPath
	if bricksCli == "" {
		bricksCli = "bricks"
	}

	out, err := exec.Command(bricksCli, what...).Output()
	if ee, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("cannot get access token: %s", string(ee.Stderr))
	}
	if err != nil {
		return nil, fmt.Errorf("cannot get access token: %v", err)
	}
	var t oauth2.Token
	err = json.Unmarshal(out, &t)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal Bricks CLI result: %w", err)
	}
	logger.Infof(context.Background(), "Refreshed OAuth token from Bricks CLI, expires on %s", t.Expiry)
	return &t, nil
}
