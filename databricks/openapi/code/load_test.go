package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	batch, err := NewFromFile("/tmp/processed-scim.json")
	assert.NoError(t, err)

	types := batch.Packages[2].Types()
	assert.Greater(t, len(types), 3)

	assert.Len(t, batch.Packages, 3)
}

func TestBasic(t *testing.T) {
	batch, err := NewFromFile("../testdata/spec.json")
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}
