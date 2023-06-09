package files

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// BEGIN TODO: remove this once FilesAPI gets OpenAPI spec

func NewFiles(client *client.DatabricksClient) *FilesAPI {
	return &FilesAPI{
		impl: &filesImpl{
			client: client,
		},
	}
}

type FilesService interface{}

type FilesAPI struct {
	impl FilesService
}

type filesImpl struct {
	client *client.DatabricksClient
}

// END TODO: remove this once FilesAPI gets OpenAPI spec

func (a *FilesAPI) Upload(ctx context.Context, path string, r io.Reader) error {
	impl, ok := a.impl.(*filesImpl)
	if !ok {
		return fmt.Errorf("wrong impl: %v", a.impl)
	}
	return impl.client.Do(ctx, "PUT", "/api/2.0/fs/files"+path, r, nil,
		func(r *http.Request) error {
			r.Header.Set("Content-Type", "application/octet-stream")
			return nil
		})
}

// WriteFile is identical to [os.WriteFile] but for Files API.
func (a *FilesAPI) WriteFile(ctx context.Context, name string, data []byte) error {
	return a.Upload(ctx, name, bytes.NewBuffer(data))
}

func (a *FilesAPI) Download(ctx context.Context, path string) (io.ReadCloser, error) {
	impl, ok := a.impl.(*filesImpl)
	if !ok {
		return nil, fmt.Errorf("wrong impl: %v", a.impl)
	}
	var buf bytes.Buffer
	err := impl.client.Do(ctx, "GET", "/api/2.0/fs/files"+path, nil, &buf)
	if err != nil {
		return nil, err
	}
	// TODO: direclty call lower-level APIs to return raw HTTP response stream
	return io.NopCloser(&buf), nil
}

// ReadFile is identical to [os.ReadFile] but for Files API.
func (a *FilesAPI) ReadFile(ctx context.Context, name string) ([]byte, error) {
	b, err := a.Download(ctx, name)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(b)
}

func (a *FilesAPI) Delete(ctx context.Context, path string) error {
	impl, ok := a.impl.(*filesImpl)
	if !ok {
		return fmt.Errorf("wrong impl: %v", a.impl)
	}
	return impl.client.Do(ctx, "DELETE", "/api/2.0/fs/files"+path, nil, nil)
}
