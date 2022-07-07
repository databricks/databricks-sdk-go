package implicit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigAttributes(t *testing.T) {
	num := len(ConfigAttributes)
	assert.Equal(t, 1, num)
}
