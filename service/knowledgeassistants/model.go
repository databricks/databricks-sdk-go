// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledgeassistants

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateKnowledgeAssistantRequest struct {
	// The Knowledge Assistant to create.
	KnowledgeAssistant KnowledgeAssistant `json:"knowledge_assistant"`
}

type CreateKnowledgeSourceRequest struct {
	KnowledgeSource KnowledgeSource `json:"knowledge_source"`
	// Parent resource where this source will be created. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent string `json:"-" url:"-"`
}

type DeleteKnowledgeAssistantRequest struct {
	// The resource name of the knowledge assistant to be deleted. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
}

type DeleteKnowledgeSourceRequest struct {
	// The resource name of the Knowledge Source to delete. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"-" url:"-"`
}

// FileTableSpec specifies a file table source configuration.
type FileTableSpec struct {
	// The name of the column containing BINARY file content to be indexed.
	FileCol string `json:"file_col"`
	// Full UC name of the table, in the format of
	// {CATALOG}.{SCHEMA}.{TABLE_NAME}.
	TableName string `json:"table_name"`
}

// FilesSpec specifies a files source configuration.
type FilesSpec struct {
	// A UC volume path that includes a list of files.
	Path string `json:"path"`
}

type GetKnowledgeAssistantRequest struct {
	// The resource name of the knowledge assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
}

type GetKnowledgeSourceRequest struct {
	// The resource name of the Knowledge Source. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"-" url:"-"`
}

// IndexSpec specifies a vector search index source configuration.
type IndexSpec struct {
	// The column that specifies a link or reference to where the information
	// came from.
	DocUriCol string `json:"doc_uri_col"`
	// Full UC name of the vector search index, in the format of
	// {CATALOG}.{SCHEMA}.{INDEX_NAME}.
	IndexName string `json:"index_name"`
	// The column that includes the document text for retrieval.
	TextCol string `json:"text_col"`
}

// Entity message that represents a knowledge assistant. Note: REQUIRED
// annotations below represent create-time requirements. For updates, required
// fields are determined by the update mask.
type KnowledgeAssistant struct {
	// Creation timestamp.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The creator of the Knowledge Assistant.
	Creator string `json:"creator,omitempty"`
	// Description of what this agent can do (user-facing). Required when
	// creating a Knowledge Assistant. When updating a Knowledge Assistant,
	// optional unless included in update_mask.
	Description string `json:"description"`
	// The display name of the Knowledge Assistant, unique at workspace level.
	// Required when creating a Knowledge Assistant. When updating a Knowledge
	// Assistant, optional unless included in update_mask.
	DisplayName string `json:"display_name"`
	// The name of the knowledge assistant agent endpoint.
	EndpointName string `json:"endpoint_name,omitempty"`
	// Error details when the Knowledge Assistant is in FAILED state.
	ErrorInfo string `json:"error_info,omitempty"`
	// The MLflow experiment ID.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The universally unique identifier (UUID) of the Knowledge Assistant.
	Id string `json:"id,omitempty"`
	// Additional global instructions on how the agent should generate answers.
	// Optional on create and update. When updating a Knowledge Assistant,
	// include this field in update_mask to modify it.
	Instructions string `json:"instructions,omitempty"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"name,omitempty"`
	// State of the Knowledge Assistant. Not returned in List responses.
	State KnowledgeAssistantState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistant) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistant) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeAssistantState string

const KnowledgeAssistantStateActive KnowledgeAssistantState = `ACTIVE`

const KnowledgeAssistantStateCreating KnowledgeAssistantState = `CREATING`

const KnowledgeAssistantStateFailed KnowledgeAssistantState = `FAILED`

