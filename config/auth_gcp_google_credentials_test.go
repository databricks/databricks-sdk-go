package config

import (
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
)

func TestGoogleCredsHappyFlow(t *testing.T) {
	assertHeaders(t, &Config{
		GoogleCredentials: "abc",
		Host:              "bcd",
		DatabricksEnvironments: []DatabricksEnvironment{
			{
				dnsZone: "bcd",
				Cloud:   CloudGCP,
			},
		},
		HTTPTransport: fixtures.MappingTransport{
			//..
		},
	}, map[string]string{
		"Authorization": "Bearer cde",
	})
}
