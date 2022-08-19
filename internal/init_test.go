package internal

import (
	"os"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
)

func init() {
	databricks.WithProduct("integration-tests", databricks.Version())
}

// GetEnvOrSkipTest proceeds with test only with that env variable
func GetEnvOrSkipTest(t *testing.T, name string) string {
	value := os.Getenv(name)
	if value == "" {
		t.Skipf("Environment variable %s is missing", name)
	}
	return value
}
