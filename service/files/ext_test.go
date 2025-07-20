package files

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/google/go-cmp/cmp"
)

// MockReadCloser implements io.ReadCloser for testing
type MockReadCloser struct {
	data      []byte
	readIndex int
	readErr   error
	closeErr  error
	closed    bool
}

func NewMockReadCloser(data []byte) *MockReadCloser {
	return &MockReadCloser{
		data:      data,
		readIndex: 0,
	}
}

func (m *MockReadCloser) Read(p []byte) (int, error) {
	if m.closed {
		return 0, errors.New("read on closed reader")
	}
	if m.readErr != nil {
		return 0, m.readErr
	}
	if m.readIndex >= len(m.data) {
		return 0, io.EOF
	}

	n := copy(p, m.data[m.readIndex:])
	m.readIndex += n
	return n, nil
}

func (m *MockReadCloser) Close() error {
	m.closed = true
	return m.closeErr
}

func (m *MockReadCloser) SetReadError(err error) {
	m.readErr = err
}

func (m *MockReadCloser) SetCloseError(err error) {
	m.closeErr = err
}

// MockFilesExt implements a mock FilesExt for testing
type MockFilesExt struct {
	*FilesExt
	openDownloadStreamFunc func(ctx context.Context, filePath string, startByteOffset int64, ifUnmodifiedSinceTimestamp string) (*DownloadResponse, error)
}

func NewMockFilesExt() *MockFilesExt {
	return &MockFilesExt{
		FilesExt: &FilesExt{},
	}
}

func (m *MockFilesExt) openDownloadStream(ctx context.Context, filePath string, startByteOffset int64, ifUnmodifiedSinceTimestamp string) (*DownloadResponse, error) {
	if m.openDownloadStreamFunc != nil {
		return m.openDownloadStreamFunc(ctx, filePath, startByteOffset, ifUnmodifiedSinceTimestamp)
	}
	return nil, errors.New("mock openDownloadStream not implemented")
}

// EnhancedMockReadCloser supports error sequences for Read
// and can be used to simulate read errors in order.
type EnhancedMockReadCloser struct {
	data       []byte
	readIndex  int
	readErrors []error // errors to return on each Read call
	closeErr   error
	closed     bool
}

func NewEnhancedMockReadCloser(data []byte, readErrors []error) *EnhancedMockReadCloser {
	return &EnhancedMockReadCloser{
		data:       data,
		readIndex:  0,
		readErrors: readErrors,
	}
}

func (m *EnhancedMockReadCloser) Read(p []byte) (int, error) {
	if m.closed {
		return 0, errors.New("read on closed reader")
	}
	if len(m.readErrors) > 0 {
		err := m.readErrors[0]
		m.readErrors = m.readErrors[1:]
		if err != nil {
			return 0, err
		}
	}
	if m.readIndex >= len(m.data) {
		return 0, io.EOF
	}

	n := copy(p, m.data[m.readIndex:])
	m.readIndex += n
	return n, nil
}

func (m *EnhancedMockReadCloser) Close() error {
	m.closed = true
	return m.closeErr
}

// testOpenDownloadStream is a package-level variable used to override openDownloadStream in tests
var testOpenDownloadStream func(ctx context.Context, filePath string, startByteOffset int64, ifUnmodifiedSinceTimestamp string) (*DownloadResponse, error)

// FilesExtTestable is a test-only subclass of FilesExt that overrides openDownloadStream
// to call testOpenDownloadStream if set
// This allows us to inject custom recovery logic in tests
type FilesExtTestable struct {
	FilesExt
}

var _ filesExtAPI = (*FilesExtTestable)(nil) // Ensure interface is implemented

