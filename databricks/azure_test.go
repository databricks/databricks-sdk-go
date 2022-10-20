package databricks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAzureLoginAppID(t *testing.T) {
	var cfg Config

	// It's not set
	cfg = Config{}
	assert.Equal(t, azureDatabricksLoginAppID, cfg.getAzureLoginAppID())

	// It's set
	cfg = Config{AzureLoginAppID: "foobar"}
	assert.Equal(t, "foobar", cfg.getAzureLoginAppID())
}
