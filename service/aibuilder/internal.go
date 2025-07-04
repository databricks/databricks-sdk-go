// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aibuilder

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func cancelCustomLlmOptimizationRunRequestToPb(st *CancelCustomLlmOptimizationRunRequest) (*cancelCustomLlmOptimizationRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelCustomLlmOptimizationRunRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type cancelCustomLlmOptimizationRunRequestPb struct {
	Id string `json:"-" url:"-"`
}

func cancelCustomLlmOptimizationRunRequestFromPb(pb *cancelCustomLlmOptimizationRunRequestPb) (*CancelCustomLlmOptimizationRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelCustomLlmOptimizationRunRequest{}
	st.Id = pb.Id

	return st, nil
}

func cancelOptimizeResponseToPb(st *CancelOptimizeResponse) (*cancelOptimizeResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cancelOptimizeResponsePb{}

	return pb, nil
}

type cancelOptimizeResponsePb struct {
}

func cancelOptimizeResponseFromPb(pb *cancelOptimizeResponsePb) (*CancelOptimizeResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelOptimizeResponse{}

	return st, nil
}

func createCustomLlmRequestToPb(st *CreateCustomLlmRequest) (*createCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCustomLlmRequestPb{}
	pb.AgentArtifactPath = st.AgentArtifactPath
	pb.Datasets = st.Datasets
	pb.Guidelines = st.Guidelines
	pb.Instructions = st.Instructions
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createCustomLlmRequestPb struct {
	AgentArtifactPath string    `json:"agent_artifact_path,omitempty"`
	Datasets          []Dataset `json:"datasets,omitempty"`
	Guidelines        []string  `json:"guidelines,omitempty"`
	Instructions      string    `json:"instructions"`
	Name              string    `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createCustomLlmRequestFromPb(pb *createCustomLlmRequestPb) (*CreateCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomLlmRequest{}
	st.AgentArtifactPath = pb.AgentArtifactPath
	st.Datasets = pb.Datasets
	st.Guidelines = pb.Guidelines
	st.Instructions = pb.Instructions
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createCustomLlmRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createCustomLlmRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func customLlmToPb(st *CustomLlm) (*customLlmPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customLlmPb{}
	pb.AgentArtifactPath = st.AgentArtifactPath
	pb.CreationTime = st.CreationTime
	pb.Creator = st.Creator
	pb.Datasets = st.Datasets
	pb.EndpointName = st.EndpointName
	pb.Guidelines = st.Guidelines
	pb.Id = st.Id
	pb.Instructions = st.Instructions
	pb.Name = st.Name
	pb.OptimizationState = st.OptimizationState

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type customLlmPb struct {
	AgentArtifactPath string    `json:"agent_artifact_path,omitempty"`
	CreationTime      string    `json:"creation_time,omitempty"`
	Creator           string    `json:"creator,omitempty"`
	Datasets          []Dataset `json:"datasets,omitempty"`
	EndpointName      string    `json:"endpoint_name,omitempty"`
	Guidelines        []string  `json:"guidelines,omitempty"`
	Id                string    `json:"id,omitempty"`
	Instructions      string    `json:"instructions"`
	Name              string    `json:"name"`
	OptimizationState State     `json:"optimization_state"`

	ForceSendFields []string `json:"-" url:"-"`
}

func customLlmFromPb(pb *customLlmPb) (*CustomLlm, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomLlm{}
	st.AgentArtifactPath = pb.AgentArtifactPath
	st.CreationTime = pb.CreationTime
	st.Creator = pb.Creator
	st.Datasets = pb.Datasets
	st.EndpointName = pb.EndpointName
	st.Guidelines = pb.Guidelines
	st.Id = pb.Id
	st.Instructions = pb.Instructions
	st.Name = pb.Name
	st.OptimizationState = pb.OptimizationState

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *customLlmPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st customLlmPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func datasetToPb(st *Dataset) (*datasetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &datasetPb{}
	pb.Table = st.Table

	return pb, nil
}

type datasetPb struct {
	Table Table `json:"table"`
}

func datasetFromPb(pb *datasetPb) (*Dataset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dataset{}
	st.Table = pb.Table

	return st, nil
}

func deleteCustomLlmRequestToPb(st *DeleteCustomLlmRequest) (*deleteCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCustomLlmRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type deleteCustomLlmRequestPb struct {
	Id string `json:"-" url:"-"`
}

func deleteCustomLlmRequestFromPb(pb *deleteCustomLlmRequestPb) (*DeleteCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCustomLlmRequest{}
	st.Id = pb.Id

	return st, nil
}

func deleteCustomLlmResponseToPb(st *DeleteCustomLlmResponse) (*deleteCustomLlmResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCustomLlmResponsePb{}

	return pb, nil
}

type deleteCustomLlmResponsePb struct {
}

func deleteCustomLlmResponseFromPb(pb *deleteCustomLlmResponsePb) (*DeleteCustomLlmResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCustomLlmResponse{}

	return st, nil
}

func getCustomLlmRequestToPb(st *GetCustomLlmRequest) (*getCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCustomLlmRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type getCustomLlmRequestPb struct {
	Id string `json:"-" url:"-"`
}

func getCustomLlmRequestFromPb(pb *getCustomLlmRequestPb) (*GetCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomLlmRequest{}
	st.Id = pb.Id

	return st, nil
}

func startCustomLlmOptimizationRunRequestToPb(st *StartCustomLlmOptimizationRunRequest) (*startCustomLlmOptimizationRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startCustomLlmOptimizationRunRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

type startCustomLlmOptimizationRunRequestPb struct {
	Id string `json:"-" url:"-"`
}

func startCustomLlmOptimizationRunRequestFromPb(pb *startCustomLlmOptimizationRunRequestPb) (*StartCustomLlmOptimizationRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartCustomLlmOptimizationRunRequest{}
	st.Id = pb.Id

	return st, nil
}

func tableToPb(st *Table) (*tablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tablePb{}
	pb.RequestCol = st.RequestCol
	pb.ResponseCol = st.ResponseCol
	pb.TablePath = st.TablePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tablePb struct {
	RequestCol  string `json:"request_col"`
	ResponseCol string `json:"response_col,omitempty"`
	TablePath   string `json:"table_path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableFromPb(pb *tablePb) (*Table, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Table{}
	st.RequestCol = pb.RequestCol
	st.ResponseCol = pb.ResponseCol
	st.TablePath = pb.TablePath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateCustomLlmRequestToPb(st *UpdateCustomLlmRequest) (*updateCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCustomLlmRequestPb{}
	pb.CustomLlm = st.CustomLlm
	pb.Id = st.Id
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

type updateCustomLlmRequestPb struct {
	CustomLlm  CustomLlm `json:"custom_llm"`
	Id         string    `json:"-" url:"-"`
	UpdateMask string    `json:"update_mask"`
}

func updateCustomLlmRequestFromPb(pb *updateCustomLlmRequestPb) (*UpdateCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCustomLlmRequest{}
	st.CustomLlm = pb.CustomLlm
	st.Id = pb.Id
	st.UpdateMask = pb.UpdateMask

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
