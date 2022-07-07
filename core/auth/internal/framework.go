package internal

import (
	"strings"
)

type Host string

// IsAzure returns true if client is configured for Azure Databricks
func (h Host) IsAzure() bool {
	return strings.Contains(string(h), ".azuredatabricks.net")
}

// IsGcp returns true if client is configured for GCP
func (h Host) IsGcp() bool {
	return strings.Contains(string(h), ".gcp.databricks.com")
}

// IsAws returns true if client is configured for AWS
func (h Host) IsAws() bool {
	return !h.IsAzure() && !h.IsGcp()
}

// IsAws returns true if client is configured for Accounts API
func (h Host) IsAccountsClient() bool {
	return strings.HasPrefix(string(h), "https://accounts.")
}
