package internal

import (
	"bytes"
	"context"
	"encoding/base64"
	"hash/fnv"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func assertAcceptance(t *testing.T) {
	env := GetEnvOrSkipTest(t, "CLOUD_ENV")
	if env == "gcp" {
		t.Skip("dbfs tests are skipped because dbfs rest api is disabled on gcp")
	}
}

func TestAccDbfsUtilities(t *testing.T) {
	assertAcceptance(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}

	path := RandomName("/tmp/databricks-go-sdk/fake")
	rand.Seed(time.Now().UnixNano())
	in := make([]byte, 1.44*1e6)
	_, _ = rand.Read(in)
	h := fnv.New32a()
	h.Write(in)
	inHash := h.Sum32()

	err := w.Dbfs.Overwrite(ctx, path, bytes.NewReader(in))
	require.NoError(t, err)

	defer w.Dbfs.Delete(ctx, dbfs.Delete{
		Path: path,
	})

	r, err := w.Dbfs.Open(ctx, path)
	require.NoError(t, err)

	download, _ := os.Create("/path/to/local")
	io.Copy(download, r)

	out, err := io.ReadAll(r)
	require.NoError(t, err)

	h = fnv.New32a()
	h.Write(out)
	outHash := h.Sum32()

	require.Equal(t, outHash, inHash)
}

func TestAccListDbfsIntegration(t *testing.T) {
	assertAcceptance(t)
	ctx := context.Background()
	w := databricks.Must(databricks.NewWorkspaceClient())
	if w.Config.IsAccountsClient() {
		t.SkipNow()
	}
	testFile1 := RandomName("test-file-1-")
	testFile2 := RandomName("test-file-2-")
	dbfsTestDirPath1 := RandomName("/tmp/databricks-go-sdk/integration/dbfs/TestDir1-")
	dbfsTestDirPath2 := RandomName("/tmp/databricks-go-sdk/integration/dbfs/TestDir2-")

	t.Cleanup(func() {
		// recursively delete the test dir1 and any test files inside it
		err := w.Dbfs.Delete(ctx, dbfs.Delete{
			Path:      dbfsTestDirPath1,
			Recursive: true,
		})
		require.NoError(t, err)
		// recursively delete the test dir2 and any test files inside it
		err = w.Dbfs.Delete(ctx, dbfs.Delete{
			Path:      dbfsTestDirPath2,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	// make test dir2 in workspace root
	err := w.Dbfs.MkdirsByPath(ctx, dbfsTestDirPath2)
	require.NoError(t, err)

	// create testFile1 in test dir2
	createdFile, err := w.Dbfs.Create(ctx, dbfs.Create{
		Path:      filepath.Join(dbfsTestDirPath2, testFile1),
		Overwrite: true,
	})
	require.NoError(t, err)

	// write 'Hello, World!' to testFile1
	err = w.Dbfs.AddBlock(ctx, dbfs.AddBlock{
		Data:   base64.StdEncoding.EncodeToString([]byte("Hello, World!")),
		Handle: createdFile.Handle,
	})
	require.NoError(t, err)

	// Close testFile1 handle
	err = w.Dbfs.CloseByHandle(ctx, createdFile.Handle)
	require.NoError(t, err)

	// Get testFile1 status
	testFileStatus, err := w.Dbfs.GetStatusByPath(ctx, filepath.Join(dbfsTestDirPath2, testFile1))
	require.NoError(t, err)
	assert.True(t, testFileStatus.Path == filepath.Join(dbfsTestDirPath2, testFile1))
	assert.True(t, testFileStatus.IsDir == false)

	// List all files in test dir2 and assert testFile1 is found
	listOfFilesInWorkspaceRoot, err := w.Dbfs.ListByPath(ctx, dbfsTestDirPath2)
	require.NoError(t, err)
	foundTestFile := false
	for _, fileInfo := range listOfFilesInWorkspaceRoot.Files {
		if fileInfo.Path == filepath.Join(dbfsTestDirPath2, testFile1) {
			foundTestFile = true
		}
	}
	assert.True(t, foundTestFile)

	// make test dir1 in workspace root
	err = w.Dbfs.MkdirsByPath(ctx, dbfsTestDirPath1)
	require.NoError(t, err)

	// move testFile1 to test dir1
	err = w.Dbfs.Move(ctx, dbfs.Move{
		SourcePath:      filepath.Join(dbfsTestDirPath2, testFile1),
		DestinationPath: filepath.Join(dbfsTestDirPath1, testFile1),
	})
	require.NoError(t, err)

	// put a new file testFile2 in test dir1
	err = w.Dbfs.Put(ctx, dbfs.Put{
		Path:      filepath.Join(dbfsTestDirPath1, testFile2),
		Contents:  base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// assert on contents of testFile1
	contentsTestFile, err := w.Dbfs.Read(ctx, dbfs.Read{
		Path: filepath.Join(dbfsTestDirPath1, testFile1),
	})
	require.NoError(t, err)
	assert.True(t, contentsTestFile.Data == base64.StdEncoding.EncodeToString([]byte("Hello, World!")))

	// assert on contents of testFile2
	contentsByeByeWorldFile, err := w.Dbfs.Read(ctx, dbfs.Read{
		Path: filepath.Join(dbfsTestDirPath1, testFile2),
	})
	require.NoError(t, err)
	assert.True(t, contentsByeByeWorldFile.Data == base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")))
}
