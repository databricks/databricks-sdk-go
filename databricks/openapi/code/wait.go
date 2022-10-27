package code

import (
	"golang.org/x/exp/slices"
)

type Wait struct {
	Method *Method
}

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
	t := w.Method.operation.Wait.Timeout
	if t == 0 {
		return defaultLongRunningTimeout
	}
	return t
}

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
		slices.SortFunc(binding, func(a, b Binding) bool {
			return a.PollField.Name < b.PollField.Name
		})
	} else {
		responseBind := true
		bind := w.Method.wait.Bind
		entity := w.Method.Response
		if entity == nil {
			entity = w.Method.Request
			responseBind = false
		}
		binding = append(binding, Binding{
			PollField:      poll.Request.Field(bind),
			Bind:           entity.Field(bind),
			IsResponseBind: responseBind,
		})
	}
	return binding
}

func (w *Wait) ForceBindRequest() bool {
	if w.Method.Response == nil {
		return false
	}
	binding := w.Binding()
	// this was specifically added for Jobs#RepairRun,
	// that does not send run_id in response
	if len(binding) == 1 && !binding[0].IsResponseBind {
		return true
	}
	return false
}

func (w *Wait) Poll() *Method {
	getStatus, ok := w.Method.Service.methods[w.Method.wait.Poll]
	if !ok {
		return nil
	}
	return getStatus
}

func (w *Wait) Success() (match []EnumEntry) {
	enum := w.enum()
	for _, v := range w.Method.wait.Success {
		match = append(match, enum[v])
	}
	return match
}

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
