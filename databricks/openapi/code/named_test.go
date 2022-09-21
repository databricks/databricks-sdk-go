package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNonConflictingCamelNames(t *testing.T) {
	n := Named{Name: "Import"}
	assert.True(t, n.IsNameReserved())
}
