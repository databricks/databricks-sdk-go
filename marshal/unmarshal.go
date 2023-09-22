package marshal

import (
	"encoding/json"
	"errors"
	"reflect"
)

// Unmarshals a JSON element and fills in the ForceSendFields field if
// the struct contains it. Only anotates basic types in the ForceSendFields.
func Unmarshal(data []byte, v any) error {
	if len(data) == 0 {
		return nil
	}
	// create a new element if the pointer is nul
	if v == nil {
		objectType := reflect.ValueOf(v).Type()
		v = reflect.New(objectType).Interface()
	}
	err := unmarshalInternal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func setForceSendFields(v any, presentFields []string) error {

	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Ptr {
		return nil
	}

	// reflex.GetFieldByName may return the field from an anonymous field
	// if the extending struct doesn't have the field.
	field := getFieldByName(v, forceSendFieldName)
	if !field.IsValid() {
		return nil
	}

	if !field.CanSet() || field.Kind() != reflect.Slice {
		return errors.New("cannot set field")
	}

	if len(presentFields) == 0 {
		return nil
	}
	presentFieldsValue := reflect.ValueOf(presentFields)
	field.Set(presentFieldsValue)

	return nil
}

func unmarshalInternal(data []byte, v any) error {
	var jsonFields map[string]json.RawMessage
	err := json.Unmarshal([]byte(data), &jsonFields)
	if err != nil {
		return err
	}

	value := reflect.ValueOf(v)
	derefer := reflect.Indirect(value)

	objectType := derefer.Type()

	foundFields := []string{}
	fields := getTypeFields(objectType)

	for i := 0; i < derefer.NumField(); i++ {
		field := derefer.Field(i)
		fieldType := field.Type()
		jsonTag := fields[i].JsonTag

		if fields[i].Anonymous {
			setField(field, data)
			continue
		}
		if jsonTag == "" {
			continue
		}
		tag, err := parseJSONTag(jsonTag)

		if err != nil {
			return err
		}

		value, ok := jsonFields[tag.name]
		if !ok {
			continue
		}

		if isBasicType(fieldType) {
			foundFields = append(foundFields, fields[i].Name)
		}

		if !field.CanAddr() {
			return errors.New("cannot address field")
		}

		err = setField(field, value)
		if err != nil {
			return err
		}
	}

	return setForceSendFields(v, foundFields)

}

func setField(field reflect.Value, value []byte) error {
	pointer := field.Addr().Interface()
	err := json.Unmarshal(value, pointer)
	if err == nil {
		return nil
	}
	// We use godss/yaml to convert YAML into JSON.
	// This library stops converting when it finds a custom Marshaller,
	// Since strings in YAML may not have quotes, they won't be added.
	// So we add them manually
	return json.Unmarshal([]byte(`"`+string(value)+`"`), pointer)
}
