package oidc

import (
	"context"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

// defaultEnv is the default environment variables that must be set to
// successfully create an Azure DevOps ID token source.
var defaultEnv = map[string]string{
	"SYSTEM_ACCESSTOKEN":                 "test-access-token",
	"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI": "test-team-foundation-collection-uri",
	"SYSTEM_PLANID":                      "test-plan-id",
	"SYSTEM_JOBID":                       "test-job-id",
	"SYSTEM_TEAMPROJECTID":               "test-team-project-id",
	"SYSTEM_HOSTTYPE":                    "test-host-type",
}

func TestNewAzureDevOpsIDTokenSource_success(t *testing.T) {
	for k, v := range defaultEnv {
		t.Setenv(k, v)
	}

	want := &azureDevOpsIDTokenSource{
		AccessToken:                 "test-access-token",
		TeamFoundationCollectionURI: "test-team-foundation-collection-uri",
		PlanID:                      "test-plan-id",
		JobID:                       "test-job-id",
		TeamProjectID:               "test-team-project-id",
		HostType:                    "test-host-type",
	}

	got, gotErr := NewAzureDevOpsIDTokenSource(nil)
	if gotErr != nil {
		t.Fatalf("NewAzureDevOpsIDTokenSource() want no error, got error: %v", gotErr)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("NewAzureDevOpsIDTokenSource() mismatch (-want +got):\n%s", diff)
	}
}

func TestNewAzureDevOpsIDTokenSource_missingEnv(t *testing.T) {
	// Guarantee that the tests are run in a deterministic order.
	sortedKeys := make([]string, 0, len(defaultEnv))
	for k := range defaultEnv {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, env := range sortedKeys {
		t.Run(fmt.Sprintf("missing %s", env), func(t *testing.T) {
			// Set all env vars except the missing one.
			for k, v := range defaultEnv {
				if k != env {
					t.Setenv(k, v)
				}
			}

			got, gotErr := NewAzureDevOpsIDTokenSource(nil)

			if gotErr == nil {
				t.Fatalf("NewAzureDevOpsIDTokenSource() want error, got none")
			}
			// SYSTEM_ACCESSTOKEN has a special error message format
			var expectedErrorSubstring string
			if env == "SYSTEM_ACCESSTOKEN" {
				expectedErrorSubstring = "SYSTEM_ACCESSTOKEN env var not found"
			} else {
				expectedErrorSubstring = fmt.Sprintf("not calling from Azure DevOps Pipeline: missing env var %s", env)
			}

			if !strings.Contains(gotErr.Error(), expectedErrorSubstring) {
				t.Errorf("NewAzureDevOpsIDTokenSource() want error containing %q, got error: %v", expectedErrorSubstring, gotErr)
			}
			if got != nil {
				t.Errorf("NewAzureDevOpsIDTokenSource() want nil, got %v", got)
			}
		})
	}
}

func TestNewAzureDevOpsIDTokenSource_IDToken(t *testing.T) {
	testCases := []struct {
		desc          string
		ts            azureDevOpsIDTokenSource
		httpTransport http.RoundTripper
		want          *IDToken
		wantErr       bool
	}{
		{
			desc: "server error response",
			ts: azureDevOpsIDTokenSource{
				AccessToken:                 "token-1337",
				TeamFoundationCollectionURI: "https://dev.azure.com/myorg",
				PlanID:                      "plan123",
				JobID:                       "job456",
				TeamProjectID:               "project789",
				HostType:                    "build",
			},
			httpTransport: fixtures.MappingTransport{
				"POST /myorg/project789/_apis/distributedtask/hubs/build/plans/plan123/jobs/job456/oidctoken?api-version=7.2-preview.1": {
					Status: http.StatusInternalServerError,
					ExpectedHeaders: map[string]string{
						"Authorization": "Bearer token-1337",
						"Accept":        "application/json",
					},
				},
			},
			wantErr: true,
		},
		{
			desc: "success with build host type",
			ts: azureDevOpsIDTokenSource{
				AccessToken:                 "token-1337",
				TeamFoundationCollectionURI: "https://dev.azure.com/myorg",
				PlanID:                      "plan123",
				JobID:                       "job456",
				TeamProjectID:               "project789",
				HostType:                    "build",
			},
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
			want: &IDToken{
				Value: "azure-devops-id-token-42",
			},
			wantErr: false,
		},
		{
			desc: "success with release host type",
			ts: azureDevOpsIDTokenSource{
				AccessToken:                 "token-1337",
				TeamFoundationCollectionURI: "https://dev.azure.com/myorg",
				PlanID:                      "plan123",
				JobID:                       "job456",
				TeamProjectID:               "project789",
				HostType:                    "release",
			},
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
			want: &IDToken{
				Value: "azure-devops-release-token-42",
			},
			wantErr: false,
		},
		{
			desc: "empty oidc token in response",
			ts: azureDevOpsIDTokenSource{
				AccessToken:                 "token-1337",
				TeamFoundationCollectionURI: "https://dev.azure.com/myorg",
				PlanID:                      "plan123",
				JobID:                       "job456",
				TeamProjectID:               "project789",
				HostType:                    "build",
			},
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
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			cli := httpclient.NewApiClient(httpclient.ClientConfig{
				Transport: tc.httpTransport,
			})
			tc.ts.Client = cli

			got, gotErr := tc.ts.IDToken(context.Background(), "test-audience")

			if tc.wantErr && gotErr == nil {
				t.Errorf("IDToken(): expected error, got none")
			}
			if !tc.wantErr && gotErr != nil {
				t.Errorf("IDToken(): got error %q, want none", gotErr)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("IDToken(): mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
