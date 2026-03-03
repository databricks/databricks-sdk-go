package config

import (
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
)

const (
	testHMHost        = "https://dummy-workspace.databricks.com"
	testHMAccHost     = "https://dummy-accounts.databricks.com"
	testHMAccountID   = "00000000-0000-0000-0000-000000000001"
	testHMWorkspaceID = "111111111111111"
)

func newTestAPIClient(transport fixtures.MappingTransport) *httpclient.ApiClient {
	return httpclient.NewApiClient(httpclient.ClientConfig{Transport: transport})
}

func TestGetHostMetadata_WorkspaceStaticOIDCEndpoint(t *testing.T) {
	client := newTestAPIClient(fixtures.MappingTransport{
		"GET /.well-known/databricks-config": {
			Status: 200,
			Response: map[string]string{
				"oidc_endpoint": testHMHost + "/oidc",
				"account_id":    testHMAccountID,
				"workspace_id":  testHMWorkspaceID,
			},
		},
	})
	meta, err := getHostMetadata(context.Background(), testHMHost, client)
	if err != nil {
		t.Fatal(err)
	}
	want := &hostMetadata{
		OIDCEndpoint: testHMHost + "/oidc",
		AccountID:    testHMAccountID,
		WorkspaceID:  testHMWorkspaceID,
	}
	if diff := cmp.Diff(want, meta); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestGetHostMetadata_AccountHostRawOIDCTemplate(t *testing.T) {
	client := newTestAPIClient(fixtures.MappingTransport{
		"GET /.well-known/databricks-config": {
			Status: 200,
			Response: map[string]string{
				"oidc_endpoint": testHMAccHost + "/oidc/accounts/{account_id}",
			},
		},
	})
	meta, err := getHostMetadata(context.Background(), testHMAccHost, client)
	if err != nil {
		t.Fatal(err)
	}
	want := &hostMetadata{
		OIDCEndpoint: testHMAccHost + "/oidc/accounts/{account_id}",
	}
	if diff := cmp.Diff(want, meta); diff != "" {
		t.Errorf("mismatch (-want +got):\n%s", diff)
	}
}

func TestGetHostMetadata_HTTPError(t *testing.T) {
	client := newTestAPIClient(fixtures.MappingTransport{
		"GET /.well-known/databricks-config": {
			Status: 404,
		},
	})
	_, err := getHostMetadata(context.Background(), testHMHost, client)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "fetching host metadata from") {
		t.Errorf("expected error containing %q, got %q", "fetching host metadata from", err.Error())
	}
}
