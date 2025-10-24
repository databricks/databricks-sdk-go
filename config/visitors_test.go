package config

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/useragent"
	"github.com/stretchr/testify/assert"
)

func TestAuthInUserAgentVisitorDefault(t *testing.T) {
	c := &Config{
		AuthType: "PAT",
	}
	request := &http.Request{}
	err := authInUserAgentVisitor(c)(request)
	assert.NoError(t, err)
	// tests may be developed and run on different versions of different things
	uac := strings.Split(useragent.FromContext(request.Context()), " ")
	assert.Contains(t, uac, "auth/PAT")
}

func TestAuthInUserAgentVisitorCustom(t *testing.T) {
	c := &Config{
		AuthType: "PAT",
	}
	request := &http.Request{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, AuthType{}, "oath")
	request = request.WithContext(ctx)
	err := authInUserAgentVisitor(c)(request)
	assert.NoError(t, err)
	// tests may be developed and run on different versions of different things
	uac := strings.Split(useragent.FromContext(request.Context()), " ")
	assert.NotContains(t, uac, "auth/PAT")
	assert.Contains(t, uac, "auth/oath")
}

func TestWorkspaceOrgIdVisitor_AddsHeader(t *testing.T) {
	cfg := &Config{WorkspaceId: "12345"}
	request := &http.Request{Header: http.Header{}}

	err := workspaceOrgIdVisitor(cfg)(request)

	assert.NoError(t, err)
	assert.Equal(t, "12345", request.Header.Get("X-Databricks-Org-Id"))
}

func TestWorkspaceOrgIdVisitor_NoHeaderWhenEmpty(t *testing.T) {
	cfg := &Config{WorkspaceId: ""}
	request := &http.Request{Header: http.Header{}}

	err := workspaceOrgIdVisitor(cfg)(request)

	assert.NoError(t, err)
	assert.Empty(t, request.Header.Get("X-Databricks-Org-Id"))
}
