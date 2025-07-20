# Enhanced Files API for Databricks Go SDK

This package provides an enhanced Files API that extends the standard Files API with advanced functionality for handling large file uploads and downloads, similar to the Python SDK's `FilesExt` class.

## Features

### 1. Multipart Upload Support
- **Automatic Detection**: Automatically chooses between one-shot upload and multipart upload based on file size
- **Cloud Provider Support**: Supports AWS S3 and Azure Blob Storage multipart uploads
- **Configurable Chunk Size**: Default 100MB chunks, configurable
- **Retry Logic**: Built-in retry mechanism for failed uploads
- **URL Expiration Handling**: Handles expired presigned URLs gracefully

### 2. Resumable Upload Support
- **GCP Support**: Implements resumable upload for Google Cloud Storage
- **Chunked Upload**: Uploads files in configurable chunks
- **Resume Capability**: Can resume interrupted uploads
- **Progress Tracking**: Tracks upload progress and handles partial completions

### 3. Resilient Download Support
- **Automatic Recovery**: Automatically recovers from network failures during download
- **Offset Tracking**: Maintains download offset for seamless resumption
- **Configurable Retries**: Configurable retry limits for download recovery
- **Streaming Support**: Supports streaming downloads with recovery

## Usage

### Basic Setup

```go
package main

import (
    "context"
    "io"
    "strings"
    
    "github.com/databricks/databricks-sdk-go/client"
    "github.com/databricks/databricks-sdk-go/config"
    "github.com/databricks/databricks-sdk-go/service/files"
)

func main() {
    // Create configuration
    cfg := &config.Config{
        Host:  "https://your-workspace.cloud.databricks.com",
        Token: "your-token",
    }
    
    // Create client
    databricksClient, err := client.New(cfg)
    if err != nil {
        panic(err)
    }
    
    // Create enhanced Files API
    filesExt := files.NewFilesExt(databricksClient)
    
    // Use the enhanced API...
}
```

### Upload Examples

#### Simple Upload (Automatic Method Selection)

```go
// The API automatically chooses the best upload method
content := strings.NewReader("Hello, World!")
err := filesExt.Upload(context.Background(), files.UploadRequest{
    FilePath:  "/Volumes/my-catalog/my-schema/my-volume/file.txt",
    Contents:  io.NopCloser(content),
    Overwrite: true,
})
if err != nil {
    panic(err)
}
```

#### Large File Upload (Multipart)

```go
// For large files, multipart upload is automatically used
largeContent := strings.NewReader(strings.Repeat("Large content ", 1000000))
err := filesExt.Upload(context.Background(), files.UploadRequest{
    FilePath:  "/Volumes/my-catalog/my-schema/my-volume/large-file.txt",
    Contents:  io.NopCloser(largeContent),
    Overwrite: true,
})
if err != nil {
    panic(err)
}
```

### Download Examples

#### Simple Download

```go
response, err := filesExt.Download(context.Background(), files.DownloadRequest{
    FilePath: "/Volumes/my-catalog/my-schema/my-volume/file.txt",
})
if err != nil {
    panic(err)
}
defer response.Contents.Close()

// Read the content
content, err := io.ReadAll(response.Contents)
if err != nil {
    panic(err)
}

println("Downloaded content:", string(content))
```

#### Streaming Download with Recovery

```go
response, err := filesExt.Download(context.Background(), files.DownloadRequest{
    FilePath: "/Volumes/my-catalog/my-schema/my-volume/large-file.txt",
})
if err != nil {
    panic(err)
}
defer response.Contents.Close()

// The download automatically recovers from network failures
buffer := make([]byte, 1024)
for {
    n, err := response.Contents.Read(buffer)
    if err == io.EOF {
        break
    }
    if err != nil {
        panic(err)
    }
    
    // Process the chunk
    processChunk(buffer[:n])
}
```

## Configuration

