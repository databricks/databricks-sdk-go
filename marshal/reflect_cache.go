package marshal

import (
	"reflect"
	"strings"
)

var typeCache = map[reflect.Type][]cachedType{}

var nameCache = map[reflect.Type]string{}

var tagCache = map[string]jsonTag{}

type cachedType struct {
	reflect.StructField
	JsonTag string
}

func getTypeFields(structType reflect.Type) []cachedType {
	if res, ok := typeCache[structType]; ok {
		return res
	}
	res := []cachedType{}
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		res = append(res, cachedType{
			StructField: field,
			JsonTag:     field.Tag.Get("json"),
		})
	}
	typeCache[structType] = res
	return res
}

func getTypeName(structType reflect.Type) string {
	if res, ok := nameCache[structType]; ok {
		return res
	}
	name := structType.Name()
	nameCache[structType] = name
	return name
}

func parseJSONTag(raw string) (jsonTag, error) {
	if tag, ok := tagCache[raw]; ok {
		return tag, nil
	}
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
	tagCache[raw] = jsonTag
	return jsonTag, nil
}
