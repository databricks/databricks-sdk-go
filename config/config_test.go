package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAccountClient_AwsAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts.cloud.databricks.com",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.True(t, c.IsAccountClient())
}

func TestIsAccountClient_AwsDodAccount(t *testing.T) {
	c := &Config{
		Host:      "https://accounts-dod.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.True(t, c.IsAccountClient())
}

func TestIsAccountClient_AwsWorkspace(t *testing.T) {
	c := &Config{
		Host:      "https://my-workspace.cloud.databricks.us",
		AccountID: "123e4567-e89b-12d3-a456-426614174000",
	}
	assert.False(t, c.IsAccountClient())
}
