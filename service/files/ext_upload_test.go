package files

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOptimizePartSize(t *testing.T) {
	const MiB = 1024 * 1024
	const GiB = 1024 * MiB

	tests := []struct {
		name              string
		contentLength     int64
		explicitPartSize  int64
		wantPartSize      int64
		wantBatchSize     int
	}{
		{
			// 5 MiB / 10 MiB = 1 part; ceil(sqrt(1)) = 1
			name:          "small file 5 MiB returns defaultPartSize batch 1",
			contentLength: 5 * MiB,
			wantPartSize:  defaultPartSize,
			wantBatchSize: 1,
		},
		{
			// 500 MiB / 10 MiB = 50 parts <= 100; ceil(sqrt(50)) = 8
			name:          "500 MiB returns 10 MiB parts batch 8",
			contentLength: 500 * MiB,
			wantPartSize:  10 * MiB,
			wantBatchSize: 8,
		},
		{
			// 5 GiB: 10 MiB = 512 parts (too many), 20 MiB = 256 (too many),
			// 50 MiB = 103 (too many), 100 MiB = 52 <= 100; ceil(sqrt(52)) = 8
			name:          "5 GiB returns 100 MiB parts batch 8",
			contentLength: 5 * GiB,
			wantPartSize:  100 * MiB,
			wantBatchSize: 8,
		},
		{
			// Explicit 20 MiB for 500 MiB file: 500/20 = 25 parts; ceil(sqrt(25)) = 5
			name:             "explicit part size 20 MiB for 500 MiB file batch 5",
			contentLength:    500 * MiB,
			explicitPartSize: 20 * MiB,
			wantPartSize:     20 * MiB,
			wantBatchSize:    5,
		},
		{
			// Unknown size (0): returns defaultPartSize, batch 1
			name:          "unknown size returns defaultPartSize batch 1",
			contentLength: 0,
			wantPartSize:  defaultPartSize,
			wantBatchSize: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotPartSize, gotBatchSize := optimizePartSize(tc.contentLength, tc.explicitPartSize)
			if gotPartSize != tc.wantPartSize {
				t.Errorf("optimizePartSize(%d, %d) partSize = %d, want %d",
					tc.contentLength, tc.explicitPartSize, gotPartSize, tc.wantPartSize)
			}
			if gotBatchSize != tc.wantBatchSize {
				t.Errorf("optimizePartSize(%d, %d) batchSize = %d, want %d",
					tc.contentLength, tc.explicitPartSize, gotBatchSize, tc.wantBatchSize)
			}
		})
	}
}

// multipartMockServer simulates the multipart upload protocol for testing.
type multipartMockServer struct {
	mu             sync.Mutex
	parts          map[int][]byte
	completed      bool
	completeParts  []completeUploadPart
	sessionToken   string
}

func newMultipartMockServer() *multipartMockServer {
	return &multipartMockServer{
		parts:        make(map[int][]byte),
		sessionToken: "test-session-token-123",
	}
}

func (m *multipartMockServer) handler(srv *httptest.Server) http.Handler {
	mux := http.NewServeMux()

	// Initiate upload
	mux.HandleFunc("/api/2.0/fs/files/", func(w http.ResponseWriter, r *http.Request) {
		action := r.URL.Query().Get("action")
		switch action {
		case "initiate-upload":
			w.Header().Set("Content-Type", "application/json")
			resp := initiateUploadResponse{
				MultipartUpload: &multipartUploadSession{
					SessionToken: m.sessionToken,
				},
			}
			json.NewEncoder(w).Encode(resp)
		case "complete-upload":
			var req completeUploadRequest
			if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			m.mu.Lock()
			m.completed = true
			m.completeParts = req.Parts
			m.mu.Unlock()
			w.WriteHeader(http.StatusOK)
		default:
			http.Error(w, "unknown action", http.StatusBadRequest)
		}
	})

	// Create upload part URLs
	mux.HandleFunc("/api/2.0/fs/create-upload-part-urls", func(w http.ResponseWriter, r *http.Request) {
		var req createUploadPartURLsRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		urls := make([]presignedURL, req.Count)
		for i := 0; i < req.Count; i++ {
			pn := req.StartPartNumber + i
			urls[i] = presignedURL{
				URL:        fmt.Sprintf("%s/upload-part/%d", srv.URL, pn),
				PartNumber: pn,
				Headers: []presignedHeader{
					{Name: "x-test-header", Value: "test-value"},
				},
			}
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(createUploadPartURLsResponse{UploadPartURLs: urls})
	})

	// Upload part (presigned URL target)
	mux.HandleFunc("/upload-part/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/upload-part/"), "/")
		partNum, err := strconv.Atoi(parts[0])
		if err != nil {
			http.Error(w, "invalid part number", http.StatusBadRequest)
			return
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		m.mu.Lock()
		m.parts[partNum] = data
		m.mu.Unlock()

		etag := fmt.Sprintf("etag-part-%d", partNum)
		w.Header().Set("ETag", etag)
		w.WriteHeader(http.StatusOK)
	})

	return mux
}

func newTestFilesAPI(t *testing.T, serverURL string) *FilesAPI {
	t.Helper()
	cfg := &config.Config{
		Host:  serverURL,
		Token: "test-token",
	}
	err := cfg.EnsureResolved()
	require.NoError(t, err)
	databricksClient, err := client.New(cfg)
	require.NoError(t, err)
	return NewFiles(databricksClient)
}

func TestUploadMultipart_FullFlow(t *testing.T) {
	mock := newMultipartMockServer()

	// Create server with a temporary handler, then replace with the real one
	srv := httptest.NewServer(http.NotFoundHandler())
	defer srv.Close()
	srv.Config.Handler = mock.handler(srv)

	api := newTestFilesAPI(t, srv.URL)

	// Create 100 KiB of test content
	contentSize := 100 * 1024
	content := make([]byte, contentSize)
	for i := range content {
		content[i] = byte(i % 256)
	}

	cfg := &UploadConfig{
		PartSize:    30 * 1024, // 30 KiB parts
		Parallelism: 2,
		Overwrite:   true,
	}

	ctx := context.Background()
	err := api.uploadMultipart(ctx, "/test/upload.bin", strings.NewReader(string(content)), int64(contentSize), cfg)
	require.NoError(t, err)

	// Verify completion
	mock.mu.Lock()
	defer mock.mu.Unlock()

	assert.True(t, mock.completed, "upload should be completed")
	assert.True(t, len(mock.parts) > 1, "should have multiple parts, got %d", len(mock.parts))

	// Verify total bytes match
	totalBytes := 0
	for _, data := range mock.parts {
		totalBytes += len(data)
	}
	assert.Equal(t, contentSize, totalBytes, "total uploaded bytes should match content size")

	// 100 KiB / 30 KiB = 4 parts (30+30+30+10)
	expectedParts := 4
	assert.Equal(t, expectedParts, len(mock.completeParts), "complete request should have correct number of parts")

	// Verify parts are sorted by part number
	for i := 1; i < len(mock.completeParts); i++ {
		assert.True(t, mock.completeParts[i].PartNumber > mock.completeParts[i-1].PartNumber,
			"parts should be sorted by part number")
	}
}
