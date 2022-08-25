package useragent

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/internal"
	"github.com/stretchr/testify/assert"
)

func TestGoVersion(t *testing.T) {
	// The function `goVersion` returns a semver canonicalized Go version,
	// but `runtime.Version` may or may not include a patch version.
	// Correct for that in the test below.
	parts := strings.SplitN(strings.TrimPrefix(runtime.Version(), "go"), ".", 3)
	if len(parts) == 2 {
		parts = append(parts, "0")
	}
	assert.Equal(t, goVersion(), strings.Join(parts, "."))
}

func TestFromContext_Default(t *testing.T) {
	userAgent := FromContext(context.Background())
	// tests may be developed and run on different versions of different things
	expectedFormat := "unknown/0.0.0 databricks-sdk-go/%s go/%s os/%s"
	expected := fmt.Sprintf(expectedFormat, internal.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected, userAgent)
}

func TestFromContext_Custom(t *testing.T) {
	WithProduct("unit-tests", "0.0.1")
	WithUserAgentExtra("pulumi", "3.8.4")
	ctx := InContext(context.Background(), "a", "b")
	ctx = InContext(ctx, "terraform", "1.2.5")
	userAgent := FromContext(ctx)

	// tests may be developed and run on different versions of different things
	expectedFormat := "unit-tests/0.0.1 databricks-sdk-go/%s go/%s os/%s pulumi/3.8.4 a/b terraform/1.2.5"
	expected := fmt.Sprintf(expectedFormat, internal.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected, userAgent)
}
