package config

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspaceIDHeaderVisitor_SetWhenPresent(t *testing.T) {
	visitor := workspaceIDHeaderVisitor(&Config{
		WorkspaceID: "6051921418418893",
	})
	req, _ := http.NewRequest("GET", "/api/2.1/jobs/list", nil)
	err := visitor(req)
	assert.NoError(t, err)
	assert.Equal(t, "6051921418418893", req.Header.Get("X-Databricks-Org-Id"))
}

func TestWorkspaceIDHeaderVisitor_NotSetWhenEmpty(t *testing.T) {
	visitor := workspaceIDHeaderVisitor(&Config{
		AccountID: "7a99b43c-b46c-432b-b0a7-814217701909",
	})
	req, _ := http.NewRequest("GET", "/api/2.0/accounts/7a99b43c/workspaces", nil)
	err := visitor(req)
	assert.NoError(t, err)
	assert.Empty(t, req.Header.Get("X-Databricks-Org-Id"))
}
