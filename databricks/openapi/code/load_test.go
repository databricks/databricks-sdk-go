package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombined(t *testing.T) {
	batch, err := NewFromFile("/private/tmp/processed-databricks-workspace-all.json")
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}

func TestBasic(t *testing.T) {
	batch, err := NewFromFile("../testdata/spec.json")
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}
