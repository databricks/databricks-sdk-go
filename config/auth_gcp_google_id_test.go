package config

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
)

func TestGoogleIdHappyFlow(t *testing.T) {
	assertHeaders(t, &Config{
		GoogleServiceAccount: "abc",
		Host:                 "bcd",
		DatabricksEnvironments: []DatabricksEnvironment{
			{
				dnsZone: "bcd",
				Cloud:   CloudGCP,
			},
		},
		HTTPTransport: fixtures.MappingTransport{
			"POST /v1/projects/-/serviceAccounts/abc:generateIdToken": {
				ExpectedRequest: map[string]any{
					"audience":     "https://bcd",
					"includeEmail": true,
				},
				Response: `{"token": "cde"}`,
			},
		},
	}, map[string]string{
		"Authorization": "Bearer cde",
	})
}
