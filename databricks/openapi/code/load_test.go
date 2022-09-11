package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	batch, err := NewFromFile("../testdata/spec.json", []string{})
	assert.NoError(t, err)

	assert.Len(t, batch.Packages, 1)
}
