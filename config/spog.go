package config

import (
	"net/url"
	"strings"
)

// IsUnifiedHost returns true if the given host is a unified host that supports
// both workspace-level and account-level APIs. This matches the SPOG domain checker
// logic from the platform.
//
// A hostname is considered a unified host if it matches one of the SPOG URL patterns:
// - *.databricks.com (without subdomains, e.g., company.databricks.com)
// - *.azuredatabricks.net (without subdomains, e.g., company.azuredatabricks.net)
func IsUnifiedHost(host string) bool {
	if host == "" {
		return false
	}

	hostname := extractHostname(host)
	if hostname == "" {
		return false
	}

	return matchesSpogUrlPattern(hostname)
}

// extractHostname parses a URL or hostname string and returns the hostname component.
// It handles both full URLs (with scheme) and bare hostnames.
func extractHostname(host string) string {
	// Parse the URL to extract just the hostname
	parsedHost, err := url.Parse(host)
	if err != nil {
		return ""
	}

	// If no host was parsed, assume the scheme wasn't included
	hostname := parsedHost.Hostname()
	if hostname == "" {
		parsedHost, err = url.Parse("https://" + host)
		if err != nil {
			return ""
		}
		hostname = parsedHost.Hostname()
	}

	return hostname
}

// matchesSpogUrlPattern checks if the hostname matches any SPOG URL pattern.
// For production (which is what the SDK primarily targets), this includes:
// - *.databricks.com
// - *.azuredatabricks.net
func matchesSpogUrlPattern(hostname string) bool {
	// SPOG URL patterns for production
	spogUrlSuffixes := []string{".databricks.com", ".azuredatabricks.net"}

	// Check if hostname matches any SPOG URL pattern
	for _, suffix := range spogUrlSuffixes {
		if hostnameMatchesSpogUrlSuffix(hostname, suffix) {
			return true
		}
	}

	return false
}

// hostnameMatchesSpogUrlSuffix checks if hostname ends with the given suffix
// and that the prefix (before the suffix) doesn't contain any dots.
// This ensures that only single-level subdomains are matched.
//
// Examples:
//   - "company.databricks.com" with suffix ".databricks.com" -> true (prefix "company" has no dots)
//   - "dbc-12345.cloud.databricks.com" with suffix ".databricks.com" -> false (prefix "dbc-12345.cloud" has dots)
func hostnameMatchesSpogUrlSuffix(hostname string, suffix string) bool {
	if !strings.HasSuffix(hostname, suffix) {
		return false
	}

	// Extract the prefix before the suffix
	prefix := strings.TrimSuffix(hostname, suffix)

	// Check that the prefix doesn't contain any dots (no subdomains)
	return !strings.Contains(prefix, ".")
}
