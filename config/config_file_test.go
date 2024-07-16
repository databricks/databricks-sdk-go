package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigFileLoad(t *testing.T) {
	f, err := LoadFile("testdata/.databrickscfg")
	assert.NoError(t, err)
	assert.NotNil(t, f)

	for _, name := range []string{
		"password-with-double-quotes",
		"password-with-single-quotes",
		"password-without-quotes",
	} {
		section := f.Section(name)
		assert.NotNil(t, section)
		assert.Equal(t, "%Y#X$Z", section.Key("password").String())
	}
}
