package internal

import (
	"context"
	"encoding/base64"
	"path/filepath"
	"testing"

	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/databricks/databricks-sdk-go/workspaces"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccListDbfsIntegration(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	t.Log(env)
	if env == "gcp" {
		t.Skip("dbfs tests are skipped because dbfs rest api is disabled on gcp")
	}

	ctx := context.Background()
	wsc := workspaces.New()
	testFile1 := RandomName("test-file-1-")
	testFile2 := RandomName("test-file-2-")
	dbfsTestDirPath1 := RandomName("/tmp/databricks-go-sdk/integration/dbfs/TestDir1-")
	dbfsTestDirPath2 := RandomName("/tmp/databricks-go-sdk/integration/dbfs/TestDir2-")

	t.Cleanup(func() {
		// recursively delete the test dir1 and any test files inside it
		err := wsc.Dbfs.Delete(ctx, dbfs.Delete{
			Path:      dbfsTestDirPath1,
			Recursive: true,
		})
		require.NoError(t, err)
		// recursively delete the test dir2 and any test files inside it
		err = wsc.Dbfs.Delete(ctx, dbfs.Delete{
			Path:      dbfsTestDirPath2,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	// make test dir2 in workspace root
	err := wsc.Dbfs.MkdirsByPath(ctx, dbfsTestDirPath2)
	require.NoError(t, err)

	// create testFile1 in test dir2
	createdFile, err := wsc.Dbfs.Create(ctx, dbfs.Create{
		Path:      filepath.Join(dbfsTestDirPath2, testFile1),
		Overwrite: true,
	})
	require.NoError(t, err)

	// write 'Hello, World!' to testFile1
	err = wsc.Dbfs.AddBlock(ctx, dbfs.AddBlock{
		Data:   base64.StdEncoding.EncodeToString([]byte("Hello, World!")),
		Handle: createdFile.Handle,
	})
	require.NoError(t, err)

	// Close testFile1 handle
	err = wsc.Dbfs.CloseByHandle(ctx, createdFile.Handle)
	require.NoError(t, err)

	// Get testFile1 status
	testFileStatus, err := wsc.Dbfs.GetStatusByPath(ctx, filepath.Join(dbfsTestDirPath2, testFile1))
	require.NoError(t, err)
	assert.True(t, testFileStatus.Path == filepath.Join(dbfsTestDirPath2, testFile1))
	assert.True(t, testFileStatus.IsDir == false)

	// List all files in test dir2 and assert testFile1 is found
	listOfFilesInWorkspaceRoot, err := wsc.Dbfs.ListByPath(ctx, dbfsTestDirPath2)
	require.NoError(t, err)
	foundTestFile := false
	for _, fileInfo := range listOfFilesInWorkspaceRoot.Files {
		if fileInfo.Path == filepath.Join(dbfsTestDirPath2, testFile1) {
			foundTestFile = true
		}
	}
	assert.True(t, foundTestFile)

	// make test dir1 in workspace root
	err = wsc.Dbfs.MkdirsByPath(ctx, dbfsTestDirPath1)
	require.NoError(t, err)

	// move testFile1 to test dir1
	err = wsc.Dbfs.Move(ctx, dbfs.Move{
		SourcePath:      filepath.Join(dbfsTestDirPath2, testFile1),
		DestinationPath: filepath.Join(dbfsTestDirPath1, testFile1),
	})
	require.NoError(t, err)

	// put a new file testFile2 in test dir1
	err = wsc.Dbfs.Put(ctx, dbfs.Put{
		Path:      filepath.Join(dbfsTestDirPath1, testFile2),
		Contents:  base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// assert on contents of testFile1
	contentsTestFile, err := wsc.Dbfs.Read(ctx, dbfs.ReadRequest{
		Path: filepath.Join(dbfsTestDirPath1, testFile1),
	})
	require.NoError(t, err)
	assert.True(t, contentsTestFile.Data == base64.StdEncoding.EncodeToString([]byte("Hello, World!")))

	// assert on contents of testFile2
	contentsByeByeWorldFile, err := wsc.Dbfs.Read(ctx, dbfs.ReadRequest{
		Path: filepath.Join(dbfsTestDirPath1, testFile2),
	})
	require.NoError(t, err)
	assert.True(t, contentsByeByeWorldFile.Data == base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")))
}
