package useragent

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/databricks/sdk-go/databricks/internal"
	"github.com/stretchr/testify/assert"
)

func TestFromContext_Default(t *testing.T) {
	userAgent := FromContext(context.Background())
	// tests may be developed and run on different versions of different things
	expectedFormat := "unknown/0.0.0 databricks-sdk-go/%s go/%s os/%s"
	expected := fmt.Sprintf(expectedFormat, internal.Version,
		strings.TrimPrefix(runtime.Version(), "go"), runtime.GOOS)
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
	expected := fmt.Sprintf(expectedFormat, internal.Version,
		strings.TrimPrefix(runtime.Version(), "go"), runtime.GOOS)
	assert.Equal(t, expected, userAgent)
}
