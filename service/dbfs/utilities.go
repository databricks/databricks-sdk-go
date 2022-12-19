package dbfs

import (
	"context"
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// Overwrite is like Put, but more friendly
func (a *DbfsAPI) Overwrite(ctx context.Context, path string, r io.Reader) (err error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "dbfs-overwrite")
	handle, err := OpenFile(ctx, a, path, DbfsWrite|DbfsOverwrite)
	if err != nil {
		return err
	}
	_, err = handle.ReadFrom(r)
	cerr := handle.Close()
	if err != nil {
		return err
	}
	if cerr != nil {
		return cerr
	}
	return nil
}

func (a *DbfsAPI) Open(ctx context.Context, path string) (*dbfsHandle, error) {
	return OpenFile(ctx, a, path, DbfsRead)
}

func (a *DbfsAPI) OpenFile(ctx context.Context, path string, mode DbfsFileMode) (*dbfsHandle, error) {
	return OpenFile(ctx, a, path, mode)
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
		if apierr.IsMissing(err) {
			// skip on path deleted during iteration
			continue
		}
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
