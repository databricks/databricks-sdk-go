package databricks

import (
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"github.com/databricks/databricks-sdk-go/databricks/version"
)

type Config config.Config

// Must panics if error is not nil. It's intended to be used with
// [databricks.NewWorkspaceClient] and [databricks.NewAccountClient].
func Must[T any](c T, err error) T {
	if err != nil {
		panic(err)
	}
	return c
}

// Version of this SDK
func Version() string {
	return version.Version
}

// WithProduct is expected to be set by developers to differentiate their app from others.
//
// Example setting is:
//
//	func init() {
//		databricks.WithProduct("your-product", "0.0.1")
//	}
func WithProduct(name, version string) {
	useragent.WithProduct(name, version)
}
