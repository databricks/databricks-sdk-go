package roll

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/openapi/render"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadsFolder(t *testing.T) {
	s, err := NewSuite("../../internal")
	require.NoError(t, err)

	methods := s.Methods()
	assert.True(t, len(methods) > 1)

	samples := s.Samples()
	assert.True(t, len(samples) > 1)
}

func TestRegenerateExamples(t *testing.T) {
	t.Skip()
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
