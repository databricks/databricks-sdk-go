package useragent

import (
	"os"
	"sync"
)

var runtimeOnce sync.Once
var runtimeVersion string

// DisableRuntimeCaching is a flag to disable caching of runtime version.
// This is useful for testing.
var DisableRuntimeCaching bool

func getRuntimeVersion() string {
	v := os.Getenv("DATABRICKS_RUNTIME_VERSION")
	return Sanitize(v)
}

func Runtime() string {
	if DisableRuntimeCaching {
		return getRuntimeVersion()
	}

	runtimeOnce.Do(func() {
		runtimeVersion = getRuntimeVersion()
	})
	return runtimeVersion
}
