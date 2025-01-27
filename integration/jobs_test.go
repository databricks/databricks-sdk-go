package integration

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/log"
	"github.com/databricks/databricks-sdk-go/jobs/v2"
	"github.com/stretchr/testify/assert"
)

func IsInDebug() bool {
	ex, _ := os.Executable()
	return strings.HasPrefix(path.Base(ex), "__debug_bin")
}

// loads debug environment from ~/.databricks/debug-env.json
func loadDebugEnvIfRunsFromIDE(t *testing.T, key string) {
	if !IsInDebug() {
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

// GetEnvOrSkipTest proceeds with test only with that env variable
func GetEnvOrSkipTest(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		skipf(t)("Environment variable %s is missing", name)
	}
	return value
}

func skipf(t *testing.T) func(format string, args ...any) {
	if IsInDebug() {
		// VSCode "debug test" feature doesn't show dlv logs,
		// so that we fail here for maintainer productivity.
		return t.Fatalf
	}
	return t.Skipf
}

// prelude for all workspace-level tests
func workspaceTest(t *testing.T) context.Context {
	loadDebugEnvIfRunsFromIDE(t, "workspace")
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	if os.Getenv("DATABRICKS_ACCOUNT_ID") != "" {
		skipf(t)("Skipping workspace test on account level")
	}
	t.Parallel()
	ctx := context.Background()
	// return ctx, databricks.Must(databricks.NewWorkspaceClient())
	return ctx
}

func init() {
	logger := log.New(log.LevelDebug) // drops trace and debug logs
	log.SetDefaultLogger(logger)
}

func TestAccJobsGetCorrectErrorNoTranspile(t *testing.T) {
	ctx := workspaceTest(t)
	client, err := jobs.NewJobsClientFromConfig()
	assert.NoError(t, err)
	assert.NoError(t, err)
	resp, err := client.Jobs.Create(ctx, jobs.CreateJob{
		Name: "test-job",
	})
	assert.NoError(t, err)
	job, err := client.Jobs.GetByJobId(ctx, resp.JobId)
	assert.NoError(t, err)
	assert.Equal(t, job.JobId, resp.JobId)
}
