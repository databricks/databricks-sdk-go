package marshal

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/exp/maps"
)

const force_send_field_name = "ForceSendFields"

// Marshal returns a JSON encoding of the given object. Included fields:
// - non-empty value
// - a basic type whose field's name is present in forceSendFields
// Our templates always populate either omitempty or the '-' tag.
// For simplicity, we assume one of those cases.
func Marshal(object any) ([]byte, error) {

	dataMap, err := structAsMap(object)
	if err != nil {
		return nil, err
	}

	return json.Marshal(dataMap)
}

// Converts the object to a map
func structAsMap(object any) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	// If the object is nil or a pointer to nil, don't do anything
	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return result, nil
	}

	includeFields, err := getForceSendFields(object)

	if err != nil {
		return nil, err
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
			anonymousFieldResult, err := structAsMap(fieldValue.Interface())
			if err != nil {
				return nil, err
			}
			result = mergeMaps(anonymousFieldResult, result)
			continue
		}

		// Skip fields which should not be included
		if !includeField(tag, fieldValue, includeFields[fieldStruct.Name]) {
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

// Determines whether a field should be included or not
func includeField(tag jsonTag, value reflect.Value, forceSend bool) bool {
	if tag.ignore {
		return false
	}
	if !tag.omitempty {
		return true
	}
	return (isBasicType(value.Type()) && forceSend) || !isEmptyValue(value)
}

// isEmptyValue returns whether v is the empty value for its type.
// This implementation is based on on the encoding/json package for consistency on the results
// https://github.com/golang/go/blob/a278550c40ef3f01a5fcbef43414dc49009201f8/src/encoding/json/encode.go#L306
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}
	return false
}

// Merges two maps. Overrides duplicated elements, which is fine because json.Unmarshal also
// does it, and a JSON should not have duplicated entries.
func mergeMaps(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	maps.Copy(merged, m2)
	maps.Copy(merged, m1)
	return merged
}

func getForceSendFields(v any) (map[string]bool, error) {
	// reflect.GetFieldByName panics if the field is inside a null anonymous field
	field := getFieldByName(v, force_send_field_name)
	if !field.IsValid() {
		return nil, nil
	}
	forceSendFields, ok := field.Interface().([]string)
	if !ok {
		return nil, fmt.Errorf("invalid type for %s field", force_send_field_name)
	}
	includeFields := make(map[string]bool)
	for _, field := range forceSendFields {
		includeFields[field] = true
	}

	return includeFields, nil

}

func getFieldByName(v any, fieldName string) reflect.Value {
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
