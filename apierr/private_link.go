package apierr

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/common/environment"
)

type privateLinkInfo struct {
	serviceName   string
	endpointName  string
	referencePage string
}

func (p privateLinkInfo) errorMessage() string {
	privateLinkValidationError := fmt.Sprintf(
		`The requested workspace has %[1]s enabled and is not accessible from
the current network. Ensure that %[1]s is properly configured and that your
device has access to the %s. For more information, see %s.`,
		p.serviceName, p.endpointName, p.referencePage)
	return strings.ReplaceAll(privateLinkValidationError, "\n", " ")
}

var privateLinkInfoMap = map[environment.Cloud]privateLinkInfo{
	environment.CloudAWS: {
		serviceName:   "AWS PrivateLink",
		endpointName:  "AWS VPC endpoint",
		referencePage: "https://docs.databricks.com/en/security/network/classic/privatelink.html",
	},
	environment.CloudAzure: {
		serviceName:   "Azure Private Link",
		endpointName:  "Azure Private Link endpoint",
		referencePage: "https://learn.microsoft.com/en-us/azure/databricks/security/network/classic/private-link-standard#authentication-troubleshooting",
	},
	environment.CloudGCP: {
		serviceName:   "Private Service Connect",
		endpointName:  "GCP VPC endpoint",
		referencePage: "https://docs.gcp.databricks.com/en/security/network/classic/private-service-connect.html",
	},
}

func PrivateLinkValidationError(url *url.URL) *APIError {
	env := environment.GetEnvironmentForHostname(url.Host)
	info := privateLinkInfoMap[env.Cloud]
	return &APIError{
		ErrorCode:  "PRIVATE_LINK_VALIDATION_ERROR",
		StatusCode: 403,
		Message:    info.errorMessage(),
		unwrap:     ErrPermissionDenied,
	}
}
