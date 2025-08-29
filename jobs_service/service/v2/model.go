// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package v2

import (
	"github.com/databricks/databricks-sdk-go/marshal"
	jobsv2 "github.com/databricks/databricks-sdk-go/protos/jobs/v2"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

type CreateJobRequest struct {
	Job *jobsv2.Job `json:"job,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *CreateJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type Job struct {
	Name string `json:"name,omitempty"`

	CreatedAt int64 `json:"createdAt,omitempty"`

	Description string `json:"description,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Job) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Job) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type DeleteJobRequest struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *DeleteJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GetJobRequest struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *GetJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListJobsRequest struct {
	PageSize int32 `json:"pageSize,omitempty"`

	PageToken string `json:"pageToken,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListJobsResponse struct {
	Jobs []*jobsv2.Job `json:"jobs,omitempty"`

	NextPageToken string `json:"nextPageToken,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListJobsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListJobsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateJobRequest struct {
	Job *jobsv2.Job `json:"job,omitempty"`

	UpdateMask *fieldmaskpb.FieldMask `json:"updateMask,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *UpdateJobRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s UpdateJobRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