The enhanced Files API supports configuration through the client configuration, similar to the Python SDK. Configuration parameters are automatically read from the client config and can be set via environment variables or programmatically.

### Client Configuration

Configuration is integrated with the Databricks client configuration system:

```go
// Create configuration with custom Files API settings
cfg := &config.Config{
    Host:  "https://your-workspace.cloud.databricks.com",
    Token: "your-token",
    
    // Files API configuration
    FilesAPIMultipartUploadMinStreamSize:                             50 * 1024 * 1024, // 50MB
    FilesAPIMultipartUploadChunkSize:                                 50 * 1024 * 1024, // 50MB
    FilesAPIMultipartUploadBatchURLCount:                             5,
    FilesAPIMultipartUploadMaxRetries:                                5,
    FilesAPIMultipartUploadSingleChunkUploadTimeoutSeconds:           600,
    FilesAPIMultipartUploadURLExpirationDurationSeconds:              7200, // 2 hours
    FilesAPIClientDownloadMaxTotalRecovers:                           15,
    FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing:         5,
}

// Create client
databricksClient, err := client.New(cfg)
if err != nil {
    panic(err)
}

// Create enhanced Files API
filesExt := files.NewFilesExt(databricksClient)

// Get current configuration
currentConfig := filesExt.GetUploadConfig()
```

### Environment Variables

You can also configure the Files API using environment variables:

```bash
export DATABRICKS_FILES_API_MULTIPART_UPLOAD_MIN_STREAM_SIZE=52428800  # 50MB
export DATABRICKS_FILES_API_MULTIPART_UPLOAD_CHUNK_SIZE=52428800       # 50MB
export DATABRICKS_FILES_API_MULTIPART_UPLOAD_BATCH_URL_COUNT=5
export DATABRICKS_FILES_API_MULTIPART_UPLOAD_MAX_RETRIES=5
export DATABRICKS_FILES_API_MULTIPART_UPLOAD_SINGLE_CHUNK_UPLOAD_TIMEOUT_SECONDS=600
export DATABRICKS_FILES_API_MULTIPART_UPLOAD_URL_EXPIRATION_DURATION_SECONDS=7200
export DATABRICKS_FILES_API_CLIENT_DOWNLOAD_MAX_TOTAL_RECOVERS=15
export DATABRICKS_FILES_API_CLIENT_DOWNLOAD_MAX_TOTAL_RECOVERS_WITHOUT_PROGRESSING=5
```

### Default Configuration

The enhanced Files API automatically sets sensible defaults in the config object when no configuration is provided. These defaults are applied during client initialization:

```go
// Default values automatically set in config object
// - FilesAPIMultipartUploadMinStreamSize: 100MB
// - FilesAPIMultipartUploadChunkSize: 100MB
// - FilesAPIMultipartUploadBatchURLCount: 10
// - FilesAPIMultipartUploadMaxRetries: 3
// - FilesAPIMultipartUploadSingleChunkUploadTimeoutSeconds: 300
// - FilesAPIMultipartUploadURLExpirationDurationSeconds: 3600 (1 hour)
// - FilesAPIClientDownloadMaxTotalRecovers: 10
// - FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3
```

The defaults are applied automatically when you create a client, so you don't need to set them manually unless you want to override them.

### Configuration Parameters

