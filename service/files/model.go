// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/files/filespb"
)

type AddBlock struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	// Wire name: 'data'
	Data string ``
	// The handle on an open stream.
	// Wire name: 'handle'
	Handle int64 ``
}

func (st AddBlock) MarshalJSON() ([]byte, error) {
	pb, err := AddBlockToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AddBlock) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.AddBlockPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AddBlockFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AddBlockToPb(st *AddBlock) (*filespb.AddBlockPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.AddBlockPb{}
	pb.Data = st.Data
	pb.Handle = st.Handle

	return pb, nil
}

func AddBlockFromPb(pb *filespb.AddBlockPb) (*AddBlock, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddBlock{}
	st.Data = pb.Data
	st.Handle = pb.Handle

	return st, nil
}

type Close struct {
	// The handle on an open stream.
	// Wire name: 'handle'
	Handle int64 ``
}

func (st Close) MarshalJSON() ([]byte, error) {
	pb, err := CloseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Close) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.ClosePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CloseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CloseToPb(st *Close) (*filespb.ClosePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.ClosePb{}
	pb.Handle = st.Handle

	return pb, nil
}

func CloseFromPb(pb *filespb.ClosePb) (*Close, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Close{}
	st.Handle = pb.Handle

	return st, nil
}

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	// Wire name: 'overwrite'
	Overwrite bool ``
	// The path of the new file. The path should be the absolute DBFS path.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Create) MarshalJSON() ([]byte, error) {
	pb, err := CreateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Create) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.CreatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateToPb(st *Create) (*filespb.CreatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.CreatePb{}
	pb.Overwrite = st.Overwrite
	pb.Path = st.Path

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateFromPb(pb *filespb.CreatePb) (*Create, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Create{}
	st.Overwrite = pb.Overwrite
	st.Path = pb.Path

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateDirectoryRequest struct {
	// The absolute path of a directory.
	// Wire name: 'directory_path'
	DirectoryPath string `tf:"-"`
}

func (st CreateDirectoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateDirectoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateDirectoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.CreateDirectoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateDirectoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateDirectoryRequestToPb(st *CreateDirectoryRequest) (*filespb.CreateDirectoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.CreateDirectoryRequestPb{}
	pb.DirectoryPath = st.DirectoryPath

	return pb, nil
}

func CreateDirectoryRequestFromPb(pb *filespb.CreateDirectoryRequestPb) (*CreateDirectoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDirectoryRequest{}
	st.DirectoryPath = pb.DirectoryPath

	return st, nil
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	// Wire name: 'handle'
	Handle          int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.CreateResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateResponseToPb(st *CreateResponse) (*filespb.CreateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.CreateResponsePb{}
	pb.Handle = st.Handle

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateResponseFromPb(pb *filespb.CreateResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	st.Handle = pb.Handle

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	// Wire name: 'path'
	Path string ``
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	// Wire name: 'recursive'
	Recursive       bool     ``
	ForceSendFields []string `tf:"-"`
}

func (st Delete) MarshalJSON() ([]byte, error) {
	pb, err := DeleteToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Delete) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.DeletePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteToPb(st *Delete) (*filespb.DeletePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.DeletePb{}
	pb.Path = st.Path
	pb.Recursive = st.Recursive

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteFromPb(pb *filespb.DeletePb) (*Delete, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Delete{}
	st.Path = pb.Path
	st.Recursive = pb.Recursive

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteDirectoryRequest struct {
	// The absolute path of a directory.
	// Wire name: 'directory_path'
	DirectoryPath string `tf:"-"`
}

func (st DeleteDirectoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteDirectoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteDirectoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.DeleteDirectoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteDirectoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteDirectoryRequestToPb(st *DeleteDirectoryRequest) (*filespb.DeleteDirectoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.DeleteDirectoryRequestPb{}
	pb.DirectoryPath = st.DirectoryPath

	return pb, nil
}

func DeleteDirectoryRequestFromPb(pb *filespb.DeleteDirectoryRequestPb) (*DeleteDirectoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDirectoryRequest{}
	st.DirectoryPath = pb.DirectoryPath

	return st, nil
}

type DeleteFileRequest struct {
	// The absolute path of the file.
	// Wire name: 'file_path'
	FilePath string `tf:"-"`
}

func (st DeleteFileRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteFileRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteFileRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.DeleteFileRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteFileRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteFileRequestToPb(st *DeleteFileRequest) (*filespb.DeleteFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.DeleteFileRequestPb{}
	pb.FilePath = st.FilePath

	return pb, nil
}

func DeleteFileRequestFromPb(pb *filespb.DeleteFileRequestPb) (*DeleteFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFileRequest{}
	st.FilePath = pb.FilePath

	return st, nil
}

type DirectoryEntry struct {
	// The length of the file in bytes. This field is omitted for directories.
	// Wire name: 'file_size'
	FileSize int64 ``
	// True if the path is a directory.
	// Wire name: 'is_directory'
	IsDirectory bool ``
	// Last modification time of given file in milliseconds since unix epoch.
	// Wire name: 'last_modified'
	LastModified int64 ``
	// The name of the file or directory. This is the last component of the
	// path.
	// Wire name: 'name'
	Name string ``
	// The absolute path of the file or directory.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st DirectoryEntry) MarshalJSON() ([]byte, error) {
	pb, err := DirectoryEntryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DirectoryEntry) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.DirectoryEntryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DirectoryEntryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DirectoryEntryToPb(st *DirectoryEntry) (*filespb.DirectoryEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.DirectoryEntryPb{}
	pb.FileSize = st.FileSize
	pb.IsDirectory = st.IsDirectory
	pb.LastModified = st.LastModified
	pb.Name = st.Name
	pb.Path = st.Path

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DirectoryEntryFromPb(pb *filespb.DirectoryEntryPb) (*DirectoryEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DirectoryEntry{}
	st.FileSize = pb.FileSize
	st.IsDirectory = pb.IsDirectory
	st.LastModified = pb.LastModified
	st.Name = pb.Name
	st.Path = pb.Path

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DownloadRequest struct {
	// The absolute path of the file.
	// Wire name: 'file_path'
	FilePath string `tf:"-"`
}

func (st DownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := DownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.DownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DownloadRequestToPb(st *DownloadRequest) (*filespb.DownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.DownloadRequestPb{}
	pb.FilePath = st.FilePath

	return pb, nil
}

func DownloadRequestFromPb(pb *filespb.DownloadRequestPb) (*DownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadRequest{}
	st.FilePath = pb.FilePath

	return st, nil
}

type DownloadResponse struct {
	// The length of the HTTP response body in bytes.
	// Wire name: 'content-length'
	ContentLength int64 `tf:"-"`

	// Wire name: 'content-type'
	ContentType string `tf:"-"`

	// Wire name: 'contents'
	Contents io.ReadCloser `tf:"-"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	// Wire name: 'last-modified'
	LastModified    string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st DownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := DownloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.DownloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DownloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DownloadResponseToPb(st *DownloadResponse) (*filespb.DownloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.DownloadResponsePb{}
	pb.ContentLength = st.ContentLength
	pb.ContentType = st.ContentType
	pb.Contents = st.Contents
	pb.LastModified = st.LastModified

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DownloadResponseFromPb(pb *filespb.DownloadResponsePb) (*DownloadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadResponse{}
	st.ContentLength = pb.ContentLength
	st.ContentType = pb.ContentType
	st.Contents = pb.Contents
	st.LastModified = pb.LastModified

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FileInfo struct {
	// The length of the file in bytes. This field is omitted for directories.
	// Wire name: 'file_size'
	FileSize int64 ``
	// True if the path is a directory.
	// Wire name: 'is_dir'
	IsDir bool ``
	// Last modification time of given file in milliseconds since epoch.
	// Wire name: 'modification_time'
	ModificationTime int64 ``
	// The absolute path of the file or directory.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st FileInfo) MarshalJSON() ([]byte, error) {
	pb, err := FileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.FileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FileInfoToPb(st *FileInfo) (*filespb.FileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.FileInfoPb{}
	pb.FileSize = st.FileSize
	pb.IsDir = st.IsDir
	pb.ModificationTime = st.ModificationTime
	pb.Path = st.Path

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FileInfoFromPb(pb *filespb.FileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	st.FileSize = pb.FileSize
	st.IsDir = pb.IsDir
	st.ModificationTime = pb.ModificationTime
	st.Path = pb.Path

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetDirectoryMetadataRequest struct {
	// The absolute path of a directory.
	// Wire name: 'directory_path'
	DirectoryPath string `tf:"-"`
}

func (st GetDirectoryMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetDirectoryMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetDirectoryMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.GetDirectoryMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetDirectoryMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetDirectoryMetadataRequestToPb(st *GetDirectoryMetadataRequest) (*filespb.GetDirectoryMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.GetDirectoryMetadataRequestPb{}
	pb.DirectoryPath = st.DirectoryPath

	return pb, nil
}

func GetDirectoryMetadataRequestFromPb(pb *filespb.GetDirectoryMetadataRequestPb) (*GetDirectoryMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDirectoryMetadataRequest{}
	st.DirectoryPath = pb.DirectoryPath

	return st, nil
}

type GetMetadataRequest struct {
	// The absolute path of the file.
	// Wire name: 'file_path'
	FilePath string `tf:"-"`
}

func (st GetMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.GetMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetMetadataRequestToPb(st *GetMetadataRequest) (*filespb.GetMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.GetMetadataRequestPb{}
	pb.FilePath = st.FilePath

	return pb, nil
}

func GetMetadataRequestFromPb(pb *filespb.GetMetadataRequestPb) (*GetMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetadataRequest{}
	st.FilePath = pb.FilePath

	return st, nil
}

type GetMetadataResponse struct {
	// The length of the HTTP response body in bytes.
	// Wire name: 'content-length'
	ContentLength int64 `tf:"-"`

	// Wire name: 'content-type'
	ContentType string `tf:"-"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	// Wire name: 'last-modified'
	LastModified    string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st GetMetadataResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetMetadataResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetMetadataResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.GetMetadataResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetMetadataResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetMetadataResponseToPb(st *GetMetadataResponse) (*filespb.GetMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.GetMetadataResponsePb{}
	pb.ContentLength = st.ContentLength
	pb.ContentType = st.ContentType
	pb.LastModified = st.LastModified

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetMetadataResponseFromPb(pb *filespb.GetMetadataResponsePb) (*GetMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetadataResponse{}
	st.ContentLength = pb.ContentLength
	st.ContentType = pb.ContentType
	st.LastModified = pb.LastModified

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	// Wire name: 'path'
	Path string `tf:"-"`
}

func (st GetStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.GetStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetStatusRequestToPb(st *GetStatusRequest) (*filespb.GetStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.GetStatusRequestPb{}
	pb.Path = st.Path

	return pb, nil
}

func GetStatusRequestFromPb(pb *filespb.GetStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	st.Path = pb.Path

	return st, nil
}

type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	// Wire name: 'path'
	Path string `tf:"-"`
}

func (st ListDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.ListDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListDbfsRequestToPb(st *ListDbfsRequest) (*filespb.ListDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.ListDbfsRequestPb{}
	pb.Path = st.Path

	return pb, nil
}

func ListDbfsRequestFromPb(pb *filespb.ListDbfsRequestPb) (*ListDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDbfsRequest{}
	st.Path = pb.Path

	return st, nil
}

type ListDirectoryContentsRequest struct {
	// The absolute path of a directory.
	// Wire name: 'directory_path'
	DirectoryPath string `tf:"-"`
	// The maximum number of directory entries to return. The response may
	// contain fewer entries. If the response contains a `next_page_token`,
	// there may be more entries, even if fewer than `page_size` entries are in
	// the response.
	//
	// We recommend not to set this value unless you are intentionally listing
	// less than the complete directory contents.
	//
	// If unspecified, at most 1000 directory entries will be returned. The
	// maximum value is 1000. Values above 1000 will be coerced to 1000.
	// Wire name: 'page_size'
	PageSize int64 `tf:"-"`
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the contents of this directory. Provide this
	// token to retrieve the next page of directory entries. When providing a
	// `page_token`, all other parameters provided to the request must match the
	// previous request. To list all of the entries in a directory, it is
	// necessary to continue requesting pages of entries until the response
	// contains no `next_page_token`. Note that the number of entries returned
	// must not be used to determine when the listing is complete.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListDirectoryContentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListDirectoryContentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListDirectoryContentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.ListDirectoryContentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListDirectoryContentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListDirectoryContentsRequestToPb(st *ListDirectoryContentsRequest) (*filespb.ListDirectoryContentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.ListDirectoryContentsRequestPb{}
	pb.DirectoryPath = st.DirectoryPath
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListDirectoryContentsRequestFromPb(pb *filespb.ListDirectoryContentsRequestPb) (*ListDirectoryContentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDirectoryContentsRequest{}
	st.DirectoryPath = pb.DirectoryPath
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	// Wire name: 'contents'
	Contents []DirectoryEntry ``
	// A token, which can be sent as `page_token` to retrieve the next page.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListDirectoryResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListDirectoryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListDirectoryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.ListDirectoryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListDirectoryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListDirectoryResponseToPb(st *ListDirectoryResponse) (*filespb.ListDirectoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.ListDirectoryResponsePb{}

	var contentsPb []filespb.DirectoryEntryPb
	for _, item := range st.Contents {
		itemPb, err := DirectoryEntryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			contentsPb = append(contentsPb, *itemPb)
		}
	}
	pb.Contents = contentsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListDirectoryResponseFromPb(pb *filespb.ListDirectoryResponsePb) (*ListDirectoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDirectoryResponse{}

	var contentsField []DirectoryEntry
	for _, itemPb := range pb.Contents {
		item, err := DirectoryEntryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			contentsField = append(contentsField, *item)
		}
	}
	st.Contents = contentsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	// Wire name: 'files'
	Files []FileInfo ``
}

func (st ListStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.ListStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListStatusResponseToPb(st *ListStatusResponse) (*filespb.ListStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.ListStatusResponsePb{}

	var filesPb []filespb.FileInfoPb
	for _, item := range st.Files {
		itemPb, err := FileInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			filesPb = append(filesPb, *itemPb)
		}
	}
	pb.Files = filesPb

	return pb, nil
}

func ListStatusResponseFromPb(pb *filespb.ListStatusResponsePb) (*ListStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStatusResponse{}

	var filesField []FileInfo
	for _, itemPb := range pb.Files {
		item, err := FileInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			filesField = append(filesField, *item)
		}
	}
	st.Files = filesField

	return st, nil
}

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	// Wire name: 'path'
	Path string ``
}

func (st MkDirs) MarshalJSON() ([]byte, error) {
	pb, err := MkDirsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *MkDirs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.MkDirsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MkDirsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func MkDirsToPb(st *MkDirs) (*filespb.MkDirsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.MkDirsPb{}
	pb.Path = st.Path

	return pb, nil
}

func MkDirsFromPb(pb *filespb.MkDirsPb) (*MkDirs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MkDirs{}
	st.Path = pb.Path

	return st, nil
}

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	// Wire name: 'destination_path'
	DestinationPath string ``
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	// Wire name: 'source_path'
	SourcePath string ``
}

func (st Move) MarshalJSON() ([]byte, error) {
	pb, err := MoveToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Move) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.MovePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MoveFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func MoveToPb(st *Move) (*filespb.MovePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.MovePb{}
	pb.DestinationPath = st.DestinationPath
	pb.SourcePath = st.SourcePath

	return pb, nil
}

func MoveFromPb(pb *filespb.MovePb) (*Move, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Move{}
	st.DestinationPath = pb.DestinationPath
	st.SourcePath = pb.SourcePath

	return st, nil
}

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	// Wire name: 'contents'
	Contents string ``
	// The flag that specifies whether to overwrite existing file/files.
	// Wire name: 'overwrite'
	Overwrite bool ``
	// The path of the new file. The path should be the absolute DBFS path.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Put) MarshalJSON() ([]byte, error) {
	pb, err := PutToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Put) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.PutPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PutFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PutToPb(st *Put) (*filespb.PutPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.PutPb{}
	pb.Contents = st.Contents
	pb.Overwrite = st.Overwrite
	pb.Path = st.Path

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PutFromPb(pb *filespb.PutPb) (*Put, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Put{}
	st.Contents = pb.Contents
	st.Overwrite = pb.Overwrite
	st.Path = pb.Path

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ReadDbfsRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	// Wire name: 'length'
	Length int64 `tf:"-"`
	// The offset to read from in bytes.
	// Wire name: 'offset'
	Offset int64 `tf:"-"`
	// The path of the file to read. The path should be the absolute DBFS path.
	// Wire name: 'path'
	Path            string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ReadDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ReadDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ReadDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.ReadDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ReadDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ReadDbfsRequestToPb(st *ReadDbfsRequest) (*filespb.ReadDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.ReadDbfsRequestPb{}
	pb.Length = st.Length
	pb.Offset = st.Offset
	pb.Path = st.Path

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ReadDbfsRequestFromPb(pb *filespb.ReadDbfsRequestPb) (*ReadDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadDbfsRequest{}
	st.Length = pb.Length
	st.Offset = pb.Offset
	st.Path = pb.Path

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	// Wire name: 'bytes_read'
	BytesRead int64 ``
	// The base64-encoded contents of the file read.
	// Wire name: 'data'
	Data            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ReadResponse) MarshalJSON() ([]byte, error) {
	pb, err := ReadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ReadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.ReadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ReadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ReadResponseToPb(st *ReadResponse) (*filespb.ReadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.ReadResponsePb{}
	pb.BytesRead = st.BytesRead
	pb.Data = st.Data

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ReadResponseFromPb(pb *filespb.ReadResponsePb) (*ReadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadResponse{}
	st.BytesRead = pb.BytesRead
	st.Data = pb.Data

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UploadRequest struct {

	// Wire name: 'contents'
	Contents io.ReadCloser `tf:"-"`
	// The absolute path of the file.
	// Wire name: 'file_path'
	FilePath string `tf:"-"`
	// If true or unspecified, an existing file will be overwritten. If false,
	// an error will be returned if the path points to an existing file.
	// Wire name: 'overwrite'
	Overwrite       bool     `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st UploadRequest) MarshalJSON() ([]byte, error) {
	pb, err := UploadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UploadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &filespb.UploadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UploadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UploadRequestToPb(st *UploadRequest) (*filespb.UploadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filespb.UploadRequestPb{}
	pb.Contents = st.Contents
	pb.FilePath = st.FilePath
	pb.Overwrite = st.Overwrite

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UploadRequestFromPb(pb *filespb.UploadRequestPb) (*UploadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UploadRequest{}
	st.Contents = pb.Contents
	st.FilePath = pb.FilePath
	st.Overwrite = pb.Overwrite

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
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
	fm := strings.Split(*s, ",")
	return &fm, nil
}
