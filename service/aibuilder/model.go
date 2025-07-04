// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package aibuilder

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type CancelCustomLlmOptimizationRunRequest struct {
	Id string `json:"-" tf:"-"`
}

func (st CancelCustomLlmOptimizationRunRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := cancelCustomLlmOptimizationRunRequestToPb(&st)
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

func (st *CancelCustomLlmOptimizationRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelCustomLlmOptimizationRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelCustomLlmOptimizationRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelCustomLlmOptimizationRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := cancelCustomLlmOptimizationRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CancelOptimizeResponse struct {
}

func (st CancelOptimizeResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := cancelOptimizeResponseToPb(&st)
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

func (st *CancelOptimizeResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cancelOptimizeResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cancelOptimizeResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CancelOptimizeResponse) MarshalJSON() ([]byte, error) {
	pb, err := cancelOptimizeResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateCustomLlmRequest struct {
	// Optional: UC path for agent artifacts. If you are using a dataset that
	// you only have read permissions, please provide a destination path where
	// you have write permissions. Please provide this in catalog.schema format.
	// Wire name: 'agent_artifact_path'
	AgentArtifactPath string `json:"agent_artifact_path,omitempty"`
	// Datasets used for training and evaluating the model, not for inference.
	// Currently, only 1 dataset is accepted.
	// Wire name: 'datasets'
	Datasets []Dataset `json:"datasets,omitempty"`
	// Guidelines for the custom LLM to adhere to
	// Wire name: 'guidelines'
	Guidelines []string `json:"guidelines,omitempty"`
	// Instructions for the custom LLM to follow
	// Wire name: 'instructions'
	Instructions string `json:"instructions"`
	// Name of the custom LLM. Only alphanumeric characters and dashes allowed.
	// Wire name: 'name'
	Name string `json:"name"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateCustomLlmRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createCustomLlmRequestToPb(&st)
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

func (st *CreateCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := createCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CustomLlm struct {

	// Wire name: 'agent_artifact_path'
	AgentArtifactPath string `json:"agent_artifact_path,omitempty"`
	// Creation timestamp of the custom LLM
	// Wire name: 'creation_time'
	CreationTime string `json:"creation_time,omitempty"`
	// Creator of the custom LLM
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// Datasets used for training and evaluating the model, not for inference
	// Wire name: 'datasets'
	Datasets []Dataset `json:"datasets,omitempty"`
	// Name of the endpoint that will be used to serve the custom LLM
	// Wire name: 'endpoint_name'
	EndpointName string `json:"endpoint_name,omitempty"`
	// Guidelines for the custom LLM to adhere to
	// Wire name: 'guidelines'
	Guidelines []string `json:"guidelines,omitempty"`

	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// Instructions for the custom LLM to follow
	// Wire name: 'instructions'
	Instructions string `json:"instructions"`
	// Name of the custom LLM
	// Wire name: 'name'
	Name string `json:"name"`
	// If optimization is kicked off, tracks the state of the custom LLM
	// Wire name: 'optimization_state'
	OptimizationState State `json:"optimization_state"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CustomLlm) EncodeValues(key string, v *url.Values) error {
	pb, err := customLlmToPb(&st)
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

func (st *CustomLlm) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customLlmPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customLlmFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomLlm) MarshalJSON() ([]byte, error) {
	pb, err := customLlmToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type Dataset struct {

	// Wire name: 'table'
	Table Table `json:"table"`
}

func (st Dataset) EncodeValues(key string, v *url.Values) error {
	pb, err := datasetToPb(&st)
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

func (st *Dataset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &datasetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := datasetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Dataset) MarshalJSON() ([]byte, error) {
	pb, err := datasetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteCustomLlmRequest struct {
	// The id of the custom llm
	Id string `json:"-" tf:"-"`
}

func (st DeleteCustomLlmRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCustomLlmRequestToPb(&st)
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

func (st *DeleteCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteCustomLlmResponse struct {
}

func (st DeleteCustomLlmResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteCustomLlmResponseToPb(&st)
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

func (st *DeleteCustomLlmResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCustomLlmResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCustomLlmResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCustomLlmResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteCustomLlmResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetCustomLlmRequest struct {
	// The id of the custom llm
	Id string `json:"-" tf:"-"`
}

func (st GetCustomLlmRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getCustomLlmRequestToPb(&st)
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

func (st *GetCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type StartCustomLlmOptimizationRunRequest struct {
	// The Id of the tile.
	Id string `json:"-" tf:"-"`
}

func (st StartCustomLlmOptimizationRunRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := startCustomLlmOptimizationRunRequestToPb(&st)
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

func (st *StartCustomLlmOptimizationRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &startCustomLlmOptimizationRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := startCustomLlmOptimizationRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st StartCustomLlmOptimizationRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := startCustomLlmOptimizationRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

type Table struct {
	// Name of the request column
	// Wire name: 'request_col'
	RequestCol string `json:"request_col"`
	// Optional: Name of the response column if the data is labeled
	// Wire name: 'response_col'
	ResponseCol string `json:"response_col,omitempty"`
	// Full UC table path in catalog.schema.table_name format
	// Wire name: 'table_path'
	TablePath string `json:"table_path"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Table) EncodeValues(key string, v *url.Values) error {
	pb, err := tableToPb(&st)
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

func (st *Table) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &tablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := tableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Table) MarshalJSON() ([]byte, error) {
	pb, err := tableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateCustomLlmRequest struct {
	// The CustomLlm containing the fields which should be updated.
	// Wire name: 'custom_llm'
	CustomLlm CustomLlm `json:"custom_llm"`
	// The id of the custom llm
	Id string `json:"-" tf:"-"`
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
	UpdateMask string `json:"update_mask"`
}

func (st UpdateCustomLlmRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateCustomLlmRequestToPb(&st)
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

func (st *UpdateCustomLlmRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCustomLlmRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCustomLlmRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCustomLlmRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateCustomLlmRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
