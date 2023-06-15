package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/logger"
)

const fullCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const hexCharset = "0123456789abcdef"

func init() {
	databricks.WithProduct("integration-tests", databricks.Version())
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelDebug,
	}
}

// prelude for all workspace-level tests
func workspaceTest(t *testing.T) (context.Context, *databricks.WorkspaceClient) {
	loadDebugEnvIfRunsFromIDE(t, "workspace")
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	if os.Getenv("DATABRICKS_ACCOUNT_ID") != "" {
		skipf(t)("Skipping workspace test on account level")
	}
	t.Parallel()
	ctx := context.Background()
	return ctx, databricks.Must(databricks.NewWorkspaceClient())
}

// prelude for all workspace-level UC tests
func ucwsTest(t *testing.T) (context.Context, *databricks.WorkspaceClient) {
	loadDebugEnvIfRunsFromIDE(t, "ucws")
	if os.Getenv("DATABRICKS_ACCOUNT_ID") != "" {
		skipf(t)("Skipping workspace test on account level")
	}
	GetEnvOrSkipTest(t, "TEST_METASTORE_ID")
	t.Parallel()
	ctx := context.Background()
	return ctx, databricks.Must(databricks.NewWorkspaceClient())
}

// prelude for all account-level tests
func accountTest(t *testing.T) (context.Context, *databricks.AccountClient) {
	loadDebugEnvIfRunsFromIDE(t, "account")
	cfg := &config.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	}
	err := cfg.EnsureResolved()
	if err != nil {
		skipf(t)("error: %s", err)
	}
	if !cfg.IsAccountClient() {
		skipf(t)("Not in account env: %s/%s", cfg.AccountID, cfg.Host)
	}
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()
	ctx := context.Background()
	return ctx, databricks.Must(databricks.NewAccountClient(
		(*databricks.Config)(cfg)))
}

// prelude for all UC account-level tests
func ucacctTest(t *testing.T) (context.Context, *databricks.AccountClient) {
	loadDebugEnvIfRunsFromIDE(t, "ucacct")
	cfg := &config.Config{
		AccountID: GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID"),
	}
	err := cfg.EnsureResolved()
	if err != nil {
		skipf(t)("error: %s", err)
	}
	if !cfg.IsAccountClient() {
		skipf(t)("Not in account env: %s/%s", cfg.AccountID, cfg.Host)
	}
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	t.Parallel()
	ctx := context.Background()
	return ctx, databricks.Must(databricks.NewAccountClient(
		(*databricks.Config)(cfg)))
}

// GetEnvOrSkipTest proceeds with test only with that env variable
func GetEnvOrSkipTest(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		skipf(t)("Environment variable %s is missing", name)
	}
	return value
}

func MustParseInt64(v string) int64 {
	i, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		panic(fmt.Sprintf("`%s` is not int64: %s", v, err))
	}
	return i
}

// RandomEmail generates random email
func RandomEmail(prefix ...string) string {
	return fmt.Sprintf("%s@example.com", RandomName(
		append([]string{"sdk-go-"}, prefix...)...))
}

// RandomName gives random name with optional prefix. e.g. qa.RandomName("tf-")
func RandomName(prefix ...string) string {
	rand.Seed(time.Now().UnixNano())
	randLen := 12
	b := make([]byte, randLen)
	for i := range b {
		b[i] = fullCharset[rand.Intn(randLen)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", strings.Join(prefix, ""), b)
	}
	return string(b)
}

func RandomHex(prefix string, randLen int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, randLen)
	for i := range b {
		b[i] = hexCharset[rand.Intn(randLen)%len(hexCharset)]
	}
	if len(prefix) > 0 {
		return fmt.Sprintf("%s%s", prefix, b)
	}
	return string(b)
}

func skipf(t *testing.T) func(format string, args ...any) {
	if isInDebug() {
		// VSCode "debug test" feature doesn't show dlv logs,
		// so that we fail here for maintainer productivity.
		return t.Fatalf
	}
	return t.Skipf
}

// detects if test is run from "debug test" feature in VSCode
func isInDebug() bool {
	ex, _ := os.Executable()
	return path.Base(ex) == "__debug_bin"
}

// loads debug environment from ~/.databricks/debug-env.json
func loadDebugEnvIfRunsFromIDE(t *testing.T, key string) {
	if !isInDebug() {
		return
	}
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("cannot find user home: %s", err)
	}
	raw, err := os.ReadFile(filepath.Join(home, ".databricks/debug-env.json"))
	if err != nil {
		t.Fatalf("cannot load ~/.databricks/debug-env.json: %s", err)
	}
	var conf map[string]map[string]string
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		t.Fatalf("cannot parse ~/.databricks/debug-env.json: %s", err)
	}
	vars, ok := conf[key]
	if !ok {
		t.Fatalf("~/.databricks/debug-env.json#%s not configured", key)
	}
	for k, v := range vars {
		os.Setenv(k, v)
	}
}
