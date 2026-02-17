package config

import (
	"testing"
)

func TestHostnameMatchesSpogUrlSuffix(t *testing.T) {
	testCases := []struct {
		name     string
		hostname string
		suffix   string
		expected bool
	}{
		// Valid SPOG patterns - single subdomain
		{
			name:     "valid databricks.com",
			hostname: "my-company.databricks.com",
			suffix:   ".databricks.com",
			expected: true,
		},
		{
			name:     "valid azuredatabricks.net",
			hostname: "my-company.azuredatabricks.net",
			suffix:   ".azuredatabricks.net",
			expected: true,
		},
		// Invalid - multiple subdomains
		{
			name:     "invalid - subdomain in databricks.com",
			hostname: "my-company.subdomain.databricks.com",
			suffix:   ".databricks.com",
			expected: false,
		},
		{
			name:     "invalid - subdomain in azuredatabricks.net",
			hostname: "my-company.subdomain.azuredatabricks.net",
			suffix:   ".azuredatabricks.net",
			expected: false,
		},
		{
			name:     "invalid - cloud subdomain",
			hostname: "dbc-12345.cloud.databricks.com",
			suffix:   ".databricks.com",
			expected: false,
		},
		// Invalid - doesn't end with suffix
		{
			name:     "invalid - wrong suffix",
			hostname: "my-company.staging.databricks.com",
			suffix:   ".databricks.com",
			expected: false,
		},
		{
			name:     "invalid - no match",
			hostname: "my-company.example.com",
			suffix:   ".databricks.com",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := hostnameMatchesSpogUrlSuffix(tc.hostname, tc.suffix)
			if result != tc.expected {
				t.Errorf("hostnameMatchesSpogUrlSuffix(%q, %q) = %v, want %v", tc.hostname, tc.suffix, result, tc.expected)
			}
		})
	}
}

func TestMatchesSpogUrlPattern(t *testing.T) {
	testCases := []struct {
		name     string
		hostname string
		expected bool
	}{
		// Production patterns - valid
		{
			name:     "valid - databricks.com",
			hostname: "my-company.databricks.com",
			expected: true,
		},
		{
			name:     "valid - azuredatabricks.net",
			hostname: "my-company.azuredatabricks.net",
			expected: true,
		},
		{
			name:     "valid - with dashes",
			hostname: "my-company-name.databricks.com",
			expected: true,
		},
		// Production patterns - invalid (subdomains)
		{
			name:     "invalid - subdomain databricks.com",
			hostname: "my-company.subdomain.databricks.com",
			expected: false,
		},
		{
			name:     "invalid - subdomain azuredatabricks.net",
			hostname: "my-company.subdomain.azuredatabricks.net",
			expected: false,
		},
		{
			name:     "invalid - cloud subdomain",
			hostname: "dbc-12345.cloud.databricks.com",
			expected: false,
		},
		// Staging patterns (not matched in production)
		{
			name:     "staging - not matched",
			hostname: "my-company.staging.databricks.com",
			expected: false,
		},
		{
			name:     "staging cloud - not matched",
			hostname: "my-company-spog.staging.cloud.databricks.com",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := matchesSpogUrlPattern(tc.hostname)
			if result != tc.expected {
				t.Errorf("matchesSpogUrlPattern(%q) = %v, want %v", tc.hostname, result, tc.expected)
			}
		})
	}
}

func TestExtractHostname(t *testing.T) {
	testCases := []struct {
		name     string
		host     string
		expected string
	}{
		{
			name:     "full URL with https",
			host:     "https://my-company.databricks.com",
			expected: "my-company.databricks.com",
		},
		{
			name:     "full URL with http",
			host:     "http://my-company.databricks.com",
			expected: "my-company.databricks.com",
		},
		{
			name:     "bare hostname",
			host:     "my-company.databricks.com",
			expected: "my-company.databricks.com",
		},
		{
			name:     "URL with port",
			host:     "https://my-company.databricks.com:443",
			expected: "my-company.databricks.com",
		},
		{
			name:     "URL with path",
			host:     "https://my-company.databricks.com/path/to/resource",
			expected: "my-company.databricks.com",
		},
		{
			name:     "empty string",
			host:     "",
			expected: "",
		},
		{
			name:     "invalid URL",
			host:     "://invalid",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := extractHostname(tc.host)
			if result != tc.expected {
				t.Errorf("extractHostname(%q) = %q, want %q", tc.host, result, tc.expected)
			}
		})
	}
}

