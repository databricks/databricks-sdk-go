package files

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/useragent"
)

const (
	// minMultipartUploadSize is the minimum file size (in bytes) to trigger multipart upload.
	minMultipartUploadSize = 50 * 1024 * 1024 // 50 MiB

	// defaultPartSize is the default part size for multipart uploads.
	defaultPartSize = 10 * 1024 * 1024 // 10 MiB

	// maxPartSize is the maximum part size for multipart uploads (Azure limit).
	maxPartSize = 4 * 1024 * 1024 * 1024 // 4 GiB

	// defaultParallelism is the default number of concurrent upload workers.
	defaultParallelism = 10

	// maxPartsTarget is the target maximum number of parts for an upload.
	maxPartsTarget = 100

	// urlExpirationDuration is how long presigned URLs are valid.
	urlExpirationDuration = 1 * time.Hour

	// maxURLExpirationRetries is the maximum number of retries for URL expiration.
	maxURLExpirationRetries = 3
)

// partSizeOptions lists the candidate part sizes in ascending order.
var partSizeOptions = []int64{
	10 * 1024 * 1024,       // 10 MiB
	20 * 1024 * 1024,       // 20 MiB
	50 * 1024 * 1024,       // 50 MiB
	100 * 1024 * 1024,      // 100 MiB
	200 * 1024 * 1024,      // 200 MiB
	500 * 1024 * 1024,      // 500 MiB
	1 * 1024 * 1024 * 1024, // 1 GiB
	2 * 1024 * 1024 * 1024, // 2 GiB
	4 * 1024 * 1024 * 1024, // 4 GiB
}

// UploadConfig holds configuration for a multipart upload.
type UploadConfig struct {
	PartSize    int64
	Parallelism int
	Overwrite   bool
}

// UploadOption is a functional option for configuring a multipart upload.
type UploadOption func(*UploadConfig)

// WithPartSize sets the part size for a multipart upload.
func WithPartSize(partSize int64) UploadOption {
	return func(c *UploadConfig) {
		c.PartSize = partSize
	}
}

// WithParallelism sets the number of concurrent upload workers.
func WithParallelism(parallelism int) UploadOption {
	return func(c *UploadConfig) {
		c.Parallelism = parallelism
	}
}

// WithOverwrite sets whether to overwrite an existing file.
func WithOverwrite(overwrite bool) UploadOption {
	return func(c *UploadConfig) {
		c.Overwrite = overwrite
	}
}

// initiateUploadResponse is the response from initiating a multipart upload.
type initiateUploadResponse struct {
	MultipartUpload *multipartUploadSession `json:"multipart_upload,omitempty"`
	ResumableUpload *resumableUploadSession `json:"resumable_upload,omitempty"`
}

// multipartUploadSession holds the state for a multipart upload session.
type multipartUploadSession struct {
	SessionToken string `json:"session_token"`
}

// resumableUploadSession holds the state for a resumable upload session.
type resumableUploadSession struct {
	SessionToken string `json:"session_token"`
}

// presignedURL represents a presigned URL for uploading a part.
type presignedURL struct {
	URL        string            `json:"url"`
	PartNumber int               `json:"part_number"`
	Headers    []presignedHeader `json:"headers"`
}

// presignedHeader is a header to include when uploading to a presigned URL.
type presignedHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// createUploadPartURLsRequest is the request to create presigned URLs for upload parts.
type createUploadPartURLsRequest struct {
	Path            string `json:"path"`
	SessionToken    string `json:"session_token"`
	StartPartNumber int    `json:"start_part_number"`
	Count           int    `json:"count"`
	ExpireTime      string `json:"expire_time"`
}

// createUploadPartURLsResponse is the response containing presigned URLs for upload parts.
type createUploadPartURLsResponse struct {
	UploadPartURLs []presignedURL `json:"upload_part_urls"`
}

// completeUploadPart represents a completed upload part.
type completeUploadPart struct {
	PartNumber int    `json:"part_number"`
	ETag       string `json:"etag"`
}

// completeUploadRequest is the request to complete a multipart upload.
type completeUploadRequest struct {
	Parts []completeUploadPart `json:"parts"`
}

