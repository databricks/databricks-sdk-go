// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
	libraryv2 "github.com/databricks/databricks-sdk-go/protos/library/v2"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type GetLibraryRequest struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetLibraryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLibraryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListLibrariesRequest struct {
	PageSize int32 `json:"pageSize,omitempty"`

	PageToken string `json:"pageToken,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListLibrariesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListLibrariesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListLibrariesResponse struct {
	Libraries []*libraryv2.Library `json:"libraries,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListLibrariesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListLibrariesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateLibraryRequest struct {
	Library *libraryv2.Library `json:"library,omitempty"`

	UpdateMask *fieldmaskpb.FieldMask `json:"updateMask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateLibraryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateLibraryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type CreateLibraryRequest struct {
	Library *libraryv2.Library `json:"library,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateLibraryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateLibraryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Library struct {
	Name string `json:"name,omitempty"`

	CreatedAt int64 `json:"createdAt,omitempty"`

	Description string `json:"description,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Library) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Library) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteLibraryRequest struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteLibraryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteLibraryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
