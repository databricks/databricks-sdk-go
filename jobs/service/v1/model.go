// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v1

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateJobRequest struct {
	Job *Job `json:"-"`
}

type DeleteJobRequest struct {
	Name string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetJobRequest struct {
	Name string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Job struct {
	CreatedAt int64 `json:"-"`

	Description string `json:"-"`

	Name string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Job) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Job) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListJobsRequest struct {
	PageSize int `json:"-"`

	PageToken string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListJobsResponse struct {
	Jobs []Job `json:"-"`

	NextPageToken string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateJobRequest struct {
	Job *Job `json:"-"`

	UpdateMask string `json:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
