package marshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

const FORCE_SEND_FIELD_NAME = "ForceSendFields"

// Marshal returns a JSON encoding of the given object. Included fields:
// - non-empty value
// - a basic type whose field's name is present in forceSendFields
// Our templates always populate either omitempty or the '-' tag.
// For simplicity, we assume one of those cases.
func Marshal(object interface{}) ([]byte, error) {

	dataMap, err := structAsMap(object)
	if err != nil {
		return nil, err
	}

	return json.Marshal(dataMap)
}

// Converts the object to a map
func structAsMap(object interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// If the object is nil or a pointer to nil, don't do anything
	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return result, nil
	}

	forceSendFields, err := getForceSendFields(object)

	if err != nil {
		return nil, err
	}

	includeFields := make(map[string]bool)
	for _, field := range forceSendFields {
		includeFields[field] = true
	}

	value := reflect.ValueOf(object)
	value = reflect.Indirect(value)
	objectType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		jsonTag := objectType.Field(i).Tag.Get("json")
		tag, err := parseJSONTag(jsonTag)
		if err != nil {
			return nil, err
		}

		fieldValue := value.Field(i)
		fieldStruct := objectType.Field(i)

		// Anonymous fields should be marshalled using the same JSON, and then merged into the same map
		if fieldStruct.Anonymous && fieldValue.IsValid() {
			dataMap2, err := structAsMap(fieldValue.Interface())
			if err != nil {
				return nil, err
			}
			result = mergeMaps(result, dataMap2)
			continue
		}

		// Skip fields which should not be included
		if !includeField(tag, fieldValue, fieldStruct, includeFields) {
			continue
		}

		if tag.asString {
			result[tag.name] = formatAsString(fieldValue, fieldStruct.Type.Kind())
		} else {
			result[tag.name] = fieldValue.Interface()
		}
	}
	return result, nil
}

// returns the element as string
func formatAsString(v reflect.Value, kind reflect.Kind) string {
	if kind == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}

	return fmt.Sprintf("%v", v.Interface())
}

type jsonTag struct {
	name      string
	asString  bool
	ignore    bool
	omitempty bool
}

func parseJSONTag(raw string) (jsonTag, error) {
	if raw == "-" || raw == "" {
		return jsonTag{ignore: true}, nil
	}

	parts := strings.Split(raw, ",")

	jsonTag := jsonTag{
		name: parts[0],
	}

	for _, v := range parts {
		switch v {
		case "omitempty":
			jsonTag.omitempty = true
		case "string":
			jsonTag.asString = true
		}
	}
	return jsonTag, nil
}

// Determines wether a field should be indluded or not
func includeField(tag jsonTag, value reflect.Value, field reflect.StructField, mustInclude map[string]bool) bool {
	if tag.ignore {
		return false
	}
	if !tag.omitempty {
		return true
	}
	return (isBasicType(value.Type()) && mustInclude[field.Name]) || !isEmptyValue(value)
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Array, reflect.Map, reflect.Slice:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

// Merges two maps. Overrides duplicated elements, which is fine because json.Unmarshal also
// does it, and a JSON should not have duplicated entries.
func mergeMaps(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for k, v := range m1 {
		merged[k] = v
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}

func getForceSendFields(v interface{}) ([]string, error) {
	// reflect.GetFieldByName panics if the field is inside a null anonymous field
	field := getFieldByName(v, FORCE_SEND_FIELD_NAME)
	if !field.IsValid() {
		return nil, nil
	}
	result, ok := field.Interface().([]string)
	if !ok {
		return nil, fmt.Errorf("invalid type for %s field", FORCE_SEND_FIELD_NAME)
	}
	return result, nil

}

func getFieldByName(v interface{}, fieldName string) reflect.Value {
	value := reflect.ValueOf(v)
	value = reflect.Indirect(value)
	objectType := value.Type()

	for i := 0; i < value.NumField(); i++ {
		name := objectType.Field(i).Name
		if name == fieldName {
			return value.Field(i)
		}
	}
	return reflect.Value{}
}
