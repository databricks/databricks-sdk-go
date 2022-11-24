package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccListWorkspaceIntegration(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	testDirPath := RandomName("/tmp/databricks-go-sdk/integration/workspace/TestDir-")
	testFileName := RandomName("test-file-")

	t.Cleanup(func() {
		// Delete the test directory
		err := w.Workspace.Delete(ctx, workspace.Delete{
			Path:      testDirPath,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	// Make test directory
	err := w.Workspace.MkdirsByPath(ctx, testDirPath)
	require.NoError(t, err)

	// Import the test notebook
	err = w.Workspace.Import(ctx, workspace.Import{
		Path:      filepath.Join(testDirPath, testFileName),
		Format:    "SOURCE",
		Language:  "PYTHON",
		Content:   base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// Get test notebook status
	getStatusResponse, err := w.Workspace.GetStatusByPath(ctx, filepath.Join(testDirPath, testFileName))
	require.NoError(t, err)
	assert.True(t, getStatusResponse.Language == "PYTHON")
	assert.True(t, getStatusResponse.Path == filepath.Join(testDirPath, testFileName))
	assert.True(t, string(getStatusResponse.ObjectType) == "NOTEBOOK")

	// Export the notebook and assert the contents
	exportResponse, err := w.Workspace.Export(ctx, workspace.ExportRequest{
		DirectDownload: false,
		Format:         "SOURCE",
		Path:           filepath.Join(testDirPath, testFileName),
	})
	require.NoError(t, err)
	assert.True(t, exportResponse.Content == base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")))

	// Assert the test notebook is present in test dir using list api
	objects, err := w.Workspace.ListAll(ctx, workspace.ListRequest{
		Path: testDirPath,
	})
	require.NoError(t, err)
	foundTestNotebook := false
	for _, objectInfo := range objects {
		if objectInfo.Path == filepath.Join(testDirPath, testFileName) {
			foundTestNotebook = true
		}
	}
	assert.True(t, foundTestNotebook)
}
