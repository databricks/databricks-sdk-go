package config

import (
	"fmt"
	"strings"
)

type Cloud string

const (
	CloudUnspecified Cloud = "Unspecified"
	CloudAWS         Cloud = "AWS"
	CloudAzure       Cloud = "Azure"
	CloudGCP         Cloud = "GCP"
)

type DatabricksEnvironment struct {
	Cloud              Cloud
	dnsZone            string
	azureApplicationID string
	azureEnvironment   *azureEnvironment
}

func (de DatabricksEnvironment) DeploymentURL(name string) string {
	return fmt.Sprintf("https://%s%s", name, de.dnsZone)
}

func (de DatabricksEnvironment) AzureServiceManagementEndpoint() string {
	if de.azureEnvironment == nil {
		return ""
	}
	return de.azureEnvironment.ServiceManagementEndpoint
}

func (de DatabricksEnvironment) AzureResourceManagerEndpoint() string {
	if de.azureEnvironment == nil {
		return ""
	}
	return de.azureEnvironment.ResourceManagerEndpoint
}

func (de DatabricksEnvironment) AzureActiveDirectoryEndpoint() string {
	if de.azureEnvironment == nil {
		return ""
	}
	return de.azureEnvironment.ActiveDirectoryEndpoint
}

// we default to AWS Prod environment since this case will be a hit for PVC
var defaultEnvironment = DatabricksEnvironment{
	Cloud:   CloudAWS,
	dnsZone: ".cloud.databricks.com",
}

var envs = []DatabricksEnvironment{
	{Cloud: CloudUnspecified, dnsZone: "localhost"},

	{Cloud: CloudAWS, dnsZone: ".dev.databricks.com"},
	{Cloud: CloudAWS, dnsZone: ".staging.cloud.databricks.com"},
	{Cloud: CloudAWS, dnsZone: ".cloud.databricks.us"},
	defaultEnvironment,

	{Cloud: CloudAzure, dnsZone: ".dev.azuredatabricks.net", azureApplicationID: "62a912ac-b58e-4c1d-89ea-b2dbfc7358fc", azureEnvironment: &publicCloud},
	{Cloud: CloudAzure, dnsZone: ".staging.azuredatabricks.net", azureApplicationID: "4a67d088-db5c-48f1-9ff2-0aace800ae68", azureEnvironment: &publicCloud},
	{Cloud: CloudAzure, dnsZone: ".azuredatabricks.net", azureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", azureEnvironment: &publicCloud},
	{Cloud: CloudAzure, dnsZone: ".databricks.azure.us", azureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", azureEnvironment: &usGovernmentCloud},
	{Cloud: CloudAzure, dnsZone: ".databricks.azure.cn", azureApplicationID: "2ff814a6-3304-4ab8-85cb-cd0e6f879c1d", azureEnvironment: &chinaCloud},

	{Cloud: CloudGCP, dnsZone: ".dev.gcp.databricks.com"},
	{Cloud: CloudGCP, dnsZone: ".staging.gcp.databricks.com"},
	{Cloud: CloudGCP, dnsZone: ".gcp.databricks.com"},
}

func (c *Config) Environment() DatabricksEnvironment {
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
			if v.azureEnvironment.Name != azureEnv {
				continue
			}
			if strings.HasPrefix(v.dnsZone, ".dev") || strings.HasPrefix(v.dnsZone, ".staging") {
				continue
			}
			return v
		}
	}
	hostname := c.CanonicalHostName()
	for _, e := range append(c.DatabricksEnvironments, envs...) {
		if strings.HasSuffix(hostname, e.dnsZone) {
			return e
		}
	}
	return defaultEnvironment
}
