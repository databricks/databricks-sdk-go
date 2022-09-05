package openapi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromJson(t *testing.T) {
	f, err := os.Open("testdata/spec.json")
	assert.NoError(t, err)
	spec, err := NewFromReader(f)
	assert.NoError(t, err)
	assert.Equal(t, "Jobs", spec.Tags[0].Name)
}
