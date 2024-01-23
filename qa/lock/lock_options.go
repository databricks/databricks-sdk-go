package lock

import (
	"os"
	"reflect"
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

type lockState struct {
	Test            string
	Start           time.Time
	Expiry          time.Time
	Lockable        string
	LeaseId         string
	IsInDebug       bool
	User            string `env:"USER"`
	GithubActions   string `env:"GITHUB_ACTIONS"`
	GithubRef       string `env:"GITHUB_REF"`
	GithubRunId     string `env:"GITHUB_RUN_ID"`
	GithubRunNumber string `env:"GITHUB_RUN_NUMBER"`
	GithubSha       string `env:"GITHUB_SHA"`
	GithubWorkflow  string `env:"GITHUB_WORKFLOW"`
}

func (l *lockState) loadFromEnv() {
	rv := reflect.ValueOf(l).Elem()
	t := rv.Type()
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		tag := t.Field(i).Tag.Get("env")
		if tag == "" {
			continue
		}
		val := os.Getenv(tag)
		if val != "" {
			field.SetString(val)
		}
	}
}

func (l *lockState) IsValid() bool {
	return time.Now().Before(l.Expiry)
}

func (l *lockState) Extend(duration time.Duration) {
	l.Expiry = time.Now().Add(duration)
}

func newLockState(lockable Lockable, leaseId string, leaseDuration time.Duration, t *testing.T) *lockState {
	now := time.Now()
	state := &lockState{
		Start:     now,
		Expiry:    now.Add(leaseDuration),
		Lockable:  lockable.GetLockId(),
		LeaseId:   leaseId,
		IsInDebug: qa.IsInDebug(),
	}
	state.loadFromEnv()
	if t != nil {
		state.Test = t.Name()
	}
	return state
}