// optimizePartSize selects the best part size and batch size for a multipart upload.
//
// If explicitPartSize > 0, it is used directly and the batch size is computed
// as max(1, ceil(sqrt(numParts))).
//
// If contentLength <= 0 (unknown), defaultPartSize and batch size 1 are returned.
//
// Otherwise, the smallest part size from partSizeOptions where the number of
// parts is <= maxPartsTarget is selected. If no option satisfies the constraint,
// maxPartSize is used as a fallback.
func optimizePartSize(contentLength int64, explicitPartSize int64) (int64, int) {
	if explicitPartSize > 0 {
		numParts := int(math.Ceil(float64(contentLength) / float64(explicitPartSize)))
		if numParts < 1 {
			numParts = 1
		}
		batchSize := int(math.Ceil(math.Sqrt(float64(numParts))))
		if batchSize < 1 {
			batchSize = 1
		}
		return explicitPartSize, batchSize
	}

	if contentLength <= 0 {
		return defaultPartSize, 1
	}

	for _, partSize := range partSizeOptions {
		numParts := int(math.Ceil(float64(contentLength) / float64(partSize)))
		if numParts <= maxPartsTarget {
			batchSize := int(math.Ceil(math.Sqrt(float64(numParts))))
			if batchSize < 1 {
				batchSize = 1
			}
			return partSize, batchSize
		}
	}

	// Fallback to maxPartSize
	numParts := int(math.Ceil(float64(contentLength) / float64(maxPartSize)))
	if numParts < 1 {
		numParts = 1
	}
	batchSize := int(math.Ceil(math.Sqrt(float64(numParts))))
	if batchSize < 1 {
		batchSize = 1
	}
	return maxPartSize, batchSize
}

// uploadURLExpireTime returns the expiration time for a presigned URL as an RFC3339 string.
func uploadURLExpireTime() string {
	return time.Now().Add(urlExpirationDuration).UTC().Format(time.RFC3339)
}

// initiateMultipartUpload starts a multipart upload session for the given file path.
func (a *FilesAPI) initiateMultipartUpload(ctx context.Context, filePath string, overwrite bool) (*initiateUploadResponse, error) {
	var resp initiateUploadResponse
	apiPath := fmt.Sprintf("/api/2.0/fs/files%s", httpclient.EncodeMultiSegmentPathParameter(filePath))
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	queryParams := make(map[string]any)
	queryParams["action"] = "initiate-upload"
	if overwrite {
		queryParams["overwrite"] = "true"
	}
	err := a.filesImpl.client.Do(ctx, http.MethodPost, apiPath, headers, queryParams, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// getUploadPartURLs fetches presigned URLs for uploading parts.
func (a *FilesAPI) getUploadPartURLs(ctx context.Context, filePath, sessionToken string, startPartNumber, count int) ([]presignedURL, error) {
	var resp createUploadPartURLsResponse
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	queryParams := make(map[string]any)
	req := createUploadPartURLsRequest{
		Path:            filePath,
		SessionToken:    sessionToken,
		StartPartNumber: startPartNumber,
		Count:           count,
		ExpireTime:      uploadURLExpireTime(),
	}
	err := a.filesImpl.client.Do(ctx, http.MethodPost, "/api/2.0/fs/create-upload-part-urls", headers, queryParams, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.UploadPartURLs, nil
}

// isURLExpiredResponse returns true if the HTTP response indicates the presigned URL has expired.
func isURLExpiredResponse(statusCode int, body []byte) bool {
	if statusCode != http.StatusForbidden {
		return false
	}
	s := string(body)
	return strings.Contains(s, "AccessDenied") || strings.Contains(s, "AuthenticationFailed")
}

// uploadOnePart uploads a single part to a presigned URL and returns the ETag.
func (a *FilesAPI) uploadOnePart(ctx context.Context, presigned presignedURL, partData io.ReadSeeker, contentLength int64) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, presigned.URL, partData)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.ContentLength = contentLength
	req.Header.Set("Content-Type", "application/octet-stream")
	for _, h := range presigned.Headers {
		req.Header.Set(h.Name, h.Value)
	}

	// Use a client without a total timeout — context cancellation handles
	// overall deadline, and large parts may legitimately take a long time.
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("upload part %d failed: %w", presigned.PartNumber, err)
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		etag := resp.Header.Get("ETag")
		resp.Body.Close()
		return etag, nil
	}

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	return "", &partUploadError{
		partNumber: presigned.PartNumber,
		statusCode: resp.StatusCode,
		body:       body,
	}
}

// partUploadError represents a failed part upload with status and body.
type partUploadError struct {
	partNumber int
	statusCode int
	body       []byte
}

