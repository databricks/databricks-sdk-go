package commands

import (
	"context"
	"fmt"
	"html"
	"regexp"
	"strings"
)

var (
	// IPython's output prefixes
	outRE = regexp.MustCompile(`Out\[[\d\s]+\]:\s`)
	// HTML tags
	tagRE = regexp.MustCompile(`<[^>]*>`)
	// just exception content without exception name
	exceptionRE = regexp.MustCompile(`.*Exception:\s+(.*)`)
	// execution errors resulting from http errors are sometimes hidden in these keys
	executionErrorRE = regexp.MustCompile(`ExecutionError: ([\s\S]*)\n(StatusCode=[0-9]*)\n(StatusDescription=.*)\n`)
	// usual error message explanation is hidden in this key
	errorMessageRE = regexp.MustCompile(`ErrorMessage=(.+)\n`)
)

// CommandMock mocks the execution of command
type CommandMock func(commandStr string) CommandResults

// CommandExecutor creates a spark context and executes a command and then closes context
type CommandExecutor interface {
	Execute(ctx context.Context, clusterID, language, commandStr string) CommandResults
}

// CommandResults captures results of a command
type CommandResults struct {
	ResultType   string      `json:"resultType,omitempty"`
	Summary      string      `json:"summary,omitempty"`
	Cause        string      `json:"cause,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Schema       interface{} `json:"schema,omitempty"`
	Truncated    bool        `json:"truncated,omitempty"`
	IsJSONSchema bool        `json:"isJsonSchema,omitempty"`
	pos          int
}

// Failed tells if command execution failed
func (cr *CommandResults) Failed() bool {
	return cr.ResultType == "error"
}

// Text returns plain text results
func (cr *CommandResults) Text() string {
	if cr.ResultType != "text" {
		return ""
	}
	return outRE.ReplaceAllLiteralString(cr.Data.(string), "")
}

// Err returns error type
func (cr *CommandResults) Err() error {
	if !cr.Failed() {
		return nil
	}
	return fmt.Errorf(cr.Error())
}

// Error returns error in a bit more friendly way
func (cr *CommandResults) Error() string {
	if cr.ResultType != "error" {
		return ""
	}
	summary := tagRE.ReplaceAllLiteralString(cr.Summary, "")
	summary = html.UnescapeString(summary)

	exceptionMatches := exceptionRE.FindStringSubmatch(summary)
	if len(exceptionMatches) == 2 {
		summary = strings.ReplaceAll(exceptionMatches[1], "; nested exception is:", "")
		summary = strings.TrimRight(summary, " ")
		return summary
	}

	executionErrorMatches := executionErrorRE.FindStringSubmatch(cr.Cause)
	if len(executionErrorMatches) == 4 {
		return strings.Join(executionErrorMatches[1:], "\n")
	}

	errorMessageMatches := errorMessageRE.FindStringSubmatch(cr.Cause)
	if len(errorMessageMatches) == 2 {
		return errorMessageMatches[1]
	}

	return summary
}

// Scan scans for results
func (cr *CommandResults) Scan(dest ...interface{}) bool {
	if cr.ResultType != "table" {
		return false
	}
	if rows, ok := cr.Data.([]interface{}); ok {
		if cr.pos >= len(rows) {
			return false
		}
		if cols, ok := rows[cr.pos].([]interface{}); ok {
			for i := range dest {
				switch d := dest[i].(type) {
				case *string:
					*d = cols[i].(string)
				case *int:
					*d = cols[i].(int)
				case *bool:
					*d = cols[i].(bool)
				}
			}
			cr.pos++
			return true
		}
	}
	return false
}