// String representation for [fmt.Print]
func (f *KnowledgeAssistantState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KnowledgeAssistantState) Set(v string) error {
	switch v {
	case `ACTIVE`, `CREATING`, `FAILED`:
		*f = KnowledgeAssistantState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "CREATING", "FAILED"`, v)
	}
}

// Values returns all possible values for KnowledgeAssistantState.
//
// There is no guarantee on the order of the values in the slice.
func (f *KnowledgeAssistantState) Values() []KnowledgeAssistantState {
	return []KnowledgeAssistantState{
		KnowledgeAssistantStateActive,
		KnowledgeAssistantStateCreating,
		KnowledgeAssistantStateFailed,
	}
}

// Type always returns KnowledgeAssistantState to satisfy [pflag.Value] interface
func (f *KnowledgeAssistantState) Type() string {
	return "KnowledgeAssistantState"
}

// KnowledgeSource represents a source of knowledge for the KnowledgeAssistant.
// Used in create/update requests and returned in Get/List responses. Note:
// REQUIRED annotations below represent create-time requirements. For updates,
// required fields are determined by the update mask.
type KnowledgeSource struct {
	// Timestamp when this knowledge source was created.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// Description of the knowledge source. Required when creating a Knowledge
	// Source. When updating a Knowledge Source, optional unless included in
	// update_mask.
	Description string `json:"description"`
	// Human-readable display name of the knowledge source. Required when
	// creating a Knowledge Source. When updating a Knowledge Source, optional
	// unless included in update_mask.
	DisplayName string `json:"display_name"`

	FileTable *FileTableSpec `json:"file_table,omitempty"`

	Files *FilesSpec `json:"files,omitempty"`

	Id string `json:"id,omitempty"`

	Index *IndexSpec `json:"index,omitempty"`
	// Timestamp representing the cutoff before which content in this knowledge
	// source is being ingested.
	KnowledgeCutoffTime *time.Time `json:"knowledge_cutoff_time,omitempty"`
	// Full resource name:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"name,omitempty"`
	// The type of the source: "index", "files", or "file_table". Required when
	// creating a Knowledge Source. When updating a Knowledge Source, this field
	// is ignored.
	SourceType string `json:"source_type"`

	State KnowledgeSourceState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeSource) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeSource) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type KnowledgeSourceState string

const KnowledgeSourceStateFailedUpdate KnowledgeSourceState = `FAILED_UPDATE`

const KnowledgeSourceStateUpdated KnowledgeSourceState = `UPDATED`

const KnowledgeSourceStateUpdating KnowledgeSourceState = `UPDATING`

// String representation for [fmt.Print]
func (f *KnowledgeSourceState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *KnowledgeSourceState) Set(v string) error {
	switch v {
	case `FAILED_UPDATE`, `UPDATED`, `UPDATING`:
		*f = KnowledgeSourceState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_UPDATE", "UPDATED", "UPDATING"`, v)
	}
}

// Values returns all possible values for KnowledgeSourceState.
//
// There is no guarantee on the order of the values in the slice.
func (f *KnowledgeSourceState) Values() []KnowledgeSourceState {
	return []KnowledgeSourceState{
		KnowledgeSourceStateFailedUpdate,
		KnowledgeSourceStateUpdated,
		KnowledgeSourceStateUpdating,
	}
}

// Type always returns KnowledgeSourceState to satisfy [pflag.Value] interface
func (f *KnowledgeSourceState) Type() string {
	return "KnowledgeSourceState"
}

type ListKnowledgeAssistantsRequest struct {
	// The maximum number of knowledge assistants to return. If unspecified, at
	// most 100 knowledge assistants will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListKnowledgeAssistants` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeAssistantsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeAssistantsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Knowledge Assistants.
type ListKnowledgeAssistantsResponse struct {
	KnowledgeAssistants []KnowledgeAssistant `json:"knowledge_assistants,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeAssistantsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeAssistantsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListKnowledgeSourcesRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent resource to list from. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeSourcesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeSourcesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListKnowledgeSourcesResponse struct {
	KnowledgeSources []KnowledgeSource `json:"knowledge_sources,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListKnowledgeSourcesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListKnowledgeSourcesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SyncKnowledgeSourcesRequest struct {
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
}

type UpdateKnowledgeAssistantRequest struct {
	// The Knowledge Assistant update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Assistant fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeAssistant KnowledgeAssistant `json:"knowledge_assistant"`
	// The resource name of the Knowledge Assistant. Format:
	// knowledge-assistants/{knowledge_assistant_id}
	Name string `json:"-" url:"-"`
	// Comma-delimited list of fields to update on the Knowledge Assistant.
	// Allowed values: `display_name`, `description`, `instructions`. Examples:
	// - `display_name` - `description,instructions`
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type UpdateKnowledgeSourceRequest struct {
	// The Knowledge Source update payload. Only fields listed in update_mask
	// are updated. REQUIRED annotations on Knowledge Source fields describe
	// create-time requirements and do not mean all those fields are required
	// for update.
	KnowledgeSource KnowledgeSource `json:"knowledge_source"`
	// The resource name of the Knowledge Source to update. Format:
	// knowledge-assistants/{knowledge_assistant_id}/knowledge-sources/{knowledge_source_id}
	Name string `json:"-" url:"-"`
	// Comma-delimited list of fields to update on the Knowledge Source. Allowed
	// values: `display_name`, `description`. Examples: - `display_name` -
	// `display_name,description`
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}
