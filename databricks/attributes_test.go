package databricks

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTestSetAttribute(t *testing.T) {
	config := Config{}
	attrs := discoverAttributesOf(&config, "config", map[string]string{})
	rv := reflect.ValueOf(&config)
	attrs[0].Set(rv, "http://localhost")
	assert.Equal(t, "http://localhost", config.Host)
}

func TestTestSetAttributeOnPat(t *testing.T) {
	// TODO: this test may go away
	pat := PatCredentials{}
	attrs := discoverAttributesOf(pat, "pat", map[string]string{})
	rv := reflect.ValueOf(&pat)
	attrs[0].Set(rv, "abc12345678")
	assert.Equal(t, "abc12345678", pat.Token)
}
