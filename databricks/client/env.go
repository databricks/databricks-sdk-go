package client

import (
	"os"
	"strings"
	"sync"
)

var (
	envMutex     sync.Mutex
	onceClient   sync.Once
	commonClient *DatabricksClient

	DefaultClient = &DatabricksClient{}
)

// ResetCommonEnvironmentClient resets test dummy
func ResetCommonEnvironmentClient() {
	commonClient = nil
	onceClient = sync.Once{}
}

// CommonEnvironmentClient configured once per run of application
func CommonEnvironmentClient() *DatabricksClient {
	if commonClient != nil {
		return commonClient
	}
	onceClient.Do(func() {
		commonClient = &DatabricksClient{}
		// err := commonClient.Authenticate(context.Background())
		// if err != nil {
		// 	panic(err)
		// }
	})
	return commonClient
}

// CleanupEnvironment backs up environment - use as `defer CleanupEnvironment()()`
// clears it and restores it in the end. It's meant strictly for "unit" tests
// as last resort, because it slows down parallel execution with mutex.
func CleanupEnvironment() func() {
	// make a backed-up pristine environment
	envMutex.Lock()
	prevEnv := os.Environ()
	oldPath := os.Getenv("PATH")
	pwd := os.Getenv("PWD")
	os.Clearenv()
	os.Setenv("PATH", oldPath)
	os.Setenv("HOME", pwd)
	// and return restore function
	return func() {
		for _, kv := range prevEnv {
			kvs := strings.SplitN(kv, "=", 2)
			os.Setenv(kvs[0], kvs[1])
		}
		envMutex.Unlock()
	}
}
