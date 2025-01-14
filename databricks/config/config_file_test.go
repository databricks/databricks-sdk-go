package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigFileLoad(t *testing.T) {
	f, err := LoadFile("testdata/.databrickscfg")
	require.NoError(t, err)
	assert.NotNil(t, f)

	for _, name := range []string{
		"password-with-double-quotes",
		"password-with-single-quotes",
		"password-without-quotes",
	} {
		section := f.Section(name)
		require.NotNil(t, section)
		assert.Equal(t, "%Y#X$Z", section.Key("password").String())
	}
}
