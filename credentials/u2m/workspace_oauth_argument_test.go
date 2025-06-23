package u2m

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateHost(t *testing.T) {
	tests := []struct {
		host string
		want string
	}{
		// Valid hosts
		{"https://some-host.com", ""},
		{"http://127.0.0.1", ""},
		{"http://127.0.0.1:5656", ""},

		// Invalid hosts
		{"http://some-host.com", "host must start with 'https://': http://some-host.com"},
		{"https://some-host.com/", "host must not have a trailing slash: https://some-host.com/"},
	}

	for _, test := range tests {
		err := validateHost(test.host)
		if test.want == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.want)
		}

		_, err = NewBasicWorkspaceOAuthArgument(test.host)
		if test.want == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.want)
		}

		_, err = NewBasicAccountOAuthArgument(test.host, "123")
		if test.want == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, test.want)
		}
	}
}
