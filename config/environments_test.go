package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverriddenEnvironmentIsReturnedInTesting(t *testing.T) {
	c := &Config{
		Host: "something.else",
		DatabricksEnvironments: []DatabricksEnvironment{
			{Cloud: CloudAzure, DnsZone: "holla"},
		},
	}
	c.WithTesting()
	assert.Equal(t, "holla", c.Environment().DnsZone)
}

func TestOverriddenEnvironmentOverrides(t *testing.T) {
	c := &Config{
		Host: "my.workspace.holla",
		DatabricksEnvironments: []DatabricksEnvironment{
			{Cloud: CloudAzure, DnsZone: "holla"},
		},
	}
	assert.Equal(t, "holla", c.Environment().DnsZone)
}

func TestEnvironmentFallback(t *testing.T) {
	c := &Config{
		Host: "a.dev.databricks.com",
		DatabricksEnvironments: []DatabricksEnvironment{
			{Cloud: CloudAzure, DnsZone: "holla"},
		},
	}
	assert.Equal(t, ".dev.databricks.com", c.Environment().DnsZone)
}
