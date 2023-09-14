package marshal

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
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

	return recordOriginalFields(data, v)
}

func recordOriginalFields(data []byte, v any) error {

	value := reflect.ValueOf(v)
	switch value.Kind() {
	case reflect.Ptr:
		// reflex.GetFieldByName may return the field for an anonymous field
		// if the extending struct doesn't have the field
		field := getFieldByName(v, force_send_field_name)
		if !field.IsValid() {
			return nil
		}
		// A Value can be changed only if it is
		// addressable and was not obtained by
		// the use of unexported struct fields.
		if field.CanSet() {
			// change value of N
			if field.Kind() == reflect.Slice {
				presentFields, err := getFieldsInJson(data, v)
				if err != nil {
					return err
				}
				if len(presentFields) == 0 {
					return nil
				}
				presentFieldsValue := reflect.ValueOf(presentFields)
				field.Set(presentFieldsValue)
			}
		} else {
			return errors.New("cannot set field")
		}
	}
	return nil
}

// Returns the list of fields found in the json
func getFieldsInJson(data []byte, v any) ([]string, error) {
	var jsonFields map[string]interface{}
	err := json.Unmarshal([]byte(data), &jsonFields)
	if err != nil {
		return []string{}, err
	}

	schema := reflect.ValueOf(v)
	derefer := reflect.Indirect(schema)

	objectType := derefer.Type()

	elements := []string{}

	for i := 0; i < derefer.NumField(); i++ {
		fieldType := objectType.Field(i).Type

		if !isBasicType(fieldType) {
			continue
		}

		jsonTag := objectType.Field(i).Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		jsonName, err := getJsonName(jsonTag)
		if err != nil {
			return nil, err
		}

		if _, ok := jsonFields[jsonName]; ok {
			name := objectType.Field(i).Name
			elements = append(elements, name)
		}
	}

	return elements, nil
}

// Extract the json name from the json tag
func getJsonName(val string) (string, error) {
	if val == "-" {
		return "-", nil
	}

	i := strings.Index(val, ",")

	if i == -1 {
		return val, nil
	}
	if val[:i] == "" {
		return "", fmt.Errorf("malformed json tag: %s", val)
	}

	return val[:i], nil
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

	for i := 0; i < derefer.NumField(); i++ {
		field := derefer.Field(i)
		jsonTag := objectType.Field(i).Tag.Get("json")

		if objectType.Field(i).Anonymous {
			setField(field, data)
			continue
		}
		if jsonTag == "" {
			continue
		}
		jsonName, err := getJsonName(jsonTag)

		if err != nil {
			return err
		}

		value, ok := jsonFields[jsonName]
		if !ok {
			continue
		}

		if !field.CanAddr() {
			return errors.New("cannot address field")
		}

		err = setField(field, value)
		if err != nil {
			return err
		}
	}

	return nil

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
