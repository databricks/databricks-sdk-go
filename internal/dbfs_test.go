package internal

import (
	"bytes"
	"encoding/base64"
	"hash/fnv"
	"io"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/service/dbfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccDbfsUtilities(t *testing.T) {
	ctx, w := workspaceTest(t)
	if w.Config.IsGcp() {
		t.Skip("dbfs not available on gcp")
	}

	path := RandomName("/tmp/.sdk/fake")
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
	ctx, w := workspaceTest(t)
	if w.Config.IsGcp() {
		t.Skip("dbfs not available on gcp")
	}

	testFile1 := RandomName("f001-")
	testFile2 := RandomName("f002-")
	testPath1 := RandomName("/tmp/.sdk/001-")
	testPath2 := RandomName("/tmp/.sdk/002-")

	t.Cleanup(func() {
		// recursively delete the test dir1 and any test files inside it
		err := w.Dbfs.Delete(ctx, dbfs.Delete{
			Path:      testPath1,
			Recursive: true,
		})
		require.NoError(t, err)
		// recursively delete the test dir2 and any test files inside it
		err = w.Dbfs.Delete(ctx, dbfs.Delete{
			Path:      testPath2,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	// make test dir2 in workspace root
	err := w.Dbfs.MkdirsByPath(ctx, testPath2)
	require.NoError(t, err)

	// create testFile1 in test dir2
	createdFile, err := w.Dbfs.Create(ctx, dbfs.Create{
		Path:      filepath.Join(testPath2, testFile1),
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
	testFileStatus, err := w.Dbfs.GetStatusByPath(ctx, filepath.Join(testPath2, testFile1))
	require.NoError(t, err)
	assert.True(t, testFileStatus.Path == filepath.Join(testPath2, testFile1))
	assert.True(t, testFileStatus.IsDir == false)

	// List all files in test dir2 and assert testFile1 is found
	listOfFilesInWorkspaceRoot, err := w.Dbfs.ListByPath(ctx, testPath2)
	require.NoError(t, err)
	foundTestFile := false
	for _, fileInfo := range listOfFilesInWorkspaceRoot.Files {
		if fileInfo.Path == filepath.Join(testPath2, testFile1) {
			foundTestFile = true
		}
	}
	assert.True(t, foundTestFile)

	// make test dir1 in workspace root
	err = w.Dbfs.MkdirsByPath(ctx, testPath1)
	require.NoError(t, err)

	// move testFile1 to test dir1
	err = w.Dbfs.Move(ctx, dbfs.Move{
		SourcePath:      filepath.Join(testPath2, testFile1),
		DestinationPath: filepath.Join(testPath1, testFile1),
	})
	require.NoError(t, err)

	// put a new file testFile2 in test dir1
	err = w.Dbfs.Put(ctx, dbfs.Put{
		Path:      filepath.Join(testPath1, testFile2),
		Contents:  base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// assert on contents of testFile1
	contentsTestFile, err := w.Dbfs.Read(ctx, dbfs.Read{
		Path: filepath.Join(testPath1, testFile1),
	})
	require.NoError(t, err)
	assert.True(t, contentsTestFile.Data == base64.StdEncoding.EncodeToString([]byte("Hello, World!")))

	// assert on contents of testFile2
	contentsByeByeWorldFile, err := w.Dbfs.Read(ctx, dbfs.Read{
		Path: filepath.Join(testPath1, testFile2),
	})
	require.NoError(t, err)
	assert.True(t, contentsByeByeWorldFile.Data == base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")))

	files, err := w.Dbfs.RecursiveList(ctx, path.Dir(testPath1))
	require.NoError(t, err)
	assert.True(t, len(files) >= 2)
}
