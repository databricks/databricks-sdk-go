package internal

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccListDbfsIntegration(t *testing.T) {
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))
	ctx := context.Background()
	wsc := workspaces.New()
	dbfsTestFilePath := RandomName("/test-file-")
	dbfsTestDirPath := RandomName("/tmp/databricks-go-sdk/integration/dbfs/TestDir-")

	// create file
	createdFile, err := wsc.Dbfs.Create(ctx,
		dbfs.CreateRequest{
			Path:      dbfsTestFilePath,
			Overwrite: true,
		},
	)
	require.NoError(t, err)

	// write 'Hello, World!' to file
	err = wsc.Dbfs.AddBlock(ctx,
		dbfs.AddBlockRequest{
			Data:   base64.StdEncoding.EncodeToString([]byte("Hello, World!")),
			Handle: createdFile.Handle,
		},
	)
	require.NoError(t, err)

	// Close file handle
	err = wsc.Dbfs.Close(ctx,
		dbfs.CloseRequest{
			Handle: createdFile.Handle,
		},
	)
	require.NoError(t, err)

	// Get file status
	testFileStatus, err := wsc.Dbfs.GetStatus(ctx,
		dbfs.GetStatusRequest{
			Path: dbfsTestFilePath,
		},
	)
	require.NoError(t, err)
	assert.True(t, testFileStatus.Path == dbfsTestFilePath)
	assert.True(t, testFileStatus.IsDir == false)

	// List all files in workspace root and assert test file is found
	listOfFilesInWorkspaceRoot, err := wsc.Dbfs.ListStatus(ctx,
		dbfs.ListStatusRequest{
			Path: "/",
		},
	)
	require.NoError(t, err)
	foundTestFile := false
	for _, fileInfo := range listOfFilesInWorkspaceRoot.Files {
		if fileInfo.Path == dbfsTestFilePath {
			foundTestFile = true
		}
	}
	assert.True(t, foundTestFile)

	// make test dir in workspace root
	err = wsc.Dbfs.MkDirs(ctx,
		dbfs.MkDirsRequest{
			Path: dbfsTestDirPath,
		},
	)
	require.NoError(t, err)

	// move test file to test dir
	err = wsc.Dbfs.Move(ctx,
		dbfs.MoveRequest{
			SourcePath:      dbfsTestFilePath,
			DestinationPath: dbfsTestDirPath + dbfsTestFilePath,
		},
	)
	require.NoError(t, err)

	// put a new file (byebye-world.txt) in test dir
	err = wsc.Dbfs.Put(ctx,
		dbfs.PutRequest{
			Path:      dbfsTestDirPath + "/byebye-world.txt",
			Contents:  base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")),
			Overwrite: true,
		},
	)
	require.NoError(t, err)

	// assert on contents of original test file
	contentsTestFile, err := wsc.Dbfs.Read(ctx,
		dbfs.ReadRequest{
			Path: dbfsTestDirPath + dbfsTestFilePath,
		},
	)
	require.NoError(t, err)
	assert.True(t, contentsTestFile.Data == base64.StdEncoding.EncodeToString([]byte("Hello, World!")))

	// assert on contents of byebye-world.txt
	contentsByeByeWorldFile, err := wsc.Dbfs.Read(ctx,
		dbfs.ReadRequest{
			Path: dbfsTestDirPath + "/byebye-world.txt",
		},
	)
	require.NoError(t, err)
	assert.True(t, contentsByeByeWorldFile.Data == base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")))

	// recursively delete the test dir and the test files inside it
	err = wsc.Dbfs.Delete(ctx,
		dbfs.DeleteRequest{
			Path:      dbfsTestDirPath,
			Recursive: true,
		},
	)
	require.NoError(t, err)
}
