package lock

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/qa"
)

type LockOption func(*LockOptions)

func WithBackend(backend LockBackend) LockOption {
	return func(opts *LockOptions) {
		opts.Backend = backend
	}
}

func WithLeaseDuration(duration time.Duration) LockOption {
	return func(opts *LockOptions) {
		opts.LeaseDuration = duration
	}
}

func InTest(t *testing.T) LockOption {
	return func(opts *LockOptions) {
		opts.T = t
	}
}

type LockOptions struct {
	LeaseDuration time.Duration
	Backend       LockBackend
	T             *testing.T
}

func (opts LockOptions) getLockContents(l Lockable) ([]byte, error) {
	info := make(map[string]string)

	if opts.T != nil {
		info["Test"] = opts.T.Name()
	}
	info["Start"] = time.Now().Format(time.RFC3339)
	info["Lockable"] = l.GetLockId()
	info["LeaseDuration"] = opts.LeaseDuration.String()
	info["IsInDebug"] = strconv.FormatBool(qa.IsInDebug())
	envVars := []string{
		"USER",
		"GITHUB_ACTIONS",
		"GITHUB_REF",
		"GITHUB_RUN_ID",
		"GITHUB_RUN_NUMBER",
		"GITHUB_SHA",
		"GITHUB_WORKFLOW",
	}
	for _, envVar := range envVars {
		val := os.Getenv(envVar)
		if val != "" {
			info[envVar] = val
		}
	}

	return json.MarshalIndent(info, "", "  ")
}
