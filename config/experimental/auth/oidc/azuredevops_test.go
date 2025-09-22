package oidc

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

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
			if !strings.Contains(gotErr.Error(), fmt.Sprintf("not calling from Azure DevOps Pipeline: missing env var %s", env)) {
				t.Errorf("NewAzureDevOpsIDTokenSource() want error containing %q, got error: %v", env, gotErr)
			}
			if got != nil {
				t.Errorf("NewAzureDevOpsIDTokenSource() want nil, got %v", got)
			}
		})
	}
}

func message(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func TestNewAzureDevOpsIDTokenSource_IDToken(t *testing.T) {
	testCases := []struct {
		desc       string
		ts         azureDevOpsIDTokenSource
		want       *IDToken
		wantErrMsg string
	}{
		// TODO: add test cases
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, gotErr := tc.ts.IDToken(context.Background(), "")
			gotErrMsg := message(gotErr)

			if tc.wantErrMsg != gotErrMsg {
				t.Errorf("IDToken(): want error message %q, got %q", tc.wantErrMsg, gotErrMsg)
			}
			if !cmp.Equal(got, tc.want) {
				t.Errorf("IDToken(): want %v, got %v", tc.want, got)
			}
		})
	}
}
