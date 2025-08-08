// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package agentbricks

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/agentbricks/agentbrickspb"
)

type CancelCustomLlmOptimizationRunRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st CancelCustomLlmOptimizationRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := CancelCustomLlmOptimizationRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CancelCustomLlmOptimizationRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.CancelCustomLlmOptimizationRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CancelCustomLlmOptimizationRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CancelCustomLlmOptimizationRunRequestToPb(st *CancelCustomLlmOptimizationRunRequest) (*agentbrickspb.CancelCustomLlmOptimizationRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.CancelCustomLlmOptimizationRunRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func CancelCustomLlmOptimizationRunRequestFromPb(pb *agentbrickspb.CancelCustomLlmOptimizationRunRequestPb) (*CancelCustomLlmOptimizationRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CancelCustomLlmOptimizationRunRequest{}
	st.Id = pb.Id

	return st, nil
}

type CreateCustomLlmRequest struct {
	// Optional: UC path for agent artifacts. If you are using a dataset that
	// you only have read permissions, please provide a destination path where
	// you have write permissions. Please provide this in catalog.schema format.
	// Wire name: 'agent_artifact_path'
	AgentArtifactPath string ``
	// Datasets used for training and evaluating the model, not for inference.
	// Currently, only 1 dataset is accepted.
	// Wire name: 'datasets'
	Datasets []Dataset ``
	// Guidelines for the custom LLM to adhere to
	// Wire name: 'guidelines'
	Guidelines []string ``
	// Instructions for the custom LLM to follow
	// Wire name: 'instructions'
	Instructions string ``
	// Name of the custom LLM. Only alphanumeric characters and dashes allowed.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.CreateCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCustomLlmRequestToPb(st *CreateCustomLlmRequest) (*agentbrickspb.CreateCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.CreateCustomLlmRequestPb{}
	pb.AgentArtifactPath = st.AgentArtifactPath

	var datasetsPb []agentbrickspb.DatasetPb
	for _, item := range st.Datasets {
		itemPb, err := DatasetToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			datasetsPb = append(datasetsPb, *itemPb)
		}
	}
	pb.Datasets = datasetsPb
	pb.Guidelines = st.Guidelines
	pb.Instructions = st.Instructions
	pb.Name = st.Name

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateCustomLlmRequestFromPb(pb *agentbrickspb.CreateCustomLlmRequestPb) (*CreateCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCustomLlmRequest{}
	st.AgentArtifactPath = pb.AgentArtifactPath

	var datasetsField []Dataset
	for _, itemPb := range pb.Datasets {
		item, err := DatasetFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			datasetsField = append(datasetsField, *item)
		}
	}
	st.Datasets = datasetsField
	st.Guidelines = pb.Guidelines
	st.Instructions = pb.Instructions
	st.Name = pb.Name

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CustomLlm struct {

	// Wire name: 'agent_artifact_path'
	AgentArtifactPath string ``
	// Creation timestamp of the custom LLM
	// Wire name: 'creation_time'
	CreationTime string `` //legacy
	// Creator of the custom LLM
	// Wire name: 'creator'
	Creator string ``
	// Datasets used for training and evaluating the model, not for inference
	// Wire name: 'datasets'
	Datasets []Dataset ``
	// Name of the endpoint that will be used to serve the custom LLM
	// Wire name: 'endpoint_name'
	EndpointName string ``
	// Guidelines for the custom LLM to adhere to
	// Wire name: 'guidelines'
	Guidelines []string ``

	// Wire name: 'id'
	Id string ``
	// Instructions for the custom LLM to follow
	// Wire name: 'instructions'
	Instructions string ``
	// Name of the custom LLM
	// Wire name: 'name'
	Name string ``
	// If optimization is kicked off, tracks the state of the custom LLM
	// Wire name: 'optimization_state'
	OptimizationState State    ``
	ForceSendFields   []string `tf:"-"`
}

func (st CustomLlm) MarshalJSON() ([]byte, error) {
	pb, err := CustomLlmToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CustomLlm) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.CustomLlmPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CustomLlmFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CustomLlmToPb(st *CustomLlm) (*agentbrickspb.CustomLlmPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.CustomLlmPb{}
	pb.AgentArtifactPath = st.AgentArtifactPath
	pb.CreationTime = st.CreationTime
	pb.Creator = st.Creator

	var datasetsPb []agentbrickspb.DatasetPb
	for _, item := range st.Datasets {
		itemPb, err := DatasetToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			datasetsPb = append(datasetsPb, *itemPb)
		}
	}
	pb.Datasets = datasetsPb
	pb.EndpointName = st.EndpointName
	pb.Guidelines = st.Guidelines
	pb.Id = st.Id
	pb.Instructions = st.Instructions
	pb.Name = st.Name
	optimizationStatePb, err := StateToPb(&st.OptimizationState)
	if err != nil {
		return nil, err
	}
	if optimizationStatePb != nil {
		pb.OptimizationState = *optimizationStatePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CustomLlmFromPb(pb *agentbrickspb.CustomLlmPb) (*CustomLlm, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomLlm{}
	st.AgentArtifactPath = pb.AgentArtifactPath
	st.CreationTime = pb.CreationTime
	st.Creator = pb.Creator

	var datasetsField []Dataset
	for _, itemPb := range pb.Datasets {
		item, err := DatasetFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			datasetsField = append(datasetsField, *item)
		}
	}
	st.Datasets = datasetsField
	st.EndpointName = pb.EndpointName
	st.Guidelines = pb.Guidelines
	st.Id = pb.Id
	st.Instructions = pb.Instructions
	st.Name = pb.Name
	optimizationStateField, err := StateFromPb(&pb.OptimizationState)
	if err != nil {
		return nil, err
	}
	if optimizationStateField != nil {
		st.OptimizationState = *optimizationStateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Dataset struct {

	// Wire name: 'table'
	Table Table ``
}

func (st Dataset) MarshalJSON() ([]byte, error) {
	pb, err := DatasetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Dataset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.DatasetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatasetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DatasetToPb(st *Dataset) (*agentbrickspb.DatasetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.DatasetPb{}
	tablePb, err := TableToPb(&st.Table)
	if err != nil {
		return nil, err
	}
	if tablePb != nil {
		pb.Table = *tablePb
	}

	return pb, nil
}

func DatasetFromPb(pb *agentbrickspb.DatasetPb) (*Dataset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dataset{}
	tableField, err := TableFromPb(&pb.Table)
	if err != nil {
		return nil, err
	}
	if tableField != nil {
		st.Table = *tableField
	}

	return st, nil
}

type DeleteCustomLlmRequest struct {
	// The id of the custom llm
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st DeleteCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.DeleteCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteCustomLlmRequestToPb(st *DeleteCustomLlmRequest) (*agentbrickspb.DeleteCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.DeleteCustomLlmRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteCustomLlmRequestFromPb(pb *agentbrickspb.DeleteCustomLlmRequestPb) (*DeleteCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCustomLlmRequest{}
	st.Id = pb.Id

	return st, nil
}

type GetCustomLlmRequest struct {
	// The id of the custom llm
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st GetCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.GetCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetCustomLlmRequestToPb(st *GetCustomLlmRequest) (*agentbrickspb.GetCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.GetCustomLlmRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func GetCustomLlmRequestFromPb(pb *agentbrickspb.GetCustomLlmRequestPb) (*GetCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCustomLlmRequest{}
	st.Id = pb.Id

	return st, nil
}

type StartCustomLlmOptimizationRunRequest struct {
	// The Id of the tile.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st StartCustomLlmOptimizationRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := StartCustomLlmOptimizationRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *StartCustomLlmOptimizationRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.StartCustomLlmOptimizationRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := StartCustomLlmOptimizationRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func StartCustomLlmOptimizationRunRequestToPb(st *StartCustomLlmOptimizationRunRequest) (*agentbrickspb.StartCustomLlmOptimizationRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.StartCustomLlmOptimizationRunRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func StartCustomLlmOptimizationRunRequestFromPb(pb *agentbrickspb.StartCustomLlmOptimizationRunRequestPb) (*StartCustomLlmOptimizationRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartCustomLlmOptimizationRunRequest{}
	st.Id = pb.Id

	return st, nil
}

// States of Custom LLM optimization lifecycle.
type State string

const StateCancelled State = `CANCELLED`

const StateCompleted State = `COMPLETED`

const StateCreated State = `CREATED`

const StateFailed State = `FAILED`

const StatePending State = `PENDING`

const StateRunning State = `RUNNING`

// String representation for [fmt.Print]
func (f *State) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *State) Set(v string) error {
	switch v {
	case `CANCELLED`, `COMPLETED`, `CREATED`, `FAILED`, `PENDING`, `RUNNING`:
		*f = State(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "COMPLETED", "CREATED", "FAILED", "PENDING", "RUNNING"`, v)
	}
}

// Values returns all possible values for State.
//
// There is no guarantee on the order of the values in the slice.
func (f *State) Values() []State {
	return []State{
		StateCancelled,
		StateCompleted,
		StateCreated,
		StateFailed,
		StatePending,
		StateRunning,
	}
}

// Type always returns State to satisfy [pflag.Value] interface
func (f *State) Type() string {
	return "State"
}

func StateToPb(st *State) (*agentbrickspb.StatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := agentbrickspb.StatePb(*st)
	return &pb, nil
}

func StateFromPb(pb *agentbrickspb.StatePb) (*State, error) {
	if pb == nil {
		return nil, nil
	}
	st := State(*pb)
	return &st, nil
}

type Table struct {
	// Name of the request column
	// Wire name: 'request_col'
	RequestCol string ``
	// Optional: Name of the response column if the data is labeled
	// Wire name: 'response_col'
	ResponseCol string ``
	// Full UC table path in catalog.schema.table_name format
	// Wire name: 'table_path'
	TablePath       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Table) MarshalJSON() ([]byte, error) {
	pb, err := TableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Table) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.TablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TableToPb(st *Table) (*agentbrickspb.TablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.TablePb{}
	pb.RequestCol = st.RequestCol
	pb.ResponseCol = st.ResponseCol
	pb.TablePath = st.TablePath

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TableFromPb(pb *agentbrickspb.TablePb) (*Table, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Table{}
	st.RequestCol = pb.RequestCol
	st.ResponseCol = pb.ResponseCol
	st.TablePath = pb.TablePath

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateCustomLlmRequest struct {
	// The CustomLlm containing the fields which should be updated.
	// Wire name: 'custom_llm'
	CustomLlm CustomLlm ``
	// The id of the custom llm
	// Wire name: 'id'
	Id string `tf:"-"`
	// The list of the CustomLlm fields to update. These should correspond to
	// the values (or lack thereof) present in `custom_llm`.
	//
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	// Wire name: 'update_mask'
	UpdateMask string `` //legacy

}

func (st UpdateCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &agentbrickspb.UpdateCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCustomLlmRequestToPb(st *UpdateCustomLlmRequest) (*agentbrickspb.UpdateCustomLlmRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &agentbrickspb.UpdateCustomLlmRequestPb{}
	customLlmPb, err := CustomLlmToPb(&st.CustomLlm)
	if err != nil {
		return nil, err
	}
	if customLlmPb != nil {
		pb.CustomLlm = *customLlmPb
	}
	pb.Id = st.Id
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

func UpdateCustomLlmRequestFromPb(pb *agentbrickspb.UpdateCustomLlmRequestPb) (*UpdateCustomLlmRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCustomLlmRequest{}
	customLlmField, err := CustomLlmFromPb(&pb.CustomLlm)
	if err != nil {
		return nil, err
	}
	if customLlmField != nil {
		st.CustomLlm = *customLlmField
	}
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
