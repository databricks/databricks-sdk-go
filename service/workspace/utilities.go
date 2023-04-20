package workspace

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/compute"
	"github.com/databricks/databricks-sdk-go/useragent"
)

var b64 = base64.StdEncoding

// PythonNotebookOverwrite crafts Python import notebook request
// also by trimming the code specified in the second argument
func PythonNotebookOverwrite(path, content string) Import {
	content = compute.TrimLeadingWhitespace(content)
	request, _ := PythonNotebookOverwriteReader(path,
		strings.NewReader(content))
	return request
}

func PythonNotebookOverwriteReader(path string, r io.Reader) (Import, error) {
	raw, err := io.ReadAll(r)
	if err != nil {
		return Import{}, fmt.Errorf("read: %w", err)
	}
	return Import{
		Path:      path,
		Overwrite: true,
		Format:    ExportFormatSource,
		Language:  LanguagePython,
		Content:   b64.EncodeToString(raw),
	}, nil
}

func (r *ExportResponse) Bytes() ([]byte, error) {
	return b64.DecodeString(r.Content)
}

// RecursiveList traverses the workspace tree and returns all non-directory
// objects under the path
func (a *WorkspaceAPI) RecursiveList(ctx context.Context, path string) ([]ObjectInfo, error) {
	ctx = useragent.InContext(ctx, "sdk-feature", "recursive-list")
	var results []ObjectInfo
	queue := []string{path}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		batch, err := a.ListAll(ctx, ListWorkspaceRequest{
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
			if v.ObjectType == ObjectTypeDirectory {
				queue = append(queue, v.Path)
				continue
			}
			results = append(results, v)
		}
	}
	return results, nil
}
