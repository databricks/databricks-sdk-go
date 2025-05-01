// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type AddBlock struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data string
	// The handle on an open stream.
	Handle int64
}

func addBlockToPb(st *AddBlock) (*addBlockPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addBlockPb{}
	dataPb, err := identity(&st.Data)
	if err != nil {
		return nil, err
	}
	if dataPb != nil {
		pb.Data = *dataPb
	}

	handlePb, err := identity(&st.Handle)
	if err != nil {
		return nil, err
	}
	if handlePb != nil {
		pb.Handle = *handlePb
	}

	return pb, nil
}

func (st *AddBlock) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &addBlockPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := addBlockFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AddBlock) MarshalJSON() ([]byte, error) {
	pb, err := addBlockToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type addBlockPb struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	Data string `json:"data"`
	// The handle on an open stream.
	Handle int64 `json:"handle"`
}

func addBlockFromPb(pb *addBlockPb) (*AddBlock, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddBlock{}
	dataField, err := identity(&pb.Data)
	if err != nil {
		return nil, err
	}
	if dataField != nil {
		st.Data = *dataField
	}
	handleField, err := identity(&pb.Handle)
	if err != nil {
		return nil, err
	}
	if handleField != nil {
		st.Handle = *handleField
	}

	return st, nil
}

type AddBlockResponse struct {
}

func addBlockResponseToPb(st *AddBlockResponse) (*addBlockResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &addBlockResponsePb{}

	return pb, nil
}

