package beep

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/openapi/render"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadsFolder(t *testing.T) {
	t.Skip()
	s, err := NewSuite("../../internal")
	require.NoError(t, err)

	methods := s.Methods()
	assert.True(t, len(methods) > 1)

	samples := s.Samples()
	assert.True(t, len(samples) > 1)

	target := "../.."
	pass := render.NewPass(target, s.ServicesExamples(), map[string]string{
		".codegen/examples_test.go.tmpl": "service/{{.Package}}/{{.SnakeName}}_usage_test.go",
	})
	err = pass.Run()
	assert.NoError(t, err)

	err = render.Fomratter(target, pass.Filenames, "go fmt ./... && go run golang.org/x/tools/cmd/goimports@latest -w $FILENAMES")
	assert.NoError(t, err)
}
