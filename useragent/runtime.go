package useragent

import (
	"os"
	"sync"
)

var runtimeOnce sync.Once
var runtimeVersion string

func getRuntimeVersion() string {
	v := os.Getenv("DATABRICKS_RUNTIME_VERSION")
	return Sanitize(v)
}

func Runtime() string {
	runtimeOnce.Do(func() {
		runtimeVersion = getRuntimeVersion()
	})
	return runtimeVersion
}
