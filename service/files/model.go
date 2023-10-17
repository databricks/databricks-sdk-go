// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

type AddBlock struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data string `json:"data"`
	// The handle on an open stream.
	Handle int64 `json:"handle"`
}

type Close struct {
	// The handle on an open stream.
	Handle int64 `json:"handle"`
}

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path string `json:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *Create) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Create) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle int64 `json:"handle,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *CreateResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path string `json:"path"`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *Delete) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Delete) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a file or directory
type DeleteFileRequest struct {
	// The absolute path of the file or directory in DBFS.
	FilePath string `json:"-" url:"-"`
}

// Download a file
type DownloadRequest struct {
	// The absolute path of the file or directory in DBFS.
	FilePath string `json:"-" url:"-"`
}

type DownloadResponse struct {
	Contents io.ReadCloser `json:"-"`
}

type FileInfo struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize int64 `json:"file_size,omitempty"`
	// True if the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// Last modification time of given file in milliseconds since epoch.
	ModificationTime int64 `json:"modification_time,omitempty"`
	// The absolute path of the file or directory.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *FileInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FileInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the information of a file or directory
type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `json:"-" url:"path"`
}

// List directory contents or file details
type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `json:"-" url:"path"`
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files []FileInfo `json:"files,omitempty"`
}

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path string `json:"path"`
}

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath string `json:"destination_path"`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath string `json:"source_path"`
}

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents string `json:"contents,omitempty"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path string `json:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *Put) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Put) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get the contents of a file
type ReadDbfsRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length int `json:"-" url:"length,omitempty"`
	// The offset to read from in bytes.
	Offset int `json:"-" url:"offset,omitempty"`
	// The path of the file to read. The path should be the absolute DBFS path.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-"`
}

func (s *ReadDbfsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadDbfsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ReadResponse struct {
	// The number of bytes read (could be less than `length` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead int64 `json:"bytes_read,omitempty"`
	// The base64-encoded contents of the file read.
	Data string `json:"data,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ReadResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ReadResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Upload a file
type UploadRequest struct {
	Contents io.ReadCloser `json:"-"`
	// The absolute path of the file or directory in DBFS.
	FilePath string `json:"-" url:"-"`
	// If true, an existing file will be overwritten.
	Overwrite bool `json:"-" url:"overwrite,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *UploadRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UploadRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
