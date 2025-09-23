package oidc

import (
	"context"
	"fmt"
	"os"

	"github.com/databricks/databricks-sdk-go/httpclient"
)

// NewAzureDevOpsIDTokenSource returns a new IDTokenSource that retrieves an
// IDToken from an Azure DevOps environment.
//
// This IDTokenSource is only valid when running in Azure DevOps Pipelines.
func NewAzureDevOpsIDTokenSource(client *httpclient.ApiClient) (IDTokenSource, error) {
	ts := &azureDevOpsIDTokenSource{Client: client}

	if err := setFromAzureDevOpsEnv(&ts.AccessToken, "SYSTEM_ACCESSTOKEN"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.TeamFoundationCollectionURI, "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.PlanID, "SYSTEM_PLANID"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.JobID, "SYSTEM_JOBID"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.TeamProjectID, "SYSTEM_TEAMPROJECTID"); err != nil {
		return nil, err
	}
	if err := setFromAzureDevOpsEnv(&ts.HostType, "SYSTEM_HOSTTYPE"); err != nil {
		return nil, err
	}

	return ts, nil
}

func setFromAzureDevOpsEnv(s *string, envVar string) error {
	v := os.Getenv(envVar)
	if v == "" {
		if envVar == "SYSTEM_ACCESSTOKEN" {
			return fmt.Errorf("SYSTEM_ACCESSTOKEN env var not found, if calling from Azure DevOps Pipeline, please set this env var following https://learn.microsoft.com/en-us/azure/devops/pipelines/build/variables?view=azure-devops&tabs=yaml#systemaccesstoken")
		}
		return fmt.Errorf("not calling from Azure DevOps Pipeline: missing env var %s", envVar)
	}
	*s = v
	return nil
}

// azureDevOpsIDTokenSource retrieves JWT Tokens from Azure DevOps Pipelines.
type azureDevOpsIDTokenSource struct {
	Client                      *httpclient.ApiClient
	AccessToken                 string
	TeamFoundationCollectionURI string
	PlanID                      string
	JobID                       string
	TeamProjectID               string
	HostType                    string
}

// IDToken returns a JWT Token for the specified audience. For Azure DevOps OIDC,
// the audience parameter is ignored as Azure DevOps tokens always use "api://AzureADTokenExchange".
// It will return an error if not running in Azure DevOps Pipelines.
func (a *azureDevOpsIDTokenSource) IDToken(ctx context.Context, audience string) (*IDToken, error) {
	// Azure DevOps OIDC endpoint format
	// Reference: https://learn.microsoft.com/en-us/rest/api/azure/devops/distributedtask/oidctoken/create?view=azure-devops-rest-7.1
	// Use azureDevOpsHostType to determine the hub name dynamically (e.g., "build", "release", etc.)
	requestUrl := fmt.Sprintf("%s/%s/_apis/distributedtask/hubs/%s/plans/%s/jobs/%s/oidctoken?api-version=7.2-preview.1",
		a.TeamFoundationCollectionURI,
		a.TeamProjectID,
		a.HostType,
		a.PlanID,
		a.JobID,
	)

	// Azure DevOps returns {"oidcToken":"***"}, not {"value":"***"}.
	var azureResp struct {
		OIDCToken string `json:"oidcToken"`
	}

	err := a.Client.Do(ctx, "POST", requestUrl,
		httpclient.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken)),
		httpclient.WithResponseUnmarshal(&azureResp),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to request ID token from Azure DevOps: %w", err)
	}
	if azureResp.OIDCToken == "" {
		return nil, fmt.Errorf("empty OIDC token received from Azure DevOps")
	}

	return &IDToken{Value: azureResp.OIDCToken}, nil
}
