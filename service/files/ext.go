package files

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
)

// FilesExt extends the FilesAPI with enhanced functionality for large file uploads and downloads
type FilesExt struct {
	*FilesAPI
	config *config.Config
}

// NewFilesExt creates a new FilesExt instance
func NewFilesExt(client *client.DatabricksClient) *FilesExt {
	return &FilesExt{
		FilesAPI: NewFiles(client),
		config:   client.Config,
	}
}

// UploadConfig contains configuration for multipart uploads
type UploadConfig struct {
	// Minimum stream size to trigger multipart upload (default: 100MB)
	MultipartUploadMinStreamSize int64
	// Chunk size for multipart uploads (default: 100MB)
	MultipartUploadChunkSize int64
	// Number of upload URLs to request in a batch (default: 10)
	MultipartUploadBatchURLCount int64
	// Maximum number of retries for multipart upload (default: 3)
	MultipartUploadMaxRetries int64
	// Timeout for single chunk upload in seconds (default: 300)
	MultipartUploadSingleChunkUploadTimeoutSeconds int64
	// URL expiration duration (default: 1 hour)
	MultipartUploadURLExpirationDuration time.Duration
	// Maximum total recovers for downloads (default: 10)
	FilesAPIClientDownloadMaxTotalRecovers int64
	// Maximum recovers without progressing for downloads (default: 3)
	FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing int64
}

// DefaultUploadConfig returns default configuration for uploads
func DefaultUploadConfig() *UploadConfig {
	return &UploadConfig{
		MultipartUploadMinStreamSize:                             100 * 1024 * 1024, // 100MB
		MultipartUploadChunkSize:                                 100 * 1024 * 1024, // 100MB
		MultipartUploadBatchURLCount:                             10,
		MultipartUploadMaxRetries:                                3,
		MultipartUploadSingleChunkUploadTimeoutSeconds:           300,
		MultipartUploadURLExpirationDuration:                     time.Hour,
		FilesAPIClientDownloadMaxTotalRecovers:                   10,
		FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing: 3,
	}
}

// Upload uploads a file with enhanced multipart upload support
func (f *FilesExt) Upload(ctx context.Context, request UploadRequest) error {
	config := DefaultUploadConfig()

	// Read a small buffer to determine if we should use multipart upload
	preReadBuffer := make([]byte, config.MultipartUploadMinStreamSize)
	n, err := io.ReadFull(request.Contents, preReadBuffer)
	if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
		return fmt.Errorf("failed to read from input stream: %w", err)
	}

	// If the file is smaller than the minimum size, use one-shot upload
	if n < int(config.MultipartUploadMinStreamSize) {
		logger.Debugf(ctx, "Using one-shot upload for input stream of size %d below %d bytes", n, config.MultipartUploadMinStreamSize)
		return f.FilesAPI.Upload(ctx, UploadRequest{
			FilePath:  request.FilePath,
			Contents:  io.NopCloser(bytes.NewReader(preReadBuffer[:n])),
			Overwrite: request.Overwrite,
		})
	}

	// Initiate multipart upload
	query := map[string]any{"action": "initiate-upload"}
	if request.Overwrite {
		query["overwrite"] = request.Overwrite
	}

	var initiateResponse map[string]any
	err = f.client.Do(ctx, "POST", fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(request.FilePath)),
		nil, query, nil, &initiateResponse)
	if err != nil {
		return fmt.Errorf("failed to initiate upload: %w", err)
	}

	// Create a new reader that includes the pre-read buffer
	combinedReader := io.MultiReader(bytes.NewReader(preReadBuffer[:n]), request.Contents)

	if multipartUpload, ok := initiateResponse["multipart_upload"].(map[string]any); ok {
		sessionToken, ok := multipartUpload["session_token"].(string)
		if !ok {
			return fmt.Errorf("unexpected server response: missing session_token")
		}

		cloudProviderSession := f.createCloudProviderSession()
		err = f.performMultipartUpload(ctx, request.FilePath, combinedReader, sessionToken, preReadBuffer[:n], cloudProviderSession, config)
		if err != nil {
			logger.Infof(ctx, "Aborting multipart upload on error: %v", err)
			abortErr := f.abortMultipartUpload(ctx, request.FilePath, sessionToken, cloudProviderSession, config)
			if abortErr != nil {
				logger.Debugf(ctx, "Failed to abort upload: %v", abortErr)
			}
			return err
		}
	} else if resumableUpload, ok := initiateResponse["resumable_upload"].(map[string]any); ok {
		sessionToken, ok := resumableUpload["session_token"].(string)
		if !ok {
			return fmt.Errorf("unexpected server response: missing session_token")
		}

		cloudProviderSession := f.createCloudProviderSession()
		err = f.performResumableUpload(ctx, request.FilePath, combinedReader, sessionToken, request.Overwrite, preReadBuffer[:n], cloudProviderSession, config)
		if err != nil {
			logger.Infof(ctx, "Aborting resumable upload on error: %v", err)
			// Note: Resumable upload abort is handled differently
			return err
		}
	} else {
		return fmt.Errorf("unexpected server response: %v", initiateResponse)
	}

	return nil
}

