// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

// all definitions in this file are in alphabetical order

type CancelCommand struct {
	ClusterId string `json:"clusterId,omitempty"`

	CommandId string `json:"commandId,omitempty"`

	ContextId string `json:"contextId,omitempty"`
}

type Command struct {
	// Running cluster id
	ClusterId string `json:"clusterId,omitempty"`
	// Executable code
	Command string `json:"command,omitempty"`
	// Running context id
	ContextId string `json:"contextId,omitempty"`

	Language Language `json:"language,omitempty"`
}

type CommandStatus string

const CommandStatusCancelled CommandStatus = `Cancelled`

const CommandStatusCancelling CommandStatus = `Cancelling`

const CommandStatusError CommandStatus = `Error`

const CommandStatusFinished CommandStatus = `Finished`

const CommandStatusQueued CommandStatus = `Queued`

const CommandStatusRunning CommandStatus = `Running`

// Get command info
type CommandStatusRequest struct {
	ClusterId string `json:"-" url:"clusterId"`

	CommandId string `json:"-" url:"commandId"`

	ContextId string `json:"-" url:"contextId"`
}

type CommandStatusResponse struct {
	Id string `json:"id,omitempty"`

	Results *Results `json:"results,omitempty"`

	Status CommandStatus `json:"status,omitempty"`
}

type ContextStatus string

const ContextStatusError ContextStatus = `Error`

const ContextStatusPending ContextStatus = `Pending`

const ContextStatusRunning ContextStatus = `Running`

// Get status
type ContextStatusRequest struct {
	ClusterId string `json:"-" url:"clusterId"`

	ContextId string `json:"-" url:"contextId"`
}

type ContextStatusResponse struct {
	Id string `json:"id,omitempty"`

	Status ContextStatus `json:"status,omitempty"`
}

type CreateContext struct {
	// Running cluster id
	ClusterId string `json:"clusterId,omitempty"`

	Language Language `json:"language,omitempty"`
}

type Created struct {
	Id string `json:"id,omitempty"`
}

type DestroyContext struct {
	ClusterId string `json:"clusterId"`

	ContextId string `json:"contextId"`
}

type Language string

const LanguagePython Language = `python`

const LanguageScala Language = `scala`

const LanguageSql Language = `sql`

type ResultType string

const ResultTypeError ResultType = `error`

const ResultTypeImage ResultType = `image`

const ResultTypeImages ResultType = `images`

const ResultTypeTable ResultType = `table`

const ResultTypeText ResultType = `text`

type Results struct {
	// The cause of the error
	Cause string `json:"cause,omitempty"`

	Data any/* MISSING TYPE */ `json:"data,omitempty"`
	// The image filename
	FileName string `json:"fileName,omitempty"`

	FileNames []string `json:"fileNames,omitempty"`
	// true if a JSON schema is returned instead of a string representation of
	// the Hive type.
	IsJsonSchema bool `json:"isJsonSchema,omitempty"`
	// internal field used by SDK
	Pos int `json:"pos,omitempty"`

	ResultType ResultType `json:"resultType,omitempty"`
	// The table schema
	Schema [][]any/* MISSING TYPE */ `json:"schema,omitempty"`
	// The summary of the error
	Summary string `json:"summary,omitempty"`
	// true if partial results are returned.
	Truncated bool `json:"truncated,omitempty"`
}
