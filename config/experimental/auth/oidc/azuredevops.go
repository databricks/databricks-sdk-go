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
func NewAzureDevOpsIDTokenSource(client *httpclient.ApiClient, systemAccessToken, systemTeamFoundationCollectionUri, systemPlanId, systemJobId, systemProject string) IDTokenSource {
	return &azureDevOpsIDTokenSource{
		systemAccessToken:                 systemAccessToken,
		systemTeamFoundationCollectionUri: systemTeamFoundationCollectionUri,
		refreshClient:                     client,
		systemPlanId:                      systemPlanId,
		systemJobId:                       systemJobId,
		systemProject:                     systemProject,
	}
}

// azureDevOpsIDTokenSource retrieves JWT Tokens from Azure DevOps Pipelines.
type azureDevOpsIDTokenSource struct {
	systemAccessToken                 string
	systemTeamFoundationCollectionUri string
	refreshClient                     *httpclient.ApiClient
	systemPlanId                      string
	systemJobId                       string
	systemProject                     string
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
	if a.systemProject == "" {
		logger.Debugf(ctx, "Missing SYSTEM_TEAMPROJECT, likely not calling from Azure DevOps Pipeline")
		return nil, errors.New("missing SYSTEM_TEAMPROJECT")
	}

	resp := &IDToken{}
	// Azure DevOps OIDC endpoint format
	// Reference: https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/oidctoken/create?view=azure-devops-rest-7.1
	requestUrl := fmt.Sprintf("%s/%s/_apis/distributedtask/hubs/build/plans/%s/jobs/%s/oidctoken?api-version=7.1-preview.1",
		a.systemTeamFoundationCollectionUri,
		a.systemProject,
		a.systemPlanId,
		a.systemJobId)

	// Create a struct for the request body
	requestBody := struct {
		Audience string `json:"audience,omitempty"`
	}{
		Audience: audience,
	}

	err := a.refreshClient.Do(ctx, "POST", requestUrl,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", a.systemAccessToken)),
		httpclient.WithRequestData(requestBody),
		httpclient.WithResponseUnmarshal(resp),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request ID token from Azure DevOps: %w", err)
	}

	return resp, nil
}