// Download downloads a file with enhanced resilient download support
func (f *FilesExt) Download(ctx context.Context, request DownloadRequest) (*DownloadResponse, error) {
	initialResponse, err := f.openDownloadStream(ctx, request.FilePath, 0, "")
	if err != nil {
		return nil, err
	}

	wrappedResponse := f.wrapStream(ctx, request.FilePath, initialResponse)
	initialResponse.Contents = wrappedResponse
	return initialResponse, nil
}

// performMultipartUpload performs multipart upload using presigned URLs
func (f *FilesExt) performMultipartUpload(ctx context.Context, targetPath string, inputStream io.Reader, sessionToken string, preReadBuffer []byte, cloudProviderSession *http.Client, config *UploadConfig) error {
	currentPartNumber := int64(1)
	etags := make(map[int64]string)
	buffer := preReadBuffer
	chunkOffset := int64(0)
	retryCount := 0

	for {
		// Fill buffer if needed
		buffer = f.fillBuffer(buffer, config.MultipartUploadChunkSize, inputStream)
		if len(buffer) == 0 {
			break
		}

		logger.Debugf(ctx, "Multipart upload: requesting next %d upload URLs starting from part %d", config.MultipartUploadBatchURLCount, currentPartNumber)

		body := map[string]any{
			"path":              targetPath,
			"session_token":     sessionToken,
			"start_part_number": currentPartNumber,
			"count":             config.MultipartUploadBatchURLCount,
			"expire_time":       f.getURLExpireTime(config),
		}

		var uploadPartURLsResponse map[string]any
		err := f.client.Do(ctx, "POST", "/api/2.0/fs/create-upload-part-urls",
			map[string]string{"Content-Type": "application/json"}, nil, body, &uploadPartURLsResponse)
		if err != nil {
			return fmt.Errorf("failed to get upload part URLs: %w", err)
		}

		uploadPartURLs, ok := uploadPartURLsResponse["upload_part_urls"].([]any)
		if !ok || len(uploadPartURLs) == 0 {
			return fmt.Errorf("unexpected server response: %v", uploadPartURLsResponse)
		}

		for _, uploadPartURL := range uploadPartURLs {
			urlData, ok := uploadPartURL.(map[string]any)
			if !ok {
				return fmt.Errorf("invalid upload part URL data")
			}

			buffer = f.fillBuffer(buffer, config.MultipartUploadChunkSize, inputStream)
			actualBufferLength := len(buffer)
			if actualBufferLength == 0 {
				break
			}

			url, ok := urlData["url"].(string)
			if !ok {
				return fmt.Errorf("invalid upload URL")
			}

			partNumber, ok := urlData["part_number"].(float64)
			if !ok || int64(partNumber) != currentPartNumber {
				return fmt.Errorf("invalid part number")
			}

			requiredHeaders, _ := urlData["headers"].([]any)
			headers := map[string]string{"Content-Type": "application/octet-stream"}
			for _, h := range requiredHeaders {
				if headerData, ok := h.(map[string]any); ok {
					if name, ok := headerData["name"].(string); ok {
						if value, ok := headerData["value"].(string); ok {
							headers[name] = value
						}
					}
				}
			}

			actualChunkLength := min(actualBufferLength, int(config.MultipartUploadChunkSize))
			logger.Debugf(ctx, "Uploading part %d: [%d, %d]", currentPartNumber, chunkOffset, chunkOffset+int64(actualChunkLength)-1)

			chunk := bytes.NewReader(buffer[:actualChunkLength])

			uploadResponse, err := f.retryIdempotentOperation(ctx, func() (*http.Response, error) {
				req, err := http.NewRequestWithContext(ctx, "PUT", url, chunk)
				if err != nil {
					return nil, err
				}
				for k, v := range headers {
					req.Header.Set(k, v)
				}
				return cloudProviderSession.Do(req)
			}, func() {
				chunk.Seek(0, 0)
			}, config)

			if err != nil {
				return fmt.Errorf("failed to upload part %d: %w", currentPartNumber, err)
			}

			if uploadResponse.StatusCode == 200 || uploadResponse.StatusCode == 201 {
				chunkOffset += int64(actualChunkLength)
				etag := uploadResponse.Header.Get("ETag")
				etags[currentPartNumber] = etag
				buffer = buffer[actualChunkLength:]
				retryCount = 0
			} else if f.isURLExpiredResponse(uploadResponse) {
				if retryCount < int(config.MultipartUploadMaxRetries) {
					retryCount++
					logger.Debugf(ctx, "Upload URL expired, retrying")
					continue
				} else {
					return fmt.Errorf("upload URL expired after %d retries", config.MultipartUploadMaxRetries)
				}
			} else {
				return fmt.Errorf("unsuccessful chunk upload. Response status: %d", uploadResponse.StatusCode)
			}

			currentPartNumber++
		}
	}

	logger.Debugf(ctx, "Completing multipart upload after uploading %d parts of up to %d bytes", len(etags), config.MultipartUploadChunkSize)

	// Complete the upload
	parts := make([]map[string]any, 0, len(etags))
	for partNumber, etag := range etags {
		parts = append(parts, map[string]any{
			"part_number": partNumber,
			"etag":        etag,
		})
	}

	body := map[string]any{"parts": parts}
	query := map[string]any{
		"action":        "complete-upload",
		"upload_type":   "multipart",
		"session_token": sessionToken,
	}

	err := f.client.Do(ctx, "POST", fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(targetPath)),
		map[string]string{"Content-Type": "application/json"}, query, body, nil)
	if err != nil {
		return fmt.Errorf("failed to complete multipart upload: %w", err)
	}

	return nil
}

