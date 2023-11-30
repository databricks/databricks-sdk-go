package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultEnvironmentIsReturned(t *testing.T) {
	c := &Config{}
	assert.Equal(t, ".cloud.databricks.com", c.Environment().DnsZone)
}

func TestOverriddenEnvironmentIsReturned(t *testing.T) {
	c := &Config{
		Host: "something.else",
		DatabricksEnvironment: &DatabricksEnvironment{
			Cloud:   CloudAzure,
			DnsZone: "holla",
		},
	}
	assert.Equal(t, "holla", c.Environment().DnsZone)
}
