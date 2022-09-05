package internal

import (
	"context"
	"encoding/base64"
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
	testFileName := RandomName("/test-file-")

	// Make test directory
	err := wsc.Workspace.Mkdirs(ctx,
		workspace.MkdirsRequest{
			Path: testDirPath,
		},
	)
	require.NoError(t, err)

	// Import the test notebook
	err = wsc.Workspace.Import(ctx,
		workspace.ImportRequest{
			Path:      testDirPath + testFileName,
			Format:    "SOURCE",
			Language:  "PYTHON",
			Content:   base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")),
			Overwrite: true,
		},
	)
	require.NoError(t, err)

	// Get test notebook status
	getStatusResponse, err := wsc.Workspace.GetStatus(ctx,
		workspace.GetStatusRequest{
			Path: testDirPath + testFileName,
		},
	)
	require.NoError(t, err)
	assert.True(t, getStatusResponse.Language == "PYTHON")
	assert.True(t, getStatusResponse.Path == testDirPath+testFileName)
	assert.True(t, string(getStatusResponse.ObjectType) == "NOTEBOOK")

	// Export the notebook and assert the contents
	exportResponse, err := wsc.Workspace.Export(ctx,
		workspace.ExportRequest{
			DirectDownload: false,
			Format:         "SOURCE",
			Path:           testDirPath + testFileName,
		},
	)
	require.NoError(t, err)
	assert.True(t, exportResponse.Content == base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")))

	// Assert the test notebook is present in test dir using list api
	listReponse, err := wsc.Workspace.List(ctx,
		workspace.ListRequest{
			Path: testDirPath,
		},
	)
	require.NoError(t, err)
	foundTestNotebook := false
	for _, objectInfo := range listReponse.Objects {
		if objectInfo.Path == testDirPath+testFileName {
			foundTestNotebook = true
		}
	}
	assert.True(t, foundTestNotebook)

	// Delete the test directory
	err = wsc.Workspace.Delete(ctx,
		workspace.DeleteRequest{
			Path:      testDirPath,
			Recursive: true,
		},
	)
	require.NoError(t, err)
}