// performResumableUpload performs resumable upload (GCP)
func (f *FilesExt) performResumableUpload(ctx context.Context, targetPath string, inputStream io.Reader, sessionToken string, overwrite bool, preReadBuffer []byte, cloudProviderSession *http.Client, config *UploadConfig) error {
	body := map[string]any{
		"path":          targetPath,
		"session_token": sessionToken,
	}

	var resumableUploadURLResponse map[string]any
	err := f.client.Do(ctx, "POST", "/api/2.0/fs/create-resumable-upload-url",
		map[string]string{"Content-Type": "application/json"}, nil, body, &resumableUploadURLResponse)
	if err != nil {
		return fmt.Errorf("failed to create resumable upload URL: %w", err)
	}

	resumableUploadURLNode, ok := resumableUploadURLResponse["resumable_upload_url"].(map[string]any)
	if !ok {
		return fmt.Errorf("unexpected server response: %v", resumableUploadURLResponse)
	}

	resumableUploadURL, ok := resumableUploadURLNode["url"].(string)
	if !ok {
		return fmt.Errorf("unexpected server response: %v", resumableUploadURLResponse)
	}

	requiredHeaders, _ := resumableUploadURLNode["headers"].([]any)
	headers := make(map[string]string)
	for _, h := range requiredHeaders {
		if headerData, ok := h.(map[string]any); ok {
			if name, ok := headerData["name"].(string); ok {
				if value, ok := headerData["value"].(string); ok {
					headers[name] = value
				}
			}
		}
	}

	// Buffer for one chunk + read-ahead block
	minBufferSize := config.MultipartUploadChunkSize + 1
	buffer := preReadBuffer
	uploadedBytesCount := int64(0)
	chunkOffset := int64(0)

	for {
		// Fill buffer if needed
		bytesToRead := max(0, minBufferSize-(int64(len(buffer))-uploadedBytesCount))
		nextBuf := make([]byte, bytesToRead)
		n, err := io.ReadFull(inputStream, nextBuf)
		if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
			return fmt.Errorf("failed to read from input stream: %w", err)
		}
		buffer = append(buffer[uploadedBytesCount:], nextBuf[:n]...)

		if int64(n) < bytesToRead {
			// This is the last chunk
			actualChunkLength := len(buffer)
			fileSize := chunkOffset + int64(actualChunkLength)
			contentRangeHeader := fmt.Sprintf("bytes %d-%d/%d", chunkOffset, chunkOffset+int64(actualChunkLength)-1, fileSize)
			logger.Debugf(ctx, "Uploading final chunk: %s", contentRangeHeader)

			uploadHeaders := map[string]string{"Content-Type": "application/octet-stream"}
			for k, v := range headers {
				uploadHeaders[k] = v
			}
			uploadHeaders["Content-Range"] = contentRangeHeader

			req, err := http.NewRequestWithContext(ctx, "PUT", resumableUploadURL, bytes.NewReader(buffer[:actualChunkLength]))
			if err != nil {
				return fmt.Errorf("failed to create upload request: %w", err)
			}
			for k, v := range uploadHeaders {
				req.Header.Set(k, v)
			}

			uploadResponse, err := cloudProviderSession.Do(req)
			if err != nil {
				return fmt.Errorf("failed to upload final chunk: %w", err)
			}

			if uploadResponse.StatusCode == 200 || uploadResponse.StatusCode == 201 {
				break // Upload complete
			} else {
				return fmt.Errorf("unsuccessful final chunk upload. Response status: %d", uploadResponse.StatusCode)
			}
		} else {
			// More chunks expected
			actualChunkLength := config.MultipartUploadChunkSize
			contentRangeHeader := fmt.Sprintf("bytes %d-%d/*", chunkOffset, chunkOffset+actualChunkLength-1)
			logger.Debugf(ctx, "Uploading chunk: %s", contentRangeHeader)

			uploadHeaders := map[string]string{"Content-Type": "application/octet-stream"}
			for k, v := range headers {
				uploadHeaders[k] = v
			}
			uploadHeaders["Content-Range"] = contentRangeHeader

			req, err := http.NewRequestWithContext(ctx, "PUT", resumableUploadURL, bytes.NewReader(buffer[:actualChunkLength]))
			if err != nil {
				return fmt.Errorf("failed to create upload request: %w", err)
			}
			for k, v := range uploadHeaders {
				req.Header.Set(k, v)
			}

			uploadResponse, err := cloudProviderSession.Do(req)
			if err != nil {
				return fmt.Errorf("failed to upload chunk: %w", err)
			}

			if uploadResponse.StatusCode == 308 {
				// Chunk accepted, determine received offset
				rangeString := uploadResponse.Header.Get("Range")
				confirmedOffset := f.extractRangeOffset(rangeString)
				logger.Debugf(ctx, "Received confirmed offset: %d", confirmedOffset)

				if confirmedOffset != nil {
					if *confirmedOffset < chunkOffset-1 || *confirmedOffset > chunkOffset+actualChunkLength-1 {
						return fmt.Errorf("unexpected received offset: %d is outside of expected range", *confirmedOffset)
					}
					nextChunkOffset := *confirmedOffset + 1
					uploadedBytesCount = nextChunkOffset - chunkOffset
					chunkOffset = nextChunkOffset
				} else {
					if chunkOffset > 0 {
						return fmt.Errorf("unexpected received offset: %v is outside of expected range", confirmedOffset)
					}
					uploadedBytesCount = 0
				}
			} else if uploadResponse.StatusCode == 412 && !overwrite {
				return fmt.Errorf("the file being created already exists")
			} else {
				return fmt.Errorf("unsuccessful chunk upload. Response status: %d", uploadResponse.StatusCode)
			}
		}
	}

	return nil
}

