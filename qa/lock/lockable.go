package lock

import (
	"bytes"
	"fmt"
	"reflect"
)

type Lockable interface {
	GetLockId() string
}

// GitCredentials are unique to the user and workspace.
type GitCredentials struct {
	WorkspaceHost string
	Username      string
}

func (g GitCredentials) GetLockId() string {
	return generateBlobName(g)
}

// Some operations require locking the entire workspace.
type Workspace struct {
	WorkspaceId string
}

func (w Workspace) GetLockId() string {
	return generateBlobName(w)
}

// LockableImpl is a default implementation of Lockable. If r is a struct, it uses
// reflection to generate a unique lock ID based on the name and fields of the struct.
// If r is a string, it uses the string as the lock ID.
type LockableImpl[R any] struct {
	r R
}

func NewLockable[R any](r R) LockableImpl[R] {
	return LockableImpl[R]{r: r}
}
func (l LockableImpl[R]) GetLockId() string {
	return generateBlobName(l.r)
}

func generateBlobName[R any](r R) string {
	rv := reflect.ValueOf(r)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	switch rv.Kind() {
	case reflect.String:
		return rv.String()
	case reflect.Struct:
		id := bytes.Buffer{}
		t := rv.Type()
		if rv.NumField() == 0 {
			return t.Name()
		}
		id.WriteString(fmt.Sprintf("%s:", t.Name()))
		for i := 0; i < rv.NumField(); i++ {
			fieldName := t.Field(i).Name
			id.WriteString(fmt.Sprintf("%s=", fieldName))
			switch v := rv.Field(i).Interface().(type) {
			case string:
				id.WriteString(v)
			case int, int8, int16, int32, int64:
				id.WriteString(fmt.Sprintf("%d", v))
			default:
				panic(fmt.Errorf("unsupported type %T for field %s, must be string or int", v, fieldName))
			}
			if i < rv.NumField()-1 {
				id.WriteString(";")
			}
		}
		return id.String()
	default:
		panic(fmt.Errorf("unsupported type %T, must be string or struct (or pointer to string or pointer to struct)", r))
	}
}
