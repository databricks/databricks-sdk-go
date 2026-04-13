package files

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/logger"
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

	// uploadInactivityTimeout is the timeout for upload inactivity.
	uploadInactivityTimeout = 60 * time.Second
)

// partSizeOptions lists the candidate part sizes in ascending order.
var partSizeOptions = []int64{
	10 * 1024 * 1024,         // 10 MiB
	20 * 1024 * 1024,         // 20 MiB
	50 * 1024 * 1024,         // 50 MiB
	100 * 1024 * 1024,        // 100 MiB
	200 * 1024 * 1024,        // 200 MiB
	500 * 1024 * 1024,        // 500 MiB
	1 * 1024 * 1024 * 1024,   // 1 GiB
	2 * 1024 * 1024 * 1024,   // 2 GiB
	4 * 1024 * 1024 * 1024,   // 4 GiB
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
	MultipartUpload  *multipartUploadSession  `json:"multipart_upload,omitempty"`
	ResumableUpload  *resumableUploadSession  `json:"resumable_upload,omitempty"`
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
	URL        string           `json:"url"`
	PartNumber int              `json:"part_number"`
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

// uploadOnePart uploads a single part to a presigned URL and returns the ETag.
func (a *FilesAPI) uploadOnePart(ctx context.Context, presigned presignedURL, partData io.ReadSeeker) (string, error) {
	// Determine content length
	currentPos, err := partData.Seek(0, io.SeekCurrent)
	if err != nil {
		return "", fmt.Errorf("failed to get current position: %w", err)
	}
	endPos, err := partData.Seek(0, io.SeekEnd)
	if err != nil {
		return "", fmt.Errorf("failed to seek to end: %w", err)
	}
	contentLength := endPos - currentPos
	if _, err := partData.Seek(currentPos, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to seek back: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, presigned.URL, partData)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.ContentLength = contentLength
	req.Header.Set("Content-Type", "application/octet-stream")
	for _, h := range presigned.Headers {
		req.Header.Set(h.Name, h.Value)
	}

	httpClient := &http.Client{Timeout: uploadInactivityTimeout}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("upload part %d failed: %w", presigned.PartNumber, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("upload part %d failed with status %d: %s", presigned.PartNumber, resp.StatusCode, string(body))
	}

	etag := resp.Header.Get("ETag")
	return etag, nil
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

// uploadPartResult holds the result of uploading a single part.
type uploadPartResult struct {
	PartNumber int
	ETag       string
	Err        error
}

// uploadMultipart orchestrates a full multipart upload.
func (a *FilesAPI) uploadMultipart(ctx context.Context, filePath string, content io.ReadSeeker, contentLength int64, cfg *UploadConfig) error {
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
		// Check if a previous upload failed
		mu.Lock()
		if uploadErr != nil {
			mu.Unlock()
			break
		}
		mu.Unlock()

		// Read the next part
		buf := make([]byte, partSize)
		n, readErr := io.ReadFull(content, buf)
		if n == 0 && readErr != nil {
			break
		}
		buf = buf[:n]

		currentPartNumber := partNumber
		partNumber++

		sem <- struct{}{}
		go func(pn int, data []byte) {
			defer func() { <-sem }()

			// Check for prior error
			mu.Lock()
			if uploadErr != nil {
				mu.Unlock()
				return
			}
			mu.Unlock()

			// Get presigned URL for this part
			urls, err := a.getUploadPartURLs(ctx, filePath, sessionToken, pn, 1)
			if err != nil {
				mu.Lock()
				if uploadErr == nil {
					uploadErr = fmt.Errorf("failed to get upload URL for part %d: %w", pn, err)
				}
				mu.Unlock()
				return
			}
			if len(urls) == 0 {
				mu.Lock()
				if uploadErr == nil {
					uploadErr = fmt.Errorf("no upload URL returned for part %d", pn)
				}
				mu.Unlock()
				return
			}

			// Upload the part
			etag, err := a.uploadOnePart(ctx, urls[0], bytes.NewReader(data))
			if err != nil {
				mu.Lock()
				if uploadErr == nil {
					uploadErr = fmt.Errorf("failed to upload part %d: %w", pn, err)
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
		return uploadErr
	}

	// Phase 3: Complete
	logger.Debugf(ctx, "completing multipart upload with %d parts", len(etags))
	return a.completeMultipartUpload(ctx, filePath, sessionToken, etags)
}