// openDownloadStream opens a download stream from given offset
func (f *FilesExt) openDownloadStream(ctx context.Context, filePath string, startByteOffset int64, ifUnmodifiedSinceTimestamp string) (*DownloadResponse, error) {
	headers := map[string]string{"Accept": "application/octet-stream"}

	if startByteOffset > 0 && ifUnmodifiedSinceTimestamp == "" {
		return nil, fmt.Errorf("if_unmodified_since_timestamp is required if start_byte_offset is specified")
	}

	if startByteOffset > 0 {
		headers["Range"] = fmt.Sprintf("bytes=%d-", startByteOffset)
	}

	if ifUnmodifiedSinceTimestamp != "" {
		headers["If-Unmodified-Since"] = ifUnmodifiedSinceTimestamp
	}

	var response DownloadResponse
	err := f.client.Do(ctx, "GET", fmt.Sprintf("/api/2.0/fs/files%v", httpclient.EncodeMultiSegmentPathParameter(filePath)),
		headers, nil, nil, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// wrapStream wraps the download response with resilient functionality
func (f *FilesExt) wrapStream(ctx context.Context, filePath string, downloadResponse *DownloadResponse) io.ReadCloser {
	return &ResilientResponse{
		api:                f,
		filePath:           filePath,
		fileLastModified:   downloadResponse.LastModified,
		offset:             0,
		underlyingResponse: downloadResponse.Contents,
	}
}

// fillBuffer tries to fill the given buffer to contain at least desiredMinSize bytes
func (f *FilesExt) fillBuffer(buffer []byte, desiredMinSize int64, inputStream io.Reader) []byte {
	bytesToRead := max(0, desiredMinSize-int64(len(buffer)))
	if bytesToRead > 0 {
		nextBuf := make([]byte, bytesToRead)
		n, err := io.ReadFull(inputStream, nextBuf)
		if err != nil && err != io.ErrUnexpectedEOF && err != io.EOF {
			return buffer
		}
		return append(buffer, nextBuf[:n]...)
	}
	return buffer
}

// isURLExpiredResponse checks if response matches known "URL expired" responses
func (f *FilesExt) isURLExpiredResponse(response *http.Response) bool {
	if response.StatusCode != 403 {
		return false
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return false
	}

	var xmlRoot struct {
		XMLName xml.Name `xml:"Error"`
		Code    string   `xml:"Code"`
		Message string   `xml:"Message"`
		Details string   `xml:"AuthenticationErrorDetail"`
	}

	err = xml.Unmarshal(body, &xmlRoot)
	if err != nil {
		return false
	}

	if xmlRoot.Code == "AuthenticationFailed" {
		// Azure
		if strings.Contains(xmlRoot.Details, "Signature not valid in the specified time frame") {
			return true
		}
	}

	if xmlRoot.Code == "AccessDenied" {
		// AWS
		if xmlRoot.Message == "Request has expired" {
			return true
		}
	}

	return false
}

// extractRangeOffset parses the response range header to extract the last byte
func (f *FilesExt) extractRangeOffset(rangeString string) *int64 {
	if rangeString == "" {
		return nil
	}

	re := regexp.MustCompile(`bytes=0-(\d+)`)
	match := re.FindStringSubmatch(rangeString)
	if len(match) == 2 {
		if offset, err := strconv.ParseInt(match[1], 10, 64); err == nil {
			return &offset
		}
	}

	return nil
}

// getURLExpireTime generates expiration time in the required format
func (f *FilesExt) getURLExpireTime(config *UploadConfig) string {
	expireTime := time.Now().UTC().Add(config.MultipartUploadURLExpirationDuration)
	return expireTime.Format("2006-01-02T15:04:05Z")
}

// abortMultipartUpload aborts ongoing multipart upload session
func (f *FilesExt) abortMultipartUpload(ctx context.Context, targetPath string, sessionToken string, cloudProviderSession *http.Client, config *UploadConfig) error {
	body := map[string]any{
		"path":          targetPath,
		"session_token": sessionToken,
		"expire_time":   f.getURLExpireTime(config),
	}

	var abortURLResponse map[string]any
	err := f.client.Do(ctx, "POST", "/api/2.0/fs/create-abort-upload-url",
		map[string]string{"Content-Type": "application/json"}, nil, body, &abortURLResponse)
	if err != nil {
		return fmt.Errorf("failed to create abort upload URL: %w", err)
	}

	abortUploadURLNode, ok := abortURLResponse["abort_upload_url"].(map[string]any)
	if !ok {
		return fmt.Errorf("unexpected server response: %v", abortURLResponse)
	}

	abortURL, ok := abortUploadURLNode["url"].(string)
	if !ok {
		return fmt.Errorf("unexpected server response: %v", abortURLResponse)
	}

	requiredHeaders, _ := abortUploadURLNode["headers"].([]any)
	headers := map[string]string{"Content-Type": "application/octet-stream"}
	for _, h := range requiredHeaders {
		if headerData, ok := h.(map[string]any); ok {
			if name, ok := headerData["name"].(string); ok {
				if value, ok := headerData["value"].(string); ok {
					headers[name] = value
				}
			}
		}
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", abortURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create abort request: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	abortResponse, err := cloudProviderSession.Do(req)
	if err != nil {
		return fmt.Errorf("failed to abort upload: %w", err)
	}

	if abortResponse.StatusCode != 200 && abortResponse.StatusCode != 201 {
		return fmt.Errorf("failed to abort upload: status %d", abortResponse.StatusCode)
	}

	return nil
}

// createCloudProviderSession creates a separate session for cloud provider requests
func (f *FilesExt) createCloudProviderSession() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     180 * time.Second,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   time.Duration(DefaultUploadConfig().MultipartUploadSingleChunkUploadTimeoutSeconds) * time.Second,
	}
}

// retryIdempotentOperation performs given idempotent operation with necessary retries
func (f *FilesExt) retryIdempotentOperation(ctx context.Context, operation func() (*http.Response, error), beforeRetry func(), config *UploadConfig) (*http.Response, error) {
	retryableStatusCodes := []int{408, 429, 500, 502, 503, 504}

	var lastResponse *http.Response
	var lastErr error

	for attempt := 0; attempt <= int(config.MultipartUploadMaxRetries); attempt++ {
		if attempt > 0 && beforeRetry != nil {
			beforeRetry()
		}

		response, err := operation()
		if err == nil {
			lastResponse = response
			lastErr = nil
			break
		}

		lastErr = err

		// Check if error is retryable
		isRetryable := false
		if response != nil {
			for _, code := range retryableStatusCodes {
				if response.StatusCode == code {
					isRetryable = true
					break
				}
			}
		}

		if !isRetryable {
			break
		}

		if attempt < int(config.MultipartUploadMaxRetries) {
			// Wait before retry
			time.Sleep(time.Duration(attempt+1) * time.Second)
		}
	}

	if lastErr != nil {
		return nil, lastErr
	}

	return lastResponse, nil
}

// ResilientResponse wraps the underlying response with resilient functionality
type ResilientResponse struct {
	api                *FilesExt
	filePath           string
	fileLastModified   string
	offset             int64
	underlyingResponse io.ReadCloser
	iterator           *ResilientIterator
}

func (r *ResilientResponse) Read(p []byte) (int, error) {
	if r.iterator == nil {
		r.iterator = &ResilientIterator{
			underlyingIterator: r.underlyingResponse,
			api:                r.api,
			filePath:           r.filePath,
			fileLastModified:   r.fileLastModified,
			offset:             r.offset,
			chunkSize:          len(p),
			config:             DefaultUploadConfig(),
		}
	}

	return r.iterator.Read(p)
}

func (r *ResilientResponse) Close() error {
	if r.iterator != nil {
		r.iterator.Close()
	}
	if r.underlyingResponse != nil {
		return r.underlyingResponse.Close()
	}
	return nil
}

// ResilientIterator provides resilient iteration over the response content
type ResilientIterator struct {
	underlyingIterator io.ReadCloser
	api                *FilesExt
	filePath           string
	fileLastModified   string
	offset             int64
	chunkSize          int
	config             *UploadConfig

	totalRecoversCount              int64
	recoversWithoutProgressingCount int64
	closed                          bool
}

func (r *ResilientIterator) Read(p []byte) (int, error) {
	if r.closed {
		return 0, fmt.Errorf("I/O operation on closed file")
	}

	for {
		n, err := r.underlyingIterator.Read(p)
		if err == nil {
			r.offset += int64(n)
			r.recoversWithoutProgressingCount = 0
			return n, nil
		}

		if err == io.EOF {
			return n, err
		}

		// Try to recover from the error
		if !r.shouldRecover() {
			return n, err
		}

		if !r.recover() {
			return n, err
		}
	}
}

func (r *ResilientIterator) shouldRecover() bool {
	if r.totalRecoversCount >= r.config.FilesAPIClientDownloadMaxTotalRecovers {
		logger.Debugf(context.Background(), "Total recovers limit exceeded")
		return false
	}
	if r.recoversWithoutProgressingCount >= r.config.FilesAPIClientDownloadMaxTotalRecoversWithoutProgressing {
		logger.Debugf(context.Background(), "No progression recovers limit exceeded")
		return false
	}
	return true
}

func (r *ResilientIterator) recover() bool {
	if !r.shouldRecover() {
		return false
	}

	r.totalRecoversCount++
	r.recoversWithoutProgressingCount++

	if r.underlyingIterator != nil {
		r.underlyingIterator.Close()
	}

	logger.Debugf(context.Background(), "Trying to recover from offset %d", r.offset)

	downloadResponse, err := r.api.openDownloadStream(context.Background(), r.filePath, r.offset, r.fileLastModified)
	if err != nil {
		return false
	}

	r.underlyingIterator = downloadResponse.Contents
	logger.Debugf(context.Background(), "Recover succeeded")
	return true
}

func (r *ResilientIterator) Close() error {
	if r.underlyingIterator != nil {
		return r.underlyingIterator.Close()
	}
	r.closed = true
	return nil
}

// Helper functions
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
