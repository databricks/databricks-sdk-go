package useragent

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/version"
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
	expected := fmt.Sprintf(expectedFormat, version.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected, userAgent)
}

func TestFromContext_Custom(t *testing.T) {
	WithProduct("unit-tests", "0.0.1")
	WithUserAgentExtra("pulumi", "3.8.4")
	WithUserAgentExtra("terraform", "1.3.4")
	WithUserAgentExtra("terraform", "1.3.5")
	WithUserAgentExtra("terraform", "1.3.6")

	ctx := InContext(context.Background(), "a", "b")
	ctx = InContext(ctx, "a", "d")
	ctx = InContext(ctx, "a", "e")

	ctx2 := InContext(ctx, "a", "f")
	ctx2 = InContext(ctx2, "foo", "bar")

	userAgent := FromContext(ctx)
	userAgent2 := FromContext(ctx2)

	// tests may be developed and run on different versions of different things
	expectedFormat := "unit-tests/0.0.1 databricks-sdk-go/%s go/%s os/%s pulumi/3.8.4 terraform/1.3.6 a/e"
	expected := fmt.Sprintf(expectedFormat, version.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected, userAgent)

	expectedFormat2 := "unit-tests/0.0.1 databricks-sdk-go/%s go/%s os/%s pulumi/3.8.4 terraform/1.3.6 a/f foo/bar"
	expected2 := fmt.Sprintf(expectedFormat2, version.Version, goVersion(), runtime.GOOS)
	assert.Equal(t, expected2, userAgent2)
}

func TestDefaultsAreValid(t *testing.T) {
	WithProduct(productName, productVersion)
}

func TestUpstreamUserAgent(t *testing.T) {
	// Set the environment variables to test the upstream user agent
	// and check that it is included in the user agent string.
	os.Setenv("DATABRICKS_SDK_UPSTREAM", "my-upstream-sdk")
	os.Setenv("DATABRICKS_SDK_UPSTREAM_VERSION", "1.2.3")
	userAgent := FromContext(context.Background())
	assert.Contains(t, userAgent, "upstream/my-upstream-sdk upstream-version/1.2.3")
	os.Unsetenv("DATABRICKS_SDK_UPSTREAM")
	os.Unsetenv("DATABRICKS_SDK_UPSTREAM_VERSION")
}
