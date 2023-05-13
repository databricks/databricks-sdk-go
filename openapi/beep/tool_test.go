package beep

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
	assert.True(t, len(methods) == 1)

	samples := s.Samples()
	assert.True(t, len(samples) == 1)
}

func TestLoadsClusters(t *testing.T) {
	s, err := NewSuite("../../internal/clusters_test.go")
	require.NoError(t, err)

	methods := s.Methods()
	assert.True(t, len(methods) == 1)

	samples := s.Samples()
	assert.True(t, len(samples) == 1)

	pass := render.NewPass("/Users/serge.smertin/git/databricks/databricks-sdk-go", []*suite{s}, map[string]string{
		".codegen/examples_test.go.tmpl": "examples/examples_xxx.go",
	})
	err = pass.Run()
	assert.NoError(t, err)
}