func (e *partUploadError) Error() string {
	return fmt.Sprintf("upload part %d failed with status %d: %s", e.partNumber, e.statusCode, string(e.body))
}

func (e *partUploadError) isExpiredURL() bool {
	return isURLExpiredResponse(e.statusCode, e.body)
}

// uploadOnePartWithRetry uploads a part, re-fetching the presigned URL if it expires.
func (a *FilesAPI) uploadOnePartWithRetry(ctx context.Context, filePath, sessionToken string, partNumber int, partData io.ReadSeeker, contentLength int64) (string, error) {
	for attempt := 0; attempt <= maxURLExpirationRetries; attempt++ {
		// Rewind reader for retries.
		if attempt > 0 {
			if _, err := partData.Seek(0, io.SeekStart); err != nil {
				return "", fmt.Errorf("failed to rewind part data: %w", err)
			}
		}

		// Fetch a (fresh) presigned URL.
		urls, err := a.getUploadPartURLs(ctx, filePath, sessionToken, partNumber, 1)
		if err != nil {
			return "", fmt.Errorf("failed to get upload URL for part %d: %w", partNumber, err)
		}
		if len(urls) == 0 {
			return "", fmt.Errorf("no upload URL returned for part %d", partNumber)
		}

		etag, err := a.uploadOnePart(ctx, urls[0], partData, contentLength)
		if err != nil {
			var pErr *partUploadError
			if errors.As(err, &pErr) && pErr.isExpiredURL() && attempt < maxURLExpirationRetries {
				logger.Debugf(ctx, "presigned URL expired for part %d (attempt %d/%d), fetching new URL",
					partNumber, attempt+1, maxURLExpirationRetries)
				continue
			}
			return "", err
		}
		return etag, nil
	}

	return "", fmt.Errorf("upload part %d: presigned URL expired after %d retries", partNumber, maxURLExpirationRetries)
}

// completeMultipartUpload finalizes a multipart upload with the given ETags.
func (a *FilesAPI) completeMultipartUpload(ctx context.Context, filePath, sessionToken string, etags map[int]string) error {
	apiPath := fmt.Sprintf("/api/2.0/fs/files%s", httpclient.EncodeMultiSegmentPathParameter(filePath))
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	queryParams := make(map[string]any)
	queryParams["action"] = "complete-upload"
	queryParams["upload_type"] = "multipart"
	queryParams["session_token"] = sessionToken

	parts := make([]completeUploadPart, 0, len(etags))
	for partNum, etag := range etags {
		parts = append(parts, completeUploadPart{
			PartNumber: partNum,
			ETag:       etag,
		})
	}
	sort.Slice(parts, func(i, j int) bool {
		return parts[i].PartNumber < parts[j].PartNumber
	})

	req := completeUploadRequest{Parts: parts}
	return a.filesImpl.client.Do(ctx, http.MethodPost, apiPath, headers, queryParams, req, nil)
}

// uploadMultipart orchestrates a full multipart upload.
func (a *FilesAPI) uploadMultipart(ctx context.Context, filePath string, content io.ReadSeeker, _ int64, cfg *UploadConfig) error {
	// Phase 1: Initiate
	initResp, err := a.initiateMultipartUpload(ctx, filePath, cfg.Overwrite)
	if err != nil {
		return fmt.Errorf("failed to initiate multipart upload: %w", err)
	}
	if initResp.MultipartUpload == nil {
		return fmt.Errorf("multipart upload not supported for this path (GCP is not supported)")
	}
	sessionToken := initResp.MultipartUpload.SessionToken
	logger.Debugf(ctx, "initiated multipart upload with session token")

	partSize := cfg.PartSize
	parallelism := cfg.Parallelism
	if parallelism < 1 {
		parallelism = defaultParallelism
	}

	// Phase 2: Upload parts in parallel
	sem := make(chan struct{}, parallelism)
	var mu sync.Mutex
	etags := make(map[int]string)
	var uploadErr error

	partNumber := 1
	for {
		// Check context cancellation and prior upload errors.
		if err := ctx.Err(); err != nil {
			break
		}
		mu.Lock()
		if uploadErr != nil {
			mu.Unlock()
			break
		}
		mu.Unlock()

		// Read the next part.
		buf := make([]byte, partSize)
		n, readErr := io.ReadFull(content, buf)
		if n == 0 && readErr != nil {
			break
		}
		buf = buf[:n]
		partDataLen := int64(n)

		currentPartNumber := partNumber
		partNumber++

		sem <- struct{}{}
		go func(pn int, data []byte) {
			defer func() { <-sem }()

			// Check for prior error or context cancellation.
			mu.Lock()
			if uploadErr != nil {
				mu.Unlock()
				return
			}
			mu.Unlock()
			if ctx.Err() != nil {
				return
			}

			// Upload the part with automatic URL refresh on expiration.
			etag, err := a.uploadOnePartWithRetry(ctx, filePath, sessionToken, pn, bytes.NewReader(data), partDataLen)
			if err != nil {
				mu.Lock()
				if uploadErr == nil {
					uploadErr = err
				}
				mu.Unlock()
				return
			}

			mu.Lock()
			etags[pn] = etag
			mu.Unlock()

			logger.Debugf(ctx, "uploaded part %d", pn)
		}(currentPartNumber, buf)

		if readErr != nil {
			break
		}
	}

	// Wait for all goroutines to finish
	for i := 0; i < parallelism; i++ {
		sem <- struct{}{}
	}

	if uploadErr != nil {
		a.abortMultipartUpload(ctx, filePath, sessionToken)
		return uploadErr
	}

	// Phase 3: Complete
	logger.Debugf(ctx, "completing multipart upload with %d parts", len(etags))
	return a.completeMultipartUpload(ctx, filePath, sessionToken, etags)
}

