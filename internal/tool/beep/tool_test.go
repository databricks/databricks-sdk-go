package beep

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadsBilling(t *testing.T) {
	s, err := NewSuite("../../billing_test.go")
	require.NoError(t, err)

	svcs := s.ImprotedServices()
	assert.True(t, len(svcs) > 0)
}

func TestLoadsClusters(t *testing.T) {
	s, err := NewSuite("../../clusters_test.go")
	require.NoError(t, err)

	samples := s.Sample("clusters", "pin")
	assert.True(t, len(samples) == 1)
}