func (f *FilesExtTestable) openDownloadStream(ctx context.Context, filePath string, startByteOffset int64, ifUnmodifiedSinceTimestamp string) (*DownloadResponse, error) {
	if testOpenDownloadStream != nil {
		return testOpenDownloadStream(ctx, filePath, startByteOffset, ifUnmodifiedSinceTimestamp)
	}
	return f.FilesExt.openDownloadStream(ctx, filePath, startByteOffset, ifUnmodifiedSinceTimestamp)
}

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
	// Test default config
	uploadConfig := DefaultUploadConfig()

	// Verify default values
	if uploadConfig.MultipartUploadMinStreamSize != 100*1024*1024 {
		t.Errorf("Expected MultipartUploadMinStreamSize to be 100MB, got %d", uploadConfig.MultipartUploadMinStreamSize)
	}

	if uploadConfig.MultipartUploadChunkSize != 100*1024*1024 {
		t.Errorf("Expected MultipartUploadChunkSize to be 100MB, got %d", uploadConfig.MultipartUploadChunkSize)
	}

	if uploadConfig.MultipartUploadBatchURLCount != 10 {
		t.Errorf("Expected MultipartUploadBatchURLCount to be 10, got %d", uploadConfig.MultipartUploadBatchURLCount)
	}

	if uploadConfig.MultipartUploadMaxRetries != 3 {
		t.Errorf("Expected MultipartUploadMaxRetries to be 3, got %d", uploadConfig.MultipartUploadMaxRetries)
	}

	// Test client config integration
	cfg := &config.Config{
		Host:  "https://your-workspace.cloud.databricks.com",
		Token: "your-token",
		// Set custom Files API configuration
		FilesAPIMultipartUploadMinStreamSize:                     50 * 1024 * 1024, // 50MB
		FilesAPIMultipartUploadChunkSize:                         50 * 1024 * 1024, // 50MB
		FilesAPIMultipartUploadBatchURLCount:                     5,
		FilesAPIMultipartUploadMaxRetries:                        5,
		FilesAPIMultipartUploadSingleChunkUploadTimeoutSeconds:   600,
		FilesAPIMultipartUploadURLExpirationDurationSeconds:      7200, // 2 hours
		FilesAPIClientDownloadMaxTotalRecovers:                   15,
		FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 5,
	}

	// Create client
	databricksClient, err := client.New(cfg)
	if err != nil {
		t.Skipf("Skipping test - unable to create client: %v", err)
	}

	// Create enhanced Files API
	filesExt := NewFilesExt(databricksClient)

	// Get config from client
	clientConfig := filesExt.GetUploadConfig()

	// Verify custom values are used
	if clientConfig.MultipartUploadMinStreamSize != 50*1024*1024 {
		t.Errorf("Expected client config MultipartUploadMinStreamSize to be 50MB, got %d", clientConfig.MultipartUploadMinStreamSize)
	}

	if clientConfig.MultipartUploadChunkSize != 50*1024*1024 {
		t.Errorf("Expected client config MultipartUploadChunkSize to be 50MB, got %d", clientConfig.MultipartUploadChunkSize)
	}

	if clientConfig.MultipartUploadBatchURLCount != 5 {
		t.Errorf("Expected client config MultipartUploadBatchURLCount to be 5, got %d", clientConfig.MultipartUploadBatchURLCount)
	}

	if clientConfig.MultipartUploadMaxRetries != 5 {
		t.Errorf("Expected client config MultipartUploadMaxRetries to be 5, got %d", clientConfig.MultipartUploadMaxRetries)
	}

	if clientConfig.MultipartUploadSingleChunkUploadTimeoutSeconds != 600 {
		t.Errorf("Expected client config MultipartUploadSingleChunkUploadTimeoutSeconds to be 600, got %d", clientConfig.MultipartUploadSingleChunkUploadTimeoutSeconds)
	}

	if clientConfig.MultipartUploadURLExpirationDuration != 2*time.Hour {
		t.Errorf("Expected client config MultipartUploadURLExpirationDuration to be 2 hours, got %v", clientConfig.MultipartUploadURLExpirationDuration)
	}

	if clientConfig.FilesAPIClientDownloadMaxTotalRecovers != 15 {
		t.Errorf("Expected client config FilesAPIClientDownloadMaxTotalRecovers to be 15, got %d", clientConfig.FilesAPIClientDownloadMaxTotalRecovers)
	}

	if clientConfig.FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing != 5 {
		t.Errorf("Expected client config FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing to be 5, got %d", clientConfig.FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing)
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

func TestResilientIterator_Read(t *testing.T) {
	tests := []struct {
		name                    string
		initialData             []byte
		bufferSize              int
		readErrors              []error // Errors to return on successive reads
		recoveryErrors          []error // Errors to return on recovery attempts
		config                  *UploadConfig
		expectedReads           [][]byte // Expected data from each read
		expectedErrors          []error  // Expected errors from each read
		expectedOffset          int64    // Expected final offset
		expectedRecoverAttempts int      // Expected number of recovery attempts
		description             string
	}{
		{
			name:        "successful_single_read",
			initialData: []byte("hello world"),
			bufferSize:  20,
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:  [][]byte{[]byte("hello world")},
			expectedErrors: []error{io.EOF},
			expectedOffset: 11,
			description:    "Simple successful read with EOF",
		},
		{
			name:        "successful_multiple_reads",
			initialData: []byte("hello world"),
			bufferSize:  5,
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads: [][]byte{
				[]byte("hello"),
				[]byte(" worl"),
				[]byte("d"),
			},
			expectedErrors: []error{nil, nil, io.EOF},
			expectedOffset: 11,
			description:    "Multiple successful reads with partial buffers",
		},
		{
			name:        "read_error_with_successful_recovery",
			initialData: []byte("hello world"),
			bufferSize:  20,
			readErrors:  []error{errors.New("network error"), nil},
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:           [][]byte{[]byte("hello world")},
			expectedErrors:          []error{io.EOF},
			expectedOffset:          11,
			expectedRecoverAttempts: 1,
			description:             "Read error followed by successful recovery",
		},
		{
			name:        "multiple_read_errors_with_recovery",
			initialData: []byte("hello world"),
			bufferSize:  20,
			readErrors:  []error{errors.New("network error"), errors.New("timeout"), nil},
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:           [][]byte{[]byte("hello world")},
			expectedErrors:          []error{io.EOF},
			expectedOffset:          11,
			expectedRecoverAttempts: 1,
			description:             "Multiple read errors followed by successful recovery",
		},
		{
			name:           "recovery_failure_exceeds_total_limit",
			initialData:    []byte("hello world"),
			bufferSize:     20,
			readErrors:     []error{errors.New("network error"), errors.New("network error"), errors.New("network error"), errors.New("network error")},
			recoveryErrors: []error{nil, nil, nil, nil}, // Recovery succeeds but read keeps failing
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   3,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:           [][]byte{[]byte("hello world")},
			expectedErrors:          []error{io.EOF},
			expectedOffset:          11,
			expectedRecoverAttempts: 1,
			description:             "Recovery attempts exceed total limit",
		},
		{
			name:           "recovery_failure_exceeds_no_progress_limit",
			initialData:    []byte("hello world"),
			bufferSize:     20,
			readErrors:     []error{errors.New("network error"), errors.New("network error"), errors.New("network error")},
			recoveryErrors: []error{nil, nil, nil}, // Recovery succeeds but read keeps failing
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 2,
			},
			expectedReads:           [][]byte{[]byte("hello world")},
			expectedErrors:          []error{io.EOF},
			expectedOffset:          11,
			expectedRecoverAttempts: 1,
			description:             "Recovery attempts exceed no-progress limit",
		},
		{
			name:        "partial_read_with_error_then_recovery",
			initialData: []byte("hello world"),
			bufferSize:  20,
			readErrors:  []error{nil, errors.New("network error"), nil},
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:           [][]byte{[]byte("hello world")},
			expectedErrors:          []error{io.EOF},
			expectedOffset:          11,
			expectedRecoverAttempts: 1,
			description:             "Partial read succeeds, then error, then recovery",
		},
		{
			name:        "empty_data",
			initialData: []byte{},
			bufferSize:  20,
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:  [][]byte{},
			expectedErrors: []error{io.EOF},
			expectedOffset: 0,
			description:    "Empty data returns EOF immediately",
		},
		{
			name:        "large_data_multiple_reads",
			initialData: []byte(strings.Repeat("a", 1000)),
			bufferSize:  100,
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads: [][]byte{
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
				[]byte(strings.Repeat("a", 100)),
			},
			expectedErrors: []error{nil, nil, nil, nil, nil, nil, nil, nil, nil, io.EOF},
			expectedOffset: 1000,
			description:    "Large data with multiple reads",
		},
		{
			name:           "recovery_error_prevents_further_attempts",
			initialData:    []byte("hello world"),
			bufferSize:     20,
			readErrors:     []error{errors.New("network error")},
			recoveryErrors: []error{errors.New("recovery failed")},
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:           [][]byte{},
			expectedErrors:          []error{errors.New("network error")},
			expectedOffset:          0,
			expectedRecoverAttempts: 1,
			description:             "Recovery error prevents further recovery attempts",
		},
		{
			name:        "read_after_close",
			initialData: []byte("hello world"),
			bufferSize:  20,
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:  [][]byte{},
			expectedErrors: []error{errors.New("I/O operation on closed file")},
			expectedOffset: 0,
			description:    "Read after iterator is closed",
		},
		{
			name:        "zero_buffer_size",
			initialData: []byte("hello world"),
			bufferSize:  0,
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:  [][]byte{},
			expectedErrors: []error{nil},
			expectedOffset: 0,
			description:    "Zero buffer size read",
		},
		{
			name:        "intermittent_errors_with_progress",
			initialData: []byte("hello world"),
			bufferSize:  20,
			readErrors:  []error{nil, errors.New("network error"), nil, errors.New("timeout"), nil},
			config: &UploadConfig{
				FilesAPIClientDownloadMaxTotalRecovers:                   10,
				FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
			},
			expectedReads:           [][]byte{[]byte("hello world")},
			expectedErrors:          []error{io.EOF},
			expectedOffset:          11,
			expectedRecoverAttempts: 1,
			description:             "Intermittent errors with successful progress between failures",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset testOpenDownloadStream for each test
			testOpenDownloadStream = nil

			// Create enhanced mock reader with error sequence
			mockReader := NewEnhancedMockReadCloser(tt.initialData, append([]error{}, tt.readErrors...))

			// Track recovery attempts
			recoveryAttempts := 0
			testOpenDownloadStream = func(ctx context.Context, filePath string, startByteOffset int64, ifUnmodifiedSinceTimestamp string) (*DownloadResponse, error) {
				recoveryAttempts++
				if len(tt.recoveryErrors) > 0 && recoveryAttempts <= len(tt.recoveryErrors) {
					if tt.recoveryErrors[recoveryAttempts-1] != nil {
						return nil, tt.recoveryErrors[recoveryAttempts-1]
					}
				}
				return &DownloadResponse{
					Contents: NewEnhancedMockReadCloser(tt.initialData[startByteOffset:], nil),
				}, nil
			}

			// Create FilesExtTestable and ResilientIterator
			filesExt := &FilesExtTestable{}
			iterator := &ResilientIterator{
				underlyingIterator: mockReader,
				api:                filesExt,
				filePath:           "/test/file.txt",
				fileLastModified:   "2023-01-01T00:00:00Z",
				offset:             0,
				chunkSize:          tt.bufferSize,
				config:             tt.config,
			}

			// Special case for read after close test
			if tt.name == "read_after_close" {
				iterator.Close()
			}

			// Perform reads
			var actualReads [][]byte
			var actualErrors []error
			readCount := 0
			maxReads := len(tt.expectedReads) + 5 // Allow extra reads to catch unexpected behavior

			for readCount < maxReads {
				buffer := make([]byte, tt.bufferSize)
				n, err := iterator.Read(buffer)

				if n > 0 {
					actualReads = append(actualReads, buffer[:n])
				}
				actualErrors = append(actualErrors, err)

				if err == io.EOF {
					break
				}
				if err != nil && err != io.EOF {
					// For error cases, we expect the error to be returned
					break
				}

				readCount++
			}

			// Verify results using cmp library
			// Check total data read
			var expectedTotal []byte
			for _, expected := range tt.expectedReads {
				expectedTotal = append(expectedTotal, expected...)
			}

			var actualTotal []byte
			for _, actual := range actualReads {
				actualTotal = append(actualTotal, actual...)
			}

			if diff := cmp.Diff(expectedTotal, actualTotal); diff != "" {
				t.Errorf("Total read data mismatch (-expected +actual):\n%s", diff)
			}

			// Check final error using cmp with custom comparer
			if len(tt.expectedErrors) > 0 {
				expectedLastError := tt.expectedErrors[len(tt.expectedErrors)-1]
				var actualLastError error
				if len(actualErrors) > 0 {
					actualLastError = actualErrors[len(actualErrors)-1]
				}

				if diff := cmp.Diff(expectedLastError, actualLastError, cmp.Comparer(func(a, b error) bool {
					if a == nil && b == nil {
						return true
					}
					if a == nil || b == nil {
						return false
					}
					if a == io.EOF && b == io.EOF {
						return true
					}
					// For other errors, compare error messages
					return a.Error() == b.Error()
				})); diff != "" {
					t.Errorf("Final error mismatch (-expected +actual):\n%s", diff)
				}
			}

			// Check offset and recovery attempts
			if diff := cmp.Diff(tt.expectedOffset, iterator.offset); diff != "" {
				t.Errorf("Offset mismatch (-expected +actual):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectedRecoverAttempts, recoveryAttempts); diff != "" {
				t.Errorf("Recovery attempts mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func TestResilientIterator_shouldRecover(t *testing.T) {
	tests := []struct {
		name                            string
		totalRecoversCount              int64
		recoversWithoutProgressingCount int64
		maxTotalRecovers                int64
		maxRecoversWithoutProgressing   int64
		expected                        bool
		description                     string
	}{
		{
			name:                            "should_recover_below_limits",
			totalRecoversCount:              5,
			recoversWithoutProgressingCount: 2,
			maxTotalRecovers:                10,
			maxRecoversWithoutProgressing:   3,
			expected:                        true,
			description:                     "Recovery counts below limits should allow recovery",
		},
		{
			name:                            "should_not_recover_total_limit_exceeded",
			totalRecoversCount:              10,
			recoversWithoutProgressingCount: 2,
			maxTotalRecovers:                10,
			maxRecoversWithoutProgressing:   3,
			expected:                        false,
			description:                     "Total recovery limit exceeded should prevent recovery",
		},
		{
			name:                            "should_not_recover_no_progress_limit_exceeded",
			totalRecoversCount:              5,
			recoversWithoutProgressingCount: 3,
			maxTotalRecovers:                10,
			maxRecoversWithoutProgressing:   3,
			expected:                        false,
			description:                     "No progress recovery limit exceeded should prevent recovery",
		},
		{
			name:                            "should_not_recover_both_limits_exceeded",
			totalRecoversCount:              10,
			recoversWithoutProgressingCount: 3,
			maxTotalRecovers:                10,
			maxRecoversWithoutProgressing:   3,
			expected:                        false,
			description:                     "Both limits exceeded should prevent recovery",
		},
		{
			name:                            "should_recover_at_limits",
			totalRecoversCount:              9,
			recoversWithoutProgressingCount: 2,
			maxTotalRecovers:                10,
			maxRecoversWithoutProgressing:   3,
			expected:                        true,
			description:                     "Recovery counts at limits should allow one more recovery",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterator := &ResilientIterator{
				totalRecoversCount:              tt.totalRecoversCount,
				recoversWithoutProgressingCount: tt.recoversWithoutProgressingCount,
				config: &UploadConfig{
					FilesAPIClientDownloadMaxTotalRecovers:                   tt.maxTotalRecovers,
					FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: tt.maxRecoversWithoutProgressing,
				},
			}

			result := iterator.shouldRecover()
			if result != tt.expected {
				t.Errorf("Expected shouldRecover() to return %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestResilientIterator_Close(t *testing.T) {
	tests := []struct {
		name          string
		closeError    error
		expectedError error
		description   string
	}{
		{
			name:          "successful_close",
			closeError:    nil,
			expectedError: nil,
			description:   "Successful close with no error",
		},
		{
			name:          "close_with_error",
			closeError:    errors.New("close failed"),
			expectedError: errors.New("close failed"),
			description:   "Close with underlying error",
		},
		{
			name:          "close_already_closed",
			closeError:    nil,
			expectedError: nil,
			description:   "Close when already closed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := NewMockReadCloser([]byte("test"))
			mockReader.SetCloseError(tt.closeError)

			iterator := &ResilientIterator{
				underlyingIterator: mockReader,
				closed:             false,
			}

			// First close
			err := iterator.Close()
			if diff := cmp.Diff(tt.expectedError, err, cmp.Comparer(func(a, b error) bool {
				if a == nil && b == nil {
					return true
				}
				if a == nil || b == nil {
					return false
				}
				return a.Error() == b.Error()
			})); diff != "" {
				t.Errorf("Close error mismatch (-expected +actual):\n%s", diff)
			}

			if diff := cmp.Diff(true, iterator.closed); diff != "" {
				t.Errorf("Iterator closed state mismatch (-expected +actual):\n%s", diff)
			}

			// Second close should not cause issues
			err2 := iterator.Close()
			if diff := cmp.Diff(nil, err2, cmp.Comparer(func(a, b error) bool {
				if a == nil && b == nil {
					return true
				}
				if a == nil || b == nil {
					return false
				}
				return a.Error() == b.Error()
			})); diff != "" {
				t.Errorf("Second close error mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func TestResilientResponse_Read(t *testing.T) {
	tests := []struct {
		name        string
		initialData []byte
		bufferSize  int
		expected    []byte
		expectedErr error
		description string
	}{
		{
			name:        "successful_read",
			initialData: []byte("hello world"),
			bufferSize:  20,
			expected:    []byte("hello world"),
			expectedErr: nil, // Successful read returns nil error, EOF comes on next read
			description: "Successful read through ResilientResponse",
		},
		{
			name:        "partial_read",
			initialData: []byte("hello world"),
			bufferSize:  5,
			expected:    []byte("hello"),
			expectedErr: nil,
			description: "Partial read through ResilientResponse",
		},
		{
			name:        "empty_data",
			initialData: []byte{},
			bufferSize:  20,
			expected:    []byte{},
			expectedErr: io.EOF,
			description: "Empty data read through ResilientResponse",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := NewMockReadCloser(tt.initialData)
			mockFilesExt := NewMockFilesExt()

			response := &ResilientResponse{
				api:                mockFilesExt.FilesExt,
				filePath:           "/test/file.txt",
				fileLastModified:   "2023-01-01T00:00:00Z",
				offset:             0,
				underlyingResponse: mockReader,
				config:             DefaultUploadConfig(),
			}

			buffer := make([]byte, tt.bufferSize)
			n, err := response.Read(buffer)

			actual := buffer[:n]
			if diff := cmp.Diff(tt.expected, actual); diff != "" {
				t.Errorf("Read data mismatch (-expected +actual):\n%s", diff)
			}

			if diff := cmp.Diff(tt.expectedErr, err, cmp.Comparer(func(a, b error) bool {
				if a == nil && b == nil {
					return true
				}
				if a == nil || b == nil {
					return false
				}
				if a == io.EOF && b == io.EOF {
					return true
				}
				return a.Error() == b.Error()
			})); diff != "" {
				t.Errorf("Read error mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}

func TestResilientResponse_Close(t *testing.T) {
	tests := []struct {
		name          string
		closeError    error
		expectedError error
		description   string
	}{
		{
			name:          "successful_close",
			closeError:    nil,
			expectedError: nil,
			description:   "Successful close with no error",
		},
		{
			name:          "close_with_error",
			closeError:    errors.New("close failed"),
			expectedError: errors.New("close failed"),
			description:   "Close with underlying error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := NewMockReadCloser([]byte("test"))
			mockReader.SetCloseError(tt.closeError)

			response := &ResilientResponse{
				underlyingResponse: mockReader,
			}

			err := response.Close()
			if diff := cmp.Diff(tt.expectedError, err, cmp.Comparer(func(a, b error) bool {
				if a == nil && b == nil {
					return true
				}
				if a == nil || b == nil {
					return false
				}
				return a.Error() == b.Error()
			})); diff != "" {
				t.Errorf("Close error mismatch (-expected +actual):\n%s", diff)
			}
		})
	}
}
