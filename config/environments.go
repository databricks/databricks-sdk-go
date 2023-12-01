package config

import (
	"fmt"
	"strings"
)

type Cloud string

const (
	CloudAWS   Cloud = "AWS"
	CloudAzure Cloud = "Azure"
	CloudGCP   Cloud = "GCP"
)

type DatabricksEnvironment struct {
	Cloud              Cloud
	DnsZone            string
	AzureApplicationID string
	AzureEnvironment   *azureEnvironment
}

func (de DatabricksEnvironment) DeploymentURL(name string) string {
	return fmt.Sprintf("https://%s%s", name, de.DnsZone)
}

func (de DatabricksEnvironment) AzureServiceManagementEndpoint() string {
	if de.AzureEnvironment == nil {
		return ""
	}
	return de.AzureEnvironment.ServiceManagementEndpoint
}

func (de DatabricksEnvironment) AzureResourceManagerEndpoint() string {
	if de.AzureEnvironment == nil {
		return ""
	}
	return de.AzureEnvironment.ResourceManagerEndpoint
}

func (de DatabricksEnvironment) AzureActiveDirectoryEndpoint() string {
	if de.AzureEnvironment == nil {
		return ""
	}
	return de.AzureEnvironment.ActiveDirectoryEndpoint
}

// we default to AWS Prod environment since this case will be a hit for PVC
var defaultEnvironment = DatabricksEnvironment{
	Cloud:   CloudAWS,
	DnsZone: ".cloud.databricks.com",
}

var envs = []DatabricksEnvironment{
	{Cloud: CloudAWS, DnsZone: ".dev.databricks.com"},
	{Cloud: CloudAWS, DnsZone: ".staging.cloud.databricks.com"},
	{Cloud: CloudAWS, DnsZone: ".cloud.databricks.us"},
	defaultEnvironment,

	{Cloud: CloudAzure, DnsZone: ".dev.azuredatabricks.net", AzureApplicationID: "62a912ac-b58e-4c1d-89ea-b2dbfc7358fc", AzureEnvironment: &AzurePublicCloud},
	{Cloud: CloudAzure, DnsZone: ".staging.azuredatabricks.net", AzureApplicationID: "4a67d088-db5c-48f1-9ff2-0aace800ae68", AzureEnvironment: &AzurePublicCloud},
	{Cloud: CloudAzure, DnsZone: ".azuredatabricks.net", AzureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", AzureEnvironment: &AzurePublicCloud},
	{Cloud: CloudAzure, DnsZone: ".databricks.azure.us", AzureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", AzureEnvironment: &AzureUsGovernmentCloud},
	{Cloud: CloudAzure, DnsZone: ".databricks.azure.cn", AzureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", AzureEnvironment: &AzureChinaCloud},

	{Cloud: CloudGCP, DnsZone: ".dev.gcp.databricks.com"},
	{Cloud: CloudGCP, DnsZone: ".staging.gcp.databricks.com"},
	{Cloud: CloudGCP, DnsZone: ".gcp.databricks.com"},
}

func (c *Config) Environment() DatabricksEnvironment {
	// Use the provided environment if specified. Tests may configure the client with different hostnames,
	// like localhost, which are not resolvable to a known environment, while needing to mock a specific environment.
	if c.DatabricksEnvironment != nil {
		return *c.DatabricksEnvironment
	}
	if c.Host == "" && c.AzureResourceID != "" {
		// azure resource ID can also be used in lieu of host by some
		// of the clients, like Terraform. However, in this case, the workspace
		// is assumed to be a production workspace.
		azureEnv := strings.ToUpper(c.AzureEnvironment)
		if azureEnv == "" {
			azureEnv = "PUBLIC"
		}
		for _, v := range envs {
			if v.Cloud != CloudAzure {
				continue
			}
			if v.AzureEnvironment.Name != azureEnv {
				continue
			}
			if strings.HasPrefix(v.DnsZone, ".dev") || strings.HasPrefix(v.DnsZone, ".staging") {
				// we can't support host-less Azure CLI auth for dev & staging environments, as users will get errors like
				// ... `The user or administrator has not consented to use the application with ID '...' named
				// 'Microsoft Azure CLI'.`.
				continue
			}
			return v
		}
	}
	hostname := c.CanonicalHostName()
	for _, e := range envs {
		if strings.HasSuffix(hostname, e.DnsZone) {
			return e
		}
	}
	return defaultEnvironment
}