func (st *AddBlockResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &addBlockResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := addBlockResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AddBlockResponse) MarshalJSON() ([]byte, error) {
	pb, err := addBlockResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type addBlockResponsePb struct {
}

func addBlockResponseFromPb(pb *addBlockResponsePb) (*AddBlockResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AddBlockResponse{}

	return st, nil
}

type Close struct {
	// The handle on an open stream.
	Handle int64
}

func closeToPb(st *Close) (*closePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &closePb{}
	handlePb, err := identity(&st.Handle)
	if err != nil {
		return nil, err
	}
	if handlePb != nil {
		pb.Handle = *handlePb
	}

	return pb, nil
}

func (st *Close) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &closePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := closeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Close) MarshalJSON() ([]byte, error) {
	pb, err := closeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type closePb struct {
	// The handle on an open stream.
	Handle int64 `json:"handle"`
}

func closeFromPb(pb *closePb) (*Close, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Close{}
	handleField, err := identity(&pb.Handle)
	if err != nil {
		return nil, err
	}
	if handleField != nil {
		st.Handle = *handleField
	}

	return st, nil
}

type CloseResponse struct {
}

func closeResponseToPb(st *CloseResponse) (*closeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &closeResponsePb{}

	return pb, nil
}

func (st *CloseResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &closeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := closeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CloseResponse) MarshalJSON() ([]byte, error) {
	pb, err := closeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type closeResponsePb struct {
}

func closeResponseFromPb(pb *closeResponsePb) (*CloseResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CloseResponse{}

	return st, nil
}

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool
	// The path of the new file. The path should be the absolute DBFS path.
	Path string

	ForceSendFields []string
}

func createToPb(st *Create) (*createPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPb{}
	overwritePb, err := identity(&st.Overwrite)
	if err != nil {
		return nil, err
	}
	if overwritePb != nil {
		pb.Overwrite = *overwritePb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Create) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Create) MarshalJSON() ([]byte, error) {
	pb, err := createToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createPb struct {
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createFromPb(pb *createPb) (*Create, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Create{}
	overwriteField, err := identity(&pb.Overwrite)
	if err != nil {
		return nil, err
	}
	if overwriteField != nil {
		st.Overwrite = *overwriteField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Create a directory
type CreateDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath string
}

func createDirectoryRequestToPb(st *CreateDirectoryRequest) (*createDirectoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDirectoryRequestPb{}
	directoryPathPb, err := identity(&st.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathPb != nil {
		pb.DirectoryPath = *directoryPathPb
	}

	return pb, nil
}

func (st *CreateDirectoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createDirectoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createDirectoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateDirectoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := createDirectoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createDirectoryRequestPb struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" url:"-"`
}

func createDirectoryRequestFromPb(pb *createDirectoryRequestPb) (*CreateDirectoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDirectoryRequest{}
	directoryPathField, err := identity(&pb.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathField != nil {
		st.DirectoryPath = *directoryPathField
	}

	return st, nil
}

type CreateDirectoryResponse struct {
}

func createDirectoryResponseToPb(st *CreateDirectoryResponse) (*createDirectoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createDirectoryResponsePb{}

	return pb, nil
}

func (st *CreateDirectoryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createDirectoryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createDirectoryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateDirectoryResponse) MarshalJSON() ([]byte, error) {
	pb, err := createDirectoryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createDirectoryResponsePb struct {
}

func createDirectoryResponseFromPb(pb *createDirectoryResponsePb) (*CreateDirectoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateDirectoryResponse{}

	return st, nil
}

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle int64

	ForceSendFields []string
}

func createResponseToPb(st *CreateResponse) (*createResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createResponsePb{}
	handlePb, err := identity(&st.Handle)
	if err != nil {
		return nil, err
	}
	if handlePb != nil {
		pb.Handle = *handlePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateResponse) MarshalJSON() ([]byte, error) {
	pb, err := createResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createResponsePb struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	Handle int64 `json:"handle,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createResponseFromPb(pb *createResponsePb) (*CreateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateResponse{}
	handleField, err := identity(&pb.Handle)
	if err != nil {
		return nil, err
	}
	if handleField != nil {
		st.Handle = *handleField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path string
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive bool

	ForceSendFields []string
}

func deleteToPb(st *Delete) (*deletePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePb{}
	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	recursivePb, err := identity(&st.Recursive)
	if err != nil {
		return nil, err
	}
	if recursivePb != nil {
		pb.Recursive = *recursivePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Delete) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deletePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Delete) MarshalJSON() ([]byte, error) {
	pb, err := deleteToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deletePb struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	Path string `json:"path"`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteFromPb(pb *deletePb) (*Delete, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Delete{}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}
	recursiveField, err := identity(&pb.Recursive)
	if err != nil {
		return nil, err
	}
	if recursiveField != nil {
		st.Recursive = *recursiveField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deletePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deletePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Delete a directory
type DeleteDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath string
}

func deleteDirectoryRequestToPb(st *DeleteDirectoryRequest) (*deleteDirectoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDirectoryRequestPb{}
	directoryPathPb, err := identity(&st.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathPb != nil {
		pb.DirectoryPath = *directoryPathPb
	}

	return pb, nil
}

func (st *DeleteDirectoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDirectoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDirectoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDirectoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteDirectoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteDirectoryRequestPb struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" url:"-"`
}

func deleteDirectoryRequestFromPb(pb *deleteDirectoryRequestPb) (*DeleteDirectoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDirectoryRequest{}
	directoryPathField, err := identity(&pb.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathField != nil {
		st.DirectoryPath = *directoryPathField
	}

	return st, nil
}

type DeleteDirectoryResponse struct {
}

func deleteDirectoryResponseToPb(st *DeleteDirectoryResponse) (*deleteDirectoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteDirectoryResponsePb{}

	return pb, nil
}

func (st *DeleteDirectoryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteDirectoryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteDirectoryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteDirectoryResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteDirectoryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteDirectoryResponsePb struct {
}

func deleteDirectoryResponseFromPb(pb *deleteDirectoryResponsePb) (*DeleteDirectoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteDirectoryResponse{}

	return st, nil
}

// Delete a file
type DeleteFileRequest struct {
	// The absolute path of the file.
	FilePath string
}

func deleteFileRequestToPb(st *DeleteFileRequest) (*deleteFileRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteFileRequestPb{}
	filePathPb, err := identity(&st.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathPb != nil {
		pb.FilePath = *filePathPb
	}

	return pb, nil
}

func (st *DeleteFileRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteFileRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteFileRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteFileRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteFileRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteFileRequestPb struct {
	// The absolute path of the file.
	FilePath string `json:"-" url:"-"`
}

func deleteFileRequestFromPb(pb *deleteFileRequestPb) (*DeleteFileRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFileRequest{}
	filePathField, err := identity(&pb.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathField != nil {
		st.FilePath = *filePathField
	}

	return st, nil
}

type DeleteResponse struct {
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
}

func (st *DeleteResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
}

type DirectoryEntry struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize int64
	// True if the path is a directory.
	IsDirectory bool
	// Last modification time of given file in milliseconds since unix epoch.
	LastModified int64
	// The name of the file or directory. This is the last component of the
	// path.
	Name string
	// The absolute path of the file or directory.
	Path string

	ForceSendFields []string
}

func directoryEntryToPb(st *DirectoryEntry) (*directoryEntryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &directoryEntryPb{}
	fileSizePb, err := identity(&st.FileSize)
	if err != nil {
		return nil, err
	}
	if fileSizePb != nil {
		pb.FileSize = *fileSizePb
	}

	isDirectoryPb, err := identity(&st.IsDirectory)
	if err != nil {
		return nil, err
	}
	if isDirectoryPb != nil {
		pb.IsDirectory = *isDirectoryPb
	}

	lastModifiedPb, err := identity(&st.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedPb != nil {
		pb.LastModified = *lastModifiedPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DirectoryEntry) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &directoryEntryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := directoryEntryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DirectoryEntry) MarshalJSON() ([]byte, error) {
	pb, err := directoryEntryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type directoryEntryPb struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize int64 `json:"file_size,omitempty"`
	// True if the path is a directory.
	IsDirectory bool `json:"is_directory,omitempty"`
	// Last modification time of given file in milliseconds since unix epoch.
	LastModified int64 `json:"last_modified,omitempty"`
	// The name of the file or directory. This is the last component of the
	// path.
	Name string `json:"name,omitempty"`
	// The absolute path of the file or directory.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func directoryEntryFromPb(pb *directoryEntryPb) (*DirectoryEntry, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DirectoryEntry{}
	fileSizeField, err := identity(&pb.FileSize)
	if err != nil {
		return nil, err
	}
	if fileSizeField != nil {
		st.FileSize = *fileSizeField
	}
	isDirectoryField, err := identity(&pb.IsDirectory)
	if err != nil {
		return nil, err
	}
	if isDirectoryField != nil {
		st.IsDirectory = *isDirectoryField
	}
	lastModifiedField, err := identity(&pb.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedField != nil {
		st.LastModified = *lastModifiedField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *directoryEntryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st directoryEntryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Download a file
type DownloadRequest struct {
	// The absolute path of the file.
	FilePath string
}

func downloadRequestToPb(st *DownloadRequest) (*downloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &downloadRequestPb{}
	filePathPb, err := identity(&st.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathPb != nil {
		pb.FilePath = *filePathPb
	}

	return pb, nil
}

func (st *DownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &downloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := downloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := downloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type downloadRequestPb struct {
	// The absolute path of the file.
	FilePath string `json:"-" url:"-"`
}

func downloadRequestFromPb(pb *downloadRequestPb) (*DownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadRequest{}
	filePathField, err := identity(&pb.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathField != nil {
		st.FilePath = *filePathField
	}

	return st, nil
}

type DownloadResponse struct {
	// The length of the HTTP response body in bytes.
	ContentLength int64

	ContentType string

	Contents io.ReadCloser
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified string

	ForceSendFields []string
}

func downloadResponseToPb(st *DownloadResponse) (*downloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &downloadResponsePb{}
	contentLengthPb, err := identity(&st.ContentLength)
	if err != nil {
		return nil, err
	}
	if contentLengthPb != nil {
		pb.ContentLength = *contentLengthPb
	}

	contentTypePb, err := identity(&st.ContentType)
	if err != nil {
		return nil, err
	}
	if contentTypePb != nil {
		pb.ContentType = *contentTypePb
	}

	contentsPb, err := identity(&st.Contents)
	if err != nil {
		return nil, err
	}
	if contentsPb != nil {
		pb.Contents = *contentsPb
	}

	lastModifiedPb, err := identity(&st.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedPb != nil {
		pb.LastModified = *lastModifiedPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &downloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := downloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := downloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type downloadResponsePb struct {
	// The length of the HTTP response body in bytes.
	ContentLength int64 `json:"-" url:"-" header:"content-length,omitempty"`

	ContentType string `json:"-" url:"-" header:"content-type,omitempty"`

	Contents io.ReadCloser `json:"-"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified string `json:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func downloadResponseFromPb(pb *downloadResponsePb) (*DownloadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DownloadResponse{}
	contentLengthField, err := identity(&pb.ContentLength)
	if err != nil {
		return nil, err
	}
	if contentLengthField != nil {
		st.ContentLength = *contentLengthField
	}
	contentTypeField, err := identity(&pb.ContentType)
	if err != nil {
		return nil, err
	}
	if contentTypeField != nil {
		st.ContentType = *contentTypeField
	}
	contentsField, err := identity(&pb.Contents)
	if err != nil {
		return nil, err
	}
	if contentsField != nil {
		st.Contents = *contentsField
	}
	lastModifiedField, err := identity(&pb.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedField != nil {
		st.LastModified = *lastModifiedField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *downloadResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st downloadResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileInfo struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize int64
	// True if the path is a directory.
	IsDir bool
	// Last modification time of given file in milliseconds since epoch.
	ModificationTime int64
	// The absolute path of the file or directory.
	Path string

	ForceSendFields []string
}

func fileInfoToPb(st *FileInfo) (*fileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileInfoPb{}
	fileSizePb, err := identity(&st.FileSize)
	if err != nil {
		return nil, err
	}
	if fileSizePb != nil {
		pb.FileSize = *fileSizePb
	}

	isDirPb, err := identity(&st.IsDir)
	if err != nil {
		return nil, err
	}
	if isDirPb != nil {
		pb.IsDir = *isDirPb
	}

	modificationTimePb, err := identity(&st.ModificationTime)
	if err != nil {
		return nil, err
	}
	if modificationTimePb != nil {
		pb.ModificationTime = *modificationTimePb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileInfo) MarshalJSON() ([]byte, error) {
	pb, err := fileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type fileInfoPb struct {
	// The length of the file in bytes. This field is omitted for directories.
	FileSize int64 `json:"file_size,omitempty"`
	// True if the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// Last modification time of given file in milliseconds since epoch.
	ModificationTime int64 `json:"modification_time,omitempty"`
	// The absolute path of the file or directory.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileInfoFromPb(pb *fileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	fileSizeField, err := identity(&pb.FileSize)
	if err != nil {
		return nil, err
	}
	if fileSizeField != nil {
		st.FileSize = *fileSizeField
	}
	isDirField, err := identity(&pb.IsDir)
	if err != nil {
		return nil, err
	}
	if isDirField != nil {
		st.IsDir = *isDirField
	}
	modificationTimeField, err := identity(&pb.ModificationTime)
	if err != nil {
		return nil, err
	}
	if modificationTimeField != nil {
		st.ModificationTime = *modificationTimeField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get directory metadata
type GetDirectoryMetadataRequest struct {
	// The absolute path of a directory.
	DirectoryPath string
}

func getDirectoryMetadataRequestToPb(st *GetDirectoryMetadataRequest) (*getDirectoryMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDirectoryMetadataRequestPb{}
	directoryPathPb, err := identity(&st.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathPb != nil {
		pb.DirectoryPath = *directoryPathPb
	}

	return pb, nil
}

func (st *GetDirectoryMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDirectoryMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDirectoryMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDirectoryMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := getDirectoryMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getDirectoryMetadataRequestPb struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" url:"-"`
}

func getDirectoryMetadataRequestFromPb(pb *getDirectoryMetadataRequestPb) (*GetDirectoryMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDirectoryMetadataRequest{}
	directoryPathField, err := identity(&pb.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathField != nil {
		st.DirectoryPath = *directoryPathField
	}

	return st, nil
}

type GetDirectoryMetadataResponse struct {
}

func getDirectoryMetadataResponseToPb(st *GetDirectoryMetadataResponse) (*getDirectoryMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getDirectoryMetadataResponsePb{}

	return pb, nil
}

func (st *GetDirectoryMetadataResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getDirectoryMetadataResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getDirectoryMetadataResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetDirectoryMetadataResponse) MarshalJSON() ([]byte, error) {
	pb, err := getDirectoryMetadataResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getDirectoryMetadataResponsePb struct {
}

func getDirectoryMetadataResponseFromPb(pb *getDirectoryMetadataResponsePb) (*GetDirectoryMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetDirectoryMetadataResponse{}

	return st, nil
}

// Get file metadata
type GetMetadataRequest struct {
	// The absolute path of the file.
	FilePath string
}

func getMetadataRequestToPb(st *GetMetadataRequest) (*getMetadataRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetadataRequestPb{}
	filePathPb, err := identity(&st.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathPb != nil {
		pb.FilePath = *filePathPb
	}

	return pb, nil
}

func (st *GetMetadataRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getMetadataRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getMetadataRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetMetadataRequest) MarshalJSON() ([]byte, error) {
	pb, err := getMetadataRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getMetadataRequestPb struct {
	// The absolute path of the file.
	FilePath string `json:"-" url:"-"`
}

func getMetadataRequestFromPb(pb *getMetadataRequestPb) (*GetMetadataRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetadataRequest{}
	filePathField, err := identity(&pb.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathField != nil {
		st.FilePath = *filePathField
	}

	return st, nil
}

type GetMetadataResponse struct {
	// The length of the HTTP response body in bytes.
	ContentLength int64

	ContentType string
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified string

	ForceSendFields []string
}

func getMetadataResponseToPb(st *GetMetadataResponse) (*getMetadataResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetadataResponsePb{}
	contentLengthPb, err := identity(&st.ContentLength)
	if err != nil {
		return nil, err
	}
	if contentLengthPb != nil {
		pb.ContentLength = *contentLengthPb
	}

	contentTypePb, err := identity(&st.ContentType)
	if err != nil {
		return nil, err
	}
	if contentTypePb != nil {
		pb.ContentType = *contentTypePb
	}

	lastModifiedPb, err := identity(&st.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedPb != nil {
		pb.LastModified = *lastModifiedPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetMetadataResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getMetadataResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getMetadataResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetMetadataResponse) MarshalJSON() ([]byte, error) {
	pb, err := getMetadataResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getMetadataResponsePb struct {
	// The length of the HTTP response body in bytes.
	ContentLength int64 `json:"-" url:"-" header:"content-length,omitempty"`

	ContentType string `json:"-" url:"-" header:"content-type,omitempty"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified string `json:"-" url:"-" header:"last-modified,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getMetadataResponseFromPb(pb *getMetadataResponsePb) (*GetMetadataResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetadataResponse{}
	contentLengthField, err := identity(&pb.ContentLength)
	if err != nil {
		return nil, err
	}
	if contentLengthField != nil {
		st.ContentLength = *contentLengthField
	}
	contentTypeField, err := identity(&pb.ContentType)
	if err != nil {
		return nil, err
	}
	if contentTypeField != nil {
		st.ContentType = *contentTypeField
	}
	lastModifiedField, err := identity(&pb.LastModified)
	if err != nil {
		return nil, err
	}
	if lastModifiedField != nil {
		st.LastModified = *lastModifiedField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getMetadataResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getMetadataResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get the information of a file or directory
type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string
}

func getStatusRequestToPb(st *GetStatusRequest) (*getStatusRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getStatusRequestPb{}
	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	return pb, nil
}

func (st *GetStatusRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getStatusRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getStatusRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetStatusRequest) MarshalJSON() ([]byte, error) {
	pb, err := getStatusRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getStatusRequestPb struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `json:"-" url:"path"`
}

func getStatusRequestFromPb(pb *getStatusRequestPb) (*GetStatusRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetStatusRequest{}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	return st, nil
}

// List directory contents or file details
type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string
}

func listDbfsRequestToPb(st *ListDbfsRequest) (*listDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDbfsRequestPb{}
	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	return pb, nil
}

func (st *ListDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listDbfsRequestPb struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `json:"-" url:"path"`
}

func listDbfsRequestFromPb(pb *listDbfsRequestPb) (*ListDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDbfsRequest{}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	return st, nil
}

// List directory contents
type ListDirectoryContentsRequest struct {
	// The absolute path of a directory.
	DirectoryPath string
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
	PageSize int64
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the contents of this directory. Provide this
	// token to retrieve the next page of directory entries. When providing a
	// `page_token`, all other parameters provided to the request must match the
	// previous request. To list all of the entries in a directory, it is
	// necessary to continue requesting pages of entries until the response
	// contains no `next_page_token`. Note that the number of entries returned
	// must not be used to determine when the listing is complete.
	PageToken string

	ForceSendFields []string
}

func listDirectoryContentsRequestToPb(st *ListDirectoryContentsRequest) (*listDirectoryContentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDirectoryContentsRequestPb{}
	directoryPathPb, err := identity(&st.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathPb != nil {
		pb.DirectoryPath = *directoryPathPb
	}

	pageSizePb, err := identity(&st.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizePb != nil {
		pb.PageSize = *pageSizePb
	}

	pageTokenPb, err := identity(&st.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenPb != nil {
		pb.PageToken = *pageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListDirectoryContentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDirectoryContentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDirectoryContentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDirectoryContentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listDirectoryContentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listDirectoryContentsRequestPb struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" url:"-"`
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
	PageSize int64 `json:"-" url:"page_size,omitempty"`
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the contents of this directory. Provide this
	// token to retrieve the next page of directory entries. When providing a
	// `page_token`, all other parameters provided to the request must match the
	// previous request. To list all of the entries in a directory, it is
	// necessary to continue requesting pages of entries until the response
	// contains no `next_page_token`. Note that the number of entries returned
	// must not be used to determine when the listing is complete.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDirectoryContentsRequestFromPb(pb *listDirectoryContentsRequestPb) (*ListDirectoryContentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDirectoryContentsRequest{}
	directoryPathField, err := identity(&pb.DirectoryPath)
	if err != nil {
		return nil, err
	}
	if directoryPathField != nil {
		st.DirectoryPath = *directoryPathField
	}
	pageSizeField, err := identity(&pb.PageSize)
	if err != nil {
		return nil, err
	}
	if pageSizeField != nil {
		st.PageSize = *pageSizeField
	}
	pageTokenField, err := identity(&pb.PageToken)
	if err != nil {
		return nil, err
	}
	if pageTokenField != nil {
		st.PageToken = *pageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDirectoryContentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDirectoryContentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	Contents []DirectoryEntry
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken string

	ForceSendFields []string
}

func listDirectoryResponseToPb(st *ListDirectoryResponse) (*listDirectoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listDirectoryResponsePb{}

	var contentsPb []directoryEntryPb
	for _, item := range st.Contents {
		itemPb, err := directoryEntryToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			contentsPb = append(contentsPb, *itemPb)
		}
	}
	pb.Contents = contentsPb

	nextPageTokenPb, err := identity(&st.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenPb != nil {
		pb.NextPageToken = *nextPageTokenPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListDirectoryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listDirectoryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listDirectoryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListDirectoryResponse) MarshalJSON() ([]byte, error) {
	pb, err := listDirectoryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listDirectoryResponsePb struct {
	// Array of DirectoryEntry.
	Contents []directoryEntryPb `json:"contents,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listDirectoryResponseFromPb(pb *listDirectoryResponsePb) (*ListDirectoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListDirectoryResponse{}

	var contentsField []DirectoryEntry
	for _, itemPb := range pb.Contents {
		item, err := directoryEntryFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			contentsField = append(contentsField, *item)
		}
	}
	st.Contents = contentsField
	nextPageTokenField, err := identity(&pb.NextPageToken)
	if err != nil {
		return nil, err
	}
	if nextPageTokenField != nil {
		st.NextPageToken = *nextPageTokenField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listDirectoryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listDirectoryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files []FileInfo
}

func listStatusResponseToPb(st *ListStatusResponse) (*listStatusResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listStatusResponsePb{}

	var filesPb []fileInfoPb
	for _, item := range st.Files {
		itemPb, err := fileInfoToPb(&item)
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

func (st *ListStatusResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listStatusResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listStatusResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListStatusResponse) MarshalJSON() ([]byte, error) {
	pb, err := listStatusResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listStatusResponsePb struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	Files []fileInfoPb `json:"files,omitempty"`
}

func listStatusResponseFromPb(pb *listStatusResponsePb) (*ListStatusResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListStatusResponse{}

	var filesField []FileInfo
	for _, itemPb := range pb.Files {
		item, err := fileInfoFromPb(&itemPb)
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
	Path string
}

func mkDirsToPb(st *MkDirs) (*mkDirsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mkDirsPb{}
	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	return pb, nil
}

func (st *MkDirs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mkDirsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := mkDirsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MkDirs) MarshalJSON() ([]byte, error) {
	pb, err := mkDirsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type mkDirsPb struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	Path string `json:"path"`
}

func mkDirsFromPb(pb *mkDirsPb) (*MkDirs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MkDirs{}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	return st, nil
}

type MkDirsResponse struct {
}

func mkDirsResponseToPb(st *MkDirsResponse) (*mkDirsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mkDirsResponsePb{}

	return pb, nil
}

func (st *MkDirsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mkDirsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := mkDirsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MkDirsResponse) MarshalJSON() ([]byte, error) {
	pb, err := mkDirsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type mkDirsResponsePb struct {
}

func mkDirsResponseFromPb(pb *mkDirsResponsePb) (*MkDirsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MkDirsResponse{}

	return st, nil
}

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath string
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath string
}

func moveToPb(st *Move) (*movePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &movePb{}
	destinationPathPb, err := identity(&st.DestinationPath)
	if err != nil {
		return nil, err
	}
	if destinationPathPb != nil {
		pb.DestinationPath = *destinationPathPb
	}

	sourcePathPb, err := identity(&st.SourcePath)
	if err != nil {
		return nil, err
	}
	if sourcePathPb != nil {
		pb.SourcePath = *sourcePathPb
	}

	return pb, nil
}

func (st *Move) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &movePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := moveFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Move) MarshalJSON() ([]byte, error) {
	pb, err := moveToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type movePb struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	DestinationPath string `json:"destination_path"`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	SourcePath string `json:"source_path"`
}

func moveFromPb(pb *movePb) (*Move, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Move{}
	destinationPathField, err := identity(&pb.DestinationPath)
	if err != nil {
		return nil, err
	}
	if destinationPathField != nil {
		st.DestinationPath = *destinationPathField
	}
	sourcePathField, err := identity(&pb.SourcePath)
	if err != nil {
		return nil, err
	}
	if sourcePathField != nil {
		st.SourcePath = *sourcePathField
	}

	return st, nil
}

type MoveResponse struct {
}

func moveResponseToPb(st *MoveResponse) (*moveResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &moveResponsePb{}

	return pb, nil
}

func (st *MoveResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &moveResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := moveResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st MoveResponse) MarshalJSON() ([]byte, error) {
	pb, err := moveResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type moveResponsePb struct {
}

func moveResponseFromPb(pb *moveResponsePb) (*MoveResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &MoveResponse{}

	return st, nil
}

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents string
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool
	// The path of the new file. The path should be the absolute DBFS path.
	Path string

	ForceSendFields []string
}

func putToPb(st *Put) (*putPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putPb{}
	contentsPb, err := identity(&st.Contents)
	if err != nil {
		return nil, err
	}
	if contentsPb != nil {
		pb.Contents = *contentsPb
	}

	overwritePb, err := identity(&st.Overwrite)
	if err != nil {
		return nil, err
	}
	if overwritePb != nil {
		pb.Overwrite = *overwritePb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Put) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Put) MarshalJSON() ([]byte, error) {
	pb, err := putToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type putPb struct {
	// This parameter might be absent, and instead a posted file will be used.
	Contents string `json:"contents,omitempty"`
	// The flag that specifies whether to overwrite existing file/files.
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path.
	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func putFromPb(pb *putPb) (*Put, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Put{}
	contentsField, err := identity(&pb.Contents)
	if err != nil {
		return nil, err
	}
	if contentsField != nil {
		st.Contents = *contentsField
	}
	overwriteField, err := identity(&pb.Overwrite)
	if err != nil {
		return nil, err
	}
	if overwriteField != nil {
		st.Overwrite = *overwriteField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *putPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st putPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PutResponse struct {
}

func putResponseToPb(st *PutResponse) (*putResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putResponsePb{}

	return pb, nil
}

func (st *PutResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutResponse) MarshalJSON() ([]byte, error) {
	pb, err := putResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type putResponsePb struct {
}

func putResponseFromPb(pb *putResponsePb) (*PutResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutResponse{}

	return st, nil
}

// Get the contents of a file
type ReadDbfsRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length int64
	// The offset to read from in bytes.
	Offset int64
	// The path of the file to read. The path should be the absolute DBFS path.
	Path string

	ForceSendFields []string
}

func readDbfsRequestToPb(st *ReadDbfsRequest) (*readDbfsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &readDbfsRequestPb{}
	lengthPb, err := identity(&st.Length)
	if err != nil {
		return nil, err
	}
	if lengthPb != nil {
		pb.Length = *lengthPb
	}

	offsetPb, err := identity(&st.Offset)
	if err != nil {
		return nil, err
	}
	if offsetPb != nil {
		pb.Offset = *offsetPb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ReadDbfsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &readDbfsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := readDbfsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReadDbfsRequest) MarshalJSON() ([]byte, error) {
	pb, err := readDbfsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type readDbfsRequestPb struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length int64 `json:"-" url:"length,omitempty"`
	// The offset to read from in bytes.
	Offset int64 `json:"-" url:"offset,omitempty"`
	// The path of the file to read. The path should be the absolute DBFS path.
	Path string `json:"-" url:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func readDbfsRequestFromPb(pb *readDbfsRequestPb) (*ReadDbfsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadDbfsRequest{}
	lengthField, err := identity(&pb.Length)
	if err != nil {
		return nil, err
	}
	if lengthField != nil {
		st.Length = *lengthField
	}
	offsetField, err := identity(&pb.Offset)
	if err != nil {
		return nil, err
	}
	if offsetField != nil {
		st.Offset = *offsetField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *readDbfsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st readDbfsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead int64
	// The base64-encoded contents of the file read.
	Data string

	ForceSendFields []string
}

func readResponseToPb(st *ReadResponse) (*readResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &readResponsePb{}
	bytesReadPb, err := identity(&st.BytesRead)
	if err != nil {
		return nil, err
	}
	if bytesReadPb != nil {
		pb.BytesRead = *bytesReadPb
	}

	dataPb, err := identity(&st.Data)
	if err != nil {
		return nil, err
	}
	if dataPb != nil {
		pb.Data = *dataPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ReadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &readResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := readResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ReadResponse) MarshalJSON() ([]byte, error) {
	pb, err := readResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type readResponsePb struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	BytesRead int64 `json:"bytes_read,omitempty"`
	// The base64-encoded contents of the file read.
	Data string `json:"data,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func readResponseFromPb(pb *readResponsePb) (*ReadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReadResponse{}
	bytesReadField, err := identity(&pb.BytesRead)
	if err != nil {
		return nil, err
	}
	if bytesReadField != nil {
		st.BytesRead = *bytesReadField
	}
	dataField, err := identity(&pb.Data)
	if err != nil {
		return nil, err
	}
	if dataField != nil {
		st.Data = *dataField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *readResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st readResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Upload a file
type UploadRequest struct {
	Contents io.ReadCloser
	// The absolute path of the file.
	FilePath string
	// If true or unspecified, an existing file will be overwritten. If false,
	// an error will be returned if the path points to an existing file.
	Overwrite bool

	ForceSendFields []string
}

func uploadRequestToPb(st *UploadRequest) (*uploadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &uploadRequestPb{}
	contentsPb, err := identity(&st.Contents)
	if err != nil {
		return nil, err
	}
	if contentsPb != nil {
		pb.Contents = *contentsPb
	}

	filePathPb, err := identity(&st.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathPb != nil {
		pb.FilePath = *filePathPb
	}

	overwritePb, err := identity(&st.Overwrite)
	if err != nil {
		return nil, err
	}
	if overwritePb != nil {
		pb.Overwrite = *overwritePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UploadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &uploadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := uploadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UploadRequest) MarshalJSON() ([]byte, error) {
	pb, err := uploadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type uploadRequestPb struct {
	Contents io.ReadCloser `json:"-"`
	// The absolute path of the file.
	FilePath string `json:"-" url:"-"`
	// If true or unspecified, an existing file will be overwritten. If false,
	// an error will be returned if the path points to an existing file.
	Overwrite bool `json:"-" url:"overwrite,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func uploadRequestFromPb(pb *uploadRequestPb) (*UploadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UploadRequest{}
	contentsField, err := identity(&pb.Contents)
	if err != nil {
		return nil, err
	}
	if contentsField != nil {
		st.Contents = *contentsField
	}
	filePathField, err := identity(&pb.FilePath)
	if err != nil {
		return nil, err
	}
	if filePathField != nil {
		st.FilePath = *filePathField
	}
	overwriteField, err := identity(&pb.Overwrite)
	if err != nil {
		return nil, err
	}
	if overwriteField != nil {
		st.Overwrite = *overwriteField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *uploadRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st uploadRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UploadResponse struct {
}

func uploadResponseToPb(st *UploadResponse) (*uploadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &uploadResponsePb{}

	return pb, nil
}

func (st *UploadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &uploadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := uploadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UploadResponse) MarshalJSON() ([]byte, error) {
	pb, err := uploadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type uploadResponsePb struct {
}

func uploadResponseFromPb(pb *uploadResponsePb) (*UploadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UploadResponse{}

	return st, nil
}
