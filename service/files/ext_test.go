package files

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
)

func TestFilesExt_Upload(t *testing.T) {
	// This is a test to demonstrate the usage of the enhanced Files API
	// In a real scenario, you would need a valid Databricks client

	// Example configuration
	cfg := &config.Config{
		Host:  "https://your-workspace.cloud.databricks.com",
		Token: "your-token",
	}

	// Create client
	databricksClient, err := client.New(cfg)
	if err != nil {
		t.Skipf("Skipping test - unable to create client: %v", err)
	}

	// Create enhanced Files API
	filesExt := NewFilesExt(databricksClient)

	// Example 1: Upload a small file (uses one-shot upload)
	smallContent := strings.NewReader("Hello, World!")
	err = filesExt.Upload(context.Background(), UploadRequest{
		FilePath:  "/Volumes/my-catalog/my-schema/my-volume/small-file.txt",
		Contents:  io.NopCloser(smallContent),
		Overwrite: true,
	})
	if err != nil {
		t.Logf("Upload failed: %v", err)
	}

	// Example 2: Upload a large file (uses multipart upload)
	largeContent := strings.NewReader(strings.Repeat("Large file content ", 1000000)) // ~20MB
	err = filesExt.Upload(context.Background(), UploadRequest{
		FilePath:  "/Volumes/my-catalog/my-schema/my-volume/large-file.txt",
		Contents:  io.NopCloser(largeContent),
		Overwrite: true,
	})
	if err != nil {
		t.Logf("Large file upload failed: %v", err)
	}
}

func TestFilesExt_Download(t *testing.T) {
	// Example configuration
	cfg := &config.Config{
		Host:  "https://your-workspace.cloud.databricks.com",
		Token: "your-token",
	}

	// Create client
	databricksClient, err := client.New(cfg)
	if err != nil {
		t.Skipf("Skipping test - unable to create client: %v", err)
	}

	// Create enhanced Files API
	filesExt := NewFilesExt(databricksClient)

	// Example: Download a file with resilient download
	response, err := filesExt.Download(context.Background(), DownloadRequest{
		FilePath: "/Volumes/my-catalog/my-schema/my-volume/file.txt",
	})
	if err != nil {
		t.Logf("Download failed: %v", err)
		return
	}

	// Read the content
	content, err := io.ReadAll(response.Contents)
	if err != nil {
		t.Logf("Failed to read content: %v", err)
		return
	}

	t.Logf("Downloaded file size: %d bytes", len(content))
	response.Contents.Close()
}

func TestUploadConfig(t *testing.T) {
	config := DefaultUploadConfig()

	// Verify default values
	if config.MultipartUploadMinStreamSize != 100*1024*1024 {
		t.Errorf("Expected MultipartUploadMinStreamSize to be 100MB, got %d", config.MultipartUploadMinStreamSize)
	}

	if config.MultipartUploadChunkSize != 100*1024*1024 {
		t.Errorf("Expected MultipartUploadChunkSize to be 100MB, got %d", config.MultipartUploadChunkSize)
	}

	if config.MultipartUploadBatchURLCount != 10 {
		t.Errorf("Expected MultipartUploadBatchURLCount to be 10, got %d", config.MultipartUploadBatchURLCount)
	}

	if config.MultipartUploadMaxRetries != 3 {
		t.Errorf("Expected MultipartUploadMaxRetries to be 3, got %d", config.MultipartUploadMaxRetries)
	}
}

func TestFillBuffer(t *testing.T) {
	filesExt := &FilesExt{}

	// Test with buffer smaller than desired size
	buffer := []byte("hello")
	input := strings.NewReader("world")

	result := filesExt.fillBuffer(buffer, 10, input)
	expected := []byte("helloworld")

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	// Test with buffer already large enough
	buffer = []byte("hello world")
	input = strings.NewReader("extra")

	result = filesExt.fillBuffer(buffer, 5, input)
	expected = []byte("hello world")

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestExtractRangeOffset(t *testing.T) {
	filesExt := &FilesExt{}

	// Test valid range string
	rangeStr := "bytes=0-1023"
	offset := filesExt.extractRangeOffset(rangeStr)
	if offset == nil || *offset != 1023 {
		t.Errorf("Expected offset 1023, got %v", offset)
	}

	// Test empty range string
	offset = filesExt.extractRangeOffset("")
	if offset != nil {
		t.Errorf("Expected nil offset, got %v", offset)
	}

	// Test invalid range string
	offset = filesExt.extractRangeOffset("invalid")
	if offset != nil {
		t.Errorf("Expected nil offset, got %v", offset)
	}
}

func TestGetURLExpireTime(t *testing.T) {
	filesExt := &FilesExt{}
	config := DefaultUploadConfig()

	expireTime := filesExt.getURLExpireTime(config)

	// Verify the format is correct (RFC 3339)
	if len(expireTime) != 20 || !strings.HasSuffix(expireTime, "Z") {
		t.Errorf("Expected RFC 3339 format, got %s", expireTime)
	}
}
