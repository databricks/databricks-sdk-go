// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package commands

import "fmt"

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

// String representation for [fmt.Print]
func (cs *CommandStatus) String() string {
	return string(*cs)
}

// Set raw string value and validate it against allowed values
func (cs *CommandStatus) Set(v string) error {
	switch v {
	case `Cancelled`, `Cancelling`, `Error`, `Finished`, `Queued`, `Running`:
		*cs = CommandStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Cancelled", "Cancelling", "Error", "Finished", "Queued", "Running"`, v)
	}
}

// Type always returns CommandStatus to satisfy [pflag.Value] interface
func (cs *CommandStatus) Type() string {
	return "CommandStatus"
}

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

// String representation for [fmt.Print]
func (cs *ContextStatus) String() string {
	return string(*cs)
}

// Set raw string value and validate it against allowed values
func (cs *ContextStatus) Set(v string) error {
	switch v {
	case `Error`, `Pending`, `Running`:
		*cs = ContextStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Error", "Pending", "Running"`, v)
	}
}

// Type always returns ContextStatus to satisfy [pflag.Value] interface
func (cs *ContextStatus) Type() string {
	return "ContextStatus"
}

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

// String representation for [fmt.Print]
func (l *Language) String() string {
	return string(*l)
}

// Set raw string value and validate it against allowed values
func (l *Language) Set(v string) error {
	switch v {
	case `python`, `scala`, `sql`:
		*l = Language(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "python", "scala", "sql"`, v)
	}
}

// Type always returns Language to satisfy [pflag.Value] interface
func (l *Language) Type() string {
	return "Language"
}

type ResultType string

const ResultTypeError ResultType = `error`

const ResultTypeImage ResultType = `image`

const ResultTypeImages ResultType = `images`

const ResultTypeTable ResultType = `table`

const ResultTypeText ResultType = `text`

// String representation for [fmt.Print]
func (rt *ResultType) String() string {
	return string(*rt)
}

// Set raw string value and validate it against allowed values
func (rt *ResultType) Set(v string) error {
	switch v {
	case `error`, `image`, `images`, `table`, `text`:
		*rt = ResultType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "error", "image", "images", "table", "text"`, v)
	}
}

// Type always returns ResultType to satisfy [pflag.Value] interface
func (rt *ResultType) Type() string {
	return "ResultType"
}

type Results struct {
	// The cause of the error
	Cause string `json:"cause,omitempty"`

	Data any `json:"data,omitempty"`
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
	Schema []map[string]any `json:"schema,omitempty"`
	// The summary of the error
	Summary string `json:"summary,omitempty"`
	// true if partial results are returned.
	Truncated bool `json:"truncated,omitempty"`
}
