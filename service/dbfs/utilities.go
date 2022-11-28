package dbfs

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/useragent"
)

var b64 = base64.StdEncoding

// Overwrite is like Put, but more friendly
func (a *DbfsAPI) Overwrite(ctx context.Context, path string, r io.Reader) (err error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "dbfs-overwrite")
	handle, err := a.Create(ctx, Create{
		Path:      path,
		Overwrite: true,
	})
	if err != nil {
		return fmt.Errorf("create: %w", err)
	}
	defer func() {
		cerr := a.CloseByHandle(ctx, handle.Handle)
		if cerr != nil {
			err = fmt.Errorf("close: %w", cerr)
		}
	}()
	buffer := make([]byte, 1e6)
	for {
		n, err := r.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		err = a.AddBlock(ctx, AddBlock{
			Data:   b64.EncodeToString(buffer[0:n]),
			Handle: handle.Handle,
		})
		if err != nil {
			return fmt.Errorf("add block: %w", err)
		}
	}
	return err
}

type FileReader struct {
	Size   int64
	ctx    context.Context
	api    *DbfsAPI
	path   string
	offset int64
}

func (r *FileReader) Read(p []byte) (n int, err error) {
	if r.api == nil {
		panic("invalid call")
	}
	resp, err := r.api.Read(r.ctx, Read{
		Path:   r.path,
		Length: len(p),
		Offset: int(r.offset), // TODO: make int32/in64 work properly
	})
	if err != nil {
		return 0, fmt.Errorf("dbfs read: %w", err)
	}
	if resp.BytesRead == 0 {
		return 0, io.EOF
	}
	r.offset += resp.BytesRead
	chunk, err := b64.DecodeString(resp.Data)
	if err != nil {
		return 0, fmt.Errorf("base64 decode: %w", err)
	}
	return copy(p, chunk), nil
}

func (a *DbfsAPI) Open(ctx context.Context, path string) (*FileReader, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "dbfs-open")
	info, err := a.GetStatusByPath(ctx, path)
	if err != nil {
		return nil, fmt.Errorf("get status: %w", err)
	}
	return &FileReader{
		Size: info.FileSize,
		path: path,
		ctx:  ctx,
		api:  a,
	}, nil
}

// RecursiveList traverses the DBFS tree and returns all non-directory
// objects under the path
func (a DbfsAPI) RecursiveList(ctx context.Context, path string) ([]FileInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "recursive-list")
	var results []FileInfo
	queue := []string{path}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		batch, err := a.ListAll(ctx, List{
			Path: path,
		})
		if err != nil {
			return nil, fmt.Errorf("list %s: %w", path, err)
		}
		for _, v := range batch {
			if v.IsDir {
				queue = append(queue, v.Path)
				continue
			}
			results = append(results, v)
		}
	}
	return results, nil
}
