package workspace

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
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
		Format:    ImportFormatSource,
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

type UploadOption = func(*Import)

func UploadOverwrite() UploadOption {
	return func(i *Import) {
		i.Overwrite = true
	}
}

func UploadLanguage(l Language) UploadOption {
	return func(i *Import) {
		i.Language = l
	}
}

func UploadFormat(f ImportFormat) UploadOption {
	return func(i *Import) {
		i.Format = f
	}
}

// Upload a workspace object (for example, a notebook or file) or the contents
// of an entire directory (`DBC` format).
//
// Errors:
//
//   - RESOURCE_ALREADY_EXISTS: if `path` already exists no `overwrite=True`.
//   - INVALID_PARAMETER_VALUE: if `format` and `content` values are not compatible.
//
// By default, workspace.UploadFormat(workspace.ImportFormatSource). If using
// workspace.UploadFormat(workspace.ImportFormatAuto) the `path` is imported or
// exported as either a workspace file or a notebook, depending on an analysis
// of the `path`’s extension and the header content provided in the request.
// In addition, if the `path` is imported as a notebook, then the `item`’s
// extension is automatically removed.
//
// workspace.UploadLanguage(...) is only required if source format.
func (a *WorkspaceAPI) Upload(ctx context.Context, path string, r io.Reader, opts ...UploadOption) error {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)
	err := w.WriteField("path", path)
	if err != nil {
		return fmt.Errorf("write path: %w", err)
	}
	content, err := w.CreateFormFile("content", "content")
	if err != nil {
		return fmt.Errorf("write content: %w", err)
	}
	_, err = io.Copy(content, r)
	if err != nil {
		return fmt.Errorf("copy io: %w", err)
	}
	i := &Import{}
	for _, v := range opts {
		v(i)
	}
	if i.Format == "" || i.Format == ImportFormatSource {
		for sfx, lang := range map[string]Language{
			".py":    LanguagePython,
			".sql":   LanguageSql,
			".scala": LanguageScala,
			".R":     LanguageR,
		} {
			if !strings.HasSuffix(path, sfx) {
				continue
			}
			i.Language = lang
		}
	}
	if i.Format != "" {
		err = w.WriteField("format", i.Format.String())
		if err != nil {
			return fmt.Errorf("write format: %w", err)
		}
	}
	if i.Language != "" {
		err = w.WriteField("language", i.Language.String())
		if err != nil {
			return fmt.Errorf("write language: %w", err)
		}
	}
	if i.Overwrite {
		err = w.WriteField("overwrite", "true")
		if err != nil {
			return fmt.Errorf("write overwrite: %w", err)
		}
	}
	err = w.Close()
	if err != nil {
		return fmt.Errorf("write close: %w", err)
	}
	impl, ok := a.impl.(*workspaceImpl)
	if !ok {
		return fmt.Errorf("wrong impl: %v", a.impl)
	}
	return impl.client.Do(ctx, "POST", "/api/2.0/workspace/import", buf.Bytes(), nil,
		func(r *http.Request) error {
			r.Header.Set("Content-Type", w.FormDataContentType())
			return nil
		})
}

// WriteFile is identical to [os.WriteFile] but for Workspace File.
// Keep in mind: It doesn't upload the notebook, but the file and does
// always overwrite it.
func (a *WorkspaceAPI) WriteFile(ctx context.Context, name string, data []byte) error {
	return a.Upload(ctx, name, bytes.NewBuffer(data),
		UploadFormat(ImportFormatAuto),
		UploadOverwrite())
}

type DownloadOption = func(q map[string]any)

func DownloadFormat(f ExportFormat) func(q map[string]any) {
	return func(q map[string]any) {
		q["format"] = f.String()
	}
}

// Download a notebook or file from the workspace by path.
//
// By default, it acts as if workspace.DownloadFormat(workspace.ExportFormatSource) option is supplied. When using
// workspace.ExportFormatAuto, the `path` is imported or exported as either a workspace file or a notebook, depending
// on an analysis of the `item`’s extension and the file content header provided in the request.
//
// Returns [bytes.Buffer] of the path contents.
func (a *WorkspaceAPI) Download(ctx context.Context, path string, opts ...DownloadOption) (io.ReadCloser, error) {
	impl, ok := a.impl.(*workspaceImpl)
	if !ok {
		return nil, fmt.Errorf("wrong impl: %v", a.impl)
	}
	var buf bytes.Buffer
	query := map[string]any{"path": path, "direct_download": true}
	for _, v := range opts {
		v(query)
	}
	err := impl.client.Do(ctx, "GET", "/api/2.0/workspace/export", query, &buf)
	if err != nil {
		return nil, err
	}
	return io.NopCloser(&buf), nil
}

// ReadFile is identical to [os.ReadFile] but for workspace files.
func (a *WorkspaceAPI) ReadFile(ctx context.Context, name string) ([]byte, error) {
	b, err := a.Download(ctx, name)
	if err != nil {
		return nil, err
	}
	defer b.Close()
	return io.ReadAll(b)
}
