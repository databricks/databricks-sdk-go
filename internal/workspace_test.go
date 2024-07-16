package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/ml"
	"github.com/databricks/databricks-sdk-go/service/workspace"
	"github.com/stretchr/testify/assert"
)

func myNotebookPath(t *testing.T, w *databricks.WorkspaceClient) string {
	ctx := context.Background()
	testDir := filepath.Join("/Users", me(t, w).UserName, ".sdk", RandomName("t-"))
	notebook := filepath.Join(testDir, RandomName("n-"))

	err := w.Workspace.MkdirsByPath(ctx, testDir)
	assert.NoError(t, err)
	t.Cleanup(func() {
		err = w.Workspace.Delete(ctx, workspace.Delete{
			Path:      testDir,
			Recursive: true,
		})
		assert.NoError(t, err)
	})

	return notebook
}

// TODO: Enable this testing after testing infrastructure is moved to OAuth
func TestGetOAuthToken(t *testing.T) {
	t.Skip("This test only works with OAuth credentials. Automatic testing uses PAT tokens")
	// Create an Endpoint that can be used to query
	ctx, w := workspaceTest(t)
	// For manual testing using Databricks CLI credentials
	// w = databricks.Must(databricks.NewWorkspaceClient())

	created, err := w.ModelRegistry.CreateModel(ctx, ml.CreateModelRequest{
		Name: RandomName("go-sdk-"),
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		deleteModel(t, w, ctx, created)
	})

	model, err := w.ModelRegistry.GetModel(ctx, ml.GetModelRequest{
		Name: created.RegisteredModel.Name,
	})
	assert.NoError(t, err)

	err = w.ModelRegistry.UpdateModel(ctx, ml.UpdateModelRequest{
		Name:        model.RegisteredModelDatabricks.Name,
		Description: RandomName("comment "),
	})
	assert.NoError(t, err)

	all, err := w.ModelRegistry.ListModelsAll(ctx, ml.ListModelsRequest{})
	assert.NoError(t, err)
	assert.True(t, len(all) >= 1)

	// Enable once the API returns the DataplaneInfo
	// endpoint, err := w.ServingEndpoints.Create(ctx, serving.CreateServingEndpoint{
	// 	Name: RandomName("go-sdk-"),
	// 	Config: serving.EndpointCoreConfigInput{
	// 		ServedModels: []serving.ServedModelInput{
	// 			{
	// 				ModelName:          model.RegisteredModelDatabricks.Name,
	// 				ModelVersion:       model.RegisteredModelDatabricks.LatestVersions[0].Version,
	// 				Name:               RandomName("go-sdk-"),
	// 				ScaleToZeroEnabled: true,
	// 				WorkloadSize:       serving.ServedModelInputWorkloadSizeSmall,
	// 			},
	// 		},
	// 	},
	// })
	// require.NoError(t, err)
	// t.Cleanup(func() {
	// 	err := w.ServingEndpoints.Delete(ctx, serving.DeleteServingEndpointRequest{
	// 		Name: endpoint.Name,
	// 	})
	// 	assert.NoError(t, err)
	// })
	// readyEndpoint, err := endpoint.Get()
	// require.NoError(t, err)

	// _, err = w.GetOAuthToken(readyEndpoint.DataplaneInfo.AuthorizationDetails)
	// require.NoError(t, err)
	// // TODO: Query DataPlane API with the token
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
	assert.NoError(t, err)

	// Get test notebook status
	getStatusResponse, err := w.Workspace.GetStatusByPath(ctx, notebook)
	assert.NoError(t, err)
	assert.True(t, getStatusResponse.Language == workspace.LanguagePython)
	assert.True(t, getStatusResponse.Path == notebook)
	assert.True(t, getStatusResponse.ObjectType == workspace.ObjectTypeNotebook)

	// Export the notebook and assert the contents
	exportResponse, err := w.Workspace.Export(ctx, workspace.ExportRequest{
		Format: workspace.ExportFormatSource,
		Path:   notebook,
	})
	assert.NoError(t, err)
	assert.True(t, exportResponse.Content == base64.StdEncoding.EncodeToString([]byte("# Databricks notebook source\nprint('hello from job')")))

	// Assert the test notebook is present in test dir using list api
	objects, err := w.Workspace.ListAll(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	assert.NoError(t, err)

	paths, err := w.Workspace.ObjectInfoPathToObjectIdMap(ctx, workspace.ListWorkspaceRequest{
		Path: filepath.Dir(notebook),
	})
	assert.NoError(t, err)
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
		assert.NoError(t, err)
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
		assert.NoError(t, err)
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
		assert.NoError(t, err)
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
	assert.NoError(t, err)

	allMyNotebooks, err := w.Workspace.RecursiveList(ctx, filepath.Join("/Users", me(t, w).UserName))
	assert.NoError(t, err)
	assert.True(t, len(allMyNotebooks) >= 1)
}
