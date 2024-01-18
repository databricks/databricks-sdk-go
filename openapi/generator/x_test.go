package generator

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"github.com/stretchr/testify/require"
)

func run() error {
	ctx := context.Background()
	home, _ := os.UserHomeDir()
	spec, err := code.NewFromFile(ctx, filepath.Join(home,
		"Downloads/RenderWidgetSpec.openapi.json"))
	if err != nil {
		return fmt.Errorf("spec: %w", err)
	}
	gen, err := NewGenerator(filepath.Join(home,
		"git/labs/ucx/src/databricks/labs/ucx/framework/lakeview"))
	if err != nil {
		return fmt.Errorf("config: %w", err)
	}
	return gen.Apply(ctx, spec, nil)
}

func TestX(t *testing.T) {
	// t.SkipNow()
	err := run()
	require.NoError(t, err)
}
