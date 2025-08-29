// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateLibraryRequest struct {
	Library *Library `json:"-"`
}

type DeleteLibraryRequest struct {
	Name string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteLibraryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteLibraryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetLibraryRequest struct {
	Name string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetLibraryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetLibraryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Library struct {
	CreatedAt int64 `json:"-"`

	Description string `json:"-"`

	Name string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Library) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Library) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListLibrariesRequest struct {
	PageSize int `json:"-"`

	PageToken string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListLibrariesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListLibrariesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListLibrariesResponse struct {
	Libraries []Library `json:"-"`

	NextPageToken string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListLibrariesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListLibrariesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateLibraryRequest struct {
	Library *Library `json:"-"`

	UpdateMask string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateLibraryRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateLibraryRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
