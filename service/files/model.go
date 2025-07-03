// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/google/go-querystring/query"
)

type AddBlock struct {
	// The base64-encoded data to append to the stream. This has a limit of 1
	// MB.
	// Wire name: 'data'
	Data string `json:"data"`
	// The handle on an open stream.
	// Wire name: 'handle'
	Handle int64 `json:"handle"`
}

func (st *AddBlock) EncodeValues(key string, v *url.Values) error {
	pb, err := addBlockToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Close struct {
	// The handle on an open stream.
	// Wire name: 'handle'
	Handle int64 `json:"handle"`
}

func (st *Close) EncodeValues(key string, v *url.Values) error {
	pb, err := closeToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Create struct {
	// The flag that specifies whether to overwrite existing file/files.
	// Wire name: 'overwrite'
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path.
	// Wire name: 'path'
	Path string `json:"path"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Create) EncodeValues(key string, v *url.Values) error {
	pb, err := createToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" tf:"-"`
}

func (st *CreateDirectoryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createDirectoryRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type CreateResponse struct {
	// Handle which should subsequently be passed into the AddBlock and Close
	// calls when writing to a file through a stream.
	// Wire name: 'handle'
	Handle int64 `json:"handle,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *CreateResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := createResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Delete struct {
	// The path of the file or directory to delete. The path should be the
	// absolute DBFS path.
	// Wire name: 'path'
	Path string `json:"path"`
	// Whether or not to recursively delete the directory's contents. Deleting
	// empty directories can be done without providing the recursive flag.
	// Wire name: 'recursive'
	Recursive bool `json:"recursive,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Delete) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteDirectoryRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" tf:"-"`
}

func (st *DeleteDirectoryRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteDirectoryRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DeleteFileRequest struct {
	// The absolute path of the file.
	FilePath string `json:"-" tf:"-"`
}

func (st *DeleteFileRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteFileRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DirectoryEntry struct {
	// The length of the file in bytes. This field is omitted for directories.
	// Wire name: 'file_size'
	FileSize int64 `json:"file_size,omitempty"`
	// True if the path is a directory.
	// Wire name: 'is_directory'
	IsDirectory bool `json:"is_directory,omitempty"`
	// Last modification time of given file in milliseconds since unix epoch.
	// Wire name: 'last_modified'
	LastModified int64 `json:"last_modified,omitempty"`
	// The name of the file or directory. This is the last component of the
	// path.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The absolute path of the file or directory.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DirectoryEntry) EncodeValues(key string, v *url.Values) error {
	pb, err := directoryEntryToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DownloadRequest struct {
	// The absolute path of the file.
	FilePath string `json:"-" tf:"-"`
}

func (st *DownloadRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := downloadRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type DownloadResponse struct {
	// The length of the HTTP response body in bytes.
	ContentLength int64 `json:"-" tf:"-"`

	ContentType string `json:"-" tf:"-"`

	Contents io.ReadCloser `json:"-" tf:"-"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *DownloadResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := downloadResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type FileInfo struct {
	// The length of the file in bytes. This field is omitted for directories.
	// Wire name: 'file_size'
	FileSize int64 `json:"file_size,omitempty"`
	// True if the path is a directory.
	// Wire name: 'is_dir'
	IsDir bool `json:"is_dir,omitempty"`
	// Last modification time of given file in milliseconds since epoch.
	// Wire name: 'modification_time'
	ModificationTime int64 `json:"modification_time,omitempty"`
	// The absolute path of the file or directory.
	// Wire name: 'path'
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *FileInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := fileInfoToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetDirectoryMetadataRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" tf:"-"`
}

func (st *GetDirectoryMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getDirectoryMetadataRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetMetadataRequest struct {
	// The absolute path of the file.
	FilePath string `json:"-" tf:"-"`
}

func (st *GetMetadataRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getMetadataRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetMetadataResponse struct {
	// The length of the HTTP response body in bytes.
	ContentLength int64 `json:"-" tf:"-"`

	ContentType string `json:"-" tf:"-"`
	// The last modified time of the file in HTTP-date (RFC 7231) format.
	LastModified string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *GetMetadataResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getMetadataResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type GetStatusRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `json:"-" tf:"-"`
}

func (st *GetStatusRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getStatusRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListDbfsRequest struct {
	// The path of the file or directory. The path should be the absolute DBFS
	// path.
	Path string `json:"-" tf:"-"`
}

func (st *ListDbfsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listDbfsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListDirectoryContentsRequest struct {
	// The absolute path of a directory.
	DirectoryPath string `json:"-" tf:"-"`
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
	PageSize int64 `json:"-" tf:"-"`
	// An opaque page token which was the `next_page_token` in the response of
	// the previous request to list the contents of this directory. Provide this
	// token to retrieve the next page of directory entries. When providing a
	// `page_token`, all other parameters provided to the request must match the
	// previous request. To list all of the entries in a directory, it is
	// necessary to continue requesting pages of entries until the response
	// contains no `next_page_token`. Note that the number of entries returned
	// must not be used to determine when the listing is complete.
	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListDirectoryContentsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := listDirectoryContentsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListDirectoryResponse struct {
	// Array of DirectoryEntry.
	// Wire name: 'contents'
	Contents []DirectoryEntry `json:"contents,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListDirectoryResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listDirectoryResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ListStatusResponse struct {
	// A list of FileInfo's that describe contents of directory or file. See
	// example above.
	// Wire name: 'files'
	Files []FileInfo `json:"files,omitempty"`
}

func (st *ListStatusResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listStatusResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type MkDirs struct {
	// The path of the new directory. The path should be the absolute DBFS path.
	// Wire name: 'path'
	Path string `json:"path"`
}

func (st *MkDirs) EncodeValues(key string, v *url.Values) error {
	pb, err := mkDirsToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Move struct {
	// The destination path of the file or directory. The path should be the
	// absolute DBFS path.
	// Wire name: 'destination_path'
	DestinationPath string `json:"destination_path"`
	// The source path of the file or directory. The path should be the absolute
	// DBFS path.
	// Wire name: 'source_path'
	SourcePath string `json:"source_path"`
}

func (st *Move) EncodeValues(key string, v *url.Values) error {
	pb, err := moveToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type Put struct {
	// This parameter might be absent, and instead a posted file will be used.
	// Wire name: 'contents'
	Contents string `json:"contents,omitempty"`
	// The flag that specifies whether to overwrite existing file/files.
	// Wire name: 'overwrite'
	Overwrite bool `json:"overwrite,omitempty"`
	// The path of the new file. The path should be the absolute DBFS path.
	// Wire name: 'path'
	Path string `json:"path"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *Put) EncodeValues(key string, v *url.Values) error {
	pb, err := putToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ReadDbfsRequest struct {
	// The number of bytes to read starting from the offset. This has a limit of
	// 1 MB, and a default value of 0.5 MB.
	Length int64 `json:"-" tf:"-"`
	// The offset to read from in bytes.
	Offset int64 `json:"-" tf:"-"`
	// The path of the file to read. The path should be the absolute DBFS path.
	Path string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ReadDbfsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := readDbfsRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type ReadResponse struct {
	// The number of bytes read (could be less than ``length`` if we hit end of
	// file). This refers to number of bytes read in unencoded version (response
	// data is base64-encoded).
	// Wire name: 'bytes_read'
	BytesRead int64 `json:"bytes_read,omitempty"`
	// The base64-encoded contents of the file read.
	// Wire name: 'data'
	Data string `json:"data,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ReadResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := readResponseToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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

type UploadRequest struct {
	Contents io.ReadCloser `json:"-" tf:"-"`
	// The absolute path of the file.
	FilePath string `json:"-" tf:"-"`
	// If true or unspecified, an existing file will be overwritten. If false,
	// an error will be returned if the path points to an existing file.
	Overwrite bool `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *UploadRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := uploadRequestToPb(st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
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
