package code

type Wait struct {
	Method *Method
}

// Bind assumes the same name of field either in request/request or response/response
func (w *Wait) Bind() *Field {
	return w.Poll().Request.Field(w.Method.wait.Bind)
}

func (w *Wait) ForceBindRequest() bool {
	// this was specifically added for Jobs#RepairRun,
	// that does not send run_id in response
	return w.Method.wait.ForceBindRequest
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
	pathToMessage := w.Method.wait.Message
	current := pollMethod.Response
	for {
		fieldName := pathToMessage[0]
		pathToMessage = pathToMessage[1:]
		field := current.Field(fieldName)
		path = append(path, field)
		current = field.Entity
		if len(pathToMessage) == 0 {
			break
		}
	}
	return path
}
