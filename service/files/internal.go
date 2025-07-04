// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func addBlockToPb(st *AddBlock) (*addBlockPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addBlockPb{}
	pb.Data = st.Data
	pb.Handle = st.Handle

	return pb, nil
}

type addBlockPb struct {
	Data   string `json:"data"`
	Handle int64  `json:"handle"`
}

func addBlockFromPb(pb *addBlockPb) (*AddBlock, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddBlock{}
	st.Data = pb.Data
	st.Handle = pb.Handle

	return st, nil
}

func closeToPb(st *Close) (*closePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &closePb{}
	pb.Handle = st.Handle

	return pb, nil
}

type closePb struct {
	Handle int64 `json:"handle"`
}

func closeFromPb(pb *closePb) (*Close, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Close{}
	st.Handle = pb.Handle

	return st, nil
}

func createToPb(st *Create) (*createPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPb{}
	pb.Overwrite = st.Overwrite
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPb struct {
	Overwrite bool   `json:"overwrite,omitempty"`
	Path      string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFromPb(pb *createPb) (*Create, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Create{}
	st.Overwrite = pb.Overwrite
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createDirectoryRequestToPb(st *CreateDirectoryRequest) (*createDirectoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDirectoryRequestPb{}
	pb.DirectoryPath = st.DirectoryPath

	return pb, nil
}

type createDirectoryRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
}

func createDirectoryRequestFromPb(pb *createDirectoryRequestPb) (*CreateDirectoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDirectoryRequest{}
	st.DirectoryPath = pb.DirectoryPath

	return st, nil
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}
	pb.Handle = st.Handle

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createResponsePb struct {
	Handle int64 `json:"handle,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	st.Handle = pb.Handle

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteToPb(st *Delete) (*deletePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePb{}
	pb.Path = st.Path
	pb.Recursive = st.Recursive

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type deletePb struct {
	Path      string `json:"path"`
	Recursive bool   `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteFromPb(pb *deletePb) (*Delete, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Delete{}
	st.Path = pb.Path
	st.Recursive = pb.Recursive

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deletePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deletePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deleteDirectoryRequestToPb(st *DeleteDirectoryRequest) (*deleteDirectoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDirectoryRequestPb{}
	pb.DirectoryPath = st.DirectoryPath

	return pb, nil
}

type deleteDirectoryRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
}

func deleteDirectoryRequestFromPb(pb *deleteDirectoryRequestPb) (*DeleteDirectoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDirectoryRequest{}
	st.DirectoryPath = pb.DirectoryPath

	return st, nil
}

func deleteFileRequestToPb(st *DeleteFileRequest) (*deleteFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFileRequestPb{}
	pb.FilePath = st.FilePath

	return pb, nil
}

type deleteFileRequestPb struct {
	FilePath string `json:"-" url:"-"`
}

func deleteFileRequestFromPb(pb *deleteFileRequestPb) (*DeleteFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFileRequest{}
	st.FilePath = pb.FilePath

	return st, nil
}

func directoryEntryToPb(st *DirectoryEntry) (*directoryEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &directoryEntryPb{}
	pb.FileSize = st.FileSize
	pb.IsDirectory = st.IsDirectory
	pb.LastModified = st.LastModified
	pb.Name = st.Name
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type directoryEntryPb struct {
	FileSize     int64  `json:"file_size,omitempty"`
	IsDirectory  bool   `json:"is_directory,omitempty"`
	LastModified int64  `json:"last_modified,omitempty"`
	Name         string `json:"name,omitempty"`
	Path         string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func directoryEntryFromPb(pb *directoryEntryPb) (*DirectoryEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DirectoryEntry{}
	st.FileSize = pb.FileSize
	st.IsDirectory = pb.IsDirectory
	st.LastModified = pb.LastModified
	st.Name = pb.Name
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *directoryEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st directoryEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func downloadRequestToPb(st *DownloadRequest) (*downloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &downloadRequestPb{}
	pb.FilePath = st.FilePath

	return pb, nil
}

type downloadRequestPb struct {
	FilePath string `json:"-" url:"-"`
}

func downloadRequestFromPb(pb *downloadRequestPb) (*DownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadRequest{}
	st.FilePath = pb.FilePath

	return st, nil
}

func downloadResponseToPb(st *DownloadResponse) (*downloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &downloadResponsePb{}
	pb.ContentLength = st.ContentLength
	pb.ContentType = st.ContentType
	pb.Contents = st.Contents
	pb.LastModified = st.LastModified

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type downloadResponsePb struct {
	ContentLength int64         `json:"-" url:"-" header:"content-length,omitempty"`
	ContentType   string        `json:"-" url:"-" header:"content-type,omitempty"`
	Contents      io.ReadCloser `json:"-"`
	LastModified  string        `json:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func downloadResponseFromPb(pb *downloadResponsePb) (*DownloadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadResponse{}
	st.ContentLength = pb.ContentLength
	st.ContentType = pb.ContentType
	st.Contents = pb.Contents
	st.LastModified = pb.LastModified

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *downloadResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st downloadResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func fileInfoToPb(st *FileInfo) (*fileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileInfoPb{}
	pb.FileSize = st.FileSize
	pb.IsDir = st.IsDir
	pb.ModificationTime = st.ModificationTime
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type fileInfoPb struct {
	FileSize         int64  `json:"file_size,omitempty"`
	IsDir            bool   `json:"is_dir,omitempty"`
	ModificationTime int64  `json:"modification_time,omitempty"`
	Path             string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileInfoFromPb(pb *fileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	st.FileSize = pb.FileSize
	st.IsDir = pb.IsDir
	st.ModificationTime = pb.ModificationTime
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getDirectoryMetadataRequestToPb(st *GetDirectoryMetadataRequest) (*getDirectoryMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDirectoryMetadataRequestPb{}
	pb.DirectoryPath = st.DirectoryPath

	return pb, nil
}

type getDirectoryMetadataRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
}

func getDirectoryMetadataRequestFromPb(pb *getDirectoryMetadataRequestPb) (*GetDirectoryMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDirectoryMetadataRequest{}
	st.DirectoryPath = pb.DirectoryPath

	return st, nil
}

func getMetadataRequestToPb(st *GetMetadataRequest) (*getMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetadataRequestPb{}
	pb.FilePath = st.FilePath

	return pb, nil
}

type getMetadataRequestPb struct {
	FilePath string `json:"-" url:"-"`
}

func getMetadataRequestFromPb(pb *getMetadataRequestPb) (*GetMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetadataRequest{}
	st.FilePath = pb.FilePath

	return st, nil
}

func getMetadataResponseToPb(st *GetMetadataResponse) (*getMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetadataResponsePb{}
	pb.ContentLength = st.ContentLength
	pb.ContentType = st.ContentType
	pb.LastModified = st.LastModified

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getMetadataResponsePb struct {
	ContentLength int64  `json:"-" url:"-" header:"content-length,omitempty"`
	ContentType   string `json:"-" url:"-" header:"content-type,omitempty"`
	LastModified  string `json:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getMetadataResponseFromPb(pb *getMetadataResponsePb) (*GetMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetadataResponse{}
	st.ContentLength = pb.ContentLength
	st.ContentType = pb.ContentType
	st.LastModified = pb.LastModified

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getMetadataResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getMetadataResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getStatusRequestToPb(st *GetStatusRequest) (*getStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatusRequestPb{}
	pb.Path = st.Path

	return pb, nil
}

type getStatusRequestPb struct {
	Path string `json:"-" url:"path"`
}

func getStatusRequestFromPb(pb *getStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	st.Path = pb.Path

	return st, nil
}

func listDbfsRequestToPb(st *ListDbfsRequest) (*listDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDbfsRequestPb{}
	pb.Path = st.Path

	return pb, nil
}

type listDbfsRequestPb struct {
	Path string `json:"-" url:"path"`
}

func listDbfsRequestFromPb(pb *listDbfsRequestPb) (*ListDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDbfsRequest{}
	st.Path = pb.Path

	return st, nil
}

func listDirectoryContentsRequestToPb(st *ListDirectoryContentsRequest) (*listDirectoryContentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDirectoryContentsRequestPb{}
	pb.DirectoryPath = st.DirectoryPath
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDirectoryContentsRequestPb struct {
	DirectoryPath string `json:"-" url:"-"`
	PageSize      int64  `json:"-" url:"page_size,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDirectoryContentsRequestFromPb(pb *listDirectoryContentsRequestPb) (*ListDirectoryContentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDirectoryContentsRequest{}
	st.DirectoryPath = pb.DirectoryPath
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDirectoryContentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDirectoryContentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listDirectoryResponseToPb(st *ListDirectoryResponse) (*listDirectoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDirectoryResponsePb{}
	pb.Contents = st.Contents
	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listDirectoryResponsePb struct {
	Contents      []DirectoryEntry `json:"contents,omitempty"`
	NextPageToken string           `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDirectoryResponseFromPb(pb *listDirectoryResponsePb) (*ListDirectoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDirectoryResponse{}
	st.Contents = pb.Contents
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDirectoryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDirectoryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listStatusResponseToPb(st *ListStatusResponse) (*listStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStatusResponsePb{}
	pb.Files = st.Files

	return pb, nil
}

type listStatusResponsePb struct {
	Files []FileInfo `json:"files,omitempty"`
}

func listStatusResponseFromPb(pb *listStatusResponsePb) (*ListStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStatusResponse{}
	st.Files = pb.Files

	return st, nil
}

func mkDirsToPb(st *MkDirs) (*mkDirsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mkDirsPb{}
	pb.Path = st.Path

	return pb, nil
}

type mkDirsPb struct {
	Path string `json:"path"`
}

func mkDirsFromPb(pb *mkDirsPb) (*MkDirs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MkDirs{}
	st.Path = pb.Path

	return st, nil
}

func moveToPb(st *Move) (*movePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &movePb{}
	pb.DestinationPath = st.DestinationPath
	pb.SourcePath = st.SourcePath

	return pb, nil
}

type movePb struct {
	DestinationPath string `json:"destination_path"`
	SourcePath      string `json:"source_path"`
}

func moveFromPb(pb *movePb) (*Move, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Move{}
	st.DestinationPath = pb.DestinationPath
	st.SourcePath = pb.SourcePath

	return st, nil
}

func putToPb(st *Put) (*putPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putPb{}
	pb.Contents = st.Contents
	pb.Overwrite = st.Overwrite
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type putPb struct {
	Contents  string `json:"contents,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty"`
	Path      string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func putFromPb(pb *putPb) (*Put, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Put{}
	st.Contents = pb.Contents
	st.Overwrite = pb.Overwrite
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *putPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st putPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func readDbfsRequestToPb(st *ReadDbfsRequest) (*readDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &readDbfsRequestPb{}
	pb.Length = st.Length
	pb.Offset = st.Offset
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type readDbfsRequestPb struct {
	Length int64  `json:"-" url:"length,omitempty"`
	Offset int64  `json:"-" url:"offset,omitempty"`
	Path   string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func readDbfsRequestFromPb(pb *readDbfsRequestPb) (*ReadDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadDbfsRequest{}
	st.Length = pb.Length
	st.Offset = pb.Offset
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *readDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st readDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func readResponseToPb(st *ReadResponse) (*readResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &readResponsePb{}
	pb.BytesRead = st.BytesRead
	pb.Data = st.Data

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type readResponsePb struct {
	BytesRead int64  `json:"bytes_read,omitempty"`
	Data      string `json:"data,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func readResponseFromPb(pb *readResponsePb) (*ReadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadResponse{}
	st.BytesRead = pb.BytesRead
	st.Data = pb.Data

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *readResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st readResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func uploadRequestToPb(st *UploadRequest) (*uploadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &uploadRequestPb{}
	pb.Contents = st.Contents
	pb.FilePath = st.FilePath
	pb.Overwrite = st.Overwrite

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type uploadRequestPb struct {
	Contents  io.ReadCloser `json:"-"`
	FilePath  string        `json:"-" url:"-"`
	Overwrite bool          `json:"-" url:"overwrite,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func uploadRequestFromPb(pb *uploadRequestPb) (*UploadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UploadRequest{}
	st.Contents = pb.Contents
	st.FilePath = pb.FilePath
	st.Overwrite = pb.Overwrite

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *uploadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st uploadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339Nano)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339Nano, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	if *s == "" {
		return &[]string{}, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
