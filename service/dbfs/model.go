// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbfs

// all definitions in this file are in alphabetical order

type AddBlockRequest struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data string `json:"data"`
	// The handle on an open stream.
	Handle int64 `json:"handle"`
}

type CloseRequest struct {
	// The handle on an open stream.
	Handle int64 `json:"handle"`
}

type CreateRequest struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path (e.g.
	// &#34;/mnt/foo.txt&#34;).
	Path string `json:"path"`
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle int64 `json:"handle,omitempty"`
}

type DeleteRequest struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path (e.g. &#34;/mnt/foo/&#34;).
	Path string `json:"path"`
	// Whether or not to recursively delete the directory&#39;s contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive bool `json:"recursive,omitempty"`
}

type FileInfo struct {
	// The length of the file in bytes or zero if the path is a directory.
	FileSize int64 `json:"file_size,omitempty"`
	// True if the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// Last modification time of given file/dir in milliseconds since Epoch.
	ModificationTime int64 `json:"modification_time,omitempty"`
	// The path of the file or directory.
	Path string `json:"path,omitempty"`
}

type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path (e.g. &#34;/mnt/foo/&#34;).
	Path string ` url:"path,omitempty"`
}

type GetStatusResponse struct {
	// The length of the file in bytes or zero if the path is a directory.
	FileSize int64 `json:"file_size,omitempty"`
	// True if the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// Last modification time of given file/dir in milliseconds since Epoch.
	ModificationTime int64 `json:"modification_time,omitempty"`
	// The path of the file or directory.
	Path string `json:"path,omitempty"`
}

type ListStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path (e.g. &#34;/mnt/foo/&#34;).
	Path string ` url:"path,omitempty"`
}

type ListStatusResponse struct {
	// A list of FileInfo&#39;s that describe contents of directory or file. See
	// example above.
	Files []FileInfo `json:"files,omitempty"`
}

type MkDirsRequest struct {
	// The path of the new directory. The path should be the absolute DBFS path
	// (e.g. &#34;/mnt/foo/&#34;).
	Path string `json:"path"`
}

type MoveRequest struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path (e.g. &#34;/mnt/bar/&#34;).
	DestinationPath string `json:"destination_path"`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path (e.g. &#34;/mnt/foo/&#34;).
	SourcePath string `json:"source_path"`
}

type PutRequest struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents string `json:"contents,omitempty"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path (e.g.
	// &#34;/mnt/foo/&#34;).
	Path string `json:"path"`
}

type ReadRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length int ` url:"length,omitempty"`
	// The offset to read from in bytes.
	Offset int ` url:"offset,omitempty"`
	// The path of the file to read. The path should be the absolute DBFS path
	// (e.g. &#34;/mnt/foo/&#34;).
	Path string ` url:"path,omitempty"`
}

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead int64 `json:"bytes_read,omitempty"`
	// The base64-encoded contents of the file read.
	Data string `json:"data,omitempty"`
}