func TestIsUnifiedHost(t *testing.T) {
	testCases := []struct {
		name     string
		host     string
		expected bool
	}{
		// Valid unified hosts - databricks.com
		{
			name:     "valid - company.databricks.com",
			host:     "company.databricks.com",
			expected: true,
		},
		{
			name:     "valid - my-company.databricks.com",
			host:     "my-company.databricks.com",
			expected: true,
		},
		{
			name:     "valid - with https scheme",
			host:     "https://my-company.databricks.com",
			expected: true,
		},
		{
			name:     "valid - with http scheme",
			host:     "http://my-company.databricks.com",
			expected: true,
		},
		{
			name:     "valid - with port",
			host:     "https://my-company.databricks.com:443",
			expected: true,
		},
		{
			name:     "valid - with path",
			host:     "https://my-company.databricks.com/api/2.0",
			expected: true,
		},
		// Valid unified hosts - azuredatabricks.net
		{
			name:     "valid - my-company.azuredatabricks.net",
			host:     "my-company.azuredatabricks.net",
			expected: true,
		},
		{
			name:     "valid - azure with https",
			host:     "https://my-company.azuredatabricks.net",
			expected: true,
		},
		// Invalid - subdomains (regional URLs)
		{
			name:     "invalid - subdomain databricks.com",
			host:     "my-company.subdomain.databricks.com",
			expected: false,
		},
		{
			name:     "invalid - cloud subdomain",
			host:     "dbc-12345.cloud.databricks.com",
			expected: false,
		},
		{
			name:     "invalid - regional azure",
			host:     "adb-123456.12.azuredatabricks.net",
			expected: false,
		},
		{
			name:     "invalid - subdomain azuredatabricks.net",
			host:     "my-company.subdomain.azuredatabricks.net",
			expected: false,
		},
		// Invalid - staging/dev patterns
		{
			name:     "invalid - staging",
			host:     "my-company.staging.databricks.com",
			expected: false,
		},
		{
			name:     "invalid - dev",
			host:     "my-company.dev.databricks.com",
			expected: false,
		},
		{
			name:     "invalid - staging cloud",
			host:     "my-company-spog.staging.cloud.databricks.com",
			expected: false,
		},
		// Invalid - other domains
		{
			name:     "invalid - gcp domain",
			host:     "12345.gcp.databricks.com",
			expected: false,
		},
		{
			name:     "invalid - different TLD",
			host:     "my-company.databricks.io",
			expected: false,
		},
		{
			name:     "invalid - completely different domain",
			host:     "example.com",
			expected: false,
		},
		// Edge cases
		{
			name:     "invalid - empty string",
			host:     "",
			expected: false,
		},
		{
			name:     "invalid - just domain",
			host:     "databricks.com",
			expected: false,
		},
		{
			name:     "invalid - just azure domain",
			host:     "azuredatabricks.net",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsUnifiedHost(tc.host)
			if result != tc.expected {
				t.Errorf("IsUnifiedHost(%q) = %v, want %v", tc.host, result, tc.expected)
			}
		})
	}
}

func TestIsUnifiedHostComprehensiveCases(t *testing.T) {
	testCases := []struct {
		name     string
		host     string
		expected bool
		comment  string
	}{
		// Production - databricks.com patterns
		{
			name:     "prod - simple databricks.com",
			host:     "my-company.databricks.com",
			expected: true,
			comment:  "matches SPOG pattern",
		},
		{
			name:     "prod - subdomain databricks.com",
			host:     "my-company.subdomain.databricks.com",
			expected: false,
			comment:  "has subdomain, not SPOG",
		},
		// Production - azuredatabricks.net patterns
		{
			name:     "prod - simple azuredatabricks.net",
			host:     "my-company.azuredatabricks.net",
			expected: true,
			comment:  "matches SPOG pattern",
		},
		{
			name:     "prod - subdomain azuredatabricks.net",
			host:     "my-company.subdomain.azuredatabricks.net",
			expected: false,
			comment:  "has subdomain, not SPOG",
		},
		// Staging patterns (should not match in production-focused SDK)
		{
			name:     "staging - .staging.databricks.com",
			host:     "my-company.staging.databricks.com",
			expected: false,
			comment:  "staging pattern not matched",
		},
		{
			name:     "staging - spog suffix in staging cloud",
			host:     "my-company-spog.staging.cloud.databricks.com",
			expected: false,
			comment:  "staging -spog suffix not matched",
		},
		{
			name:     "staging - spog subdomain in staging cloud",
			host:     "my-company-spog.subdomain.staging.cloud.databricks.com",
			expected: false,
			comment:  "extra subdomains in staging",
		},
		// Development patterns (should not match in production-focused SDK)
		{
			name:     "dev - spog suffix in dev",
			host:     "my-company-spog.dev.databricks.com",
			expected: false,
			comment:  "dev -spog suffix not matched",
		},
		{
			name:     "dev - spog subdomain in dev",
			host:     "my-company-spog.subdomain.dev.databricks.com",
			expected: false,
			comment:  "extra subdomains in dev",
		},
		// Regional/workspace URLs (should not match)
		{
			name:     "regional - cloud subdomain",
			host:     "dbc-12345.cloud.databricks.com",
			expected: false,
			comment:  "regional URL with cloud subdomain",
		},
		{
			name:     "regional - azure workspace",
			host:     "adb-123456.12.azuredatabricks.net",
			expected: false,
			comment:  "azure regional workspace URL",
		},
		// Real-world examples
		{
			name:     "real - typical customer unified host",
			host:     "abc-corp.databricks.com",
			expected: true,
			comment:  "typical customer unified host",
		},
		{
			name:     "real - aws workspace",
			host:     "dbc-1234abcd-5678.cloud.databricks.com",
			expected: false,
			comment:  "typical AWS workspace URL",
		},
		{
			name:     "real - azure workspace",
			host:     "adb-1234567890123456.12.azuredatabricks.net",
			expected: false,
			comment:  "typical Azure workspace URL",
		},
		{
			name:     "real - gcp workspace",
			host:     "1234567890123456.7.gcp.databricks.com",
			expected: false,
			comment:  "typical GCP workspace URL",
		},
		{
			name:     "real - account console",
			host:     "accounts.cloud.databricks.com",
			expected: false,
			comment:  "account console is not a unified host",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsUnifiedHost(tc.host)
			if result != tc.expected {
				t.Errorf("IsUnifiedHost(%q) = %v, want %v (%s)", tc.host, result, tc.expected, tc.comment)
			}
		})
	}
}
