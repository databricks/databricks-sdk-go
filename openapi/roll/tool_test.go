package roll

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/openapi/code"
	"github.com/databricks/databricks-sdk-go/openapi/render"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadsFolder(t *testing.T) {
	// ../../internal is the folder with integration tests
	s, err := NewSuite("../../internal")
	require.NoError(t, err)

	methods := s.Methods()
	assert.True(t, len(methods) > 1)

	examples := s.ServicesExamples()
	assert.True(t, len(examples) > 1)

	for _, v := range examples {
		for _, sa := range v.Samples {
			for _, ca := range sa.Calls {
				// verify no panic
				_ = ca.Original()

				// verify no panic
				_ = ca.Request()
			}
		}
	}
}

func TestOptimize(t *testing.T) {
	t.Skip()
	home, _ := os.UserHomeDir()
	batch, err := code.NewFromFile(filepath.Join(home,
		"universe/bazel-bin/openapi/all-internal.json"))
	require.NoError(t, err)

	s, err := NewSuite("../../internal")
	require.NoError(t, err)

	err = s.OptimizeWithApiSpec(batch)
	require.NoError(t, err)
}

func TestRegenerateExamples(t *testing.T) {
	t.Skip() // temporary measure
	s, err := NewSuite("../../internal")
	require.NoError(t, err)

	target := "../.."
	pass := render.NewPass(target, s.ServicesExamples(), map[string]string{
		".codegen/examples_test.go.tmpl": "service/{{.Package}}/{{.SnakeName}}_usage_test.go",
	})
	err = pass.Run()
	assert.NoError(t, err)

	err = render.Fomratter(target, pass.Filenames, "go fmt ./... && go run golang.org/x/tools/cmd/goimports@latest -w $FILENAMES")
	assert.NoError(t, err)
}

func TestRegeneratePythonExamples(t *testing.T) {
	t.Skip()
	s, err := NewSuite("../../internal")
	require.NoError(t, err)

	target := "../../../databricks-sdk-py"
	pass := render.NewPass(target, s.Samples(), map[string]string{
		".codegen/example.py.tmpl": "examples/{{.Service.SnakeName}}/{{.Method.SnakeName}}_{{.SnakeName}}.py",
	})
	err = pass.Run()
	assert.NoError(t, err)

	err = render.Fomratter(target, pass.Filenames, "yapf -pri $FILENAMES && autoflake -i $FILENAMES && isort $FILENAMES")
	assert.NoError(t, err)
}
