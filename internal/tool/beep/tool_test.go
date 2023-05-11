package beep

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoads(t *testing.T) {
	s, err := NewSuite("../../billing_test.go")
	require.NoError(t, err)

	svcs := s.ImprotedServices()
	assert.True(t, len(svcs) > 0)
}
