package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccListWorkspaceIntegration(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ctx := context.Background()
	wsc := workspaces.New()
	testDirPath := RandomName("/tmp/databricks-go-sdk/integration/workspace/TestDir-")
	testFileName := RandomName("test-file-")

	t.Cleanup(func() {
		// Delete the test directory
		err := wsc.Workspace.Delete(ctx, workspace.Delete{
			Path:      testDirPath,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	// Make test directory
	err := wsc.Workspace.MkdirsByPath(ctx, testDirPath)
	require.NoError(t, err)

	// Import the test notebook
	err = wsc.Workspace.Import(ctx, workspace.Import{
		Path:      filepath.Join(testDirPath, testFileName),
		Format:    "SOURCE",
		Language:  "PYTHON",
		Content:   base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// Get test notebook status
	getStatusResponse, err := wsc.Workspace.GetStatusByPath(ctx, filepath.Join(testDirPath, testFileName))
	require.NoError(t, err)
	assert.True(t, getStatusResponse.Language == "PYTHON")
	assert.True(t, getStatusResponse.Path == filepath.Join(testDirPath, testFileName))
	assert.True(t, string(getStatusResponse.ObjectType) == "NOTEBOOK")

	// Export the notebook and assert the contents
	exportResponse, err := wsc.Workspace.Export(ctx, workspace.ExportRequest{
		DirectDownload: false,
		Format:         "SOURCE",
		Path:           filepath.Join(testDirPath, testFileName),
	})
	require.NoError(t, err)
	assert.True(t, exportResponse.Content == base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")))

	// Assert the test notebook is present in test dir using list api
	listReponse, err := wsc.Workspace.List(ctx, workspace.ListRequest{
		Path: testDirPath,
	})
	require.NoError(t, err)
	foundTestNotebook := false
	for _, objectInfo := range listReponse.Objects {
		if objectInfo.Path == filepath.Join(testDirPath, testFileName) {
			foundTestNotebook = true
		}
	}
	assert.True(t, foundTestNotebook)
}
