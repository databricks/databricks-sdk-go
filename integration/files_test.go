package integration

import (
	"bytes"
	"context"
	"encoding/base64"
	"hash/fnv"
	"io"
	"math/rand"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/catalog/v2"
	"github.com/databricks/databricks-sdk-go/databricks/config"
	"github.com/databricks/databricks-sdk-go/files/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type hashable []byte

func (buf hashable) Hash() uint32 {
	h := fnv.New32a()
	h.Write(buf)
	return h.Sum32()
}

func TestUcAccFilesUploadAndDownload(t *testing.T) {
	ctx, cfg, volume := setupUCVolume(t)

	filePath := RandomName("/Volumes/" + volume.CatalogName + "/" + volume.SchemaName + "/" + volume.Name + "/files-with-?-and-#-")
	FilesAPI, err := files.NewFilesClient(cfg)
	require.NoError(t, err)
	_, err = FilesAPI.Upload(ctx, files.UploadRequest{
		FilePath: filePath,
		Contents: io.NopCloser(strings.NewReader("abcd")),
	})
	require.NoError(t, err)
	fileHead, err := FilesAPI.GetMetadata(ctx, files.GetMetadataRequest{
		FilePath: filePath,
	})
	require.NoError(t, err)
	require.Equal(t, fileHead.ContentType, "application/octet-stream")
	require.Equal(t, fileHead.ContentLength, int64(4))
	t.Cleanup(func() {
		_, err = FilesAPI.DeleteByFilePath(ctx, filePath)
		assert.NoError(t, err)
	})
	raw, err := FilesAPI.DownloadByFilePath(ctx, filePath)
	require.NoError(t, err)
	contents, err := io.ReadAll(raw.Contents)
	require.NoError(t, err)
	defer raw.Contents.Close()

	assert.Equal(t, "abcd", string(contents))
}

func TestUcAccFilesDelete(t *testing.T) {
	ctx, cfg, volume := setupUCVolume(t)

	FilesAPI, err := files.NewFilesClient(cfg)
	require.NoError(t, err)
	filePath := RandomName("/Volumes/" + volume.CatalogName + "/" + volume.SchemaName + "/" + volume.Name + "/file-")
	require.NoError(t, createFile(t, ctx, FilesAPI, filePath, "Hello, world!"))

	_, err = FilesAPI.DeleteByFilePath(ctx, filePath)
	require.NoError(t, err)
}

func TestUcAccFilesGetMetadata(t *testing.T) {
	ctx, cfg, volume := setupUCVolume(t)

	FilesAPI, err := files.NewFilesClient(cfg)
	require.NoError(t, err)
	filePath := RandomName("/Volumes/" + volume.CatalogName + "/" + volume.SchemaName + "/" + volume.Name + "/file-")
	require.NoError(t, createFile(t, ctx, FilesAPI, filePath, "12345"))

	metadata, err := FilesAPI.GetMetadataByFilePath(ctx, filePath)
	require.NoError(t, err)

	assert.Equal(t, "application/octet-stream", metadata.ContentType)
	assert.Equal(t, int64(5), metadata.ContentLength)
	assert.NotEmpty(t, metadata.LastModified)
}

func TestUcAccFilesCreateDirectory(t *testing.T) {
	ctx, cfg, volume := setupUCVolume(t)

	FilesAPI, err := files.NewFilesClient(cfg)
	require.NoError(t, err)
	directoryPath := RandomName("/Volumes/" + volume.CatalogName + "/" + volume.SchemaName + "/" + volume.Name + "/directory-")
	require.NoError(t, createDirectory(t, ctx, FilesAPI, directoryPath))
}

func TestUcAccFilesListDirectoryContents(t *testing.T) {
	ctx, cfg, volume := setupUCVolume(t)

	FilesAPI, err := files.NewFilesClient(cfg)
	require.NoError(t, err)
	directoryPath := RandomName("/Volumes/" + volume.CatalogName + "/" + volume.SchemaName + "/" + volume.Name + "/directory-")
	require.NoError(t, createFile(t, ctx, FilesAPI, directoryPath+"/file01.txt", "Hello, world!"))
	require.NoError(t, createFile(t, ctx, FilesAPI, directoryPath+"/file02.txt", "Hello, world!"))
	require.NoError(t, createFile(t, ctx, FilesAPI, directoryPath+"/file03.txt", "Hello, world!"))

	response, err := FilesAPI.ListDirectoryContentsAll(ctx, files.ListDirectoryContentsRequest{DirectoryPath: directoryPath})
	require.NoError(t, err)

	assert.Len(t, response, 3)
}

func TestUcAccFilesDeleteDirectory(t *testing.T) {
	ctx, cfg, volume := setupUCVolume(t)

	FilesAPI, err := files.NewFilesClient(cfg)
	require.NoError(t, err)
	directoryPath := RandomName("/Volumes/" + volume.CatalogName + "/" + volume.SchemaName + "/" + volume.Name + "/directory-")
	require.NoError(t, createDirectory(t, ctx, FilesAPI, directoryPath))

	_, err = FilesAPI.DeleteDirectoryByDirectoryPath(ctx, directoryPath)
	assert.NoError(t, err)
}

func TestUcAccFilesGetDirectoryMetadata(t *testing.T) {
	ctx, cfg, volume := setupUCVolume(t)

	FilesAPI, err := files.NewFilesClient(cfg)
	require.NoError(t, err)
	directoryPath := RandomName("/Volumes/" + volume.CatalogName + "/" + volume.SchemaName + "/" + volume.Name + "/directory-")
	require.NoError(t, createDirectory(t, ctx, FilesAPI, directoryPath))

	_, err = FilesAPI.GetDirectoryMetadataByDirectoryPath(ctx, directoryPath)
	require.NoError(t, err)
}

func TestAccDbfsOpen(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	DbfsAPI, err := files.NewDbfsClient(cfg)
	require.NoError(t, err)
	if cfg.IsGcp() {
		t.Skip("dbfs not available on gcp")
	}

	path := RandomName("/tmp/.sdk/fake")
	rand.Seed(time.Now().UnixNano())
	in := make([]byte, 1.44*1e6)
	_, _ = rand.Read(in)

	defer DbfsAPI.Delete(ctx, files.Delete{
		Path: path,
	})

	// Upload through [io.Writer].
	{
		handle, err := DbfsAPI.Open(ctx, path, files.FileModeWrite)
		require.NoError(t, err)
		n, err := handle.Write(in)
		require.NoError(t, err)
		assert.Equal(t, len(in), int(n))
		require.NoError(t, handle.Close())

		// Verify contents hash.
		out, err := DbfsAPI.ReadFile(ctx, path)
		require.NoError(t, err)
		assert.Equal(t, hashable(in).Hash(), hashable(out).Hash())
	}

	// Upload through [io.Writer] should fail because the file exists.
	{
		_, err := DbfsAPI.Open(ctx, path, files.FileModeWrite)
		require.ErrorContains(t, err, "dbfs open: A file or directory already exists at the input path")
	}

	// Upload through [io.ReadFrom] with overwrite bit set.
	{
		handle, err := DbfsAPI.Open(ctx, path, files.FileModeWrite|files.FileModeOverwrite)
		require.NoError(t, err)
		n, err := handle.ReadFrom(bytes.NewReader(in))
		require.NoError(t, err)
		assert.Equal(t, len(in), int(n))
		require.NoError(t, handle.Close())

		// Verify contents hash.
		out, err := DbfsAPI.ReadFile(ctx, path)
		require.NoError(t, err)
		assert.Equal(t, hashable(in).Hash(), hashable(out).Hash())
	}

	// Download through [io.Reader] and let [io.ReadAll] determine buffer size.
	{
		handle, err := DbfsAPI.Open(ctx, path, files.FileModeRead)
		require.NoError(t, err)

		// Note: [io.ReadAll] always calls into the [io.Reader] interface.
		out, err := io.ReadAll(handle)
		require.NoError(t, err)

		// Verify contents hash.
		assert.Equal(t, hashable(in).Hash(), hashable(out).Hash())
	}

	// Download through [io.WriterTo].
	{
		handle, err := DbfsAPI.Open(ctx, path, files.FileModeRead)
		require.NoError(t, err)

		var buf bytes.Buffer
		_, err = handle.WriteTo(&buf)
		require.NoError(t, err)

		// Verify contents hash.
		assert.Equal(t, hashable(in).Hash(), hashable(buf.Bytes()).Hash())
	}
}

func TestAccDbfsOpenDirectory(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	DbfsAPI, err := files.NewDbfsClient(cfg)
	require.NoError(t, err)
	if cfg.IsGcp() {
		t.Skip("dbfs not available on gcp")
	}

	path := RandomName("/tmp/.sdk/fake")

	defer DbfsAPI.Delete(ctx, files.Delete{
		Path: path,
	})

	// Create directory.
	_, err = DbfsAPI.MkdirsByPath(ctx, path)
	require.NoError(t, err)

	// Try to open the directory for reading.
	_, err = DbfsAPI.Open(ctx, path, files.FileModeRead)
	assert.ErrorContains(t, err, "dbfs open: cannot open directory for reading")

	// Try to open the directory for writing.
	_, err = DbfsAPI.Open(ctx, path, files.FileModeWrite)
	assert.ErrorContains(t, err, "dbfs open: A file or directory already exists")

	// Try to open the directory for writing with overwrite flag set.
	_, err = DbfsAPI.Open(ctx, path, files.FileModeWrite|files.FileModeOverwrite)
	assert.ErrorContains(t, err, "dbfs open: A file or directory already exists")
}

func TestAccDbfsReadFileWriteFile(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	DbfsAPI, err := files.NewDbfsClient(cfg)
	require.NoError(t, err)
	if cfg.IsGcp() {
		t.Skip("dbfs not available on gcp")
	}

	path := RandomName("/tmp/.sdk/fake")
	rand.Seed(time.Now().UnixNano())
	in := make([]byte, 1.44*1e6)
	_, _ = rand.Read(in)

	defer DbfsAPI.Delete(ctx, files.Delete{
		Path: path,
	})

	// Write file to DBFS.
	err = DbfsAPI.WriteFile(ctx, path, in)
	require.NoError(t, err)

	// Verify contents hash.
	out, err := DbfsAPI.ReadFile(ctx, path)
	require.NoError(t, err)
	assert.Equal(t, hashable(in).Hash(), hashable(out).Hash())

	hello := []byte("Hello world!")

	// Writing to the same path should truncate the existing file.
	err = DbfsAPI.WriteFile(ctx, path, hello)
	require.NoError(t, err)

	// Verify contents hash.
	out, err = DbfsAPI.ReadFile(ctx, path)
	require.NoError(t, err)
	assert.Equal(t, hashable(hello).Hash(), hashable(out).Hash())
}

func TestAccListDbfsIntegration(t *testing.T) {
	ctx := workspaceTest(t)
	cfg := &config.Config{}
	DbfsAPI, err := files.NewDbfsClient(cfg)
	require.NoError(t, err)
	if cfg.IsGcp() {
		t.Skip("dbfs not available on gcp")
	}

	testFile1 := RandomName("f001-")
	testFile2 := RandomName("f002-")
	testPath1 := RandomName("/tmp/.sdk/001-")
	testPath2 := RandomName("/tmp/.sdk/002-")

	t.Cleanup(func() {
		// recursively delete the test dir1 and any test files inside it
		_, err := DbfsAPI.Delete(ctx, files.Delete{
			Path:      testPath1,
			Recursive: true,
		})
		require.NoError(t, err)
		// recursively delete the test dir2 and any test files inside it
		_, err = DbfsAPI.Delete(ctx, files.Delete{
			Path:      testPath2,
			Recursive: true,
		})
		require.NoError(t, err)
	})

	// make test dir2 in workspace root
	_, err = DbfsAPI.MkdirsByPath(ctx, testPath2)
	require.NoError(t, err)

	// create testFile1 in test dir2
	createdFile, err := DbfsAPI.Create(ctx, files.Create{
		Path:      filepath.Join(testPath2, testFile1),
		Overwrite: true,
	})
	require.NoError(t, err)

	// write 'Hello, World!' to testFile1
	_, err = DbfsAPI.AddBlock(ctx, files.AddBlock{
		Data:   base64.StdEncoding.EncodeToString([]byte("Hello, World!")),
		Handle: createdFile.Handle,
	})
	require.NoError(t, err)

	// Close testFile1 handle
	_, err = DbfsAPI.CloseByHandle(ctx, createdFile.Handle)
	require.NoError(t, err)

	// Get testFile1 status
	testFileStatus, err := DbfsAPI.GetStatusByPath(ctx, filepath.Join(testPath2, testFile1))
	require.NoError(t, err)
	assert.True(t, testFileStatus.Path == filepath.Join(testPath2, testFile1))
	assert.True(t, testFileStatus.IsDir == false)

	// List all files in test dir2 and assert testFile1 is found
	listOfFilesInWorkspaceRoot, err := DbfsAPI.ListByPath(ctx, testPath2)
	require.NoError(t, err)
	foundTestFile := false
	for _, fileInfo := range listOfFilesInWorkspaceRoot.Files {
		if fileInfo.Path == filepath.Join(testPath2, testFile1) {
			foundTestFile = true
		}
	}
	assert.True(t, foundTestFile)

	// make test dir1 in workspace root
	_, err = DbfsAPI.MkdirsByPath(ctx, testPath1)
	require.NoError(t, err)

	// move testFile1 to test dir1
	_, err = DbfsAPI.Move(ctx, files.Move{
		SourcePath:      filepath.Join(testPath2, testFile1),
		DestinationPath: filepath.Join(testPath1, testFile1),
	})
	require.NoError(t, err)

	// put a new file testFile2 in test dir1
	_, err = DbfsAPI.Put(ctx, files.Put{
		Path:      filepath.Join(testPath1, testFile2),
		Contents:  base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")),
		Overwrite: true,
	})
	require.NoError(t, err)

	// assert on contents of testFile1
	contentsTestFile, err := DbfsAPI.Read(ctx, files.ReadDbfsRequest{
		Path: filepath.Join(testPath1, testFile1),
	})
	require.NoError(t, err)
	assert.True(t, contentsTestFile.Data == base64.StdEncoding.EncodeToString([]byte("Hello, World!")))

	// assert on contents of testFile2
	contentsByeByeWorldFile, err := DbfsAPI.Read(ctx, files.ReadDbfsRequest{
		Path: filepath.Join(testPath1, testFile2),
	})
	require.NoError(t, err)
	assert.True(t, contentsByeByeWorldFile.Data == base64.StdEncoding.EncodeToString([]byte("Bye Bye, World!")))

	files, err := DbfsAPI.RecursiveList(ctx, path.Dir(testPath1))
	require.NoError(t, err)
	assert.True(t, len(files) >= 2)
}

func setupUCVolume(t *testing.T) (context.Context, *config.Config, *catalog.VolumeInfo) {
	ctx := ucwsTest(t)
	cfg := &config.Config{}
	SchemasAPI, err := catalog.NewSchemasClient(cfg)
	require.NoError(t, err)
	schema, err := SchemasAPI.Create(ctx, catalog.CreateSchema{
		CatalogName: "main",
		Name:        RandomName("files-"),
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		SchemasAPI.Delete(ctx, catalog.DeleteSchemaRequest{
			FullName: schema.FullName,
		})
	})

	VolumesAPI, err := catalog.NewVolumesClient(nil)
	require.NoError(t, err)
	volume, err := VolumesAPI.Create(ctx, catalog.CreateVolumeRequestContent{
		CatalogName: "main",
		SchemaName:  schema.Name,
		Name:        RandomName("files-"),
		VolumeType:  catalog.VolumeTypeManaged,
	})
	require.NoError(t, err)
	t.Cleanup(func() {
		VolumesAPI.Delete(ctx, catalog.DeleteVolumeRequest{
			Name: volume.FullName,
		})
	})

	return ctx, cfg, volume
}

// createDirectory creates a directory with the given contents and registers a cleanup function to delete it.
func createDirectory(t *testing.T, ctx context.Context, FilesAPI *files.FilesClient, directory string) error {
	t.Helper()

	_, err := FilesAPI.CreateDirectory(ctx, files.CreateDirectoryRequest{
		DirectoryPath: directory,
	})
	if err != nil {
		return err
	}

	t.Cleanup(func() {
		if _, err := FilesAPI.DeleteDirectoryByDirectoryPath(ctx, directory); err != nil {
			t.Log("failed to clean up directory:", err)
		}
	})
	return nil
}

// createFile creates a file with the given contents and registers a cleanup function to delete it.
func createFile(t *testing.T, ctx context.Context, FilesAPI *files.FilesClient, filePath, contents string) error {
	t.Helper()

	_, err := FilesAPI.Upload(ctx, files.UploadRequest{
		FilePath: filePath,
		Contents: io.NopCloser(strings.NewReader(contents)),
	})
	if err != nil {
		return err
	}

	t.Cleanup(func() {
		if _, err := FilesAPI.DeleteByFilePath(ctx, filePath); err != nil {
			t.Log("failed to clean up file:", err)
		}
	})
	return nil
}
