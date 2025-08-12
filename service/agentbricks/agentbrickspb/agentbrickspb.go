// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentbrickspb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CancelCustomLlmOptimizationRunRequestPb struct {
	Id string `json:"-" url:"-"`
}

type CreateCustomLlmRequestPb struct {
	AgentArtifactPath string      `json:"agent_artifact_path,omitempty"`
	Datasets          []DatasetPb `json:"datasets,omitempty"`
	Guidelines        []string    `json:"guidelines,omitempty"`
	Instructions      string      `json:"instructions"`
	Name              string      `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateCustomLlmRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateCustomLlmRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomLlmPb struct {
	AgentArtifactPath string      `json:"agent_artifact_path,omitempty"`
	CreationTime      string      `json:"creation_time,omitempty"`
	Creator           string      `json:"creator,omitempty"`
	Datasets          []DatasetPb `json:"datasets,omitempty"`
	EndpointName      string      `json:"endpoint_name,omitempty"`
	Guidelines        []string    `json:"guidelines,omitempty"`
	Id                string      `json:"id,omitempty"`
	Instructions      string      `json:"instructions"`
	Name              string      `json:"name"`
	OptimizationState StatePb     `json:"optimization_state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CustomLlmPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CustomLlmPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatasetPb struct {
	Table TablePb `json:"table"`
}

type DeleteCustomLlmRequestPb struct {
	Id string `json:"-" url:"-"`
}

type GetCustomLlmRequestPb struct {
	Id string `json:"-" url:"-"`
}

type StartCustomLlmOptimizationRunRequestPb struct {
	Id string `json:"-" url:"-"`
}

type StatePb string

const StateCancelled StatePb = `CANCELLED`

const StateCompleted StatePb = `COMPLETED`

const StateCreated StatePb = `CREATED`

const StateFailed StatePb = `FAILED`

const StatePending StatePb = `PENDING`

const StateRunning StatePb = `RUNNING`

type TablePb struct {
	RequestCol  string `json:"request_col"`
	ResponseCol string `json:"response_col,omitempty"`
	TablePath   string `json:"table_path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateCustomLlmRequestPb struct {
	CustomLlm  CustomLlmPb `json:"custom_llm"`
	Id         string      `json:"-" url:"-"`
	UpdateMask string      `json:"update_mask"`
}
