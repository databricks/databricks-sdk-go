package files

import (
	"math"
	"time"
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