// abortMultipartUpload attempts to abort an in-progress multipart upload.
// This is a best-effort cleanup; errors are logged but not returned.
func (a *FilesAPI) abortMultipartUpload(ctx context.Context, filePath, sessionToken string) {
	apiPath := "/api/2.0/fs/create-abort-upload-url"
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	body := map[string]string{
		"path":          filePath,
		"session_token": sessionToken,
		"expire_time":   uploadURLExpireTime(),
	}

	var resp struct {
		AbortUploadURL struct {
			URL     string            `json:"url"`
			Headers []presignedHeader `json:"headers,omitempty"`
		} `json:"abort_upload_url"`
	}
	err := a.filesImpl.client.Do(ctx, http.MethodPost, apiPath, headers, nil, body, &resp)
	if err != nil {
		logger.Debugf(ctx, "failed to get abort URL: %v", err)
		return
	}
	if resp.AbortUploadURL.URL == "" {
		logger.Debugf(ctx, "no abort URL returned")
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, resp.AbortUploadURL.URL, nil)
	if err != nil {
		logger.Debugf(ctx, "failed to create abort request: %v", err)
		return
	}
	for _, h := range resp.AbortUploadURL.Headers {
		req.Header.Set(h.Name, h.Value)
	}

	httpClient := &http.Client{Timeout: 30 * time.Second}
	abortResp, err := httpClient.Do(req)
	if err != nil {
		logger.Debugf(ctx, "failed to abort multipart upload: %v", err)
		return
	}
	abortResp.Body.Close()
	logger.Debugf(ctx, "aborted multipart upload (status %d)", abortResp.StatusCode)
}

// UploadWithChunking uploads a file to the given path, automatically choosing
// between single-shot and multipart upload based on the content length.
// For files smaller than 50 MiB, a single PUT request is used. For larger files,
// a multipart upload is performed with configurable part size and parallelism.
func (a *FilesAPI) UploadWithChunking(ctx context.Context, filePath string, content io.ReadSeeker, contentLength int64, opts ...UploadOption) error {
	ctx = useragent.InContext(ctx, "sdk-feature", "multipart-upload")

	cfg := &UploadConfig{
		Parallelism: defaultParallelism,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	// Auto-select part size if not explicitly set
	cfg.PartSize, _ = optimizePartSize(contentLength, cfg.PartSize)

	if contentLength < minMultipartUploadSize {
		return a.Upload(ctx, UploadRequest{
			FilePath:  filePath,
			Contents:  io.NopCloser(content),
			Overwrite: cfg.Overwrite,
		})
	}

	return a.uploadMultipart(ctx, filePath, content, contentLength, cfg)
}

// UploadFromFile uploads a local file to the given path, automatically choosing
// between single-shot and multipart upload based on the file size.
func (a *FilesAPI) UploadFromFile(ctx context.Context, filePath string, sourcePath string, opts ...UploadOption) error {
	f, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to stat source file: %w", err)
	}

	return a.UploadWithChunking(ctx, filePath, f, info.Size(), opts...)
}