| Parameter | Environment Variable | Default | Description |
|-----------|---------------------|---------|-------------|
| `FilesAPIMultipartUploadMinStreamSize` | `DATABRICKS_FILES_API_MULTIPART_UPLOAD_MIN_STREAM_SIZE` | 100MB | Minimum stream size to trigger multipart upload |
| `FilesAPIMultipartUploadChunkSize` | `DATABRICKS_FILES_API_MULTIPART_UPLOAD_CHUNK_SIZE` | 100MB | Chunk size for multipart uploads |
| `FilesAPIMultipartUploadBatchURLCount` | `DATABRICKS_FILES_API_MULTIPART_UPLOAD_BATCH_URL_COUNT` | 10 | Number of upload URLs to request in a batch |
| `FilesAPIMultipartUploadMaxRetries` | `DATABRICKS_FILES_API_MULTIPART_UPLOAD_MAX_RETRIES` | 3 | Maximum number of retries for multipart upload |
| `FilesAPIMultipartUploadSingleChunkUploadTimeoutSeconds` | `DATABRICKS_FILES_API_MULTIPART_UPLOAD_SINGLE_CHUNK_UPLOAD_TIMEOUT_SECONDS` | 300 | Timeout for single chunk upload in seconds |
| `FilesAPIMultipartUploadURLExpirationDurationSeconds` | `DATABRICKS_FILES_API_MULTIPART_UPLOAD_URL_EXPIRATION_DURATION_SECONDS` | 3600 | URL expiration duration in seconds |
| `FilesAPIClientDownloadMaxTotalRecovers` | `DATABRICKS_FILES_API_CLIENT_DOWNLOAD_MAX_TOTAL_RECOVERS` | 10 | Maximum total recovers for downloads |
| `FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing` | `DATABRICKS_FILES_API_CLIENT_DOWNLOAD_MAX_TOTAL_RECOVERS_WITHOUT_PROGRESSING` | 3 | Maximum recovers without progressing for downloads |

## Architecture

### Multipart Upload Flow

1. **Initiation**: Call the initiate-upload endpoint
2. **Method Detection**: Server responds with either multipart or resumable upload details
3. **Chunk Upload**: Upload file in chunks using presigned URLs
4. **Completion**: Complete the upload with all ETags

### Resumable Upload Flow

1. **Initiation**: Create resumable upload URL
2. **Chunked Upload**: Upload in chunks with Content-Range headers
3. **Progress Tracking**: Track confirmed bytes from server responses
4. **Completion**: Final chunk marks upload as complete

### Resilient Download Flow

1. **Initial Request**: Start download from offset 0
2. **Streaming**: Stream content with automatic offset tracking
3. **Error Recovery**: On failure, restart download from current offset
4. **Retry Limits**: Respect configured retry limits

## Error Handling

The enhanced Files API provides robust error handling:

- **Network Failures**: Automatic retry with exponential backoff
- **URL Expiration**: Automatic URL refresh for expired presigned URLs
- **Partial Failures**: Graceful handling of partial upload/download failures
- **Resource Cleanup**: Automatic cleanup of incomplete uploads

## Best Practices

1. **Use Unity Catalog Volumes**: Prefer Unity Catalog volumes for better performance and security
2. **Configure Timeouts**: Set appropriate timeouts for your network conditions
3. **Monitor Progress**: Use logging to monitor upload/download progress
4. **Handle Errors**: Always check for errors and implement appropriate error handling
5. **Resource Management**: Always close response streams to prevent resource leaks

## Comparison with Python SDK

This Go implementation provides feature parity with the Python SDK's `FilesExt` class:

| Feature | Python SDK | Go SDK |
|---------|------------|--------|
| Multipart Upload | ✅ | ✅ |
| Resumable Upload | ✅ | ✅ |
| Resilient Download | ✅ | ✅ |
| Automatic Method Selection | ✅ | ✅ |
| Configurable Retries | ✅ | ✅ |
| URL Expiration Handling | ✅ | ✅ |
| Progress Tracking | ✅ | ✅ |

## Limitations

- **Cloud Provider Support**: Currently supports AWS S3, Azure Blob Storage, and Google Cloud Storage
- **File Size Limits**: Subject to cloud provider limits (typically 5TB for multipart uploads)
- **Concurrent Uploads**: No built-in support for concurrent uploads of the same file
- **Resume Across Sessions**: Download resume only works within the same session

## Contributing

When contributing to this package:

1. Follow Go coding standards
2. Add tests for new functionality
3. Update documentation for API changes
4. Ensure backward compatibility
5. Test with different cloud providers 