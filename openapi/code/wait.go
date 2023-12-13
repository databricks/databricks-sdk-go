package code

import (
	"fmt"
	"sort"
)

// Wait represents a long-running operation, that requires multiple RPC calls
type Wait struct {
	Named
	// represents a method that triggers the start of the long-running operation
	Method *Method
}

// Binding connects fields in generated code across multiple requests
type Binding struct {
	// Polling method request field
	PollField *Field

	// Wrapped method either response or request body field
	Bind *Field

	// Is wrapped method response used?
	IsResponseBind bool
}

// reasonable default timeout for the most of long-running operations
const defaultLongRunningTimeout = 20

// Timeout returns timeout in minutes, defaulting to 20
func (w *Wait) Timeout() int {
	t := w.Method.Operation.Wait.Timeout
	if t == 0 {
		return defaultLongRunningTimeout
	}
	return t
}

// Binding returns a slice of request and response connections
func (w *Wait) Binding() (binding []Binding) {
	poll := w.Poll()
	if w.Method.wait.Binding != nil {
		for pollRequestField, b := range w.Method.wait.Binding {
			var bind *Field
			if b.Response != "" {
				bind = w.Method.Response.Field(b.Response)
			} else {
				bind = w.Method.Request.Field(b.Request)
			}
			binding = append(binding, Binding{
				PollField:      poll.Request.Field(pollRequestField),
				Bind:           bind,
				IsResponseBind: b.Response != "",
			})
		}
		// ensure generated code is deterministic
		// Java SDK relies on bind parameter order.
		sort.Slice(binding, func(a, b int) bool {
			return binding[a].PollField.Name < binding[b].PollField.Name
		})
	} else {
		responseBind := true
		bind := w.Method.wait.Bind
		entity := w.Method.Response
		if entity == nil {
			entity = w.Method.Request
			responseBind = false
		}
		pollField := poll.Request.Field(bind)
		if pollField == nil {
			panic(fmt.Errorf("cannot bind response field: %s", bind))
		}
		binding = append(binding, Binding{
			PollField:      pollField,
			Bind:           entity.Field(bind),
			IsResponseBind: responseBind,
		})
	}
	return binding
}

// ForceBindRequest is a workaround for Jobs#RepairRun,
// that does not send run_id in response
func (w *Wait) ForceBindRequest() bool {
	if w.Method.Response == nil {
		return false
	}
	binding := w.Binding()
	if len(binding) == 1 && !binding[0].IsResponseBind {
		return true
	}
	return false
}

// Poll returns method definition for checking the state of
// the long running operation
func (w *Wait) Poll() *Method {
	getStatus, ok := w.Method.Service.methods[w.Method.wait.Poll]
	if !ok {
		return nil
	}
	return getStatus
}

// Success holds the successful end-states of the operation
func (w *Wait) Success() (match []EnumEntry) {
	enum := w.enum()
	for _, v := range w.Method.wait.Success {
		match = append(match, enum[v])
	}
	return match
}

// Failure holds the failed end-states of the operation
func (w *Wait) Failure() (match []EnumEntry) {
	enum := w.enum()
	for _, v := range w.Method.wait.Failure {
		match = append(match, enum[v])
	}
	return match
}

func (w *Wait) enum() map[string]EnumEntry {
	statusPath := w.StatusPath()
	statusField := statusPath[len(statusPath)-1]
	return statusField.Entity.enum
}

// StatusPath holds the path to the field of polled entity,
// that holds current state of the long-running operation
func (w *Wait) StatusPath() (path []*Field) {
	pollMethod := w.Poll()
	pathToStatus := w.Method.wait.Field
	current := pollMethod.Response
	for {
		fieldName := pathToStatus[0]
		pathToStatus = pathToStatus[1:]
		field := current.Field(fieldName)
		path = append(path, field)
		current = field.Entity
		if len(pathToStatus) == 0 {
			break
		}
	}
	return path
}

// MessagePath holds the path to the field of polled entity,
// that can tell about current inner status of the long-running operation
func (w *Wait) MessagePath() (path []*Field) {
	pollMethod := w.Poll()
	current := pollMethod.Response
	for _, fieldName := range w.Method.wait.Message {
		field := current.Field(fieldName)
		path = append(path, field)
		current = field.Entity
	}
	return path
}

func (w *Wait) Status() *Field {
	path := w.StatusPath()
	if path == nil {
		// unreachable
		return nil
	}
	return path[len(path)-1]
}

func (w *Wait) ComplexMessagePath() bool {
	return len(w.Method.wait.Message) > 1
}

func (w *Wait) MessagePathHead() *Field {
	path := w.MessagePath()
	if len(path) == 0 {
		panic("message path is empty for " + w.Method.Operation.OperationId)
	}
	return path[0]
}
