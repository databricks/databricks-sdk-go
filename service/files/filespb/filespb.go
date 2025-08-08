// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package filespb

import (
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AddBlockPb struct {
	Data   string `json:"data"`
	Handle int64  `json:"handle"`
}

type AddBlockResponsePb struct {
}

type ClosePb struct {
	Handle int64 `json:"handle"`
}

type CloseResponsePb struct {
}

type CreatePb struct {
	Overwrite bool   `json:"overwrite,omitempty"`
	Path      string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateDirectoryRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
}

type CreateResponsePb struct {
	Handle int64 `json:"handle,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeletePb struct {
	Path      string `json:"path"`
	Recursive bool   `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DeletePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DeletePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteDirectoryRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
}

type DeleteFileRequestPb struct {
	FilePath string `json:"-" url:"-"`
}

type DeleteResponsePb struct {
}

type DirectoryEntryPb struct {
	FileSize     int64  `json:"file_size,omitempty"`
	IsDirectory  bool   `json:"is_directory,omitempty"`
	LastModified int64  `json:"last_modified,omitempty"`
	Name         string `json:"name,omitempty"`
	Path         string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DirectoryEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DirectoryEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DownloadRequestPb struct {
	FilePath string `json:"-" url:"-"`
}

type DownloadResponsePb struct {
	ContentLength int64         `json:"-" url:"-" header:"content-length,omitempty"`
	ContentType   string        `json:"-" url:"-" header:"content-type,omitempty"`
	Contents      io.ReadCloser `json:"-"`
	LastModified  string        `json:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DownloadResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DownloadResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileInfoPb struct {
	FileSize         int64  `json:"file_size,omitempty"`
	IsDir            bool   `json:"is_dir,omitempty"`
	ModificationTime int64  `json:"modification_time,omitempty"`
	Path             string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetDirectoryMetadataRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
}

type GetMetadataRequestPb struct {
	FilePath string `json:"-" url:"-"`
}

type GetMetadataResponsePb struct {
	ContentLength int64  `json:"-" url:"-" header:"content-length,omitempty"`
	ContentType   string `json:"-" url:"-" header:"content-type,omitempty"`
	LastModified  string `json:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetMetadataResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetMetadataResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetStatusRequestPb struct {
	Path string `json:"-" url:"path"`
}

type ListDbfsRequestPb struct {
	Path string `json:"-" url:"path"`
}

type ListDirectoryContentsRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
	PageSize      int64  `json:"-" url:"page_size,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDirectoryContentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDirectoryContentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDirectoryResponsePb struct {
	Contents      []DirectoryEntryPb `json:"contents,omitempty"`
	NextPageToken string             `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListDirectoryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListDirectoryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListStatusResponsePb struct {
	Files []FileInfoPb `json:"files,omitempty"`
}

type MkDirsPb struct {
	Path string `json:"path"`
}

type MkDirsResponsePb struct {
}

type MovePb struct {
	DestinationPath string `json:"destination_path"`
	SourcePath      string `json:"source_path"`
}

type MoveResponsePb struct {
}

type PutPb struct {
	Contents  string `json:"contents,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty"`
	Path      string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PutPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PutPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PutResponsePb struct {
}

type ReadDbfsRequestPb struct {
	Length int64  `json:"-" url:"length,omitempty"`
	Offset int64  `json:"-" url:"offset,omitempty"`
	Path   string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ReadDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ReadDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ReadResponsePb struct {
	BytesRead int64  `json:"bytes_read,omitempty"`
	Data      string `json:"data,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ReadResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ReadResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UploadRequestPb struct {
	Contents  io.ReadCloser `json:"-"`
	FilePath  string        `json:"-" url:"-"`
	Overwrite bool          `json:"-" url:"overwrite,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UploadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UploadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
