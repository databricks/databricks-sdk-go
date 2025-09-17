package oidc

import (
	"context"
	"errors"
	"fmt"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

// NewAzureDevOpsIDTokenSource returns a new IDTokenSource that retrieves an IDToken
// from the Azure DevOps environment. This IDTokenSource is only valid when
// running in Azure DevOps Pipelines with OIDC enabled.
func NewAzureDevOpsIDTokenSource(client *httpclient.ApiClient, systemAccessToken, systemTeamFoundationCollectionUri, systemPlanId, systemJobId, systemTeamProjectId, systemHostType string) IDTokenSource {
	return &azureDevOpsIDTokenSource{
		systemAccessToken:                 systemAccessToken,
		systemTeamFoundationCollectionUri: systemTeamFoundationCollectionUri,
		refreshClient:                     client,
		systemPlanId:                      systemPlanId,
		systemJobId:                       systemJobId,
		systemTeamProjectId:               systemTeamProjectId,
		systemHostType:                    systemHostType,
	}
}

// azureDevOpsIDTokenSource retrieves JWT Tokens from Azure DevOps Pipelines.
type azureDevOpsIDTokenSource struct {
	systemAccessToken                 string
	systemTeamFoundationCollectionUri string
	refreshClient                     *httpclient.ApiClient
	systemPlanId                      string
	systemJobId                       string
	systemTeamProjectId               string
	systemHostType                    string
}

// IDToken returns a JWT Token for the specified audience. It will return
// an error if not running in Azure DevOps Pipelines.
func (a *azureDevOpsIDTokenSource) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	if a.systemAccessToken == "" {
		logger.Debugf(ctx, "Missing SYSTEM_ACCESSTOKEN, likely not calling from Azure DevOps Pipeline")
		return nil, errors.New("missing SYSTEM_ACCESSTOKEN")
	}
	if a.systemTeamFoundationCollectionUri == "" {
		logger.Debugf(ctx, "Missing SYSTEM_TEAMFOUNDATIONCOLLECTIONURI, likely not calling from Azure DevOps Pipeline")
		return nil, errors.New("missing SYSTEM_TEAMFOUNDATIONCOLLECTIONURI")
	}
	if a.systemPlanId == "" {
		logger.Debugf(ctx, "Missing SYSTEM_PLANID, likely not calling from Azure DevOps Pipeline")
		return nil, errors.New("missing SYSTEM_PLANID")
	}
	if a.systemJobId == "" {
		logger.Debugf(ctx, "Missing SYSTEM_JOBID, likely not calling from Azure DevOps Pipeline")
		return nil, errors.New("missing SYSTEM_JOBID")
	}
	if a.systemTeamProjectId == "" {
		logger.Debugf(ctx, "Missing SYSTEM_TEAMPROJECTID, likely not calling from Azure DevOps Pipeline")
		return nil, errors.New("missing SYSTEM_TEAMPROJECTID")
	}
	if a.systemHostType == "" {
		logger.Debugf(ctx, "Missing SYSTEM_HOSTTYPE, likely not calling from Azure DevOps Pipeline")
		return nil, errors.New("missing SYSTEM_HOSTTYPE")
	}

	// Azure DevOps OIDC endpoint format
	// Reference: https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/oidctoken/create?view=azure-devops-rest-7.1
	// Use systemHostType to determine the hub name dynamically (e.g., "build", "release", etc.)
	requestUrl := fmt.Sprintf("%s/%s/_apis/distributedtask/hubs/%s/plans/%s/jobs/%s/oidctoken?api-version=7.2-preview.1",
		a.systemTeamFoundationCollectionUri,
		a.systemTeamProjectId,
		a.systemHostType,
		a.systemPlanId,
		a.systemJobId)

	// Create request body with audience - Azure DevOps expects {"audience":"value"} format
	requestBody := struct {
		Audience string `json:"audience"`
	}{
		Audience: audience,
	}

	// Azure DevOps returns {"oidcToken":"***"} format, not {"value":"***"}
	var azureResp struct {
		OidcToken string `json:"oidcToken"`
	}

	err := a.refreshClient.Do(ctx, "POST", requestUrl,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", a.systemAccessToken)),
		httpclient.WithRequestData(requestBody),
		httpclient.WithResponseUnmarshal(&azureResp),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request ID token from Azure DevOps: %w", err)
	}

	if azureResp.OidcToken == "" {
		return nil, fmt.Errorf("empty OIDC token received from Azure DevOps")
	}

	return &IDToken{Value: azureResp.OidcToken}, nil
}
