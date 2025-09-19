package oidc

import (
	"context"
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

func TestAzureDevOpsIDTokenSource(t *testing.T) {
	testCases := []struct {
		desc                                   string
		azureDevOpsAccessToken                 string
		azureDevOpsTeamFoundationCollectionUri string
		azureDevOpsPlanId                      string
		azureDevOpsJobId                       string
		azureDevOpsTeamProjectId               string
		azureDevOpsHostType                    string
		audience                               string
		httpTransport                          http.RoundTripper
		wantToken                              *IDToken
		wantError                              bool
	}{
		{
			desc:                                   "missing access token",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsJobId:                       "job456",
			azureDevOpsTeamProjectId:               "project789",
			azureDevOpsHostType:                    "build",
			wantError:                              true,
		},
		{
			desc:                     "missing team foundation collection uri",
			azureDevOpsAccessToken:   "token-1337",
			azureDevOpsPlanId:        "plan123",
			azureDevOpsJobId:         "job456",
			azureDevOpsTeamProjectId: "project789",
			azureDevOpsHostType:      "build",
			wantError:                true,
		},
		{
			desc:                                   "missing plan id",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsJobId:                       "job456",
			azureDevOpsTeamProjectId:               "project789",
			azureDevOpsHostType:                    "build",
			wantError:                              true,
		},
		{
			desc:                                   "missing job id",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsTeamProjectId:               "project789",
			azureDevOpsHostType:                    "build",
			wantError:                              true,
		},
		{
			desc:                                   "missing team project id",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsJobId:                       "job456",
			azureDevOpsHostType:                    "build",
			wantError:                              true,
		},
		{
			desc:                                   "missing host type",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsJobId:                       "job456",
			azureDevOpsTeamProjectId:               "project789",
			wantError:                              true,
		},
		{
			desc:                                   "error getting token - server error",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsJobId:                       "job456",
			azureDevOpsTeamProjectId:               "project789",
			azureDevOpsHostType:                    "build",
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/project789/_apis/distributedtask/hubs/build/plans/plan123/jobs/job456/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusInternalServerError,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer token-1337",
						"Accept":        "application/json",
					},
				},
			},
			wantError: true,
		},
		{
			desc:                                   "success with build host type",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsJobId:                       "job456",
			azureDevOpsTeamProjectId:               "project789",
			azureDevOpsHostType:                    "build",
			audience:                               "test-audience",
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/project789/_apis/distributedtask/hubs/build/plans/plan123/jobs/job456/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer token-1337",
						"Accept":        "application/json",
					},
					Response: `{"oidcToken": "azure-devops-id-token-42"}`,
				},
			},
			wantToken: &IDToken{
				Value: "azure-devops-id-token-42",
			},
			wantError: false,
		},
		{
			desc:                                   "success with release host type",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsJobId:                       "job456",
			azureDevOpsTeamProjectId:               "project789",
			azureDevOpsHostType:                    "release",
			audience:                               "test-audience",
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/project789/_apis/distributedtask/hubs/release/plans/plan123/jobs/job456/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer token-1337",
						"Accept":        "application/json",
					},
					Response: `{"oidcToken": "azure-devops-release-token-42"}`,
				},
			},
			wantToken: &IDToken{
				Value: "azure-devops-release-token-42",
			},
			wantError: false,
		},
		{
			desc:                                   "success with empty oidc token in response",
			azureDevOpsAccessToken:                 "token-1337",
			azureDevOpsTeamFoundationCollectionUri: "https://dev.azure.com/myorg",
			azureDevOpsPlanId:                      "plan123",
			azureDevOpsJobId:                       "job456",
			azureDevOpsTeamProjectId:               "project789",
			azureDevOpsHostType:                    "build",
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/project789/_apis/distributedtask/hubs/build/plans/plan123/jobs/job456/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusOK,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer token-1337",
						"Accept":        "application/json",
					},
					Response: `{"oidcToken": ""}`,
				},
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			cli := httpclient.NewApiClient(httpclient.ClientConfig{
				Transport: tc.httpTransport,
			})
			p := &azureDevOpsIDTokenSource{
				azureDevOpsAccessToken:                 tc.azureDevOpsAccessToken,
				azureDevOpsTeamFoundationCollectionUri: tc.azureDevOpsTeamFoundationCollectionUri,
				azureDevOpsPlanId:                      tc.azureDevOpsPlanId,
				azureDevOpsJobId:                       tc.azureDevOpsJobId,
				azureDevOpsTeamProjectId:               tc.azureDevOpsTeamProjectId,
				azureDevOpsHostType:                    tc.azureDevOpsHostType,
				refreshClient:                          cli,
			}
			token, gotErr := p.IDToken(context.Background(), tc.audience)

			if tc.wantError && gotErr == nil {
				t.Errorf("IDToken(): expected error, got none")
			}
			if !tc.wantError && gotErr != nil {
				t.Errorf("IDToken(): got error %q, want none", gotErr)
			}
			if diff := cmp.Diff(tc.wantToken, token); diff != "" {
				t.Errorf("IDToken(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
