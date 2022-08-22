package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	batch, err := NewFromFile("/tmp/processed-jobs-2.1.json")
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}

func TestBasic(t *testing.T) {
	batch, err := NewFromFile("../testdata/spec.json")
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}
