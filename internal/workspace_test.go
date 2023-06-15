package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func myNotebookPath(t *testing.T, w *databricks.WorkspaceClient) string {
	ctx := context.Background()
	testDir := filepath.Join("/Users", me(t, w).UserName, ".sdk", RandomName("t-"))
	notebook := filepath.Join(testDir, RandomName("n-"))

	err := w.Workspace.MkdirsByPath(ctx, testDir)
	require.NoError(t, err)
	t.Cleanup(func() {
		err = w.Workspace.Delete(ctx, workspace.Delete{
			Path:      testDir,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	return notebook
}

func TestAccWorkspaceIntegration(t *testing.T) {
	ctx, w := workspaceTest(t)
	notebook := myNotebookPath(t, w)

	// Import the test notebook
	err := w.Workspace.Import(ctx, workspace.Import{
		Path:      notebook,
		Format:    workspace.ImportFormatSource,
		Language:  workspace.LanguagePython,
		Content:   base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// Get test notebook status
	getStatusResponse, err := w.Workspace.GetStatusByPath(ctx, notebook)
	require.NoError(t, err)
	assert.True(t, getStatusResponse.Language == workspace.LanguagePython)
	assert.True(t, getStatusResponse.Path == notebook)
	assert.True(t, getStatusResponse.ObjectType == workspace.ObjectTypeNotebook)

	// Export the notebook and assert the contents
	exportResponse, err := w.Workspace.Export(ctx, workspace.ExportRequest{
		Format: workspace.ExportFormatSource,
		Path:   notebook,
	})
	require.NoError(t, err)
	assert.True(t, exportResponse.Content == base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")))

	// Assert the test notebook is present in test dir using list api
	objects, err := w.Workspace.ListAll(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	require.NoError(t, err)

	paths, err := w.Workspace.ObjectInfoPathToObjectIdMap(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	require.NoError(t, err)
	assert.Equal(t, len(objects), len(paths))
	assert.Contains(t, paths, notebook)
}

func TestAccWorkspaceUploadNotebookWithFileExtensionNoTranspile(t *testing.T) {
	// TODO: remove NoTranspile suffix once other languages get Upload/Donwload features
	ctx, w := workspaceTest(t)

	notebookPath := filepath.Join("/Users", me(t, w).UserName, RandomName("notebook-")+".py")

	err := w.Workspace.Upload(ctx, notebookPath, strings.NewReader("print(1)"))
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Workspace.Delete(ctx, workspace.Delete{
			Path: notebookPath,
		})
		require.NoError(t, err)
	})

	info, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	assert.NoError(t, err)
	assert.Equal(t, workspace.LanguagePython, info.Language)
	assert.Equal(t, workspace.ObjectTypeNotebook, info.ObjectType)

	contents, err := w.Workspace.ReadFile(ctx, notebookPath)
	assert.NoError(t, err)

	assert.Equal(t, "# Databricks notebook source\nprint(1)", string(contents))
}

func TestAccWorkspaceUploadNotebookWithFileNoExtensionNoTranspile(t *testing.T) {
	// TODO: remove NoTranspile suffix once other languages get Upload/Donwload features
	ctx, w := workspaceTest(t)

	notebookPath := filepath.Join("/Users", me(t, w).UserName, RandomName("notebook-"))

	err := w.Workspace.Upload(ctx, notebookPath, strings.NewReader("print(1)"),
		workspace.UploadLanguage(workspace.LanguagePython))
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Workspace.Delete(ctx, workspace.Delete{
			Path: notebookPath,
		})
		require.NoError(t, err)
	})

	info, err := w.Workspace.GetStatusByPath(ctx, notebookPath)
	assert.NoError(t, err)
	assert.Equal(t, workspace.LanguagePython, info.Language)
	assert.Equal(t, workspace.ObjectTypeNotebook, info.ObjectType)

	contents, err := w.Workspace.ReadFile(ctx, notebookPath)
	assert.NoError(t, err)

	assert.Equal(t, "# Databricks notebook source\nprint(1)", string(contents))
}

func TestAccWorkspaceUploadFileNoTranspile(t *testing.T) {
	// TODO: remove NoTranspile suffix once other languages get Upload/Donwload features
	ctx, w := workspaceTest(t)

	txtPath := filepath.Join("/Users", me(t, w).UserName, RandomName("txt-"))

	err := w.Workspace.Upload(ctx, txtPath, strings.NewReader("print(1)"),
		workspace.UploadFormat(workspace.ImportFormatAuto))
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Workspace.Delete(ctx, workspace.Delete{
			Path: txtPath,
		})
		require.NoError(t, err)
	})

	info, err := w.Workspace.GetStatusByPath(ctx, txtPath)
	assert.NoError(t, err)
	assert.Equal(t, workspace.ObjectTypeFile, info.ObjectType)

	contents, err := w.Workspace.ReadFile(ctx, txtPath)
	assert.NoError(t, err)

	assert.Equal(t, "print(1)", string(contents))
}

func TestAccWorkspaceRecursiveListNoTranspile(t *testing.T) {
	ctx, w := workspaceTest(t)
	notebook := myNotebookPath(t, w)

	// Import the test notebook
	err := w.Workspace.Upload(ctx, notebook+".py", strings.NewReader("print(1)"),
		workspace.UploadOverwrite())
	require.NoError(t, err)

	allMyNotebooks, err := w.Workspace.RecursiveList(ctx, filepath.Join("/Users", me(t, w).UserName))
	require.NoError(t, err)
	assert.True(t, len(allMyNotebooks) >= 1)
}
