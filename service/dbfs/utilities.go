package dbfs

import (
	"bufio"
	"context"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/useragent"
)

// FileMode conveys user intent when opening a file.
type FileMode int

const (
	// Exactly one of FileModeRead or FileModeWrite must be specified.
	FileModeRead FileMode = 1 << iota
	FileModeWrite
	FileModeOverwrite
)

// Maximum read or write length for the DBFS API.
const maxDbfsBlockSize = 1024 * 1024

// Internal only state for a read handle.
type dbfsHandleReader struct {
	size   int64
	offset int64
}

// Internal only state for a write handle.
type dbfsHandleWriter struct {
	handle int64
}

// Internal only state for a DBFS file handle.
type dbfsHandle struct {
	ctx  context.Context
	api  DbfsAPI
	path string

	reader *dbfsHandleReader
	writer *dbfsHandleWriter
}

// Handle defines the interface of the object returned by [DbfsAPI.Open].
type Handle interface {
	io.ReadWriteCloser
	io.WriterTo
	io.ReaderFrom
}

// Implement the [io.Reader] interface.
func (h *dbfsHandle) Read(p []byte) (int, error) {
	r := h.reader
	if r == nil {
		return 0, fmt.Errorf("dbfs: file not open for reading")
	}

	var ntotal int
	for ntotal < len(p) {
		if r.offset >= r.size {
			return ntotal, io.EOF
		}

		chunk := p[ntotal:]
		if len(chunk) > maxDbfsBlockSize {
			chunk = chunk[:maxDbfsBlockSize]
		}

		res, err := h.api.Read(h.ctx, Read{
			Path:   h.path,
			Length: len(chunk),
			Offset: int(r.offset), // TODO: make int32/in64 work properly
		})
		if err != nil {
			return ntotal, fmt.Errorf("dbfs: read: %w", err)
		}

		// The guard against offset >= size happens above, so this can only happen
		// if the file is modified or truncated while reading. If this happens,
		// the read contents will likely be corrupted, so we return an error.
		if res.BytesRead == 0 {
			return ntotal, fmt.Errorf("dbfs: read: unexpected EOF at offset %d (size %d)", r.offset, r.size)
		}

		nread, err := base64.StdEncoding.Decode(chunk, []byte(res.Data))
		if err != nil {
			return ntotal, err
		}

		ntotal += nread
		r.offset += int64(nread)
	}

	return ntotal, nil
}

// Implement the [io.Writer] interface.
func (h *dbfsHandle) Write(p []byte) (int, error) {
	w := h.writer
	if w == nil {
		return 0, fmt.Errorf("dbfs: file not open for writing")
	}

	var ntotal int
	for ntotal < len(p) {
		chunk := p[ntotal:]
		if len(chunk) > maxDbfsBlockSize {
			chunk = chunk[:maxDbfsBlockSize]
		}

		err := h.api.AddBlock(h.ctx, AddBlock{
			Data:   base64.StdEncoding.EncodeToString(chunk),
			Handle: w.handle,
		})
		if err != nil {
			return ntotal, fmt.Errorf("dbfs: add block: %w", err)
		}

		ntotal += len(chunk)
	}

	return ntotal, nil
}

// Implement the [io.Closer] interface.
func (h *dbfsHandle) Close() error {
	w := h.writer
	if w == nil {
		return fmt.Errorf("dbfs: file not open for writing")
	}

	err := h.api.CloseByHandle(h.ctx, w.handle)
	if err != nil {
		return fmt.Errorf("dbfs: close: %w", err)
	}

	return nil
}

// Implement the [io.WriterTo] interface.
func (h *dbfsHandle) WriteTo(w io.Writer) (int64, error) {
	if h.reader == nil {
		return 0, fmt.Errorf("dbfs: file not open for reading")
	}

	// Limit types to io.Reader and io.Writer to avoid recursion
	// into WriteTo or ReadFrom functions on underlying types.
	ior := struct{ io.Reader }{h}
	iow := struct{ io.Writer }{w}
	return bufio.NewReaderSize(ior, maxDbfsBlockSize).WriteTo(iow)
}

// Implement the [io.ReaderFrom] interface.
func (h *dbfsHandle) ReadFrom(r io.Reader) (int64, error) {
	if h.writer == nil {
		return 0, fmt.Errorf("dbfs: file not open for writing")
	}

	// Limit types to io.Reader and io.Writer to avoid recursion
	// into WriteTo or ReadFrom functions on underlying types.
	ior := struct{ io.Reader }{r}
	iow := struct{ io.Writer }{h}
	bw := bufio.NewWriterSize(iow, maxDbfsBlockSize)
	n, err := bw.ReadFrom(ior)
	if err != nil {
		return n, err
	}
	return n, bw.Flush()
}

func (h *dbfsHandle) openForRead(mode FileMode) error {
	res, err := h.api.GetStatusByPath(h.ctx, h.path)
	if err != nil {
		return err
	}
	h.reader = &dbfsHandleReader{
		size: res.FileSize,
	}
	return nil
}

func (h *dbfsHandle) openForWrite(mode FileMode) error {
	res, err := h.api.Create(h.ctx, Create{
		Path:      h.path,
		Overwrite: (mode & FileModeOverwrite) != 0,
	})
	if err != nil {
		return err
	}
	h.writer = &dbfsHandleWriter{
		handle: res.Handle,
	}
	return nil
}

// Open opens a remote DBFS file for reading or writing.
// The returned object implements relevant [io] interfaces for convenient
// integration with other code that reads or writes bytes.
//
// The [io.WriterTo] interface is provided and maximizes throughput for
// bulk reads by reading data with the DBFS maximum read chunk size of 1MB.
// Similarly, the [io.ReaderFrom] interface is provided for bulk writing.
//
// A file opened for writing must always be closed.
func (a DbfsAPI) Open(ctx context.Context, path string, mode FileMode) (Handle, error) {
	h := &dbfsHandle{
		ctx:  useragent.InContext(ctx, "sdk-feature", "dbfs-handle"),
		api:  a,
		path: path,
	}

	isRead := (mode & FileModeRead) != 0
	isWrite := (mode & FileModeWrite) != 0
	if isRead && isWrite {
		return nil, fmt.Errorf("dbfs: cannot open file for reading and writing at the same time")
	}

	var err error
	switch {
	case isRead:
		err = h.openForRead(mode)
	case isWrite:
		err = h.openForWrite(mode)
	default:
		// No mode specifed. The caller should be explicit so we return an error.
		return nil, fmt.Errorf("dbfs: must specify dbfs.FileModeRead or dbfs.FileModeWrite")
	}

	if err != nil {
		return nil, fmt.Errorf("dbfs: open: %w", err)
	}

	return h, nil
}

// ReadFile is identical to [os.ReadFile] but for DBFS.
func (a DbfsAPI) ReadFile(ctx context.Context, name string) ([]byte, error) {
	h, err := a.Open(ctx, name, FileModeRead)
	if err != nil {
		return nil, err
	}

	h_ := h.(*dbfsHandle)
	buf := make([]byte, h_.reader.size)
	_, err = h.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf, nil
}

// WriteFile is identical to [os.WriteFile] but for DBFS.
func (a DbfsAPI) WriteFile(ctx context.Context, name string, data []byte) error {
	h, err := a.Open(ctx, name, FileModeWrite|FileModeOverwrite)
	if err != nil {
		return err
	}

	_, err = h.Write(data)
	cerr := h.Close()
	if err == nil && cerr != nil {
		err = cerr
	}
	return err
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
