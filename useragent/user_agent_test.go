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

	// upstream and upstream-version only appear when both are set
	assert.NotContains(t, userAgent, "upstream/")
	os.Setenv("DATABRICKS_SDK_UPSTREAM", "my-upstream-sdk")
	defer os.Unsetenv("DATABRICKS_SDK_UPSTREAM")
	userAgent3 := FromContext(ctx)
	assert.NotContains(t, userAgent3, "upstream/")
	assert.NotContains(t, userAgent3, "upstream-version/")

	os.Setenv("DATABRICKS_SDK_UPSTREAM_VERSION", "1.2.3")
	defer os.Unsetenv("DATABRICKS_SDK_UPSTREAM_VERSION")
	userAgent4 := FromContext(context.Background())
	assert.Contains(t, userAgent4, "upstream/my-upstream-sdk upstream-version/1.2.3")
}

func TestDefaultsAreValid(t *testing.T) {
	WithProduct(productName, productVersion)
}

func TestUserAgentValidate(t *testing.T) {
	assert.EqualError(t, validate("foobar!", ""), "expected user agent key to be alphanumeric: \"foobar!\"")
	assert.EqualError(t, validate("foo", "invalid!"), "expected user agent value for key \"foo\" to be alphanumeric: \"invalid!\"")
	assert.EqualError(t, validate("foo", "whatever#!@"), "expected user agent value for key \"foo\" to be alphanumeric: \"whatever#!@\"")

	assert.NoError(t, validate("foo", "7.3"))
	assert.NoError(t, validate("foo", "client.7"))
	assert.NoError(t, validate("foo", "1.1.1"))
}

func TestSanitize(t *testing.T) {
	// Valid values
	assert.True(t, isValid("foo"))
	assert.Equal(t, "foo", Sanitize("foo"))

	assert.True(t, isValid("FOO"))
	assert.Equal(t, "FOO", Sanitize("FOO"))

	assert.True(t, isValid("FOO123"))
	assert.Equal(t, "FOO123", Sanitize("FOO123"))

	assert.True(t, isValid("foo_bar"))
	assert.Equal(t, "foo_bar", Sanitize("foo_bar"))

	assert.True(t, isValid("foo-bar"))
	assert.Equal(t, "foo-bar", Sanitize("foo-bar"))

	assert.True(t, isValid("foo.bar"))
	assert.Equal(t, "foo.bar", Sanitize("foo.bar"))

	assert.True(t, isValid("1.2.3"))
	assert.Equal(t, "1.2.3", Sanitize("1.2.3"))

	assert.True(t, isValid("client.0"))
	assert.Equal(t, "client.0", Sanitize("client.0"))

	// Invalid Values, being sanitized correctly.
	assert.False(t, isValid("1@2#3?4,5/6!7 8 "))
	assert.Equal(t, "1-2-3-4-5-6-7-8-", Sanitize("1@2#3?4,5/6!7 8 "))

	assert.False(t, isValid("foo bar"))
	assert.Equal(t, "foo-bar", Sanitize("foo bar"))

	assert.False(t, isValid("foo/bar"))
	assert.Equal(t, "foo-bar", Sanitize("foo/bar"))

	assert.False(t, isValid("foo:)bar"))
	assert.Equal(t, "foo--bar", Sanitize("foo:)bar"))

	assert.False(t, isValid("foo😊bar"))
	assert.Equal(t, "foo-bar", Sanitize("foo😊bar"))
}
