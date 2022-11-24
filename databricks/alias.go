package databricks

import (
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/databricks/internal"
	"github.com/databricks/databricks-sdk-go/databricks/useragent"
)

type Config config.Config

// Must panics if error is not nil. It's intended to be used with
// [databricks.NewWorkspaceClient] for variable initializations
func Must[T any](c T, err error) T {
	if err != nil {
		panic(err)
	}
	return c
}

// Version of this SDK
func Version() string {
	return internal.Version
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
